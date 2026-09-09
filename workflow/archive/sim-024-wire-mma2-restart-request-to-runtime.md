# SIM-024 — Wire MMA2 Restart Request to the Deployed Runtime

Status: COMPLETED 2026-09-09 — independently supervised MMA2 reload verified end to end.
Previous: none
Next: none

## Primary outcome

Make a Simulator structural **Save & Apply** actually restart the independently managed MMA2 runtime, reload the committed shared MMA2 configuration, and return ready on the newly configured listener port.

## Observed failure

The current UI reports:

```text
Save & Apply failed: MMA2 structural apply failed: mma2 restart not ready within timeout on 127.0.0.1:5020: dial tcp 127.0.0.1:5020: i/o timeout
```

## Root cause established from current repository

- `simulator/SchedulerApplier.ApplyStructural` composes the shared MMA2 config, writes `$OSJS_DATA_DIR/config/mma2/restart-request.yaml`, then immediately waits for the configured MMA2 listener ports to become reachable.
- `simulator/restart.go` only persists the restart request and polls readiness; it does not and must not own the MMA2 process.
- Production repository code contains no consumer of `restart-request.yaml`; the only consumers found are tests/harnesses that simulate an external restart actor.
- `deploy/docker-compose.yml` currently defines `osjs-shell` and `modbus-simulator-runtime` only. It does not define the independently managed MMA2 runtime that must consume the shared config and restart request.
- Therefore a structural Save & Apply can commit the new listener configuration and request a restart, but nothing in the deployed stack performs that restart. The readiness loop then truthfully times out on the new port.

## Scope

1. Add the independently managed MMA2 runtime to the deployed MCS.OSJS stack using the shared `osjs-data` volume and the effective config at `/data/config/mma2/config.yaml`.
2. Add the smallest appliance/deployment-side restart actuator that consumes the Simulator's `restart-request.yaml` contract and restarts/re-execs MMA2 with the committed effective config.
3. Keep process ownership outside `simulator/`. The Simulator remains limited to the existing restart request contract; it must not gain START, STOP, SPAWN, KILL, REPLACE, Docker-control, or general MMA2 lifecycle ownership.
4. Ensure one restart request produces exactly one MMA2 restart/reload. The actuator must not enter a restart loop while the request file remains present during the Simulator readiness window.
5. Preserve the existing success contract: after MMA2 is reachable on every composed Simulator port, the Simulator clears the request and arms schedules.
6. Preserve negative paths: rejected config, failed composition, no-change apply, and timing-only apply must not restart MMA2.

## Non-scope

- No general MMA2 control API.
- No restoration of the removed Simulator-owned `MMA2Lifecycle` / `MMA2_BINARY` process-management design.
- No Simulator START or STOP command.
- No UI redesign or new status vocabulary.
- No changes to Modbus/raw-ingest semantics except what is required to prove the restarted MMA2 instance is serving the committed configuration.

## Acceptance criteria

1. `docker compose` starts MMA2 independently of the Simulator and MMA2 loads `/data/config/mma2/config.yaml`.
2. With the stack running, changing the Simulator listener to port `5020` and pressing **Save & Apply** causes exactly one restart/reload of MMA2 and the apply returns success instead of `MMA2_NOT_READY` timeout.
3. After the apply succeeds, `127.0.0.1:5020` accepts Modbus connections using the committed Simulator entry.
4. The restart actuator consumes each request once and does not repeatedly restart MMA2 while the same request remains pending.
5. Rejected/invalid structural config causes no restart.
6. Timing-only and no-change apply cause no restart.
7. Static inspection confirms `simulator/` still contains no process spawn/signal/kill/container-control implementation.
8. Existing Simulator restart/readiness tests continue to pass, and new tests cover the production restart-consumer behavior.

## Required repository-native verification

At minimum, run the native gates for every changed executable surface:

```text
cd simulator && go test -count=1 ./...
cd MMA2 && go test -count=1 ./...
docker compose -f deploy/docker-compose.yml config
docker compose -f deploy/docker-compose.yml build
```

Then perform the real deployed verification:

```text
docker compose -f deploy/docker-compose.yml up -d
# structural Save & Apply to a new listener such as 5020
# prove MMA2 restarted/reloaded exactly once
# prove the new listener is reachable and the apply succeeds
```

If the repository-native Docker/deployment verification cannot run, keep SIM-024 ACTIVE and report verification unavailable; do not archive or claim completion.

## Completion evidence required

Record:

- changed production paths that own the independent MMA2 runtime/restart actuator;
- exact restart-request consumption behavior;
- proof that one request caused one restart/reload;
- proof the new configured listener became reachable after Save & Apply;
- proof Simulator process-ownership boundaries remain intact;
- exact native verification commands and results.

## Completion evidence

- `deploy/docker-compose.yml` now runs independent `mcs-mma2`; `MMA2/Dockerfile.supervised`, `cmd/mma2-supervisor`, and `internal/restartwatch` own appliance-side process supervision and exact-once request observation.
- The Simulator writes a config-SHA restart request, removes stale acknowledgement, and waits for the appliance to acknowledge that exact SHA before accepting listener readiness. MMA2 writes the acknowledgement only after consuming the request and launching the replacement process; Simulator clears request and acknowledgement after readiness.
- A real UI structural apply changed Sim-PLC-1 from port 5020 to 5021. Logs contain exactly one `restart request consumed`, followed by `ingress sim-5021-1 listening`; Save & Apply returned success, TCP 5021 was reachable, and the browser showed both MMA2 and Simulator `RUNNING` from accepted Raw Ingest evidence.
- The canonical device document persists port 5021 and both restart artifacts are cleared. Static inspection finds process control only in test harnesses under `simulator/`; production Simulator ownership remains unchanged.
- `cd simulator && go test -count=1 ./...`, `cd MMA2 && go test -count=1 ./...`, `docker compose -f deploy/docker-compose.yml config`, and `docker compose -f deploy/docker-compose.yml build` pass. The MMA2 gate also restored its referenced but missing `test/policy_test.yaml` fixture.
