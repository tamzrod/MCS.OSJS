# UMIG-004-V — VERIFY: Live Memory Behavior

Status: ACTIVE — UNIT/BUILD predecessor PASS reviewed; isolation BLUEPRINT PREPARED but external target UNVERIFIED; live VERIFY NOT RUN / NOT PASS.
Stage / owner: VERIFY / OpenHands/JR via `operation cwal.md`; environment source and evidence review / ChatGPT
Previous: UMIG-004-T (COMPLETE, `workflow/archive/umig-004-t-memory-adapter.md`)
Next: UMIG-005 (QUEUED; advance only after actual live VERIFY PASS reviewed by ChatGPT)

## Outcome
Independently observe Toolkit Memory's canonical load, explicit apply and selected-device status against real Simulator and MMA2 in a safely disposable environment; verify persisted None/Random semantics, truthful errors and bounded cleanup.

## Prepared prerequisite — not yet satisfied
Human authorized creation of a separate disposable Simulator/MMA2/OS.js test environment. ChatGPT committed `deploy/verify/compose.yaml` and `deploy/verify/README.md` (source-only UMIG-004-E documented in `workflow/archive/umig-004-e-disposable-target.md`). The test-only Compose definition has project-scoped data, no container names/production mounts/host networking, internal MMA2 network, Simulator sharing only that test MMA2's namespace, OS.js on a distinct bridge with a loopback-only UI port. This definition has NOT been executed or validated by Docker on a known isolated host. Operator's running stack and `osjs-data` remain forbidden test targets; running a second Compose project on the operator Docker daemon is NOT isolation.

## Current JR packet and stop condition
Execute ONLY the current `## JR TEST TASK — CURRENT: UMIG-004-V TARGET PREFLIGHT ONLY` in `handoff.md`. On a separately owned disposable Linux VM/sandbox, record real host/daemon identity and test ownership, nonmutating Compose config, resolved ports/network/volumes and absence of existing project resources. No `docker compose up`, dependency install, backend/socket or browser access, configuration writes or cleanup until ChatGPT reviews preflight and publishes a separate executable packet. No demonstrably distinct Docker host, missing daemon, contested project resources or port = BLOCKED, not PASS. Do not run JR merely to repeat a known missing-host blocker.

## Acceptance once a safe target is independently verified
On that exact documented disposable host and only after a new JR packet, record clean browser GUI canonical load; intentional test-only change and manual Save & Apply, runtime acknowledgement and persisted reload; selected-device status and None/Random/error/unavailable handling, screenshots/logs, test config bytes and isolated resource cleanup. Test synthetic port 15020/unit 1 only; do not contact customer devices. No target = BLOCKED.

## Non-scope and dependencies
No production/user data, live volume/socket, operator Docker context, external Modbus endpoint, Replicator work, source fixes or legacy cutover. Prerequisite UMIG-004-T PASS is met; independent host verification and executable test instructions remain outstanding. Never use `docker compose down -v`.

## Sizing
Surface 0, environment 1, behavior 0, verification 1, recovery 1 = 3.
