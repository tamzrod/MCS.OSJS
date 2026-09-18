//go:build !windows

package simulator

import (
	"net"
	"time"
)

func runtimeTestSocketPath(root string) string {
	return RuntimeSocketPath(root)
}

func dialRuntimeSocket(endpoint string, timeout time.Duration) (net.Conn, error) {
	return net.DialTimeout("unix", endpoint, timeout)
}
