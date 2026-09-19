'use strict';

// JR-only deterministic local fake runtime. Never connect to Docker or user data.
const assert = require('assert');
const fs = require('fs');
const os = require('os');
const path = require('path');
const net = require('net');
const createProvider = require('../src/packages/MCSModbusToolkit/server');
const request = {version: 1, request_id: 'toolkit-local-test', operation: 'load', payload: {}};
const provider = createProvider();
const call = (ws, value) => new Promise(resolve => provider.onmessage(ws, resolve, [value]));

(async () => {
  const unauthorized = await call({_osjs_client: false}, request);
  assert.strictEqual(unauthorized.ok, false);
  assert.strictEqual(unauthorized.error.code, 'INVALID_REQUEST');
  const unallowed = await call({_osjs_client: true}, {...request, operation: 'restart'});
  assert.strictEqual(unallowed.ok, false);
  const badId = await call({_osjs_client: true}, {...request, request_id: ''});
  assert.strictEqual(badId.error.code, 'INVALID_REQUEST');
  console.log('provider requires authenticated OS.js session, v1 ID and allowlisted operation: checked');

  const root = fs.mkdtempSync(path.join(os.tmpdir(), 'toolkit-memory-relay-'));
  const previous = process.env.OSJS_DATA_DIR;
  const socketPath = path.join(root, 'run', 'modbus-simulator.sock');
  fs.mkdirSync(path.dirname(socketPath));
  const received = [];
  const server = net.createServer(conn => {
    let buffer = Buffer.alloc(0);
    conn.on('data', chunk => {
      buffer = Buffer.concat([buffer, chunk]);
      if (buffer.length < 4) return;
      const size = buffer.readUInt32BE(0);
      if (buffer.length < size + 4) return;
      const payload = JSON.parse(buffer.slice(4, 4 + size).toString('utf8'));
      received.push(payload);
      const reply = Buffer.from(JSON.stringify({version: 1, request_id: payload.request_id, ok: true,
        result: {document: {devices: []}}}));
      const header = Buffer.alloc(4); header.writeUInt32BE(reply.length);
      conn.end(Buffer.concat([header, reply]));
    });
  });
  try {
    process.env.OSJS_DATA_DIR = root;
    await new Promise((resolve, reject) => {
      server.once('error', reject);
      server.listen(socketPath, resolve);
    });
    const response = await call({_osjs_client: true}, request);
    assert.strictEqual(response.version, 1);
    assert.strictEqual(response.request_id, request.request_id);
    assert.strictEqual(response.ok, true);
    assert.deepStrictEqual(response.result.document, {devices: []});
    assert.deepStrictEqual(received, [request]);
    console.log('Toolkit-owned provider preserves v1 framing and response envelope on isolated Unix socket: checked');
    await new Promise((resolve, reject) => server.close(error => error ? reject(error) : resolve()));
    const unavailable = await call({_osjs_client: true}, {...request, request_id: 'no-runtime'});
    assert.strictEqual(unavailable.ok, false);
    assert.strictEqual(unavailable.error.code, 'RUNTIME_UNAVAILABLE');
    assert.strictEqual(unavailable.request_id, 'no-runtime');
    console.log('unavailable runtime returns correlated explicit error: checked');
    console.log('UMIG-004 Toolkit Memory relay cases complete');
  } finally {
    if (server.listening) await new Promise(resolve => server.close(resolve));
    if (previous === undefined) delete process.env.OSJS_DATA_DIR;
    else process.env.OSJS_DATA_DIR = previous;
    fs.rmSync(root, {recursive: true, force: true}); // Only our own mkdtemp path.
  }
})().catch(error => { console.error(error); process.exitCode = 1; });
