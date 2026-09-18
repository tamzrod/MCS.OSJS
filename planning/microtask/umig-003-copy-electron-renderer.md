# UMIG-003 — Copy the Approved Electron Renderer into Toolkit

Status: PLANNED / BLOCKED — donor freeze and human promotion required. Not ACTIVE.
Previous: UMIG-002
Next: UMIG-004

## Primary outcome
Display a self-contained copy of the approved Electron Toolkit UI inside the single OS.js window using fixture data.

## Scope
Copy the pinned Electron `renderer/index.html`, `renderer/style.css`, `renderer/app.js` and needed UI assets into OS.js-owned files inside the new package. Adapt only host bootstrap and style scoping so desktop-wide `html, body` rules cannot affect the OS.js shell. Supply a fixture/mock host interface for rendering; preserve Memory, Replicator and Diagnostics tabs and the donor's layout/controls. Document necessary host differences rather than redesigning the UI. This is a one-time source copy: the Windows Electron app and OS.js Toolkit remain independently maintained, built and deployed. The OS.js copy must not reference files or assets from `electron/` at build time or runtime.

## Non-scope
No shared renderer package, symlink, cross-directory import, generated link, automatic upstream synchronization, real runtime calls, Electron preload/main process, Windows service management, existing app deletion or new UI features.

## Acceptance
1. All three donor tabs render within one OS.js window from the pinned, locally copied source.
2. Styles are scoped to Toolkit and do not overwrite desktop/taskbar/window chrome.
3. Fixture rendering does not require `window.mcsDesktop`, Electron installation or native Electron APIs.
4. OS.js bundle references only OS.js-owned Toolkit files/assets; no shared Electron source or build dependency remains.

## Verification
Run the OS.js Toolkit build/static render check against fixture data; inspect resulting DOM, CSS scoping and bundle/import/asset paths for independence. Real backend acceptance belongs to later tasks.

## Dependencies
UMIG-002; donor commit pinned by UMIG-001. If the donor changes, stop and obtain renewed human approval rather than silently chasing its latest files.

## Sizing
Surface 2, environment 0, behavior 0, verification 1, recovery 1 = 4 (one bounded donor-copy/static-render workflow; runtime is split out).
