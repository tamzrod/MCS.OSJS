'use strict';

const SNAPSHOT_VERSION = 1;
const WINDOW_EVENTS = [
  'moved',
  'resized',
  'minimize',
  'maximize',
  'restore',
  'focus'
];

const integer = value => Number.isFinite(value) && Number.isInteger(value);

const validateEntry = entry => {
  if (!entry || typeof entry !== 'object' ||
      typeof entry.package !== 'string' || entry.package.length === 0 ||
      !entry.position || !entry.dimension ||
      !integer(entry.position.left) || !integer(entry.position.top) ||
      !integer(entry.dimension.width) || !integer(entry.dimension.height) ||
      entry.dimension.width <= 0 || entry.dimension.height <= 0 ||
      typeof entry.minimized !== 'boolean' ||
      typeof entry.maximized !== 'boolean' ||
      !integer(entry.order) || entry.order < 0) return null;

  return {
    package: entry.package,
    position: {left: entry.position.left, top: entry.position.top},
    dimension: {width: entry.dimension.width, height: entry.dimension.height},
    minimized: entry.minimized,
    maximized: entry.maximized,
    order: entry.order
  };
};

const validateSnapshot = value => {
  if (!value || typeof value !== 'object' || value.version !== SNAPSHOT_VERSION ||
      !Array.isArray(value.windows)) return null;
  return {
    version: SNAPSHOT_VERSION,
    windows: value.windows.map(validateEntry).filter(Boolean)
  };
};

const serializeWindow = (record, order) => {
  const state = record.window && record.window.state;
  const position = state && state.position;
  const dimension = state && state.dimension;
  if (!state || !position || !dimension ||
      !integer(position.left) || !integer(position.top) ||
      !integer(dimension.width) || !integer(dimension.height) ||
      dimension.width <= 0 || dimension.height <= 0) {
    return null;
  }

  return {
    package: record.packageName,
    position: {left: position.left, top: position.top},
    dimension: {width: dimension.width, height: dimension.height},
    minimized: state.minimized === true,
    maximized: state.maximized === true,
    order
  };
};

class SessionSnapshotCapture {
  constructor(saveSnapshot, options = {}) {
    this.saveSnapshot = saveSnapshot;
    this.delay = options.delay === undefined ? 200 : options.delay;
    // Browser timer functions require their Window receiver in some engines.
    // Wrap the globals so invoking them through this capture object cannot
    // accidentally supply SessionSnapshotCapture as `this` (DRES-LIVE-001).
    this.setTimer = options.setTimer || ((fn, delay) => setTimeout(fn, delay));
    this.clearTimer = options.clearTimer || (timer => clearTimeout(timer));
    this.records = [];
    this.nextOrder = 0;
    this.timer = null;
    this.destroyed = false;
  }

  track(packageName, process) {
    const window = process && Array.isArray(process.windows) && process.windows[0];
    if (typeof packageName !== 'string' || packageName.length === 0 ||
        !window || typeof window.on !== 'function') return false;

    const record = {packageName, window, order: this.nextOrder++, listeners: []};
    this.records.push(record);
    WINDOW_EVENTS.forEach(event => {
      const listener = () => {
        if (event === 'focus') record.order = this.nextOrder++;
        this.schedule();
      };
      window.on(event, listener);
      record.listeners.push([event, listener]);
    });
    const destroy = () => {
      this.records = this.records.filter(candidate => candidate !== record);
      this.schedule();
    };
    window.on('destroy', destroy);
    record.listeners.push(['destroy', destroy]);
    this.schedule();
    return true;
  }

  snapshot() {
    const records = this.records.slice().sort((a, b) => a.order - b.order);
    const windows = [];
    records.forEach(record => {
      try {
        const entry = serializeWindow(record, windows.length);
        if (entry) windows.push(entry);
      } catch (error) {
        // A third-party window can expose throwing state accessors. It must not
        // prevent healthy application windows from being persisted.
      }
    });
    return {version: SNAPSHOT_VERSION, windows};
  }

  schedule() {
    if (this.destroyed) return;
    if (this.timer !== null) this.clearTimer(this.timer);
    this.timer = this.setTimer(() => {
      this.timer = null;
      this.flush();
    }, this.delay);
  }

  flush() {
    if (this.destroyed) return Promise.resolve(false);
    if (this.timer !== null) {
      this.clearTimer(this.timer);
      this.timer = null;
    }
    try {
      return Promise.resolve(this.saveSnapshot(this.snapshot())).catch(() => false);
    } catch (error) {
      return Promise.resolve(false);
    }
  }

  destroy() {
    if (this.timer !== null) this.clearTimer(this.timer);
    this.timer = null;
    this.records.forEach(record => {
      if (typeof record.window.off === 'function') {
        record.listeners.forEach(([event, listener]) => record.window.off(event, listener));
      }
    });
    this.records = [];
    this.destroyed = true;
  }
}

export {
  SessionSnapshotCapture,
  SNAPSHOT_VERSION,
  WINDOW_EVENTS,
  serializeWindow,
  validateEntry,
  validateSnapshot
};
