# Handoff

## Current direction — 2026-09-18

Human priority: preserve the MCS.OSJS desktop and backend functionality while migrating to one `MCS Modbus Toolkit`. MMA2, Go services, persistent data and Windows Electron remain in scope. The early Toolkit placeholder build and rendered one-window checks are complete. Service-free portable Electron and Modpoll are discussed concepts only, not promoted implementation tasks.

## Roles and authoritative execution state

ChatGPT owns CODE, source checkpoints, review of independent TEST/VERIFY evidence and workflow advancement. OpenHands/JR performs only explicitly authorized tests via `operation cwal.md` and never fixes source, changes workflow or edits ICC. Only BLACK SHEEP WALL may maintain ICC.

**ACTIVE: DOCKER-001 — restore Linux runtime IPC.** The user explicitly authorized this bounded deployment regression fix and a push to `main`. Platform-specific Go transport source is authored and read back, but Linux Go builds, Windows build/fixtures, Docker image construction, runtime connectivity and full Compose deployment have NOT been executed or passed for this patch. The source checkpoint must not be mistaken for deployment PASS. There is no current JR TEST TASK or successor promotion. UMIG-001 and UMIG-003 onward remain in Planning; previously parked Windows, Memory and Replicator features remain unchanged.

## Reported deployment failure and bounded source repair

User ran `cd deploy && docker compose down`, then pulled `cd50830..66fed78` and attempted `docker compose up -d --build`. Linux Simulator Docker compilation failed with `runtime_server.go:205:25: undefined: winio.ListenPipe` and `:205:53: undefined: winio.PipeConfig`; Replicator and OS.js builds were canceled. The Compose stack is down; `docker compose down` without `-v` does not remove the named data volume.

The Go Simulator and Replicator runtimes had been switched to Windows-only named pipes, while OS.js server relays still dialed `$OSJS_DATA_DIR/run/modbus-simulator.sock` and `modbus-replicator.sock`. DOCKER-001 separates each platform transport using Go build constraints: Windows retains its original named-pipe paths and security descriptors; Linux uses Unix sockets in the shared data root with owner protection, stale-socket handling and cleanup. Simulator's framing tests now have per-platform dial helpers; Linux Replicator listener tests cover path, live owner, stale socket and cleanup. The version-1 length-prefixed JSON protocol and OS.js relay sources were not altered. No MMA2, Electron, Docker Compose or persistent configuration was changed. See `workflow/active_work/docker-001-restore-linux-runtime-ipc.md` for the bounded task and pending gates.

Source edits were re-read and commit patches inspected. The editing environment has Go 1.23 only, lacks a repository checkout and Docker, and cannot fetch GitHub with container DNS; the modules require Go 1.25. Therefore no Go/Docker product build is claimed. This is an **unverified source checkpoint**, not a completed or deployed repair.

## Existing verified checkpoints

- UMIG-002 CODE: disconnected Toolkit placeholder source archived. UMIG-002-R added the missing package manifest at `cd67e15`; original independent FAIL remains at https://github.com/tamzrod/MCS.OSJS/blob/752a54108ecaf91d9e53440ed344f8f31a1ce8de/handoff.md .
- UMIG-002-T TEST: independent build/discovery PASS at `e9d25e3`; archive `workflow/archive/umig-002-t-build-discover.md`; full report at https://github.com/tamzrod/MCS.OSJS/blob/e9d25e3d4c1b832126a161a6071fd4abf8b5546b/handoff.md .
- UMIG-002-V VERIFY: independent rendered one-window PASS at `7349caa`; archived. Clean Chromium session showed exactly one Toolkit placeholder window after one menu click, working desktop, taskbar and Start menu, and both legacy application entries present. Raw report at https://github.com/tamzrod/MCS.OSJS/blob/7349caace4762cedcef8109371fb8b02f7a0e8c2/handoff.md . The preceding ICC-gate BLOCKED attempt remains historical at `3831e33`. These gates prove only build/discovery and placeholder rendering, not backend connectivity or final migration.

## Observations and limits

Four earlier boot-time missing icon/sound assets remain recorded outside the previous one-window gate, not fixed or part of DOCKER-001. The Go transport repair has not verified connected Simulator/Replicator functionality, Windows packaging, service-free portable behavior or any new Modpoll tab. ICC remains a stale optional semantic cache and is not current task authority; only BLACK SHEEP WALL may refresh relevant context when actually needed. Do not treat ICC staleness as a reason to block an otherwise self-contained deployment check.

## Next action and recommendation

Next action (operator after pulling latest `main`): from `deploy/`, run `docker compose up -d --build` (do not use `down -v`), then `docker compose ps` and `docker compose logs --tail=100 modbus-simulator-runtime modbus-replicator-runtime osjs-shell mma2`; report raw output and whether `/data/run/modbus-simulator.sock` and `/data/run/modbus-replicator.sock` are present and the OS.js relay calls work. If builds fail, report the first actual error and stop rather than claiming success. A Windows-capable runner must separately confirm Windows cross-build/named-pipe fixtures. Recommendation: preserve the named data volume and do not archive DOCKER-001 or start an unrelated task until required verification evidence is reviewed.
