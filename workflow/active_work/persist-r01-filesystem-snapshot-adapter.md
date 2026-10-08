# PERSIST-R01 — Filesystem Snapshot Adapter

Status: ACTIVE — HUMAN ASSIGNED 2026-10-08
Stage: CODE
Owner: OpenHands JR DEV
Previous: PERSIST-021
Next: PERSIST-R02

## Primary outcome
Implement the concrete appliance filesystem adapter for persistence snapshot payloads and manifest metadata using the existing persistence snapshot/store/source contracts.

## Scope
- Implement concrete `PersistenceSnapshotStore` / `PersistenceSnapshotSource` behavior needed by the runtime.
- Root persistence data only beneath the authoritative appliance data root using deterministic Port → Unit ID → Area paths.
- Reuse the existing raw snapshot format, manifest metadata, size validation and checksum contracts; do not invent a second format.

## Non-scope
No RBE runtime subscription, startup restore orchestration, runtime status wiring, Electron UI, repair/retry loop, production/customer data, or ICC edits.

## Acceptance
1. Snapshot payload + manifest can be written/read deterministically for one `(Port, UnitID, Area)` without escaping the appliance data root.
2. Existing manifest/length/checksum validation remains authoritative; malformed/missing files surface explicit errors rather than fabricated defaults.
3. Focused tests cover deterministic paths, round-trip, missing/corrupt content and safe replacement behavior without changing existing snapshot semantics.

## Evidence / handoff
Record exact changed paths, focused tests, bounded regression and delivered source SHA. This is JR DEV self-check only. After genuine delivery, automatically arm PERSIST-R02 for the next invocation and STOP.

## Dependencies
PERSIST-001..021 CODE complete. Human has promoted the full PERSIST-R01..R05 + PERSIST-UI01 chain.

## Sizing
2/0/1/1/0=4.

## CWAL
This is the sole current ACTIVE product task. Execute exactly this task, deliver it, update workflow evidence, automatically activate PERSIST-R02 for the next invocation, and STOP.
