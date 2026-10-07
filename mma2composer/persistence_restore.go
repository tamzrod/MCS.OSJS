package mma2composer

import "fmt"

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
		response, err := writer.WritePersistenceRawIngest(plan.Key, code, area.Start, area.Count, area.Payload)
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
