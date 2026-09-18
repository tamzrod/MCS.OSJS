# DOCKER-001 — CODE: Restore Linux runtime IPC without changing Windows pipes

Status: QUEUED — paused by human priority change on 2026-09-18; operator reports Docker deployment working, but required independent Linux/Windows build and IPC evidence is not yet recorded. Not COMPLETE or PASS.
Stage / owner: CODE / ChatGPT
Previous: none — independent human-authorized deployment regression repair
Next: none — independent parked verification; do not auto-activate from the UMIG chain.

## Repro and scope
The user's `git pull` reached `66fed78`, then `docker compose up -d --build` failed compiling Simulator on Linux: `undefined: winio.ListenPipe` and `undefined: winio.PipeConfig`; Replicator and OS.js builds were canceled. OS.js relays still use `$OSJS_DATA_DIR/run/modbus-simulator.sock` and `modbus-replicator.sock` whereas the Go runtimes used Windows pipes. The Compose stack was taken down without `-v`, so no volume removal is authorized. Subsequently the operator reported that Docker deployment is fine; no command/output or Windows evidence was supplied in that report.

## Implemented source checkpoint
- `simulator/runtime_server.go` retains its protocol, service logic and accept loop, delegating listener/path selection to `runtime_transport_windows.go` (same named pipe and ACL) or `runtime_transport_unix.go` (Unix socket under the shared data root, existing-owner protection, stale-socket handling and cleanup).
- `replicator/runtime_api.go` and runtime `main.go` delegate path and listener to Windows/Unix implementations. Windows pipe name and ACL remain unchanged; Linux uses `run/modbus-replicator.sock`, matching the unchanged OS.js relay.
- Simulator framing tests now dial via Windows/Unix test helpers rather than importing go-winio unconditionally; Linux Replicator tests cover relay path, live-owner protection, cleanup and stale-socket recovery.
- The Go framing protocol, MMA2, OS.js UI/relays, Electron and Docker Compose were not changed. Source files and patches were re-read; the actual Go/Docker gates have NOT been evidenced for this checkpoint.

## Verification required before COMPLETE
When explicitly resumed, collect independent `cd simulator && go test ./... && go build -o /tmp/mcs-simulator-test ./cmd/modbus-simulator-runtime`, `cd replicator && go test ./... && go build -o /tmp/mcs-replicator-test ./cmd/modbus-replicator-runtime`, Docker Compose build/status/logs and shared Unix-socket connectivity; Windows compilation and named-pipe framing fixtures need a Windows-capable runner. A user report of working deployment is recorded, not substituted for all gates. Preserve the named Docker volume; never use `docker compose down -v` or change production data to force testing.

## Continuation
Remain QUEUED and unverified while UMIG-001 is ACTIVE. Resume only by explicit human selection, maintaining exactly one ACTIVE. Only BLACK SHEEP WALL edits ICC; no automatic ICC refresh is a deployment prerequisite.
