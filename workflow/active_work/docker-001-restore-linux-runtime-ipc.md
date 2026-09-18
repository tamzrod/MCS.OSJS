# DOCKER-001 — CODE: Restore Linux runtime IPC without changing Windows pipes

Status: ACTIVE — human explicitly authorized fix and push to main on 2026-09-18; implementation and independent verification pending.
Stage / owner: CODE / ChatGPT
Previous: none — independent deployment regression repair
Next: none — stop after source checkpoint and request deployment verification; no unrelated task promotion.

## Repro and scope
After `docker compose down`, `git pull` to `66fed78` and `docker compose up -d --build`, Linux Simulator compilation failed at `simulator/runtime_server.go` with undefined `winio.ListenPipe` and `winio.PipeConfig`; Replicator and OS.js builds were canceled. The Linux OS.js relays still connect to Unix sockets at `$OSJS_DATA_DIR/run/modbus-{simulator,replicator}.sock`, but Go runtimes were switched to Windows-only named pipes. Original evidence is user-supplied deployment log; no product test has passed for this repair.

## Bounded repair
Separate Go transport implementations by OS: retain the exact existing Windows named-pipe names, ACLs and JSON framing; restore Linux Unix-socket paths, listener ownership/cleanup and matching relay endpoints for Simulator and Replicator. Ensure cross-platform Simulator test code does not unconditionally import Windows-only go-winio. Avoid MMA2 protocol, OS.js UI, Electron installer, configuration, persistent volume and unrelated changes.

## Gates
Read back all authored files; inspect diff for Windows behavior and Linux relay compatibility. Required independent evidence: `cd simulator && go test ./... && go build -o /tmp/mcs-simulator-test ./cmd/modbus-simulator-runtime`; `cd replicator && go test ./... && go build -o /tmp/mcs-replicator-test ./cmd/modbus-replicator-runtime`; `cd deploy && docker compose up -d --build`, `docker compose ps` and appropriate runtime socket/relay checks; check Windows cross-compilation and named-pipe fixtures on a Windows-capable runner. Do not call a source-only patch or a generic parser equivalent to these gates. Do not change ICC except via separate BLACK SHEEP WALL directive.

## Continuation
If direct execution is unavailable, push the bounded source checkpoint as unverified, keep task ACTIVE and ask for independent TEST evidence. Never archive or advance on an unexecuted build. The user retains control of live Docker lifecycle and persistent data; never run `docker compose down -v`.
