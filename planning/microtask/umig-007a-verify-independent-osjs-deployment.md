# UMIG-007A — Verify Independent OS.js Toolkit Deployment

Status: PLANNED / BLOCKED — donor freeze and human promotion required. Not ACTIVE.
Previous: UMIG-007
Next: UMIG-008

## Primary outcome
Prove the OS.js Toolkit can be built and deployed without the Electron application or its build artifacts.

## Scope
Use an isolated OS.js deployment/test target containing the OS.js-owned Toolkit copy and its normal OS.js backend services. Execute one OS.js package/build-and-launch smoke workflow with no installed Electron binary, Electron `node_modules`, Electron renderer directory or Electron build output on that target. Inspect Toolkit bundle assets and configuration for cross-deployment source imports, symlinks, packaging references, network calls or runtime IPC to the Windows Electron installation. Do not change either product to make this verification pass; record blockers for separately scoped fixes.

## Non-scope
No shared code, shared package, common deployment pipeline, Electron rebuild, backend redesign, new runtime behavior, launcher cutover or legacy UI removal.

## Acceptance
1. OS.js Toolkit builds and launches from OS.js-owned sources/assets alone on the isolated test target.
2. OS.js Toolkit uses its own deployed services and has no build/runtime dependency on an Electron installation or artifacts.
3. Electron's standalone Windows packaging and source files remain unchanged by the OS.js deployment workflow.

## Verification
Run and record an isolated OS.js deployment smoke workflow and inspect the produced bundle/configuration plus changed paths. A visual comparison alone is insufficient; an OS.js runtime test does not claim Windows Electron acceptance.

## Dependencies
UMIG-007 completed, focused UMIG-004/005/006 runtime gates passed, and explicit human promotion. This independent-deployment gate must pass before UMIG-008 launcher cutover.

## Sizing
Surface 0, environment 1, behavior 0, verification 1, recovery 1 = 3.
