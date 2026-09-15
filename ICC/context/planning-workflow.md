# Planning Workflow

Baseline commit: ee19b8a
Working tree: dirty; no planning/workflow overlay. See `ICC/INDEX.md` for the audited Electron overlay.
Source dependencies: planning/README.md, planning/Brainstorm/README.md, planning/microtask/README.md, planning/microtask/rules.md, workflow/active_work/README.md
Parent: L0-project
Zoom In: brainstorm-topics
Zoom Out: L0-project

## Boundary

Planning is human-owned and non-executable. Brainstorm captures exploration; `planning/microtask/` contains independently sized task records; promotion moves a human-approved task into `workflow/active_work/` and synchronizes `handoff.md`.

## Current inventory

- Brainstorm contains no topic files; only its README remains.
- Microtask contains no executable task records; only its README and sizing rules remain.
- Prior Simulator, MMA2, and Replicator planning records were removed or promoted/archived by `ee19b8a`.

## Rules that matter for current review

- One microtask has one primary outcome and one detailed task file.
- Scores use five dimensions, each restricted to 0-2.
- More than three independent acceptance outcomes, multiple verification workflows, or a total of 6+ require splitting.
- Planning never grants execution authority; only `workflow/active_work/` does.

## Current planning-to-active relationship

There are no Electron or OS.js task records in Planning. ELECTRON-001 and ELECTRON-002 were retired by human decision after the initial Electron goal was accepted as achieved; their records are preserved in `workflow/archive/` without claiming their remaining verification gates passed.
