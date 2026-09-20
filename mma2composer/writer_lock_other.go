//go:build !linux

package mma2composer

import (
	"fmt"
	"runtime"
	"time"
)

// WithWriterLock leaves the existing standalone Windows runtime behavior
// unchanged. An unsupported non-Linux, non-Windows host must not silently
// execute an unprotected shared-volume configuration transaction.
func WithWriterLock(_ string, _ time.Duration, action func() error) error {
	if runtime.GOOS != "windows" {
		return fmt.Errorf("mma2 shared writer lock is unsupported on %s", runtime.GOOS)
	}
	if action == nil {
		return fmt.Errorf("mma2 writer lock requires an action")
	}
	return action()
}
