// ICON-002 — deliberate generic fallback for launcher image failures.

const GENERIC_ICON = '/icons/NamelessClassicIcons/icons/application-x-executable.svg';
const LAUNCHER_IMAGE = [
  '.osjs-desktop-iconview__entry img',
  '#osjs-context-menu img',
  '.osjs-panel-item[data-name="menu"] img'
].join(',');

export default class NamelessIconFallbackServiceProvider {
  constructor() {
    this.onError = null;
  }

  provides() { return []; }

  init() {
    this.onError = event => {
      const image = event.target;
      if (!image || typeof image.matches !== 'function' ||
          !image.matches(LAUNCHER_IMAGE) || image.dataset.namelessIconFallback) return;
      image.dataset.namelessIconFallback = '1';
      image.src = GENERIC_ICON;
    };
    document.addEventListener('error', this.onError, true);
    return Promise.resolve(true);
  }

  start() { return Promise.resolve(true); }

  destroy() {
    if (this.onError) document.removeEventListener('error', this.onError, true);
    this.onError = null;
  }
}

export {GENERIC_ICON, LAUNCHER_IMAGE};
