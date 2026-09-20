# UMIG-EM-003-L — CODE: Shared Linux writer-lock primitive

Status: ACTIVE — human approved UMIG-EM-003 CODE splitting and subsequent OpenCode TEST trial on 2026-09-20.
Stage / owner: CODE / ChatGPT.
Previous: UMIG-EM-002-V (archived COMPLETE / live VERIFY PASS).
Next: UMIG-EM-003-S (QUEUED CODE).

## Primary outcome
Introduce one reusable, bounded, crash-released Linux advisory transaction lock under the shared `OSJS_DATA_DIR/config/mma2` volume, without yet changing existing writer call sites. Define a single API for the later Simulator, Replicator and MMA-manager transaction boundaries. The Linux lock file must be the same for all processes using the same data root. Windows standalone behavior must not be changed.

## Scope and acceptance
1. Add the Linux lock primitive to `mma2composer`: reject empty/relative roots, nonpositive timeout and non-regular/symlink lock files; use one shared fixed lock path and an exclusive kernel lock with bounded wait, clear timeout/error, and release on every exit (including action error). Kernel-held lock must release on process exit; no stale PID-file ownership guesses.
2. Add a platform-guarded non-Linux implementation preserving existing standalone Windows behavior, while failing closed on unsupported non-Windows targets. Document that every outer transaction acquires ONCE and nested composer helpers must not reacquire; shell command approval is NOT a security sandbox.
3. Source-only checkpoint: inspect changed paths and read back exact authored files, record commit and remaining integration scope in handoff. Do not claim unit, race, live, or runtime PASS before independent TEST.

Non-scope: no callers rewired yet, no MMA2 runtime/restart modifications, no installed/operator files, no Docker, no RBE output, no Electron UI or production deployment. This intermediate CODE task does not prove concurrency safety until `UMIG-EM-003-S`, `UMIG-EM-003` and their JR TEST gate pass.

Dependencies: approved `docs/TOOLKIT_ELECTRON_MODEL_CONTRACT.md`, shared `mma2composer/composer.go`, Linux shared mounted data root. Sizing implementation/environment/behavior/verification/recovery = 1/1/1/0/1 = 4 (one tightly coupled primitive). Only coding agent archives/advances after source readback; ICC only through BLACK SHEEP WALL.
