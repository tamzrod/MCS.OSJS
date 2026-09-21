# OSJT-009 — Hydration projection

Status: COMPLETE
Stage: CODE
Owner: coding agent
Previous: OSJT-008
Next: OSJT-010

## Outcome

Added a bounded advanced-settings projection helper that selects the matching effective memory by
listener port and unit ID, hydrates only omitted policy, state-sealing, RBE, and unknown extension
values, preserves explicit values including non-nil empty maps, and returns a contextual error for
malformed recognized inherited values. Load-path integration remains deferred to OSJT-011.

## Source checkpoint

- Revision: `c289f2f3872b792c45813ac7a515f8fe90e77671`
- CODE paths: `simulator/advanced_projection.go`, `simulator/advanced_projection_test.go`
- No dependency, production data, Docker, service, browser, or external-runtime changes

## Preliminary coding-agent evidence

`go test -count=1 -timeout=90s -v -run '^TestAdvancedProjectionMerge' .` visibly ran the named
test and its three focused subtests and exited 0. This is preliminary CODE evidence, not the
independent OSJT-010 verdict.
