package replicator

import (
	"context"
	"encoding/binary"
	"io"
	"net"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"testing"
	"time"

	"github.com/tamzrod/MCS.OSJS/mma2composer"
	"github.com/tamzrod/MCS.OSJS/mma2raw"
	"github.com/tamzrod/MCS.OSJS/simulator"
)

func TestSimulatorToReplicatorE2E(t *testing.T) {
	root := t.TempDir()
	sourcePort := freeTCPPort(t)
	destinationPort := freeTCPPort(t)

	simStore := simulator.Store{Root: root}
	simDoc := simulator.Document{Devices: []simulator.DeviceDefinition{{
		Name:    "rep007-source",
		Enabled: true,
		MMA2: simulator.MMA2Params{
			Port:   sourcePort,
			UnitID: 1,
			FC3:    simulator.Area{Start: 23, Count: 2},
		},
		RandomRuntime: simulator.RandomRuntimeParams{FC3IntervalMS: 1000},
	}}}
	if err := simStore.SaveDocument(simDoc); err != nil {
		t.Fatalf("save simulator document: %v", err)
	}
	if err := simStore.ComposeDocument(simDoc); err != nil {
		t.Fatalf("compose simulator reservation: %v", err)
	}

	repStore := Store{Root: root}
	repCfg := Config{
		Source: SourceConfig{
			Host: "127.0.0.1", Port: sourcePort, UnitID: 1,
			Function: 3, Start: 23, Count: 2, PollIntervalMS: 30,
		},
		Destination: DestinationConfig{
			ListenerPort: destinationPort, UnitID: 2,
			Area: "fc3", Start: 40, Count: 2,
		},
	}
	if err := repStore.Save(repCfg); err != nil {
		t.Fatalf("save replicator config: %v", err)
	}
	if err := repStore.composeDestination(repCfg.Destination); err != nil {
		t.Fatalf("compose replicator destination: %v", err)
	}

	composer := mma2composer.New(root, ProducerReplicator)
	owners, err := composer.LoadOwners()
	if err != nil {
		t.Fatal(err)
	}
	if !hasOwner(owners, sourcePort, 1, "simulator") || !hasOwner(owners, destinationPort, 2, ProducerReplicator) {
		t.Fatalf("distinct ownership missing: %#v", owners.Reservations)
	}

	binaryPath := filepath.Join(root, "mma2")
	build := exec.Command("go", "build", "-o", binaryPath, "../MMA2/cmd/mma2")
	build.Env = append(os.Environ(), "GOCACHE="+filepath.Join(root, "go-cache"))
	if output, err := build.CombinedOutput(); err != nil {
		t.Fatalf("build MMA2: %v\n%s", err, output)
	}

	cmd := exec.Command(binaryPath, composer.EffectiveConfigPath())
	if err := cmd.Start(); err != nil {
		t.Fatalf("start MMA2: %v", err)
	}
	defer func() {
		if cmd.Process != nil {
			_ = cmd.Process.Signal(os.Interrupt)
			_, _ = cmd.Process.Wait()
		}
	}()

	waitTCP(t, sourcePort, 3*time.Second)
	waitTCP(t, destinationPort, 3*time.Second)

	sourceClient := mma2raw.NewClient(net.JoinHostPort("127.0.0.1", portText(sourcePort)), 1, map[mma2raw.Area]mma2raw.Range{
		mma2raw.HoldingRegisters: {Start: 23, Count: 2},
	})
	first := []uint16{1111, 2222}
	if err := sourceClient.Send(mma2raw.HoldingRegisters, mma2raw.Values{Registers: first}); err != nil {
		t.Fatalf("seed simulator source: %v", err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	runtime := NewRuntime(repStore)
	done := make(chan error, 1)
	go func() { done <- runtime.Run(ctx) }()

	waitRegisters(t, destinationPort, 2, 40, first, 2*time.Second)

	second := []uint16{3333, 4444}
	if err := sourceClient.Send(mma2raw.HoldingRegisters, mma2raw.Values{Registers: second}); err != nil {
		cancel()
		<-done
		t.Fatalf("change simulator source: %v", err)
	}
	waitRegisters(t, destinationPort, 2, 40, second, 2*time.Second)

	cancel()
	if err := <-done; err != nil {
		t.Fatalf("runtime stop: %v", err)
	}

	ownersAfter, err := composer.LoadOwners()
	if err != nil {
		t.Fatal(err)
	}
	if !hasOwner(ownersAfter, sourcePort, 1, "simulator") || !hasOwner(ownersAfter, destinationPort, 2, ProducerReplicator) {
		t.Fatalf("ownership changed unexpectedly: %#v", ownersAfter.Reservations)
	}
}

func hasOwner(doc mma2composer.OwnershipDoc, port, unitID uint16, owner string) bool {
	for _, entry := range doc.Reservations {
		if entry.Port == port && entry.UnitID == unitID && entry.Owner == owner {
			return true
		}
	}
	return false
}

func freeTCPPort(t *testing.T) uint16 {
	t.Helper()
	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	defer ln.Close()
	_, port := splitTestAddress(t, ln.Addr().String())
	return port
}

func waitTCP(t *testing.T, port uint16, timeout time.Duration) {
	t.Helper()
	deadline := time.Now().Add(timeout)
	addr := net.JoinHostPort("127.0.0.1", portText(port))
	for time.Now().Before(deadline) {
		conn, err := net.DialTimeout("tcp", addr, 50*time.Millisecond)
		if err == nil {
			_ = conn.Close()
			return
		}
		time.Sleep(20 * time.Millisecond)
	}
	t.Fatalf("listener %s did not become ready", addr)
}

func waitRegisters(t *testing.T, port, unitID, start uint16, want []uint16, timeout time.Duration) {
	t.Helper()
	deadline := time.Now().Add(timeout)
	for time.Now().Before(deadline) {
		got, err := readHoldingRegisters(port, unitID, start, uint16(len(want)))
		if err == nil && reflect.DeepEqual(got, want) {
			return
		}
		time.Sleep(20 * time.Millisecond)
	}
	got, err := readHoldingRegisters(port, unitID, start, uint16(len(want)))
	t.Fatalf("destination registers = %#v err=%v, want %#v", got, err, want)
}

func readHoldingRegisters(port, unitID, start, count uint16) ([]uint16, error) {
	conn, err := net.DialTimeout("tcp", net.JoinHostPort("127.0.0.1", portText(port)), time.Second)
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	_ = conn.SetDeadline(time.Now().Add(time.Second))

	req := make([]byte, 12)
	binary.BigEndian.PutUint16(req[0:2], 1)
	binary.BigEndian.PutUint16(req[4:6], 6)
	req[6] = byte(unitID)
	req[7] = 3
	binary.BigEndian.PutUint16(req[8:10], start)
	binary.BigEndian.PutUint16(req[10:12], count)
	if _, err := conn.Write(req); err != nil {
		return nil, err
	}
	header := make([]byte, 7)
	if _, err := io.ReadFull(conn, header); err != nil {
		return nil, err
	}
	pdu := make([]byte, int(binary.BigEndian.Uint16(header[4:6]))-1)
	if _, err := io.ReadFull(conn, pdu); err != nil {
		return nil, err
	}
	if len(pdu) != 2+int(count)*2 || pdu[0] != 3 || int(pdu[1]) != int(count)*2 {
		return nil, io.ErrUnexpectedEOF
	}
	values := make([]uint16, count)
	for i := range values {
		values[i] = binary.BigEndian.Uint16(pdu[2+i*2 : 4+i*2])
	}
	return values, nil
}

func portText(port uint16) string {
	if port == 0 {
		return "0"
	}
	var buf [5]byte
	i := len(buf)
	for port > 0 {
		i--
		buf[i] = byte('0' + port%10)
		port /= 10
	}
	return string(buf[i:])
}
