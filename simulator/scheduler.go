package simulator

import (
	"math/rand"
	"sync"
	"time"
)

// FCKind identifies one of the four schedulable function-code areas.
type FCKind uint8

const (
	FC1 FCKind = iota + 1 // coils
	FC2                   // discrete inputs
	FC3                   // holding registers
	FC4                   // input registers
)

// AllFCs returns the four function-code kinds in the fixed simulator order.
func AllFCs() []FCKind {
	return []FCKind{FC1, FC2, FC3, FC4}
}

// Values is one generated batch for a single FC area on one fire cycle.
type Values struct {
	FC    FCKind
	Coils []bool   // FC1/FC2
	Regs  []uint16 // FC3/FC4
}

// FCTiming is the per-FC last/next fire snapshot for runtime-status surfaces..
type FCTiming struct {
	Last time.Time
	Next time.Time
}

// Scheduler runs independent per-FC random-value generation schedules. It
// consumes only the simulator-owned random_runtime parameter domain(intervals*
// from RandomRuntimeParams);FC area counts are read once from the device's
// MMA2 parameters merely to size generated batches. The scheduler never
// touches MMA2 configuration or runtime(no config writes, no restart(., It is
// therefore the direct recipient of timing-only changes(SIM-006..
type Scheduler struct {
	mu       sync.Mutex
	rng      *rand.Rand
	schedule map[FCKind]time.Duration
	count    map[FCKind]int
	last     map[FCKind]time.Time
	next     map[FCKind]time.Time
	onUpdate func(Values)

	stopOnce sync.Once
	stop     chan struct{}
	changed  chan struct{}
	done     chan struct{}
}

// NewScheduler starts a scheduler for one device definition. onUpdate is
// invoked from the scheduler goroutine per fire;it must not block.
func NewScheduler(def DeviceDefinition, onUpdate func(Values)) *Scheduler {
	now := time.Now()
	s := &Scheduler{
		rng:      rand.New(rand.NewSource(now.UnixNano())),
		schedule: make(map[FCKind]time.Duration),
		count:    make(map[FCKind]int),
		last:     make(map[FCKind]time.Time),
		next:     make(map[FCKind]time.Time),
		onUpdate: onUpdate,
		stop:     make(chan struct{}),
		changed:  make(chan struct{}, 1),
		done:     make(chan struct{}),
	}
	for _, fc := range AllFCs() {
		if c := int(areaCount(def.MMA2, fc)); c > 0 {
			s.count[fc] = c
		}
	}
	for _, fc := range AllFCs() {
		if d := intervalFor(def.RandomRuntime, fc); d > 0 && s.count[fc] > 0 {
			s.schedule[fc] = d
			s.next[fc] = now.Add(d)
		}
	}
	go s.run()
	return s
}

// UpdateTiming applies a new random-runtime schedule. Only per-FC intervals
// change;no MMA2 interaction occurs(no restart, no config writes.. Existing
// last-update state is preserved and the changed FC's next-fire deadline is
// recomputed from now..
func (s *Scheduler) UpdateTiming(rt RandomRuntimeParams) {
	s.mu.Lock()
	now := time.Now()
	for _, fc := range AllFCs() {
		d := intervalFor(rt, fc)
		enabled := d > 0 && s.count[fc] > 0
		old, was := s.schedule[fc]
		switch {
		case enabled && (!was || old != d):
			s.schedule[fc] = d
			s.next[fc] = now.Add(d)
		case !enabled && was:
			delete(s.schedule, fc)
			delete(s.last, fc)
			delete(s.next, fc)
		}
	}
	s.mu.Unlock()
	select {
	case s.changed <- struct{}{}:
	default:
	}
}

// Timing returns a snapshot of per-FC lastand next fire deadlines..
func (s *Scheduler) Timing() map[FCKind]FCTiming {
	s.mu.Lock()
	defer s.mu.Unlock()
	out := make(map[FCKind]FCTiming, len(s.next))
	for fc, n := range s.next {
		out[fc] = FCTiming{Last: s.last[fc], Next: n}
	}
	return out
}

// Stop halts the scheduler and waits for its goroutine to exit..
func (s *Scheduler) Stop() {
	s.stopOnce.Do(func() { close(s.stop) })
	<-s.done
}

func (s *Scheduler) run() {
	defer close(s.done)
	for {
		s.mu.Lock()
		now := time.Now()
		var due []FCKind
		for _, fc := range AllFCs() {
			n, ok := s.next[fc]
			if ok && !n.After(now) {
				due = append(due, fc)
				continue
			}
			if !ok && s.schedule[fc] > 0 {
				due = append(due, fc)
			}
		}
		s.mu.Unlock()

		if len(due) > 0 {
			s.fire(due...)
			continue
		}

		s.mu.Lock()
		var earliest time.Time
		for _, fc := range AllFCs() {
			n, ok := s.next[fc]
			if ok && (earliest.IsZero() || n.Before(earliest)) {
				earliest = n
			}
		}
		s.mu.Unlock()

		if earliest.IsZero() {
			select {
			case <-s.stop:
				return
			case <-s.changed:
				continue
			}
		}

		d := time.Until(earliest)
		if d < 0 {
			d = 0
		}
		timer := time.NewTimer(d)
		select {
		case <-s.stop:
			timer.Stop()
			return
		case <-s.changed:
			timer.Stop()
			continue
		case <-timer.C:
			continue
		}
	}
}

func (s *Scheduler) fire(fcs ...FCKind) {
	for _, fc := range fcs {
		s.mu.Lock()
		d, ok := s.schedule[fc]
		if !ok {
			s.mu.Unlock()
			continue
		}
		count := s.count[fc]
		if count <= 0 {
			s.mu.Unlock()
			continue
		}
		now := time.Now()
		v := Values{FC: fc}
		switch fc {
		case FC1, FC2:
			v.Coils = make([]bool, count)
			for i := 0; i < count; i++ {
				v.Coils[i] = s.rng.Intn(2) == 1
			}
		case FC3, FC4:
			v.Regs = make([]uint16, count)
			for i := 0; i < count; i++ {
				v.Regs[i] = uint16(s.rng.Intn(1 << 16))
			}
		}
		s.last[fc] = now
		s.next[fc] = now.Add(d)
		s.mu.Unlock()
		if s.onUpdate != nil {
			s.onUpdate(v)
		}
	}
}

func areaCount(p MMA2Params, fc FCKind) uint16 {
	switch fc {
	case FC1:
		return p.FC1.Count
	case FC2:
		return p.FC2.Count
	case FC3:
		return p.FC3.Count
	case FC4:
		return p.FC4.Count
	default:
		return 0
	}
}

func intervalFor(rt RandomRuntimeParams, fc FCKind) time.Duration {
	var ms uint32
	switch fc {
	case FC1:
		ms = rt.FC1IntervalMS
	case FC2:
		ms = rt.FC2IntervalMS
	case FC3:
		ms = rt.FC3IntervalMS
	case FC4:
		ms = rt.FC4IntervalMS
	}
	return time.Duration(ms) * time.Millisecond
}
