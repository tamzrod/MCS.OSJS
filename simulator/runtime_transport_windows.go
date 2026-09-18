//go:build windows

package simulator

import (
	"net"

	"github.com/Microsoft/go-winio"
)

const RuntimePipePath = `\\.\pipe\mcs-modbus-simulator`

func RuntimeSocketPath(_ string) string {
	return RuntimePipePath
}

func listenRuntime(endpoint string) (net.Listener, func(), error) {
	listener, err := winio.ListenPipe(endpoint, &winio.PipeConfig{
		SecurityDescriptor: "D:P(A;;GA;;;SY)(A;;GA;;;BA)(A;;GRGW;;;AU)",
	})
	return listener, func() {}, err
}
