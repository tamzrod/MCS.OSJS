package simulator

import (
	"encoding/binary"
	"io"
	"net"
	"sync"
	"testing"
)

type rawFixture struct {
	ln      net.Listener
	mu      sync.Mutex
	codes   []byte
	lastPkt []byte
}

func newRawFixture(t *testing.T, codes ...byte) *rawFixture {
	t.Helper()
	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	f := &rawFixture{ln: ln, codes: codes}
	go f.serve(t)
	t.Cleanup(func() { _ = ln.Close() })
	return f
}

func (f *rawFixture) serve(t *testing.T) {
	for {
		conn, err := f.ln.Accept()
		if err != nil {
			return
		}
		go func(c net.Conn) {
			defer c.Close()
			hdr := make([]byte, rawHeaderLen)
			if _, err := io.ReadFull(c, hdr); err != nil {
				return
			}
			plen := int(binary.BigEndian.Uint16(hdr[8:10])) * 2
			if hdr[3] == 1 || hdr[3] == 2 {
				plen = (int(binary.BigEndian.Uint16(hdr[8:10])) + 7) / 8
			}
			payload := make([]byte, plen)
			if _, err := io.ReadFull(c, payload); err != nil {
				return
			}
			f.mu.Lock()
			f.lastPkt = append([]byte(nil), hdr...)
			f.lastPkt = append(f.lastPkt, payload...)
			var code byte
			if len(f.codes) > 0 {
				code, f.codes = f.codes[0], f.codes[1:]
			}
			f.mu.Unlock()
			_, _ = c.Write([]byte{code})
		}(conn)
	}
}

func (f *rawFixture) lastPacket(t *testing.T) []byte {
	t.Helper()
	f.mu.Lock()
	defer f.mu.Unlock()
	return append([]byte(nil), f.lastPkt...)
}

func assertRawFrame(t *testing.T, pkt []byte, unitID, addr, count uint16, area byte, payloadLen int) {
	t.Helper()
	if len(pkt) != rawHeaderLen+payloadLen {
		t.Fatalf("packet length %d != %d", len(pkt), rawHeaderLen+payloadLen)
	}
	if pkt[0] != 'R' || pkt[1] != 'I' || pkt[2] != rawVersion {
		t.Fatalf("packet is not a v1 raw-ingest frame")
	}
	if pkt[3] != area {
		t.Fatalf("area byte %d != %d", pkt[3], area)
	}
	if got := binary.BigEndian.Uint16(pkt[4:6]); got != unitID {
		t.Fatalf("unit id %d != %d", got, unitID)
	}
	if got := binary.BigEndian.Uint16(pkt[6:8]); got != addr {
		t.Fatalf("addr %d != %d", got, addr)
	}
	if got := binary.BigEndian.Uint16(pkt[8:10]); got != count {
		t.Fatalf("count %d != %d", got, count)
	}
}
func TestRawIngestClientSendOK(t *testing.T) {
	cl := newRawFixture(t, rawRespOK)
	v := Values{FC: FC3, Regs: []uint16{0x5432}}
	def := validDevice()
	cli := NewRawIngestClient(def)
	cli.addr = cl.ln.Addr().String()
	if err := cli.Send(v); err != nil {
		t.Fatal(err)
	}
	pkt := cl.lastPacket(t)
	assertRawFrame(t, pkt, def.MMA2.UnitID, def.MMA2.FC3.Start, 1, 3, 2)
	if pkt[rawHeaderLen] != 0x54 || pkt[rawHeaderLen+1] != 0x32 {
		t.Fatalf("reg payload wrong: % x", pkt[rawHeaderLen:])
	}
}

func TestRawIngestClientSendRejectsUnconfiguredFC(t *testing.T) {
	cl := newRawFixture(t, rawRespOK)
	v := Values{FC: FC4, Regs: []uint16{1}}
	def := validDevice()
	def.MMA2.FC4.Count = 0
	cli := NewRawIngestClient(def)
	cli.addr = cl.ln.Addr().String()
	if err := cli.Send(v); err == nil {
		t.Fatal("expected error for FC4 (not configured for this fixture)")
	}
}

func TestRawIngestClientSendResponseError(t *testing.T) {
	cl := newRawFixture(t, 0x21)
	v := Values{FC: FC3, Regs: []uint16{1}}
	def := validDevice()
	cli := NewRawIngestClient(def)
	cli.addr = cl.ln.Addr().String()
	if err := cli.Send(v); err == nil {
		t.Fatal("expected rejected response error")
	}
}
