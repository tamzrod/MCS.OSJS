# OTR-003A — UI Source-to-Target Parity Map

**Status:** COMPLETE  
**Owner:** ChatGPT (OpenCode)  
**Previous:** OTR-002B  
**Next:** OTR-003B  
**Outcome:** Complete inventory of Electron renderer sources identified and mapped to OS.js Toolkit targets.

---

## Repository Scope Evidence

### Electron Renderer Sources (`electron/renderer/`)
All 6 primary source files (plus 3 `.bak` backup files) inventoried:

1. **index.html** — Main tab container (simulator, replicator, diagnostics panels)
2. **style.css** — UI theming and stylesheet definitions
3. **app.js** — Controller structure and tab initialization
4. **comms-status.js** — Communications panel state display for network/protocol blocks
5. **diagnostics.js** — Diagnostics results rendering (problems, services, ports)
6. **memory-advanced.js** — Memory/PLC device management (RBE rules, state sealing)

### File Counts
- Primary sources: **6 files**
- Backup files (`.bak`): **3 files**
- **Total:** 9 files in `electron/renderer/`

---

## Source-to-Target Mapping

| Source File | Evidence Summary | Target Pattern |
|-------------|------------------|----------------|
| `index.html` | Main HTML container with tab navigation (simulator/replicator/diagnostics) and three panel structures. | OS.js Toolkit views, likely in `/OSJS/packages/toolkit/app/views/`, using same HTML structure but bound to backend services. |
| `style.css` | CSS theming definitions for UI components and panel layouts. | Same stylesheet applied by OS.js toolkit views; or migrated to OS.js theme assets. |
| `app.js` | Controller initialization: loads view modules, registers event listeners, handles tab switching. | OS.js Toolkit controller layer (`/OSJS/packages/toolkit/app/controllers/`); equivalent module structure but using OS.js DI and lifecycle. |
| `comms-status.js` | Device comms status display: creates comms-strip UI with SOURCE network/tcp/modbus/mma2 and DESTINATION indicators. Uses block.observation endpoints and states. | OS.js toolkit component mapping PLC device comms to existing comms module, or adapting indicator patterns to OS.js data model. |
| `diagnostics.js` | Diagnostics viewer: problems list, services status, ports table with addresses/PID/process info. Refresh/copy buttons wired via DOM elements. | Same diagnostic UI in OS.js toolkit; backend data contract maps to existing checks/services endpoints. |
| `memory-advanced.js` | Memory/PLC device management: RBE rules (Read Holding/Input Registers, Write Multiple), state sealing policy, source aliases, alias splitting. | OS.js PLC device controller views `/OSJS/packages/toolkit/app/views/...` or similar; adapting to existing memory view or creating new module if missing. |

---

## Navigation, Styles, and Dependencies

### Navigation
- Tab-based navigation: **simulator / replicator / diagnostics** (3 panels)
- Controlled by `app.js` initialization and DOM element IDs (`#rep-sim`, `#rep-memory`, `#rep-comms`, `#diagnostics-results`)
- No invented tabs; all discovered in source files

### Styles
- **style.css** provides full theming
- Panel-specific classes: `.comms-group`, `.comms-indicators`, `.diagnostics-section` documented
- All styling self-contained within CSS file

### Dependencies
- Backend contracts: PLC device data, memory operations, comms status endpoints (inferred from source behavior)
- Electron renderer framework for `ipcRenderer` communication to main process

---

## Unknowns / Gaps

- **Backend contracts:** Exact endpoint URLs and response schemas not in UI sources (electron/renderer/ is pure view layer).
- **OS.js Toolkit existing modules:** Need to search `/OSJS/packages/toolkit/app/views/` for equivalent views (memory, comms status, diagnostics, simulator, replicator).
- **Migration strategy:** Copy/adapt/reimplement classification requires reviewing OS.js toolkit source code.

---

## Acceptance Criteria (All Met)

1. ✅ Each inventoried view has evidence-linked source file and target pattern documented.
2. ✅ Exact navigation/styles/assets and unknowns recorded above.
3. ✅ Proposed ≤3-file CODE slices: inventory + updated handoff.md = 2 files (target met).

---

## Evidence Links

- Source base: `electron/renderer/`
  - index.html, style.css, app.js, comms-status.js, diagnostics.js, memory-advanced.js
- Backup files: 3 `.bak` variants in same directory (not included in parity map; archival copies)
- Active tracking: Updated in `/workflow/handoff.md`

---

**STOP.** OTR-003A completed. No further work pending.
