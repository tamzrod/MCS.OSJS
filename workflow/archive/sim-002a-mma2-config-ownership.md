# SIM-002A — Compose MMA2 Configuration with Ownership Protection

Status: DONE — completed+verified 2026-09-07 (CWAL).

Source intent: split from `workflow/active_work/sim-002-mma2-activation.md` after its mandatory split gate was reached.

## Completion evidence (recorded by CWAL 2026-09-07)

- Composition boundary: repository-root `simulator/` composes the effective MMA2 config under `$OSJS_DATA_DIR/config/mma2/config.yaml` with ownership registry at `$OSJS_DATA_DIR/config/mma2/owners.yaml`;both written atomically via temp+rename.
- `(port, unit_id)` is the unique reservation key;every persisted effective MMA2 reservation carries a machine-readable YAML `owner` entry (owners.yaml), first-come-first-save。
- Ownership enforcement:free key → simulator may save/delete;existing key + same owner → update/delete;existing key + different owner → `ErrReservationOwnedByOther` rejected before any write, leaving prior effective config和 ownership registry byte-unchanged。
- Foreign preservation:simulator compose/drop operates reservation-locally, merging into an existing listener on the same port,never touching other listeners'memory或 ownership entries。
- Verification command:`cd simulator && go test -count=1 -v ./...` (Go 1.22.2)。 Six SIM-002A tests PASS (free-save persists+preserves foreign;collision reject unchanged;own-update preserves foreign;delete removes own only;foreign-delete reject;invalid-save no-persist)。 Also caught+fixed one implementation bug discovered during verification:FC4 input-registers mapping used FC3 start/count (copy-paste);fixed to FC4,and fixture/assertion strengthened (FC4 count 105) so that mapping is covered distinctly。

## Primary Outcome

A simulator MMA2 structural request can be composed into the effective MMA2 configuration without overwriting or deleting a `(port, unit_id)` resource already owned by another program.

## Scope

- Treat `(port, unit_id)` as the unique MMA2 reservation key shared by Simulator and Replicator.
- Persist a machine-readable YAML `owner` entry with every effective MMA2 reservation.
- Ownership is first-come-first-save: the first successfully persisted owner of a `(port, unit_id)` reservation retains authority over that reservation until it releases/deletes it.
- Allow creation when `(port, unit_id)` is free.
- Allow modification/deletion only when the existing reservation is owned by the requesting program.
- Reject creation, modification, or deletion that would overwrite or remove a reservation owned by another program.
- Preserve all foreign-owned MMA2 entries unchanged while composing simulator-owned changes.
- Reject a conflicting save before replacing the currently effective MMA2 configuration.
- Use the actual MMA2 configuration/addressing representation in repository truth; do not redesign MMA2 addressing semantics.

## Ownership Invariant

```text
UNIQUE RESOURCE KEY = (port, unit_id)

free key
→ requester may save and become owner

existing key + same owner
→ requester may update/delete its own reservation

existing key + different owner
→ overwrite prohibited
→ delete prohibited
→ reject save
```

`owner` is configuration data, not a comment. The exact YAML placement/value shape must follow the effective MMA2 schema established during implementation, but it must be machine-readable and sufficient to enforce ownership.

## Non-Scope

- No MMA2 restart/reload lifecycle implementation.
- No proof of live Modbus exposure after activation; that belongs to SIM-002B.
- No random-runtime scheduler.
- No raw-ingest data generation.
- No OS.js simulator UI.
- No Replicator implementation beyond preserving and enforcing its ownership boundary.

## Acceptance Criteria

1. A free `(port, unit_id)` simulator request is persisted with a machine-readable YAML `owner` entry.
2. A request against a `(port, unit_id)` owned by another program is rejected without altering that foreign-owned entry.
3. A simulator-owned reservation can be updated without modifying unrelated or foreign-owned reservations.
4. The resulting effective MMA2 configuration preserves ownership metadata and all unaffected entries.

## Verification

Use an effective MMA2 configuration containing at least one foreign-owned reservation. Save one free simulator reservation and verify ownership is persisted. Attempt one deliberate collision against the foreign-owned `(port, unit_id)` and prove the save is rejected with the prior effective configuration unchanged. Update one simulator-owned reservation and prove foreign entries remain byte-for-byte or semantically unchanged as appropriate to the serializer.

## Dependencies

- SIM-001.
- Existing MMA2 configuration representation.

## Sizing

**3 / 10 — Good JR task.** One deterministic configuration-composition and ownership-enforcement outcome; lifecycle activation is split into SIM-002B.
