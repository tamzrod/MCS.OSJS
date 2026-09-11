package replicator

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type persistedDocument struct {
	Devices []persistedDevice `yaml:"devices"`
}

type persistedDevice struct {
	Name        string               `yaml:"name"`
	Enabled     bool                 `yaml:"enabled"`
	Endpoint    string               `yaml:"endpoint"`
	UnitID      uint16               `yaml:"unit_id"`
	PullBlock   *PullBlock           `yaml:"pull_block"`
	Function    uint8                `yaml:"function"`
	Start       uint16               `yaml:"start"`
	Count       uint16               `yaml:"count"`
	ScanRateMS  uint32               `yaml:"scan_rate_ms"`
	Destination DestinationSelection `yaml:"destination"`
}

func (s Store) DocumentPath() string {
	return filepath.Join(s.ReplicatorDir(), DocumentFile)
}

// LoadDocument returns an empty document when the UI configuration has not yet
// been created. Existing pre-pull-block documents are migrated in memory and
// become canonical the next time SaveDocument succeeds.
func (s Store) LoadDocument() (Document, error) {
	path := s.DocumentPath()
	b, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return Document{Devices: []DeviceDefinition{}}, nil
		}
		return Document{}, err
	}
	var stored persistedDocument
	if err := yaml.Unmarshal(b, &stored); err != nil {
		return Document{}, fmt.Errorf("load %s: %w", path, err)
	}
	doc := Document{Devices: make([]DeviceDefinition, 0, len(stored.Devices))}
	for _, item := range stored.Devices {
		block := PullBlock{Function: item.Function, Start: item.Start, Count: item.Count, ScanRateMS: item.ScanRateMS}
		if item.PullBlock != nil {
			block = *item.PullBlock
		}
		doc.Devices = append(doc.Devices, DeviceDefinition{
			Name:        item.Name,
			Enabled:     item.Enabled,
			Endpoint:    item.Endpoint,
			UnitID:      item.UnitID,
			PullBlock:   block,
			Destination: item.Destination,
		})
	}
	return doc, nil
}

// SaveDocument validates and atomically persists the operator-facing document.
func (s Store) SaveDocument(doc Document) error {
	if err := ValidateDocument(doc); err != nil {
		return err
	}
	if err := os.MkdirAll(s.ReplicatorDir(), 0o755); err != nil {
		return err
	}
	b, err := yaml.Marshal(&doc)
	if err != nil {
		return err
	}
	final := s.DocumentPath()
	tmp := final + ".tmp"
	if err := os.WriteFile(tmp, b, 0o644); err != nil {
		return err
	}
	if err := os.Rename(tmp, final); err != nil {
		_ = os.Remove(tmp)
		return err
	}
	return nil
}
