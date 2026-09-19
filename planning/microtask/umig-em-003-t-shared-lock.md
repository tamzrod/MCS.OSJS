# UMIG-EM-003-T — TEST: Shared writer lock
Status: PLANNED. Stage/owner TEST / OpenHands JR. Previous: UMIG-EM-003. Next: UMIG-EM-003-V.

Outcome: prove source change compiles and concurrent writer regression tests pass without touching running services. Proposed commands root: `cd mma2composer && go test ./...`, `cd simulator && go test ./...`, `cd replicator && go test ./...` plus exact focused race/concurrency tests named in CODE checkpoint. Require Go >=1.25, temp directories only; record command exits and evidence that lock timeout/crash and foreign owners fail closed, no reentrant hang. No product edits or Docker/installed volume. Exact current packet required at activation; FAIL/BLOCKED + STOP. Size 0/1/0/1/1=3.
