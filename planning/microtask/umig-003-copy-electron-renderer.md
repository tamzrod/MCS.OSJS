# UMIG-003 — Copy the Approved Electron Renderer into Toolkit

Status: PLANNED / BLOCKED — donor freeze and human promotion required. Not ACTIVE.
Previous: UMIG-002
Next: UMIG-004

## Primary outcome
Display the approved Electron Toolkit UI inside the single OS.js window using fixture data.

## Scope
Copy the pinned Electron `renderer/index.html`, `renderer/style.css`, `renderer/app.js` and needed UI assets into the new OS.js package, adapting only the host bootstrap and style scoping so desktop-wide `html, body` rules cannot affect the OS.js shell. Supply a fixture/mock host interface for rendering; preserve Memory, Replicator and Diagnostics tabs and the donor's layout/controls. Document necessary host differences rather than redesigning the UI.

## Non-scope
No real runtime calls, Electron preload/main process, Windows service management, existing app deletion or new UI features.

## Acceptance
1. All three donor tabs render within one OS.js window from the pinned source.
2. Styles are scoped to Toolkit and do not overwrite desktop/taskbar/window chrome.
3. Fixture rendering does not require `window.mcsDesktop` or native Electron APIs.

## Verification
Run the Toolkit build/static render check against fixture data; inspect resulting DOM and CSS scoping. Real backend acceptance belongs to later tasks.

## Dependencies
UMIG-002; donor commit pinned by UMIG-001. If the donor changes, stop and obtain renewed human approval rather than silently chasing its latest files.

## Sizing
Surface 2, environment 0, behavior 0, verification 1, recovery 1 = 4 (one bounded donor-copy/static-render workflow; runtime is split out).
