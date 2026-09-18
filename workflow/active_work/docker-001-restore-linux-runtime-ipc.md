# DOCKER-001 — CODE: Restore Linux runtime IPC without changing Windows pipes

Status: ACTIVE — bounded source repair authored and read back; Linux/Windows build and live deployment verification still pending.
Stage / owner: CODE / ChatGPT
Previous: none — independent human-authorized deployment regression repair
Next: none — do not select unrelated work or archive before the required gate.

## Repro and scope
The user's `git pull` reached `66fed78`, then `docker compose up -d --build` failed compiling Simulator on Linux: `undefined: winio.ListenPipe` and `undefined: winio.PipeConfig`; Replicator and OS.js builds were canceled. OS.js relays still use `$OSJS_DATA_DIR/run/modbus-simulator.sock` and `modbus-replicator.sock` whereas the Go runtimes used Windows pipes. The Compose stack was taken down without `-v`, so no volume removal is authorized.

## Implemented source checkpoint
- `simulator/runtime_server.go` retains its protocol, service logic and accept loop, delegating listener/path selection to `runtime_transport_windows.go` (same named pipe and ACL) or `runtime_transport_unix.go` (Unix socket under the shared data root, existing-owner protection, stale-socket handling and cleanup).
- `replicator/runtime_api.go` and runtime `main.go` delegate path and listener to Windows/Unix implementations. Windows pipe name and ACL remain unchanged; Linux uses `run/modbus-replicator.sock`, matching the unchanged OS.js relay.
- Simulator framing tests now dial via Windows/Unix test helpers rather than importing go-winio unconditionally; Linux Replicator tests cover relay path, live-owner protection, cleanup and stale-socket recovery.
- The Go framing protocol, MMA2, OS.js UI/relays, Electron and Docker Compose were not changed. Source files and patches were re-read; the actual Go/Docker gates have NOT run in the editing environment (Go 1.23 only; repository checkout and Docker unavailable).

## Verification required before COMPLETE
Use a real checkout with Go 1.25 and Docker. Record independent `cd simulator && go test ./... && go build -o /tmp/mcs-simulator-test ./cmd/modbus-simulator-runtime`, `cd replicator && go test ./... && go build -o /tmp/mcs-replicator-test ./cmd/modbus-replicator-runtime`, then `cd deploy && docker compose up -d --build && docker compose ps` and runtime logs / shared Unix socket connectivity. Check Windows compilation and named-pipe framing fixtures on a Windows-capable runner. Report exact command outputs; a rebuild alone cannot prove the relays work. Do not use `docker compose down -v` or alter persistent data.

## Continuation
Keep this task ACTIVE until independent evidence proves the required deployment gate. If verification fails, report the exact failure, do not relabel the source checkpoint PASS. Only BLACK SHEEP WALL may edit ICC, and no ICC refresh is an automatic prerequisite for the deployment test.
