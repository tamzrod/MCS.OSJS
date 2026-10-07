package mma2composer

import (
	"errors"
	"testing"
)

type rawIngestCall struct {
	key     PersistenceMemoryKey
	area    PersistenceRawIngestArea
	start   uint16
	count   uint16
	payload []byte
}

type fakeRawIngestWriter struct {
	calls     []rawIngestCall
	responses map[PersistenceRawIngestArea]byte
	err       error
}

func (f *fakeRawIngestWriter) WritePersistenceRawIngest(key PersistenceMemoryKey, area PersistenceRawIngestArea, start, count uint16, payload []byte) (byte, error) {
	f.calls = append(f.calls, rawIngestCall{key: key, area: area, start: start, count: count, payload: append([]byte(nil), payload...)})
	if f.err != nil {
		return 0, f.err
	}
	if f.responses == nil {
		return PersistenceRawIngestOK, nil
	}
	if code, ok := f.responses[area]; ok {
		return code, nil
	}
	return PersistenceRawIngestOK, nil
}

// restoreReadyPlan builds a Ready plan for one register area and one bit area.
func restoreReadyPlan(t *testing.T) (PersistenceRestorePlan, []byte, []byte) {
	t.Helper()
	key, areas, source := loaderFixture(t)
	plan, err := LoadPersistenceSnapshots(key, true, true, areas, source)
	if err != nil {
		t.Fatal(err)
	}
	if plan.State != PersistenceRestoreReady {
		t.Fatalf("fixture must be ready: %s", plan.State)
	}
	return plan, source.payloads["holding_registers"], source.payloads["coils"]
}

// PERSIST-016 self-check: every ready area is written back to the exact
// area/start/count identity it was loaded from, using the existing Raw Ingest
// area codes, and success requires a success response for every area.
func TestPersistenceRestoreWritesEachAreaToItsIdentity(t *testing.T) {
	plan, regs, bits := restoreReadyPlan(t)
	writer := &fakeRawIngestWriter{}

	result, err := RestorePersistencePlan(plan, writer)
	if err != nil {
		t.Fatal(err)
	}
	if !result.Completed || result.Written != 2 || result.Areas != 2 {
		t.Fatalf("complete restore expected: %+v", result)
	}
	if len(writer.calls) != 2 {
		t.Fatalf("expected exactly one write per area, got %d", len(writer.calls))
	}
	for _, call := range writer.calls {
		if call.key != plan.Key {
			t.Fatalf("write used wrong memory identity: %+v", call)
		}
	}
	byArea := map[PersistenceRawIngestArea]rawIngestCall{}
	for _, call := range writer.calls {
		byArea[call.area] = call
	}
	hr, ok := byArea[RawIngestHoldingRegisters]
	if !ok || hr.start != 10 || hr.count != 2 || len(hr.payload) != len(regs) {
		t.Fatalf("holding registers restored to wrong identity: %+v", hr)
	}
	if string(hr.payload) != string(regs) {
		t.Fatalf("holding register payload not carried unchanged: %v", hr.payload)
	}
	coils, ok := byArea[RawIngestCoils]
	if !ok || coils.start != 0 || coils.count != 9 || len(coils.payload) != len(bits) {
		t.Fatalf("coils restored to wrong identity: %+v", coils)
	}
	if string(coils.payload) != string(bits) {
		t.Fatalf("coil payload not carried unchanged: %v", coils.payload)
	}
}

// No cross-area mirroring: the restore writes exactly the plan's areas and never
// duplicates a payload into another area.
func TestPersistenceRestoreAddsNoCrossAreaMirroring(t *testing.T) {
	plan, regs, bits := restoreReadyPlan(t)
	writer := &fakeRawIngestWriter{}

	if _, err := RestorePersistencePlan(plan, writer); err != nil {
		t.Fatal(err)
	}
	if len(writer.calls) != len(plan.Areas) {
		t.Fatalf("write count must equal configured area count: %d vs %d", len(writer.calls), len(plan.Areas))
	}
	seen := map[PersistenceRawIngestArea]int{}
	for _, call := range writer.calls {
		seen[call.area]++
		switch call.area {
		case RawIngestHoldingRegisters:
			if string(call.payload) != string(regs) {
				t.Fatalf("register area carried a foreign payload: %v", call.payload)
			}
		case RawIngestCoils:
			if string(call.payload) != string(bits) {
				t.Fatalf("bit area carried a foreign payload: %v", call.payload)
			}
		default:
			t.Fatalf("unexpected mirrored area written: %d", call.area)
		}
	}
	for area, count := range seen {
		if count != 1 {
			t.Fatalf("area %d written %d times", area, count)
		}
	}
}

// Every Raw Ingest response is checked: a non-zero response aborts the restore
// immediately and no later area is written.
func TestPersistenceRestoreAbortsOnNonZeroResponse(t *testing.T) {
	for _, code := range []byte{0x10, 0x12, 0x14, 0x20, 0x21, 0x30} {
		plan, _, _ := restoreReadyPlan(t)
		// holding_registers is written first in the fixture's area order.
		writer := &fakeRawIngestWriter{responses: map[PersistenceRawIngestArea]byte{RawIngestHoldingRegisters: code}}

		result, err := RestorePersistencePlan(plan, writer)
		if err != nil {
			t.Fatalf("a response code must not surface as a transport error: %v", err)
		}
		if result.Completed {
			t.Fatalf("code 0x%02X must not complete the restore", code)
		}
		if result.Written != 0 {
			t.Fatalf("code 0x%02X must abort before recording a successful write: %+v", code, result)
		}
		if len(writer.calls) != 1 {
			t.Fatalf("code 0x%02X must abort after the failing write, calls=%d", code, len(writer.calls))
		}
		if writer.calls[0].area != RawIngestHoldingRegisters {
			t.Fatalf("abort must occur at the failing area: %+v", writer.calls[0])
		}
	}
}

// A non-ready plan is refused without writing anything.
func TestPersistenceRestoreRefusesNonReadyPlan(t *testing.T) {
	key, areas, source := loaderFixture(t)
	delete(source.snapshots, "coils")
	plan, err := LoadPersistenceSnapshots(key, true, true, areas, source)
	if err != nil {
		t.Fatal(err)
	}
	if plan.State == PersistenceRestoreReady {
		t.Fatal("fixture must not be ready")
	}
	writer := &fakeRawIngestWriter{}
	result, err := RestorePersistencePlan(plan, writer)
	if err != nil {
		t.Fatal(err)
	}
	if result.Completed || result.Written != 0 || len(writer.calls) != 0 {
		t.Fatalf("non-ready plan must not write: %+v calls=%d", result, len(writer.calls))
	}
}

// A transport error surfaces and aborts the restore.
func TestPersistenceRestoreSurfacesTransportError(t *testing.T) {
	plan, _, _ := restoreReadyPlan(t)
	writer := &fakeRawIngestWriter{err: errors.New("raw ingest unavailable")}
	result, err := RestorePersistencePlan(plan, writer)
	if err == nil {
		t.Fatal("transport error must surface")
	}
	if result.Completed || result.Written != 0 {
		t.Fatalf("transport error must abort restore: %+v", result)
	}
}

// A missing writer is rejected for a ready plan.
func TestPersistenceRestoreRequiresWriter(t *testing.T) {
	plan, _, _ := restoreReadyPlan(t)
	if _, err := RestorePersistencePlan(plan, nil); err == nil {
		t.Fatal("nil writer must be rejected for a ready plan")
	}
}

// The Raw Ingest area codes match the v1 wire contract.
func TestPersistenceRestoreAreaCodes(t *testing.T) {
	cases := map[string]PersistenceRawIngestArea{
		"coils":             RawIngestCoils,
		"discrete_inputs":   RawIngestDiscreteInputs,
		"holding_registers": RawIngestHoldingRegisters,
		"input_registers":   RawIngestInputRegisters,
	}
	for name, want := range cases {
		got, ok := rawIngestArea(name)
		if !ok || got != want {
			t.Fatalf("area %q mapped to %d, want %d", name, got, want)
		}
	}
	if RawIngestCoils != 1 || RawIngestDiscreteInputs != 2 || RawIngestHoldingRegisters != 3 || RawIngestInputRegisters != 4 {
		t.Fatal("raw ingest area codes drifted from the v1 contract")
	}
	if _, ok := rawIngestArea("nonsense"); ok {
		t.Fatal("unknown area must not map")
	}
}
