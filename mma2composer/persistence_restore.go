package mma2composer

import (
	"fmt"
	"strings"
)

// PersistenceRawIngestArea names one Raw Ingest v1 area. The numeric values are
// the wire contract (coils 1, discrete inputs 2, holding registers 3, input
// registers 4) and must not be changed here.
type PersistenceRawIngestArea uint8

const (
	RawIngestCoils PersistenceRawIngestArea = iota + 1
	RawIngestDiscreteInputs
	RawIngestHoldingRegisters
	RawIngestInputRegisters
)

// rawIngestArea maps a canonical memory area key to its Raw Ingest v1 area code.
func rawIngestArea(area string) (PersistenceRawIngestArea, bool) {
	switch area {
	case "coils":
		return RawIngestCoils, true
	case "discrete_inputs":
		return RawIngestDiscreteInputs, true
	case "holding_registers":
		return RawIngestHoldingRegisters, true
	case "input_registers":
		return RawIngestInputRegisters, true
	default:
		return 0, false
	}
}

// PersistenceRawIngestWriter submits one Raw Ingest v1 write and reports the
// single-byte response code. Implementations carry the existing protocol; the
// restore path never re-encodes or extends it. A response code of
// PersistenceRawIngestOK means the write committed; any other code means no
// write occurred.
type PersistenceRawIngestWriter interface {
	WritePersistenceRawIngest(key PersistenceMemoryKey, area PersistenceRawIngestArea, start, count uint16, payload []byte) (byte, error)
}

// PersistenceRawIngestOK is the Raw Ingest v1 success response code (0x00).
const PersistenceRawIngestOK = byte(0x00)

// persistenceUnsealValue is the State Sealing flag value that means unsealed
// (MMA2 state sealing: 0 = sealed, 1 = unsealed). Writing it is the explicit
// final commit action of a completed restore.
const persistenceUnsealValue = true

// PersistenceSealingFlag is the authoritative State Sealing flag location: the
// coils bit that encodes sealed (0) or unsealed (1). It is derived from the
// configured State Sealing block, which remains the single source of truth; no
// second sealing location is introduced. Address is the memory-wide coil
// address (the State Sealing address is validated against the coil area).
type PersistenceSealingFlag struct {
	Address uint16
}

// PersistenceSealingFlagFromExtra derives the authoritative sealing flag
// location from a memory's configuration block. It returns ok=false when state
// sealing is absent or disabled (nothing to protect), or when the block does
// not describe a valid coil sealing address. It mirrors MMA2's sealing
// enablement rule (absent block disabled; absent enabled flag defaults to
// enabled; explicit flag wins) and requires area "coil" and a non-empty address
// field, so it never invents a second source of truth. It is read-only.
func PersistenceSealingFlagFromExtra(extra map[string]interface{}) (PersistenceSealingFlag, bool) {
	block, ok := extra["state_sealing"].(map[string]interface{})
	if !ok {
		return PersistenceSealingFlag{}, false
	}
	if enabled, present := block["enabled"]; present {
		flag, ok := enabled.(bool)
		if !ok || !flag {
			return PersistenceSealingFlag{}, false
		}
	}
	area, _ := block["area"].(string)
	if strings.ToLower(strings.TrimSpace(area)) != "coil" {
		return PersistenceSealingFlag{}, false
	}
	addr, ok := stateSealingAddress(block["address"])
	if !ok {
		return PersistenceSealingFlag{}, false
	}
	return PersistenceSealingFlag{Address: addr}, true
}

// stateSealingAddress reads the sealing address from either an int or a uint16
// YAML value.
func stateSealingAddress(value interface{}) (uint16, bool) {
	switch v := value.(type) {
	case int:
		if v < 0 || v > 0xFFFF {
			return 0, false
		}
		return uint16(v), true
	case uint16:
		return v, true
	default:
		return 0, false
	}
}

// forcePersistenceSealingFlag clears the sealing flag bit within a coils payload
// that begins at the given area start, so a restored snapshot can never unseal
// the memory mid-restore. It writes only that one bit (sealed = 0); every other
// restored bit is preserved. When the flag lies outside the written range the
// write cannot touch it, so the payload is returned unchanged. A flag inside the
// range but beyond the payload length is an integrity failure and returns an
// error. The payload is never mutated in place.
func forcePersistenceSealingFlag(areaStart, areaCount uint16, payload []byte, flag PersistenceSealingFlag) ([]byte, error) {
	out := append([]byte(nil), payload...)
	if uint32(flag.Address) < uint32(areaStart) || uint32(flag.Address) >= uint32(areaStart)+uint32(areaCount) {
		return out, nil
	}
	offset := uint32(flag.Address) - uint32(areaStart)
	needed := int(offset/8) + 1
	if len(out) < needed {
		return nil, fmt.Errorf("coil payload length %d cannot hold state sealing address %d", len(payload), flag.Address)
	}
	out[offset/8] &^= 1 << (offset % 8)
	return out, nil
}

// PersistenceRestoreFailure classifies, deterministically, why a startup
// persistence restore did not safely commit. A non-zero value means the memory
// was left sealed; PersistenceRestoreFailureNone is used only when the restore
// completed and committed.
type PersistenceRestoreFailure int

const (
	// PersistenceRestoreFailureNone: no failure; the restore committed.
	PersistenceRestoreFailureNone PersistenceRestoreFailure = iota
	// PersistenceRestoreFailureDisabled: persistence is not enabled.
	PersistenceRestoreFailureDisabled
	// PersistenceRestoreFailureUnsealed: persistence is enabled but the memory
	// was not sealed at startup, so restore must not proceed.
	PersistenceRestoreFailureUnsealed
	// PersistenceRestoreFailureEmpty: no area is configured to restore.
	PersistenceRestoreFailureEmpty
	// PersistenceRestoreFailureMissingSnapshot: no snapshot is present for a
	// configured area.
	PersistenceRestoreFailureMissingSnapshot
	// PersistenceRestoreFailureInvalidSnapshot: a snapshot is corrupt, incomplete
	// or unreadable.
	PersistenceRestoreFailureInvalidSnapshot
	// PersistenceRestoreFailureIncompatibleSnapshot: a snapshot does not match
	// the configured format/identity/area/layout.
	PersistenceRestoreFailureIncompatibleSnapshot
	// PersistenceRestoreFailureAreaNotReady: a configured area in the plan is not
	// ready to restore.
	PersistenceRestoreFailureAreaNotReady
	// PersistenceRestoreFailureUnknownArea: a configured area has no Raw Ingest
	// mapping.
	PersistenceRestoreFailureUnknownArea
	// PersistenceRestoreFailureRawIngestWrite: the Raw Ingest write for an area
	// failed at the transport level.
	PersistenceRestoreFailureRawIngestWrite
	// PersistenceRestoreFailureRawIngestResponse: a Raw Ingest write for an area
	// was rejected with a non-zero response code.
	PersistenceRestoreFailureRawIngestResponse
	// PersistenceRestoreFailureIncomplete: verification found a required area
	// missing or failed.
	PersistenceRestoreFailureIncomplete
	// PersistenceRestoreFailureNoSealingFlag: the restore completed but no State
	// Sealing flag is configured to unseal.
	PersistenceRestoreFailureNoSealingFlag
	// PersistenceRestoreFailureUnsealWrite: the unseal write failed at the
	// transport level.
	PersistenceRestoreFailureUnsealWrite
	// PersistenceRestoreFailureUnsealResponse: the unseal write was rejected with
	// a non-zero response code.
	PersistenceRestoreFailureUnsealResponse
)

func (f PersistenceRestoreFailure) String() string {
	switch f {
	case PersistenceRestoreFailureNone:
		return "none"
	case PersistenceRestoreFailureDisabled:
		return "disabled"
	case PersistenceRestoreFailureUnsealed:
		return "unsealed"
	case PersistenceRestoreFailureEmpty:
		return "empty"
	case PersistenceRestoreFailureMissingSnapshot:
		return "missing_snapshot"
	case PersistenceRestoreFailureInvalidSnapshot:
		return "invalid_snapshot"
	case PersistenceRestoreFailureIncompatibleSnapshot:
		return "incompatible_snapshot"
	case PersistenceRestoreFailureAreaNotReady:
		return "area_not_ready"
	case PersistenceRestoreFailureUnknownArea:
		return "unknown_area"
	case PersistenceRestoreFailureRawIngestWrite:
		return "raw_ingest_write"
	case PersistenceRestoreFailureRawIngestResponse:
		return "raw_ingest_response"
	case PersistenceRestoreFailureIncomplete:
		return "incomplete"
	case PersistenceRestoreFailureNoSealingFlag:
		return "no_sealing_flag"
	case PersistenceRestoreFailureUnsealWrite:
		return "unseal_write"
	case PersistenceRestoreFailureUnsealResponse:
		return "unseal_response"
	default:
		return "unknown"
	}
}

// persistenceRestoreFailureForArea maps a loader area outcome to its restore
// failure classification.
func persistenceRestoreFailureForArea(outcome PersistenceRestoreOutcome) PersistenceRestoreFailure {
	switch outcome {
	case PersistenceRestoreMissing:
		return PersistenceRestoreFailureMissingSnapshot
	case PersistenceRestoreIncompatible:
		return PersistenceRestoreFailureIncompatibleSnapshot
	case PersistenceRestoreInvalid:
		return PersistenceRestoreFailureInvalidSnapshot
	case PersistenceRestoreDisabled:
		return PersistenceRestoreFailureDisabled
	case PersistenceRestoreUnsealed:
		return PersistenceRestoreFailureUnsealed
	case PersistenceRestoreEmpty:
		return PersistenceRestoreFailureEmpty
	default:
		return PersistenceRestoreFailureAreaNotReady
	}
}

// PersistenceRestoreResult reports what one memory's restore did. RequiredAreas
// is the deterministic set of configured area keys that must be restored, in the
// plan's canonical order; AcknowledgedAreas is the subset that was written and
// acknowledged with the Raw Ingest success code. Completed is true only when
// both sets are identical — every required area was written and acknowledged —
// so a missing or failed area can never be reported as a completed restore.
//
// Committed is the separate, later final step (PERSIST-019): it is true only
// when the authoritative State Sealing flag was written to unsealed (1) after
// verification success. A restore that is Completed but not Committed has been
// written but remains sealed, so it is still not exposed to Modbus.
//
// Failure is the deterministic classification of why the restore did not
// commit (PERSIST-020); it is PersistenceRestoreFailureNone only when Committed
// is true. Detail carries a human-readable reason and Sealed reports whether the
// memory was left sealed as a result.
type PersistenceRestoreResult struct {
	Key               PersistenceMemoryKey
	State             PersistenceRestoreOutcome
	Areas             int
	Written           int
	RequiredAreas     []string
	AcknowledgedAreas []string
	Completed         bool
	Committed         bool
	Failure           PersistenceRestoreFailure
	Sealed            bool
	Detail            string
}

// failedPersistenceRestore marks a result as a deterministic, sealed failure and
// returns it. Every non-committing path in RestorePersistencePlan goes through
// here, so a failure always leaves the memory sealed and carries a classified
// reason. It never fabricates a default, retries or unseals.
func failedPersistenceRestore(result PersistenceRestoreResult, failure PersistenceRestoreFailure, detail string) PersistenceRestoreResult {
	result.Failure = failure
	result.Sealed = true
	result.Detail = detail
	return result
}

// persistenceRequiredAreas returns the deterministic required-area set for a
// plan in plan order. It never mutates the plan.
func persistenceRequiredAreas(plan PersistenceRestorePlan) []string {
	required := make([]string, 0, len(plan.Areas))
	for _, area := range plan.Areas {
		required = append(required, area.Area)
	}
	return required
}

// VerifyPersistenceRestore reports whether a restore completed: true only when
// there is a non-empty deterministic required-area set and it is identical to
// the acknowledged-area set (same size and same members). This is the
// restore-completion gate; it is a pure read-only check over already-observed
// restore results and never writes, commits or unseals.
func VerifyPersistenceRestore(result PersistenceRestoreResult) bool {
	if len(result.RequiredAreas) == 0 || len(result.RequiredAreas) != len(result.AcknowledgedAreas) {
		return false
	}
	required := make(map[string]int, len(result.RequiredAreas))
	for _, area := range result.RequiredAreas {
		required[area]++
	}
	for _, area := range result.AcknowledgedAreas {
		if required[area] == 0 {
			return false
		}
		required[area]--
	}
	for _, count := range required {
		if count != 0 {
			return false
		}
	}
	return true
}

// RestorePersistencePlan writes a validated restore plan back into the matching
// MMA2 memory through the existing Raw Ingest v1 contract. Each ready area is
// submitted to the same area/start/count identity it was loaded from; the Raw
// Ingest area code is derived from the area key, never authored independently.
//
// A plan that is not Ready is refused without any write, so a partially valid
// snapshot set can never be exposed. Every Raw Ingest response is checked; the
// first non-success response, transport error or unknown area aborts the restore
// immediately and is reported. The restore writes only the areas named by the
// plan and adds no cross-area mirroring, and it never unseals memory.
//
// When the plan carries an authoritative State Sealing flag (PERSIST-017), the
// sealing bit is forced to sealed (0) within the restored coils payload before
// it is written, so a snapshot captured while unsealed cannot unseal the memory
// mid-restore. Only that one bit is altered; the flag location comes solely from
// configuration and no second sealing source of truth is introduced.
//
// Restore completion is gated (PERSIST-018): the result records the
// deterministic required-area set and the acknowledged-area set, and Completed
// is emitted only when VerifyPersistenceRestore confirms they are identical. A
// missing or failed area therefore prevents completion.
//
// The final commit action (PERSIST-019) is an explicit write of the
// authoritative State Sealing flag to unsealed (1), performed only after
// verification success. The flag location comes solely from configuration. Until
// this step the memory remains sealed, so no earlier restore step can expose it
// to Modbus; Committed reports whether this final unseal committed.
//
// Restore-failure behavior (PERSIST-020): every path that does not safely commit
// leaves the memory sealed and returns a deterministic Failure classification
// plus a human-readable Detail. Missing, corrupt or incompatible snapshots, a
// Raw Ingest or verification failure, and an unseal/commit failure all keep the
// memory sealed; the restore never fabricates a default, retries or unseals on
// failure.
func RestorePersistencePlan(plan PersistenceRestorePlan, writer PersistenceRawIngestWriter) (PersistenceRestoreResult, error) {
	result := PersistenceRestoreResult{
		Key:           plan.Key,
		State:         plan.State,
		Areas:         len(plan.Areas),
		RequiredAreas: persistenceRequiredAreas(plan),
	}
	if plan.State != PersistenceRestoreReady {
		return failedPersistenceRestore(result, persistenceRestoreFailureForArea(plan.State),
			fmt.Sprintf("restore plan is %s, not ready; memory remains sealed", plan.State)), nil
	}
	if len(plan.Areas) == 0 {
		result.State = PersistenceRestoreEmpty
		return failedPersistenceRestore(result, PersistenceRestoreFailureEmpty,
			"restore plan has no configured areas; memory remains sealed"), nil
	}
	if writer == nil {
		return PersistenceRestoreResult{}, fmt.Errorf("persistence restore requires a raw ingest writer")
	}
	for _, area := range plan.Areas {
		if area.Outcome != PersistenceRestoreReady {
			return failedPersistenceRestore(result, persistenceRestoreFailureForArea(area.Outcome),
				fmt.Sprintf("area %q is %s, not ready; memory remains sealed", area.Area, area.Outcome)), nil
		}
		code, ok := rawIngestArea(area.Area)
		if !ok {
			return failedPersistenceRestore(result, PersistenceRestoreFailureUnknownArea,
				fmt.Sprintf("area %q has no raw ingest mapping; memory remains sealed", area.Area)), nil
		}
		payload := area.Payload
		if plan.SealingFlag != nil && area.Area == "coils" {
			protected, err := forcePersistenceSealingFlag(area.Start, area.Count, payload, *plan.SealingFlag)
			if err != nil {
				return failedPersistenceRestore(result, PersistenceRestoreFailureInvalidSnapshot,
					fmt.Sprintf("area %q seal-flag protection failed: %v; memory remains sealed", area.Area, err)), nil
			}
			payload = protected
		}
		response, err := writer.WritePersistenceRawIngest(plan.Key, code, area.Start, area.Count, payload)
		if err != nil {
			failed := failedPersistenceRestore(result, PersistenceRestoreFailureRawIngestWrite,
				fmt.Sprintf("area %q raw ingest write failed: %v; memory remains sealed", area.Area, err))
			return failed, err
		}
		if response != PersistenceRawIngestOK {
			return failedPersistenceRestore(result, PersistenceRestoreFailureRawIngestResponse,
				fmt.Sprintf("area %q raw ingest response 0x%02X aborted restore; memory remains sealed", area.Area, response)), nil
		}
		result.Written++
		result.AcknowledgedAreas = append(result.AcknowledgedAreas, area.Area)
	}
	// Restore-completion gate: success is emitted only when the acknowledged
	// area set exactly covers the deterministic required-area set.
	result.Completed = VerifyPersistenceRestore(result)
	if !result.Completed {
		return failedPersistenceRestore(result, PersistenceRestoreFailureIncomplete,
			"restore did not acknowledge every required persistence area; memory remains sealed"), nil
	}
	// Atomic unseal / commit step (PERSIST-019): only after verification success,
	// write the authoritative State Sealing flag to unsealed (1) as the explicit
	// final commit action. The location comes solely from configuration; no
	// earlier step can expose memory to Modbus because it stays sealed until here.
	if plan.SealingFlag == nil {
		return failedPersistenceRestore(result, PersistenceRestoreFailureNoSealingFlag,
			"restore completed but no state sealing flag is configured to unseal; memory remains sealed"), nil
	}
	unsealed, err := encodePersistenceSealingFlag(persistenceUnsealValue)
	if err != nil {
		return failedPersistenceRestore(result, PersistenceRestoreFailureUnsealWrite,
			fmt.Sprintf("state sealing unseal encoding failed: %v; memory remains sealed", err)), nil
	}
	response, err := writer.WritePersistenceRawIngest(plan.Key, RawIngestCoils, plan.SealingFlag.Address, 1, unsealed)
	if err != nil {
		failed := failedPersistenceRestore(result, PersistenceRestoreFailureUnsealWrite,
			fmt.Sprintf("state sealing unseal write failed: %v; memory remains sealed", err))
		return failed, err
	}
	if response != PersistenceRawIngestOK {
		return failedPersistenceRestore(result, PersistenceRestoreFailureUnsealResponse,
			fmt.Sprintf("state sealing unseal response 0x%02X aborted commit; memory remains sealed", response)), nil
	}
	result.Committed = true
	return result, nil
}

// encodePersistenceSealingFlag encodes the single coils write payload for one
// State Sealing flag value. The location comes from the authoritative State
// Sealing configuration; this function never introduces an alternate flag or
// location.
func encodePersistenceSealingFlag(value bool) ([]byte, error) {
	return EncodePersistenceBits([]bool{value}, 1)
}
