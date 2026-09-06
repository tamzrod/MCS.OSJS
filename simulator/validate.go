package simulator

import (
	"fmt"
	"strings"
)

// ValidateDevice checks field types and ranges from repository-established
// MMA2 structural limits and the approved simulator brainstorm.
//
// MMA2 truth (MMA2/internal/config/validate.go):
//   - port must be > 0
//   - unit_id must be <= 255
//   - unused FC area has count == 0
//   - start+count must not exceed the 16-bit address space (end exclusive 0x10000)
//
// Brainstorm truth (planning/Brainstorm/osjs-modbus-simulator.md):
//   - Randomize Every (ms) must be a positive integer for a configured FC
func ValidateDevice(d DeviceDefinition) error {
	if strings.TrimSpace(d.Name) == "" {
		return fmt.Errorf("name is required")
	}
	if d.MMA2.Port == 0 {
		return fmt.Errorf("mma2.port must be > 0")
	}
	if d.MMA2.UnitID > 0xFF {
		return fmt.Errorf("mma2.unit_id must be <= 255")
	}
	if err := validateArea("fc1", d.MMA2.FC1, d.RandomRuntime.FC1IntervalMS); err != nil {
		return err
	}
	if err := validateArea("fc2", d.MMA2.FC2, d.RandomRuntime.FC2IntervalMS); err != nil {
		return err
	}
	if err := validateArea("fc3", d.MMA2.FC3, d.RandomRuntime.FC3IntervalMS); err != nil {
		return err
	}
	if err := validateArea("fc4", d.MMA2.FC4, d.RandomRuntime.FC4IntervalMS); err != nil {
		return err
	}
	return nil
}

func validateArea(name string, a Area, intervalMS uint32) error {
	if a.Count == 0 {
		return nil
	}
	end := uint32(a.Start) + uint32(a.Count)
	if end > 0x10000 {
		return fmt.Errorf("mma2.%s: start(%d)+count(%d) exceeds 16-bit address space", name, a.Start, a.Count)
	}
	if intervalMS == 0 {
		return fmt.Errorf("random_runtime.%s_interval_ms must be > 0 when %s count > 0", name, name)
	}
	return nil
}
