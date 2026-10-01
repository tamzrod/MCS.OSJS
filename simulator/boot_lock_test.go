//go:build linux

package simulator

import (
	"os"
	"testing"
	"time"

	"github.com/tamzrod/MCS.OSJS/mma2composer"
)

func TestSimulatorBootReadinessDoesNotBlockConfigWrites(t *testing.T) {
	store := Store{Root: t.TempDir()}
	device := validDevice()
	device.MMA2.Port = freePort(t)
	if err := store.SaveDocument(Document{Devices: []DeviceDefinition{device}}); err != nil {
		t.Fatal(err)
	}
	type bootResult struct {
		scheduler *SchedulerApplier
		err       error
	}
	done := make(chan bootResult, 1)
	go func() {
		_, scheduler, err := newRuntimeApplyRouter(store, 2*time.Second)
		done <- bootResult{scheduler, err}
	}()
	defer func() {
		select {
		case result := <-done:
			if result.err != nil {
				t.Errorf("boot failed: %v", result.err)
			}
			if result.scheduler != nil {
				defer result.scheduler.Stop()
				if result.scheduler.ready {
					t.Error("unavailable MMA2 reported ready")
				}
			}
		case <-time.After(3 * time.Second):
			t.Error("boot did not finish")
		}
	}()
	deadline := time.Now().Add(time.Second)
	for {
		if _, err := os.Stat(store.OwnershipPath()); err == nil {
			break
		} else if !os.IsNotExist(err) {
			t.Fatal(err)
		}
		if time.Now().After(deadline) {
			t.Fatal("boot did not compose")
		}
		time.Sleep(5 * time.Millisecond)
	}
	if err := mma2composer.WithWriterLock(store.Root, 250*time.Millisecond, func() error {
		composer := store.composer()
		cfg, err := composer.LoadEffective()
		if err != nil {
			return err
		}
		owners, err := composer.LoadOwners()
		if err != nil {
			return err
		}
		cfg.Extra = map[string]interface{}{"debug": true}
		return composer.Commit(cfg, owners)
	}); err != nil {
		t.Fatalf("readiness blocked a config writer: %v", err)
	}
	select {
	case result := <-done:
		done <- result
		t.Fatal("boot finished before concurrent write")
	default:
	}
	cfg, err := store.loadEffective()
	if err != nil {
		t.Fatal(err)
	}
	if cfg.Extra["debug"] != true || len(cfg.Listeners) != 1 {
		t.Fatal("unrelated config update lost memory or root setting")
	}
}
