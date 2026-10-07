const memoryUI = (() => {
  const presets = {'Read Only': [1, 2, 3, 4, 43], 'Write Only': [5, 6, 15, 16], 'Read/Write': [1, 2, 3, 4, 5, 6, 15, 16, 43]};
  const functions = {1: 'Read Coils', 2: 'Read Discrete Inputs', 3: 'Read Holding Registers', 4: 'Read Input Registers', 5: 'Write Single Coil', 6: 'Write Single Register', 15: 'Write Multiple Coils', 16: 'Write Multiple Registers'};
  const areas = {coils: 'Coils', discrete_inputs: 'Discrete Inputs', holding_registers: 'Holding Registers', input_registers: 'Input Registers'};
  functions[43] = 'Read Device Identification';
  const identityFields = {vendor_name: 'Vendor Name', product_code: 'Product Code', major_minor_revision: 'Major / Minor Revision'};
  const validateIdentity = params => {
    if (params.fc43 == null) return null;
    if (typeof params.fc43 !== 'object' || Array.isArray(params.fc43)) return 'Device Identification: FC43 must be an object.';
    for (const [key, label] of Object.entries(identityFields)) {
      if (!Object.prototype.hasOwnProperty.call(params.fc43, key)) continue;
      const value = params.fc43[key];
      if (typeof value !== 'string' || !value.length || value.length > 244 || /[^\x00-\x7f]/.test(value)) {
        return `Device Identification: ${label} must contain 1–244 ASCII bytes.`;
      }
    }
    return null;
  };
  const states = new WeakMap();
  const sourceAliases = {'All IPv4': '0.0.0.0/0', 'All IPv6': '::/0'};
  const splitSources = value => String(value).split(',').map(item => item.trim()).filter(Boolean).map(item => sourceAliases[item] || item);
  const sealingEnabled = params => Boolean(params.state_sealing && params.state_sealing.enabled !== false);
  const persistenceEnabled = params => Boolean(params.persistence && params.persistence.enabled === true);
  // Persistence requires state sealing. This returns an operator-facing reason
  // when persistence is enabled without an enabled state-sealing block, else null.
  const persistenceError = params => {
    if (!persistenceEnabled(params)) return null;
    if (sealingEnabled(params)) return null;
    return 'Persistence requires state sealing to be present and enabled. Enable state sealing first.';
  };
  // One derived persistence RBE rule per present area, start/count taken from the
  // authoritative area. Mirrors the Go projection; used for the locked display.
  const derivedPersistenceRules = params => Object.keys(areas).flatMap((area, index) => {
    const layout = params[`fc${index + 1}`];
    return layout && layout.count > 0 ? [{area, start: layout.start, count: layout.count, system_owned: true}] : [];
  });
  const defaults = () => ({state_sealing: {enabled: false}, policy: {rules: [{id: 'all-addresses', source_ip: ['0.0.0.0/0', '::/0'], allow_fc: [...presets['Read/Write']]}]}});
  const accessMode = codes => Object.keys(presets).find(key => codes.length === presets[key].length && presets[key].every(code => codes.includes(code))) || 'Custom';
  const replicatorParams = device => {
    device.mma2_advanced ||= {};
    const params = {};
    for (const key of ['policy', 'rbe', 'state_sealing', 'fc43', 'persistence']) Object.defineProperty(params, key, {
      get: () => device.mma2_advanced[key], set: value => { device.mma2_advanced[key] = value; }
    });
    for (let number = 1; number <= 4; number++) {
      const blocks = device.pull_blocks.filter(block => Number(block.function) === number);
      const start = blocks.length ? Math.min(...blocks.map(block => Number(block.start))) : 0;
      const end = blocks.length ? Math.max(...blocks.map(block => Number(block.start) + Number(block.count))) : 0;
      params[`fc${number}`] = {start, count: end - start};
    }
    return params;
  };
  const allRules = params => Object.keys(areas).flatMap(area => (params.rbe?.[area] || []).map(rule => ({area, rule})));
  const nextID = devices => {
    const used = new Set(devices.flatMap(device => allRules(device.mma2).map(item => Number(item.rule.id))));
    for (let id = 1; id <= 255; id++) if (!used.has(id)) return id;
    return null;
  };
  const assignCopiedIDs = (params, devices) => {
    const used = new Set(devices.flatMap(device => allRules(device.mma2).map(item => Number(item.rule.id))));
    const available = Array.from({length: 255}, (_, index) => index + 1).filter(id => !used.has(id));
    const rules = allRules(params);
    if (available.length < rules.length) throw new Error('Not enough unused RBE rule IDs to duplicate this device.');
    rules.forEach(({rule}, index) => { rule.id = available[index]; });
  };
  const widgets = document => {
    const element = (tag, text, className) => {
      const node = document.createElement(tag);
      if (text !== undefined) node.textContent = text;
      if (className) node.className = className;
      return node;
    };
    const button = (text, callback, disabled = false) => {
      const node = element('button', text, 'tool-button');
      node.type = 'button'; node.disabled = disabled;
      node.addEventListener('click', callback);
      return node;
    };
    const input = (label, value, callback, type = 'text', disabled = false) => {
      const wrapper = element('label', undefined, 'tool-field');
      const control = element('input');
      control.type = type; control.value = value ?? ''; control.disabled = disabled;
      control.setAttribute('aria-label', label);
      if (type === 'number') { control.min = 0; control.step = 1; }
      control.addEventListener('input', () => callback(type === 'number' ? Number(control.value) : control.value, control));
      wrapper.append(element('span', label, 'tool-field-label'), control);
      return wrapper;
    };
    const select = (label, value, choices, callback, disabled = false) => {
      const wrapper = element('label', undefined, 'tool-field');
      const control = element('select');
      control.setAttribute('aria-label', label); control.disabled = disabled;
      for (const [key, text] of choices) {
        const option = element('option', text); option.value = String(key); control.append(option);
      }
      control.value = String(value);
      control.addEventListener('change', () => callback(control.value));
      wrapper.append(element('span', label, 'tool-field-label'), control);
      return wrapper;
    };
    const checkbox = (label, checked, callback) => {
      const wrapper = element('label', undefined, 'advanced-check');
      const control = element('input'); control.type = 'checkbox'; control.checked = checked;
      control.setAttribute('aria-label', label);
      control.addEventListener('change', () => callback(control.checked));
      wrapper.append(control, element('span', label)); return wrapper;
    };
    return {element, button, input, select, checkbox};
  };

  const mount = (root, params, options) => {
    const {element, button, input, select, checkbox} = widgets(options.document);
    const state = states.get(params) || {tab: 'RBE Rules', selected: 0, custom: new WeakSet()};
    states.set(params, state);
    const draw = () => {
      root.replaceChildren();
      const tabs = element('nav', undefined, 'memory-subtabs');
      for (const title of ['RBE Rules', 'State Sealing', 'Access Policy', 'Device Identification', 'Persistence']) {
        const tab = button(title, () => { state.tab = title; draw(); });
        tab.setAttribute('aria-pressed', state.tab === title ? 'true' : 'false'); tabs.append(tab);
      }
      root.append(tabs);
      if (state.tab === 'Device Identification') {
        root.append(element('h3', 'Device Identification'), element('p', 'FC43 / MEI14 – Read Device Identification. Overrides apply to this memory. Omitted fields use MMA2 defaults; changes require restart.'));
        const form = element('div', undefined, 'advanced-form');
        for (const [key, label] of Object.entries(identityFields)) {
          const field = input(label, params.fc43?.[key], value => { params.fc43 ||= {}; params.fc43[key] = value; });
          field.querySelector('input').placeholder = 'MMA2 default (omitted)';
          form.append(field, button(`Use default for ${label}`, () => {
            params.fc43 ||= {}; delete params.fc43[key]; draw();
          }));
        }
        root.append(form, button('Use MMA2 Defaults', () => { params.fc43 = {}; draw(); }));
        return;
      }
      if (state.tab === 'Persistence') {
        const enabled = persistenceEnabled(params);
        root.append(checkbox('Enable persistence', enabled, checked => {
          if (checked) {
            params.persistence ||= {};
            params.persistence.enabled = true;
          } else if (params.persistence) {
            params.persistence.enabled = false;
          }
          draw();
        }));
        const error = persistenceError(params);
        if (error) {
          const message = element('p', error, 'tool-validation');
          message.setAttribute('role', 'alert');
          root.append(message);
        }
        root.append(element('p', 'Persistence-owned RBE rules are derived from the memory layout and are system-owned; they are locked and cannot be edited or deleted here. User RBE rules remain editable on the RBE Rules tab.', 'advanced-note'));
        const table = element('div', undefined, 'advanced-rules');
        const header = element('div', undefined, 'rbe-rule-row');
        for (const title of ['Area', 'Start', 'Count', 'Ownership']) header.append(element('span', title));
        table.append(header);
        const derived = derivedPersistenceRules(params);
        for (const rule of derived) {
          const row = element('div', undefined, 'rbe-rule-row');
          row.append(element('span', areas[rule.area]), element('span', String(rule.start)), element('span', String(rule.count)), element('span', 'System-owned (locked)'));
          table.append(row);
        }
        if (!derived.length) table.append(element('div', 'No memory areas configured.', 'tool-empty'));
        root.append(table);
        return;
      }
      if (state.tab === 'State Sealing') {
        const seal = params.state_sealing;
        const enabled = Boolean(seal && seal.enabled !== false);
        const ensure = () => params.state_sealing ||= {enabled: false, area: 'coil', address: params.fc1.start, exception: 6};
        root.append(checkbox('Enable state sealing', enabled, checked => {
          const value = ensure(); value.enabled = checked;
          value.area ||= 'coil'; value.address ??= params.fc1.start; value.exception ??= 6; draw();
        }));
        const form = element('div', undefined, 'advanced-form');
        const area = input('Control area', 'Coils', () => {}, 'text', true);
        form.append(area, input('Control address', seal?.address ?? params.fc1.start, value => { ensure().address = value; }, 'number', !enabled));
        form.append(select('Sealed response', seal?.exception ?? 6, [
          [1, '0x01 — Illegal Function'], [2, '0x02 — Illegal Data Address'], [3, '0x03 — Illegal Data Value'],
          [4, '0x04 — Server Device Failure'], [5, '0x05 — Acknowledge'], [6, '0x06 — Server Device Busy'],
          [8, '0x08 — Memory Parity Error'], [10, '0x0A — Gateway Path Unavailable'], [11, '0x0B — Gateway Target Failed to Respond']
        ], value => { ensure().exception = Number(value); }, !enabled));
        root.append(form); return;
      }
      if (state.tab === 'RBE Rules') {
        const listen = options.outputListen;
        const port = listen?.match(/:(\d+)$/)?.[1];
        root.append(element('div', `RBE TCP Port: ${options.outputLoaded ? port || 'Not configured' : 'Unavailable'}`, 'runtime-row'));
        if (options.configureOutput) root.append(button('RBE TCP Settings...', options.configureOutput));
        const actions = element('div', undefined, 'tool-actions');
        actions.append(button('Add', () => {
          const id = nextID(options.devices);
          if (id === null) return;
          const area = Object.keys(areas).find((key, index) => params[`fc${index + 1}`]?.count > 0);
          if (!area) return;
          params.rbe ||= {}; params.rbe[area] ||= [];
          params.rbe[area].push({id, name: `Rule-${id}`, start: params[`fc${Object.keys(areas).indexOf(area) + 1}`].start, count: 1}); draw();
        }, nextID(options.devices) === null || ![1, 2, 3, 4].some(number => params[`fc${number}`]?.count > 0)));
        root.append(actions);
        const table = element('div', undefined, 'advanced-rules');
        for (const {area, rule} of allRules(params)) {
          const row = element('div', undefined, 'rbe-rule-row');
          row.append(input('ID', rule.id, value => { rule.id = value; }, 'number'),
            input('Name', rule.name, value => { rule.name = value; }),
            select('Area', area, Object.entries(areas), value => {
              params.rbe[area].splice(params.rbe[area].indexOf(rule), 1);
              params.rbe[value] ||= []; params.rbe[value].push(rule); draw();
            }), input('Start', rule.start, value => { rule.start = value; }, 'number'),
            input('Count', rule.count, value => { rule.count = value; }, 'number'),
            button('Delete', () => {
              params.rbe[area].splice(params.rbe[area].indexOf(rule), 1);
              if (!allRules(params).length) params.rbe = null;
              draw();
            }));
          table.append(row);
        }
        if (!allRules(params).length) table.append(element('div', 'No RBE rules.', 'tool-empty'));
        root.append(table); return;
      }
      const rules = params.policy?.rules || [];
      state.selected = Math.min(state.selected, Math.max(0, rules.length - 1));
      const actions = element('div', undefined, 'tool-actions');
      const move = offset => {
        const target = state.selected + offset;
        [rules[state.selected], rules[target]] = [rules[target], rules[state.selected]]; state.selected = target; draw();
      };
      actions.append(button('Add', () => {
        params.policy ||= {rules: []}; params.policy.rules ||= [];
        params.policy.rules.push({id: `Rule-${rules.length + 1}`, source_ip: ['0.0.0.0/0', '::/0'], allow_fc: [...presets['Read/Write']]});
        state.selected = params.policy.rules.length - 1; draw();
      }), button('Delete', () => { rules.splice(state.selected, 1); draw(); }, !rules.length),
      button('↑', () => move(-1), !rules.length || state.selected === 0),
      button('↓', () => move(1), !rules.length || state.selected === rules.length - 1));
      root.append(actions);
      const table = element('div', undefined, 'policy-rules');
      const header = element('div', undefined, 'policy-columns fc-header');
      for (const title of ['Order', 'Rule ID', 'Source IP / CIDR', 'Access']) header.append(element('span', title));
      table.append(header);
      rules.forEach((rule, index) => {
        const row = button('', () => { state.selected = index; draw(); });
        row.append(element('span', String(index + 1)), element('span', rule.id), element('span', (rule.source_ip || []).join(', ')), element('span', accessMode(rule.allow_fc || [])));
        row.className = `tool-button policy-rule policy-columns${index === state.selected ? ' selected' : ''}`; table.append(row);
      });
      root.append(table);
      const rule = rules[state.selected];
      if (!rule) { root.append(element('div', 'No access rules.', 'tool-empty')); return; }
      const form = element('div', undefined, 'advanced-form');
      form.append(input('Rule ID', rule.id, value => { rule.id = value; }));
      const sourceDrafts = [...rule.source_ip || []];
      sourceDrafts.forEach((source, index) => {
        const row = element('div', undefined, 'source-row');
        const display = Object.keys(sourceAliases).find(key => sourceAliases[key] === source) || source;
        const update = value => {
          sourceDrafts[index] = value === 'Custom' ? '' : value;
          rule.source_ip = sourceDrafts.flatMap(splitSources);
        };
        const field = input('Source IP / CIDR', display, (value, control) => {
          update(value);
          if (value === 'Custom') control.value = '';
        });
        const control = field.querySelector('input');
        control.placeholder = 'IP / CIDR, separated by commas';
        const presetsField = select('Source presets', '', [['', '▼'], ...['All IPv4', 'All IPv6', 'Custom'].map(value => [value, value])], value => {
          if (!value) return;
          update(value); control.value = value === 'Custom' ? '' : value;
          presetsField.querySelector('select').value = '';
          control.focus?.();
        });
        presetsField.className = 'source-presets';
        row.append(field, presetsField, button('Remove source', () => {
          sourceDrafts.splice(index, 1); rule.source_ip = sourceDrafts.flatMap(splitSources); draw();
        })); form.append(row);
      });
      form.append(button('Add source', () => { rule.source_ip ||= []; rule.source_ip.push(''); draw(); }));
      const mode = state.custom.has(rule) ? 'Custom' : accessMode(rule.allow_fc || []);
      form.append(select('Access', mode, [...Object.keys(presets), 'Custom'].map(value => [value, value]), value => {
        if (value === 'Custom') state.custom.add(rule);
        else { state.custom.delete(rule); rule.allow_fc = [...presets[value]]; }
        draw();
      }));
      if (mode === 'Custom') {
        const checks = element('div', undefined, 'function-checks');
        for (const [code, title] of Object.entries(functions)) {
          checks.append(checkbox(`FC${code} ${title}`, (rule.allow_fc || []).includes(Number(code)), checked => {
            state.custom.add(rule);
            rule.allow_fc = (rule.allow_fc || []).filter(value => value !== Number(code));
            if (checked) rule.allow_fc.push(Number(code));
            rule.allow_fc.sort((left, right) => left - right);
          }));
        }
        form.append(checks);
      }
      root.append(form);
    };
    draw();
  };

  const mountShared = (root, settings, document) => {
    const {element, input, checkbox} = widgets(document);
    const draw = () => {
      root.replaceChildren();
      root.append(element('h2', 'MMA Settings'));
      root.append(checkbox('Enable RBE TCP output', Boolean(settings.rbe), checked => {
        settings.rbe = checked ? {tcp: {listen: ':9001'}} : null; draw();
      }));
      if (settings.rbe) root.append(input('RBE TCP listen address (IP:port)', settings.rbe.tcp?.listen || '', value => { settings.rbe.tcp ||= {}; settings.rbe.tcp.listen = value.trim(); }));
      root.append(checkbox('Debug logging', Boolean(settings.debug), checked => { settings.debug = checked; }));
      root.append(checkbox('Access-event logging', Boolean(settings.access_events?.enabled), checked => {
        settings.access_events ||= {mode: 'rate', window: 5, key_fields: ['src_ip', 'function_code', 'action', 'status', 'port', 'unit'], include_counter: true, limits: {max_keys: 10000, ttl: 60}, output: {type: 'http_stream', listen: '', path: '/events'}};
        settings.access_events.enabled = checked; draw();
      }));
      if (settings.access_events?.enabled) {
        const config = settings.access_events;
        config.output ||= {}; config.limits ||= {};
        const form = element('div', undefined, 'advanced-form');
        form.append(input('HTTP listen address', config.output.listen, value => { config.output.listen = value; }),
          input('HTTP path', config.output.path, value => { config.output.path = value; }),
          input('Window (seconds)', config.window, value => { config.window = value; }, 'number'),
          input('Maximum keys', config.limits.max_keys, value => { config.limits.max_keys = value; }, 'number'),
          input('TTL (seconds)', config.limits.ttl, value => { config.limits.ttl = value; }, 'number'),
          checkbox('Include counter', Boolean(config.include_counter), checked => { config.include_counter = checked; }));
        root.append(form);
      }
    };
    draw();
  };
  return {defaults, accessMode, mount, mountShared, nextID, assignCopiedIDs, splitSources, replicatorParams, validateIdentity, sealingEnabled, persistenceEnabled, persistenceError, derivedPersistenceRules};
})();
if (typeof module !== 'undefined') module.exports = memoryUI;
if (typeof window !== 'undefined') window.mcsMemoryUI = memoryUI;
