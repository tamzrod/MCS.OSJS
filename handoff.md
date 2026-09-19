# Handoff

## Direction and authority — 2026-09-19

ChatGPT owns CODE, task-specific JR packet authoring, independent evidence review, and advancement. OpenHands/JR executes only the exact currently active `handoff.md` packet under the general `operation cwal.md`, writes only the authorized JR report section, commits/pushes only `handoff.md`, then STOPS. Only BLACK SHEEP WALL edits ICC. Do not alter `operation cwal.md` to compensate for packet authoring gaps. Never touch production Docker/`osjs-data`, operator/customer configurations, legacy apps, Windows Electron or retained test volumes; no `down -v`.

## UMIG-006-V adjudication — COMPLETE / PASS

Independent JR live report is preserved in full at immutable `handoff.md` commit `f6b7549b2c31e4fcf22427d8dee00b181f4b7186`. GitHub compare from previous main `f4994d435e1b3f02d05f2355abcd2e5c09b14901` shows JR modified ONLY `handoff.md`. Reviewer read the current packet/report and accepted isolation, five-service startup, real browser Diagnostics showing canonical device RUNNING/IDLE and RUNNING/OK, UNKNOWN global services, disabled controls, read-only refresh without hash/restart changes, real synthetic connection-refused ERROR and recovery, six bounded explicit applies, baseline restoration and project-owned teardown with no volume removal. Critically, actual Toolkit window chrome CLOSE -> zero windows -> Start-menu RELAUNCH -> fresh canonical Replicator load and unchanged files/owners/restart count were directly reported; this closes the prior UMIG-005-V window-lifecycle evidence gap. This is review of JR evidence, not a rerun and not production, global service-health, Windows or visual-parity acceptance. Archive: `workflow/archive/umig-006-v-diagnostics-ui.md`. Retained test volumes `mcsverify-1789784184-232_verify-data` (if present), `mcsverify-rep-005-20260919_verify-data` and `mcsverify-diag-006-20260919_verify-data` MUST NOT be reused or removed without distinct authorization.

## Workflow stop gate — NO ACTIVE TASK / NO JR PACKET

UMIG-006-V is archived COMPLETE/PASS. UMIG-007 (visual parity) is QUEUED and cannot yet be promoted because its required *approved retrievable Electron donor screenshots* are not identified. Donor **source**, not screenshots, is frozen at `1c971b9a6e00bafadf329df8821421a40cfc079c` in `workflow/archive/umig-001-freeze-electron-ui-donor.md`. The earlier UMIG-003-V archive says browser screenshots lived outside the repo at `/tmp/jr-shots/`, not a durable donor baseline. Obtain approved Memory, Replicator and Diagnostics donor screenshots at defined window dimensions and comparable device state, OR explicit human approval of an independently reproducible donor capture procedure and pin its durable artifacts. Do not manufacture a baseline or silently treat frozen HTML/CSS as visual evidence. The current task file is `workflow/active_work/umig-007-verify-visual-parity.md` and remains QUEUED; UMIG-007A remains QUEUED; UMIG-008/009 cutover/removal in planning. Zero ACTIVE intentionally means STOP, not permission for JR to select a task.

No `## JR TEST TASK — CURRENT` exists at this HEAD. Do NOT invoke `OPERATION CWAL` until ChatGPT checks the missing donor baseline, promotes UMIG-007 as sole ACTIVE and authors the exact bounded packet. No further Docker/browser runs, product/Compose edits, ICC or general CWAL changes are authorized by this handoff.

## Next action and recommendation

Next action: human supply or authorize a pinned three-tab donor screenshot baseline; ChatGPT then promote UMIG-007 and publish its exact JR packet. Recommendation: do not rerun the verified Diagnostics stack or ask JR to guess the donor reference; visual parity and any future cutover remain separate evidence/approval gates.
