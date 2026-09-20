package simulator

import (
	"fmt"

	"github.com/tamzrod/MCS.OSJS/mma2composer"
)

const ProducerSimulator = "simulator"

var ErrReservationOwnedByOther = mma2composer.ErrReservationOwnedByOther

type OwnershipEntry = mma2composer.OwnershipEntry
type OwnershipDoc = mma2composer.OwnershipDoc
type MMA2Area = mma2composer.Area
type MMA2PolicyRule = mma2composer.PolicyRule
type MMA2Policy = mma2composer.Policy
type MMA2Memory = mma2composer.Memory
type MMA2Listener = mma2composer.Listener
type EffectiveMMA2Config = mma2composer.EffectiveConfig

type mma2Key struct {
	port   uint16
	unitID uint16
}

func (s Store) composer() *mma2composer.Composer { return mma2composer.New(s.Root, ProducerSimulator) }
func (s Store) MMA2ConfigDir() string            { return s.composer().ConfigDir() }
func (s Store) EffectiveConfigPath() string      { return s.composer().EffectiveConfigPath() }
func (s Store) OwnershipPath() string            { return s.composer().OwnershipPath() }

// SaveAndCompose is a legacy public entry point: lock ONCE through both shared
// artifacts and the Simulator document. Do not call another locking Store API.
func (s Store) SaveAndCompose(def DeviceDefinition) error {
	if err := ValidateDevice(def); err != nil { return err }
	return s.withWriterLock(func() error {
		composer := s.composer()
		cfg, err := composer.LoadEffective()
		if err != nil { return err }
		owners, err := composer.LoadOwners()
		if err != nil { return err }
		if err := mma2composer.Collision(def.MMA2.Port, def.MMA2.UnitID, ProducerSimulator, owners); err != nil { return err }
		inheritMemorySettings(&def.MMA2, cfg)
		cfg, owners = composer.DropProducerReservations(cfg, owners)
		if hasMMA2Areas(def.MMA2) {
			cfg = addReservation(cfg, def.MMA2)
			owners.Reservations = append(owners.Reservations, OwnershipEntry{Port: def.MMA2.Port, UnitID: def.MMA2.UnitID, Owner: ProducerSimulator})
		}
		if err := composer.Commit(cfg, owners); err != nil { return err }
		return s.replace(Document{Devices: []DeviceDefinition{def}})
	})
}

// ComposeDocument is the public standalone entry. Live structural apply and
// boot restoration call composeDocumentLocked under their own outer boundary.
func (s Store) ComposeDocument(doc Document) error {
	return s.withWriterLock(func() error { return s.composeDocumentLocked(doc) })
}

// composeDocumentLocked must be called only with the shared writer lock held.
// Validation remains here because boot and direct structural calls also use it.
func (s Store) composeDocumentLocked(doc Document) error {
	for i := range doc.Devices {
		if err := ValidateDevice(doc.Devices[i]); err != nil {
			return fmt.Errorf("device %d: %w", i, err)
		}
	}
	composer := s.composer()
	cfg, err := composer.LoadEffective()
	if err != nil { return err }
	owners, err := composer.LoadOwners()
	if err != nil { return err }
	for index := range doc.Devices {
		def := &doc.Devices[index]
		if err := mma2composer.Collision(def.MMA2.Port, def.MMA2.UnitID, ProducerSimulator, owners); err != nil { return err }
		inheritMemorySettings(&def.MMA2, cfg)
	}
	cfg, owners = composer.DropProducerReservations(cfg, owners)
	seen := make(map[mma2Key]bool)
	for _, def := range doc.Devices {
		if !def.Enabled || !hasMMA2Areas(def.MMA2) { continue }
		key := mma2Key{port: def.MMA2.Port, unitID: def.MMA2.UnitID}
		if seen[key] { return fmt.Errorf("duplicate simulator MMA2 reservation (%d,%d)", key.port, key.unitID) }
		seen[key] = true
		cfg = addReservation(cfg, def.MMA2)
		owners.Reservations = append(owners.Reservations, OwnershipEntry{Port: key.port, UnitID: key.unitID, Owner: ProducerSimulator})
	}
	return composer.Commit(cfg, owners)
}

func (s Store) DeleteAndCompose(port, unitID uint16) error {
	return s.withWriterLock(func() error {
		composer := s.composer()
		cfg, err := composer.LoadEffective()
		if err != nil { return err }
		owners, err := composer.LoadOwners()
		if err != nil { return err }
		if err := mma2composer.Collision(port, unitID, ProducerSimulator, owners); err != nil { return err }
		cfg, owners = composer.DropOneReservation(cfg, owners, port, unitID)
		return composer.Commit(cfg, owners)
	})
}

func (s Store) loadEffective() (EffectiveMMA2Config, error) { return s.composer().LoadEffective() }
func (s Store) loadOwners() (OwnershipDoc, error)           { return s.composer().LoadOwners() }
func (s Store) saveEffective(cfg EffectiveMMA2Config) error {
	return s.withWriterLock(func() error { return s.composer().SaveEffective(cfg) })
}
func (s Store) saveOwners(doc OwnershipDoc) error {
	return s.withWriterLock(func() error { return s.composer().SaveOwners(doc) })
}
// replaceMMA2File is a lock-free INTERNAL helper used by restart-request
// operations while their enclosing transaction already owns the writer lock.
func (s Store) replaceMMA2File(path string, b []byte) error {
	return s.composer().WriteFileAtomic(path, b)
}

func addReservation(cfg EffectiveMMA2Config, p MMA2Params) EffectiveMMA2Config {
	return mma2composer.AddMemory(cfg, fmt.Sprintf("sim-%d-%d", p.Port, p.UnitID), fmt.Sprintf("0.0.0.0:%d", p.Port), memoryFromMMA2Params(p))
}

func memoryFromMMA2Params(p MMA2Params) MMA2Memory {
	mem := MMA2Memory{UnitID: p.UnitID}
	if p.FC1.Count > 0 { mem.Coils = &MMA2Area{Start: p.FC1.Start, Count: p.FC1.Count} }
	if p.FC2.Count > 0 { mem.DiscreteInputs = &MMA2Area{Start: p.FC2.Start, Count: p.FC2.Count} }
	if p.FC3.Count > 0 { mem.HoldingRegs = &MMA2Area{Start: p.FC3.Start, Count: p.FC3.Count} }
	if p.FC4.Count > 0 { mem.InputRegs = &MMA2Area{Start: p.FC4.Start, Count: p.FC4.Count} }
	mem.Policy = &MMA2Policy{Rules: []MMA2PolicyRule{{
		ID: "simulator-fc-access",
		SourceIP: []string{"0.0.0.0/0", "::/0", "127.0.0.1", "::1"},
		AllowFC: []uint8{1, 2, 3, 4, 5, 6, 15, 16},
	}}}
	if p.Policy != nil { mem.Policy = p.Policy }
	mem.Extra = make(map[string]interface{})
	for key, value := range p.Extra { mem.Extra[key] = value }
	if p.StateSealing != nil { mem.Extra["state_sealing"] = p.StateSealing }
	if p.RBE != nil { mem.Extra["rbe"] = p.RBE }
	return mem
}

func inheritMemorySettings(params *MMA2Params, cfg EffectiveMMA2Config) {
	for _, listener := range cfg.Listeners {
		if mma2composer.ListenPort(listener.Listen) != params.Port { continue }
		for _, memory := range listener.Memory {
			if memory.UnitID != params.UnitID { continue }
			if params.Policy == nil { params.Policy = memory.Policy }
			if params.Extra == nil { params.Extra = make(map[string]interface{}) }
			for key, value := range memory.Extra {
				switch key {
				case "state_sealing":
					if params.StateSealing == nil { params.StateSealing, _ = value.(map[string]interface{}) }
				case "rbe":
					if params.RBE == nil { params.RBE, _ = value.(map[string]interface{}) }
				default:
					if _, exists := params.Extra[key]; !exists { params.Extra[key] = value }
				}
			}
		}
	}
}

func hasMMA2Areas(p MMA2Params) bool {
	return p.FC1.Count > 0 || p.FC2.Count > 0 || p.FC3.Count > 0 || p.FC4.Count > 0
}

func listenPort(listen string) uint16 { return mma2composer.ListenPort(listen) }
