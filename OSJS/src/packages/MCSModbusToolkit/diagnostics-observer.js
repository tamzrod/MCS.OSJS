'use strict';

// Read-only, canonical observation. The existing Toolkit contracts own v1
// validation and relay; this module NEVER applies, suggests, probes or controls
// a service. The first persisted device is an explicitly labelled sample, not
// evidence of fleet-wide or supervisor health.
const {mapDiagnostics} = require('./diagnostics-model');

const messageOf = error => error && typeof error.message === 'string' && error.message.trim() ?
  error.message : String(error || 'Unknown observation failure');

const observe = async (contract, label) => {
  let name = null;
  let count = null;
  try {
    const loaded = await contract.load();
    if (!loaded || !loaded.document || !Array.isArray(loaded.document.devices)) {
      throw new Error('canonical device document is invalid');
    }
    count = loaded.document.devices.length;
    if (!count) return {name, count, observation: null};
    const first = loaded.document.devices[0];
    if (!first || typeof first.name !== 'string' || !first.name.trim()) {
      throw new Error('first canonical device has no valid name');
    }
    name = first.name;
    const result = await contract.status(name);
    return {name, count, observation: {fresh: true, expectedName: name, result}};
  } catch (error) {
    return {name, count, observation: {error: {message: `${label}: ${messageOf(error)}`}}};
  }
};

const collectDiagnostics = async ({memory, replicator}) => {
  if (!memory || typeof memory.load !== 'function' || typeof memory.status !== 'function' ||
      !replicator || typeof replicator.load !== 'function' || typeof replicator.status !== 'function') {
    throw new TypeError('Diagnostics requires the existing read-only Memory and Replicator contracts');
  }
  // Independent failures must not conceal a healthy observation of the other
  // device. No information from the Toolkit's fixture renderer is consumed.
  const [memoryResult, replicatorResult] = await Promise.all([
    observe(memory, 'Memory'), observe(replicator, 'Replicator')
  ]);
  return {
    memory: memoryResult,
    replicator: replicatorResult,
    view: mapDiagnostics({memory: memoryResult.observation, replicator: replicatorResult.observation})
  };
};

module.exports = {collectDiagnostics};
