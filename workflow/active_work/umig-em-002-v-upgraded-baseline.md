# UMIG-EM-002-V — VERIFY: upgraded Go/Toolkit baseline in disposable stack

Status: ACTIVE — approved successor to independently reviewed `UMIG-EM-002-R` PASS on 2026-09-20.
Stage / owner: VERIFY / independent OpenHands JR; packet, evidence review and advancement / ChatGPT.
Previous: UMIG-EM-002-R (archived COMPLETE / TEST PASS; original transcript at `handoff.md` commit `0290e6b3bb29552c04fedbc6651d99b399a570d0`).
Next: UMIG-EM-003 (PLANNED CODE, not in the currently authorized queued sequence; separate human promotion required).

## Primary outcome
Independently verify that the upgraded, regression-tested Go/MMA2/Replicator baseline and existing OS.js Toolkit work together using ONE freshly owned, empty, isolated Linux Docker Compose project and real browser. Prior historical disposable stack PASS is not evidence for this upgraded backend.

## Three acceptance outcomes

1. Prove fresh checkout, local disposable daemon, unique empty project/volume, safe rendered `deploy/verify/compose.yaml`, free loopback UI port and ZERO host Modbus/RBE/access-event publication BEFORE ONE startup; fail closed when provenance/labels are ambiguous. Never use operator services, customer data, old volumes or production Compose.
2. In the new five-service stack, directly check `/healthz`, both sockets and seed config; use real Chromium Toolkit UI for ONE synthetic Memory source `127.0.0.1:15020` and ONE Replicator FC3 Pull Block, destination 15021/unit 1. Confirm canonical YAML + MMA2 owners, real source/poll statuses and Diagnostics truthfulness with no fabricated global health. Restore both via no more than FOUR total explicit Save & Apply operations (two each); no RBE output, source rewiring, unrelated tests or listener enablement.
3. Record pre/post config bytes/hashes and scoped, label-verified Compose `down --remove-orphans` WITHOUT `-v`, retained new test volume, no change to other resources and clean tracked tree. STOP on first product contradiction; never improvise restoration or repair.

## Exact authority and scope limit

The sole current `JR TEST TASK` in `handoff.md` supplies complete ordered commands, read-only ownership gates, one-shot startup, UI actions, bounded timeouts, cleanup, report evidence and report-only commit authority. JR operates ONLY through `operation cwal.md`, may report PASS/FAIL/BLOCKED, edits ONLY the authorized report, never changes product, workflow or ICC, and stops. This task's VERIFIED result will not mean advanced OS.js UI parity, multiwriter lock, privileged MMA manager, true four-layer COMMS, enabled RBE TCP, live customer systems, Electron deployment, visual parity or launcher cutover.

Sizing: implementation 0 / environment 2 / behavior 0 / verification 1 / recovery 1 = 4. Tight coupling is bounded by the existing test-only Compose topology, previously demonstrated synthetic workflow, one fixed project and four-apply maximum; this is one end-to-end compatibility workflow. If daemon provenance or ownership requires independent discovery, STOP BLOCKED and split a separately human-promoted environment-preflight task; no improvisation.
