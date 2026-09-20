//go:build linux

package mma2composer

import (
	"errors"
	"sync"
	"testing"
	"time"
)

// Two producers racing for the exact same pair must serialize their complete
// read/collision/compose/commit transactions. Exactly one owner survives.
func TestWriterLockConcurrentForeignOwnershipConflict(t *testing.T) {
	root := t.TempDir()
	start := make(chan struct{})
	results := make(chan error, 2)
	var wg sync.WaitGroup
	for _, producer := range []string{"simulator", "replicator"} {
		wg.Add(1)
		go func(producer string) {
			defer wg.Done()
			<-start
			results <- WithWriterLock(root, 5*time.Second, func() error {
				composer := New(root, producer)
				cfg, err := composer.LoadEffective()
				if err != nil { return err }
				owners, err := composer.LoadOwners()
				if err != nil { return err }
				if err := Collision(61001, 1, producer, owners); err != nil { return err }
				cfg = AddMemory(cfg, "test-61001", "0.0.0.0:61001", validCandidate(61001, 1).Listeners[0].Memory[0])
				owners.Reservations = append(owners.Reservations, OwnershipEntry{Port: 61001, UnitID: 1, Owner: producer})
				return composer.Commit(cfg, owners)
			})
		}(producer)
	}
	close(start)
	wg.Wait()
	close(results)
	passed, conflicted := 0, 0
	for err := range results {
		switch {
		case err == nil: passed++
		case errors.Is(err, ErrReservationOwnedByOther): conflicted++
		default: t.Fatalf("unexpected writer error: %v", err)
		}
	}
	if passed != 1 || conflicted != 1 { t.Fatalf("want one commit and one foreign collision, got %d/%d", passed, conflicted) }
	composer := New(root, "auditor")
	owners, err := composer.LoadOwners()
	if err != nil { t.Fatal(err) }
	cfg, err := composer.LoadEffective()
	if err != nil { t.Fatal(err) }
	if len(owners.Reservations) != 1 || len(cfg.Listeners) != 1 || len(cfg.Listeners[0].Memory) != 1 {
		t.Fatalf("racing writers corrupted shared state: owners=%+v cfg=%+v", owners, cfg)
	}
}
