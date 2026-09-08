// Package configvalidate exposes MMA2's authoritative configuration validation
// to local configuration composers without starting the MMA2 runtime.
package configvalidate

import (
	"fmt"

	"gopkg.in/yaml.v3"
	internalconfig "mma2/internal/config"
)

// YAML validates one complete MMA2 configuration document.
func YAML(data []byte) error {
	var cfg internalconfig.Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return fmt.Errorf("parse MMA2 config: %w", err)
	}
	return internalconfig.Validate(&cfg)
}
