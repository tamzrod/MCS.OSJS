'use strict';

// The donor Diagnostics surface, backed ONLY by the two existing read-only
// Toolkit contracts. No Docker health probe, logs endpoint, Windows IPC,
// native path discovery, service-control call or fixture fallback exists.
const {UNKNOWN, mapDiagnostics} = require('./diagnostics-model');
const {collectDiagnostics} = require('./diagnostics-observer');

const createDiagnosticsEditor = (doc, root, memory, replicator, options = {}) => {
  const setTimer = options.setTimer || setTimeout;
  const clearTimer = options.clearTimer || clearTimeout;
  const interval = options.interval || 5000;
  let closed = false;
  let generation = 0;
  let timer = null;

  const h = (tag, className = '', text) => {
    const node = doc.createElement(tag);
    if (className) node.className = className;
    if (text !== undefined) node.textContent = String(text);
    return node;
  };
  const row = (parent, label, value) => {
    const line = h('div');
    line.append(h('strong', '', label), doc.createTextNode(' '), h('span', '', value));
    parent.appendChild(line);
  };
  const selectorLabel = observation => {
    if (!observation) return 'Not yet observed';
    if (observation.name) return `${observation.name} (first of ${observation.count} canonical device(s))`;
    if (observation.count === 0) return 'No canonical devices configured';
    return 'Canonical selection unavailable';
  };
  const render = (snapshot, note) => {
    if (closed) return;
    const mapped = snapshot ? snapshot.view : mapDiagnostics();
    const shell = h('div');
    // The original panel used a flex-column layout with a growing log pane.
    // Keep that donor behavior inside this one ShadowRoot, not global CSS.
    shell.style.cssText = 'display:flex; flex-direction:column; height:100%; min-height:0;';
    const controls = h('div', 'diagnostic-actions');
    ['Start runtimes', 'Stop runtimes'].forEach(label => {
      const button = h('button', 'tool-button', label);
      button.type = 'button';
      button.disabled = true;
      button.title = 'Unavailable: this Toolkit has no service-control capability.';
      button.setAttribute('aria-label', `${label}: unavailable; no service-control capability`);
      controls.appendChild(button);
    });
    const refresh = h('button', 'tool-button', 'Refresh observations');
    refresh.type = 'button';
    refresh.title = 'Read canonical device definitions and device status only; no write or service control.';
    refresh.addEventListener('click', requestRefresh);
    controls.appendChild(refresh);
    shell.appendChild(controls);

    const paths = h('div', 'paths');
    row(paths, 'MMA2 service:', mapped.runtime.mma2);
    row(paths, 'Simulator service:', mapped.runtime.simulator);
    row(paths, 'Replicator service:', mapped.runtime.replicator);
    row(paths, 'Runtime mode:', mapped.diagnostics.runtime_mode);
    row(paths, 'Binary folder:', mapped.diagnostics.bin_path);
    row(paths, 'Data folder:', mapped.diagnostics.data_path);
    shell.appendChild(paths);

    const devices = h('div', 'paths');
    row(devices, 'Memory device:', selectorLabel(snapshot && snapshot.memory));
    row(devices, 'Memory device MMA2:', mapped.devices.memory.mma2);
    row(devices, 'Memory simulation:', mapped.devices.memory.simulator);
    row(devices, 'Replicator device:', selectorLabel(snapshot && snapshot.replicator));
    row(devices, 'Replicator device runtime:', mapped.devices.replicator.runtime);
    row(devices, 'Replicator source:', mapped.devices.replicator.source);
    mapped.devices.replicator.blocks.forEach(block => {
      row(devices, `Pull Block ${block.index + 1}:`, `${block.source} / poller ${block.running}`);
    });
    shell.appendChild(devices);
    const notice = h('div', 'tool-status', note ||
      'READ ONLY — first canonical device of each type; no global service-health or COMMS probes.');
    shell.appendChild(notice);

    const messages = ['Read-only device observations; not runtime service logs.'];
    if (mapped.errors.memory) messages.push(`Memory: ${mapped.errors.memory}`);
    if (mapped.errors.replicator) messages.push(`Replicator: ${mapped.errors.replicator}`);
    const rep = snapshot && snapshot.replicator;
    // Only a freshly correlated device status may contribute backend detail.
    if (rep && rep.observation && rep.observation.fresh === true &&
        rep.observation.expectedName === rep.name) {
      const status = rep.observation.result;
      if (status && status.name === rep.name && Array.isArray(status.blocks)) {
        if (typeof status.last_error === 'string' && status.last_error) {
          messages.push(`Replicator source: ${status.last_error}`);
        }
        status.blocks.forEach((block, index) => {
          if (block && block.index === index) {
            const last = typeof block.last_poll === 'string' && block.last_poll ? block.last_poll : UNKNOWN;
            const error = typeof block.last_error === 'string' && block.last_error ? block.last_error : 'no reported error';
            messages.push(`Block ${index + 1}: last poll ${last}; ${error}`);
          }
        });
      }
    }
    if (!snapshot) messages.push('No status received yet.');
    messages.push('Windows service controls and binary/data paths: UNAVAILABLE. Docker/MMA2 supervisor status: UNKNOWN.');
    const log = h('pre', '', messages.join('\n'));
    log.id = 'log';
    shell.appendChild(log);
    root.replaceChildren(shell);
  };

  function requestRefresh() {
    if (closed) return;
    const token = ++generation;
    if (timer !== null) clearTimer(timer);
    timer = null;
    // Invalidate displayed observations immediately; never display cached
    // green status while newer requests are outstanding or after teardown.
    render(null, 'Refreshing read-only canonical device observations...');
    Promise.resolve().then(() => collectDiagnostics({memory, replicator})).then(snapshot => {
      if (!closed && token === generation) render(snapshot);
    }).catch(error => {
      if (closed || token !== generation) return;
      const message = error && error.message ? error.message : String(error);
      render({view: mapDiagnostics({memory: {error: {message}}, replicator: {error: {message}}}),
        memory: null, replicator: null}, 'Diagnostics observation unavailable.');
    }).then(() => {
      if (!closed && token === generation) timer = setTimer(requestRefresh, interval);
    });
  }

  render(null);
  requestRefresh();
  return {destroy: () => {
    if (closed) return;
    closed = true;
    generation++;
    if (timer !== null) clearTimer(timer);
    timer = null;
    root.replaceChildren();
  }};
};

module.exports = {createDiagnosticsEditor};
