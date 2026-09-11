package replicator

import (
	"encoding/binary"
	"io"
	"net"
	"reflect"
	"testing"
	"time"
)

func TestReadSourceRangeFC3(t *testing.T) {
	addr, stop := startTestModbusServer(t, 3, 12, []uint16{100, 200, 300})
	defer stop()

	host, port := splitTestAddress(t, addr)
	got, err := readSourceRange(SourceConfig{
		Host: host, Port: port, UnitID: 7, Function: 3, Start: 12, Count: 3,
	}, time.Second)
	if err != nil {
		t.Fatalf("readSourceRange: %v", err)
	}

	want := RegisterValues{Function: 3, Start: 12, Values: []uint16{100, 200, 300}}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got %#v want %#v", got, want)
	}
}

func TestReadSourceRangeFC4(t *testing.T) {
	addr, stop := startTestModbusServer(t, 4, 21, []uint16{0x1234, 0xabcd})
	defer stop()

	host, port := splitTestAddress(t, addr)
	got, err := readSourceRange(SourceConfig{
		Host: host, Port: port, UnitID: 7, Function: 4, Start: 21, Count: 2,
	}, time.Second)
	if err != nil {
		t.Fatalf("readSourceRange: %v", err)
	}
	if !reflect.DeepEqual(got.Values, []uint16{0x1234, 0xabcd}) {
		t.Fatalf("values = %#v", got.Values)
	}
}

func TestReadSourceRangeRejectsNonRegisterFunction(t *testing.T) {
	_, err := ReadSourceRange(SourceConfig{Host: "127.0.0.1", Port: 502, UnitID: 1, Function: 1, Start: 0, Count: 1})
	if err == nil {
		t.Fatal("expected unsupported-function error")
	}
}

func TestReadSourceRangeConnectionFailure(t *testing.T) {
	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	addr := ln.Addr().String()
	_ = ln.Close()

	host, port := splitTestAddress(t, addr)
	_, err = readSourceRange(SourceConfig{Host: host, Port: port, UnitID: 1, Function: 3, Start: 0, Count: 1}, 100*time.Millisecond)
	if err == nil {
		t.Fatal("expected connection error")
	}
}

func startTestModbusServer(t *testing.T, function uint8, start uint16, values []uint16) (string, func()) {
	t.Helper()
	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	done := make(chan struct{})
	go func() {
		defer close(done)
		conn, err := ln.Accept()
		if err != nil {
			return
		}
		defer conn.Close()

		request := make([]byte, 12)
		if _, err := io.ReadFull(conn, request); err != nil {
			return
		}
		if request[7] != function || binary.BigEndian.Uint16(request[8:10]) != start || binary.BigEndian.Uint16(request[10:12]) != uint16(len(values)) {
			return
		}

		payloadBytes := len(values) * 2
		response := make([]byte, 9+payloadBytes)
		copy(response[0:2], request[0:2])
		binary.BigEndian.PutUint16(response[2:4], 0)
		binary.BigEndian.PutUint16(response[4:6], uint16(3+payloadBytes))
		response[6] = request[6]
		response[7] = function
		response[8] = byte(payloadBytes)
		for i, value := range values {
			binary.BigEndian.PutUint16(response[9+i*2:11+i*2], value)
		}
		_, _ = conn.Write(response)
	}()

	return ln.Addr().String(), func() {
		_ = ln.Close()
		<-done
	}
}

func splitTestAddress(t *testing.T, addr string) (string, uint16) {
	t.Helper()
	host, portText, err := net.SplitHostPort(addr)
	if err != nil {
		t.Fatal(err)
	}
	var port uint64
	for _, ch := range portText {
		port = port*10 + uint64(ch-'0')
	}
	return host, uint16(port)
}
