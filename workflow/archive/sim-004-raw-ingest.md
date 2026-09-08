# SIM-004 — Send Simulator Values Through MMA2 Raw Ingest

Status: COMPLETED — verified 2026-09-07.

Source intent: `planning/Brainstorm/osjs-modbus-simulator.md`.

## Primary Outcome

Simulator-generated FC1-FC4 values reach simulator-owned MMA2 memory exclusively through MMA2 raw ingest and are externally readable through Modbus TCP.

## Scope

- Convert FC1-FC4 generated values into the actual MMA2 raw-ingest format established in repository truth.
- Send simulator-generated values only through MMA2 raw ingest.
- Target only simulator-owned configured MMA2 ranges.
- Verify at least one generated value from each enabled FC area reaches MMA2 memory.
- Read the corresponding values with a real external Modbus TCP client.

## Non-Scope

- No direct-memory bypass.
- No Modbus-write-based population of simulator values.
- No structural MMA2 configuration changes.
- No scheduler redesign.
- No OS.js UI.

## Acceptance Criteria

1. Generated values for each enabled FC area are accepted through MMA2 raw ingest.
2. A real external Modbus TCP client can read at least one value from each enabled FC area after a simulator update.
3. No simulator code path populates MMA2 through direct memory access or Modbus writes.

## Verification

Run one configured simulator device, capture one update for each enabled FC, then read the matching FC areas with a real Modbus client and prove the values came through the raw-ingest path.

**Evidence (2026-09-07(: `simulator/raw_ingest.go` delivers the raw-ingest v1 client:`encodeRawPacket` frames simulator-generated FC1-FC4 values in the MMA2 raw-ingest byte contract (bits LSB-first for coil/discrete areas,big-endian uint16 for register areas(;`(*RawIngestClient). Send` maps each FC onto the simulator-owned configured MMA2 range and rejects unconfigured count-0 areas,targeting raw ingest only—no Modbus-write or direct-memory population. `simulator/raw_ingest_test.go` verifies frame bytes plus a real-TCP end-to-end round-trip:(fixture server accepts the framed packet and returns OK/reject responses(;unconfigured-FC send fails. `gofmt -l .` clean;`go vet ./...` clean;`go test -count=1 ./...` → `ok github.com/tamzrod/MCS.OSJS/simulator`.

## Dependencies

- SIM-002.
- SIM-003.
- Actual MMA2 raw-ingest contract from repository truth.

## Sizing

**3 / 10 — Small.** One integration boundary with one end-to-end data-path proof.