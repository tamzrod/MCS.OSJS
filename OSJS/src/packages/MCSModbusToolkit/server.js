'use strict';

// Independent Toolkit provider: no imports from the legacy Simulator UI.
// OS.js authenticated application WebSocket -> existing Go runtime Unix socket.
const net = require('net');
const path = require('path');
const VERSION = 1;
const MAX_MESSAGE = 1024 * 1024;
const ALLOWED = new Set(['load', 'apply', 'status']);
const runtimeSocket = () => path.join(process.env.OSJS_DATA_DIR || process.cwd(), 'run', 'modbus-simulator.sock');

const callRuntime = request => new Promise((resolve, reject) => {
  let body;
  try { body = Buffer.from(JSON.stringify(request)); } catch (error) { reject(error); return; }
  if (!body.length || body.length > MAX_MESSAGE) { reject(new Error('Simulator request is too large or empty')); return; }
  const header = Buffer.alloc(4);
  header.writeUInt32BE(body.length);
  const socket = net.createConnection(runtimeSocket());
  let buffer = Buffer.alloc(0);
  let expected = null;
  let settled = false;
  const timeout = setTimeout(() => socket.destroy(new Error('Simulator runtime request timed out')), 25000);
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
      if (!expected || expected > MAX_MESSAGE) { finish(new Error('Invalid Simulator runtime response length')); return; }
    }
    if (expected !== null && buffer.length >= expected + 4) {
      try { finish(null, JSON.parse(buffer.slice(4, expected + 4).toString('utf8'))); }
      catch (error) { finish(error); }
    } else if (buffer.length > MAX_MESSAGE + 4) {
      finish(new Error('Simulator runtime response is too large'));
    }
  });
  socket.on('error', error => finish(error));
  socket.on('end', () => finish(new Error('Simulator runtime closed before responding')));
  socket.on('close', () => finish(new Error('Simulator runtime connection closed')));
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
        !ALLOWED.has(request.operation) || !request.payload ||
        typeof request.payload !== 'object' || Array.isArray(request.payload)) {
      rejectRequest(respond, request, 'INVALID_REQUEST', 'Rejected Toolkit Memory request');
      return;
    }
    callRuntime(request).then(response => respond(response)).catch(error => {
      rejectRequest(respond, request, 'RUNTIME_UNAVAILABLE', error.message || 'Simulator runtime unavailable');
    });
  }
});
