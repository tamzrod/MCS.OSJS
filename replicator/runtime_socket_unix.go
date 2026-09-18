//go:build !windows

package replicator

import "path/filepath"

func RuntimeSocketPath(root string) string {
	return filepath.Join(root, "run", "modbus-replicator.sock")
}
