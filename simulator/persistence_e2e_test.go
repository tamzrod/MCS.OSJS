package simulator

import (
	"encoding/binary"
	"fmt"
	"io"
	"net"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"testing"
	"time"

	"github.com/tamzrod/MCS.OSJS/mma2composer"
)

// PERSIST-R05 — Disposable Persistence E2E Harness.
//
// This committed, repository-native disposable harness is what PERSIST-022
// independent JR runs once. It exercises the real runtime lifecycle end to end
// against a real MMA2 built from source, using only disposable loopback ports
// and a t.TempDir() data root:
//
//	known value -> RBE save -> restart sealed -> Modbus rejected while sealed
//	-> restore -> verify -> final unseal -> Modbus reads the restored value
//
// plus one deterministic failed-restore case that remains sealed.
//
// Exact command (from the repository root):
//
//	cd simulator && MCS_RUN_PERSIST_E2E=1 go test -mod=readonly -count=1 \
//	    -run TestPersistenceDisposableEndToEnd -v .
//
// It is gated behind MCS_RUN_PERSIST_E2E so ordinary `go test ./...` does not
// build MMA2; it builds MMA2 itself from ../MMA2/cmd/mma2 into the temp dir and
// cleans up only its own disposable resources (temp dir + processes).
func TestPersistenceDisposableEndToEnd(t *testing.T) {
	if os.Getenv("MCS_RUN_PERSIST_E2E") != "1" {
		t.Skip("set MCS_RUN_PERSIST_E2E=1 to run the real-MMA2 disposable persistence harness")
	}
	root := t.TempDir()
	binaryPath := buildDisposableMMA2(t, root)

	port := freePort(t)
	store := Store{Root: root}
	device := persistenceE2EDevice(port)
	if err := store.SaveDocument(Document{Devices: []DeviceDefinition{device}}); err != nil {
		t.Fatalf("save simulator document: %v", err)
	}
	if err := store.ComposeDocument(Document{Devices: []DeviceDefinition{device}}); err != nil {
		t.Fatalf("compose MMA2 config: %v", err)
	}

	// -------- Phase A: live operation, known value, RBE save --------
	// current always points at the live disposable process; each is stopped
	// exactly once (a second stop would block on an already-consumed channel).
	var current *applianceProcess
	defer func() {
		if current != nil {
			current.stop(t)
		}
	}()
	current = startAppliance(t, binaryPath, store.EffectiveConfigPath())
	if err := WaitMMA2Ready([]uint16{port}, 5*time.Second); err != nil {
		t.Fatalf("MMA2 not ready: %v", err)
	}

	// First boot: no snapshot exists, so the persistence-enabled memory cannot
	// restore and stays sealed (fail-closed).
	applierA := newDisposableRuntime(t, store)
	if status := mustPersistenceStatus(t, applierA, device.Name); !status.Sealed {
		t.Fatalf("first boot with no snapshot must stay sealed: %+v", status)
	}
	if _, code, err := readModbusResult(port, 1, 3, 0, 1); err != nil || code != 0x06 {
		t.Fatalf("sealed memory must reject Modbus read with 0x06: code=0x%02X err=%v", code, err)
	}

	// Write the known value and unseal via the authoritative unsealing path (Raw
	// Ingest bypasses sealing by design), then operate live.
	rawWriter := NewPersistenceRawIngestWriter("127.0.0.1", port, 1)
	key := mma2composer.PersistenceMemoryKey{Port: port, UnitID: 1}
	if code, err := rawWriter.WritePersistenceRawIngest(key, mma2composer.RawIngestHoldingRegisters, 0, 1, []byte{0x12, 0x34}); err != nil || code != 0x00 {
		t.Fatalf("write known value via raw ingest: code=0x%02X err=%v", code, err)
	}
	if code, err := rawWriter.WritePersistenceRawIngest(key, mma2composer.RawIngestCoils, 0, 1, []byte{0x01}); err != nil || code != 0x00 {
		t.Fatalf("unseal via raw ingest: code=0x%02X err=%v", code, err)
	}
	if payload, code, err := readModbusResult(port, 1, 3, 0, 1); err != nil || code != 0x00 || binary.BigEndian.Uint16(payload) != 0x1234 {
		t.Fatalf("live unsealed read must return the known value: code=0x%02X payload=%x err=%v", code, payload, err)
	}

	// Trigger the real persistence RBE save through the runtime's subscription
	// point (the same call the RBE subscriber delivers). This reads the
	// authoritative area over Modbus and persists the changed bytes.
	ids := applierA.persistenceSave.SubscribedRuleIDs()
	if len(ids) < 1 {
		t.Fatalf("expected subscribed persistence rules, got %v", ids)
	}
	for _, id := range ids {
		applierA.PublishPersistenceEvent(id)
	}
	if status := mustPersistenceStatus(t, applierA, device.Name); status.LastSave == nil {
		t.Fatalf("a real save must be observed: %+v", status)
	}
	applierA.Stop()
	current.stop(t)
	current = nil

	// -------- Phase B: restart sealed -> restore -> unseal -> Modbus reads --------
	current = startAppliance(t, binaryPath, store.EffectiveConfigPath())
	if err := WaitMMA2Ready([]uint16{port}, 5*time.Second); err != nil {
		t.Fatalf("MMA2 not ready after restart: %v", err)
	}
	// Before the runtime restores, a fresh MMA2 has the memory sealed.
	if _, code, err := readModbusResult(port, 1, 3, 0, 1); err != nil || code != 0x06 {
		t.Fatalf("restarted memory must be sealed before restore: code=0x%02X err=%v", code, err)
	}

	// Building the runtime runs the real startup restore: load, restore, verify,
	// final unseal.
	applierB := newDisposableRuntime(t, store)
	statusB := mustPersistenceStatus(t, applierB, device.Name)
	if !statusB.Healthy || statusB.Sealed {
		t.Fatalf("restore must commit and unseal: %+v", statusB)
	}
	if statusB.LastRestore == nil || !statusB.LastRestore.Committed {
		t.Fatalf("last restore observation must appear: %+v", statusB)
	}
	if payload, code, err := readModbusResult(port, 1, 3, 0, 1); err != nil || code != 0x00 || binary.BigEndian.Uint16(payload) != 0x1234 {
		t.Fatalf("Modbus must read the restored value after unseal: code=0x%02X payload=%x err=%v", code, payload, err)
	}
	applierB.Stop()
	current.stop(t)
	current = nil

	// -------- Phase C: failed restore stays sealed --------
	corruptRoot := t.TempDir()
	corruptBinary := buildDisposableMMA2(t, corruptRoot)
	corruptStore := Store{Root: corruptRoot}
	if err := corruptStore.SaveDocument(Document{Devices: []DeviceDefinition{device}}); err != nil {
		t.Fatal(err)
	}
	if err := corruptStore.ComposeDocument(Document{Devices: []DeviceDefinition{device}}); err != nil {
		t.Fatal(err)
	}
	copyDisposableSnapshots(t, filepath.Join(root, mma2composer.RelPersistenceDir), filepath.Join(corruptRoot, mma2composer.RelPersistenceDir))
	corruptSnapshotFile(t, filepath.Join(corruptRoot, mma2composer.RelPersistenceDir))

	corruptProcess := startAppliance(t, corruptBinary, corruptStore.EffectiveConfigPath())
	defer func() { corruptProcess.stop(t) }()
	if err := WaitMMA2Ready([]uint16{port}, 5*time.Second); err != nil {
		t.Fatalf("corrupt MMA2 not ready: %v", err)
	}
	applierC := newDisposableRuntime(t, corruptStore)
	statusC := mustPersistenceStatus(t, applierC, device.Name)
	if statusC.Healthy || !statusC.Sealed {
		t.Fatalf("a failed restore must stay sealed and unhealthy: %+v", statusC)
	}
	if statusC.RestoreOutcome == "" || statusC.RestoreOutcome == "none" {
		t.Fatalf("a failed restore must surface a classified reason: %+v", statusC)
	}
	if _, code, err := readModbusResult(port, 1, 3, 0, 1); err != nil || code != 0x06 {
		t.Fatalf("a failed restore must keep Modbus sealed: code=0x%02X err=%v", code, err)
	}
	applierC.Stop()
}

// persistenceE2EDevice builds the disposable device: one persistence-enabled
// holding-register memory with State Sealing, no random runtime.
func persistenceE2EDevice(port uint16) DeviceDefinition {
	device := validDevice()
	device.Name = "persist-e2e"
	device.Enabled = true
	device.MMA2.Port = port
	device.MMA2.UnitID = 1
	device.MMA2.FC1 = Area{Start: 0, Count: 8}
	device.MMA2.FC2 = Area{}
	device.MMA2.FC3 = Area{Start: 0, Count: 2}
	device.MMA2.FC4 = Area{}
	device.MMA2.StateSealing = map[string]interface{}{"enabled": true, "area": "coil", "address": 0, "exception": 6}
	device.MMA2.Persistence = &mma2composer.Persistence{Enabled: boolPtr(true)}
	device.RandomRuntime = RandomRuntimeParams{}
	return device
}

func buildDisposableMMA2(t *testing.T, dir string) string {
	t.Helper()
	binaryPath := filepath.Join(dir, "mma2")
	build := exec.Command("go", "build", "-o", binaryPath, "../MMA2/cmd/mma2")
	build.Env = append(os.Environ(), "GOCACHE="+filepath.Join(dir, "go-cache"))
	if output, err := build.CombinedOutput(); err != nil {
		t.Fatalf("build MMA2: %v\n%s", err, output)
	}
	return binaryPath
}

// newDisposableRuntime builds the real runtime (which arms persistence save and
// runs the real startup restore) and registers cleanup.
func newDisposableRuntime(t *testing.T, store Store) *SchedulerApplier {
	t.Helper()
	_, applier, err := newRuntimeApplyRouter(store, 5*time.Second)
	if err != nil {
		t.Fatalf("build runtime: %v", err)
	}
	t.Cleanup(applier.Stop)
	return applier
}

func mustPersistenceStatus(t *testing.T, applier *SchedulerApplier, name string) mma2composer.PersistenceRuntimeStatus {
	t.Helper()
	status, err := applier.RuntimeStatus(name)
	if err != nil {
		t.Fatalf("runtime status: %v", err)
	}
	if status.Persistence == nil {
		t.Fatalf("runtime status must carry persistence health: %+v", status)
	}
	return *status.Persistence
}

// readModbusResult performs one Modbus read and returns the payload plus the
// exception code (0 when the read succeeded). It returns an error instead of
// failing the test so sealed rejection can be asserted directly.
func readModbusResult(port uint16, unit, fc uint8, start, count uint16) ([]byte, byte, error) {
	conn, err := net.DialTimeout("tcp", net.JoinHostPort("127.0.0.1", strconv.Itoa(int(port))), 2*time.Second)
	if err != nil {
		return nil, 0, err
	}
	defer conn.Close()
	_ = conn.SetDeadline(time.Now().Add(2 * time.Second))
	req := make([]byte, 12)
	binary.BigEndian.PutUint16(req[0:2], 1)
	binary.BigEndian.PutUint16(req[4:6], 6)
	req[6], req[7] = unit, fc
	binary.BigEndian.PutUint16(req[8:10], start)
	binary.BigEndian.PutUint16(req[10:12], count)
	if _, err := conn.Write(req); err != nil {
		return nil, 0, err
	}
	header := make([]byte, 7)
	if _, err := io.ReadFull(conn, header); err != nil {
		return nil, 0, err
	}
	pdu := make([]byte, int(binary.BigEndian.Uint16(header[4:6]))-1)
	if _, err := io.ReadFull(conn, pdu); err != nil {
		return nil, 0, err
	}
	if len(pdu) >= 1 && pdu[0]&0x80 != 0 {
		code := byte(0)
		if len(pdu) >= 2 {
			code = pdu[1]
		}
		return nil, code, nil
	}
	if len(pdu) < 2 || pdu[0] != fc {
		return nil, 0, fmt.Errorf("unexpected FC%d response %x", fc, pdu)
	}
	return pdu[2:], 0, nil
}

// copyDisposableSnapshots copies a snapshot tree between disposable roots so the
// corrupt-case can start from a valid snapshot and then damage it.
func copyDisposableSnapshots(t *testing.T, src, dst string) {
	t.Helper()
	walk := func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return nil
		}
		rel, _ := filepath.Rel(src, path)
		target := filepath.Join(dst, rel)
		if info.IsDir() {
			return os.MkdirAll(target, 0o755)
		}
		data, rerr := os.ReadFile(path)
		if rerr != nil {
			return rerr
		}
		return os.WriteFile(target, data, 0o644)
	}
	if err := filepath.Walk(src, walk); err != nil {
		t.Fatalf("copy snapshots: %v", err)
	}
}

// corruptSnapshotFile damages the first raw snapshot found, so restore validation
// must fail closed.
func corruptSnapshotFile(t *testing.T, persistenceDir string) {
	t.Helper()
	var target string
	if err := filepath.Walk(persistenceDir, func(path string, info os.FileInfo, err error) error {
		if err == nil && !info.IsDir() && filepath.Base(path) == "snapshot.bin" && target == "" {
			target = path
		}
		return nil
	}); err != nil {
		t.Fatal(err)
	}
	if target == "" {
		t.Fatalf("no snapshot.bin found under %s to corrupt", persistenceDir)
	}
	data, err := os.ReadFile(target)
	if err != nil {
		t.Fatal(err)
	}
	if len(data) == 0 {
		t.Fatalf("snapshot %s is empty", target)
	}
	data[0] ^= 0xFF
	if err := os.WriteFile(target, data, 0o644); err != nil {
		t.Fatal(err)
	}
}
