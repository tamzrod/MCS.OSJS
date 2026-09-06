# Microtask List — OS.js Base Web App Initialization

Source planning: `planning/Brainstorm/osjs-base-webapp-init.md`

Human clarification supersedes the brainstorm's sample-app candidate: MCS.OSJS needs the basic OS.js desktop only. Nameless SCADA applications are not part of this initialization.

## OSJS-001 — Initialize the basic OS.js desktop in `OSJS/`

### Primary Outcome

A runnable basic OS.js desktop from Nameless SCADA exists under repository-root `OSJS/`, without Nameless SCADA application packages.

### Scope

- Use `tamzrod/namelessscada` → `desktop/osjs-prototype/` as the donor.
- Create `OSJS/` and copy only the files required for the basic OS.js desktop runtime/build.
- Exclude Nameless SCADA application packages and application-specific integration code.
- Make only the minimum reference/removal adjustments required for the basic desktop to build and start without the excluded applications.

### Non-Scope

- No SCADA applications, sample application, Orchestrator, Replicator, MMA2, Modbus, DNP3, Tag Manager, or Ingestor integration.
- No port change.
- No UI redesign or unrelated refactor.

### Acceptance Criteria

1. `OSJS/` contains the basic OS.js desktop runtime/build structure.
2. The copied desktop builds, starts, and renders the basic OS.js desktop.
3. Nameless SCADA application packages are absent and are not loaded by the desktop.

### Verification

Build/start the desktop from `OSJS/`, confirm the desktop renders, and confirm Nameless SCADA applications are absent from package discovery and the launcher/menu.

### Dependencies

- Access to the Nameless SCADA donor repository.
- Human promotion before execution.

### Sizing Assessment

- One primary outcome: establish the runnable basic desktop.
- Three bounded implementation actions: copy the required base, exclude donor applications, make minimum boot-reference adjustments.
- One verification workflow.
- Size: 3–4; acceptable as one JR task. Split only if execution reveals an independent blocking problem.

---

## OSJS-002 — Apply the donor desktop color schemes

### Primary Outcome

The basic MCS.OSJS desktop uses the selected color-scheme/theme configuration copied from the Nameless SCADA basic desktop donor material.

### Scope

- Identify the donor color-scheme/theme configuration used by the basic desktop.
- Copy the required theme/color-scheme files into the existing `OSJS/` base.
- Wire only the configuration needed for those schemes to be available to the basic desktop.

### Non-Scope

- No donor applications.
- No new theme design.
- No application-specific icons/assets unless required by the basic desktop theme itself.

### Acceptance Criteria

1. Required donor color-scheme/theme files are present under `OSJS/`.
2. The basic desktop loads the copied scheme/theme without errors.
3. No Nameless SCADA application package is introduced by the theme copy.

### Verification

Start the OS.js desktop and verify the copied color scheme/theme is applied or available as intended, with no donor applications added.

### Dependencies

- OSJS-001.
- Human promotion before execution.

### Sizing Assessment

- One primary outcome: carry over the basic desktop appearance.
- One verification workflow.
- Size: 2–3; good JR task.

---

## OSJS-003 — Assign and document the OS.js management port

### Primary Outcome

The MCS.OSJS basic desktop runs on one chosen random management/UI port, and that chosen port is recorded in project documentation.

### Scope

- Choose one available random port for the OS.js management/UI endpoint.
- Configure the OS.js base runtime/deployment to use that port.
- Record the chosen port in `docs/NETWORK_EXPOSURE.md` and the nearest OSJS runtime/readme reference if one exists.

### Non-Scope

- No MMA2 port assignment or changes.
- No dynamic-per-boot port discovery system.
- No broader port-authority architecture.

### Acceptance Criteria

1. One OS.js management/UI port is configured.
2. The OS.js desktop is reachable on the chosen port.
3. The chosen port is documented consistently.

### Verification

Start the OS.js desktop, connect to the configured port, and compare the runtime value with the documented value.

### Dependencies

- OSJS-001.
- Human promotion before execution.

### Sizing Assessment

- One primary outcome: establish the documented OS.js management endpoint.
- Three bounded implementation actions: choose, configure, document.
- One verification workflow.
- Size: 2–3; good JR task.

## Sequence

```text
OSJS-001 basic desktop
    ↓
OSJS-002 donor color schemes
    ↓
OSJS-003 management port + documentation
```

These are planning tasks only until human promotion moves a selected task into `workflow/active_work/`.
