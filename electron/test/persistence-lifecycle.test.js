const test = require('node:test');
const assert = require('node:assert/strict');
const fs = require('fs');
const os = require('os');
const path = require('path');
const net = require('net');
const persistence = require('../persistence');

const tempRoot = () => fs.mkdtempSync(path.join(os.tmpdir(), 'mcs-persist-electron-'));

const fakeMMA2 = async ({register = 0x1234, rejectUnseal = false} = {}) => {
  const packets = [];
  let sealed = false;
  const server = net.createServer(socket => {
    const chunks = [];
    socket.on('data', chunk => {
      chunks.push(chunk);
      const req = Buffer.concat(chunks);
      if (req.length < 2) return;

      if (req[0] === 0x52 && req[1] === 0x49) {
        if (req.length < 10) return;
        const area = req[3];
        const address = req.readUInt16BE(6);
        const count = req.readUInt16BE(8);
        const payloadLength = (area === 1 || area === 2) ? Math.ceil(count / 8) : count * 2;
        if (req.length < 10 + payloadLength) return;
        const payload = Buffer.from(req.subarray(10, 10 + payloadLength));
        packets.push({area, address, count, payload});
        if (area === 1 && count === 1 && payload[0] === 0x01) {
          if (rejectUnseal) return socket.end(Buffer.from([0x21]));
          sealed = false;
        }
        return socket.end(Buffer.from([0x00]));
      }

      if (req.length < 12) return;
      const fc = req[7];
      const count = req.readUInt16BE(10);
      if (sealed) {
        const response = Buffer.alloc(9);
        response.writeUInt16BE(req.readUInt16BE(0), 0);
        response.writeUInt16BE(3, 4);
        response[6] = req[6];
        response[7] = fc | 0x80;
        response[8] = 0x06;
        return socket.end(response);
      }
      let payload;
      if (fc === 3 || fc === 4) {
        payload = Buffer.alloc(count * 2);
        for (let i = 0; i < count; i++) payload.writeUInt16BE(register, i * 2);
      } else {
        payload = Buffer.alloc(Math.ceil(count / 8));
      }
      const response = Buffer.alloc(9 + payload.length);
      response.writeUInt16BE(req.readUInt16BE(0), 0);
      response.writeUInt16BE(3 + payload.length, 4);
      response[6] = req[6];
      response[7] = fc;
      response[8] = payload.length;
      payload.copy(response, 9);
      socket.end(response);
    });
  });
  await new Promise((resolve, reject) => {
    server.once('error', reject);
    server.listen(0, '127.0.0.1', resolve);
  });
  const port = server.address().port;
  return {
    port,
    packets,
    seal: () => { sealed = true; },
    close: () => new Promise(resolve => server.close(resolve))
  };
};

const device = (port, persistenceEnabled) => ({
  name: 'PLC',
  enabled: true,
  mma2: {
    port,
    unit_id: 1,
    fc1: {start: 0, count: 8},
    fc2: {start: 0, count: 0},
    fc3: {start: 0, count: 1},
    fc4: {start: 0, count: 0},
    persistence: persistenceEnabled ? {enabled: true} : {enabled: false},
    state_sealing: persistenceEnabled
      ? {enabled: true, area: 'coil', address: 0, exception: 6}
      : {enabled: false, area: 'coil', address: 0, exception: 6}
  }
});

test('first enable captures live snapshots then restores and final-unseals', async t => {
  const fake = await fakeMMA2();
  t.after(() => fake.close());
  const root = tempRoot();
  t.after(() => fs.rmSync(root, {recursive: true, force: true}));

  const before = device(fake.port, false);
  const after = device(fake.port, true);
  await persistence.captureInitialSnapshots(root, {devices: [before]}, {devices: [after]});

  assert.equal(fs.existsSync(path.join(root, 'persistence', 'snapshots', `port-${fake.port}`, 'unit-1', 'area-holding_registers', 'snapshot.bin')), true);

  fake.seal();
  await persistence.restoreAndUnseal(root, after);

  const last = fake.packets.at(-1);
  assert.equal(last.area, 1);
  assert.equal(last.address, 0);
  assert.equal(last.count, 1);
  assert.deepEqual([...last.payload], [1]);
});

test('capture failure is surfaced before a sealed restart can be reported successful', async () => {
  const root = tempRoot();
  const before = device(65534, false);
  const after = device(65534, true);
  await assert.rejects(
    persistence.captureInitialSnapshots(root, {devices: [before]}, {devices: [after]}),
    /connect|ECONNREFUSED|timed out/i
  );
  fs.rmSync(root, {recursive: true, force: true});
});

test('final unseal rejection fails closed', async t => {
  const fake = await fakeMMA2({rejectUnseal: true});
  t.after(() => fake.close());
  const root = tempRoot();
  t.after(() => fs.rmSync(root, {recursive: true, force: true}));

  const before = device(fake.port, false);
  const after = device(fake.port, true);
  await persistence.captureInitialSnapshots(root, {devices: [before]}, {devices: [after]});
  fake.seal();
  await assert.rejects(persistence.restoreAndUnseal(root, after), /Raw Ingest rejected with 0x21/);
});


test('every persistence Save & Apply capture overwrites the previous snapshot', async t => {
  const fake = await fakeMMA2({register: 0x1234});
  t.after(() => fake.close());
  const root = tempRoot();
  t.after(() => fs.rmSync(root, {recursive: true, force: true}));

  const current = device(fake.port, true);
  await persistence.captureInitialSnapshots(root, {devices: [current]}, {devices: [current]});

  const snapshot = path.join(root, 'persistence', 'snapshots', `port-${fake.port}`, 'unit-1', 'area-holding_registers', 'snapshot.bin');
  fs.writeFileSync(snapshot, Buffer.from([0x00, 0x00]));

  await persistence.captureInitialSnapshots(root, {devices: [current]}, {devices: [current]});
  assert.deepEqual([...fs.readFileSync(snapshot)], [0x12, 0x34]);
});

test('Save & Apply snapshot capture recovers a previously sealed persistence runtime', async t => {
  const fake = await fakeMMA2({register: 0x4321});
  t.after(() => fake.close());
  const root = tempRoot();
  t.after(() => fs.rmSync(root, {recursive: true, force: true}));

  const current = device(fake.port, true);
  fake.seal();

  await persistence.captureInitialSnapshots(root, {devices: [current]}, {devices: [current]});

  const snapshot = path.join(root, 'persistence', 'snapshots', `port-${fake.port}`, 'unit-1', 'area-holding_registers', 'snapshot.bin');
  assert.deepEqual([...fs.readFileSync(snapshot)], [0x43, 0x21]);
  const unseal = fake.packets.find(packet => packet.area === 1 && packet.count === 1 && packet.payload[0] === 0x01);
  assert.ok(unseal, 'sealed runtime must be unsealed through Raw Ingest before fresh snapshot capture');
});
