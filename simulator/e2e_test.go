package simulator

import (
	"bytes"
	"encoding/binary"
	"errors"
	"io"
	"net"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"sync"
	"testing"
	"time"
)

type applianceProcess struct {
	cmd  *exec.Cmd
	done chan error
}

func startAppliance(t *testing.T, binaryPath, configPath string) *applianceProcess {
	t.Helper()
	cmd := exec.Command(binaryPath, configPath)
	var output bytes.Buffer
	cmd.Stdout, cmd.Stderr = &output, &output
	if err := cmd.Start(); err != nil {
		t.Fatal(err)
	}
	p := &applianceProcess{cmd: cmd, done: make(chan error, 1)}
	go func() { p.done <- cmd.Wait() }()
	return p
}

func (p *applianceProcess) stop(t *testing.T) {
	t.Helper()
	if p == nil || p.cmd.Process == nil {
		return
	}
	_ = p.cmd.Process.Signal(os.Interrupt)
	select {
	case <-p.done:
	case <-time.After(2 * time.Second):
		_ = p.cmd.Process.Kill()
		<-p.done
	}
}

func readModbus(t *testing.T, port uint16, unit, fc uint8, start, count uint16) []byte {
	t.Helper()
	conn, err := net.DialTimeout("tcp", net.JoinHostPort("127.0.0.1", fmtPort(port)), time.Second)
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()
	req := make([]byte, 12)
	binary.BigEndian.PutUint16(req[0:2], 1)
	binary.BigEndian.PutUint16(req[4:6], 6)
	req[6], req[7] = unit, fc
	binary.BigEndian.PutUint16(req[8:10], start)
	binary.BigEndian.PutUint16(req[10:12], count)
	if _, err := conn.Write(req); err != nil {
		t.Fatal(err)
	}
	header := make([]byte, 7)
	if _, err := io.ReadFull(conn, header); err != nil {
		t.Fatal(err)
	}
	pdu := make([]byte, int(binary.BigEndian.Uint16(header[4:6]))-1)
	if _, err := io.ReadFull(conn, pdu); err != nil {
		t.Fatal(err)
	}
	if len(pdu) < 2 || pdu[0] != fc || pdu[0]&0x80 != 0 {
		t.Fatalf("FC%d response=%x", fc, pdu)
	}
	return pdu[2:]
}

func fmtPort(port uint16) string { return strconv.Itoa(int(port)) }

func TestEndToEndSimulatorMMA2Architecture(t *testing.T) {
	if os.Getenv("MCS_RUN_E2E") != "1" {
		t.Skip("set MCS_RUN_E2E=1 for real MMA2 verification")
	}
	root := t.TempDir()
	binaryPath := filepath.Join(root, "mma2")
	build := exec.Command("go", "build", "-o", binaryPath, "../MMA2/cmd/mma2")
	build.Env = append(os.Environ(), "GOCACHE="+filepath.Join(root, "go-cache"))
	if output, err := build.CombinedOutput(); err != nil {
		t.Fatalf("build MMA2: %v\n%s", err, output)
	}

	store := Store{Root: root}
	foreignPort, firstPort, secondPort := freePort(t), freePort(t), freePort(t)
	foreign := MMA2Memory{UnitID: 77, HoldingRegs: &MMA2Area{Start: 0, Count: 2}, Policy: memoryFromMMA2Params(simMMA2Params()).Policy}
	if err := store.saveEffective(EffectiveMMA2Config{Listeners: []MMA2Listener{{ID: "foreign", Listen: "0.0.0.0:" + fmtPort(foreignPort), Memory: []MMA2Memory{foreign}}}}); err != nil {
		t.Fatal(err)
	}
	if err := store.saveOwners(OwnershipDoc{Reservations: []OwnershipEntry{{Port: foreignPort, UnitID: 77, Owner: "replicator"}}}); err != nil {
		t.Fatal(err)
	}
	device := validDevice()
	device.Name, device.Enabled, device.MMA2.Port, device.MMA2.UnitID = "e2e-device", true, firstPort, 1
	device.MMA2.FC1, device.MMA2.FC2 = Area{Start: 3, Count: 8}, Area{Start: 13, Count: 8}
	device.MMA2.FC3, device.MMA2.FC4 = Area{Start: 23, Count: 2}, Area{Start: 33, Count: 2}
	device.RandomRuntime = RandomRuntimeParams{FC1IntervalMS: 10, FC2IntervalMS: 10, FC3IntervalMS: 10, FC4IntervalMS: 10}
	doc := Document{Devices: []DeviceDefinition{device}}
	if err := store.SaveDocument(doc); err != nil {
		t.Fatal(err)
	}
	if err := store.ComposeDocument(doc); err != nil {
		t.Fatal(err)
	}

	process := startAppliance(t, binaryPath, store.EffectiveConfigPath())
	defer func() { process.stop(t) }()
	if err := WaitMMA2Ready([]uint16{foreignPort, firstPort}, 3*time.Second); err != nil {
		t.Fatal(err)
	}
	router, applier, err := newRuntimeApplyRouter(store, 2*time.Second)
	if err != nil {
		t.Fatal(err)
	}
	defer applier.Stop()
	time.Sleep(80 * time.Millisecond)
	for fc, start := range map[uint8]uint16{1: 3, 2: 13, 3: 23, 4: 33} {
		if got := readModbus(t, firstPort, 1, fc, start, map[uint8]uint16{1: 8, 2: 8, 3: 2, 4: 2}[fc]); len(got) == 0 {
			t.Fatalf("FC%d returned no values", fc)
		}
	}
	before := readModbus(t, firstPort, 1, 3, 23, 2)
	changed := false
	for i := 0; i < 20; i++ {
		time.Sleep(15 * time.Millisecond)
		if !bytes.Equal(before, readModbus(t, firstPort, 1, 3, 23, 2)) {
			changed = true
			break
		}
	}
	if !changed {
		t.Fatal("scheduled FC3 values did not change")
	}

	var mu sync.Mutex
	restarts := 0
	supervisorDone := make(chan error, 1)
	go func() {
		deadline := time.Now().Add(3 * time.Second)
		for time.Now().Before(deadline) {
			if _, found, _ := store.LoadRestartRequest(); found {
				mu.Lock()
				restarts++
				mu.Unlock()
				process.stop(t)
				process = startAppliance(t, binaryPath, store.EffectiveConfigPath())
				supervisorDone <- nil
				return
			}
			time.Sleep(10 * time.Millisecond)
		}
		supervisorDone <- errors.New("restart request not observed")
	}()
	edited := cloneDocument(doc)
	edited.Devices[0].MMA2.Port = secondPort
	if _, err := router.Apply(edited); err != nil {
		t.Fatal(err)
	}
	if err := <-supervisorDone; err != nil {
		t.Fatal(err)
	}
	mu.Lock()
	restartCount := restarts
	mu.Unlock()
	if restartCount != 1 {
		t.Fatalf("restart requests=%d", restartCount)
	}
	if err := WaitMMA2Ready([]uint16{foreignPort, secondPort}, 2*time.Second); err != nil {
		t.Fatal(err)
	}
	readModbus(t, secondPort, 1, 4, 33, 2)

	configBefore, _ := os.ReadFile(store.EffectiveConfigPath())
	rejected := cloneDocument(edited)
	duplicate := rejected.Devices[0]
	duplicate.Name = "collision"
	rejected.Devices = append(rejected.Devices, duplicate)
	if _, err := router.Apply(rejected); err == nil {
		t.Fatal("expected rejected collision")
	}
	configAfter, _ := os.ReadFile(store.EffectiveConfigPath())
	if !bytes.Equal(configBefore, configAfter) {
		t.Fatal("rejected config damaged shared config")
	}
	if _, found, _ := store.LoadRestartRequest(); found {
		t.Fatal("rejected config requested restart")
	}

	applier.Stop()
	process.stop(t)
	process = startAppliance(t, binaryPath, store.EffectiveConfigPath())
	if err := WaitMMA2Ready([]uint16{foreignPort, secondPort}, 3*time.Second); err != nil {
		t.Fatal(err)
	}
	_, restored, err := newRuntimeApplyRouter(store, 2*time.Second)
	if err != nil {
		t.Fatal(err)
	}
	defer restored.Stop()
	time.Sleep(50 * time.Millisecond)
	status, err := restored.RuntimeStatus("e2e-device")
	if err != nil || status.Device != "RUNNING" || status.RawIngest != "OK" {
		t.Fatalf("reboot restore status=%+v err=%v", status, err)
	}
	if _, found, _ := store.LoadRestartRequest(); found {
		t.Fatal("ordinary reboot created restart request")
	}
	readModbus(t, secondPort, 1, 1, 3, 8)
	readModbus(t, foreignPort, 77, 3, 0, 2)
}
