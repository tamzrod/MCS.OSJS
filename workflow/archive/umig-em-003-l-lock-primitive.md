# UMIG-EM-003-L — CODE: Shared Linux writer-lock primitive

Status: COMPLETE / SOURCE-ONLY checkpoint 2026-09-20. NOT TESTED; independent TEST remains required after both writers are integrated.
Stage / owner: CODE / ChatGPT.
Previous: UMIG-EM-002-V (archived live VERIFY PASS).
Next: UMIG-EM-003-S (human-approved QUEUED CODE at source readback).

Outcome: introduced fixed shared `<OSJS_DATA_DIR>/config/mma2/.writer.lock` advisory Linux `flock` in `mma2composer.WithWriterLock`, requiring an absolute root, positive bounded timeout, private regular non-symlink file, exclusive kernel lock, release on callback error/return, and `ErrWriterLockTimeout` on contention. Explicit contract requires the outer transaction to acquire ONCE and its nested helpers not to reacquire. Platform-guarded fallback leaves standalone Windows behavior unchanged and fails closed on other non-Linux platforms. Added unit test cases for timed contention, action-error release, process-exit release, invalid root and symlink refusal; these tests were WRITTEN, not executed.

Source checkpoint: common API `mma2composer/writer_lock.go` commit `a61dabc60226d97c822fc55673146ce7543279f5`; Linux implementation commit `5717ff0c8451a888a441d7af789ed91616d7e1c6`; non-Linux implementation `ada0dda2b4ffe2434ec06311f4319a58a22b37f8`; focused tests `6dfb746e152d92c856e495691b0eac2bee182245`. GitHub source readback performed for all four paths and connector compare from baseline `c9baf86` confirmed exactly those four product test/source paths plus authorized workflow tasks at that point. No Go module command, race test, build, Docker, production or live verification ran. No caller uses the lock YET; this source checkpoint is not multiwriter protection and cannot be promoted directly to JR TEST before the Simulator and Replicator tasks.

Non-scope: existing Go writer call sites, MMA2 restart protocol, installed data and Electron UI remain untouched. Further authorized code tasks: `UMIG-EM-003-S` Simulator transaction boundary then `UMIG-EM-003` Replicator boundary/audit, then independently activated `UMIG-EM-003-T` OpenCode JR TEST trial. ICC only through BLACK SHEEP WALL. Original sizing 6 split into stages L/S/003; L sizing 1/1/1/0/1=4, one tightly coupled primitive.
