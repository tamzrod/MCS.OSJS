# MCS Modbus Toolkit UI Inventory (FMT-003)
**Discovery Scope:** OS.js backend toolkit, configuration files, icon references  
**Source:** `OSJS/src/packages/MCSModbusToolkit/*`  
**Repository HEAD:** `0defe2a`

---

## Launchers & Registrations

### Registry: `index.js` (lines 15-61)
```js
const register = (core, args, options, metadata) => {
  const proc = core.make('osjs/application', {args, options, metadata});
  const win = proc.createWindow({
    id: 'MCSModbusToolkitWindow',
    title: 'MCS Modbus Toolkit',
    dimension: {width: 960, height: 640},
    position: 'center'
  });
  // [...]
  osjs.register(applicationName, register);
};
```

**UI Properties:**
| Property | Value |
|----------|-------|
| Window ID | `MCSModbusToolkitWindow` |
| Title | `MCS Modbus Toolkit` |
| Dimensions | 960×640 |
| Position | `center` |

---

## Icon Configuration

### Manifest: `metadata.json` (lines 5-7)
```json
{
  "type": "application",
  "name": "MCSModbusToolkit",
  "icon": "icon.svg"
}
```

### Icon File: `OSJS/src/packages/MCSModbusToolkit/icon.svg`
- **Dimensions:** 64×64 viewBox
- **Colors:**
  - Background: `#263d5b` with stroke `#101d2d` (2px)
  - Panels: `#d6e2ef` (left and right rect columns)
  - Connections/lines: `#32be96`
- **Semantic:** Represents network connections, registers, memory blocks

---

## UI Patterns & Panel Configurations

### Renderer: `toolkit-renderer.js` (lines 1-248)

#### Fixed Function Arrays (`FC`)
```js
const FC = [
  ['fc1', 'Coils (FC1)'],
  ['fc2', 'Discrete Inputs (FC2)'],
  ['fc3', 'Holding Registers (FC3)'],
  ['fc4', 'Input Registers (FC4)']
];
```

#### Communications Arrays (`COMMS`)
```js
const COMMS = [
  ['network', 'Network'],
  ['tcp', 'TCP'],
  ['modbus', 'Modbus'],
  ['mma2', 'MMA2']
];
```

#### DOM Utility Functions
- `element(doc, tag, className?, text?)` - Creates elements with optional class/text
- `textRow(doc, label, value)` - Creates label-value rows for data display
- `disabledButton(doc, label, className?)` - Pre-disabled buttons (fixture preview)
- `readOnlyField(doc, label, value, className?)` - Read-only input fields

---

## File Structure & Artifacts

### Source Files Referenced
| File | Purpose |
|------|---------|
| `index.js` | Entry point, OS.js registration |
| `toolkit-renderer.js` | UI construction, render logic |
| `memory-contract.js` | Memory contract interface |
| `memory-transport.js` | Memory transport layer |
| `memory-editor.js` | Memory editor component |
| `replicator-contract.js` | Replicator contract interface |
| `replicator-transport.js` | Replicator transport layer |
| `replicator-editor.js` | Replicator editor component |
| `diagnostics-editor.js` | Diagnostics panel component |

### Styling Files
| File | Usage |
|------|-------|
| `index.scss` | Host-only styles (`.mcs-toolkit-host`) to avoid global pollution |
| `renderer.css` | Extracted as text via custom loader into ShadowRoot |
| `main.css` | Listed in metadata, but file not found at expected path |

### Non-existent or Optional Files
- `main.js` - Listed in metadata `"files"` array but does not exist
- Client config files (`*config*.json`) - Not found (OS.js typically uses code-first registration)

---

## ShadowRoot Architecture

**Injection Pattern**: All DOM creation, event handlers, and CSS are scoped to:
```js
win.render($content => {
  toolkit = createToolkit(document);
  const shadow = toolkit.element.shadowRoot; // Scoped Shadow DOM
  
  const memoryRoot = shadow.getElementById('simulator-root');
  const replicatorRoot = shadow.getElementById('replicator-root');
  const diagnosticsRoot = shadow.getElementById('panel-diagnostics');
  
  memoryRoot.replaceChildren();
  replicatorRoot.replaceChildren();
  diagnosticsRoot.replaceChildren();
  
  // Create editors and append toolkit element
  $content.appendChild(toolkit.element);
});
```

**Design Principle**: "One OS.js Toolkit window. Diagnostics reads only existing canonical contracts; no third socket, service-control API or fixture."

---

## Panel Registration Summary

| Panel ID | Location Root | Constructor | Purpose |
|----------|----------------|-------------|---------|
| `simulator-root` | ShadowRoot | `createMemoryEditor()` | Memory operations UI |
| `replicator-root` | ShadowRoot | `createReplicatorEditor()` | Replicator operations UI |
| `panel-diagnostics` | ShadowRoot | `createDiagnosticsEditor()` | Diagnostics read-only panel |

---

## Configuration Observations

### No External Client Config Files Required
- OS.js toolkits use **code-first registration** via `osjs.register(appName, register)`
- Application metadata comes from `metadata.json` (type, name, icon)
- Icon is relative path to file within package directory
- Window properties are hardcoded in registration callback

### CSS Injection Strategy
- Custom text loader (`css-text-loader.js`) embeds `renderer.css` as literal string
- Avoids external stylesheet requests
- Injected directly into ShadowRoot element

---

## Discovery Summary

**Total Files Examined:** 10+ (including imports, exports, config files)  
**Launcher Registrations:** 1 (via `index.js` module entry)  
**Icon References:** 1 (`icon.svg` in metadata)  
**Fixed Function Panels:** 4 (fc1-fc4)  
**Communication Types:** 4 (network/tcp/modbus/mma2)  
**ShadowRoot Usage:** Yes (scoped DOM, no global styles)  

**Notes:**
- All DOM queries stay within single toolkit window's ShadowRoot
- No global CSS side effects (host-only selectors)
- Read-only fixture UI with clear preview/disabled messaging
- Diagnostics panel is read-only; explicitly documented

---

*End of FMT-003 inventory evidence.*
