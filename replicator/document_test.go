package replicator

import (
	"errors"
	"fmt"
	"net"
	"os"
	"reflect"
	"strings"
	"testing"

	"github.com/tamzrod/MCS.OSJS/mma2composer"
)

func validDeviceDefinition(name string) DeviceDefinition {
	return DeviceDefinition{
		Name:     name,
		Enabled:  true,
		Endpoint: "127.0.0.1:5020",
		UnitID:   1,
		PullBlock: PullBlock{
			Function:   3,
			Start:      10,
			Count:      4,
			ScanRateMS: 250,
		},
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

func TestLoadDocumentMigratesLegacyRangeIntoPullBlock(t *testing.T) {
	store := Store{Root: t.TempDir()}
	if err := os.MkdirAll(store.ReplicatorDir(), 0o755); err != nil {
		t.Fatal(err)
	}
	legacy := []byte(`devices:
  - name: PLC-old
    enabled: true
    endpoint: 127.0.0.1:5020
    unit_id: 7
    function: 4
    start: 123
    count: 9
    scan_rate_ms: 750
    destination:
      port: 5021
      unit_id: 2
      auto_port: false
      auto_unit_id: false
`)
	if err := os.WriteFile(store.DocumentPath(), legacy, 0o644); err != nil {
		t.Fatal(err)
	}
	doc, err := store.LoadDocument()
	if err != nil {
		t.Fatal(err)
	}
	if len(doc.Devices) != 1 {
		t.Fatalf("devices = %d", len(doc.Devices))
	}
	block := doc.Devices[0].PullBlock
	if block.Function != 4 || block.Start != 123 || block.Count != 9 || block.ScanRateMS != 750 {
		t.Fatalf("migrated pull block = %#v", block)
	}
}

func TestSuggestDestinationSkipsOccupiedPortsAndUnits(t *testing.T) {
	store := Store{Root: t.TempDir()}
	composer := mma2composer.New(store.Root, "simulator")
	owners := mma2composer.OwnershipDoc{Reservations: []mma2composer.OwnershipEntry{
		{Port: DefaultDestinationPort, UnitID: 1, Owner: "simulator"},
		{Port: DefaultDestinationPort + 1, UnitID: 2, Owner: ProducerReplicator},
	}}
	if err := composer.SaveOwners(owners); err != nil {
		t.Fatal(err)
	}
	got, err := store.SuggestDestination()
	if err != nil {
		t.Fatal(err)
	}
	if got.Port != DefaultDestinationPort+2 || got.UnitID != 3 || got.Status != "AVAILABLE" {
		t.Fatalf("suggestion = %#v, want port %d unit 3 AVAILABLE", got, DefaultDestinationPort+2)
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
	if !strings.Contains(err.Error(), "simulator") || !strings.Contains(err.Error(), "5020") {
		t.Fatalf("error = %v, want actual foreign owner and reservation", err)
	}
}

func TestResolveRejectsForeignPortWithDifferentUnit(t *testing.T) {
	store := Store{Root: t.TempDir()}
	composer := mma2composer.New(store.Root, "simulator")
	if err := composer.SaveOwners(mma2composer.OwnershipDoc{Reservations: []mma2composer.OwnershipEntry{{Port: 5020, UnitID: 1, Owner: "simulator"}}}); err != nil {
		t.Fatal(err)
	}
	device := validDeviceDefinition("PLC-port-conflict")
	device.Destination = DestinationSelection{Port: 5020, UnitID: 2}
	_, err := store.ResolveDocumentDestinations(Document{Devices: []DeviceDefinition{device}})
	if !errors.Is(err, mma2composer.ErrReservationOwnedByOther) || !strings.Contains(err.Error(), "port 5020") || !strings.Contains(err.Error(), "simulator") {
		t.Fatalf("error = %v, want foreign port ownership collision", err)
	}
}

func TestResolveRejectsForeignUnitWithDifferentPort(t *testing.T) {
	store := Store{Root: t.TempDir()}
	composer := mma2composer.New(store.Root, "simulator")
	if err := composer.SaveOwners(mma2composer.OwnershipDoc{Reservations: []mma2composer.OwnershipEntry{{Port: 5020, UnitID: 1, Owner: "simulator"}}}); err != nil {
		t.Fatal(err)
	}
	device := validDeviceDefinition("PLC-unit-conflict")
	device.Destination = DestinationSelection{Port: 5021, UnitID: 1}
	_, err := store.ResolveDocumentDestinations(Document{Devices: []DeviceDefinition{device}})
	if !errors.Is(err, mma2composer.ErrReservationOwnedByOther) || !strings.Contains(err.Error(), "Unit ID 1") || !strings.Contains(err.Error(), "simulator") {
		t.Fatalf("error = %v, want foreign Unit ID ownership collision", err)
	}
}

func TestSaveTimeOwnershipGuardRejectsLatestForeignOwner(t *testing.T) {
	store := Store{Root: t.TempDir()}
	composer := mma2composer.New(store.Root, "simulator")
	owners := mma2composer.OwnershipDoc{Reservations: []mma2composer.OwnershipEntry{{Port: 5020, UnitID: 1, Owner: "simulator"}}}
	if err := composer.SaveOwners(owners); err != nil {
		t.Fatal(err)
	}
	for _, destination := range []DestinationSelection{{Port: 5020, UnitID: 2}, {Port: 5021, UnitID: 1}} {
		device := validDeviceDefinition("PLC-1")
		device.Destination = destination
		if err := store.CheckDocumentOwnership(Document{Devices: []DeviceDefinition{device}}); !errors.Is(err, mma2composer.ErrReservationOwnedByOther) {
			t.Fatalf("CheckDocumentOwnership(%#v) error = %v", destination, err)
		}
	}
	after, err := composer.LoadOwners()
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(after, owners) {
		t.Fatalf("foreign owners mutated: got %#v want %#v", after, owners)
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
	first.Destination = DestinationSelection{Port: 5021, UnitID: 2}
	second := validDeviceDefinition("PLC-2")
	second.Endpoint = "127.0.0.1:5022"
	second.UnitID = 2
	second.PullBlock.Function = 4
	second.Destination = DestinationSelection{Port: 5022, UnitID: 3}

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
		seen[fmt.Sprintf("%d/%d", entry.Port, entry.UnitID)] = entry.Owner
	}
	if seen["5020/1"] != "simulator" {
		t.Fatalf("foreign ownership lost: %#v", owners.Reservations)
	}
	if seen["5021/2"] != ProducerReplicator || seen["5022/3"] != ProducerReplicator {
		t.Fatalf("Replicator ownership incomplete: %#v", owners.Reservations)
	}

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

func TestDeviceRuntimeConfigMapsPullBlockOneToOne(t *testing.T) {
	device := validDeviceDefinition("PLC-1")
	device.Endpoint = "10.20.30.40:1502"
	device.UnitID = 7
	device.PullBlock = PullBlock{Function: 4, Start: 123, Count: 9, ScanRateMS: 750}
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
