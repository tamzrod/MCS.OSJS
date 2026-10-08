package mma2composer

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

// RelPersistenceDir is the appliance-relative directory that holds persistence
// snapshot payloads and manifest metadata beneath the authoritative appliance
// data root. It is the only directory the adapter reads or writes.
const RelPersistenceDir = "persistence/snapshots"

// PersistenceSnapshotConfig registers one configured persisted area under the
// adapter. It carries the same authoritative area identity the loader/restore
// contracts use (kind/start/count) so the adapter never authors an independent
// range or format. Key is the Port -> Unit ID -> Memory identity.
type PersistenceSnapshotConfig struct {
	Key   PersistenceMemoryKey
	Area  string
	Kind  PersistenceAreaKind
	Start uint16
	Count uint16
}

type persistenceAreaIdentity struct {
	Port   uint16
	UnitID uint16
	Area   string
}

// PersistenceFilesystemAdapter is the concrete appliance filesystem
// implementation of the persistence snapshot store/source contracts. It roots
// every file beneath the configured appliance data root using deterministic
// Port -> Unit ID -> Area paths, reuses the existing raw snapshot format and
// manifest metadata, and surfaces malformed/missing content as explicit errors
// rather than fabricated defaults. It never invents a second format or range.
type PersistenceFilesystemAdapter struct {
	root  string
	areas map[persistenceAreaIdentity]PersistenceSnapshotConfig
}

// NewPersistenceFilesystemAdapter builds an adapter rooted at the appliance data
// root. Every config area must be a canonical Modbus area key; an unknown area
// or a non-canonical key is rejected so a configured path can never escape the
// root through an authored name.
func NewPersistenceFilesystemAdapter(root string, configs []PersistenceSnapshotConfig) (*PersistenceFilesystemAdapter, error) {
	if root == "" {
		return nil, fmt.Errorf("persistence filesystem adapter requires an appliance data root")
	}
	areas := make(map[persistenceAreaIdentity]PersistenceSnapshotConfig, len(configs))
	for _, cfg := range configs {
		if !PersistenceCanonicalArea(cfg.Area) {
			return nil, fmt.Errorf("persistence area %q is not a canonical modbus area", cfg.Area)
		}
		kind, _ := persistenceAreaKind(cfg.Area)
		if kind != cfg.Kind {
			return nil, fmt.Errorf("persistence area %q kind %d does not match canonical kind %d", cfg.Area, cfg.Kind, kind)
		}
		if cfg.Count == 0 {
			return nil, fmt.Errorf("persistence area %q has an empty range", cfg.Area)
		}
		id := persistenceAreaIdentity{Port: cfg.Key.Port, UnitID: cfg.Key.UnitID, Area: cfg.Area}
		if _, exists := areas[id]; exists {
			return nil, fmt.Errorf("persistence area %q for (%d,%d) is registered twice", cfg.Area, cfg.Key.Port, cfg.Key.UnitID)
		}
		areas[id] = cfg
	}
	return &PersistenceFilesystemAdapter{root: root, areas: areas}, nil
}

// PersistenceCanonicalArea reports whether area is one of the four canonical
// Modbus memory area keys. It is read-only.
func PersistenceCanonicalArea(area string) bool {
	_, ok := persistenceAreaKind(area)
	return ok
}

func (a *PersistenceFilesystemAdapter) areaDir(key PersistenceMemoryKey, area string) string {
	return filepath.Join(a.root, RelPersistenceDir,
		fmt.Sprintf("port-%d", key.Port),
		fmt.Sprintf("unit-%d", key.UnitID),
		"area-"+area,
	)
}

// RawSnapshotPath returns the deterministic path of the raw snapshot payload for
// one configured (Port, UnitID, Area). It is exported so callers and tests can
// assert the deterministic layout.
func (a *PersistenceFilesystemAdapter) RawSnapshotPath(key PersistenceMemoryKey, area string) string {
	return filepath.Join(a.areaDir(key, area), "snapshot.bin")
}

// ManifestPath returns the deterministic path of the compatibility manifest for
// one configured (Port, UnitID, Area).
func (a *PersistenceFilesystemAdapter) ManifestPath(key PersistenceMemoryKey, area string) string {
	return filepath.Join(a.areaDir(key, area), "manifest.json")
}

func (a *PersistenceFilesystemAdapter) config(key PersistenceMemoryKey, area string) (PersistenceSnapshotConfig, error) {
	cfg, ok := a.areas[persistenceAreaIdentity{Port: key.Port, UnitID: key.UnitID, Area: area}]
	if !ok {
		return PersistenceSnapshotConfig{}, fmt.Errorf("persistence area %q for (%d,%d) is not configured", area, key.Port, key.UnitID)
	}
	return cfg, nil
}

// WritePersistenceBytes implements PersistenceSnapshotStore: it splices the
// changed bytes into the area's raw snapshot image at the given offset, then
// atomically replaces the raw file and its manifest (payload length and
// SHA-256 recomputed over the whole image). The first write for an area may
// cover the whole image; later writes carry only changed runs. Malformed
// existing state (wrong size, unreadable) fails closed rather than being
// silently overwritten.
func (a *PersistenceFilesystemAdapter) WritePersistenceBytes(key PersistenceMemoryKey, area string, offset int, data []byte) error {
	cfg, err := a.config(key, area)
	if err != nil {
		return err
	}
	size, err := PersistenceSnapshotSize(cfg.Kind, cfg.Count)
	if err != nil {
		return err
	}
	if offset < 0 || len(data) == 0 || offset+len(data) > size {
		return fmt.Errorf("persistence write for area %q range [%d..%d) exceeds snapshot size %d", area, offset, offset+len(data), size)
	}
	image := make([]byte, size)
	rawPath := a.RawSnapshotPath(key, area)
	if existing, err := os.ReadFile(rawPath); err == nil {
		if len(existing) != size {
			return fmt.Errorf("persistence snapshot %s has size %d, expected %d", rawPath, len(existing), size)
		}
		copy(image, existing)
	} else if !os.IsNotExist(err) {
		return err
	}
	copy(image[offset:], data)

	manifest := NewPersistenceSnapshotManifest(key, area, cfg.Start, cfg.Count, image)
	manifestBytes, err := json.MarshalIndent(manifest, "", "  ")
	if err != nil {
		return err
	}
	dir := a.areaDir(key, area)
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return err
	}
	if err := writeFileAtomic(rawPath, image); err != nil {
		return err
	}
	return writeFileAtomic(a.ManifestPath(key, area), manifestBytes)
}

// ReadPersistenceSnapshot implements PersistenceSnapshotSource: it reads the raw
// payload and manifest for one configured area. present is false only when no
// snapshot exists for the area; a manifest that exists but is unreadable, or a
// missing raw file beside a manifest, surfaces an explicit error (incomplete
// content) rather than fabricated defaults. The existing manifest/length/
// checksum validation remains authoritative and is performed by the loader, not
// here.
func (a *PersistenceFilesystemAdapter) ReadPersistenceSnapshot(key PersistenceMemoryKey, area string) (PersistenceSnapshotManifest, []byte, bool, error) {
	if _, err := a.config(key, area); err != nil {
		return PersistenceSnapshotManifest{}, nil, false, err
	}
	manifestBytes, err := os.ReadFile(a.ManifestPath(key, area))
	if err != nil {
		if os.IsNotExist(err) {
			return PersistenceSnapshotManifest{}, nil, false, nil
		}
		return PersistenceSnapshotManifest{}, nil, false, err
	}
	var manifest PersistenceSnapshotManifest
	if err := json.Unmarshal(manifestBytes, &manifest); err != nil {
		return PersistenceSnapshotManifest{}, nil, false, fmt.Errorf("parse persistence manifest %s: %w", a.ManifestPath(key, area), err)
	}
	payload, err := os.ReadFile(a.RawSnapshotPath(key, area))
	if err != nil {
		return PersistenceSnapshotManifest{}, nil, false, fmt.Errorf("persistence manifest %s has no readable raw snapshot: %w", a.ManifestPath(key, area), err)
	}
	return manifest, payload, true, nil
}

// writeFileAtomic replaces path with data via a temp file and rename in the same
// directory, so a crash never leaves a partially written snapshot file.
func writeFileAtomic(path string, data []byte) error {
	tmp := path + ".tmp"
	if err := os.WriteFile(tmp, data, 0o644); err != nil {
		return err
	}
	if err := os.Rename(tmp, path); err != nil {
		_ = os.Remove(tmp)
		return err
	}
	return nil
}
