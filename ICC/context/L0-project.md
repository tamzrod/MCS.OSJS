# L0 — MCS.OSJS project

Parent: none
Zoom Out: none
Zoom In: [governance](governance.md), [donor-licensing](donor-licensing.md), [network-exposure](network-exposure.md), [planning-workflow](planning-workflow.md), [osjs-shell](osjs-shell.md), [simulator-device-config](simulator-device-config.md), [replicator](replicator.md), [electron-memory-layout](electron-memory-layout.md), [electron-replicator-advanced](electron-replicator-advanced.md), [electron-replicator-leds](electron-replicator-leds.md), [electron-settings-owner](electron-settings-owner.md), [active-work](active-work.md)
Source dependencies: `README.md`, `PROJECT_IDENTITY.md`, `handoff.md`, `OSJS/src/packages/MCSModbusToolkit/index.js`.

## Identity and boundaries

MCS.OSJS is the standalone OS.js-based Modbus Consolidation System. The OS.js desktop hosts the Toolkit; MMA2, Simulator and Replicator are distinct runtimes. The Electron Toolkit is a separate deployment target. Consult authoritative source for deployment-specific implementation details rather than generalizing Windows transports or paths to OS.js.

## Verified remote snapshot, not a runtime acceptance

Observed `main` HEAD: `45cd3d2d81831306c6943f844e49b7deccbfc5ad` (2026-09-23 review). Local worktree and uncommitted overlay were not inspected. At this committed snapshot, `OSJS/src/packages/MCSModbusToolkit/index.js` registers one Toolkit window and mounts Memory, Replicator, Diagnostics and shared-settings UI/contract components. It is **not** the historical placeholder described by the superseded ICC snapshot. This source inspection proves module wiring exists, not that live Linux/Docker operation passes.

The committed `handoff.md` names OTR-001A as the current read-only task for that main snapshot and describes the OTR successor sequence. Read [active-work](active-work.md) for the scoped summary and the actual handoff/packet for task authority. A different branch or subsequent checkout may have different task state.

## Navigation and evidence

For desktop and Toolkit integration, zoom into [osjs-shell](osjs-shell.md). For task routing, use [active-work](active-work.md); workflow status does not determine product architecture. Follow other children only when the requested question requires them. ICC caches verified findings; Git and task packets remain authoritative. Do not infer any product TEST/VERIFY PASS, OS.js cutover, deployment readiness, or successor activation from this node.

## Refresh rule

If a source dependency changes after this snapshot, refresh only the facts affected by that delta. No whole-project rescan merely because HEAD advanced. Preserve an older verified node as historical/possibly stale until its specific dependencies are checked. Bootstrap is not a default fallback.
