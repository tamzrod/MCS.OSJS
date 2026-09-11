package replicator

import (
	"encoding/binary"
	"io"
	"net"
	"reflect"
	"testing"
	"time"
)

func TestReadConfiguredSourceUsesPersistedConfig(t *testing.T) {
	addr, stop := startTestModbusServer(t, 3, 30, []uint16{11, 22})
	defer stop()
	host, port := splitTestAddress(t, addr)
	store := Store{Root: t.TempDir()}
	cfg := Config{
		Source: SourceConfig{Host: host, Port: port, UnitID: 7, Function: 3, Start: 30, Count: 2, PollIntervalMS: 100},
		Destination: DestinationConfig{ListenerPort: 1502, UnitID: 9, Area: "fc3", Start: 100, Count: 2},
	}
	if err := store.Save(cfg); err != nil { t.Fatalf("Save: %v", err) }
	got, err := store.ReadConfiguredSource()
	if err != nil { t.Fatalf("ReadConfiguredSource: %v", err) }
	want := RegisterValues{Function: 3, Start: 30, Values: []uint16{11, 22}}
	if !reflect.DeepEqual(got, want) { t.Fatalf("got %#v want %#v", got, want) }
}

func TestReadSourceRangeFC1AndFC2(t *testing.T) {
	for _, function := range []uint8{1, 2} {
		wantBits := []bool{true, false, true, true, false, false, true, false, true, true}
		addr, stop := startTestBitModbusServer(t, function, 12, wantBits)
		host, port := splitTestAddress(t, addr)
		got, err := readSourceRange(SourceConfig{Host: host, Port: port, UnitID: 7, Function: function, Start: 12, Count: uint16(len(wantBits))}, time.Second)
		stop()
		if err != nil { t.Fatalf("FC%d readSourceRange: %v", function, err) }
		if !reflect.DeepEqual(got.Bits, wantBits) || len(got.Values) != 0 { t.Fatalf("FC%d got %#v", function, got) }
	}
}

func TestReadSourceRangeFC3(t *testing.T) {
	addr, stop := startTestModbusServer(t, 3, 12, []uint16{100, 200, 300})
	defer stop()
	host, port := splitTestAddress(t, addr)
	got, err := readSourceRange(SourceConfig{Host: host, Port: port, UnitID: 7, Function: 3, Start: 12, Count: 3}, time.Second)
	if err != nil { t.Fatalf("readSourceRange: %v", err) }
	want := RegisterValues{Function: 3, Start: 12, Values: []uint16{100, 200, 300}}
	if !reflect.DeepEqual(got, want) { t.Fatalf("got %#v want %#v", got, want) }
}

func TestReadSourceRangeFC4(t *testing.T) {
	addr, stop := startTestModbusServer(t, 4, 21, []uint16{0x1234, 0xabcd})
	defer stop()
	host, port := splitTestAddress(t, addr)
	got, err := readSourceRange(SourceConfig{Host: host, Port: port, UnitID: 7, Function: 4, Start: 21, Count: 2}, time.Second)
	if err != nil { t.Fatalf("readSourceRange: %v", err) }
	if !reflect.DeepEqual(got.Values, []uint16{0x1234, 0xabcd}) { t.Fatalf("values = %#v", got.Values) }
}

func TestReadSourceRangeRejectsUnsupportedFunction(t *testing.T) {
	_, err := ReadSourceRange(SourceConfig{Host: "127.0.0.1", Port: 502, UnitID: 1, Function: 5, Start: 0, Count: 1})
	if err == nil { t.Fatal("expected unsupported-function error") }
}

func TestReadSourceRangeConnectionFailure(t *testing.T) {
	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil { t.Fatal(err) }
	addr := ln.Addr().String()
	_ = ln.Close()
	host, port := splitTestAddress(t, addr)
	_, err = readSourceRange(SourceConfig{Host: host, Port: port, UnitID: 1, Function: 3, Start: 0, Count: 1}, 100*time.Millisecond)
	if err == nil { t.Fatal("expected connection error") }
}

func startTestBitModbusServer(t *testing.T, function uint8, start uint16, values []bool) (string, func()) {
	t.Helper()
	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil { t.Fatal(err) }
	done := make(chan struct{})
	go func() {
		defer close(done)
		conn, err := ln.Accept()
		if err != nil { return }
		defer conn.Close()
		request := make([]byte, 12)
		if _, err := io.ReadFull(conn, request); err != nil { return }
		if request[7] != function || binary.BigEndian.Uint16(request[8:10]) != start || binary.BigEndian.Uint16(request[10:12]) != uint16(len(values)) { return }
		payloadBytes := (len(values) + 7) / 8
		response := make([]byte, 9+payloadBytes)
		copy(response[0:2], request[0:2])
		binary.BigEndian.PutUint16(response[2:4], 0)
		binary.BigEndian.PutUint16(response[4:6], uint16(3+payloadBytes))
		response[6], response[7], response[8] = request[6], function, byte(payloadBytes)
		for i, on := range values { if on { response[9+i/8] |= 1 << uint(i%8) } }
		_, _ = conn.Write(response)
	}()
	return ln.Addr().String(), func() { _ = ln.Close(); <-done }
}

func startTestModbusServer(t *testing.T, function uint8, start uint16, values []uint16) (string, func()) {
	t.Helper()
	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil { t.Fatal(err) }
	done := make(chan struct{})
	go func() {
		defer close(done)
		conn, err := ln.Accept()
		if err != nil { return }
		defer conn.Close()
		request := make([]byte, 12)
		if _, err := io.ReadFull(conn, request); err != nil { return }
		if request[7] != function || binary.BigEndian.Uint16(request[8:10]) != start || binary.BigEndian.Uint16(request[10:12]) != uint16(len(values)) { return }
		payloadBytes := len(values) * 2
		response := make([]byte, 9+payloadBytes)
		copy(response[0:2], request[0:2])
		binary.BigEndian.PutUint16(response[2:4], 0)
		binary.BigEndian.PutUint16(response[4:6], uint16(3+payloadBytes))
		response[6], response[7], response[8] = request[6], function, byte(payloadBytes)
		for i, value := range values { binary.BigEndian.PutUint16(response[9+i*2:11+i*2], value) }
		_, _ = conn.Write(response)
	}()
	return ln.Addr().String(), func() { _ = ln.Close(); <-done }
}

func splitTestAddress(t *testing.T, addr string) (string, uint16) {
	t.Helper()
	host, portText, err := net.SplitHostPort(addr)
	if err != nil { t.Fatal(err) }
	var port uint64
	for _, ch := range portText { port = port*10 + uint64(ch-'0') }
	return host, uint16(port)
}
