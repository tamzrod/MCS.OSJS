package mma2composer

import (
	"os"
	"testing"
)

func TestRebuildRetainsListenerAndForeignAdvancedSettings(t *testing.T) {
	composer := New(t.TempDir(), "simulator")
	foreign := Memory{UnitID: 2, Coils: &Area{Count: 4}, Extra: map[string]interface{}{
		"state_sealing": map[string]interface{}{"enabled": false},
		"rbe":           map[string]interface{}{"coils": []interface{}{map[string]interface{}{"id": 2, "name": "foreign", "start": 0, "count": 1}}},
	}}
	cfg := EffectiveConfig{Listeners: []Listener{{ID: "private", Listen: "127.0.0.1:5502", Extra: map[string]interface{}{"custom": "keep"}, Memory: []Memory{{UnitID: 1, Coils: &Area{Count: 4}}, foreign}}}, Extra: map[string]interface{}{"rbe": map[string]interface{}{"tcp": map[string]interface{}{"listen": "127.0.0.1:9001"}}}}
	owners := OwnershipDoc{Reservations: []OwnershipEntry{{Port: 5502, UnitID: 1, Owner: "simulator"}, {Port: 5502, UnitID: 2, Owner: "replicator"}}}
	if err := composer.Commit(cfg, owners); err != nil {
		t.Fatal(err)
	}
	cfg, owners = composer.DropProducerReservations(cfg, owners)
	cfg = AddMemory(cfg, "generated", "0.0.0.0:5502", Memory{UnitID: 1, Coils: &Area{Count: 8}})
	owners.Reservations = append(owners.Reservations, OwnershipEntry{Port: 5502, UnitID: 1, Owner: "simulator"})
	if err := composer.Commit(cfg, owners); err != nil {
		t.Fatal(err)
	}
	loaded, err := composer.LoadEffective()
	if err != nil {
		t.Fatal(err)
	}
	listener := loaded.Listeners[0]
	if listener.Listen != "127.0.0.1:5502" || listener.Extra["custom"] != "keep" || listener.ID != "private" {
		t.Fatal("listener settings lost")
	}
	if listener.Memory[0].UnitID != 2 || listener.Memory[0].Extra["rbe"] == nil || listener.Memory[0].Extra["state_sealing"] == nil {
		t.Fatal("foreign settings lost")
	}
	if loaded.Extra["rbe"] == nil {
		t.Fatal("RBE output lost")
	}
	before, _ := os.ReadFile(composer.EffectiveConfigPath())
	loaded.Listeners[0].Memory[1].Extra = map[string]interface{}{"rbe": map[string]interface{}{"coils": []interface{}{map[string]interface{}{"id": 2, "name": "duplicate", "start": 0, "count": 1}}}}
	if err := composer.Commit(loaded, owners); err == nil {
		t.Fatal("duplicate cross-producer ID accepted")
	}
	after, _ := os.ReadFile(composer.EffectiveConfigPath())
	if string(before) != string(after) {
		t.Fatal("invalid composition modified config")
	}
}
