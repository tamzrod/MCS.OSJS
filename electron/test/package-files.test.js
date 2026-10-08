const test = require('node:test');
const assert = require('node:assert/strict');
const path = require('path');
const fs = require('fs');

test('Electron package whitelist includes every main-process local module', () => {
  const root = path.join(__dirname, '..');
  const pkg = JSON.parse(fs.readFileSync(path.join(root, 'package.json'), 'utf8'));
  const files = new Set(pkg.build.files || []);

  const main = fs.readFileSync(path.join(root, 'main.js'), 'utf8');
  const localRequires = [...main.matchAll(/require\(['"]\.\/([^'"]+)['"]\)/g)]
    .map(match => `${match[1]}.js`);

  for (const file of localRequires) {
    assert.ok(files.has(file), `${file} is required by main.js but missing from electron-builder build.files`);
  }
});

// NPE-05: MMA2 owns persistence runtime, so no Electron-owned persistence
// runtime/restore engine may remain.
test('Electron retains no legacy persistence runtime or restore engine', () => {
  const root = path.join(__dirname, '..');
  assert.equal(fs.existsSync(path.join(root, 'persistence.js')), false, 'electron/persistence.js must be removed');
  const sources = ['main.js', 'preload.js', 'replicator-runtime.js', 'replicator-ipc.js', 'runtime-status.js', 'diagnostics.js', 'memory-settings.js',
    ...fs.readdirSync(path.join(root, 'renderer')).filter(name => name.endsWith('.js')).map(name => path.join('renderer', name))];
  for (const relative of sources) {
    const source = fs.readFileSync(path.join(root, relative), 'utf8');
    for (const banned of ['captureSnapshots', 'restoreAndUnseal', 'snapshotsComplete', "require('./persistence')", 'persistence manager watchdog']) {
      assert.ok(!source.includes(banned), `${relative} must not contain the retired persistence runtime marker ${banned}`);
    }
  }
});
