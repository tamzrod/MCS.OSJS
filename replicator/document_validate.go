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

func validatePullBlock(block PullBlock, index int) error {
	prefix := fmt.Sprintf("pull_blocks[%d]", index)
	if block.Function != 3 && block.Function != 4 {
		return fmt.Errorf("%s.function must be FC3 or FC4", prefix)
	}
	if block.Count == 0 {
		return fmt.Errorf("%s.count must be > 0", prefix)
	}
	if uint32(block.Start)+uint32(block.Count) > 0x10000 {
		return fmt.Errorf("%s start(%d)+count(%d) exceeds 16-bit address space", prefix, block.Start, block.Count)
	}
	if block.ScanRateMS == 0 {
		return fmt.Errorf("%s.scan_rate_ms must be > 0", prefix)
	}
	return nil
}

func ValidateDeviceDefinition(device DeviceDefinition) error {
	if strings.TrimSpace(device.Name) == "" {
		return fmt.Errorf("name is required")
	}
	blocks := device.blocks()
	if len(blocks) == 0 {
		return fmt.Errorf("at least one pull block is required")
	}
	for i, block := range blocks {
		if err := validatePullBlock(block, i); err != nil {
			return err
		}
	}
	if !device.Destination.AutoPort && device.Destination.Port == 0 {
		return fmt.Errorf("destination.port must be > 0 when automatic port allocation is disabled")
	}

	probe := device
	if probe.Destination.AutoPort && probe.Destination.Port == 0 {
		probe.Destination.Port = DefaultDestinationPort
	}
	if probe.Destination.AutoUnitID && probe.Destination.UnitID > 0xFF {
		probe.Destination.UnitID = DefaultDestinationUnit
	}
	for i, block := range blocks {
		cfg, err := probe.runtimeConfigForBlock(block)
		if err != nil {
			return err
		}
		if err := ValidateConfig(cfg); err != nil {
			return fmt.Errorf("pull_blocks[%d]: %w", i, err)
		}
	}
	return nil
}
