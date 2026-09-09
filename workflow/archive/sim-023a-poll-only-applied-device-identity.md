# SIM-023A — Poll Only the Applied Device Identity

Status: COMPLETED 2026-09-09 — poll target resolved from persisted/applied identity; new/renamed-unapplied rows expose no poll target.
Previous: none
Next: SIM-023B

## Primary outcome

Make selected-device runtime-status polling target only an identity that exists in the persisted/applied Simulator document.

## Scope

- Resolve the status poll target from the persisted/applied selected device identity rather than directly from the editable working-copy name.
- For a newly added unsaved device, expose no runtime poll target and issue no `status` request for that local-only name.
- When an existing device name is edited but not yet applied, continue to associate runtime polling with its last persisted identity instead of the edited name.
- Keep ordinary selection changes among already persisted devices polling their corresponding applied names.
- Add focused client/poller coverage for new-device and rename-before-apply cases.

## Non-scope

- No changes to backend `RuntimeStatus` lookup semantics.
- No Save & Apply polling pause/resume behavior; SIM-023B owns that lifecycle.
- No status-row wording or placeholder changes; SIM-023C owns presentation.
- No MMA2 lifecycle/control changes.

## Acceptance criteria

1. Selecting a newly added unsaved device issues no runtime `status` request for its local-only name.
2. Editing the name of a persisted selected device does not cause status polling to switch to the unapplied edited name.
3. Selecting an unchanged persisted device continues to poll its applied device name normally.

## Verification

Use a focused client/poller test with a stubbed runtime requester that records requested names, plus the smallest relevant JavaScript syntax/package build check.

## Dependencies

- Existing SIM-020 persisted-versus-working-copy client state.
- Existing SIM-021C selected-device status poller.

## Sizing

Implementation 1, environment 0, behavioral 1, verification 1, decision/recovery 0 = 3. One bounded identity-selection behavior inside the existing Modbus Simulator client branch.
