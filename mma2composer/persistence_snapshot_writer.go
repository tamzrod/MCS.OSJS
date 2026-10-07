package mma2composer

import (
	"fmt"
	"sync"
)

// PersistenceAreaKind distinguishes packed bit areas from register areas, which
// determines the snapshot diff granularity (byte vs register word).
type PersistenceAreaKind int

const (
	PersistenceBits PersistenceAreaKind = iota
	PersistenceRegisters
)

// PersistenceMemoryKey is the Port -> Unit ID -> Memory identity of one
// configured memory instance.
type PersistenceMemoryKey struct {
	Port   uint16
	UnitID uint16
}

// PersistenceSnapshotRule maps one persistence-owned RBE rule ID to the memory
// area it covers, so an event can locate the authoritative state to compare.
type PersistenceSnapshotRule struct {
	ID     uint8
	Memory PersistenceMemoryKey
	Area   string
	Kind   PersistenceAreaKind
	Start  uint16
	Count  uint16
}

// PersistenceAreaReader supplies the authoritative current bytes for an area
// range: LSB-first packed bits for bit areas, big-endian uint16 words for
// register areas. Implementations read live memory on demand; they are not
// polled.
type PersistenceAreaReader interface {
	ReadPersistenceArea(key PersistenceMemoryKey, area string, kind PersistenceAreaKind, start, count uint16) ([]byte, error)
}

// PersistenceSnapshotStore receives only changed bytes at a byte offset within
// an area's snapshot image. It is never called for unchanged state.
type PersistenceSnapshotStore interface {
	WritePersistenceBytes(key PersistenceMemoryKey, area string, offset int, data []byte) error
}

// PersistenceWriteResult reports what one persistence RBE event did.
type PersistenceWriteResult struct {
	Matched      bool
	BytesWritten int
	StoreCalls   int
}

type persistenceImageKey struct {
	memory PersistenceMemoryKey
	area   string
}

// PersistenceSnapshotWriter reacts to persistence-owned RBE events and updates
// the corresponding snapshot image without continuous polling. On each event it
// reads the authoritative area state, compares it to the last snapshot image and
// writes only the changed bytes/register words. Unchanged state causes no store
// (disk) call at all.
type PersistenceSnapshotWriter struct {
	mu     sync.Mutex
	rules  map[uint8]PersistenceSnapshotRule
	reader PersistenceAreaReader
	store  PersistenceSnapshotStore
	images map[persistenceImageKey][]byte
}

// NewPersistenceSnapshotWriter builds a writer for the given persistence rules.
func NewPersistenceSnapshotWriter(rules []PersistenceSnapshotRule, reader PersistenceAreaReader, store PersistenceSnapshotStore) (*PersistenceSnapshotWriter, error) {
	if reader == nil || store == nil {
		return nil, fmt.Errorf("persistence snapshot writer requires a reader and a store")
	}
	byID := make(map[uint8]PersistenceSnapshotRule, len(rules))
	for _, rule := range rules {
		if rule.ID == 0 {
			return nil, fmt.Errorf("persistence snapshot rule ID 0 is reserved")
		}
		if rule.Count == 0 {
			return nil, fmt.Errorf("persistence snapshot rule %d has empty range", rule.ID)
		}
		if _, exists := byID[rule.ID]; exists {
			return nil, fmt.Errorf("persistence snapshot rule ID %d is duplicated", rule.ID)
		}
		byID[rule.ID] = rule
	}
	return &PersistenceSnapshotWriter{rules: byID, reader: reader, store: store, images: map[persistenceImageKey][]byte{}}, nil
}

// Publish implements the one-byte RBE sink contract: it is the subscription
// point for persistence-owned RBE events. Errors are not surfaced here (the
// signal path is non-blocking); use OnPersistenceEvent for observable results.
func (w *PersistenceSnapshotWriter) Publish(id uint8) {
	_, _ = w.OnPersistenceEvent(id)
}

// OnPersistenceEvent handles one persistence RBE event and reports the outcome.
// A non-persistence ID is ignored (Matched=false). On error nothing is written
// and the image is not advanced.
func (w *PersistenceSnapshotWriter) OnPersistenceEvent(id uint8) (PersistenceWriteResult, error) {
	w.mu.Lock()
	defer w.mu.Unlock()

	rule, ok := w.rules[id]
	if !ok {
		return PersistenceWriteResult{}, nil
	}
	current, err := w.reader.ReadPersistenceArea(rule.Memory, rule.Area, rule.Kind, rule.Start, rule.Count)
	if err != nil {
		return PersistenceWriteResult{Matched: true}, err
	}
	key := persistenceImageKey{memory: rule.Memory, area: rule.Area}
	unit := 1
	if rule.Kind == PersistenceRegisters {
		unit = 2
	}
	runs, err := changedRuns(w.images[key], current, unit)
	if err != nil {
		return PersistenceWriteResult{Matched: true}, err
	}
	result := PersistenceWriteResult{Matched: true}
	for _, run := range runs {
		if err := w.store.WritePersistenceBytes(rule.Memory, rule.Area, run.offset, current[run.offset:run.offset+run.length]); err != nil {
			return result, err
		}
		result.StoreCalls++
		result.BytesWritten += run.length
	}
	if len(runs) > 0 {
		w.images[key] = append([]byte(nil), current...)
	}
	return result, nil
}

type persistenceByteRun struct {
	offset int
	length int
}

// changedRuns compares the snapshot image to the current bytes at the given
// granularity (unit = 1 byte for bit areas, 2 bytes for register words) and
// returns the contiguous changed runs. A nil or resized image counts as fully
// changed. No run means the state is unchanged.
func changedRuns(image, current []byte, unit int) ([]persistenceByteRun, error) {
	if len(current)%unit != 0 {
		return nil, fmt.Errorf("persistence area length %d is not a multiple of unit %d", len(current), unit)
	}
	if image == nil || len(image) != len(current) {
		if len(current) == 0 {
			return nil, nil
		}
		return []persistenceByteRun{{offset: 0, length: len(current)}}, nil
	}
	var runs []persistenceByteRun
	for start := 0; start < len(current); start += unit {
		changed := false
		for i := 0; i < unit; i++ {
			if image[start+i] != current[start+i] {
				changed = true
				break
			}
		}
		if !changed {
			continue
		}
		if len(runs) > 0 && runs[len(runs)-1].offset+runs[len(runs)-1].length == start {
			runs[len(runs)-1].length += unit
		} else {
			runs = append(runs, persistenceByteRun{offset: start, length: unit})
		}
	}
	return runs, nil
}

// persistenceAreaKind maps a canonical area name to its snapshot granularity.
func persistenceAreaKind(area string) (PersistenceAreaKind, bool) {
	switch area {
	case "coils", "discrete_inputs":
		return PersistenceBits, true
	case "holding_registers", "input_registers":
		return PersistenceRegisters, true
	default:
		return 0, false
	}
}

// PersistenceSnapshotRules builds writer rules from the persistence-owned RBE
// rules (PERSIST-003/009) for one memory instance, mapping each rule's area to
// its diff granularity.
func PersistenceSnapshotRules(key PersistenceMemoryKey, rules []PersistenceRBERule) ([]PersistenceSnapshotRule, error) {
	out := make([]PersistenceSnapshotRule, 0, len(rules))
	for _, rule := range rules {
		kind, ok := persistenceAreaKind(rule.Area)
		if !ok {
			return nil, fmt.Errorf("persistence RBE rule %d has unknown area %q", rule.ID, rule.Area)
		}
		out = append(out, PersistenceSnapshotRule{ID: rule.ID, Memory: key, Area: rule.Area, Kind: kind, Start: rule.Start, Count: rule.Count})
	}
	return out, nil
}
