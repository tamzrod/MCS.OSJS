# OTR-001 — Electron Toolkit reference inventory
Status: PLANNING / UNDER REVIEW. Stage: DISCOVERY. Owner: OpenCode (assigned only after promotion). Previous: brainstorm `planning/Brainstorm/osjs-toolkit-electron-replica.md`. Next: OTR-002.

## Primary outcome
Produce an evidence-linked inventory of the Electron Toolkit's actual UI screens and implementation files; establish the reference without guessing OS.js architecture.

## Scope
Read-only inspect `electron/` Toolkit entry, renderer modules, CSS/assets and navigation. Record exact paths and line/function anchors for Simulator, Replicator, MMA memory, dialogs, tabs, state variants and available screenshot references. Write a concise inventory as the sole deliverable under the approved planning/evidence path at execution time; no product modifications.

## Non-scope
No OS.js edits, code port, backend changes, builds, runtime tests, installs or live service operations. No claims about behavior not supported by code or observed evidence.

## Acceptance (max 3)
1. Inventory names every observed Toolkit screen/navigation path and marks unverified UI states explicitly.
2. Maps each observed screen to concrete Electron source and style/asset paths, including relevant entry points.
3. Separates reusable renderer logic from Electron/Windows-specific calls with cited source anchors; unresolved contracts are listed.

## Evidence / handoff
Return exact inspected paths, relevant line/function anchors, source revision and inventory document; source inspection only, no test PASS. Stop after this task. Do not automatically execute OTR-002.

## Dependencies and size
Brainstorm only. Five-dimension size: implementation 0, environment 1, behavior 1, verification 1, decision/recovery 1 = 4; tightly coupled read-only inventory with one output. Exact file allowlist and artifact path to be pinned at review/promotion; do not expand scope silently.
