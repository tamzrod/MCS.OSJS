const test = require('node:test');
const assert = require('node:assert/strict');
const fs = require('node:fs');
const path = require('node:path');
const vm = require('node:vm');
const settings = require('../memory-settings');

const fixture = () => ({rbe: {tcp: {listen: '127.0.0.1:9001'}}, custom: 'keep', listeners: [{
  id: 'local', listen: '127.0.0.1:5020', customListener: 7, memory: [{unit_id: 1,
    coils: {start: 0, count: 4}, customMemory: {value: 8},
    state_sealing: {enabled: false, area: 'coil', address: 0},
    policy: {rules: [{id: 'read', source_ip: ['::/0'], allow_fc: [1, 3]}]},
    rbe: {coils: [{id: 1, name: 'coil', start: 0, count: 1}]}
  }]
}]});

test('legacy documents inherit advanced fields and explicit edits win without mutating inputs', () => {
  const config = fixture();
  const params = {port: 5020, unit_id: 1, fc1: {start: 0, count: 4}};
  const hydrated = settings.inheritSettings(params, config);
  assert.deepEqual(hydrated.rbe, config.listeners[0].memory[0].rbe);
  assert.equal(hydrated.customMemory.value, 8);
  assert.equal(params.policy, undefined);
  const changed = settings.inheritSettings({...params, policy: {rules: []}, state_sealing: {enabled: false}, rbe: {}}, config);
  assert.deepEqual(changed.policy, {rules: []});
  assert.deepEqual(changed.rbe, {});
  const memory = settings.applySettings({unit_id: 1, coils: params.fc1}, hydrated);
  assert.equal(memory.port, undefined);
  assert.equal(memory.fc1, undefined);
  assert.deepEqual(memory.policy, hydrated.policy);
  hydrated.policy.rules[0].id = 'edited';
  assert.equal(config.listeners[0].memory[0].policy.rules[0].id, 'read');
});

test('persistence.enabled round-trips through inherit and apply without a second authority', () => {
  const config = fixture();
  config.listeners[0].memory[0].persistence = {enabled: true};
  const params = {port: 5020, unit_id: 1, fc1: {start: 0, count: 4}};
  const hydrated = settings.inheritSettings(params, config);
  assert.deepEqual(hydrated.persistence, {enabled: true});
  // An explicit draft edit wins.
  const edited = settings.applySettings({unit_id: 1, coils: params.fc1}, {...hydrated, persistence: {enabled: false}});
  assert.deepEqual(edited.persistence, {enabled: false});
  // A device with no persisted persistence config stays absent.
  const absent = settings.inheritSettings(params, fixture());
  assert.equal(absent.persistence, undefined);
});

test('shared settings preserve memories and unknown fields; listener binding survives composition', () => {
  const previous = fixture();
  const updated = settings.updateSharedSettings(previous, {debug: true, rbe: {tcp: {listen: '[::1]:9100'}}});
  assert.deepEqual(updated.listeners, previous.listeners);
  assert.equal(updated.custom, 'keep');
  assert.equal(settings.sharedSettings(updated).rbe.tcp.listen, '[::1]:9100');
  assert.throws(() => settings.updateSharedSettings(previous, {listeners: []}), /Unsupported/);
  const candidate = {listeners: [{id: 'new', listen: '0.0.0.0:5020', memory: []}]};
  settings.preserveListeners(candidate, previous);
  assert.equal(candidate.listeners[0].listen, '127.0.0.1:5020');
  assert.equal(candidate.listeners[0].customListener, 7);
  assert.equal(settings.updateSharedSettings(previous, {rbe: null}).rbe, undefined);
});

test('Electron composition validates before writes and preserves foreign settings on success', () => {
  const writes = [];
  const config = fixture();
  config.listeners[0].memory.push({unit_id: 9, holding_registers: {start: 0, count: 2}, custom: 'foreign'});
  let reject = true;
  let candidate;
  const fakeFS = {
    readFileSync: file => JSON.stringify(file.endsWith('owners.yaml')
      ? {reservations: [{port: 5020, unit_id: 1, owner: 'simulator'}, {port: 5020, unit_id: 9, owner: 'other'}]}
      : config),
    mkdirSync: () => {}, rmSync: () => {},
    writeFileSync: (file, body) => writes.push({file, body}), renameSync: () => {}
  };
  const mockRequire = name => {
    if (name === 'electron') return {app: {isPackaged: false, getPath: () => '/isolated'}, ipcMain: {}};
    if (name === 'fs') return fakeFS;
    if (name === 'js-yaml') return {load: JSON.parse, dump: JSON.stringify};
    if (name === 'child_process') return {execFileSync: (file, args, options) => {
      assert.deepEqual(Array.from(args), ['--validate-stdin']);
      candidate = JSON.parse(options.input);
      if (reject) throw new Error('invalid fixture');
    }};
    if (name === './memory-settings') return settings;
    if (name.startsWith('./')) return {};
    return require(name);
  };
  const context = vm.createContext({require: mockRequire, __dirname: path.join(__dirname, '..'), process});
  const source = fs.readFileSync(path.join(__dirname, '../main.js'), 'utf8');
  vm.runInContext(source.slice(0, source.indexOf('const portStatus =')), context);
  const compose = "composeAll({devices:[{name:'PLC',enabled:true,mma2:{port:5020,unit_id:1,fc1:{start:0,count:4}}}]},{devices:[]})";
  assert.throws(() => vm.runInContext(compose, context), /validation failed/);
  assert.equal(writes.length, 0);
  reject = false;
  vm.runInContext(compose, context);
  assert.equal(writes.length, 3);
  const memories = candidate.listeners[0].memory;
  assert.equal(memories.find(memory => memory.unit_id === 9).custom, 'foreign');
  assert.deepEqual(memories.find(memory => memory.unit_id === 1).policy, config.listeners[0].memory[0].policy);
  assert.equal(candidate.rbe.tcp.listen, '127.0.0.1:9001');
});

// NPE-03 harness: load the main-process compose/validate/save region in an
// isolated VM. The './persistence' module is intentionally unavailable so any
// accidental Save & Apply dependency on the legacy external persistence
// lifecycle fails loudly.
const loadMain = (config, owners) => {
  const writes = [];
  const candidate = {value: null};
  const fakeFS = {
    readFileSync: file => JSON.stringify(file.endsWith('owners.yaml') ? owners : config),
    mkdirSync: () => {}, rmSync: () => {},
    writeFileSync: (file, body) => writes.push({file, body}), renameSync: () => {}
  };
  const mockRequire = name => {
    if (name === 'electron') return {app: {isPackaged: false, getPath: () => '/isolated'}, ipcMain: {}};
    if (name === 'fs') return fakeFS;
    if (name === 'js-yaml') return {load: JSON.parse, dump: JSON.stringify};
    if (name === 'child_process') return {execFileSync: (file, args, options) => {
      assert.deepEqual(Array.from(args), ['--validate-stdin']);
      candidate.value = JSON.parse(options.input);
    }};
    if (name === './memory-settings') return settings;
    if (name === './persistence') throw new Error('Save & Apply must not require the legacy persistence module');
    if (name.startsWith('./')) return {};
    return require(name);
  };
  const context = vm.createContext({require: mockRequire, __dirname: path.join(__dirname, '..'), process});
  const source = fs.readFileSync(path.join(__dirname, '../main.js'), 'utf8');
  vm.runInContext(source.slice(0, source.indexOf('const portStatus =')), context);
  return {context, writes, candidate};
};

const composeSimulator = (h, mma2) => {
  const sim = {devices: [{name: 'PLC', enabled: true, mma2}]};
  vm.runInContext(`composeAll(${JSON.stringify(sim)},{devices:[]})`, h.context);
  return h.candidate.value.listeners
    .find(listener => Number(String(listener.listen).match(/:(\d+)$/)[1]) === Number(mma2.port))
    .memory.find(memory => Number(memory.unit_id) === Number(mma2.unit_id));
};

const persistenceOwners = () => ({reservations: [{port: 5020, unit_id: 1, owner: 'simulator'}]});

test('Save & Apply emits native persistence with no forced directory or ranges', () => {
  const h = loadMain(fixture(), persistenceOwners());
  const memory = composeSimulator(h, {port: 5020, unit_id: 1, fc1: {start: 0, count: 4}, persistence: {enabled: true}});
  // Enabled with default settings carries no directory and no ranges override.
  assert.deepEqual(memory.persistence, {enabled: true});
});

test('custom persistence directory and ranges round-trip through Save & Apply', () => {
  const h = loadMain(fixture(), persistenceOwners());
  const memory = composeSimulator(h, {
    port: 5020, unit_id: 1, fc1: {start: 0, count: 4},
    persistence: {enabled: true, directory: '/custom/mma2/unit1', ranges: {coils: [{start: 1, count: 2}]}}
  });
  assert.deepEqual(memory.persistence, {enabled: true, directory: '/custom/mma2/unit1', ranges: {coils: [{start: 1, count: 2}]}});
});

test('disabling persistence changes only the persistence block', () => {
  const h = loadMain(fixture(), persistenceOwners());
  const memory = composeSimulator(h, {
    port: 5020, unit_id: 1, fc1: {start: 0, count: 4},
    persistence: {enabled: false},
    state_sealing: {enabled: false, area: 'coil', address: 0},
    rbe: {coils: [{id: 1, name: 'coil', start: 0, count: 1}]}
  });
  assert.deepEqual(memory.persistence, {enabled: false});
  // Unrelated memory settings survive composition untouched.
  assert.deepEqual(memory.state_sealing, {enabled: false, area: 'coil', address: 0});
  assert.deepEqual(memory.rbe, {coils: [{id: 1, name: 'coil', start: 0, count: 1}]});
});

test('Save & Apply preserves unrelated and foreign memory settings', () => {
  const config = fixture();
  config.listeners[0].memory.push({unit_id: 9, holding_registers: {start: 0, count: 2}, custom: 'foreign'});
  const owners = {reservations: [{port: 5020, unit_id: 1, owner: 'simulator'}, {port: 5020, unit_id: 9, owner: 'other'}]};
  const h = loadMain(config, owners);
  const memory = composeSimulator(h, {port: 5020, unit_id: 1, fc1: {start: 0, count: 4}, persistence: {enabled: true}});
  assert.deepEqual(memory.persistence, {enabled: true});
  const candidate = h.candidate.value;
  assert.equal(candidate.listeners[0].memory.find(entry => entry.unit_id === 9).custom, 'foreign');
  assert.deepEqual(candidate.listeners[0].memory.find(entry => entry.unit_id === 1).policy, config.listeners[0].memory[0].policy);
});

test('Save & Apply contains no external persistence restore/unseal lifecycle', () => {
  const source = fs.readFileSync(path.join(__dirname, '../main.js'), 'utf8');
  const block = source.slice(source.indexOf('const applySimulator'), source.indexOf('const portStatus'));
  for (const banned of ['captureSnapshots', 'restoreAndUnseal', 'snapshotsComplete', 'persistence.', 'watchdog', 'unseal', 'Unlock']) {
    assert.ok(!block.includes(banned), `Save & Apply must not contain ${banned}`);
  }
  assert.ok(!source.includes("require('./persistence')"), 'main process must not require the legacy persistence module');
});

