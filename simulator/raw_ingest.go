package simulator

import (
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"net"
	"strconv"
	"time"
)

// RawIngest v1 protocol constants per MMA2/docs/RAW_INGEST.md.
const (
	rawMagic0      = byte('R')
	rawMagic1      = byte('I')
	rawVersion     = byte(0x01)
	rawRespOK      = byte(0x00)
	rawHeaderLen   = 10 // Magic(2) + Version(1) + Area(1) + UnitID(2) + Addr(2) + Count(2)
	rawSendTimeout = 5 * time.Second
)

// rawArea maps the simulator FC kind to the MMA2 raw-ingest area code:
// FC1=coils,FC2=discrete inputs,FC3=holding regs,FC4=input regs
func rawArea(fc FCKind) uint8 {
	switch fc {
	case FC1:
		return 1
	case FC2:
		return 2
	case FC3:
		return 3
	case FC4:
		return 4
	default:
		return 0
	}
}

// rawAreaStart returns the configured simulator-owned start address for an FC area
func rawAreaStart(p MMA2Params, fc FCKind) uint16 {
	switch fc {
	case FC1:
		return p.FC1.Start
	case FC2:
		return p.FC2.Start
	case FC3:
		return p.FC3.Start
	case FC4:
		return p.FC4.Start
	default:
		return 0
	}
}

// rawAreaConfigured reports whether the FC area has a nonzero count in the device's
// simulator-owned MMA2 range
func rawAreaConfigured(p MMA2Params, fc FCKind) bool {
	switch fc {
	case FC1:
		return p.FC1.Count > 0
	case FC2:
		return p.FC2.Count > 0
	case FC3:
		return p.FC3.Count > 0
	case FC4:
		return p.FC4.Count > 0
	default:
		return false
	}
}

// encodeRawPacket encodes one v1 raw-ingest write packet for a single FC batch.
// Bit areas(FC1/FC2) are packed LSB-first into ceil(count/8) bytes;reg areas
// (FC3/FC4) are big-endian uint16 words, 2 bytes each. See RAW_INGEST.md..
func encodeRawPacket(unitID, addr, count uint16, area uint8, values Values) ([]byte, error) {
	if area == 0 || area > 4 {
		return nil, fmt.Errorf("unsupported FC area: %d", area)
	}
	if count == 0 {
		return nil, errors.New("cannot encode an empty raw-ingest batch")
	}

	plen := int(count) * 2
	if values.FC == FC1 || values.FC == FC2 {
		plen = int((count + 7) / 8)
	}

	payload := make([]byte, plen)
	switch values.FC {
	case FC1, FC2:
		// LSB-first packed bits:value i occupies byte i/8, bit position i%8..
		for i, on := range values.Coils {
			if on {
				payload[i/8] |= byte(1 << (i % 8))
			}
		}
	case FC3, FC4:
		for i, v := range values.Regs {
			binary.BigEndian.PutUint16(payload[i*2:], v)
		}
	default:
		return nil, fmt.Errorf("unsupported FC kind in values: %d", values.FC)
	}

	pkt := make([]byte, rawHeaderLen+len(payload))
	pkt[0] = rawMagic0
	pkt[1] = rawMagic1
	pkt[2] = rawVersion
	pkt[3] = area
	binary.BigEndian.PutUint16(pkt[4:6], unitID)
	binary.BigEndian.PutUint16(pkt[6:8], addr)
	binary.BigEndian.PutUint16(pkt[8:10], count)
	copy(pkt[rawHeaderLen:], payload)
	return pkt, nil
}

// sendRawPacket delivers one raw-ingest write to the MMA2 listener and returns
// the server's 1-byte response. A non-0x00 response code means the write was
// rejected and produces an error, per RAW_INGEST.md response semantics..
func sendRawPacket(addr string, pkt []byte) (byte, error) {
	conn, err := net.DialTimeout("tcp", addr, rawSendTimeout)
	if err != nil {
		return rawRespOK, fmt.Errorf("raw ingest dial %s: %w", addr, err)
	}
	defer conn.Close()
	_ = conn.SetDeadline(time.Now().Add(rawSendTimeout))

	if _, err := conn.Write(pkt); err != nil {
		return rawRespOK, fmt.Errorf("raw ingest write: %w", err)
	}
	var resp [1]byte
	if _, err := io.ReadFull(conn, resp[:]); err != nil {
		return rawRespOK, fmt.Errorf("raw ingest response: %w", err)
	}
	return resp[0], nil
}

// RawIngestClient sends simulator-generated values to one simulator-owned MMA2
// reservation through the MMA2 raw-ingest transport. It is constructed per device
// definition and targets only that device's configured port/unit/FC ranges..
type RawIngestClient struct {
	addr   string
	unitID uint16
	params MMA2Params
}

// NewRawIngestClient builds a client for one device's simulator-owned MMA2
// reservation. The destination is the configured listener port on loopback..
func NewRawIngestClient(def DeviceDefinition) *RawIngestClient {
	return &RawIngestClient{
		addr:   net.JoinHostPort("127.0.0.1", strconv.Itoa(int(def.MMA2.Port))),
		unitID: def.MMA2.UnitID,
		params: def.MMA2,
	}
}

// Send converts one scheduler batch into the raw-ingest format and writes it to
// MMA2. Only simulator-owned configured FC ranges are ever targeted..
func (c *RawIngestClient) Send(v Values) error {
	if !rawAreaConfigured(c.params, v.FC) {
		return fmt.Errorf("FC%d area is not configured for this device", v.FC)
	}
	area := rawArea(v.FC)
	count := uint16(len(v.Coils) + len(v.Regs))
	if count == 0 {
		return errors.New("cannot send an empty raw-ingest batch")
	}
	if int(count) > int(areaCount(c.params, v.FC)) {
		return fmt.Errorf("FC%d batch count %d exceeds configured count", v.FC, count)
	}

	addr := rawAreaStart(c.params, v.FC)
	pkt, err := encodeRawPacket(c.unitID, addr, count, area, v)
	if err != nil {
		return err
	}
	resp, err := sendRawPacket(c.addr, pkt)
	if err != nil {
		return err
	}
	if resp != rawRespOK {
		return fmt.Errorf("raw ingest rejected: response code 0x%02X", resp)
	}
	return nil
}
