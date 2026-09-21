# OTR-002 — OS.js Toolkit baseline inventory
Status: PLANNING / UNDER REVIEW. Stage: DISCOVERY. Owner: OpenCode. Previous: OTR-001. Next: OTR-003.

Outcome: Produce a read-only inventory of the existing unified OS.js Toolkit, UI entry points, styles, config/backend interfaces and test harness, anchored to exact paths and source revision. Scope: inspect `OSJS/src/packages/MCSModbusToolkit/` and directly imported integration files only; document existing Simulator, Replicator and MMA views and what is actually implemented. Non-scope: Electron changes, implementation, live tests, installs, service or configuration writes.

Acceptance: (1) Evidence-linked OS.js UI/file map. (2) Existing backend/config capabilities distinguished from missing/unknown ones. (3) Reusable components and parity gaps recorded against OTR-001 without assuming hidden tab infrastructure.

Evidence: paths, function/line anchors, commit, inventory artifact and unknowns; inspection is not a test PASS. Dependencies: OTR-001 inventory. Size (implementation/environment/behavior/verification/decision): 0/1/1/1/1 = 4, one bounded inventory. At promotion pin exact inspection boundary and evidence file. STOP; do not execute successor.
