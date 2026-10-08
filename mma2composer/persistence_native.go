package mma2composer

import (
	"fmt"
	"strings"
)

// persistenceAreaOrder is the canonical area order used when enumerating
// explicit persistence ranges, matching the MMA2 memory layout.
var persistenceAreaOrder = []string{"coils", "discrete_inputs", "holding_registers", "input_registers"}

// namedPersistenceRange pairs an explicit persisted range with its area key.
type namedPersistenceRange struct {
	area string
	r    PersistenceArea
}

// persistenceRangeSelections enumerates the explicit ranges in canonical area
// order. A nil Ranges contributes nothing.
func persistenceRangeSelections(r *PersistenceRanges) []namedPersistenceRange {
	if r == nil {
		return nil
	}
	var out []namedPersistenceRange
	add := func(area string, ranges []PersistenceArea) {
		for _, x := range ranges {
			out = append(out, namedPersistenceRange{area: area, r: x})
		}
	}
	add("coils", r.Coils)
	add("discrete_inputs", r.DiscreteInputs)
	add("holding_registers", r.HoldingRegs)
	add("input_registers", r.InputRegs)
	return out
}

// allocatedPersistenceArea returns the allocated window for one canonical area
// key. An area is allocated only when present with a non-zero count; its
// start/count remain the single source of truth.
func allocatedPersistenceArea(memory Memory, area string) (Area, bool) {
	var a *Area
	switch area {
	case "coils":
		a = memory.Coils
	case "discrete_inputs":
		a = memory.DiscreteInputs
	case "holding_registers":
		a = memory.HoldingRegs
	case "input_registers":
		a = memory.InputRegs
	}
	if a == nil || a.Count == 0 {
		return Area{}, false
	}
	return *a, true
}

// ValidateMemoryPersistence validates the native per-memory persistence block
// for one memory candidate. An absent block, or an explicit enabled:false,
// disables persistence for that memory only and is accepted. An enabled block
// requires a nonempty directory in the same memory entry, and any explicit
// ranges must be nonempty, contained within the owning memory's allocated area
// and nonoverlapping (per area). Validation is read-only and never mutates the
// caller's memory. Persistence never carries range identity beyond what is
// validated here.
func ValidateMemoryPersistence(memory Memory) error {
	p := memory.Persistence
	if p == nil || p.Enabled == nil || !*p.Enabled {
		return nil
	}
	if strings.TrimSpace(p.Directory) == "" {
		return fmt.Errorf("persistence.enabled requires a nonempty directory")
	}
	if p.Ranges == nil {
		return nil
	}
	selected := persistenceRangeSelections(p.Ranges)
	if len(selected) == 0 {
		return fmt.Errorf("persistence.ranges cannot be empty")
	}
	seen := make(map[string][]PersistenceArea)
	for _, sel := range selected {
		if sel.r.Count == 0 {
			return fmt.Errorf("persistence.ranges.%s: count must be > 0", sel.area)
		}
		alloc, ok := allocatedPersistenceArea(memory, sel.area)
		if !ok {
			return fmt.Errorf("persistence.ranges.%s: area is not allocated", sel.area)
		}
		start := uint32(sel.r.Start)
		count := uint32(sel.r.Count)
		if start < uint32(alloc.Start) || start+count > uint32(alloc.Start)+uint32(alloc.Count) {
			return fmt.Errorf("persistence.ranges.%s: range [%d..%d) is outside allocated memory [%d..%d)",
				sel.area, sel.r.Start, start+count, alloc.Start, uint32(alloc.Start)+uint32(alloc.Count))
		}
		for _, prev := range seen[sel.area] {
			a := uint32(sel.r.Start)
			b := a + uint32(sel.r.Count)
			x := uint32(prev.Start)
			y := x + uint32(prev.Count)
			if a < y && x < b {
				return fmt.Errorf("persistence.ranges.%s: overlapping ranges", sel.area)
			}
		}
		seen[sel.area] = append(seen[sel.area], sel.r)
	}
	return nil
}

// ValidateCandidatePersistence rejects a root-level persistence key in a
// composed effective configuration. Persistence is valid only at
// listeners[].memory[].persistence; a root block must fail validation rather
// than be silently ignored.
func ValidateCandidatePersistence(cfg EffectiveConfig) error {
	if _, ok := cfg.Extra["persistence"]; ok {
		return fmt.Errorf("root-level persistence is unsupported; configure listeners[].memory[].persistence")
	}
	return nil
}
