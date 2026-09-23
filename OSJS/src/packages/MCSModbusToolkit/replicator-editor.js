'use strict';

const advanced = require('./advanced-editor');
const comms = require('./comms-status');

// Toolkit-owned canonical Replicator editor. Fixtures are never passed in;
// a failed load leaves the tab explicitly unavailable rather than editable.
const {copy, blankBlock, blankDevice, normalizeDocument, validateDocument, displayStatus} = require('./replicator-adapter');
const FALLBACK = 'UNKNOWN';
const createReplicatorEditor = (doc, root, replicator, options = {}) => {
  let section = 'Device Definition';
  let advancedSupported = false;
  let receivedAt = 0;
  let documentValue = null;
  let persisted = null;
  let suggestion = null;
  let selected = null;
  let selectedBlock = 0;
  let loading = true;
  let saving = false;
  let closed = false;
  let message = 'Loading canonical Replicator definitions...';
  let messageError = false;
  let status = null;
  let statusError = null;
  let inspection = null;
  let timer = null;
  let generation = 0;
  const h = (tag, className = '', text) => {
    const node = doc.createElement(tag);
    if (className) node.className = className;
    if (text !== undefined) node.textContent = String(text);
    return node;
  };
  const button = (label, action, disabled = false, extra = '') => {
    const node = h('button', `tool-button ${extra}`.trim(), label);
    node.type = 'button'; node.dataset.replicatorAction = action; node.disabled = disabled;
    return node;
  };
  const device = () => documentValue && selected !== null ? documentValue.devices[selected] : null;
  const appliedName = () => {
    const current = device();
    const old = persisted && selected !== null && persisted.devices[selected];
    return current && old && current.name === old.name ? old.name : null;
  };
  const manualKey = () => {
    const current = device();
    const d = current && current.destination;
    return d && !d.auto_port && !d.auto_unit_id ? `${d.port}/${d.unit_id}` : null;
  };
  const inspectConflict = () => inspection && inspection.key === manualKey() && inspection.status === 'IN USE';
  const validation = () => validateDocument(documentValue) || (inspectConflict() ?
    `Destination ${inspection.key} is owned by ${inspection.owner || 'another producer'}.` : null);
  const patchValidation = () => {
    const node = root.querySelector('[data-replicator-validation]');
    if (node) node.textContent = validation() || '';
    const save = root.querySelector('[data-replicator-action="save"]');
    if (save) save.disabled = loading || saving || Boolean(validation());
    const ownership = root.querySelector('[data-replicator-ownership]');
    if (ownership) ownership.textContent = inspection && inspection.key === manualKey() ?
      `${inspection.owner} / ${inspection.status}` : 'UNKNOWN — select Check ownership for a manual pair';
  };
  const patchStatus = () => {
    const name = appliedName();
    const fresh = Date.now() - receivedAt < 6000;
    const observed = displayStatus(name && fresh ? status : null, name, name ? statusError : null);
    comms.update(root.querySelector('#rep-comms'), name && fresh && device().enabled ? status : null,
      name, statusError ? statusError.message || String(statusError) : 'Status is stale, unavailable or belongs to an unsaved device');
    if (options.onStatus) options.onStatus(observed.replicator);
    const rep = root.querySelector('[data-replicator-status="runtime"]');
    const source = root.querySelector('[data-replicator-status="source"]');
    const error = root.querySelector('[data-replicator-status="error"]');
    if (rep) rep.textContent = observed.replicator;
    if (source) source.textContent = observed.source;
    if (error) error.textContent = observed.error ? `Runtime status: ${observed.error}` : '';
    root.querySelectorAll('[data-replicator-block-status]').forEach(node => {
      const index = Number(node.dataset.replicatorBlockStatus);
      const block = observed.blocks[index];
      node.textContent = block && block.index === index && typeof block.source_status === 'string' ?
        `Block ${index + 1}: ${block.source_status}; last poll ${block.last_poll || 'UNKNOWN'}; ${block.last_error || 'no reported error'}` :
        `Block ${index + 1}: ${observed.replicator === 'UNAVAILABLE' ? 'UNAVAILABLE' : FALLBACK}`;
    });
  };
  const stopPolling = () => {
    generation++;
    if (timer !== null) clearTimeout(timer);
    timer = null; status = null; statusError = null; patchStatus();
  };
  const poll = token => {
    const name = appliedName();
    if (closed || loading || saving || token !== generation || !name) return;
    Promise.resolve(replicator.status(name)).then(value => {
      if (closed || token !== generation || appliedName() !== name) return;
      status = value; receivedAt = Date.now(); statusError = null; patchStatus();
    }).catch(error => {
      if (closed || token !== generation || appliedName() !== name) return;
      status = null; statusError = error; patchStatus();
    }).then(() => {
      if (!closed && !saving && token === generation && appliedName() === name) {
        timer = setTimeout(() => poll(token), 1000);
      }
    });
  };
  const restartPolling = () => {
    stopPolling();
    if (!closed && !loading && !saving && appliedName()) poll(generation);
  };
  const field = (label, value, update, options = {}) => {
    const wrapper = h('label', options.cell ? 'cell-field' : 'tool-field');
    wrapper.appendChild(h('span', 'tool-field-label', label));
    const input = h('input'); input.type = options.type || 'number';
    input.value = value === undefined || value === null ? '' : String(value);
    input.setAttribute('aria-label', label); input.disabled = saving;
    input.readOnly = Boolean(options.readOnly);
    if (options.min !== undefined) input.min = String(options.min);
    if (options.max !== undefined) input.max = String(options.max);
    input.addEventListener('input', () => {
      update(input.type === 'number' ? (input.value.trim() ? Number(input.value) : NaN) : input.value);
      if (options.ownership) inspection = null;
      patchValidation();
      if (options.identity) restartPolling();
    });
    wrapper.appendChild(input); return wrapper;
  };
  const checkbox = (label, value, update) => {
    const wrapper = h('label', 'tool-checkbox');
    const input = h('input'); input.type = 'checkbox'; input.checked = Boolean(value); input.disabled = saving;
    input.setAttribute('aria-label', label);
    input.addEventListener('change', () => { update(input.checked); inspection = null; render(); restartPolling(); });
    wrapper.append(input, h('span', '', label)); return wrapper;
  };
  const render = () => {
    if (closed) return;
    const shell = h('div', 'tool-layout replicator-live');
    const side = h('aside', 'tool-sidebar');
    side.appendChild(h('h2', '', 'Replicator Devices'));
    const actions = h('div', 'tool-actions');
    actions.append(button('Add', 'add', loading || saving), button('Duplicate', 'duplicate', loading || saving || selected === null),
      button('Delete', 'delete', loading || saving || selected === null, 'tool-danger'));
    side.appendChild(actions);
    const list = h('div', 'tool-list');
    if (documentValue) documentValue.devices.forEach((item, index) => {
      const row = button(item.name || 'Unnamed', 'select', saving, `tool-row${index === selected ? ' selected' : ''}`);
      row.dataset.index = String(index);
      row.appendChild(h('span', '', `${item.endpoint} / ${Array.isArray(item.pull_blocks) ? item.pull_blocks.length : 0} Pull Blocks`));
      list.appendChild(row);
    });
    if (!list.children.length) list.appendChild(h('div', 'tool-empty', loading ? 'Loading...' :
      documentValue ? 'No devices configured. Choose Add.' : 'No canonical data available.'));
    side.appendChild(list);
    const editor = h('section', 'tool-editor');
    editor.appendChild(advanced.tabs(doc, section, saving, value => { section = value; render(); }));
    const current = device();
    if (current) {
      const com = comms.create(doc);
      editor.appendChild(com);
      const runtime = h('div', 'runtime-row');
      runtime.append(h('span', '', 'Replicator:'), h('strong', '', FALLBACK), h('span', '', 'Source:'), h('strong', '', FALLBACK));
      runtime.children[1].dataset.replicatorStatus = 'runtime';
      runtime.children[3].dataset.replicatorStatus = 'source';
      editor.appendChild(runtime);
      const statusLine = h('div', 'tool-validation'); statusLine.dataset.replicatorStatus = 'error';
      editor.appendChild(statusLine);
      if (section === 'Advanced Settings') {
        if (!advancedSupported) editor.appendChild(h('div', 'tool-validation', 'Update and restart the Replicator backend to edit advanced settings.'));
        else advanced.mount(doc, editor, advanced.paramsFor(current),
          documentValue.devices.map(entry => ({mma2: advanced.paramsFor(entry)})), saving, options.shared);
      } else {
        const identity = h('div', 'tool-grid');
        identity.append(field('Name', current.name, value => { current.name = value; }, {type: 'text', identity: true}),
          checkbox('Enabled', current.enabled, value => { current.enabled = value; }),
          field('Endpoint', current.endpoint, value => { current.endpoint = value; }, {type: 'text'}),
          field('Source Unit ID', current.unit_id, value => { current.unit_id = value; }, {min: 0, max: 255}));
        editor.append(identity, h('h3', '', 'Destination'));
        const d = current.destination;
        const dest = h('div', 'tool-grid');
        dest.append(field('Port', d.port, value => { d.port = value; }, {min: 1, max: 65535, readOnly: d.auto_port, ownership: true}),
          checkbox('Auto Port', d.auto_port, value => { d.auto_port = value; }),
          field('Unit ID', d.unit_id, value => { d.unit_id = value; }, {min: 0, max: 255, readOnly: d.auto_unit_id, ownership: true}),
          checkbox('Auto Unit ID', d.auto_unit_id, value => { d.auto_unit_id = value; }));
        editor.appendChild(dest);
        const ownership = h('div', 'runtime-row'); ownership.appendChild(h('strong', '', 'Destination ownership:'));
        const owner = h('span', '', 'UNKNOWN'); owner.dataset.replicatorOwnership = 'true'; ownership.appendChild(owner);
        ownership.appendChild(button('Check ownership', 'inspect', saving || loading || !manualKey()));
        editor.append(ownership, h('h3', '', 'Pull Blocks'));
        const blockActions = h('div', 'tool-actions');
        blockActions.append(button('Add Block', 'add-block', saving), button('Duplicate Block', 'duplicate-block', saving || !current.pull_blocks.length),
          button('Delete Block', 'delete-block', saving || current.pull_blocks.length <= 1, 'tool-danger'));
        editor.appendChild(blockActions);
        const table = h('div', 'block-table');
        const header = h('div', 'block-row block-header');
        ['#', 'FC', 'Start', 'Count', 'Scan Rate (ms)', ''].forEach(title => header.appendChild(h('span', '', title)));
        table.appendChild(header);
        current.pull_blocks.forEach((block, index) => {
          const row = h('div', `block-row${index === selectedBlock ? ' selected' : ''}`);
          row.appendChild(button(String(index + 1), `block-${index}`, saving));
          const select = h('select'); select.setAttribute('aria-label', `Pull Block ${index + 1} Function`); select.disabled = saving;
          [1, 2, 3, 4].forEach(fc => { const option = h('option', '', `FC${fc}`); option.value = String(fc); select.appendChild(option); });
          select.value = String(block.function);
          select.addEventListener('change', () => { block.function = Number(select.value); patchValidation(); });
          row.append(select, field(`Block ${index + 1} Start`, block.start, value => { block.start = value; }, {cell: true, min: 0, max: 65535}),
            field(`Block ${index + 1} Count`, block.count, value => { block.count = value; }, {cell: true, min: 1, max: 65535}),
            field(`Block ${index + 1} Scan Rate`, block.scan_rate_ms, value => { block.scan_rate_ms = value; }, {cell: true, min: 1, max: 4294967295}),
            button('-', `remove-block-${index}`, saving || current.pull_blocks.length <= 1, 'tool-danger tool-row-delete'));
          const blockStatus = h('div', 'runtime-row', `Block ${index + 1}: UNKNOWN`);
          blockStatus.dataset.replicatorBlockStatus = String(index);
          table.append(row, blockStatus);
        });
        editor.appendChild(table);
      }
    } else editor.appendChild(h('div', 'tool-empty', loading ? 'Loading canonical definitions...' : 'Select a device or choose Add.'));
    if (documentValue) {
      const warning = h('div', 'tool-validation'); warning.dataset.replicatorValidation = 'true';
      editor.appendChild(warning);
      const controls = h('div', 'editor-actions');
      controls.append(button(saving ? 'Saving...' : 'Save & Apply', 'save', saving || Boolean(validation()), 'tool-primary'),
        button('Discard', 'discard', saving));
      editor.appendChild(controls);
    }
    editor.appendChild(h('div', `tool-status${messageError ? ' error' : ''}`, message));
    shell.append(side, editor); root.replaceChildren(shell);
    patchValidation(); patchStatus();
    root.querySelectorAll('[data-replicator-action]').forEach(node => node.addEventListener('click', onAction));
  };
  const setMessage = (text, error = false) => { message = text; messageError = error; render(); };
  const load = async () => {
    try {
      const result = await replicator.load();
      if (closed) return;
      suggestion = result.suggestion;
      advancedSupported = Boolean(result.capabilities && result.capabilities.mma2_advanced);
      persisted = copy(normalizeDocument(result.document)); documentValue = copy(persisted);
      selected = documentValue.devices.length ? 0 : null;
      message = documentValue.devices.length ? 'Canonical Replicator definitions loaded.' : 'No devices configured. Choose Add to begin.';
      messageError = false;
    } catch (error) {
      if (closed) return;
      persisted = null; documentValue = null; selected = null;
      message = `Replicator runtime unavailable: ${error.message || error}`; messageError = true;
    } finally { if (!closed) { loading = false; render(); restartPolling(); } }
  };
  const save = async () => {
    if (closed || loading || saving || !documentValue) return;
    const error = validation();
    if (error) { setMessage(error, true); return; }
    saving = true; stopPolling(); setMessage('Applying Replicator definitions through one runtime transaction...');
    try {
      const result = await replicator.apply(copy(documentValue));
      if (closed) return;
      persisted = copy(normalizeDocument(result.document)); documentValue = copy(persisted);
      selected = selected === null || !documentValue.devices.length ? null : Math.min(selected, documentValue.devices.length - 1);
      inspection = null;
      message = `${result.message} Applied at ${result.completed_at}.`; messageError = false;
    } catch (failure) {
      if (closed) return;
      message = `Save & Apply failed (${failure.code || 'ERROR'}): ${failure.message || failure}`; messageError = true;
    } finally { saving = false; if (!closed) { render(); restartPolling(); } }
  };
  const inspect = async () => {
    const key = manualKey();
    const current = device();
    if (!key || !current || saving || closed) return;
    const port = current.destination.port; const unit = current.destination.unit_id;
    inspection = null; setMessage('Inspecting destination ownership...');
    try {
      const result = await replicator.suggest({inspect: true, port, unit_id: unit});
      if (closed || device() !== current || manualKey() !== key) return;
      inspection = {key, owner: result.owner, status: result.status};
      message = `Destination ${key}: ${result.status} (${result.owner}).`; messageError = result.status === 'IN USE';
    } catch (error) {
      if (closed || device() !== current || manualKey() !== key) return;
      message = `Ownership inspection unavailable: ${error.message || error}`; messageError = true;
    }
    render();
  };
  const add = async duplicate => {
    if (closed || loading || saving || !documentValue) return;
    const source = duplicate && device() ? copy(device()) : null;
    saving = true; setMessage('Requesting destination suggestion...');
    try {
      const proposed = await replicator.suggest();
      if (closed) return;
      const item = source || blankDevice(documentValue.devices.length + 1, proposed);
      if (source) {
        item.name = `${item.name} (copy)`;
        item.destination = {port: proposed.port, unit_id: proposed.unit_id, auto_port: true, auto_unit_id: true};
      }
      documentValue.devices.push(item);
      selected = documentValue.devices.length - 1; selectedBlock = 0; inspection = null;
      message = 'Destination suggestion is advisory, not reserved; Save & Apply validates ownership.'; messageError = false;
    } catch (error) {
      if (closed) return;
      message = `Destination suggestion unavailable: ${error.message || error}`; messageError = true;
    } finally { saving = false; if (!closed) { render(); restartPolling(); } }
  };
  function onAction(event) {
    const action = event.currentTarget.dataset.replicatorAction;
    if (closed || loading || saving || !documentValue) return;
    if (action === 'save') { save(); return; }
    if (action === 'inspect') { inspect(); return; }
    if (action === 'add' || action === 'duplicate') { add(action === 'duplicate'); return; }
    if (action === 'select') { selected = Number(event.currentTarget.dataset.index); selectedBlock = 0; inspection = null; }
    else if (action === 'delete' && device()) {
      documentValue.devices.splice(selected, 1);
      selected = documentValue.devices.length ? Math.min(selected, documentValue.devices.length - 1) : null;
      selectedBlock = 0; inspection = null;
    } else if (action === 'discard') {
      documentValue = copy(persisted);
      selected = documentValue.devices.length ? Math.min(selected || 0, documentValue.devices.length - 1) : null;
      selectedBlock = 0; inspection = null;
      message = 'Unsaved Replicator changes discarded.'; messageError = false;
    } else if (action === 'add-block' && device()) {
      device().pull_blocks.push(blankBlock()); selectedBlock = device().pull_blocks.length - 1;
    } else if (action === 'duplicate-block' && device()) {
      const blocks = device().pull_blocks;
      blocks.splice(selectedBlock + 1, 0, copy(blocks[selectedBlock])); selectedBlock++;
    } else if ((action === 'delete-block' || action.startsWith('remove-block-')) && device()) {
      const blocks = device().pull_blocks;
      if (blocks.length > 1) {
        const index = action.startsWith('remove-block-') ? Number(action.slice(13)) : selectedBlock;
        blocks.splice(index, 1); selectedBlock = Math.min(selectedBlock, blocks.length - 1);
      }
    } else if (action.startsWith('block-')) selectedBlock = Number(action.slice(6));
    render(); restartPolling();
  }
  render(); load();
  const staleTimer = setInterval(patchStatus, 1000);
  const unsubscribeShared = options.shared ? options.shared.subscribe(render) : () => {};
  return {destroy: () => { closed = true; unsubscribeShared(); clearInterval(staleTimer); stopPolling(); root.replaceChildren(); }};
};
module.exports = {createReplicatorEditor};
