'use strict';

const memoryUI = require('./memory-advanced');
const paramsByDevice = new WeakMap();
const paramsFor = device => {
  if (!paramsByDevice.has(device)) paramsByDevice.set(device, memoryUI.replicatorParams(device));
  const params = paramsByDevice.get(device);
  const ranges = memoryUI.replicatorParams(device);
  for (let number = 1; number <= 4; number++) params[`fc${number}`] = ranges[`fc${number}`];
  return params;
};
const tabs = (doc, selected, disabled, change) => {
  const strip = doc.createElement('nav');
  strip.className = 'memory-subtabs';
  strip.setAttribute('aria-label', 'Device editor sections');
  for (const title of ['Device Definition', 'Advanced Settings']) {
    const button = doc.createElement('button');
    button.type = 'button'; button.textContent = title; button.disabled = disabled;
    button.setAttribute('aria-pressed', String(title === selected));
    button.addEventListener('click', () => change(title));
    strip.appendChild(button);
  }
  return strip;
};
const mount = (doc, editor, params, devices, disabled, shared) => {
  const fields = doc.createElement('fieldset');
  fields.className = 'memory-advanced'; fields.disabled = disabled;
  const output = shared ? shared.output() : {outputLoaded: false};
  memoryUI.mount(fields, params, {document: doc, devices, ...output,
    configureOutput: shared ? () => shared.open() : undefined});
  if (shared) {
    const open = doc.createElement('button'); open.type = 'button'; open.className = 'tool-button';
    open.textContent = 'MMA Settings...'; open.disabled = disabled;
    open.addEventListener('click', () => shared.open()); editor.appendChild(open);
    editor.appendChild(fields); return;
  }
  const note = doc.createElement('p');
  note.textContent = 'Shared RBE TCP settings are not exposed by this OS.js runtime API. Existing output configuration is preserved; enabling rules requires a configured output.';
  editor.append(fields, note);
};
module.exports = {tabs, mount, paramsFor, memoryUI};
