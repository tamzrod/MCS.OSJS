# Handoff: OpenCode repair for OSJT-008

## Result and authority

- Baseline: `fc755a9` on main. OSJT-008 is FAIL, not complete. Do not advance OSJT-009.
- Human requested fixing OSJT-008 or handing the repair to OpenCode. This is the repair handoff; no product or test source was changed here.
- Next operation: bounded CODE repair of the test fixture, not source editing during TEST. Record the repair in `workflow/active_work/` and reconcile its active status with this handoff before editing. OSJT-007 and OSJT-008 currently still say QUEUED; archive claims are not test evidence.
- This packet supplies the repair scope and checks; no additional human packet is needed for this deterministic local repair. No live deployment, services, Docker, browser, network outputs, dependency upgrades, commit or push.

## Actual failure

Windows local reproduction, Go 1.25.0, at the baseline:

```text
cd simulator
go test -count=1 -timeout=90s -v -run '^TestAdvancedProjectionPresence' .
=== RUN   TestAdvancedProjectionPresence
    osjs_toolkit_settings_test.go:42: failed to save device with advanced settings: validate complete MMA2 candidate: listeners[0].memory[0].rbe: root rbe output is required
--- FAIL: TestAdvancedProjectionPresence (0.03s)
FAIL
FAIL github.com/tamzrod/MCS.OSJS/simulator 1.629s
FAIL
```

Exit code 1. This is not Ubuntu acceptance. The previous instruction to close a verified-incomplete gate is withdrawn: incomplete or failing checks never count as PASS.

## Repair, one step at a time

1. Initially edit only `simulator/osjs_toolkit_settings_test.go`. The fixture enables memory RBE rules without a root output. Reuse the configuration-only fixture in `simulator/advanced_settings_test.go`, `TestAdvancedSettingsRoundTripAndCompose`:

   ```go
   cfg := EffectiveMMA2Config{Extra: map[string]interface{}{
       "rbe": map[string]interface{}{"tcp": map[string]interface{}{"listen": "127.0.0.1:9001"}},
       "custom_root": "keep",
   }}
   ```

   Call `composer.Commit(cfg, OwnershipDoc{})` in the test's temporary store. This is fixture data, not an actual listener. Do not enable automatic RBE output in production or weaken validation.

2. Repair the other incorrect assertions in that same test:
   - Replace the no-op `_ = true` nil-policy check with real separate cases for omitted, explicit null, false sealing and empty advanced values. Assert the representation/projection contract from OSJT-007 without inventing new apply behavior.
   - Unknown `corrupted_rbe` is not a malformed recognized field. Use a genuinely invalid recognized value through validating SaveAndCompose; assert the error and unchanged device/effective file contents. The existing advanced-settings test has an invalid-sealing example.
   - Memory extensions belong in the matching `loaded.Listeners[...].Memory[...].Extra`, not root `loaded.Extra`. Identify the correct memory before asserting.
   - Check all marshal/load/save/read errors and collection lengths before dereferencing. Reading files alone does not establish unchanged contents.
   - Remove the unused `main()`. Do not remove required cases, skip tests, or weaken assertions to get PASS.

3. If real presence cases expose missing OSJT-007 production behavior, report the exact failing cases for a separate bounded CODE repair. Do not expand this fixture task into a production redesign. Fixing the RBE seed alone does not prove presence semantics.

## Ubuntu checks and return packet

Run from OpenCode's repository checkout, capturing each command's output and exit code independently:

```sh
git rev-parse HEAD
git status --short
go version
cd simulator
go test -count=1 -timeout=90s -v -run '^TestAdvancedProjectionPresence' .
go test -count=1 -timeout=90s -v -run '^TestAdvancedSettingsRoundTripAndCompose$' .
cd ..
git diff --check
git diff -- simulator/osjs_toolkit_settings_test.go
```

- Temporary test data only; no operator installation or external runtime target needed.
- If Go/dependencies are unavailable, report the exact environment limitation, not PASS. A later successful command must not hide an earlier failure.
- Return source SHA, changed paths/diff, Go/OS version, full named-test output, exit codes and remaining failures. Zero matching tests is not PASS.
- Coding-agent checks are preliminary evidence. Keep OSJT-008 incomplete until its exact check passes on the repaired revision and its required independent review is recorded. Do not archive or activate OSJT-009 early.
- Preserve the stash named `Preserve local OSJT clarification audit before main update`; do not blindly apply its stale task statuses.
