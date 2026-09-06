// OS.js client bootstrap for the neutral desktop shell (SHELL-001).
// Registers only the standard OS.js client service providers. No SCADA
// bridge or domain coupling lives client-side.
//
// REWORK-002: import the OS.js base provider stylesheets here (in the JS
// entry) so webpack inlines them into the bundle CSS. (Doing it via SCSS
// `@import "~@osjs/.../dist/main.css"` does NOT inline under dart-sass,
// which treats `.css` imports as plain CSS @import statements — that left
// dist/osjs.css empty and was the root cause of all shell visual defects.)
// The theme/icons packages supply their own runtime <link> overrides after
// the bundle, so they still win the cascade.

import {
  Core,
  CoreServiceProvider,
  DesktopServiceProvider,
  VFSServiceProvider,
  NotificationServiceProvider,
  SettingsServiceProvider,
  AuthServiceProvider
} from '@osjs/client';

import {PanelServiceProvider} from '@osjs/panels';
import {GUIServiceProvider} from '@osjs/gui';
import {DialogServiceProvider} from '@osjs/dialogs';

// OSUI-006: Start Menu panel item whose menu anchor is corrected so the menu
// opens entirely above the bottom taskbar (stock geometry covers the Start
// button and overlaps the taskbar). Replaces the stock `menu` registry entry.
import NamelessMenuPanelItem from './panels/nameless-menu-panel-item.js';

import NamelessAutoStartServiceProvider from './providers/nameless-autostart.js';
// DRES-002: persist the operator's live desktop window snapshot through the
// server settings adapter. Restore is deliberately introduced by DRES-003.
import NamelessSessionServiceProvider from './providers/nameless-session.js';
// OSUI-011: restores the persisted Start button icon (nameless/taskbar
// settings) when the desktop loads.
import NamelessTaskbarServiceProvider from './providers/nameless-taskbar.js';

// OSUI-NEW-006: movable desktop icons (drag-and-drop placement persisted
// per user) plus the Windows-style Auto Arrange desktop context-menu entry.
import NamelessDesktopIconsServiceProvider from './providers/nameless-desktop-icons.js';
import NamelessIconFallbackServiceProvider from './providers/nameless-icon-fallback.js';

// UIX-001: browser-tab icon —the donated shell factory/building plus
// heartbeat SVG replaces the browser's default (generic globe).
import faviconUrl from '../assets/nameless-scada-favicon.svg';

// Base provider CSS — inlined into the bundle by webpack's CSS pipeline.
import '@osjs/client/dist/main.css';
import '@osjs/gui/dist/main.css';
import '@osjs/dialogs/dist/main.css';
import '@osjs/panels/dist/main.css';

import config from './config.js';
import './index.scss';

// OSIA-013: reflect the stock osjs/contextmenu open/close state on the Start
// button. The stock menu panel item only *toggles* a contextmenu and keeps
// no open state, so menu-open is derived from the presence of the stock
// .osjs-contextmenu element in the DOM and mirrored as a class on the panel
// item for theme styling. The attribute guard prevents observer re-entry.
const MENU_SELECTOR = '.osjs-panel-item[data-name=menu] .osjs-panel-item--clickable';

const reflectMenuState = () => {
  const open = !!document.querySelector('.osjs-contextmenu');
  const item = document.querySelector(MENU_SELECTOR);
  if (item && item.getAttribute('data-menu-open') !== String(open)) {
    item.setAttribute('data-menu-open', String(open));
    item.classList.toggle('nameless-menu-open', open);
  }
};

const init = () => {
  const osjs = new Core(config, {});

  osjs.register(CoreServiceProvider);
  osjs.register(DesktopServiceProvider);
  osjs.register(VFSServiceProvider);
  osjs.register(NotificationServiceProvider);
  // OSUI-003: server adapter so settings persist in the user's VFS home
  // (operator profile), not browser localStorage.
  osjs.register(SettingsServiceProvider, {before: true, args: {adapter: 'server'}});
  osjs.register(AuthServiceProvider, {before: true});
  osjs.register(PanelServiceProvider, {args: {registry: {menu: NamelessMenuPanelItem}}});
  osjs.register(DialogServiceProvider);
  osjs.register(GUIServiceProvider);
  osjs.register(NamelessSessionServiceProvider);
  osjs.register(NamelessAutoStartServiceProvider);
  osjs.register(NamelessTaskbarServiceProvider);
  osjs.register(NamelessDesktopIconsServiceProvider);
  osjs.register(NamelessIconFallbackServiceProvider);

  // UIX-001: point the browser tab at the shell favicon (renders the donated
  // factory/building + heartbeat artwork shipped in src/assets..
  const favicon = document.createElement('link');
  favicon.rel = 'icon';
  favicon.type = 'image/svg+xml';
  favicon.href = faviconUrl;
  document.head.appendChild(favicon);

  osjs.boot();

  // DRES-003: a valid session snapshot (including an empty one) is
  // authoritative. Auto-Start remains the clean-start fallback only.
  osjs.on('osjs/core:started', async () => {
    const restored = await osjs.make('nameless/session').apply();
    if (!restored) osjs.make('nameless/autostart').apply();
  });

  new MutationObserver(reflectMenuState).observe(document.body, {
    childList: true,
    subtree: true
  });
};

window.addEventListener('DOMContentLoaded', () => init());
