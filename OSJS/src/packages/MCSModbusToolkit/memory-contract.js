'use strict';

// UMIG-CF-001: transport-neutral Simulator v1 contract, staged but NOT wired.
// No OS.js messaging, runtime connection or configuration write occurs here.
const VERSION = 1;
let clientSequence = 0;

class MemoryContractError extends Error {
  constructor(code, message) {
    super(message);
    this.name = 'MemoryContractError';
    this.code = code;
  }
}

const invalidResponse = message => new MemoryContractError('INVALID_RESPONSE', message);
const isRecord = value => value !== null && typeof value === 'object' && !Array.isArray(value);

// `send` is supplied by a future Toolkit-owned transport. Its only contract:
// accept a version-1 request and return the corresponding response or reject.
// This module is intentionally not imported by any Toolkit entry point yet.
const createMemoryContract = send => {
  if (typeof send !== 'function') throw new TypeError('Memory transport must be a function');
  const instance = ++clientSequence;
  let sequence = 0;

  const request = async (operation, payload) => {
    const requestId = `mcs-memory-${Date.now()}-${instance}-${++sequence}`;
    const reply = await send({version: VERSION, request_id: requestId, operation, payload});
    if (!isRecord(reply) || reply.version !== VERSION || reply.request_id !== requestId || typeof reply.ok !== 'boolean') {
      throw invalidResponse(`Invalid Simulator ${operation} response envelope`);
    }
    if (!reply.ok) {
      const error = isRecord(reply.error) ? reply.error : {};
      const code = typeof error.code === 'string' && error.code ? error.code : 'RUNTIME_ERROR';
      const message = typeof error.message === 'string' && error.message ? error.message : 'Simulator runtime request failed';
      throw new MemoryContractError(code, message);
    }
    if (!isRecord(reply.result)) throw invalidResponse(`Simulator ${operation} response has no result`);
    if ((operation === 'load' || operation === 'apply') && (!isRecord(reply.result.document) || !Array.isArray(reply.result.document.devices))) {
      throw invalidResponse(`Simulator ${operation} response has no valid document`);
    }
    if (operation === 'status' && !isRecord(reply.result.status)) {
      throw invalidResponse('Simulator status response has no status');
    }
    return reply.result;
  };

  return Object.freeze({
    load: () => request('load', {}),
    apply: document => {
      if (!isRecord(document) || !Array.isArray(document.devices)) {
        return Promise.reject(new MemoryContractError('INVALID_REQUEST', 'Memory apply requires a document with devices'));
      }
      // Snapshot at dispatch time so a caller cannot mutate an in-flight apply.
      let copy;
      try {
        copy = JSON.parse(JSON.stringify(document));
      } catch (_) {
        return Promise.reject(new MemoryContractError('INVALID_REQUEST', 'Memory document must be JSON-serializable'));
      }
      return request('apply', {document: copy});
    },
    status: name => {
      if (typeof name !== 'string' || !name.trim()) {
        return Promise.reject(new MemoryContractError('INVALID_REQUEST', 'Memory status requires a device name'));
      }
      return request('status', {name});
    }
  });
};

module.exports = {VERSION, MemoryContractError, createMemoryContract};
