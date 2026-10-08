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

// PersistenceRestoreResult reports what one memory's restore did. Completed is
// true only when every configured area was written through Raw Ingest and every
// response was the success code. Nothing is unsealed here: the final unseal is
// a later, separate commit step.
type PersistenceRestoreResult struct {
	Key       PersistenceMemoryKey
	State     PersistenceRestoreOutcome
	Areas     int
	Written   int
	Completed bool
	Detail    string
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
func RestorePersistencePlan(plan PersistenceRestorePlan, writer PersistenceRawIngestWriter) (PersistenceRestoreResult, error) {
	result := PersistenceRestoreResult{Key: plan.Key, State: plan.State, Areas: len(plan.Areas)}
	if plan.State != PersistenceRestoreReady {
		result.Detail = fmt.Sprintf("restore plan is %s, not ready", plan.State)
		return result, nil
	}
	if len(plan.Areas) == 0 {
		result.State = PersistenceRestoreEmpty
		result.Detail = "restore plan has no configured areas"
		return result, nil
	}
	if writer == nil {
		return PersistenceRestoreResult{}, fmt.Errorf("persistence restore requires a raw ingest writer")
	}
	for _, area := range plan.Areas {
		if area.Outcome != PersistenceRestoreReady {
			result.Detail = fmt.Sprintf("area %q is %s, not ready", area.Area, area.Outcome)
			return result, nil
		}
		code, ok := rawIngestArea(area.Area)
		if !ok {
			result.Detail = fmt.Sprintf("area %q has no raw ingest mapping", area.Area)
			return result, nil
		}
		payload := area.Payload
		if plan.SealingFlag != nil && area.Area == "coils" {
			protected, err := forcePersistenceSealingFlag(area.Start, area.Count, payload, *plan.SealingFlag)
			if err != nil {
				result.Detail = fmt.Sprintf("area %q seal-flag protection failed: %v", area.Area, err)
				return result, nil
			}
			payload = protected
		}
		response, err := writer.WritePersistenceRawIngest(plan.Key, code, area.Start, area.Count, payload)
		if err != nil {
			result.Detail = fmt.Sprintf("area %q raw ingest write failed: %v", area.Area, err)
			return result, err
		}
		if response != PersistenceRawIngestOK {
			result.Detail = fmt.Sprintf("area %q raw ingest response 0x%02X aborted restore", area.Area, response)
			return result, nil
		}
		result.Written++
	}
	result.Completed = true
	return result, nil
}
