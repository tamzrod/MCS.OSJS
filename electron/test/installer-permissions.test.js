const test = require('node:test');
const assert = require('node:assert/strict');
const fs = require('node:fs');
const path = require('node:path');

const source = fs.readFileSync(path.join(__dirname, '../build/installer.nsh'), 'utf8');

test('installer requires an explicit account and stores the resolved SID for upgrades', () => {
  assert.match(source, /Page custom MCSSettingsPageCreate MCSSettingsPageLeave/);
  assert.match(source, /FileWriteUTF16LE \$1 "\$MCSSettingsAccount"/);
  assert.match(source, /-mode resolve -account-file/);
  assert.match(source, /ReadRegStr \$MCSPreviousSettingsOwner HKLM "Software\\MCS Modbus Toolkit" "SettingsOwnerSID"/);
  assert.match(source, /WriteRegStr HKLM "Software\\MCS Modbus Toolkit" "SettingsOwnerSID" "\$MCSSettingsOwner"/);
  assert.match(source, /A settings owner is required/);
  assert.doesNotMatch(source, /S-1-5-32-545|S-1-1-0|ROD\\rod/i);
});

test('installer targets config only and stops when permission setup fails', () => {
  assert.match(source, /-mode grant -config "\$MCSRuntimeRoot\\config" -owner "\$MCSSettingsOwner" -previous "\$MCSPreviousSettingsOwner"/);
  assert.match(source, /Settings permissions failed:[\s\S]*?Abort/);
  const pkg = JSON.parse(fs.readFileSync(path.join(__dirname, '../package.json'), 'utf8'));
  assert.match(pkg.scripts['dist:win'], /^npm run build:settings-access && /);
});
