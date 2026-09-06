package simulator

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

const (
	// RelConfigDir is the brainstorm conceptual layout under the verified
	// host-mounted data root: <OSJS_DATA_DIR>/config/simulator
	RelConfigDir  = "config/simulator"
	DevicesFile   = "devices.yaml"
	envDataDirKey = "OSJS_DATA_DIR"
)

// Store persists simulator-owned device definitions under the verified
// host-mounted configuration root (OSJS_DATA_DIR from deploy/docker-compose.yml
// and OSJS/src/server/config.js). It never writes into MMA2/ or OSJS package
// source trees.
type Store struct {
	Root string // OSJS_DATA_DIR value
}

// ConfigRootFromEnv returns the verified host-mounted data root.
// OSJS_DATA_DIR is the only established mount; this function does not invent
// a fallback host path when the variable is unset.
func ConfigRootFromEnv() (string, error) {
	root := os.Getenv(envDataDirKey)
	if root == "" {
		return "", fmt.Errorf("%s is unset; refusing to invent a host configuration path", envDataDirKey)
	}
	return root, nil
}

func (s Store) SimulatorDir() string {
	return filepath.Join(s.Root, RelConfigDir)
}

func (s Store) DevicesPath() string {
	return filepath.Join(s.SimulatorDir(), DevicesFile)
}

// Load reads the last valid persisted document. Missing file is an empty document.
func (s Store) Load() (Document, error) {
	path := s.DevicesPath()
	b, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return Document{}, nil
		}
		return Document{}, err
	}
	var doc Document
	if err := yaml.Unmarshal(b, &doc); err != nil {
		return Document{}, fmt.Errorf("load %s: %w", path, err)
	}
	return doc, nil
}

// SaveOne validates def and replaces the persisted document with that single
// definition. Invalid input leaves the previous file bytes unchanged.
func (s Store) SaveOne(def DeviceDefinition) error {
	if err := ValidateDevice(def); err != nil {
		return err
	}
	return s.replace(Document{Devices: []DeviceDefinition{def}})
}

func (s Store) replace(doc Document) error {
	if err := os.MkdirAll(s.SimulatorDir(), 0o755); err != nil {
		return err
	}
	b, err := yaml.Marshal(&doc)
	if err != nil {
		return err
	}
	final := s.DevicesPath()
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
