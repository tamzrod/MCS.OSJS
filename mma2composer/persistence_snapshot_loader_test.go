package mma2composer

import (
	"errors"
	"testing"
)

type fakeSnapshotSource struct {
	snapshots map[string]PersistenceSnapshotManifest
	payloads  map[string][]byte
	err       error
	reads     int
}

func (f *fakeSnapshotSource) ReadPersistenceSnapshot(key PersistenceMemoryKey, area string) (PersistenceSnapshotManifest, []byte, bool, error) {
	f.reads++
	if f.err != nil {
		return PersistenceSnapshotManifest{}, nil, false, f.err
	}
	manifest, ok := f.snapshots[area]
	if !ok {
		return PersistenceSnapshotManifest{}, nil, false, nil
	}
	return manifest, append([]byte(nil), f.payloads[area]...), true, nil
}

func loaderArea(name string, kind PersistenceAreaKind, start, count uint16) PersistenceSnapshotArea {
	return PersistenceSnapshotArea{Area: name, Kind: kind, Start: start, Count: count}
}

func loaderFixture(t *testing.T) (PersistenceMemoryKey, []PersistenceSnapshotArea, *fakeSnapshotSource) {
	t.Helper()
	key := PersistenceMemoryKey{Port: 5020, UnitID: 4}
	regs, err := EncodePersistenceRegisters([]uint16{0x0102, 0x0304}, 2)
	if err != nil {
		t.Fatal(err)
	}
	bits, err := EncodePersistenceBits([]bool{true, false, true, false, true, false, true, false, true}, 9)
	if err != nil {
		t.Fatal(err)
	}
	source := &fakeSnapshotSource{
		snapshots: map[string]PersistenceSnapshotManifest{
			"holding_registers": NewPersistenceSnapshotManifest(key, "holding_registers", 10, 2, regs),
			"coils":             NewPersistenceSnapshotManifest(key, "coils", 0, 9, bits),
		},
		payloads: map[string][]byte{"holding_registers": regs, "coils": bits},
	}
	areas := []PersistenceSnapshotArea{
		loaderArea("holding_registers", PersistenceRegisters, 10, 2),
		loaderArea("coils", PersistenceBits, 0, 9),
	}
	return key, areas, source
}

func areaOutcome(plan PersistenceRestorePlan, area string) PersistenceAreaRestore {
	for _, a := range plan.Areas {
		if a.Area == area {
			return a
		}
	}
	return PersistenceAreaRestore{}
}

// PERSIST-015 self-check: a complete, compatible, integrity-verified snapshot
// set loads into an explicit Ready plan without writing or unsealing.
func TestPersistenceSnapshotLoaderReadyForCompleteSet(t *testing.T) {
	key, areas, source := loaderFixture(t)

	plan, err := LoadPersistenceSnapshots(key, true, true, areas, source)
	if err != nil {
		t.Fatal(err)
	}
	if plan.State != PersistenceRestoreReady {
		t.Fatalf("complete snapshot set must be Ready: %s", plan.State)
	}
	if !plan.Enabled || !plan.Sealed {
		t.Fatalf("plan must retain enablement/sealed facts: %+v", plan)
	}
	if len(plan.Areas) != 2 || source.reads != 2 {
		t.Fatalf("expected one read per configured area: areas=%d reads=%d", len(plan.Areas), source.reads)
	}
	regs := areaOutcome(plan, "holding_registers")
	if regs.Outcome != PersistenceRestoreReady || len(regs.Payload) != 4 {
		t.Fatalf("register area not ready: %+v", regs)
	}
	decoded, err := DecodePersistenceRegisters(regs.Payload, 2)
	if err != nil || decoded[0] != 0x0102 || decoded[1] != 0x0304 {
		t.Fatalf("register payload not preserved: %+v err=%v", decoded, err)
	}
	coils := areaOutcome(plan, "coils")
	if coils.Outcome != PersistenceRestoreReady || len(coils.Payload) != 2 {
		t.Fatalf("bit area not ready: %+v", coils)
	}
}

// The loader acts only for persistence-enabled sealed memories.
func TestPersistenceSnapshotLoaderActsOnlyForEnabledSealed(t *testing.T) {
	key, areas, source := loaderFixture(t)

	disabled, err := LoadPersistenceSnapshots(key, false, true, areas, source)
	if err != nil {
		t.Fatal(err)
	}
	if disabled.State != PersistenceRestoreDisabled || len(disabled.Areas) != 0 {
		t.Fatalf("disabled persistence must not load: %+v", disabled)
	}
	if source.reads != 0 {
		t.Fatalf("disabled persistence must not read snapshots: %d", source.reads)
	}

	unsealed, err := LoadPersistenceSnapshots(key, true, false, areas, source)
	if err != nil {
		t.Fatal(err)
	}
	if unsealed.State != PersistenceRestoreUnsealed || len(unsealed.Areas) != 0 {
		t.Fatalf("unsealed memory must not load: %+v", unsealed)
	}
	if source.reads != 0 {
		t.Fatalf("unsealed memory must not read snapshots: %d", source.reads)
	}

	empty, err := LoadPersistenceSnapshots(key, true, true, nil, source)
	if err != nil {
		t.Fatal(err)
	}
	if empty.State != PersistenceRestoreEmpty {
		t.Fatalf("no configured areas must be Empty: %s", empty.State)
	}
}

// Missing, invalid and incompatible snapshots each produce a distinct explicit
// restore state and never a Ready plan.
func TestPersistenceSnapshotLoaderExplicitFailureStates(t *testing.T) {
	key, areas, _ := loaderFixture(t)

	t.Run("missing", func(t *testing.T) {
		_, _, source := loaderFixture(t)
		delete(source.snapshots, "coils")
		plan, err := LoadPersistenceSnapshots(key, true, true, areas, source)
		if err != nil {
			t.Fatal(err)
		}
		if plan.State != PersistenceRestoreMissing {
			t.Fatalf("missing snapshot must be Missing: %s", plan.State)
		}
		if got := areaOutcome(plan, "coils"); got.Outcome != PersistenceRestoreMissing {
			t.Fatalf("missing area state wrong: %+v", got)
		}
		if got := areaOutcome(plan, "holding_registers"); got.Outcome != PersistenceRestoreReady {
			t.Fatalf("unaffected area must stay Ready: %+v", got)
		}
	})

	t.Run("corrupt-payload", func(t *testing.T) {
		_, _, source := loaderFixture(t)
		source.payloads["coils"][0] ^= 0x01
		plan, err := LoadPersistenceSnapshots(key, true, true, areas, source)
		if err != nil {
			t.Fatal(err)
		}
		if plan.State != PersistenceRestoreInvalid {
			t.Fatalf("corrupt payload must be Invalid: %s", plan.State)
		}
		if got := areaOutcome(plan, "coils"); got.Outcome != PersistenceRestoreInvalid || len(got.Payload) != 0 {
			t.Fatalf("corrupt area must be Invalid with no payload: %+v", got)
		}
	})

	t.Run("truncated-payload", func(t *testing.T) {
		_, _, source := loaderFixture(t)
		source.payloads["coils"] = source.payloads["coils"][:1]
		plan, err := LoadPersistenceSnapshots(key, true, true, areas, source)
		if err != nil {
			t.Fatal(err)
		}
		if plan.State != PersistenceRestoreInvalid {
			t.Fatalf("incomplete payload must be Invalid: %s", plan.State)
		}
	})

	t.Run("incompatible-version", func(t *testing.T) {
		_, _, source := loaderFixture(t)
		manifest := source.snapshots["coils"]
		manifest.FormatVersion++
		source.snapshots["coils"] = manifest
		plan, err := LoadPersistenceSnapshots(key, true, true, areas, source)
		if err != nil {
			t.Fatal(err)
		}
		if plan.State != PersistenceRestoreIncompatible {
			t.Fatalf("version mismatch must be Incompatible: %s", plan.State)
		}
	})

	t.Run("incompatible-layout", func(t *testing.T) {
		_, _, source := loaderFixture(t)
		manifest := source.snapshots["coils"]
		manifest.Count = 10
		source.snapshots["coils"] = manifest
		plan, err := LoadPersistenceSnapshots(key, true, true, areas, source)
		if err != nil {
			t.Fatal(err)
		}
		if plan.State != PersistenceRestoreIncompatible {
			t.Fatalf("layout mismatch must be Incompatible: %s", plan.State)
		}
	})

	t.Run("incompatible-identity", func(t *testing.T) {
		_, _, source := loaderFixture(t)
		manifest := source.snapshots["coils"]
		manifest.UnitID = 5
		source.snapshots["coils"] = manifest
		plan, err := LoadPersistenceSnapshots(key, true, true, areas, source)
		if err != nil {
			t.Fatal(err)
		}
		if plan.State != PersistenceRestoreIncompatible {
			t.Fatalf("identity mismatch must be Incompatible: %s", plan.State)
		}
	})

	t.Run("read-error", func(t *testing.T) {
		_, _, source := loaderFixture(t)
		source.err = errors.New("snapshot store unavailable")
		plan, err := LoadPersistenceSnapshots(key, true, true, areas, source)
		if err != nil {
			t.Fatal(err)
		}
		if plan.State != PersistenceRestoreInvalid {
			t.Fatalf("unreadable snapshot must be Invalid: %s", plan.State)
		}
	})
}

// A mixed set reports the most fundamental problem and is never Ready, so a
// partially loaded set can never be exposed as restorable.
func TestPersistenceSnapshotLoaderNeverReadyWithAnyFailure(t *testing.T) {
	key, areas, _ := loaderFixture(t)

	_, _, source := loaderFixture(t)
	manifest := source.snapshots["coils"]
	manifest.FormatVersion++
	source.snapshots["coils"] = manifest
	delete(source.snapshots, "holding_registers")
	plan, err := LoadPersistenceSnapshots(key, true, true, areas, source)
	if err != nil {
		t.Fatal(err)
	}
	if plan.State == PersistenceRestoreReady {
		t.Fatal("mixed snapshot set must never be Ready")
	}
	if plan.State != PersistenceRestoreIncompatible {
		t.Fatalf("most fundamental problem must win: %s", plan.State)
	}
	for _, area := range plan.Areas {
		if area.Outcome == PersistenceRestoreReady {
			continue
		}
		if len(area.Payload) != 0 {
			t.Fatalf("non-ready area must carry no payload: %+v", area)
		}
	}
}

// The startup load set is derived from the persistence-owned RBE rules.
func TestPersistenceSnapshotAreasDerivedFromRules(t *testing.T) {
	areas, err := PersistenceSnapshotAreas([]PersistenceRBERule{
		{ID: 1, Area: "coils", Start: 0, Count: 8, SystemOwned: true},
		{ID: 2, Area: "holding_registers", Start: 5, Count: 3, SystemOwned: true},
	})
	if err != nil {
		t.Fatal(err)
	}
	if len(areas) != 2 ||
		areas[0].Area != "coils" || areas[0].Kind != PersistenceBits || areas[0].Count != 8 ||
		areas[1].Area != "holding_registers" || areas[1].Kind != PersistenceRegisters || areas[1].Start != 5 {
		t.Fatalf("derived areas wrong: %+v", areas)
	}

	if _, err := PersistenceSnapshotAreas([]PersistenceRBERule{{ID: 1, Area: "nonsense", Count: 1}}); err == nil {
		t.Fatal("unknown area must be rejected")
	}
	if _, err := PersistenceSnapshotAreas([]PersistenceRBERule{
		{ID: 1, Area: "coils", Count: 1},
		{ID: 2, Area: "coils", Count: 1},
	}); err == nil {
		t.Fatal("duplicate area must be rejected")
	}
}

func TestPersistenceSnapshotLoaderRequiresSource(t *testing.T) {
	key, areas, _ := loaderFixture(t)
	if _, err := LoadPersistenceSnapshots(key, true, true, areas, nil); err == nil {
		t.Fatal("nil source must be rejected for an enabled sealed memory")
	}
}
