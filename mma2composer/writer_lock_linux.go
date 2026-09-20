//go:build linux

package mma2composer

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"syscall"
	"time"
)

// WithWriterLock serializes one complete configuration transaction among
// cooperating Linux processes using the same absolute OSJS_DATA_DIR root.
func WithWriterLock(root string, timeout time.Duration, action func() error) (result error) {
	if root == "" || !filepath.IsAbs(root) {
		return fmt.Errorf("mma2 writer lock requires an absolute data root")
	}
	if timeout <= 0 {
		return fmt.Errorf("mma2 writer lock requires a positive timeout")
	}
	if action == nil {
		return fmt.Errorf("mma2 writer lock requires an action")
	}

	dir := filepath.Join(filepath.Clean(root), "config", "mma2")
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return fmt.Errorf("create mma2 lock directory: %w", err)
	}
	path := filepath.Join(dir, ".writer.lock")
	fd, err := syscall.Open(path, syscall.O_RDWR|syscall.O_CREAT|syscall.O_CLOEXEC|syscall.O_NOFOLLOW, 0o600)
	if err != nil {
		return fmt.Errorf("open mma2 writer lock: %w", err)
	}
	file := os.NewFile(uintptr(fd), path)
	defer func() {
		if err := file.Close(); err != nil {
			result = errors.Join(result, fmt.Errorf("close mma2 writer lock: %w", err))
		}
	}()
	info, err := file.Stat()
	if err != nil {
		return fmt.Errorf("stat mma2 writer lock: %w", err)
	}
	if !info.Mode().IsRegular() || info.Mode().Perm()&0o077 != 0 {
		return fmt.Errorf("mma2 writer lock must be a private regular file: %s", path)
	}

	deadline := time.Now().Add(timeout)
	for {
		err := syscall.Flock(fd, syscall.LOCK_EX|syscall.LOCK_NB)
		if err == nil {
			break
		}
		if err != syscall.EWOULDBLOCK && err != syscall.EAGAIN {
			return fmt.Errorf("acquire mma2 writer lock: %w", err)
		}
		remaining := time.Until(deadline)
		if remaining <= 0 {
			return fmt.Errorf("%w after %s", ErrWriterLockTimeout, timeout)
		}
		if remaining > 20*time.Millisecond {
			remaining = 20 * time.Millisecond
		}
		time.Sleep(remaining)
	}
	defer func() {
		if err := syscall.Flock(fd, syscall.LOCK_UN); err != nil {
			result = errors.Join(result, fmt.Errorf("release mma2 writer lock: %w", err))
		}
	}()
	return action()
}
