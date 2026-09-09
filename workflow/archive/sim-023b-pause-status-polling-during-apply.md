# SIM-023B — Pause Status Polling During Save & Apply

Status: COMPLETED 2026-09-09 — polling paused via select(null,) during Save & Apply; in-flight responses dropped; forced refresh resumes after the apply settles.
Previous: SIM-023A
Next: SIM-023C

## Primary outcome

Prevent selected-device runtime-status polling from racing a Save & Apply transaction.

## Scope

- While `state.saving` is true, do not issue selected-device `status` requests.
- Preserve the last applied runtime status while the apply request is in flight; do not replace it with an unavailable result caused only by the apply transition.
- After a successful apply finishes and the persisted document has been replaced, force one immediate status refresh for the applied selected device.
- After a failed apply finishes, resume polling the previously applied target without treating the failed edited identity as runtime truth.

## Non-scope

- No backend runtime-service changes.
- No MMA2 lifecycle/control changes.
- No new status vocabulary or visual redesign.
- No changes to device add/rename target selection; SIM-023A owns applied-identity selection.

## Acceptance criteria

1. No `status` request is started while Save & Apply is in progress.
2. A successful apply resumes polling and immediately refreshes the newly applied selected device.
3. A failed apply resumes polling the prior applied target and does not poll an unapplied edited identity.

## Verification

Use a focused client/poller test with a controllable apply-in-progress flag and stubbed runtime requester, plus the smallest relevant JavaScript syntax/package build check.

## Dependencies

- SIM-023A.
- Existing SIM-020 Save & Apply flow and SIM-021C status poller.

## Sizing

Implementation 1, environment 0, behavioral 1, verification 1, decision/recovery 0 = 3. One bounded polling-lifecycle change around the existing apply transaction.
