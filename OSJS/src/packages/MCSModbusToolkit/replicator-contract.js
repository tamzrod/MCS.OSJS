'use strict';

// UMIG-CF-002: dormant, transport-injected Replicator v1 contract.
// This module is intentionally NOT imported by the Toolkit UI or OS.js runtime.
const VERSION = 1;
let clientSequence = 0;

class ReplicatorContractError extends Error {
  constructor(code, message) {
    super(message);
    this.name = 'ReplicatorContractError';
    this.code = code;
  }
}

const record = value => value !== null && typeof value === 'object' && !Array.isArray(value);
const invalidRequest = message => new ReplicatorContractError('INVALID_REQUEST', message);
const invalidResponse = message => new ReplicatorContractError('INVALID_RESPONSE', message);
const validDocument = value => record(value) && Array.isArray(value.devices);
const validSuggestion = value => record(value) &&
  Number.isInteger(value.port) && value.port >= 1 && value.port <= 65535 &&
  Number.isInteger(value.unit_id) && value.unit_id >= 0 && value.unit_id <= 65535 &&
  typeof value.owner === 'string' && typeof value.status === 'string';

// The future Toolkit-owned OS.js transport supplies `send` and owns its timeout
// and disposal lifecycle. No socket, timer or persistence code is used here.
const createReplicatorContract = send => {
  if (typeof send !== 'function') throw new TypeError('Replicator transport must be a function');
  const instance = ++clientSequence;
  let sequence = 0;

  const request = async (operation, payload) => {
    const requestId = `mcs-replicator-${Date.now()}-${instance}-${++sequence}`;
    const response = await send({version: VERSION, request_id: requestId, operation, payload});
    if (!record(response) || response.version !== VERSION || response.request_id !== requestId || typeof response.ok !== 'boolean') {
      throw invalidResponse(`Invalid Replicator ${operation} response envelope`);
    }
    if (!response.ok) {
      const error = record(response.error) ? response.error : {};
      const code = typeof error.code === 'string' && error.code ? error.code : 'RUNTIME_ERROR';
      const message = typeof error.message === 'string' && error.message ? error.message : 'Replicator runtime request failed';
      throw new ReplicatorContractError(code, message);
    }
    const result = response.result;
    if (!record(result)) throw invalidResponse(`Replicator ${operation} response has no result`);
    if (operation === 'load' && (!validDocument(result.document) || !validSuggestion(result.suggestion))) {
      throw invalidResponse('Replicator load response requires a document and destination suggestion');
    }
    if (operation === 'apply' && (!validDocument(result.document) || typeof result.structural !== 'boolean' || typeof result.message !== 'string' || typeof result.completed_at !== 'string')) {
      throw invalidResponse('Replicator apply response is incomplete');
    }
    if (operation === 'status' && (result.name !== payload.name || typeof result.enabled !== 'boolean' || typeof result.running !== 'boolean' || typeof result.source_status !== 'string' || !Array.isArray(result.blocks))) {
      throw invalidResponse('Replicator status response is incomplete or belongs to another device');
    }
    if (operation === 'suggest' && !validSuggestion(result)) {
      throw invalidResponse('Replicator destination suggestion response is invalid');
    }
    return result;
  };

  return Object.freeze({
    load: () => request('load', {}),
    apply: document => {
      if (!validDocument(document)) return Promise.reject(invalidRequest('Replicator apply requires a document with devices'));
      let copy;
      try {
        copy = JSON.parse(JSON.stringify(document));
      } catch (_) {
        return Promise.reject(invalidRequest('Replicator document must be JSON-serializable'));
      }
      return request('apply', {document: copy});
    },
    status: name => {
      if (typeof name !== 'string' || !name.trim()) return Promise.reject(invalidRequest('Replicator status requires a device name'));
      return request('status', {name});
    },
    suggest: inspection => {
      if (inspection === undefined) return request('suggest', {});
      if (!record(inspection) || inspection.inspect !== true ||
          !Number.isInteger(inspection.port) || inspection.port < 1 || inspection.port > 65535 ||
          !Number.isInteger(inspection.unit_id) || inspection.unit_id < 0 || inspection.unit_id > 65535) {
        return Promise.reject(invalidRequest('Replicator inspection requires a valid port and unit_id'));
      }
      return request('suggest', {inspect: true, port: inspection.port, unit_id: inspection.unit_id});
    }
  });
};

module.exports = {VERSION, ReplicatorContractError, createReplicatorContract};
