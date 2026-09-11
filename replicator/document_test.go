package replicator

import (
	"errors"
	"net"
	"reflect"
	"strings"
	"testing"

	"github.com/tamzrod/MCS.OSJS/mma2composer"
)

func validDeviceDefinition(name string) DeviceDefinition {
	return DeviceDefinition{
		Name:       name,
		Enabled:    true,
		Endpoint:   "127.0.0.1:5020",
		UnitID:     1,
		Function:   3,
		Start:      10,
		Count:      4,
		ScanRateMS: 250,
		Destination: DestinationSelection{
			Port:       DefaultDestinationPort,
			UnitID:     DefaultDestinationUnit,
			AutoPort:   true,
			AutoUnitID: true,
		},
	}
}

func TestDocumentSaveLoadRoundTrip(t *testing.T) {
	store := Store{Root: t.TempDir()}
	doc := Document{Devices: []DeviceDefinition{validDeviceDefinition("PLC-1")}}
	if err := store.SaveDocument(doc); err != nil {
		t.Fatalf("SaveDocument: %v", err)
	}
	got, err := store.LoadDocument()
	if err != nil {
		t.Fatalf("LoadDocument: %v", err)
	}
	if !reflect.DeepEqual(got, doc) {
		t.Fatalf("round trip mismatch:\n got: %#v\nwant: %#v", got, doc)
	}
}

func TestSuggestDestinationSkipsOccupiedReservations(t *testing.T) {
	store := Store{Root: t.TempDir()}
	composer := mma2composer.New(store.Root, "simulator")
	owners := mma2composer.OwnershipDoc{Reservations: []mma2composer.OwnershipEntry{
		{Port: DefaultDestinationPort, UnitID: 1, Owner: "simulator"},
		{Port: DefaultDestinationPort, UnitID: 2, Owner: ProducerReplicator},
	}}
	if err := composer.SaveOwners(owners); err != nil {
		t.Fatal(err)
	}
	got, err := store.SuggestDestination()
	if err != nil {
		t.Fatal(err)
	}
	if got.Port != DefaultDestinationPort || got.UnitID != 3 || got.Status != "AVAILABLE" {
		t.Fatalf("suggestion = %#v, want port %d unit 3 AVAILABLE", got, DefaultDestinationPort)
	}
}

func TestResolveManualForeignCollisionReportsOwner(t *testing.T) {
	store := Store{Root: t.TempDir()}
	composer := mma2composer.New(store.Root, "simulator")
	owners := mma2composer.OwnershipDoc{Reservations: []mma2composer.OwnershipEntry{{
		Port: 5020, UnitID: 1, Owner: "simulator",
	}}}
	if err := composer.SaveOwners(owners); err != nil {
		t.Fatal(err)
	}
	device := validDeviceDefinition("PLC-1")
	device.Destination = DestinationSelection{Port: 5020, UnitID: 1}
	_, err := store.ResolveDocumentDestinations(Document{Devices: []DeviceDefinition{device}})
	if !errors.Is(err, mma2composer.ErrReservationOwnedByOther) {
		t.Fatalf("error = %v, want ownership collision", err)
	}
	if !strings.Contains(err.Error(), "simulator") {
		t.Fatalf("error = %v, want actual foreign owner", err)
	}
}

func TestComposeDocumentPreservesForeignAndAllReplicatorReservations(t *testing.T) {
	store := Store{Root: t.TempDir()}
	foreignComposer := mma2composer.New(store.Root, "simulator")
	foreignMemory := mma2composer.Memory{
		UnitID:      1,
		HoldingRegs: &mma2composer.Area{Start: 0, Count: 2},
		Policy: &mma2composer.Policy{Rules: []mma2composer.PolicyRule{{
			ID: "foreign", SourceIP: []string{"0.0.0.0/0"}, AllowFC: []uint8{3},
		}}},
	}
	foreignCfg := mma2composer.AddMemory(mma2composer.EffectiveConfig{}, "simulator-5020-1", net.JoinHostPort("0.0.0.0", "5020"), foreignMemory)
	foreignOwners := mma2composer.OwnershipDoc{Reservations: []mma2composer.OwnershipEntry{{Port: 5020, UnitID: 1, Owner: "simulator"}}}
	if err := foreignComposer.Commit(foreignCfg, foreignOwners); err != nil {
		t.Fatal(err)
	}

	first := validDeviceDefinition("PLC-1")
	first.Destination = DestinationSelection{Port: 5021, UnitID: 1}
	second := validDeviceDefinition("PLC-2")
	second.Endpoint = "127.0.0.1:5022"
	second.UnitID = 2
	second.Function = 4
	second.Destination = DestinationSelection{Port: 5021, UnitID: 2}

	resolved, _, err := store.ComposeDocumentDestinations(Document{Devices: []DeviceDefinition{first, second}})
	if err != nil {
		t.Fatalf("ComposeDocumentDestinations: %v", err)
	}
	if len(resolved.Devices) != 2 {
		t.Fatalf("resolved device count = %d", len(resolved.Devices))
	}
	owners, err := foreignComposer.LoadOwners()
	if err != nil {
		t.Fatal(err)
	}
	if len(owners.Reservations) != 3 {
		t.Fatalf("owners = %#v, want foreign + 2 Replicator reservations", owners.Reservations)
	}
	seen := make(map[string]string)
	for _, entry := range owners.Reservations {
		seen[net.JoinHostPort("", portUnit(entry.Port, entry.UnitID))] = entry.Owner
	}
	if seen[net.JoinHostPort("", portUnit(5020, 1))] != "simulator" {
		t.Fatalf("foreign ownership lost: %#v", owners.Reservations)
	}
	if seen[net.JoinHostPort("", portUnit(5021, 1))] != ProducerReplicator || seen[net.JoinHostPort("", portUnit(5021, 2))] != ProducerReplicator {
		t.Fatalf("Replicator ownership incomplete: %#v", owners.Reservations)
	}

	// Rebuild with PLC-2 only. The removed Replicator reservation must disappear
	// while the Simulator reservation remains untouched.
	if _, _, err := store.ComposeDocumentDestinations(Document{Devices: []DeviceDefinition{second}}); err != nil {
		t.Fatalf("recompose: %v", err)
	}
	after, err := foreignComposer.LoadOwners()
	if err != nil {
		t.Fatal(err)
	}
	if len(after.Reservations) != 2 {
		t.Fatalf("owners after delete = %#v", after.Reservations)
	}
	for _, entry := range after.Reservations {
		if entry.Port == 5020 && entry.UnitID == 1 && entry.Owner == "simulator" {
			return
		}
	}
	t.Fatalf("Simulator reservation was removed: %#v", after.Reservations)
}

func TestDeviceRuntimeConfigMapsOneToOneRange(t *testing.T) {
	device := validDeviceDefinition("PLC-1")
	device.Endpoint = "10.20.30.40:1502"
	device.UnitID = 7
	device.Function = 4
	device.Start = 123
	device.Count = 9
	device.ScanRateMS = 750
	device.Destination = DestinationSelection{Port: 2502, UnitID: 8}
	cfg, err := device.runtimeConfig()
	if err != nil {
		t.Fatal(err)
	}
	if cfg.Source.Host != "10.20.30.40" || cfg.Source.Port != 1502 || cfg.Source.UnitID != 7 || cfg.Source.Function != 4 {
		t.Fatalf("source mapping = %#v", cfg.Source)
	}
	if cfg.Destination.ListenerPort != 2502 || cfg.Destination.UnitID != 8 || cfg.Destination.Area != "fc4" || cfg.Destination.Start != 123 || cfg.Destination.Count != 9 {
		t.Fatalf("destination mapping = %#v", cfg.Destination)
	}
	if cfg.Source.PollIntervalMS != 750 {
		t.Fatalf("poll interval = %d", cfg.Source.PollIntervalMS)
	}
}

func portUnit(port, unit uint16) string {
	return string(rune(port)) + "/" + string(rune(unit))
}
