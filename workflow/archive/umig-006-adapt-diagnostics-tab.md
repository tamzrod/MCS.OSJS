# UMIG-006 — CODE: Adapt OS.js Diagnostics

Status: COMPLETE — source-only checkpoint 2026-09-19; independent TEST/VERIFY NOT RUN.
Stage / owner: CODE / ChatGPT
Previous: UMIG-005-V (core live behavior archived, real window reopen remains unverified)
Next: UMIG-006-T (independent unit/build TEST, sole ACTIVE)

## Outcome
Toolkit `index.js` replaces the Diagnostics fixture before DOM mount and creates a Toolkit-owned read-only editor using the already-existing Memory and Replicator contracts/transports. `diagnostics-observer.js` loads each canonical document, explicitly labels the FIRST configured device of each type and requests that specific device's status only; independent load/status failures yield UNAVAILABLE and never mask the other service. An empty canonical document yields UNKNOWN without invoking status. It NEVER calls apply/suggest or writes config.

Prepared `diagnostics-model.js` now is live-imported via observer/editor and accepts actual Simulator None/IDLE. Its global MMA2/Simulator/Replicator service-health fields remain UNKNOWN: per-device status is not proof of supervisor health. Paths/runtime mode remain UNAVAILABLE, and Start/Stop controls remain disabled without handlers. `diagnostics-editor.js` preserves the donor actions/paths/log structure and flex pane within Toolkit ShadowRoot, renders device-scoped statuses plus genuine Replicator per-block errors, explicitly labels its text as observations NOT service logs, refreshes read-only every five seconds, invalidates stale responses and stops polling on window destruction. No fixture fallback, native Windows API, privileged operation, new HTTP route, service control or direct probe.

## Source-only evidence / scope
Baseline `a32776608317669d75bea2b87369f69dfb518b55`; CODE checkpoint `c996794ff14dd65479e84123724f6b9fd9235ee6`. GitHub compare confirmed exactly six Toolkit code/test paths: modified `index.js`, `diagnostics-model.js`; added `diagnostics-observer.js`, `diagnostics-editor.js`, `OSJS/tests/toolkit-diagnostics-observer.test.js`, `OSJS/tests/toolkit-diagnostics-editor.test.js`. Existing `OSJS/tests/toolkit-diagnostics-model.test.js` is included in future independent run. Readback of source/imports performed. NO test execution, build, browser or live Diagnostics acceptance by CODE. No Memory/Replicator implementation, Go, MMA2, production Compose, legacy app, Electron, ICC or general `operation cwal.md` edit.

## Limitations and successor
The pane observes only the FIRST persisted device of each type, clearly labelled; multiple-device aggregate service health, actual global supervisor health, native Windows services, filesystem paths, Docker controls and runtime log retrieval are UNSUPPORTED/UNKNOWN/UNAVAILABLE rather than simulated. UMIG-006-T must independently test pure mapping, canonical selection/partial failures, disabled controls, stale refresh/disposal, existing Memory/Replicator regression, build/discovery and import isolation. UMIG-006-V remains QUEUED; its future actual-browser packet must also observe the real Toolkit window CLOSE and Start-menu RELAUNCH omitted by UMIG-005-V. Do not claim that reopen gate passed.

## Sizing
Surface 1, environment 0, behavior 1, verification 0, recovery 0 = 2.
