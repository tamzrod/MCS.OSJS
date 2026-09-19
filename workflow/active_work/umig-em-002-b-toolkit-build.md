# UMIG-EM-002-B — TEST: Upgraded Toolkit Node/build regression

Status: ACTIVE — human approved promotion 2026-09-19. Stage / owner TEST / OpenHands JR; packet, adjudication, advancement / ChatGPT.
Previous: UMIG-EM-002-T (archived COMPLETE / Go UNIT PASS). Next: UMIG-EM-002-V (PLANNED ONLY; separately human-promote).

## Primary outcome
Independently establish Node unit/fixture and local package build/discovery/desktop bundle evidence for the CURRENT OS.js MCS Modbus Toolkit. Product-source checkpoint: `8b5541fcefeba8df82bea675b15bd0d03c7c6c20`; its only changes since the previous Go-review checkpoint `425efaac4b06289723dd0d84bee7a0e873b34a88` concern Electron Windows installer permissions and a scoped ICC context, not OS.js or Go. Future commits require a guarded diff before testing. This TEST is not live runtime or GUI acceptance.

## Acceptance (the current handoff.md packet specifies all commands)
1. JR's own clean disposable checkout; confirm source checkpoint ancestry, allow only task/handoff documentation deltas, and perform a single guarded clean fast-forward if behind. Supported sandbox-local Node 16.x/npm, with dependency preparation limited to ignored local install artifacts. Unsafe state BLOCKED.
2. Execute separately the TWELVE existing `OSJS/tests/toolkit-*.test.js` files with Node, capturing command, exit and assertion evidence. Actual files are under `OSJS/tests/`, not the old Planning placeholder `OSJS/src/packages/MCSModbusToolkit/test/`. Any failing executed test FAIL and STOP, no repair/retries.
3. Separately execute `npm run build:local-packages`, `npm run package:discover`, `npm run build` inside `OSJS` and check the Toolkit bundle, manifest and shell HTML outputs; all gates exit 0 with clean tracked tree after. Generated ignored dist/package files are sandbox-only, not commits.

Non-scope: Go tests, backend sockets, Docker, rendered GUI, RBE/access-event exposure, Windows, production/customer data, retained volumes, source fixes, ICC, general CWAL or other task promotion. The only JR edit allowed is the exact report section of `handoff.md` with handoff-only commit/push and STOP. Go UNIT PASS from prior stage does not certify this TEST. Size: surface 0 + environment 1 + behavior 0 + verification 1 + recovery 1 = 3.
