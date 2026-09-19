'use strict';

// JR-only deterministic DOM/lifecycle smoke: no browser, Docker or network.
const assert = require('assert');
const {createDiagnosticsEditor} = require('../src/packages/MCSModbusToolkit/diagnostics-editor');
class Node {
  constructor(tag) {this.tag = tag; this.children = []; this.attributes = {}; this.handlers = {}; this.textContent = ''; this.disabled = false; this.style = {};}
  append(...nodes) {this.children.push(...nodes);}
  appendChild(node) {this.children.push(node); return node;}
  replaceChildren(...nodes) {this.children = nodes;}
  setAttribute(name, value) {this.attributes[name] = value;}
  addEventListener(name, handler) {this.handlers[name] = handler;}
}
const doc = {createElement: tag => new Node(tag), createTextNode: text => {const node = new Node('#text'); node.textContent = text; return node;}};
const all = node => [node, ...node.children.flatMap(all)];
const text = node => all(node).map(entry => entry.textContent).join(' ');
const settle = () => new Promise(resolve => setImmediate(resolve));
const memory = {load: async () => ({document: {devices: [{name: 'Sim-D'}]}}),
  status: async () => ({status: {name: 'Sim-D', mma2_status: 'RUNNING', device_status: 'IDLE'}})};
let offline = false;
const replicator = {load: async () => ({document: {devices: [{name: 'Rep-D'}]}}),
  status: async () => offline ? Promise.reject(new Error('replicator socket unavailable')) : ({
    name: 'Rep-D', enabled: true, running: true, source_status: 'ERROR', last_error: 'source down',
    blocks: [{index: 0, running: true, source_status: 'ERROR', last_poll: '2026-09-19T00:00:00Z', last_error: 'connection refused'}]
  })};

(async () => {
  const root = new Node('root');
  const pending = [];
  let cancels = 0;
  const editor = createDiagnosticsEditor(doc, root, memory, replicator,
    {setTimer: fn => {pending.push(fn); return fn;}, clearTimer: () => {cancels++;}});
  await settle();
  const buttons = all(root).filter(node => node.tag === 'button');
  assert.deepStrictEqual(buttons.map(node => node.textContent), ['Start runtimes', 'Stop runtimes', 'Refresh observations']);
  assert(buttons[0].disabled && buttons[1].disabled && !buttons[2].disabled);
  assert(!buttons[0].handlers.click && !buttons[1].handlers.click);
  assert.match(text(root), /Sim-D \(first of 1 canonical device/);
  assert.match(text(root), /Memory simulation:.*IDLE/);
  assert.match(text(root), /Rep-D \(first of 1 canonical device/);
  assert.match(text(root), /Replicator source:.*ERROR/);
  assert.match(text(root), /connection refused/);
  assert.match(text(root), /MMA2 service:.*UNKNOWN/);
  assert.match(text(root), /Binary folder:.*UNAVAILABLE/);
  assert.match(text(root), /not runtime service logs/);
  assert.strictEqual(pending.length, 1);
  console.log('read-only donor controls, canonical per-device data and no fabricated global health: checked');

  offline = true;
  buttons[2].handlers.click();
  assert.doesNotMatch(text(root), /connection refused/); // stale evidence cleared immediately.
  await settle();
  assert.match(text(root), /Replicator device runtime:.*UNAVAILABLE/);
  assert.match(text(root), /replicator socket unavailable/);
  assert.match(text(root), /Memory simulation:.*IDLE/);
  editor.destroy();
  assert.strictEqual(root.children.length, 0);
  assert(cancels >= 1);
  assert.strictEqual(pending.length, 2);
  console.log('independent outage, stale-clear, interval disposal and disabled native actions: checked');

  let release;
  const delayed = {load: () => new Promise(resolve => {release = resolve;}), status: memory.status};
  const lateRoot = new Node('root');
  const late = createDiagnosticsEditor(doc, lateRoot, delayed, replicator,
    {setTimer: fn => {pending.push(fn); return fn;}, clearTimer: () => {cancels++;}});
  await settle(); // load is now pending.
  late.destroy();
  release({document: {devices: [{name: 'Sim-D'}]}});
  await settle();
  assert.strictEqual(lateRoot.children.length, 0);
  assert.strictEqual(pending.length, 2); // no polling resumed after destroy.
  console.log('pending request cannot repaint or restart polling after window teardown: checked');
  console.log('UMIG-006 Diagnostics editor checks complete');
})().catch(error => {console.error(error); process.exitCode = 1;});
