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
