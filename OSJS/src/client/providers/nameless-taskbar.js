// OSUI-011 — Restore the persisted Start button icon at desktop load.
//
// The Taskbar Settings utility (OSUI-010) persists the operator's Start
// button icon choice in the per-user settings store under
// `nameless/taskbar` → {startIcon: <name>} (server adapter, so it travels
// with the workstation like the Auto-Start selection). On `osjs/core:started`
// this provider reads the selection and applies it to the Start button
// image rendered by the panel module. Settings are loaded before the
// started event fires (same guarantee the autostart provider relies on);
// the panel img may lag the event, so the apply retries briefly.

const SETTINGS_NS = 'nameless/taskbar';
const ICONS_BASE = '/icons/NamelessClassicIcons/icons/';

export default class NamelessTaskbarServiceProvider {
  constructor(core) {
    this.core = core;
  }

  provides() {
    return [];
  }

  async init() {
    this.core.on('osjs/core:started', () => this.applySavedIcon());
    return true;
  }

  start() {}

  destroy() {}

  applySavedIcon() {
    let startIcon, clickAnimation, hoverEffect, iconRatio;
    try {
      startIcon = this.core.make('osjs/settings').get(SETTINGS_NS, 'startIcon', null);
      clickAnimation = this.core.make('osjs/settings').get(SETTINGS_NS, 'clickAnimation', false);
      hoverEffect = this.core.make('osjs/settings').get(SETTINGS_NS, 'hoverEffect', 'none');
      iconRatio = Number(this.core.make('osjs/settings').get(SETTINGS_NS, 'iconRatio', 100));
    } catch (e) {
      return;
    }
    // OSUI-012: restore the open/close animation flag (default false =
    // stock instant toggle preserved).
    document.body.classList.toggle('nameless-menu-animation', clickAnimation === true);
    // OSUI-013: restore the hover effect (default 'none' = the classic
    // no-effect treatment).
    document.body.classList.toggle('nameless-hover-animate', hoverEffect === 'animate');
    document.body.classList.toggle('nameless-hover-color', hoverEffect === 'color');
    // OSUI-NEW-003: restore the taskbar icon/button ratio (default 100 =
    // today's normal icon size; the theme scales icons by this property
    // and clips overflow at the button boundary).
    if (Number.isFinite(iconRatio)) {
      document.body.style.setProperty('--nameless-taskbar-icon-ratio', String(iconRatio / 100));
    }
    if (typeof startIcon !== 'string' || !startIcon) return;
    const src = `${ICONS_BASE}${startIcon}.svg`;
    let tries = 0;
    const apply = () => {
      const $img = document.querySelector('.osjs-panel-item[data-name="menu"] img');
      if ($img) {
        $img.src = src;
        return;
      }
      if (++tries < 40) setTimeout(apply, 250);
    };
    apply();
  }
}
