//go:build linux

package mma2composer

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

func TestWriterLockTimeoutAndRelease(t *testing.T) {
	root := t.TempDir()
	acquired := make(chan struct{})
	unblock := make(chan struct{})
	done := make(chan error, 1)
	sentinel := errors.New("action failed")
	go func() {
		done <- WithWriterLock(root, time.Second, func() error {
			close(acquired)
			<-unblock
			return sentinel
		})
	}()
	select {
	case <-acquired:
	case <-time.After(2 * time.Second):
		t.Fatal("first writer did not acquire lock")
	}
	run := false
	err := WithWriterLock(root, 60*time.Millisecond, func() error { run = true; return nil })
	if !errors.Is(err, ErrWriterLockTimeout) || run {
		t.Fatalf("contending writer must time out without action; err=%v run=%t", err, run)
	}
	close(unblock)
	if err := <-done; !errors.Is(err, sentinel) {
		t.Fatalf("action error not returned: %v", err)
	}
	if err := WithWriterLock(root, time.Second, func() error { run = true; return nil }); err != nil || !run {
		t.Fatalf("released lock could not be reacquired: err=%v run=%t", err, run)
	}
}

func TestWriterLockCrashRelease(t *testing.T) {
	if os.Getenv("MMA2_LOCK_CRASH_CHILD") == "1" {
		err := WithWriterLock(os.Getenv("MMA2_LOCK_TEST_ROOT"), time.Second, func() error {
			fmt.Fprintln(os.Stdout, "LOCKED")
			os.Exit(0) // no deferred unlock: kernel must release on process exit
			return nil
		})
		t.Fatalf("child failed to acquire writer lock: %v", err)
	}
	root := t.TempDir()
	cmd := exec.Command(os.Args[0], "-test.run=^TestWriterLockCrashRelease$")
	cmd.Env = append(os.Environ(), "MMA2_LOCK_CRASH_CHILD=1", "MMA2_LOCK_TEST_ROOT="+root)
	output, err := cmd.CombinedOutput()
	if err != nil || !strings.Contains(string(output), "LOCKED") {
		t.Fatalf("child failed: err=%v output=%q", err, output)
	}
	if err := WithWriterLock(root, time.Second, func() error { return nil }); err != nil {
		t.Fatalf("process-exit lock was not released: %v", err)
	}
}

func TestWriterLockRejectsAmbiguousPath(t *testing.T) {
	if err := WithWriterLock("", time.Second, func() error { t.Fatal("must not execute"); return nil }); err == nil {
		t.Fatal("empty root accepted")
	}
	if err := WithWriterLock("relative", time.Second, func() error { t.Fatal("must not execute"); return nil }); err == nil {
		t.Fatal("relative root accepted")
	}
	root := t.TempDir()
	lockDir := filepath.Join(root, "config", "mma2")
	if err := os.MkdirAll(lockDir, 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.Symlink(filepath.Join(root, "elsewhere"), filepath.Join(lockDir, ".writer.lock")); err != nil {
		t.Fatal(err)
	}
	if err := WithWriterLock(root, time.Second, func() error { t.Fatal("must not execute"); return nil }); err == nil {
		t.Fatal("symlink lock path accepted")
	}
}
