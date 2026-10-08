package simulator

import (
	"fmt"

	"github.com/tamzrod/MCS.OSJS/mma2composer"
)

// PersistenceAreaReaderFunc adapts a function to the persistence area reader
// contract, so a runtime can supply the authoritative current bytes of an area
// (LSB-first packed bits for bit areas, big-endian uint16 words for register
// areas) without a concrete type.
type PersistenceAreaReaderFunc func(key mma2composer.PersistenceMemoryKey, area string, kind mma2composer.PersistenceAreaKind, start, count uint16) ([]byte, error)

// ReadPersistenceArea implements mma2composer.PersistenceAreaReader.
func (f PersistenceAreaReaderFunc) ReadPersistenceArea(key mma2composer.PersistenceMemoryKey, area string, kind mma2composer.PersistenceAreaKind, start, count uint16) ([]byte, error) {
	return f(key, area, kind, start, count)
}

// PersistenceRuntimeSaveMemory registers one persistence-enabled memory's save
// wiring: its authoritative persistence-owned RBE rules (whose IDs the runtime
// subscribes) and the reader that supplies its current area bytes.
type PersistenceRuntimeSaveMemory struct {
	Key    mma2composer.PersistenceMemoryKey
	Rules  []mma2composer.PersistenceRBERule
	Reader mma2composer.PersistenceAreaReader
}

// PersistenceRuntimeSaveHost holds the concrete runtime save components for the
// persistence-enabled memories of one appliance. It connects persistence-owned
// RBE events to the PERSIST-R01 filesystem adapter beneath dataRoot and routes
// only the subscribed persistence rule IDs. It is observational and read/write
// only to the snapshot store; it never touches MMA2 memory authority, sealing,
// restore or user-owned RBE.
type PersistenceRuntimeSaveHost struct {
	components []*mma2composer.PersistenceRuntimeSave
	ids        map[uint8]bool
}

// NewPersistenceRuntimeSaveHost builds the runtime save components beneath
// dataRoot for the given persistence memories. It derives each adapter
// registration from the authoritative rules and refuses a memory that subscribes
// a user-owned rule.
func NewPersistenceRuntimeSaveHost(dataRoot string, memories []PersistenceRuntimeSaveMemory) (*PersistenceRuntimeSaveHost, error) {
	host := &PersistenceRuntimeSaveHost{ids: map[uint8]bool{}}
	for _, mem := range memories {
		configs, err := mma2composer.PersistenceSnapshotConfigs(mem.Key, mem.Rules)
		if err != nil {
			return nil, err
		}
		if len(configs) == 0 {
			// Persistence-enabled memory with no persisted area has nothing to
			// save; skip it rather than registering an empty adapter.
			continue
		}
		adapter, err := mma2composer.NewPersistenceFilesystemAdapter(dataRoot, configs)
		if err != nil {
			return nil, err
		}
		component, err := mma2composer.NewPersistenceRuntimeSave(mma2composer.PersistenceRuntimeSaveConfig{
			Key:    mem.Key,
			Rules:  mem.Rules,
			Reader: mem.Reader,
			Store:  adapter,
		})
		if err != nil {
			return nil, err
		}
		for _, id := range component.SubscribedRuleIDs() {
			host.ids[id] = true
		}
		host.components = append(host.components, component)
	}
	return host, nil
}

// Handles reports whether a one-byte RBE ID belongs to a persistence-owned rule
// of this host. Only these IDs are routed to the save path; a user-owned rule ID
// returns false.
func (h *PersistenceRuntimeSaveHost) Handles(id uint8) bool {
	if h == nil {
		return false
	}
	return h.ids[id]
}

// SubscribedRuleIDs returns the union of subscribed persistence rule IDs.
func (h *PersistenceRuntimeSaveHost) SubscribedRuleIDs() []uint8 {
	if h == nil {
		return nil
	}
	ids := make([]uint8, 0, len(h.ids))
	for id := range h.ids {
		ids = append(ids, id)
	}
	return ids
}

// Publish routes one persistence RBE event to the component that owns the rule
// ID. Non-persistence IDs are ignored.
func (h *PersistenceRuntimeSaveHost) Publish(id uint8) {
	if h == nil {
		return
	}
	for _, c := range h.components {
		if c.Handles(id) {
			c.Publish(id)
			return
		}
	}
}

// Status returns the combined observational save status for all components.
func (h *PersistenceRuntimeSaveHost) Status() mma2composer.PersistenceSaveStatus {
	combined := mma2composer.PersistenceSaveStatus{}
	if h == nil {
		return combined
	}
	for _, c := range h.components {
		s := c.Status()
		combined.ConfiguredRules += s.ConfiguredRules
		combined.Events += s.Events
		combined.Saves += s.Saves
		combined.BytesWritten += s.BytesWritten
		if s.LastError != "" {
			combined.LastError = s.LastError
		}
	}
	return combined
}

// persistenceSaveContext builds the runtime save host for a set of live devices,
// using the real appliance data root. It returns nil for a nil/empty data root or
// no persistence-enabled memory, so the runtime simply runs without persistence
// save wiring rather than inventing a path.
func persistenceSaveContext(dataRoot string, devices []DeviceDefinition) (*PersistenceRuntimeSaveHost, error) {
	if dataRoot == "" {
		return nil, nil
	}
	var memories []PersistenceRuntimeSaveMemory
	for _, device := range devices {
		if !persistenceConfigured(device) {
			continue
		}
		key := mma2composer.PersistenceMemoryKey{Port: device.MMA2.Port, UnitID: device.MMA2.UnitID}
		rules, err := persistenceOwnedRulesForMemory(memoryFromMMA2Params(device.MMA2))
		if err != nil {
			return nil, fmt.Errorf("device %q: %w", device.Name, err)
		}
		if len(rules) == 0 {
			continue
		}
		memories = append(memories, PersistenceRuntimeSaveMemory{
			Key:    key,
			Rules:  rules,
			Reader: persistenceAreaReader(device),
		})
	}
	if len(memories) == 0 {
		return nil, nil
	}
	return NewPersistenceRuntimeSaveHost(dataRoot, memories)
}

// persistenceOwnedRulesForMemory derives the persistence-owned RBE projection for
// a memory from its authoritative layout.
func persistenceOwnedRulesForMemory(mem mma2composer.Memory) ([]mma2composer.PersistenceRBERule, error) {
	if !mma2composer.PersistenceEnabled(mem) {
		return nil, nil
	}
	// Derive from the layout, then assign globally unique RBE v1 IDs (1..255). The
	// caller owns user-rule IDs; PERSIST-R02 has no user-rule source here, so the
	// derived set is allocated from the free ID space deterministically.
	return mma2composer.AllocatePersistenceRBEIDs(mma2composer.DerivePersistenceRBE(mem))
}
