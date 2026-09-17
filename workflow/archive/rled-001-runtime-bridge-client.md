# RLED-001 — Establish Replicator Runtime Bridge Client

Status: COMPLETED — Windows named-pipe fixture verification passed 2026-09-17
Previous: none
Next: RLED-002

## Primary outcome
Provide a bounded Electron-side client for the existing Go Replicator runtime named-pipe protocol.

## Scope
Implement a small reusable Node client matching existing version-1 length-prefixed JSON frames, request IDs, max frame size and Windows named-pipe path. Test success, runtime failures and invalid responses with a local named-pipe fixture. Do not wire the UI yet.

## Non-scope
No new daemon, HTTP API, IPC operation or Go protocol/version change. Real Windows end-to-end is RLED-011.

## Acceptance
1. Calls return the Go API result and reject API errors without declaring health.
2. Connection, incomplete-frame, timeout and invalid-response cases reject cleanly.
3. Named-pipe path and frame contract agree with `replicator/runtime_api.go` and its runtime server.

## Verification
`node --test electron/test/replicator-runtime.test.js` on a repository checkout; inspect wire contract and re-read written source. Unavailable checks remain unverified.

Verification evidence: On the full Windows repository checkout, node --test electron/test/replicator-runtime.test.js passed all seven fixtures. The client and Go service agree on the Windows Replicator named pipe; success, runtime error, mismatched ID, invalid length, truncated response, timeout, unavailable pipe, unsupported operation and oversized request behavior were exercised. go vet ./... passed for both Replicator and Simulator modules. The focused Simulator named-pipe framing test also passed. Full module suites still contain unrelated environment-dependent failures documented in the handoff.

## Dependencies
None. MEM-004 is paused QUEUED, not completed.

## Sizing
Surface 1, environment 0, behavior 1, verification 1, recovery 0 = 3.
