package mma2composer

import (
	"encoding/binary"
	"fmt"
)

// PersistenceSnapshotSize returns the exact raw snapshot byte length for a
// configured area count. The raw file contains payload only; start/count and
// compatibility metadata belong to the manifest layer (PERSIST-014).
func PersistenceSnapshotSize(kind PersistenceAreaKind, count uint16) (int, error) {
	switch kind {
	case PersistenceBits:
		return (int(count) + 7) / 8, nil
	case PersistenceRegisters:
		return int(count) * 2, nil
	default:
		return 0, fmt.Errorf("unknown persistence area kind %d", kind)
	}
}

// EncodePersistenceBits packs exactly count bit values LSB-first within each
// byte, matching Modbus bit packing. Unused high bits in the final byte are zero.
func EncodePersistenceBits(values []bool, count uint16) ([]byte, error) {
	if len(values) != int(count) {
		return nil, fmt.Errorf("persistence bit value count %d does not match configured count %d", len(values), count)
	}
	size, _ := PersistenceSnapshotSize(PersistenceBits, count)
	out := make([]byte, size)
	for i, value := range values {
		if value {
			out[i/8] |= 1 << uint(i%8)
		}
	}
	return out, nil
}

// DecodePersistenceBits validates the raw image length and expands exactly count
// LSB-first packed bits. Padding bits beyond count are ignored.
func DecodePersistenceBits(raw []byte, count uint16) ([]bool, error) {
	expected, _ := PersistenceSnapshotSize(PersistenceBits, count)
	if len(raw) != expected {
		return nil, fmt.Errorf("persistence bit snapshot length %d does not match expected %d", len(raw), expected)
	}
	out := make([]bool, int(count))
	for i := range out {
		out[i] = raw[i/8]&(1<<uint(i%8)) != 0
	}
	return out, nil
}

// EncodePersistenceRegisters writes exactly count uint16 values in big-endian
// byte order, matching Modbus register representation.
func EncodePersistenceRegisters(values []uint16, count uint16) ([]byte, error) {
	if len(values) != int(count) {
		return nil, fmt.Errorf("persistence register value count %d does not match configured count %d", len(values), count)
	}
	size, _ := PersistenceSnapshotSize(PersistenceRegisters, count)
	out := make([]byte, size)
	for i, value := range values {
		binary.BigEndian.PutUint16(out[i*2:i*2+2], value)
	}
	return out, nil
}

// DecodePersistenceRegisters validates the raw image length and decodes exactly
// count big-endian uint16 values.
func DecodePersistenceRegisters(raw []byte, count uint16) ([]uint16, error) {
	expected, _ := PersistenceSnapshotSize(PersistenceRegisters, count)
	if len(raw) != expected {
		return nil, fmt.Errorf("persistence register snapshot length %d does not match expected %d", len(raw), expected)
	}
	out := make([]uint16, int(count))
	for i := range out {
		out[i] = binary.BigEndian.Uint16(raw[i*2 : i*2+2])
	}
	return out, nil
}

// ValidatePersistenceSnapshotLength checks that raw payload length matches the
// deterministic size for the configured area kind/count.
func ValidatePersistenceSnapshotLength(kind PersistenceAreaKind, count uint16, raw []byte) error {
	expected, err := PersistenceSnapshotSize(kind, count)
	if err != nil {
		return err
	}
	if len(raw) != expected {
		return fmt.Errorf("persistence snapshot length %d does not match expected %d", len(raw), expected)
	}
	return nil
}
