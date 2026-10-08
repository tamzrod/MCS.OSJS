package simulator

import (
	"encoding/binary"
	"fmt"
	"io"
	"net"
	"time"

	"github.com/tamzrod/MCS.OSJS/mma2composer"
)

const persistenceRawIngestTimeout = 2 * time.Second

// persistenceRawIngestWriterAdapter implements
// mma2composer.PersistenceRawIngestWriter over the existing MMA2 Raw Ingest v1
// contract. It submits one Raw Ingest v1 packet per area (magic "RI", version 1,
// area/unit/address/count, packed payload) and returns the single-byte response
// code. The restore path infers start/count from the payload, so the writer
// derives them deterministically and never re-encodes or extends the protocol.
type persistenceRawIngestWriterAdapter struct {
	host   string
	port   uint16
	unitID uint16
}

// NewPersistenceRawIngestWriter returns a writer for the MMA2 Raw Ingest v1
// endpoint of one memory.
func NewPersistenceRawIngestWriter(host string, port, unitID uint16) mma2composer.PersistenceRawIngestWriter {
	return persistenceRawIngestWriterAdapter{host: host, port: port, unitID: unitID}
}

func (w persistenceRawIngestWriterAdapter) WritePersistenceRawIngest(_ mma2composer.PersistenceMemoryKey, area mma2composer.PersistenceRawIngestArea, start, count uint16, payload []byte) (byte, error) {
	expected, err := persistenceRawIngestPayloadLen(area, count)
	if err != nil {
		return 0, err
	}
	if len(payload) != expected {
		return 0, fmt.Errorf("raw ingest area %d payload length %d does not match count %d", area, len(payload), count)
	}
	pkt := make([]byte, 10+len(payload))
	pkt[0], pkt[1], pkt[2], pkt[3] = 'R', 'I', 0x01, byte(area)
	binary.BigEndian.PutUint16(pkt[4:6], w.unitID)
	binary.BigEndian.PutUint16(pkt[6:8], start)
	binary.BigEndian.PutUint16(pkt[8:10], count)
	copy(pkt[10:], payload)

	address := net.JoinHostPort(w.host, fmt.Sprintf("%d", w.port))
	conn, err := net.DialTimeout("tcp", address, persistenceRawIngestTimeout)
	if err != nil {
		return 0, fmt.Errorf("raw ingest dial %s: %w", address, err)
	}
	defer conn.Close()
	if err := conn.SetDeadline(time.Now().Add(persistenceRawIngestTimeout)); err != nil {
		return 0, fmt.Errorf("raw ingest set deadline: %w", err)
	}
	if _, err := conn.Write(pkt); err != nil {
		return 0, fmt.Errorf("raw ingest write: %w", err)
	}
	var response [1]byte
	if _, err := io.ReadFull(conn, response[:]); err != nil {
		return 0, fmt.Errorf("raw ingest response: %w", err)
	}
	return response[0], nil
}

// persistenceRawIngestPayloadLen returns the deterministic payload length for a
// Raw Ingest v1 area write.
func persistenceRawIngestPayloadLen(area mma2composer.PersistenceRawIngestArea, count uint16) (int, error) {
	switch area {
	case mma2composer.RawIngestCoils, mma2composer.RawIngestDiscreteInputs:
		return (int(count) + 7) / 8, nil
	case mma2composer.RawIngestHoldingRegisters, mma2composer.RawIngestInputRegisters:
		return int(count) * 2, nil
	default:
		return 0, fmt.Errorf("raw ingest area %d is unknown", area)
	}
}
