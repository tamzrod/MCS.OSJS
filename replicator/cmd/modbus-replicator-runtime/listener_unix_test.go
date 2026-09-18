//go:build !windows

package main

import (
	"net"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/tamzrod/MCS.OSJS/replicator"
)

func TestUnixRuntimeListenerMatchesRelayAndProtectsLiveOwner(t *testing.T) {
	root := t.TempDir()
	endpoint := replicator.RuntimeSocketPath(root)
	if want := filepath.Join(root, "run", "modbus-replicator.sock"); endpoint != want {
		t.Fatalf("runtime endpoint = %q, want %q", endpoint, want)
	}
	listener, cleanup, err := listenRuntime(endpoint)
	if err != nil {
		t.Fatal(err)
	}
	defer cleanup()
	defer listener.Close()
	conn, err := net.DialTimeout("unix", endpoint, time.Second)
	if err != nil {
		t.Fatalf("relay could not dial runtime: %v", err)
	}
	_ = conn.Close()
	if _, _, err := listenRuntime(endpoint); err == nil {
		t.Fatal("second listener stole socket from live owner")
	}
	if err := listener.Close(); err != nil {
		t.Fatal(err)
	}
	cleanup()
	if _, err := os.Lstat(endpoint); !os.IsNotExist(err) {
		t.Fatalf("socket retained after cleanup: %v", err)
	}
}

func TestUnixRuntimeListenerReplacesStaleSocket(t *testing.T) {
	endpoint := replicator.RuntimeSocketPath(t.TempDir())
	if err := os.MkdirAll(filepath.Dir(endpoint), 0o755); err != nil {
		t.Fatal(err)
	}
	stale, err := net.Listen("unix", endpoint)
	if err != nil {
		t.Fatal(err)
	}
	stale.(*net.UnixListener).SetUnlinkOnClose(false)
	if err := stale.Close(); err != nil {
		t.Fatal(err)
	}
	listener, cleanup, err := listenRuntime(endpoint)
	if err != nil {
		t.Fatalf("stale socket was not replaced: %v", err)
	}
	_ = listener.Close()
	cleanup()
}
