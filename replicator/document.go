package replicator

import (
	"fmt"
	"net"
	"strconv"
	"strings"
)

const (
	DocumentFile           = "devices.yaml"
	DefaultDestinationPort = uint16(5021)
	DefaultDestinationUnit = uint16(1)
)

// Document is the human-facing Replicator configuration used by the OS.js UI.
// Each device is one independent source -> MMA2 replication pipeline.
type Document struct {
	Devices []DeviceDefinition `yaml:"devices" json:"devices"`
}

// PullBlock is the one explicit Modbus range a Replicator device polls in this
// milestone. The model is intentionally a block even while only one block per
// device is supported, so runtime reads are never inferred from loose fields.
type PullBlock struct {
	Function   uint8  `yaml:"function" json:"function"`
	Start      uint16 `yaml:"start" json:"start"`
	Count      uint16 `yaml:"count" json:"count"`
	ScanRateMS uint32 `yaml:"scan_rate_ms" json:"scan_rate_ms"`
}

// DeviceDefinition keeps endpoint identity, one explicit source pull block,
// and the resolved MMA2 destination reservation chosen by the allocator.
type DeviceDefinition struct {
	Name        string               `yaml:"name" json:"name"`
	Enabled     bool                 `yaml:"enabled" json:"enabled"`
	Endpoint    string               `yaml:"endpoint" json:"endpoint"`
	UnitID      uint16               `yaml:"unit_id" json:"unit_id"`
	PullBlock   PullBlock            `yaml:"pull_block" json:"pull_block"`
	Destination DestinationSelection `yaml:"destination" json:"destination"`
}

// DestinationSelection stores both the resolved reservation and whether the
// operator wants either dimension to keep following automatic allocation.
type DestinationSelection struct {
	Port       uint16 `yaml:"port" json:"port"`
	UnitID     uint16 `yaml:"unit_id" json:"unit_id"`
	AutoPort   bool   `yaml:"auto_port" json:"auto_port"`
	AutoUnitID bool   `yaml:"auto_unit_id" json:"auto_unit_id"`
	Owner      string `yaml:"-" json:"owner,omitempty"`
	Status     string `yaml:"-" json:"status,omitempty"`
}

func (d DeviceDefinition) runtimeConfig() (Config, error) {
	host, portString, err := net.SplitHostPort(strings.TrimSpace(d.Endpoint))
	if err != nil {
		return Config{}, fmt.Errorf("endpoint must be host:port: %w", err)
	}
	if strings.TrimSpace(host) == "" {
		return Config{}, fmt.Errorf("endpoint host is required")
	}
	portNumber, err := strconv.Atoi(portString)
	if err != nil || portNumber < 1 || portNumber > 65535 {
		return Config{}, fmt.Errorf("endpoint port must be between 1 and 65535")
	}
	block := d.PullBlock
	area := fmt.Sprintf("fc%d", block.Function)
	cfg := Config{
		Source: SourceConfig{
			Host:           host,
			Port:           uint16(portNumber),
			UnitID:         d.UnitID,
			Function:       block.Function,
			Start:          block.Start,
			Count:          block.Count,
			PollIntervalMS: block.ScanRateMS,
		},
		Destination: DestinationConfig{
			ListenerPort: d.Destination.Port,
			UnitID:       d.Destination.UnitID,
			Area:         area,
			Start:        block.Start,
			Count:        block.Count,
		},
	}
	return cfg, nil
}
