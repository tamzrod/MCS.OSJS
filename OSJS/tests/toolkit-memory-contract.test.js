'use strict';

// Future JR test for dormant code-first Memory contract. NOT run by CODE.
const assert = require('assert');
const {VERSION, MemoryContractError, createMemoryContract} = require('../src/packages/MCSModbusToolkit/memory-contract');

const device = {name: 'Test Sim', enabled: true, mma2: {port: 5020, unit_id: 1}, random_runtime: {fc1_interval_ms: 0}};
const document = {devices: [device]};
const respond = (request, result) => ({version: VERSION, request_id: request.request_id, ok: true, result});

(async () => {
  assert.strictEqual(VERSION, 1);
  assert.throws(() => createMemoryContract(null), TypeError);
  console.log('version and injected transport: checked');

  const sent = [];
  const memory = createMemoryContract(async request => {
    sent.push(request);
    if (request.operation === 'status') return respond(request, {status: {name: request.payload.name, mma2_status: 'UNKNOWN'}});
    return respond(request, {document: request.payload.document || document});
  });
  assert.deepStrictEqual((await memory.load()).document, document);
  assert.deepStrictEqual(await memory.status('Test Sim'), {status: {name: 'Test Sim', mma2_status: 'UNKNOWN'}});
  const edited = JSON.parse(JSON.stringify(document));
  const applying = memory.apply(edited);
  edited.devices[0].name = 'changed after dispatch';
  assert.strictEqual((await applying).document.devices[0].name, 'Test Sim');
  assert.deepStrictEqual(sent.map(item => item.operation), ['load', 'status', 'apply']);
  assert.deepStrictEqual(sent.map(item => item.payload.name), [undefined, 'Test Sim', undefined]);
  assert.strictEqual(new Set(sent.map(item => item.request_id)).size, 3);
  assert(sent.every(item => item.version === 1 && typeof item.request_id === 'string' && item.request_id.length));
  console.log('load, status, snapshot apply, unique IDs and protocol envelope: checked');

  const invalidCallCount = sent.length;
  await assert.rejects(memory.apply(null), error => error instanceof MemoryContractError && error.code === 'INVALID_REQUEST');
  await assert.rejects(memory.status('  '), error => error.code === 'INVALID_REQUEST');
  assert.strictEqual(sent.length, invalidCallCount);
  console.log('invalid requests do not reach transport: checked');

  const runtimeFailure = createMemoryContract(request => ({version: 1, request_id: request.request_id, ok: false, error: {code: 'MMA2_NOT_READY', message: 'restart pending'}}));
  await assert.rejects(runtimeFailure.load(), error => error.code === 'MMA2_NOT_READY' && error.message === 'restart pending');
  const unavailable = createMemoryContract(() => Promise.reject(new Error('socket closed')));
  await assert.rejects(unavailable.load(), /socket closed/);
  console.log('runtime errors and transport failures propagated: checked');

  const wrongId = createMemoryContract(request => ({version: 1, request_id: request.request_id + '-old', ok: true, result: {document}}));
  await assert.rejects(wrongId.load(), error => error.code === 'INVALID_RESPONSE');
  const missingStatus = createMemoryContract(request => respond(request, {}));
  await assert.rejects(missingStatus.status('Test Sim'), error => error.code === 'INVALID_RESPONSE');
  const wrongVersion = createMemoryContract(request => ({...respond(request, {document}), version: 2}));
  await assert.rejects(wrongVersion.load(), error => error.code === 'INVALID_RESPONSE');
  console.log('wrong correlation, missing data and wrong protocol fail closed: checked');
  console.log('UMIG-CF-001 Memory contract cases complete');
})().catch(error => {
  console.error(error);
  process.exitCode = 1;
});
