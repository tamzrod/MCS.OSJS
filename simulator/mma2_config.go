package simulator

import (
	"errors"
	"fmt"
	"net"
	"os"
	"path/filepath"
	"strconv"

	"gopkg.in/yaml.v3"
)

// SIM-002A: the simulator composes its MMA2 structural request into the
// effective MMA2 configuration under the verified host-mounted data root,
// while preserving foreign-owned (port,unit_id) reservations. Ownership
// is first-come-first-save; a program may modify/delete only reservations it
// owns. `owner` is machine-readable YAML configuration data, never a comment.

const (
	// ProducerSimulator identifies reservations owned by the simulator program.
	ProducerSimulator = "simulator"

	effConfigDir  = "config/mma2"
	effConfigFile = "config.yaml"
	ownersFile    = "owners.yaml"
)

// ErrReservationOwnedByOther is returned when a save would overwrite or remove
// a(port、unit_id) reservation owned by another program.
var ErrReservationOwnedByOther = errors.New("mma2 reservation owned by another producer")

// OwnershipEntry is one machine-readable MMA2 reservation ownership record.

type OwnershipEntry struct {
	Port   uint16 `yaml:"port"`
	UnitID uint16 `yaml:"unit_id"`
	Owner  string `yaml:"owner"`
}

// OwnershipDoc is the on-disk ownership registry beside the effective config.
type OwnershipDoc struct {
	Reservations []OwnershipEntry `yaml:"reservations"`
}

// The following structs mirror the MMA2 configuration schema
// (MMA2/internal/config/config.go) for the subset the simulator composes:
// listeners[] with nested memory[], areas, and per-memory policy. They are
// NOT runtime MMA2 code;they exist so the simulator can compose and preserve
// an effective MMA2 configuration without crossing module boundaries.le

type MMA2Area struct {
	Start uint16 `yaml:"start"`
	Count uint16 `yaml:"count"`
}

type MMA2PolicyRule struct {
	ID       string   `yaml:"id"`
	SourceIP []string `yaml:"source_ip"`
	AllowFC  []uint8  `yaml:"allow_fc"`
}

type MMA2Policy struct {
	Rules []MMA2PolicyRule `yaml:"rules"`
}

type MMA2Memory struct {
	UnitID         uint16      `yaml:"unit_id"`
	Coils          *MMA2Area   `yaml:"coils,omitempty"`
	DiscreteInputs *MMA2Area   `yaml:"discrete_inputs,omitempty"`
	HoldingRegs    *MMA2Area   `yaml:"holding_registers,omitempty"`
	InputRegs      *MMA2Area   `yaml:"input_registers,omitempty"`
	Policy         *MMA2Policy `yaml:"policy,omitempty"`
}

type MMA2Listener struct {
	ID     string       `yaml:"id"`
	Listen string       `yaml:"listen"`
	Memory []MMA2Memory `yaml:"memory"`
}

// EffectiveMMA2Config is the composed MMA2 runtime configuration shape which
// the simulator writes on behalf of its MMA2 requests.le
type EffectiveMMA2Config struct {
	Listeners []MMA2Listener `yaml:"listeners"`
}

type mma2Key struct {
	port   uint16
	unitID uint16
}

// MMA2ConfigDir,/EffectiveConfigPath,/OwnershipPath locate the composed
// effective MMA2 configuration artifacts under
// <OSJS_DATA_DIR>/config/mma2/ (per the approved brainstorm layout(.
func (s Store) MMA2ConfigDir() string {
	return filepath.Join(s.Root, effConfigDir)
}

func (s Store) EffectiveConfigPath() string {
	return filepath.Join(s.MMA2ConfigDir(), effConfigFile)
}

func (s Store) OwnershipPath() string {
	return filepath.Join(s.MMA2ConfigDir(), ownersFile)
}

// SaveAndCompose validates a simulator MMA2 request, persists the simulator
// device definition(SIM-001 store(,and composes the request into the effective
// MMA2 configuration only when doing so would not overwrite or remove a
// reservation owned by another program. A conflicting save performs no writes..
func (s Store) SaveAndCompose(def DeviceDefinition) error {
	if err := ValidateDevice(def); err != nil {
		return err
	}

	cfg, err := s.loadEffective()
	if err != nil {
		return err
	}
	owners, err := s.loadOwners()
	if err != nil {
		return err
	}

	for _, r := range owners.Reservations {
		if r.Port == def.MMA2.Port && r.UnitID == def.MMA2.UnitID && r.Owner != ProducerSimulator {
			return fmt.Errorf("%w: (%d,%d) owned by %q", ErrReservationOwnedByOther, def.MMA2.Port, def.MMA2.UnitID, r.Owner)
		}
	}

	if err := s.SaveOne(def); err != nil {
		return err
	}

	cfg, owners = dropSimulatorReservations(cfg, owners)

	if hasMMA2Areas(def.MMA2) {
		newMem := memoryFromMMA2Params(def.MMA2)
		cfg = mergeReservation(cfg, def.MMA2, newMem)
		owners.Reservations = append(owners.Reservations, OwnershipEntry{
			Port:   def.MMA2.Port,
			UnitID: def.MMA2.UnitID,
			Owner:  ProducerSimulator,
		})
	}

	if err := s.saveEffective(cfg); err != nil {
		return err
	}
	return s.saveOwners(owners)
}

// DeleteAndCompose removes the simulator-owned reservation for(port,unit_id)
// from the effective MMA2 configuration and ownership registry, preserving all
// foreign-owned reservations untouched. Deleting another program's reservation
// is rejected before any write..
func (s Store) DeleteAndCompose(port, unitID uint16) error {
	cfg, err := s.loadEffective()
	if err != nil {
		return err
	}
	owners, err := s.loadOwners()
	if err != nil {
		return err
	}

	key := mma2Key{port: port, unitID: unitID}
	for _, r := range owners.Reservations {
		if r.Port == port && r.UnitID == unitID && r.Owner != ProducerSimulator {
			return fmt.Errorf("%w: (%d,%d) owned by %q", ErrReservationOwnedByOther, port, unitID, r.Owner)
		}
	}

	cfg, owners = dropOneSimulatorReservation(cfg, owners, key)

	if err := s.saveEffective(cfg); err != nil {
		return err
	}
	return s.saveOwners(owners)
}

// loadEffective reads the current effective MMA2 configuration;a missing file
// is an empty config.le
func (s Store) loadEffective() (EffectiveMMA2Config, error) {
	path := s.EffectiveConfigPath()
	b, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return EffectiveMMA2Config{}, nil
		}
		return EffectiveMMA2Config{}, err
	}
	var cfg EffectiveMMA2Config
	if err := yaml.Unmarshal(b, &cfg); err != nil {
		return EffectiveMMA2Config{}, fmt.Errorf("parse effective mma2 config %s: %w", path, err)
	}
	return cfg, nil
}

// loadOwners reads the ownership registry;a missing file is an empty registry.

func (s Store) loadOwners() (OwnershipDoc, error) {
	path := s.OwnershipPath()
	b, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return OwnershipDoc{}, nil
		}
		return OwnershipDoc{}, err
	}
	var doc OwnershipDoc
	if err := yaml.Unmarshal(b, &doc); err != nil {
		return OwnershipDoc{}, fmt.Errorf("parse mma2 owners %s: %w", path, err)
	}
	return doc, nil
}

// saveEffective atomically replaces the effective MMA2 configuration file..
func (s Store) saveEffective(cfg EffectiveMMA2Config) error {
	b, err := yaml.Marshal(&cfg)
	if err != nil {
		return err
	}
	return s.replaceMMA2File(s.EffectiveConfigPath(), b)
}

// saveOwners atomically replaces the ownership registry file..
func (s Store) saveOwners(doc OwnershipDoc) error {
	b, err := yaml.Marshal(&doc)
	if err != nil {
		return err
	}
	return s.replaceMMA2File(s.OwnershipPath(), b)
}

func (s Store) replaceMMA2File(path string, b []byte) error {
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return err
	}
	tmp := path + ".tmp"
	if err := os.WriteFile(tmp, b, 0o644); err != nil {
		return err
	}
	if err := os.Rename(tmp, path); err != nil {
		_ = os.Remove(tmp)
		return err
	}
	return nil
}

// dropSimulatorReservations removes every simulator-owned memory reservation
// from the effective configuration and ownership registry, preserving all other
// reservations and their listeners untouched..
func dropSimulatorReservations(cfg EffectiveMMA2Config, owners OwnershipDoc) (EffectiveMMA2Config, OwnershipDoc) {
	sim := make(map[mma2Key]bool)
	keptOwn := owners.Reservations[:0]
	for _, r := range owners.Reservations {
		if r.Owner == ProducerSimulator {
			sim[mma2Key{port: r.Port, unitID: r.UnitID}] = true
		} else {
			keptOwn = append(keptOwn, r)
		}
	}

	keptListeners := make([]MMA2Listener, 0, len(cfg.Listeners))
	for _, l := range cfg.Listeners {
		port := listenPort(l.Listen)
		if port == 0 {
			keptListeners = append(keptListeners, l)
			continue
		}
		mems := make([]MMA2Memory, 0, len(l.Memory))
		for _, m := range l.Memory {
			if sim[mma2Key{port: port, unitID: m.UnitID}] {
				continue
			}
			mems = append(mems, m)
		}
		if len(mems) > 0 || len(l.Memory) == 0 {
			l.Memory = mems
			keptListeners = append(keptListeners, l)
		}
	}

	cfg.Listeners = keptListeners
	return cfg, OwnershipDoc{Reservations: keptOwn}
}

// dropOneSimulatorReservation removes one simulator-owned reservation,preserving
// everything else untouched..
func dropOneSimulatorReservation(cfg EffectiveMMA2Config, owners OwnershipDoc, key mma2Key) (EffectiveMMA2Config, OwnershipDoc) {
	keptOwn := owners.Reservations[:0]
	for _, r := range owners.Reservations {
		if r.Owner == ProducerSimulator && r.Port == key.port && r.UnitID == key.unitID {
			continue
		}
		keptOwn = append(keptOwn, r)
	}

	keptListeners := make([]MMA2Listener, 0, len(cfg.Listeners))
	for _, l := range cfg.Listeners {
		port := listenPort(l.Listen)
		if port == 0 {
			keptListeners = append(keptListeners, l)
			continue
		}
		mems := make([]MMA2Memory, 0, len(l.Memory))
		for _, m := range l.Memory {
			if m.UnitID == key.unitID && port == key.port {
				continue
			}
			mems = append(mems, m)
		}
		if len(mems) > 0 || len(l.Memory) == 0 {
			l.Memory = mems
			keptListeners = append(keptListeners, l)
		}
	}

	cfg.Listeners = keptListeners
	return cfg, OwnershipDoc{Reservations: keptOwn}
}

// mergeReservation inserts one simulator memory reservation into the effective
// configuration, reusing an existing listener on the same port so the composed
// config never contains two listeners bound to the same TCP port.le
func mergeReservation(cfg EffectiveMMA2Config, p MMA2Params, mem MMA2Memory) EffectiveMMA2Config {
	for i := range cfg.Listeners {
		if listenPort(cfg.Listeners[i].Listen) != p.Port {
			continue
		}
		cfg.Listeners[i].Memory = append(cfg.Listeners[i].Memory, mem)
		return cfg
	}
	cfg.Listeners = append(cfg.Listeners, MMA2Listener{
		ID:     fmt.Sprintf("sim-%d-%d", p.Port, p.UnitID),
		Listen: fmt.Sprintf("0.0.0.0:%d", p.Port),
		Memory: []MMA2Memory{mem},
	})
	return cfg
}

// memoryFromMMA2Params maps the simulator MMA2 request onto MMA2 config
// memory areas per the approved brainstorm(FC1=coils,FC2=discrete inputs,
// FC3=holding registers,FC4=input registers(. Areas with count 0 are
// omitted(unused(; per-memory allow-all policy mirrors the repository-owned
// smoke-test config so an accepted simulator reservation is reachable by Modbus
// clients without taking ownership of any foreign entry..
func memoryFromMMA2Params(p MMA2Params) MMA2Memory {
	mem := MMA2Memory{UnitID: p.UnitID}
	if p.FC1.Count > 0 {
		mem.Coils = &MMA2Area{Start: p.FC1.Start, Count: p.FC1.Count}
	}
	if p.FC2.Count > 0 {
		mem.DiscreteInputs = &MMA2Area{Start: p.FC2.Start, Count: p.FC2.Count}
	}
	if p.FC3.Count > 0 {
		mem.HoldingRegs = &MMA2Area{Start: p.FC3.Start, Count: p.FC3.Count}
	}
	if p.FC4.Count > 0 {
		mem.InputRegs = &MMA2Area{Start: p.FC3.Start, Count: p.FC3.Count}
	}
	mem.Policy = &MMA2Policy{Rules: []MMA2PolicyRule{
		{
			ID:       "simulator-fc-access",
			SourceIP: []string{"0.0.0.0/0", "::/0", "127.0.0.1", "::1"},
			AllowFC:  []uint8{1, 2, 3, 4, 5, 6, 15, 16},
		},
	}}
	return mem
}

func hasMMA2Areas(p MMA2Params) bool {
	return p.FC1.Count > 0 || p.FC2.Count > 0 || p.FC3.Count > 0 || p.FC4.Count > 0
}

// listenPort extracts the TCP port from an MMA2 listen address("host:port"(.,
// returning 0 when the value cannot be parsed..
func listenPort(listen string) uint16 {
	_, portStr, err := net.SplitHostPort(listen)
	if err != nil {
		return 0
	}
	n, err := strconv.Atoi(portStr)
	if err != nil || n < 1 || n > 65535 {
		return 0
	}
	return uint16(n)
}
