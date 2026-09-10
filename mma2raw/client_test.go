package mma2raw

import (
	"encoding/binary"
	"io"
	"net"
	"testing"
)

func TestEncodeAreaFixtures(t *testing.T) {
	tests := []struct {
		name    string
		area    Area
		values  Values
		count   uint16
		payload []byte
	}{
		{"coils", Coils, Values{Bits: []bool{true, false, true, true, false, false, false, false, true}}, 9, []byte{0x0d, 0x01}},
		{"discrete inputs", DiscreteInputs, Values{Bits: []bool{false, true, false, false, true}}, 5, []byte{0x12}},
		{"holding registers", HoldingRegisters, Values{Registers: []uint16{0x1234, 0x00ab, 0xffff}}, 3, []byte{0x12, 0x34, 0x00, 0xab, 0xff, 0xff}},
		{"input registers", InputRegisters, Values{Registers: []uint16{0x0102, 0xa0b0}}, 2, []byte{0x01, 0x02, 0xa0, 0xb0}},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			pkt, err := Encode(7, 11, tc.area, tc.values)
			if err != nil {
				t.Fatal(err)
			}
			if string(pkt[:4]) != string([]byte{'R', 'I', Version, byte(tc.area)}) || binary.BigEndian.Uint16(pkt[4:6]) != 7 || binary.BigEndian.Uint16(pkt[6:8]) != 11 || binary.BigEndian.Uint16(pkt[8:10]) != tc.count {
				t.Fatalf("unexpected header: % x", pkt[:HeaderLen])
			}
			if string(pkt[HeaderLen:]) != string(tc.payload) {
				t.Fatalf("payload % x != % x", pkt[HeaderLen:], tc.payload)
			}
		})
	}
}

func TestClientSendAndResponse(t *testing.T) {
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	defer listener.Close()
	received := make(chan []byte, 1)
	go func() {
		conn, acceptErr := listener.Accept()
		if acceptErr != nil {
			return
		}
		defer conn.Close()
		packet := make([]byte, HeaderLen+2)
		_, _ = io.ReadFull(conn, packet)
		received <- packet
		_, _ = conn.Write([]byte{ResponseOK})
	}()
	client := NewClient(listener.Addr().String(), 3, map[Area]Range{HoldingRegisters: {Start: 9, Count: 2}})
	if err := client.Send(HoldingRegisters, Values{Registers: []uint16{0x5432}}); err != nil {
		t.Fatal(err)
	}
	packet := <-received
	if binary.BigEndian.Uint16(packet[4:6]) != 3 || binary.BigEndian.Uint16(packet[6:8]) != 9 || string(packet[HeaderLen:]) != string([]byte{0x54, 0x32}) {
		t.Fatalf("unexpected packet: % x", packet)
	}
}
