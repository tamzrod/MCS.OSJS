# NPE-06W — Windows Electron Native Persistence Acceptance

Status: PROMOTED / CURRENT
Stage: OPERATOR ACCEPTANCE / VERIFY
Owner: Human Operator
Previous: NPE-06 BLOCKED (headless Linux environment)
Next: NPO-01 on PASS

## Purpose
Complete the Electron native-persistence gate on the product's real Windows desktop environment.

NPE-06 was not a product failure. It was blocked because the Independent JR environment had no graphical display and could not launch Electron. Native MMA2 persistence itself passed the committed disposable donor harness.

## Target
Use the rebuilt MCS Modbus Toolkit Electron application on a disposable/safe Windows test memory.

Do not test against production/customer memory.

## Acceptance — default native persistence

1. Open the Electron application.
2. Create/select one disposable Memory device.
3. Ensure State Sealing is DISABLED for this case.
4. Ensure no RBE TCP output is required/configured for persistence.
5. Open Advanced Settings → Persistence.
6. Enable Persistence.
7. Leave Storage at Default / directory omitted.
8. Leave Persisted Memory at All Allocated Areas.
9. Save & Apply.
10. Using Modbus Poll or another normal Modbus client, write a distinctive holding-register value.
11. Wait a few seconds for native MMA2 persistence flush.
12. Restart ONLY the MMA2 service/process. Do not restart the Electron application or any external persistence helper.
13. Read the same register.
14. PASS condition: the value is restored after MMA2-only restart.

## Acceptance — disabled persistence

1. Disable Persistence for a fresh disposable memory or explicitly reset the test memory to a known baseline.
2. Save & Apply.
3. Write a distinctive value.
4. Restart ONLY MMA2.
5. PASS condition: the value is NOT restored from persistence.

Do not delete or manipulate snapshot files manually to manufacture this result.

## Acceptance — custom range

1. Enable Persistence.
2. Select custom ranges.
3. Configure one valid holding-register subrange and leave at least one nearby allocated register outside that persisted subrange.
4. Save & Apply.
5. Write distinct values inside and outside the persisted subrange.
6. Restart ONLY MMA2.
7. PASS condition:
   - inside-range value restores;
   - outside-range value follows normal non-persistent initialization behavior.

## Required observations

Record:
- application build/commit if shown;
- exact memory Port and Unit ID used;
- whether Persistence was enabled/disabled;
- whether directory was Default or custom;
- whether State Sealing was disabled;
- whether RBE TCP was absent/disabled;
- test register address(es);
- written value(s);
- value(s) read after MMA2-only restart;
- exact MMA2 service/process name restarted;
- PASS/FAIL for each of the three cases.

Screenshots are optional but useful.

## Verdict

PASS only if:
- Electron Save & Apply accepts native persistence configuration;
- default native persistence restores after MMA2-only restart;
- no State Sealing dependency is required;
- no RBE TCP dependency is required;
- disabled persistence does not restore;
- custom-range behavior is correct.

FAIL if any executed product behavior contradicts those expectations.

BLOCKED only if the Windows target itself cannot execute the required steps.

On PASS, route directly to NPO-01 — OS.js Toolkit Persistence UI.
On FAIL, stop and report the exact failed step without starting OS.js Toolkit work.
