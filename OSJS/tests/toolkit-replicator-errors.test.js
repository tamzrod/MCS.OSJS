'use strict';

// Future JR unit test: Go-shaped ownership, validation, missing-device and
// destination-inspection results through the actual injected v1 contract.
const assert = require('assert');
const {createReplicatorContract} = require('../src/packages/MCSModbusToolkit/replicator-contract');
const document = {devices: [{name: 'Rep-A', enabled: true, endpoint: '127.0.0.1:5020', unit_id: 1,
  destination: {port: 5021, unit_id: 1, auto_port: false, auto_unit_id: false},
  pull_blocks: [{function: 3, start: 0, count: 16, scan_rate_ms: 1000}]}]};
(async () => {
  const calls = [];
  const api = createReplicatorContract(request => {
    calls.push(request);
    const response = {version: 1, request_id: request.request_id};
    if (request.operation === 'apply') return {...response, ok: false,
      error: {code: 'APPLY_FAILED', message: 'mma2 reservation owned by another producer'}};
    if (request.operation === 'status') return {...response, ok: false,
      error: {code: 'STATUS_FAILED', message: 'device "missing" not found'}};
    if (request.operation === 'suggest') return {...response, ok: true,
      result: {port: 5021, unit_id: 1, owner: 'simulator', status: 'IN USE'}};
    return {...response, ok: true, result: {document, suggestion: {port: 5022, unit_id: 1, owner: 'replicator', status: 'AVAILABLE'}}};
  });
  const collision = await api.apply(document).then(() => null, error => error);
  assert.strictEqual(collision.code, 'APPLY_FAILED');
  assert.match(collision.message, /owned by another producer/);
  const missing = await api.status('missing').then(() => null, error => error);
  assert.strictEqual(missing.code, 'STATUS_FAILED');
  assert.match(missing.message, /not found/);
  const inspected = await api.suggest({inspect: true, port: 5021, unit_id: 1});
  assert.strictEqual(inspected.owner, 'simulator');
  assert.strictEqual(inspected.status, 'IN USE');
  assert.strictEqual(calls.filter(call => call.operation === 'apply').length, 1);
  assert.deepStrictEqual(calls.map(call => call.operation), ['apply', 'status', 'suggest']);
  assert.deepStrictEqual(calls[2].payload, {inspect: true, port: 5021, unit_id: 1});
  console.log('single apply, ownership collision, missing device and Go destination inspection: checked');
  console.log('UMIG-005 Toolkit Replicator typed-error cases complete');
})().catch(error => { console.error(error); process.exitCode = 1; });
