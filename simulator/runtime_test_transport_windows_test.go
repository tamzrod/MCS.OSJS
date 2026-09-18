//go:build windows

package simulator

import (
	"fmt"
	"net"
	"time"

	"github.com/Microsoft/go-winio"
)

func runtimeTestSocketPath(_ string) string {
	return fmt.Sprintf(`\\.\pipe\mcs-simulator-test-%d`, time.Now().UnixNano())
}

func dialRuntimeSocket(endpoint string, timeout time.Duration) (net.Conn, error) {
	return winio.DialPipe(endpoint, &timeout)
}
