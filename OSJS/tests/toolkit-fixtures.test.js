// UMIG-003-T fixture contract. JR executes this; CODE does not claim a run.
const assert = require('assert');
const {UNKNOWN, snapshot} = require('../src/packages/MCSModbusToolkit/fixtures');

assert.strictEqual(UNKNOWN, 'UNKNOWN');
console.log('fixture unknown constant: checked');
const first = snapshot();
assert.deepStrictEqual(first.runtime, {mma2: UNKNOWN, simulator: UNKNOWN, replicator: UNKNOWN});
assert.deepStrictEqual(first.comms, {network: UNKNOWN, tcp: UNKNOWN, modbus: UNKNOWN, mma2: UNKNOWN});
assert(Object.values(first.diagnostics).every(value => value === 'UNAVAILABLE'));
console.log('runtime, COMMS and diagnostics fail closed: checked');
assert.strictEqual(first.memory.devices.length, 1);
assert.strictEqual(first.memory.devices[0].name, 'Fixture Sim-PLC-1');
assert.strictEqual(first.memory.devices[0].random_runtime.fc1_interval_ms, 0);
assert.strictEqual(first.memory.devices[0].random_runtime.fc3_interval_ms, 1000);
assert.strictEqual(first.replicator.devices.length, 1);
assert.strictEqual(first.replicator.devices[0].name, 'Fixture Rep-PLC-1');
assert.strictEqual(first.replicator.devices[0].destination.owner, UNKNOWN);
assert.strictEqual(first.replicator.devices[0].pull_blocks[0].function, 3);
console.log('Memory and Replicator example shapes: checked');
first.memory.devices[0].name = 'mutated';
first.comms.network = 'OK';
const second = snapshot();
assert.strictEqual(second.memory.devices[0].name, 'Fixture Sim-PLC-1');
assert.strictEqual(second.comms.network, UNKNOWN);
console.log('fresh fixture snapshots cannot mutate later windows: checked');
console.log('UMIG-003 fixture contract checks complete');
