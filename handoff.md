# Handoff

## Human direction and authority — 2026-09-19

Human approved creating a separate disposable Simulator/MMA2/OS.js verification environment after UMIG-004-T independent PASS. ChatGPT owns environment preparation/CODE, independent JR evidence review, packet publication and workflow advancement; OpenHands/JR owns TEST/VERIFY under `operation cwal.md` and may edit only the current authorized report and push only `handoff.md`. Only BLACK SHEEP WALL edits ICC. Preserve operator's working Docker deployment, production/customer configs and `osjs-data`, legacy apps, and separate Windows Electron. NEVER invoke `docker compose down -v` or target the operator Docker daemon for this test.

## Current selection and prior gates

**SOLE ACTIVE: UMIG-004-V — real Memory VERIFY, isolation preflight pending; live VERIFY NOT RUN/NOT PASS.** UMIG-004-T is COMPLETE: JR UNIT/BUILD PASS evidence at immutable `0eba36e61ec30254ddccb19bb8ff88ff403131cf` (`handoff.md`); archived `workflow/archive/umig-004-t-memory-adapter.md`. UMIG-004 CODE source checkpoint `508c6b031e675c696de3ff7bdb4291dd005d54ae` is archived; UMIG-003-T and UMIG-003-V fixture/rendered passes are archived. UMIG-005 and later remain QUEUED; DOCKER-001 queued/paused; launcher cutover/legacy removal remain Planning. Unit/build PASS does not prove GUI, actual backend, safe writes, Docker health or deployed behavior.

## Prepared target blueprint — NOT a verified running environment

Human-authorized test-only source files: `deploy/verify/compose.yaml` and `deploy/verify/README.md`; provenance `workflow/archive/umig-004-e-disposable-target.md`. Read README before preflight. This separate Compose file contains `seed` (empty volume only; initializes test MMA2 `listeners: []`), independently supervised test MMA2, test Simulator in only that MMA2's network namespace for local readiness/raw ingest, and test OS.js on a separate bridge. MMA2 network is internal, no Modbus port published, UI only `127.0.0.1:${MCS_VERIFY_OSJS_PORT:-18219}:18209`. Volume is project-scoped, no explicit container names, external volumes, host network or production mounts. Never combine this file with `deploy/docker-compose.yml`. Coding environment lacks Docker; these definitions were inspected but have NOT been executed or Docker-validated. The operator's live host is NOT an approved target merely because a second project name can be chosen.

**Target prerequisite remains UNSATISFIED:** OpenHands must first demonstrate a separate, disposable Linux host/VM with its OWN Docker daemon (not the operator's), browser availability, unambiguous resource ownership, a free loopback UI port, brand-new project/volume and safe scoped cleanup. ChatGPT reviews the recorded target and MUST publish a different live-test packet before JR starts any container. If any target proof is missing, report BLOCKED and stop. No backup/restore of production is authorized.

## JR TEST TASK — CURRENT: UMIG-004-V TARGET PREFLIGHT ONLY

GOAL: Independently verify a reproducible, distinct, disposable target and this test-only Compose configuration WITHOUT starting containers, connecting to a runtime/browser, or writing configuration. This preflight can report READY/BLOCKED; it CANNOT establish UMIG-004-V PASS.

SAFE PRECONDITIONS: Work only in a fresh tracked-clean disposable checkout of `tamzrod/MCS.OSJS` at current `origin/main` on a separately owned disposable Linux VM/sandbox. Record full HEAD, `git status --porcelain`, `hostname`, host/VM ownership/provenance and proof that its Docker daemon is NOT the operator/production daemon. Read `deploy/verify/README.md`, this packet, `workflow/active_work/umig-004-v-memory-runtime.md`, archived UMIG-004-T and UMIG-004-E. Verify `git merge-base --is-ancestor 0eba36e61ec30254ddccb19bb8ff88ff403131cf HEAD` exits 0, UMIG-004-V SOLE ACTIVE and UMIG-005 QUEUED. Unknown host ownership, missing Docker/Compose, an operator daemon, dirty tracked checkout or conflicting task state -> BLOCKED immediately without daemon-mutating actions. Do not install dependencies, pull images, copy configs, reset/clean, or touch ICC.

EXACT NONMUTATING ACTIONS (in order; capture each exit/output; stop on uncertainty):
1. `hostname; uname -a; docker context show; docker info --format '{{.ID}} {{.Name}} {{.DockerRootDir}}'; docker compose version; docker version` on the distinct disposable host. Record daemon endpoint/context and host ownership/provenance, compare to any available operator identity; do NOT treat an arbitrary context name or absent production container as sufficient isolation proof.
2. `docker ps -a --format '{{.Names}}'; docker volume ls --format '{{.Name}}'; docker network ls --format '{{.Name}}'`. Confirm the four production names (`mcs-osjs-shell`, `mcs-mma2`, `mcs-modbus-simulator-runtime`, `mcs-modbus-replicator-runtime`) and production volume `osjs-data` are absent. If an existing MCS service, unknown target, or any risk of shared daemon is found, BLOCKED. Record loopback TCP port 18219 availability using `ss -ltn` (or equivalent read-only listener listing); if busy, BLOCKED, do not kill or remap an existing process.
3. Set a NEW unique identifier in the disposable shell, for example `PROJECT="mcsverify-$(date +%s)-$$"`; record it. Use only `docker compose --project-name "$PROJECT" -f deploy/verify/compose.yaml config` (with `MCS_VERIFY_OSJS_PORT=18219` if an explicit env is needed); capture exit 0 and resolved YAML. Confirm exactly seed/mma2/modbus-simulator-runtime/osjs-shell; correct `../../MMA2`, repository-root simulator and `../../OSJS` build contexts; seed empty-volume refusal; scoped `verify-data` (no external/name override); no `container_name`, host network, bind/production mounts or published Modbus ports; Simulator `network_mode: service:mma2`; MMA2 internal-only network; OS.js on separate bridge and loopback-only 18219->18209 mapping. Any mismatch = BLOCKED, not permission to experiment.
4. Read-only check `docker compose --project-name "$PROJECT" -f deploy/verify/compose.yaml ps -a`, `docker ps -a --filter "label=com.docker.compose.project=$PROJECT"`, `docker volume ls --filter "label=com.docker.compose.project=$PROJECT"`, and equivalent network listing filtered by project label. Expect zero pre-existing resources; no `up`, `build`, `pull`, `down`, volume removal, service/socket request or browser operation. Record clean `git status --porcelain`.

EXPECTED / EVIDENCE: If distinct independently owned Docker host plus all nonmutating checks and resolved config prove isolation, report **TARGET PREFLIGHT READY; LIVE VERIFY NOT RUN** with exact host/context/daemon proof, project token, port, resulting Compose config, zero existing resources, HEAD, ancestor, clean status and any caveats. If host lacks Docker or independence/resource invariants cannot be proved, report **BLOCKED** with reason. A successful `docker compose config` alone is insufficient to declare target safe. Neither outcome may be called UMIG-004-V PASS. No production access, images pulled, container start or data write.

REPORT-WRITE AUTHORITY: Replace ONLY `## JR TEST REPORT — UMIG-004-V` with factual preflight READY/BLOCKED and evidence, commit/push ONLY `handoff.md`, then STOP. JR may not edit product, Compose, workflow, ICC or other handoff sections. ChatGPT reviews report and alone authorizes a different exact executable live VERIFY packet if a safe target is demonstrated.

## JR TEST REPORT — UMIG-004-V

Verdict: **BLOCKED** — TARGET PREFLIGHT could not establish a separately owned, disposable Docker host. No independent Docker daemon is present in this execution environment, so preflight step 1 (`docker info` server identity) and step 2 (daemon-scoped container/volume/network listing) cannot be executed and no host/daemon isolation can be proved. Per packet, "missing Docker" is an immediate BLOCKED condition without daemon-mutating actions. Live browser/backend VERIFY: **NOT RUN / NOT PASS**. A successful `docker compose config` is recorded below but is explicitly insufficient to declare the target safe.

### Checkout + task state (verified)
- HEAD: `39c2a3a5f6e4863360b161337c41a82158f98495` = `origin/main` (`tamzrod/MCS.OSJS`, branch `main`).
- `git status --porcelain` empty at start and after all actions → tracked-clean. No tracked, product, Compose, workflow or ICC file was modified.
- Precondition note: the initial workspace was a **1-commit grafted shallow clone** (`.git/shallow` = `39c2a3a…`, `git rev-list --count HEAD` = 1), so the mandated ancestor check could not run as-is. Resolved by the read-only transport action `git fetch --unshallow origin` (exit 0), which added no working-tree changes. No repository file was rewritten, reset or cleaned.
- `git merge-base --is-ancestor 0eba36e61ec30254ddccb19bb8ff88ff403131cf HEAD` → **exit 0** (target commit is `UMIG-004-T: JR independent Memory adapter/relay contract TEST report — PASS`, 2026-09-19 01:19:09 +0000).
- Task state: `workflow/active_work/umig-004-v-memory-runtime.md` = ACTIVE; `workflow/active_work/umig-005-connect-replicator-tab.md` = QUEUED (Previous = UMIG-004-V). UMIG-004-V is SOLE ACTIVE and UMIG-005 is QUEUED, as required.
- Read as required: `deploy/verify/README.md`, this packet, `workflow/active_work/umig-004-v-memory-runtime.md`, `workflow/archive/umig-004-t-memory-adapter.md` (state), archived UMIG-004-E blueprint reference.

### Step 1 — host / daemon identity: FAILED (no daemon)
- `hostname` → `runtime-ylrnihathakoylzg-68df6b6fb9-6cgp6`
- `uname -a` → `Linux … 6.8.0-1055-gke #61-Ubuntu SMP … x86_64 GNU/Linux` (GKE-hosted OpenHands runtime sandbox)
- `docker context show` → `default`; `docker context ls` → only `default *` (`unix:///var/run/docker.sock`). No separate/disposable context exists.
- `docker info --format '{{.ID}} {{.Name}} {{.DockerRootDir}}'` → **exit 1**: `failed to connect to the docker API at unix:///var/run/docker.sock … no such file or directory`. No server identity obtainable.
- `docker compose version` → `Docker Compose version v5.5.1` (exit 0). `docker version` → client 29.8.0 (API 1.56) but server section **exit 1**, same socket error.
- Socket absence confirmed: `/var/run/docker.sock` and `/run/docker.sock` do not exist. `DOCKER_HOST` unset. Identity `uid=10001(openhands)`.
- Consequence: **no Docker daemon is reachable**, therefore host/daemon ownership and independence cannot be recorded, and the operator-daemon comparison the packet demands cannot be performed. This environment is the shared OpenHands runtime sandbox, not a demonstrably separately owned disposable VM with its own daemon. Ownership/provenance of this host is not established by evidence available in-session → BLOCKED.

### Step 2 — daemon resource listing: NOT EXECUTABLE
- `docker ps -a`, `docker volume ls`, `docker network ls` cannot run without a daemon (would repeat the same exit-1 socket error). Production-name absence (`mcs-osjs-shell`, `mcs-mma2`, `mcs-modbus-simulator-runtime`, `mcs-modbus-replicator-runtime`) and volume `osjs-data` absence are therefore **UNVERIFIED**, not confirmed.
- Loopback port availability (read-only): `ss -ltn` filter for `:18219` → **no listener** (port free). This is a host-level observation only and does not substitute for daemon isolation proof.
- No `up`/`build`/`pull`/`down`/volume removal, no service/socket request, no browser operation, no daemon-mutating action was performed.

### Step 3 — nonmutating resolved Compose config: PASS (definition-only; not isolation proof)
- Project token used: `mcsverify-1789782862-232`, invoked as `docker compose --project-name "$PROJECT" -f deploy/verify/compose.yaml config` with `MCS_VERIFY_OSJS_PORT=18219`. **Exit 0.**
- Resolved services exactly `seed`, `mma2`, `modbus-simulator-runtime`, `osjs-shell`.
- Build contexts correct: `mma2` → `/workspace/project/MCS.OSJS/MMA2` (`Dockerfile.supervised`); `modbus-simulator-runtime` → repo root `/workspace/project/MCS.OSJS` (`simulator/Dockerfile`); `osjs-shell` → `/workspace/project/MCS.OSJS/OSJS` (`Dockerfile`).
- Invariants confirmed in resolved output: no `container_name`; `seed` uses `network_mode: none` with empty-volume refusal guard; `verify-runtime` network `internal: true`; `verify-ui` separate bridge; `modbus-simulator-runtime` uses `network_mode: service:mma2`; no host-network mode; no published Modbus ports (5020/15020 absent); OS.js port published only `127.0.0.1:18219 -> 18209`; single project-scoped volume `mcsverify-1789782862-232_verify-data` (no external/name override); networks project-scoped `…_verify-runtime` / `…_verify-ui`.
- This proves only that the test-only Compose definition resolves as authored. Per packet and `deploy/verify/README.md`, it is **insufficient** to declare the target safe.

### Step 4 — project resource existence check: NOT EXECUTABLE
- `docker compose … ps -a` and label-filtered `docker ps/volume/network ls --filter com.docker.compose.project=$PROJECT` require a daemon and could not run. Zero-preexisting-resource expectation is **UNVERIFIED**.
- `git status --porcelain` after all actions → empty (clean).

### BLOCKER (exact)
1. No Docker daemon/engine reachable in this environment (`/var/run/docker.sock` absent; `docker info`/`docker version` server exit 1).
2. No demonstrably separate, OpenHands-owned disposable Linux host/VM with its own daemon — only the shared OpenHands runtime sandbox (`runtime-ylrnihathakoylzg-…`, GKE kernel), whose ownership/independence from the operator environment cannot be proved from in-session evidence.
3. Therefore daemon identity, host independence, production-resource absence and zero-project-resource checks cannot be established, and no daemon isolation can be proved.

### Not performed (per authorization)
No containers started, images pulled/built, configs written, dependencies installed, browser/backend/socket contacted, volumes removed, ICC touched, or any reset/clean. Only authorized report section edited. Neither outcome claims UMIG-004-V PASS.

## Next action and recommendation

Next action: run only the nonmutating target preflight on an independently disposable, Docker-capable test VM; if OpenHands has no such VM/daemon, report that blocker without interacting with production. Recommendation: preserve operator Docker and data; await actual target evidence before starting the test stack or issuing any Save & Apply.
