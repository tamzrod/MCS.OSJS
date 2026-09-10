// Package mma2raw implements the producer-neutral MMA2 Raw Ingest v1 client.
package mma2raw

import (
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"net"
	"time"
)

const (
	Version    = byte(0x01)
	ResponseOK = byte(0x00)
	HeaderLen  = 10
)

type Area uint8

const (
	Coils Area = iota + 1
	DiscreteInputs
	HoldingRegisters
	InputRegisters
)

type Range struct {
	Start uint16
	Count uint16
}

type Values struct {
	Bits      []bool
	Registers []uint16
}

type Client struct {
	Addr    string
	UnitID  uint16
	Ranges  map[Area]Range
	Timeout time.Duration
}

func NewClient(addr string, unitID uint16, ranges map[Area]Range) *Client {
	return &Client{Addr: addr, UnitID: unitID, Ranges: ranges, Timeout: 5 * time.Second}
}

func Encode(unitID, addr uint16, area Area, values Values) ([]byte, error) {
	if area < Coils || area > InputRegisters {
		return nil, fmt.Errorf("unsupported raw-ingest area: %d", area)
	}
	count := len(values.Registers)
	bitArea := area == Coils || area == DiscreteInputs
	if bitArea {
		count = len(values.Bits)
		if len(values.Registers) != 0 {
			return nil, errors.New("bit-area write cannot contain registers")
		}
	} else if len(values.Bits) != 0 {
		return nil, errors.New("register-area write cannot contain bits")
	}
	if count == 0 {
		return nil, errors.New("cannot encode an empty raw-ingest batch")
	}
	if count > 65535 {
		return nil, errors.New("raw-ingest batch exceeds uint16 count")
	}

	payloadLen := count * 2
	if bitArea {
		payloadLen = (count + 7) / 8
	}
	pkt := make([]byte, HeaderLen+payloadLen)
	pkt[0], pkt[1], pkt[2], pkt[3] = 'R', 'I', Version, byte(area)
	binary.BigEndian.PutUint16(pkt[4:6], unitID)
	binary.BigEndian.PutUint16(pkt[6:8], addr)
	binary.BigEndian.PutUint16(pkt[8:10], uint16(count))
	if bitArea {
		for i, on := range values.Bits {
			if on {
				pkt[HeaderLen+i/8] |= byte(1 << (i % 8))
			}
		}
	} else {
		for i, value := range values.Registers {
			binary.BigEndian.PutUint16(pkt[HeaderLen+i*2:], value)
		}
	}
	return pkt, nil
}

func (c *Client) Send(area Area, values Values) error {
	configured, ok := c.Ranges[area]
	if !ok || configured.Count == 0 {
		return fmt.Errorf("raw-ingest area %d is not configured", area)
	}
	count := len(values.Registers)
	if area == Coils || area == DiscreteInputs {
		count = len(values.Bits)
	}
	if count > int(configured.Count) {
		return fmt.Errorf("raw-ingest area %d batch count %d exceeds configured count", area, count)
	}
	pkt, err := Encode(c.UnitID, configured.Start, area, values)
	if err != nil {
		return err
	}
	timeout := c.Timeout
	if timeout <= 0 {
		timeout = 5 * time.Second
	}
	conn, err := net.DialTimeout("tcp", c.Addr, timeout)
	if err != nil {
		return fmt.Errorf("raw ingest dial %s: %w", c.Addr, err)
	}
	defer conn.Close()
	_ = conn.SetDeadline(time.Now().Add(timeout))
	if _, err := conn.Write(pkt); err != nil {
		return fmt.Errorf("raw ingest write: %w", err)
	}
	var response [1]byte
	if _, err := io.ReadFull(conn, response[:]); err != nil {
		return fmt.Errorf("raw ingest response: %w", err)
	}
	if response[0] != ResponseOK {
		return fmt.Errorf("raw ingest rejected: response code 0x%02X", response[0])
	}
	return nil
}
