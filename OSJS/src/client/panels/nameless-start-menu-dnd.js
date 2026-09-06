// OSUI-NEW-007 — Start Menu drag-and-drop and the New Folder context action.
//
// Interaction layer on top of the OPEN, stock-rendered Start Menu element
// (#osjs-context-menu). Like the band and the desktop icons, the hyperapp-
// managed vnode tree is never modified — drags move a detached ghost clone
// and drops translate into pure layout operations that persist through
// settings; the menu is then re-rendered by a close/reopen cycle through
// the normal Start-button path (which re-reads the layout).
//
// - 5px movement threshold keeps plain clicks (launch) and submenu hovers
//   untouched; a captured click is suppressed once after a real drag so the
//   dragged app is not launched by accident.
// - Right-click on the open Start Menu exposes a single context action:
//   "New Folder" (a small overlay menu, like the ROD DESKTOP band, rendered
//   OUTSIDE the hyperapp subtree), which prompts for a folder label.

import {
  readLayout,
  writeLayout,
  materializeLayout,
  splitStockTree,
  moveRoot,
  moveToFolder,
  addFolder
} from './nameless-start-menu-layout.js';

const DRAG_THRESHOLD = 5;
const STOCK_TREE_KEY = '_namelessStockTree';

// The panel item hands us the stock tree it intercepted; drops materialize
// layouts relative to it.
export function rememberStockTree(menuEl, stock) {
  if (menuEl) {
    menuEl[STOCK_TREE_KEY] = stock;
  }
}

const nameOfEntry = (menuEl, $entry) => {
  const label = ($entry.querySelector('.osjs-gui-menu-label') || {}).textContent;
  // The customized tree keeps the stock label for each app; resolve the
  // package name via the stock app map (labels are the stable surface here).
  const stock = menuEl[STOCK_TREE_KEY] || [];
  const {apps} = splitStockTree(stock);
  for (const [name, node] of apps) {
    if (node.label === label) {
      return name;
    }
  }
  return null;
};

// Root entries that correspond to LAYOUT nodes (apps and folders) — the
// system tail (separators, Save/Log Out) is excluded so drop indices are
// computed in layout space, not DOM space.
const layoutEntries = (menuEl) =>
  Array.from(menuEl.querySelectorAll(':scope > ul > li.osjs-gui-menu-entry'))
    .filter(($e) => nameOfEntry(menuEl, $e) !== null || $e.querySelector(':scope > .osjs-gui-menu-container > ul'));

// Re-render: close the shared context menu, then re-click the Start button
// so the whole customization chain (intercept + transform) runs again.
// The stock ContextMenu's toggle guard resets on a 0ms timeout after hide()
// (see @osjs/gui ContextMenu actions), so re-clicking in the same tick would
// be read as a toggle and close the menu again — the reopen is deferred.
const reopenMenu = (core) => {
  try {
    core.make('osjs/contextmenu').hide();
  } catch (e) {
    // already closed
  }
  setTimeout(() => {
    const $btn = document.querySelector('.osjs-panel-item[data-name=menu] .osjs-panel-item--clickable');
    if ($btn) {
      $btn.click();
    }
  }, 100);
};

const persistAndReopen = (core, menuEl, layout) => {
  writeLayout(core, layout).then(() => reopenMenu(core));
};

function installDrag(core, menuEl) {
  menuEl.addEventListener('pointerdown', (ev) => {
    if (ev.button !== 0) return;
    const $entry = ev.target.closest('li.osjs-gui-menu-entry');
    if (!$entry || !menuEl.contains($entry)) return;
    const name = nameOfEntry(menuEl, $entry);
    if (!name) return; // separators / session actions stay non-draggable

    const startX = ev.clientX;
    const startY = ev.clientY;
    let dragging = false;
    let ghost = null;
    let candidate = null;

    const onMove = (mev) => {
      const dx = mev.clientX - startX;
      const dy = mev.clientY - startY;
      if (!dragging && Math.hypot(dx, dy) < DRAG_THRESHOLD) return;
      if (!dragging) {
        dragging = true;
        document.body.classList.add('nameless-menu-dnd');
        const container = $entry.querySelector('.osjs-gui-menu-container');
        ghost = document.createElement('div');
        ghost.className = 'nameless-menu-dnd__ghost';
        ghost.textContent = container ? container.textContent : name;
        document.body.appendChild(ghost);
      }
      ghost.style.left = `${mev.clientX + 8}px`;
      ghost.style.top = `${mev.clientY - 12}px`;

      // hover feedback over candidate targets
      if (candidate) {
        candidate.classList.remove('nameless-menu-dnd__target');
        candidate = null;
      }
      const hit = dropTargetAt(menuEl, mev.clientX, mev.clientY, $entry);
      if (hit && hit.entry) {
        candidate = hit.entry;
        candidate.classList.add('nameless-menu-dnd__target');
      }
    };

    const onUp = (uev) => {
      window.removeEventListener('pointermove', onMove);
      window.removeEventListener('pointerup', onUp);
      document.body.classList.remove('nameless-menu-dnd');
      if (candidate) {
        candidate.classList.remove('nameless-menu-dnd__target');
        candidate = null;
      }
      if (ghost) {
        ghost.remove();
        ghost = null;
      }
      if (!dragging) return;

      // the entry's own click must not launch the app we just moved
      const suppress = (cev) => {
        cev.stopPropagation();
        cev.preventDefault();
      };
      window.addEventListener('click', suppress, {capture: true, once: true});

      const hit = dropTargetAt(menuEl, uev.clientX, uev.clientY, $entry);
      if (!hit) return;

      const stock = menuEl[STOCK_TREE_KEY] || [];
      const current = readLayout(core) || materializeLayout(stock);
      if (hit.type === 'folder') {
        persistAndReopen(core, menuEl, moveToFolder(current, name, hit.folderName));
      } else if (hit.type === 'root') {
        persistAndReopen(core, menuEl, moveRoot(current, name, hit.index));
      }
    };

    window.addEventListener('pointermove', onMove);
    window.addEventListener('pointerup', onUp);
  });
}

// Resolve the drop target under the pointer. A dragged app lands IN a
// folder when it is released over that folder's open submenu or the MIDDLE
// of the folder row. Releasing near the TOP/BOTTOM EDGE of a folder row
// reorders at root (before/after) — a folder row's middle still accepts
// into-folder drops (Windows-style), while its edges stay reorder-strips so
// dragging OUT of a folder into the root list is always possible.
// App rows order before/after by midpoint. Anything else — the
// session-action tail, empty desktop — cancels the drop.
const EDGE = 0.25;
function dropTargetAt(menuEl, x, y, draggedEntry) {
  const entries = layoutEntries(menuEl);
  for (const $e of entries) {
    if ($e === draggedEntry) continue;
    const r = $e.getBoundingClientRect();
    const $sub = $e.querySelector(':scope > .osjs-gui-menu-container > ul');
    if ($sub) {
      const label = ($e.querySelector('.osjs-gui-menu-label') || {}).textContent;
      if (contains($sub.getBoundingClientRect(), x, y)) {
        return {type: 'folder', entry: $e, folderName: label};
      }
      if (contains(r, x, y)) {
        const h = r.height;
        if (y > r.top + h * EDGE && y < r.bottom - h * EDGE) {
          return {type: 'folder', entry: $e, folderName: label};
        }
      }
    }
    if (contains(r, x, y)) {
      const before = y < r.top + r.height / 2;
      let index = entries.indexOf($e);
      if (!before) index += 1;
      const draggedIndex = entries.indexOf(draggedEntry);
      if (draggedIndex !== -1 && draggedIndex < index) index -= 1;
      return {type: 'root', entry: $e, index};
    }
  }
  return null;
}

const contains = (r, x, y) => x >= r.left && x <= r.right && y >= r.top && y <= r.bottom;

function installNewFolderContext(core, menuEl) {
  menuEl.addEventListener('contextmenu', (ev) => {
    ev.preventDefault();
    ev.stopPropagation();
    openNewFolderOverlay(core, menuEl, ev.clientX, ev.clientY);
  }, true);
}

function openNewFolderOverlay(core, menuEl, x, y) {
  dismissNewFolderOverlay();
  const $menu = document.createElement('div');
  $menu.className = 'nameless-new-folder-menu';
  const $item = document.createElement('button');
  $item.type = 'button';
  $item.className = 'nameless-new-folder-menu__item';
  $item.textContent = 'New Folder';
  $item.addEventListener('click', () => {
    dismissNewFolderOverlay();
    core.make('osjs/dialog', 'prompt', {
      title: 'New Folder',
      message: 'Folder label:',
      value: 'New Folder'
    }, (btn, value) => {
      if (btn !== 'ok' && btn !== 'yes') return;
      const stock = menuEl[STOCK_TREE_KEY] || [];
      const current = readLayout(core) || materializeLayout(stock);
      persistAndReopen(core, menuEl, addFolder(current, value));
    });
  });
  $menu.appendChild($item);
  $menu.style.left = `${x}px`;
  $menu.style.top = `${y}px`;
  document.body.appendChild($menu);
  $menu.dataset.namelessNewFolder = '1';

  const dismiss = (ev) => {
    if ($menu.contains(ev.target)) return;
    dismissNewFolderOverlay();
    window.removeEventListener('pointerdown', dismiss, true);
  };
  window.addEventListener('pointerdown', dismiss, true);
}

function dismissNewFolderOverlay() {
  Array.from(document.querySelectorAll('[data-nameless-new-folder]')).forEach(($m) => $m.remove());
}

// Attach the interaction layer to an OPEN Start Menu element. Idempotent
// per element instance; the shared context-menu element is re-created by
// the stock renderer, so this runs on every open.
export function attachStartMenuDnD(core, menuEl, stockTree) {
  if (!menuEl) return;
  menuEl._namelessCore = core;
  rememberStockTree(menuEl, stockTree);
  if (!menuEl.dataset.namelessMenuDnd) {
    menuEl.dataset.namelessMenuDnd = '1';
    installDrag(core, menuEl);
    installNewFolderContext(core, menuEl);
  }
}
