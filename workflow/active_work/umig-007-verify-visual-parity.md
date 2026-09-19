# UMIG-007 — VERIFY: Three-Tab Visual Parity

Status: QUEUED — predecessor UMIG-006-V COMPLETE/PASS; BLOCKED from activation until approved three-tab donor screenshot baseline is retrievable and an exact JR packet can be authored.
Stage / owner: VERIFY / OpenHands (JR), packet and advancement / ChatGPT
Previous: UMIG-006-V (archived COMPLETE/PASS at `workflow/archive/umig-006-v-diagnostics-ui.md`)
Next: UMIG-007A

## Primary outcome
Independently verify single-window OS.js Toolkit appearance against the human-approved Electron donor.

## Verified reference and missing prerequisite
Frozen donor source is pinned at `1c971b9a6e00bafadf329df8821421a40cfc079c` in `workflow/archive/umig-001-freeze-electron-ui-donor.md`, including renderer HTML/CSS/JS. This is source provenance, NOT approved screenshots. `workflow/archive/umig-003-v-renderer-scope.md` says its screenshots were stored outside the repository at `/tmp/jr-shots/`; no durable approved comparable Memory, Replicator, Diagnostics donor screenshot set is identified. Request the three approved donor images or an explicit human-approved reproducible donor capture/baseline procedure with fixed dimensions and test state, then pin its accessible path/hash. Do not invent images, conflate the frozen source SHA with screenshots, or claim parity from a source diff. The actual OS.js tabs now load canonical data, not donor fixtures: define comparable synthetic device state and explicitly record any unavoidable data differences before issuing the packet.

## Verification action after prerequisite is met
Using a separately authorized, safe isolated OS.js test desktop, compare all three tabs at matched dimensions against the approved donor baseline: forms, spacing, selection, scrolling, device content, OS.js desktop/taskbar/window chrome, and no global CSS leakage. Keep actual appearance comparison separate from runtime success. Require donor and current screenshots/observations with dimensions and discrepancy list. Unavailable required visuals = BLOCKED; concrete mismatch = FAIL pending a separately authorized CODE fix. JR does not debug/fix product or improvise screenshots.

## Non-scope and authority
No code edits, backend protocol changes, destructive tests, production/user-data access, legacy UI removal, cutover or `operation cwal.md` edits. This QUEUED task and this descriptive file are NOT an executable JR instruction. ChatGPT must first verify reference availability, promote this as the sole ACTIVE task and issue the task-specific `handoff.md` JR TEST TASK; until then no `OPERATION CWAL` invocation.

## Dependencies
UMIG-006-V PASS (met); approved retrievable donor screenshots or approved reproducible capture procedure (missing); exact JR packet (missing).

## Sizing
Surface 0, environment 1, behavior 0, verification 1, recovery 1 = 3.
