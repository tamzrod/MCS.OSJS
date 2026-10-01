//go:build linux

package mma2composer

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestConcurrentWritersPreserveFC43AndForeignMemory(t *testing.T) {
	root := t.TempDir()
	start := make(chan struct{})
	results := make(chan error, 2)
	var wg sync.WaitGroup
	for index, producer := range []string{"simulator", "replicator"} {
		wg.Add(1)
		go func(index int, producer string) {
			defer wg.Done()
			<-start
			results <- WithWriterLock(root, time.Second, func() error {
				composer := New(root, producer)
				cfg, err := composer.LoadEffective()
				if err != nil {
					return err
				}
				owners, err := composer.LoadOwners()
				if err != nil {
					return err
				}
				port := uint16(61001 + index)
				if err := Collision(port, 1, producer, owners); err != nil {
					return err
				}
				cfg, owners = composer.DropProducerReservations(cfg, owners)
				memory := Memory{UnitID: 1, HoldingRegs: &Area{Count: 10}}
				if producer == "simulator" {
					memory.Extra = map[string]interface{}{"fc43": map[string]interface{}{"product_code": "ConcurrentIdentity"}}
				}
				cfg = AddMemory(cfg, producer, fmt.Sprintf("0.0.0.0:%d", port), memory)
				owners.Reservations = append(owners.Reservations, OwnershipEntry{Port: port, UnitID: 1, Owner: producer})
				return composer.Commit(cfg, owners)
			})
		}(index, producer)
	}
	close(start)
	wg.Wait()
	close(results)
	for err := range results {
		if err != nil {
			t.Fatal(err)
		}
	}
	composer := New(root, "auditor")
	cfg, err := composer.LoadEffective()
	if err != nil {
		t.Fatal(err)
	}
	owners, err := composer.LoadOwners()
	if err != nil {
		t.Fatal(err)
	}
	if len(cfg.Listeners) != 2 || len(owners.Reservations) != 2 {
		t.Fatalf("lost a committed update: cfg=%+v owners=%+v", cfg, owners)
	}
	found := false
	for _, listener := range cfg.Listeners {
		if ListenPort(listener.Listen) == 61001 {
			fields, ok := listener.Memory[0].Extra["fc43"].(map[string]interface{})
			found = ok && fields["product_code"] == "ConcurrentIdentity"
		}
	}
	if !found {
		t.Fatal("concurrent composition lost FC43")
	}
}
