package replicator

import (
	"strings"
	"testing"
	"time"
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

func TestApplyRestartFailureRestoresPreviousPollers(t *testing.T) {
	store := Store{Root: t.TempDir()}
	device := validDeviceDefinition("PLC-recovery")
	device.Endpoint = "127.0.0.1:1"
	device.Destination = DestinationSelection{Port: 55021, UnitID: 7}
	previous := Document{Devices: []DeviceDefinition{device}}
	if err := store.SaveDocument(previous); err != nil {
		t.Fatal(err)
	}

	manager := NewRuntimeManager(store)
	manager.timeout = 60 * time.Millisecond
	defer manager.Stop()
	if err := manager.replaceRuntimes(previous); err != nil {
		t.Fatal(err)
	}

	edited := cloneDocument(previous)
	edited.Devices[0].PullBlocks[0].ScanRateMS = 50
	edited.Devices[0].PullBlock = edited.Devices[0].PullBlocks[0]
	if _, _, err := manager.Apply(edited); err == nil {
		t.Fatal("expected Apply to fail without an MMA2 restart acknowledgement")
	}

	deadline := time.Now().Add(500 * time.Millisecond)
	for {
		status, err := manager.Status(device.Name)
		if err != nil {
			t.Fatal(err)
		}
		if status.Running && len(status.Blocks) == 1 && status.Blocks[0].Running {
			break
		}
		if time.Now().After(deadline) {
			t.Fatalf("previous poller was not restored after failed apply: %#v", status)
		}
		time.Sleep(10 * time.Millisecond)
	}

	current := manager.Document()
	if len(current.Devices) != 1 || current.Devices[0].PullBlocks[0].ScanRateMS != previous.Devices[0].PullBlocks[0].ScanRateMS {
		t.Fatalf("manager document changed after rejected apply: %#v", current)
	}
}
