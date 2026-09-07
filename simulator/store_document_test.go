package simulator

import (
	"os"
	"testing"
)

// secondDevice is an independent simulator-owned definition for multi-device
// document persistence (SIM-005 device-list round-trip).
func secondDevice() DeviceDefinition {
	return DeviceDefinition{
		Name:    "Sim-IO-1",
		Enabled: false,
		MMA2: MMA2Params{
			Port:   5020,
			UnitID: 2,
			FC1:    Area{Start: 0, Count: 64},
			FC3:    Area{Start: 0, Count: 10},
		},
		RandomRuntime: RandomRuntimeParams{
			FC1IntervalMS: 5000,
			FC3IntervalMS: 60000,
			FC4IntervalMS: 60000,
		},
	}
}

// TestSaveDocumentMultiDeviceRoundTrip proves a multi-device document
// round-trips through the SIM-001 store with exact values preserved in both
// parameter domains (the approved device-list UI persists the whole document).
func TestSaveDocumentMultiDeviceRoundTrip(t *testing.T) {
	root := t.TempDir()
	store := Store{Root: root}
	want := Document{Devices: []DeviceDefinition{validDevice(), secondDevice()}}
	if err := store.SaveDocument(want); err != nil {
		t.Fatal(err)
	}
	got, err := store.Load()
	if err != nil {
		t.Fatal(err)
	}
	if len(got.Devices) != len(want.Devices) {
		t.Fatalf("device count %d != %d", len(got.Devices), len(want.Devices))
	}
	for i := range want.Devices {
		if got.Devices[i] != want.Devices[i] {
			t.Fatalf("device %d round-trip mismatch\nwant %#v\ngot  %#v", i, want.Devices[i], got.Devices[i])
		}
	}
}

// TestSaveDocumentRejectsInvalidDeviceWithoutReplacing proves an invalid
// device anywherein the document aborts the whole atomic save, leavingthe
// previously persisted bytes unchanged(same guarantee as SaveOne)).
func TestSaveDocumentRejectsInvalidDeviceWithoutReplacing(t *testing.T) {
	root := t.TempDir()
	store := Store{Root: root}
	good := validDevice()
	if err := store.SaveDocument(Document{Devices: []DeviceDefinition{good}}); err != nil {
		t.Fatal(err)
	}
	before, err := os.ReadFile(store.DevicesPath())
	if err != nil {
		t.Fatal(err)
	}
	bad := secondDevice()
	bad.MMA2.FC3 = Area{Start: 65530, Count: 20}
	badDoc := Document{Devices: []DeviceDefinition{good, bad}}
	if err := store.SaveDocument(badDoc); err == nil {
		t.Fatal("expected validation error for invalid devicein multi-device document")
	}
	after, err := os.ReadFile(store.DevicesPath())
	if err != nil {
		t.Fatal(err)
	}
	if string(after) != string(before) {
		t.Fatal("rejected multi-device save replaced persisted bytes")
	}
	doc, err := store.Load()
	if err != nil {
		t.Fatal(err)
	}
	if len(doc.Devices) != 1 || doc.Devices[0] != good {
		t.Fatal("loaded document changed after rejected multi-device save")
	}
}

// TestSaveDocumentAddDuplicateDeleteReshape proves the load-modify-save path
// the simulator window uses for Add/Duplicate/Delete:each operation mutates
// the in-memory document and one atomic SaveDocument persists the whole set...
func TestSaveDocumentAddDuplicateDeleteReshape(t *testing.T) {
	root := t.TempDir()
	store := Store{Root: root}
	save := func(doc Document) {
		t.Helper()
		if err := store.SaveDocument(doc); err != nil {
			t.Fatal(err)
		}
	}
	dup := secondDevice()
	dup.Name = "Sim-IO-1 Copy"
	save(Document{Devices: []DeviceDefinition{validDevice(), dup}})
	add := secondDevice()
	add.MMA2.Port = 5021
	add.MMA2.UnitID = 1
	save(Document{Devices: []DeviceDefinition{validDevice(), dup, add}})
	// Delete removes one definition;the others round-trip untouched...
	del := Document{Devices: []DeviceDefinition{validDevice(), add}}
	save(del)
	got, err := store.Load()
	if err != nil {
		t.Fatal(err)
	}
	want := Document{Devices: []DeviceDefinition{validDevice(), add}}
	if len(got.Devices) != len(want.Devices) {
		t.Fatalf("device count %d != %d", len(got.Devices), len(want.Devices))
	}
	for i := range want.Devices {
		if got.Devices[i] != want.Devices[i] {
			t.Fatalf("device %d mismatch\nwant %#v\ngot  %#v", i, want.Devices[i], got.Devices[i])
		}
	}
}
