package main

import (
	"fmt"
	"testing"
)

func TestWatchDocumentReloadsNoneScheduleWithTrace(t *testing.T) {
	store := Store{Root: t.TempDir()}
	watched := DeviceDefinition{
		Name:    "watched",
		Enabled: true,
		MMA2: MMA2Params{
			Port:   5020,
			UnitID: 1,
			FC3:    Area{Start: 0, Count: 4},
		},
		RandomRuntime: RandomRuntimeParams{FC3IntervalMS: 60000},
	}
	if err := store.SaveDocument(Document{Devices: []DeviceDefinition{watched}}); err != nil {
		t.Fatal(err)
	}
	fmt.Printf("Initial FC3IntervalMS=%d\n", watched.RandomRuntime.FC3IntervalMS)

	scheduler := NewSchedulerApplier(store, Document{Devices: []DeviceDefinition{watched}})
	defer scheduler.Stop()

	watchedFC3 := 60000
	watchedRandomFC3IntervalMS := storedRandomFC3IntervalMS(watched, store)
	fmt.Printf("watched.MMA2.FC3.Count=%d watched.RandomRuntime.FC3IntervalMS=%d\n", watched.MMA2.FC3.Count, watchedRandomFC3IntervalMS)

	watched.RandomRuntime.FC3IntervalMS = 0
	savedWatched := storedDevice(watched, store)
	fmt.Printf("saved device RandomRuntime.FC3IntervalMS=%d\n", savedWatched.RandomRuntime.FC3IntervalMS)

	ticker := time.NewTicker(10 * time.Millisecond)
	defer ticker.Stop()
	go watchDocument(context.Background(), store, scheduler)

	i := 0
	for ; i < 60; i++ {
		status, err := scheduler.RuntimeStatus(watched.Name)
		if err != nil || status.RawIngest == "NOT REQUIRED" && len(status.FC) == 0 {
			break
		}
		time.Sleep(25 * time.Millisecond)
	}
	fmt.Printf("Final status: %+v\n", status)
	status, _ = scheduler.RuntimeStatus(watched.Name)
	if status.RawIngest != "NOT REQUIRED" || len(status.FC) != 0 {
		t.Fatalf("RawIngest=%q (expected NOT REQUIRED), FC count=%d (expected 0)\n", status.RawIngest, len(status.FC))
	}
	fmt.Println("TEST PASSED")
}

func storedRandomFC3IntervalMS(dev DeviceDefinition, store Store) int {
	var stored DeviceDefinition
	if err := store.withWriterLock(func() error {
		if loaded, err := store.Load(); err == nil {
			for _, d := range loaded.Devices { if d.Name == dev.Name { stored = d } }
		}
		return nil
	}); err != nil {
		t.Fatal(err)
	}
	return stored.RandomRuntime.FC3IntervalMS
}

func savedDevice(dev DeviceDefinition, store Store) DeviceDefinition {
	var result DeviceDefinition
	if err := store.withWriterLock(func() error {
		if loaded, err := store.Load(); err == nil {
			for _, d := range loaded.Devices { if d.Name == dev.Name { result = d } }
		}
		return nil
	}); err != nil {
		t.Fatal(err)
	}
	return result
}

