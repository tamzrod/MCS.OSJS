// OSUI-006: Start Menu panel item with corrected menu geometry.
//
// The stock @osjs/panels MenuPanelItem opens the application menu via
// osjs/contextmenu.show({position: ev.target, ...}). The GUI ContextMenu
// resolves an Element position to the element's BOTTOM-left corner and, when
// the menu would overflow the root, clamps it so the menu's bottom edge sits
// at the requested top. With the taskbar anchored to the bottom edge, the
// Start button's bottom edge is effectively the viewport bottom, so the
// clamped menu hangs DOWN over the taskbar and covers the Start button.
//
// This subclass keeps the stock item's rendering, menu contents and behavior
// untouched and only corrects the anchor handed to the contextmenu: for a
// bottom taskbar it passes the taskbar's TOP edge, so the stock clamp logic
// places the menu's bottom edge exactly on the taskbar's top edge — entirely
// above the taskbar, never covering the Start button. For any other panel
// position the original event (and stock geometry) is used unchanged.
//
// OSUI-008: when the operator-configurable flag
// `application.classicStartMenuBand` is not false (default on), the first
// Start-button click also starts a watcher that keeps the classic
// ROD DESKTOP vertical band attached to the Start Menu's left side (see
// nameless-start-menu-band.js; the shared hyperapp renderer drops
// externally inserted siblings on re-render, so the watcher re-attaches,
// and the shared context-menu element is gated on the Start Menu marker).
// The band carries no behavior; geometry stays governed by the same
// taskbar-top anchor, so the OSUI-006/OSUI-007 invariant still holds with
// the band attached.

import {MenuPanelItem} from '@osjs/panels';
import {watchStartMenuBand} from './nameless-start-menu-band.js';
import {readLayout, renderLayout, isStartMenuTree} from './nameless-start-menu-layout.js';
import {attachStartMenuDnD} from './nameless-start-menu-dnd.js';

export default class NamelessMenuPanelItem extends MenuPanelItem {
  render(state, actions) {
    const vnode = super.render(state, actions);
    const clickable = (vnode.children || []).find(
      (child) => child && child.attributes && typeof child.attributes.onclick === 'function'
    );

    if (!clickable) {
      return vnode;
    }

    const stockOnclick = clickable.attributes.onclick;
    clickable.attributes.onclick = (ev) => {
      const panel = this.panel && this.panel.$element;
      const hasBand = this.core && typeof this.core.config === 'function'
        ? this.core.config('application.classicStartMenuBand') !== false
        : true;

      if (!panel || panel.getAttribute('data-position') !== 'bottom') {
        return stockOnclick(ev);
      }

      const anchor = ev && ev.currentTarget && ev.currentTarget.getBoundingClientRect
        ? ev.currentTarget.getBoundingClientRect()
        : {left: 0};

      // OSUI-NEW-007: interpose on the stock handler's contextmenu.show
      // call so the factory-built tree passes through the operator layout
      // (folders/ordering). The interposition lives ONLY for the duration
      // of this stock onclick, the transform is marker-gated on the Start
      // Menu, and with no saved layout the stock tree passes through
      // verbatim — factory behavior is untouched for everyone else.
      const ctx = this.core.make('osjs/contextmenu');
      const stockShow = typeof ctx.show === 'function' ? ctx.show.bind(ctx) : null;
      if (stockShow) {
        const core = this.core;
        ctx.show = (args) => {
          const stockTree = args && args.menu;
          const transformed = stockTree && isStartMenuTree(stockTree)
            ? renderLayout(readLayout(core), stockTree)
            : stockTree;
          const shown = stockShow(args && transformed !== stockTree
            ? Object.assign({}, args, {menu: transformed})
            : args);
          // the shared element is rendered synchronously by show()
          attachStartMenuDnD(core, document.querySelector('#osjs-context-menu'), stockTree);
          return shown;
        };
      }

      // The stock handler only reads `ev.target` (to derive the menu anchor),
      // and ContextMenu.show accepts a plain {left, top} position verbatim.
      let result;
      try {
        result = stockOnclick({
          target: {
            left: anchor.left,
            top: panel.getBoundingClientRect().top
          }
        });
      } finally {
        if (stockShow) {
          ctx.show = stockShow;
        }
      }

      // OSUI-008: start the band watcher on the first Start-button click
      // (the shared #osjs-context-menu element exists hidden from boot, so
      // an open-transition test cannot gate this); attachment itself is
      // marker- and visibility-gated inside the helper.
      if (hasBand && !this._bandWatcher) {
        this._bandWatcher = watchStartMenuBand(this.core);
      }

      return result;
    };

    return vnode;
  }
}
