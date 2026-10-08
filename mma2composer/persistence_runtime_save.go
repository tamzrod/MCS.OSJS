package mma2composer

import (
	"fmt"
	"time"
)

// PersistenceRuntimeSave is the concrete runtime wiring that connects
// persistence-owned RBE events to the existing PersistenceSnapshotWriter and a
// PersistenceSnapshotStore (the PERSIST-R01 filesystem adapter in production).
//
// Only the system-derived persistence RBE projection is subscribed: each
// persistence-owned rule ID maps to the memory/area it covers. A user-owned RBE
// rule never publishes a persistence event, so user RBE remains independent and
// persistence ranges stay derived and system-owned. The save path only persists
// changed bytes/words (change-only), exactly as PersistenceSnapshotWriter does.
//
// Status is observational: it reports save activity and last error, and never
// grants control authority over restore, sealing or MMA2.
type PersistenceRuntimeSave struct {
	writer *PersistenceSnapshotWriter
	key    PersistenceMemoryKey
	ids    map[uint8]PersistenceSnapshotRule
	status PersistenceSaveStatus
}

// PersistenceSaveStatus is the read-only observation of the runtime save path.
type PersistenceSaveStatus struct {
	// ConfiguredRules is the number of persistence-owned rules subscribed.
	ConfiguredRules int `json:"configured_rules"`
	// Events is the number of persistence events observed.
	Events int `json:"events"`
	// Saves is the number of events that wrote at least one changed byte/word.
	Saves int `json:"saves"`
	// BytesWritten is the cumulative number of changed bytes/words persisted.
	BytesWritten int `json:"bytes_written"`
	// LastSaveAt is the observed instant of the most recent successful save
	// (a run that wrote at least one changed byte/word). It is empty when nothing
	// has been saved yet and is never fabricated.
	LastSaveAt string `json:"last_save_at,omitempty"`
	// LastError is the most recent save error, empty when the last save path was
	// clean. It never fabricates a success.
	LastError string `json:"last_error,omitempty"`
}

// PersistenceRuntimeSaveConfig registers one memory's persistence save wiring.
type PersistenceRuntimeSaveConfig struct {
	Key    PersistenceMemoryKey
	Rules  []PersistenceRBERule
	Reader PersistenceAreaReader
	Store  PersistenceSnapshotStore
}

// NewPersistenceRuntimeSave builds the runtime save component: it derives the
// change-only writer rules from the system-owned persistence RBE projection and
// returns the set of subscribed persistence rule IDs. It refuses a nil reader or
// store and rejects a caller that tries to subscribe a user-owned (non-system)
// rule as persistence.
func NewPersistenceRuntimeSave(cfg PersistenceRuntimeSaveConfig) (*PersistenceRuntimeSave, error) {
	if cfg.Reader == nil || cfg.Store == nil {
		return nil, fmt.Errorf("persistence runtime save requires a reader and a store")
	}
	for _, rule := range cfg.Rules {
		if !rule.SystemOwned {
			return nil, fmt.Errorf("persistence runtime save cannot subscribe user-owned rule for area %q", rule.Area)
		}
	}
	rules, err := PersistenceSnapshotRules(cfg.Key, cfg.Rules)
	if err != nil {
		return nil, err
	}
	writer, err := NewPersistenceSnapshotWriter(rules, cfg.Reader, cfg.Store)
	if err != nil {
		return nil, err
	}
	ids := make(map[uint8]PersistenceSnapshotRule, len(rules))
	for _, rule := range rules {
		ids[rule.ID] = rule
	}
	return &PersistenceRuntimeSave{
		writer: writer,
		key:    cfg.Key,
		ids:    ids,
		status: PersistenceSaveStatus{ConfiguredRules: len(rules)},
	}, nil
}

// Key returns the Port -> Unit ID -> Memory identity this component saves.
func (s *PersistenceRuntimeSave) Key() PersistenceMemoryKey { return s.key }

// SubscribedRuleIDs returns the persistence-owned RBE rule IDs this component
// listens for, so a runtime can route exactly those one-byte events here.
func (s *PersistenceRuntimeSave) SubscribedRuleIDs() []uint8 {
	ids := make([]uint8, 0, len(s.ids))
	for id := range s.ids {
		ids = append(ids, id)
	}
	return ids
}

// Handles reports whether a one-byte RBE ID belongs to a persistence-owned rule.
// A user-owned rule ID returns false and never reaches the save path.
func (s *PersistenceRuntimeSave) Handles(id uint8) bool {
	_, ok := s.ids[id]
	return ok
}

// Publish implements the one-byte RBE sink contract: the runtime routes
// persistence-owned RBE events here. A non-persistence ID is ignored and never
// touches storage. It observes the outcome so operators can see save activity;
// it never performs control actions.
func (s *PersistenceRuntimeSave) Publish(id uint8) {
	if !s.Handles(id) {
		return
	}
	result, err := s.writer.OnPersistenceEvent(id)
	s.status.Events++
	if err != nil {
		s.status.LastError = err.Error()
		return
	}
	if result.StoreCalls > 0 {
		s.status.Saves++
		s.status.BytesWritten += result.BytesWritten
		s.status.LastSaveAt = time.Now().UTC().Format(time.RFC3339Nano)
	}
	s.status.LastError = ""
}

// Status returns a read-only copy of the observed save status.
func (s *PersistenceRuntimeSave) Status() PersistenceSaveStatus {
	return s.status
}
