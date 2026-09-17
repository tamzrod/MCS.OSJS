# Handoff

## Current

COMPLETED + ARCHIVED:
- RLED-001 — Establish Replicator Runtime Bridge Client — Windows named-pipe fixture verification passed.

ACTIVE: RLED-002 — Replace Fake Electron Replicator Status.

QUEUED RLED sequence (approved): RLED-003 → RLED-004 → RLED-005 → RLED-006 → RLED-007 → RLED-008 → RLED-009 → RLED-010 → RLED-011. Each detailed record in workflow/active_work/ has explicit Previous/Next; RLED-011 is final Windows acceptance.

PAUSED QUEUED Memory chain: MEM-004 → MEM-005 → MEM-006 → MEM-007 → MEM-008. MEM-001–003 completed and archived. MEM-004 is NOT complete: cd simulator && go test -count=1 ./... && go vet ./... still requires recorded passing evidence. Resume only after deliberate selection while maintaining exactly one ACTIVE task.

Other QUEUED work unchanged: REP-BLOCK-002 and REP-BLOCK-003 retests.

## Approved Windows transport

Electron and its Go backend runtimes are Windows-only. Local runtime IPC uses secured Windows named pipes while preserving the version-1 four-byte big-endian length-prefixed JSON protocol:

- Replicator: Windows pipe mcs-modbus-replicator
- Simulator: Windows pipe mcs-modbus-simulator

Pipe security grants full access to SYSTEM and Administrators and read/write access to authenticated users. Filesystem Unix sockets, stale-socket cleanup and chmod are not used.

## Approved Replicator COMMS design

Selected source device: SOURCE Network, TCP, Modbus LEDs; DESTINATION MMA2 LED. Permanent display is labels plus four small circles only. Circle hover/focus/tap reveals actual diagnostics. Green=observed healthy, yellow=Modbus exception/warning, red=confirmed connection/write failure or Modbus response timeout, gray=disabled/not tested/unknown/stale. ICMP cannot override a working TCP/Modbus connection. Source TCP connection is per poll (green means recent success, not persistent socket). MMA2 flashes only after its existing Raw Ingest positive acknowledgement; TCP flashes on actual activity, never from repeated status polling. Preserve per-FC/block errors and distinguish source from destination failures.

## RLED-001 evidence

On the full Windows checkout, node --test electron/test/replicator-runtime.test.js passed all seven named-pipe fixtures. Success, runtime errors, mismatched request IDs, invalid lengths, truncated responses, timeout, unavailable pipe, unsupported operation and oversized request behavior were verified. The focused Simulator named-pipe framing test passed. go vet ./... passed for both modules.

The full Replicator suite remains blocked by pre-existing environment requirements: missing temporary MMA2 executable, missing restart acknowledgement, and configured Windows data-root expectations. The full Simulator suite remains blocked by pre-existing ownership-map and configured Windows data-root expectations. These failures are not substituted for the passing RLED-001 task-defined gate.

## Next execution

RLED-002 must replace Electron's hardcoded status with callReplicatorRuntime using the status operation, preserve load/apply behavior, propagate unavailable-runtime errors, add a focused IPC integration test, and verify packaged module inclusion.

Existing MMA2 allocation, Raw Ingest protocol, Modbus semantics, Poll Blocks, Memory functionality and OS.js UI remain out of scope.