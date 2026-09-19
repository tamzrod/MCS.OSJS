# UMIG-005-V — VERIFY: Live Replicator Behavior

Status: ACTIVE — promoted 2026-09-19 after independent UMIG-005-T UNIT/BUILD PASS review; LIVE VERIFY NOT RUN / NO PASS.
Stage / owner: VERIFY / OpenHands JR; evidence review / ChatGPT
Previous: UMIG-005-T (COMPLETE/PASS, `workflow/archive/umig-005-t-replicator-adapter.md`)
Next: UMIG-006 (QUEUED; only after ChatGPT accepts a real live VERIFY PASS)

## Outcome
Direct real-browser and actual sandbox Go runtime verification of Toolkit Replicator canonical load, destination suggest/inspect including foreign ownership conflict, explicit Save & Apply, source/Pull Block status, loss/recovery and restore, using ONLY a new disposable project and synthetic local source/destination. No fabricated COMMS health.

## Target and exact authority
Test-only `deploy/verify/compose.yaml` now adds a Replicator service sharing ONLY the private MMA2 network namespace and the disposable test volume, plus a canonical empty Replicator seed. These changes are CODE preparation, NOT executed or proven. The prior Memory test volume is retained and forbidden to reuse. The existing ephemeral OpenHands sandbox is the intended target, without a separate-VM precondition. Re-establish daemon/project ownership before startup and permit sandbox-local dockerd preparation. The sole binding commands/actions, expected results, evidence, report authority and cleanup are `## JR TEST TASK — CURRENT: UMIG-005-V LIVE VERIFY` in `handoff.md`.

## Acceptance
Real UI starts without Replicator fixture leakage or automatic write; synthetic Simulator source on private 15020/1 and Replicator destination on private 15021/1; real ownership inspection on simulator-owned 15020/1 reports IN USE and cannot claim it; one authoritative Go Replicator apply produces persisted canonical config and matching ownership plus real per-block source status. An unreachable synthetic loopback endpoint produces a truthful ERROR rather than fabricated green, recovery restores successful polling, and explicit deletion releases the Replicator destination. Inspect actual logs/config/browser evidence; safely tear down only project-owned containers/networks while retaining named test volume.

## Boundary
No production/customer endpoint or port, no real operator Docker or persistent volume, no host-published Modbus port, no unscheduled source changes, no Windows COMMS proof, no Diagnostics integration, no launcher cutover, no deletion of any volume. Unit/build PASS is not live VERIFY PASS.

## Sizing
Surface 0, environment 1, behavior 0, verification 1, recovery 1 = 3.
