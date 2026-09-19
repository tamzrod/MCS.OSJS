'use strict';

// Future JR test: pure Toolkit Replicator adapter, no backend or configuration write.
const assert = require('assert');
const {copy, blankBlock, blankDevice, normalizeDocument, validateDocument, displayStatus} =
  require('../src/packages/MCSModbusToolkit/replicator-adapter');
const suggestion = {port: 5021, unit_id: 1, owner: 'replicator', status: 'AVAILABLE'};
const make = () => blankDevice(1, suggestion);

assert.deepStrictEqual(blankBlock(), {function: 3, start: 0, count: 16, scan_rate_ms: 1000});
const device = make();
assert.strictEqual(validateDocument({devices: [device]}), null);
assert.deepStrictEqual(normalizeDocument({devices: [{...device, pull_blocks: null, pull_block: blankBlock()}]}).devices[0].pull_blocks, [blankBlock()]);
const cloned = copy(device); cloned.pull_blocks[0].count = 99;
assert.strictEqual(device.pull_blocks[0].count, 16);
console.log('canonical defaults, legacy pull block normalization and copy isolation: checked');

const bad = copy(device); bad.endpoint = 'bad-endpoint';
assert.match(validateDocument({devices: [bad]}), /endpoint/);
bad.endpoint = '127.0.0.1:5020'; bad.pull_blocks[0].function = 5;
assert.match(validateDocument({devices: [bad]}), /Function/);
bad.pull_blocks[0].function = 3; bad.pull_blocks[0].start = 65535; bad.pull_blocks[0].count = 2;
assert.match(validateDocument({devices: [bad]}), /range/);
bad.pull_blocks[0].start = 0; bad.pull_blocks[0].count = 16; bad.pull_blocks[0].scan_rate_ms = 0;
assert.match(validateDocument({devices: [bad]}), /Scan Rate/);
bad.pull_blocks[0].scan_rate_ms = 1000;
bad.pull_blocks.push({function: 3, start: 18, count: 2, scan_rate_ms: 1000});
assert.match(validateDocument({devices: [bad]}), /gaps/);
bad.pull_blocks[1].start = 16;
assert.strictEqual(validateDocument({devices: [bad]}), null);
assert.match(validateDocument({devices: [device, {...device, name: 'rep-plc-1'}]}), /Duplicate/);
console.log('source, FC1-FC4 block/range/gap/cadence and case-insensitive duplicate validation: checked');

const sample = {name: device.name, running: true, source_status: 'ERROR', last_error: 'source timeout',
  blocks: [{index: 0, source_status: 'ERROR', last_error: 'source timeout'}]};
const observed = displayStatus(sample, device.name, null);
assert.strictEqual(observed.replicator, 'RUNNING');
assert.strictEqual(observed.source, 'ERROR');
assert.strictEqual(observed.blocks[0].source_status, 'ERROR');
assert.strictEqual(displayStatus(sample, 'other', null).replicator, 'UNKNOWN');
assert.strictEqual(displayStatus(sample, device.name, new Error('socket missing')).replicator, 'UNAVAILABLE');
assert.strictEqual(displayStatus(null, device.name, null).source, 'UNKNOWN');
console.log('direct runtime status, per-block errors and fail-closed unavailable/unknown: checked');
console.log('UMIG-005 Toolkit Replicator adapter checks complete');
