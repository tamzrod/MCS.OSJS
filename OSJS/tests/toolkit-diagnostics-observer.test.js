'use strict';

// JR runs this in UMIG-006-T. Source author has not executed it.
const assert = require('assert');
const {collectDiagnostics} = require('../src/packages/MCSModbusToolkit/diagnostics-observer');
const memoryStatus = name => ({status: {name, mma2_status: 'RUNNING', device_status: 'IDLE'}});
const repStatus = name => ({name, enabled: true, running: true, source_status: 'ERROR', last_error: 'source down',
  blocks: [{index: 0, running: true, source_status: 'ERROR', last_error: 'connection refused'}]});

(async () => {
  const calls = [];
  const memory = {load: async () => {calls.push('memory.load'); return {document: {devices: [{name: 'Sim-1'}, {name: 'Sim-2'}]}};},
    status: async name => {calls.push(`memory.status:${name}`); return memoryStatus(name);}};
  const replicator = {load: async () => {calls.push('replicator.load'); return {document: {devices: [{name: 'Rep-1'}]}};},
    status: async name => {calls.push(`replicator.status:${name}`); return repStatus(name);}};
  const live = await collectDiagnostics({memory, replicator});
  assert.strictEqual(live.memory.name, 'Sim-1');
  assert.strictEqual(live.memory.count, 2);
  assert.strictEqual(live.replicator.name, 'Rep-1');
  assert.strictEqual(live.view.devices.memory.simulator, 'IDLE');
  assert.strictEqual(live.view.devices.memory.mma2, 'RUNNING');
  assert.strictEqual(live.view.devices.replicator.runtime, 'RUNNING');
  assert.strictEqual(live.view.devices.replicator.source, 'ERROR');
  assert.deepStrictEqual(live.view.runtime, {mma2: 'UNKNOWN', simulator: 'UNKNOWN', replicator: 'UNKNOWN'});
  assert.deepStrictEqual(calls.sort(), ['memory.load', 'memory.status:Sim-1', 'replicator.load', 'replicator.status:Rep-1'].sort());
  console.log('read-only canonical selection, actual IDLE/error and unknown global services: checked');

  const empty = await collectDiagnostics({memory: {load: async () => ({document: {devices: []}}), status: () => {throw Error('unexpected Memory status');}},
    replicator: {load: async () => ({document: {devices: []}}), status: () => {throw Error('unexpected Replicator status');}}});
  assert.strictEqual(empty.memory.count, 0);
  assert.strictEqual(empty.replicator.count, 0);
  assert.strictEqual(empty.view.devices.memory.mma2, 'UNKNOWN');
  assert.strictEqual(empty.view.devices.replicator.source, 'UNKNOWN');
  console.log('empty canonical documents produce no fictional status call: checked');

  const outage = await collectDiagnostics({memory: {load: async () => ({document: {devices: [{name: 'Sim-1'}]}}),
    status: async () => {throw new Error('simulator socket missing');}}, replicator});
  assert.strictEqual(outage.memory.name, 'Sim-1');
  assert.strictEqual(outage.view.devices.memory.simulator, 'UNAVAILABLE');
  assert.match(outage.view.errors.memory, /simulator socket missing/);
  assert.strictEqual(outage.view.devices.replicator.source, 'ERROR');
  console.log('independent backend error cannot hide other device evidence: checked');

  const malformed = await collectDiagnostics({memory: {load: async () => ({document: {devices: [{}]}}),
    status: () => {throw Error('unexpected status');}}, replicator: {load: async () => {throw Error('replicator offline');},
    status: () => {throw Error('unexpected status');}}});
  assert.strictEqual(malformed.view.devices.memory.mma2, 'UNAVAILABLE');
  assert.strictEqual(malformed.view.devices.replicator.runtime, 'UNAVAILABLE');
  assert.match(malformed.view.errors.memory, /first canonical device/);
  assert.match(malformed.view.errors.replicator, /offline/);
  const mismatch = await collectDiagnostics({memory: {load: memory.load, status: async () => memoryStatus('other')}, replicator});
  assert.strictEqual(mismatch.view.devices.memory.simulator, 'UNKNOWN');
  await assert.rejects(collectDiagnostics({memory: {}, replicator}), /requires/);
  console.log('malformed, wrong-device, offline and missing contract fail closed: checked');
  console.log('UMIG-006 Diagnostics observer checks complete');
})().catch(error => {console.error(error); process.exitCode = 1;});
