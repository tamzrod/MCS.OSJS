# UMIG-006 — CODE: Adapt OS.js Diagnostics

Status: ACTIVE — promoted 2026-09-19 after UMIG-005-V core live Replicator behavioral PASS review with explicit window-reopen evidence gap carried to UMIG-006-V.
Stage / owner: CODE / ChatGPT
Previous: UMIG-005-V (archived; exact Toolkit close/reopen NOT yet verified)
Next: UMIG-006-T (QUEUED)

## Primary outcome
Author Diagnostics mapping that shows only truthful OS.js-available observations.

## Scope
Preserve donor layout; reuse supported Toolkit Memory/Replicator/MMA2 response/status/error contracts without inventing backend capabilities; disable Windows-only native service controls and filesystem paths when unsupported; show UNKNOWN/UNAVAILABLE rather than green by default. Keep Memory and Replicator source behavior unchanged; no production data or Docker edits.

## Non-scope
No live testing by CODE, privileged new endpoint, backend service controls, Windows IPC, desktop redesign, legacy cutover, or ICC edits. Do not claim UMIG-005-V Toolkit window close/reopen was tested: JR instead switched tabs. The missing check is owned by future UMIG-006-V rendered verification.

## Coding acceptance / handoff
1. Supported Diagnostics observations map actual status/error contracts; absent measurements are UNKNOWN/UNAVAILABLE.
2. Unsupported native controls cannot call Windows APIs; no extra endpoint or privileged operation.
3. Record exact source files, code commit and unsupported capabilities; author narrowly scoped unit tests and separate JR TEST packet without running independent tests or prematurely promoting the successor.

## Dependencies
Observed UMIG-005-V core behavior accepted with limitation at archived task/report `1a04664e5fdc21e3bd323a97a4f12f3f9d39abdd`; any new service-control capability needs separate approval. Sole ACTIVE task must remain UMIG-006 until CODE checkpoint and workflow promotion.

## Sizing
Surface 1, environment 0, behavior 1, verification 0, recovery 0 = 2.
