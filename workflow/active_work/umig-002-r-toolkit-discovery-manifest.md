# UMIG-002-R — CODE: Repair Toolkit Package Discovery Manifest

Status: ACTIVE — human authorized 2026-09-18 after UMIG-002-T FAIL.
Stage / owner: CODE / ChatGPT
Previous: UMIG-002-T (failed test; repair interruption)
Next: UMIG-002-T (retry same test; resume existing queued task)

## Primary outcome
Make the existing independently built `MCSModbusToolkit` discoverable by OS.js by supplying its missing local-package manifest.

## Scope
Add exactly `OSJS/src/packages/MCSModbusToolkit/package.json`, using the same local-package convention as the existing Simulator/Replicator packages: stable lowercase npm package name, version, private flag, description and `osjs.type = package`. Retain Toolkit `metadata.json` as the application metadata. Inspect/read back the committed file and record the source commit. Resume the original UMIG-002-T test with an updated exact JR packet; do not reinterpret the prior FAIL as PASS.

## Acceptance / evidence
1. New manifest is valid JSON and includes the OS.js discovery marker, with no unrelated product source changes.
2. Committed source file is read back and source SHA recorded; this establishes only source checkpoint, not build/discovery PASS.
3. Existing UMIG-002-T is restored to ACTIVE only after CODE repair is archived and a fresh retest packet is placed in `handoff.md`.

## Non-scope
No UI feature additions, renderer refactor, OS.js discovery-engine changes, OS.js runtime/GUI or backend tests, Electron/Windows edits, task advancement beyond reactivation of the already-authorized UMIG-002-T, or ICC writes. TEST belongs to OpenHands/JR under Operation CWAL. The queued UMIG-002-V remains unauthorized until TEST PASS and explicit coding-agent advancement.

## Dependencies
The previous TEST failure is documented in `handoff.md` at `752a541`; original UMIG-002-T is temporarily QUEUED. Preserve the failed report as historical evidence, and have BLACK SHEEP WALL refresh affected ICC context before JR rerun.

## Sizing
Surface 1, environment 0, behavior 0, verification 0, recovery 1 = 2.
