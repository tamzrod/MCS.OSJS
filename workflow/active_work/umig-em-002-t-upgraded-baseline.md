# UMIG-EM-002-T — TEST: Upgraded Go backend baseline

Status: ACTIVE — first regression stage authorized after human approval of UMIG-EM-001 contract on 2026-09-19. Stage/owner: TEST / OpenHands JR; packet and adjudication / ChatGPT.
Previous: UMIG-EM-001 (archived COMPLETE design); Next: UMIG-EM-002-B (PLANNED only; separately human-promote before advancement).

## Primary outcome
Independently verify Go unit/regression tests for the already merged MMA2 RBE, producer composer, Simulator advanced-field persistence and Replicator measured COMMS source. The OS.js Node/build checks are a different workflow in UMIG-EM-002-B; live behavior in UMIG-EM-002-V. No CODE stage or new source implementation is implied.

## Exact scope and acceptance
1. In a clean disposable checkout at latest main descended from approved checkpoint `68c41227d841609e6d73d77e1f87eb325499773c`, establish Go >=1.25 using sandbox-local tooling if needed. Do not run on production/operator checkout or use installed data.
2. Execute four separately reported full Go suite commands (MMA2, mma2composer, simulator, replicator), all exit 0; execute four pinned focused Go regression commands, all exit 0, with observed named test PASS (RBE validation; composer partial-write restoration; advanced-field persistence; per-cycle COMMS/error/aggregation). Exact shell commands and order, stop behavior, timeouts and evidence are exclusively in `handoff.md`.
3. Show clean tracked worktree before and after and identify only local test execution. A failed product test is FAIL + raw evidence + STOP; unpreparable sandbox toolchain is BLOCKED. No reruns to force PASS and no substituted old commit test output.

Non-scope: OS.js Node/build, Docker/live GUI, RBE listener activation/exposure, production endpoints, Windows installer, source fixes, workflow/ICC/CWAL edits, claims about cross-process lock/shared manager or COMMS in OS.js. Only authorized edit is current JR report in `handoff.md` and handoff-only commit/push.

Evidence: raw commands, exits, per-module package results, focused test names, Go/toolchain provenance, git HEAD/status and unexpected findings. Size surface 0, environment 1, behavior 0, verification 1, recovery 1 = 3.
