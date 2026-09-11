package replicator

import (
	"encoding/binary"
	"fmt"
	"io"
	"net"
	"time"
)

const defaultModbusTimeout = 2 * time.Second

// RegisterValues carries one Modbus source payload. Bits is used for FC1/FC2;
// Values is used for FC3/FC4. The legacy type name is retained for compatibility.
type RegisterValues struct {
	Function uint8
	Start    uint16
	Bits     []bool
	Values   []uint16
}

func (s Store) ReadConfiguredSource() (RegisterValues, error) {
	cfg, err := s.Load()
	if err != nil {
		return RegisterValues{}, fmt.Errorf("load replicator config: %w", err)
	}
	if err := ValidateConfig(cfg); err != nil {
		return RegisterValues{}, fmt.Errorf("validate replicator config: %w", err)
	}
	return ReadSourceRange(cfg.Source)
}

// ReadSourceRange performs one deterministic Modbus TCP read for FC1-FC4.
func ReadSourceRange(source SourceConfig) (RegisterValues, error) {
	return readSourceRange(source, defaultModbusTimeout)
}

func readSourceRange(source SourceConfig, timeout time.Duration) (RegisterValues, error) {
	if source.Function < 1 || source.Function > 4 {
		return RegisterValues{}, fmt.Errorf("source.function must be FC1, FC2, FC3, or FC4")
	}
	if source.Port == 0 {
		return RegisterValues{}, fmt.Errorf("source.port must be > 0")
	}
	if source.UnitID > 0xFF {
		return RegisterValues{}, fmt.Errorf("source.unit_id must be <= 255")
	}
	if source.Count == 0 {
		return RegisterValues{}, fmt.Errorf("source.count must be > 0")
	}
	if uint32(source.Start)+uint32(source.Count) > 0x10000 {
		return RegisterValues{}, fmt.Errorf("source range exceeds 16-bit address space")
	}

	address := net.JoinHostPort(source.Host, fmt.Sprintf("%d", source.Port))
	conn, err := net.DialTimeout("tcp", address, timeout)
	if err != nil {
		return RegisterValues{}, fmt.Errorf("connect %s: %w", address, err)
	}
	defer conn.Close()
	if err := conn.SetDeadline(time.Now().Add(timeout)); err != nil {
		return RegisterValues{}, fmt.Errorf("set deadline: %w", err)
	}

	const transactionID uint16 = 1
	request := make([]byte, 12)
	binary.BigEndian.PutUint16(request[0:2], transactionID)
	binary.BigEndian.PutUint16(request[2:4], 0)
	binary.BigEndian.PutUint16(request[4:6], 6)
	request[6] = byte(source.UnitID)
	request[7] = source.Function
	binary.BigEndian.PutUint16(request[8:10], source.Start)
	binary.BigEndian.PutUint16(request[10:12], source.Count)
	if _, err := conn.Write(request); err != nil {
		return RegisterValues{}, fmt.Errorf("write Modbus request: %w", err)
	}

	header := make([]byte, 7)
	if _, err := io.ReadFull(conn, header); err != nil {
		return RegisterValues{}, fmt.Errorf("read Modbus header: %w", err)
	}
	if binary.BigEndian.Uint16(header[0:2]) != transactionID {
		return RegisterValues{}, fmt.Errorf("unexpected transaction id")
	}
	if binary.BigEndian.Uint16(header[2:4]) != 0 {
		return RegisterValues{}, fmt.Errorf("unexpected Modbus protocol id")
	}
	if header[6] != byte(source.UnitID) {
		return RegisterValues{}, fmt.Errorf("unexpected unit id %d", header[6])
	}
	length := binary.BigEndian.Uint16(header[4:6])
	if length < 3 {
		return RegisterValues{}, fmt.Errorf("invalid Modbus response length %d", length)
	}
	pdu := make([]byte, int(length)-1)
	if _, err := io.ReadFull(conn, pdu); err != nil {
		return RegisterValues{}, fmt.Errorf("read Modbus PDU: %w", err)
	}
	if pdu[0] == source.Function|0x80 {
		if len(pdu) < 2 {
			return RegisterValues{}, fmt.Errorf("malformed Modbus exception response")
		}
		return RegisterValues{}, fmt.Errorf("Modbus exception code %d", pdu[1])
	}
	if pdu[0] != source.Function {
		return RegisterValues{}, fmt.Errorf("unexpected function code %d", pdu[0])
	}
	if len(pdu) < 2 {
		return RegisterValues{}, fmt.Errorf("malformed Modbus response")
	}

	if source.Function == 1 || source.Function == 2 {
		expectedBytes := (int(source.Count) + 7) / 8
		if int(pdu[1]) != expectedBytes || len(pdu) != expectedBytes+2 {
			return RegisterValues{}, fmt.Errorf("unexpected bit payload length: byte_count=%d pdu_len=%d expected=%d", pdu[1], len(pdu), expectedBytes)
		}
		bits := make([]bool, source.Count)
		for i := range bits {
			bits[i] = pdu[2+i/8]&(1<<uint(i%8)) != 0
		}
		return RegisterValues{Function: source.Function, Start: source.Start, Bits: bits}, nil
	}

	expectedBytes := int(source.Count) * 2
	if int(pdu[1]) != expectedBytes || len(pdu) != expectedBytes+2 {
		return RegisterValues{}, fmt.Errorf("unexpected register payload length: byte_count=%d pdu_len=%d expected=%d", pdu[1], len(pdu), expectedBytes)
	}
	values := make([]uint16, source.Count)
	for i := range values {
		offset := 2 + i*2
		values[i] = binary.BigEndian.Uint16(pdu[offset : offset+2])
	}
	return RegisterValues{Function: source.Function, Start: source.Start, Values: values}, nil
}
