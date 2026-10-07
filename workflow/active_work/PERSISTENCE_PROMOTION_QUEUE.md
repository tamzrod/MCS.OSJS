# Persistence autonomous promotion queue

Human-approved and promoted: PERSIST-001..022 on 2026-10-07.

All packets in this queue are authorized Active Work for their written scope. Promotion does NOT mean simultaneous execution. Repository rule remains: exactly ONE ACTIVE/current task per OPERATION CWAL invocation, selected by `handoff.md`.

Human explicitly selected the persistence workstream on 2026-10-07. **PERSIST-001 is CODE COMPLETE at `faa33929429a0382a64b78bc773f0074e11cb6b1`. PERSIST-002 is CODE COMPLETE and delivered at `fe3bda9e5c863c959500514e8c6083d2514526d1`. PERSIST-003 is now the sole current ACTIVE task, assigned to OpenHands JR DEV.** PERSIST-004..022 remain QUEUED and dependency-gated.

Ordered persistence sequence:

- [x] PERSIST-001 — Persistence Configuration Schema (CODE COMPLETE — `faa3392`; independent TEST/VERIFY deferred to PERSIST-011)
- [x] PERSIST-002 — State Sealing Prerequisite Validation (CODE COMPLETE — `fe3bda9`; independent TEST/VERIFY deferred to PERSIST-011)
- [>] PERSIST-003 — Derived Persistence RBE Generation (CODE, ACTIVE — OpenHands JR DEV)
- [ ] PERSIST-004 — Locked System RBE Behavior (CODE)
- [ ] PERSIST-005 — Memory Range Synchronization (CODE)
- [ ] PERSIST-006 — Memory Area Removal Synchronization (CODE)
- [ ] PERSIST-007 — Persistence Disable Cleanup (CODE)
- [ ] PERSIST-008 — User RBE Compatibility (CODE)
- [ ] PERSIST-009 — RBE ID Collision Handling (CODE)
- [ ] PERSIST-010 — Persistence Configuration UI (CODE)
- [ ] PERSIST-011 — Persistence Configuration Tests (TEST)
- [ ] PERSIST-012 — RBE-Triggered Snapshot Writer (CODE)
- [ ] PERSIST-013 — Snapshot File Format (CODE)
- [ ] PERSIST-014 — Snapshot Manifest / Compatibility Metadata (CODE)
- [ ] PERSIST-015 — Startup Snapshot Loader (CODE)
- [ ] PERSIST-016 — Raw Ingest Restore (CODE)
- [ ] PERSIST-017 — Seal-Flag Protection During Restore (CODE)
- [ ] PERSIST-018 — Restore Verification (CODE)
- [ ] PERSIST-019 — Atomic Unseal / Commit Step (CODE)
- [ ] PERSIST-020 — Restore Failure Behavior (CODE)
- [ ] PERSIST-021 — Persistence Runtime Status (CODE)
- [ ] PERSIST-022 — End-to-End Persistence Verification (VERIFY)

## Routing
This queue is selected. PERSIST-001 has recorded CODE completion evidence. PERSIST-002 repository delivery is confirmed at `fe3bda9`. Execute PERSIST-003 now as OpenHands JR DEV. Run one packet per invocation and STOP. OpenHands JR DEV may implement assigned CODE/DISCOVERY packets but does not self-certify independent TEST/VERIFY. TEST/VERIFY requires a separate independent JR invocation with an exact current handoff packet and never fixes product source.

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
