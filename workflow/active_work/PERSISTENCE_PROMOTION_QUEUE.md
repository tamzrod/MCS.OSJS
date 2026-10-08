# Persistence autonomous promotion queue

Human-approved and promoted: PERSIST-001..022 on 2026-10-07.

All packets in this queue are authorized Active Work for their written scope. Promotion does NOT mean simultaneous execution. Repository rule remains: exactly ONE ACTIVE/current task per OPERATION CWAL invocation, selected by `handoff.md`.

Human explicitly selected the persistence workstream on 2026-10-07. **PERSIST-001..010 CODE tasks are complete and delivered through `190e464f6106c3e640d21b1f9352328447464bef`. PERSIST-011 independently PASSed on tested HEAD `313e3be60a29fff6c967d78c0cf677b8b888ed56`. PERSIST-012 is CODE COMPLETE and delivered at `2f18f8196f5497c5435b697b8b1570f59cab2642`. PERSIST-013 is CODE COMPLETE and delivered at `ccccb891b737995f6a24a58e844e5bd1d794b609`. PERSIST-014 is CODE COMPLETE and delivered at `98a952a89b087f12a9b99b6bd648ce97ac44e6f2` via connector recovery after transport failure. PERSIST-015 is CODE COMPLETE and delivered at `a497d9d4ca35cab8de832bc92fc3b5d90c7a4b33` (OpenHands JR DEV; independent TEST/VERIFY deferred to PERSIST-022). PERSIST-016 is CODE COMPLETE and delivered at `2fbabc3481b6435ca38596ea7dd060e36ed0655d` (OpenHands JR DEV; independent TEST/VERIFY deferred to PERSIST-022). PERSIST-017 is CODE COMPLETE and delivered at `fa1276c26cbb639f061436c3f111351232037b57` (OpenHands JR DEV; independent TEST/VERIFY deferred to PERSIST-022). PERSIST-018 is CODE COMPLETE and delivered at `3eb118f507688d483c5d4985f74fe6a98b531bc3` (OpenHands JR DEV; independent TEST/VERIFY deferred to PERSIST-022). PERSIST-019 is CODE COMPLETE and delivered at `e22459ecdd10433ef21f63df25e3a22a43fb3662` (OpenHands JR DEV; independent TEST/VERIFY deferred to PERSIST-022). PERSIST-020 is CODE COMPLETE and delivered at `1528d3e2451f3588de46f7539127164588878f1a` (OpenHands JR DEV; independent TEST/VERIFY deferred to PERSIST-022). PERSIST-021 is the sole current ACTIVE task, assigned to OpenHands JR DEV.** PERSIST-022 remains QUEUED and dependency-gated.

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
- [>] PERSIST-021 — Persistence Runtime Status (CODE, ACTIVE — OpenHands JR DEV)
- [ ] PERSIST-022 — End-to-End Persistence Verification (VERIFY)

## Routing
This queue is selected. PERSIST-020 delivery is confirmed at `1528d3e2451f3588de46f7539127164588878f1a`. Execute PERSIST-021 now as OpenHands JR DEV. Run one packet per invocation and STOP.

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

It must still execute only one product task per invocation and STOP after preparing the successor. TEST/VERIFY packets switch to independent JR mode and require their exact packet; JR DEV does not certify them.
