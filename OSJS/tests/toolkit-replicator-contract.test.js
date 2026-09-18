'use strict';

// Future JR-only focused test for dormant UMIG-CF-002 code; NOT run by CODE.
const assert = require('assert');
const {VERSION, ReplicatorContractError, createReplicatorContract} = require('../src/packages/MCSModbusToolkit/replicator-contract');

const device = {name: 'Test Rep', enabled: true, endpoint: '192.0.2.1:502', unit_id: 1,
  pull_blocks: [{function: 3, start: 0, count: 16, scan_rate_ms: 1000}], destination: {port: 5021, unit_id: 1}};
const document = {devices: [device]};
const suggestion = {port: 5021, unit_id: 1, owner: 'replicator', status: 'AVAILABLE'};
const status = {name: 'Test Rep', enabled: true, running: false, cycles: 0, source_status: 'WAITING', blocks: []};
const respond = (req, result) => ({version: VERSION, request_id: req.request_id, ok: true, result});

(async () => {
  assert.strictEqual(VERSION, 1);
  assert.throws(() => createReplicatorContract(null), TypeError);
  console.log('protocol and injected transport: checked');

  const sent = [];
  const rep = createReplicatorContract(async req => {
    sent.push(req);
    if (req.operation === 'load') return respond(req, {document, suggestion});
    if (req.operation === 'apply') return respond(req, {document: req.payload.document, structural: false, message: 'applied', completed_at: '2026-09-19T00:00:00Z'});
    if (req.operation === 'status') return respond(req, status);
    return respond(req, suggestion);
  });
  assert.deepStrictEqual(await rep.load(), {document, suggestion});
  assert.deepStrictEqual(await rep.status('Test Rep'), status); // direct Go result, NOT {status}
  assert.deepStrictEqual(await rep.suggest(), suggestion);
  assert.deepStrictEqual(await rep.suggest({inspect: true, port: 5021, unit_id: 1}), suggestion);
  const edited = JSON.parse(JSON.stringify(document));
  const applying = rep.apply(edited);
  edited.devices[0].name = 'mutated after dispatch';
  assert.strictEqual((await applying).document.devices[0].name, 'Test Rep');
  assert.deepStrictEqual(sent.map(req => req.operation), ['load', 'status', 'suggest', 'suggest', 'apply']);
  assert.deepStrictEqual(sent[2].payload, {});
  assert.deepStrictEqual(sent[3].payload, {inspect: true, port: 5021, unit_id: 1});
  assert.strictEqual(new Set(sent.map(req => req.request_id)).size, 5);
  assert(sent.every(req => req.version === VERSION && typeof req.request_id === 'string' && req.request_id));
  console.log('load, direct status, suggest/inspect, apply snapshot, envelopes: checked');

  const before = sent.length;
  await assert.rejects(rep.apply(null), error => error instanceof ReplicatorContractError && error.code === 'INVALID_REQUEST');
  await assert.rejects(rep.status('  '), error => error.code === 'INVALID_REQUEST');
  await assert.rejects(rep.suggest({inspect: true, port: 0, unit_id: 1}), error => error.code === 'INVALID_REQUEST');
  await assert.rejects(rep.suggest({inspect: false, port: 5021, unit_id: 1}), error => error.code === 'INVALID_REQUEST');
  assert.strictEqual(sent.length, before);
  console.log('invalid requests cannot reach transport: checked');

  const failed = createReplicatorContract(req => respond(req, null));
  await assert.rejects(failed.load(), error => error.code === 'INVALID_RESPONSE');
  const runtimeError = createReplicatorContract(req => ({version: 1, request_id: req.request_id, ok: false, error: {code: 'APPLY_FAILED', message: 'reservation conflict'}}));
  await assert.rejects(runtimeError.apply(document), error => error.code === 'APPLY_FAILED' && error.message === 'reservation conflict');
  const unavailable = createReplicatorContract(() => Promise.reject(new Error('socket closed')));
  await assert.rejects(unavailable.status('Test Rep'), /socket closed/);
  const wrongId = createReplicatorContract(req => ({...respond(req, {document, suggestion}), request_id: req.request_id + '-wrong'}));
  await assert.rejects(wrongId.load(), error => error.code === 'INVALID_RESPONSE');
  const wrongVersion = createReplicatorContract(req => ({...respond(req, {document, suggestion}), version: 2}));
  await assert.rejects(wrongVersion.load(), error => error.code === 'INVALID_RESPONSE');
  const wrongStatus = createReplicatorContract(req => respond(req, {...status, name: 'Another Rep'}));
  await assert.rejects(wrongStatus.status('Test Rep'), error => error.code === 'INVALID_RESPONSE');
  const invalidSuggestion = createReplicatorContract(req => respond(req, {port: 0, unit_id: 1, owner: '', status: 'AVAILABLE'}));
  await assert.rejects(invalidSuggestion.suggest(), error => error.code === 'INVALID_RESPONSE');
  const incompleteApply = createReplicatorContract(req => respond(req, {document}));
  await assert.rejects(incompleteApply.apply(document), error => error.code === 'INVALID_RESPONSE');
  console.log('runtime/transport failures and malformed replies fail closed: checked');
  console.log('UMIG-CF-002 Replicator contract cases complete');
})().catch(error => {
  console.error(error);
  process.exitCode = 1;
});
