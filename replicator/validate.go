package replicator

import (
	"fmt"
	"strings"
)

// ValidateConfig rejects invalid required parameters before persistence.
func ValidateConfig(cfg Config) error {
	if strings.TrimSpace(cfg.Source.Host) == "" {
		return fmt.Errorf("source.host is required")
	}
	if cfg.Source.Port == 0 {
		return fmt.Errorf("source.port must be > 0")
	}
	if cfg.Source.UnitID > 0xFF {
		return fmt.Errorf("source.unit_id must be <= 255")
	}
	if cfg.Source.Function < 1 || cfg.Source.Function > 4 {
		return fmt.Errorf("source.function must be one of 1, 2, 3, or 4")
	}
	if cfg.Source.Count == 0 {
		return fmt.Errorf("source.count must be > 0")
	}
	if err := validateRange("source", cfg.Source.Start, cfg.Source.Count); err != nil {
		return err
	}
	if cfg.Source.PollIntervalMS == 0 {
		return fmt.Errorf("source.poll_interval_ms must be > 0")
	}

	if cfg.Destination.ListenerPort == 0 {
		return fmt.Errorf("destination.listener_port must be > 0")
	}
	if cfg.Destination.UnitID > 0xFF {
		return fmt.Errorf("destination.unit_id must be <= 255")
	}
	if !validArea(cfg.Destination.Area) {
		return fmt.Errorf("destination.area must be one of fc1, fc2, fc3, or fc4")
	}
	if cfg.Destination.Count == 0 {
		return fmt.Errorf("destination.count must be > 0")
	}
	if err := validateRange("destination", cfg.Destination.Start, cfg.Destination.Count); err != nil {
		return err
	}
	if cfg.Source.Count != cfg.Destination.Count {
		return fmt.Errorf("source.count and destination.count must match for 1:1 replication")
	}
	return nil
}

func validateRange(name string, start, count uint16) error {
	if uint32(start)+uint32(count) > 0x10000 {
		return fmt.Errorf("%s: start(%d)+count(%d) exceeds 16-bit address space", name, start, count)
	}
	return nil
}

func validArea(area string) bool {
	switch strings.ToLower(strings.TrimSpace(area)) {
	case "fc1", "fc2", "fc3", "fc4":
		return true
	default:
		return false
	}
}
