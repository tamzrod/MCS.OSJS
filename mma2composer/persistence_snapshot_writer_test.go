package mma2composer

import (
	"errors"
	"testing"
)

type fakeAreaReader struct {
	data  map[string][]byte
	err   error
	reads int
}

func (f *fakeAreaReader) ReadPersistenceArea(key PersistenceMemoryKey, area string, kind PersistenceAreaKind, start, count uint16) ([]byte, error) {
	f.reads++
	if f.err != nil {
		return nil, f.err
	}
	return append([]byte(nil), f.data[area]...), nil
}

type storeCall struct {
	area   string
	offset int
	data   []byte
}

type fakeSnapshotStore struct {
	calls []storeCall
}

func (f *fakeSnapshotStore) WritePersistenceBytes(key PersistenceMemoryKey, area string, offset int, data []byte) error {
	f.calls = append(f.calls, storeCall{area: area, offset: offset, data: append([]byte(nil), data...)})
	return nil
}

func snapshotWriterFixture(t *testing.T) (*PersistenceSnapshotWriter, *fakeAreaReader, *fakeSnapshotStore) {
	t.Helper()
	key := PersistenceMemoryKey{Port: 5020, UnitID: 1}
	rules, err := PersistenceSnapshotRules(key, []PersistenceRBERule{
		{ID: 1, Area: "holding_registers", Start: 0, Count: 3, SystemOwned: true},
		{ID: 2, Area: "coils", Start: 0, Count: 8, SystemOwned: true},
	})
	if err != nil {
		t.Fatal(err)
	}
	reader := &fakeAreaReader{data: map[string][]byte{
		"holding_registers": {0, 1, 0, 2, 0, 3},
		"coils":             {0x00},
	}}
	store := &fakeSnapshotStore{}
	writer, err := NewPersistenceSnapshotWriter(rules, reader, store)
	if err != nil {
		t.Fatal(err)
	}
	return writer, reader, store
}

// PERSIST-012 self-check: events drive reads and only changed bytes are written;
// unchanged state causes no store (disk) call.
func TestPersistenceSnapshotWriterEventDrivenChangedBytesOnly(t *testing.T) {
	writer, reader, store := snapshotWriterFixture(t)

	// First event: no image yet -> whole area written.
	result, err := writer.OnPersistenceEvent(1)
	if err != nil {
		t.Fatal(err)
	}
	if !result.Matched || result.StoreCalls != 1 || result.BytesWritten != 6 {
		t.Fatalf("initial write wrong: %+v", result)
	}
	if reader.reads != 1 {
		t.Fatalf("event must read authoritative state once, reads=%d", reader.reads)
	}
	if store.calls[0].offset != 0 || len(store.calls[0].data) != 6 {
		t.Fatalf("initial store call wrong: %+v", store.calls[0])
	}

	// Unchanged state: no disk write at all.
	before := len(store.calls)
	result, err = writer.OnPersistenceEvent(1)
	if err != nil {
		t.Fatal(err)
	}
	if result.StoreCalls != 0 || result.BytesWritten != 0 {
		t.Fatalf("unchanged state wrote: %+v", result)
	}
	if len(store.calls) != before {
		t.Fatal("unchanged state caused a store call")
	}

	// One register word changes: only that word (2 bytes) is written.
	reader.data["holding_registers"] = []byte{0, 1, 0, 9, 0, 3}
	result, err = writer.OnPersistenceEvent(1)
	if err != nil {
		t.Fatal(err)
	}
	if result.StoreCalls != 1 || result.BytesWritten != 2 {
		t.Fatalf("single-word change wrote %+v", result)
	}
	last := store.calls[len(store.calls)-1]
	if last.offset != 2 || len(last.data) != 2 || last.data[0] != 0 || last.data[1] != 9 {
		t.Fatalf("changed word written at wrong place: %+v", last)
	}

	// Two non-adjacent words change: two runs, 4 bytes total.
	before = len(store.calls)
	reader.data["holding_registers"] = []byte{7, 7, 0, 9, 8, 8}
	result, err = writer.OnPersistenceEvent(1)
	if err != nil {
		t.Fatal(err)
	}
	if result.StoreCalls != 2 || result.BytesWritten != 4 {
		t.Fatalf("two separated words wrote %+v", result)
	}
	if len(store.calls) != before+2 {
		t.Fatalf("expected two store calls, got %d", len(store.calls)-before)
	}

	// A non-persistence ID is ignored: no read, no write.
	reads, calls := reader.reads, len(store.calls)
	result, err = writer.OnPersistenceEvent(99)
	if err != nil {
		t.Fatal(err)
	}
	if result.Matched || reader.reads != reads || len(store.calls) != calls {
		t.Fatalf("unknown ID was not ignored: %+v", result)
	}

	// Bit area event: single byte written, then unchanged -> no write.
	result, err = writer.OnPersistenceEvent(2)
	if err != nil {
		t.Fatal(err)
	}
	if !result.Matched || result.StoreCalls != 1 || result.BytesWritten != 1 {
		t.Fatalf("bit area initial write wrong: %+v", result)
	}
	result, _ = writer.OnPersistenceEvent(2)
	if result.StoreCalls != 0 {
		t.Fatalf("unchanged bit area wrote: %+v", result)
	}
}

// A read failure writes nothing and does not advance the image.
func TestPersistenceSnapshotWriterReadErrorWritesNothing(t *testing.T) {
	writer, reader, store := snapshotWriterFixture(t)
	reader.err = errors.New("memory unavailable")
	result, err := writer.OnPersistenceEvent(1)
	if err == nil {
		t.Fatal("expected read error to surface")
	}
	if !result.Matched || len(store.calls) != 0 {
		t.Fatalf("read error caused a write: %+v", result)
	}
	// Recovery: once the read succeeds the image initializes normally.
	reader.err = nil
	result, err = writer.OnPersistenceEvent(1)
	if err != nil || result.StoreCalls != 1 {
		t.Fatalf("recovery write wrong: %+v err=%v", result, err)
	}
}
