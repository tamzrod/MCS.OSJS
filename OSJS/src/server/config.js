// OS.js server configuration for the neutral desktop shell (SHELL-001/002).
const path = require('path');
const root = path.resolve(__dirname, '../../');

// `OSJS_DATA_DIR` relocates the cwd-relative persistence files (VFS home
// roots + connect-loki session store) to a stable data directory, which the
// container deployment mounts as a persistent volume (see deploy/docker-compose.yml).
// When unset, runtime behavior matches the source checkout (data next to the app).
const dataDir = process.env.OSJS_DATA_DIR || process.cwd();

module.exports = {
  root,
  // OSJS-003: deterministic unoccupied management/UI port for the shell.

  // `PORT` env var overrides when a deployment needs a different host port。
  port: Number(process.env.PORT) || 18209,
  public: path.resolve(root, 'dist'),
  // VFS default root is `{cwd}/vfs`; relocate alongside the session store when
  // a data dir is configured so both persist together.
 vfs: {
   root: path.join(dataDir, 'vfs')
  },
  // connect-loki default path is `{cwd}/session-store.db`; keep it beside vfs.

 session: {
   store: {
     options: {
      autosave: true,
       path: path.join(dataDir, 'session-store.db')
     }
    }
  }
};
