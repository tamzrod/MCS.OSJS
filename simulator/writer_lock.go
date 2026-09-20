package simulator

import "github.com/tamzrod/MCS.OSJS/mma2composer"

// withWriterLock protects the outermost Simulator configuration transaction.
// Inner composer and Store helpers must not acquire a second flock on this path.
func (s Store) withWriterLock(action func() error) error {
	return mma2composer.WithWriterLock(s.Root, mma2composer.DefaultWriterLockTimeout, action)
}
