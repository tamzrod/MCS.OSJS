# Simulator None Mode + Named-Pipe Runtime

Baseline commit: ee19b8a
Working tree: clean
Audited overlay: none
Source dependencies: simulator/validate.go, simulator/apply.go, simulator/store.go, simulator/memory_none_test.go, simulator/runtime_server.go, simulator/runtime_server_test.go, simulator/cmd/modbus-simulator-runtime/main.go, simulator/cmd/modbus-simulator-runtime/main_test.go, simulator/go.mod, simulator/go.sum, simulator/Dockerfile, deploy/docker-compose.yml, OSJS/src/packages/ModbusSimulator/server.js
Parent: simulator-device-config
Zoom In: none; leaf node
Zoom Out: simulator-device-config

## Boundary

This node owns the two things that changed in the Simulator runtime after the SIM-024
baseline and that the parent `simulator-device-config` node does not describe: the None
simulation mode (MEM-004 → MEM-008) and the Windows named-pipe runtime transport (7b26c4b).

## Established truth (None mode)

- A configured MMA2 area may keep `count > 0` with a zero random interval. `validateArea` in
  `simulator/validate.go` keeps the interval parameter but ignores it, so a zero interval is
  valid regardless of Count. The previous "interval_ms must be > 0 when count > 0" rejection
  is gone, and the corresponding "fc1 interval 0" invalid-input fixture was removed from
  `simulator/store_test.go`.
- Zero interval means None: the area stays allocated and externally served over Modbus, but
  no generator writes to it. New Memory devices default all four intervals to 0; previously
  saved positive intervals still load as Random (MEM-005/MEM-006).
- Truthful status (MEM-007): `needsRandomIngest` in `simulator/apply.go` is true only for an
  allocated area with positive timing. When false, `RuntimeStatus` reports Raw Ingest
  `NOT REQUIRED` with no raw error and returns `IDLE` while MMA2 is `RUNNING`, otherwise
  `WAITING`. This supersedes the parent SIM-021A rule that an enabled device with no accepted
  ingest stays `WAITING`: an all-None device now reaches `IDLE`. Mixed None/Random devices
  still require a successful ingest, and MMA2 unreachability still blocks any `RUNNING` state.
- Coverage: `simulator/memory_none_test.go` proves all-None areas validate, persist and reload
  with memory structure and intervals unchanged, produce no scheduled FC and no generated
  value, and that a mixed None/Random device schedules only the random area.
  `simulator/cmd/modbus-simulator-runtime/main_test.go` proves the runtime reloads a persisted
  interval change into a None schedule.

## Established truth (runtime document watcher)

- `modbus-simulator-runtime` now starts `watchDocument` (in
  `simulator/cmd/modbus-simulator-runtime/main.go`) before serving. It polls
  `store.DevicesPath()` every 250 ms and compares size plus modification time; on a change it
  reloads, validates every device, and only then calls `scheduler.ArmSchedules(doc)`. An
  unreadable or invalid document is logged and skipped, so prior schedules stay armed.

## Established truth (Windows named-pipe transport, 7b26c4b)

- The Simulator runtime no longer listens on a Unix-domain socket.
  `simulator/runtime_server.go` calls `winio.ListenPipe` on the fixed pipe
  `\\.\pipe\mcs-modbus-simulator` with security descriptor
  `D:P(A;;GA;;;SY)(A;;GA;;;BA)(A;;GRGW;;;AU)`; the existing-live-owner probe, stale-socket
  removal and `os.Chmod` logic were deleted. `RuntimeSocketPath` ignores its root argument and
  returns the pipe path; the `RuntimeSocketRelPath` constant is gone.
- `simulator/go.mod` adds `github.com/Microsoft/go-winio v0.6.2` plus indirect
  `golang.org/x/sys v0.10.0`. No `//go:build windows` constraint exists on these packages, and
  `winio.ListenPipe` is a Windows named-pipe API. Whether the Linux Go build still succeeds was NOT
  verified here (no Go toolchain is available in this environment), so this node records the Linux
  buildability of the simulator and replicator packages as unverified rather than failed.
- Unresolved transport mismatch: the OS.js relay
  `OSJS/src/packages/ModbusSimulator/server.js` still resolves
  `path.join(OSJS_DATA_DIR, 'run', 'modbus-simulator.sock')` and connects over TCP-style
  `net.createConnection`. The committed relay therefore no longer matches the committed runtime
  pipe name. The Replicator pair is consistent on the Electron side
  (`electron/replicator-runtime.js` uses `\\.\pipe\mcs-modbus-replicator`) but its OS.js package
  (`OSJS/src/packages/ModbusReplicator/server.js`) has the same stale unix-socket path. This node
  records the mismatch; it does not decide the fix.
- Docker images still build on `golang:1.25-bookworm` and run on `debian:bookworm-slim`
  (`simulator/Dockerfile`, `replicator/Dockerfile`). That Linux image path is now mismatched with a
  Windows-only pipe transport. Reported as inconsistent, not resolved.
