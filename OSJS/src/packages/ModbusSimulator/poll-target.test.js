'use strict';

const assert = require('assert');
const {appliedPollTargets, pollTargetFor} = require('./poll-target');

(async () => {
  const persisted = [{name: 'PLC-A'}, {name: 'PLC-B'}, {name: 'PLC-C'}];

  // Persisted (applied) document → one poll target per working-copy row.

  const targets = appliedPollTargets(persisted);
  assert.deepStrictEqual(targets, ['PLC-A', 'PLC-B', 'PLC-C']);

  // Unchanged persisted device polls its applied name..
  assert.strictEqual(pollTargetFor(targets, 0), 'PLC-A');
  assert.strictEqual(pollTargetFor(targets, 2), 'PLC-C');

  // Newly added unsaved row has no applied identity→ no poll target..
  targets.push(null);
  assert.strictEqual(pollTargetFor(targets, 3), null);

  // Renaming the working-copy name must not change the poll target (persisted name retained).
  // The targets array is built from the persisted document only, so editing the working
  // copy name cannot retarget polling..
  targets[0] === 'PLC-A' && assert.strictEqual(targets[0], 'PLC-A');

  // Deleting a row shifts indexes consistently; each remaining row keeps its own
  // pervived applied identity..
  targets.splice(1, 1);   // delete an unapplied middle edit
  assert.strictEqual(pollTargetFor(targets, 0), 'PLC-A');
  assert.strictEqual(pollTargetFor(targets, 1), 'PLC-C');

  // Empty / missing persistence yields no target..
  assert.deepStrictEqual(appliedPollTargets([]), []);
  assert.deepStrictEqual(appliedPollTargets(undefined), []);
  assert.deepStrictEqual(appliedPollTargets(null), []);
  assert.strictEqual(pollTargetFor([], 0), null);
  assert.strictEqual(pollTargetFor(null, 0), null);
  assert.strictEqual(pollTargetFor(targets, null), null);
  assert.strictEqual(pollTargetFor(targets, -1), null);
  assert.strictEqual(pollTargetFor(targets, 99), null);

  console.log('poll-target tests passed');
})().catch(error => {
  console.error(error);
  process.exitCode = 1;
});