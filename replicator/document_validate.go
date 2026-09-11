package replicator

import (
	"fmt"
	"strings"
)

func ValidateDocument(doc Document) error {
	names := make(map[string]bool, len(doc.Devices))
	for i, device := range doc.Devices {
		if err := ValidateDeviceDefinition(device); err != nil {
			return fmt.Errorf("device %d: %w", i, err)
		}
		key := strings.ToLower(strings.TrimSpace(device.Name))
		if names[key] {
			return fmt.Errorf("device %d: duplicate name %q", i, device.Name)
		}
		names[key] = true
	}
	return nil
}

func ValidateDeviceDefinition(device DeviceDefinition) error {
	if strings.TrimSpace(device.Name) == "" {
		return fmt.Errorf("name is required")
	}
	if device.Function != 3 && device.Function != 4 {
		return fmt.Errorf("function must be FC3 or FC4")
	}
	if device.Count == 0 {
		return fmt.Errorf("count must be > 0")
	}
	if uint32(device.Start)+uint32(device.Count) > 0x10000 {
		return fmt.Errorf("start(%d)+count(%d) exceeds 16-bit address space", device.Start, device.Count)
	}
	if device.ScanRateMS == 0 {
		return fmt.Errorf("scan_rate_ms must be > 0")
	}
	if !device.Destination.AutoPort && device.Destination.Port == 0 {
		return fmt.Errorf("destination.port must be > 0 when automatic port allocation is disabled")
	}

	// Endpoint parsing and all runtime mapping invariants are centralized by the
	// existing Config validator. Automatic destination fields receive temporary
	// valid placeholders until ResolveDocumentDestinations assigns real values.
	probe := device
	if probe.Destination.AutoPort && probe.Destination.Port == 0 {
		probe.Destination.Port = DefaultDestinationPort
	}
	if probe.Destination.AutoUnitID && probe.Destination.UnitID > 0xFF {
		probe.Destination.UnitID = DefaultDestinationUnit
	}
	cfg, err := probe.runtimeConfig()
	if err != nil {
		return err
	}
	return ValidateConfig(cfg)
}
