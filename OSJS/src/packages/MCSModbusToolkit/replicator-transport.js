'use strict';

// Independent request lifecycle for Replicator. Preserve the entire v1 reply
// for replicator-contract.js to validate request ID, Go error code and shape.
const createReplicatorTransport = (proc, {setTimer = setTimeout, clearTimer = clearTimeout, timeoutMs = 26000} = {}) => {
  const pending = new Map();
  let closed = false;
  proc.on('ws:message', response => {
    if (closed || !response || typeof response.request_id !== 'string') return;
    const entry = pending.get(response.request_id);
    if (!entry) return;
    pending.delete(response.request_id);
    clearTimer(entry.timer);
    entry.resolve(response);
  });
  const send = request => new Promise((resolve, reject) => {
    if (closed) { reject(new Error('Toolkit Replicator transport closed')); return; }
    const id = request && request.request_id;
    if (typeof id !== 'string' || !id.startsWith('mcs-replicator-') || pending.has(id)) {
      reject(new Error('Invalid or duplicate Toolkit Replicator request ID')); return;
    }
    const timer = setTimer(() => {
      pending.delete(id);
      reject(new Error('Replicator runtime request timed out'));
    }, timeoutMs);
    pending.set(id, {resolve, reject, timer});
    try { proc.send(request); } catch (error) {
      clearTimer(timer); pending.delete(id); reject(error);
    }
  });
  const close = () => {
    if (closed) return;
    closed = true;
    pending.forEach(entry => {
      clearTimer(entry.timer);
      entry.reject(new Error('Toolkit Replicator window closed'));
    });
    pending.clear();
  };
  return {send, close};
};
module.exports = {createReplicatorTransport};
