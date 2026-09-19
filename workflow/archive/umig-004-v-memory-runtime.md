# UMIG-004-V — VERIFY: Live Memory Behavior

Status: COMPLETE — independent live Memory BEHAVIOR ACCEPTED WITH DOCUMENTED TEST-PACKET DEVIATION, reviewed 2026-09-19.
Stage / owner: VERIFY / OpenHands JR; evidence adjudication / ChatGPT
Previous: UMIG-004-T (COMPLETE/PASS)
Next: UMIG-005 (subsequent sole ACTIVE CODE task)

## Evidence and reviewer decision
Immutable JR real-browser/live-backend report: `handoff.md` at `d230f60f231086bdba95b981555ad02bd015e677`; source/test baseline `bf489986ac6b1498ead80cf02226467244000489`. GitHub comparison `bf48998..d230f60` confirmed only the authorized handoff report changed. A freshly initialized OpenHands sandbox-local daemon hosted project-owned test volume, private MMA2 network, loopback-only UI and no production containers/data. Real Chromium observed empty canonical Memory without fixture substitution/automatic write, three explicit Save & Apply operations creating synthetic `VERIFY-SIM-1` on private port 15020/unit 1 and persisting None→FC3 Random 1000ms→None. Runtime classifications, MMA2 restart, ownership file, config hashes/reload and selected RUNNING/IDLE were observed. Intentional Simulator outage showed UNAVAILABLE and no fixture substitution; recovery returned the persisted device and RUNNING/IDLE. Label-checked cleanup removed four test containers and two networks, retaining only the explicitly named test-only volume `mcsverify-1789784184-232_verify-data`; tracked checkout clean. All three stage-specific observable behavioral acceptance outcomes are supported. Does NOT establish production Docker acceptance or missing-file `devices:null` behavior.

## Mandatory qualification — NOT an exact-packet/all-commands PASS
Original step-9 `docker compose --project-name "$PROJECT" -f deploy/verify/compose.yaml start modbus-simulator-runtime` FAILED exit 1 twice: Compose reran the one-shot seed against a populated test volume, which correctly refused reuse. JR then deviated from the exact packet/stop rule and used `docker start mcsverify-1789784184-232-modbus-simulator-runtime-1`, exit 0, on the same project-owned container. The necessary live UI/socket recovery observation was directly obtained; no mock or prior unit test substituted for behavioral evidence. Reviewer accepts stage-specific Memory live behavior, separately classifies the bad restart command as a test harness/runbook defect, and does not claim the failed Compose start or exact packet execution passed. Correction documented in test-only `deploy/verify/README.md` for future explicitly authorized packets; seed safety guard preserved. No additional tests or retest claimed.

## Boundary
UMIG-004-T unit/build PASS separate. Go, MMA2 and production Compose/source, user data/`osjs-data`, legacy OS.js applications, Windows Electron, ICC untouched. Replicator and Diagnostics fixture-only and not verified. UMIG-005 must execute its CODE/TEST/VERIFY gates independently; cutover not approved.

## Sizing
Surface 0, environment 1, behavior 0, verification 1, recovery 1 = 3.
