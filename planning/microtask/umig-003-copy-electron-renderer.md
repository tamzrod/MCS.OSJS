# UMIG-003 — Copy the Approved Electron Renderer into Toolkit

Status: PLANNED / BLOCKED — UMIG-002 scaffold and UMIG-001 donor freeze must pass; human promotion required. Not ACTIVE.
Previous: UMIG-001
Next: UMIG-004

## Primary outcome
Display a self-contained copy of the approved Electron Toolkit UI inside the single OS.js window using fixture data.

## Scope
Copy the pinned Electron renderer HTML/CSS/JS, communications display script and required UI assets into OS.js-owned files within the new package. Adapt only host bootstrap and style scoping so desktop-wide `html, body` rules cannot affect the OS.js shell. Use fixture data for rendering; preserve Memory, Replicator and Diagnostics tabs and approved layout/controls. Document host differences and known status-contract gaps without inventing healthy LED values. This is a one-time source copy; Windows Electron and OS.js remain independently maintained, built and deployed.

## Non-scope
No shared renderer package, symlink, cross-directory import, generated link, automatic synchronization, real backend calls, Electron preload/main process, Windows service management, legacy app deletion or new UI features.

## Acceptance
1. All three donor tabs render within one OS.js window from the pinned, locally copied source.
2. Styles are scoped and do not alter OS.js desktop, taskbar or window chrome.
3. Fixture rendering and the bundle require no Electron installation, `window.mcsDesktop`, cross-deployment source or native Electron APIs.

## Verification
Run the OS.js Toolkit build and static fixture rendering; inspect DOM, CSS scoping and bundle/import/asset paths. Verify unknown/unavailable LED state is truthful when evidence is missing. Backend acceptance belongs to later tasks.

## Dependencies
UMIG-002 accepted and UMIG-001 approved donor SHA pinned, then human promotion. If donor changes, stop and obtain renewed approval rather than silently tracking the latest Electron files.

## Sizing
Surface 2, environment 0, behavior 0, verification 1, recovery 1 = 4; bounded copy/static-render workflow, separated from runtime wiring.
