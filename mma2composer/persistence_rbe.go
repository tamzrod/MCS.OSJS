package mma2composer

// PersistenceRBERule is one system-owned persistence RBE rule derived from a
// memory area. It has no ID, name or independent range: start/count are copied
// from the authoritative memory area, which remains the single source of truth.
// Global ID assignment, locking and user-rule collision policy are separate
// tasks and are deliberately not expressed here.
type PersistenceRBERule struct {
	Area  string `yaml:"area" json:"area"`
	Start uint16 `yaml:"start" json:"start"`
	Count uint16 `yaml:"count" json:"count"`
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
			Area:  entry.name,
			Start: entry.area.Start,
			Count: entry.area.Count,
		})
	}
	return rules
}
