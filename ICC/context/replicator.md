# Replicator Authorized Work

Baseline commit: ee19b8a
Working tree: clean; the previously audited Electron overlay is now committed, so no
uncommitted Replicator overlay exists.
Source dependencies: handoff.md, workflow/active_work/rep-block-002-independent-block-pollers.md, workflow/active_work/rep-block-003-tabbed-block-editor.md, workflow/archive/rep-*.md, replicator/runtime_api.go, replicator/cmd/modbus-replicator-runtime/main.go, replicator/go.mod
Zoom In: none
Zoom Out: active-work

## Current boundary

The original REP-002 through REP-006 implementation sequence is archived. Current authorized
Replicator work remains queued, with no active predecessor:

- REP-BLOCK-002: multi-block persistence, independent pollers, mixed FC1-FC4 replication,
  external destination serving, and end-to-end operational status. The record says
  implementation exists but JR retest is pending.
- REP-BLOCK-003: classic Device/Pull Blocks folder tabs and compact spreadsheet rows. It
  depends on REP-BLOCK-002 and awaits rendered retest. Its sizing totals 5, which the current
  rules allow only as a tightly coupled single workflow.

Neither Replicator record is ACTIVE, and neither declares `Previous`/`Next` links.

## Committed Replicator runtime truth (7b26c4b)

- `replicator/runtime_api.go` no longer derives a Unix-socket path.
  `RuntimeSocketPath` ignores its root argument and returns the fixed Windows named pipe
  `\\.\pipe\mcs-modbus-replicator`.
- `replicator/cmd/modbus-replicator-runtime/main.go` listens with `winio.ListenPipe` using
  security descriptor `D:P(A;;GA;;;SY)(A;;GA;;;BA)(A;;GRGW;;;AU)`. The previous directory
  creation, stale-socket removal and `os.Chmod 0o666` calls were removed.
- `replicator/go.mod` adds `github.com/Microsoft/go-winio v0.6.2` and indirect
  `golang.org/x/sys v0.10.0`. There is no `//go:build windows` constraint; Linux buildability of
  this package after the pipe change was not verified here and is recorded as unverified.
- The Electron client `electron/replicator-runtime.js` matches the pipe
  (`\\.\pipe\mcs-modbus-replicator`) and is consistent with the Go runtime. The OS.js package
  `OSJS/src/packages/ModbusReplicator/server.js` still resolves
  `$OSJS_DATA_DIR/run/modbus-replicator.sock`, so that relay no longer reaches the runtime.
  This node records the mismatch; it does not resolve it.

## Relevant adjacent Windows chain

The Replicator-facing Windows work is broader than the REP-BLOCK pair. RREC-001, RREC-002 and
RREC-003 are recorded COMPLETE in `workflow/active_work/`, covering the explicit ProgramData
runtime root, routing Electron load/apply/status through the live runtime transaction, and
truthful service/runtime errors. RLED-001 and RLED-002 are archived as completed. RREC-004
(installed-package proof) and RLED-003 through RLED-011 are QUEUED and PAUSED by the human
switch to the OS.js Toolkit migration; no Windows installer or COMMS-LED acceptance is
established.

## Workflow integrity observation

REP-BLOCK-002 declares sizing behavior 3, above the 0-2 range `planning/microtask/rules.md`
allows; it totals 8 and is oversized under the current rules.
