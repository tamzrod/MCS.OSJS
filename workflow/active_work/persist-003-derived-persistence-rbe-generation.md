# PERSIST-003 — Derived Persistence RBE Generation

Status: CODE COMPLETE — 2026-10-07 (OpenHands JR DEV); awaiting separate independent TEST/VERIFY (PERSIST-011)
Stage: CODE
Owner: OpenHands JR DEV
Previous: PERSIST-002
Next: PERSIST-004

## Primary outcome
Generate one system-owned persistence RBE rule per configured memory area from the authoritative memory layout.

## Scope
Implement or verify only the primary outcome above against the current authoritative MCS.OSJS source and existing MMA2 contracts. Preserve the Port → Unit ID → Memory authority model and existing State Sealing/RBE semantics unless this packet explicitly says otherwise.

## Non-scope
Do not implement UI locking, user-rule collision policy, or disk persistence.

## Acceptance
1. Each present area gets exactly one derived persistence RBE.
2. Derived start/count equal the area's current start/count.
3. Persistence does not introduce independent editable range fields.

## Evidence / handoff
Coding owner records the exact source diff, changed paths, targeted self-checks and resulting source checkpoint. Do not claim independent TEST/VERIFY PASS. STOP after this task.

## Dependencies
Requires genuine completion evidence for PERSIST-002. Being present in `workflow/active_work/` authorizes this written scope but does not bypass predecessor or current-task selection gates.

## Sizing
2/0/2/1/1=6; bounded to derivation only.

## Coding evidence (PERSIST-003, OpenHands JR DEV)

- Source checkpoint base: `9c01bb8c0dc88ee270a6d5d393c8cbfd6c56997b` (`main`, clean, = `origin/main` at run).
- Changed/added paths (product):
  - A `mma2composer/persistence_rbe.go` — `PersistenceRBERule{Area,Start,Count}` + `DerivePersistenceRBE(memory Memory) []PersistenceRBERule`. One rule per present area in canonical order (coils, discrete_inputs, holding_registers, input_registers); start/count copied from `Area`; count-0/nil areas omitted; no ID/name/lock/range fields.
  - A `mma2composer/persistence_rbe_test.go` — focused self-check (one-per-present-area, range tracks layout, empty when no areas).
- Self-check: `cd mma2composer && go test -run TestDerivePersistenceRBE .` → `ok`, exit 0 (3 subtests).
- Bounded regression (`-mod=readonly`, no manifest mutation): full `mma2composer`, `simulator`, `replicator` suites → all `ok`, exit 0. `gofmt -l` clean; `go vet` exit 0.
- Acceptance mapping: (1) exactly one rule per present area; (2) derived start/count equal the area start/count (asserted, incl. after a layout change); (3) no independent editable range — the rule carries only a copy and the helper reads straight from the area.
- Deferred by design (out of scope here): global RBE ID assignment, system-owned/locked representation, user-rule collision policy, filesystem persistence, restore.
- Invariant lives on the shared schema owner (`mma2composer`), applying to Simulator and Replicator alike.
- This is a JR DEV self-check only. It is **not** an independent TEST/VERIFY PASS; PERSIST-011 remains the separate independent JR gate.
- Delivered source commit: `4e8f774eb946cc9a946938e4f3a28eaa13de5c9d`.

## CWAL
PERSIST-002 is delivered on GitHub main at `fe3bda9e5c863c959500514e8c6083d2514526d1`. This packet is the sole current ACTIVE assignment for OpenHands JR DEV. Execute exactly this task, deliver it under standing JR DEV authority, prepare the next already-promoted eligible packet for a later invocation, and STOP.
