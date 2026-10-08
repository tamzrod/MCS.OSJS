package mma2composer

import (
	"reflect"
	"sort"
	"testing"
)

func runtimeSaveConfig(t *testing.T) (PersistenceRuntimeSaveConfig, *fakeAreaReader, *fakeSnapshotStore) {
	t.Helper()
	key := PersistenceMemoryKey{Port: 5020, UnitID: 1}
	rules := []PersistenceRBERule{
		{ID: 3, Area: "holding_registers", Start: 0, Count: 3, SystemOwned: true},
		{ID: 7, Area: "coils", Start: 0, Count: 8, SystemOwned: true},
	}
	reader := &fakeAreaReader{data: map[string][]byte{
		"holding_registers": {0, 1, 0, 2, 0, 3},
		"coils":             {0x00},
	}}
	store := &fakeSnapshotStore{}
	return PersistenceRuntimeSaveConfig{Key: key, Rules: rules, Reader: reader, Store: store}, reader, store
}

// PERSIST-R02 self-check: only the system-derived persistence RBE projection is
// subscribed, and a real event for a persistence rule reaches the writer for the
// correct memory/area identity.
func TestPersistenceRuntimeSaveSubscribesOnlyPersistenceRules(t *testing.T) {
	cfg, _, store := runtimeSaveConfig(t)
	save, err := NewPersistenceRuntimeSave(cfg)
	if err != nil {
		t.Fatal(err)
	}
	ids := save.SubscribedRuleIDs()
	sort.Slice(ids, func(i, j int) bool { return ids[i] < ids[j] })
	if !reflect.DeepEqual(ids, []uint8{3, 7}) {
		t.Fatalf("subscribed IDs wrong: %v", ids)
	}
	if save.Handles(99) || save.Handles(0) {
		t.Fatal("a non-persistence ID must not be handled")
	}
	// A persistence event reaches the writer and persists the area.
	save.Publish(3)
	if len(store.calls) != 1 || store.calls[0].area != "holding_registers" {
		t.Fatalf("persistence event did not reach the store: %+v", store.calls)
	}
	// A user-owned ID never reaches the save path.
	before := len(store.calls)
	save.Publish(99)
	if len(store.calls) != before {
		t.Fatal("a user/non-persistence ID must not write")
	}
}

// A caller cannot subscribe a user-owned rule as persistence.
func TestPersistenceRuntimeSaveRejectsUserOwnedRule(t *testing.T) {
	cfg, _, _ := runtimeSaveConfig(t)
	cfg.Rules = append(cfg.Rules, PersistenceRBERule{ID: 9, Area: "coils", Start: 0, Count: 8, SystemOwned: false})
	if _, err := NewPersistenceRuntimeSave(cfg); err == nil {
		t.Fatal("a user-owned persistence rule must be rejected")
	}
	if _, err := NewPersistenceRuntimeSave(PersistenceRuntimeSaveConfig{}); err == nil {
		t.Fatal("nil reader/store must be rejected")
	}
}

// PERSIST-R02 self-check: changed state saves; unchanged state causes no
// unnecessary rewrite. The status observes events/saves/changed bytes.
func TestPersistenceRuntimeSaveChangeOnly(t *testing.T) {
	cfg, reader, store := runtimeSaveConfig(t)
	save, err := NewPersistenceRuntimeSave(cfg)
	if err != nil {
		t.Fatal(err)
	}
	// First event: whole area written.
	save.Publish(3)
	if len(store.calls) != 1 || len(store.calls[0].data) != 6 {
		t.Fatalf("initial save wrong: %+v", store.calls)
	}
	// Unchanged: no store call.
	save.Publish(3)
	if len(store.calls) != 1 {
		t.Fatal("unchanged state must not rewrite the snapshot")
	}
	// One word changed: only that word written.
	reader.data["holding_registers"] = []byte{0, 1, 0, 9, 0, 3}
	save.Publish(3)
	if len(store.calls) != 2 || len(store.calls[1].data) != 2 {
		t.Fatalf("changed word save wrong: %+v", store.calls)
	}

	status := save.Status()
	if status.ConfiguredRules != 2 || status.Events != 3 {
		t.Fatalf("status event accounting wrong: %+v", status)
	}
	if status.Saves != 2 || status.BytesWritten != 8 {
		t.Fatalf("status save accounting wrong: %+v", status)
	}
	if status.LastError != "" {
		t.Fatalf("no error expected: %+v", status)
	}
}

// A read error surfaces in the observed status and writes nothing.
func TestPersistenceRuntimeSaveReadErrorObserved(t *testing.T) {
	cfg, reader, store := runtimeSaveConfig(t)
	save, err := NewPersistenceRuntimeSave(cfg)
	if err != nil {
		t.Fatal(err)
	}
	reader.err = errFake("memory unavailable")
	save.Publish(3)
	if len(store.calls) != 0 {
		t.Fatal("a read error must not write")
	}
	if save.Status().LastError == "" {
		t.Fatalf("a read error must be observable: %+v", save.Status())
	}
}

// PERSIST-R04 self-check: the save status records a genuine last-save instant
// only after a real save, never on a non-persistence event.
func TestPersistenceRuntimeSaveLastSaveAt(t *testing.T) {
	cfg, _, _ := runtimeSaveConfig(t)
	save, err := NewPersistenceRuntimeSave(cfg)
	if err != nil {
		t.Fatal(err)
	}
	if save.Status().LastSaveAt != "" {
		t.Fatal("no save yet must report no last-save instant")
	}
	save.Publish(3)
	first := save.Status().LastSaveAt
	if first == "" {
		t.Fatal("a real save must record a last-save instant")
	}
	save.Publish(99)
	if save.Status().LastSaveAt != first {
		t.Fatal("a non-persistence event must not change the last-save instant")
	}
}

type errFake string

func (e errFake) Error() string { return string(e) }
