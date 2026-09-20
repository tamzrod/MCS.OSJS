# Handoff

## Current authority — 2026-09-20

Human-approved sequence: finish `UMIG-EM-003-V` independent live VERIFY with OpenHands, then prepare `UMIG-EM-004` OpenCode coding trial after accepted VERIFY and separate human promotion. Previous Go race TEST `UMIG-EM-003-R-T` is archived PASS; that is not live VERIFY. SOLE ACTIVE: `workflow/active_work/umig-em-003-v-shared-lock.md`. `UMIG-EM-004` is PLANNED; no coding authority. Only BLACK SHEEP WALL updates ICC.

**SOLE OPENHANDS ENTRY POINT: `OPERATION CWAL`.** OpenHands must not be instructed to execute a preparation document or any other workflow directly. Operation CWAL reads this CURRENT JR packet and follows the exact subordinate commands named below. No second directive, new sandbox or new clone is required. This packet authorizes preparation and report delivery only, NOT the live concurrency test.

## JR TEST TASK — CURRENT: UMIG-EM-003-V — CWAL PREPARATION GATE ONLY

GOAL: in the existing disposable OpenHands environment, unblock the missing local Docker daemon once, confirm the verify-only topology and collect the remaining evidence needed for the subsequent exact live VERIFY. This is a PREPARATION stage of the sole ACTIVE VERIFY task, not a product test or a VERIFY PASS.

TARGET / SOURCE: reuse the existing clean `/tmp/em003v-discovery-clone` if present, otherwise the already existing OpenHands checkout. Product source checkpoint `538324a472eb15ff8ef97ee66f826fa02ba9462f` must be an ancestor of the final preparation HEAD. Require a clean checkout and `HEAD == origin/main == git ls-remote origin refs/heads/main` after any strictly authorized fast-forward. The current activation revision is the live GitHub main revision proven by this equality at run time; record its full SHA. No operator/Legion Docker, remote daemon, other deployments or old test volumes.

EXACT COMMANDS / ACTIONS, IN ORDER: Under `OPERATION CWAL`, read `workflow/cwal/umig-em-003-v-sandbox-preflight.md` and execute its numbered steps **1 through 6 exactly once**. That file is incorporated into this packet as the exact command list, NOT a separately invocable directive. It explicitly authorizes ONE clean local-main fast-forward if stale and ONE bounded sandbox-local `sudo -n dockerd` startup if the daemon is absent; it does not permit other Git repairs or privileged operations. No application container start/build, synthetic applies, product tests, production Compose, installation, source edits or live concurrency. If an unsafe prerequisite fails, stop actions and preserve actual BLOCKED evidence; perform only safe post-checks from step 6 where possible.

EXPECTED: source clean/current; sandbox-local daemon accessible with identity; verify-only rendered Compose has no host-network or operator mounts and no public Modbus bind; inventory shows no collisions; local runtime apply/restart interfaces and unknowns reported; final checkout unchanged by preparation (except permitted clean fast-forward). Missing environment or ambiguous provenance = BLOCKED, not product FAIL. Full completion of preparation = `PREPARATION COMPLETE`, **never** `UMIG-EM-003-V PASS`. Interrupted or undelivered mandatory report = `INCOMPLETE` with observed underlying result preserved.

EVIDENCE: command strings, original stdout/stderr or appropriately secret-redacted durable transcript, exit codes, initial/final source SHA/status and live remote, daemon startup result and provenance, rendered verify topology, port/resource inventory, interface file paths and schemas, any side effects. Do not print token-bearing remote URLs. No need to rerun earlier completed product tests or create another sandbox.

REPORT / DELIVERY EXCEPTION: If the checkout is fresh, clean and source-verified after safe post-check, replace ONLY the section `## JR PREPARATION REPORT — UMIG-EM-003-V` below with actual evidence and verdict (including a reportable BLOCKED after a successful freshness check). Preserve all other handoff bytes. Check `git diff --check` and require `git diff --name-only` to equal ONLY `handoff.md`; `git add -- handoff.md`; require `git diff --cached --name-only` to equal ONLY `handoff.md`; `git commit -m 'JR: EM-003-V sandbox preparation report' -- handoff.md`; verify the new commit changes only `handoff.md`. Require live `git ls-remote --exit-code origin refs/heads/main` still equals the tested preparation HEAD immediately before ONE non-force `git push origin HEAD:refs/heads/main`. Re-query live remote and require it equals the new report commit SHA; check clean worktree. No force, no retry, no unrelated commits. If stale/dirty prevents safe report, return evidence IN CHAT and stop; if transport fails, preserve verdict and state INCOMPLETE. Return report commit SHA and STOP. JR cannot archive/promote tasks or authorize live execution.

NEXT (not authorized now): coding agent reviews the pushed preparation evidence and authors a distinct complete live VERIFY packet with exact synthetic requests, bounds, hashes, acknowledgement and label-scoped cleanup. No automatic advance to EM-004.

## JR PREPARATION REPORT — UMIG-EM-003-V

VERDICT: PREPARATION COMPLETE

Executed by independent OpenHands JR under `OPERATION CWAL` (preparation gate only; NOT live VERIFY, NOT a product PASS). Preflight steps 1–6 of `workflow/cwal/umig-em-003-v-sandbox-preflight.md` were run exactly once in the existing disposable checkout `/tmp/em003v-discovery-clone`. No application container was started, no synthetic apply was sent, no source edited, no install performed.

### Step 1 — source freshness (one authorized fast-forward)

| Command | Exit | Result |
|---|---|---|
| `pwd -P` | 0 | `/tmp/em003v-discovery-clone` |
| `git status --porcelain --untracked-files=all` (before) | 0 | empty (clean) |
| `git branch --show-current` | 0 | `main` |
| `git rev-parse HEAD` (before) | 0 | `f58033e02db80b8af50cf05558d1415314986ebf` (stale) |
| `git rev-parse origin/main` | 0 | `4b840d9d4a481b09f210dd3e279e7b2da636b2ae` |
| `git ls-remote --exit-code origin refs/heads/main` | 0 | `4b840d9d…` |
| `git merge-base --is-ancestor HEAD origin/main` | 0 | pass (fast-forward safe) |

Authorized `git merge --ff-only origin/main` -> exit 0 (`f58033e..4b840d9`, 3 files). Post-check: `git status --porcelain --untracked-files=all` empty; `HEAD` = `origin/main` = live `git ls-remote` = `4b840d9d4a481b09f210dd3e279e7b2da636b2ae`; `git merge-base --is-ancestor 538324a472eb15ff8ef97ee66f826fa02ba9462f HEAD` exit 0. Sole ACTIVE task confirmed `UMIG-EM-003-V`; current CWAL packet reread after fast-forward. No reset/rebase/clean/force/second fetch.

### Step 2 — sandbox / daemon provenance

- `uname -s` -> `Linux`; `docker context show` -> `default`; `docker context ls` -> only `default` -> `unix:///var/run/docker.sock`. `env | grep -E '^(DOCKER_HOST|DOCKER_CONTEXT|COMPOSE_PROJECT_NAME|MCS_VERIFY_OSJS_PORT)='` -> empty (exit 1); no remote/operator host indicated.
- Client present (Docker Engine 29.8.0, API 1.56), but daemon was initially absent (`docker version`/`docker info` exit 1; no socket, no `dockerd`/`containerd`, `/tmp/mcs-em003v-dockerd.pid` absent) — matching the prior known state.
- Exactly ONE bounded sandbox-local startup was run, verbatim: guard `test ! -S /var/run/docker.sock && ! pgrep -x dockerd && test ! -e /tmp/mcs-em003v-dockerd.pid && sudo -n sh -c 'nohup dockerd > /tmp/mcs-em003v-dockerd.log 2>&1 < /dev/null & echo $! > /tmp/mcs-em003v-dockerd.pid'` -> exit 0, PID `717`. Then `timeout 90s sh -c 'until sudo -n docker version >/dev/null 2>&1; do sleep 2; done'` -> exit 0.
- `sudo -n docker version` -> Server Docker Engine 29.8.0 (containerd v2.3.5, runc 1.5.1). `sudo -n docker info --format '{{.Name}} {{.DockerRootDir}}'` -> `runtime-dfpoqximlibnvfbn-89d77c974-j29f4 /var/lib/docker` (daemon name equals this sandbox hostname; local root `/var/lib/docker`) — proven sandbox-local. `sudo docker compose version` -> `v5.5.1`. No second startup attempt.

### Step 3 — verify-only rendered topology

- `sudo -n docker compose -f deploy/verify/compose.yaml -p mcs-em003v-discovery config` -> exit 0 (rendered above).
- Topology confirmed: 5 services `seed`/`mma2`/`modbus-simulator-runtime`/`modbus-replicator-runtime`/`osjs-shell`; project-scoped named volume `mcs-em003v-discovery_verify-data`; seed refuses any nonempty volume and writes `listeners: []` + `devices: []` + `devices: []`; `mma2` on `internal: true` `verify-runtime`; BOTH Go runtimes `network_mode: service:mma2`; `osjs-shell` only published `127.0.0.1:18219:18209` on `verify-ui`; no fixed `container_name`, no privileged mode, no host networking, no operator/bind mounts or external volumes, no public Modbus/RBE/access-event port.
- Port occupancy: `ss` not installed (exit 127); read-only fallback `cat /proc/net/tcp`/`tcp6` shows loopback `127.0.0.1:18219` (0x472B) FREE. No `compose up/build/down` executed.

### Step 4 — read-only inventory

- `sudo -n docker ps -a` -> empty. `sudo -n docker network ls` -> only default `bridge`/`host`/`none`. `sudo -n docker volume ls` -> empty. `sudo -n docker compose … ps -a` -> empty. No collisions; discovery project `mcs-em003v-discovery` distinct from the later executable project name. No volume contents inspected; nothing deleted/reused.

### Step 5 — local apply/restart interfaces (source-read only)

- Sockets: Simulator `<OSJS_DATA_DIR>/run/modbus-simulator.sock` (`simulator/runtime_transport_unix.go: const RuntimeSocketRelPath = "run/modbus-simulator.sock"`); Replicator `<OSJS_DATA_DIR>/run/modbus-replicator.sock` (`replicator/runtime_socket_unix.go`). In this stack `OSJS_DATA_DIR=/data`, shared `verify-data`.
- Framing (both runtimes): 4-byte big-endian uint32 length + JSON body, one request per connection. Simulator: `simulator/runtime_server.go` `handleRuntimeConn`/`writeRuntimeResponse`, `maxRuntimeMessage = 1<<20`, 30s conn deadline. Replicator: `replicator/cmd/modbus-replicator-runtime/main.go` (`binary.BigEndian.Uint32` header, `HandleRuntimeRequest`). Toolkit relay `OSJS/src/packages/MCSModbusToolkit/server.js` uses the same 4-byte BE framing, `MAX_MESSAGE = 1<<20`, 25s timeout.
- Request envelope: `{version:1, request_id:"<non-empty>", operation, payload}`; response `{version, request_id, ok, result, error{code,message}}`. Simulator operations `load`/`apply`/`status` (`payload.document` for apply; `payload.name` for status). Replicator operations `load`/`suggest`/`status`/`apply` (`replicator/runtime_api.go`; `apply` payload `{document:…}`).
- Toolkit relay routing: OS.js WebSocket provider; `request_id` prefixed `mcs-replicator-` -> Replicator socket, otherwise Simulator socket (`server.js`). It is not a plain HTTP endpoint. Hence a bounded synthetic parallel test must speak the Unix socket protocol directly (or drive the relay through OS.js auth).
- Shared lock primitive: `mma2composer.WithWriterLock(root, DefaultWriterLockTimeout=5s, action)` on the fixed file `<OSJS_DATA_DIR>/config/mma2/.writer.lock` (Linux kernel advisory flock, non-reentrant); both `simulator/writer_lock.go` and `replicator/writer_lock.go` delegate to it. Replicator `RuntimeManager.Apply` (`replicator/manager.go:120`) and Simulator `applyStructuralLocked` (`simulator/apply.go:83`) acquire it once at the outer transaction.
- MMA2 restart/ack: Simulator structural apply writes `config/mma2/restart-request.yaml` (fields `requested_at`, `reason`, `config_sha256`, `ports`) after config commit, awaits `config/mma2/restart-ack` equal to the request `config_sha256` within `DefaultRestartReadyTimeout=20s` (`simulator/restart.go`), then waits TCP readiness on each composed `127.0.0.1:<port>`. Deterministic ack = comparing ack bytes to the config SHA.
- Uncertainties for the later packet: (a) the Toolkit relay is an authenticated OS.js WebSocket provider, not a direct socket client — choose raw Unix-socket requests or authenticated relay; (b) MMA2 ingress binds `0.0.0.0:<port>` inside the container only (no host publish); (c) confirm `MAX_MESSAGE`/deadline parity between relay and runtimes across the shared volume. These must be pinned by the coding agent's later live packet.

### Step 6 — safe post-check / side effects

- Final `git status --porcelain --untracked-files=all` empty; `git rev-parse HEAD` = `4b840d9d…`; live `git ls-remote --exit-code origin refs/heads/main` = `4b840d9d…` (unchanged). Only side effect: sandbox-local daemon started, PID `717`, log `/tmp/mcs-em003v-dockerd.log`; `docker ps -a` and `docker volume ls` still empty (no app containers/volumes created).

Report commit: see pushed SHA below (transport section). This preparation establishes `source clean/current`, sandbox-local daemon, verify-only topology, clean inventory, free loopback port `18219`, and the local apply/restart interfaces. It does NOT establish `UMIG-EM-003-V PASS`.
