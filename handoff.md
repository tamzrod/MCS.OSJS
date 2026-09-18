# Handoff

## Current direction — 2026-09-18

Human priority: staged OS.js replacement with one `MCS Modbus Toolkit`. Preserve the OS.js desktop, Start menu, taskbar and clock; the final default desktop has exactly one Toolkit icon. Do not remove legacy UI packages before verified cutover. MMA2, Go services, shared memory, user configuration and Windows Electron remain in scope and are not decommissioned by this UI migration. Windows LED/installer work is a separate parked track.

## Roles and authoritative state

ChatGPT owns CODE, source checkpoints, repair and workflow advancement. OpenHands is JR for separate TEST and VERIFY stages via `operation cwal.md`: JR never codes, fixes, advances tasks or updates ICC, and may edit only an explicitly authorized test-report section in this file. Only BLACK SHEEP WALL updates ICC.

COMPLETE: `UMIG-002-T` — independent Toolkit build/discovery retest PASS at `e9d25e3d4c1b832126a161a6071fd4abf8b5546b`, reviewed and archived in `workflow/archive/umig-002-t-build-discover.md`. Its initial FAIL at `752a541` remains historical, not reclassified; the one-file manifest CODE repair `UMIG-002-R` is archived. ACTIVE: `UMIG-002-V` — rendered one-window VERIFY, **not yet run**. No other ACTIVE task. UMIG-001 and UMIG-003 onward remain in Planning; parked Windows/Memory/Replicator work remains unchanged. No donor UI copy, launcher cutover, legacy app deletion, backend runtime acceptance, or Windows Electron change is authorized by the current packet.

## UMIG-002 test history and reviewed evidence

Original JR FAIL: `752a54108ecaf91d9e53440ed344f8f31a1ce8de`; both commands exited zero and Toolkit assets existed but it was missing from package discovery and generated manifests. Source-only `UMIG-002-R` repair added the missing `OSJS/src/packages/MCSModbusToolkit/package.json` (`osjs.type: package`) at `cd67e150139b60f3914db8c6f1998c96c5b8da07`. BLACK SHEEP WALL refreshed affected ICC to baseline `9775593` and pushed `24c00de78ccbdf1773c1c9749f2062ff578c1b4f`. JR's subsequent retest at that HEAD used Node 16.20.2 and npm 8.19.4 in a disposable checkout: `npm run build:local-packages` exited 0 with `main.js` 1909 bytes and `main.css` 433 bytes; `npm run package:discover` exited 0 with seven packages and `mcs-modbus-toolkit as MCSModbusToolkit`. `OSJS/packages.json`, `OSJS/dist/metadata.json`, and `OSJS/dist/apps/MCSModbusToolkit/` all contained the expected entry/assets; pre/post tracked checkout clean. Full independently authored JR report and raw output are preserved at https://github.com/tamzrod/MCS.OSJS/blob/e9d25e3d4c1b832126a161a6071fd4abf8b5546b/handoff.md ; initial full FAIL at https://github.com/tamzrod/MCS.OSJS/blob/752a54108ecaf91d9e53440ed344f8f31a1ce8de/handoff.md . ChatGPT reviewed every required TEST result and archived UMIG-002-T. **Only build and discovery passed; no rendered UI or runtime result is claimed.**

## ICC prerequisite

`ICC/INDEX.md` was refreshed to baseline `9775593` by BLACK SHEEP WALL at `24c00de`. Since then the JR PASS report and the coding-agent archive/promotion/handoff changed the active-work context. Before JR runs this new VERIFY packet, request a **bounded BLACK SHEEP WALL refresh of only the affected ICC branch and its necessary parents**; do not treat the old `UMIG-002-T` execution summary as current or refresh unrelated branches. Only BLACK SHEEP WALL edits ICC. If this prerequisite remains unsatisfied, JR reports BLOCKED rather than bypassing it.

## JR TEST TASK — CURRENT: UMIG-002-V

GOAL / TARGET: Independent *rendered* verification that the committed/discovered `MCSModbusToolkit` launches once as exactly one OS.js window showing `NOT CONNECTED — PLACEHOLDER ONLY`, while desktop, Start menu, taskbar and legacy Simulator/Replicator availability remain intact. Task authority: `workflow/active_work/umig-002-v-window-launch.md`. This is VERIFY, not a rerun of UMIG-002-T or a functional backend test.

REPOSITORY STATE: Create a disposable checkout of current `origin/main` of `tamzrod/MCS.OSJS`; record HEAD and pre-test `git status --short`. Verify `e9d25e3d4c1b832126a161a6071fd4abf8b5546b` (independent TEST PASS) is an ancestor of HEAD, and that `UMIG-002-V` is the sole ACTIVE task. Confirm the bounded ICC refresh requested above is recorded. Missing predecessor, stale ICC, unexpectedly dirty tracked checkout, unavailable safe disposable environment or GUI = BLOCKED; do not reset, clean, restore or modify product files to force execution.

OPTIONAL SAFE SETUP: Use sandbox-local compatible Node 16 / npm (repository engines Node >=10 <17). From a disposable checkout, `cd OSJS && npm install --no-audit --no-fund` is allowed to prepare dependencies if needed, provided tracked files remain unchanged. Use a fresh disposable OSJS_DATA_DIR (for example create it with `mktemp -d` outside the checkout) for VFS/session data; never use installed appliance or user data. Check whether localhost port 18209 is free before starting the server; if unavailable, report BLOCKED and request a packet change rather than selecting a different port. Prepare a browser/GUI test session if available without changing product source or user state.

EXACT COMMAND / ACTION: In that disposable checkout run, in this order, recording actual commands/output/exit codes: (1) `cd OSJS && npm run build:local-packages`; (2) `cd OSJS && npm run package:discover`; (3) `cd OSJS && npm run build`; (4) from `OSJS/`, start `OSJS_DATA_DIR="$TEST_DATA_DIR" PORT=18209 npm run serve` in a dedicated session, where `TEST_DATA_DIR` is the fresh directory created in safe setup, and retain startup logs. If a required build/start step fails, report FAIL with evidence and do not fabricate later observations. In a real browser/GUI navigate to `http://127.0.0.1:18209/`, observe the desktop, open the existing Start menu, locate and click `MCS Modbus Toolkit` **exactly once**, then inspect the rendered result: exactly one Toolkit window; literal disconnected placeholder; desktop, Start menu and taskbar remain functional; legacy Simulator and Replicator applications remain listed/available. Do not launch them or start any backend. Record how the window count and legacy availability were observed. If the browser/GUI or Start menu cannot be operated, or an alternate launch method seems necessary, report BLOCKED and request an exact packet update; no source/metadata/HTTP-only substitutes for rendered observation. Stop only the server process started for this test and preserve evidence; remove solely disposable test data if safe and unambiguous. Record post-test `git status --short` and any errors.

EXPECTED RESULT: Repo-native preparation/build/discovery/serve actions complete successfully; `http://127.0.0.1:18209/` renders the OS.js desktop; exactly one Toolkit window appears after one Start-menu launch, displaying `NOT CONNECTED — PLACEHOLDER ONLY`; Start menu, taskbar and desktop still function; legacy Simulator and Replicator remain available. No backend, Electron, Windows or final migration acceptance implied. A build/discovery exit 0 alone is never a VERIFY PASS.

EVIDENCE TO RETURN: HEAD SHA, ICC readiness, exact toolchain/setup and commands with exit codes, pre/post tracked status, server URL/startup output, direct rendered screenshots if available or detailed first-hand GUI observations and browser method, launch interaction, Toolkit window count and placeholder text, desktop/Start/taskbar and legacy application observations, any errors, and safe cleanup. PASS only if **all** required UI observations are directly confirmed; FAIL if an executed product result contradicts expectations; BLOCKED if an indispensable UI or safety prerequisite is unavailable. No manual claims inferred from source or prior TEST.

REPORT-WRITE AUTHORITY: OpenHands/JR may replace **only** the `## JR TEST REPORT — UMIG-002-V` section immediately below, preserving every other section. JR may commit/push only `handoff.md` for that report and must STOP. No product fixes, workflow advancement, ICC edits, icon cutover, legacy app deletion, backend launches or next-task selection. ChatGPT reviews the evidence and decides completion independently.

## JR TEST REPORT — UMIG-002-V

VERDICT: **BLOCKED** — the packet's mandatory ICC prerequisite is not satisfied. No build, discovery, serve or rendered UI step was executed.

### Block reason

The packet (`REPOSITORY STATE`) makes a recorded **bounded BLACK SHEEP WALL ICC refresh** a gate, and `workflow/active_work/umig-002-v-window-launch.md` independently requires that affected ICC context be refreshed before JR execution. That refresh has not happened.

- The last committed ICC change is `24c00de78ccbdf1773c1c9749f2062ff578c1b4f` ("BLACK SHEEP WALL: refresh ICC to baseline 9775593"). `git log --oneline 9775593..HEAD -- ICC/` returns exactly that one commit.
- The five later commits (`e9d25e3`, `e90f2fe`, `f54d277`, `f3fe611`, `9a9daf0`) changed `handoff.md` and `workflow/` but did not touch `ICC/`.
- Consequently committed ICC still describes the **superseded** state: `ICC/INDEX.md` reads "Overlay state: exactly one ACTIVE task, `UMIG-002-T` (Toolkit build and discovery TEST), **retest pending with no PASS**" and "`UMIG-002-V` is QUEUED", and `ICC/context/active-work.md` says "`handoff.md` agrees with this directory: `UMIG-002-T` is ACTIVE and `UMIG-002-V` is QUEUED" and, under Verification state, "UMIG-002-T: the retest has produced no evidence yet. `handoff.md` carries a `PENDING` `JR TEST REPORT — UMIG-002-T RETEST` placeholder, not a result."

That committed ICC statement is now factually wrong, and it is stale for the exact delta it declares a dependency on. `ICC/INDEX.md` registers `handoff.md` and `workflow/active_work/*.md` as source dependencies of both `L0-project` and `active-work`, and `ICC/context/active-work.md` repeats that dependency list. The delta since the last ICC write is precisely those dependency files (`git diff --name-status 9775593..HEAD` shows `M handoff.md`, `M workflow/active_work/umig-002-v-window-launch.md`, `D workflow/active_work/umig-002-t-build-discover.md`, `A workflow/archive/umig-002-t-build-discover.md`). The dependency set therefore intersects the delta, so ICC is stale rather than merely irrelevant, and the packet forbids treating the old `UMIG-002-T` execution summary as current.

Honoring the recorded gate: jr may not refresh ICC (`operation cwal.md` §2 and §4; only BLACK SHEEP WALL writes ICC), may not proceed on unverified prerequisite evidence, and may not self-authorize a cross-directive BLACK SHEEP WALL run under the Directive Composition Rule. Verdict is BLOCKED, not PASS and not FAIL: the requested product test was never executed, so no product result contradicts any expectation.

### Repository-state checks that did pass

- HEAD: `9a9daf09e4190004e4b27cd6d63417170b52233a`; `origin/main` is at the same SHA.
- Pre-test `git status --short`: empty (tracked checkout clean).
- `UMIG-002-V` predecessor: `git merge-base --is-ancestor e9d25e3d4c1b832126a161a6071fd4abf8b5546b HEAD` exits 0, so the independent TEST PASS commit is an ancestor of HEAD.
- Exactly one ACTIVE task in `workflow/active_work/`: `grep -rln "^Status: ACTIVE"` returns only `workflow/active_work/umig-002-v-window-launch.md`; 17 files are QUEUED.
- Port 18209 was free before execution.

None of these compensate for the unsatisfied ICC gate.

### Commands executed (evidence only; no product test run)

```text
git rev-parse HEAD                              -> 9a9daf09e4190004e4b27cd6d63417170b52233a
git status --short                              -> (empty)
git merge-base --is-ancestor e9d25e3 HEAD       -> exit 0
git log --oneline 9775593..HEAD                 -> 9a9daf0, f3fe611, f54d277, e90f2fe, 5a7155e, e9d25e3, 24c00de
git log --oneline 9775593..HEAD -- ICC/         -> 24c00de only
git diff --name-status 9775593..HEAD            -> ICC/INDEX.md, ICC/context/{L0-project,active-work,osjs-shell}.md,
                                                   electron/renderer/app.js, electron/test/memory-settings.test.js,
                                                   handoff.md, workflow/active_work/umig-002-t-build-discover.md (D),
                                                   workflow/active_work/umig-002-v-window-launch.md,
                                                   workflow/archive/umig-002-t-build-discover.md (A)
grep -rln "^Status: ACTIVE" workflow/active_work/ -> umig-002-v-window-launch.md
git config user.name / user.email               -> unset in this disposable checkout
```

`git ls-remote origin main` also returns `9a9daf09e4190004e4b27cd6d63417170b52233a`, confirming no newer remote HEAD carried an ICC refresh.

### Environment notes (not the block reason)

- Sandbox-local toolchain present: Node `v22.23.2`, npm `10.9.8`. The repository declares `engines` Node `>=10 <17`, so the optional safe-setup step expected a Node 16 install; a webpack 4 / MD4 build may additionally require `NODE_OPTIONS=--openssl-legacy-provider` on Node 17+, as the `osjs-shell` ICC node records. This was not tested, because the gate above failed first.
- No repository file was modified. No `npm install`, build, discovery, serve, browser or GUI action was performed. No server was started and no test data directory was created, so no cleanup was required.
- Post-test `git status --short`: empty (unchanged from pre-test). The only working-tree change this run makes is this authorized report section.

### Not verified / not claimed

No rendered Toolkit window, placeholder text, Start-menu launch, window count, desktop/taskbar integrity, legacy Simulator/Replicator availability, backend health, Electron, Windows or final migration acceptance is claimed. This report records a blocked prerequisite only.

### Unblock path (for the coding agent; JR cannot do this)

1. Run the bounded BLACK SHEEP WALL refresh for the affected ICC branch and its necessary parents (`active-work`, `L0-project`, `osjs-shell`, plus this index), so committed ICC records `UMIG-002-T` COMPLETE/PASS at `e9d25e3`, `UMIG-002-V` as the sole ACTIVE task, and the current `PENDING` report state. Commit/push it.
2. Confirm port 18209 remains free and the disposable-checkout plan is unchanged.
3. Re-issue or confirm the `UMIG-002-V` packet, then re-invoke JR. The packet text itself needs no correction; only its prerequisite must be met.

### Report commit

Report-write authority was used for this section only. `handoff.md` is committed and pushed with this report; no other file is included.

## Evidence limits

UMIG-002-T independent build/discovery PASS is established. UMIG-002-V rendered UI acceptance, Windows COMMS/installer acceptance and final migration acceptance remain unverified.
