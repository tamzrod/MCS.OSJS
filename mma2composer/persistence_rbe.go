package mma2composer

import "fmt"

// PersistenceRBERule is one system-owned persistence RBE rule derived from a
// memory area. It has no ID, name or independent range: start/count are copied
// from the authoritative memory area, which remains the single source of truth.
// Global ID assignment and user-rule collision policy are separate tasks.
//
// SystemOwned is the lock marker: persistence rules are system-owned and are
// non-editable/non-deletable through supported configuration paths. Ordinary
// user-authored RBE rules live elsewhere and are never marked system-owned.
type PersistenceRBERule struct {
	Area        string `yaml:"area" json:"area"`
	Start       uint16 `yaml:"start" json:"start"`
	Count       uint16 `yaml:"count" json:"count"`
	SystemOwned bool   `yaml:"system_owned" json:"system_owned"`
}

// persistenceAreas pairs each canonical MMA2 RBE area key with its layout in
// canonical order. A nil or zero-count area is absent.
func persistenceAreas(memory Memory) []struct {
	name string
	area *Area
} {
	return []struct {
		name string
		area *Area
	}{
		{"coils", memory.Coils},
		{"discrete_inputs", memory.DiscreteInputs},
		{"holding_registers", memory.HoldingRegs},
		{"input_registers", memory.InputRegs},
	}
}

// DerivePersistenceRBE returns one system-owned persistence RBE rule for every
// present memory area, in canonical order. The derived range is read directly
// from the area, so it can never diverge from the authoritative layout and no
// independent editable range is introduced. Areas with count 0 are omitted.
func DerivePersistenceRBE(memory Memory) []PersistenceRBERule {
	var rules []PersistenceRBERule
	for _, entry := range persistenceAreas(memory) {
		if entry.area == nil || entry.area.Count == 0 {
			continue
		}
		rules = append(rules, PersistenceRBERule{
			Area:        entry.name,
			Start:       entry.area.Start,
			Count:       entry.area.Count,
			SystemOwned: true,
		})
	}
	return rules
}

// PersistenceRBERuleKey identifies one derived persistence rule by area key.
// Range is intentionally excluded: a rule's identity is its area, and its range
// is derived from that area rather than authored.
func PersistenceRBERuleKey(rule PersistenceRBERule) string { return rule.Area }

// SynchronizePersistenceRBE keeps the persistence-owned RBE projection aligned
// with the authoritative memory layout. The returned projection is exactly the
// layout-derived set for the current areas, so an existing rule's start/count
// follows any area start/count change with no second edit, newly allocated
// areas appear, and removed areas drop out. Every rule remains system-owned.
//
// Callers persist the returned set; they never hand-author a range. Ordinary
// user-owned RBE rules are a separate set and are never touched here.
func SynchronizePersistenceRBE(memory Memory) []PersistenceRBERule {
	return DerivePersistenceRBE(memory)
}

// presentPersistenceAreas returns the set of canonical area keys currently
// present (non-nil and count > 0) in the authoritative memory layout.
func presentPersistenceAreas(memory Memory) map[string]bool {
	present := make(map[string]bool)
	for _, entry := range persistenceAreas(memory) {
		if entry.area != nil && entry.area.Count > 0 {
			present[entry.name] = true
		}
	}
	return present
}

// PruneRemovedPersistenceRBE drops the persistence RBE rule for each area no
// longer present in the authoritative memory layout, leaving every other
// projection (and its range) byte-for-byte unchanged. It is the removal
// counterpart to SynchronizePersistenceRBE: removing one area removes exactly
// its own system persistence RBE, while other persisted-area projections
// remain. The input slice is never mutated.
//
// Only the persistence projection is passed here; ordinary user-owned RBE rules
// are a separate set and are untouched. Persistence-disable behavior is not
// expressed by this function.
func PruneRemovedPersistenceRBE(projection []PersistenceRBERule, memory Memory) []PersistenceRBERule {
	present := presentPersistenceAreas(memory)
	kept := make([]PersistenceRBERule, 0, len(projection))
	for _, rule := range projection {
		if present[rule.Area] {
			kept = append(kept, rule)
		}
	}
	return kept
}

// PersistenceEnabled reports whether persistence is enabled for this memory.
// Persistence is off unless it is explicitly enabled.
func PersistenceEnabled(memory Memory) bool {
	return memory.Persistence != nil && memory.Persistence.Enabled != nil && *memory.Persistence.Enabled
}

// PersistenceRBEProjection returns the persistence-owned RBE projection for a
// memory as a function of its persistence enablement and its authoritative
// layout:
//
//   - persistence disabled -> no persistence-owned RBE (cleanup); the caller's
//     existing projection is dropped wholesale;
//   - persistence enabled  -> one layout-derived system-owned rule per present
//     area (regeneration), identical to SynchronizePersistenceRBE.
//
// Disabling removal and re-enabling regeneration therefore use the same source
// of truth and never require a hand-authored range. Only persistence-owned rules
// are represented here; ordinary user-owned RBE entries are a separate set and
// are never deleted or rewritten. This is not a file or runtime operation.
func PersistenceRBEProjection(memory Memory) []PersistenceRBERule {
	if !PersistenceEnabled(memory) {
		return nil
	}
	return DerivePersistenceRBE(memory)
}

// rejectUserOwnedRules is a fail-closed guard for the supported configuration
// path: every persistence rule must be system-owned. A caller cannot submit
// user-owned (editable/deletable) persistence rules.
func rejectUserOwnedRules(rules []PersistenceRBERule) error {
	for _, rule := range rules {
		if !rule.SystemOwned {
			return fmt.Errorf("persistence RBE rule for area %q must be system-owned", rule.Area)
		}
	}
	return nil
}

// ValidatePersistenceRuleMutation guards an edit to the persistence RBE set.
// The authoritative set is always the layout-derived projection, so a caller
// cannot:
//   - delete a system-owned rule (an expected area missing from the request);
//   - add or change a rule's range or ownership (any requested rule that is not
//     byte-identical to the layout-derived system-owned rule);
//   - submit a user-owned persistence rule.
//
// It never mutates the request or the derived set. Ordinary user-authored RBE
// rules are not part of this set and are unaffected.
func ValidatePersistenceRuleMutation(requested []PersistenceRBERule, derived []PersistenceRBERule) error {
	if err := rejectUserOwnedRules(requested); err != nil {
		return err
	}
	derivedByArea := make(map[string]PersistenceRBERule, len(derived))
	for _, rule := range derived {
		derivedByArea[rule.Area] = rule
	}
	requestedByArea := make(map[string]PersistenceRBERule, len(requested))
	for _, rule := range requested {
		expected, ok := derivedByArea[rule.Area]
		if !ok || rule != expected {
			return fmt.Errorf("persistence RBE rule for area %q is system-owned and cannot be modified", rule.Area)
		}
		requestedByArea[rule.Area] = rule
	}
	for _, rule := range derived {
		if _, ok := requestedByArea[rule.Area]; !ok {
			return fmt.Errorf("persistence RBE rule for area %q is system-owned and cannot be removed", rule.Area)
		}
	}
	return nil
}
