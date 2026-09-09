'use strict';

const assert = require('assert');
const {createStatusPoller, createRuntimeMessageController} = require('./runtime-status-poller');

const deferred = () => {
  let resolve;
  let reject;
  const promise = new Promise((yes, no) => { resolve = yes; reject = no; });
  return {promise, resolve, reject};
};
const flush = () => new Promise(resolve => setImmediate(resolve));

(async () => {
  const calls = [];
  const pending = [];
  const statuses = [];
  const failures = [];
  const timers = new Map();
  let timerSequence = 0;
  const poller = createStatusPoller({
    requestStatus: name => {
      calls.push(name);
      const next = deferred();
      pending.push(next);
      return next.promise;
    },
    onStatus: (name, status) => statuses.push({name, status}),
    onUnavailable: (name, error) => failures.push({name, error: error.message}),
    setTimer: (callback, delay) => {
      const id = ++timerSequence;
      timers.set(id, {callback, delay});
      return id;
    },
    clearTimer: id => timers.delete(id)
  });

  poller.select('device-a');
  assert.deepStrictEqual(calls, ['device-a'], 'selection must request immediately');
  pending[0].resolve({device_status: 'WAITING'});
  await flush();
  assert.deepStrictEqual(statuses, [{name: 'device-a', status: {device_status: 'WAITING'}}]);
  assert.strictEqual(timers.size, 1);
  const cadence = [...timers.values()][0];
  assert.strictEqual(cadence.delay, 1000);
  timers.clear();
  cadence.callback();
  assert.deepStrictEqual(calls, ['device-a', 'device-a'], 'timer must poll only the selected device');

  poller.select('device-b');
  assert.deepStrictEqual(calls, ['device-a', 'device-a', 'device-b'], 'new selection must replace and refresh immediately');
  pending[1].resolve({device_status: 'RUNNING'});
  await flush();
  assert.strictEqual(statuses.length, 1, 'stale response must be ignored');

  pending[2].reject(new Error('runtime unavailable'));
  await flush();
  assert.deepStrictEqual(failures, [{name: 'device-b', error: 'runtime unavailable'}]);
  assert.strictEqual(statuses.length, 1, 'failure must not fabricate a running status');

  poller.stop();
  assert.strictEqual(timers.size, 0, 'stop must cancel the refresh loop');
  poller.refresh();
  assert.strictEqual(calls.length, 3, 'stopped poller must not issue requests');

  const messages = [];
  const controller = createRuntimeMessageController({publish: (message, error) => messages.push({message, error})});
  const ingestError = {device_status: 'ERROR', raw_ingest_error: 'Raw Ingest rejected frame'};
  controller.status(ingestError);
  controller.status(ingestError);
  assert.deepStrictEqual(messages, [{message: 'Simulator error: Raw Ingest rejected frame', error: true}], 'identical polling errors must be deduplicated');
  controller.unavailable(new Error('relay offline'));
  controller.unavailable(new Error('relay offline'));
  assert.strictEqual(messages.length, 2, 'identical unavailable results must be deduplicated');
  controller.status({device_status: 'RUNNING'});
  controller.status({device_status: 'RUNNING'});
  assert.deepStrictEqual(messages[2], {message: 'Simulator runtime status recovered.', error: false});
  assert.strictEqual(messages.length, 3, 'steady running must not add bottom-bar messages');
})().catch(error => {
  console.error(error);
  process.exitCode = 1;
});
