# Brainstorm: Electron Post-Test Polish

Status: brainstorm material only. Non-authoritative. No implementation authorized until human promotion.

Date: 2026-09-12

## Test observations

The packaged Electron build is now running successfully on Windows, but hands-on testing exposed three usability / packaging issues that should be resolved before considering the desktop packaging work finished.

### 1. Parameter editing currently requires Run as administrator

Observed behavior:

- Electron can launch normally.
- Editing / applying parameters only works reliably when the application is launched as Administrator.

Desired behavior:

- Electron remains a normal non-elevated desktop application.
- Ordinary parameter editing must not require UAC elevation.
- Installation / service registration may remain elevated, but routine operator use should not be.

Candidate direction:

- Verify the writable runtime/config path used by the packaged Windows deployment.
- Give normal users only the required Modify permission on the mutable MCS runtime/config location.
- Do not make the installed binaries or application directory generally writable.
- Confirm the Electron UI and the NSSM-managed backend services are using the same intended runtime/config location.

### 2. FC dropdown appears to lose focus while runtime status is active

Observed behavior:

- The FC dropdown can unexpectedly lose focus / close while the user is interacting with it.
- The problem appears periodic.
- Initial suspicion is that a running-status / activity update, possibly the Simulator activity indicator, is triggering a UI update at the same time.

Important constraint:

- Runtime / activity indicators may continue updating visually, but those updates must never rebuild or replace editable controls currently being used by the operator.

Candidate investigation:

- Instrument or inspect periodic status/activity callbacks while the FC dropdown is open.
- Check for any `render()`, `replaceChildren()`, editor reconstruction, or equivalent DOM replacement occurring during polling / LED activity.
- Confirm whether the FC `<select>` node is being destroyed and recreated.
- Prefer patching only the status text / LED class/state during periodic updates.
- Preserve focus and open interactive controls during background status updates.

The blinking LED is only a suspect at this stage; do not remove the indicator unless testing proves that the animation itself is responsible.

### 3. Replace the default Electron icon

Observed behavior:

- The packaged application still uses the default Electron icon.

Desired behavior:

- Use a simple icon appropriate for a small industrial Modbus utility / toolkit.
- The same identity should appear on the packaged executable, Windows shortcut / Start Menu entry, installer where applicable, and application window/taskbar.

Candidate direction:

- Add a Windows `.ico` asset under the Electron build resources.
- Configure `electron-builder` Windows icon metadata.
- Configure the Electron `BrowserWindow` icon where needed.
- Prefer a simple industrial/network/device motif that remains legible at small Windows icon sizes.

## Candidate completion order

1. Remove the normal-use Administrator requirement.
2. Reproduce and isolate the FC dropdown focus loss.
3. Make periodic status/activity rendering focus-safe.
4. Replace the default Electron icon.
5. Rebuild the Windows installer.
6. Re-test parameter editing as a standard user while Simulator/MMA2 status activity is running.

## Candidate acceptance checks

- Installed Electron application launches and edits/applies parameters as a standard Windows user.
- No routine operator action requires Run as administrator.
- FC dropdown remains open and usable while background runtime/status/activity updates continue.
- Status LEDs continue to update without rebuilding the editor controls.
- Packaged executable and Windows shortcuts no longer show the default Electron icon.

## Scope note

These are post-test packaging/UI polish findings only. They do not authorize implementation and do not change the existing backend/runtime architecture.