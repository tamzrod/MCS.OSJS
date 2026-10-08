const memoryUI = (() => {
  const presets = {'Read Only': [1, 2, 3, 4], 'Write Only': [5, 6, 15, 16], 'Read/Write': [1, 2, 3, 4, 5, 6, 15, 16]};
  const functions = {1: 'Read Coils', 2: 'Read Discrete Inputs', 3: 'Read Holding Registers', 4: 'Read Input Registers', 5: 'Write Single Coil', 6: 'Write Single Register', 15: 'Write Multiple Coils', 16: 'Write Multiple Registers'};
  const areas = {coils: 'Coils', discrete_inputs: 'Discrete Inputs', holding_registers: 'Holding Registers', input_registers: 'Input Registers'};
  const states = new WeakMap();
  const sourceAliases = {'All IPv4': '0.0.0.0/0', 'All IPv6': '::/0'};
  const splitSources = value => String(value).split(',').map(item => item.trim()).filter(Boolean).map(item => sourceAliases[item] || item);
  const persistenceEnabled = params => Boolean(params.persistence && params.persistence.enabled === true);
  // Canonical area order maps fc1..fc4 to the four Modbus areas.
  const areaOrder = ['coils', 'discrete_inputs', 'holding_registers', 'input_registers'];
  // Allocated areas of this memory (count > 0), in canonical order. Native
  // persistence ranges are bounded by these authoritative areas.
  const allocatedAreas = params => areaOrder.flatMap((area, index) => {
    const layout = params[`fc${index + 1}`];
    return layout && Number(layout.count) > 0 ? [{area, start: Number(layout.start) || 0, count: Number(layout.count)}] : [];
  });
  // Native persistence is independent of State Sealing and RBE. Explicit ranges,
  // when present, must be nonempty, bounded by the owning allocated area and
  // non-overlapping per area; an omitted ranges block means all allocated areas.
  const persistenceRangeError = params => {
    if (!persistenceEnabled(params)) return null;
    const ranges = params.persistence.ranges;
    if (ranges === undefined || ranges === null) return null;
    const allocated = new Map(allocatedAreas(params).map(entry => [entry.area, entry]));
    let total = 0;
    for (const area of areaOrder) {
      const list = ranges[area];
      if (!list || !list.length) continue;
      const alloc = allocated.get(area);
      if (!alloc) return `${areas[area]} is not an allocated area.`;
      for (const range of list) {
        const start = Number(range.start), count = Number(range.count);
        if (!(count > 0)) return `${areas[area]} range count must be greater than 0.`;
        if (start < alloc.start || start + count > alloc.start + alloc.count) return `${areas[area]} range is outside the allocated area.`;
      }
      const sorted = [...list].sort((left, right) => Number(left.start) - Number(right.start));
      for (let index = 1; index < sorted.length; index++) {
        if (Number(sorted[index].start) < Number(sorted[index - 1].start) + Number(sorted[index - 1].count)) return `${areas[area]} ranges overlap.`;
      }
      total += list.length;
    }
    if (total === 0) return 'Select at least one range, or use All Allocated Areas.';
    return null;
  };
  const defaults = () => ({state_sealing: {enabled: false}, policy: {rules: [{id: 'all-addresses', source_ip: ['0.0.0.0/0', '::/0'], allow_fc: [...presets['Read/Write']]}]}});
  const accessMode = codes => Object.keys(presets).find(key => codes.length === presets[key].length && presets[key].every(code => codes.includes(code))) || 'Custom';
  const replicatorParams = device => {
    device.mma2_advanced ||= {};
    const params = {};
    for (const key of ['policy', 'rbe', 'state_sealing', 'persistence']) Object.defineProperty(params, key, {
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
      const tabTitles = options.persistenceSupported
        ? ['RBE Rules', 'State Sealing', 'Persistence', 'Access Policy']
        : ['RBE Rules', 'State Sealing', 'Access Policy'];
      for (const title of tabTitles) {
        // State Sealing stays an independent feature; the Persistence tab never
        // manages or blocks it.
        const tab = button(title, () => { state.tab = title; draw(); });
        tab.setAttribute('aria-pressed', state.tab === title ? 'true' : 'false');
        tabs.append(tab);
      }
      root.append(tabs);
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
        root.append(element('p',
          'Native MMA2 persistence for this memory. State Sealing and RBE are independent features configured on their own tabs. MMA2 owns snapshot, restore, flush, backup and recovery.',
          'advanced-note'));
        if (!enabled) return;
        const allocated = allocatedAreas(params);
        const usingAll = params.persistence.ranges === undefined || params.persistence.ranges === null;
        const settings = element('div', undefined, 'advanced-form');
        settings.append(
          input('Directory', params.persistence.directory ?? '', value => {
            const text = String(value).trim();
            if (text) params.persistence.directory = text; else delete params.persistence.directory;
          }),
          element('p', 'Leave Directory empty to use MMA2\'s native default location beside the loaded YAML.', 'advanced-note'),
          select('Persisted memory', usingAll ? 'all' : 'selected', [['all', 'All Allocated Areas'], ['selected', 'Selected Ranges']], value => {
            if (value === 'all') delete params.persistence.ranges;
            else if (usingAll) params.persistence.ranges = {};
            draw();
          }));
        root.append(settings);
        const error = persistenceRangeError(params);
        if (error) {
          const message = element('p', error, 'tool-validation');
          message.setAttribute('role', 'alert');
          root.append(message);
        }
        if (!allocated.length) { root.append(element('div', 'No allocated memory areas to persist.', 'tool-empty')); return; }
        if (usingAll) {
          const table = element('div', undefined, 'advanced-rules');
          const header = element('div', undefined, 'rbe-rule-row');
          for (const title of ['Area', 'Start', 'Count']) header.append(element('span', title));
          table.append(header);
          for (const entry of allocated) {
            const row = element('div', undefined, 'rbe-rule-row');
            row.append(element('span', areas[entry.area]), element('span', String(entry.start)), element('span', String(entry.count)));
            table.append(row);
          }
          root.append(table); return;
        }
        for (const entry of allocated) {
          const list = params.persistence.ranges[entry.area] || [];
          root.append(element('h3', areas[entry.area]));
          const table = element('div', undefined, 'advanced-rules');
          list.forEach((range, index) => {
            const row = element('div', undefined, 'rbe-rule-row');
            row.append(
              input('Start', range.start, value => { range.start = Number(value); draw(); }, 'number'),
              input('Count', range.count, value => { range.count = Number(value); draw(); }, 'number'),
              button('Delete', () => { list.splice(index, 1); if (!list.length) delete params.persistence.ranges[entry.area]; draw(); }));
            table.append(row);
          });
          table.append(button('Add range', () => { (params.persistence.ranges[entry.area] ||= []).push({start: entry.start, count: 1}); draw(); }));
          root.append(table);
        }
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
  return {defaults, accessMode, mount, mountShared, nextID, assignCopiedIDs, splitSources, replicatorParams, persistenceEnabled, persistenceRangeError, allocatedAreas};
})();
if (typeof module !== 'undefined') module.exports = memoryUI;
if (typeof window !== 'undefined') window.mcsMemoryUI = memoryUI;
