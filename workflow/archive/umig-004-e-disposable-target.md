# UMIG-004-E — CODE: Prepare Disposable UMIG-004-V Test Stack

Status: COMPLETE — source-only preparation and bounded follow-up, 2026-09-19; no production deployment or live VERIFY PASS.
Stage / owner: CODE / ChatGPT. Human approved creating a separate disposable Simulator/MMA2/OS.js environment after UMIG-004-T PASS.
Previous: UMIG-004-T (COMPLETE/PASS)
Next: UMIG-004-V (sole ACTIVE, executable live JR packet prepared separately)

## Outcome and source
- `deploy/verify/compose.yaml`: four test-only services, project-scoped volume, empty-volume-refusing seed, internal MMA2 network, Simulator shares only test MMA2's namespace, separate OS.js loopback UI bridge. No fixed container names, production mounts, host networking or exposed host Modbus ports.
- Follow-up after corrected JR preflight `37dc61f`: seed an explicit empty canonical Simulator `devices: []` document alongside empty MMA2 config. Test-only MMA2 container user `0:0` permits its supervisor to create `restart-ack` in `/data/config/mma2` created/updated by root seed and Simulator; no privileged container and no production Dockerfile/Compose change. The missing-file nil-slice representation remains unverified, not silently claimed fixed.
- `deploy/verify/README.md`: accepts freshly initialized sandbox-local Docker daemon as the target (no new VM), requires resource/port/daemon ownership recheck before startup and restricts synthetic test config plus scoped teardown.
- Original source-only checkpoint `39c2a3a5f6e4863360b161337c41a82158f98495` was read back; corrected JR nonmutating target preflight READY at `37dc61f533c8e7221610c8bf5858ddbe01790312`. This follow-up is source-only; Compose update and full live VERIFY have NOT RUN here.

## Boundary
Operator's Docker/user data and original `deploy/docker-compose.yml` untouched. The corrected preflight proves an eligible disposable sandbox, not the running test stack or real Memory functionality. UMIG-004-V remains ACTIVE until actual independent VERIFY result. Do not use `docker compose down -v`, delete any data volume without separate permission, or infer production acceptance.

## Sizing
Surface 1, environment 1, behavior 0, verification 0, recovery 1 = 3.
