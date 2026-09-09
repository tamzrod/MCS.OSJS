# SIM-023C — Render Neutral Status for Unapplied Selection

Status: COMPLETED 2026-09-09 — neutral status placeholders for rows without an applied runtime identity(new/renamed-unapplied; genuine request failures stay UNAVAILABLE via the existing error path.
Previous: SIM-023B

## Primary outcome

Keep the status row truthful when the selected editor entry has no applied runtime identity to query.

## Scope

- When the selected working-copy device has no applicable persisted runtime target, render neutral status placeholders instead of `UNAVAILABLE` for MMA2 and Simulator.
- Keep actual runtime/relay request failures visibly `UNAVAILABLE` through the existing error path.
- Preserve the existing bottom status bar message for local states such as new unsaved device, edited device, Saving, apply failure, and apply success.
- Ensure a later successful runtime-status response replaces the neutral placeholders normally.

## Non-scope

- No new runtime status enum values such as `APPLYING` or `NOT APPLIED`.
- No backend status-contract changes.
- No polling target or save lifecycle changes; SIM-023A and SIM-023B own those behaviors.
- No layout or CSS redesign beyond what is necessary to keep the current row truthful.

## Acceptance criteria

1. A newly added unsaved device does not display MMA2 or Simulator as `UNAVAILABLE` solely because no runtime target exists yet.
2. A renamed-but-unapplied working copy likewise shows neutral placeholders rather than fabricated unavailability.
3. A genuine runtime status request failure still displays `UNAVAILABLE` and its diagnostic through the existing message path.

## Verification

Use focused rendering/state tests for neutral-without-target versus unavailable-on-request-failure, plus the smallest relevant JavaScript syntax/package build check.

## Dependencies

- SIM-023B.
- Existing SIM-021D compact runtime-status row and SIM-021E status-error presentation.

## Sizing

Implementation 1, environment 0, behavioral 1, verification 1, decision/recovery 0 = 3. One presentation-state distinction with one deterministic verification workflow.
