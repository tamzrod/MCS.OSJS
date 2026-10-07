# PERSIST-010 — Persistence Configuration UI

Status: CODE COMPLETE — 2026-10-07 (OpenHands JR DEV); awaiting separate independent TEST/VERIFY (PERSIST-011)
Stage: CODE
Owner: OpenHands JR DEV
Previous: PERSIST-009
Next: PERSIST-011

## Primary outcome
Expose persistence enablement and locked derived RBE state in the existing memory Advanced Settings UI.

## Scope
Implement or verify only the primary outcome above against the current authoritative MCS.OSJS source and existing MMA2 contracts. Preserve the Port → Unit ID → Memory authority model and existing State Sealing/RBE semantics unless this packet explicitly says otherwise.

## Non-scope
Do not implement disk/runtime status UI yet.

## Acceptance
1. User can enable/disable persistence only through supported draft/apply flow.
2. UI rejects persistence ON when State Sealing is disabled and explains why.
3. Derived persistence RBE entries display as locked/system-owned while user RBE remains editable.

## Evidence / handoff
Coding owner records the exact source diff, changed paths, targeted self-checks and resulting source checkpoint. Do not claim independent TEST/VERIFY PASS. STOP after this task.

## Dependencies
Requires genuine completion evidence for PERSIST-009. Being present in `workflow/active_work/` authorizes this written scope but does not bypass predecessor or current-task selection gates.

## Sizing
2/0/2/1/1=6; UI only.

## Coding evidence (PERSIST-010, OpenHands JR DEV)

- Source checkpoint base: `4c53efbe4b75269ec3c4245085741d68fe68ca99` (`main`, clean, = `origin/main` at run).
- Changed/added paths (product):
  - M `OSJS/src/packages/MCSModbusToolkit/memory-advanced.js` (blob `723a7cdcd5776af46ba18671f0becd765dd40f4d`) — adds a **Persistence** tab to the existing Advanced Settings UI: an "Enable persistence" control (draft/apply via the same `params` binding), an inline explanation when persistence is on without enabled state sealing, and a read-only table of layout-derived system-owned ("System-owned (locked)") rules. Adds helpers `sealingEnabled`, `persistenceEnabled`, `persistenceError`, `derivedPersistenceRules`; `replicatorParams` now binds `persistence`.
  - A `OSJS/tests/toolkit-persistence-ui.test.js` (blob `7ba7d4589523166c2b9a2f921a6b3521bcfff09d`) — UI self-check.
- Self-check: `cd OSJS && node tests/toolkit-persistence-ui.test.js` → PASS, exit 0.
- Bounded regression: all 16 `OSJS/tests/*.test.js` → PASS (including `toolkit-ui-parity` and `toolkit-fc43`, which exercise the same editor). `node --check` on the module → OK.
- Acceptance mapping: (1) persistence enable/disable only through the draft/apply control (mutates `params.persistence`, the editor's draft model); (2) persistence ON without sealing shows an operator-facing reason (`persistenceError` + `role="alert"`); (3) derived persistence RBE entries render as "System-owned (locked)" with no editable inputs, while user RBE remains editable on the RBE Rules tab.
- Non-scope preserved: no disk/runtime status UI.
- Test dependency note: `jsdom` (already an OSJS devDependency) was installed locally via `npm install --no-save` to run the tests; `OSJS/package.json` is unchanged and `node_modules` is git-ignored.
- This is a JR DEV self-check only. It is **not** an independent TEST/VERIFY PASS; PERSIST-011 remains the separate independent JR gate.
- Delivered source commit: `190e464f6106c3e640d21b1f9352328447464bef`.

## CWAL
PERSIST-009 is delivered on GitHub main at `19057f02d94ff404b74a9f01348fc8c31235e374`. This packet is the sole current ACTIVE assignment for OpenHands JR DEV. Execute exactly this task, deliver it under standing JR DEV authority, prepare PERSIST-011 for the next invocation as an independent JR TEST task, and STOP.
