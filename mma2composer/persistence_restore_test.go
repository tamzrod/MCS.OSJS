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

// coilCall returns the coils write captured by the fake writer.
func coilCall(t *testing.T, writer *fakeRawIngestWriter) rawIngestCall {
	t.Helper()
	for _, call := range writer.calls {
		if call.area == RawIngestCoils {
			return call
		}
	}
	t.Fatal("no coils write was made")
	return rawIngestCall{}
}

func areaPayload(plan PersistenceRestorePlan, area string) []byte {
	for _, a := range plan.Areas {
		if a.Area == area {
			return a.Payload
		}
	}
	return nil
}

// PERSIST-017 self-check: the sealing flag location is derived from the
// authoritative State Sealing configuration only.
func TestPersistenceSealingFlagFromExtra(t *testing.T) {
	cases := []struct {
		name  string
		extra map[string]interface{}
		want  uint16
		ok    bool
	}{
		{"absent-block", map[string]interface{}{}, 0, false},
		{"disabled-explicit", map[string]interface{}{"state_sealing": map[string]interface{}{"enabled": false, "area": "coil", "address": 5}}, 0, false},
		{"enabled-default", map[string]interface{}{"state_sealing": map[string]interface{}{"area": "coil", "address": 5}}, 5, true},
		{"enabled-true", map[string]interface{}{"state_sealing": map[string]interface{}{"enabled": true, "area": "coil", "address": 5}}, 5, true},
		{"area-case-insensitive", map[string]interface{}{"state_sealing": map[string]interface{}{"area": " Coil ", "address": 7}}, 7, true},
		{"wrong-area", map[string]interface{}{"state_sealing": map[string]interface{}{"area": "holding_register", "address": 5}}, 0, false},
		{"no-area", map[string]interface{}{"state_sealing": map[string]interface{}{"address": 5}}, 0, false},
		{"uint16-address", map[string]interface{}{"state_sealing": map[string]interface{}{"area": "coil", "address": uint16(9)}}, 9, true},
		{"negative-address", map[string]interface{}{"state_sealing": map[string]interface{}{"area": "coil", "address": -1}}, 0, false},
		{"address-too-large", map[string]interface{}{"state_sealing": map[string]interface{}{"area": "coil", "address": 70000}}, 0, false},
		{"missing-address", map[string]interface{}{"state_sealing": map[string]interface{}{"area": "coil"}}, 0, false},
	}
	for _, tc := range cases {
		got, ok := PersistenceSealingFlagFromExtra(tc.extra)
		if ok != tc.ok || (ok && got.Address != tc.want) {
			t.Fatalf("%s: got (%+v,%v), want address %d ok=%v", tc.name, got, ok, tc.want, tc.ok)
		}
	}
}

// PERSIST-017 self-check: a snapshot captured while unsealed (flag=1) cannot
// unseal during restore — the sealing bit is forced to sealed (0) while every
// other restored bit is preserved, and the source payload is not mutated.
func TestPersistenceRestoreKeepsSealingFlagSealed(t *testing.T) {
	key, areas, source := loaderFixture(t)
	plan, err := LoadPersistenceSnapshots(key, true, true, areas, source)
	if err != nil {
		t.Fatal(err)
	}
	if plan.State != PersistenceRestoreReady {
		t.Fatalf("fixture must be ready: %s", plan.State)
	}
	// Coils snapshot bits are [1,0,1,0,1,0,1,0,1]; index 4 is unsealed (1).
	flag := PersistenceSealingFlag{Address: 4}
	plan.SealingFlag = &flag
	originalCoils := append([]byte(nil), source.payloads["coils"]...)
	originalRegs := append([]byte(nil), source.payloads["holding_registers"]...)

	writer := &fakeRawIngestWriter{}
	result, err := RestorePersistencePlan(plan, writer)
	if err != nil {
		t.Fatal(err)
	}
	if !result.Completed {
		t.Fatalf("restore must complete: %+v", result)
	}

	coils := coilCall(t, writer)
	decoded, err := DecodePersistenceBits(coils.payload, 9)
	if err != nil {
		t.Fatal(err)
	}
	if decoded[4] {
		t.Fatalf("sealing flag bit 4 must be sealed (0) after restore: %v", decoded)
	}
	wantBits := []bool{true, false, true, false, false, false, true, false, true}
	for i := range wantBits {
		if decoded[i] != wantBits[i] {
			t.Fatalf("bit %d changed unexpectedly: got %v want %v", i, decoded, wantBits)
		}
	}
	for _, call := range writer.calls {
		if call.area == RawIngestHoldingRegisters && string(call.payload) != string(originalRegs) {
			t.Fatalf("register payload was altered: %v", call.payload)
		}
	}
	if string(source.payloads["coils"]) != string(originalCoils) {
		t.Fatalf("source coil payload was mutated: %v", source.payloads["coils"])
	}
	if string(areaPayload(plan, "coils")) != string(originalCoils) {
		t.Fatalf("plan coil payload was mutated: %v", areaPayload(plan, "coils"))
	}
}

// Without an authoritative sealing flag the coils payload is written unchanged.
func TestPersistenceRestoreNoSealingFlagUnchanged(t *testing.T) {
	plan, _, bits := restoreReadyPlan(t)
	writer := &fakeRawIngestWriter{}
	if _, err := RestorePersistencePlan(plan, writer); err != nil {
		t.Fatal(err)
	}
	if string(coilCall(t, writer).payload) != string(bits) {
		t.Fatalf("coils payload must be unchanged without a sealing flag")
	}
}

// A sealing flag outside the restored coil range cannot be affected by the write,
// so the payload is left unchanged rather than aborting.
func TestPersistenceRestoreSealingFlagOutsideCoilRangeUnchanged(t *testing.T) {
	plan, _, bits := restoreReadyPlan(t)
	flag := PersistenceSealingFlag{Address: 100}
	plan.SealingFlag = &flag
	writer := &fakeRawIngestWriter{}
	result, err := RestorePersistencePlan(plan, writer)
	if err != nil {
		t.Fatal(err)
	}
	if !result.Completed {
		t.Fatalf("restore must still complete: %+v", result)
	}
	if string(coilCall(t, writer).payload) != string(bits) {
		t.Fatalf("out-of-range sealing flag must not alter the payload")
	}
}

// forcePersistenceSealingFlag clears exactly one bit, handles byte boundaries,
// preserves every other bit, and fails closed when the payload is too short.
func TestForcePersistenceSealingFlag(t *testing.T) {
	got, err := forcePersistenceSealingFlag(0, 8, []byte{0xFF}, PersistenceSealingFlag{Address: 0})
	if err != nil || got[0] != 0xFE {
		t.Fatalf("address 0: got %v err=%v", got, err)
	}
	got, err = forcePersistenceSealingFlag(0, 8, []byte{0xFF}, PersistenceSealingFlag{Address: 7})
	if err != nil || got[0] != 0x7F {
		t.Fatalf("address 7: got %v err=%v", got, err)
	}
	got, err = forcePersistenceSealingFlag(0, 16, []byte{0xFF, 0xFF}, PersistenceSealingFlag{Address: 8})
	if err != nil || got[0] != 0xFF || got[1] != 0xFE {
		t.Fatalf("address 8: got %v err=%v", got, err)
	}
	got, err = forcePersistenceSealingFlag(10, 2, []byte{0x03}, PersistenceSealingFlag{Address: 10})
	if err != nil || got[0] != 0x02 {
		t.Fatalf("start 10 address 10: got %v err=%v", got, err)
	}
	if _, err := forcePersistenceSealingFlag(0, 16, []byte{0xFF}, PersistenceSealingFlag{Address: 8}); err == nil {
		t.Fatal("short payload must fail closed")
	}
	src := []byte{0xFF}
	got, err = forcePersistenceSealingFlag(0, 1, src, PersistenceSealingFlag{Address: 5})
	if err != nil || got[0] != 0xFF {
		t.Fatalf("out-of-range: got %v err=%v", got, err)
	}
	if &got[0] == &src[0] {
		t.Fatal("must return a copy, not the same slice")
	}
}
