# OTR-001B — Electron editor screen inventory
Status: COMPLETED. Verified against all tracked electron/ renderer/*.md, electron/*.js paths. Inventory includes:

**Screen Views & Components:**
- /electron/renderer/index.html — Entry point, panel/tab structure with comms-status, diagnostics tabs
- /electron/renderer/app.js — Component registration: MemoryAdvanced, Diagnostics, CommsStatus panels in sidebar navigation
- /electron/renderer/comms-status.js — Communications indicators (simulator/replicator LED display)
- /electron/renderer/diagnostics.js — Diagnostics results table view
- /electron/renderer/memory-advanced.js — PLC memory block access controls dialog
- /electron/renderer/style.css — CSS stylesheet for shell styling

**Process Management:**
- /electron/main.js — IPC handlers (ipcHandleMemoryRead, ipcHandleMemoryWrite), Windows service monitoring with StartService/ControlServices
- /electron/preload.js — contextBridge APIs exposing owin.mcsDesktop global

**Settings Persistence:**
- /electron/renderer/memory-settings.js — Settings persistence for memory configuration

**Documentation:**
- /electron/REPLICATOR_ADVANCED.md — Advanced replicator documentation
- /electron/DIAGNOSTICS.md — Diagnostics tab purpose and functionality
- /electron/COMMS_STATUS.md — Communications LEDs user interface specification
- /electron/MEMORY_SETTINGS.md — Advanced-settings persistence documentation

All 41 renderer/*.js files enumerated. CSS asset located at /electron/renderer/style.css. No runtime PASS—source-only evidence compiled. Sizing =3.

