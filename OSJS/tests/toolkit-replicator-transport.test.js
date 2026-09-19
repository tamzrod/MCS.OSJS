'use strict';

// Future JR test for Toolkit-owned Replicator transport. No sockets/services.
const assert = require('assert');
const {createReplicatorTransport} = require('../src/packages/MCSModbusToolkit/replicator-transport');
(async () => {
  let receive;
  const sent = [];
  const timers = [];
  const proc = {on(name, handler) { assert.strictEqual(name, 'ws:message'); receive = handler; }, send(request) { sent.push(request); }};
  const transport = createReplicatorTransport(proc, {setTimer: fn => { timers.push(fn); return fn; }, clearTimer: () => {}});
  const request = {version: 1, request_id: 'mcs-replicator-test-1', operation: 'load', payload: {}};
  const pending = transport.send(request);
  await assert.rejects(transport.send(request), /duplicate/);
  await assert.rejects(transport.send({...request, request_id: 'memory-id'}), /Invalid/);
  receive({version: 1, request_id: 'unrelated', ok: true, result: {}});
  const reply = {version: 1, request_id: request.request_id, ok: false, error: {code: 'LOAD_FAILED', message: 'unavailable'}};
  receive(reply);
  assert.deepStrictEqual(await pending, reply); // Full envelope, not unwrapped.
  assert.deepStrictEqual(sent, [request]);
  console.log('correlated replies, duplicate rejection, full envelope and namespace: checked');

  const timed = transport.send({...request, request_id: 'mcs-replicator-test-2'});
  timers[timers.length - 1]();
  await assert.rejects(timed, /timed out/);
  const closing = transport.send({...request, request_id: 'mcs-replicator-test-3'});
  transport.close(); transport.close();
  await assert.rejects(closing, /closed/);
  await assert.rejects(transport.send({...request, request_id: 'mcs-replicator-test-4'}), /closed/);
  console.log('timeout, disposal and send-after-close: checked');
  console.log('UMIG-005 Toolkit Replicator transport checks complete');
})().catch(error => { console.error(error); process.exitCode = 1; });
