# ELECTRON-001 — NSIS + NSSM Service Installer

Status: ACTIVE

## Outcome

Produce one assisted Windows Setup.exe for the standalone Electron deployment. The installer must install the Electron UI, MMA2, Simulator runtime, Replicator runtime, and NSSM; provide an embedded service-configuration wizard page; register selected backend runtimes as Windows services; start selected services after install; register a normal Windows uninstaller; and remove the NSSM services during uninstall.

## Required behavior

- Electron remains a normal desktop application, not a service.
- Backend services are managed by NSSM.
- Service names:
  - `MCS-MMA2`
  - `MCS-Simulator`
  - `MCS-Replicator`
- Assisted NSIS installer, per-machine/elevated installation.
- Embedded wizard page with checkboxes for MMA2, Simulator, Replicator and a Start services after installation option.
- MMA2 defaults selected; Simulator and Replicator default selected.
- Runtime binaries and `nssm.exe` are packaged under Electron `resources/bin`.
- Services use automatic startup.
- Uninstaller stops/removes all MCS NSSM services before files are removed.
- Closing Electron must not stop installed Windows services.
- Packaged Windows Electron must not spawn duplicate backend child processes.
- Development mode may continue using direct child processes.

## Build acceptance

From `electron/` on Windows:

```text
npm install
npm run dist:win
```

Expected artifact:

```text
dist/MCS-Modbus-Toolkit-<version>-Setup.exe
```

Installer verification:

1. Setup displays normal assisted wizard.
2. Service Configuration page is visible.
3. Install completes without a visible command shell.
4. Selected services appear in Windows Services and use Automatic startup.
5. Selected services start successfully when requested.
6. Electron opens without spawning duplicate runtime processes.
7. Windows Installed Apps contains MCS Modbus Toolkit with an uninstaller.
8. Uninstall stops/removes MCS-MMA2, MCS-Simulator, and MCS-Replicator and removes application files.

## Scope guard

This task covers Windows installer/service packaging only. It does not port the full Simulator/Replicator editor UIs into Electron.