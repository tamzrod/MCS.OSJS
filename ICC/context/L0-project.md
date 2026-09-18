# L0 Project Context

Baseline commit: ee19b8a
Working tree: clean
Source dependencies: README.md, PROJECT_IDENTITY.md, handoff.md
Parent: none
Zoom In: governance, donor-licensing, network-exposure, planning-workflow, osjs-shell, active-work
Zoom Out: none

## Identity

MCS.OSJS = Modbus Consolidation System; standalone repo/lineage rebuilt around an OS.js application shell
and proven Modbus runtime components. Planned as single-container appliance: OS.js presentation, Orchestrator
lifecycle/control authority, Modbus Replicator acquisition, MMA2 deterministic Modbus memory appliance. It is NOT part of Nameless SCADA, but reuses proven generic parts from it( and may later be migrated into it after proving standalone-first.

## Construction Strategy

Staged: scaffold -> donor inventory(Nameless SCADA, then Modbus Replicator stack( -> program architecture+rewiring -> UI plan -> microtasks -> promotion -> Operation CWAL execution -> verify appliance. Completed so far:  scaffolding, Nameless SCADA donor inventory + OS.js shell harvest (executed,head of sim/MMA2 work(; Modbus Replicator stack inventory and remaining architecture/UI stages proceed through planning+workflow.

## Development Model

Brainstorm -> microtask -> promotion -> CWAL -> verified change. Context maintenance via BLACK SHEEP WALL and ICC. One task = one primary outcome. Reuse proven machinery, but preserve explicit boundaries; no component gains authority merely for convenience. Standalone-first:  stabilize the standalone appliance before any optional Nameless SCADA migration.



## Authority Boundaries

Project identity document defines identity/direction only:  no implementation authority, no settled architecture. Detailed architecture, donor selection, rewiring, UI behavior, and execution scope established via planning + workflow docs. Repository files are authoritative; conversation memory is not a substitute.

## Current Execution State (handoff)

Handoff status: a new staged UI migration is in progress and is now the human priority. The
2026-09-18 direction is a staged OS.js replacement with a single `MCS Modbus Toolkit`: preserve the
OS.js desktop, Start menu, taskbar and clock; end with exactly one Toolkit desktop icon; remove legacy
UI packages only after verified replacement and cutover. MMA2, the Go services, shared memory, user
configuration and the Windows Electron app are explicitly NOT decommissioned by this migration.

Role split recorded in `handoff.md`: ChatGPT owns CODE, source changes, source checkpoints and task
advancement; OpenHands is JR for the separate TEST and VERIFY stages under `operation cwal.md`, does
not code, fix failures, promote/archive tasks, or write ICC. Only BLACK SHEEP WALL updates ICC.

Current position: CODE `UMIG-002` is archived as a source-only checkpoint (placeholder
`MCSModbusToolkit` package); `UMIG-002-T` (Toolkit build and discovery TEST) is the sole ACTIVE task
with a current `JR TEST TASK` packet but NO execution result; `UMIG-002-V` (rendered window VERIFY) is
QUEUED. No donor SHA is approved, no Electron renderer has been copied, no launcher switch happened,
and no legacy OS.js app has been deleted. Detail in Zoom In `active-work`; planning inventory in
`planning-workflow`.

Historical baseline still relevant where untouched: MMA2-001/002 and SIM-001 through SIM-024 are
completed and verified, with SIM-017's real-MMA2 capstone at `57c8714` and predecessors archived at
`6069fef`; SIM-009 was superseded and SIM-022 retired. Lineage detail is now stale and lives outside
this node's scope; use Zoom In `simulator-device-config` when Simulator truth is actually required.

## Component Boundaries (starting hypotheses until revised by repository authority)

- OS.js:  presentation/desktop shell — runnable base shell exists at OSJS/ (neutral, no SCADA backend coupling(.
- Orchestrator:  lifecycle/control authority — not yet implemented.
- Modbus Replicator:  acquisition/replication — runtime exists in `replicator/`; its workflow records are queued/archived rather than complete.
- MMA2: deterministic Modbus memory appliance — imported, built, runtime-smoke-tested, and activated for simulator-owned configuration and raw ingest.
- Simulator (new program): OS.js Modbus device simulator using MMA2 as memory/runtime engine — complete through SIM-024, plus None mode and a Windows named-pipe runtime transport; not part of the base four-component identity.
- Windows Electron Toolkit: standalone desktop deployment target in `electron/`, not part of the OS.js desktop; its installer and live COMMS acceptance remain unverified.

## Unresolved / Open

- Orchestrator/Replicator architecture, internal transport, and persistence model remain planning-stage questions.
- MMA2 Modbus TCP port numbers remain architecture-task decisions governed by the network directive.
- The internal runtime transport is committed as a Windows named pipe, but the OS.js relay packages still target a Unix socket; which side changes is unresolved.
- The staged Toolkit migration is mid-flight: only the UMIG-002 placeholder exists, and its build/discovery test has not run.
