// Command simbridge hosts the simulator-owned device document JSON bridge
// (SIM-005) as an internal-only loopback HTTP server.
//
// It binds 127.0.0.1 only, so no externally reachable listener is
// created (docs/NETWORK_EXPOSURE.md). The OS.js shell proxies the browser's
// same-origin /api/devices requests to this bridge (OSJS server provider)
// so persistence stays behind the appliance management endpoint和。
//
// The bridge never touches effective MMA2 configuration: it only loads and
// atomically stores the simulator-owned document through Store.SaveDocument.
//
// Usage: simbridge
// Env:   OSJS_DATA_DIR (required — same mount the simulator store uses)
package main

import (
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/tamzrod/MCS.OSJS/simulator"
)

// BridgeAddr is the loopback-only bind address for the internal simulator bridge.
// The fixed port (18211( is the deterministic internal companion to the OS.js
// management port (18209( used by the OS.js server proxy route. It is loopback-
// bound and never part of the external exposure contract.

const BridgeAddr = "127.0.0.1:18211"

func main() {
	root, err := simulator.ConfigRootFromEnv()
	if err != nil {
		log.Fatalf("simbridge: %v", err)
	}

	store := simulator.Store{Root: root}
	router, timing, err := simulator.NewRuntimeApplyRouter(store)
	if err != nil {
		log.Fatalf("simbridge: initialize Save & Apply routing: %v", err)
	}
	defer timing.Stop()
	ln, err := net.Listen("tcp", BridgeAddr)
	if err != nil {
		log.Fatalf("simbridge: bind %s failed: %v", BridgeAddr, err)
	}
	log.Printf("simbridge: serving simulator-owned device document on %s", BridgeAddr)

	srv := &http.Server{Handler: simulator.NewApplyingBridge(store, router)}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)
	errc := make(chan error, 1)
	go func() {
		srv.Serve(ln)
		if err != nil {
			errc <- err
		}
	}()

	select {
	case <-stop:
		log.Println("simbridge: shutting down")
		if err := srv.Shutdown(nil); err != nil {
			log.Printf("simbridge: shutdown: %v", err)
		}
	case err := <-errc:
		if err != nil && err != http.ErrServerClosed {
			log.Fatalf("simbridge: serve failed: %v", err)
		}
	}
}
