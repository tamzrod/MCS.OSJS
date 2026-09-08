package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/tamzrod/MCS.OSJS/simulator"
)

func main() {
	root, err := simulator.ConfigRootFromEnv()
	if err != nil {
		log.Fatal(err)
	}
	service, scheduler, err := simulator.NewLiveRuntimeService(simulator.Store{Root: root})
	if err != nil {
		log.Fatal(err)
	}
	defer scheduler.Stop()
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()
	if err := simulator.ServeRuntime(ctx, simulator.RuntimeSocketPath(root), service); err != nil {
		log.Fatal(err)
	}
}
