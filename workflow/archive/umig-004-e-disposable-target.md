# UMIG-004-E — CODE: Prepare Disposable UMIG-004-V Test Stack

Status: COMPLETE — source-only preparation, 2026-09-19; not deployed, not runtime-verified.
Stage / owner: CODE / ChatGPT. Human explicitly approved creating a separate disposable Simulator/MMA2/OS.js environment after UMIG-004-T PASS.
Previous: UMIG-004-T (COMPLETE/PASS)
Next: UMIG-004-V (remains sole ACTIVE; verified external target prerequisite outstanding)

## Single outcome
Add reproducible, test-only isolated Compose definition and safety guide, without touching production deployment.

## Source evidence
- `deploy/verify/compose.yaml`: four separate test services with project-scoped volume; `seed` refuses existing data and writes empty MMA2 config; MMA2 on private internal bridge, Simulator shares ONLY test MMA2's network namespace for localhost readiness/Raw Ingest, OS.js on separate bridge with only loopback UI publishing. No fixed container names, production mounts, host-network use or Modbus host ports.
- `deploy/verify/README.md`: requires a separately owned disposable Linux host/VM with a distinct Docker daemon; mandatory nonmutating owner/daemon/port/volume preflight, explicit unique Compose project, no production access, restricted later synthetic port 15020 test, safe scoped teardown and no blanket volume cleanup.
- Both files read back after creation. `deploy/docker-compose.yml`, existing product and operator data unchanged by this source preparation.

## Boundary / handoff
This is a blueprint committed to GitHub, NOT a created/running VM and NOT proof of a safe execution target. Coding environment has no Docker daemon. UMIG-004-V stays ACTIVE with a preflight-only JR packet until an independent disposable host and exact resolved Compose target are demonstrated, reviewed by ChatGPT and replaced with an executable runtime VERIFY packet. No user-config writes, live service launch, browser tests or PASS claimed.

## Sizing
Surface 1, environment 1, behavior 0, verification 0, recovery 1 = 3.
