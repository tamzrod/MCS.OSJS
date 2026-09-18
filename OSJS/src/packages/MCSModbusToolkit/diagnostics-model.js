'use strict';

// UMIG-CF-003: dormant, pure Diagnostics mapping. Not imported by the UI.
// Inputs must come from a future verified Toolkit-owned adapter, never fixtures.
const UNKNOWN = 'UNKNOWN';
const UNAVAILABLE = 'UNAVAILABLE';
const SIM_STATES = new Set(['RUNNING', 'WAITING', 'STOPPED', 'ERROR']);
const SOURCE_STATES = new Set(['WAITING', 'OK', 'ERROR', 'DISABLED']);
const record = value => value !== null && typeof value === 'object' && !Array.isArray(value);
const state = (value, allowed) => allowed.has(value) ? value : UNKNOWN;
const observed = item => record(item) && item.fresh === true &&
  typeof item.expectedName === 'string' && item.expectedName.trim() !== '' &&
  record(item.result);
const failureText = item => record(item.error) && typeof item.error.message === 'string' &&
  item.error.message.trim() ? item.error.message : 'Runtime observation unavailable';

// Simulator status() returns {status:{name, device_status, mma2_status}}.
// Replicator status() returns the device status directly, NOT {status}.
// Neither response proves overall Docker service or MMA2 supervisor health.
const mapDiagnostics = ({memory, replicator} = {}) => {
  const view = {
    runtime: {mma2: UNKNOWN, simulator: UNKNOWN, replicator: UNKNOWN},
    devices: {
      memory: {mma2: UNKNOWN, simulator: UNKNOWN},
      replicator: {runtime: UNKNOWN, source: UNKNOWN, blocks: []}
    },
    errors: {memory: null, replicator: null},
    diagnostics: {runtime_mode: UNAVAILABLE, bin_path: UNAVAILABLE, data_path: UNAVAILABLE},
    controls: {start: false, stop: false}
  };

  if (record(memory) && memory.error) {
    view.devices.memory.mma2 = UNAVAILABLE;
    view.devices.memory.simulator = UNAVAILABLE;
    view.errors.memory = failureText(memory);
  } else if (observed(memory)) {
    const result = memory.result.status;
    if (record(result) && result.name === memory.expectedName) {
      view.devices.memory.mma2 = state(result.mma2_status, SIM_STATES);
      view.devices.memory.simulator = state(result.device_status, SIM_STATES);
    }
  }

  if (record(replicator) && replicator.error) {
    view.devices.replicator.runtime = UNAVAILABLE;
    view.devices.replicator.source = UNAVAILABLE;
    view.errors.replicator = failureText(replicator);
  } else if (observed(replicator)) {
    const result = replicator.result;
    if (result.name === replicator.expectedName && typeof result.enabled === 'boolean' &&
        typeof result.running === 'boolean' && Array.isArray(result.blocks)) {
      // Report device/poller observations only; never extrapolate service status.
      view.devices.replicator.runtime = result.running && !result.enabled ? UNKNOWN :
        result.running ? 'RUNNING' : 'STOPPED';
      view.devices.replicator.source = state(result.source_status, SOURCE_STATES);
      view.devices.replicator.blocks = result.blocks.map((block, index) => ({
        index,
        source: record(block) && block.index === index ? state(block.source_status, SOURCE_STATES) : UNKNOWN,
        running: record(block) && block.index === index && typeof block.running === 'boolean' ?
          (block.running ? 'RUNNING' : 'STOPPED') : UNKNOWN
      }));
    }
  }
  return view;
};

module.exports = {UNKNOWN, UNAVAILABLE, mapDiagnostics};
