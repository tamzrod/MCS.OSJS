# PERSIST-013 — Snapshot File Format

Status: CODE COMPLETE — 2026-10-08; delivered via GitHub connector recovery; awaiting independent TEST/VERIFY (PERSIST-022)
Stage: CODE
Owner: OpenHands JR DEV
Previous: PERSIST-012
Next: PERSIST-014

## Primary outcome
Define and implement deterministic raw snapshot encoding for all supported Modbus memory areas.

## Scope
Implement or verify only the primary outcome above against the current authoritative MCS.OSJS source and existing MMA2 contracts. Preserve the Port → Unit ID → Memory authority model and existing State Sealing/RBE semantics unless this packet explicitly says otherwise.

## Non-scope
Do not add compatibility manifest or restore sequencing.

## Acceptance
1. Coils/discrete inputs use LSB-first packed bits.
2. Holding/input registers use big-endian uint16 words.
3. Encoding/decoding is deterministic for configured start/count.

## Evidence / handoff
Coding owner records the exact source diff, changed paths, targeted self-checks and resulting source checkpoint. Do not claim independent TEST/VERIFY PASS. STOP after this task.

## Dependencies
Requires genuine completion evidence for PERSIST-012. Being present in `workflow/active_work/` authorizes this written scope but does not bypass predecessor or current-task selection gates.

## Sizing
1/0/2/1/1=5.

## Coding evidence / recovery

OpenHands completed and self-checked this task locally but transport failed because its runtime GitHub credential returned HTTP 401. Local-only commits were `9343d9e70bdb2d557b837b3d9f3d6604bf54f9a1` (product) and `ab27b27e4d9734b9bb28a5a0c38d2582d4cba7a2` (workflow); they were never delivered.

The GitHub connector recovered the product change directly onto authoritative main as:
- `ccccb891b737995f6a24a58e844e5bd1d794b609` — `PERSIST-013: deterministic snapshot file format`
- product paths: `mma2composer/persistence_snapshot_format.go`, `mma2composer/persistence_snapshot_format_test.go`

Delivered behavior:
- bit snapshots pack values LSB-first, with deterministic ceil(count/8) size;
- register snapshots encode uint16 words big-endian, deterministic 2*count size;
- decode/validation rejects wrong raw lengths;
- raw payload contains no compatibility metadata; that belongs to PERSIST-014.

OpenHands' reported local self-check before transport failure:
- `cd mma2composer && go test -run TestPersistenceSnapshot .` → ok, exit 0
- full mma2composer/simulator/replicator regression → all ok
- gofmt clean; go vet exit 0

Because connector recovery recreated equivalent remote source rather than pushing the exact local commit object, the authoritative delivered SHA is `ccccb891...`. Independent verification remains deferred to PERSIST-022.

## CWAL

PERSIST-013 is complete and no longer ACTIVE. If a reused OpenHands workspace still contains local-only `9343d9e` / `ab27b27`, treat them as superseded duplicate history, not new work. PERSIST-014 is the next current task.
