# UMIG-003 — CODE: Copy Approved Renderer into OS.js

Status: QUEUED — promoted 2026-09-18; blocked from execution until UMIG-001 records an approved donor SHA.
Stage / owner: CODE / ChatGPT
Previous: UMIG-001
Next: UMIG-003-T

## Primary outcome
Author a self-contained, OS.js-owned copy of the approved three-tab Toolkit renderer.

## Scope
Copy files/assets from the pinned donor SHA into `OSJS/src/packages/MCSModbusToolkit`, adapt OS.js window bootstrap, scope CSS away from OS.js chrome, provide fixture-only data and explicit unknown LED states. Retain Memory, Replicator and Diagnostics tab layout. The Toolkit is built within the existing `osjs-shell` Docker image; do not add an Electron runtime or a new Toolkit container.

## Non-scope
No build/tests, real backend calls, Electron preload/main, legacy UI deletion, Modpoll implementation, new feature or visual-parity PASS.

## Coding acceptance / handoff
1. Self-owned three-tab renderer and required assets authored and committed from the approved SHA.
2. No Electron runtime/global CSS or legacy-package dependency added; missing status remains UNKNOWN.
3. Read back changed files and record commit, donor provenance, fixtures and known host gaps for OpenHands; do not claim tested.

## Dependencies
UMIG-001 human-approved donor SHA and accepted source-only donor note. Only execute when this task becomes sole ACTIVE.

## Sizing
Surface 2, environment 0, behavior 0, verification 0, recovery 1 = 3.
