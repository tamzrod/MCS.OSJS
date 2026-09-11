package replicator

import "testing"

func TestDestinationMemoryForBlocksAllowsModbusServing(t *testing.T) {
	memory, err := destinationMemoryForBlocks(7, []PullBlock{
		{Function: 3, Start: 0, Count: 16, ScanRateMS: 1000},
		{Function: 4, Start: 0, Count: 16, ScanRateMS: 1000},
	})
	if err != nil {
		t.Fatal(err)
	}
	if memory.Policy == nil || len(memory.Policy.Rules) == 0 {
		t.Fatal("Replicator destination memory has no Modbus access policy")
	}
	rule := memory.Policy.Rules[0]
	if rule.ID != "replicator-fc-access" {
		t.Fatalf("policy rule id = %q", rule.ID)
	}
	allowed := map[uint8]bool{}
	for _, fc := range rule.AllowFC {
		allowed[fc] = true
	}
	if !allowed[3] || !allowed[4] {
		t.Fatalf("Replicator serving policy must allow FC3 and FC4 reads; got %#v", rule.AllowFC)
	}
	if len(rule.SourceIP) == 0 {
		t.Fatal("Replicator serving policy has no allowed source networks")
	}
}
