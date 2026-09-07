package simulator

import (
	"testing"
	"time"
)

type fireCounter struct {
	buckets map[FCKind]int
	last    map[FCKind]time.Time
	seq     chan Values
}

func newFireCounter() *fireCounter {
	return &fireCounter{
		buckets: make(map[FCKind]int, 8),
		last:    make(map[FCKind]time.Time, 8),
		seq:     make(chan Values, 9),
	}
}

func (c *fireCounter) record(v Values) {
	c.seq <- v
}

func (c *fireCounter) drain() {
	for {
		select {
		case v := <-c.seq:
			c.buckets[v.FC]++
			c.last[v.FC] = time.Now()
		default:
			return
		}
	}
}

func (c *fireCounter) fireCounts() map[FCKind]int {
	c.drain()
	out := make(map[FCKind]int, 8)
	for k, v := range c.buckets {
		out[k] = v
	}
	return out
}

func schedulerDevice(a uint32, b uint32, c2 uint32, d uint32) DeviceDefinition {
	x := validDevice()
	x.RandomRuntime.FC1IntervalMS = a
	x.RandomRuntime.FC2IntervalMS = b
	x.RandomRuntime.FC3IntervalMS = c2
	x.RandomRuntime.FC4IntervalMS = d
	return x
}

func (c *fireCounter) waitFires(t *testing.T, fcs ...FCKind) {
	t.Helper()
	deadline := time.Now().Add(5 * time.Second)
	for {
		c.drain()
		ok := true
		for _, fc := range fcs {
			if c.buckets[fc] == 0 {
				ok = false
				break
			}
		}
		if ok {
			return
		}
		if time.Now().After(deadline) {
			t.Fatalf("timed out waiting for fires: buckets=%v", c.buckets)
		}
		time.Sleep(10 * time.Millisecond)
	}
}

func TestSchedulerRunsAllFCsConcurrentlyAtOwnCadence(t *testing.T) {
	c := newFireCounter()
	d := schedulerDevice(5, 10, 0, 0)
	s := NewScheduler(d, c.record)
	defer s.Stop()

	c.waitFires(t, FC1, FC2)
	time.Sleep(105 * time.Millisecond)
	b := c.fireCounts()
	c1 := b[FC1]
	c2 := b[FC2]
	if c1 < 8 || c2 < 4 {
		t.Fatalf("cadence too low: fc1=%d fc2=%d", c1, c2)
	}
	if c1 <= c2*3/2 {
		t.Fatalf("FC1 cadence not ~2x FC2: fc1=%d fc2=%d", c1, c2)
	}
	c3 := b[FC3]
	c4 := b[FC4]
	if c3 != 0 || c4 != 0 {
		t.Fatalf("unconfigured FC3/FC4 fired: %v", b)
	}
}

func TestSchedulerIntervalChangeUpdatesScheduleWithoutRestart(t *testing.T) {
	c := newFireCounter()
	d := schedulerDevice(4, 9, 9, 9)
	s := NewScheduler(d, c.record)
	defer s.Stop()

	c.waitFires(t, FC1, FC2, FC3, FC4)

	before := s.Timing()
	time.Sleep(20 * time.Millisecond)
	s.UpdateTiming(RandomRuntimeParams{
		FC1IntervalMS: 4,
		FC2IntervalMS: 9,
		FC3IntervalMS: 2,
		FC4IntervalMS: 9,
	})
	after := s.Timing()

	if after[FC1].Next.IsZero() {
		t.Fatalf("FC1 schedule lost")
	}
	if after[FC2].Next.IsZero() {
		t.Fatalf("FC2 schedule lost")
	}
	if after[FC4].Next.IsZero() {
		t.Fatalf("FC4 schedule lost")
	}
	if after[FC3].Next.IsZero() {
		t.Fatalf("FC3 schedule lost")
	}
	if !after[FC3].Next.After(before[FC3].Next) {
		t.Fatalf("FC3 next deadline not recomputed")
	}
	c.drain()
	if c.buckets[FC2] == 0 {
		t.Fatalf("FC2 stopped firing after FC3 interval change")
	}
}
