# Active Work — OS.js Base Web App Initialization

Human-promoted from `planning/microtask/osjs-base-webapp-init.md`.

Execute in sequence. Complete and verify each task before continuing to the next.

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

1. Static scope check: expected base files exist; excluded donor application directories are absent; no load/require references remain; diff stays in authorized scope.
2. Clean dependency/build test: install from clean state, run discovery/build, and fail on unresolved imports or excluded-app references. Existing generated artifacts are not proof.
3. Fresh runtime smoke test: start from `OSJS/`, confirm process remains running, request web root, and verify desktop shell/panel/taskbar/desktop interaction without fatal client errors.
4. Negative application-presence test: verify excluded apps are absent from discovery, launcher/menu, autostart, runtime requests, and server output.
5. Restart test: stop cleanly, start again without manual repair, repeat web-root/desktop checks, and confirm the same clean basic desktop state.
6. Evidence: record exact install/build/start/verification commands, runtime URL, successful build/load evidence, excluded-app absence, and restart success. File-copy evidence alone is insufficient.

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

1. Required donor color-scheme/theme files are present under `OSJS/` and referenced without introducing donor applications.
2. A clean build/start succeeds and the intended theme loads without missing-asset, stylesheet, or theme-resolution errors.
3. After restart, the same theme remains available/applied and donor application packages remain absent.

### Test and Verification

1. Compare selected theme files against donor source and confirm only required basic-desktop material was copied.
2. Run a clean build and fail on missing theme imports, stylesheets, assets, or package-discovery errors.
3. Start the desktop and verify the scheme/theme resolves at runtime, not merely on disk.
4. Verify expected panel/window/desktop styling with no fatal client-side theme errors.
5. Verify excluded Nameless SCADA applications remain absent.
6. Restart and confirm the theme still resolves and renders correctly.
7. Record build/start commands and concise theme-resolution evidence.

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

1. Exactly one OS.js management/UI port is configured and the runtime binds successfully to it.
2. The desktop is reachable through that port after a clean start and restart, with no unintended second OS.js management listener introduced.
3. Runtime configuration and project documentation record the same chosen port.

### Test and Verification

1. Verify the chosen port is unoccupied before configuration.
2. Start OS.js and verify successful bind to the chosen port.
3. Request the endpoint and verify successful response and desktop load.
4. Inspect listeners/process output and verify no unintended second management port.
5. Stop/restart and verify the same configured port remains reachable.
6. Compare runtime/config value with every documentation location changed; values must match exactly.
7. Record chosen port, commands/checks, endpoint evidence, listener evidence, and restart evidence.

## Execution Sequence

```text
OSJS-001 basic desktop
    ↓
OSJS-002 donor color schemes
    ↓
OSJS-003 management port + documentation
```
