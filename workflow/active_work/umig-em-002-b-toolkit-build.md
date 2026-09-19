# UMIG-EM-002-B — TEST: Upgraded Toolkit Node/build regression

Status: ACTIVE — human approved promotion 2026-09-19. Stage / owner TEST / OpenHands JR; packet, evidence adjudication and advancement / ChatGPT.
Previous: UMIG-EM-002-T (archived COMPLETE / Go UNIT PASS at `425efaac4b06289723dd0d84bee7a0e873b34a88`). Next: UMIG-EM-002-V (PLANNED ONLY; requires separate human promotion).

## Primary outcome
Independently establish deterministic Node unit/fixture and local package build/discovery/desktop bundle evidence for the CURRENT OS.js MCS Modbus Toolkit source, following the Electron model design. This is not Go live-runtime or GUI acceptance. Baseline product source `425efaac4b06289723dd0d84bee7a0e873b34a88`; future Git commits must be explicitly guarded for product-source changes before testing.

## Acceptance (exact packet in handoff.md)
1. In JR's own clean disposable checkout, prove baseline ancestry, unchanged product source, and safe guarded one-time fast-forward if necessary. Use supported sandbox-local Node 16.x/npm (OSJS engine >=10 <17); prepare package dependencies without changing tracked manifests or project config. Missing safe prep BLOCKED, not an automatic product FAIL.
2. Execute each of the TWELVE existing `OSJS/tests/toolkit-*.test.js` scripts separately with `node`, recording every exit and assertion output. The authoritative test files live in `OSJS/tests/` (not `OSJS/src/packages/MCSModbusToolkit/test/`, the original Planning placeholder). Each must exit 0; stop on first executed contradiction, no source repairs or retry to force PASS.
3. Separately run `npm run build:local-packages`, `npm run package:discover`, `npm run build` from `OSJS`, all exit 0; prove nonempty Toolkit JS/CSS, discovered Toolkit manifest and desktop HTML output; show clean tracked tree after. Discovery and builds may create ignored sandbox build outputs only. Capture warnings, exact commands and output; no substituted syntax check.

Non-scope: Go suites, backend sockets, Docker, running GUI, production/retained volumes, RBE or access-event network output, Windows, source changes, ICC or general CWAL edits. Go UNIT result from prior stage does not prove Node/build or live runtime. JR's sole edit is the explicitly authorized `handoff.md` report, committed/pushed alone, then STOP. If FAIL/BLOCKED, no autonomous next stage. Size surface 0 + environment 1 + behavior 0 + verification 1 + recovery 1 = 3.
