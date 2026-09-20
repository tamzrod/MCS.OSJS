package replicator

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/tamzrod/MCS.OSJS/mma2composer"
)

// Once MMA2 has been committed, a missing restart ACK cannot produce a
// successful Save & Apply or restart stale pollers against the new mapping.
func TestManagerCommittedUnacknowledgedRestartFailsClosed(t *testing.T) {
	store := Store{Root: t.TempDir()}
	manager := NewRuntimeManager(store)
	manager.timeout = 80 * time.Millisecond
	defer manager.Stop()
	device := validDeviceDefinition("RECOVERY-REP")
	device.Destination = DestinationSelection{Port: 15021, UnitID: 1}
	_, _, err := manager.Apply(Document{Devices: []DeviceDefinition{device}})
	if err == nil || !strings.Contains(err.Error(), "recovery required") {
		t.Fatalf("committed but unacknowledged restart must fail with recovery state: %v", err)
	}
	if len(manager.runtimes) != 0 { t.Fatalf("unacknowledged config must not arm pollers: %+v", manager.runtimes) }
	composer := mma2composer.New(store.Root, ProducerReplicator)
	owners, err := composer.LoadOwners()
	if err != nil { t.Fatal(err) }
	if len(owners.Reservations) != 1 || owners.Reservations[0].Owner != ProducerReplicator {
		t.Fatalf("committed owner was lost or changed: %+v", owners)
	}
	if _, err := os.Stat(filepath.Join(composer.ConfigDir(), restartRequestFile)); err != nil {
		t.Fatalf("pending restart request must be retained for reconciliation: %v", err)
	}
	if _, err := os.Stat(store.DocumentPath()); !os.IsNotExist(err) {
		t.Fatalf("unacknowledged apply must not persist a success document: %v", err)
	}
}
