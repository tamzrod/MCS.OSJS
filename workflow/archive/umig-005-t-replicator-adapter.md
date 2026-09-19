# UMIG-005-T — TEST: Replicator Adapter Contract

Status: COMPLETE — independent UNIT/BUILD PASS reviewed 2026-09-19 by ChatGPT.
Stage / owner: TEST / OpenHands JR; evidence review / ChatGPT
Previous: UMIG-005 (archived CODE source-only)
Next: UMIG-005-V (sole ACTIVE independent live VERIFY)

## Evidence reviewed
Immutable JR report: `handoff.md` at `4bbe4df852c5ab38c8081abdc15047279b8ceff4`; source/test baseline `f28d5f6b9d46d280ce6486f72de0afaec927ffe4`. GitHub compare showed `f28d5f6..4bbe4df` modified ONLY `handoff.md` and the diff replaced only the authorized JR report section. JR used a fresh clean disposable checkout, read-only unshallow, sandbox-local Node 16/npm 8 and isolated npm install; reports all nine Node tests and all three build/discovery commands exit 0, clean tracked tree, Toolkit bundle `main.js` 47134 bytes / `main.css` 121 bytes, valid discovery/metadata and bounded source diff. Contract, typed ownership/missing-device errors, relay per-socket routing, one apply, validation, per-block status and UNKNOWN COMMS, transport timeout/close and Memory regression were reported as passing. No actual Go runtime or product Docker was started. ChatGPT accepts this UNIT/BUILD stage ONLY, not GUI/backend/production acceptance.

## Boundary
Replicator live behavior, real browser, MMA2 reservation and source polling/recovery remain unverified, and four COMMS LEDs have no independent probes. The previous Memory test volume `mcsverify-1789784184-232_verify-data` is retained and must not be reused or removed without separate authorization. `UMIG-005-V` requires new project-owned disposable source and destination, direct runtime/browser evidence and scoped cleanup; no operator data, legacy/UI cutover, ICC or production Compose edits.

## Sizing
Surface 0, environment 0, behavior 0, verification 1, recovery 0 = 1.
