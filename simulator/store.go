package simulator

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"gopkg.in/yaml.v3"
)

const (
	// RelConfigDir is the simulator-owned configuration path below the data root.
	RelConfigDir  = "config/simulator"
	DevicesFile   = "devices.yaml"
	envDataDirKey = "OSJS_DATA_DIR"
)

// Store persists simulator-owned device definitions under the verified
// host-mounted configuration root. It never writes into MMA2 or OS.js source.
type Store struct {
	Root string // OSJS_DATA_DIR value or standalone Windows data root
}

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

func (s Store) SimulatorDir() string { return filepath.Join(s.Root, RelConfigDir) }
func (s Store) DevicesPath() string  { return filepath.Join(s.SimulatorDir(), DevicesFile) }

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

func (s Store) SaveOne(def DeviceDefinition) error {
	return s.SaveDocument(Document{Devices: []DeviceDefinition{def}})
}

// SaveDocument is the standalone document writer. Transaction owners call
// replace only AFTER their own validation and while holding the outer lock;
// reacquiring here would deadlock on Linux's non-reentrant flock.
func (s Store) SaveDocument(doc Document) error {
	for i := range doc.Devices {
		if err := ValidateDevice(doc.Devices[i]); err != nil {
			return err
		}
	}
	return s.withWriterLock(func() error { return s.replace(doc) })
}

// replace is private and lock-free; callers must hold the writer lock.
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
