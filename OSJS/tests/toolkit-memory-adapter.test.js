'use strict';

// JR executes this in UMIG-004-T; CODE authors but does not claim TEST PASS.
const assert = require('assert');
const {EventEmitter} = require('events');
const {createMemoryContract} = require('../src/packages/MCSModbusToolkit/memory-contract');
const {createMemoryTransport} = require('../src/packages/MCSModbusToolkit/memory-transport');
const {blankDevice, normalizeDocument, validateDocument, statusWord, changeMode} = require('../src/packages/MCSModbusToolkit/memory-editor');

const respond = (request, result) => ({version: 1, request_id: request.request_id, ok: true, result});
const record = {devices: [blankDevice(1)]};
record.devices[0].random_runtime.fc3_interval_ms = 1250;

(async () => {
  const proc = new EventEmitter();
  const sent = [];
  proc.send = request => { sent.push(request); };
  const transport = createMemoryTransport(proc);
  const contract = createMemoryContract(transport.send);
  const loading = contract.load();
  assert.strictEqual(sent.length, 1);
  assert.strictEqual(sent[0].operation, 'load');
  assert.deepStrictEqual(sent[0].payload, {});
  proc.emit('ws:message', respond({...sent[0], request_id: 'unrelated'}, {document: {devices: []}}));
  proc.emit('ws:message', respond(sent[0], {document: record}));
  assert.deepStrictEqual((await loading).document, record);
  console.log('Toolkit provider envelope, correlated load, stale reply ignored: checked');

  const editing = JSON.parse(JSON.stringify(record));
  const applying = contract.apply(editing);
  editing.devices[0].random_runtime.fc3_interval_ms = 999;
  assert.strictEqual(sent[1].operation, 'apply');
  assert.strictEqual(sent[1].payload.document.devices[0].random_runtime.fc3_interval_ms, 1250);
  proc.emit('ws:message', respond(sent[1], {document: record, message: 'Random-runtime timing updated without restarting MMA2.', completed_at: '2026-09-19T00:00:00Z'}));
  assert.strictEqual((await applying).document.devices[0].random_runtime.fc3_interval_ms, 1250);
  const observing = contract.status('Sim-PLC-1');
  assert.deepStrictEqual(sent[2].payload, {name: 'Sim-PLC-1'});
  proc.emit('ws:message', respond(sent[2], {status: {name: 'Sim-PLC-1', device_status: 'IDLE', mma2_status: 'RUNNING'}}));
  assert.strictEqual((await observing).status.device_status, 'IDLE');
  console.log('explicit apply snapshot, canonical result, status payload and IDLE: checked');

  const failing = contract.apply(record);
  proc.emit('ws:message', {version: 1, request_id: sent[3].request_id, ok: false,
    error: {code: 'MMA2_NOT_READY', message: 'restart pending'}});
  await assert.rejects(failing, error => error.code === 'MMA2_NOT_READY' && error.message === 'restart pending');
  const unavailable = contract.status('Sim-PLC-1');
  proc.emit('ws:message', {version: 1, request_id: sent[4].request_id, ok: false,
    error: {code: 'RUNTIME_UNAVAILABLE', message: 'socket closed'}});
  await assert.rejects(unavailable, error => error.code === 'RUNTIME_UNAVAILABLE' && error.message === 'socket closed');
  const interrupted = contract.load();
  transport.close();
  await assert.rejects(interrupted, /window closed/);
  await assert.rejects(contract.load(), /transport closed/);
  console.log('typed runtime failures, unavailable transport and teardown: checked');

  const none = blankDevice(2);
  assert.deepStrictEqual(Object.values(none.random_runtime), [0, 0, 0, 0]);
  assert.strictEqual(validateDocument({devices: [none]}), null);
  none.random_runtime.fc3_interval_ms = 1250;
  const remembered = new Map();
  changeMode(none, 'fc3', 'none', remembered);
  assert.strictEqual(none.random_runtime.fc3_interval_ms, 0);
  assert.strictEqual(validateDocument({devices: [none]}), null);
  changeMode(none, 'fc3', 'random', remembered);
  assert.strictEqual(none.random_runtime.fc3_interval_ms, 1250);
  const fresh = blankDevice(4);
  changeMode(fresh, 'fc4', 'random', new Map());
  assert.strictEqual(fresh.random_runtime.fc4_interval_ms, 1000);
  assert.strictEqual(validateDocument({devices: [none]}), null);
  none.mma2.fc4.start = 65535;
  none.mma2.fc4.count = 2;
  assert.match(validateDocument({devices: [none]}), /16-bit/);
  assert.match(validateDocument({devices: [blankDevice(1), blankDevice(1)]}), /Duplicate/);
  assert.strictEqual(validateDocument({devices: []}), null); // Delete-last can be committed explicitly.
  assert.deepStrictEqual(normalizeDocument(null), {devices: []});
  console.log('None/Random round-trip, fresh defaults, empty document and validation: checked');

  assert.strictEqual(statusWord(null, 'Sim-PLC-1', 'device_status', null), 'UNKNOWN');
  assert.strictEqual(statusWord({name: 'Other', device_status: 'RUNNING'}, 'Sim-PLC-1', 'device_status', null), 'UNKNOWN');
  assert.strictEqual(statusWord({name: 'Sim-PLC-1', device_status: 'RUNNING'}, 'Sim-PLC-1', 'device_status', 'socket closed'), 'UNAVAILABLE');
  assert.strictEqual(statusWord({name: 'Sim-PLC-1', device_status: 'IDLE'}, 'Sim-PLC-1', 'device_status', null), 'IDLE');
  console.log('unknown, wrong-device, unavailable and observed status mapping: checked');
  console.log('UMIG-004 Memory adapter contract cases complete');
})().catch(error => { console.error(error); process.exitCode = 1; });
