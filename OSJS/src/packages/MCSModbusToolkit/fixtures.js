// UMIG-003: deterministic display-only fixtures. No runtime or persistence API.
// These examples must never be represented as loaded customer configuration.
const UNKNOWN = 'UNKNOWN';

const snapshot = () => ({
  runtime: {mma2: UNKNOWN, simulator: UNKNOWN, replicator: UNKNOWN},
  memory: {
    devices: [{
      name: 'Fixture Sim-PLC-1', enabled: true,
      mma2: {
        port: 5020, unit_id: 1,
        fc1: {start: 0, count: 16}, fc2: {start: 0, count: 16},
        fc3: {start: 0, count: 16}, fc4: {start: 0, count: 16}
      },
      random_runtime: {
        fc1_interval_ms: 0, fc2_interval_ms: 0,
        fc3_interval_ms: 1000, fc4_interval_ms: 0
      }
    }]
  },
  replicator: {
    devices: [{
      name: 'Fixture Rep-PLC-1', enabled: true,
      endpoint: '192.0.2.1:502', unit_id: 1,
      destination: {port: 5021, unit_id: 1, auto_port: true, auto_unit_id: true, owner: UNKNOWN, status: UNKNOWN},
      pull_blocks: [{function: 3, start: 0, count: 16, scan_rate_ms: 1000}]
    }]
  },
  comms: {network: UNKNOWN, tcp: UNKNOWN, modbus: UNKNOWN, mma2: UNKNOWN},
  diagnostics: {runtime_mode: 'UNAVAILABLE', bin_path: 'UNAVAILABLE', data_path: 'UNAVAILABLE'}
});

module.exports = {UNKNOWN, snapshot};
