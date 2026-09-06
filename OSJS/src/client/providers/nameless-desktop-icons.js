// OSUI-NEW-006 — Movable desktop icons with Windows-style Auto Arrange.
//
// The stock desktop iconview renders entries as in-flow inline-blocks with
// no positioning model. This provider adds free drag-and-drop placement on
// top of it WITHOUT touching the hyperapp renderer (the same overlay
// discipline as the Start Menu band — inserting into the stock vnode tree
// was measured to crash its index-based diff):
//
// - Positions are plain inline absolute styles applied to the rendered
//   entry elements AFTER render, keyed by the entry's label text. hyperapp
//   only patches attributes present in its vnodes (entries carry no style
//   attribute), so the applied positions survive selection re-renders; a
//   MutationObserver reapplies them whenever the entry set is re-rendered
//   (reload, add, remove).
// - Dragging uses pointer events with a small movement threshold, so
//   click-to-select and double-click-to-launch keep working untouched.
// - Positions persist per user through the existing settings mechanism
//   (server adapter) under nameless/desktop -> positions.
// - Auto Arrange is registered through the stock desktop context-menu
//   extension point (osjs/desktop.addContextMenuEntries); it clears every
//   manual position and re-lays all entries into a deterministic
//   label-sorted grid (column-first, like Windows), then persists.
//
// Application registration, launch behavior, the taskbar, and the stock
// file context menus are all untouched.

const SETTINGS_NS = 'nameless/desktop';
const DRAG_THRESHOLD = 5;
const CELL_W = 96;   // grid cell: stock entry is 5em wide + 1em margins
const CELL_H = 120;  // grid cell: stock entry is 6.5em high + 1em margins

const readPositions = (core) => {
  try {
    const saved = core.make('osjs/settings').get(SETTINGS_NS, 'positions', {});
    return saved && typeof saved === 'object' && !Array.isArray(saved) ? saved : {};
  } catch (e) {
    return {};
  }
};

const writePositions = (core, positions) => {
  const settings = core.make('osjs/settings');
  settings.set(SETTINGS_NS, 'positions', positions);
  return settings.save().catch(() => false);
};

const entryLabel = ($entry) => {
  const $label = $entry.querySelector('.osjs-desktop-iconview__entry__label');
  return $label ? $label.textContent : null;
};

const wrapper = () => document.querySelector('.osjs-desktop-iconview__wrapper');

// Applies saved positions to the currently rendered entries. Idempotent;
// entries without a saved position stay in the stock flow layout.
const applyPositions = (core) => {
  const $wrapper = wrapper();
  if (!$wrapper) return;
  const positions = readPositions(core);
  Array.from($wrapper.querySelectorAll('.osjs-desktop-iconview__entry')).forEach(($entry) => {
    const label = entryLabel($entry);
    const pos = label && positions[label];
    if (pos && Number.isFinite(pos.x) && Number.isFinite(pos.y)) {
      $entry.style.position = 'absolute';
      $entry.style.left = `${Math.round(pos.x)}px`;
      $entry.style.top = `${Math.round(pos.y)}px`;
      $entry.style.margin = '0';
    } else if ($entry.dataset.namelessPlaced) {
      // position was cleared (Auto Arrange re-render) — back to flow
      $entry.style.position = '';
      $entry.style.left = '';
      $entry.style.top = '';
      $entry.style.margin = '';
      delete $entry.dataset.namelessPlaced;
    }
    if (pos) {
      $entry.dataset.namelessPlaced = '1';
    }
  });
};

// Deterministic label-sorted grid, column-first (Windows Auto Arrange).
const autoArrange = (core) => {
  const $wrapper = wrapper();
  if (!$wrapper) return;
  const labels = Array.from($wrapper.querySelectorAll('.osjs-desktop-iconview__entry'))
    .map(entryLabel)
    .filter(Boolean)
    .sort((a, b) => a.localeCompare(b));
  const rows = Math.max(1, Math.floor($wrapper.clientHeight / CELL_H));
  const positions = {};
  labels.forEach((label, i) => {
    positions[label] = {
      x: Math.floor(i / rows) * CELL_W,
      y: (i % rows) * CELL_H
    };
  });
  writePositions(core, positions).then(() => applyPositions(core));
};

const clamp = (v, min, max) => Math.min(max, Math.max(min, v));

const installDragging = (core) => {
  const $wrapper = wrapper();
  if (!$wrapper || $wrapper.dataset.namelessDragInstalled) return;
  $wrapper.dataset.namelessDragInstalled = '1';

  $wrapper.addEventListener('pointerdown', (ev) => {
    if (ev.button !== 0) return;
    const $entry = ev.target.closest('.osjs-desktop-iconview__entry');
    if (!$entry) return;
    const label = entryLabel($entry);
    if (!label) return;

    const startX = ev.clientX;
    const startY = ev.clientY;
    const wr = $wrapper.getBoundingClientRect();
    const er = $entry.getBoundingClientRect();
    // origin of the entry inside the wrapper (flow or absolute)
    const baseX = er.left - wr.left;
    const baseY = er.top - wr.top;
    let dragging = false;

    const onMove = (mev) => {
      const dx = mev.clientX - startX;
      const dy = mev.clientY - startY;
      if (!dragging && Math.hypot(dx, dy) < DRAG_THRESHOLD) return;
      if (!dragging) {
        dragging = true;
        document.body.classList.add('nameless-icon-dragging');
        $entry.style.position = 'absolute';
        $entry.style.margin = '0';
        $entry.style.zIndex = '10';
      }
      $entry.style.left = `${clamp(baseX + dx, 0, Math.max(0, $wrapper.clientWidth - er.width))}px`;
      $entry.style.top = `${clamp(baseY + dy, 0, Math.max(0, $wrapper.clientHeight - er.height))}px`;
    };

    const onUp = (uev) => {
      window.removeEventListener('pointermove', onMove);
      window.removeEventListener('pointerup', onUp);
      document.body.classList.remove('nameless-icon-dragging');
      if (!dragging) return; // plain click: selection/launch untouched
      $entry.style.zIndex = '';
      const positions = readPositions(core);
      positions[label] = {
        x: clamp(baseX + (uev.clientX - startX), 0, Math.max(0, $wrapper.clientWidth - er.width)),
        y: clamp(baseY + (uev.clientY - startY), 0, Math.max(0, $wrapper.clientHeight - er.height))
      };
      writePositions(core, positions);
    };

    window.addEventListener('pointermove', onMove);
    window.addEventListener('pointerup', onUp);
  });
};

export default class NamelessDesktopIconsServiceProvider {
  constructor(core) {
    this.core = core;
    this.observer = null;
  }

  provides() {
    return [];
  }

  async init() {
    this.core.on('osjs/core:started', () => this.start_());
    return true;
  }

  start() {}

  destroy() {
    if (this.observer) {
      this.observer.disconnect();
      this.observer = null;
    }
  }

  start_() {
    const core = this.core;
    // Auto Arrange via the stock desktop context-menu extension point.
    try {
      core.make('osjs/desktop').addContextMenuEntries([
        {
          label: 'Auto Arrange',
          onclick: () => autoArrange(core)
        },
        {
          label: 'Reset Desktop Session',
          onclick: () => core.make('nameless/session').reset()
        }
      ]);
    } catch (e) {
      // desktop service unavailable — arrange simply never registers
    }

    // The iconview renders AFTER osjs/core:started (async VFS readdir),
    // and the stock renderer re-renders entries on every reload/add/remove.
    // Watch the whole body (always present) and let the idempotent
    // install/apply pair pick the wrapper up whenever it appears or
    // changes; it also reapplies saved positions after re-renders.
    this.observer = new MutationObserver(() => {
      installDragging(core);
      applyPositions(core);
    });
    this.observer.observe(document.body, {childList: true, subtree: true});
    installDragging(core);
    applyPositions(core);
  }
}
