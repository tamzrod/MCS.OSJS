// OS.js server configuration for the neutral desktop shell (SHELL-001/002).
const path = require('path');
const root = path.resolve(__dirname, '../../');

module.exports = {
  root,
  // OSJS-003: deterministic unoccupied management/UI port for the shell.
  // `PORT` env var overrides when a deployment needs a different host port.
  port: Number(process.env.PORT) || 18209,
  public: path.resolve(root, 'dist')
};
