'use strict';

// Keep the complete v1 envelope for memory-contract.js to validate, unlike
// the legacy Simulator UI's helper which unwraps response.result itself.
const createMemoryTransport = (proc, {setTimer = setTimeout, clearTimer = clearTimeout, timeoutMs = 26000} = {}) => {
  const pending = new Map();
  let closed = false;

  proc.on('ws:message', response => {
    if (closed || !response || typeof response.request_id !== 'string') return;
    const entry = pending.get(response.request_id);
    if (!entry) return; // Stale, unsolicited, or different-request response.
    pending.delete(response.request_id);
    clearTimer(entry.timer);
    entry.resolve(response); // Contract validates version, ID, ok and result.
  });

  const send = request => new Promise((resolve, reject) => {
    if (closed) { reject(new Error('Toolkit Memory transport closed')); return; }
    const id = request && request.request_id;
    if (typeof id !== 'string' || !id || pending.has(id)) {
      reject(new Error('Invalid or duplicate Toolkit Memory request ID'));
      return;
    }
    const timer = setTimer(() => {
      pending.delete(id);
      reject(new Error('Simulator runtime request timed out'));
    }, timeoutMs);
    pending.set(id, {resolve, reject, timer});
    try { proc.send(request); } catch (error) {
      clearTimer(timer);
      pending.delete(id);
      reject(error);
    }
  });

  const close = () => {
    if (closed) return;
    closed = true;
    pending.forEach(entry => {
      clearTimer(entry.timer);
      entry.reject(new Error('Toolkit Memory window closed'));
    });
    pending.clear();
  };

  return {send, close};
};

module.exports = {createMemoryTransport};
