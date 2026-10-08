const crypto = require('crypto');
const fs = require('fs');
const net = require('net');
const path = require('path');

const AREAS = [
  {fc: 'fc1', area: 'coils', code: 1, readFC: 1, bits: true},
  {fc: 'fc2', area: 'discrete_inputs', code: 2, readFC: 2, bits: true},
  {fc: 'fc3', area: 'holding_registers', code: 3, readFC: 3, bits: false},
  {fc: 'fc4', area: 'input_registers', code: 4, readFC: 4, bits: false}
];

const delay = ms => new Promise(resolve => setTimeout(resolve, ms));
const num = value => Number(value) || 0;
const enabled = device => Boolean(device?.mma2?.persistence?.enabled === true);
const hasArea = area => area && num(area.count) > 0;

const snapshotDir = (root, port, unitID, area) =>
  path.join(root, 'persistence', 'snapshots', `port-${port}`, `unit-${unitID}`, `area-${area}`);

const writeAtomic = (file, data) => {
  fs.mkdirSync(path.dirname(file), {recursive: true});
  const tmp = `${file}.tmp`;
  fs.writeFileSync(tmp, data);
  fs.renameSync(tmp, file);
};

const snapshotSize = (bits, count) => bits ? Math.ceil(count / 8) : count * 2;

const manifestFor = (device, spec, payload) => ({
  format_version: 1,
  port: num(device.mma2.port),
  unit_id: num(device.mma2.unit_id),
  area: spec.area,
  start: num(device.mma2[spec.fc].start),
  count: num(device.mma2[spec.fc].count),
  payload_bytes: payload.length,
  sha256: crypto.createHash('sha256').update(payload).digest('hex')
});

const validateManifest = (device, spec, payload, manifest) => {
  const area = device.mma2[spec.fc];
  const expected = snapshotSize(spec.bits, num(area.count));
  const checks = [
    [manifest?.format_version === 1, 'format version'],
    [num(manifest?.port) === num(device.mma2.port), 'port'],
    [num(manifest?.unit_id) === num(device.mma2.unit_id), 'unit id'],
    [manifest?.area === spec.area, 'area'],
    [num(manifest?.start) === num(area.start), 'start'],
    [num(manifest?.count) === num(area.count), 'count'],
    [num(manifest?.payload_bytes) === expected, 'payload size'],
    [payload.length === expected, 'raw snapshot size'],
    [manifest?.sha256 === crypto.createHash('sha256').update(payload).digest('hex'), 'sha256']
  ];
  const failed = checks.find(([ok]) => !ok);
  if (failed) throw new Error(`${device.name}: persistence snapshot ${spec.area} ${failed[1]} mismatch`);
};

const transact = (port, request, responseReader, timeout = 5000) => new Promise((resolve, reject) => {
  const socket = net.createConnection({host: '127.0.0.1', port: num(port)});
  const timer = setTimeout(() => socket.destroy(new Error('persistence socket timed out')), timeout);
  socket.on('connect', () => socket.write(request));
  socket.on('data', chunk => {
    try {
      const result = responseReader(chunk);
      clearTimeout(timer);
      socket.destroy();
      resolve(result);
    } catch (error) {
      clearTimeout(timer);
      socket.destroy();
      reject(error);
    }
  });
  socket.on('error', error => {
    clearTimeout(timer);
    reject(error);
  });
});

const readModbus = async (device, spec) => {
  const area = device.mma2[spec.fc];
  if (!hasArea(area)) return null;
  const request = Buffer.alloc(12);
  request.writeUInt16BE(1, 0);
  request.writeUInt16BE(0, 2);
  request.writeUInt16BE(6, 4);
  request[6] = num(device.mma2.unit_id);
  request[7] = spec.readFC;
  request.writeUInt16BE(num(area.start), 8);
  request.writeUInt16BE(num(area.count), 10);
  return transact(device.mma2.port, request, chunk => {
    if (chunk.length < 9) throw new Error(`${device.name}: short Modbus response while reading ${spec.area}`);
    const pdu = chunk.subarray(7);
    if (pdu[0] & 0x80) throw new Error(`${device.name}: Modbus ${spec.area} read rejected with 0x${(pdu[1] || 0).toString(16).padStart(2, '0')}`);
    if (pdu[0] !== spec.readFC) throw new Error(`${device.name}: unexpected Modbus function while reading ${spec.area}`);
    const length = pdu[1];
    const payload = Buffer.from(pdu.subarray(2, 2 + length));
    const expected = snapshotSize(spec.bits, num(area.count));
    if (payload.length !== expected) throw new Error(`${device.name}: ${spec.area} payload length ${payload.length}, expected ${expected}`);
    return payload;
  });
};

const rawIngest = async (device, areaCode, start, count, payload) => {
  const packet = Buffer.alloc(10 + payload.length);
  packet[0] = 0x52; packet[1] = 0x49; packet[2] = 0x01; packet[3] = areaCode;
  packet.writeUInt16BE(num(device.mma2.unit_id), 4);
  packet.writeUInt16BE(num(start), 6);
  packet.writeUInt16BE(num(count), 8);
  payload.copy(packet, 10);
  return transact(device.mma2.port, packet, chunk => {
    const code = chunk[0];
    if (code !== 0x00) throw new Error(`${device.name}: Raw Ingest rejected with 0x${(code || 0).toString(16).padStart(2, '0')}`);
    return code;
  });
};

const snapshotsComplete = (root, device) => {
  try {
    for (const spec of AREAS) {
      const area = device.mma2[spec.fc];
      if (!hasArea(area)) continue;
      const dir = snapshotDir(root, device.mma2.port, device.mma2.unit_id, spec.area);
      const payload = Buffer.from(fs.readFileSync(path.join(dir, 'snapshot.bin')));
      const manifest = JSON.parse(fs.readFileSync(path.join(dir, 'manifest.json'), 'utf8'));
      validateManifest(device, spec, payload, manifest);
    }
    return AREAS.some(spec => hasArea(device.mma2[spec.fc]));
  } catch (_) {
    return false;
  }
};

const unsealForBootstrap = async device => {
  const sealing = device?.mma2?.state_sealing;
  if (!sealing || sealing.enabled === false) return;
  if (String(sealing.area || '').toLowerCase() !== 'coil') {
    throw new Error(`${device.name}: State Sealing bootstrap requires coil area`);
  }
  await rawIngest(device, 1, num(sealing.address), 1, Buffer.from([0x01]));
};

const captureInitialSnapshots = async (root, previousDoc, editedDoc) => {
  const before = new Map((previousDoc.devices || []).map(device => [device.name, device]));
  for (const after of editedDoc.devices || []) {
    if (!enabled(after)) continue;

    const old = before.get(after.name);
    if (!old) throw new Error(`${after.name}: persistence Save & Apply requires a running memory to snapshot`);
    if (num(old.mma2.port) !== num(after.mma2.port) || num(old.mma2.unit_id) !== num(after.mma2.unit_id)) {
      throw new Error(`${after.name}: snapshot current persistence state before changing Port/Unit ID`);
    }
    for (const spec of AREAS) {
      const oldArea = old.mma2[spec.fc] || {start: 0, count: 0};
      const newArea = after.mma2[spec.fc] || {start: 0, count: 0};
      if (num(oldArea.start) !== num(newArea.start) || num(oldArea.count) !== num(newArea.count)) {
        throw new Error(`${after.name}: snapshot current persistence state before changing memory ranges`);
      }
    }

    // Save & Apply always starts by taking a fresh full persistence snapshot.
    // If a previous failed apply left the running memory sealed, Raw Ingest is
    // the authoritative bypass used only to make the current state readable
    // before capture. Existing snapshot files are deliberately overwritten;
    // runtime RBE saves may update them again afterward.
    if (enabled(old)) await unsealForBootstrap(old);

    for (const spec of AREAS) {
      const newArea = after.mma2[spec.fc] || {start: 0, count: 0};
      if (!hasArea(newArea)) continue;
      const payload = await readModbus(old, spec);
      const dir = snapshotDir(root, after.mma2.port, after.mma2.unit_id, spec.area);
      writeAtomic(path.join(dir, 'snapshot.bin'), payload);
      writeAtomic(path.join(dir, 'manifest.json'), Buffer.from(JSON.stringify(manifestFor(after, spec, payload), null, 2)));
    }

    if (!snapshotsComplete(root, after)) {
      throw new Error(`${after.name}: persistence snapshot set is incomplete after Save & Apply capture`);
    }
  }
};

const waitRestartAcknowledged = async (ackPath, expectedSHA, timeout = 20000) => {
  const deadline = Date.now() + timeout;
  for (;;) {
    try {
      if (fs.readFileSync(ackPath, 'utf8') === expectedSHA) return;
    } catch (error) {
      if (error.code !== 'ENOENT') throw error;
    }
    if (Date.now() >= deadline) throw new Error('MMA2 restart request was not acknowledged within timeout');
    await delay(25);
  }
};

const waitPort = async (port, timeout = 20000) => {
  const deadline = Date.now() + timeout;
  for (;;) {
    const ready = await new Promise(resolve => {
      const socket = net.createConnection({host: '127.0.0.1', port: num(port)});
      socket.setTimeout(250);
      socket.on('connect', () => { socket.destroy(); resolve(true); });
      socket.on('timeout', () => { socket.destroy(); resolve(false); });
      socket.on('error', () => resolve(false));
    });
    if (ready) return;
    if (Date.now() >= deadline) throw new Error(`MMA2 restart not ready on port ${port}`);
    await delay(25);
  }
};

const restoreAndUnseal = async (root, device) => {
  if (!enabled(device)) return;
  const sealing = device.mma2.state_sealing;
  if (!sealing || sealing.enabled === false || String(sealing.area || '').toLowerCase() !== 'coil') {
    throw new Error(`${device.name}: persistence requires enabled coil State Sealing`);
  }

  for (const spec of AREAS) {
    const area = device.mma2[spec.fc];
    if (!hasArea(area)) continue;
    const dir = snapshotDir(root, device.mma2.port, device.mma2.unit_id, spec.area);
    const payload = Buffer.from(fs.readFileSync(path.join(dir, 'snapshot.bin')));
    const manifest = JSON.parse(fs.readFileSync(path.join(dir, 'manifest.json'), 'utf8'));
    validateManifest(device, spec, payload, manifest);
    let writePayload = payload;
    if (spec.area === 'coils') {
      const offset = num(sealing.address) - num(area.start);
      if (offset >= 0 && offset < num(area.count)) {
        writePayload = Buffer.from(payload);
        writePayload[Math.floor(offset / 8)] &= ~(1 << (offset % 8));
      }
    }
    await rawIngest(device, spec.code, area.start, area.count, writePayload);
  }

  await rawIngest(device, 1, num(sealing.address), 1, Buffer.from([0x01]));

  // Final proof that the memory is no longer sealed: one normal Modbus read
  // from the first configured area must succeed after the commit write.
  const proof = AREAS.find(spec => hasArea(device.mma2[spec.fc]));
  if (proof) await readModbus(device, proof);
};

module.exports = {
  AREAS,
  snapshotsComplete,
  captureInitialSnapshots,
  restoreAndUnseal,
  waitRestartAcknowledged,
  waitPort
};
