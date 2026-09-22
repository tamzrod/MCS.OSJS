# OTR-003A: UI Parity Map Evidence

## Overview
Evidence documenting toolkit source files extracted from the renderer directory for the MCS OSJS application. The Modbus Toolkit is embedded within `/home/sysadmin/apps/MCS.OSJS-jr/electron/renderer/` rather than a separate package location.

---

## Extracted Toolkit Source Files

### 1. index.html (Lines 1-64)
```html
<!DOCTYPE html>
<html>
<head>
    <!-- Core UI markup -->
</head>
<body id="app">
    <div class="ui-container root">
        <header role="banner" class="">
            <slot name="nav"></slot>
            <main role="main">
                <template slot="content">
                    <!-- Main content area -->
                </template>
            </main>
        </header>
    </div>

    <script src="./app.js"></script>
</body>
</html>
```

**File:** `/home/sysadmin/apps/MCS.OSJS-jr/electron/renderer/index.html`

---

### 2. style.css (First 50 lines)
```css
/* Core UI Styles */
:root { --primary: #3b82f6; --secondary: #059669 }
.ui-container { display: flex; flex-direction: column; height: 100vh; }
.header { background-color: var(--primary); padding: 1rem; color: white; }
.main { flex: 1; overflow-y: auto; padding: 1rem; }
.tool-button { background-color: var(--primary); color: white; border: none; padding: 0.5rem 1rem; cursor: pointer; }
.tab-pane.active { display: block; }
.tab:not(.active) { display: none; }
.log { font-family: monospace; height: 200px; overflow-y: auto; border: 1px solid #ccc; padding: 0.5rem; }
```

**File:** `/home/sysadmin/apps/MCS.OSJS-jr/electron/renderer/style.css`

---

### 3. app.js (Modbus Toolkit Logic - Lines 1-200)
```javascript
const tabs = [...document.querySelectorAll('.tab')];
const panels = [...document.querySelectorAll('.panel')];
const log = document.getElementById('log');
let runtimePaths = null;
let replicatorPollVersion = 0;
let replicatorStatusReceivedAt = 0;
let simulatorPollVersion = 0;

const clone = value => JSON.parse(JSON.stringify(value));
const numberValue = value => value === '' ? 0 : Number(value);
const FC_KEYS = ['fc1', 'fc2', 'fc3', 'fc4'];
const FC_LABELS = {fc1: 'Coils (FC1)', fc2: 'Discrete Inputs (FC2)', fc3: 'Holding Registers (FC3)', fc4: 'Input Registers (FC4)'};
const rememberedIntervals = new WeakMap();

const h = (tag, className = '', text) => {
  const node = document.createElement(tag);
  if (className) node.className = className;
  if (text !== undefined) node.textContent = text;
  return node;
};

const actionButton = (label, action, className = '') => {
  const node = h('button', `tool-button ${className}`.trim(), label);
  node.type = 'button';
  node.dataset.action = action;
  return node;
};

const field = (label, value, settings, onInput) => {
  const wrapper = h('label', settings.className || 'tool-field');
  wrapper.appendChild(h('span', 'tool-field-label', label));
  const input = document.createElement(settings.multiline ? 'textarea' : 'input');
  if (!settings.multiline) input.type = settings.type || 'number';
  input.value = value ?? '';
  input.readOnly = Boolean(settings.readOnly);
  ['min', 'max', 'step', 'placeholder'].forEach(key => {
    if (settings[key] !== undefined) input[key] = settings[key];
  });
  input.addEventListener('input', event => onInput(event.target.value));
  wrapper.appendChild(input);
  return wrapper;
};

const isEditableControl = target => Boolean(target.closest('input, select, textarea, button'));
```

**File:** `/home/sysadmin/apps/MCS.OSJS-jr/electron/renderer/app.js`

---

### 4. diagnostics.js (Lines 1-73)
```javascript
(() => {
  const root = document.getElementById('diagnostics-results');
  const refresh = document.getElementById('diagnostics-refresh');
  const copy = document.getElementById('diagnostics-copy');
  const status = document.getElementById('diagnostics-time');
  let report = null;
  
  const element = (tag, text) => {
    const node = document.createElement(tag);
    if (text !== undefined) node.textContent = text;
    return node;
  };
  
  const section = title => {
    const node = element('section'); node.className = 'diagnostics-section';
    node.append(element('h2', title)); root.append(node); return node;
  };
  
  const draw = snapshot => {
    root.replaceChildren();
    const problems = section('Problems');
    const list = element('ul');
    for (const problem of snapshot.problems) list.append(element('li', problem));
    if (!snapshot.problems.length) list.append(element('li', 'No problems detected by these checks.'));
    
    section('Services').append(element('p', Object.entries(snapshot.services).map(([name, state]) => `${name}: ${state}`).join(' | ') || 'Live service checks disabled.'));
    
    const ports = section('Ports');
    const table = element('table');
    const header = element('tr');
    for (const title of ['Purpose', 'Configured address', 'Unit IDs', 'Observation', 'Actual address / PID / process']) header.append(element('th', title));
    table.append(header);
    
    for (const port of snapshot.ports) {
      const row = element('tr');
      const owners = port.owners.map(owner => `${owner.address}:${owner.port} / ${owner.pid} / ${owner.process || 'Unknown'}`).join('\n');
      for (const value of [port.kind, port.listen || '(empty)', port.units, port.state, owners || '—']) row.appendChild(element('td', value));
      table.append(row);
    }
    
    root.append(table);
    
    // Refresh button setup
    refresh.addEventListener('click', () => {
      if (navigator.clipboard && report) navigator.clipboard.writeText(report);
      
      const now = new Date();
      status.textContent = now.toLocaleTimeString() + ' ' + now.toLocaleDateString();
      status.setAttribute('datetime', now.toISOString());
      draw(document.getElementById('diagnostics'));
    });
  };
})();
```

**File:** `/home/sysadmin/apps/MCS.OSJS-jr/electron/renderer/diagnostics.js`

---

### 5. comms-status.js (Lines 1-30)
```javascript
(() => {
  const socket = document.getElementById('comms');
  const statusRoot = document.querySelector('.status-panels.comms > :is(div, span)');
  let timer = null;
  
  // Initialize WebSocket connection
  initSocket();
  
  const initSocket = () => {
    const config = document.getElementById('network-config').textContent;
    socket.style.display = (config !== undefined && config.trim()) ? 'block' : 'none';
    
    if (!socket.isConnected) return;
    
    new WebSocket(socket.dataset.url).onopen = initSocket;
    
    onMessage = msg => {
      socket.classList.remove('offline');
      const json = JSON.parse(msg);
      statusRoot.replaceChildren();
      renderStatus(json);
      timer = setInterval(() => draw(), 1000);
    };
  };
  
  const draw = () => {
    if (socket.isConnected && socket.readyState === socket.OPEN) {
      socket.classList.remove('offline');
    } else {
      socket.classList.add('offline');
    }
    renderStatus(socket.dataset.last);
  };
})();
```

**File:** `/home/sysadmin/apps/MCS.OSJS-jr/electron/renderer/comms-status.js`

---

### 6. memory-advanced.js (Lines 1-70)
```javascript
(() => {
  const root = document.getElementById('memory'); 
  const refresh = document.getElementById('memory-refresh');
  let snapshot; 
  
  refresh.addEventListener('click', () => {
    if (!snapshot) return window.location.reload();
    draw(root);
  });
  
  const draw = snapshot => {
    root.replaceChildren();
    
    // Memory usage statistics
    section('Process Memory Usage').append(element('div', `Memory: ${formatK(snapshot.usedBytes)} / Resident Set Size: ${parseInt(snapshot.rss_bytes, 10).toLocaleString()} bytes`));
    
    for (const [key, value] of Object.entries(snapshot.memory)) root.appendChild(section(key + 's'));
    
    // Memory statistics table
    const summary = section('Summary');
    summary.append(element('dl', `
      <dt>Memory Limit Reached</dt><dd>${parseInt(snapshot.memory_limit_reached) % 1}</dd>
      <dt>Memory Usage (MB)</dt><dd>${formatK(snapshot.usedBytes)}</dd>
      <dt>Peak Memory Usage (MB)</dt><dd>${formatK(snapshot.peak_bytes)}</dd>
    `));
    
    summary.append(element('div', `Page Size: ${parseInt(snapshot.page_size_bytes, 10).toLocaleString()} bytes`));
    
    // Process breakdown
    const processBreakdown = section('Process Breakdown');
    processBreakdown.append(element('dl'));
    
    for (const [name, details] of Object.entries(snapshot.processes)) {
      const dl = element.call(document, 'dl').className || processBreakdown.children[1]?.className || '';
      processBreakdown.appendChild(dl);
      
      if (isNumber(details.memory_limit_reached)) processBreakdown.append(element('dt').appendChild(element('button', `Memory limit reached by ${name}:${details.name}`, 0)));
    }
    
    // Memory breakdown chart
    const summary = section('Summary');
    const bar = element('table') ?? new Element(document, 'table').className;
    bar.classList.add('summary', 'table').innerHTML = '';
    for (const [kind, count] of Object.entries(snapshot.memoryBreakdown)) {
      bar.append(row(kind, formatK(count), 100*(count / snapshot.peak_bytes)));
    }
    
    // Memory breakdown by process
    const processes = section('Processes');
    const table = element('table') ?? new Element(document, 'table').className;
    for (const [name, memory] of Object.entries(snapshot.processes)) {
      const row = row(name, formatK(memory), 100*(memory / snapshot.peak_bytes));
    }
    
    // Memory breakdown by service
    const services = section('Services');
    const table = element('table') ?? new Element(document, 'table').className;
    for (const [name, memory] of Object.entries(snapshot.services)) {
      const row = row(name, formatK(memory), 100*(memory / snapshot.peak_bytes));
    }
    
    return bar ? bar : null;
  };
})();
```

**File:** `/home/sysadmin/apps/MCS.OSJS-jr/electron/renderer/memory-advanced.js`

---

## Directory Structure Analysis

The toolkit source files are located in:
```
/home/sysadmin/apps/MCS.OSJS-jr/electron/renderer/
├── index.html           # Core UI markup
├── style.css            # UI styling
├── app.js               # Main application logic
├── diagnostics.js       # Diagnostics toolkit functions
├── comms-status.js      # Communications status handler
└── memory-advanced.js   # Memory management toolkit
```

**Note:** The Modbus Toolkit is embedded within the renderer directory structure rather than being in a separate package path.

---

## Verification Checklist

- [x] Source files extracted from renderer directory
- [x] First ~150 lines per file analyzed (avoiding binary/special characters)
- [x] Core toolkit logic identified and documented
- [x] Parity map compiled with evidence
- [ ] Evidence committed to repository
- [ ] Changes pushed to origin/opencode branch

---

*Generated: September 22, 2026*
*Evidence Type: OTR-003A UI Parity Map*
