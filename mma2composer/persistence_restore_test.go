package mma2composer

import (
	"errors"
	"reflect"
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

// PERSIST-018 self-check: the required-area set is tracked deterministically in
// plan order, and a successful restore acknowledges exactly that set.
func TestPersistenceRestoreTracksRequiredAndAcknowledgedAreas(t *testing.T) {
	plan, _, _ := restoreReadyPlan(t)
	writer := &fakeRawIngestWriter{}
	result, err := RestorePersistencePlan(plan, writer)
	if err != nil {
		t.Fatal(err)
	}
	want := []string{"holding_registers", "coils"}
	if !reflect.DeepEqual(result.RequiredAreas, want) {
		t.Fatalf("required set must be deterministic plan order: got %v want %v", result.RequiredAreas, want)
	}
	if !reflect.DeepEqual(result.AcknowledgedAreas, want) {
		t.Fatalf("acknowledged set must equal required set: got %v want %v", result.AcknowledgedAreas, want)
	}
	if !result.Completed || !VerifyPersistenceRestore(result) {
		t.Fatalf("full acknowledgement must complete and verify: %+v", result)
	}
}

// PERSIST-018 self-check: any failed area prevents restore completion and never
// appears as acknowledged.
func TestPersistenceRestoreAnyFailedAreaPreventsCompletion(t *testing.T) {
	for _, code := range []byte{0x10, 0x12, 0x14, 0x20, 0x21, 0x30} {
		plan, _, _ := restoreReadyPlan(t)
		writer := &fakeRawIngestWriter{responses: map[PersistenceRawIngestArea]byte{RawIngestHoldingRegisters: code}}
		result, err := RestorePersistencePlan(plan, writer)
		if err != nil {
			t.Fatal(err)
		}
		if result.Completed || VerifyPersistenceRestore(result) {
			t.Fatalf("code 0x%02X must prevent completion: %+v", code, result)
		}
		if len(result.RequiredAreas) != 2 {
			t.Fatalf("required set must still be tracked: %v", result.RequiredAreas)
		}
		if len(result.AcknowledgedAreas) != 0 {
			t.Fatalf("failed area must not be acknowledged: %v", result.AcknowledgedAreas)
		}
		if result.Detail == "" {
			t.Fatal("failure must report a detail")
		}
	}
}

// A partial acknowledgement (first area succeeds, a later area fails) must not
// complete: the acknowledged set is a strict subset of the required set.
func TestPersistenceRestorePartialAcknowledgementDoesNotComplete(t *testing.T) {
	plan, _, _ := restoreReadyPlan(t)
	// Fixture order is [holding_registers, coils]; fail the second area.
	writer := &fakeRawIngestWriter{responses: map[PersistenceRawIngestArea]byte{RawIngestCoils: 0x21}}
	result, err := RestorePersistencePlan(plan, writer)
	if err != nil {
		t.Fatal(err)
	}
	if result.Completed || VerifyPersistenceRestore(result) {
		t.Fatalf("partial acknowledgement must not complete: %+v", result)
	}
	if len(result.AcknowledgedAreas) != 1 || result.AcknowledgedAreas[0] != "holding_registers" {
		t.Fatalf("only the acknowledged area may be recorded: %v", result.AcknowledgedAreas)
	}
	if result.Written != 1 {
		t.Fatalf("written count must reflect the acknowledged area only: %d", result.Written)
	}
}

// A transport error prevents completion and acknowledges nothing.
func TestPersistenceRestoreTransportErrorPreventsCompletion(t *testing.T) {
	plan, _, _ := restoreReadyPlan(t)
	writer := &fakeRawIngestWriter{err: errors.New("raw ingest unavailable")}
	result, err := RestorePersistencePlan(plan, writer)
	if err == nil {
		t.Fatal("transport error must surface")
	}
	if result.Completed || VerifyPersistenceRestore(result) || len(result.AcknowledgedAreas) != 0 {
		t.Fatalf("transport error must prevent completion: %+v", result)
	}
}

// The completion gate is strict: it requires a non-empty required set that
// exactly equals the acknowledged set (no missing, extra or duplicate areas).
func TestVerifyPersistenceRestoreGate(t *testing.T) {
	ok := PersistenceRestoreResult{
		RequiredAreas:     []string{"coils", "holding_registers"},
		AcknowledgedAreas: []string{"holding_registers", "coils"},
		Completed:         true,
	}
	if !VerifyPersistenceRestore(ok) {
		t.Fatal("identical sets must verify")
	}
	cases := map[string]PersistenceRestoreResult{
		"empty-required":     {RequiredAreas: nil, AcknowledgedAreas: nil},
		"missing-area":       {RequiredAreas: []string{"coils", "holding_registers"}, AcknowledgedAreas: []string{"coils"}},
		"extra-area":         {RequiredAreas: []string{"coils"}, AcknowledgedAreas: []string{"coils", "holding_registers"}},
		"foreign-area":       {RequiredAreas: []string{"coils"}, AcknowledgedAreas: []string{"holding_registers"}},
		"duplicate-required": {RequiredAreas: []string{"coils", "coils"}, AcknowledgedAreas: []string{"coils"}},
		"no-acknowledged":    {RequiredAreas: []string{"coils"}, AcknowledgedAreas: nil},
	}
	for name, result := range cases {
		if VerifyPersistenceRestore(result) {
			t.Fatalf("%s must not verify: %+v", name, result)
		}
	}
}

// A non-ready plan is refused, tracks the required set, and never completes.
func TestPersistenceRestoreNonReadyTracksRequiredWithoutCompletion(t *testing.T) {
	key, areas, source := loaderFixture(t)
	delete(source.snapshots, "coils")
	plan, err := LoadPersistenceSnapshots(key, true, true, areas, source)
	if err != nil {
		t.Fatal(err)
	}
	writer := &fakeRawIngestWriter{}
	result, err := RestorePersistencePlan(plan, writer)
	if err != nil {
		t.Fatal(err)
	}
	if result.Completed || VerifyPersistenceRestore(result) {
		t.Fatalf("non-ready plan must not complete: %+v", result)
	}
	if len(result.RequiredAreas) != 2 || len(writer.calls) != 0 {
		t.Fatalf("required set tracked, no writes: required=%v calls=%d", result.RequiredAreas, len(writer.calls))
	}
}

// restoreReadyPlanWithFlag builds a Ready plan carrying an authoritative State
// Sealing flag so the PERSIST-019 commit step is exercised.
func restoreReadyPlanWithFlag(t *testing.T, address uint16) PersistenceRestorePlan {
	t.Helper()
	plan, _, _ := restoreReadyPlan(t)
	plan.SealingFlag = &PersistenceSealingFlag{Address: address}
	return plan
}

// unsealCalls returns the Raw Ingest writes that target a single coil (the
// sealing flag write), i.e. the PERSIST-019 commit step.
func unsealCalls(writer *fakeRawIngestWriter) []rawIngestCall {
	var out []rawIngestCall
	for _, call := range writer.calls {
		if call.area == RawIngestCoils && call.count == 1 {
			out = append(out, call)
		}
	}
	return out
}

// PERSIST-019 self-check: the final successful restore action is an explicit
// write of the authoritative State Sealing flag to unsealed (1), at the
// configured flag address, and it happens only after every area is restored.
func TestPersistenceRestoreUnsealsOnlyAfterVerification(t *testing.T) {
	plan := restoreReadyPlanWithFlag(t, 4)
	writer := &fakeRawIngestWriter{}
	result, err := RestorePersistencePlan(plan, writer)
	if err != nil {
		t.Fatal(err)
	}
	if !result.Completed || !result.Committed {
		t.Fatalf("restore must complete and commit: %+v", result)
	}
	unseals := unsealCalls(writer)
	if len(unseals) != 1 {
		t.Fatalf("expected exactly one sealing flag write, got %d", len(unseals))
	}
	write := unseals[0]
	if write.start != 4 || write.count != 1 {
		t.Fatalf("unseal must target the configured flag address: %+v", write)
	}
	if len(write.payload) != 1 || write.payload[0] != 0x01 {
		t.Fatalf("unseal payload must be the sealed->unsealed value 1: %v", write.payload)
	}
	// The unseal must be the very last write, after all area restores.
	last := writer.calls[len(writer.calls)-1]
	if last.area != RawIngestCoils || last.count != 1 || last.start != 4 {
		t.Fatalf("unseal must be the final write action: %+v", last)
	}
}

// PERSIST-019 self-check: a failed area prevents completion and therefore the
// unseal/commit step never runs.
func TestPersistenceRestoreNoUnsealOnFailure(t *testing.T) {
	plan := restoreReadyPlanWithFlag(t, 4)
	writer := &fakeRawIngestWriter{responses: map[PersistenceRawIngestArea]byte{RawIngestHoldingRegisters: 0x21}}
	result, err := RestorePersistencePlan(plan, writer)
	if err != nil {
		t.Fatal(err)
	}
	if result.Completed || result.Committed {
		t.Fatalf("failed area must prevent completion and commit: %+v", result)
	}
	if len(unsealCalls(writer)) != 0 {
		t.Fatalf("no unseal write may occur on failure: %+v", writer.calls)
	}
}

// PERSIST-019 self-check: a non-zero unseal response leaves the restore
// completed but not committed (the memory stays sealed).
func TestPersistenceRestoreUnsealResponseFailureNotCommitted(t *testing.T) {
	plan := restoreReadyPlanWithFlag(t, 4)
	// unsealFailWriter rejects only the single-coil sealing flag write; every
	// area write still succeeds.
	failing := &unsealFailWriter{inner: &fakeRawIngestWriter{}, flagAddress: 4}
	result, err := RestorePersistencePlan(plan, failing)
	if err != nil {
		t.Fatal(err)
	}
	if !result.Completed {
		t.Fatalf("areas succeeded, restore must complete: %+v", result)
	}
	if result.Committed {
		t.Fatalf("a rejected unseal response must not commit: %+v", result)
	}
	if result.Detail == "" {
		t.Fatal("unseal failure must report a detail")
	}
}

// unsealFailWriter fails only the single-coil sealing flag write, letting all
// area writes succeed.
type unsealFailWriter struct {
	inner       *fakeRawIngestWriter
	flagAddress uint16
}

func (w *unsealFailWriter) WritePersistenceRawIngest(key PersistenceMemoryKey, area PersistenceRawIngestArea, start, count uint16, payload []byte) (byte, error) {
	if area == RawIngestCoils && count == 1 && start == w.flagAddress {
		w.inner.calls = append(w.inner.calls, rawIngestCall{key: key, area: area, start: start, count: count, payload: append([]byte(nil), payload...)})
		return 0x21, nil
	}
	return w.inner.WritePersistenceRawIngest(key, area, start, count, payload)
}

// PERSIST-019 self-check: without a configured State Sealing flag the restore
// completes but does not commit, and no unseal write is issued.
func TestPersistenceRestoreNoFlagCompletesWithoutCommit(t *testing.T) {
	plan, _, _ := restoreReadyPlan(t) // SealingFlag nil
	writer := &fakeRawIngestWriter{}
	result, err := RestorePersistencePlan(plan, writer)
	if err != nil {
		t.Fatal(err)
	}
	if !result.Completed {
		t.Fatalf("restore must complete: %+v", result)
	}
	if result.Committed {
		t.Fatalf("no flag configured means no commit: %+v", result)
	}
	if len(unsealCalls(writer)) != 0 {
		t.Fatalf("no unseal write without a configured flag: %+v", writer.calls)
	}
}

// PERSIST-019 self-check: the unseal payload encoder sets exactly the unsealed
// value 1 and never a different flag value.
func TestEncodePersistenceSealingFlag(t *testing.T) {
	on, err := encodePersistenceSealingFlag(true)
	if err != nil || len(on) != 1 || on[0] != 0x01 {
		t.Fatalf("unsealed encode must be 0x01: %v err=%v", on, err)
	}
	off, err := encodePersistenceSealingFlag(false)
	if err != nil || len(off) != 1 || off[0] != 0x00 {
		t.Fatalf("sealed encode must be 0x00: %v err=%v", off, err)
	}
}

// PERSIST-020 self-check: a committed restore has no failure classification and
// is not left sealed.
func TestPersistenceRestoreCommittedHasNoFailure(t *testing.T) {
	plan := restoreReadyPlanWithFlag(t, 4)
	writer := &fakeRawIngestWriter{}
	result, err := RestorePersistencePlan(plan, writer)
	if err != nil {
		t.Fatal(err)
	}
	if !result.Committed {
		t.Fatalf("restore must commit: %+v", result)
	}
	if result.Failure != PersistenceRestoreFailureNone || result.Sealed {
		t.Fatalf("a committed restore must have no failure and not be sealed: %+v", result)
	}
}

// brokenLoaderPlan builds a non-ready plan from a loader source the caller has
// corrupted, so the restore-failure classification can be exercised.
func brokenLoaderPlan(t *testing.T, mutate func(*fakeSnapshotSource)) PersistenceRestorePlan {
	t.Helper()
	key, areas, source := loaderFixture(t)
	mutate(source)
	plan, err := LoadPersistenceSnapshots(key, true, true, areas, source)
	if err != nil {
		t.Fatal(err)
	}
	return plan
}

// PERSIST-020 self-check: missing, corrupt and incompatible snapshots each keep
// the memory sealed with a deterministic, distinct failure classification.
func TestPersistenceRestoreSnapshotFailuresKeepSealed(t *testing.T) {
	cases := []struct {
		name    string
		mutate  func(*fakeSnapshotSource)
		failure PersistenceRestoreFailure
	}{
		{"missing", func(s *fakeSnapshotSource) { delete(s.snapshots, "coils") }, PersistenceRestoreFailureMissingSnapshot},
		{"corrupt", func(s *fakeSnapshotSource) { s.payloads["coils"][0] ^= 0x01 }, PersistenceRestoreFailureInvalidSnapshot},
		{"incompatible-version", func(s *fakeSnapshotSource) {
			m := s.snapshots["coils"]
			m.FormatVersion++
			s.snapshots["coils"] = m
		}, PersistenceRestoreFailureIncompatibleSnapshot},
		{"incompatible-layout", func(s *fakeSnapshotSource) {
			m := s.snapshots["coils"]
			m.Count = 10
			s.snapshots["coils"] = m
		}, PersistenceRestoreFailureIncompatibleSnapshot},
		{"incompatible-identity", func(s *fakeSnapshotSource) {
			m := s.snapshots["coils"]
			m.UnitID = 5
			s.snapshots["coils"] = m
		}, PersistenceRestoreFailureIncompatibleSnapshot},
	}
	for _, tc := range cases {
		plan := brokenLoaderPlan(t, tc.mutate)
		writer := &fakeRawIngestWriter{}
		result, err := RestorePersistencePlan(plan, writer)
		if err != nil {
			t.Fatalf("%s: unexpected error: %v", tc.name, err)
		}
		if result.Committed || result.Completed {
			t.Fatalf("%s: must not complete or commit: %+v", tc.name, result)
		}
		if result.Failure != tc.failure || !result.Sealed {
			t.Fatalf("%s: want failure %s sealed, got %+v", tc.name, tc.failure, result)
		}
		if result.Detail == "" {
			t.Fatalf("%s: failure must surface a reason", tc.name)
		}
		if len(writer.calls) != 0 {
			t.Fatalf("%s: a non-ready plan must write nothing: %d calls", tc.name, len(writer.calls))
		}
	}
}

// PERSIST-020 self-check: a Raw Ingest failure (transport error or non-zero
// response) keeps the memory sealed with the matching classification.
func TestPersistenceRestoreRawIngestFailureKeepsSealed(t *testing.T) {
	plan := restoreReadyPlanWithFlag(t, 4)
	writer := &fakeRawIngestWriter{responses: map[PersistenceRawIngestArea]byte{RawIngestHoldingRegisters: 0x21}}
	result, err := RestorePersistencePlan(plan, writer)
	if err != nil {
		t.Fatal(err)
	}
	if result.Failure != PersistenceRestoreFailureRawIngestResponse || !result.Sealed || result.Committed {
		t.Fatalf("non-zero response must keep sealed: %+v", result)
	}
	if len(unsealCalls(writer)) != 0 {
		t.Fatal("no unseal may occur after a Raw Ingest failure")
	}

	plan2 := restoreReadyPlanWithFlag(t, 4)
	writer2 := &fakeRawIngestWriter{err: errors.New("raw ingest unavailable")}
	result2, err := RestorePersistencePlan(plan2, writer2)
	if err == nil {
		t.Fatal("transport error must surface")
	}
	if result2.Failure != PersistenceRestoreFailureRawIngestWrite || !result2.Sealed || result2.Committed {
		t.Fatalf("transport error must keep sealed: %+v", result2)
	}
}

// PERSIST-020 self-check: a completed-but-uncommitted restore (no flag, or a
// rejected unseal) is deterministically classified and left sealed.
func TestPersistenceRestoreCommitFailuresKeepSealed(t *testing.T) {
	// No sealing flag configured.
	plan, _, _ := restoreReadyPlan(t)
	result, err := RestorePersistencePlan(plan, &fakeRawIngestWriter{})
	if err != nil {
		t.Fatal(err)
	}
	if !result.Completed || result.Committed {
		t.Fatalf("expected completed-not-committed: %+v", result)
	}
	if result.Failure != PersistenceRestoreFailureNoSealingFlag || !result.Sealed {
		t.Fatalf("missing flag must be classified and sealed: %+v", result)
	}

	// Unseal rejected with a non-zero response.
	plan2 := restoreReadyPlanWithFlag(t, 4)
	result2, err := RestorePersistencePlan(plan2, &unsealFailWriter{inner: &fakeRawIngestWriter{}, flagAddress: 4})
	if err != nil {
		t.Fatal(err)
	}
	if !result2.Completed || result2.Committed {
		t.Fatalf("expected completed-not-committed: %+v", result2)
	}
	if result2.Failure != PersistenceRestoreFailureUnsealResponse || !result2.Sealed {
		t.Fatalf("rejected unseal must be classified and sealed: %+v", result2)
	}

	// Unseal transport error.
	plan3 := restoreReadyPlanWithFlag(t, 4)
	writer3 := &unsealErrWriter{inner: &fakeRawIngestWriter{}, flagAddress: 4, err: errors.New("unseal transport down")}
	result3, err := RestorePersistencePlan(plan3, writer3)
	if err == nil {
		t.Fatal("unseal transport error must surface")
	}
	if !result3.Completed || result3.Committed {
		t.Fatalf("expected completed-not-committed: %+v", result3)
	}
	if result3.Failure != PersistenceRestoreFailureUnsealWrite || !result3.Sealed {
		t.Fatalf("unseal transport error must be classified and sealed: %+v", result3)
	}
}

// unsealErrWriter fails only the single-coil sealing flag write with a transport
// error, letting all area writes succeed.
type unsealErrWriter struct {
	inner       *fakeRawIngestWriter
	flagAddress uint16
	err         error
}

func (w *unsealErrWriter) WritePersistenceRawIngest(key PersistenceMemoryKey, area PersistenceRawIngestArea, start, count uint16, payload []byte) (byte, error) {
	if area == RawIngestCoils && count == 1 && start == w.flagAddress {
		return 0, w.err
	}
	return w.inner.WritePersistenceRawIngest(key, area, start, count, payload)
}

// PERSIST-020 self-check: the disabled and empty plan states are classified and
// leave the memory sealed.
func TestPersistenceRestoreDisabledAndEmptyKeepSealed(t *testing.T) {
	key, areas, source := loaderFixture(t)
	disabled, err := LoadPersistenceSnapshots(key, false, true, areas, source)
	if err != nil {
		t.Fatal(err)
	}
	result, err := RestorePersistencePlan(disabled, &fakeRawIngestWriter{})
	if err != nil {
		t.Fatal(err)
	}
	if result.Failure != PersistenceRestoreFailureDisabled || !result.Sealed {
		t.Fatalf("disabled must be classified and sealed: %+v", result)
	}

	empty, err := LoadPersistenceSnapshots(key, true, true, nil, source)
	if err != nil {
		t.Fatal(err)
	}
	result2, err := RestorePersistencePlan(empty, &fakeRawIngestWriter{})
	if err != nil {
		t.Fatal(err)
	}
	if result2.Failure != PersistenceRestoreFailureEmpty || !result2.Sealed {
		t.Fatalf("empty must be classified and sealed: %+v", result2)
	}
}

// The failure classification has a stable string form for each value.
func TestPersistenceRestoreFailureString(t *testing.T) {
	want := map[PersistenceRestoreFailure]string{
		PersistenceRestoreFailureNone:                 "none",
		PersistenceRestoreFailureDisabled:             "disabled",
		PersistenceRestoreFailureUnsealed:             "unsealed",
		PersistenceRestoreFailureEmpty:                "empty",
		PersistenceRestoreFailureMissingSnapshot:      "missing_snapshot",
		PersistenceRestoreFailureInvalidSnapshot:      "invalid_snapshot",
		PersistenceRestoreFailureIncompatibleSnapshot: "incompatible_snapshot",
		PersistenceRestoreFailureAreaNotReady:         "area_not_ready",
		PersistenceRestoreFailureUnknownArea:          "unknown_area",
		PersistenceRestoreFailureRawIngestWrite:       "raw_ingest_write",
		PersistenceRestoreFailureRawIngestResponse:    "raw_ingest_response",
		PersistenceRestoreFailureIncomplete:           "incomplete",
		PersistenceRestoreFailureNoSealingFlag:        "no_sealing_flag",
		PersistenceRestoreFailureUnsealWrite:          "unseal_write",
		PersistenceRestoreFailureUnsealResponse:       "unseal_response",
	}
	for value, text := range want {
		if got := value.String(); got != text {
			t.Fatalf("failure %d string = %q, want %q", value, got, text)
		}
	}
	if got := PersistenceRestoreFailure(999).String(); got != "unknown" {
		t.Fatalf("unrecognized failure must map to unknown, got %q", got)
	}
}

// A plan whose area is not ready is classified as area_not_ready and sealed.
func TestPersistenceRestoreAreaNotReadyClassification(t *testing.T) {
	plan := PersistenceRestorePlan{
		Key: PersistenceMemoryKey{Port: 1, UnitID: 1}, Enabled: true, Sealed: true, State: PersistenceRestoreReady,
		Areas: []PersistenceAreaRestore{{Area: "holding_registers", Kind: PersistenceRegisters, Start: 0, Count: 2, Outcome: PersistenceRestoreInvalid, Detail: "corrupt"}},
	}
	result, err := RestorePersistencePlan(plan, &fakeRawIngestWriter{})
	if err != nil {
		t.Fatal(err)
	}
	if result.Failure != PersistenceRestoreFailureInvalidSnapshot || !result.Sealed {
		t.Fatalf("non-ready area must be classified from its outcome and sealed: %+v", result)
	}
}
