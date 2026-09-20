# Handoff

## Current authority — 2026-09-20

Human approved sequence: promote `UMIG-EM-003-V` first, then prepare (NOT activate) an OpenCode `UMIG-EM-004` CODE trial only after independent live verification and further promotion. `UMIG-EM-003-R-T` focused Replicator race TEST is archived COMPLETE/PASS; full original report at `workflow/archive/umig-em-003-r-t-passed-report.md`, report commit `04a31475c6427b185c09767703564894bd2791ed`. Earlier failed report remains archived. Unit tests do not prove cross-process runtime serialization.

SOLE ACTIVE: `workflow/active_work/umig-em-003-v-shared-lock.md` (VERIFY; HUMAN-PROMOTED, EXECUTION BLOCKED). No other task ACTIVE. `UMIG-EM-004` is PLANNED only; OpenCode coding trial NOT authorized to edit source, commit or push. Only BLACK SHEEP WALL edits ICC.

## JR TEST TASK — CURRENT: UMIG-EM-003-V — BLOCKED / NO EXECUTABLE PACKET

GOAL: independently verify shared writer-lock correctness under real simultaneous synthetic Simulator/Replicator applies in a NEW isolated disposable Docker project.

BLOCKER: exact new safe sandbox Docker host/daemon, unique project token, unshared volumes/ports, request commands and bounds, baseline/restore hashes and permitted label-scoped cleanup have not been confirmed for this run. The earlier OpenHands sandbox is not assumed to persist. The Legion `MCS.OSJS-jr` directory is a Git worktree, not a Docker/security sandbox. The current human approval promotes the task and permits planning, NOT arbitrary Docker operations or repurposing a running deployment.

EXACT ACTION NOW: Do NOT invoke Operation CWAL or execute any Docker, runtime/service, config-write or synthetic apply command. Coding agent must first identify and pin a genuinely disposable safe target, construct source-pinned exact preflight/execute/postcheck/cleanup/evidence/report packet, obtain any further safety approval required for that target, and then update this handoff. If invoked before then, JR returns BLOCKED: missing executable target and commands; JR STOP, no shell operations.

Pending transition: On independently accepted VERIFY PASS, coding agent may archive this task. `planning/microtask/umig-em-004-canonical-hydration.md` remains PLANNED; user approval is required for promotion and the separate OpenCode coding agent experiment. The independent JR identity for VERIFY stays OpenHands unless human specifically authorizes another one-task exception.
