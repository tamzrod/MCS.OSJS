const test = require('node:test');
const assert = require('node:assert/strict');
const ui = require('../renderer/memory-advanced');

class Element {
  constructor(tag) { this.tagName = tag; this.children = []; this.events = {}; this.attributes = {}; }
  append(...children) { this.children.push(...children); }
  replaceChildren() { this.children = []; }
  setAttribute(key, value) { this.attributes[key] = value; }
  addEventListener(key, callback) { this.events[key] = callback; }
  querySelector(selector) { return collect(this).find(node => node !== this && node.tagName === selector) || null; }
}
const collect = root => [root, ...root.children.flatMap(collect)];
const document = {createElement: tag => new Element(tag)};
const find = (root, label) => collect(root).find(node => node.attributes['aria-label'] === label);
const click = (root, title) => {
  const button = collect(root).find(node => node.tagName === 'button' && node.textContent === title);
  assert.ok(button, `missing button ${title}`); assert.equal(button.disabled, false); button.events.click();
};
const change = (node, value) => { node.value = value; node.events.change(); };
const type = (node, value) => { node.value = value; node.events.input(); };
const setup = (params = {...ui.defaults(), fc1: {start: 10, count: 16}}, options = {}) => {
  const root = new Element('div');
  const devices = [{mma2: params}];
  const merged = {document, devices, outputLoaded: true, outputListen: '127.0.0.1:19001', persistenceSupported: true, rbeAvailable: true, ...options};
  ui.mount(root, params, merged);
  return {root, params, options: merged};
};

test('new defaults are independent and opening advanced tabs does not alter saved settings', () => {
  const first = ui.defaults(); const second = ui.defaults();
  assert.deepEqual(first.policy.rules[0].source_ip, ['0.0.0.0/0', '::/0']);
  assert.equal(first.state_sealing.enabled, false);
  first.policy.rules[0].source_ip.pop(); assert.equal(second.policy.rules[0].source_ip.length, 2);
  const params = {fc1: {start: 10, count: 16}};
  const before = JSON.stringify(params);
  const {root} = setup(params);
  click(root, 'State Sealing'); click(root, 'Access Policy'); click(root, 'RBE Rules');
  assert.equal(JSON.stringify(params), before);
});

test('Replicator advanced editor uses destination area unions and persists only advanced fields', () => {
  const device = {pull_blocks: [{function: 1, start: 10, count: 4}, {function: 1, start: 20, count: 2}, {function: 3, start: 100, count: 8}], mma2_advanced: ui.defaults()};
  const params = ui.replicatorParams(device);
  assert.deepEqual(params.fc1, {start: 10, count: 12});
  assert.deepEqual(params.fc3, {start: 100, count: 8});
  const {root} = setup(params);
  click(root, 'Add');
  assert.equal(device.mma2_advanced.rbe.coils[0].start, 10);
  click(root, 'State Sealing');
  const enabled = find(root, 'Enable state sealing'); enabled.checked = true; enabled.events.change();
  type(find(root, 'Control address'), '11');
  assert.equal(device.mma2_advanced.state_sealing.address, 11);
  assert.equal(device.mma2_advanced.fc1, undefined);
  assert.equal(ui.replicatorParams(device).rbe.coils[0].id, 1);
});

test('state sealing controls are disabled while off and emit the selected values without help text', () => {
  const {root, params} = setup(); click(root, 'State Sealing');
  assert.equal(find(root, 'Control address').disabled, true);
  const enabled = find(root, 'Enable state sealing'); enabled.checked = true; enabled.events.change();
  type(find(root, 'Control address'), '12'); change(find(root, 'Sealed response'), '2');
  assert.deepEqual(params.state_sealing, {enabled: true, area: 'coil', address: 12, exception: 2});
  assert.equal(find(root, 'Control area').disabled, true);
  assert.ok(!collect(root).some(node => /0 =|1 =|Modbus reads/.test(node.textContent || '')));
});

test('access checkboxes are Custom-only and preset/source selections use canonical values', () => {
  const {root, params} = setup(); click(root, 'Access Policy');
  assert.equal(find(root, 'Access').value, 'Read/Write');
  assert.equal(collect(root).filter(node => node.type === 'checkbox').length, 0);
  change(find(root, 'Access'), 'Read Only'); assert.deepEqual(params.policy.rules[0].allow_fc, [1, 2, 3, 4]);
  change(find(root, 'Access'), 'Write Only'); assert.deepEqual(params.policy.rules[0].allow_fc, [5, 6, 15, 16]);
  change(find(root, 'Access'), 'Custom');
  assert.equal(collect(root).filter(node => node.type === 'checkbox').length, 8);
  const coil = find(root, 'FC1 Read Coils'); coil.checked = true; coil.events.change();
  assert.ok(params.policy.rules[0].allow_fc.includes(1));
  const source = find(root, 'Source IP / CIDR');
  assert.equal(source.attributes.list, undefined);
  type(source, 'All IPv6'); assert.equal(params.policy.rules[0].source_ip[0], '::/0');
  change(find(root, 'Source presets'), 'Custom'); assert.equal(source.value, '');
  type(source, '192.168.4.0/24'); assert.equal(params.policy.rules[0].source_ip[0], '192.168.4.0/24');
  change(find(root, 'Source presets'), 'All IPv4');
  assert.equal(params.policy.rules[0].source_ip[0], '0.0.0.0/0');
  change(find(root, 'Access'), 'Read/Write');
  assert.equal(collect(root).filter(node => node.type === 'checkbox').length, 0);
});

test('comma-separated sources normalize without shifting later source rows while typing', () => {
  const {root, params} = setup(); click(root, 'Access Policy');
  const source = find(root, 'Source IP / CIDR');
  type(source, '192.168.1.5, 10.0.0.0/8, ,2001:db8::/32');
  assert.deepEqual(params.policy.rules[0].source_ip, ['192.168.1.5', '10.0.0.0/8', '2001:db8::/32', '::/0']);
  type(source, '192.168.1.6, 10.0.0.0/8');
  assert.deepEqual(params.policy.rules[0].source_ip, ['192.168.1.6', '10.0.0.0/8', '::/0']);
  change(find(root, 'Source presets'), 'All IPv6');
  assert.deepEqual(params.policy.rules[0].source_ip, ['::/0', '::/0']);
});

test('policy ordering, removal and draft edits persist across subtab changes', () => {
  const {root, params} = setup(); click(root, 'Access Policy'); click(root, 'Add');
  type(find(root, 'Rule ID'), 'second'); click(root, '↑');
  assert.equal(params.policy.rules[0].id, 'second');
  click(root, 'State Sealing'); click(root, 'Access Policy');
  assert.equal(find(root, 'Rule ID').value, 'second');
  click(root, 'Remove source'); assert.equal(params.policy.rules[0].source_ip.length, 1);
  click(root, 'Add source'); assert.equal(params.policy.rules[0].source_ip.length, 2);
  click(root, 'Delete'); assert.equal(params.policy.rules.length, 1);
});

test('RBE displays the configured port and edits independent rules with unused IDs', () => {
  const {root, params, options} = setup();
  assert.ok(collect(root).some(node => node.textContent === 'RBE TCP Port: 19001'));
  options.devices.push({mma2: {rbe: {coils: [{id: 1}]}}});
  click(root, 'Add'); assert.equal(params.rbe.coils[0].id, 2);
  type(find(root, 'Name'), 'temperature'); type(find(root, 'Start'), '11');
  click(root, 'State Sealing'); click(root, 'RBE Rules');
  assert.equal(find(root, 'Name').value, 'temperature');
  assert.equal(params.rbe.coils[0].start, 11);
  click(root, 'Delete'); assert.equal(params.rbe, null);
  options.outputListen = undefined; ui.mount(root, params, options);
  assert.ok(collect(root).some(node => node.textContent === 'RBE TCP Port: Not configured'));
});

test('shared output edits stay separate from per-device data', () => {
  const root = new Element('div'); const settings = {};
  ui.mountShared(root, settings, document);
  const enabled = find(root, 'Enable RBE TCP output'); enabled.checked = true; enabled.events.change();
  assert.equal(settings.rbe.tcp.listen, ':9001');
  type(find(root, 'RBE TCP listen address (IP:port)'), '[::1]:9900');
  assert.equal(settings.rbe.tcp.listen, '[::1]:9900');
  const disabled = find(root, 'Enable RBE TCP output'); disabled.checked = false; disabled.events.change();
  assert.equal(settings.rbe, null);
});

test('device copies get unused global RBE IDs and exhaustion leaves the copy unchanged', () => {
  const original = {rbe: {coils: [{id: 1}, {id: 3}]}};
  const copied = JSON.parse(JSON.stringify(original));
  ui.assignCopiedIDs(copied, [{mma2: original}]);
  assert.deepEqual(copied.rbe.coils.map(rule => rule.id), [2, 4]);
  assert.deepEqual(original.rbe.coils.map(rule => rule.id), [1, 3]);
  const full = {rbe: {coils: Array.from({length: 255}, (_, index) => ({id: index + 1}))}};
  assert.throws(() => ui.assignCopiedIDs(copied, [{mma2: full}]), /unused RBE/);
  assert.deepEqual(copied.rbe.coils.map(rule => rule.id), [2, 4]);
});

test('persistence tab round-trips persistence.enabled through the same draft', () => {
  const params = {...ui.defaults(), fc1: {start: 10, count: 16}};
  const {root} = setup(params);
  const tabOrder = collect(root).filter(node => node.tagName === 'button' && ['RBE Rules', 'State Sealing', 'Persistence', 'Access Policy'].includes(node.textContent)).map(node => node.textContent);
  assert.deepEqual(tabOrder, ['RBE Rules', 'State Sealing', 'Persistence', 'Access Policy']);
  click(root, 'Persistence');
  const enable = find(root, 'Enable persistence');
  assert.equal(enable.checked, false);
  enable.checked = true; enable.events.change();
  assert.deepEqual(params.persistence, {enabled: true});
  // NPE-02: enabling persistence must not touch State Sealing or RBE.
  assert.equal(params.state_sealing.enabled, false);
  assert.equal(params.rbe, undefined);
  // Disabling only flips the persistence flag.
  enable.checked = false; enable.events.change();
  assert.deepEqual(params.persistence, {enabled: false});
  assert.equal(params.state_sealing.enabled, false);
  // Navigating away and back preserves the draft value.
  click(root, 'RBE Rules'); click(root, 'Persistence');
  assert.equal(find(root, 'Enable persistence').checked, false);
});

test('persistence is hidden where unsupported and shows all allocated areas by default', () => {
  const params = {...ui.defaults(), fc1: {start: 10, count: 16}, fc3: {start: 100, count: 8}};
  const {root} = setup(params, {persistenceSupported: false});
  assert.ok(!collect(root).some(node => node.tagName === 'button' && node.textContent === 'Persistence'));
  // Support it again: default is All Allocated Areas with no explicit ranges.
  ui.mount(root, params, {document, devices: [{mma2: params}], outputLoaded: true, persistenceSupported: true});
  click(root, 'Persistence');
  const enable = find(root, 'Enable persistence'); enable.checked = true; enable.events.change();
  assert.equal(find(root, 'Persisted memory').value, 'all');
  assert.equal(params.persistence.ranges, undefined);
  assert.ok(collect(root).some(node => /State Sealing and RBE are independent/.test(node.textContent || '')));
  const rows = collect(root).filter(node => node.className === 'rbe-rule-row');
  const text = rows.map(row => collect(row).map(node => node.textContent).join('|'));
  assert.ok(text.some(entry => /Coils|10|16/.test(entry)));
  assert.ok(text.some(entry => /Holding Registers|100|8/.test(entry)));
  // Default view is read-only: only the enable checkbox and the directory field.
  assert.equal(collect(root).filter(node => node.tagName === 'input').length, 2);
  // No persistence RBE, lock coil, unseal, restore or snapshot controls remain.
  assert.ok(!collect(root).some(node => /Lock coil|Unseal|Restore|Snapshot|RBE rules are derived/.test(node.textContent || '')));
});

test('persistence directory is optional and custom overrides round-trip', () => {
  const params = {...ui.defaults(), fc1: {start: 10, count: 16}};
  const {root} = setup(params);
  click(root, 'Persistence');
  const enable = find(root, 'Enable persistence'); enable.checked = true; enable.events.change();
  assert.equal(params.persistence.directory, undefined);
  type(find(root, 'Directory'), '/var/lib/mma2/unit1');
  assert.equal(params.persistence.directory, '/var/lib/mma2/unit1');
  type(find(root, 'Directory'), '   ');
  assert.equal(params.persistence.directory, undefined);
});

test('selected custom ranges validate against allocated areas and round-trip', () => {
  const params = {...ui.defaults(), fc1: {start: 10, count: 16}, fc3: {start: 100, count: 8}};
  const {root} = setup(params);
  click(root, 'Persistence');
  const enable = find(root, 'Enable persistence'); enable.checked = true; enable.events.change();
  change(find(root, 'Persisted memory'), 'selected');
  assert.deepEqual(params.persistence.ranges, {});
  // Empty explicit ranges are invalid until a range exists.
  const text = () => collect(root).map(node => node.textContent).join('|');
  assert.ok(/Select at least one range, or use All Allocated Areas\./.test(text()));
  // Add a range under Coils.
  const addButtons = collect(root).filter(node => node.tagName === 'button' && node.textContent === 'Add range');
  assert.equal(addButtons.length, 2); // one per allocated area
  addButtons[0].events.click();
  assert.deepEqual(params.persistence.ranges.coils, [{start: 10, count: 1}]);
  const starts = collect(root).filter(node => node.attributes['aria-label'] === 'Start');
  const counts = collect(root).filter(node => node.attributes['aria-label'] === 'Count');
  type(starts[0], '12'); type(counts[0], '4');
  assert.deepEqual(params.persistence.ranges.coils, [{start: 12, count: 4}]);
  // Out-of-area range is rejected (allocated Coils is [10..26)).
  type(starts[0], '24'); type(counts[0], '4');
  assert.ok(/Coils range is outside the allocated area\./.test(text()));
  type(starts[0], '12'); type(counts[0], '4');
  assert.ok(!/Coils range is outside the allocated area\./.test(text()));
  // Switching back to All Allocated Areas clears explicit ranges.
  change(find(root, 'Persisted memory'), 'all');
  assert.equal(params.persistence.ranges, undefined);
});

test('persistence validates independently of State Sealing and RBE', () => {
  // Range validation is a pure function of the persistence block and layout.
  const on = {persistence: {enabled: true}, state_sealing: {enabled: true}, fc1: {start: 10, count: 16}};
  assert.equal(ui.persistenceRangeError(on), null);
  assert.equal(ui.persistenceRangeError({...on, state_sealing: {enabled: false}}), null);
  assert.equal(ui.persistenceRangeError({...on, rbe: {tcp: {listen: ':9001'}}}), null);
  assert.match(ui.persistenceRangeError({persistence: {enabled: true, ranges: {coils: [{start: 0, count: 4}]}}, fc1: {start: 10, count: 16}}), /outside the allocated area/);

  const params = {...ui.defaults(), fc1: {start: 10, count: 16}};
  const {root} = setup(params);
  click(root, 'Persistence');
  const enable = find(root, 'Enable persistence'); enable.checked = true; enable.events.change();
  // State Sealing tab stays independently usable while persistence is enabled.
  click(root, 'State Sealing');
  const sealingTab = collect(root).find(node => node.tagName === 'button' && node.textContent === 'State Sealing');
  assert.equal(sealingTab.disabled, false);
  const seal = find(root, 'Enable state sealing'); seal.checked = true; seal.events.change();
  assert.equal(params.state_sealing.enabled, true);
  // Sealing state does not affect the persistence draft.
  assert.deepEqual(params.persistence, {enabled: true});
});

test('Replicator destination exposes native persistence parity bounded to pull-block areas', () => {
  const device = {pull_blocks: [{function: 1, start: 10, count: 4}, {function: 3, start: 100, count: 8}], mma2_advanced: ui.defaults()};
  const params = ui.replicatorParams(device);
  const root = new Element('div');
  ui.mount(root, params, {document, devices: [{mma2: params}], outputLoaded: true, persistenceSupported: true});
  // The destination allocation is the pull-block union.
  assert.deepEqual(ui.allocatedAreas(params).map(entry => [entry.area, entry.start, entry.count]), [['coils', 10, 4], ['holding_registers', 100, 8]]);
  click(root, 'Persistence');
  const enable = find(root, 'Enable persistence'); enable.checked = true; enable.events.change();
  // The native block is written to the Replicator's own advanced settings.
  assert.deepEqual(device.mma2_advanced.persistence, {enabled: true});
  change(find(root, 'Persisted memory'), 'selected');
  assert.deepEqual(device.mma2_advanced.persistence.ranges, {});
  const addButtons = collect(root).filter(node => node.tagName === 'button' && node.textContent === 'Add range');
  assert.equal(addButtons.length, 2);
  addButtons[0].events.click();
  assert.deepEqual(device.mma2_advanced.persistence.ranges.coils, [{start: 10, count: 1}]);
  const starts = collect(root).filter(node => node.attributes['aria-label'] === 'Start');
  const counts = collect(root).filter(node => node.attributes['aria-label'] === 'Count');
  const text = () => collect(root).map(node => node.textContent).join('|');
  type(starts[0], '12'); type(counts[0], '4'); // 12..16 exceeds destination coils [10..14)
  assert.ok(/Coils range is outside the allocated area\./.test(text()));
  type(starts[0], '11'); type(counts[0], '3'); // 11..14 within
  assert.ok(!/Coils range is outside the allocated area\./.test(text()));
  assert.deepEqual(device.mma2_advanced.persistence.ranges.coils, [{start: 11, count: 3}]);
});
