# CLEANUP-001-OC — OpenCode review of legacy retirement and default Auto Arrange

**Status**: READY FOR EXPLICIT ASSIGNMENT  
**Stage**: CODE REVIEW / TARGETED SELF-TEST (not an independent JR TEST/VERIFY verdict)  
**Owner**: OpenCode  
**Product checkpoint**: `6465ec8b9f558184ab9e53e2a214b5b95a6f173c`  
**Prepared branch**: `cleanup/retire-legacy-components`

## Objective

Review the cleanup produced at the product checkpoint and verify that the intended retirements are complete without removing the current MCS Modbus simulator runtime. Verify the fresh-profile desktop behavior added for Auto Arrange. Make only narrowly scoped corrective source edits if a defect is found; otherwise report the checks and STOP.

This packet does **not** authorize branch switching, merging, cherry-picking, resetting, stashing, cleaning, pushing, or creating a PR. OpenCode must remain on its mandatory `opencode` branch. If the product checkpoint is not already an ancestor of the current `opencode` checkout, report `BLOCKED: cleanup checkpoint not present in opencode checkout` and STOP. A human may integrate/synchronize the checkpoint separately.

## Required preflight

Run exactly:

```bash
git branch --show-current
git rev-parse HEAD
git status --short
git worktree list
git merge-base --is-ancestor 6465ec8b9f558184ab9e53e2a214b5b95a6f173c HEAD
```

Expected:
- current branch is exactly `opencode`;
- workspace is safe for the bounded check;
- ancestry command exits 0.

## Verification scope

1. **Modbus Sim retirement**
   - `deploy/pretest/` and the old `modbus-sim` pretest service must be absent.
   - Do **not** remove or rename `modbus-simulator-runtime`, `simulator/`, or MCS Toolkit runtime integration.

2. **DNP3 Sim retirement**
   - the old `dnp3-sim` pretest service and `deploy/pretest/dnp3-sim/` must be absent.
   - Historical/provenance text that merely says DNP3 material was excluded is not an active simulator and need not be deleted.

3. **Calculator retirement**
   - no `@osjs/calculator-application` dependency/package registration/default shortcut remains;
   - no calculator-specific package-discovery/test/doc/icon residue remains;
   - `accessories-calculator.svg`, its icon metadata entry, and its CSS mapping are absent.

4. **TES/Test Application retirement**
   - no operational `TestAppApplication` / **Test Application** package remains.
   - Historical third-party provenance that states TestApp was excluded is not an operational app.

5. **Auto Arrange default**
   - manual desktop context-menu **Auto Arrange** remains available;
   - an existing non-empty `nameless/desktop.positions` map is respected and not overwritten on boot;
   - when no positions exist, the provider waits for the initial icon render to settle, runs the deterministic label-sorted column-first layout, and persists the resulting positions;
   - drag persistence still writes a moved icon back to the same settings map.

## Exact source checks

Run:

```bash
test -z "$(git ls-tree -r --name-only HEAD -- deploy/pretest)"
! git grep -nE '(^|[^[:alnum:]_-])modbus-sim([^[:alnum:]_-]|$)|dnp3-sim' -- .
! git grep -nE '@osjs/calculator-application|accessories-calculator' -- .
! git grep -nE 'TestAppApplication|Test Application' -- ':(exclude)THIRD_PARTY_NOTICES.md'
git grep -nE 'DEFAULT_ARRANGE_DELAY|defaultArrangeTimer|defaultArrangePending|autoArrange\(core\)|writePositions\(core, positions\)' -- OSJS/src/client/providers/nameless-desktop-icons.js
```

The negative greps must exit 0 because `!` inverts an expected no-match exit. The final grep must show the fresh-profile initialization/persistence anchors.

## Build check

Without installing or downloading dependencies, run:

```bash
cd OSJS
NODE_OPTIONS=--openssl-legacy-provider npm run build
cd ..
git status --short
```

Expected: build exits 0. The final status must contain no unexpected tracked source changes. Do not clean or reset generated output; if tracked files changed, report them and STOP.

## Review result

Return:
- tested HEAD;
- preflight outputs;
- exact source-check outputs/exits;
- build exit and concise output summary;
- final `git status --short`;
- any narrowly scoped defect/correction made.

This is an OpenCode review/self-test of work authored by another agent, not an independent JR certification. Do not edit ICC, handoff, queues, unrelated OTR/FMT tasks, production YAML, or runtime/device configuration. STOP after this one task.
