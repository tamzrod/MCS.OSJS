const net = require('net');
const path = require('path');

const VERSION = 1;
const MAX_MESSAGE = 1024 * 1024;
const ALLOWED = new Set(['load', 'apply', 'status']);

const runtimeSocket = () => path.join(process.env.OSJS_DATA_DIR || process.cwd(), 'run', 'modbus-simulator.sock');

const callRuntime = request => new Promise((resolve, reject) => {
  const body = Buffer.from(JSON.stringify(request));
  if (body.length > MAX_MESSAGE) return reject(new Error('Simulator request is too large'));
  const header = Buffer.alloc(4);
  header.writeUInt32BE(body.length);
  const socket = net.createConnection(runtimeSocket());
  const chunks = [];
  let expected = null;
  let received = 0;
  const timeout = setTimeout(() => socket.destroy(new Error('Simulator runtime request timed out')), 25000);
  const finish = error => {
    clearTimeout(timeout);
    if (error) reject(error);
  };
  socket.on('connect', () => socket.write(Buffer.concat([header, body])));
  socket.on('data', chunk => {
    chunks.push(chunk);
    received += chunk.length;
    const data = Buffer.concat(chunks);
    if (expected === null && data.length >= 4) expected = data.readUInt32BE(0);
    if (expected !== null && (expected > MAX_MESSAGE || expected < 1)) return socket.destroy(new Error('Invalid Simulator runtime response'));
    if (expected !== null && received >= expected + 4) {
      try {
        resolve(JSON.parse(data.slice(4, expected + 4).toString('utf8')));
        clearTimeout(timeout);
        socket.end();
      } catch (error) {
        socket.destroy(error);
      }
    }
  });
  socket.on('error', finish);
});

module.exports = () => ({
  onmessage(ws, respond, args) {
    const request = args && args[0];
    if (!ws._osjs_client || !request || request.version !== VERSION || !ALLOWED.has(request.operation)) {
      respond({version: VERSION, request_id: request && request.request_id || '', ok: false, error: {code: 'INVALID_REQUEST', message: 'Rejected Simulator request'}});
      return;
    }
    callRuntime(request)
      .then(response => respond(response))
      .catch(error => respond({version: VERSION, request_id: request.request_id || '', ok: false, error: {code: 'RUNTIME_UNAVAILABLE', message: error.message}}));
  }
});
