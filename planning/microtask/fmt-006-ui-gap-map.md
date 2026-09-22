# FMT-006 — UI gap map
Task ID: FMT-006
Task Name: Map Electron-to-OS.js shell and navigation gaps
Blocker Task: FMT-001, FMT-002, FMT-003 [RESOLVED]
Status: BLOCKED->COMPLETING (blockers satisfied, awaiting explicit build command)
Assigned Agent: OpenCode — read-only design
Stage: COMPLETING

## Objective
Compare only source-verified shell/navigation behavior, not all individual editor views, and identify exact gaps without rewriting working code.
## Scope
Read `ICC/INDEX.md`, accepted FMT-001/002/003 evidence and only their pinned shell/navigation source files. Write only `planning/microtask/evidence/fmt-006-shell-navigation-gap-map.md` when promoted. No source, ICC, historical evidence or runtime changes.
## Execution
1. Verify blocker reports and HEAD relevance; identify exact Electron and OS.js shell/navigation source paths.
2. Map each source-defined navigation item to target implementation: existing, missing, differing or unknown.
3. For each proven gap propose one bounded future CODE outcome with exact affected paths, or separate discovery if unknown; defer editor internals and screenshots.
## Acceptance Criteria
1. Source-linked shell/navigation comparison with supported classifications.
2. Only proven gaps generate candidate work; existing features explicitly marked no-change.
3. Rendered parity remains unverified unless separately evidenced; view-specific gaps are deferred.
## Evidence
`planning/microtask/evidence/fmt-006-shell-navigation-gap-map.md`: SHA, per-item source/target lines, gap and no-change decisions, unknowns.
## Completion and delivery
COMPLETE when all outcomes evidenced; FAIL if comparison unsupported; BLOCKED if prerequisites stale/unverified. Report and authorized task status only; no commit/push authority until promotion supplies exact ref/command. STOP.
## Sizing
Implementation 0; environment 0; behavior 1; verification 1; decision/recovery 1 = 3/10. Shell/navigation only, not entire UI.