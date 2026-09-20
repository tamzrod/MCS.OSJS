package replicator

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type persistedDocument struct { Devices []persistedDevice `yaml:"devices"` }
type persistedDevice struct {
	MMA2Advanced map[string]interface{} `yaml:"mma2_advanced"`
	Name string `yaml:"name"`
	Enabled bool `yaml:"enabled"`
	Endpoint string `yaml:"endpoint"`
	UnitID uint16 `yaml:"unit_id"`
	PullBlocks []PullBlock `yaml:"pull_blocks"`
	PullBlock *PullBlock `yaml:"pull_block"`
	Function uint8 `yaml:"function"`
	Start uint16 `yaml:"start"`
	Count uint16 `yaml:"count"`
	ScanRateMS uint32 `yaml:"scan_rate_ms"`
	Destination DestinationSelection `yaml:"destination"`
}

func (s Store) DocumentPath() string { return filepath.Join(s.ReplicatorDir(), DocumentFile) }

func (s Store) LoadDocument() (Document, error) {
	path := s.DocumentPath()
	b, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) { return Document{Devices: []DeviceDefinition{}}, nil }
		return Document{}, err
	}
	var stored persistedDocument
	if err := yaml.Unmarshal(b, &stored); err != nil { return Document{}, fmt.Errorf("load %s: %w", path, err) }
	doc := Document{Devices: make([]DeviceDefinition, 0, len(stored.Devices))}
	for _, item := range stored.Devices {
		blocks := append([]PullBlock(nil), item.PullBlocks...)
		if len(blocks) == 0 && item.PullBlock != nil { blocks = []PullBlock{*item.PullBlock} }
		if len(blocks) == 0 && (item.Function != 0 || item.Count != 0 || item.ScanRateMS != 0) {
			blocks = []PullBlock{{Function: item.Function, Start: item.Start, Count: item.Count, ScanRateMS: item.ScanRateMS}}
		}
		device := DeviceDefinition{
			MMA2Advanced: item.MMA2Advanced, Name: item.Name, Enabled: item.Enabled,
			Endpoint: item.Endpoint, UnitID: item.UnitID, PullBlocks: blocks, Destination: item.Destination,
		}
		if len(blocks) > 0 { device.PullBlock = blocks[0] }
		doc.Devices = append(doc.Devices, device)
	}
	return doc, nil
}

func canonicalDocument(doc Document) Document {
	out := cloneDocument(doc)
	for i := range out.Devices {
		device := &out.Devices[i]
		if len(device.PullBlocks) == 0 { device.PullBlocks = append([]PullBlock(nil), device.blocks()...) }
		device.PullBlock = PullBlock{}
	}
	return out
}

// SaveDocument is a standalone entry point. Manager Boot/Apply call the inner
// helper while holding the SAME lock for composition, restart and persistence.
func (s Store) SaveDocument(doc Document) error {
	return s.withWriterLock(func() error { return s.saveDocumentLocked(doc) })
}

// saveDocumentLocked is not independently callable across process boundaries.
func (s Store) saveDocumentLocked(doc Document) error {
	canonical := canonicalDocument(doc)
	if err := ValidateDocument(canonical); err != nil { return err }
	if err := os.MkdirAll(s.ReplicatorDir(), 0o755); err != nil { return err }
	b, err := yaml.Marshal(&canonical)
	if err != nil { return err }
	final := s.DocumentPath()
	tmp := final + ".tmp"
	if err := os.WriteFile(tmp, b, 0o644); err != nil { return err }
	if err := os.Rename(tmp, final); err != nil { _ = os.Remove(tmp); return err }
	return nil
}
