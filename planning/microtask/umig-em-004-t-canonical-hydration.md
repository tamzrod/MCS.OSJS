# UMIG-EM-004-T — TEST: Hydrated Simulator load
Status: PLANNED. Stage/owner TEST / OpenHands JR. Previous: UMIG-EM-004. Next: UMIG-EM-004-V.

Outcome: run `cd simulator && go test ./...` and `cd mma2composer && go test ./...` plus exact hydration regression command named by CODE checkpoint. Require temp files: legacy missing vs explicit false/empty, root/listener extra, corrupt config and no-on-load mutation with original bytes/hashes. Record HEAD, command exits and diff boundaries. No Go fix, Docker, ICC or existing volumes. Follow exact task-specific packet only after promotion; FAIL/BLOCKED STOP. Size 0/1/0/1/1=3.
