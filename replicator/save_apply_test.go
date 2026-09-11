package replicator

import (
	"errors"
	"reflect"
	"strings"
	"testing"

	"github.com/tamzrod/MCS.OSJS/mma2composer"
)

func TestSaveApplyRejectsForeignReservationBeforeMutation(t *testing.T) {
	root := t.TempDir()
	store := Store{Root: root}
	previousDevice := validDeviceDefinition("PLC-1")
	previousDevice.Destination = DestinationSelection{Port: 5021, UnitID: 1}
	previous := Document{Devices: []DeviceDefinition{previousDevice}}
	if err := store.SaveDocument(previous); err != nil {
		t.Fatal(err)
	}

	composer := mma2composer.New(root, "simulator")
	owners := mma2composer.OwnershipDoc{Reservations: []mma2composer.OwnershipEntry{
		{Port: 5020, UnitID: 1, Owner: "simulator"},
		{Port: 5021, UnitID: 1, Owner: ProducerReplicator},
	}}
	if err := composer.SaveOwners(owners); err != nil {
		t.Fatal(err)
	}

	// Ownership collision is the exact (port, unit_id) pair. Reusing Unit ID 1
	// on another port is valid; requesting Simulator's exact 5020/1 is not.
	editedDevice := previousDevice
	editedDevice.Destination = DestinationSelection{Port: 5020, UnitID: 1}
	manager := NewRuntimeManager(store)
	_, _, err := manager.Apply(Document{Devices: []DeviceDefinition{editedDevice}})
	if !errors.Is(err, mma2composer.ErrReservationOwnedByOther) {
		t.Fatalf("Apply error = %v, want ownership conflict", err)
	}
	if !strings.Contains(err.Error(), "simulator") || !strings.Contains(err.Error(), "(5020,1)") {
		t.Fatalf("Apply error = %v, want owner and exact reservation", err)
	}

	persisted, loadErr := store.LoadDocument()
	if loadErr != nil {
		t.Fatal(loadErr)
	}
	if !reflect.DeepEqual(persisted, previous) {
		t.Fatalf("persisted document changed on conflict:\n got %#v\nwant %#v", persisted, previous)
	}
	after, loadOwnersErr := composer.LoadOwners()
	if loadOwnersErr != nil {
		t.Fatal(loadOwnersErr)
	}
	if !reflect.DeepEqual(after, owners) {
		t.Fatalf("owners changed on conflict:\n got %#v\nwant %#v", after, owners)
	}
}
