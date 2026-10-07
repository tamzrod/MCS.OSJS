package mma2composer

import (
	"strings"
	"testing"
)

func TestPersistenceSnapshotManifestAcceptsCompatiblePayload(t *testing.T) {
	key := PersistenceMemoryKey{Port: 5020, UnitID: 7}
	payload, err := EncodePersistenceRegisters([]uint16{0x1234, 0xabcd}, 2)
	if err != nil {
		t.Fatal(err)
	}
	manifest := NewPersistenceSnapshotManifest(key, "holding_registers", 100, 2, payload)

	if manifest.FormatVersion != PersistenceSnapshotFormatVersion ||
		manifest.Port != 5020 || manifest.UnitID != 7 ||
		manifest.Area != "holding_registers" ||
		manifest.Start != 100 || manifest.Count != 2 ||
		manifest.PayloadBytes != 4 || manifest.SHA256 == "" {
		t.Fatalf("manifest fields not bound correctly: %+v", manifest)
	}
	if err := ValidatePersistenceSnapshotManifest(manifest, key, "holding_registers", PersistenceRegisters, 100, 2, payload); err != nil {
		t.Fatalf("compatible snapshot rejected: %v", err)
	}
}

func TestPersistenceSnapshotManifestRejectsCorruptionAndIncompletePayload(t *testing.T) {
	key := PersistenceMemoryKey{Port: 502, UnitID: 1}
	payload, err := EncodePersistenceBits([]bool{true, false, true, false, true, false, true, false, true}, 9)
	if err != nil {
		t.Fatal(err)
	}
	manifest := NewPersistenceSnapshotManifest(key, "coils", 0, 9, payload)

	corrupt := append([]byte(nil), payload...)
	corrupt[0] ^= 0x01
	if err := ValidatePersistenceSnapshotManifest(manifest, key, "coils", PersistenceBits, 0, 9, corrupt); err == nil || !strings.Contains(err.Error(), "checksum") {
		t.Fatalf("expected checksum rejection, got %v", err)
	}

	short := payload[:1]
	if err := ValidatePersistenceSnapshotManifest(manifest, key, "coils", PersistenceBits, 0, 9, short); err == nil || !strings.Contains(err.Error(), "length") {
		t.Fatalf("expected incomplete payload rejection, got %v", err)
	}
}

func TestPersistenceSnapshotManifestRejectsIncompatibleLayoutAndIdentity(t *testing.T) {
	key := PersistenceMemoryKey{Port: 1502, UnitID: 3}
	payload, err := EncodePersistenceRegisters([]uint16{1, 2, 3}, 3)
	if err != nil {
		t.Fatal(err)
	}
	manifest := NewPersistenceSnapshotManifest(key, "input_registers", 10, 3, payload)

	cases := []struct {
		name  string
		m     PersistenceSnapshotManifest
		key   PersistenceMemoryKey
		area  string
		start uint16
		count uint16
	}{
		{"version", func() PersistenceSnapshotManifest { x := manifest; x.FormatVersion++; return x }(), key, "input_registers", 10, 3},
		{"port", manifest, PersistenceMemoryKey{Port: 2502, UnitID: 3}, "input_registers", 10, 3},
		{"unit", manifest, PersistenceMemoryKey{Port: 1502, UnitID: 4}, "input_registers", 10, 3},
		{"area", manifest, key, "holding_registers", 10, 3},
		{"start", manifest, key, "input_registers", 11, 3},
		{"count", manifest, key, "input_registers", 10, 4},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if err := ValidatePersistenceSnapshotManifest(tc.m, tc.key, tc.area, PersistenceRegisters, tc.start, tc.count, payload); err == nil {
				t.Fatalf("expected incompatible %s to be rejected", tc.name)
			}
		})
	}
}

func TestPersistenceSnapshotManifestRejectsManifestSizeMismatch(t *testing.T) {
	key := PersistenceMemoryKey{Port: 502, UnitID: 1}
	payload, err := EncodePersistenceRegisters([]uint16{1, 2}, 2)
	if err != nil {
		t.Fatal(err)
	}
	manifest := NewPersistenceSnapshotManifest(key, "holding_registers", 0, 2, payload)
	manifest.PayloadBytes = 6

	if err := ValidatePersistenceSnapshotManifest(manifest, key, "holding_registers", PersistenceRegisters, 0, 2, payload); err == nil || !strings.Contains(err.Error(), "configured size") {
		t.Fatalf("expected manifest size rejection, got %v", err)
	}
}
