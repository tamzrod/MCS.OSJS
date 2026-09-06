// OSUI-005 — Auto-Start startup launcher.
//
// Applies the operator-configured `nameless/autostart` settings when the
// desktop starts: every ENABLED application in the map is launched with its
// configured startup window state (OSUI-004). Nothing is hard-coded — the
// list comes entirely from the persisted settings (home:/.osjs/settings.json
// via the server settings adapter), which the Auto-Start Start Menu utility
// writes. Auto-Start itself launches only when the operator enabled it.
//
// Window states:
//   minimized → minimize every window of the application after launch
//   maximized → maximize every window
//   last      → restore the last recorded window geometry
//               (apps[name].lastWindow, captured on window destroy below)
//
// "Last State" bookkeeping: when a window of an enabled, last-state
// application is destroyed, its geometry is recorded back into settings so
// the next desktop start opens it where the operator left it. Settings saves
// are fire-and-forget here (a lost geometry update degrades to the previous
// state, not to an error).

const SETTINGS_NS = 'nameless/autostart';

export default class NamelessAutoStartServiceProvider {
  constructor(core) {
    this.core = core;
  }

  provides() {
    return ['nameless/autostart'];
  }

  init() {
    this.core.singleton('nameless/autostart', () => ({
      read: () => this.read(),
      apply: () => this.apply()
    }));
    return Promise.resolve();
  }

  start() {
    this.core.on('osjs/application:launched', (name, proc) => this.track(name, proc));
    return Promise.resolve();
  }

  destroy() {}

  read() {
    return this.core.make('osjs/settings').get(SETTINGS_NS, 'apps', {}) || {};
  }

  // Launches every enabled application with its configured window state.
  // Called on 'osjs/core:started' — settings are already loaded (the settings
  // provider initializes before desktop start).
  apply() {
    const apps = this.read();
    const packages = this.core.make('osjs/packages');
    Object.keys(apps).forEach((name) => {
      const entry = apps[name];
      if (!entry || entry.enabled !== true) return;
      packages.launch(name)
        .then((proc) => (proc.windows || []).forEach((win) => this.whenRendered(win, entry)))
        .catch(() => {}); // uninstalled/renamed package — skip, never crash boot
    });
  }

  // Window DOM mutations (maximize/minimize/geometry) are lost when applied
  // before the window's hyperapp render created $element, so state
  // application waits (briefly) for render to land.
  whenRendered(win, entry, attempts = 20) {
    if (win.$element || attempts <= 0) {
      if (win.$element) this.applyState(win, entry);
      return;
    }
    setTimeout(() => this.whenRendered(win, entry, attempts - 1), 50);
  }

  applyState(win, entry) {
    const state = entry.state || 'last';
    if (state === 'minimized') {
      win.minimize();
      return;
    }
    // A non-maximizable window (e.g. the stock Calculator, fixed dimension)
    // ignores maximize() — the launcher must not lie about the outcome, so it
    // falls back to the window's normal geometry.
    if (state === 'maximized') {
      if (win.attributes.maximizable !== false) win.maximize();
      return;
    }
    const last = entry.lastWindow;
    if (!last) return;
    if (last.maximized && win.attributes.maximizable !== false) {
      win.maximize();
      return;
    }
    if (last.dimension) win.setDimension(last.dimension);
    if (last.position) win.setPosition(last.position);
  }

  // Records last window geometry for enabled last-state applications.
  track(name, proc) {
    const entry = this.read()[name];
    if (!entry || entry.enabled !== true || (entry.state || 'last') !== 'last') return;
    (proc.windows || []).forEach((win) => win.on('destroy', () => {
      const apps = this.read();
      const latest = apps[name];
      if (!latest || latest.enabled !== true) return;
      latest.lastWindow = {
        position: win.state.position,
        dimension: win.state.dimension,
        maximized: !!win.state.maximized
      };
      // Whole-namespace replace (see AutoStart package comment) so a previous
      // lastWindow entry doesn't linger on a partial write.
      this.core.make('osjs/settings').set(SETTINGS_NS, null, {apps}).save().catch(() => {});
    }));
  }
}
