// OSUI-NEW-007 — Operator-organizable Start Menu (folders + ordering).
//
// Model: the persisted layout is a flat array of ROOT nodes describing the
// application part of the Start Menu:
//   {type: 'app',    name: '<package name>'}                  — app entry
//   {type: 'folder', name: '<label>', items: ['<name>', ...]} — folder
//
// - null layout ⇒ the factory tree from the stock panels builder is used
//   VERBATIM (deterministic factory list; nothing is customized).
// - A layout only ever reorganizes the APPLICATION entries. The system tail
//   of the factory tree (separators, Save Session & Log Out, Log Out) is
//   always appended unchanged, so launching and session actions are
//   structurally unaffected.
// - Apps missing from the layout (newly installed, or a stale layout) are
//   appended at the root, label-sorted, so every app stays reachable.
// - Folders render through the STOCK submenu machinery; one nesting level.
//
// The tree produced by the stock builder is never built again here — the
// customization is a pure transformation of the stock tree, so factory
// fidelity (icons, labels, sorting, i18n) is preserved by construction.

const SETTINGS_NS = 'nameless/startmenu';
const START_MENU_ACTION = 'Log Out'; // stock system-tail marker

export function readLayout(core) {
  try {
    const layout = core.make('osjs/settings').get(SETTINGS_NS, 'layout', null);
    return Array.isArray(layout) ? layout : null;
  } catch (e) {
    return null;
  }
}

export function writeLayout(core, layout) {
  const settings = core.make('osjs/settings');
  settings.set(SETTINGS_NS, 'layout', layout);
  return settings.save().then(() => true).catch(() => false);
}

// Split a stock tree into its application entries and its system tail.
// Application entries are nodes carrying data.name (root apps, category
// items, pinned apps); the tail is separators and session actions.
export function splitStockTree(stock) {
  const apps = new Map(); // name -> {icon, label, data}
  const tail = [];
  const collect = (nodes) => {
    (nodes || []).forEach((node) => {
      if (node && node.items) {
        collect(node.items);
      } else if (node && node.data && node.data.name) {
        apps.set(node.data.name, {icon: node.icon, label: node.label, data: node.data});
      } else if (node && (node.type === 'separator' || (node.data && node.data.action))) {
        tail.push(node);
      }
    });
  };
  collect(stock);
  return {apps, tail};
}

// Materialize a layout from the CURRENT stock tree — used the first time an
// operator drags in factory mode, so category folders become named folders
// and the visible order is preserved.
export function materializeLayout(stock) {
  const layout = [];
  (stock || []).forEach((node) => {
    if (!node || node.type === 'separator' || (node.data && node.data.action)) return;
    if (node.items) {
      layout.push({
        type: 'folder',
        name: String(node.label || 'Folder'),
        items: node.items
          .filter((item) => item && item.data && item.data.name)
          .map((item) => item.data.name)
      });
    } else if (node.data && node.data.name) {
      layout.push({type: 'app', name: node.data.name});
    }
  });
  return layout;
}

// Render a layout against the stock app map; unknown names are skipped,
// unlisted apps are appended label-sorted, the system tail is untouched.
export function renderLayout(layout, stock) {
  const {apps, tail} = splitStockTree(stock);
  if (!layout) {
    return stock;
  }

  const used = new Set();
  const appNode = (name) => {
    const app = apps.get(name);
    if (!app) return null;
    used.add(name);
    return app;
  };

  const root = [];
  layout.forEach((node) => {
    if (!node || typeof node !== 'object') return;
    if (node.type === 'folder') {
      const items = (Array.isArray(node.items) ? node.items : [])
        .map(appNode)
        .filter(Boolean);
      root.push({
        icon: {name: 'folder'},
        label: String(node.name || 'Folder'),
        items
      });
    } else if (node.type === 'app') {
      const app = appNode(node.name);
      if (app) root.push(app);
    }
  });

  // reachability: apps absent from the layout append label-sorted
  const missing = Array.from(apps.entries())
    .filter(([name]) => !used.has(name))
    .map(([, node]) => node)
    .sort((a, b) => String(a.label).toLowerCase().localeCompare(String(b.label).toLowerCase()));

  return root.concat(missing, tail);
}

// Whether a stock tree is the Start Menu (vs. some other context menu).
export function isStartMenuTree(stock) {
  return (stock || []).some((node) => node && node.data && node.data.action === 'logOut') ||
    (stock || []).some((node) => node && node.label === START_MENU_ACTION);
}

// ---- pure layout operations (return a NEW layout) ----

export function moveRoot(layout, name, index) {
  const rest = layout.filter((node) => !(node.type === 'app' && node.name === name));
  const i = Math.max(0, Math.min(rest.length, index));
  return rest.slice(0, i).concat([{type: 'app', name}], rest.slice(i));
}

export function moveToFolder(layout, name, folderName) {
  const without = removeEverywhere(layout, name);
  return without.map((node) => node.type === 'folder' && node.name === folderName
    ? {...node, items: node.items.concat([name])}
    : node);
}

export function removeEverywhere(layout, name) {
  return layout
    .filter((node) => !(node.type === 'app' && node.name === name))
    .map((node) => node.type === 'folder'
      ? {...node, items: node.items.filter((item) => item !== name)}
      : node);
}

export function addFolder(layout, requestedName) {
  const base = String(requestedName || '').trim() || 'New Folder';
  const taken = new Set(layout.filter((n) => n.type === 'folder').map((n) => n.name));
  let name = base;
  for (let i = 2; taken.has(name); i++) {
    name = `${base} (${i})`;
  }
  return layout.concat([{type: 'folder', name, items: []}]);
}
