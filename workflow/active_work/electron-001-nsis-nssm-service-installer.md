# ELECTRON-001 — NSIS + NSSM Service Installer

Status: ACTIVE

## Outcome

Produce one assisted Windows Setup.exe for the standalone Electron deployment. The installer must install the Electron UI, MMA2, Simulator runtime, Replicator runtime, and NSSM; provide an embedded service-configuration wizard page; register selected backend runtimes as Windows services; start selected services after install; register a normal Windows uninstaller; remove the NSSM services during uninstall; and act as the maintenance entry point for an existing installation.

## Required behavior

- Electron remains a normal desktop application, not a service.
- Backend services are managed by NSSM.
- Service names:
  - `MCS-MMA2`
  - `MCS-Simulator`
  - `MCS-Replicator`
- Assisted NSIS installer, per-machine/elevated installation.
- When MCS Modbus Toolkit is not installed, Setup presents an Install option.
- When an existing per-machine installation is detected, Setup presents Repair / reconfigure and Uninstall options instead of behaving like a blind fresh install.
- Repair targets the detected installation directory, reinstalls application files, preserves existing runtime configuration, and allows the backend service selection to be reviewed.
- Repair defaults the service checkboxes from the services currently installed so it does not blindly enable previously unselected runtimes.
- Uninstall from the maintenance page launches the installed Windows uninstaller; the normal uninstaller remains registered in Windows Installed Apps.
- Embedded wizard page with checkboxes for MMA2, Simulator, Replicator and a Start services after installation option.
- On a fresh install, MMA2 defaults selected; Simulator and Replicator default selected.
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

1. On a clean Windows machine, Setup displays Install and proceeds through the assisted wizard.
2. Service Configuration page is visible during install/repair.
3. Install completes without a visible command shell.
4. Selected services appear in Windows Services and use Automatic startup.
5. Selected services start successfully when requested.
6. Electron opens without spawning duplicate runtime processes.
7. Windows Installed Apps contains MCS Modbus Toolkit with an uninstaller.
8. Running the same Setup.exe after installation detects the existing per-machine installation and offers Repair / reconfigure and Uninstall.
9. Repair uses the existing installation directory and preserves runtime/configuration data.
10. Repair defaults the service selection to the services that currently exist and does not create duplicate Windows services.
11. Choosing Uninstall from Setup launches the installed uninstaller.
12. Uninstall stops/removes MCS-MMA2, MCS-Simulator, and MCS-Replicator and removes application files.

## Verification status

The maintenance-mode implementation may be authored outside Windows, but the build and behavioral acceptance above remain Windows-only gates. Do not mark this task complete or archive it until `npm run dist:win` and the clean-install / repair / uninstall checks have been executed successfully on Windows.

## Scope guard

This task covers Windows installer/service packaging only. It does not port the full Simulator/Replicator editor UIs into Electron.