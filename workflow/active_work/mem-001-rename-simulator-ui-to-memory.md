# MEM-001 — Rename Simulator UI to Memory

Status: ACTIVE
Previous: none
Next: MEM-002

## Primary Outcome
Rename the operator-facing Simulator workspace into Memory without changing internal simulator services, APIs, persistence paths, or runtime behavior.

## Scope
- Rename the top-level `Simulator` tab label to `Memory` while retaining the existing internal tab/panel identifiers.
- Rename `Simulator Devices` to `Devices` in the left tree.
- Rename the memory table heading `Function` to `Area`.
- Update nearby operator-facing subtitle wording from Simulator to Memory where it describes the workspace rather than the internal runtime.

## Non-Scope
- No backend changes.
- No service/process/API/path renames.
- No Simulation mode selector yet.
- No status behavior changes.

## Acceptance Criteria
1. Top-level operator tab reads `Memory`.
2. Left tree heading reads `Devices`.
3. Device Definition memory table first column reads `Area`.
4. Existing device editing, Save & Apply, runtime calls, and persisted document shape are unchanged.

## Verification
Static source re-read of `electron/renderer/index.html` and `electron/renderer/app.js`; final rendered verification is covered by MEM-008.

## Dependencies
None.

## Sizing
Implementation 1; environment 0; behavior 0; verification 1; decision/recovery 0. Total 2.