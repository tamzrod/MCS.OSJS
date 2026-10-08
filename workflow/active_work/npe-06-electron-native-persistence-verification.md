# NPE-06 — Electron Native Persistence Verification

Status: BLOCKED — headless Linux environment; superseded by NPE-06W Windows acceptance
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


## Blocked execution record — 2026-10-08

Independent JR could not launch the Electron product because DISPLAY was unset and no Xvfb/xvfb-run environment was available. Therefore the required Electron product acceptance path could not be exercised.

Supplementary native-runtime evidence was real and successful:
`python3 MMA2/test/persistence_manual/test.py` → exit 0, PASS (primary restore + backup recovery).

This does not constitute NPE-06 PASS. Product-level Electron verification is moved to:
`workflow/active_work/npe-06w-windows-electron-native-persistence-acceptance.md`.

Do not execute this packet again unless specifically reactivated.
