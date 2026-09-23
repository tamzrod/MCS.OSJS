'use strict';

const memoryUI = require('./memory-advanced');
const copy = value => JSON.parse(JSON.stringify(value));
const createSharedSettings = (doc, host, memory) => {
  let persisted = {};
  let draft = {};
  let revision = null;
  let loaded = false;
  let busy = false;
  let uncertain = false;
  let closed = false;
  let dialog = null;
  let returnFocus = null;
  let message = '';
  const listeners = new Set();
  const notify = () => listeners.forEach(callback => callback());
  const button = (title, action, disabled = false) => {
    const node = doc.createElement('button');
    node.type = 'button'; node.className = 'tool-button'; node.textContent = title;
    node.disabled = disabled; node.addEventListener('click', action);
    return node;
  };
  const close = () => {
    if (!dialog || busy) return;
    if (dialog.close) dialog.close();
    dialog.remove(); dialog = null;
    if (returnFocus && returnFocus.isConnected) returnFocus.focus();
  };
  const render = () => {
    if (!dialog || closed) return;
    const fields = doc.createElement('fieldset');
    fields.disabled = busy || !loaded || uncertain;
    memoryUI.mountShared(fields, draft, doc);
    const status = doc.createElement('p'); status.setAttribute('role', 'status');
    status.textContent = message;
    const actions = doc.createElement('div'); actions.className = 'editor-actions';
    actions.append(button('Save & Apply', save, busy || !loaded || uncertain),
      button('Discard', () => { draft = copy(persisted); render(); }, busy || !loaded),
      button('Reload saved settings', load, busy), button('Close', close, busy));
    dialog.replaceChildren(fields, status, actions);
  };
  const load = async () => {
    if (closed || busy) return;
    busy = true; message = 'Loading shared MMA settings...'; render();
    try {
      const result = await memory.loadShared();
      if (closed) return;
      persisted = copy(result.settings); draft = copy(persisted); revision = result.revision;
      loaded = true; uncertain = false; message = 'Shared MMA settings affect all devices. Changes are separate from device drafts.';
    } catch (error) {
      if (closed) return;
      loaded = false; message = `Shared settings unavailable: ${error.message || error}. Update and restart the Simulator backend if unsupported.`;
    } finally { busy = false; if (!closed) { render(); notify(); } }
  };
  async function save() {
    if (closed || busy || !loaded || uncertain) return;
    busy = true; message = 'Saving shared configuration and waiting for restart acknowledgment...'; render();
    try {
      const result = await memory.applyShared(revision, copy(draft));
      if (closed) return;
      persisted = copy(result.settings); draft = copy(persisted); revision = result.revision;
      message = result.message;
    } catch (error) {
      if (closed) return;
      uncertain = error.code !== 'VALIDATION_FAILED' && error.code !== 'INVALID_REQUEST';
      if (error.result && error.result.committed) {
        persisted = copy(error.result.settings); revision = error.result.revision;
      }
      message = `${error.code || 'UNAVAILABLE'}: ${error.message || error}${uncertain ? ' Reload saved settings before another save; the draft has been retained.' : ''}`;
    } finally { busy = false; if (!closed) { render(); notify(); } }
  }
  const api = {
    output: () => ({outputLoaded: loaded && !uncertain, outputListen: persisted.rbe && persisted.rbe.tcp && persisted.rbe.tcp.listen}),
    subscribe: callback => { listeners.add(callback); return () => listeners.delete(callback); },
    open: () => {
      if (closed || dialog) return;
      returnFocus = host.activeElement || doc.activeElement;
      dialog = doc.createElement('dialog'); dialog.className = 'mma-dialog';
      dialog.setAttribute('aria-label', 'Shared MMA Settings'); dialog.setAttribute('aria-modal', 'true');
      dialog.addEventListener('cancel', event => { event.preventDefault(); close(); });
      host.appendChild(dialog); render();
      if (dialog.showModal) dialog.showModal(); else dialog.setAttribute('open', '');
      if (!loaded) load();
    },
    destroy: () => { closed = true; listeners.clear(); if (dialog) dialog.remove(); dialog = null; },
    load
  };
  return api;
};
module.exports = {createSharedSettings};
