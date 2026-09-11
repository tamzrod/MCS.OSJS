package replicator

import "testing"

func TestDestinationMemorySupportsFC1AndFC2(t *testing.T) {
	memory, err := destinationMemoryForBlocks(7, []PullBlock{
		{Function: 1, Start: 0, Count: 16, ScanRateMS: 100},
		{Function: 2, Start: 32, Count: 8, ScanRateMS: 200},
	})
	if err != nil {
		t.Fatal(err)
	}
	if memory.UnitID != 7 {
		t.Fatalf("unit id = %d", memory.UnitID)
	}
	if memory.Coils == nil || memory.Coils.Start != 0 || memory.Coils.Count != 16 {
		t.Fatalf("coils = %#v", memory.Coils)
	}
	if memory.DiscreteInputs == nil || memory.DiscreteInputs.Start != 32 || memory.DiscreteInputs.Count != 8 {
		t.Fatalf("discrete inputs = %#v", memory.DiscreteInputs)
	}
	if memory.Policy == nil || len(memory.Policy.Rules) == 0 {
		t.Fatal("Replicator destination must remain externally serveable")
	}
}

func TestValidateDeviceAllowsAllReadFunctions(t *testing.T) {
	device := validDeviceDefinition("PLC-all-fc")
	device.PullBlocks = []PullBlock{
		{Function: 1, Start: 0, Count: 8, ScanRateMS: 100},
		{Function: 2, Start: 0, Count: 8, ScanRateMS: 100},
		{Function: 3, Start: 0, Count: 8, ScanRateMS: 100},
		{Function: 4, Start: 0, Count: 8, ScanRateMS: 100},
	}
	device.PullBlock = device.PullBlocks[0]
	if err := ValidateDeviceDefinition(device); err != nil {
		t.Fatalf("all FC1-FC4 blocks must validate: %v", err)
	}
}
