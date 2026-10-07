# PERSIST-001 — Persistence Configuration Schema

Status: CODE COMPLETE — 2026-10-07 (OpenHands JR DEV); awaiting separate independent TEST/VERIFY (PERSIST-011)
Stage: CODE
Owner: OpenHands JR DEV
Previous: none
Next: PERSIST-002

## Primary outcome
Add `persistence.enabled` per MMA2 memory instance without adding persistence-owned port, unit, area, start or count fields.

## Scope
Implement or verify only the primary outcome above against the current authoritative MCS.OSJS source and existing MMA2 contracts. Preserve the Port → Unit ID → Memory authority model and existing State Sealing/RBE semantics unless this packet explicitly says otherwise.

## Non-scope
Do not implement RBE derivation, filesystem persistence, restore, or runtime behavior.

## Acceptance
1. Persistence is represented per existing Port → Unit ID → Memory identity.
2. `persistence.enabled` round-trips without changing existing configurations.
3. No duplicate memory-range source of truth is introduced.

## Evidence / handoff
Coding owner records the exact source diff, changed paths, targeted self-checks and resulting source checkpoint. Do not claim independent TEST/VERIFY PASS. STOP after this task.

## Dependencies
Requires genuine completion evidence for no predecessor. Being present in `workflow/active_work/` authorizes this written scope but does not bypass predecessor or current-task selection gates.

## Sizing
1/0/1/1/1=4; tightly coupled schema/round-trip change.

## Coding evidence (PERSIST-001, OpenHands JR DEV)

- Source checkpoint base: `a82a2d038f1d7cecb4127208595523dc4954ef28` (`main`, clean, = `origin/main` at run).
- Changed/added paths (exact working-tree state):
  - M `mma2composer/composer.go` (blob `099148366f50a532bfa49a05e334783dba8cb0b6`) — added `Persistence{Enabled *bool}` and `Memory.Persistence *Persistence`.
  - M `simulator/device.go` (blob `1949b5144610e37b6b924464b60fadc93c2174b6`) — added `MMA2Params.Persistence *mma2composer.Persistence`.
  - M `simulator/mma2_config.go` (blob `2a19e684fd1cc8a532afd23ec9d1bf6e905d96ec`) — `memoryFromMMA2Params` copies the typed pointer; `inheritMemorySettings` inherits it when absent.
  - A `simulator/persistence_config_schema_test.go` (blob `744c5884ebb38ebba022294e9562572e645feab9`) — focused self-check.
- Self-check command (Go 1.25.3, sandbox-local): `cd simulator && go test -run TestPersistenceConfigurationSchemaRoundTrip .` → `ok`, exit 0.
- Bounded regression (readonly, no manifest mutation): `mma2composer`, `simulator`, `replicator` full `go test -mod=readonly ./...` → all `ok`, exit 0. `gofmt -l` clean; `go vet` exit 0.
- Design notes: `persistence.enabled` is typed per memory and serialized as its own key; it does **not** enter the generic `Extra` inline map, so it cannot become a range source of truth. Absent => off; explicit `enabled: false` preserved. Persistence fields are not yet emitted into the composed MMA2 effective config because `configvalidate` rejects unknown memory keys (PERSIST-003 derives system RBE; out of scope here).
- This is a JR DEV self-check only. It is **not** an independent TEST/VERIFY PASS; PERSIST-011 remains the separate independent JR gate.

## CWAL
CODE COMPLETE at `faa33929429a0382a64b78bc773f0074e11cb6b1`. This packet is no longer the current ACTIVE assignment. Its independent TEST/VERIFY coverage remains deferred to PERSIST-011.
