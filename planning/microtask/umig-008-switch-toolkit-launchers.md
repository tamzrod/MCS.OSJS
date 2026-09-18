# UMIG-008 — Switch OS.js Launchers to One Toolkit

Status: PLANNED / BLOCKED — donor freeze and human promotion required. Not ACTIVE.
Previous: UMIG-007
Next: UMIG-009

## Primary outcome
Make MCS Modbus Toolkit the single normal launch target in the OS.js workstation.

## Scope
Change project-owned desktop shortcuts and any configured default launch/start-menu or auto-start references from legacy Modbus app names to Toolkit. Preserve unrelated OS.js shell packages/settings. Handle existing saved session references deliberately so users are not forced into broken legacy launch targets.

## Non-scope
Do not delete legacy UI packages yet, alter backend services or change Electron launch behavior.

## Acceptance
1. Normal MCS desktop launch opens one Toolkit window.
2. New auto-start/launcher references target Toolkit, not retired application names.
3. Legacy package source remains available as fallback until the following retirement task.

## Verification
Inspect launcher/auto-start mappings and launch from a clean OS.js session; check one-window result and record any saved-session migration requirement.

## Dependencies
UMIG-007 visual gate and successful focused runtime gates UMIG-004/005/006; human approval of cutover/promotion.

## Sizing
Surface 1, environment 0, behavior 1, verification 1, recovery 1 = 4 (one launcher cutover outcome).
