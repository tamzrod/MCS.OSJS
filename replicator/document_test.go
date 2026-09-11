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
	block := PullBlock{Function: 3, Start: 10, Count: 4, ScanRateMS: 250}
	return DeviceDefinition{
		Name:       name,
		Enabled:    true,
		Endpoint:   "127.0.0.1:5020",
		UnitID:     1,
		PullBlocks: []PullBlock{block},
		PullBlock:  block,
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

func TestLoadDocumentMigratesLegacyRangeIntoPullBlocks(t *testing.T) {
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
	if len(doc.Devices) != 1 || len(doc.Devices[0].PullBlocks) != 1 {
		t.Fatalf("migrated document = %#v", doc)
	}
	block := doc.Devices[0].PullBlocks[0]
	if block.Function != 4 || block.Start != 123 || block.Count != 9 || block.ScanRateMS != 750 {
		t.Fatalf("migrated pull block = %#v", block)
	}
}

func TestLoadDocumentMigratesSinglePullBlockIntoCollection(t *testing.T) {
	store := Store{Root: t.TempDir()}
	if err := os.MkdirAll(store.ReplicatorDir(), 0o755); err != nil {
		t.Fatal(err)
	}
	legacy := []byte(`devices:
  - name: PLC-old
    enabled: true
    endpoint: 127.0.0.1:5020
    unit_id: 7
    pull_block:
      function: 3
      start: 20
      count: 5
      scan_rate_ms: 300
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
	if len(doc.Devices[0].PullBlocks) != 1 || doc.Devices[0].PullBlocks[0].Start != 20 {
		t.Fatalf("single block migration = %#v", doc.Devices[0].PullBlocks)
	}
}

func TestMultiplePullBlocksPersistInOrder(t *testing.T) {
	store := Store{Root: t.TempDir()}
	device := validDeviceDefinition("PLC-multi")
	device.PullBlocks = []PullBlock{
		{Function: 3, Start: 0, Count: 8, ScanRateMS: 100},
		{Function: 4, Start: 100, Count: 4, ScanRateMS: 750},
	}
	device.PullBlock = device.PullBlocks[0]
	if err := store.SaveDocument(Document{Devices: []DeviceDefinition{device}}); err != nil {
		t.Fatal(err)
	}
	got, err := store.LoadDocument()
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(got.Devices[0].PullBlocks, device.PullBlocks) {
		t.Fatalf("pull block order changed: got %#v want %#v", got.Devices[0].PullBlocks, device.PullBlocks)
	}
}

func TestSuggestDestinationSkipsOccupiedPairOnly(t *testing.T) {
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
	owners := mma2composer.OwnershipDoc{Reservations: []mma2composer.OwnershipEntry{{Port: 5020, UnitID: 1, Owner: "simulator"}}}
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

func TestResolveAllowsSameUnitOnDifferentPort(t *testing.T) {
	store := Store{Root: t.TempDir()}
	composer := mma2composer.New(store.Root, "simulator")
	if err := composer.SaveOwners(mma2composer.OwnershipDoc{Reservations: []mma2composer.OwnershipEntry{{Port: 5020, UnitID: 1, Owner: "simulator"}}}); err != nil {
		t.Fatal(err)
	}
	device := validDeviceDefinition("PLC-same-unit-different-port")
	device.Destination = DestinationSelection{Port: 5022, UnitID: 1}
	resolved, err := store.ResolveDocumentDestinations(Document{Devices: []DeviceDefinition{device}})
	if err != nil {
		t.Fatalf("different port with same Unit ID must be allowed: %v", err)
	}
	if got := resolved.Devices[0].Destination; got.Port != 5022 || got.UnitID != 1 {
		t.Fatalf("resolved destination = %#v", got)
	}
}

func TestResolveAllowsSamePortWithDifferentUnit(t *testing.T) {
	store := Store{Root: t.TempDir()}
	composer := mma2composer.New(store.Root, "simulator")
	if err := composer.SaveOwners(mma2composer.OwnershipDoc{Reservations: []mma2composer.OwnershipEntry{{Port: 5020, UnitID: 1, Owner: "simulator"}}}); err != nil {
		t.Fatal(err)
	}
	device := validDeviceDefinition("PLC-same-port-different-unit")
	device.Destination = DestinationSelection{Port: 5020, UnitID: 2}
	resolved, err := store.ResolveDocumentDestinations(Document{Devices: []DeviceDefinition{device}})
	if err != nil {
		t.Fatalf("same port with different Unit ID must be allowed: %v", err)
	}
	if got := resolved.Devices[0].Destination; got.Port != 5020 || got.UnitID != 2 {
		t.Fatalf("resolved destination = %#v", got)
	}
}

func TestSaveTimeOwnershipGuardRejectsExactForeignPairOnly(t *testing.T) {
	store := Store{Root: t.TempDir()}
	composer := mma2composer.New(store.Root, "simulator")
	owners := mma2composer.OwnershipDoc{Reservations: []mma2composer.OwnershipEntry{{Port: 5020, UnitID: 1, Owner: "simulator"}}}
	if err := composer.SaveOwners(owners); err != nil {
		t.Fatal(err)
	}
	conflict := validDeviceDefinition("PLC-conflict")
	conflict.Destination = DestinationSelection{Port: 5020, UnitID: 1}
	if err := store.CheckDocumentOwnership(Document{Devices: []DeviceDefinition{conflict}}); !errors.Is(err, mma2composer.ErrReservationOwnedByOther) {
		t.Fatalf("exact pair conflict error = %v", err)
	}
	for _, destination := range []DestinationSelection{{Port: 5020, UnitID: 2}, {Port: 5022, UnitID: 1}} {
		device := validDeviceDefinition("PLC-allowed")
		device.Destination = destination
		if err := store.CheckDocumentOwnership(Document{Devices: []DeviceDefinition{device}}); err != nil {
			t.Fatalf("CheckDocumentOwnership(%#v) = %v, want allowed", destination, err)
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
		Policy:      &mma2composer.Policy{Rules: []mma2composer.PolicyRule{{ID: "foreign", SourceIP: []string{"0.0.0.0/0"}, AllowFC: []uint8{3}}}},
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
	second.PullBlocks[0].Function = 4
	second.PullBlock = second.PullBlocks[0]
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
		seen[fmt.Sprintf("%d/%d", entry.Port, entry.UnitID)] = entry.Owner
	}
	if seen["5020/1"] != "simulator" || seen["5021/1"] != ProducerReplicator || seen["5021/2"] != ProducerReplicator {
		t.Fatalf("ownership incomplete: %#v", owners.Reservations)
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
	foundForeign := false
	for _, entry := range after.Reservations {
		if entry.Port == 5020 && entry.UnitID == 1 && entry.Owner == "simulator" {
			foundForeign = true
		}
	}
	if !foundForeign {
		t.Fatalf("Simulator reservation was removed: %#v", after.Reservations)
	}
}

func TestDestinationMemorySpansAllBlocksByArea(t *testing.T) {
	memory, err := destinationMemoryForBlocks(7, []PullBlock{
		{Function: 3, Start: 10, Count: 4, ScanRateMS: 100},
		{Function: 3, Start: 14, Count: 5, ScanRateMS: 200},
		{Function: 4, Start: 100, Count: 2, ScanRateMS: 300},
	})
	if err != nil {
		t.Fatal(err)
	}
	if memory.UnitID != 7 || memory.HoldingRegs == nil || memory.HoldingRegs.Start != 10 || memory.HoldingRegs.Count != 9 {
		t.Fatalf("holding area = %#v", memory.HoldingRegs)
	}
	if memory.InputRegs == nil || memory.InputRegs.Start != 100 || memory.InputRegs.Count != 2 {
		t.Fatalf("input area = %#v", memory.InputRegs)
	}
}

func TestDeviceRuntimeConfigMapsPullBlockOneToOne(t *testing.T) {
	device := validDeviceDefinition("PLC-1")
	device.Endpoint = "10.20.30.40:1502"
	device.UnitID = 7
	block := PullBlock{Function: 4, Start: 123, Count: 9, ScanRateMS: 750}
	device.PullBlocks = []PullBlock{block}
	device.PullBlock = block
	device.Destination = DestinationSelection{Port: 2502, UnitID: 8}
	cfg, err := device.runtimeConfigForBlock(block)
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
