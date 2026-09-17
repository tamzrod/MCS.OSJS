package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

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
	go watchDocument(ctx, simulator.Store{Root: root}, scheduler)
	if err := simulator.ServeRuntime(ctx, simulator.RuntimeSocketPath(root), service); err != nil {
		log.Fatal(err)
	}
}
func watchDocument(ctx context.Context, store simulator.Store, scheduler *simulator.SchedulerApplier) {
	var lastSize int64 = -1
	var lastModified int64
	if info, err := os.Stat(store.DevicesPath()); err == nil {
		lastSize = info.Size()
		lastModified = info.ModTime().UnixNano()
	}
	ticker := time.NewTicker(250 * time.Millisecond)
	defer ticker.Stop()
	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			info, err := os.Stat(store.DevicesPath())
			if err != nil {
				if !os.IsNotExist(err) {
					log.Printf("watch simulator document: %v", err)
				}
				continue
			}
			modified := info.ModTime().UnixNano()
			if info.Size() == lastSize && modified == lastModified {
				continue
			}
			doc, err := store.Load()
			if err != nil {
				log.Printf("reload simulator document: %v", err)
				continue
			}
			valid := true
			for _, device := range doc.Devices {
				if err := simulator.ValidateDevice(device); err != nil {
					log.Printf("reload simulator document: %v", err)
					valid = false
					break
				}
			}
			if !valid {
				continue
			}
			scheduler.ArmSchedules(doc)
			lastSize = info.Size()
			lastModified = modified
		}
	}
}
