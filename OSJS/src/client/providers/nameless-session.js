// DRES-002 — server-backed desktop session snapshot capture.

import {SessionSnapshotCapture, validateSnapshot} from './session-snapshot.js';

const SETTINGS_NS = 'nameless/session';
const SETTINGS_KEY = 'snapshot';
export default class NamelessSessionServiceProvider {
  constructor(core) {
    this.core = core;
    this.capture = null;
    this.onApplicationLaunched = null;
    this.onPageHide = null;
  }

  provides() {
    return ['nameless/session'];
  }

  init() {
    const settings = this.core.make('osjs/settings');
    this.capture = new SessionSnapshotCapture(value => {
      settings.set(SETTINGS_NS, SETTINGS_KEY, value);
      return settings.save();
    });
    this.core.singleton('nameless/session', () => ({
      flush: () => this.capture.flush(),
      apply: () => this.apply(),
      reset: () => this.reset()
    }));
    return Promise.resolve();
  }

  start() {
    this.onApplicationLaunched = (name, process) => this.capture.track(name, process);
    this.core.on('osjs/application:launched', this.onApplicationLaunched);
    this.onPageHide = () => this.capture.flush();
    window.addEventListener('pagehide', this.onPageHide);
    return Promise.resolve();
  }

  destroy() {
    if (this.onApplicationLaunched && typeof this.core.off === 'function') {
      this.core.off('osjs/application:launched', this.onApplicationLaunched);
    }
    if (this.onPageHide) window.removeEventListener('pagehide', this.onPageHide);
    if (this.capture) this.capture.destroy();
  }

  read() {
    return validateSnapshot(
      this.core.make('osjs/settings').get(SETTINGS_NS, SETTINGS_KEY, null)
    );
  }

  whenRendered(window, attempts = 20) {
    return new Promise(resolve => {
      const inspect = remaining => {
        if (window.$element || remaining <= 0) {
          resolve(window.$element ? window : null);
        } else {
          setTimeout(() => inspect(remaining - 1), 50);
        }
      };
      inspect(attempts);
    });
  }

  async applyEntry(packages, entry) {
    try {
      const process = await packages.launch(entry.package);
      const window = process && Array.isArray(process.windows) && process.windows[0];
      if (!window) return null;
      const rendered = await this.whenRendered(window);
      if (!rendered) return null;
      window.setDimension(entry.dimension);
      window.setPosition(entry.position);
      if (entry.maximized && window.attributes.maximizable !== false) window.maximize();
      if (entry.minimized) window.minimize();
      return window;
    } catch (error) {
      return null;
    }
  }

  // Returns false only when there is no valid authoritative snapshot. The boot
  // seam uses that result to fall back to the existing Auto-Start behavior.
  async apply() {
    const snapshot = this.read();
    if (!snapshot) return false;
    const packages = this.core.make('osjs/packages');
    const entries = snapshot.windows.slice().sort((a, b) => a.order - b.order);
    let focus = null;
    for (const entry of entries) {
      const window = await this.applyEntry(packages, entry);
      if (window && !entry.minimized) focus = window;
    }
    if (focus && typeof focus.focus === 'function') focus.focus();
    return true;
  }

  reset() {
    // Stop pending/current-page capture first so a queued debounce cannot put
    // the snapshot back after the operator explicitly removed it.
    this.capture.destroy();
    const settings = this.core.make('osjs/settings');
    const raw = settings.get(SETTINGS_NS, null, {});
    const namespace = raw && typeof raw === 'object' && !Array.isArray(raw)
      ? {...raw}
      : {};
    delete namespace[SETTINGS_KEY];
    settings.set(SETTINGS_NS, null, namespace);
    return settings.save().then(() => true).catch(() => false);
  }
}
