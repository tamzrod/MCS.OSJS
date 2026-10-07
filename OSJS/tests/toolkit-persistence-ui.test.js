'use strict';
const assert = require('assert');
const {JSDOM} = require('jsdom');
const ui = require('../src/packages/MCSModbusToolkit/memory-advanced');

// PERSIST-010 self-check: persistence enablement and locked derived RBE state in
// the Advanced Settings UI, with the state-sealing prerequisite explained.
const dom = new JSDOM('<div id="root"></div>');
const doc = dom.window.document;
const root = doc.getElementById('root');
const click = title => {
  const node = [...root.querySelectorAll('button')].find(node => node.textContent === title);
  assert.ok(node, title); node.click();
};
const edit = (label, value, event = 'input') => {
  const node = root.querySelector(`[aria-label="${label}"]`);
  assert.ok(node, label);
  if (node.type === 'checkbox') node.checked = value; else node.value = value;
  node.dispatchEvent(new dom.window.Event(event, {bubbles: true}));
};
const alertText = () => root.querySelector('[role="alert"]')?.textContent || '';

const params = {fc1: {start: 0, count: 16}, fc3: {start: 0, count: 8}, policy: {rules: [{id: 'existing', source_ip: ['0.0.0.0/0'], allow_fc: [1, 2, 3, 4]}]}, rbe: {coils: [{id: 1, name: 'UserRule', start: 0, count: 1}]}};
ui.mount(root, params, {document: doc, devices: []});
click('Persistence');

// Default: persistence off, no error, derived rules shown as system-owned.
assert.ok(!('persistence' in params), 'Opening persistence must not create a persistence block');
assert.strictEqual(ui.persistenceEnabled(params), false);
assert.strictEqual(ui.persistenceError(params), null);
assert.ok(root.textContent.includes('System-owned (locked)'), 'derived rules must be shown as locked');

// Enable persistence while state sealing is disabled: reject + explain.
edit('Enable persistence', true, 'change');
assert.deepStrictEqual(params.persistence, {enabled: true});
assert.strictEqual(ui.persistenceEnabled(params), true);
assert.match(ui.persistenceError(params), /state sealing/i);
assert.match(alertText(), /state sealing/i, 'UI must explain the sealing prerequisite');

// Enable state sealing: the error clears and persistence stays enabled.
click('State Sealing');
edit('Enable state sealing', true, 'change');
assert.strictEqual(ui.persistenceError(params), null);
click('Persistence');
assert.strictEqual(alertText(), '', 'error must clear once sealing is enabled');
assert.strictEqual(ui.persistenceEnabled(params), true);

// Disabling persistence through the same draft/apply control clears it.
edit('Enable persistence', false, 'change');
assert.deepStrictEqual(params.persistence, {enabled: false});
assert.strictEqual(ui.persistenceEnabled(params), false);

// Derived persistence rules: one per present area, ranges from the layout.
assert.deepStrictEqual(ui.derivedPersistenceRules(params), [
  {area: 'coils', start: 0, count: 16, system_owned: true},
  {area: 'holding_registers', start: 0, count: 8, system_owned: true}
]);

// The persistence table is read-only: it contains no editable inputs.
const persistenceInputs = [...root.querySelectorAll('input, select')].filter(node => node.getAttribute('aria-label') !== 'Enable persistence');
assert.strictEqual(persistenceInputs.length, 0, 'locked persistence rules must not be editable');

// User RBE remains editable on the RBE Rules tab.
click('RBE Rules');
const idField = root.querySelector('[aria-label="ID"]');
assert.ok(idField && !idField.disabled, 'user RBE rule ID must remain editable');
edit('Count', 5);
assert.strictEqual(ui.derivedPersistenceRules(params).length, 2, 'user edits do not change the derived projection');

dom.window.close();
console.log('Persistence UI: draft/apply enablement, sealing prerequisite explanation, locked system RBE and editable user RBE PASS');
