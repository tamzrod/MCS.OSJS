# Microtask OSJS-001 — Copy Nameless SCADA Basic OS.js Desktop

## ID and Title

- ID: `OSJS-001`
- Title: Copy the Nameless SCADA basic OS.js desktop into `OSJS/`

## Primary Outcome

MCS.OSJS contains a runnable basic OS.js desktop copied from the Nameless SCADA donor, without copying the donor SCADA applications.

## Scope

- Source: `tamzrod/namelessscada` → `desktop/osjs-prototype/`.
- Destination: repository-root `OSJS/`.
- Copy only the donor files required to build, start, and render the basic OS.js desktop shell.
- Preserve the basic desktop appearance/configuration required for the shell to run.
- Keep only runtime/build files and shell assets that are actually required by the basic desktop.

## Explicit Exclusions

Do not copy donor application packages or application-specific integration code.

At minimum, exclude the application content under `desktop/osjs-prototype/src/packages/`, including SCADA/editor/management applications such as Ingestor, ModbusEditor, Dnp3Editor, TagManager, AutoStart, TaskbarSettings, and other donor applications.

Also exclude application-specific client/server providers, tests, scripts, or assets when they exist only to support those excluded applications.

## Non-Scope

- No Orchestrator, Replicator, MMA2, Modbus, DNP3, Tag Manager, or Ingestor integration.
- No port changes.
- No new application development.
- No UI redesign.
- No architectural rewiring beyond removing references required to prevent excluded donor applications from being loaded.
- No unrelated repository changes.

## Acceptance Criteria

1. `OSJS/` exists at the MCS.OSJS repository root and contains the basic OS.js desktop runtime/build structure from the donor.
2. The OS.js desktop can build/start and render the basic desktop shell.
3. Donor SCADA/application packages are not present or loaded in the copied desktop.

## Verification

- Build/start the copied OS.js desktop from `OSJS/` using the donor's applicable base runtime workflow.
- Verify the OS.js desktop shell renders successfully.
- Verify excluded donor applications are absent from the desktop/package discovery and are not available from the desktop menus/application launcher.
- Verify the implementation diff is confined to `OSJS/`.

## Dependencies

- Human promotion into `workflow/active_work/` before execution.
- Access to the Nameless SCADA donor repository.

## Sizing Assessment

- One primary outcome: produce a runnable basic OS.js desktop from the existing donor shell.
- Selection is bounded by a clear rule: shell/runtime required for desktop boot stays; donor applications and their integration code do not.
- One verification workflow: build/start desktop and confirm donor applications are absent.
- Size: good JR task.
