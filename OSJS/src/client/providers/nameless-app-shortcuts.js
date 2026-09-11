// Application shortcuts that are intentionally pinned to the workstation
// desktop. The stock Start/Application menu remains package-driven; this
// provider only adds the small set of explicit desktop launchers requested by
// MCS.OSJS.

const SHORTCUTS = [
  {
    name: 'ModbusReplicator',
    label: 'Modbus Replicator',
    icon: '/icons/NamelessClassicIcons/icons/application-x-executable.svg'
  }
];

const ensureShortcuts = core => {
  const wrapper = document.querySelector('.osjs-desktop-iconview__wrapper');
  if (!wrapper) return;

  SHORTCUTS.forEach(shortcut => {
    if (wrapper.querySelector(`[data-mcs-shortcut="${shortcut.name}"]`)) return;

    const entry = document.createElement('div');
    entry.className = 'osjs-desktop-iconview__entry';
    entry.dataset.mcsShortcut = shortcut.name;
    entry.tabIndex = 0;
    entry.title = shortcut.label;

    const image = document.createElement('img');
    image.src = shortcut.icon;
    image.alt = '';

    const label = document.createElement('div');
    label.className = 'osjs-desktop-iconview__entry__label';
    label.textContent = shortcut.label;

    const launch = () => core.make('osjs/packages').launch(shortcut.name).catch(() => {});
    entry.addEventListener('dblclick', launch);
    entry.addEventListener('keydown', event => {
      if (event.key === 'Enter') launch();
    });
    entry.append(image, label);
    wrapper.appendChild(entry);
  });
};

export default class NamelessAppShortcutsServiceProvider {
  constructor(core) {
    this.core = core;
    this.observer = null;
  }

  provides() { return []; }

  init() {
    this.core.on('osjs/core:started', () => {
      ensureShortcuts(this.core);
      this.observer = new MutationObserver(() => ensureShortcuts(this.core));
      this.observer.observe(document.body, {childList: true, subtree: true});
    });
    return Promise.resolve(true);
  }

  start() { return Promise.resolve(true); }

  destroy() {
    if (this.observer) this.observer.disconnect();
    this.observer = null;
  }
}

export {SHORTCUTS, ensureShortcuts};
