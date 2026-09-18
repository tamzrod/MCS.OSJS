# UMIG-006 — Adapt Toolkit Diagnostics to OS.js

Status: PLANNED / BLOCKED — donor freeze and human promotion required. Not ACTIVE.
Previous: UMIG-005
Next: UMIG-007

## Primary outcome
Show truthful OS.js-available diagnostic information in the copied Diagnostics tab.

## Scope
Retain the approved Electron tab layout, map available Simulator/Replicator/MMA2 status and errors through supported OS.js interfaces, and clearly disable or omit Windows-only service start/stop and native filesystem paths when no equivalent authorized Linux API exists. Provide explicit unavailable/unknown states rather than default green.

## Non-scope
No new privileged service-control endpoint, backend restart commands, Windows IPC or desktop redesign.

## Acceptance
1. Diagnostic status reflects actual available observations and errors.
2. Unsupported native controls cannot invoke Windows-only APIs.
3. Missing diagnostic data is explicitly unknown or unavailable.

## Verification
Focused diagnostics fixture checks for healthy, stopped, unreachable and unsupported-capability conditions inside OS.js.

## Dependencies
UMIG-005 and human promotion. Any proposed new service-control capability needs a separately approved task.

## Sizing
Surface 1, environment 0, behavior 1, verification 1, recovery 0 = 3.
