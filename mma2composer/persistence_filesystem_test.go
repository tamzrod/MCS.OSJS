package mma2composer

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func fsAdapter(t *testing.T, configs ...PersistenceSnapshotConfig) *PersistenceFilesystemAdapter {
	t.Helper()
	a, err := NewPersistenceFilesystemAdapter(t.TempDir(), configs)
	if err != nil {
		t.Fatal(err)
	}
	return a
}

func regConfig(key PersistenceMemoryKey, area string, start, count uint16) PersistenceSnapshotConfig {
	return PersistenceSnapshotConfig{Key: key, Area: area, Kind: PersistenceRegisters, Start: start, Count: count}
}

func bitConfig(key PersistenceMemoryKey, area string, start, count uint16) PersistenceSnapshotConfig {
	return PersistenceSnapshotConfig{Key: key, Area: area, Kind: PersistenceBits, Start: start, Count: count}
}

// PERSIST-R01 self-check: paths are deterministic and rooted beneath the
// appliance data root for a Port -> Unit ID -> Area identity.
func TestPersistenceFilesystemDeterministicPaths(t *testing.T) {
	key := PersistenceMemoryKey{Port: 5020, UnitID: 4}
	a := fsAdapter(t, regConfig(key, "holding_registers", 10, 2))
	rawPath := a.RawSnapshotPath(key, "holding_registers")
	manifestPath := a.ManifestPath(key, "holding_registers")
	if !strings.Contains(rawPath, filepath.Join("persistence", "snapshots", "port-5020", "unit-4", "area-holding_registers")) {
		t.Fatalf("raw path is not the deterministic layout: %s", rawPath)
	}
	if filepath.Dir(rawPath) != filepath.Dir(manifestPath) {
		t.Fatalf("raw and manifest must share the area directory: %s vs %s", rawPath, manifestPath)
	}
	if filepath.Base(rawPath) != "snapshot.bin" || filepath.Base(manifestPath) != "manifest.json" {
		t.Fatalf("unexpected file names: %s / %s", rawPath, manifestPath)
	}
}

// PERSIST-R01 self-check: a registered adapter rejects an unknown/non-canonical
// area, an empty range and a kind mismatch, so a path can never escape the root
// through an authored area name or a duplicated registration.
func TestPersistenceFilesystemRejectsInvalidConfig(t *testing.T) {
	key := PersistenceMemoryKey{Port: 1, UnitID: 1}
	if _, err := NewPersistenceFilesystemAdapter("", nil); err == nil {
		t.Fatal("blank root must be rejected")
	}
	if _, err := NewPersistenceFilesystemAdapter("", []PersistenceSnapshotConfig{regConfig(key, "holding_registers", 0, 1)}); err == nil {
		t.Fatal("blank root must be rejected even with configs")
	}
	if _, err := NewPersistenceFilesystemAdapter(t.TempDir(), []PersistenceSnapshotConfig{
		{Key: key, Area: "../../etc/passwd", Kind: PersistenceRegisters, Start: 0, Count: 1},
	}); err == nil {
		t.Fatal("non-canonical area must be rejected")
	}
	if _, err := NewPersistenceFilesystemAdapter(t.TempDir(), []PersistenceSnapshotConfig{regConfig(key, "coils", 0, 1)}); err == nil {
		t.Fatal("kind mismatch must be rejected")
	}
	if _, err := NewPersistenceFilesystemAdapter(t.TempDir(), []PersistenceSnapshotConfig{regConfig(key, "holding_registers", 0, 0)}); err == nil {
		t.Fatal("empty range must be rejected")
	}
	if _, err := NewPersistenceFilesystemAdapter(t.TempDir(), []PersistenceSnapshotConfig{
		regConfig(key, "holding_registers", 0, 1), regConfig(key, "holding_registers", 5, 1),
	}); err == nil {
		t.Fatal("duplicate registration must be rejected")
	}
}

// PERSIST-R01 self-check: write then read round-trips the raw payload and the
// manifest for one (Port, UnitID, Area), and the manifest/length/checksum
// validation accepts it.
func TestPersistenceFilesystemRoundTrip(t *testing.T) {
	key := PersistenceMemoryKey{Port: 5020, UnitID: 4}
	a := fsAdapter(t, regConfig(key, "holding_registers", 10, 2))
	payload, err := EncodePersistenceRegisters([]uint16{0x0102, 0x0304}, 2)
	if err != nil {
		t.Fatal(err)
	}
	if err := a.WritePersistenceBytes(key, "holding_registers", 0, payload); err != nil {
		t.Fatal(err)
	}
	manifest, got, present, err := a.ReadPersistenceSnapshot(key, "holding_registers")
	if err != nil {
		t.Fatal(err)
	}
	if !present || string(got) != string(payload) {
		t.Fatalf("round-trip payload wrong: present=%v got=%v", present, got)
	}
	if err := ValidatePersistenceSnapshotManifest(manifest, key, "holding_registers", PersistenceRegisters, 10, 2, got); err != nil {
		t.Fatalf("round-tripped manifest must validate: %v", err)
	}
}

// A bit area round-trips through the same adapter with LSB-first packing.
func TestPersistenceFilesystemBitAreaRoundTrip(t *testing.T) {
	key := PersistenceMemoryKey{Port: 1502, UnitID: 3}
	a := fsAdapter(t, bitConfig(key, "coils", 0, 9))
	payload, err := EncodePersistenceBits([]bool{true, false, true, false, true, false, true, false, true}, 9)
	if err != nil {
		t.Fatal(err)
	}
	if err := a.WritePersistenceBytes(key, "coils", 0, payload); err != nil {
		t.Fatal(err)
	}
	manifest, got, present, err := a.ReadPersistenceSnapshot(key, "coils")
	if err != nil || !present {
		t.Fatalf("coil read failed: present=%v err=%v", present, err)
	}
	if err := ValidatePersistenceSnapshotManifest(manifest, key, "coils", PersistenceBits, 0, 9, got); err != nil {
		t.Fatalf("coil manifest must validate: %v", err)
	}
}

// PERSIST-R01 self-check: a missing snapshot is reported as absent, not as an
// error or a fabricated default.
func TestPersistenceFilesystemMissingIsAbsent(t *testing.T) {
	key := PersistenceMemoryKey{Port: 1, UnitID: 1}
	a := fsAdapter(t, regConfig(key, "holding_registers", 0, 2))
	_, _, present, err := a.ReadPersistenceSnapshot(key, "holding_registers")
	if err != nil || present {
		t.Fatalf("missing snapshot must be absent with no error: present=%v err=%v", present, err)
	}
}

// A manifest present without its raw payload is an explicit incomplete-content
// error, never a fabricated empty snapshot.
func TestPersistenceFilesystemManifestWithoutPayloadFailsClosed(t *testing.T) {
	key := PersistenceMemoryKey{Port: 1, UnitID: 1}
	a := fsAdapter(t, regConfig(key, "holding_registers", 0, 2))
	dir := a.areaDir(key, "holding_registers")
	if err := os.MkdirAll(dir, 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(a.ManifestPath(key, "holding_registers"), []byte(`{"format_version":1}`), 0o644); err != nil {
		t.Fatal(err)
	}
	if _, _, _, err := a.ReadPersistenceSnapshot(key, "holding_registers"); err == nil {
		t.Fatal("a manifest without a raw payload must fail closed")
	}
}

// Corrupt manifest JSON surfaces an explicit parse error.
func TestPersistenceFilesystemCorruptManifestFailsClosed(t *testing.T) {
	key := PersistenceMemoryKey{Port: 1, UnitID: 1}
	a := fsAdapter(t, regConfig(key, "holding_registers", 0, 2))
	dir := a.areaDir(key, "holding_registers")
	if err := os.MkdirAll(dir, 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(a.ManifestPath(key, "holding_registers"), []byte("not json"), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(a.RawSnapshotPath(key, "holding_registers"), []byte{0, 1, 0, 2}, 0o644); err != nil {
		t.Fatal(err)
	}
	if _, _, _, err := a.ReadPersistenceSnapshot(key, "holding_registers"); err == nil {
		t.Fatal("a corrupt manifest must surface an explicit error")
	}
}

// A concrete appliance error: an unknown (unregistered) area is refused
// explicitly by both store and source paths.
func TestPersistenceFilesystemUnregisteredAreaRefused(t *testing.T) {
	key := PersistenceMemoryKey{Port: 1, UnitID: 1}
	a := fsAdapter(t, regConfig(key, "holding_registers", 0, 2))
	if err := a.WritePersistenceBytes(key, "input_registers", 0, []byte{0, 0}); err == nil {
		t.Fatal("writing an unregistered area must be refused")
	}
	if _, _, _, err := a.ReadPersistenceSnapshot(key, "input_registers"); err == nil {
		t.Fatal("reading an unregistered area must be refused")
	}
	// A different (Port, UnitID) is a different identity and is also unregistered.
	other := PersistenceMemoryKey{Port: 2, UnitID: 1}
	if _, _, _, err := a.ReadPersistenceSnapshot(other, "holding_registers"); err == nil {
		t.Fatal("a different memory identity must not resolve to a configured area")
	}
}

// PERSIST-R01 self-check: a partial (changed-run) write splices into the
// existing image and the manifest reflects the whole image, matching the
// PersistenceSnapshotWriter's change-only store contract.
func TestPersistenceFilesystemPartialWriteSplicesImage(t *testing.T) {
	key := PersistenceMemoryKey{Port: 1, UnitID: 1}
	a := fsAdapter(t, regConfig(key, "holding_registers", 0, 3))
	first, _ := EncodePersistenceRegisters([]uint16{0x0102, 0x0304, 0x0506}, 3)
	if err := a.WritePersistenceBytes(key, "holding_registers", 0, first); err != nil {
		t.Fatal(err)
	}
	// Change only the middle word (offset 2, two bytes), as the writer would.
	if err := a.WritePersistenceBytes(key, "holding_registers", 2, []byte{0x09, 0x09}); err != nil {
		t.Fatal(err)
	}
	manifest, got, _, err := a.ReadPersistenceSnapshot(key, "holding_registers")
	if err != nil {
		t.Fatal(err)
	}
	want := []byte{0x01, 0x02, 0x09, 0x09, 0x05, 0x06}
	if string(got) != string(want) {
		t.Fatalf("partial write did not splice image: got %v want %v", got, want)
	}
	if err := ValidatePersistenceSnapshotManifest(manifest, key, "holding_registers", PersistenceRegisters, 0, 3, got); err != nil {
		t.Fatalf("manifest must validate the spliced image: %v", err)
	}
}

// An out-of-range write is refused rather than silently truncating.
func TestPersistenceFilesystemOutOfRangeWriteRefused(t *testing.T) {
	key := PersistenceMemoryKey{Port: 1, UnitID: 1}
	a := fsAdapter(t, regConfig(key, "holding_registers", 0, 2))
	if err := a.WritePersistenceBytes(key, "holding_registers", 3, []byte{0, 0}); err == nil {
		t.Fatal("out-of-range write must be refused")
	}
	if err := a.WritePersistenceBytes(key, "holding_registers", 0, []byte{0, 0, 0, 0, 0}); err == nil {
		t.Fatal("oversized write must be refused")
	}
}

// A pre-existing raw file of the wrong size fails closed rather than being
// silently overwritten/repaired.
func TestPersistenceFilesystemMalformedExistingImageFailsClosed(t *testing.T) {
	key := PersistenceMemoryKey{Port: 1, UnitID: 1}
	a := fsAdapter(t, regConfig(key, "holding_registers", 0, 2))
	dir := a.areaDir(key, "holding_registers")
	if err := os.MkdirAll(dir, 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(a.RawSnapshotPath(key, "holding_registers"), []byte{0x01}, 0o644); err != nil {
		t.Fatal(err)
	}
	if err := a.WritePersistenceBytes(key, "holding_registers", 0, []byte{0, 0}); err == nil {
		t.Fatal("a wrong-size existing image must fail closed")
	}
}

// PERSIST-R01 self-check: safe replacement is atomic; after a write the tmp file
// is gone and only the final snapshot + manifest exist in the area directory.
func TestPersistenceFilesystemAtomicReplacement(t *testing.T) {
	key := PersistenceMemoryKey{Port: 1, UnitID: 1}
	a := fsAdapter(t, regConfig(key, "holding_registers", 0, 2))
	payload, _ := EncodePersistenceRegisters([]uint16{0xAAAA, 0xBBBB}, 2)
	if err := a.WritePersistenceBytes(key, "holding_registers", 0, payload); err != nil {
		t.Fatal(err)
	}
	// Re-writing replaces both files cleanly.
	second, _ := EncodePersistenceRegisters([]uint16{0x1111, 0x2222}, 2)
	if err := a.WritePersistenceBytes(key, "holding_registers", 0, second); err != nil {
		t.Fatal(err)
	}
	entries, err := os.ReadDir(a.areaDir(key, "holding_registers"))
	if err != nil {
		t.Fatal(err)
	}
	names := map[string]bool{}
	for _, e := range entries {
		names[e.Name()] = true
	}
	if names["snapshot.bin.tmp"] || names["manifest.json.tmp"] {
		t.Fatalf("atomic write left a temp file behind: %v", names)
	}
	if !names["snapshot.bin"] || !names["manifest.json"] {
		t.Fatalf("expected snapshot + manifest, got %v", names)
	}
	_, got, _, _ := a.ReadPersistenceSnapshot(key, "holding_registers")
	if string(got) != string(second) {
		t.Fatalf("replacement did not take effect: %v", got)
	}
}

// The adapter implements both persistence contracts, so the runtime can use it
// for save (store) and startup load (source).
func TestPersistenceFilesystemImplementsContracts(t *testing.T) {
	var _ PersistenceSnapshotStore = (*PersistenceFilesystemAdapter)(nil)
	var _ PersistenceSnapshotSource = (*PersistenceFilesystemAdapter)(nil)
}

// PERSIST-R01 self-check: the concrete adapter drives the existing
// PersistenceSnapshotWriter change-only path end to end — events read the
// authoritative area and persist changed bytes only, and a kept snapshot then
// loads back through the same adapter.
func TestPersistenceFilesystemWithSnapshotWriter(t *testing.T) {
	key := PersistenceMemoryKey{Port: 5020, UnitID: 1}
	adapter := fsAdapter(t, regConfig(key, "holding_registers", 0, 3))
	rules, err := PersistenceSnapshotRules(key, []PersistenceRBERule{
		{ID: 1, Area: "holding_registers", Start: 0, Count: 3, SystemOwned: true},
	})
	if err != nil {
		t.Fatal(err)
	}
	reader := &fakeAreaReader{data: map[string][]byte{"holding_registers": {0, 1, 0, 2, 0, 3}}}
	writer, err := NewPersistenceSnapshotWriter(rules, reader, adapter)
	if err != nil {
		t.Fatal(err)
	}
	// First event persists the whole area; so does the adapter's manifest.
	result, err := writer.OnPersistenceEvent(1)
	if err != nil {
		t.Fatal(err)
	}
	if result.StoreCalls != 1 || result.BytesWritten != 6 {
		t.Fatalf("initial save wrong: %+v", result)
	}
	// Unchanged state causes no store call (no rewrite).
	result, _ = writer.OnPersistenceEvent(1)
	if result.StoreCalls != 0 {
		t.Fatalf("unchanged state must not rewrite the snapshot: %+v", result)
	}
	// One word change writes only that word through the adapter.
	reader.data["holding_registers"] = []byte{0, 1, 0, 9, 0, 3}
	result, err = writer.OnPersistenceEvent(1)
	if err != nil {
		t.Fatal(err)
	}
	if result.StoreCalls != 1 || result.BytesWritten != 2 {
		t.Fatalf("changed word wrote %+v", result)
	}
	// The persisted snapshot loads back and validates through the adapter.
	manifest, payload, present, err := adapter.ReadPersistenceSnapshot(key, "holding_registers")
	if err != nil || !present {
		t.Fatalf("load after writer failed: present=%v err=%v", present, err)
	}
	if string(payload) != string([]byte{0, 1, 0, 9, 0, 3}) {
		t.Fatalf("adapter persisted wrong image: %v", payload)
	}
	if err := ValidatePersistenceSnapshotManifest(manifest, key, "holding_registers", PersistenceRegisters, 0, 3, payload); err != nil {
		t.Fatalf("persisted manifest must validate: %v", err)
	}
}
