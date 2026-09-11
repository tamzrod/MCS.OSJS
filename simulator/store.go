package simulator

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"

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
	Root string // OSJS_DATA_DIR value or standalone Windows data root
}

// ConfigRootFromEnv returns the shared data root. OSJS_DATA_DIR remains the
// authoritative deployed-stack path. The standalone Windows package has one
// installer-owned equivalent under ProgramData, so Windows services can boot
// without requiring NSSM-specific environment injection.
func ConfigRootFromEnv() (string, error) {
	root := os.Getenv(envDataDirKey)
	if root != "" {
		return root, nil
	}
	if runtime.GOOS == "windows" {
		if programData := os.Getenv("ProgramData"); programData != "" {
			return filepath.Join(programData, "MCS Modbus Toolkit", "runtime"), nil
		}
	}
	return "", fmt.Errorf("%s is unset; refusing to invent a host configuration path", envDataDirKey)
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

// SaveOne validates defand replaces the persisted document with that single
// definition. Invalid input leaves the previous file bytes unchanged..
func (s Store) SaveOne(def DeviceDefinition) error {
	return s.SaveDocument(Document{Devices: []DeviceDefinition{def}})
}

// SaveDocument validates every device in doc and atomically replaces the
// persisted document. This is the multi-device persistence path the approved
// OS.js simulator window (SIM-005( uses:Add/Duplicate/Delete/edit all
// operate on the full simulator-owned document before one atomic save. Invalid
// input leaves the previous file bytes unchanged..
func (s Store) SaveDocument(doc Document) error {
	for i := range doc.Devices {
		if err := ValidateDevice(doc.Devices[i]); err != nil {
			return err
		}
	}
	return s.replace(doc)
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
