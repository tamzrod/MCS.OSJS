# UMIG-EM-002-B — TEST: Upgraded Toolkit Node/build regression

Status: PLANNED — not executable until separate human promotion. Stage/owner TEST / OpenHands JR. Previous: UMIG-EM-002-T. Next: UMIG-EM-002-V (PLANNED).

Primary outcome: Independently re-run Toolkit Node unit/fixture regressions and OS.js local package build/discovery/bundle against the already merged new Go/MMA2/Replicator baseline. Separate from Go test workflow to localize failures.

Scope/acceptance: inspect supported Node >=10 <17 (use sandbox-local Node 16 if necessary), clean checkout; exact commands must be authored in a future JR packet: Node test files for Memory, Replicator, Diagnostics and fixtures in `OSJS/src/packages/MCSModbusToolkit/test/` (first verify actual paths), then from `OSJS` separately `npm run build:local-packages`, `npm run package:discover`, `npm run build`. Capture each exit, discovery/bundle evidence, source-only UNKNOWN COMMS and disabled native controls; do not claim live Go/API success. No code edits or automated repair.

Non-scope Go test execution, Docker, production, ICC, general CWAL, advancing workflow or product changes. Evidence: raw commands/outputs and clean tracked diff. Sizing 0/1/0/1/1=3.
