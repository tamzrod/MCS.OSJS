# Planning Workflow

Baseline commit: ee19b8a
Working tree: clean; no planning/workflow overlay remains uncommitted.
Source dependencies: planning/README.md, planning/Brainstorm/README.md, planning/microtask/README.md, planning/microtask/rules.md, workflow/active_work/README.md
Parent: L0-project
Zoom In: brainstorm-topics
Zoom Out: L0-project

## Boundary

Planning is human-owned and non-executable. Brainstorm captures exploration; `planning/microtask/`
contains independently sized task records; promotion moves a human-approved task into
`workflow/active_work/` and synchronizes `handoff.md`.

## Current inventory

- Brainstorm contains no topic files; only its README remains.
- Microtask now holds the staged OS.js Toolkit replacement series `UMIG-001` and
  `UMIG-003` through `UMIG-009` with their `-T` (TEST) and `-V` (VERIFY) stages, 23 files plus
  README and rules. All are PLANNED/BLOCKED pending human promotion. `UMIG-001` is the human
  donor-approval gate and `UMIG-007`/`UMIG-007A` are verification-only; neither manufactures
  coding work.
- The earlier Simulator, MMA2 and Replicator planning records were removed or
  promoted/archived by `ee19b8a` and remain absent.

## Rules that matter for current review

- One microtask has one primary outcome and one detailed task file.
- Scores use five dimensions, each restricted to 0-2.
- More than three independent acceptance outcomes, multiple verification workflows, or a total
  of 6+ require splitting.
- Coding, TEST and VERIFY are distinct by default, even for a small feature; only coding and
  either test stage are never combined.
- Planning never grants execution authority; only `workflow/active_work/` does.

## Current planning-to-active relationship

Only `UMIG-002` (CODE) and `UMIG-002-T`/`UMIG-002-V` have moved out of Planning: UMIG-002 is
archived as a source-only checkpoint and UMIG-002-T is the sole ACTIVE task. Every other UMIG
record stays in Planning, including `UMIG-001` donor approval and `UMIG-003` onward; no donor SHA
is approved and no renderer copy, launcher switch or legacy-package retirement is authorized.

`planning/microtask/README.md` and `rules.md` were both updated by this delta to add the TEST/VERIFY
stage split, stage ownership and the ICC-first rule restated for the Toolkit work.
