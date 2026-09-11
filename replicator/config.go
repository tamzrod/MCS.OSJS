package replicator

// Config is one Replicator-owned persisted definition. REP-003 intentionally
// supports exactly one external Modbus source range and one MMA2 destination
// range. Runtime connection and polling behavior belong to later tasks.
type Config struct {
	Source      SourceConfig      `yaml:"source" json:"source"`
	Destination DestinationConfig `yaml:"destination" json:"destination"`
}

// SourceConfig identifies one Modbus range to read later.
type SourceConfig struct {
	Host           string `yaml:"host" json:"host"`
	Port           uint16 `yaml:"port" json:"port"`
	UnitID         uint16 `yaml:"unit_id" json:"unit_id"`
	Function       uint8  `yaml:"function" json:"function"`
	Start          uint16 `yaml:"start" json:"start"`
	Count          uint16 `yaml:"count" json:"count"`
	PollIntervalMS uint32 `yaml:"poll_interval_ms" json:"poll_interval_ms"`
}

// DestinationConfig identifies the Replicator-owned MMA2 reservation that
// receives the source range in later replication tasks.
type DestinationConfig struct {
	ListenerPort uint16 `yaml:"listener_port" json:"listener_port"`
	UnitID       uint16 `yaml:"unit_id" json:"unit_id"`
	Area         string `yaml:"area" json:"area"`
	Start        uint16 `yaml:"start" json:"start"`
	Count        uint16 `yaml:"count" json:"count"`
}
