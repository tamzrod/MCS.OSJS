// ICON-002 — explicit public route for the installed Nameless Classic assets.

const fs = require('fs');
const path = require('path');

const ICON_NAME = /^[A-Za-z0-9_-]+\.svg$/;

const resolveIcon = (root, name) => {
  if (typeof name !== 'string' || !ICON_NAME.test(name)) return null;
  // Use the package-owned dist directly. The discovery output is a symlink and
  // can retain an absolute host target when a local dist/ is copied into an
  // image before discovery; the package path is stable in both host and image.
  const file = path.resolve(root, 'src', 'packages', 'NamelessClassicIcons', 'dist', 'icons', name);
  return fs.existsSync(file) && fs.statSync(file).isFile() ? file : null;
};

class ClassicIconsRouteProvider {
  constructor(core) {
    this.core = core;
  }

  provides() { return []; }

  init() {
    const express = this.core.make('osjs/express');
    const root = this.core.config('root');
    express.route('get', '/icons/NamelessClassicIcons/icons/:name', (req, res) => {
      const file = resolveIcon(root, req.params.name);
      if (!file) return res.status(404).end();
      return res.type('image/svg+xml').sendFile(file);
    });
    return Promise.resolve(true);
  }

  start() { return Promise.resolve(true); }
  destroy() {}
}

module.exports = {ClassicIconsRouteProvider, resolveIcon};
