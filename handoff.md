# Handoff

## Current

ACTIVE: RLED-001 — Establish Replicator Runtime Bridge Client. Human approved switching to COMMS LED coding.

QUEUED RLED sequence (approved): RLED-002 → RLED-003 → RLED-004 → RLED-005 → RLED-006 → RLED-007 → RLED-008 → RLED-009 → RLED-010 → RLED-011. Each detailed record in `workflow/active_work/` has explicit Previous/Next; RLED-011 is final Windows acceptance.

PAUSED QUEUED Memory chain: MEM-004 → MEM-005 → MEM-006 → MEM-007 → MEM-008. MEM-001–003 completed and archived. MEM-004 is NOT complete: `cd simulator && go test -count=1 ./... && go vet ./...` still requires recorded passing evidence. Previously implemented MEM-004–007 code remains; no Memory task may be archived without its gate. Resume the Memory chain only after deliberate selection, maintaining exactly one ACTIVE task.

Other QUEUED work unchanged: REP-BLOCK-002 and REP-BLOCK-003 retests.

## Approved Replicator COMMS design

Selected source device: SOURCE Network, TCP, Modbus LEDs; DESTINATION MMA2 LED. Permanent display is labels plus four small circles only. Circle hover/focus/tap reveals actual diagnostics. Green=observed healthy, yellow=Modbus exception/warning, red=confirmed connection/write failure or Modbus response timeout, gray=disabled/not tested/unknown/stale. ICMP cannot override a working TCP/Modbus connection. Source TCP connection is per poll (green means recent success, not persistent socket). MMA2 flashes only after its existing Raw Ingest positive acknowledgement; TCP flashes on actual activity, never from repeated status polling. Preserve per-FC/block errors and distinguish source from destination failures.

## Coding order and gates

RLED-001 first establishes an Electron Node client for the existing Go runtime's version-1 length-framed JSON Unix-domain socket, rooted at the Electron data directory. RLED-002 replaces Electron's current hardcoded `{running:true,source_status:'CONFIGURED',last_poll:''}` with the real runtime query. Go Replicator existing per-block Source/LastPoll/LastError is insufficient to infer separate TCP/Modbus/MMA2 state; RLED-003–007 add truthful observations, RLED-008–010 implement LED UI/activity. RLED-011 requires real Windows proof. Follow task-defined gates and stop formal advancement at first unverified gate.

Existing MMA2 allocation, Raw Ingest protocol, Modbus semantics, Poll Blocks, Memory functionality and OS.js UI are out of scope. Preserve parked local/uncommitted Electron overlays on checkout; GitHub main describes only the committed baseline.
