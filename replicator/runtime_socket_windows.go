//go:build windows

package replicator

func RuntimeSocketPath(_ string) string {
	return `\\.\pipe\mcs-modbus-replicator`
}
