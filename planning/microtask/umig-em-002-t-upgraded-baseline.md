# UMIG-EM-002-T — TEST: Changed Go / Toolkit baseline
Status: PLANNED; not executable until human promotion and exact current handoff packet. Stage/owner TEST / OpenHands JR. Previous: UMIG-EM-001 DESIGN. Next: UMIG-EM-002-V.

Outcome: Independently establish unit/build evidence for ALREADY merged MMA2 RBE, Simulator advanced persistence, Replicator measured COMMS and current Toolkit regressions; no artificial CODE stage.

Proposed isolated commands from repository root, separately record exit/stdout: `cd MMA2 && go test ./...`; `cd mma2composer && go test ./...`; `cd simulator && go test ./...`; `cd replicator && go test ./...`; `cd OSJS && npm run build:local-packages && npm run package:discover && npm run build`. Use Go >=1.25 modules, supported OS.js Node >=10 <17; separate commands in exact promotion packet. No Go network production dependencies, installer or installed data. Check changed backend JSON fields, RBE default disabled/validation, no config writes by read-only tests; report failures without fixes.

Evidence: full HEAD/clean precondition, actual individual command exits, paths and output, exact schema/source inspections and no runtime claim. No product/ICC/CWAL/workflow edits; JR only authorized report. Size surface 0/env 1/behavior 0/verify 1/recovery 1 =3.
