# Handoff

## Current authority — 2026-09-20

`UMIG-EM-002-V` OpenHands independent live VERIFY was reviewed and archived PASS at `c9baf86a0284a0bb2db1ae9cb7b29fb123efc64b`, bounded to a synthetic disposable stack only. No production, COMMS parity, privileged MMA manager, RBE network exposure or deployment claim.

Human approved the split UMIG-EM-003 CODE sequence and one subsequent OpenCode V2 JR TEST trial. `UMIG-EM-003-L` Linux writer-lock primitive is archived SOURCE ONLY (checkpoint `6dfb746e152d92c856e495691b0eac2bee182245`); `UMIG-EM-003-S` Simulator integration is archived SOURCE ONLY (checkpoint `b9d674759b6f1f3ea7d8e58389efb1553075875e` and its archived task records behavior and source diff). No Go builds/tests or runtime verifies have run against this new code. Source compilation correctness and all claimed locking properties remain unverified until the independent TEST.

SOLE ACTIVE: `workflow/active_work/umig-em-003-shared-lock.md`, CODE / ChatGPT. Previous S is archived; `workflow/active_work/umig-em-003-t-shared-lock.md` is QUEUED TEST. Next action: finish Replicator shared transaction lock and source audit/test cases, read back and checkpoint, archive CODE and activate only its named queued TEST with an exact JR packet. `UMIG-EM-003-V` is PLANNED, separate human permission required. ICC stale; only BLACK SHEEP WALL can edit ICC.

## JR TEST TASK — NONE / NOT AUTHORIZED

Operation CWAL is not yet authorized because ACTIVE is CODE. Experimental OpenCode on Legion `~/apps/MCS.OSJS-jr` is authorized for only ONE future local safe unit TEST when UMIG-EM-003-T is ACTIVE and a fresh packet exists. Remove the experimental profile's earlier no-CWAL sentence only at that gate; retain mandatory shell approval and edit denial, no --auto. Worktree is NOT a security sandbox; do not touch Legion Docker/sudo/Modbus/services/operator data. OpenHands is fallback only. JR may not edit code, workflow, ICC or advance state.

## Next action

ChatGPT implements and checkpoints Replicator without tests/production actions, then authors exact source-pinned safe local OpenCode TEST packet. Do not run stale UMIG-EM-002-V packet or infer a test PASS from source readback.
