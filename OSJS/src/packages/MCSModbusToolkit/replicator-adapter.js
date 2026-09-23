'use strict';

// Toolkit-owned Replicator document and status rules. Go remains the final
// authority for destination ownership, source validation, and apply commits.
const copy = value => JSON.parse(JSON.stringify(value));
const integer = (value, min, max) => Number.isSafeInteger(value) && value >= min && value <= max;
const blankBlock = () => ({function: 3, start: 0, count: 16, scan_rate_ms: 1000});
const normalizeDocument = value => ({devices: Array.isArray(value && value.devices) ? value.devices.map(entry => {
  const device = copy(entry);
  // Only historical persisted pull_block is migrated, not fixture defaults.
  if (!Array.isArray(device.pull_blocks) || !device.pull_blocks.length) {
    device.pull_blocks = device.pull_block ? [copy(device.pull_block)] : [];
  }
  delete device.pull_block;
  return device;
}) : []});
const blankDevice = (sequence, suggestion) => ({
  name: `Rep-PLC-${sequence}`, enabled: true, endpoint: '127.0.0.1:5020', unit_id: 1,
  pull_blocks: [blankBlock()], destination: {
    port: suggestion.port, unit_id: suggestion.unit_id, auto_port: true, auto_unit_id: true
  }
});
const endpointValid = value => {
  if (typeof value !== 'string') return false;
  const match = /^(\[[^\]\s]+\]|[^:\s]+):([0-9]+)$/.exec(value.trim());
  return Boolean(match && integer(Number(match[2]), 1, 65535));
};
const validateDevice = device => {
  if (!device || typeof device.name !== 'string' || !device.name.trim()) return 'Name is required.';
  if (!endpointValid(device.endpoint)) return 'Source endpoint must be host:port, with a valid port.';
  if (!integer(device.unit_id, 0, 255)) return 'Source Unit ID must be 0–255.';
  if (!device.destination) return 'Destination is required.';
  const dest = device.destination;
  if (!integer(dest.port, dest.auto_port ? 0 : 1, 65535)) return 'Destination Port must be 1–65535 (or 0 for Auto).';
  if (!integer(dest.unit_id, 0, 255)) return 'Destination Unit ID must be 0–255.';
  if (typeof dest.auto_port !== 'boolean' || typeof dest.auto_unit_id !== 'boolean') return 'Destination allocation flags are invalid.';
  if (!Array.isArray(device.pull_blocks) || !device.pull_blocks.length) return 'At least one Pull Block is required.';
  const groups = new Map();
  for (let i = 0; i < device.pull_blocks.length; i++) {
    const block = device.pull_blocks[i];
    const prefix = `Pull Block ${i + 1}`;
    if (!block || !integer(block.function, 1, 4)) return `${prefix}: Function must be FC1–FC4.`;
    if (!integer(block.start, 0, 65535) || !integer(block.count, 1, 65535) || block.start + block.count > 65536) {
      return `${prefix}: invalid Start/Count or address range.`;
    }
    if (!integer(block.scan_rate_ms, 1, 4294967295)) return `${prefix}: Scan Rate must be a positive uint32.`;
    if (!groups.has(block.function)) groups.set(block.function, []);
    groups.get(block.function).push(block);
  }
  for (const [fc, blocks] of groups) {
    blocks.sort((a, b) => a.start - b.start);
    let end = blocks[0].start + blocks[0].count;
    for (const block of blocks.slice(1)) {
      if (block.start > end) return `FC${fc} Pull Blocks must overlap or touch; gaps cannot be represented by MMA2.`;
      end = Math.max(end, block.start + block.count);
    }
  }
  return null;
};
const validateDocument = document => {
  if (!document || !Array.isArray(document.devices)) return 'Replicator document unavailable.';
  const names = new Set();
  for (const device of document.devices) {
    const error = validateDevice(device);
    if (error) return error;
    const key = device.name.trim().toLowerCase();
    if (names.has(key)) return `Duplicate Replicator name: ${device.name}`;
    names.add(key);
  }
  return null;
};
const displayStatus = (status, selectedName, error) => {
  const unavailable = Boolean(error);
  const valid = !unavailable && status && status.name === selectedName &&
    typeof status.running === 'boolean' && typeof status.source_status === 'string';
  if (!valid) return {replicator: unavailable ? 'UNAVAILABLE' : 'UNKNOWN', source: unavailable ? 'UNAVAILABLE' : 'UNKNOWN',
    blocks: [], error: error ? (error.message || String(error)) : ''};
  return {replicator: status.running ? 'RUNNING' : 'STOPPED', source: status.source_status,
    blocks: Array.isArray(status.blocks) ? status.blocks : [], error: status.last_error || ''};
};
module.exports = {copy, blankBlock, blankDevice, normalizeDocument, validateDevice, validateDocument, displayStatus};
