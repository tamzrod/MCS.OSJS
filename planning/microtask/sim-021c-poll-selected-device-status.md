# SIM-021C — Poll Selected-Device Runtime Status

Status: PLANNING ONLY — not authorized for implementation until human promotion.

## Primary outcome

Make the Modbus Simulator window request current runtime status for the selected device while the window is open.

## Scope

- Reuse the existing application WebSocket runtime request helper/path established by SIM-019/SIM-020.
- Start a bounded 1-second status refresh loop only while the Modbus Simulator window is alive.
- Request status only for the currently selected device.
- Refresh immediately after selection changes and after a successful Save & Apply.
- Stop/replace stale polling when the selected device changes or the window closes.
- Store the latest successful status response separately from transient bottom-bar messages.
- Treat relay/runtime request failure as unavailable status data; do not fabricate `RUNNING` from the saved device model.

## Non-scope

- No status-row HTML/CSS presentation.
- No new server route or transport.
- No MMA2 control.
- No FC Last/Next display.
- No register values.

## Acceptance criteria

1. With one selected device, the client sends one status request immediately and then at approximately 1-second cadence.
2. Selecting another device switches status requests to the new device without continuing stale requests for the old device.
3. Closing the window stops the refresh loop.
4. A failed status request does not cause the client to synthesize `RUNNING` from `Enabled` or cached configuration.

## Verification

Use focused client/static tests where available plus `node --check`/package build; verify request lifecycle with an instrumented or stubbed runtime requester.

## Dependencies

- SIM-021B.
- SIM-020 runtime request plumbing.

## Sizing

Implementation 1, environment 0, behavioral 1, verification 1, decision/recovery 0 = 3. One bounded client lifecycle behavior.
