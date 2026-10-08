# NPO-04 — OS.js Toolkit Native Persistence Verification

Status: PROMOTED / QUEUED
Stage: TEST/VERIFY
Owner: Independent JR TEST/VERIFY
Previous: NPO-03
Next: NPF-01

## Purpose
Independently verify OS.js Toolkit configuration of native MMA2 persistence.

## Acceptance
Verify with executed evidence:
- enable/default persistence through Toolkit;
- optional custom range/directory round-trip;
- known Modbus value survives MMA2-only restart;
- RBE TCP not required;
- State Sealing not required;
- disabled-memory isolation;
- no external Toolkit persistence runtime is involved.

STOP after verdict.
