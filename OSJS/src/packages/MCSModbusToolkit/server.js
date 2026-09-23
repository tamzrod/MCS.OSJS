'use strict';

// One Toolkit-owned authenticated OS.js WebSocket provider, two isolated Go
// Unix sockets. No legacy package import, HTTP endpoint or direct config write.
const net = require('net');
const path = require('path');
const VERSION = 1;
const MAX_MESSAGE = 1024 * 1024;
const MEMORY_OPS = new Set(['load', 'apply', 'status', 'mma-load', 'mma-apply']);
const REPLICATOR_OPS = new Set(['load', 'apply', 'status', 'suggest']);
const runtimeSocket = service => path.join(process.env.OSJS_DATA_DIR || process.cwd(), 'run',
  service === 'replicator' ? 'modbus-replicator.sock' : 'modbus-simulator.sock');

const callRuntime = (request, service) => new Promise((resolve, reject) => {
  let body;
  try { body = Buffer.from(JSON.stringify(request)); } catch (error) { reject(error); return; }
  if (!body.length || body.length > MAX_MESSAGE) { reject(new Error(`${service} request is too large or empty`)); return; }
  const header = Buffer.alloc(4);
  header.writeUInt32BE(body.length);
  const socket = net.createConnection(runtimeSocket(service));
  let buffer = Buffer.alloc(0);
  let expected = null;
  let settled = false;
  const timeout = setTimeout(() => socket.destroy(new Error(`${service} runtime request timed out`)),
    request.operation === 'mma-apply' ? 35000 : 25000);
  const finish = (error, response) => {
    if (settled) return;
    settled = true;
    clearTimeout(timeout);
    if (error) reject(error);
    else resolve(response);
    socket.destroy();
  };
  socket.on('connect', () => socket.write(Buffer.concat([header, body])));
  socket.on('data', chunk => {
    buffer = Buffer.concat([buffer, chunk]);
    if (expected === null && buffer.length >= 4) {
      expected = buffer.readUInt32BE(0);
      if (!expected || expected > MAX_MESSAGE) { finish(new Error(`Invalid ${service} runtime response length`)); return; }
    }
    if (buffer.length > MAX_MESSAGE + 4) { finish(new Error(`${service} runtime response is too large`)); return; }
    if (expected !== null && buffer.length >= expected + 4) {
      try { finish(null, JSON.parse(buffer.slice(4, expected + 4).toString('utf8'))); }
      catch (error) { finish(error); }
    }
  });
  socket.on('error', error => finish(error));
  socket.on('end', () => finish(new Error(`${service} runtime closed before responding`)));
  socket.on('close', () => finish(new Error(`${service} runtime connection closed`)));
});
const rejectRequest = (respond, request, code, message) => respond({
  version: VERSION,
  request_id: request && typeof request.request_id === 'string' ? request.request_id : '',
  ok: false,
  error: {code, message}
});
module.exports = () => ({
  onmessage(ws, respond, args) {
    const request = args && args[0];
    if (!ws._osjs_client || !request || request.version !== VERSION ||
        typeof request.request_id !== 'string' || !request.request_id.trim() ||
        !request.payload || typeof request.payload !== 'object' || Array.isArray(request.payload)) {
      rejectRequest(respond, request, 'INVALID_REQUEST', 'Rejected Toolkit runtime request');
      return;
    }
    // Replicator contract generates its own namespaced ID. Legacy Memory test
    // IDs remain accepted and continue to resolve ONLY to the Simulator socket.
    const service = request.request_id.startsWith('mcs-replicator-') ? 'replicator' : 'simulator';
    const allowed = service === 'replicator' ? REPLICATOR_OPS : MEMORY_OPS;
    if (!allowed.has(request.operation)) {
      rejectRequest(respond, request, 'INVALID_REQUEST', 'Rejected Toolkit runtime operation');
      return;
    }
    callRuntime(request, service).then(response => respond(response)).catch(error => {
      rejectRequest(respond, request, 'RUNTIME_UNAVAILABLE', error.message || `${service} runtime unavailable`);
    });
  }
});
