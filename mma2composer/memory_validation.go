package mma2composer

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"mma2/pkg/configvalidate"
)

// ValidateFC43 supplements the bundled validator, which predates per-memory
// identification. Values are measured in bytes and never normalized.
func ValidateFC43(value interface{}) error {
	if value == nil {
		return nil
	}
	fields, ok := value.(map[string]interface{})
	if !ok {
		return fmt.Errorf("Device Identification: fc43 must be an object")
	}
	for _, field := range []struct{ key, label string }{
		{"vendor_name", "Vendor Name"}, {"product_code", "Product Code"}, {"major_minor_revision", "Major / Minor Revision"},
	} {
		raw, exists := fields[field.key]
		if !exists {
			continue
		}
		text, ok := raw.(string)
		valid := ok && len(text) >= 1 && len(text) <= 244
		for i := 0; i < len(text); i++ {
			if text[i] > 127 {
				valid = false
			}
		}
		if !valid {
			return fmt.Errorf("Device Identification: %s must contain 1–244 ASCII bytes", field.label)
		}
	}
	return nil
}

// ValidateMemory checks editor settings even for disabled memories. Commit
// separately checks the real shared output. Persistence and State Sealing
// validate independently: a persistence-enabled memory does not require State
// Sealing, and validation never enables, mutates or inspects sealing on
// persistence's behalf.
func ValidateMemory(memory Memory) error {
	if err := ValidateMemoryPersistence(memory); err != nil {
		return err
	}
	if err := ValidateFC43(memory.Extra["fc43"]); err != nil {
		return err
	}
	cfg := EffectiveConfig{Listeners: []Listener{{ID: "validation", Listen: "0.0.0.0:502", Memory: []Memory{memory}}},
		Extra: map[string]interface{}{"rbe": map[string]interface{}{"tcp": map[string]interface{}{"listen": "127.0.0.1:9001"}}}}
	data, err := yaml.Marshal(cfg)
	if err != nil {
		return err
	}
	return configvalidate.YAML(data)
}
