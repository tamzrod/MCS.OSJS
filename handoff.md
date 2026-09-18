# Handoff

## Current direction — 2026-09-18

Human priority: staged OS.js replacement with one `MCS Modbus Toolkit`. Windows LED repairs, if any, are a separate track. Preserve the OS.js desktop, Start menu, taskbar and clock; final default desktop has exactly one Toolkit icon. Legacy UI packages are removed only after verified replacement and cutover. MMA2, Go services, shared memory, user configuration and Windows Electron are not decommissioned by this UI migration.

## Roles and authoritative state

ChatGPT owns CODE, source changes, source checkpoint and task advancement. OpenHands is JR for separate TEST and VERIFY stages, invoked using `operation cwal.md` and governed by its current `JR TEST TASK`; OpenHands does not code, fix failures, promote/archive tasks or update ICC. Only BLACK SHEEP WALL updates ICC.

CODE `UMIG-002`: source-only milestone archived. Placeholder source added at `f83f1e716a3ab4c12446c04a7d264e9eb280bc06`; scoped disconnected display commits `1812e1c17a29dff5b56d7c4cb078da6b9a099ac9` and `9410edccd3f8914cc76a365980c0d0f1e9bd8bbb`. Read back `OSJS/src/packages/MCSModbusToolkit/{index.js,index.scss,metadata.json,icon.svg,webpack.config.js}`. Compared coding delta: only Toolkit `index.js` and `index.scss` modified after the earlier scaffold. This evidence proves source authorship only, NOT a successful build or UI launch.

ACTIVE: `UMIG-002-T` — OpenHands build and discovery TEST. QUEUED: `UMIG-002-V` — separate rendered one-window VERIFY, not authorized for execution by the current packet. Other parked Windows/Memory/Replicator tasks are unchanged; UMIG-001 donor approval and UMIG-003 onward stay in Planning. No donor SHA approved, no UI copy, launcher switch or old app deletion.

## ICC prerequisite

`ICC/INDEX.md` still records an earlier baseline and its Planning/Active Work views are stale. Request a bounded BLACK SHEEP WALL refresh for affected Planning and Active Work branches before invoking OpenHands; JR cannot perform that refresh. If this prerequisite has not been satisfied, report BLOCKED rather than treating stale ICC as current. Do not modify unrelated ICC branches or assume a clean Windows checkout.

## JR TEST TASK — CURRENT: UMIG-002-T

GOAL / TARGET: Verify only the independent OS.js `MCSModbusToolkit` package build and discovery from committed source. Task authority: `workflow/active_work/umig-002-t-build-discover.md`. Do not run UMIG-002-V in this invocation.

REPOSITORY STATE: Use a disposable checkout of `tamzrod/MCS.OSJS` at current `origin/main`, including source commit `9410edccd3f8914cc76a365980c0d0f1e9bd8bbb`. Record `git rev-parse HEAD` and `git status --short` before starting. If the required baseline is absent, checkout has unexpected tracked changes, or ICC prerequisite above remains unresolved, report BLOCKED; do not reset/clean/restore product files to force execution.

EXACT COMMAND / ACTION: From repository root, run `cd OSJS && npm run build:local-packages && npm run package:discover`. Record each command's actual exit code and raw output separately. Then inspect `OSJS/src/packages/MCSModbusToolkit/dist/main.js`, `dist/main.css` and OS.js package-discovery output/manifest to find `MCSModbusToolkit`. Record `git status --short` after. Do not launch GUI or invoke Electron tools.

OPTIONAL SAFE SETUP: Check the local Node/npm toolchain and install missing dependencies only within the disposable sandbox through the existing OS.js package manager/lockfile as permitted by `operation cwal.md`; do not edit tracked manifests, locks or source. If no safe local setup exists, report BLOCKED. Generated ignored build output may be inspected, but unexpected tracked modifications must be reported and not cleaned.

EXPECTED RESULT: The two exact OS.js commands exit zero, `main.js` and `main.css` exist in Toolkit `dist/`, and OS.js discovery lists the Toolkit package. No Electron build/install, backend service startup, saved configuration mutation or desktop change is required. No browser or backend PASS is implied.

EVIDENCE TO RETURN: HEAD SHA, pre/post `git status --short`, toolchain/setup notes if needed, verbatim output and exit code for each command, artifact paths and observed existence, discovered Toolkit metadata/path, and unexpected changes/errors. Overall verdict must be PASS only if all required results are observed; FAIL for an executed product test contradicting expectations; BLOCKED for unobservable prerequisites. Never substitute static/source inspection for build/discovery evidence.

REPORT-WRITE AUTHORITY: OpenHands/JR may append or replace only the `## JR TEST REPORT — UMIG-002-T` section below in `handoff.md`. Preserve every other section. It may commit/push only `handoff.md` if the current packet permits this; this packet permits a handoff-only report commit/push. Do not change task status, source or ICC. After reporting, STOP. ChatGPT reviews the evidence and, if PASS, separately activates `UMIG-002-V` with a new exact packet.

## JR TEST REPORT — UMIG-002-T

PENDING — no OpenHands test has been executed or observed by ChatGPT. OpenHands replaces this section with its raw evidence and PASS / FAIL / BLOCKED verdict when the packet runs.

## Evidence limits

Windows RLED appearance was accepted by the human, but installed live COMMS status and RREC-004 installer repair/hash/ProgramData verification are not established by this OS.js source milestone. Do not claim Windows tests or the final OS.js UI migration passed.
