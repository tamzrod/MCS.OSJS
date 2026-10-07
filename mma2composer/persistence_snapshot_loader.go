package mma2composer

import "fmt"

// PersistenceSnapshotSource reads the durable snapshot for one persisted memory
// area: the compatibility manifest (PERSIST-014) and the raw payload
// (PERSIST-013). present is false when no snapshot exists for the area. The
// source is read-only; startup loading never writes, restores or unseals.
type PersistenceSnapshotSource interface {
	ReadPersistenceSnapshot(key PersistenceMemoryKey, area string) (PersistenceSnapshotManifest, []byte, bool, error)
}

// PersistenceSnapshotArea is one configured persisted area to load at startup.
// Its identity/layout is derived from the authoritative memory area, never
// authored independently.
type PersistenceSnapshotArea struct {
	Area  string
	Kind  PersistenceAreaKind
	Start uint16
	Count uint16
}

// PersistenceRestoreOutcome is the explicit, deterministic startup state of one
// area or of a whole memory. Only PersistenceRestoreReady means a complete,
// integrity-verified snapshot set is available for the later restore stage.
type PersistenceRestoreOutcome int

const (
	// PersistenceRestoreReady: every configured area has a present, compatible,
	// integrity-verified snapshot. Nothing is written or unsealed.
	PersistenceRestoreReady PersistenceRestoreOutcome = iota
	// PersistenceRestoreDisabled: persistence is not enabled, so the loader does
	// not act.
	PersistenceRestoreDisabled
	// PersistenceRestoreUnsealed: persistence is enabled but the memory is not
	// sealed, so restore must not proceed.
	PersistenceRestoreUnsealed
	// PersistenceRestoreEmpty: persistence is enabled and sealed but no area is
	// configured, so there is nothing to restore.
	PersistenceRestoreEmpty
	// PersistenceRestoreMissing: no snapshot is present for an area.
	PersistenceRestoreMissing
	// PersistenceRestoreInvalid: a snapshot exists but is corrupt, incomplete or
	// unreadable.
	PersistenceRestoreInvalid
	// PersistenceRestoreIncompatible: a snapshot does not match the configured
	// format version, identity, area or layout.
	PersistenceRestoreIncompatible
)

func (o PersistenceRestoreOutcome) String() string {
	switch o {
	case PersistenceRestoreReady:
		return "ready"
	case PersistenceRestoreDisabled:
		return "disabled"
	case PersistenceRestoreUnsealed:
		return "unsealed"
	case PersistenceRestoreEmpty:
		return "empty"
	case PersistenceRestoreMissing:
		return "missing"
	case PersistenceRestoreInvalid:
		return "invalid"
	case PersistenceRestoreIncompatible:
		return "incompatible"
	default:
		return "unknown"
	}
}

// PersistenceAreaRestore is the startup result for one configured area. Payload
// carries the validated bytes only when the outcome is Ready.
type PersistenceAreaRestore struct {
	Area    string
	Kind    PersistenceAreaKind
	Start   uint16
	Count   uint16
	Outcome PersistenceRestoreOutcome
	Payload []byte
	Detail  string
}

// PersistenceRestorePlan is the complete startup restore plan for one memory. It
// is a plan only: the loader performs no Raw Ingest write and no unseal, so a
// partially valid snapshot set can never expose state through Modbus. State is
// Ready only when every configured area is Ready.
type PersistenceRestorePlan struct {
	Key     PersistenceMemoryKey
	Enabled bool
	Sealed  bool
	State   PersistenceRestoreOutcome
	Areas   []PersistenceAreaRestore
}

// PersistenceSnapshotAreas derives the startup load set from the
// persistence-owned RBE rules (PERSIST-003/009) for one memory.
func PersistenceSnapshotAreas(rules []PersistenceRBERule) ([]PersistenceSnapshotArea, error) {
	out := make([]PersistenceSnapshotArea, 0, len(rules))
	seen := make(map[string]bool, len(rules))
	for _, rule := range rules {
		kind, ok := persistenceAreaKind(rule.Area)
		if !ok {
			return nil, fmt.Errorf("persistence RBE rule %d has unknown area %q", rule.ID, rule.Area)
		}
		if seen[rule.Area] {
			return nil, fmt.Errorf("persistence RBE area %q is duplicated", rule.Area)
		}
		seen[rule.Area] = true
		out = append(out, PersistenceSnapshotArea{Area: rule.Area, Kind: kind, Start: rule.Start, Count: rule.Count})
	}
	return out, nil
}

// LoadPersistenceSnapshots loads and validates every configured persisted area
// during startup. It acts only when persistence is enabled and the memory is
// still sealed; otherwise it returns the matching explicit state without reading
// any snapshot. Each area is checked for compatibility (format version,
// identity, area, layout) and integrity (length, checksum) before it is marked
// Ready. The loader writes nothing and never unseals, so the aggregate State is
// Ready only when every configured area is Ready.
func LoadPersistenceSnapshots(key PersistenceMemoryKey, persistenceEnabled, sealed bool, areas []PersistenceSnapshotArea, source PersistenceSnapshotSource) (PersistenceRestorePlan, error) {
	plan := PersistenceRestorePlan{Key: key, Enabled: persistenceEnabled, Sealed: sealed}
	if !persistenceEnabled {
		plan.State = PersistenceRestoreDisabled
		return plan, nil
	}
	if !sealed {
		plan.State = PersistenceRestoreUnsealed
		return plan, nil
	}
	if source == nil {
		return PersistenceRestorePlan{}, fmt.Errorf("persistence snapshot loader requires a source")
	}
	if len(areas) == 0 {
		plan.State = PersistenceRestoreEmpty
		return plan, nil
	}
	plan.State = PersistenceRestoreReady
	plan.Areas = make([]PersistenceAreaRestore, 0, len(areas))
	for _, area := range areas {
		result := loadPersistenceArea(key, area, source)
		plan.Areas = append(plan.Areas, result)
		if result.Outcome != PersistenceRestoreReady && outcomePriority(result.Outcome) > outcomePriority(plan.State) {
			plan.State = result.Outcome
		}
	}
	return plan, nil
}

func loadPersistenceArea(key PersistenceMemoryKey, area PersistenceSnapshotArea, source PersistenceSnapshotSource) PersistenceAreaRestore {
	result := PersistenceAreaRestore{Area: area.Area, Kind: area.Kind, Start: area.Start, Count: area.Count}
	if area.Count == 0 {
		result.Outcome = PersistenceRestoreInvalid
		result.Detail = "configured persistence area has an empty range"
		return result
	}
	manifest, payload, present, err := source.ReadPersistenceSnapshot(key, area.Area)
	if err != nil {
		result.Outcome = PersistenceRestoreInvalid
		result.Detail = err.Error()
		return result
	}
	if !present {
		result.Outcome = PersistenceRestoreMissing
		result.Detail = "no snapshot is present"
		return result
	}
	if detail := snapshotIncompatibility(manifest, key, area); detail != "" {
		result.Outcome = PersistenceRestoreIncompatible
		result.Detail = detail
		return result
	}
	if err := ValidatePersistenceSnapshotManifest(manifest, key, area.Area, area.Kind, area.Start, area.Count, payload); err != nil {
		result.Outcome = PersistenceRestoreInvalid
		result.Detail = err.Error()
		return result
	}
	result.Outcome = PersistenceRestoreReady
	result.Payload = append([]byte(nil), payload...)
	return result
}

// snapshotIncompatibility reports a non-empty detail when the manifest's declared
// format version or identity/layout does not match the configured area.
func snapshotIncompatibility(manifest PersistenceSnapshotManifest, key PersistenceMemoryKey, area PersistenceSnapshotArea) string {
	switch {
	case manifest.FormatVersion != PersistenceSnapshotFormatVersion:
		return fmt.Sprintf("snapshot format version %d is incompatible with supported version %d", manifest.FormatVersion, PersistenceSnapshotFormatVersion)
	case manifest.Port != key.Port || manifest.UnitID != key.UnitID:
		return fmt.Sprintf("snapshot memory identity (%d,%d) does not match configured (%d,%d)", manifest.Port, manifest.UnitID, key.Port, key.UnitID)
	case manifest.Area != area.Area:
		return fmt.Sprintf("snapshot area %q does not match configured %q", manifest.Area, area.Area)
	case manifest.Start != area.Start || manifest.Count != area.Count:
		return fmt.Sprintf("snapshot layout start/count %d/%d does not match configured %d/%d", manifest.Start, manifest.Count, area.Start, area.Count)
	}
	return ""
}

// outcomePriority orders non-Ready outcomes by severity so a mixed snapshot set
// reports the most fundamental problem deterministically.
func outcomePriority(o PersistenceRestoreOutcome) int {
	switch o {
	case PersistenceRestoreIncompatible:
		return 3
	case PersistenceRestoreInvalid:
		return 2
	case PersistenceRestoreMissing:
		return 1
	default:
		return 0
	}
}
