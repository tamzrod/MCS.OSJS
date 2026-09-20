# Handoff: OSJT-009 Hydration projection

## Current authority

- Sole ACTIVE task: `OSJT-009` in `workflow/active_work/osjt-009-hydration-projection.md`.
- Stage: CODE. Owner: coding agent.
- Predecessor OSJT-008 independently PASSed at `644a50cc21e984b562e6add1d6bf9c600cbe1274`.
- Goal: project omitted policy, sealing, RBE, and unknown extensions from the matching effective
  port/unit memory while preserving explicit device values; malformed recognized projection data
  must return an error.

## Bounded scope and stop rule

- Authorized task paths: `simulator/advanced_projection.go`,
  `simulator/advanced_projection_test.go`, and `simulator/osjs_toolkit_settings_test.go`.
- No Electron, legacy removal, production/operator data, dependency upgrades, Docker, services,
  browser, external runtime, or unrelated cleanup.
- If integration requires editing a production file outside the task's named scope, or exposes a
  new architecture decision, stop and report the smallest task repair/split needed before editing.

## Source-only completion gate

- Inspect the focused diff and regression cases.
- Run only repository-authorized formatting or compile/test checks needed to validate authored code;
  record them as preliminary evidence, not independent TEST PASS.
- Record changed paths and create a source checkpoint only when the bounded implementation is coherent.
- Do not activate OSJT-010 until OSJT-009's source-only gate is satisfied and its exact successor
  linkage is confirmed.
