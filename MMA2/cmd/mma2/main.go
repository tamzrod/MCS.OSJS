// cmd/mma2/main.go
package main

import (
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"syscall"

	"mma2/internal/accessevents"
	"mma2/internal/authority"
	"mma2/internal/config"
	"mma2/internal/ingress"
	"mma2/internal/notify"
	"mma2/internal/transport/modbus"
	"mma2/internal/transport/rawingest"
	"mma2/internal/version"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalf("usage: mma2 <config.yaml>")
	}

	cfgPath := os.Args[1]

	ext := strings.ToLower(filepath.Ext(cfgPath))
	if ext != ".yaml" && ext != ".yml" {
		log.Fatalf("config file must be .yaml or .yml: %s", cfgPath)
	}

	cfg, err := config.Load(cfgPath)
	if err != nil {
		log.Fatalf("config load failed: %v", err)
	}
	if err := config.Validate(cfg); err != nil {
		log.Fatalf("config validation failed: %v", err)
	}

	log.Printf("[mma2 %s] config loaded and validated successfully", version.Version)

	store, err := config.BuildStore(cfg)
	if err != nil {
		log.Fatalf("build store: %v", err)
	}

	auth, err := authority.New(cfg)
	if err != nil {
		log.Fatalf("authority init failed: %v", err)
	}
	log.Println("authority policies loaded")

	notifier, err := notify.New(cfg, store)
	if err != nil {
		log.Fatalf("notify init failed: %v", err)
	}
	defer notifier.Close()
	log.Println("notify engine enabled")

	ae, err := accessevents.New(cfg)
	if err != nil {
		log.Fatalf("access events init failed: %v", err)
	}
	if ae != nil {
		defer ae.Close()
	}

	if cfg.AccessEvents != nil && cfg.AccessEvents.Enabled {
		mux := http.NewServeMux()
		mux.HandleFunc("/events", ae.HandleEvents)

		ln, err := net.Listen("tcp", cfg.AccessEvents.Output.Listen)
		if err != nil {
			log.Fatalf("access events: failed to bind %s: %v", cfg.AccessEvents.Output.Listen, err)
		}

		go func() {
			log.Printf("access events HTTP listening on %s", cfg.AccessEvents.Output.Listen)
			if err := http.Serve(ln, mux); err != nil {
				log.Fatalf("access events HTTP server failed: %v", err)
			}
		}()

		log.Println("access events engine started")
	} else {
		log.Println("access events disabled")
	}

	// --------------------
	// Start ingress
	// --------------------

	for _, gate := range cfg.Ingress {

		onModbus := func(conn net.Conn) {
			modbus.HandleConn(conn, store, auth, notifier, ae, cfg.Debug)
		}

		onRawIngest := func(conn net.Conn) {
			rawingest.HandleConn(conn, store, notifier)
		}

		l := ingress.NewListener(gate)

		go func(g ingress.Listener) {
			if err := g.ListenAndServe(onModbus, onRawIngest); err != nil {
				log.Fatalf("ingress %s failed: %v", gate.ID, err)
			}
		}(*l)
	}

	log.Println("mma2 ingress started")

	// An empty valid configuration has no ingress goroutine. A bare `select {}`
	// makes the Go runtime treat that state as a deadlock and abort the process.
	// Block on real process signals instead so MMA2 can stay alive while idle and
	// still terminate cleanly when the supervisor requests a restart/shutdown.
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop
	signal.Stop(stop)
	log.Println("mma2 shutdown requested")
}
