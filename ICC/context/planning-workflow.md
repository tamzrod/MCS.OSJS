# Planning Workflow

Baseline commit: 5a89194e4acc710032c00a26bb5d2a1a095eacd1
Working tree: clean outside this ICC maintenance pass; no planning/workflow overlay remains uncommitted.
Source dependencies: planning/README.md, planning/Brainstorm/README.md, planning/microtask/README.md, planning/microtask/rules.md, workflow/active_work/README.md
Parent: L0-project
Zoom In: brainstorm-topics
Zoom Out: L0-project

## Boundary

Planning is human-owned and non-executable. Brainstorm captures exploration; `planning/microtask/`
contains independently sized task records; promotion moves a human-approved task into
`workflow/active_work/` and synchronizes `handoff.md`.

## Current inventory

- Brainstorm retains its README plus one live topic, `osjs-toolkit-electron-replica.md` (OS.js Toolkit as an Electron-Toolkit replica; BRAINSTORM/REVIEW ONLY, promoting no microtask).
- `planning/microtask/` holds the OTR series that succeeded the retired UMIG planning backlog: `otr-001`/`001a`/`001b`/`001c`, `otr-002`/`002a`/`002b`, `otr-003`/`003a`/`003b`, the `otr-004`..`otr-014` parent templates, plus `otr-remaining-atomic-decomposition.md`, `otr-sizing-and-decomposition-review.md` and `rules.md` — 25 Markdown files including the README. The earlier UMIG planning records no longer live under `planning/`; they were archived under `workflow/archive/`. All `planning/microtask/` items remain non-executable until human promotion.
- OTR-004..014 are parent topics, not executable children; they must be materialized as source-pinned atomic children after inventories.

## Rules that matter for current review

- One microtask has one primary outcome and one detailed task file.
- Scores use five dimensions, each restricted to 0-2.
- More than three independent acceptance outcomes, multiple verification workflows, or a total of 6+ require splitting.
- Coding, TEST and VERIFY are distinct by default, even for a small feature.
- Planning never grants execution authority; only promotion into `workflow/active_work/` does. `rules.md` now also routes its ICC-first rule through `workflow/ICC_CONSUMER.md` with per-node baseline/overlay/validity rather than an index-wide snapshot, and equally states that a migrated `unverified` node does not justify a global audit.

## Current planning-to-active relationship

`workflow/active_work/README.md` records that the human approved the OTR roadmap and promotion and that this directory now contains the ACTIVE OTR-001A packet, queued discovery packets and copied OTR-004..014 parent planning templates. The authoritative selector is `handoff.md` plus `workflow/active_work/OTR_PROMOTION_QUEUE.md`; this node does not resolve task assignment. The former UMIG-002/UMIG-002-T assertion previously cached here is superseded and must not be restated as current.
