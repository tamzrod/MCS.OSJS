//go:build linux

package simulator

import (
	"errors"
	"os"
	"testing"
	"time"

	"github.com/tamzrod/MCS.OSJS/mma2composer"
)

// A competing process using the same root must prevent even the first
// Simulator artifact mutation; no optimistic compose outside the shared lock.
func TestSimulatorComposeRejectsBusySharedWriterLock(t *testing.T) {
	store := Store{Root: t.TempDir()}
	acquired := make(chan struct{})
	release := make(chan struct{})
	finished := make(chan error, 1)
	go func() {
		finished <- mma2composer.WithWriterLock(store.Root, 10*time.Second, func() error {
			close(acquired)
			<-release
			return nil
		})
	}()
	select {
	case <-acquired:
	case <-time.After(2 * time.Second):
		t.Fatal("first writer did not acquire the lock")
	}
	device := validDevice()
	err := store.ComposeDocument(Document{Devices: []DeviceDefinition{device}})
	if !errors.Is(err, mma2composer.ErrWriterLockTimeout) {
		close(release)
		<-finished
		t.Fatalf("busy shared lock must fail closed: %v", err)
	}
	for _, path := range []string{store.EffectiveConfigPath(), store.OwnershipPath(), store.DevicesPath()} {
		if _, err := os.Stat(path); !os.IsNotExist(err) {
			close(release)
			<-finished
			t.Fatalf("blocked writer modified %s: %v", path, err)
		}
	}
	close(release)
	if err := <-finished; err != nil { t.Fatal(err) }
	if err := store.ComposeDocument(Document{Devices: []DeviceDefinition{device}}); err != nil {
		t.Fatalf("released lock must allow a new transaction: %v", err)
	}
}
