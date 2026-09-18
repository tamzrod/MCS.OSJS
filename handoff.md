# Handoff

## Current direction — 2026-09-18

Human priority: staged OS.js replacement with one `MCS Modbus Toolkit`. Preserve the OS.js desktop, Start menu, taskbar and clock; the final default desktop has exactly one Toolkit icon. Do not remove legacy UI packages before verified cutover. MMA2, Go services, shared memory, user configuration and Windows Electron remain in scope and are not decommissioned by this UI migration. Windows LED/installer work is a separate parked track.

## Roles and authoritative state

ChatGPT owns CODE, source checkpoints, repair and workflow advancement. OpenHands is JR for separate TEST and VERIFY stages via `operation cwal.md`: JR never codes, fixes, advances tasks or updates ICC, and may edit only an explicitly authorized test-report section in this file. Only BLACK SHEEP WALL updates ICC.

ACTIVE: `UMIG-002-T` — independent OS.js Toolkit build and discovery TEST, **retest pending, no PASS**. QUEUED: `UMIG-002-V` — separate rendered one-window VERIFY, not authorized by the current packet. The original UMIG-002 CODE task and bounded repair `UMIG-002-R` are archived as source-only checkpoints. All other parked work is unchanged; no donor UI copy, launcher switch, legacy app removal or Windows Electron change is authorized by this packet.

## Failed test and coding repair

Original JR test verdict: **FAIL**, commit `752a54108ecaf91d9e53440ed344f8f31a1ce8de`. Both build and discovery commands exited 0 and Toolkit `dist/main.js`/`main.css` existed, but Toolkit was missing from discovered package list, `packages.json`, `dist/metadata.json` and `dist/apps/`. JR identified the absent Toolkit `package.json` required by OS.js discovery. The full original raw evidence remains preserved in the immutable report at https://github.com/tamzrod/MCS.OSJS/blob/752a54108ecaf91d9e53440ed344f8f31a1ce8de/handoff.md ; this FAIL is not erased or reclassified.

Human authorized repair on 2026-09-18. `UMIG-002-R` added only `OSJS/src/packages/MCSModbusToolkit/package.json` with `name: mcs-modbus-toolkit`, `version: 0.1.0`, `private: true`, and `osjs.type: package`, matching sibling OS.js package discovery convention. Source commit: `cd67e150139b60f3914db8c6f1998c96c5b8da07`. Coding agent read the committed file back and compared the source delta; this is **source-only evidence**, NOT an executed build/discovery or rendered UI PASS. Repair record: `workflow/archive/umig-002-r-toolkit-discovery-manifest.md`.

## ICC prerequisite

`ICC/INDEX.md` records baseline `7029e41`; subsequent JR report, task-state changes, and Toolkit package manifest affect handoff/active-work and OS.js shell context. Request a **bounded BLACK SHEEP WALL refresh of the affected branches** and record its evidence before invoking OpenHands. JR must not edit ICC or treat the prior baseline as synchronized. If this required refresh is still pending, return BLOCKED rather than testing against a knowingly stale prerequisite. Do not refresh unrelated context branches.

## JR TEST TASK — CURRENT: UMIG-002-T (RETEST)

GOAL / TARGET: Rerun only the existing Toolkit build-and-discovery TEST against committed repair `cd67e150139b60f3914db8c6f1998c96c5b8da07`. Task authority: `workflow/active_work/umig-002-t-build-discover.md`. This is a retest after prior FAIL, not UMIG-002-V or a new feature test.

REPOSITORY STATE: In a disposable checkout of `tamzrod/MCS.OSJS` at current `origin/main`, record HEAD and `git status --short` before starting. Verify required repair commit `cd67e150139b60f3914db8c6f1998c96c5b8da07` is an ancestor of HEAD and `OSJS/src/packages/MCSModbusToolkit/package.json` is present. If missing, unexpectedly dirty, or the ICC prerequisite above has not been satisfied, report BLOCKED without reset/clean/restore of product files.

EXACT COMMAND / ACTION: From the repository root, execute the two original product commands in order, capturing independent exit codes and raw output for each: (1) `cd OSJS && npm run build:local-packages`; (2) `cd OSJS && npm run package:discover`. Inspect Toolkit `dist/main.js`, `dist/main.css`; verify Toolkit appears in discovery command output, `OSJS/packages.json`, `OSJS/dist/metadata.json`, and `OSJS/dist/apps/MCSModbusToolkit/`. Record pre/post `git status --short`. If command 1 fails, report FAIL with its output; do not pretend command 2 ran. Do not launch GUI, invoke Electron tools or run other product tests.

OPTIONAL SAFE SETUP: Check Node/npm toolchain; `OSJS/package.json` declares Node >=10 <17, so prefer compatible sandbox-local Node 16 when safely available. JR may install missing dependencies only inside its disposable sandbox using the existing package manager without edits to tracked source, manifests, locks or configuration, per `operation cwal.md`. Note actual tool versions. If safe setup is unavailable, BLOCKED. Generated ignored build outputs can be inspected; unexpected tracked modifications must be reported without cleanup.

EXPECTED RESULT: Both commands exit zero, Toolkit `main.js` and `main.css` exist, discovery output lists `mcs-modbus-toolkit as MCSModbusToolkit`, `packages.json` references `src/packages/MCSModbusToolkit`, `dist/metadata.json` contains the application, and `dist/apps/MCSModbusToolkit/` exists. A zero exit code alone is never sufficient for PASS. No GUI, backend or Windows verification is implied.

EVIDENCE TO RETURN: HEAD SHA; pre/post status; ICC-prerequisite status; toolchain/setup; each command's raw output and exit code; actual artifact existence and paths; exact relevant discovered-list/manifest entries and directory listing; errors or unexpected changes. PASS only when **all** mandatory results are observed, FAIL for an executed contradictory product result, BLOCKED for a missing prerequisite or unobservable requirement. Preserve the initial FAIL separately as history.

REPORT-WRITE AUTHORITY: OpenHands/JR may replace only the `## JR TEST REPORT — UMIG-002-T RETEST` section immediately below; preserve every other section, including initial FAIL reference and this packet. JR may commit/push **only `handoff.md`** for its report, then STOP. Do not edit product source, workflow or ICC or activate UMIG-002-V. ChatGPT reviews the new evidence and alone decides subsequent task advancement.

## JR TEST REPORT — UMIG-002-T RETEST

PENDING — the repaired source has not yet been retested. JR replaces only this section with observed PASS / FAIL / BLOCKED and raw evidence after the ICC prerequisite is satisfied and the current packet executes.

## Evidence limits

No independent retest PASS, rendered one-window verification, Windows COMMS/install verification, or final OS.js UI migration acceptance has been established by the source-only manifest repair.
