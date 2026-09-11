package replicator

import (
	"encoding/binary"
	"errors"
	"io"
	"net"
	"reflect"
	"testing"
	"time"

	"github.com/tamzrod/MCS.OSJS/mma2composer"
	"github.com/tamzrod/MCS.OSJS/mma2raw"
)

func TestRunOnceCopiesConfiguredRegisters(t *testing.T) {
	sourceAddr, stopSource := startTestModbusServer(t, 3, 12, []uint16{111, 222, 333})
	defer stopSource()
	sourceHost, sourcePort := splitTestAddress(t, sourceAddr)

	destAddr, received, stopDest := startRawCaptureServer(t, 3)
	defer stopDest()
	_, destPort := splitTestAddress(t, destAddr)

	store := Store{Root: t.TempDir()}
	cfg := Config{
		Source: SourceConfig{
			Host: sourceHost, Port: sourcePort, UnitID: 7, Function: 3,
			Start: 12, Count: 3, PollIntervalMS: 250,
		},
		Destination: DestinationConfig{
			ListenerPort: destPort, UnitID: 9, Area: "fc3", Start: 400, Count: 3,
		},
	}
	if err := store.Save(cfg); err != nil {
		t.Fatal(err)
	}

	result, err := store.RunOnce()
	if err != nil {
		t.Fatalf("RunOnce: %v", err)
	}
	wantResult := CycleResult{Function: 3, SourceStart: 12, DestinationArea: "fc3", DestinationStart: 400, Count: 3}
	if !reflect.DeepEqual(result, wantResult) {
		t.Fatalf("result = %#v, want %#v", result, wantResult)
	}

	packet := <-received
	if len(packet) != mma2raw.HeaderLen+6 {
		t.Fatalf("raw packet length = %d", len(packet))
	}
	if string(packet[0:2]) != "RI" || packet[2] != mma2raw.Version || packet[3] != byte(mma2raw.HoldingRegisters) {
		t.Fatalf("unexpected raw header: %v", packet[:4])
	}
	if got := binary.BigEndian.Uint16(packet[4:6]); got != 9 {
		t.Fatalf("unit id = %d", got)
	}
	if got := binary.BigEndian.Uint16(packet[6:8]); got != 400 {
		t.Fatalf("destination start = %d", got)
	}
	if got := binary.BigEndian.Uint16(packet[8:10]); got != 3 {
		t.Fatalf("destination count = %d", got)
	}
	gotValues := []uint16{
		binary.BigEndian.Uint16(packet[10:12]),
		binary.BigEndian.Uint16(packet[12:14]),
		binary.BigEndian.Uint16(packet[14:16]),
	}
	if !reflect.DeepEqual(gotValues, []uint16{111, 222, 333}) {
		t.Fatalf("destination values = %#v", gotValues)
	}

	owners, err := mma2composer.New(store.Root, ProducerReplicator).LoadOwners()
	if err != nil {
		t.Fatal(err)
	}
	if len(owners.Reservations) != 1 || owners.Reservations[0].Owner != ProducerReplicator || owners.Reservations[0].Port != destPort || owners.Reservations[0].UnitID != 9 {
		t.Fatalf("unexpected ownership registry: %#v", owners.Reservations)
	}
}

func TestRunOnceRejectsForeignOwnedDestination(t *testing.T) {
	destLn, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	_, destPort := splitTestAddress(t, destLn.Addr().String())
	_ = destLn.Close()

	store := Store{Root: t.TempDir()}
	cfg := validConfig()
	cfg.Source.Function = 3
	cfg.Source.Count = 1
	cfg.Destination.ListenerPort = destPort
	cfg.Destination.UnitID = 4
	cfg.Destination.Area = "fc3"
	cfg.Destination.Count = 1
	if err := store.Save(cfg); err != nil {
		t.Fatal(err)
	}

	composer := mma2composer.New(store.Root, "simulator")
	effective := mma2composer.AddMemory(mma2composer.EffectiveConfig{}, "sim-foreign", net.JoinHostPort("0.0.0.0", portString(destPort)), mma2composer.Memory{
		UnitID:      4,
		HoldingRegs: &mma2composer.Area{Start: 0, Count: 1},
		Policy: &mma2composer.Policy{Rules: []mma2composer.PolicyRule{{
			ID: "foreign", SourceIP: []string{"0.0.0.0/0"}, AllowFC: []uint8{3},
		}}},
	})
	owners := mma2composer.OwnershipDoc{Reservations: []mma2composer.OwnershipEntry{{Port: destPort, UnitID: 4, Owner: "simulator"}}}
	if err := composer.Commit(effective, owners); err != nil {
		t.Fatal(err)
	}

	_, err = store.RunOnce()
	if !errors.Is(err, mma2composer.ErrReservationOwnedByOther) {
		t.Fatalf("RunOnce error = %v, want ownership collision", err)
	}

	after, err := composer.LoadOwners()
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(after, owners) {
		t.Fatalf("foreign ownership changed: got %#v want %#v", after, owners)
	}
}

func TestValidateCycleMappingRejectsCrossArea(t *testing.T) {
	cfg := validConfig()
	cfg.Source.Function = 3
	cfg.Destination.Area = "fc4"
	if err := validateCycleMapping(cfg); err == nil {
		t.Fatal("expected cross-area mapping error")
	}
}

func startRawCaptureServer(t *testing.T, count uint16) (string, <-chan []byte, func()) {
	t.Helper()
	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	received := make(chan []byte, 1)
	done := make(chan struct{})
	go func() {
		defer close(done)
		conn, err := ln.Accept()
		if err != nil {
			return
		}
		defer conn.Close()
		_ = conn.SetDeadline(time.Now().Add(2 * time.Second))
		packet := make([]byte, mma2raw.HeaderLen+int(count)*2)
		if _, err := io.ReadFull(conn, packet); err != nil {
			return
		}
		received <- packet
		_, _ = conn.Write([]byte{mma2raw.ResponseOK})
	}()
	return ln.Addr().String(), received, func() {
		_ = ln.Close()
		<-done
	}
}

func portString(port uint16) string {
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
