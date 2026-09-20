# Handoff

## Sole ACTIVE and scope — 2026-09-20

`UMIG-EM-002-R` is archived COMPLETE / TEST PASS after ChatGPT reviewed independent JR's original one-time Go race/regression report (`9f9746e0a2a3d99bc09dc27259440b316f297653`) AND the formal Operation CWAL verbatim-original-output accounting (`0290e6b3bb29552c04fedbc6651d99b399a570d0`). The complete original report, its preflight and 128-line original test transcript are permanently readable at that latter commit's `handoff.md`; this new handoff is task continuation, not a deletion of historical evidence. Go test PASS is NOT live VERIFY. Latest approved PRODUCT checkpoint remains `c98eacef8a6af8b0786a09766a4a034e25e82ce6`; subsequent changes through this task activation are workflow/handoff only. The only ACTIVE task is `UMIG-EM-002-V`, independent disposable live VERIFY. `UMIG-EM-003` is PLANNED CODE and is NOT automatically promotable. No product fixes, operator Docker, customer volumes, enabled RBE/access-event output, production, Electron deployment or launcher cutover. ICC is stale; only BLACK SHEEP WALL edits ICC.

## JR TEST TASK — CURRENT: UMIG-EM-002-V disposable upgraded-baseline live VERIFY (Operation CWAL)

GOAL: Independently confirm the upgraded Go/MMA2/Replicator and existing OS.js Toolkit can start and execute ONE end-to-end synthetic Memory -> Replicator -> MMA2/Diagnostics workflow in a NEW, EMPTY, isolated Linux disposable stack. This is not new UI-feature parity, full COMMS certification, production or network-listener authorization. Execute this packet only through `operation cwal.md`; previous archived packets are background, not authority. JR must STOP after its report; ChatGPT reviews results and owns advancement.

TARGET / SAFE BOUNDARY: JR's own clean disposable repository checkout AND its own provably local disposable Docker daemon only, never the operator's Legion, a remote/production daemon, existing service, live customer Modbus endpoint, existing `osjs-data` mount or any previously retained verification volume. Read `AGENTS.md`, `operation cwal.md`, this handoff, `workflow/active_work/umig-em-002-v-upgraded-baseline.md`, `deploy/verify/README.md` and `deploy/verify/compose.yaml`; read-only inspection of those files only. Execute commands from repository root unless shown otherwise; record commands and exit codes. On essential unknown provenance BLOCKED/STOP; do not invent a fallback.

A. REPOSITORY PREFLIGHT, in order, before ANY Docker action:
```sh
git fetch origin main
git status --porcelain
git rev-parse HEAD
git rev-parse origin/main
git merge-base --is-ancestor 0290e6b3bb29552c04fedbc6651d99b399a570d0 origin/main
git merge-base --is-ancestor c98eacef8a6af8b0786a09766a4a034e25e82ce6 origin/main
git diff --name-only c98eacef8a6af8b0786a09766a4a034e25e82ce6 origin/main
git merge-base --is-ancestor HEAD origin/main
```
All commands must exit 0 and status be empty. Net diff must contain EXACTLY these five paths and NO product/source/config/Compose paths: `handoff.md`, `workflow/active_work/umig-em-002-b-toolkit-build.md` (removed), `workflow/archive/umig-em-002-b-toolkit-build.md`, `workflow/archive/umig-em-002-r-current-go-regression.md`, `workflow/active_work/umig-em-002-v-upgraded-baseline.md`. Only if ancestry fails SOLELY due to a confirmed shallow repository (`git rev-parse --is-shallow-repository` -> true), one `git fetch --unshallow origin` and recheck only failed ancestry commands is permitted. If clean, source diff exact, HEAD ancestor and HEAD differs, do ONE `git merge --ff-only origin/main`; then repeat HEAD, origin/main and status. Require equal SHAs and empty tracked status. No reset/clean/rebase/cherry-pick/force. Confirm ONE ACTIVE file marked `UMIG-EM-002-V`; predecessor archived PASS; no other ACTIVE. If mismatch, BLOCKED/STOP. Do not edit source to make any check pass.

B. SAFE DAEMON / PROJECT PREFLIGHT, in order, read-only until all gates pass:
```sh
hostname
id
docker context show
printenv DOCKER_HOST
sudo docker info --format '{{.ID}} {{.Name}} {{.DockerRootDir}}'
sudo docker compose version
sudo docker ps -a
sudo docker volume ls
sudo docker network ls
ss -ltn
```
`printenv DOCKER_HOST` exit 1 when unset is acceptable; `ss` missing permits read-only `cat /proc/net/tcp /proc/net/tcp6` instead. Capture actual identity and inspect context, daemon root and resources to PROVE this is JR's disposable local daemon and NOT a production/remote/operator daemon; previously retained test volumes may exist but must remain untouched. If dockerd idle, start it only when provenance establishes sandbox-local daemon AND `deploy/verify/README.md` permits safe local startup; otherwise BLOCKED (no operator fallback). Require port 127.0.0.1:18219 free; if occupied BLOCKED, do not probe another port or stop the owner. No host-published Modbus port or production service names in the proposed verification target. Unknown identity/permissions => BLOCKED before creation.

Choose ONCE and record a fresh collision-resistant project identifier, in the SAME shell session used for Compose commands:
```sh
PROJECT="mcsverify-em002v-$(date -u +%Y%m%d%H%M%S)-$$"
export PROJECT MCS_VERIFY_OSJS_PORT=18219
printf 'PROJECT=%s\n' "$PROJECT"
sudo docker compose --project-name "$PROJECT" -f deploy/verify/compose.yaml config
sudo docker ps -a --filter "label=com.docker.compose.project=$PROJECT"
sudo docker volume ls --filter "label=com.docker.compose.project=$PROJECT"
sudo docker network ls --filter "label=com.docker.compose.project=$PROJECT"
```
All three project-resource inventories MUST be EMPTY before startup. Record rendered Compose; require exactly seed/mma2/modbus-simulator-runtime/modbus-replicator-runtime/osjs-shell, new named `verify-data`, seed nonempty-volume refusal and empty initial YAML, MMA2 internal `verify-runtime` network, BOTH Go runtimes `network_mode: service:mma2`, no operator bind mounts/external volumes/privileged mode/fixed names/host network, and OS.js ONLY published `127.0.0.1:18219:18209` on `verify-ui`. No host Modbus, RBE or access-event port; if discrepancy or resource collision, BLOCKED/STOP. NEVER retry with an alternative project token merely to evade collision. No production Compose file.

C. SINGLE STARTUP and read-only readiness (ONLY after A+B PASS):
```sh
sudo docker compose --project-name "$PROJECT" -f deploy/verify/compose.yaml up -d --build
sudo docker compose --project-name "$PROJECT" -f deploy/verify/compose.yaml ps -a
sudo docker compose --project-name "$PROJECT" -f deploy/verify/compose.yaml logs --tail=120 seed mma2 modbus-simulator-runtime modbus-replicator-runtime osjs-shell
curl -fsS http://127.0.0.1:18219/healthz
sudo docker compose --project-name "$PROJECT" -f deploy/verify/compose.yaml exec -T osjs-shell sh -c 'test -S /data/run/modbus-simulator.sock && test -S /data/run/modbus-replicator.sock'
sudo docker compose --project-name "$PROJECT" -f deploy/verify/compose.yaml exec -T osjs-shell sh -c 'cat /data/config/simulator/devices.yaml /data/config/replicator/devices.yaml /data/config/mma2/config.yaml; sha256sum /data/config/simulator/devices.yaml /data/config/replicator/devices.yaml /data/config/mma2/config.yaml'
```
ONE `up -d --build` only. Wait at most 180s for readiness without repeating `up`/`start` (seed refuses a populated volume). Expect seed Exited(0), other FOUR services Up, healthz HTTP 200 / status ok, two real Unix sockets, initial Simulator/Replicator `devices: []` and MMA2 `listeners: []`, and record the three seed-file SHA256 fingerprints. Confirm from rendered config + `ps` there is no published host Modbus/RBE/access-event port. Fail actual startup/health/sockets/seed contradiction => FAIL; missing essential safe environment => BLOCKED. Capture outputs/exits, not inferred green states.

D. ONE synthetic real-browser compatibility workflow; use actual Chromium and new profile pointed ONLY at `http://127.0.0.1:18219`. Never directly edit container YAML, mock browser state, substitute unit-test output, click an apply twice because delayed, stop services, or access other endpoints. Document browser UI states/screenshots or direct DOM observations and browser console failures. If input rerenders, reacquire element state; maximum two corrections to an unsaved field, no repeated apply. STOP product actions on first genuine product contradiction.

D1. Open OS.js Start > Development > MCS Modbus Toolkit ONCE, confirm one Toolkit window with Memory / Replicator / Diagnostics. Empty canonical lists, no fixture devices; Diagnostics global service health UNKNOWN and unsupported native paths UNAVAILABLE/controls disabled, NOT fabricated RUNNING. Opening/refreshing UI must not mutate initial three seed hashes. Record real browser observations.

D2. Memory GUI only: create ONE test device named `VERIFY-EM002V-SIM`, port 15020, unit 1, FC1/FC2/FC3/FC4 each start 0, count 16, simulation None and interval 0. Memory Save & Apply ONCE (Memory apply #1). Within <=60s verify canonical Simulator YAML exactly one synthetic test device, private MMA2 listener and owner `(15020,1)` simulator, Memory device MMA2 RUNNING and Simulation IDLE. Confirm no host port 15020 published. Read config via test `osjs-shell` `cat`/`sha256sum`; capture test Compose logs and UI. Any ownership of a foreign device or external endpoint => STOP/FAIL, no workaround.

D3. Replicator GUI only: create ONE device `VERIFY-EM002V-REP`, source `127.0.0.1:15020`, unit 1; ONE FC3 Pull Block start 0/count 16; scan interval 1000ms; turn destination Auto Port and Auto Unit ID OFF and use destination port 15021/unit 1. Check availability once, then Replicator Save & Apply ONCE (Replicator apply #1). Within <=90s require canonical Replicator YAML one test device and private MMA2 `owners.yaml`/effective config preserve the Simulator `(15020,1)` reservation and add Replicator `(15021,1)` without foreign-owner mutation. Directly observe real Replicator runtime RUNNING, source OK, Pull Block 1 OK with nonempty actual `last_poll` and matching source/block identity; do not infer success from `source_status` alone. Backend v1 status or genuine rendered UI may substantiate detailed per-block fields, but report the actual source/surface. No host 15021. Capture actual config and logs.

D4. Diagnostics actual browser: select/refresh ONCE, wait <=30s; confirm test Memory RUNNING/IDLE, test Replicator RUNNING/source OK/Block 1 OK with real poll timestamp, and globals MMA2/Simulator/Replicator still UNKNOWN unless real supported service probes now prove otherwise (no fabricated green); native paths UNAVAILABLE, Start/Stop disabled, observation pane labelled observations not service logs. Before and after this read-only Diagnostics refresh, record `sha256sum` of the three canonical YAMLs AND `/data/config/mma2/owners.yaml` from test `osjs-shell`; require identical bytes and no MMA2 restart request caused by refresh, verified from test service logs. Record UI/console and direct status evidence; if false health, config change or unsupported control enabled => FAIL.

D5. Restore test state by GUI only, at most Replicator Save & Apply ONCE (#2) after deleting ONLY `VERIFY-EM002V-REP`, then Memory Save & Apply ONCE (#2) after deleting ONLY `VERIFY-EM002V-SIM`. No new device or reconfiguration. Confirm Replicator/Simulator YAML `devices: []`, MMA2 `listeners: []`, any ownership reservations empty and original THREE seeded SHA256 values restored EXACTLY; no RBE/access-event listener configured or published, Diagnostics empty canonical lists, tracked `git status --porcelain` empty. TOTAL at most FOUR explicit Save & Apply clicks (Memory two; Replicator two). If any earlier product FAIL, do NOT attempt unscripted GUI restoration; preserve test data and proceed only to safely scoped cleanup.

For EVERY relevant canonical evidence read after D2/D3/D4/D5, use this read-only test-container command; log expected vs observed values without editing them:
```sh
sudo docker compose --project-name "$PROJECT" -f deploy/verify/compose.yaml exec -T osjs-shell sh -c 'cat /data/config/simulator/devices.yaml /data/config/replicator/devices.yaml /data/config/mma2/config.yaml; if test -f /data/config/mma2/owners.yaml; then cat /data/config/mma2/owners.yaml; fi; sha256sum /data/config/simulator/devices.yaml /data/config/replicator/devices.yaml /data/config/mma2/config.yaml; if test -f /data/config/mma2/owners.yaml; then sha256sum /data/config/mma2/owners.yaml; fi'
```
Check actual owner/effective listener contents, not hashes alone. Request log counters by read-only `sudo docker compose --project-name "$PROJECT" -f deploy/verify/compose.yaml logs --tail=120 mma2 modbus-simulator-runtime modbus-replicator-runtime` at D3/D4 and after D5; do not assert an unsupported restart ACK. Any unobservable required per-block telemetry => BLOCKED (not PASS); contradictory product output => FAIL.

E. OWNED-RESOURCE CLEANUP (mandatory after resources exist, even FAIL/BLOCKED, only if ownership proven):
```sh
sudo docker ps -a --filter "label=com.docker.compose.project=$PROJECT"
sudo docker volume ls --filter "label=com.docker.compose.project=$PROJECT"
sudo docker network ls --filter "label=com.docker.compose.project=$PROJECT"
```
Individually verify every target container, network AND volume's exact `com.docker.compose.project=$PROJECT` ownership with `sudo docker inspect` / `sudo docker volume inspect` / `sudo docker network inspect` on ONLY the concrete IDs/names returned above; confirm Compose down's deletion set contains no other project. If ownership cannot be proven STOP cleanup and report BLOCKED with inventory; NEVER remove/modify an ambiguous resource. Otherwise run ONCE:
```sh
sudo docker compose --project-name "$PROJECT" -f deploy/verify/compose.yaml down --remove-orphans
sudo docker ps -a --filter "label=com.docker.compose.project=$PROJECT"
sudo docker network ls --filter "label=com.docker.compose.project=$PROJECT"
sudo docker volume ls --filter "label=com.docker.compose.project=$PROJECT"
git status --porcelain
```
Require only this project's containers and networks absent, newly named project `verify-data` volume retained with exact project label, other resources untouched and tracked tree clean. NEVER use `-v`, `prune`, wildcard cleanup, volume rm, old-volume cleanup, a second Compose up/start, production/host fallback or extra test. If failure occurs before creation, report inventory and no teardown. Retain any test volume even on FAIL.

VERDICT / REPORT: PASS only if all A+B gates, one successful C startup, D1-D5 real synthetic runtime/UI/ownership/telemetry/empty-restoration checks and E scoped teardown were directly evidenced. Actual product contradiction => FAIL. Unsafe environment/ownership, impossible required observation, insufficient browser, or ambiguous cleanup => BLOCKED; distinguish any concurrent product FAIL in the report. Record complete preflight SHAs/net paths/daemon identity/project token/rendered config and exits, test port/volume evidence, `up` result/health/sockets/logs, initial/final YAML and hashes, four apply counts, actual per-block telemetry/timestamp/source, browser evidence/console, absence of output exposure, teardown resource labels/retained volume, warnings and side effects. Do not extrapolate to advanced UI, true four-layer COMMS, external network listeners, production or final visual parity. DO NOT FIX SOURCE OR CONFIG. No retries except bounded health waiting and unsaved-field correction specified above.

REPORT-WRITE AUTHORITY: JR may replace ONLY `## JR TEST REPORT — UMIG-EM-002-V` below. Preserve every other handoff section and all project files. Before report transport re-fetch origin/main; if it advanced, do NOT merge/rebase/force to deliver report: report BLOCKED transport race in chat and retain evidence. If unchanged, commit/push ONLY `handoff.md`, verify HEAD=origin/main and tracked tree clean, then STOP. NO task advancement, ICC, workflow, product, Compose or general CWAL edits.

## JR TEST REPORT — UMIG-EM-002-V

PENDING independent Operation CWAL execution. No current-upgraded-baseline live VERIFY PASS or production claim.
