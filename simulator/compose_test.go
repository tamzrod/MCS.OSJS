package simulator

import (
	"errors"
	"os"
	"path/filepath"
	"strconv"
	"testing"
)

// foreignFixture seeds an effective MMA2 config plus ownership registry with
// one foreign-owned(port,unit_id) reservation, e.g. a future Replicator.
func foreignFixture(s Store) error {
	cfg := EffectiveMMA2Config{
		Listeners: []MMA2Listener{
			{
				ID:     "replicator-5020-1",
				Listen: "0.0.0.0:61000",
				Memory: []MMA2Memory{
					{
						UnitID:      1,
						HoldingRegs: &MMA2Area{Start: 0, Count: 10},
						Policy: &MMA2Policy{Rules: []MMA2PolicyRule{
							{ID: "replicator-fc-access", SourceIP: []string{"0.0.0.0/0"}, AllowFC: []uint8{1, 2, 3, 4, 5, 6, 15, 16}},
						}},
					},
				},
			},
		},
	}
	owners := OwnershipDoc{Reservations: []OwnershipEntry{
		{Port: 61000, UnitID: 1, Owner: "replicator"},
	}}
	if err := s.saveEffective(cfg); err != nil {
		return err
	}
	return s.saveOwners(owners)
}

func simMMA2Params() MMA2Params {
	return MMA2Params{
		Port:   61001,
		UnitID: 1,
		FC1:    Area{Start: 0, Count: 64},
		FC2:    Area{Start: 0, Count: 64},
		FC3:    Area{Start: 0, Count: 100},
		FC4:    Area{Start: 0, Count: 105},
	}
}

func TestSaveAndComposeFreeReservationPersistsOwner(t *testing.T) {
	root := t.TempDir()
	s := Store{Root: root}
	if err := foreignFixture(s); err != nil {
		t.Fatal(err)
	}

	def := validDevice()
	def.MMA2 = simMMA2Params()
	if err := s.SaveAndCompose(def); err != nil {
		t.Fatal(err)
	}

	doc, err := s.Load()
	if err != nil {
		t.Fatal(err)
	}
	if len(doc.Devices) != 1 || doc.Devices[0].Name != def.Name {
		t.Fatalf("store device not persisted: %+v", doc.Devices)
	}

	cfg, err := s.loadEffective()
	if err != nil {
		t.Fatal(err)
	}
	if len(cfg.Listeners) != 2 {
		t.Fatalf("want 2 listeners, got %d: %+v", len(cfg.Listeners), cfg.Listeners)
	}
	foreignFound := false
	simFound := false
	for _, l := range cfg.Listeners {

		if listenPort(l.Listen) == 61000 {
			if len(l.Memory) != 1 || l.Memory[0].UnitID != 1 || l.Memory[0].HoldingRegs == nil {
				t.Fatalf("foreign memory altered: %+v", l.Memory)
			}
			foreignFound = true
		}

		if listenPort(l.Listen) == 61001 {
			if len(l.Memory) != 1 || l.Memory[0].UnitID != 1 {
				t.Fatalf("simulator memory missing: %+v", l.Memory)
			}
			mem := l.Memory[0]
			if mem.Coils == nil || mem.Coils.Count != 64 {
				t.Fatalf("fc1->coils mapping wrong: %+v", mem.Coils)
			}
			if mem.DiscreteInputs == nil || mem.DiscreteInputs.Count != 64 {
				t.Fatalf("fc2->discrete_inputs mapping wrong: %+v", mem.DiscreteInputs)
			}
			if mem.HoldingRegs == nil || mem.HoldingRegs.Count != 100 {
				t.Fatalf("fc3->holding_registers mapping wrong: %+v", mem.HoldingRegs)
			}
			if mem.InputRegs == nil || mem.InputRegs.Count != 105 {
				t.Fatalf("fc4->input_registers mapping wrong: %+v", mem.InputRegs)
			}
			if mem.Policy == nil || len(mem.Policy.Rules) != 1 {
				t.Fatalf("per-memory policy missing: %+v", mem.Policy)
			}
			simFound = true
		}
	}
	if !foreignFound || !simFound {
		t.Fatalf("effective config missing reservations: foreign=%v sim=%v", foreignFound, simFound)
	}

	owners, err := s.loadOwners()
	if err != nil {
		t.Fatal(err)
	}
	if len(owners.Reservations) != 2 {
		t.Fatalf("want 2 ownership entries, got %+v", owners.Reservations)
	}
	seen := map[string]string{}
	for _, r := range owners.Reservations {
		seen[filepath.Join(strconv.Itoa(int(r.Port)), strconv.Itoa(int(r.UnitID)))] = r.Owner
	}
	if seen["61000/1"] != "replicator" || seen["61001/1"] != ProducerSimulator {
		t.Fatalf("ownership entries wrong: %+v", seen)
	}
}
func TestComposeRejectsForeignCollisionUnchanged(t *testing.T) {
	root := t.TempDir()
	s := Store{Root: root}
	if err := foreignFixture(s); err != nil {
		t.Fatal(err)
	}
	cfgBefore, err := os.ReadFile(s.EffectiveConfigPath())
	if err != nil {
		t.Fatal(err)
	}
	ownBefore, err := os.ReadFile(s.OwnershipPath())
	if err != nil {
		t.Fatal(err)
	}

	def := validDevice()
	def.MMA2.Port = 61000
	def.MMA2.UnitID = 1
	if err := s.SaveAndCompose(def); !errors.Is(err, ErrReservationOwnedByOther) {
		t.Fatalf("want ErrReservationOwnedByOther, got %v", err)
	}

	cfgAfter, _ := os.ReadFile(s.EffectiveConfigPath())
	ownAfter, _ := os.ReadFile(s.OwnershipPath())
	if string(cfgAfter) != string(cfgBefore) || string(ownAfter) != string(ownBefore) {
		t.Fatal("conflicting save modified the prior effective configuration")
	}
	if _, err := os.Stat(s.DevicesPath()); !os.IsNotExist(err) {
		t.Fatalf("conflicting save persisted a simulator definition: %v", err)
	}
}
func TestSaveAndComposeUpdatesOwnReservationPreservesForeign(t *testing.T) {
	root := t.TempDir()
	s := Store{Root: root}
	if err := foreignFixture(s); err != nil {
		t.Fatal(err)
	}

	def := validDevice()
	def.MMA2 = simMMA2Params()
	if err := s.SaveAndCompose(def); err != nil {
		t.Fatal(err)
	}

	def2 := def
	def2.MMA2.FC3 = Area{Start: 0, Count: 200}
	if err := s.SaveAndCompose(def2); err != nil {
		t.Fatal(err)
	}

	cfg, err := s.loadEffective()
	if err != nil {
		t.Fatal(err)
	}
	simUpdated := false
	foreignIntact := false
	for _, l := range cfg.Listeners {
		switch listenPort(l.Listen) {
		case 61000:
			if len(l.Memory) != 1 || l.Memory[0].HoldingRegs == nil || l.Memory[0].HoldingRegs.Count != 10 {
				t.Fatalf("foreign reservation changed after simulator update: %+v", l.Memory)
			}
			foreignIntact = true
		case 61001:
			if len(l.Memory) != 1 || l.Memory[0].HoldingRegs == nil || l.Memory[0].HoldingRegs.Count != 200 {
				t.Fatalf("simulator reservation not updated: %+v", l.Memory)
			}
			simUpdated = true
		}
	}
	if !simUpdated || !foreignIntact {
		t.Fatalf("update lost a reservation: sim=%v foreign=%v", simUpdated, foreignIntact)
	}

	owners, err := s.loadOwners()
	if err != nil {
		t.Fatal(err)
	}
	if len(owners.Reservations) != 2 {
		t.Fatalf("ownership registry corrupted: %+v", owners.Reservations)
	}
	seen := map[string]string{}
	for _, r := range owners.Reservations {
		seen[filepath.Join(strconv.Itoa(int(r.Port)), strconv.Itoa(int(r.UnitID)))] = r.Owner
	}
	if seen["61000/1"] != "replicator" || seen["61001/1"] != ProducerSimulator {
		t.Fatalf("ownership owners wrong: %+v", seen)
	}
}
func TestDeleteAndComposeRemovesOnlyOwn(t *testing.T) {
	root := t.TempDir()
	s := Store{Root: root}
	if err := foreignFixture(s); err != nil {
		t.Fatal(err)
	}
	def := validDevice()
	def.MMA2 = simMMA2Params()
	if err := s.SaveAndCompose(def); err != nil {
		t.Fatal(err)
	}

	if err := s.DeleteAndCompose(61001, 1); err != nil {
		t.Fatal(err)
	}

	cfg, err := s.loadEffective()
	if err != nil {
		t.Fatal(err)
	}
	if len(cfg.Listeners) != 1 {
		t.Fatalf("want only foreign listener remaining, got %+v", cfg.Listeners)
	}
	if listenPort(cfg.Listeners[0].Listen) != 61000 {
		t.Fatalf("foreign reservation removed: %+v", cfg.Listeners)
	}
	owners, err := s.loadOwners()
	if err != nil {
		t.Fatal(err)
	}
	if len(owners.Reservations) != 1 || owners.Reservations[0].Owner != "replicator" {
		t.Fatalf("ownership registry wrong after delete: %+v", owners.Reservations)
	}
}
func TestDeleteAndComposeRejectsForeignOwned(t *testing.T) {
	root := t.TempDir()
	s := Store{Root: root}
	if err := foreignFixture(s); err != nil {
		t.Fatal(err)
	}
	if err := s.DeleteAndCompose(61000, 1); !errors.Is(err, ErrReservationOwnedByOther) {
		t.Fatalf("want ErrReservationOwnedByOther, got %v", err)
	}
}

func TestSaveAndComposeRejectsInvalidBeforeAnyPersist(t *testing.T) {
	root := t.TempDir()
	s := Store{Root: root}
	bad := validDevice()
	bad.MMA2.Port = 0
	if err := s.SaveAndCompose(bad); err == nil {
		t.Fatal("expected validation error")
	}
	if _, err := os.Stat(s.DevicesPath()); !os.IsNotExist(err) {
		t.Fatalf("invalid save persisted a definition: %v", err)
	}
	if _, err := os.Stat(s.EffectiveConfigPath()); !os.IsNotExist(err) {
		t.Fatalf("invalid save wrote an effective config: %v", err)
	}
	if _, err := os.Stat(s.OwnershipPath()); !os.IsNotExist(err) {
		t.Fatalf("invalid save wrote an ownership registry: %v", err)
	}
}
