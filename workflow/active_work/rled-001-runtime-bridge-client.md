# RLED-001 — Establish Replicator Runtime Bridge Client

Status: ACTIVE
Previous: none
Next: RLED-002

## Primary outcome
Provide a bounded Electron-side client for the existing Go Replicator runtime socket protocol.

## Scope
Implement a small reusable Node client matching existing version-1 length-prefixed JSON frames, request IDs, max frame size and runtime socket path under the existing data root. Test success, runtime failures and invalid responses with a local test socket. Do not wire the UI yet.

## Non-scope
No new daemon, HTTP API, IPC operation or Go protocol/version change. Real Windows end-to-end is RLED-011.

## Acceptance
1. Calls return the Go API result and reject API errors without declaring health.
2. Connection, incomplete-frame, timeout and invalid-response cases reject cleanly.
3. Socket path and frame contract agree with `replicator/runtime_api.go` and its runtime server.

## Verification
`node --test electron/test/replicator-runtime.test.js` on a repository checkout; inspect wire contract and re-read written source. Unavailable checks remain unverified.

## Dependencies
None. MEM-004 is paused QUEUED, not completed.

## Sizing
Surface 1, environment 0, behavior 1, verification 1, recovery 0 = 3.
