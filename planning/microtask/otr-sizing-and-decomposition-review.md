# OTR microtask sizing and decomposition review
Status: PLANNING / UNDER REVIEW; no execution authority. Date: 2026-09-21. Source: `planning/Brainstorm/osjs-toolkit-electron-replica.md`, OTR-001..014 and `planning/microtask/rules.md`. This is a review map, NOT a runnable task. Existing OTR IDs remain phase/parent references when split. No active_work/handoff/ICC edits.

## Decision criteria
Five dimensions I/E/B/V/D each 0..2. Preferred sum 0..3; 4..5 split unless genuinely one deterministic outcome; 6..7 split; 8..10 mandatory split. Independent discovery vs code, backend vs UI, multiple runtimes/screens, CODE vs JR TEST vs JR VERIFY always split. Every executable child needs its OWN detailed file, single owner/stage/outcome, <=3 production files plus tests, <=3 acceptance checks, source checkpoint, exact allowed paths, exact commands/actions and evidence. Source paths, backend contracts and UI screenshots not yet established are NOT invented here. Only promote children after predecessor inventory identifies them. No guessed PASS or implicit successor execution.

## All 14 reviewed
| Parent | Existing I/E/B/V/D | Total | Disposition / smallest honest work units |
|---|---|---:|---|
| OTR-001 Electron inventory | 0/1/1/1/1 | 4 | Split inventory into shell/navigation; editor screen/asset mapping; backend-boundary/source contract inventory. Each inspection-only artifact; no implementation. |
| OTR-002 OS.js inventory | 0/1/1/1/1 (provisional) | 4 | Split UI/launch baseline from backend/config/service capability baseline. Read-only. |
| OTR-003 parity map | 0/1/1/1/1 (provisional) | 4 | Split UI source-to-target map from config/backend contract map. Do not choose unverified adapters. |
| OTR-004 shell | 2/1/1/1/1 | 6 | Split shell structure/navigation CODE and CSS/assets visual parity CODE; separate JR static TEST and rendered VERIFY. |
| OTR-005 Simulator UI | 2/1/1/1/1 | 6 | Split screen shell/navigation, one editor view/tab per mapped reference, and draft/discard state behavior. Do not invent tabs; exact number after inventory. Separate TEST/VERIFY. |
| OTR-006 Replicator UI | 2/1/1/1/1 | 6 | Split screen shell/navigation, one mapped view/tab per child, draft/discard state. Separate TEST/VERIFY. |
| OTR-007 MMA UI | 2/1/1/1/1 | 6 | Split memory overview/list, one mapped editor view, and read-only/draft states where independently verifiable. Separate TEST/VERIFY. |
| OTR-008 read adapter | 1/1/1/1/1 | 5 | Three separate runtime read contracts: Simulator, Replicator, MMA; split backend adapter from UI wiring if >3 files or distinct outcomes. Each gets JR TEST. |
| OTR-009 validation | 1/1/1/1/1 | 5 | Separate each verified configuration schema/runtime; split validator from serializer/unknown-field roundtrip if distinct. JR TEST for each. |
| OTR-010 safe save | 2/1/1/1/1 | 6 | Per runtime: (a) disposable backend write/conflict/failure contract, (b) UI save wiring; split conflict/atomicity further if independent. JR TEST each; no production writes. |
| OTR-011 restart | 2/2/1/1/1 | 7 | First read-only mechanism/permission discovery PER runtime; only where verified, per-runtime backend restart adapter then UI action wiring; independent mock/disposable TEST and runtime VERIFY if safe. Unsupported restart remains explicitly unsupported. |
| OTR-012 status | 1/1/1/1/1 | 5 | Per runtime: status observation adapter then UI status states if independently testable. Never green from acknowledgement. JR TEST/VERIFY separate. |
| OTR-013 UI VERIFY | 0/1/1/2/1 | 5 | One exact screen and scenario per independent JR packet: shell, Simulator, Replicator, MMA; additional dialogs/errors only where distinct. Pin viewports/screenshots/actions. |
| OTR-014 E2E VERIFY | 0/2/2/2/2 | 8 | Per supported runtime: one happy-path disposable workflow; separate invalid input, discard, save conflict/write failure, restart failure, stale/applied-state failure scenarios. Exact commands, safety and report transport before promotion. |

## Optimal sequence (not blanket promotion)
1. Finish atomic read-only Electron and OS.js inventories, then UI and backend parity maps. Only those can fix true screen counts, source paths and contracts.
2. Port shell structure then visuals; port each real editor view independently; draft/discard only where not already intrinsic to that view. Avoid duplicate abstractions and arbitrary per-tab work when one inseparable component is already <=3 files.
3. For each runtime, implement read -> validate/roundtrip -> safe write -> supported restart -> independently observed status. Do not hold all three runtimes behind one large task or assert a universal restart endpoint.
4. Independent JR static/unit TEST after each code checkpoint, rendered VERIFY per screen, then disposable E2E per runtime/scenario. Failed gates produce a separately reviewed repair task.

## Review blockers, not execution blockers
OTR-001..003 are source discovery: their exact path lists and commands must be pinned at promotion; remaining implementation/verification children cannot yet be fully executable without their outputs. Existing parent OTR-004..014 are NOT promotion-ready as written, even where they say 'one selected runtime' but omit which one. Replace parent execution with its individual child task files before promoting. Never make one giant task out of this review or automatically start its next item.
