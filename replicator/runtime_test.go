package replicator

import (
	"context"
	"errors"
	"sync"
	"testing"
	"time"
)

func runtimeConfig() Config {
	cfg := validConfig()
	cfg.Source.Function = 3
	cfg.Source.Count = 1
	cfg.Source.PollIntervalMS = 10
	cfg.Destination.Area = "fc3"
	cfg.Destination.Count = 1
	return cfg
}

func TestRuntimeRepeatsWithoutOverlap(t *testing.T) {
	store := Store{Root: t.TempDir()}
	if err := store.Save(runtimeConfig()); err != nil {
		t.Fatal(err)
	}

	var mu sync.Mutex
	active := 0
	maxActive := 0
	calls := 0
	cycle := func() (CycleResult, error) {
		mu.Lock()
		active++
		calls++
		if active > maxActive {
			maxActive = active
		}
		mu.Unlock()

		time.Sleep(20 * time.Millisecond)

		mu.Lock()
		active--
		mu.Unlock()
		return CycleResult{}, nil
	}

	runtime := newRuntimeWithCycle(store, cycle)
	ctx, cancel := context.WithTimeout(context.Background(), 85*time.Millisecond)
	defer cancel()
	if err := runtime.Run(ctx); err != nil {
		t.Fatalf("Run: %v", err)
	}

	mu.Lock()
	gotCalls, gotMax := calls, maxActive
	mu.Unlock()
	if gotCalls < 3 {
		t.Fatalf("calls = %d, want at least 3", gotCalls)
	}
	if gotMax != 1 {
		t.Fatalf("max concurrent cycles = %d, want 1", gotMax)
	}
	state := runtime.Snapshot()
	if state.Running {
		t.Fatal("runtime still marked running after context cancellation")
	}
	if !state.LastSuccess || state.LastError != "" {
		t.Fatalf("unexpected final success state: %+v", state)
	}
}

func TestRuntimeRecordsCycleErrorAndContinues(t *testing.T) {
	store := Store{Root: t.TempDir()}
	if err := store.Save(runtimeConfig()); err != nil {
		t.Fatal(err)
	}

	var mu sync.Mutex
	calls := 0
	cycle := func() (CycleResult, error) {
		mu.Lock()
		calls++
		mu.Unlock()
		return CycleResult{}, errors.New("controlled cycle failure")
	}

	runtime := newRuntimeWithCycle(store, cycle)
	ctx, cancel := context.WithTimeout(context.Background(), 35*time.Millisecond)
	defer cancel()
	if err := runtime.Run(ctx); err != nil {
		t.Fatalf("Run: %v", err)
	}

	mu.Lock()
	gotCalls := calls
	mu.Unlock()
	if gotCalls < 2 {
		t.Fatalf("calls = %d, want repeated attempts after failure", gotCalls)
	}
	state := runtime.Snapshot()
	if state.LastSuccess {
		t.Fatalf("failed cycle reported success: %+v", state)
	}
	if state.LastError != "controlled cycle failure" {
		t.Fatalf("LastError = %q", state.LastError)
	}
}

func TestRuntimeCancelStopsCleanly(t *testing.T) {
	store := Store{Root: t.TempDir()}
	cfg := runtimeConfig()
	cfg.Source.PollIntervalMS = 50
	if err := store.Save(cfg); err != nil {
		t.Fatal(err)
	}

	runtime := newRuntimeWithCycle(store, func() (CycleResult, error) {
		return CycleResult{}, nil
	})
	ctx, cancel := context.WithCancel(context.Background())
	done := make(chan error, 1)
	go func() { done <- runtime.Run(ctx) }()

	deadline := time.Now().Add(time.Second)
	for {
		state := runtime.Snapshot()
		if state.Running && state.Cycles >= 1 {
			break
		}
		if time.Now().After(deadline) {
			t.Fatalf("runtime did not start: %+v", state)
		}
		time.Sleep(time.Millisecond)
	}

	cancel()
	select {
	case err := <-done:
		if err != nil {
			t.Fatalf("Run after cancel: %v", err)
		}
	case <-time.After(250 * time.Millisecond):
		t.Fatal("runtime did not stop after cancellation")
	}
	if runtime.Snapshot().Running {
		t.Fatal("runtime state remained running after stop")
	}
}
