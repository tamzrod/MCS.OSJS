package simulator

import (
	"encoding/binary"
	"io"
	"net"
	"strconv"
	"sync"
	"testing"
	"time"

	"github.com/tamzrod/MCS.OSJS/mma2composer"
)

// fakeRawIngestV1 is a minimal MMA2 Raw Ingest v1 server. It records each decoded
// packet and answers with a per-packet response code (default 0x00 = OK).
type fakeRawIngestV1 struct {
	ln      net.Listener
	mu      sync.Mutex
	packets []rawV1Packet
	respond func(rawV1Packet) byte
}

type rawV1Packet struct {
	area    byte
	unitID  uint16
	address uint16
	count   uint16
	payload []byte
}

func newFakeRawIngestV1(t *testing.T, respond func(rawV1Packet) byte) (*fakeRawIngestV1, uint16) {
	t.Helper()
	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	f := &fakeRawIngestV1{ln: ln, respond: respond}
	go f.serve()
	t.Cleanup(func() { _ = ln.Close() })
	_, portText, _ := net.SplitHostPort(ln.Addr().String())
	port, _ := strconv.Atoi(portText)
	return f, uint16(port)
}

func (f *fakeRawIngestV1) serve() {
	for {
		conn, err := f.ln.Accept()
		if err != nil {
			return
		}
		go func(c net.Conn) {
			defer c.Close()
			_ = c.SetDeadline(time.Now().Add(2 * time.Second))
			var hdr [10]byte
			if _, err := io.ReadFull(c, hdr[:]); err != nil {
				return
			}
			count := binary.BigEndian.Uint16(hdr[8:10])
			plen := int(count) * 2
			if hdr[3] == byte(mma2composer.RawIngestCoils) || hdr[3] == byte(mma2composer.RawIngestDiscreteInputs) {
				plen = (int(count) + 7) / 8
			}
			payload := make([]byte, plen)
			if _, err := io.ReadFull(c, payload); err != nil {
				return
			}
			pkt := rawV1Packet{area: hdr[3], unitID: binary.BigEndian.Uint16(hdr[4:6]), address: binary.BigEndian.Uint16(hdr[6:8]), count: count, payload: payload}
			f.mu.Lock()
			f.packets = append(f.packets, pkt)
			f.mu.Unlock()
			code := byte(0x00)
			if f.respond != nil {
				code = f.respond(pkt)
			}
			_, _ = c.Write([]byte{code})
		}(conn)
	}
}

func (f *fakeRawIngestV1) recorded() []rawV1Packet {
	f.mu.Lock()
	defer f.mu.Unlock()
	return append([]rawV1Packet(nil), f.packets...)
}

// PERSIST-R04 self-check: the runtime status carries real save/restore
// observations (not a configured-only placeholder) and stays observational.
func TestRuntimeStatusCarriesRealPersistenceObservations(t *testing.T) {
	reads := []byte{0x00, 0x2A, 0x01, 0x00}
	endpoint, port := newFakeMMA2Endpoint(t, reads, []byte{0x00})
	device := persistenceStartupDevice("r04-status", port, 1, 0, 2)
	memory := memoryFromMMA2Params(device.MMA2)
	root := t.TempDir()
	writeStartupSnapshot(t, root, device, memory, []byte{0x00, 0x2A, 0x01, 0x00})

	// Run the real startup restore, then record it on the applier.
	results, err := persistenceStartupContext(root, []DeviceDefinition{device})
	if err != nil {
		t.Fatal(err)
	}
	applier := newSchedulerApplier(Store{Root: root}, Document{Devices: []DeviceDefinition{device}}, false)
	t.Cleanup(applier.Stop)
	if err := applier.ArmPersistenceSave(Document{Devices: []DeviceDefinition{device}}); err != nil {
		t.Fatal(err)
	}
	applier.recordPersistenceStartup(results)

	status, err := applier.RuntimeStatus(device.Name)
	if err != nil {
		t.Fatal(err)
	}
	if status.Persistence == nil {
		t.Fatal("runtime status must carry persistence observations")
	}
	if !status.Persistence.Configured || !status.Persistence.Healthy || status.Persistence.Sealed {
		t.Fatalf("committed restore must be healthy and unsealed: %+v", status.Persistence)
	}
	if status.Persistence.LastRestore == nil || !status.Persistence.LastRestore.Committed {
		t.Fatalf("last restore observation must appear: %+v", status.Persistence)
	}
	_ = endpoint

	// A driven save reads the authoritative state over Modbus and appears as a
	// last-save observation.
	host := applier.persistenceSave
	ids := host.SubscribedRuleIDs()
	if len(ids) == 0 {
		t.Fatal("expected a subscribed persistence rule")
	}
	applier.PublishPersistenceEvent(ids[0])
	if host.Status().LastSaveAt == "" {
		t.Fatalf("a real save must record a last-save instant: %+v", host.Status())
	}
	status2, err := applier.RuntimeStatus(device.Name)
	if err != nil {
		t.Fatal(err)
	}
	if status2.Persistence.LastSave == nil {
		t.Fatalf("a real save must appear as a last-save observation: %+v", status2.Persistence)
	}
}

// A non-persistence device keeps no persistence observations (absent, not
// fabricated).
func TestRuntimeStatusPersistenceAbsentWithoutConfig(t *testing.T) {
	device := newStatusDevice("r04-none", true, 1, 1000)
	applier := newSchedulerApplier(Store{Root: t.TempDir()}, Document{Devices: []DeviceDefinition{device}}, false)
	t.Cleanup(applier.Stop)
	status := runtimeStatusOf(t, applier, device.Name)
	if status.Persistence == nil || status.Persistence.Configured {
		t.Fatalf("non-persistence device must report not-configured: %+v", status.Persistence)
	}
	if status.Persistence.LastSave != nil || status.Persistence.LastRestore != nil {
		t.Fatalf("non-persistence device must not fabricate observations: %+v", status.Persistence)
	}
}

// fakeMMA2Endpoint emulates the shared MMA2 listener: it classifies each
// connection by its first bytes and serves either Raw Ingest v1 writes or
// Modbus FC1/FC3 reads, so a real save (Modbus read) and a real startup restore
// (Raw Ingest write) can run against one disposable port.
type fakeMMA2Endpoint struct {
	ln        net.Listener
	mu        sync.Mutex
	packets   []rawV1Packet
	registers []byte
	bits      []byte
}

func newFakeMMA2Endpoint(t *testing.T, registers, bits []byte) (*fakeMMA2Endpoint, uint16) {
	t.Helper()
	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	f := &fakeMMA2Endpoint{ln: ln, registers: registers, bits: bits}
	go f.serve()
	t.Cleanup(func() { _ = ln.Close() })
	_, portText, _ := net.SplitHostPort(ln.Addr().String())
	port, _ := strconv.Atoi(portText)
	return f, uint16(port)
}

func (f *fakeMMA2Endpoint) serve() {
	for {
		conn, err := f.ln.Accept()
		if err != nil {
			return
		}
		go func(c net.Conn) {
			defer c.Close()
			_ = c.SetDeadline(time.Now().Add(2 * time.Second))
			var peek [2]byte
			if _, err := io.ReadFull(c, peek[:]); err != nil {
				return
			}
			if peek[0] == 'R' && peek[1] == 'I' {
				f.handleRawIngest(c)
				return
			}
			f.handleModbus(c, peek)
		}(conn)
	}
}

func (f *fakeMMA2Endpoint) handleRawIngest(c net.Conn) {
	rest := make([]byte, 8)
	if _, err := io.ReadFull(c, rest); err != nil {
		return
	}
	hdr := append([]byte{'R', 'I'}, rest...)
	count := binary.BigEndian.Uint16(hdr[8:10])
	plen := int(count) * 2
	if hdr[3] == byte(mma2composer.RawIngestCoils) || hdr[3] == byte(mma2composer.RawIngestDiscreteInputs) {
		plen = (int(count) + 7) / 8
	}
	payload := make([]byte, plen)
	if _, err := io.ReadFull(c, payload); err != nil {
		return
	}
	f.mu.Lock()
	f.packets = append(f.packets, rawV1Packet{area: hdr[3], unitID: binary.BigEndian.Uint16(hdr[4:6]), address: binary.BigEndian.Uint16(hdr[6:8]), count: count, payload: payload})
	f.mu.Unlock()
	_, _ = c.Write([]byte{0x00})
}

func (f *fakeMMA2Endpoint) handleModbus(c net.Conn, peek [2]byte) {
	rest := make([]byte, 10)
	if _, err := io.ReadFull(c, rest); err != nil {
		return
	}
	req := append(peek[:], rest...)
	function := req[7]
	data := f.registers
	if function == 1 {
		data = f.bits
		if data == nil {
			data = []byte{0}
		}
	}
	if data == nil {
		data = []byte{}
	}
	resp := make([]byte, 9+len(data))
	binary.BigEndian.PutUint16(resp[0:2], binary.BigEndian.Uint16(req[0:2]))
	binary.BigEndian.PutUint16(resp[4:6], uint16(3+len(data)))
	resp[6] = req[6]
	resp[7] = function
	resp[8] = byte(len(data))
	copy(resp[9:], data)
	_, _ = c.Write(resp)
}

func (f *fakeMMA2Endpoint) recorded() []rawV1Packet {
	f.mu.Lock()
	defer f.mu.Unlock()
	return append([]rawV1Packet(nil), f.packets...)
}

// persistenceStartupDevice builds a persistence-enabled device for the startup
// tests with one holding-register area.
func persistenceStartupDevice(name string, port uint16, unitID, start, count uint16) DeviceDefinition {
	return DeviceDefinition{
		Name:    name,
		Enabled: true,
		MMA2: MMA2Params{
			Port:         port,
			UnitID:       unitID,
			FC3:          Area{Start: start, Count: count},
			StateSealing: map[string]interface{}{"enabled": true, "area": "coil", "address": 0},
			Persistence:  &mma2composer.Persistence{Enabled: boolPtr(true)},
		},
	}
}

// writeStartupSnapshot writes a raw snapshot + manifest for a device under
// dataRoot, as the save path (PERSIST-R02/R01) would.
func writeStartupSnapshot(t *testing.T, dataRoot string, device DeviceDefinition, memory mma2composer.Memory, payload []byte) mma2composer.PersistenceMemoryKey {
	t.Helper()
	key := mma2composer.PersistenceMemoryKey{Port: device.MMA2.Port, UnitID: device.MMA2.UnitID}
	rules, err := persistenceOwnedRulesForMemory(memory)
	if err != nil {
		t.Fatal(err)
	}
	configs, err := mma2composer.PersistenceSnapshotConfigs(key, rules)
	if err != nil {
		t.Fatal(err)
	}
	adapter, err := mma2composer.NewPersistenceFilesystemAdapter(dataRoot, configs)
	if err != nil {
		t.Fatal(err)
	}
	if err := adapter.WritePersistenceBytes(key, "holding_registers", 0, payload); err != nil {
		t.Fatal(err)
	}
	return key
}

// PERSIST-R03 self-check: the real startup restore path stays sealed through all
// restore writes and unseals only after verification succeeds, against the real
// Raw Ingest v1 endpoint.
func TestPersistenceStartupRestoreSealedThenUnseals(t *testing.T) {
	fixture, port := newFakeRawIngestV1(t, nil)
	device := persistenceStartupDevice("startup-ok", port, 1, 0, 2)
	memory := memoryFromMMA2Params(device.MMA2)
	root := t.TempDir()
	writeStartupSnapshot(t, root, device, memory, []byte{0x00, 0x2A, 0x01, 0x00})

	results, err := persistenceStartupContext(root, []DeviceDefinition{device})
	if err != nil {
		t.Fatal(err)
	}
	if len(results) != 1 {
		t.Fatalf("expected one startup restore result, got %d", len(results))
	}
	r := results[0]
	if !r.Result.Completed || !r.Result.Committed {
		t.Fatalf("startup restore must complete and commit: %+v", r.Result)
	}
	packets := fixture.recorded()
	if len(packets) != 2 {
		t.Fatalf("expected 1 area write + 1 unseal, got %d: %+v", len(packets), packets)
	}
	// The final packet is the single-coil unseal (sealed -> unsealed = 1).
	last := packets[len(packets)-1]
	if last.area != byte(mma2composer.RawIngestCoils) || last.count != 1 || last.address != 0 || last.payload[0] != 0x01 {
		t.Fatalf("final packet must be the unseal: %+v", last)
	}
}

// PERSIST-R03 self-check: a missing/corrupt snapshot leaves the runtime sealed
// with the existing classified reason, and no unseal is ever issued.
func TestPersistenceStartupRestoreFailClosedMissingSnapshot(t *testing.T) {
	fixture, port := newFakeRawIngestV1(t, nil)
	device := persistenceStartupDevice("startup-missing", port, 1, 0, 2)
	root := t.TempDir()
	// No snapshot written: the loader reports missing.

	results, err := persistenceStartupContext(root, []DeviceDefinition{device})
	if err != nil {
		t.Fatal(err)
	}
	if len(results) != 1 {
		t.Fatalf("expected one result, got %d", len(results))
	}
	r := results[0]
	if r.Result.Committed {
		t.Fatalf("missing snapshot must not commit: %+v", r.Result)
	}
	if r.Result.Failure != mma2composer.PersistenceRestoreFailureMissingSnapshot || !r.Result.Sealed {
		t.Fatalf("missing snapshot must stay sealed with the missing reason: %+v", r.Result)
	}
	if len(fixture.recorded()) != 0 {
		t.Fatal("no Raw Ingest write may occur for a missing snapshot")
	}
}

// PERSIST-R03 self-check: a Raw Ingest failure during restore leaves the runtime
// sealed with the classified reason and never unseals.
func TestPersistenceStartupRestoreFailClosedRawIngest(t *testing.T) {
	fixture, port := newFakeRawIngestV1(t, func(pkt rawV1Packet) byte {
		if pkt.area == byte(mma2composer.RawIngestCoils) && pkt.count == 1 {
			return 0x21 // reject the unseal
		}
		return 0x00
	})
	device := persistenceStartupDevice("startup-raw-fail", port, 1, 0, 2)
	memory := memoryFromMMA2Params(device.MMA2)
	root := t.TempDir()
	writeStartupSnapshot(t, root, device, memory, []byte{0x00, 0x2A, 0x01, 0x00})

	results, err := persistenceStartupContext(root, []DeviceDefinition{device})
	if err != nil {
		t.Fatal(err)
	}
	r := results[0]
	if r.Result.Committed {
		t.Fatalf("a rejected unseal must not commit: %+v", r.Result)
	}
	if !r.Result.Completed || r.Result.Failure != mma2composer.PersistenceRestoreFailureUnsealResponse || !r.Result.Sealed {
		t.Fatalf("rejected unseal must be classified and sealed: %+v", r.Result)
	}
	_ = fixture
}

// PERSIST-R03 self-check: no persistence-enabled memory yields no startup work.
func TestPersistenceStartupContextSkippedWithoutPersistence(t *testing.T) {
	device := validDevice()
	results, err := persistenceStartupContext(t.TempDir(), []DeviceDefinition{device})
	if err != nil {
		t.Fatal(err)
	}
	if len(results) != 0 {
		t.Fatalf("a non-persistence device must not produce a startup restore: %+v", results)
	}
	// A blank data root is a no-op, never an invented path.
	if results, err := persistenceStartupContext("", []DeviceDefinition{device}); err != nil || results != nil {
		t.Fatalf("blank data root must be a no-op: %v %v", results, err)
	}
}
