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

VERDICT: **FAIL** — executed product test contradicts a stated expected result.

The two exact commands both exited 0, but OS.js package discovery does **not** list the Toolkit
package. The packet requires "OS.js discovery lists the Toolkit package" as a required result, so the
overall verdict cannot be PASS.

### Repository state (disposable checkout)

- Clone: `https://github.com/tamzrod/MCS.OSJS.git` → `/tmp/jr-umig-002-t` (outside the product repo).
- `git rev-parse HEAD` (pre): `db6b4edd98390d9bbb0e5d66e5525471a9e224be`
- `git status --short` (pre): *(empty — clean)*
- Required source commit `9410edccd3f8914cc76a365980c0d0f1e9bd8bbb` is an ancestor of HEAD: verified
  `git merge-base --is-ancestor 9410edccd3f8914cc76a365980c0d0f1e9bd8bbb HEAD` → included.
- `git status --short` (post): *(empty — no unexpected tracked modifications)*

Note: this checkout is one commit past the recorded ICC baseline `7029e41`. The extra commit
`db6b4ed` changes only `ICC/`, and the Toolkit source under test is byte-identical at both commits.

### Toolchain / safe setup

- Host Node was `v22.23.2` (npm `10.9.8`); `OSJS/package.json` declares `"node": ">=10.0.0 <17"`.
- Sandbox-local preparation only: `cd OSJS && npm install --no-audit --no-fund` → exit 0,
  `added 953 packages in 21s`. No tracked manifest, lockfile, config or product source was edited.
- `OSJS/.gitignore` ignores `node_modules/`, `dist/`, `packages.json`, `package-lock.json`, so all
  generated output stayed untracked and `git status` stayed clean.
- Webpack 4 / sass legacy-JS deprecation warnings were emitted; they did not fail the build.

### Command 1 — `cd OSJS && npm run build:local-packages`

Exit code: **0**

Raw output (tail):

```
> mcs-osjs-base-desktop@0.1.0 build:local-packages
> node scripts/build-local-packages.js

build-local-packages: building src/packages/MCSModbusToolkit
Hash: 1d89f44b2867d440a75b
Version: webpack 4.47.0
Time: 876ms
       Asset       Size  Chunks                   Chunk Names
    main.css  433 bytes       0  [emitted]        main
main.css.map  780 bytes       0  [emitted] [dev]  main
     main.js   1.89 KiB       0  [emitted]        main
 main.js.map   6.82 KiB       0  [emitted] [dev]  main
...
build-local-packages: built 5 local packages exactly once: MCSModbusToolkit, ModbusReplicator, ModbusSimulator, NamelessClassicIcons, NamelessWorkstationTheme
```

Toolkit build artifacts observed on disk:

```
src/packages/MCSModbusToolkit/dist/main.js    1934 bytes  -> exists: YES
src/packages/MCSModbusToolkit/dist/main.css    433 bytes  -> exists: YES
src/packages/MCSModbusToolkit/dist/main.js.map
src/packages/MCSModbusToolkit/dist/main.css.map
```

### Command 2 — `cd OSJS && npm run package:discover`

Exit code: **0**

Raw output (verbatim, complete package list):

```
> mcs-osjs-base-desktop@0.1.0 package:discover
> osjs-cli package:discover

ℹ Discovering packages...
ℹ Destination discovery map /tmp/jr-umig-002-t/OSJS/packages.json
ℹ Destination path /tmp/jr-umig-002-t/OSJS/dist
ℹ Destination manifest /tmp/jr-umig-002-t/OSJS/dist/metadata.json
ℹ Including /tmp/jr-umig-002-t/OSJS/node_modules
ℹ Including /tmp/jr-umig-002-t/OSJS/src/packages
- modbus-replicator as ModbusReplicator [symlink, local]
- nameless-classic-icons as NamelessClassicIcons [symlink, local]
- modbus-simulator as ModbusSimulator [symlink, local]
- nameless-workstation-theme as NamelessWorkstationTheme [symlink, local]
- @osjs/gnome-icons as GnomeIcons [symlink, npm]
- @osjs/standard-theme as StandardTheme [symlink, npm]
ℹ Flushing out old discoveries...
ℹ Placing packages in dist...
✔ 6 package(s) discovered.
✔ Finished in 149ms
```

`MCSModbusToolkit` is **absent** from the discovered list.

### Discovered-manifest evidence

`packages.json` (complete):

```
["src/packages/ModbusReplicator","src/packages/NamelessClassicIcons","src/packages/ModbusSimulator","src/packages/NamelessWorkstationTheme","node_modules/@osjs/gnome-icons","node_modules/@osjs/standard-theme"]
```

- `grep -c MCSModbusToolkit packages.json` → **0**
- `grep -c MCSModbusToolkit dist/metadata.json` → **0**
- `ls dist/apps/` → `ModbusReplicator  ModbusSimulator` (no Toolkit; `dist/apps/MCSModbusToolkit` does not exist)

### Observed cause (reported, not fixed)

Discovery selects local packages by globbing for `package.json`:

`node_modules/@osjs/cli/src/utils.js:38`
```js
const globs = await globby(root.replace(/\\/g, '/') + '/**/package.json', {deep: 3});
```

`OSJS/src/packages/MCSModbusToolkit/` contains **no `package.json`**; the other four local packages
each do and are discovered:

```
src/packages/MCSModbusToolkit/            package.json: NO
src/packages/ModbusReplicator/            package.json: YES
src/packages/ModbusSimulator/             package.json: YES
src/packages/NamelessClassicIcons/        package.json: YES
src/packages/NamelessWorkstationTheme/    package.json: YES
```

The Toolkit ships `metadata.json` instead, while its siblings' `package.json` carry the
`"osjs": {"type": "package"}` marker that discovery consumes. This is the reason the Toolkit is
skipped. It is a product defect; per Operation CWAL, JR reports it and does not fix it.

### Expected-result comparison

| Packet requirement | Observed | Result |
| --- | --- | --- |
| `npm run build:local-packages` exits 0 | exit 0 | pass |
| `npm run package:discover` exits 0 | exit 0 | pass |
| Toolkit `dist/main.js` exists | found, 1934 bytes | pass |
| Toolkit `dist/main.css` exists | found, 433 bytes | pass |
| OS.js discovery lists the Toolkit package | **absent from all discovered output** | **fail** |

Overall: **FAIL**. Not BLOCKED — the requested test executed and produced a result that contradicts
the stated expectation; the environment and target were available.

### Not performed (packet non-scope)

No GUI launch, no Electron build/install, no backend service startup, no saved-configuration
mutation, no desktop change, no browser PASS claim, no source fix, no workflow advancement, no ICC write.

## Evidence limits

Windows RLED appearance was accepted by the human, but installed live COMMS status and RREC-004 installer repair/hash/ProgramData verification are not established by this OS.js source milestone. Do not claim Windows tests or the final OS.js UI migration passed.
