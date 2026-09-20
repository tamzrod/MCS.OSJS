package simulator

import (
	"encoding/json"
	"fmt"
	"strings"
)

// validateExtraRBE checks if the rbe field in Extra map contains valid JSON.
// Valid values are either a boolean or an object with string keys and boolean/object values.
func validateExtraRBE(extra map[string]interface{}) error {
	rbe, ok := extra["rbe"]
	if !ok {
		return nil // rbe not present is fine
	}

	// Must be a valid JSON type (map or bool)
	switch v := rbe.(type) {
	case map[string]interface{}:
		// Map type is acceptable - no further validation needed here
		return nil
	case bool:
		// Boolean type is also acceptable
		return nil
	default:
		// Invalid type (string with malformed value, etc.)
		b, _ := json.Marshal(v)
		return fmt.Errorf("corrupted_rbe: extra[%q]=%s must be bool or object", "rbe", string(b))
	}
}

// ValidateDevice checks the MMA2 memory structure independently from optional
// random-value generation. A zero random interval means None: memory remains
// allocated but this simulator does not schedule writes to that area.
func ValidateDevice(d DeviceDefinition) error {
	if strings.TrimSpace(d.Name) == "" {
		return fmt.Errorf("name is required")
	}

	// Validate Extra map contents (RBE field validation)
	if len(d.MMA2.Extra) > 0 {
		if err := validateExtraRBE(d.MMA2.Extra); err != nil {
			return err
		}
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
