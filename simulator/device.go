package simulator

// DeviceDefinition is one simulator-owned persistent device.
// MMA2 structural parameters and random-runtime parameters are persisted
// together but remain separate domains with separate future consumers.
// This file is not MMA2 effective runtime configuration.
type DeviceDefinition struct {
	Name          string              `yaml:"name"`
	Enabled       bool                `yaml:"enabled"`
	MMA2          MMA2Params          `yaml:"mma2"`
	RandomRuntime RandomRuntimeParams `yaml:"random_runtime"`
}

// MMA2Params is the simulator's request for MMA2 listener/Unit/FC structure.
// SIM-001 persists this domain only; it does not activate MMA2.
type MMA2Params struct {
	Port   uint16 `yaml:"port"`
	UnitID uint16 `yaml:"unit_id"`
	FC1    Area   `yaml:"fc1"`
	FC2    Area   `yaml:"fc2"`
	FC3    Area   `yaml:"fc3"`
	FC4    Area   `yaml:"fc4"`
}

// Area is one function-code memory window. Count 0 means the area is unused.
type Area struct {
	Start uint16 `yaml:"start"`
	Count uint16 `yaml:"count"`
}

// RandomRuntimeParams is the simulator-owned per-FC randomization schedule.
type RandomRuntimeParams struct {
	FC1IntervalMS uint32 `yaml:"fc1_interval_ms"`
	FC2IntervalMS uint32 `yaml:"fc2_interval_ms"`
	FC3IntervalMS uint32 `yaml:"fc3_interval_ms"`
	FC4IntervalMS uint32 `yaml:"fc4_interval_ms"`
}

// Document is the on-disk simulator-owned store. It is not an MMA2 config file.
type Document struct {
	Devices []DeviceDefinition `yaml:"devices"`
}
