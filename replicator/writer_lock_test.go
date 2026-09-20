//go:build linux

package replicator

import (
	"errors"
	"os"
	"testing"
	"time"

	"github.com/tamzrod/MCS.OSJS/mma2composer"
)

// A concurrent Simulator/MMA2 writer holding the shared root's lock must
// prevent Replicator from reading stale owners or persisting a candidate.
func TestReplicatorComposeRejectsBusySharedWriterLock(t *testing.T) {
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
	case <-time.After(2 * time.Second): t.Fatal("writer failed to acquire shared lock")
	}
	device := validDeviceDefinition("LOCK-REP")
	device.Destination = DestinationSelection{Port: 15021, UnitID: 1}
	_, _, err := store.ComposeDocumentDestinations(Document{Devices: []DeviceDefinition{device}})
	if !errors.Is(err, mma2composer.ErrWriterLockTimeout) {
		close(release)
		<-finished
		t.Fatalf("replicator must timeout before any compose: %v", err)
	}
	composer := mma2composer.New(store.Root, ProducerReplicator)
	for _, path := range []string{composer.EffectiveConfigPath(), composer.OwnershipPath(), store.DocumentPath()} {
		if _, err := os.Stat(path); !os.IsNotExist(err) {
			close(release)
			<-finished
			t.Fatalf("blocked Replicator wrote %s: %v", path, err)
		}
	}
	close(release)
	if err := <-finished; err != nil { t.Fatal(err) }
	if _, _, err := store.ComposeDocumentDestinations(Document{Devices: []DeviceDefinition{device}}); err != nil {
		t.Fatalf("lock release should permit composition: %v", err)
	}
}
