package replicator

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/tamzrod/MCS.OSJS/mma2composer"
)

func TestValidateRejectsDisjointSameFCPullBlocks(t *testing.T) {
	device := validDeviceDefinition("PLC-gap")
	device.PullBlocks = []PullBlock{
		{Function: 3, Start: 0, Count: 10, ScanRateMS: 100},
		{Function: 3, Start: 100, Count: 10, ScanRateMS: 200},
	}
	device.PullBlock = device.PullBlocks[0]

	err := ValidateDeviceDefinition(device)
	if err == nil {
		t.Fatal("expected disjoint same-FC blocks to be rejected")
	}
	if !strings.Contains(err.Error(), "must overlap or be contiguous") {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestValidateAllowsContiguousSameFCPullBlocks(t *testing.T) {
	device := validDeviceDefinition("PLC-contiguous")
	device.PullBlocks = []PullBlock{
		{Function: 3, Start: 0, Count: 10, ScanRateMS: 100},
		{Function: 3, Start: 10, Count: 10, ScanRateMS: 200},
	}
	device.PullBlock = device.PullBlocks[0]

	if err := ValidateDeviceDefinition(device); err != nil {
		t.Fatalf("contiguous same-FC blocks should be representable exactly: %v", err)
	}
}

// The edited document passes advisory preflight, but an unsupported advanced
// setting fails composition BEFORE any effective-config/owners commit. Only in
// that case is it safe to resume the previous pollers.
func TestApplyPreCommitFailureRestoresPreviousPollers(t *testing.T) {
	store, manager, previous := recoveryFixture(t)
	before, err := os.ReadFile(store.DocumentPath())
	if err != nil { t.Fatal(err) }

	edited := cloneDocument(previous)
	edited.Devices[0].MMA2Advanced = map[string]interface{}{"unsupported-test-key": true}
	_, _, err = manager.Apply(edited)
	if err == nil || !strings.Contains(err.Error(), "unsupported MMA2 advanced setting") {
		t.Fatalf("expected pre-commit advanced composition rejection, got %v", err)
	}
	assertPreviousPollerRunning(t, manager, previous)
	after, err := os.ReadFile(store.DocumentPath())
	if err != nil { t.Fatal(err) }
	if !bytes.Equal(before, after) { t.Fatal("pre-commit failure modified the persisted Replicator document") }
	composer := mma2composer.New(store.Root, ProducerReplicator)
	for _, path := range []string{composer.EffectiveConfigPath(), composer.OwnershipPath(), filepath.Join(composer.ConfigDir(), restartRequestFile)} {
		if _, err := os.Stat(path); !os.IsNotExist(err) {
			t.Fatalf("pre-commit failure unexpectedly created %s: %v", path, err)
		}
	}
}

// After a successful config+ownership commit, failed restart acknowledgement
// is a recovery state. Resuming old pollers against potentially NEW MMA2 memory
// is unsafe even if the Replicator document could not be saved.
func TestApplyPostCommitRestartFailureStopsPreviousPollers(t *testing.T) {
	store, manager, previous := recoveryFixture(t)
	before, err := os.ReadFile(store.DocumentPath())
	if err != nil { t.Fatal(err) }

	edited := cloneDocument(previous)
	edited.Devices[0].PullBlocks[0].ScanRateMS = 50
	edited.Devices[0].PullBlock = edited.Devices[0].PullBlocks[0]
	if _, _, err := manager.Apply(edited); err == nil || !strings.Contains(err.Error(), "recovery required") {
		t.Fatalf("committed, unacknowledged apply must require recovery, got %v", err)
	}
	status, err := manager.Status(previous.Devices[0].Name)
	if err != nil { t.Fatal(err) }
	if status.Running || len(status.Blocks) != 1 || status.Blocks[0].Running || len(manager.runtimes) != 0 {
		t.Fatalf("post-commit failure resurrected old pollers: status=%#v runtimes=%#v", status, manager.runtimes)
	}
	after, err := os.ReadFile(store.DocumentPath())
	if err != nil { t.Fatal(err) }
	if !bytes.Equal(before, after) { t.Fatal("unacknowledged apply falsely persisted the edited document") }
	composer := mma2composer.New(store.Root, ProducerReplicator)
	cfg, err := composer.LoadEffective()
	if err != nil || len(cfg.Listeners) != 1 { t.Fatalf("committed config unavailable: cfg=%#v err=%v", cfg, err) }
	owners, err := composer.LoadOwners()
	if err != nil || len(owners.Reservations) != 1 || owners.Reservations[0].Owner != ProducerReplicator {
		t.Fatalf("committed reservation lost: owners=%#v err=%v", owners, err)
	}
	if _, err := os.Stat(filepath.Join(composer.ConfigDir(), restartRequestFile)); err != nil {
		t.Fatalf("unacknowledged restart request not retained for reconciliation: %v", err)
	}
}

func recoveryFixture(t *testing.T) (Store, *RuntimeManager, Document) {
	t.Helper()
	store := Store{Root: t.TempDir()}
	device := validDeviceDefinition("PLC-recovery")
	device.Endpoint = "127.0.0.1:1"
	device.Destination = DestinationSelection{Port: 55021, UnitID: 7}
	previous := Document{Devices: []DeviceDefinition{device}}
	if err := store.SaveDocument(previous); err != nil { t.Fatal(err) }
	manager := NewRuntimeManager(store)
	manager.timeout = 60 * time.Millisecond
	t.Cleanup(manager.Stop)
	if err := manager.replaceRuntimes(previous); err != nil { t.Fatal(err) }
	return store, manager, previous
}

func assertPreviousPollerRunning(t *testing.T, manager *RuntimeManager, previous Document) {
	t.Helper()
	deadline := time.Now().Add(500 * time.Millisecond)
	for {
		status, err := manager.Status(previous.Devices[0].Name)
		if err != nil { t.Fatal(err) }
		if status.Running && len(status.Blocks) == 1 && status.Blocks[0].Running { break }
		if time.Now().After(deadline) { t.Fatalf("previous poller not restored after pre-commit failure: %#v", status) }
		time.Sleep(10 * time.Millisecond)
	}
	current := manager.Document()
	if len(current.Devices) != 1 || current.Devices[0].PullBlocks[0].ScanRateMS != previous.Devices[0].PullBlocks[0].ScanRateMS {
		t.Fatalf("manager document changed after pre-commit rejection: %#v", current)
	}
}
