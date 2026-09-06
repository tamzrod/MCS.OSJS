// Deterministically build every local OS.js package that ships a webpack.config.js,
// exactly once, failing loudly on any build failure. Invoked by the Dockerfile
// before `npm run package:discover` so the discovered metadata never points at
// a missing dist bundle..
'use strict';

const fs = require('fs');
const path = require('path');
const {spawnSync} = require('child_process');

const root = path.resolve(__dirname, '..');
const packagesDir = path.join(root, 'src', 'packages');
const webpackBin = path.join(root, 'node_modules', '.bin', 'webpack');

const packages = fs.readdirSync(packagesDir)
  .filter((name) => fs.statSync(path.join(packagesDir, name)).isDirectory())
  .filter((name) => fs.existsSync(path.join(packagesDir, name, 'webpack.config.js')))
  .sort();
if (packages.length === 0) {
  console.error('build-local-packages: no local packages with webpack.config.js found under ' + packagesDir);
  process.exit(1);
}

if (!fs.existsSync(webpackBin)) {
  console.error('build-local-packages: webpack binary not found at ' + webpackBin);
  process.exit(1);
}

for (const name of packages) {
  const config = path.join(packagesDir, name, 'webpack.config.js');
  const label = 'src/packages/' + name;
  console.log('build-local-packages: building ' + label);
  const res = spawnSync(webpackBin, ['--config', config], {
    cwd: root,
    env: Object.assign({}, process.env, {NODE_ENV: process.env.NODE_ENV || 'production'}),
    stdio: 'inherit'
  });
  if (res.status !== 0) {
    console.error('build-local-packages: FAILED building ' + label + ' (exit ' + res.status + ')');
    process.exit(res.status === null ? 1 : res.status);
  }
}

console.log('build-local-packages: built ' + packages.length + ' local packages exactly once: ' + packages.join(', '));