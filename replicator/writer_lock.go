package replicator

import "github.com/tamzrod/MCS.OSJS/mma2composer"

// withWriterLock uses the SAME fixed host-volume file as the Simulator.
// Only outer transaction entry points call this; locked helpers must not nest.
func (s Store) withWriterLock(action func() error) error {
	return mma2composer.WithWriterLock(s.Root, mma2composer.DefaultWriterLockTimeout, action)
}
