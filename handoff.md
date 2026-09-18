# Handoff

## Current direction — 2026-09-18

Human priority: staged OS.js replacement with one `MCS Modbus Toolkit`. Preserve the OS.js desktop, Start menu, taskbar and clock; the final default desktop has exactly one Toolkit icon. Do not remove legacy UI packages before verified cutover. MMA2, Go services, shared memory, user configuration and Windows Electron remain in scope and are not decommissioned by this UI migration. Windows LED/installer work is a separate parked track.

## Roles and authoritative state

ChatGPT owns CODE, source checkpoints, repair and workflow advancement. OpenHands is JR for separate TEST and VERIFY stages via `operation cwal.md`: JR never codes, fixes, advances tasks or updates ICC, and may edit only an explicitly authorized test-report section in this file. Only BLACK SHEEP WALL updates ICC.

COMPLETE: `UMIG-002-T` — independent Toolkit build/discovery retest PASS at `e9d25e3d4c1b832126a161a6071fd4abf8b5546b`, reviewed and archived in `workflow/archive/umig-002-t-build-discover.md`. Its initial FAIL at `752a541` remains historical, not reclassified; the one-file manifest CODE repair `UMIG-002-R` is archived. ACTIVE: `UMIG-002-V` — rendered one-window VERIFY, **not yet run**. No other ACTIVE task. UMIG-001 and UMIG-003 onward remain in Planning; parked Windows/Memory/Replicator work remains unchanged. No donor UI copy, launcher cutover, legacy app deletion, backend runtime acceptance, or Windows Electron change is authorized by the current packet.

## UMIG-002 test history and reviewed evidence

Original JR FAIL: `752a54108ecaf91d9e53440ed344f8f31a1ce8de`; both commands exited zero and Toolkit assets existed but it was missing from package discovery and generated manifests. Source-only `UMIG-002-R` repair added the missing `OSJS/src/packages/MCSModbusToolkit/package.json` (`osjs.type: package`) at `cd67e150139b60f3914db8c6f1998c96c5b8da07`. BLACK SHEEP WALL refreshed affected ICC to baseline `9775593` and pushed `24c00de78ccbdf1773c1c9749f2062ff578c1b4f`. JR's subsequent retest at that HEAD used Node 16.20.2 and npm 8.19.4 in a disposable checkout: `npm run build:local-packages` exited 0 with `main.js` 1909 bytes and `main.css` 433 bytes; `npm run package:discover` exited 0 with seven packages and `mcs-modbus-toolkit as MCSModbusToolkit`. `OSJS/packages.json`, `OSJS/dist/metadata.json`, and `OSJS/dist/apps/MCSModbusToolkit/` all contained the expected entry/assets; pre/post tracked checkout clean. Full JR PASS evidence is preserved at https://github.com/tamzrod/MCS.OSJS/blob/e9d25e3d4c1b832126a161a6071fd4abf8b5546b/handoff.md ; initial full FAIL at https://github.com/tamzrod/MCS.OSJS/blob/752a54108ecaf91d9e53440ed344f8f31a1ce8de/handoff.md . ChatGPT reviewed every required TEST result and archived UMIG-002-T. **Only build and discovery passed; no rendered UI or runtime result is claimed.**

Initial UMIG-002-V invocation was **BLOCKED before any product test** because the prior packet unnecessarily demanded a BLACK SHEEP WALL refresh. Its complete JR report and evidence remain preserved at https://github.com/tamzrod/MCS.OSJS/blob/3831e33338a77fcea2e70e96d3f5127193fc5c10/handoff.md . Human authorized removal of that unnecessary gate. This replacement packet supersedes only the former prerequisite; the prior BLOCKED is not reclassified, and the rendered VERIFY still has no result.

## ICC scope — not a test prerequisite

ICC is an optional semantic context cache, not the authority for task status or this test. Its older active-work summary is known stale; do **not** use it as current truth. For this self-contained VERIFY, use the current `handoff.md`, `workflow/active_work/umig-002-v-window-launch.md`, Git state and actual runtime observations. Do not require an ICC refresh, read stale ICC as authoritative, invoke BLACK SHEEP WALL, write ICC, or report BLOCKED merely because its baseline differs from HEAD. Separate BLACK SHEEP WALL maintenance may be requested when updated semantic context is genuinely needed; it is not part of OPERATION CWAL.

## JR TEST TASK — CURRENT: UMIG-002-V

GOAL / TARGET: Independent *rendered* verification that committed/discovered `MCSModbusToolkit` launches once as exactly one OS.js window showing `NOT CONNECTED — PLACEHOLDER ONLY`, while desktop, Start menu, taskbar and legacy Simulator/Replicator availability remain intact. Task authority: `workflow/active_work/umig-002-v-window-launch.md`. This is VERIFY, not a rerun of UMIG-002-T or a functional backend test.

REPOSITORY STATE: Create a disposable checkout of current `origin/main` of `tamzrod/MCS.OSJS`; record HEAD and pre-test `git status --short`. Verify `e9d25e3d4c1b832126a161a6071fd4abf8b5546b` (independent TEST PASS) is an ancestor of HEAD and that `UMIG-002-V` is the sole ACTIVE task using the current authoritative workflow files. Missing predecessor, task-state conflict, unexpectedly dirty tracked checkout, unavailable safe disposable environment or GUI = BLOCKED; do not reset, clean, restore or modify product files to force execution. An outdated ICC summary or unrefreshed ICC baseline is not a blocker for this packet.

OPTIONAL SAFE SETUP: Use sandbox-local compatible Node 16 / npm (repository engines Node >=10 <17). From a disposable checkout, `cd OSJS && npm install --no-audit --no-fund` is allowed to prepare dependencies if needed, provided tracked files remain unchanged. Use a fresh disposable OSJS_DATA_DIR (for example create it with `mktemp -d` outside the checkout) for VFS/session data; never use installed appliance or user data. Check whether localhost port 18209 is free before starting the server; if unavailable, report BLOCKED and request a packet change rather than selecting a different port. Prepare a browser/GUI test session if available without changing product source or user state.

EXACT COMMAND / ACTION: In that disposable checkout run, in this order, recording actual commands/output/exit codes: (1) `cd OSJS && npm run build:local-packages`; (2) `cd OSJS && npm run package:discover`; (3) `cd OSJS && npm run build`; (4) from `OSJS/`, start `OSJS_DATA_DIR="$TEST_DATA_DIR" PORT=18209 npm run serve` in a dedicated session, where `TEST_DATA_DIR` is the fresh directory created in safe setup, and retain startup logs. If a required build/start step fails, report FAIL with evidence and do not fabricate later observations. In a real browser/GUI navigate to `http://127.0.0.1:18209/`, observe the desktop, open the existing Start menu, locate and click `MCS Modbus Toolkit` **exactly once**, then inspect the rendered result: exactly one Toolkit window; literal disconnected placeholder; desktop, Start menu and taskbar remain functional; legacy Simulator and Replicator applications remain listed/available. Do not launch them or start any backend. Record how the window count and legacy availability were observed. If the browser/GUI or Start menu cannot be operated, or an alternate launch method seems necessary, report BLOCKED and request an exact packet update; no source/metadata/HTTP-only substitutes for rendered observation. Stop only the server process started for this test and preserve evidence; remove solely disposable test data if safe and unambiguous. Record post-test `git status --short` and any errors.

EXPECTED RESULT: Repo-native preparation/build/discovery/serve actions complete successfully; `http://127.0.0.1:18209/` renders the OS.js desktop; exactly one Toolkit window appears after one Start-menu launch, displaying `NOT CONNECTED — PLACEHOLDER ONLY`; Start menu, taskbar and desktop still function; legacy Simulator and Replicator remain available. No backend, Electron, Windows or final migration acceptance implied. A build/discovery exit 0 alone is never a VERIFY PASS.

EVIDENCE TO RETURN: HEAD SHA, predecessor and ACTIVE-task checks, exact toolchain/setup and commands with exit codes, pre/post tracked status, server URL/startup output, direct rendered screenshots if available or detailed first-hand GUI observations and browser method, launch interaction, Toolkit window count and placeholder text, desktop/Start/taskbar and legacy application observations, any errors, and safe cleanup. PASS only if **all** required UI observations are directly confirmed; FAIL if an executed product result contradicts expectations; BLOCKED if an indispensable UI or safety prerequisite is unavailable. No manual claims inferred from source or prior TEST.

REPORT-WRITE AUTHORITY: OpenHands/JR may replace **only** the `## JR TEST REPORT — UMIG-002-V` section immediately below, preserving every other section. JR may commit/push only `handoff.md` for that report and must STOP. No product fixes, workflow advancement, ICC edits, icon cutover, legacy app deletion, backend launches or next-task selection. ChatGPT reviews the evidence and decides completion independently.

## JR TEST REPORT — UMIG-002-V

VERDICT: **PASS** — every required rendered observation was directly confirmed in a real browser against a disposable checkout of `origin/main`.

### Repository-state checks

- Disposable checkout `git clone --depth 1 file:///workspace/project/MCS.OSJS /tmp/umigv`; HEAD `758a585817ce83bec40fc0fae764132dfece3efd`, equal to `origin/main` (`git ls-remote origin main` = `758a585817ce83bec40fc0fae764132dfece3efd`).
- Pre-test `git status --short`: empty. Post-test: empty. No repository file was modified.
- `git merge-base --is-ancestor e9d25e3d4c1b832126a161a6071fd4abf8b5546b HEAD` exits 0 (independent TEST PASS is an ancestor).
- `grep -rln "^Status: ACTIVE" workflow/active_work/` returns only `umig-002-v-window-launch.md`; 17 files are QUEUED. Sole ACTIVE task confirmed from authoritative workflow files.
- Port 18209 was free before start and free again after stop.

### Toolchain and safe setup (sandbox-local only)

Node 16.20.2 / npm 8.19.4 installed under `$HOME/.local/node16` (repository `engines` = `>=10.0.0 <17`). Playwright 1.63.0 + Chromium (headless shell 153) installed user-local under `$HOME/.local` / `$HOME/.cache/ms-playwright-agent` for real-browser driving. Fresh disposable `OSJS_DATA_DIR=$(mktemp -d)` outside the checkout. `cd OSJS && npm install --no-audit --no-fund` exited 0 (954 packages); tracked checkout remained clean.

### Commands and results

```text
(1) cd OSJS && npm run build:local-packages   -> exit 0
    "built 5 local packages exactly once: MCSModbusToolkit, ModbusReplicator,
     ModbusSimulator, NamelessClassicIcons, NamelessWorkstationTheme"
    Toolkit dist: main.js 1909 bytes, main.css 433 bytes
(2) cd OSJS && npm run package:discover       -> exit 0
    "7 package(s) discovered": mcs-modbus-toolkit as MCSModbusToolkit [symlink, local]
    OSJS/packages.json contains "src/packages/MCSModbusToolkit"
    OSJS/dist/metadata.json entry 0: type application, name MCSModbusToolkit,
      title.en_EN "MCS Modbus Toolkit"
    OSJS/dist/apps/MCSModbusToolkit -> src/packages/MCSModbusToolkit/dist
(3) cd OSJS && npm run build                  -> exit 0 (webpack 4.47.0, osjs bundle emitted)
(4) OSJS_DATA_DIR="$TEST_DATA_DIR" PORT=18209 npm run serve
    "Server listening on http://0.0.0.0:18209", "WebSocket listening on ws://0.0.0.0:18209"
    GET / -> 200 ; GET /healthz -> {"status":"ok","shell":"neutral"}
```

### Rendered browser verification (real Chromium, 1280x800)

```text
desktop_entries              ["Modbus Replicator"]      (desktop iconview wrapper present)
windows_before               0                          (clean baseline)
placeholders_before          0
clock_t0 "12:18:10" -> t1 "12:18:14"; clock_advanced = true
Start menu display after open  "block"
Start menu entries         ["Development", "MCS Modbus Toolkit", "Modbus Replicator",
                            "Modbus Simulator", "Save Session & Log Out", "Log Out"]
Start menu band text       "ROD DESKTOP" (intentional Win2000-style vertical band, OSUI-008)
launch clicks by operator   exactly 1 (real coordinate mouse click on the menu entry)
windows_after              1
placeholders_after         1
window_dump                [{dataId: "MCSModbusToolkitWindow", header: "MCS Modbus Toolkit",
                             placeholders: 1,
                             status: ["NOT CONNECTED — PLACEHOLDER ONLY"]}]
placeholder full text      "MCS Modbus Toolkit\nNOT CONNECTED — PLACEHOLDER ONLY\n\nMemory,
                            Replicator and Diagnostics are not connected yet. Existing
                            applications remain available."
panel_after               "MCS Modbus Toolkit\n12:18:20"  (taskbar window button appeared)
menu reopened display      "block"; entries unchanged with both legacy apps still listed
legacy_simulator_listed    true
legacy_replicator_listed   true
windows_final              1
page_errors                []
```

Independent launch-count confirmation from the browser console during the same single session:

```text
packages_launch_toolkit_calls  ["Packages::launch() MCSModbusToolkit {} {}"]  -> count 1
window_ctor_calls              ["Window::constructor() {id: MCSModbusToolkitWindow,
                                title: MCS Modbus Toolkit, ...}"]              -> count 1
```

Taskbar liveness (real mouse input, window-chrome minimize then taskbar restore):

```text
window minimize button (data-action="minimize") clicked
  -> data-minimized "true", data-focused "false", display "none"
taskbar window entry clicked (label "MCS Modbus Toolkit")
  -> data-minimized "false", data-focused "true", display "block"
placeholder still present after restore: 1
```

Every required acceptance item is directly confirmed: exactly one Toolkit window after exactly one Start-menu launch, the literal `NOT CONNECTED — PLACEHOLDER ONLY` placeholder, functional desktop/Start menu/taskbar with a live clock, and both legacy Simulator and Replicator still listed. Legacy applications were not launched and no backend was started.

### Non-blocking observations

- Four 404s during normal OS.js boot, all pre-existing and unrelated to the Toolkit window: `/apps/MCSModbusToolkit/icon.svg`, `/apps/ModbusReplicator/icon.png`, `/apps/ModbusSimulator/icon.png`, `/sounds/FreedesktopSounds/service-login.mp3`. The Toolkit menu entry does render with an icon in the menu DOM (the `/apps/.../icon.svg` background-image is applied); the 404 is the bare fetch. Reported as an observation only, not reclassified as a product failure by JR and not part of the acceptance criteria.
- OS.js session snapshot correctly persists window state between page loads; a rebuilt session restored the prior Toolkit window. All counts above come from a single clean session on a freshly wiped `OSJS_DATA_DIR`.

### Evidence artifacts (sandbox paths; not committed)

Screenshots (1280x800 PNG, Playwright/Chromium): `/tmp/jr-umigv-shots/DEFINITIVE-01-desktop.png`, `DEFINITIVE-02-startmenu.png`, `DEFINITIVE-03-toolkit-window.png`, `DEFINITIVE-04-menu-reopened.png`, `DEFINITIVE-05-taskbar-minimized.png`, `DEFINITIVE-06-taskbar-restored.png`, `tb-3-minimized.png`, `tb-4-restored.png`. Raw JSON: `/tmp/jr-umigv-definitive.json`, `/tmp/jr-umigv-final.json`, `/tmp/jr-umigv-report.json`. Server startup log: `/tmp/umigv-serve4.log`.

### Cleanup

Only the server process started by this test was stopped (port 18209 free again). All disposable `OSJS_DATA_DIR` directories were removed. No product, workflow, task-state or ICC file was touched.

### Evidence limits

This establishes rendered one-window VERIFY PASS only. No backend/Modbus connectivity, functional Memory/Replicator/Diagnostics acceptance, icon cutover, legacy package deletion, Windows COMMS/installer or Electron acceptance, or final migration acceptance is claimed or implied. The prior ICC-gate BLOCKED report remains preserved at `3831e33338a77fcea2e70e96d3f5127193fc5c10`.

## Evidence limits

UMIG-002-T independent build/discovery PASS is established. UMIG-002-V rendered UI acceptance, Windows COMMS/installer acceptance and final migration acceptance remain unverified.
