# Native persistence integration queue — realigned 2026-10-08

Human has realigned the persistence migration into TWO sequential product phases.

CURRENT: **NPE-06W — Windows Electron Native Persistence Acceptance**

Order:

```text
NPE-01  MCS/native contract realignment
NPE-02  Electron persistence UI
NPE-03  Electron native config / Save & Apply wiring
NPE-04  Electron Replicator destination persistence parity
NPE-05  Electron legacy persistence cleanup
NPE-05A Refresh embedded MMA2 native persistence runtime
NPE-06  Electron independent native-persistence verification — BLOCKED on headless Linux
NPE-06W Windows Electron native-persistence acceptance
        ↓ ONLY AFTER PASS
NPO-01  OS.js Toolkit persistence UI
NPO-02  OS.js Toolkit native config wiring
NPO-03  OS.js Toolkit legacy cleanup / parity
NPO-04  OS.js Toolkit independent native-persistence verification
        ↓ ONLY AFTER PASS
NPF-01  workflow/documentation cleanup
```

## Routing rules

- One packet per OPERATION CWAL invocation unless a later handoff explicitly changes that rule.
- NPE-06 remains blocked until NPE-05A refreshes the stale embedded MMA2 donor copy.
- Electron is completed and independently verified BEFORE OS.js Toolkit persistence work starts.
- MMA2 owns persistence runtime: disk snapshot, runtime flush, startup restore, backup and recovery.
- Integration layers configure native per-memory persistence only.
- Do not reintroduce persistence dependence on State Sealing, RBE TCP, Raw Ingest restore, an unlock coil, external snapshot capture, or an external persistence watchdog.
- The previously promoted NP-02..NP-06 packets are SUPERSEDED and must not execute.
- Historical PERSIST-* packets remain evidence only and must not be reactivated.
- Do not delete user snapshot data.

## Completed predecessor

NP-01 historical schema work was delivered at `757fe00b6e8b057b92cb693781fe1885ce362dd6`, but its contract is partially stale versus current MMA2: notably it required a directory and retained persistence/State-Sealing assumptions. NPE-01 exists specifically to reconcile those assumptions before Electron UI implementation.


---

# Persistence autonomous promotion queue

Human-approved and promoted: PERSIST-001..022 on 2026-10-07; optimized continuation PERSIST-R01..R05 + PERSIST-UI01 promoted on 2026-10-08.

All packets in this queue are authorized Active Work for their written scope. Promotion does NOT mean simultaneous execution. Repository rule remains: exactly ONE ACTIVE/current task per OPERATION CWAL invocation, selected by `handoff.md`.

Human explicitly selected the persistence workstream on 2026-10-07. **PERSIST-001..010 CODE tasks are complete and delivered through `190e464f6106c3e640d21b1f9352328447464bef`. PERSIST-011 independently PASSed on tested HEAD `313e3be60a29fff6c967d78c0cf677b8b888ed56`. PERSIST-012 is CODE COMPLETE and delivered at `2f18f8196f5497c5435b697b8b1570f59cab2642`. PERSIST-013 is CODE COMPLETE and delivered at `ccccb891b737995f6a24a58e844e5bd1d794b609`. PERSIST-014 is CODE COMPLETE and delivered at `98a952a89b087f12a9b99b6bd648ce97ac44e6f2` via connector recovery after transport failure. PERSIST-015 is CODE COMPLETE and delivered at `a497d9d4ca35cab8de832bc92fc3b5d90c7a4b33` (OpenHands JR DEV; independent TEST/VERIFY deferred to PERSIST-022). PERSIST-016 is CODE COMPLETE and delivered at `2fbabc3481b6435ca38596ea7dd060e36ed0655d` (OpenHands JR DEV; independent TEST/VERIFY deferred to PERSIST-022). PERSIST-017 is CODE COMPLETE and delivered at `fa1276c26cbb639f061436c3f111351232037b57` (OpenHands JR DEV; independent TEST/VERIFY deferred to PERSIST-022). PERSIST-018 is CODE COMPLETE and delivered at `3eb118f507688d483c5d4985f74fe6a98b531bc3` (OpenHands JR DEV; independent TEST/VERIFY deferred to PERSIST-022). PERSIST-019 is CODE COMPLETE and delivered at `e22459ecdd10433ef21f63df25e3a22a43fb3662` (OpenHands JR DEV; independent TEST/VERIFY deferred to PERSIST-022). PERSIST-020 is CODE COMPLETE and delivered at `1528d3e2451f3588de46f7539127164588878f1a` (OpenHands JR DEV; independent TEST/VERIFY deferred to PERSIST-022). PERSIST-021 is CODE COMPLETE and delivered at `1310c1367c7f98f084187006084d624d8e20e07b` (OpenHands JR DEV). PERSIST-001..021 CODE tasks are complete. The runtime-integration gap was decomposed and human-promoted as PERSIST-R01..R05 + PERSIST-UI01. PERSIST-R01 is CODE COMPLETE and delivered at `1ac1477177ec19fa02cb344d6a20a8632fc0c88a` (OpenHands JR DEV). PERSIST-R02 is CODE COMPLETE and delivered at `407fcc32979b9acd83dd6826e86f4b700da70c1b` (OpenHands JR DEV). PERSIST-R03 is CODE COMPLETE and delivered at `2236e15277ebb17c60cac206e05d77e17e296a53` (OpenHands JR DEV). PERSIST-UI01 is CODE COMPLETE and delivered at `cff34f41ebb350fc1c4380253cc2acc02e4562ca` (OpenHands JR DEV). PERSIST-R04 is CODE COMPLETE and delivered at `d9251d3c56bf490c21655bdb372b342ee1eec7f8` (OpenHands JR DEV). PERSIST-R05 is CODE COMPLETE and delivered at `c8f31b2443b69201f36464db72497ed8271aeb7d` (OpenHands JR DEV). All CODE tasks are complete; PERSIST-022 independently PASSed on tested HEAD `6175a77688a85e4aec4706c9e01079fa235f7963` against pinned product checkpoint `c8f31b2443b69201f36464db72497ed8271aeb7d`. The persistence workstream is COMPLETE.**

Ordered persistence sequence:

- [x] PERSIST-001 — Persistence Configuration Schema (CODE COMPLETE — `faa3392`; independent TEST/VERIFY deferred to PERSIST-011)
- [x] PERSIST-002 — State Sealing Prerequisite Validation (CODE COMPLETE — `fe3bda9`; independent TEST/VERIFY deferred to PERSIST-011)
- [x] PERSIST-003 — Derived Persistence RBE Generation (CODE COMPLETE — `4e8f774`; independent TEST/VERIFY deferred to PERSIST-011)
- [x] PERSIST-004 — Locked System RBE Behavior (CODE COMPLETE — `936ab24`; independent TEST/VERIFY deferred to PERSIST-011)
- [x] PERSIST-005 — Memory Range Synchronization (CODE COMPLETE — `c743174`; independent TEST/VERIFY deferred to PERSIST-011)
- [x] PERSIST-006 — Memory Area Removal Synchronization (CODE COMPLETE — `9ac48cb`; independent TEST/VERIFY deferred to PERSIST-011)
- [x] PERSIST-007 — Persistence Disable Cleanup (CODE COMPLETE — `feb1e62`; independent TEST/VERIFY deferred to PERSIST-011)
- [x] PERSIST-008 — User RBE Compatibility (CODE COMPLETE — `4783e0d`; independent TEST/VERIFY deferred to PERSIST-011)
- [x] PERSIST-009 — RBE ID Collision Handling (CODE COMPLETE — `19057f0`; independent TEST/VERIFY deferred to PERSIST-011)
- [x] PERSIST-010 — Persistence Configuration UI (CODE COMPLETE — `190e464`; independent TEST/VERIFY deferred to PERSIST-011)
- [x] PERSIST-011 — Persistence Configuration Tests (TEST PASS — independent JR on `313e3be`; pinned product checkpoint `190e464`)
- [x] PERSIST-012 — RBE-Triggered Snapshot Writer (CODE COMPLETE — `2f18f81`; independent TEST/VERIFY deferred to PERSIST-022)
- [x] PERSIST-013 — Snapshot File Format (CODE COMPLETE — `ccccb89`; recovered/delivered via GitHub connector; independent TEST/VERIFY deferred to PERSIST-022)
- [x] PERSIST-014 — Snapshot Manifest / Compatibility Metadata (CODE COMPLETE — `98a952a`; recovered/delivered via GitHub connector; independent TEST/VERIFY deferred to PERSIST-022)
- [x] PERSIST-015 — Startup Snapshot Loader (CODE COMPLETE — `a497d9d`; OpenHands JR DEV; independent TEST/VERIFY deferred to PERSIST-022)
- [x] PERSIST-016 — Raw Ingest Restore (CODE COMPLETE — `2fbabc3`; OpenHands JR DEV; independent TEST/VERIFY deferred to PERSIST-022)
- [x] PERSIST-017 — Seal-Flag Protection During Restore (CODE COMPLETE — `fa1276c`; OpenHands JR DEV; independent TEST/VERIFY deferred to PERSIST-022)
- [x] PERSIST-018 — Restore Verification (CODE COMPLETE — `3eb118f`; OpenHands JR DEV; independent TEST/VERIFY deferred to PERSIST-022)
- [x] PERSIST-019 — Atomic Unseal / Commit Step (CODE COMPLETE — `e22459e`; OpenHands JR DEV; independent TEST/VERIFY deferred to PERSIST-022)
- [x] PERSIST-020 — Restore Failure Behavior (CODE COMPLETE — `1528d3e`; OpenHands JR DEV; independent TEST/VERIFY deferred to PERSIST-022)
- [x] PERSIST-021 — Persistence Runtime Status (CODE COMPLETE — `1310c13`; OpenHands JR DEV; independent TEST/VERIFY deferred to PERSIST-022)
- [x] PERSIST-R01 — Filesystem Snapshot Adapter (CODE COMPLETE — `1ac1477`; OpenHands JR DEV)
- [x] PERSIST-R02 — Runtime Save Wiring (CODE COMPLETE — `407fcc3`; OpenHands JR DEV)
- [x] PERSIST-R03 — Startup Restore Wiring (CODE COMPLETE — `2236e15`; OpenHands JR DEV)
- [x] PERSIST-UI01 — Electron Persistence Settings (CODE COMPLETE — `cff34f4`; OpenHands JR DEV)
- [x] PERSIST-R04 — Runtime Status Wiring (CODE COMPLETE — `d9251d3`; OpenHands JR DEV)
- [x] PERSIST-R05 — Disposable Persistence E2E Harness (CODE COMPLETE — `c8f31b2`; OpenHands JR DEV)
- [x] PERSIST-022 — End-to-End Persistence Verification (VERIFY PASS — independent JR on tested HEAD `6175a776`; pinned product checkpoint `c8f31b2`)

## Routing
This queue is selected. PERSIST-UI01 CODE is complete and delivered by this invocation at `cff34f4`. All CODE tasks are complete and PERSIST-022 independently PASSed. The persistence workstream is COMPLETE; no persistence task remains ACTIVE.

Autonomous successor chain is mandatory after genuine delivery:
`R01 → R02 → R03 → UI01 → R04 → R05 → PERSIST-022 independent JR`.

Do not request a second human promotion or task-selection confirmation for R02..R05/UI01: they are already human-promoted. **Recursive JR DEV continuation is enabled**: after each successful CODE delivery, record/push it, activate the named successor, perform the context-reset checkpoint, and continue within the same OPERATION CWAL invocation. Do not skip a blocked/failed task. Stop before executing PERSIST-022 because VERIFY requires independent JR identity.

## Architectural invariants
- Persistence requires State Sealing.
- Port → Unit ID → Memory remains the identity/authority chain.
- Memory area start/count is the single source of truth.
- Persistence-owned RBE rules are derived, system-owned and locked.
- User RBE remains independently supported, including valid overlap.
- Persistence restore occurs while sealed; explicit successful final unseal is the commit point.
- Filesystem persistence is MCS appliance behavior, not a reason to make MMA2 memorycore own a historian/workflow engine.
- Only BLACK SHEEP WALL edits ICC.


## Standing JR DEV continuation

For this already human-promoted queue, OpenHands JR DEV may autonomously:
- finish and deliver the current CODE/DISCOVERY packet;
- perform bounded safe Git sync/rebase when remote advances without overlapping product semantics;
- commit and non-force push the current task;
- record genuine completion evidence;
- mark the current packet complete;
- assign the next eligible already-promoted CODE/DISCOVERY packet to OpenHands JR DEV for the **next invocation**.

It may recursively execute successive already-promoted CODE/DISCOVERY packets in one OPERATION CWAL invocation when handoff explicitly enables recursive continuation. Each task must still be independently delivered/recorded before the next begins, with a context-reset checkpoint between tasks. TEST/VERIFY packets switch to independent JR mode and terminate JR DEV recursion; JR DEV does not certify them.
