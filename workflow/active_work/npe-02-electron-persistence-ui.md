# NPE-02 — Electron Persistence UI

Status: PROMOTED / QUEUED
Stage: CODE
Owner: Codex JR DEV
Previous: NPE-01
Next: NPE-03

## Purpose
Implement the native MMA2 persistence configuration experience in the Electron desktop first.

## Scope
In Electron Memory Advanced Settings, expose only native persistence configuration:

- Enabled
- Storage location:
  - default/native location when directory omitted
  - optional custom directory override
- Persisted memory:
  - default = All Allocated Areas
  - optional Selected Ranges
- selected ranges for coils, discrete inputs, holding registers and input registers, bounded by each allocated area.

Persistence UI must NOT expose:
- persistence RBE
- persistence lock coil
- persistence State Sealing dependency
- restore/unseal controls
- external snapshot manager controls

State Sealing remains a separate independent tab/feature.

## UX rule
The simple default is:

```
Persistence
[✓] Enabled
Persist: All Allocated Areas
Directory: Default
```

Custom ranges/directory are advanced overrides, not required fields.

## Acceptance
Focused Electron renderer tests cover default state, enable/disable, optional directory, whole-memory default, custom range editing, validation, save/reload and independence from State Sealing/RBE.

No OS.js Toolkit changes in this packet. STOP after delivery.
