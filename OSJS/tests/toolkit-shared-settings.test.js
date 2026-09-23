'use strict';

const assert = require('assert');
const fs = require('fs');
const path = require('path');
const vm = require('vm');
const {EventEmitter} = require('events');
const {JSDOM} = require('jsdom');
const {createMemoryContract} = require('../src/packages/MCSModbusToolkit/memory-contract');
const {createSharedSettings} = require('../src/packages/MCSModbusToolkit/shared-settings');
const {createMemoryEditor, blankDevice} = require('../src/packages/MCSModbusToolkit/memory-editor');
const {createReplicatorEditor} = require('../src/packages/MCSModbusToolkit/replicator-editor');
const settle = () => new Promise(resolve => setImmediate(resolve));
const revision = 'a'.repeat(64);
const nextRevision = 'b'.repeat(64);
const copy = value => JSON.parse(JSON.stringify(value));
const click = (root, text) => {
  const control = [...root.querySelectorAll('button')].find(node => node.textContent === text);
  assert.ok(control, `missing ${text}`); assert.ok(!control.matches(':disabled'), `disabled ${text}`);
  control.click();
};
const change = (root, label, value) => {
  const control = root.querySelector(`[aria-label="${label}"]`);
  assert.ok(control, `missing ${label}`);
  if (control.type === 'checkbox') control.checked = value; else control.value = value;
  control.dispatchEvent(new control.ownerDocument.defaultView.Event(control.type === 'checkbox' ? 'change' : 'input', {bubbles: true}));
};

(async () => {
  let request = null;
  let failure = null;
  const contract = createMemoryContract(async value => {
    request = value;
    return {version: 1, request_id: value.request_id, ok: !failure,
      result: {settings: {debug: false}, revision, committed: true, restart_acknowledged: !failure}, error: failure};
  });
  assert.strictEqual((await contract.loadShared()).revision, revision);
  assert.strictEqual(request.operation, 'mma-load');
  const settings = {rbe: {tcp: {listen: ':9001'}}};
  const applying = contract.applyShared(revision, settings); settings.rbe.tcp.listen = ':9002';
  await applying;
  assert.strictEqual(request.payload.settings.rbe.tcp.listen, ':9001');
  await assert.rejects(contract.applyShared('', {}), error => error.code === 'INVALID_REQUEST');
  failure = {code: 'MMA2_RESTART_FAILED', message: 'committed but restart not acknowledged'};
  await assert.rejects(contract.applyShared(revision, {}), error => error.result.committed && error.code === failure.code);
  console.log('shared contract: versioned load/apply, snapshot, validation and committed-error details PASS');

  const received = [];
  const relayModule = {exports: {}};
  const fakeNet = {createConnection: socketPath => {
    const socket = new EventEmitter(); socket.destroy = () => {};
    socket.write = buffer => {
      const request = JSON.parse(buffer.subarray(4).toString());
      received.push({socketPath, request});
      const body = Buffer.from(JSON.stringify({version: 1, request_id: request.request_id, ok: true,
        result: {settings: {}, revision}}));
      const header = Buffer.alloc(4); header.writeUInt32BE(body.length);
      process.nextTick(() => socket.emit('data', Buffer.concat([header, body])));
    };
    process.nextTick(() => socket.emit('connect')); return socket;
  }};
  vm.runInNewContext(fs.readFileSync(path.join(__dirname, '../src/packages/MCSModbusToolkit/server.js'), 'utf8'), {
    module: relayModule, require: name => name === 'net' ? fakeNet : require(name), process, Buffer, setTimeout, clearTimeout
  });
  const provider = relayModule.exports();
  const relay = (authenticated, request) => new Promise(resolve => provider.onmessage({_osjs_client: authenticated}, resolve, [request]));
  const wire = {version: 1, request_id: 'mcs-memory-shared-test', operation: 'mma-load', payload: {}};
  assert.strictEqual((await relay(false, wire)).ok, false);
  assert.strictEqual((await relay(true, {...wire, operation: 'delete-file'})).ok, false);
  assert.strictEqual((await relay(true, {...wire, request_id: 'mcs-replicator-test'})).ok, false);
  assert.strictEqual(received.length, 0);
  assert.strictEqual((await relay(true, wire)).ok, true);
  assert.strictEqual((await relay(true, {...wire, operation: 'mma-apply', payload: {revision, settings: {debug: true}}})).ok, true);
  assert.ok(received.every(item => item.socketPath.endsWith('modbus-simulator.sock')));
  assert.strictEqual(received[1].request.payload.revision, revision);
  console.log('shared relay: authenticated allowlist, Simulator-only routing and framed messages PASS (mocked socket)');

  const dom = new JSDOM('<div id="host"></div>');
  const doc = dom.window.document;
  const shadow = doc.getElementById('host').attachShadow({mode: 'open'});
  const memoryRoot = doc.createElement('div'); const repRoot = doc.createElement('div');
  shadow.append(memoryRoot, repRoot);
  let saved = {settings: {debug: false}, revision};
  let reject = null;
  let pendingApply = null;
  let sharedDraft = null;
  const api = {
    loadShared: async () => copy(saved),
    applyShared: async (expected, value) => {
      assert.strictEqual(expected, saved.revision); sharedDraft = copy(value);
      if (pendingApply) await pendingApply;
      if (reject) throw reject;
      saved = {settings: copy(value), revision: nextRevision};
      return {...copy(saved), committed: true, restart_acknowledged: true, message: 'Restart acknowledged; readiness unverified.'};
    },
    load: async () => ({document: {devices: [blankDevice(1)]}}),
    status: async name => ({status: {name, mma2_status: 'RUNNING', device_status: 'IDLE'}})
  };
  const shared = createSharedSettings(doc, shadow, api);
  const memory = createMemoryEditor(doc, memoryRoot, api, {shared});
  const repDevice = {name: 'Rep', enabled: true, endpoint: '127.0.0.1:5020', unit_id: 1,
    destination: {port: 5021, unit_id: 1, auto_port: true, auto_unit_id: true},
    pull_blocks: [{function: 3, start: 0, count: 1, scan_rate_ms: 1000}]};
  const rep = createReplicatorEditor(doc, repRoot, {
    load: async () => ({document: {devices: [repDevice]}, capabilities: {mma2_advanced: true}}),
    status: async name => ({name, enabled: true, running: true, source_status: 'OK', blocks: []})
  }, {shared});
  try {
    await shared.load(); await settle();
    change(memoryRoot, 'Name', 'Unsaved device');
    click(memoryRoot, 'Advanced Settings'); click(repRoot, 'Advanced Settings');
    assert.ok(memoryRoot.textContent.includes('RBE TCP Port: Not configured'));
    click(memoryRoot, 'RBE TCP Settings...');
    let dialog = shadow.querySelector('dialog');
    change(dialog, 'Enable RBE TCP output', true);
    assert.strictEqual(dialog.querySelector('[aria-label="RBE TCP listen address (IP:port)"]').value, ':9001');
    change(dialog, 'RBE TCP listen address (IP:port)', '127.0.0.1:39111');
    click(dialog, 'Close'); click(repRoot, 'MMA Settings...'); dialog = shadow.querySelector('dialog');
    assert.strictEqual(dialog.querySelector('[aria-label="RBE TCP listen address (IP:port)"]').value, '127.0.0.1:39111');
    reject = Object.assign(new Error('invalid settings'), {code: 'VALIDATION_FAILED'});
    click(dialog, 'Save & Apply'); await settle();
    assert.ok(dialog.textContent.includes('invalid settings'));
    assert.strictEqual(sharedDraft.rbe.tcp.listen, '127.0.0.1:39111');
    reject = null; click(dialog, 'Save & Apply'); await settle();
    assert.ok(memoryRoot.textContent.includes('RBE TCP Port: 39111'));
    assert.ok(repRoot.textContent.includes('RBE TCP Port: 39111'));
    click(dialog, 'Close'); click(memoryRoot, 'Device Definition');
    assert.strictEqual(memoryRoot.querySelector('[aria-label="Name"]').value, 'Unsaved device');
    click(repRoot, 'MMA Settings...'); dialog = shadow.querySelector('dialog');
    change(dialog, 'RBE TCP listen address (IP:port)', ':39112');
    click(dialog, 'Discard');
    assert.strictEqual(dialog.querySelector('[aria-label="RBE TCP listen address (IP:port)"]').value, '127.0.0.1:39111');
    reject = Object.assign(new Error('changed externally'), {code: 'REVISION_CONFLICT'});
    click(dialog, 'Save & Apply'); await settle();
    assert.ok([...dialog.querySelectorAll('button')].find(node => node.textContent === 'Save & Apply').disabled);
    assert.ok(dialog.textContent.includes('draft has been retained'));
    reject = null; click(dialog, 'Reload saved settings'); await settle();
    let resolveApply;
    pendingApply = new Promise(resolve => { resolveApply = resolve; });
    click(dialog, 'Save & Apply');
    assert.ok([...dialog.querySelectorAll('button')].find(node => node.textContent === 'Close').disabled);
    shared.destroy(); resolveApply(); await settle();
    assert.strictEqual(shadow.querySelector('dialog'), null);
    console.log('shared dialog: editable default port, both editors, separate drafts, discard, validation/conflict, pending-save and teardown PASS');
  } finally { memory.destroy(); rep.destroy(); shared.destroy(); dom.window.close(); }
})().catch(error => { console.error(error); process.exitCode = 1; });
