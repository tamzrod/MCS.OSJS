// UMIG-003: one-time visual adaptation of electron/renderer at 1c971b9.
// No Electron host, OS.js backend, network, file, or persistence calls here.
const donorCSS = require('!!./css-text-loader.js!./renderer.css');
const {snapshot, UNKNOWN} = require('./fixtures');

const FC = [
  ['fc1', 'Coils (FC1)'], ['fc2', 'Discrete Inputs (FC2)'],
  ['fc3', 'Holding Registers (FC3)'], ['fc4', 'Input Registers (FC4)']
];
const COMMS = [
  ['network', 'Network'], ['tcp', 'TCP'],
  ['modbus', 'Modbus'], ['mma2', 'MMA2']
];

// All DOM queries, handlers and CSS stay inside one Toolkit window's ShadowRoot.
const element = (doc, tag, className = '', text) => {
  const node = doc.createElement(tag);
  if (className) node.className = className;
  if (text !== undefined) node.textContent = String(text);
  return node;
};
const textRow = (doc, label, value) => {
  const row = element(doc, 'div');
  row.append(element(doc, 'strong', '', label), doc.createTextNode(' '), element(doc, 'span', '', value));
  return row;
};
const disabledButton = (doc, label, className = '') => {
  const node = element(doc, 'button', `tool-button ${className}`.trim(), label);
  node.type = 'button';
  node.disabled = true;
  node.title = 'Fixture preview only; no backend operation is available.';
  return node;
};
const readOnlyField = (doc, label, value, className = 'tool-field') => {
  const wrapper = element(doc, 'label', className);
  wrapper.appendChild(element(doc, 'span', 'tool-field-label', label));
  const input = element(doc, 'input');
  input.type = 'text';
  input.value = value === undefined || value === null ? '' : String(value);
  input.readOnly = true;
  input.setAttribute('aria-label', label);
  input.title = 'Sample value; not connected to configuration.';
  wrapper.appendChild(input);
  return wrapper;
};
const readOnlyCheckbox = (doc, label, checked) => {
  const wrapper = element(doc, 'label', 'tool-checkbox');
  const input = element(doc, 'input');
  input.type = 'checkbox';
  input.checked = Boolean(checked);
  input.disabled = true;
  wrapper.append(input, element(doc, 'span', '', label));
  return wrapper;
};
const disabledSelect = (doc, label, values, selected) => {
  const node = element(doc, 'select');
  node.disabled = true;
  node.setAttribute('aria-label', label);
  values.forEach(([value, title]) => {
    const option = element(doc, 'option', '', title);
    option.value = String(value);
    node.appendChild(option);
  });
  node.value = String(selected);
  return node;
};
const notice = (doc, text) => element(doc, 'div', 'tool-status', text);
const actions = (doc, labels) => {
  const bar = element(doc, 'div', 'tool-actions');
  labels.forEach(label => bar.appendChild(disabledButton(doc, label, label === 'Delete' ? 'tool-danger' : '')));
  return bar;
};
const editorActions = doc => {
  const bar = element(doc, 'div', 'editor-actions');
  bar.append(disabledButton(doc, 'Save & Apply', 'tool-primary'), disabledButton(doc, 'Discard'));
  return bar;
};
const sidebar = (doc, heading, device, subtitle) => {
  const side = element(doc, 'aside', 'tool-sidebar');
  side.append(element(doc, 'h2', '', heading), actions(doc, ['Add', 'Duplicate', 'Delete']));
  const list = element(doc, 'div', 'tool-list');
  const row = disabledButton(doc, '', 'tool-row selected');
  row.append(element(doc, 'strong', '', device.name), element(doc, 'span', '', subtitle));
  list.appendChild(row);
  side.appendChild(list);
  return side;
};
const runtimeRow = (doc, entries) => {
  const row = element(doc, 'div', 'runtime-row');
  entries.forEach(([name, status]) => row.append(element(doc, 'span', '', name), element(doc, 'strong', '', status)));
  return row;
};
const range = area => area.count ? `${area.start}-${area.start + area.count - 1}` : 'Unused';

const renderMemory = (doc, root, fixture) => {
  const device = fixture.memory.devices[0];
  const shell = element(doc, 'div', 'tool-layout');
  shell.appendChild(sidebar(doc, 'Devices', device, `Port ${device.mma2.port} / Unit ${device.mma2.unit_id} / FIXTURE`));
  const editor = element(doc, 'section', 'tool-editor');
  editor.append(element(doc, 'h2', '', 'Device Definition'), runtimeRow(doc, [['MMA2:', UNKNOWN], ['Simulation:', UNKNOWN]]));
  const identity = element(doc, 'div', 'tool-grid');
  identity.append(
    readOnlyField(doc, 'Name', device.name), readOnlyCheckbox(doc, 'Enabled', device.enabled),
    readOnlyField(doc, 'Listen Port', device.mma2.port), readOnlyField(doc, 'Unit ID', device.mma2.unit_id)
  );
  editor.appendChild(identity);
  const table = element(doc, 'div', 'fc-table memory-fc-table');
  const header = element(doc, 'div', 'fc-row fc-header');
  ['Area', 'Start', 'Count', 'Simulation', 'Interval (ms)', 'Address Range']
    .forEach(label => header.appendChild(element(doc, 'span', '', label)));
  table.appendChild(header);
  FC.forEach(([key, label]) => {
    const area = device.mma2[key];
    const interval = device.random_runtime[`${key}_interval_ms`];
    const row = element(doc, 'div', 'fc-row');
    const mode = element(doc, 'span', 'sim-mode-cell');
    mode.appendChild(disabledSelect(doc, `${label} simulation`, [['none', 'None'], ['random', 'Random']], interval ? 'random' : 'none'));
    row.append(
      element(doc, 'strong', '', label),
      readOnlyField(doc, 'Start', area.start, 'cell-field'),
      readOnlyField(doc, 'Count', area.count, 'cell-field'),
      mode,
      readOnlyField(doc, 'Interval', interval, 'cell-field'),
      element(doc, 'span', 'range-cell', range(area))
    );
    table.appendChild(row);
  });
  editor.append(table, editorActions(doc), notice(doc, 'FIXTURE ONLY — Memory definitions are examples. No load, save, apply or polling occurs.'));
  shell.appendChild(editor);
  root.appendChild(shell);
};

const commsStrip = (doc, fixture) => {
  const strip = element(doc, 'div', 'comms-strip');
  strip.setAttribute('aria-label', 'Fixture communications; no checks performed');
  [['SOURCE', COMMS.slice(0, 3)], ['DESTINATION', COMMS.slice(3)]].forEach(([heading, items]) => {
    const group = element(doc, 'div', 'comms-group');
    group.appendChild(element(doc, 'span', 'comms-heading', heading));
    const indicators = element(doc, 'div', 'comms-indicators');
    items.forEach(([key, label]) => {
      const item = element(doc, 'div', 'comms-item');
      item.appendChild(element(doc, 'span', 'comms-label', label));
      const led = element(doc, 'button', 'comms-led');
      led.type = 'button';
      led.disabled = true;
      led.dataset.state = fixture.comms[key];
      led.setAttribute('aria-label', `${label}: ${UNKNOWN.toLowerCase()} (fixture; not tested)`);
      led.title = `${label}: not tested — fixture only`;
      item.appendChild(led);
      indicators.appendChild(item);
    });
    group.appendChild(indicators);
    strip.appendChild(group);
  });
  return strip;
};
const renderReplicator = (doc, root, fixture) => {
  const device = fixture.replicator.devices[0];
  const shell = element(doc, 'div', 'tool-layout');
  shell.appendChild(sidebar(doc, 'Replicator Devices', device, `${device.endpoint} / FC${device.pull_blocks[0].function} / FIXTURE`));
  const editor = element(doc, 'section', 'tool-editor');
  editor.append(element(doc, 'h2', '', 'Device Definition'), commsStrip(doc, fixture));
  const identity = element(doc, 'div', 'tool-grid');
  identity.append(
    readOnlyField(doc, 'Name', device.name), readOnlyCheckbox(doc, 'Enabled', device.enabled),
    readOnlyField(doc, 'Endpoint', device.endpoint), readOnlyField(doc, 'Source Unit ID', device.unit_id)
  );
  editor.append(identity, element(doc, 'h3', '', 'Destination'));
  const dest = element(doc, 'div', 'tool-grid');
  dest.append(
    readOnlyField(doc, 'Port', device.destination.port), readOnlyCheckbox(doc, 'Auto Port', device.destination.auto_port),
    readOnlyField(doc, 'Unit ID', device.destination.unit_id), readOnlyCheckbox(doc, 'Auto Unit ID', device.destination.auto_unit_id)
  );
  editor.append(dest, runtimeRow(doc, [['Destination ownership:', UNKNOWN]]), element(doc, 'h3', '', 'Pull Blocks'));
  editor.appendChild(actions(doc, ['Add Block', 'Duplicate Block', 'Delete Block']));
  const table = element(doc, 'div', 'block-table');
  const header = element(doc, 'div', 'block-row block-header');
  ['#', 'FC', 'Start', 'Count', 'Scan Rate (ms)', ''].forEach(label => header.appendChild(element(doc, 'span', '', label)));
  table.appendChild(header);
  device.pull_blocks.forEach((block, index) => {
    const row = element(doc, 'div', 'block-row selected');
    const selectWrap = element(doc, 'span');
    selectWrap.appendChild(disabledSelect(doc, `Pull Block ${index + 1} Function`, [1, 2, 3, 4].map(fc => [fc, `FC${fc}`]), block.function));
    row.append(
      element(doc, 'strong', '', index + 1), selectWrap,
      readOnlyField(doc, 'Start', block.start, 'cell-field'),
      readOnlyField(doc, 'Count', block.count, 'cell-field'),
      readOnlyField(doc, 'Scan', block.scan_rate_ms, 'cell-field'),
      disabledButton(doc, '-', 'tool-row-delete tool-danger')
    );
    table.appendChild(row);
  });
  editor.append(table, editorActions(doc), notice(doc, 'FIXTURE ONLY — Replicator configuration and COMMS are not connected or tested.'));
  shell.appendChild(editor);
  root.appendChild(shell);
};
const renderDiagnostics = (doc, root, fixture) => {
  const buttons = element(doc, 'div', 'diagnostic-actions');
  ['Start runtimes', 'Stop runtimes'].forEach(label => buttons.appendChild(disabledButton(doc, label)));
  const paths = element(doc, 'div', 'paths');
  paths.append(
    textRow(doc, 'Simulator service:', fixture.runtime.simulator),
    textRow(doc, 'Runtime mode:', fixture.diagnostics.runtime_mode),
    textRow(doc, 'Binary folder:', fixture.diagnostics.bin_path),
    textRow(doc, 'Data folder:', fixture.diagnostics.data_path)
  );
  root.append(buttons, paths, notice(doc, 'FIXTURE ONLY — no Windows services, Docker controls or live diagnostic endpoints are connected.'));
  const log = element(doc, 'pre', '', 'No runtime logs: fixture preview only.');
  log.id = 'log';
  root.appendChild(log);
};

const createToolkit = doc => {
  const fixture = snapshot();
  const host = element(doc, 'section', 'mcs-toolkit-host');
  const shadow = host.attachShadow({mode: 'open'});
  const style = element(doc, 'style');
  // The donor's global selectors now apply only inside this window's ShadowRoot.
  style.textContent = `:host {display:block; width:100%; height:100%; min-height:0; overflow:hidden; color:#111; font:12px Tahoma, "Segoe UI", sans-serif;}\n${donorCSS}`;
  shadow.appendChild(style);
  const shell = element(doc, 'div', 'app-shell');
  // Static markup follows the approved donor's three-tab index.html structure.
  shell.innerHTML = '<header class="titlebar"><div><strong>MCS Modbus Toolkit</strong><span class="subtitle">Memory + Replicator — fixture preview</span></div><div class="runtime-strip"><span>MMA2 <b>UNKNOWN</b></span><span>Replicator <b>UNKNOWN</b></span></div></header><nav class="tabs" aria-label="Application tabs"><button class="tab active" type="button" data-tab="simulator" aria-selected="true">Memory</button><button class="tab" type="button" data-tab="replicator" aria-selected="false">Replicator</button><button class="tab" type="button" data-tab="diagnostics" aria-selected="false">Diagnostics</button></nav><main class="content"><section class="panel active" id="panel-simulator"><div id="simulator-root" class="tool-root"></div></section><section class="panel" id="panel-replicator"><div id="replicator-root" class="tool-root"></div></section><section class="panel" id="panel-diagnostics"></section></main>';
  shadow.appendChild(shell);
  renderMemory(doc, shadow.getElementById('simulator-root'), fixture);
  renderReplicator(doc, shadow.getElementById('replicator-root'), fixture);
  renderDiagnostics(doc, shadow.getElementById('panel-diagnostics'), fixture);
  const tabs = Array.from(shadow.querySelectorAll('.tab'));
  const panels = Array.from(shadow.querySelectorAll('.panel'));
  const listeners = tabs.map(tab => {
    const activate = () => {
      tabs.forEach(node => {
        const current = node === tab;
        node.classList.toggle('active', current);
        node.setAttribute('aria-selected', String(current));
      });
      panels.forEach(panel => panel.classList.toggle('active', panel.id === `panel-${tab.dataset.tab}`));
    };
    tab.addEventListener('click', activate);
    return () => tab.removeEventListener('click', activate);
  });
  return {
    element: host,
    destroy: () => { listeners.forEach(remove => remove()); host.remove(); }
  };
};

module.exports = {createToolkit};
