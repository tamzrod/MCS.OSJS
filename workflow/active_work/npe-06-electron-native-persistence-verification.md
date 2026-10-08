# NPE-06 — Electron Native Persistence Verification

Status: PROMOTED / QUEUED
Stage: TEST/VERIFY
Owner: Independent JR TEST/VERIFY
Previous: NPE-05
Next: NPO-01

## Purpose
Independently verify the Electron product against native MMA2 persistence before OS.js Toolkit work begins.

## Required acceptance path
Using a disposable/safe test memory:

1. configure persistence in Electron;
2. leave directory omitted/default for at least one case;
3. Save & Apply;
4. write known Modbus values;
5. allow native MMA2 persistence to flush;
6. restart ONLY MMA2;
7. read the same values and confirm restore;
8. prove persistence works with RBE TCP absent/disabled;
9. prove persistence does not require State Sealing;
10. verify disabled persistence does not create/restore persisted state;
11. verify custom range behavior if configured.

Do not change product source while acting as Independent JR except under separate explicit authorization.

Record exact commands/actions and raw evidence. PASS only on executed evidence. STOP after verdict.
