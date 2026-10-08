package simulator

import (
	"bytes"
	"context"
	"encoding/binary"
	"io"
	"net"
	"strconv"
	"testing"
	"time"

	"github.com/tamzrod/MCS.OSJS/mma2composer"
)

// fakeModbus serves one FC3/FC1 response with the given payload, so the area
// reader's Modbus decode path can be exercised without a real MMA2.
func fakeModbus(t *testing.T, fc byte, payload []byte) uint16 {
	t.Helper()
	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { _ = ln.Close() })
	go func() {
		for {
			conn, err := ln.Accept()
			if err != nil {
				return
			}
			go func(c net.Conn) {
				defer c.Close()
				_ = c.SetDeadline(time.Now().Add(2 * time.Second))
				req := make([]byte, 12)
				if _, err := io.ReadFull(c, req); err != nil {
					return
				}
				resp := make([]byte, 9+len(payload))
				binary.BigEndian.PutUint16(resp[0:2], 1)
				binary.BigEndian.PutUint16(resp[4:6], uint16(3+len(payload)))
				resp[6] = req[6]
				resp[7] = fc
				resp[8] = byte(len(payload))
				copy(resp[9:], payload)
				_, _ = c.Write(resp)
			}(conn)
		}
	}()
	_, portText, _ := net.SplitHostPort(ln.Addr().String())
	port, _ := strconv.Atoi(portText)
	return uint16(port)
}

// modbusReadArea registers round-trips register and bit payloads deterministically.
func TestModbusReadAreaRegisterAndBit(t *testing.T) {
	regPayload := []byte{0x01, 0x02, 0x03, 0x04}
	port := fakeModbus(t, 3, regPayload)
	got, err := modbusReadArea("127.0.0.1", port, 1, "holding_registers", mma2composer.PersistenceRegisters, 0, 2)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(got, regPayload) {
		t.Fatalf("register read wrong: %v", got)
	}

	bitPayload := []byte{0xAB}
	port = fakeModbus(t, 1, bitPayload)
	got, err = modbusReadArea("127.0.0.1", port, 1, "coils", mma2composer.PersistenceBits, 0, 8)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(got, bitPayload) {
		t.Fatalf("bit read wrong: %v", got)
	}
}

// A wrong-length payload is an explicit error, never a fabricated snapshot.
func TestModbusReadAreaRejectsWrongLength(t *testing.T) {
	port := fakeModbus(t, 3, []byte{0x01, 0x02}) // 1 reg, but we ask for 2
	if _, err := modbusReadArea("127.0.0.1", port, 1, "holding_registers", mma2composer.PersistenceRegisters, 0, 2); err == nil {
		t.Fatal("a wrong-length read must fail closed")
	}
}

// persistenceAreaReader routes to the device's configured port/unit.
func TestPersistenceAreaReaderUsesDeviceIdentity(t *testing.T) {
	payload := []byte{0x0A, 0x0B}
	port := fakeModbus(t, 3, payload)
	device := DeviceDefinition{MMA2: MMA2Params{Port: port, UnitID: 1}}
	got, err := persistenceAreaReader(device).ReadPersistenceArea(mma2composer.PersistenceMemoryKey{Port: port, UnitID: 1}, "holding_registers", mma2composer.PersistenceRegisters, 0, 1)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(got, payload) {
		t.Fatalf("reader payload wrong: %v", got)
	}
}

type recordingHandler struct{ ids []uint8 }

func (h *recordingHandler) PublishPersistenceEvent(id uint8) { h.ids = append(h.ids, id) }

// PERSIST-R02 self-check: the RBE subscriber forwards non-zero one-byte IDs and
// ignores the reserved zero byte.
func TestPersistenceRBESubscriberForwardsIDs(t *testing.T) {
	handler := &recordingHandler{}
	sub, err := NewPersistenceRBESubscriber(handler)
	if err != nil {
		t.Fatal(err)
	}
	if err := sub.Subscribe(context.Background(), bytes.NewReader([]byte{3, 0, 7, 99})); err != nil {
		t.Fatal(err)
	}
	if len(handler.ids) != 3 || handler.ids[0] != 3 || handler.ids[1] != 7 || handler.ids[2] != 99 {
		t.Fatalf("forwarded IDs wrong: %v", handler.ids)
	}
	if _, err := NewPersistenceRBESubscriber(nil); err == nil {
		t.Fatal("nil handler must be rejected")
	}
}

func TestPersistenceRBESubscriberDialAndSubscribe(t *testing.T) {
	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { _ = ln.Close() })
	go func() {
		conn, err := ln.Accept()
		if err != nil {
			return
		}
		defer conn.Close()
		_, _ = conn.Write([]byte{5, 0, 6})
	}()

	handler := &recordingHandler{}
	sub, _ := NewPersistenceRBESubscriber(handler)
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	if err := sub.DialAndSubscribe(ctx, ln.Addr().String(), time.Second); err != nil {
		t.Fatalf("dial/subscribe failed: %v", err)
	}
	if len(handler.ids) != 2 || handler.ids[0] != 5 || handler.ids[1] != 6 {
		t.Fatalf("dialed IDs wrong: %v", handler.ids)
	}
}


func TestModbusReadAreaUsesActualFCForEachPersistenceArea(t *testing.T) {
	tests := []struct {
		area string
		kind mma2composer.PersistenceAreaKind
		fc byte
		payload []byte
		count uint16
	}{
		{"coils", mma2composer.PersistenceBits, 1, []byte{0x01}, 1},
		{"discrete_inputs", mma2composer.PersistenceBits, 2, []byte{0x01}, 1},
		{"holding_registers", mma2composer.PersistenceRegisters, 3, []byte{0x00, 0x01}, 1},
		{"input_registers", mma2composer.PersistenceRegisters, 4, []byte{0x00, 0x01}, 1},
	}
	for _, tc := range tests {
		t.Run(tc.area, func(t *testing.T) {
			ln, err := net.Listen("tcp", "127.0.0.1:0")
			if err != nil { t.Fatal(err) }
			defer ln.Close()
			go func() {
				conn, err := ln.Accept()
				if err != nil { return }
				defer conn.Close()
				req := make([]byte, 12)
				if _, err := io.ReadFull(conn, req); err != nil { return }
				if req[7] != tc.fc { return }
				resp := make([]byte, 9+len(tc.payload))
				binary.BigEndian.PutUint16(resp[0:2], binary.BigEndian.Uint16(req[0:2]))
				binary.BigEndian.PutUint16(resp[4:6], uint16(3+len(tc.payload)))
				resp[6], resp[7], resp[8] = req[6], tc.fc, byte(len(tc.payload))
				copy(resp[9:], tc.payload)
				_, _ = conn.Write(resp)
			}()
			_, portText, _ := net.SplitHostPort(ln.Addr().String())
			n, _ := strconv.Atoi(portText)
			got, err := modbusReadArea("127.0.0.1", uint16(n), 1, tc.area, tc.kind, 0, tc.count)
			if err != nil { t.Fatal(err) }
			if !bytes.Equal(got, tc.payload) { t.Fatalf("got %v want %v", got, tc.payload) }
		})
	}
}
