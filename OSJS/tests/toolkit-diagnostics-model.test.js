'use strict';

// Future OpenHands/JR focused test for dormant UMIG-CF-003. CODE has NOT run it.
const assert = require('assert');
const {UNKNOWN, UNAVAILABLE, mapDiagnostics} = require('../src/packages/MCSModbusToolkit/diagnostics-model');

const memoryStatus = {status: {name: 'Sim-1', mma2_status: 'RUNNING', device_status: 'WAITING'}};
const repStatus = {name: 'Rep-1', enabled: true, running: true, source_status: 'OK',
  blocks: [{index: 0, running: true, source_status: 'OK'}, {index: 1, running: false, source_status: 'WAITING'}]};
const memory = {fresh: true, expectedName: 'Sim-1', result: memoryStatus};
const replicator = {fresh: true, expectedName: 'Rep-1', result: repStatus};

const empty = mapDiagnostics();
assert.deepStrictEqual(empty.runtime, {mma2: UNKNOWN, simulator: UNKNOWN, replicator: UNKNOWN});
assert.deepStrictEqual(empty.devices.memory, {mma2: UNKNOWN, simulator: UNKNOWN});
assert.deepStrictEqual(empty.devices.replicator, {runtime: UNKNOWN, source: UNKNOWN, blocks: []});
assert(Object.values(empty.diagnostics).every(value => value === UNAVAILABLE));
assert.deepStrictEqual(empty.controls, {start: false, stop: false});
console.log('empty diagnostics, global unknown, paths unavailable, controls disabled: checked');

const live = mapDiagnostics({memory, replicator});
assert.deepStrictEqual(live.devices.memory, {mma2: 'RUNNING', simulator: 'WAITING'});
assert.strictEqual(live.devices.replicator.runtime, 'RUNNING');
assert.strictEqual(live.devices.replicator.source, 'OK');
assert.deepStrictEqual(live.devices.replicator.blocks, [
  {index: 0, source: 'OK', running: 'RUNNING'},
  {index: 1, source: 'WAITING', running: 'STOPPED'}
]);
assert.deepStrictEqual(live.runtime, empty.runtime);
assert.deepStrictEqual(live.controls, empty.controls);
assert.deepStrictEqual(live.diagnostics, empty.diagnostics);
console.log('matching explicit per-device observations never imply global service health: checked');

assert.deepStrictEqual(mapDiagnostics({memory: {...memory, fresh: false}, replicator: {...replicator, fresh: false}}).devices, empty.devices);
assert.deepStrictEqual(mapDiagnostics({memory: {...memory, expectedName: 'Other'}, replicator: {...replicator, expectedName: 'Other'}}).devices, empty.devices);
assert.deepStrictEqual(mapDiagnostics({memory: {fresh: true, expectedName: 'Sim-1', result: {status: {name: 'Sim-1', mma2_status: 'OK', device_status: 'RUNNING'}}}}).devices.memory,
  {mma2: UNKNOWN, simulator: 'RUNNING'});
assert.strictEqual(mapDiagnostics({replicator: {...replicator, result: {...repStatus, blocks: [{index: 99, running: true, source_status: 'OK'}]}}}).devices.replicator.blocks[0].source, UNKNOWN);
assert.strictEqual(mapDiagnostics({replicator: {...replicator, result: {...repStatus, enabled: false}}}).devices.replicator.runtime, UNKNOWN);
console.log('stale, wrong-device, unknown status and contradictory/malformed block fail closed: checked');

const failed = mapDiagnostics({memory: {error: {message: 'Simulator socket closed'}}, replicator: {error: {message: 'Replicator socket closed'}}});
assert.deepStrictEqual(failed.devices.memory, {mma2: UNAVAILABLE, simulator: UNAVAILABLE});
assert.strictEqual(failed.devices.replicator.runtime, UNAVAILABLE);
assert.strictEqual(failed.devices.replicator.source, UNAVAILABLE);
assert.deepStrictEqual(failed.errors, {memory: 'Simulator socket closed', replicator: 'Replicator socket closed'});
assert.deepStrictEqual(failed.runtime, empty.runtime);
assert.deepStrictEqual(mapDiagnostics({memory, replicator}).devices, live.devices);
console.log('explicit failures, independent result snapshots and no global fabricated status: checked');
console.log('UMIG-CF-003 Diagnostics model cases complete');
