package replicator

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

const (
	RelConfigDir  = "config/replicator"
	ConfigFile    = "config.yaml"
	envDataDirKey = "OSJS_DATA_DIR"
)

// Store persists the Replicator-owned configuration under OSJS_DATA_DIR.
type Store struct {
	Root string
}

// ConfigRootFromEnv returns the established shared data root. No fallback host
// path is invented when OSJS_DATA_DIR is unset.
func ConfigRootFromEnv() (string, error) {
	root := os.Getenv(envDataDirKey)
	if root == "" {
		return "", fmt.Errorf("%s is unset; refusing to invent a host configuration path", envDataDirKey)
	}
	return root, nil
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
