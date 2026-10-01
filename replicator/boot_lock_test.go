//go:build linux

package replicator

import (
	"errors"
	"net"
	"os"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/tamzrod/MCS.OSJS/mma2composer"
	"github.com/tamzrod/MCS.OSJS/simulator"
)

func TestBootReadinessDoesNotBlockMemoryFC43Apply(t *testing.T) {
	root := t.TempDir()
	simStore := simulator.Store{Root: root}
	ready, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	defer ready.Close()
	simPort := uint16(ready.Addr().(*net.TCPAddr).Port)
	memory := simulator.DeviceDefinition{Name: "memory", Enabled: true,
		MMA2: simulator.MMA2Params{Port: simPort, UnitID: 1, FC3: simulator.Area{Count: 10}}}
	if err := simStore.SaveAndCompose(memory); err != nil {
		t.Fatal(err)
	}
	router, scheduler, err := simulator.NewRuntimeApplyRouter(simStore)
	if err != nil {
		t.Fatal(err)
	}
	defer scheduler.Stop()

	unavailable, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	repPort := uint16(unavailable.Addr().(*net.TCPAddr).Port)
	if err := unavailable.Close(); err != nil {
		t.Fatal(err)
	}
	store := Store{Root: root}
	device := validDeviceDefinition("waiting-replicator")
	device.Destination = DestinationSelection{Port: repPort, AutoUnitID: true}
	if err := store.SaveDocument(Document{Devices: []DeviceDefinition{device}}); err != nil {
		t.Fatal(err)
	}
	manager := NewRuntimeManager(store)
	manager.timeout = 2 * time.Second
	defer manager.Stop()
	bootDone := make(chan error, 1)
	go func() { bootDone <- manager.Boot() }()
	defer func() {
		select {
		case err := <-bootDone:
			if err == nil || !strings.Contains(err.Error(), "not ready") {
				t.Errorf("expected readiness failure: %v", err)
			}
		case <-time.After(3 * time.Second):
			t.Error("boot did not return")
		}
		if len(manager.runtimes) != 0 {
			t.Error("failed readiness started pollers")
		}
	}()

	// Ownership publication establishes that boot entered its transaction.
	composer := mma2composer.New(root, ProducerReplicator)
	deadline := time.Now().Add(time.Second)
	for {
		owners, err := composer.LoadOwners()
		if err != nil {
			t.Fatal(err)
		}
		published := false
		for _, owner := range owners.Reservations {
			if owner.Port == repPort && owner.Owner == ProducerReplicator {
				published = true
			}
		}
		if published {
			break
		}
		if time.Now().After(deadline) {
			t.Fatal("boot did not publish ownership")
		}
		time.Sleep(5 * time.Millisecond)
	}
	if err := mma2composer.WithWriterLock(root, 250*time.Millisecond, func() error { return nil }); err != nil {
		t.Fatalf("readiness owns the config lock: %v", err)
	}
	persisted, err := store.LoadDocument()
	if err != nil {
		t.Fatal(err)
	}
	if len(persisted.Devices) != 1 || persisted.Devices[0].Destination.UnitID != 1 {
		t.Fatal("resolved desired state was not published before readiness")
	}
	select {
	case err := <-bootDone:
		bootDone <- err
		t.Fatal("boot was no longer waiting")
	default:
	}

	// Exercise the actual Memory backend apply, with an isolated restart ACK.
	ackDone := make(chan error, 1)
	go func() {
		deadline := time.Now().Add(time.Second)
		for time.Now().Before(deadline) {
			request, present, err := simStore.LoadRestartRequest()
			if err != nil {
				ackDone <- err
				return
			}
			if present {
				ackDone <- os.WriteFile(simStore.RestartAckPath(), []byte(request.ConfigSHA256), 0o644)
				return
			}
			time.Sleep(5 * time.Millisecond)
		}
		ackDone <- errors.New("Memory apply did not request restart")
	}()
	identity := map[string]interface{}{"product_code": "SavedDuringReadiness"}
	memory.MMA2.FC43 = &identity
	result, err := router.Apply(simulator.Document{Devices: []simulator.DeviceDefinition{memory}})
	if err != nil {
		t.Fatal(err)
	}
	if err := <-ackDone; err != nil {
		t.Fatal(err)
	}
	if result.Path != simulator.ApplyStructural {
		t.Fatalf("wrong apply path: %s", result.Path)
	}
	select {
	case err := <-bootDone:
		bootDone <- err
		t.Fatal("Memory save waited for boot to finish")
	default:
	}
	cfg, err := composer.LoadEffective()
	if err != nil {
		t.Fatal(err)
	}
	foundIdentity, foundReplicator := false, false
	for _, listener := range cfg.Listeners {
		for _, mem := range listener.Memory {
			if mma2composer.ListenPort(listener.Listen) == simPort {
				foundIdentity = reflect.DeepEqual(mem.Extra["fc43"], identity)
			}
			if mma2composer.ListenPort(listener.Listen) == repPort {
				foundReplicator = mem.UnitID == 1
			}
		}
	}
	if !foundIdentity || !foundReplicator {
		t.Fatal("Memory save lost identity or foreign reservation")
	}
	owners, err := composer.LoadOwners()
	if err != nil {
		t.Fatal(err)
	}
	if len(owners.Reservations) != 2 {
		t.Fatalf("ownership lost after unlock: %+v", owners)
	}
	// Both directions of foreign ownership must remain forbidden after unlock.
	foreignMemory := memory
	foreignMemory.MMA2.Port = repPort
	if err := simStore.SaveAndCompose(foreignMemory); !errors.Is(err, mma2composer.ErrReservationOwnedByOther) {
		t.Fatalf("Simulator claimed Replicator memory: %v", err)
	}
	device.Destination.Port = simPort
	device.Destination.UnitID = 1
	device.Destination.AutoUnitID = false
	if _, _, err := store.ComposeDocumentDestinations(Document{Devices: []DeviceDefinition{device}}); !errors.Is(err, mma2composer.ErrReservationOwnedByOther) {
		t.Fatalf("Replicator claimed Simulator memory: %v", err)
	}
}
