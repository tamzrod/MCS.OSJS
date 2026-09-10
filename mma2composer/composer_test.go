package mma2composer

import (
	"errors"
	"os"
	"testing"
)

func validCandidate(port, unitID uint16) EffectiveConfig {
	return EffectiveConfig{Listeners: []Listener{{
		ID:     "test-listener",
		Listen: "0.0.0.0:61001",
		Memory: []Memory{{
			UnitID:      unitID,
			HoldingRegs: &Area{Start: 0, Count: 2},
			Policy: &Policy{Rules: []PolicyRule{{
				ID:       "test-access",
				SourceIP: []string{"127.0.0.1"},
				AllowFC:  []uint8{3},
			}}},
		}},
	}}}
}

func TestProducerIdentityAndForeignCollision(t *testing.T) {
	owners := OwnershipDoc{Reservations: []OwnershipEntry{{Port: 61001, UnitID: 1, Owner: "simulator"}}}
	if err := Collision(61001, 1, "simulator", owners); err != nil {
		t.Fatalf("owner should retain its reservation: %v", err)
	}
	if err := Collision(61001, 1, "replicator", owners); !errors.Is(err, ErrReservationOwnedByOther) {
		t.Fatalf("foreign producer should be rejected, got %v", err)
	}
}

func TestFirstComeFirstSavePersistsProducer(t *testing.T) {
	c := New(t.TempDir(), "replicator")
	cfg := validCandidate(61001, 1)
	owners := OwnershipDoc{Reservations: []OwnershipEntry{{Port: 61001, UnitID: 1, Owner: c.Producer}}}
	if err := c.Commit(cfg, owners); err != nil {
		t.Fatal(err)
	}
	loaded, err := c.LoadOwners()
	if err != nil {
		t.Fatal(err)
	}
	if len(loaded.Reservations) != 1 || loaded.Reservations[0].Owner != "replicator" {
		t.Fatalf("producer identity not persisted: %+v", loaded.Reservations)
	}
	if err := Collision(61001, 1, "simulator", loaded); !errors.Is(err, ErrReservationOwnedByOther) {
		t.Fatalf("later producer should not take reservation, got %v", err)
	}
}

func TestDropOneReservationUsesRequestedPort(t *testing.T) {
	c := New(t.TempDir(), "simulator")
	cfg := EffectiveConfig{Listeners: []Listener{
		{ID: "first", Listen: "0.0.0.0:61001", Memory: []Memory{{UnitID: 1}}},
		{ID: "second", Listen: "0.0.0.0:61002", Memory: []Memory{{UnitID: 1}}},
	}}
	owners := OwnershipDoc{Reservations: []OwnershipEntry{
		{Port: 61001, UnitID: 1, Owner: "simulator"},
		{Port: 61002, UnitID: 1, Owner: "simulator"},
	}}
	cfg, owners = c.DropOneReservation(cfg, owners, 61001, 1)
	if len(cfg.Listeners) != 1 || ListenPort(cfg.Listeners[0].Listen) != 61002 {
		t.Fatalf("wrong listener removed: %+v", cfg.Listeners)
	}
	if len(owners.Reservations) != 1 || owners.Reservations[0].Port != 61002 {
		t.Fatalf("wrong ownership removed: %+v", owners.Reservations)
	}
}

func TestCommitRestoresConfigWhenOwnersReplaceFails(t *testing.T) {
	c := New(t.TempDir(), "simulator")
	prior := validCandidate(61001, 1)
	if err := c.SaveEffective(prior); err != nil {
		t.Fatal(err)
	}
	before, err := os.ReadFile(c.EffectiveConfigPath())
	if err != nil {
		t.Fatal(err)
	}
	if err := os.MkdirAll(c.OwnershipPath(), 0o755); err != nil {
		t.Fatal(err)
	}
	replacement := validCandidate(61002, 2)
	replacement.Listeners[0].Listen = "0.0.0.0:61002"
	err = c.Commit(replacement, OwnershipDoc{Reservations: []OwnershipEntry{{Port: 61002, UnitID: 2, Owner: "simulator"}}})
	if err == nil {
		t.Fatal("expected owners replace failure")
	}
	after, readErr := os.ReadFile(c.EffectiveConfigPath())
	if readErr != nil {
		t.Fatal(readErr)
	}
	if string(after) != string(before) {
		t.Fatal("effective config was not restored byte-for-byte")
	}
}
