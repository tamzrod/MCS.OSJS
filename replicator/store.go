package replicator

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"gopkg.in/yaml.v3"
)

const (
	RelConfigDir  = "config/replicator"
	ConfigFile    = "config.yaml"
	envDataDirKey = "OSJS_DATA_DIR"
)

// Store persists the Replicator-owned configuration under the shared data root.
type Store struct {
	Root string
}

// ConfigRootFromEnv returns the shared data root. OSJS_DATA_DIR remains the
// authoritative deployed-stack path. The standalone Windows package has one
// installer-owned equivalent under ProgramData, so Windows services can boot
// without NSSM-specific environment injection.
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

func (s Store) ReplicatorDir() string {
	return filepath.Join(s.Root, RelConfigDir)
}

func (s Store) ConfigPath() string {
	return filepath.Join(s.ReplicatorDir(), ConfigFile)
}

// Load reads and validates the last persisted configuration.
func (s Store) Load() (Config, error) {
	path := s.ConfigPath()
	b, err := os.ReadFile(path)
	if err != nil {
		return Config{}, err
	}
	var cfg Config
	if err := yaml.Unmarshal(b, &cfg); err != nil {
		return Config{}, fmt.Errorf("load %s: %w", path, err)
	}
	if err := ValidateConfig(cfg); err != nil {
		return Config{}, fmt.Errorf("load %s: %w", path, err)
	}
	return cfg, nil
}

// Save validates before touching disk and atomically replaces config.yaml.
func (s Store) Save(cfg Config) error {
	if err := ValidateConfig(cfg); err != nil {
		return err
	}
	if err := os.MkdirAll(s.ReplicatorDir(), 0o755); err != nil {
		return err
	}
	b, err := yaml.Marshal(&cfg)
	if err != nil {
		return err
	}
	final := s.ConfigPath()
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
