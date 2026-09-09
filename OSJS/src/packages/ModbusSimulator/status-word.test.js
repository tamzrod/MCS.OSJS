'use strict';

const assert = require('assert');
const { statusWord } = require('./status-word');

const NEUTRAL = '\u2014';
const UNAVAILABLE = 'UNAVAILABLE';
const RUNNING = 'RUNNING';

(async () => {
  assert.strictEqual(statusWord({ device: null, appliedName: null, runtimeStatus: null, runtimeStatusError: null, field: 'device_status' }), NEUTRAL, 'no device must be neutral');
  assert.strictEqual(statusWord({ device: { name: 'New Device' }, appliedName: null, runtimeStatus: null, runtimeStatusError: null, field: 'device_status' }), NEUTRAL, 'unsaved row must not fabricate unavailability');
  assert.strictEqual(statusWord({ device: { name: 'Renamed' }, appliedName: 'Original', runtimeStatus: null, runtimeStatusError: null, field: 'device_status' }), NEUTRAL,'renamed row must show neutral instead of unavailability');
  assert.strictEqual(statusWord({ device: { name: 'A' }, appliedName: 'A', runtimeStatus: null, runtimeStatusError: null, field: 'device_status' }), NEUTRAL,'initial pending must be neutral');
  assert.strictEqual(statusWord({ device: { name: 'A' }, appliedName: 'A', runtimeStatus: null, runtimeStatusError: 'relay offline', field: 'device_status' }), UNAVAILABLE,'request failure must stay UNAVAILABLE');
  assert.strictEqual(statusWord({ device: { name: 'A' }, appliedName: 'A', runtimeStatus: { device_status: RUNNING }, runtimeStatusError: null, field: 'device_status' }), RUNNING,'successful response must replace neutral');
  assert.strictEqual(statusWord({ device: { name: 'A' }, appliedName: 'A', runtimeStatus: { device_status: 'MAGIC' }, runtimeStatusError: null, field: 'device_status' }), UNAVAILABLE,'unknown runtime field keeps unavailable');

  console.log('status-word tests passed');
})().catch(error => {
  console.error(error);
  process.exitCode = 1;
});
