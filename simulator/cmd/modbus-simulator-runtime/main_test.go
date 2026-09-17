package main

import (
	"context"
	"testing"
	"time"

	"github.com/tamzrod/MCS.OSJS/simulator"
)

func TestWatchDocumentReloadsNoneSchedule(t *testing.T) {
	store := simulator.Store{Root: t.TempDir()}
	random := simulator.DeviceDefinition{
		Name:    "watched",
		Enabled: true,
		MMA2: simulator.MMA2Params{
			Port:   5020,
			UnitID: 1,
			FC3:    simulator.Area{Start: 0, Count: 4},
		},
		RandomRuntime: simulator.RandomRuntimeParams{FC3IntervalMS: 60000},
	}
	if err := store.SaveDocument(simulator.Document{Devices: []simulator.DeviceDefinition{random}}); err != nil {
		t.Fatal(err)
	}
	scheduler := simulator.NewSchedulerApplier(store, simulator.Document{Devices: []simulator.DeviceDefinition{random}})
	defer scheduler.Stop()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go watchDocument(ctx, store, scheduler)

	none := random
	none.RandomRuntime.FC3IntervalMS = 0
	if err := store.SaveDocument(simulator.Document{Devices: []simulator.DeviceDefinition{none}}); err != nil {
		t.Fatal(err)
	}

	deadline := time.Now().Add(2 * time.Second)
	for time.Now().Before(deadline) {
		status, err := scheduler.RuntimeStatus(none.Name)
		if err == nil && status.RawIngest == "NOT REQUIRED" && len(status.FC) == 0 {
			return
		}
		time.Sleep(25 * time.Millisecond)
	}
	status, err := scheduler.RuntimeStatus(none.Name)
	t.Fatalf("None schedule was not reloaded: status=%+v err=%v", status, err)
}
