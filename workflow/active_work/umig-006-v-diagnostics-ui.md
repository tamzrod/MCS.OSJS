# UMIG-006-V — VERIFY: Rendered Diagnostics Safety

Status: QUEUED — promoted 2026-09-18; wait for UMIG-006-T PASS.
Stage / owner: VERIFY / OpenHands (JR)
Previous: UMIG-006-T
Next: UMIG-007

## Primary outcome
Observe a truthful, safe Diagnostics tab in an actual OS.js window.

## Instruction / expected / evidence
In isolated OS.js, open Diagnostics with a reachable test runtime and one intentionally unavailable diagnostic source; inspect actual status/errors and Windows-only controls. Expected: accessible real observations, clearly unknown/unavailable missing data, no usable unsupported native service control. Record exact actions, screenshots, logs, HEAD and target. GUI/runtime not available = BLOCKED; wrong state = FAIL.

## Non-scope
No service starts/stops, source fixes, full visual parity or cutover.

## Dependencies
UMIG-006-T PASS and current JR packet.

## Sizing
Surface 0, environment 1, behavior 0, verification 1, recovery 0 = 2.
