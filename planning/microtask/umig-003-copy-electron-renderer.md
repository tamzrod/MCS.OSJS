# UMIG-003 — CODE: Copy Approved Renderer into OS.js

Status: PLANNED / BLOCKED — donor approval and human promotion required.
Stage / owner: CODE / ChatGPT
Previous: UMIG-001
Next: UMIG-003-T

## Primary outcome
Author a self-contained, OS.js-owned copy of the approved three-tab Toolkit renderer.

## Scope
Copy files/assets from the pinned donor SHA into `OSJS/src/packages/MCSModbusToolkit`, adapt OS.js window bootstrap, scope CSS away from OS.js chrome, provide fixture-only data and explicit unknown LED states. Retain Memory, Replicator and Diagnostics tab layout. No live Electron source imports, symlinks or automatic sync.

## Non-scope
No build/tests, real backend calls, Electron preload/main, legacy UI deletion, new feature or claim of visual parity.

## Coding acceptance / handoff
1. Self-owned three-tab renderer and required assets authored and committed from the approved SHA.
2. No Electron runtime/global CSS or legacy-package dependency added; missing status remains UNKNOWN.
3. Read back changed files and record commit, donor provenance and known host gaps for OpenHands; do not claim tested.

## Dependencies
UMIG-001 approved donor note and human promotion.

## Sizing
Surface 2, environment 0, behavior 0, verification 0, recovery 1 = 3.
