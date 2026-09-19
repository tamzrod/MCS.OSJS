'use strict';

// Future JR fake-Unix-runtime test. Only mkdtemp sockets; no product service,
// deployed Docker volume or actual Modbus destination is contacted.
const assert = require('assert');
const fs = require('fs');
const os = require('os');
const path = require('path');
const net = require('net');
const createProvider = require('../src/packages/MCSModbusToolkit/server');
const provider = createProvider();
const call = (ws, value) => new Promise(resolve => provider.onmessage(ws, resolve, [value]));
const suggestion = {port: 5021, unit_id: 1, owner: 'replicator', status: 'AVAILABLE'};
const rep = (operation, id) => ({version: 1, request_id: `mcs-replicator-${id}`, operation, payload: {}});
const fakeServer = (socketPath, received, response) => net.createServer(conn => {
  let buffer = Buffer.alloc(0);
  conn.on('data', chunk => {
    buffer = Buffer.concat([buffer, chunk]);
    if (buffer.length < 4) return;
    const count = buffer.readUInt32BE(0);
    if (buffer.length < count + 4) return;
    const request = JSON.parse(buffer.slice(4, count + 4).toString('utf8'));
    received.push(request);
    const data = Buffer.from(JSON.stringify({version: 1, request_id: request.request_id, ok: true, result: response(request)}));
    const header = Buffer.alloc(4); header.writeUInt32BE(data.length);
    conn.end(Buffer.concat([header, data]));
  });
});
(async () => {
  const denied = await call({_osjs_client: false}, rep('load', 'denied'));
  assert.strictEqual(denied.error.code, 'INVALID_REQUEST');
  assert.strictEqual((await call({_osjs_client: true}, rep('restart', 'bad-op'))).error.code, 'INVALID_REQUEST');
  assert.strictEqual((await call({_osjs_client: true}, {...rep('load', 'bad-id'), request_id: ''})).error.code, 'INVALID_REQUEST');
  assert.strictEqual((await call({_osjs_client: true}, {version: 1, request_id: 'memory-test', operation: 'suggest', payload: {}})).error.code, 'INVALID_REQUEST');
  console.log('authentication, request ID and per-service allowlists: checked');

  const root = fs.mkdtempSync(path.join(os.tmpdir(), 'toolkit-replicator-relay-'));
  const previous = process.env.OSJS_DATA_DIR;
  const dir = path.join(root, 'run'); fs.mkdirSync(dir);
  const repReceived = []; const memoryReceived = [];
  const repServer = fakeServer(path.join(dir, 'modbus-replicator.sock'), repReceived, req => {
    if (req.operation === 'load') return {document: {devices: []}, suggestion};
    if (req.operation === 'suggest') return suggestion;
    if (req.operation === 'status') return {name: req.payload.name, enabled: true, running: true, source_status: 'ERROR',
      blocks: [{index: 0, running: true, source_status: 'ERROR', last_error: 'source down'}]};
    return {document: req.payload.document, structural: true, message: 'applied', completed_at: '2026-09-19T00:00:00Z'};
  });
  const memServer = fakeServer(path.join(dir, 'modbus-simulator.sock'), memoryReceived, () => ({document: {devices: []}}));
  const listen = server => new Promise((resolve, reject) => { server.once('error', reject); server.listen(server === repServer ? path.join(dir, 'modbus-replicator.sock') : path.join(dir, 'modbus-simulator.sock'), resolve); });
  const close = server => server.listening ? new Promise(resolve => server.close(resolve)) : Promise.resolve();
  try {
    process.env.OSJS_DATA_DIR = root;
    await listen(repServer); await listen(memServer);
    const doc = {devices: [{name: 'test'}]};
    const requests = [rep('load', 'load'), {...rep('suggest', 'suggest'), payload: {inspect: true, port: 5021, unit_id: 1}},
      {...rep('status', 'status'), payload: {name: 'test'}}, {...rep('apply', 'apply'), payload: {document: doc}}];
    const results = [];
    for (const request of requests) results.push(await call({_osjs_client: true}, request));
    assert.deepStrictEqual(results.map(item => item.request_id), requests.map(item => item.request_id));
    assert(results.every(item => item.version === 1 && item.ok === true));
    assert.deepStrictEqual(results[0].result.suggestion, suggestion);
    assert.strictEqual(results[2].result.blocks[0].last_error, 'source down');
    assert.deepStrictEqual(results[3].result.document, doc);
    const memory = {version: 1, request_id: 'memory-original-id', operation: 'load', payload: {}};
    assert.strictEqual((await call({_osjs_client: true}, memory)).ok, true);
    assert.deepStrictEqual(repReceived, requests);
    assert.deepStrictEqual(memoryReceived, [memory]);
    console.log('separate Toolkit Unix sockets, v1 envelopes, one apply and direct per-block status: checked');
    await close(repServer);
    const unavailable = await call({_osjs_client: true}, rep('load', 'down'));
    assert.strictEqual(unavailable.ok, false);
    assert.strictEqual(unavailable.request_id, 'mcs-replicator-down');
    assert.strictEqual(unavailable.error.code, 'RUNTIME_UNAVAILABLE');
    console.log('correlated Replicator-only runtime unavailable without Memory cross-routing: checked');
    console.log('UMIG-005 Toolkit Replicator relay checks complete');
  } finally {
    await close(repServer); await close(memServer);
    if (previous === undefined) delete process.env.OSJS_DATA_DIR;
    else process.env.OSJS_DATA_DIR = previous;
    fs.rmSync(root, {recursive: true, force: true});
  }
})().catch(error => { console.error(error); process.exitCode = 1; });
