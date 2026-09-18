# UMIG-CF-002 — CODE: Prepare Dormant Replicator Runtime Contract

Status: ACTIVE — human approved the proposed Replicator code-first preparation on 2026-09-19 while OpenHands is occupied. This does not waive independent TEST/VERIFY.
Stage / owner: CODE / ChatGPT
Previous: none — separately human-promoted preparation; NOT advancement from UMIG-003-T/V or UMIG-005.
Next: none — STOP at source-only checkpoint; explicit selection required for further coding or tests.

## Goal
Prepare a Toolkit-owned, transport-injected version-1 Replicator request/response contract as a dormant module plus future JR-only unit tests. Derive wire shapes from current `replicator/runtime_api.go` and existing OS.js Replicator relay, not from the Windows Electron bridge.

## Scope
- Add `OSJS/src/packages/MCSModbusToolkit/replicator-contract.js` implementing load/apply/status/suggest request envelopes with unique IDs; validate received version, request ID, ok and operation-specific result; preserve Go error code/message and transport failures.
- For load/apply, require a document with a `devices` array. For load additionally require destination suggestion. For status return the direct runtime status object (not a Simulator `{status}` wrapper). For suggest accept an optional inspect pair and return the direct ownership suggestion. Reject invalid inputs before sending and deep-snapshot apply at dispatch time.
- Add focused isolated future JR tests in `OSJS/tests/toolkit-replicator-contract.test.js`. Record exact command as a proposal, NOT a test run.
- Module must remain unimported by Toolkit entry points and UI. No sockets, OS.js bridge registration, Go/Docker changes, live requests, enabled buttons or user-data mutation.

## Source-only gate
Read back both files and inspect bounded diff. Archive only after source inspection with no build/test/runtime claims. Keep UMIG-003-T/V QUEUED/NOT RUN; UMIG-005 remains dependency-blocked; don't mark verified or automatically advance.
