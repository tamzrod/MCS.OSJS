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
