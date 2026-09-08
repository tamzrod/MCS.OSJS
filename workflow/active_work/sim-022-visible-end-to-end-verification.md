# SIM-022 — Verify the Visible Simulator-to-MMA2 Flow End to End

Status: ACTIVE — promoted by user 2026-09-08; execute after SIM-020 and SIM-021.

## Primary outcome

Prove from the rebuilt operator UI and independent protocol observations that Save & Apply reaches MMA2 and that randomized Simulator values continue flowing through Raw Ingest.

## Scope

- Start MMA2 independently under the deployment's normal boot/lifecycle mechanism.
- Open the rebuilt Modbus Simulator window and apply one enabled four-FC definition.
- Observe the UI's apply, MMA2, Simulator, Raw Ingest, and per-FC timing states.
- Read FC1–FC4 through real Modbus TCP and prove values change across scheduled cycles.
- Verify structural restart, timing-only no-restart, rejected-change preservation, MMA2 outage/error display, recovery, and full reboot automatic resume.
- Capture repeatable test evidence in the task record.

## Non-scope

- No new product behavior.
- No acceptance based only on unit tests, saved JSON, or status labels without protocol/runtime corroboration.

## Acceptance criteria

1. A structural Save & Apply visibly completes one MMA2 restart and returns to ready/RUNNING.
2. FC1–FC4 Last timestamps advance and real Modbus reads observe scheduler-produced values.
3. Raw Ingest and MMA2 failures appear truthfully and recover after the underlying path succeeds.
4. Timing-only changes do not restart MMA2; rejected changes preserve the working config and runtime.
5. After a full MMA2 plus Simulator-runtime restart, the saved enabled simulation resumes without manual apply.

## Verification

Run a browser-driven system test plus independent restart-artifact, Raw Ingest, and Modbus-TCP assertions. Run affected Go tests with race detection and OS.js syntax/package/full builds.

## Dependencies

- SIM-020.
- SIM-021.

## Sizing

Implementation 0, environment 1, behavioral 2, verification 2, decision/recovery 0 = 5. Verification-only capstone with one coordinated end-to-end workflow.
