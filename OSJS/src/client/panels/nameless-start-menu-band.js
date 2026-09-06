// OSUI-008: Classic Start Menu ROD DESKTOP vertical band.
//
// Windows 2000-style Start Menus carry a narrow vertical product band on
// the menu's LEFT side, drawn with a blue gradient and vertical product
// text. This module renders that band as an absolutely positioned overlay
// in core.$root, aligned pixel-for-pixel to the stock GUI Menu root
// (#osjs-context-menu), while a marker class on the menu root makes the
// entry list reserve the band's width (margin) and height (min-height) —
// see src/client/index.scss. The band's only content is the literal text
// `ROD DESKTOP`; it contains no icons, no images, and no interactive
// content. Menu entries, submenus, and launch behavior are untouched.
//
// Why an overlay: inserting the band as a child of the menu root was tried
// and rejected — the installed @osjs/gui hyperapp diff reconciles the menu
// root's children by index, and the foreign sibling made its patch throw
// (verified: stock build throws zero patch exceptions, in-flow band build
// throws "removeAttribute is not a function" / "childNodes of undefined"),
// aborting re-renders and leaving stale menu content. The overlay lives
// outside every hyperapp-managed subtree, so the stock renderer is fully
// untouched; the marker CLASS it adds to the menu root is safe because
// hyperapp only rewrites attributes whose vnode value changed (the root
// class is constant).
//
// Two installed-runtime facts shape the sync logic:
//
// 1. The context menu element is SHARED across all menus (Start Menu,
//    panel right-click, window right-click), so the band is shown only
//    while the shared element hosts the Start Menu — detected via its
//    marker entry (`Log Out`). Any other menu hides the band and drops
//    the marker class. If the panels locale ever changes the label, the
//    band degrades gracefully to hidden — never fabricated onto an
//    unrelated menu.
// 2. Show/clamp/hide all flow through hyperapp state updates that mutate
//    the shared element's inline style (display/top/left), so a
//    MutationObserver scoped to the context-menu container re-syncs band
//    geometry and visibility after every render.
//
// The helper is idempotent and truth-preserving: a closed or foreign menu
// simply yields a hidden band.

export const BAND_TEXT = 'ROD DESKTOP';
export const BAND_CLASS = 'nameless-start-menu-band';
export const BANDED_CLASS = 'nameless-start-menu--banded';
export const START_MENU_MARKER = 'Log Out';

// OSUI-NEW-005: the band label is operator-customizable through Taskbar
// Settings (persisted under nameless/taskbar -> bandText); BAND_TEXT stays
// the default/fallback when nothing (or blank) is configured.
const SETTINGS_NS = 'nameless/taskbar';

function resolveBandText(core) {
  try {
    const saved = core.make('osjs/settings').get(SETTINGS_NS, 'bandText', BAND_TEXT);
    return typeof saved === 'string' && saved.trim() ? saved : BAND_TEXT;
  } catch (e) {
    return BAND_TEXT;
  }
}

function hasStartMenuMarker(el) {
  const labels = el.querySelectorAll('.osjs-gui-menu-label');
  for (let i = 0; i < labels.length; i++) {
    if (labels[i].textContent === START_MENU_MARKER) {
      return true;
    }
  }
  return false;
}

function createBand() {
  const text = document.createElement('div');
  text.className = BAND_CLASS + '__text';
  text.appendChild(document.createTextNode(BAND_TEXT));
  const band = document.createElement('div');
  band.className = BAND_CLASS;
  band.appendChild(text);
  band.style.display = 'none';
  return band;
}

export function syncStartMenuBand(core) {
  const el = core.$root.querySelector('#osjs-context-menu');
  let band = null;
  for (let i = 0; i < core.$root.children.length; i++) {
    if (core.$root.children[i].classList && core.$root.children[i].classList.contains(BAND_CLASS)) {
      band = core.$root.children[i];
      break;
    }
  }
  // OSUI-NEW-004: the OSUI-012 open/close animation forces the hidden
  // menu's computed display to block (visibility fades out on a delay), so
  // the computed-display check alone left the band painted as a leftover
  // sidebar strip after the menu closed. The inline display is the
  // truthful open/close signal — hyperapp flips it immediately on every
  // toggle, independent of the animation override — while the computed
  // display still guards the boot state (element hidden by stylesheet
  // before the first open, no inline style yet).
  const open = !!(el && getComputedStyle(el).display !== 'none' &&
    el.style.display !== 'none' && hasStartMenuMarker(el));

  if (open) {
    el.classList.add(BANDED_CLASS);
    if (!band) {
      band = createBand();
      core.$root.appendChild(band);
    }
    // OSUI-NEW-005: keep the label in sync with the configured band text.
    const label = resolveBandText(core);
    const textNode = band.querySelector('.' + BAND_CLASS + '__text');
    if (textNode && textNode.textContent !== label) {
      textNode.textContent = label;
    }
    // Align the overlay to the menu root in core.$root coordinates.
    const er = el.getBoundingClientRect();
    const rr = core.$root.getBoundingClientRect();
    band.style.left = (er.left - rr.left) + 'px';
    band.style.top = (er.top - rr.top) + 'px';
    band.style.height = er.height + 'px';
    band.style.display = '';
    return true;
  }

  if (el) {
    el.classList.remove(BANDED_CLASS);
  }
  if (band) {
    band.style.display = 'none';
  }
  return false;
}

// Keeps the band in sync with the Start Menu. Returns the observer so the
// caller can disconnect it; the caller must invoke this at most once per
// desktop session.
export function watchStartMenuBand(core) {
  const sync = () => syncStartMenuBand(core);
  const observer = new MutationObserver(sync);
  const container = core.$root.querySelector('.osjs-system-context-menu');
  if (container) {
    observer.observe(container, {childList: true, subtree: true, attributes: true, attributeFilter: ['style']});
  }
  sync();
  return observer;
}
