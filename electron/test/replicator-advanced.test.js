const test = require('node:test');
const assert = require('node:assert/strict');
const fs = require('node:fs');
const path = require('node:path');
const vm = require('node:vm');

class Element {
  constructor(tag) { this.tagName = tag; this.children = []; this.dataset = {}; this.events = {}; }
  append(...children) { this.children.push(...children); }
  appendChild(child) { this.append(child); return child; }
  replaceChildren() { this.children = []; }
  setAttribute(key, value) { this[key] = value; }
  addEventListener(key, callback) { this.events[key] = callback; }
  querySelector(selector) { return collect(this).find(node => node !== this && (node.tagName === selector || '#' + node.id === selector)) || null; }
}
const collect = root => [root, ...root.children.flatMap(collect)];

test('Replicator folder tabs retain advanced drafts through save rejection and discard', async () => {
  const root = new Element('div');
  let submitted;
  const document = {querySelectorAll: () => [], createElement: tag => new Element(tag),
    getElementById: id => id === 'replicator-root' ? root : root.querySelector('#' + id)};
  const context = vm.createContext({document, window: {
    mcsMemoryUI: require('../renderer/memory-advanced'),
    mcsComms: {create: () => new Element('div'), update: () => {}},
    mcsDesktop: {replicatorCall: (operation, payload) => {
      if (operation === 'status') return new Promise(() => {});
      submitted = payload;
      return Promise.reject(new Error('test rejection'));
    }}
  }});
  const source = fs.readFileSync(path.join(__dirname, '../renderer/app.js'), 'utf8');
  vm.runInContext(source.slice(0, source.indexOf("document.addEventListener('click'")), context);
  const run = code => vm.runInContext(code, context);
  const click = title => {
    const button = collect(root).find(node => node.tagName === 'button' && node.textContent === title);
    assert.ok(button, title); button.events.click();
  };
  run('replicatorState.document.devices = [blankReplicator(1)]; replicatorState.persisted = clone(replicatorState.document); replicatorState.selected = 0; renderReplicator();');
  click('Advanced Settings');
  assert.ok(collect(root).some(node => node.textContent === 'RBE Rules'));
  click('Access Policy');
  const control = collect(root).find(node => node['aria-label'] === 'Source IP / CIDR');
  control.value = '10.0.0.1, 192.168.0.0/16'; control.events.input();
  click('Device Definition'); click('Advanced Settings');
  assert.equal(run('selectedReplicator().mma2_advanced.policy.rules[0].source_ip[1]'), '192.168.0.0/16');
  await run('saveReplicator()');
  assert.equal(submitted.document.devices[0].mma2_advanced.policy.rules[0].source_ip[0], '10.0.0.1');
  assert.equal(run('selectedReplicator().mma2_advanced.policy.rules[0].source_ip[0]'), '10.0.0.1');
  run('replicatorState.document = clone(replicatorState.persisted); renderReplicator();');
  assert.equal(run('selectedReplicator().mma2_advanced.policy.rules[0].source_ip[0]'), '0.0.0.0/0');
});
