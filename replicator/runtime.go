package replicator

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// RuntimeState is the truthful current state of the single-range poll loop.
type RuntimeState struct {
	Running        bool
	Cycles         uint64
	LastSuccess    bool
	LastError      string
	LastCycleStart time.Time
	LastCycleEnd   time.Time
}

// Runtime runs the configured single-range replication cycle serially.
type Runtime struct {
	store Store
	cycle func() (CycleResult, error)

	mu    sync.RWMutex
	state RuntimeState
}

// NewRuntime builds a poll runtime around Store.RunOnce.
func NewRuntime(store Store) *Runtime {
	r := &Runtime{store: store}
	r.cycle = store.RunOnce
	return r
}

func newRuntimeWithCycle(store Store, cycle func() (CycleResult, error)) *Runtime {
	return &Runtime{store: store, cycle: cycle}
}

// Snapshot returns a race-safe copy of the current runtime state.
func (r *Runtime) Snapshot() RuntimeState {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return r.state
}

// Run validates persisted configuration, verifies destination ownership/config
// readiness, then runs one cycle immediately and one serial cycle per interval
// until ctx is cancelled. Cycle errors are recorded truthfully and the loop
// continues on the next tick; there is no overlapping execution.
func (r *Runtime) Run(ctx context.Context) error {
	cfg, err := r.store.Load()
	if err != nil {
		return err
	}
	if err := validateCycleMapping(cfg); err != nil {
		return err
	}
	if cfg.Source.PollIntervalMS == 0 {
		return fmt.Errorf("source.poll_interval_ms must be > 0")
	}
	if err := r.store.composeDestination(cfg.Destination); err != nil {
		return err
	}

	interval := time.Duration(cfg.Source.PollIntervalMS) * time.Millisecond
	r.setRunning(true)
	defer r.setRunning(false)

	r.runCycle()
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return nil
		case <-ticker.C:
			r.runCycle()
		}
	}
}

func (r *Runtime) runCycle() {
	start := time.Now()
	_, err := r.cycle()
	end := time.Now()

	r.mu.Lock()
	defer r.mu.Unlock()
	r.state.Cycles++
	r.state.LastCycleStart = start
	r.state.LastCycleEnd = end
	if err != nil {
		r.state.LastSuccess = false
		r.state.LastError = err.Error()
		return
	}
	r.state.LastSuccess = true
	r.state.LastError = ""
}

func (r *Runtime) setRunning(running bool) {
	r.mu.Lock()
	r.state.Running = running
	r.mu.Unlock()
}
