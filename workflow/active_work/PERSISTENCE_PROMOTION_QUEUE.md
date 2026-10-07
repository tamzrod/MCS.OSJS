# Persistence autonomous promotion queue

Human-approved and promoted: PERSIST-001..022 on 2026-10-07.

All packets in this queue are authorized Active Work for their written scope. Promotion does NOT mean simultaneous execution. Repository rule remains: exactly ONE ACTIVE/current task per OPERATION CWAL invocation, selected by `handoff.md`.

Human explicitly selected the persistence workstream on 2026-10-07 and superseded OTR-001C as the current assignment. **PERSIST-001 is the sole current ACTIVE task, assigned to OpenHands JR DEV.** PERSIST-002..022 remain QUEUED and dependency-gated.

Ordered persistence sequence:

- [>] PERSIST-001 — Persistence Configuration Schema (CODE, ACTIVE — OpenHands JR DEV)
- [ ] PERSIST-002 — State Sealing Prerequisite Validation (CODE)
- [ ] PERSIST-003 — Derived Persistence RBE Generation (CODE)
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
This queue is selected. Execute PERSIST-001 now as OpenHands JR DEV. Advance only after genuine predecessor evidence is recorded. Run one packet per invocation and STOP. OpenHands JR DEV may implement assigned CODE/DISCOVERY packets but does not self-certify independent TEST/VERIFY. TEST/VERIFY requires a separate independent JR invocation with an exact current handoff packet and never fixes product source.

## Architectural invariants
- Persistence requires State Sealing.
- Port → Unit ID → Memory remains the identity/authority chain.
- Memory area start/count is the single source of truth.
- Persistence-owned RBE rules are derived, system-owned and locked.
- User RBE remains independently supported, including valid overlap.
- Persistence restore occurs while sealed; explicit successful final unseal is the commit point.
- Filesystem persistence is MCS appliance behavior, not a reason to make MMA2 memorycore own a historian/workflow engine.
- Only BLACK SHEEP WALL edits ICC.
