package mma2composer

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

const PersistenceSnapshotFormatVersion = 1

// PersistenceSnapshotManifest binds one raw area snapshot to its exact memory
// identity/layout and payload integrity. It contains metadata only; restore and
// unseal sequencing are handled by later persistence stages.
type PersistenceSnapshotManifest struct {
	FormatVersion int    `yaml:"format_version" json:"format_version"`
	Port          uint16 `yaml:"port" json:"port"`
	UnitID        uint16 `yaml:"unit_id" json:"unit_id"`
	Area          string `yaml:"area" json:"area"`
	Start         uint16 `yaml:"start" json:"start"`
	Count         uint16 `yaml:"count" json:"count"`
	PayloadBytes  int    `yaml:"payload_bytes" json:"payload_bytes"`
	SHA256        string `yaml:"sha256" json:"sha256"`
}

// NewPersistenceSnapshotManifest creates compatibility/integrity metadata for
// one already-encoded raw area payload.
func NewPersistenceSnapshotManifest(key PersistenceMemoryKey, area string, start, count uint16, payload []byte) PersistenceSnapshotManifest {
	sum := sha256.Sum256(payload)
	return PersistenceSnapshotManifest{
		FormatVersion: PersistenceSnapshotFormatVersion,
		Port:          key.Port,
		UnitID:        key.UnitID,
		Area:          area,
		Start:         start,
		Count:         count,
		PayloadBytes:  len(payload),
		SHA256:        hex.EncodeToString(sum[:]),
	}
}

// ValidatePersistenceSnapshotManifest rejects snapshots that do not match the
// current configured memory identity/layout or whose raw payload is incomplete
// or corrupted. No truncation, remapping, restore, or unseal occurs here.
func ValidatePersistenceSnapshotManifest(
	manifest PersistenceSnapshotManifest,
	key PersistenceMemoryKey,
	area string,
	kind PersistenceAreaKind,
	start, count uint16,
	payload []byte,
) error {
	if manifest.FormatVersion != PersistenceSnapshotFormatVersion {
		return fmt.Errorf("persistence snapshot format version %d is incompatible with supported version %d", manifest.FormatVersion, PersistenceSnapshotFormatVersion)
	}
	if manifest.Port != key.Port || manifest.UnitID != key.UnitID {
		return fmt.Errorf("persistence snapshot memory identity mismatch: manifest (%d,%d), configured (%d,%d)", manifest.Port, manifest.UnitID, key.Port, key.UnitID)
	}
	if manifest.Area != area {
		return fmt.Errorf("persistence snapshot area mismatch: manifest %q, configured %q", manifest.Area, area)
	}
	if manifest.Start != start || manifest.Count != count {
		return fmt.Errorf("persistence snapshot layout mismatch: manifest start/count %d/%d, configured %d/%d", manifest.Start, manifest.Count, start, count)
	}

	expectedBytes, err := PersistenceSnapshotSize(kind, count)
	if err != nil {
		return err
	}
	if manifest.PayloadBytes != expectedBytes {
		return fmt.Errorf("persistence snapshot manifest payload size %d does not match configured size %d", manifest.PayloadBytes, expectedBytes)
	}
	if len(payload) != manifest.PayloadBytes {
		return fmt.Errorf("persistence snapshot payload length %d does not match manifest size %d", len(payload), manifest.PayloadBytes)
	}
	if err := ValidatePersistenceSnapshotLength(kind, count, payload); err != nil {
		return err
	}

	sum := sha256.Sum256(payload)
	actual := hex.EncodeToString(sum[:])
	if manifest.SHA256 != actual {
		return fmt.Errorf("persistence snapshot checksum mismatch")
	}
	return nil
}
