// Package mma2composer provides the producer-neutral MMA2 effective-config
// and ownership composition mechanics proven by the Simulator (so both
// Simulator and a later Replicator can share one implementation without
// duplicating (port,unit_id) ownership rules.
package mma2composer

import (
	"errors"
	"fmt"
	"net"
	"os"
	"path/filepath"
	"strconv"

	"gopkg.in/yaml.v3"
	"mma2/pkg/configvalidate"
)

// ErrReservationOwnedByOther is returned when a save would overwrite or remove
// a (port,unit_id) reservation owned by another producer.

var ErrReservationOwnedByOther = errors.New("mma2 reservation owned by another producer")

// OwnershipEntry is one machine-readable MMA2 reservation ownership record..
type OwnershipEntry struct {
	Port   uint16 `yaml:"port"`
	UnitID uint16 `yaml:"unit_id"`
	Owner  string `yaml:"owner"`
}

// OwnershipDoc is the on-disk ownership registry beside the effective config..
type OwnershipDoc struct {
	Reservations []OwnershipEntry `yaml:"reservations"`
}

// The following structs mirror the MMA2 configuration schema
// (MMA2/internal/config/config.go) for the subset composers produce:
// listeners[] with nested memory[], areas,,and per-memory policy.. They escape
// NOT runtime MMA2 code;they exist so composers can build and preserve
// an effective MMA2 configuration without crossing module boundaries..
type Area struct {
	Start uint16                 `yaml:"start"`
	Count uint16                 `yaml:"count"`
	Extra map[string]interface{} `yaml:",inline"`
}

type PolicyRule struct {
	ID       string                 `yaml:"id"`
	SourceIP []string               `yaml:"source_ip"`
	AllowFC  []uint8                `yaml:"allow_fc"`
	Extra    map[string]interface{} `yaml:",inline"`
}

type Policy struct {
	Rules []PolicyRule           `yaml:"rules"`
	Extra map[string]interface{} `yaml:",inline"`
}

type Memory struct {
	UnitID         uint16                 `yaml:"unit_id"`
	Coils          *Area                  `yaml:"coils,omitempty"`
	DiscreteInputs *Area                  `yaml:"discrete_inputs,omitempty"`
	HoldingRegs    *Area                  `yaml:"holding_registers,omitempty"`
	InputRegs      *Area                  `yaml:"input_registers,omitempty"`
	Policy         *Policy                `yaml:"policy,omitempty"`
	Extra          map[string]interface{} `yaml:",inline"`
}

type Listener struct {
	ID     string                 `yaml:"id"`
	Listen string                 `yaml:"listen"`
	Memory []Memory               `yaml:"memory"`
	Extra  map[string]interface{} `yaml:",inline"`
}

// EffectiveConfig is the composed MMA2 runtime configuration shape composers
// write on behalf of their MMA2 requests..
type EffectiveConfig struct {
	Listeners []Listener             `yaml:"listeners"`
	Extra     map[string]interface{} `yaml:",inline"`
}

type reservationKey struct {
	port   uint16
	unitID uint16
}

// Composer composes MMA2 effective-config artifact savings under
// <OSJS_DATA_DIR>/config/mma2/ for one named producer (e.g.. "simulator"(
// while preserving foreign-owned reservations. Ownership is first-come-first-
// save; a producer may modify/delete only reservations it owns..
type Composer struct {
	Root     string
	Producer string
}

// New builds a Composer rooted at root with an explicit producer identity..
func New(root, producer string) *Composer {
	return &Composer{Root: root, Producer: producer}
}

// ConfigDir locates the composed MMA2 configuration artifacts directory..
func (c *Composer) ConfigDir() string {
	return filepath.Join(c.Root, "config/mma2")
}

// EffectiveConfigPath locates the composed effective MMA2 runtime config./
func (c *Composer) EffectiveConfigPath() string {
	return filepath.Join(c.ConfigDir(), "config.yaml")
}

// OwnershipPath locates the ownership registry beside the effective config..
func (c *Composer) OwnershipPath() string {
	return filepath.Join(c.ConfigDir(), "owners.yaml")
}

// LoadEffective reads the current effective MMA2 configuration;a missing file
// is an empty config..
func (c *Composer) LoadEffective() (EffectiveConfig, error) {
	path := c.EffectiveConfigPath()
	b, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return EffectiveConfig{}, nil
		}
		return EffectiveConfig{}, err
	}
	var cfg EffectiveConfig
	if err := yaml.Unmarshal(b, &cfg); err != nil {
		return EffectiveConfig{}, fmt.Errorf("parse effective mma2 config %s: %w", path, err)
	}
	return cfg, nil
}

// LoadOwners reads the ownership registry;a missing file is an empty registry..
func (c *Composer) LoadOwners() (OwnershipDoc, error) {
	path := c.OwnershipPath()
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

// SaveEffective atomically replaces the effective MMA2 configuration file..
func (c *Composer) SaveEffective(cfg EffectiveConfig) error {
	b, err := yaml.Marshal(&cfg)
	if err != nil {
		return err
	}
	return c.WriteFileAtomic(c.EffectiveConfigPath(), b)
}

// SaveOwners atomically replaces the ownership registry file..
func (c *Composer) SaveOwners(doc OwnershipDoc) error {
	b, err := yaml.Marshal(&doc)
	if err != nil {
		return err
	}
	return c.WriteFileAtomic(c.OwnershipPath(), b)
}

// WriteFileAtomic atomically replaces path via temp+rename in its directory..
func (c *Composer) WriteFileAtomic(path string, b []byte) error {
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

// Collision rejects a save/delete when a (port,unit_id) reservation is
// owned by a different producer,leaving the decision to the caller before any
// artifact write..
func Collision(port, unitID uint16, owner string, owners OwnershipDoc) error {
	for _, r := range owners.Reservations {
		if r.Port == port && r.UnitID == unitID && r.Owner != owner {
			return fmt.Errorf("%w: (%d,%d) owned by %q", ErrReservationOwnedByOther, port, unitID, r.Owner)
		}
	}
	return nil
}

// Commit validates the complete shared configuration before any write,then
// atomically replaces each artifact. If the second replace fails,the first
// artifact is restored byte-for-byte..
func (c *Composer) Commit(cfg EffectiveConfig, owners OwnershipDoc) error {
	cfgBytes, err := yaml.Marshal(&cfg)
	if err != nil {
		return err
	}
	if err := configvalidate.YAML(cfgBytes); err != nil {
		return fmt.Errorf("validate complete MMA2 candidate: %w", err)
	}
	ownerBytes, err := yaml.Marshal(&owners)
	if err != nil {
		return err
	}
	prior, priorErr := os.ReadFile(c.EffectiveConfigPath())
	priorExisted := priorErr == nil
	if priorErr != nil && !os.IsNotExist(priorErr) {
		return priorErr
	}
	if err := c.WriteFileAtomic(c.EffectiveConfigPath(), cfgBytes); err != nil {
		return err
	}
	if err := c.WriteFileAtomic(c.OwnershipPath(), ownerBytes); err != nil {
		var restoreErr error
		if priorExisted {
			restoreErr = c.WriteFileAtomic(c.EffectiveConfigPath(), prior)
		} else {
			restoreErr = os.Remove(c.EffectiveConfigPath())
			if os.IsNotExist(restoreErr) {
				restoreErr = nil
			}
		}
		return errors.Join(err, restoreErr)
	}
	return nil
}

// DropProducerReservations removes every memory reservation owned by this
// composer's producer from the effective configuration and ownership registry,
// preserving all other reservations and their listeners untouched..
func (c *Composer) DropProducerReservations(cfg EffectiveConfig, owners OwnershipDoc) (EffectiveConfig, OwnershipDoc) {
	prod := make(map[reservationKey]bool)
	keptOwn := owners.Reservations[:0]
	for _, r := range owners.Reservations {
		if r.Owner == c.Producer {
			prod[reservationKey{port: r.Port, unitID: r.UnitID}] = true
		} else {
			keptOwn = append(keptOwn, r)
		}
	}

	keptListeners := make([]Listener, 0, len(cfg.Listeners))
	for _, l := range cfg.Listeners {
		port := ListenPort(l.Listen)
		if port == 0 {
			keptListeners = append(keptListeners, l)
			continue
		}
		mems := make([]Memory, 0, len(l.Memory))
		for _, m := range l.Memory {
			if prod[reservationKey{port: port, unitID: m.UnitID}] {
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

// DropOneReservation removes one reservation owned by this composer's
// producer,preserving everything else untouched;callers verify ownership via
// Collision first..
func (c *Composer) DropOneReservation(cfg EffectiveConfig, owners OwnershipDoc, port, unitID uint16) (EffectiveConfig, OwnershipDoc) {
	keptOwn := owners.Reservations[:0]
	for _, r := range owners.Reservations {
		if r.Owner == c.Producer && r.Port == port && r.UnitID == unitID {
			continue
		}
		keptOwn = append(keptOwn, r)
	}

	keptListeners := make([]Listener, 0, len(cfg.Listeners))
	for _, l := range cfg.Listeners {
		port := ListenPort(l.Listen)
		if port == 0 {
			keptListeners = append(keptListeners, l)
			continue
		}
		mems := make([]Memory, 0, len(l.Memory))
		for _, m := range l.Memory {
			if m.UnitID == unitID && port == port {
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

// AddMemory inserts one memory reservation into the effective configuration,
// reusing an existing listener on the same TCP port so the composed config never
// contains two listeners bound to the same TCP port..
func AddMemory(cfg EffectiveConfig, id, listen string, mem Memory) EffectiveConfig {
	for i := range cfg.Listeners {
		if ListenPort(cfg.Listeners[i].Listen) != ListenPort(listen) {
			continue
		}
		cfg.Listeners[i].Memory = append(cfg.Listeners[i].Memory, mem)
		return cfg
	}
	cfg.Listeners = append(cfg.Listeners, Listener{ID: id, Listen: listen, Memory: []Memory{mem}})
	return cfg
}

// ListenPort extracts the TCP port from an MMA2 listen address("host:port"(,
// returning 0 when the value cannot be parsed..
func ListenPort(listen string) uint16 {
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
