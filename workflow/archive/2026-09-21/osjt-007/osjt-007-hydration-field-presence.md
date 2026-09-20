# OSJT-007 — Hydration field presence

Status: COMPLETE
Stage: CODE
Owner: Coding agent
Previous: OSJT-006
Next: OSJT-008

## Outcome

Repaired the focused OSJT presence fixture without changing production behavior. The fixture now
separately checks omitted, explicit-null, false, and empty advanced-value projection; supplies the
required configuration-only root RBE output; verifies extensions on the matching effective memory;
rejects an invalid recognized state-sealing value with both persisted files unchanged; checks all
relevant errors and lengths; and removes the unused `main()`.

## Source checkpoint

- Revision: `7e3e4fd`
- CODE path: `simulator/osjs_toolkit_settings_test.go`
- Context/workflow paths: `ICC/INDEX.md`, `ICC/context/active-work.md`,
  `ICC/context/simulator-device-config.md`, and the active OSJT-007 record
- No production source, dependency, Docker, service, or external-runtime changes

## Preliminary coding-agent checks

- `go test -count=1 -timeout=90s -v -run '^TestAdvancedProjectionPresence' .` — PASS
- `go test -count=1 -timeout=90s -v -run '^TestAdvancedSettingsRoundTripAndCompose$' .` — PASS
- `git diff --check` — PASS after the separately approved whitespace-only cleanup was preserved
  in a named stash with the unrelated `mma2composer/handoff.md` content

These are preliminary CODE checks, not the independent OSJT-008 verdict.
