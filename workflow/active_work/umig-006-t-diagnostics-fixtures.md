# UMIG-006-T — TEST: Diagnostics Read-only Mapping and Build

Status: ACTIVE — promoted 2026-09-19 after UMIG-006 CODE source-only checkpoint; TEST NOT RUN / NO PASS.
Stage / owner: TEST / OpenHands JR via `operation cwal.md`
Previous: UMIG-006 (COMPLETE source-only, `workflow/archive/umig-006-adapt-diagnostics-tab.md`)
Next: UMIG-006-V (QUEUED; independent TEST PASS required)

## Primary outcome
Independently verify Diagnostics' canonical read-only Memory/Replicator mapping, fail-closed and disabled native controls, async teardown, and OS.js build/discovery while preserving existing Memory and Replicator behavior.

## Exact authority
Only `## JR TEST TASK — CURRENT: UMIG-006-T` in `handoff.md` is executable. Run the exact ordered commands in a disposable Node16 checkout and record raw command/exit, model/observer/editor assertions, sources/bundle, bounded diff and final clean tracked status. Tests and build have NOT been run by the CODE author. If a required product assertion/build fails, FAIL and STOP product testing; do not edit source or debug to force PASS. Missing evidence after reasonable allowed sandbox-local prep = BLOCKED.

## Non-scope
No actual backend, Docker, runtime/device/config write, browser live verification, Windows service control, network probe, Go/MMA2 changes, product fixes, ICC or workflow edits by JR. UMIG-006-V remains QUEUED. The former UMIG-005-V real Toolkit close/relaunch gap belongs ONLY to its later separately authorized browser packet, not this UNIT/BUILD task.

## Dependencies
Source checkpoint `c996794ff14dd65479e84123724f6b9fd9235ee6` is ancestor of test HEAD. Preserve production, legacy applications and retained project volumes.

## Sizing
Surface 0, environment 0, behavior 0, verification 1, recovery 0 = 1.
