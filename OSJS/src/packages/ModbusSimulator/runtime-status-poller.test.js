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

  const paused = [];
  const pausedPending = [];
  const pausedStatuses = [];
  const pausedFailures = [];
  const pausedTimers = new Map();
  let pausedTimerSequence = 0;
  const pausePoller = createStatusPoller({
    requestStatus: name => {
      paused.push(name);
      const next = deferred();
      pausedPending.push(next);
      return next.promise;
    },
    onStatus: (name, status) => pausedStatuses.push({name, status}),
    onUnavailable: (name, error) => pausedFailures.push({name, error: error.message}),
    setTimer: (callback, delay) => {
      const id = ++pausedTimerSequence;
      pausedTimers.set(id, {callback, delay});
      return id;
    },
    clearTimer: id => pausedTimers.delete(id)
  });

  pausePoller.select('applied-device');
  assert.deepStrictEqual(paused, ['applied-device'], 'apply start must snapshot the polled identity');
  pausePoller.select(null);   // Save & Apply pause
  const pausedRequest = paused.length;
  assert.strictEqual(pausedTimers.size, 0, 'pause must cancel the pending cadence');
  pausedPending[0].resolve({device_status: 'RUNNING'});
  await flush();
  assert.deepStrictEqual(pausedStatuses, [], 'in-flight response racingther apply transition must be dropped');
  assert.deepStrictEqual(pausedFailures, [], 'apply transition must not fabricate an unavailable result');
  assert.strictEqual(paused.length, pausedRequest, 'no status request may start while applying');
  pausePoller.select('applied-device', true);
  assert.strictEqual(paused.length, pausedRequest + 1, 'successful apply resume must force one immediate refresh');
  pausedPending[pausedPending.length - 1].resolve({device_status: 'WAITING'});
  await flush();
  assert.deepStrictEqual(pausedStatuses, [{name: 'applied-device', status: {device_status: 'WAITING'}}], 'resumed poll must reflect the applied device');
  assert.strictEqual(pausedTimers.size, 1, 'pause with resume must restore the refresh cadence');
  pausePoller.stop();

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
