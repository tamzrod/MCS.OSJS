# PERSIST-014 — Snapshot Manifest / Compatibility Metadata

Status: CODE COMPLETE — 2026-10-08; delivered via GitHub connector recovery; awaiting independent TEST/VERIFY (PERSIST-022)
Stage: CODE
Owner: OpenHands JR DEV
Previous: PERSIST-013
Next: PERSIST-015

## Primary outcome
Add snapshot metadata sufficient to reject incompatible or corrupted snapshots before restore.

## Scope
Implement or verify only the primary outcome above against the current authoritative MCS.OSJS source and existing MMA2 contracts. Preserve the Port → Unit ID → Memory authority model and existing State Sealing/RBE semantics unless this packet explicitly says otherwise.

## Non-scope
Do not restore or unseal memory.

## Acceptance
1. Manifest binds snapshot to Port, Unit ID, area, start, count and format version.
2. Integrity metadata detects corrupted/incomplete area payloads.
3. Changed incompatible memory layout causes explicit rejection rather than silent truncation/remap.

## Evidence / handoff
Coding owner records the exact source diff, changed paths, targeted self-checks and resulting source checkpoint. Do not claim independent TEST/VERIFY PASS. STOP after this task.

## Dependencies
Requires genuine completion evidence for PERSIST-013. Being present in `workflow/active_work/` authorizes this written scope but does not bypass predecessor or current-task selection gates.

## Sizing
2/0/2/1/2=7; metadata/validation only.

## Coding evidence / recovery

OpenHands completed and self-checked this task locally, but all push mechanisms failed because the injected `GITHUB_TOKEN` returned HTTP 401 and no valid write credential existed in that environment. Local-only commits were `5f944bd7e0c25167adce0424775230d8a420e901` (product) and `17648ce375ed3e6b46774b99b4414c0a2b78caae` (workflow); they were never delivered.

The GitHub connector recovered the product change directly onto authoritative main as:
- `98a952a89b087f12a9b99b6bd648ce97ac44e6f2` — `PERSIST-014: snapshot manifest and compatibility metadata`
- product paths: `mma2composer/persistence_snapshot_manifest.go`, `mma2composer/persistence_snapshot_manifest_test.go`

Delivered behavior:
- manifest binds format version, Port, Unit ID, area, start and count;
- manifest records deterministic payload byte length and SHA-256;
- validation rejects incompatible version/identity/area/start/count;
- incomplete payloads, configured-size mismatches and checksum corruption are rejected explicitly;
- no restore or unseal behavior is introduced.

The authoritative delivered SHA is `98a952a...`. Independent verification remains deferred to PERSIST-022.

## CWAL

PERSIST-014 is complete and no longer ACTIVE. If a reused OpenHands workspace still contains local-only `5f944bd` / `17648ce`, treat them as superseded duplicate history, not new work. PERSIST-015 is the next current task.
