package simulator

import (
	"encoding/binary"
	"fmt"
	"io"
	"net"
	"time"

	"github.com/tamzrod/MCS.OSJS/mma2composer"
)

const persistenceModbusTimeout = 2 * time.Second

// modbusReadArea performs one deterministic Modbus TCP read of a memory area and
// returns the raw snapshot encoding the persistence contracts expect: LSB-first
// packed bits for bit areas, big-endian uint16 words for register areas. It is
// the authoritative current-state source for the persistence save path; it never
// writes and never fabricates bytes (a failure is an explicit error).
func modbusReadArea(host string, port, unitID uint16, area mma2composer.PersistenceAreaKind, start, count uint16) ([]byte, error) {
	if count == 0 {
		return nil, fmt.Errorf("persistence area read requires a non-empty range")
	}
	address := net.JoinHostPort(host, fmt.Sprintf("%d", port))
	conn, err := net.DialTimeout("tcp", address, persistenceModbusTimeout)
	if err != nil {
		return nil, fmt.Errorf("connect %s: %w", address, err)
	}
	defer conn.Close()
	if err := conn.SetDeadline(time.Now().Add(persistenceModbusTimeout)); err != nil {
		return nil, fmt.Errorf("set deadline: %w", err)
	}

	function := byte(3)
	if area == mma2composer.PersistenceBits {
		function = 1
	}

	const transactionID uint16 = 1
	request := make([]byte, 12)
	binary.BigEndian.PutUint16(request[0:2], transactionID)
	binary.BigEndian.PutUint16(request[2:4], 0)
	binary.BigEndian.PutUint16(request[4:6], 6)
	request[6] = byte(unitID)
	request[7] = function
	binary.BigEndian.PutUint16(request[8:10], start)
	binary.BigEndian.PutUint16(request[10:12], count)
	if _, err := conn.Write(request); err != nil {
		return nil, fmt.Errorf("write read request: %w", err)
	}

	var mbap [7]byte
	if _, err := io.ReadFull(conn, mbap[:]); err != nil {
		return nil, fmt.Errorf("read response header: %w", err)
	}
	// MBAP length counts the unit id plus the PDU.
	pduLength := int(binary.BigEndian.Uint16(mbap[4:6]))
	if pduLength < 3 {
		return nil, fmt.Errorf("modbus response length %d is too short", pduLength)
	}
	pdu := make([]byte, pduLength-1)
	if _, err := io.ReadFull(conn, pdu); err != nil {
		return nil, fmt.Errorf("read response body: %w", err)
	}
	if pdu[0]&0x80 != 0 {
		return nil, fmt.Errorf("modbus exception function 0x%02X code 0x%02X", pdu[0], pdu[1])
	}
	byteCount := int(pdu[1])
	data := pdu[2:]
	if byteCount != len(data) {
		return nil, fmt.Errorf("modbus byte count %d does not match payload %d", byteCount, len(data))
	}

	if area == mma2composer.PersistenceBits {
		expected := (int(count) + 7) / 8
		if len(data) != expected {
			return nil, fmt.Errorf("modbus bit payload length %d does not match expected %d", len(data), expected)
		}
		return append([]byte(nil), data...), nil
	}
	expected := int(count) * 2
	if len(data) != expected {
		return nil, fmt.Errorf("modbus register payload length %d does not match expected %d", len(data), expected)
	}
	return append([]byte(nil), data...), nil
}

// persistenceAreaReader returns the authoritative current-state reader for one
// device's MMA2 memory, backed by a real Modbus TCP read against the device's
// configured listen port.
func persistenceAreaReader(device DeviceDefinition) mma2composer.PersistenceAreaReader {
	return PersistenceAreaReaderFunc(func(key mma2composer.PersistenceMemoryKey, area string, kind mma2composer.PersistenceAreaKind, start, count uint16) ([]byte, error) {
		return modbusReadArea("127.0.0.1", device.MMA2.Port, device.MMA2.UnitID, kind, start, count)
	})
}
