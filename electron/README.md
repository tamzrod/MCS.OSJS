# MCS Modbus Toolkit — Electron desktop package

This folder is a standalone Electron deployment target. It intentionally does **not** load the OS.js desktop. The application opens one normal desktop window with Simulator / Replicator tabs and bundles the native MCS runtimes beside the Electron app.

## Windows package layout

Before building the installer, place these files under `electron/bin/`:

```text
electron/bin/
├── nssm.exe
├── mma2.exe
├── modbus-simulator-runtime.exe
└── modbus-replicator-runtime.exe
```

`nssm.exe` is required for the installer build to be useful. The NSIS installer checks for it during installation and stops with a clear error if it was not packaged.

The existing OS.js Simulator and Replicator renderers are not yet copied into the standalone renderer. The current Electron shell exists so the Windows packaging/service path can be compiled and verified independently.

## Build on Windows

From PowerShell or Command Prompt:

```text
cd electron
npm install
npm start
```

Unpackaged development mode launches the three runtime binaries as Electron child processes when they are present.

To create an unpacked Windows application directory:

```text
npm run pack:win
```

To create the assisted Windows installer:

```text
npm run dist:win
```

Output is written to:

```text
electron/dist/
```

Installer artifact:

```text
MCS-Modbus-Toolkit-0.1.0-Setup.exe
```

## Installer wizard

The generated installer is a per-machine/elevated NSIS wizard. It includes an embedded **Service Configuration** page with these options selected by default:

```text
[x] MMA2 - shared Modbus memory appliance
[x] Simulator runtime
[x] Replicator runtime
[x] Start selected services after installation
```

Selected runtimes are registered through NSSM as automatic Windows services:

```text
MCS-MMA2
MCS-Simulator
MCS-Replicator
```

Simulator and Replicator are configured to depend on `MCS-MMA2` when MMA2 is selected.

Service working data is placed under:

```text
C:\ProgramData\MCS Modbus Toolkit\runtime
```

The installed Electron application does **not** spawn duplicate backend processes. It reads service state from Windows and closing Electron does not stop the services.

## Uninstall

Windows Installed Apps receives the normal MCS Modbus Toolkit uninstaller from NSIS. Uninstall performs service cleanup first:

```text
stop/remove MCS-Replicator
stop/remove MCS-Simulator
stop/remove MCS-MMA2
```

The application files and shortcuts are then removed by the normal NSIS uninstall flow. Runtime data under `C:\ProgramData\MCS Modbus Toolkit\runtime` is intentionally not deleted by this first implementation so configuration/data is not destroyed accidentally.

## Architecture

```text
Windows
├── MCS-MMA2 service          -> mma2.exe
├── MCS-Simulator service     -> modbus-simulator-runtime.exe
├── MCS-Replicator service    -> modbus-replicator-runtime.exe
└── MCS Modbus Toolkit.exe    -> normal Electron UI
    ├── Simulator tab
    ├── Replicator tab
    └── Diagnostics tab
```

No OS.js wallpaper, taskbar, desktop icons, application menu, or desktop workspace is part of this deployment.
