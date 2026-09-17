# RLED-011 — Windows COMMS Acceptance

Status: QUEUED
Previous: RLED-010
Next: none

## Primary outcome
Prove the four COMMS LEDs accurately represent live Replicator-to-MMA2 behavior in the installed Windows Electron app.

## Scope
Exercise source reachable, TCP refused/timeout, Modbus success/exception/timeout, multiple FC blocks, unreachable/blocked ICMP, MMA2 confirmed write and failure/recovery, stale/disabled states, tooltip hover/focus/tap and edit-focus persistence.

## Non-scope
No additional feature implementation or unrelated packaging redesign.

## Acceptance
1. Every color and tooltip corresponds to real evidence, including gray untested states.
2. Activity flashes occur only on actual source activity or acknowledged MMA2 writes.
3. Existing Replicator config and Memory tabs still work; final build/Windows evidence recorded.

## Verification
Real Windows `npm run dist:win` and installed-app behavioral evidence plus required repository-native Go/Node tests. If Windows environment unavailable, keep QUEUED/ACTIVE and report gate, never archive unverified.

## Dependencies
RLED-010.

## Sizing
Surface 0, environment 1, behavior 1, verification 1, recovery 0 = 3.
