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

1. `OSJS/` contains the minimum donor-derived runtime/build structure required to start the basic OS.js desktop, with no Nameless SCADA application packages included.
2. A clean build and fresh start succeed and render the basic OS.js desktop without startup, package-discovery, or missing-module errors caused by removed donor applications.
3. After restart, the desktop remains usable and no excluded Nameless SCADA application appears in package discovery, launcher/menu entries, autostart, or runtime requests.

### Test and Verification

Use one end-to-end verification workflow and record evidence for every step:

1. **Static scope check**
   - Inspect `OSJS/` and confirm the expected basic runtime/build files exist.
   - Confirm excluded donor application directories are absent, especially application content originating from donor `src/packages/`.
   - Search copied configuration/source for references to excluded Nameless SCADA applications; remove only references that would load or require them.
   - Confirm the task diff does not contain unrelated application code or changes outside the authorized task scope.

2. **Clean dependency/build test**
   - Perform the donor-appropriate dependency installation from a clean state.
   - Run package discovery/build from a clean state.
   - Build must complete without unresolved imports, missing package references, or build-time references to excluded donor applications.
   - Do not accept a previously generated build artifact as proof.

3. **Fresh runtime smoke test**
   - Start the OS.js runtime from the new `OSJS/` tree.
   - Confirm the process remains running long enough to establish successful startup rather than exiting immediately after launch.
   - Request the web root and verify a successful response and desktop boot.
   - Verify the desktop shell, panel/taskbar, desktop area, and normal base interaction render without fatal client errors.

4. **Negative application-presence test**
   - Inspect package discovery/runtime package metadata and verify excluded Nameless SCADA applications are not registered.
   - Inspect launcher/start-menu/application listings and verify excluded applications are absent.
   - Verify no excluded application is automatically started.
   - Inspect runtime/server output for attempts to load excluded application modules or application-specific providers.

5. **Restart test**
   - Stop the runtime cleanly.
   - Start it again without repairing files or regenerating the source tree manually.
   - Repeat the web-root and desktop-load smoke check.
   - Confirm the second start produces the same clean basic desktop state.

6. **Evidence requirement**
   - Record the exact commands used for dependency installation, build, start, and verification.
   - Record the runtime URL used for the smoke test.
   - Record concise evidence showing the build succeeded, the desktop loaded, excluded applications were absent, and restart succeeded.
   - A task is not complete from file-copy evidence alone.

### Dependencies

- Access to the Nameless SCADA donor repository.
- Human promotion before execution.

### Sizing Assessment

- One primary outcome: establish the runnable basic desktop.
- Three bounded implementation actions: copy the required base, exclude donor applications, make minimum boot-reference adjustments.
- Verification is one end-to-end workflow containing static, clean-build, runtime, negative-presence, and restart checks.
- Size: 4; acceptable as one JR task. Split only if execution reveals an independent blocking problem.

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

1. Required donor color-scheme/theme files are present under `OSJS/` and are referenced by the basic desktop configuration without introducing donor applications.
2. A clean build/start succeeds and the intended donor color scheme/theme loads without missing-asset, stylesheet, or theme-resolution errors.
3. After restart, the same theme remains available/applied and the basic desktop remains free of donor application packages.

### Test and Verification

Use one end-to-end theme verification workflow:

1. Compare the selected theme/color-scheme files against the donor source and confirm only required basic-desktop material was copied.
2. Run a clean build and fail the task on missing theme imports, missing stylesheets, missing assets, or package-discovery errors.
3. Start the desktop and verify the selected scheme/theme is actually resolved at runtime, not merely present on disk.
4. Verify the desktop renders the expected panel/window/desktop styling with no fatal client-side theme errors.
5. Verify excluded Nameless SCADA applications remain absent from package discovery and launcher/menu entries.
6. Restart the runtime and confirm the theme still resolves and the desktop still renders correctly.
7. Record build/start commands and concise evidence of successful theme resolution before completion.

### Dependencies

- OSJS-001.
- Human promotion before execution.

### Sizing Assessment

- One primary outcome: carry over the basic desktop appearance.
- Verification is one end-to-end workflow covering source comparison, clean build, runtime rendering, exclusion check, and restart.
- Size: 3; good JR task.

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

1. Exactly one OS.js management/UI port is configured for this base desktop and the runtime binds successfully to it.
2. The desktop is reachable through the configured port after a clean start and again after restart, with no accidental second OS.js management listener introduced by this task.
3. Runtime configuration and project documentation record the same chosen port value.

### Test and Verification

Use one endpoint verification workflow:

1. Before configuration, verify the chosen port is not already occupied in the execution environment.
2. Start the OS.js runtime and verify it binds successfully to the chosen management/UI port.
3. Request the configured endpoint and verify a successful response and desktop load.
4. Inspect listening sockets/process output and verify the OS.js management endpoint is not simultaneously exposed on an unintended second port introduced by this task.
5. Stop and restart the runtime; verify the same configured port is used and remains reachable.
6. Compare the runtime/configured value against every documentation location changed by this task; values must match exactly.
7. Record the chosen port, commands/checks used, successful endpoint evidence, listener evidence, and restart evidence before completion.

### Dependencies

- OSJS-001.
- Human promotion before execution.

### Sizing Assessment

- One primary outcome: establish the documented OS.js management endpoint.
- Three bounded implementation actions: choose, configure, document.
- Verification is one endpoint workflow covering availability, bind, reachability, listener uniqueness, restart, and documentation consistency.
- Size: 3; good JR task.

## Sequence

```text
OSJS-001 basic desktop
    ↓
OSJS-002 donor color schemes
    ↓
OSJS-003 management port + documentation
```

These are planning tasks only until human promotion moves a selected task into `workflow/active_work/`.
