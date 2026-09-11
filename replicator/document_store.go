package replicator

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

func (s Store) DocumentPath() string {
	return filepath.Join(s.ReplicatorDir(), DocumentFile)
}

// LoadDocument returns an empty document when the UI configuration has not yet
// been created. Existing REP-003 config.yaml remains supported independently.
func (s Store) LoadDocument() (Document, error) {
	path := s.DocumentPath()
	b, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return Document{Devices: []DeviceDefinition{}}, nil
		}
		return Document{}, err
	}
	var doc Document
	if err := yaml.Unmarshal(b, &doc); err != nil {
		return Document{}, fmt.Errorf("load %s: %w", path, err)
	}
	if doc.Devices == nil {
		doc.Devices = []DeviceDefinition{}
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
