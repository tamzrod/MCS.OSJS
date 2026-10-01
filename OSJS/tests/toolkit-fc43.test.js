'use strict';
const assert = require('assert');
const {JSDOM} = require('jsdom');
const ui = require('../src/packages/MCSModbusToolkit/memory-advanced');
const dom = new JSDOM('<div id="root"></div>');
const root = dom.window.document.getElementById('root');
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
const params = {fc1: {start: 0, count: 16}, policy: {rules: [{id: 'existing', source_ip: ['0.0.0.0/0'], allow_fc: [1, 2, 3, 4]}]}};
ui.mount(root, params, {document: dom.window.document, devices: []});
click('Device Identification');
assert.ok(!('fc43' in params), 'Opening identity must not create overrides');
edit('Product Code', 'MyDevice');
assert.deepStrictEqual(params.fc43, {product_code: 'MyDevice'});
edit('Vendor Name', 'github.com/tamzrod'); edit('Major / Minor Revision', '2.x');
assert.strictEqual(ui.validateIdentity(params), null);
for (const invalid of ['', 'é', 'x'.repeat(245)]) {
  edit('Product Code', invalid); assert.match(ui.validateIdentity(params), /Product Code.*ASCII bytes/);
}
edit('Product Code', 'x'.repeat(244)); assert.strictEqual(ui.validateIdentity(params), null);
click('Use default for Vendor Name'); assert.ok(!('vendor_name' in params.fc43));
click('Use MMA2 Defaults'); assert.deepStrictEqual(params.fc43, {});
click('Access Policy');
assert.strictEqual(root.querySelector('[aria-label="Access"]').value, 'Custom');
assert.deepStrictEqual(params.policy.rules[0].allow_fc, [1, 2, 3, 4]);
edit('FC43 Read Device Identification', true, 'change');
assert.deepStrictEqual(params.policy.rules[0].allow_fc, [1, 2, 3, 4, 43]);
edit('FC43 Read Device Identification', false, 'change');
assert.deepStrictEqual(params.policy.rules[0].allow_fc, [1, 2, 3, 4]);
for (const [mode, codes] of [['Read Only', [1, 2, 3, 4, 43]], ['Write Only', [5, 6, 15, 16]], ['Read/Write', [1, 2, 3, 4, 5, 6, 15, 16, 43]]]) {
  edit('Access', mode, 'change'); assert.deepStrictEqual(params.policy.rules[0].allow_fc, codes);
}
dom.window.close();
console.log('FC43 UI: omission, full/partial overrides, defaults, validation, presets and custom preservation PASS');
