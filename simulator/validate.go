package simulator

import (
	"fmt"
	"strings"
)

// ValidateDevice checks the MMA2 memory structure independently from optional
// random-value generation. A zero random interval means None: memory remains
// allocated but this simulator does not schedule writes to that area.
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

// Keep the interval parameter for existing callers. Zero is valid regardless
// of Count; NewScheduler/UpdateTiming schedule only positive intervals.
func validateArea(name string, a Area, _ uint32) error {
	if a.Count == 0 {
		return nil
	}
	end := uint32(a.Start) + uint32(a.Count)
	if end > 0x10000 {
		return fmt.Errorf("mma2.%s: start(%d)+count(%d) exceeds 16-bit address space", name, a.Start, a.Count)
	}
	return nil
}
