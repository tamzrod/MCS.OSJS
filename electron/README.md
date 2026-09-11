# MCS Modbus Toolkit — Electron desktop package

This folder is a standalone Electron deployment target. It intentionally does **not** load the OS.js desktop. The application opens one normal desktop window with Simulator / Replicator tabs and can bundle the native MCS runtimes beside the Electron app.

## Current scope

The Electron shell, Windows installer config, runtime process manager, runtime status/log panel, and two application tabs are present. The shell expects the native Windows binaries under `electron/bin/`:

```text
electron/bin/
├── mma2.exe
├── modbus-simulator-runtime.exe
└── modbus-replicator-runtime.exe
```

Missing binaries do not prevent the Electron UI itself from starting; the Diagnostics tab will show the runtimes as stopped. The existing OS.js Simulator and Replicator renderers are not yet copied into this standalone renderer. This separation is deliberate so the Windows packaging path can be compiled/tested before replacing the placeholders with the live editors.

## Build on Windows

From PowerShell or Command Prompt:

```text
cd electron
npm install
npm start
```

That launches the unpackaged desktop window.

To create an unpacked Windows application directory:

```text
npm run pack:win
```

To create the installer:

```text
npm run dist:win
```

Output is written to:

```text
electron/dist/
```

The NSIS installer name is:

```text
MCS-Modbus-Toolkit-0.1.0-Setup.exe
```

## Windows runtime binaries

The Go runtimes still need Windows builds before the packaged app can run the real Simulator/Replicator stack. Build the project runtimes for Windows and copy/rename the resulting executables into `electron/bin/` using the names above before `npm run dist:win`.

Electron sets `MCS_DATA_ROOT` to its per-user application data directory and starts MMA2, Simulator runtime, and Replicator runtime as hidden child processes. Closing the Electron app stops those child processes.

## Architecture

```text
MCS Modbus Toolkit.exe
└── one Electron window
    ├── Simulator tab
    ├── Replicator tab
    └── Diagnostics tab
         ├── mma2.exe
         ├── modbus-simulator-runtime.exe
         └── modbus-replicator-runtime.exe
```

No OS.js wallpaper, taskbar, desktop icons, application menu, or desktop workspace is part of this deployment.
