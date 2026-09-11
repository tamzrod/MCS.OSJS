package replicator

import (
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

func validConfig() Config {
	return Config{
		Source: SourceConfig{
			Host:           "127.0.0.1",
			Port:           1502,
			UnitID:         1,
			Function:       3,
			Start:          100,
			Count:          8,
			PollIntervalMS: 250,
		},
		Destination: DestinationConfig{
			ListenerPort: 2502,
			UnitID:       2,
			Area:         "fc3",
			Start:        400,
			Count:        8,
		},
	}
}

func TestConfigRootFromEnv(t *testing.T) {
	t.Setenv(envDataDirKey, "")
	if _, err := ConfigRootFromEnv(); err == nil {
		t.Fatal("expected error when OSJS_DATA_DIR is unset")
	}

	root := t.TempDir()
	t.Setenv(envDataDirKey, root)
	got, err := ConfigRootFromEnv()
	if err != nil {
		t.Fatalf("ConfigRootFromEnv: %v", err)
	}
	if got != root {
		t.Fatalf("root = %q, want %q", got, root)
	}
}

func TestSaveLoadRoundTrip(t *testing.T) {
	root := t.TempDir()
	t.Setenv(envDataDirKey, root)
	resolved, err := ConfigRootFromEnv()
	if err != nil {
		t.Fatalf("ConfigRootFromEnv: %v", err)
	}
	store := Store{Root: resolved}
	want := validConfig()

	if err := store.Save(want); err != nil {
		t.Fatalf("Save: %v", err)
	}
	wantPath := filepath.Join(root, "config", "replicator", "config.yaml")
	if store.ConfigPath() != wantPath {
		t.Fatalf("ConfigPath = %q, want %q", store.ConfigPath(), wantPath)
	}
	if _, err := os.Stat(wantPath); err != nil {
		t.Fatalf("persisted config missing: %v", err)
	}

	got, err := store.Load()
	if err != nil {
		t.Fatalf("Load: %v", err)
	}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("round trip mismatch:\n got: %#v\nwant: %#v", got, want)
	}
}

func TestInvalidConfigRejectedBeforePersistence(t *testing.T) {
	store := Store{Root: t.TempDir()}
	cfg := validConfig()
	cfg.Source.Host = "   "

	if err := store.Save(cfg); err == nil {
		t.Fatal("expected validation error")
	}
	if _, err := os.Stat(store.ConfigPath()); !os.IsNotExist(err) {
		t.Fatalf("invalid config should not be persisted; stat err = %v", err)
	}
}

func TestInvalidReplacementLeavesPriorBytesUnchanged(t *testing.T) {
	store := Store{Root: t.TempDir()}
	if err := store.Save(validConfig()); err != nil {
		t.Fatalf("initial Save: %v", err)
	}
	before, err := os.ReadFile(store.ConfigPath())
	if err != nil {
		t.Fatalf("read before: %v", err)
	}

	bad := validConfig()
	bad.Destination.Count = bad.Source.Count - 1
	if err := store.Save(bad); err == nil {
		t.Fatal("expected 1:1 count validation error")
	}
	after, err := os.ReadFile(store.ConfigPath())
	if err != nil {
		t.Fatalf("read after: %v", err)
	}
	if string(after) != string(before) {
		t.Fatal("invalid replacement changed persisted bytes")
	}
}

func TestValidateConfigRejectsInvalidRangesAndRequiredValues(t *testing.T) {
	tests := []struct {
		name string
		edit func(*Config)
	}{
		{"source port zero", func(c *Config) { c.Source.Port = 0 }},
		{"source unit too high", func(c *Config) { c.Source.UnitID = 256 }},
		{"source function invalid", func(c *Config) { c.Source.Function = 5 }},
		{"source count zero", func(c *Config) { c.Source.Count = 0 }},
		{"source range overflow", func(c *Config) { c.Source.Start = 65535; c.Source.Count = 2; c.Destination.Count = 2 }},
		{"poll interval zero", func(c *Config) { c.Source.PollIntervalMS = 0 }},
		{"destination port zero", func(c *Config) { c.Destination.ListenerPort = 0 }},
		{"destination unit too high", func(c *Config) { c.Destination.UnitID = 256 }},
		{"destination area invalid", func(c *Config) { c.Destination.Area = "fc9" }},
		{"destination count zero", func(c *Config) { c.Destination.Count = 0 }},
		{"destination range overflow", func(c *Config) { c.Destination.Start = 65535; c.Destination.Count = 2; c.Source.Count = 2 }},
		{"count mismatch", func(c *Config) { c.Destination.Count = c.Source.Count + 1 }},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg := validConfig()
			tt.edit(&cfg)
			if err := ValidateConfig(cfg); err == nil {
				t.Fatal("expected validation error")
			}
		})
	}
}
