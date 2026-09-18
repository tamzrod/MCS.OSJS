//go:build !windows

package simulator

import (
	"fmt"
	"net"
	"os"
	"path/filepath"
	"time"
)

const RuntimeSocketRelPath = "run/modbus-simulator.sock"

func RuntimeSocketPath(root string) string {
	return filepath.Join(root, RuntimeSocketRelPath)
}

func listenRuntime(endpoint string) (net.Listener, func(), error) {
	if err := os.MkdirAll(filepath.Dir(endpoint), 0o755); err != nil {
		return nil, nil, err
	}
	if info, err := os.Lstat(endpoint); err == nil {
		if info.Mode()&os.ModeSocket == 0 {
			return nil, nil, fmt.Errorf("runtime socket path is not a socket: %s", endpoint)
		}
		conn, dialErr := net.DialTimeout("unix", endpoint, 150*time.Millisecond)
		if dialErr == nil {
			_ = conn.Close()
			return nil, nil, fmt.Errorf("runtime socket already has a live owner: %s", endpoint)
		}
		if err := os.Remove(endpoint); err != nil {
			return nil, nil, fmt.Errorf("remove stale runtime socket: %w", err)
		}
	} else if !os.IsNotExist(err) {
		return nil, nil, err
	}
	listener, err := net.Listen("unix", endpoint)
	if err != nil {
		return nil, nil, err
	}
	if err := os.Chmod(endpoint, 0o666); err != nil {
		_ = listener.Close()
		_ = os.Remove(endpoint)
		return nil, nil, err
	}
	return listener, func() { _ = os.Remove(endpoint) }, nil
}
