# Handoff

## Current direction — 2026-09-19

Human reports Docker Compose deployment working and approved `1c971b9` as the one-time Electron UI donor for the Docker-hosted `MCS Modbus Toolkit`. Target: a self-owned Toolkit inside existing `osjs-shell`, reusing existing MMA2/Simulator/Replicator services. Electron stays an independent Windows app. Portable distribution and a Modpoll tab remain outside this sequence.

## Roles and authoritative task state

ChatGPT owns CODE, source checkpoints, review of independent TEST/VERIFY evidence and workflow advancement. OpenHands/JR runs only an explicit current `JR TEST TASK` under `operation cwal.md`, reports evidence, never fixes source/changes task status/edits ICC. Only BLACK SHEEP WALL edits ICC. Authoritative task selection is `workflow/active_work/` plus this handoff, not stale ICC.

**Sole ACTIVE: UMIG-003 — CODE: copy the approved renderer into OS.js.** Human explicitly approved full donor commit `1c971b9a6e00bafadf329df8821421a40cfc079c` on 2026-09-19 in response to the proposed `1c971b9`. UMIG-001's source-only donor provenance is recorded at `workflow/archive/umig-001-freeze-electron-ui-donor.md` and its human decision gate is COMPLETE. The donor commit is a Docker IPC repair commit whose snapshot contains the pinned Electron UI; approval is for that *renderer snapshot*, not evidence of a Windows test or authorization to import native Electron code. UMIG-003 is authorized for CODE only; **no product implementation is claimed yet**. The immediate next task UMIG-003-T and rest of previously promoted UMIG chain through UMIG-007A remain QUEUED and advance only after their own stage gates. There is **no current JR TEST TASK** until a valid UMIG-003 source checkpoint is committed and the coding agent activates UMIG-003-T.

`DOCKER-001` remains QUEUED/paused, not PASS/archived: operator reports functioning Docker deployment, but independent Linux/Windows build and IPC evidence was not recorded. Preserve existing Docker data volume and user configuration. Other previously queued MEM/RLED/RREC/REP-BLOCK tasks are unchanged. UMIG-008 launcher cutover and UMIG-009 legacy UI retirement plus tests stay in Planning pending separate human approvals. Existing Simulator and Replicator packages remain available.

## Frozen UI donor and exact CODE boundary

Pinned files at `1c971b9a6e00bafadf329df8821421a40cfc079c`: `electron/renderer/index.html`, `style.css`, `app.js`, `comms-status.js`; optional source assets `electron/build/replicator-logo.svg` and `icon.png` only if necessary and copied into OS.js-owned paths. Do not copy backup `.bak` files, `electron/main.js`, `preload.js`, installer assets or runtime bridge modules. The archive records each donor component, Electron IPC APIs, global-CSS collisions, Windows diagnostics/paths and false-status risks. This is a ONE-TIME frozen reference: no shared module, symlink, cross-build/runtime dependency or silent auto-sync.

UMIG-003 scope: render Memory, Replicator and Diagnostics within the already verified one-window Toolkit scaffold using fixture-only data; adapt window-owned event/timer lifecycle and scope CSS to Toolkit content so OS.js desktop/Start menu/taskbar/clock remain intact. All missing/stale status and COMMS LEDs must be UNKNOWN/UNAVAILABLE; native Diagnostics Start/Stop and Windows service/path controls must be disabled, not executed. Do not call `window.mcsDesktop`, real OS.js relays or backend apply while implementing this fixture-stage copy. Real Memory/Replicator bridge ownership and live tests follow under UMIG-004/005; Diagnostics mapping under UMIG-006; independent appearance and Electron-free Docker acceptance under UMIG-007/007A. No product tests are claimed by donor approval.

## Workflow sequence and verification boundary

UMIG-003 CODE → UMIG-003-T TEST (package build/discovery and fixtures) → UMIG-003-V VERIFY (render all three tabs and shell isolation) → UMIG-004 CODE/TEST/VERIFY (Memory) → UMIG-005 CODE/TEST/VERIFY (Replicator) → UMIG-006 CODE/TEST/VERIFY (Diagnostics) → UMIG-007 visual VERIFY → UMIG-007A independent Docker/OS.js VERIFY → STOP for separate human decision about cutover. ChatGPT, not JR, advances the exact next QUEUED task after reading its gate's real evidence; FAIL/BLOCKED never auto-advances. Preserve user data and never run destructive production tests.

## Historical evidence and limits

UMIG-002-T independent package build/discovery PASS at `e9d25e3`; UMIG-002-V independent one-window disconnected-placeholder rendered PASS at `7349caa`; both archived. These establish the original placeholder only, not tab rendering or backend connectivity. Original discovery FAIL at `752a541` and prior ICC prerequisite BLOCKED at `3831e33` remain immutable historical evidence. Four earlier icon/sound 404s are observations, not claimed fixed. Docker IPC repair source checkpoint: `1c971b9a6e00bafadf329df8821421a40cfc079c`; operator reports deployment fine without supplying independent commands/Windows evidence. ICC index remains known stale; use current Git sources, and only a specifically needed bounded BLACK SHEEP WALL run can maintain it.

## Next action and recommendation

Next action (ChatGPT CODE): implement only UMIG-003 from the pinned donor snapshot and archived provenance; commit/read back Toolkit-only source and fixture changes, then prepare the exact JR TEST TASK when activating UMIG-003-T. Recommendation: preserve the working Docker deployment, Windows Electron and legacy OS.js packages untouched; require separate independent TEST and rendered VERIFY before claiming functional migration or requesting launcher/legacy removal approval.
