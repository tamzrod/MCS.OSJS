'use strict';

// Pure read-only Diagnostics mapping. Inputs are freshly correlated responses
// from existing Toolkit contracts, NEVER fixture snapshots or global probes.
const UNKNOWN = 'UNKNOWN';
const UNAVAILABLE = 'UNAVAILABLE';
const SIM_STATES = new Set(['RUNNING', 'WAITING', 'STOPPED', 'ERROR', 'IDLE']);
const SOURCE_STATES = new Set(['WAITING', 'OK', 'ERROR', 'DISABLED']);
// Observational persistence health states (PERSIST-021). These are display-only
// classifications of the runtime's reported persistence health; they never grant
// control or bypass of restore/sealing gates.
const PERSIST_SNAPSHOT_STATES = new Set(['ready', 'missing', 'invalid', 'incompatible', 'disabled', 'unsealed', 'empty']);
const PERSIST_RESTORE_STATES = new Set(['none', 'disabled', 'unsealed', 'empty', 'missing_snapshot', 'invalid_snapshot', 'incompatible_snapshot', 'area_not_ready', 'unknown_area', 'raw_ingest_write', 'raw_ingest_response', 'incomplete', 'no_sealing_flag', 'unseal_write', 'unseal_response']);
const record = value => value !== null && typeof value === 'object' && !Array.isArray(value);
const state = (value, allowed) => allowed.has(value) ? value : UNKNOWN;
const observed = item => record(item) && item.fresh === true &&
  typeof item.expectedName === 'string' && item.expectedName.trim() !== '' &&
  record(item.result);
const failureText = item => record(item.error) && typeof item.error.message === 'string' &&
  item.error.message.trim() ? item.error.message : 'Runtime observation unavailable';

// Observational persistence health mapping (PERSIST-021). Display-only: it
// classifies the runtime's reported persistence health and never grants control
// or bypass of restore/sealing gates. It fails closed to UNKNOWN for any
// malformed field, and returns null when no persistence block was reported.
const mapPersistenceHealth = persistence => {
  if (!record(persistence)) return null;
  return {
    configured: typeof persistence.configured === 'boolean' ? persistence.configured : UNKNOWN,
    sealed: typeof persistence.sealed === 'boolean' ? persistence.sealed : UNKNOWN,
    healthy: typeof persistence.healthy === 'boolean' ? persistence.healthy : UNKNOWN,
    snapshot_health: state(persistence.snapshot_health, PERSIST_SNAPSHOT_STATES),
    restore_outcome: state(persistence.restore_outcome, PERSIST_RESTORE_STATES)
  };
};

// Simulator status() returns {status:{name, device_status, mma2_status}}.
// Replicator status() returns the device status directly, NOT {status}.
// Neither response proves overall Docker service or MMA2 supervisor health.
const mapDiagnostics = ({memory, replicator} = {}) => {
  const view = {
    runtime: {mma2: UNKNOWN, simulator: UNKNOWN, replicator: UNKNOWN},
    devices: {
      memory: {mma2: UNKNOWN, simulator: UNKNOWN, persistence: null},
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
      view.devices.memory.persistence = mapPersistenceHealth(result.persistence);
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
