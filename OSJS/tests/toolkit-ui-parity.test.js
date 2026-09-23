'use strict';

const assert = require('assert');
const fs = require('fs');
const path = require('path');
const {JSDOM} = require('jsdom');
const {createMemoryEditor, blankDevice} = require('../src/packages/MCSModbusToolkit/memory-editor');
const {createReplicatorEditor} = require('../src/packages/MCSModbusToolkit/replicator-editor');
const {defaults} = require('../src/packages/MCSModbusToolkit/memory-advanced');
const copy = value => JSON.parse(JSON.stringify(value));
const settle = () => new Promise(resolve => setImmediate(resolve));
const original = {setTimeout, clearTimeout, setInterval, clearInterval, now: Date.now};
const pending = new Map();
const intervals = new Map();
let sequence = 0;
let now = 10000;
global.setTimeout = callback => { pending.set(++sequence, callback); return sequence; };
global.clearTimeout = handle => pending.delete(handle);
global.setInterval = callback => { intervals.set(++sequence, callback); return sequence; };
global.clearInterval = handle => intervals.delete(handle);
Date.now = () => now;
const poll = async () => {
  const callbacks = [...pending.values()]; pending.clear();
  callbacks.forEach(callback => callback());
  await settle();
};
const click = (root, title) => {
  const button = [...root.querySelectorAll('button')].find(node => node.textContent === title);
  assert.ok(button, `Missing ${title}`);
  assert.ok(!button.matches(':disabled'), `${title} is disabled`);
  button.click();
};
const input = (root, label, value, event = 'input') => {
  const control = root.querySelector(`[aria-label="${label}"]`);
  assert.ok(control, `Missing ${label}`);
  if (control.type === 'checkbox') control.checked = value;
  else control.value = value;
  control.dispatchEvent(new control.ownerDocument.defaultView.Event(event, {bubbles: true}));
};

(async () => {
  const dom = new JSDOM('<div id="memory"></div><div id="replicator"></div>');
  const doc = dom.window.document;
  const style = doc.createElement('style');
  style.textContent = fs.readFileSync(path.join(__dirname, '../src/packages/MCSModbusToolkit/renderer.css'), 'utf8');
  doc.head.appendChild(style);
  const hiddenEmptyWarnings = root => {
    const warnings = [...root.querySelectorAll('.tool-validation')];
    assert.ok(warnings.length > 0);
    warnings.forEach(node => {
      assert.strictEqual(node.textContent, '');
      assert.strictEqual(dom.window.getComputedStyle(node).display, 'none');
    });
  };
  const root = doc.getElementById('memory');
  const device = blankDevice(1);
  Object.assign(device.mma2, defaults(), {custom_extension: {keep: true}});
  let saved = null;
  let rejectSave = false;
  let header = '';
  const memory = createMemoryEditor(doc, root, {
    load: async () => ({document: {devices: [device]}}),
    status: async name => ({status: {name, mma2_status: 'RUNNING', device_status: 'IDLE'}}),
    apply: async document => {
      if (rejectSave) throw new Error('candidate rejected');
      saved = copy(document);
      return {document: copy(document), message: 'Applied'};
    }
  }, {onStatus: value => { header = value; }});
  await settle();
  assert.strictEqual(header, 'RUNNING');
  hiddenEmptyWarnings(root);
  input(root, 'Name', '');
  const validation = root.querySelector('[data-memory-validation]');
  assert.strictEqual(validation.textContent, 'Name is required.');
  assert.notStrictEqual(dom.window.getComputedStyle(validation).display, 'none');
  input(root, 'Name', device.name); await settle();
  hiddenEmptyWarnings(root);
  click(root, 'Advanced Settings');
  assert.ok(root.textContent.includes('RBE TCP Port: Unavailable'));
  click(root, 'State Sealing');
  input(root, 'Enable state sealing', true, 'change');
  input(root, 'Control address', '2');
  click(root, 'Access Policy');
  input(root, 'Source IP / CIDR', '10.0.0.1, 10.0.0.0/24');
  input(root, 'Source presets', 'All IPv4', 'change');
  assert.strictEqual(root.querySelector('[aria-label="Source presets"]').value, '');
  input(root, 'Source IP / CIDR', '10.0.0.1, 10.0.0.0/24');
  assert.strictEqual(root.querySelector('.function-checks'), null);
  input(root, 'Access', 'Custom', 'change');
  assert.ok(root.querySelector('.function-checks'));
  input(root, 'Access', 'Read Only', 'change');
  assert.strictEqual(root.querySelector('.function-checks'), null);
  click(root, 'Device Definition'); click(root, 'Advanced Settings');
  click(root, 'Save & Apply'); await settle();
  assert.deepStrictEqual(saved.devices[0].mma2.policy.rules[0].source_ip, ['10.0.0.1', '10.0.0.0/24', '::/0']);
  assert.deepStrictEqual(saved.devices[0].mma2.policy.rules[0].allow_fc, [1, 2, 3, 4]);
  assert.strictEqual(saved.devices[0].mma2.state_sealing.address, 2);
  assert.deepStrictEqual(saved.devices[0].mma2.custom_extension, {keep: true});
  click(root, 'State Sealing'); input(root, 'Control address', '3');
  rejectSave = true; click(root, 'Save & Apply'); await settle();
  click(root, 'State Sealing');
  assert.strictEqual(root.querySelector('[aria-label="Control address"]').value, '3');
  assert.ok(root.textContent.includes('candidate rejected'));
  click(root, 'Discard'); click(root, 'State Sealing');
  assert.strictEqual(root.querySelector('[aria-label="Control address"]').value, '2');
  memory.destroy();
  console.log('Memory: folder tabs, source presets/comma lists, custom FC visibility, apply/failure/discard and extensions PASS');

  const repRoot = doc.getElementById('replicator');
  const repDevice = {name: 'Rep-1', enabled: true, endpoint: '127.0.0.1:5020', unit_id: 1,
    destination: {port: 5021, unit_id: 1, auto_port: true, auto_unit_id: true},
    pull_blocks: [{function: 1, start: 10, count: 4, scan_rate_ms: 1000}], mma2_advanced: defaults()};
  let reported = {name: 'Rep-1', enabled: true, running: true, source_status: 'OK', blocks: [],
    comms: {network: 'OK', tcp: 'OK', modbus: 'WARNING', mma2: 'ERROR'}};
  let failed = false;
  let repSaved = null;
  let repHeader = '';
  const transport = {
    load: async () => ({document: {devices: [repDevice]}, capabilities: {mma2_advanced: true}}),
    status: async () => { if (failed) throw new Error('socket unavailable'); return copy(reported); },
    apply: async document => { repSaved = copy(document); return {document: copy(document), message: 'Applied', completed_at: '2026-09-22T00:00:00Z'}; }
  };
  const replicator = createReplicatorEditor(doc, repRoot, transport, {onStatus: value => { repHeader = value; }});
  await settle();
  const state = layer => repRoot.querySelector(`#rep-comms-${layer}`).dataset.state;
  assert.strictEqual(repHeader, 'RUNNING');
  hiddenEmptyWarnings(repRoot);
  assert.strictEqual(state('tcp'), 'OK');
  assert.strictEqual(state('modbus'), 'WARNING');
  assert.strictEqual(state('mma2'), 'ERROR');
  repRoot.querySelector('#rep-comms-tcp').click();
  assert.strictEqual(repRoot.querySelector('#rep-comms-tcp-detail').hidden, false);
  click(repRoot, 'Advanced Settings'); click(repRoot, 'State Sealing');
  input(repRoot, 'Enable state sealing', true, 'change');
  assert.strictEqual(repRoot.querySelector('[aria-label="Control address"]').value, '10');
  input(repRoot, 'Control address', '11');
  click(repRoot, 'Save & Apply'); await settle();
  assert.strictEqual(repSaved.devices[0].mma2_advanced.state_sealing.address, 11);
  assert.strictEqual(repSaved.devices[0].mma2, undefined);
  now += 7000; intervals.forEach(callback => callback());
  assert.strictEqual(state('tcp'), 'UNKNOWN');
  assert.strictEqual(repHeader, 'UNKNOWN');
  await poll(); assert.strictEqual(state('tcp'), 'OK');
  failed = true; await poll();
  assert.strictEqual(state('tcp'), 'UNKNOWN');
  assert.strictEqual(repHeader, 'UNAVAILABLE');
  const runtimeWarning = repRoot.querySelector('[data-replicator-status="error"]');
  assert.ok(runtimeWarning.textContent.includes('socket unavailable'));
  assert.notStrictEqual(dom.window.getComputedStyle(runtimeWarning).display, 'none');
  failed = false; await poll(); assert.strictEqual(state('tcp'), 'OK');
  hiddenEmptyWarnings(repRoot);
  reported.name = 'Other'; await poll(); assert.strictEqual(state('tcp'), 'UNKNOWN');
  reported.name = 'Rep-1'; delete reported.comms; await poll();
  assert.strictEqual(state('tcp'), 'UNKNOWN');
  assert.ok(repRoot.textContent.includes('does not provide communication telemetry'));
  reported.comms = {tcp: 'OK'}; await poll();
  click(repRoot, 'Device Definition'); input(repRoot, 'Name', 'Unsaved');
  assert.strictEqual(state('tcp'), 'UNKNOWN');
  replicator.destroy();
  const old = createReplicatorEditor(doc, repRoot, {...transport,
    load: async () => ({document: {devices: [repDevice]}})});
  await settle(); click(repRoot, 'Advanced Settings');
  assert.ok(repRoot.textContent.includes('Update and restart'));
  assert.strictEqual(repRoot.querySelector('.memory-advanced'), null);
  old.destroy(); dom.window.close();
  assert.strictEqual(pending.size, 0); assert.strictEqual(intervals.size, 0);
  console.log('Replicator: advanced destination persistence, actual LEDs, tooltip, stale/failure/recovery, identity, old backend and teardown PASS');
})().catch(error => { console.error(error); process.exitCode = 1; }).finally(() => {
  global.setTimeout = original.setTimeout; global.clearTimeout = original.clearTimeout;
  global.setInterval = original.setInterval; global.clearInterval = original.clearInterval;
  Date.now = original.now;
});
