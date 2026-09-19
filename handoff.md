# Handoff

## Direction and authorities — 2026-09-19

Human approved staged MCS.OSJS Toolkit migration and disposable OpenHands testing. ChatGPT owns CODE, evidence adjudication and workflow advancement; OpenHands/JR owns independent TEST/VERIFY via `operation cwal.md`, writes ONLY an explicitly authorized `JR TEST REPORT` and stops. Only BLACK SHEEP WALL edits ICC. Operator's working Docker deployment, `osjs-data`, legacy apps, user config, Go/MMA2 runtime and Windows Electron remain untouched; do not use the production Docker daemon for tests or run `docker compose down -v`.

## Current task selection

**SOLE ACTIVE: UMIG-005 — CODE: Toolkit Replicator adapter. No UMIG-005 source implementation or test has been done yet.** Exact task: `workflow/active_work/umig-005-connect-replicator-tab.md`, Previous UMIG-004-V archived COMPLETE with a qualified reviewer acceptance, Next UMIG-005-T QUEUED. UMIG-005-V and later work remain QUEUED under their own gates. UMIG-008/009 launcher cutover/legacy removal remain Planning and require separate human approval. DOCKER-001 queued/paused. No CURRENT JR TEST TASK while CODE is ACTIVE; do not invoke `OPERATION CWAL` based on the historical packet.

## UMIG-004-V independent live VERIFY — reviewer disposition

Immutable JR report: `handoff.md` at `d230f60f231086bdba95b981555ad02bd015e677` (source/test baseline `bf489986ac6b1498ead80cf02226467244000489`). GitHub compared `bf48998..d230f60`: ONLY the authorized `handoff.md` JR report changed, and the report includes a clearly disclosed deviation. Original preflight READY at `37dc61f533c8e7221610c8bf5858ddbe01790312`; earlier incorrect BLOCKED `e5a1be9` retracted. UMIG-004-T independent unit/build PASS `0eba36e61ec30254ddccb19bb8ff88ff403131cf` remains separate; previous UMIG-003-T/V fixture gates PASS and archived.

ChatGPT accepts UMIG-004-V **stage-specific live BEHAVIOR**: real Chromium observed canonical empty load with no fixture leakage/automatic write; synthetic `VERIFY-SIM-1` on test-only 15020/unit 1 persisted after one structural Save & Apply; FC3 None→Random(1000ms)→None persisted in exactly two further applies; MMA2 RUNNING/Simulation IDLE, then UNAVAILABLE on real Simulator stop, then restored canonical device and truthful RUNNING/IDLE after service recovery. Test isolation, hashes, owner registry, private-only port and label-checked teardown observed. Four containers and two networks removed; test volume `mcsverify-1789784184-232_verify-data` intentionally RETAINED. No product code/deployment change and no production acceptance implied. Archival record: `workflow/archive/umig-004-v-memory-runtime.md`.

**Explicit test-command exception — NOT an all-commands PASS:** Original step-9 `docker compose ... start modbus-simulator-runtime` FAILED exit 1 twice because the one-shot seed re-ran on populated verify volume and refused reuse, correctly. JR deviated from the exact packet by using `docker start mcsverify-1789784184-232-modbus-simulator-runtime-1` on the same project-owned container, then directly observed real-browser recovery. Reviewer accepts behavioral gate despite this disclosed procedural deviation; does NOT claim the prescribed Compose restart succeeded or that the JR packet was executed exactly. Classify as verification-runbook/Compose lifecycle issue, not Toolkit product defect. Test-only `deploy/verify/README.md` documents corrected future restart procedure; retain seed's nonempty-volume guard. Original exact command remains historically failed. Future JR packets must preauthorize the corrected command and require ownership validation; do not automatically expand JR scope.

Unverified: missing-file Go `devices:null` behavior (fixture deliberately seeded `devices: []`), Replicator and Diagnostics backend, production Docker health/acceptance, launcher cutover, legacy UI retirement. Previously observed optional icon/sound 404s remain non-gating.

## UMIG-005 CODE boundary and next handoff

Use the existing Replicator v1 Go runtime/Unix protocol, prepared but unimported Toolkit `replicator-contract.js`, and existing Toolkit-owned Memory relay/transport patterns without importing legacy Replicator UI. Determine actual Go response/status shapes and runtime Unix socket path from source; check bridge compatibility and preserve explicit backend ownership conflicts, Pull Block validation, typed errors, and UNKNOWN/UNAVAILABLE COMMS; never fabricate green. Keep Memory and fixture Diagnostics untouched. CODE writes are limited to the authorized Replicator branch. Source diff/readback only, not TEST PASS; after CODE checkpoint publish sole ACTIVE UMIG-005-T and exact JR unit/build packet for OpenHands. No live backend test, deployment or configuration write now.

## Next action and recommendation

Next action (ChatGPT): implement the sole ACTIVE UMIG-005 CODE checkpoint. Recommendation: keep test-only restart instructions corrected, maintain strict production isolation, and ask JR to test the Replicator only after a source checkpoint and exact packet.
