package mma2composer

import (
	"bytes"
	"reflect"
	"testing"
)

func TestPersistenceSnapshotBitsLSBFirst(t *testing.T) {
	values := []bool{true, false, true, true, false, false, false, true, true}
	raw, err := EncodePersistenceBits(values, uint16(len(values)))
	if err != nil {
		t.Fatal(err)
	}
	want := []byte{0x8d, 0x01}
	if !bytes.Equal(raw, want) {
		t.Fatalf("LSB-first encoding: want % x, got % x", want, raw)
	}
	decoded, err := DecodePersistenceBits(raw, uint16(len(values)))
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(decoded, values) {
		t.Fatalf("bit round-trip: want %+v, got %+v", values, decoded)
	}
	again, _ := EncodePersistenceBits(values, uint16(len(values)))
	if !bytes.Equal(again, raw) {
		t.Fatalf("bit encoding is not deterministic: % x vs % x", raw, again)
	}
}

func TestPersistenceSnapshotRegistersBigEndian(t *testing.T) {
	values := []uint16{0x1234, 0xabcd, 0x0001}
	raw, err := EncodePersistenceRegisters(values, uint16(len(values)))
	if err != nil {
		t.Fatal(err)
	}
	want := []byte{0x12, 0x34, 0xab, 0xcd, 0x00, 0x01}
	if !bytes.Equal(raw, want) {
		t.Fatalf("big-endian encoding: want % x, got % x", want, raw)
	}
	decoded, err := DecodePersistenceRegisters(raw, uint16(len(values)))
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(decoded, values) {
		t.Fatalf("register round-trip: want %+v, got %+v", values, decoded)
	}
	again, _ := EncodePersistenceRegisters(values, uint16(len(values)))
	if !bytes.Equal(again, raw) {
		t.Fatalf("register encoding is not deterministic: % x vs % x", raw, again)
	}
}

func TestPersistenceSnapshotConfiguredLength(t *testing.T) {
	cases := []struct {
		kind  PersistenceAreaKind
		count uint16
		size  int
	}{
		{PersistenceBits, 0, 0},
		{PersistenceBits, 1, 1},
		{PersistenceBits, 8, 1},
		{PersistenceBits, 9, 2},
		{PersistenceRegisters, 0, 0},
		{PersistenceRegisters, 1, 2},
		{PersistenceRegisters, 7, 14},
	}
	for _, tc := range cases {
		got, err := PersistenceSnapshotSize(tc.kind, tc.count)
		if err != nil {
			t.Fatal(err)
		}
		if got != tc.size {
			t.Fatalf("kind=%d count=%d: want size %d, got %d", tc.kind, tc.count, tc.size, got)
		}
	}

	if _, err := DecodePersistenceBits([]byte{0, 0}, 8); err == nil {
		t.Fatal("expected bit snapshot length mismatch")
	}
	if _, err := DecodePersistenceRegisters([]byte{0, 1, 2}, 2); err == nil {
		t.Fatal("expected register snapshot length mismatch")
	}
	if _, err := EncodePersistenceBits([]bool{true}, 2); err == nil {
		t.Fatal("expected bit value count mismatch")
	}
	if _, err := EncodePersistenceRegisters([]uint16{1}, 2); err == nil {
		t.Fatal("expected register value count mismatch")
	}
}
