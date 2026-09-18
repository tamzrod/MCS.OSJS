//go:build windows

package main

import (
	"net"

	"github.com/Microsoft/go-winio"
)

func listenRuntime(endpoint string) (net.Listener, func(), error) {
	listener, err := winio.ListenPipe(endpoint, &winio.PipeConfig{
		SecurityDescriptor: "D:P(A;;GA;;;SY)(A;;GA;;;BA)(A;;GRGW;;;AU)",
	})
	return listener, func() {}, err
}
