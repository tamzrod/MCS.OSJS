# Microtask OSJS-001 — Copy Nameless SCADA OS.js Base Web App

## ID and Title

- ID: `OSJS-001`
- Title: Copy the Nameless SCADA OS.js base web app into `OSJS/`

## Primary Outcome

The existing OS.js base web app from the Nameless SCADA donor is copied into a new top-level `OSJS/` directory in MCS.OSJS, preserving the donor application structure so it can be used as the starting point for later MCS.OSJS work.

## Scope

- Source: `tamzrod/namelessscada` → `desktop/osjs-prototype/`.
- Destination: repository-root `OSJS/`.
- Copy the complete donor OS.js base web app tree into `OSJS/`.
- Preserve the donor directory structure and files as copied.

## Non-Scope

- No architectural rewiring.
- No Orchestrator, Replicator, or MMA2 integration.
- No port changes.
- No UI redesign.
- No cleanup, renaming, refactoring, or removal of donor applications/assets.
- No unrelated repository changes.

## Acceptance Criteria

1. `OSJS/` exists at the MCS.OSJS repository root.
2. The contents of Nameless SCADA `desktop/osjs-prototype/` are copied into `OSJS/` with the donor tree preserved.
3. No files outside `OSJS/` are changed by this task.

## Verification

- Compare the donor `desktop/osjs-prototype/` tree against the new `OSJS/` tree and verify the copied paths match.
- Verify key donor files such as `package.json`, client/server source, packages, assets, and build/runtime files exist under `OSJS/`.
- Verify the task diff is confined to `OSJS/`.

## Dependencies

- Human promotion into `workflow/active_work/` before execution.
- Access to the Nameless SCADA donor repository.

## Sizing Assessment

- One primary outcome: copy one existing application tree from donor to destination.
- No design decisions or rewiring.
- One verification workflow: donor-tree versus destination-tree comparison.
- Size: good JR task.
