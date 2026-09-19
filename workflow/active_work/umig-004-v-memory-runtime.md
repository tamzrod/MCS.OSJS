# UMIG-004-V — VERIFY: Live Memory Behavior

Status: ACTIVE — promoted following independently reviewed UMIG-004-T PASS (`0eba36e61ec30254ddccb19bb8ff88ff403131cf`); safe target MISSING, VERIFY NOT RUN / NOT PASS.
Stage / owner: VERIFY / OpenHands/JR via `operation cwal.md`; safe-target approval and outcome review / ChatGPT
Previous: UMIG-004-T (COMPLETE, `workflow/archive/umig-004-t-memory-adapter.md`)
Next: UMIG-005 (QUEUED; advance only after a real VERIFY PASS reviewed by ChatGPT)

## Outcome
Directly observe Toolkit Memory load/apply and selected-device status against a real Simulator runtime in an explicitly disposable, isolated environment. Confirm persisted None/Random semantics, accurate unavailable/error states and safe cleanup.

## Prerequisite: NOT SATISFIED
Operator Docker installation, `osjs-data` volume and live configuration are NOT test targets. Current `deploy/docker-compose.yml` uses fixed container names and host-networked simulator/MMA2, so a second unmodified Compose instance is NOT presumed isolated. No disposable target with separate data volume/root, isolated runtime/MMA2, collision-free ports, explicit ownership and cleanup has been supplied/verified. No production backup/restore operation is authorized. Never use `docker compose down -v`.

## Current JR boundary
Use only `## JR TEST TASK — CURRENT: UMIG-004-V TARGET PREFLIGHT ONLY` in `handoff.md`; it is a nonmutating prerequisite check. Do not launch the backend, Docker, browser, or perform any configuration write until ChatGPT documents and approves the safe target and publishes a new exact, executable VERIFY packet. Invoking JR now would produce BLOCKED (missing target), NOT a VERIFY PASS.

## Acceptance after safe target verified
On a documented isolated test system, observe real GUI load of canonical document, deliberate safe test edit and explicit Save & Apply, backend acknowledgment and persisted state, selected-device status, None/Random and error/unavailable behavior; collect source SHA, precise test target, pre/post config and safety proof, screenshots, response/log evidence and cleanup. No safe target means BLOCKED.

## Non-scope and dependencies
Never touch production/customer data, operator services, live ports or volumes; no Replicator work, backend fixes or legacy cutover. UNIT/BUILD PASS prerequisite met; verified disposable target and current executable test packet still outstanding.

## Sizing
Surface 0, environment 1, behavior 0, verification 1, recovery 1 = 3.
