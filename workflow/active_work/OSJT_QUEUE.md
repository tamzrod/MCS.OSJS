# OS.js Toolkit queue map

## Purpose

This file is the canonical compact route through unfinished OSJT work. It removes the need to infer
sequence, artifact intent, or the next deterministic gate from dozens of nearly identical task files.
The individual ACTIVE task and `handoff.md` remain execution authority.

Current state: OSJT-010 is ACTIVE. OSJT-001 through OSJT-009 are archived. OSJT-011 through
OSJT-059 are QUEUED. There must never be more than one ACTIVE task.

## Small-model execution rules

1. Open only this map, the sole ACTIVE task, its direct predecessor evidence, and the files named by
   that task. Do not scan later task implementations.
2. Treat a scope entry marked **NEW** below as permission to create exactly that file. An unmarked
   missing path is a blocker, not implicit creation authority.
3. CODE tasks must implement the stated outcome and author the exact named regression used by the
   immediately following TEST task. Before checkpointing, run that exact next-test command once as
   preliminary evidence. Do not call it an independent PASS.
4. TEST tasks make no source edits. The coding agent must first provide a source-pinned CWAL packet
   with checkout, clean-state, ancestry, changed-path, freshness, exact-command, post-check, verdict,
   and chat-only reporting rules. Run the product test exactly once.
5. VERIFY tasks require an explicit disposable target, synthetic baseline, exact UI/runtime actions,
   observations, and cleanup in `handoff.md`. If any item is absent, report BLOCKED without guessing.
6. On any failed gate, stop. Do not repair during TEST/VERIFY, skip assertions, broaden scope, install
   dependencies, touch production data, enable network outputs, or activate the successor.
7. After a reviewed PASS/completed CODE gate, archive only the current task, activate only its named
   successor, synchronize `handoff.md`, and commit/push that transition together.

## Existing versus planned artifacts

These task-scoped production files are intentional **NEW** artifacts when first reached:

- `simulator/shared_settings.go` — OSJT-028
- `OSJS/src/packages/MCSModbusToolkit/advanced-editor.js` — OSJT-015
- `OSJS/src/packages/MCSModbusToolkit/advanced-ids.js` — OSJT-023
- `OSJS/src/packages/MCSModbusToolkit/mma-authorization.js` — OSJT-036
- `OSJS/src/packages/MCSModbusToolkit/shared-editor.js` — OSJT-040
- `OSJS/src/packages/MCSModbusToolkit/diagnostics-contract.js` — OSJT-048
- `OSJS/src/packages/MCSModbusToolkit/diagnostics-backend.js` — OSJT-050

Every task-specific `OSJS/tests/toolkit-*.test.js` named below is also a **NEW** artifact when its
paired CODE task is reached. Do not pre-create later artifacts.

## Ordered queue

| Task | Stage | Concrete deliverable | Deterministic gate |
|---|---|---|---|
| OSJT-010 | TEST | Projection helper regression | `cd simulator && go test -count=1 -timeout=90s -v -run '^TestAdvancedProjectionMerge' .` |
| OSJT-011 | CODE | Call read-only advanced projection from canonical load under the existing lock; no writes or recursive lock | Author `TestAdvancedProjectionLoad` |
| OSJT-012 | TEST | Load integration regression | `cd simulator && go test -count=1 -timeout=90s -v -run '^TestAdvancedProjectionLoad' .` |
| OSJT-013 | CODE | Device Definition/Advanced Settings tabs in both editors; preserve drafts; no settings controls | Author `toolkit-advanced-tabs.test.js` |
| OSJT-014 | TEST | Advanced-tab regression | `cd OSJS && node --test tests/toolkit-advanced-tabs.test.js` |
| OSJT-015 | CODE | Ordered access rules and four policy modes; FC checkboxes only for Custom | Author `toolkit-access-modes.test.js` |
| OSJT-016 | TEST | Access-mode regression | `cd OSJS && node --test tests/toolkit-access-modes.test.js` |
| OSJT-017 | CODE | Ordered IPv4/IPv6/CIDR source editor with explicit validation errors | Author `toolkit-access-sources.test.js` |
| OSJT-018 | TEST | Access-source regression | `cd OSJS && node --test tests/toolkit-access-sources.test.js` |
| OSJT-019 | CODE | State-sealing editor preserving enabled=false and zero values | Author `toolkit-state-sealing.test.js` |
| OSJT-020 | TEST | State-sealing regression | `cd OSJS && node --test tests/toolkit-state-sealing.test.js` |
| OSJT-021 | CODE | RBE rule editor; configuration only, no output enablement | Author `toolkit-rbe-rules.test.js` |
| OSJT-022 | TEST | RBE-rule regression | `cd OSJS && node --test tests/toolkit-rbe-rules.test.js` |
| OSJT-023 | CODE | Deterministic globally unique RBE IDs across all memories | Author `toolkit-rbe-ids.test.js` |
| OSJT-024 | TEST | RBE-ID regression | `cd OSJS && node --test tests/toolkit-rbe-ids.test.js` |
| OSJT-025 | CODE | Wire both advanced editors to drafts/save/discard without losing unknown fields | Author `toolkit-advanced-persistence.test.js` |
| OSJT-026 | TEST | Advanced-persistence regression | `cd OSJS && node --test tests/toolkit-advanced-persistence.test.js` |
| OSJT-027 | VERIFY | Synthetic round trip for policy, sealing, RBE IDs, reload, and discard; RBE output stays disabled | Packet-defined UI/runtime evidence and cleanup |
| OSJT-028 | CODE | Revisioned read of shared MMA2 settings; read-only and lock-safe | Author `TestSharedSettingsLoad` |
| OSJT-029 | TEST | Shared-settings load regression | `cd simulator && go test -count=1 -timeout=90s -v -run '^TestSharedSettingsLoad' .` |
| OSJT-030 | CODE | Validate complete shared candidate before mutation | Author `TestSharedSettingsCandidate` |
| OSJT-031 | TEST | Candidate-validation regression | `cd simulator && go test -count=1 -timeout=90s -v -run '^TestSharedSettingsCandidate' .` |
| OSJT-032 | CODE | Reject stale revision without mutation | Author `TestSharedSettingsRevision` |
| OSJT-033 | TEST | Revision-guard regression | `cd simulator && go test -count=1 -timeout=90s -v -run '^TestSharedSettingsRevision' .` |
| OSJT-034 | CODE | Commit valid shared settings and require restart acknowledgment | Author `TestSharedSettingsCommit` |
| OSJT-035 | TEST | Commit/ack regression | `cd simulator && go test -count=1 -timeout=90s -v -run '^TestSharedSettingsCommit' .` |
| OSJT-036 | CODE | Explicit Toolkit authorization for shared MMA management | Author `toolkit-mma-authorization.test.js` |
| OSJT-037 | TEST | Authorization regression | `cd OSJS && node --test tests/toolkit-mma-authorization.test.js` |
| OSJT-038 | CODE | Allowlisted relay for revisioned load/validate/commit only | Author `toolkit-mma-relay.test.js` |
| OSJT-039 | TEST | Shared-settings relay regression | `cd OSJS && node --test tests/toolkit-mma-relay.test.js` |
| OSJT-040 | CODE | Shared MMA dialog with revision display, validation, save, discard, and conflict display | Author `toolkit-shared-dialog.test.js` |
| OSJT-041 | TEST | Shared-dialog regression | `cd OSJS && node --test tests/toolkit-shared-dialog.test.js` |
| OSJT-042 | VERIFY | Two-window stale-write conflict; rejected save must not mutate; restore baseline | Packet-defined UI/runtime evidence and cleanup |
| OSJT-043 | CODE | Map independently observed source/Modbus/MMA-write evidence; no inferred green | Author `toolkit-comms-evidence.test.js` |
| OSJT-044 | TEST | COMMS evidence regression | `cd OSJS && node --test tests/toolkit-comms-evidence.test.js` |
| OSJT-045 | CODE | Render non-color COMMS states and clear stale green | Author `toolkit-comms-leds.test.js` |
| OSJT-046 | TEST | COMMS presentation regression | `cd OSJS && node --test tests/toolkit-comms-leds.test.js` |
| OSJT-047 | VERIFY | Observe success, isolated source failure, truthful layer errors, and recovery | Packet-defined timestamps, observations, and cleanup |
| OSJT-048 | CODE | Bounded Linux diagnostics read contract; no arbitrary paths/commands | Author `toolkit-diagnostics-contract.test.js` |
| OSJT-049 | TEST | Diagnostics-contract regression | `cd OSJS && node --test tests/toolkit-diagnostics-contract.test.js` |
| OSJT-050 | CODE | Report configured listeners and locally observable ownership only | Author `toolkit-diagnostics-ports.test.js` |
| OSJT-051 | TEST | Port-evidence regression | `cd OSJS && node --test tests/toolkit-diagnostics-ports.test.js` |
| OSJT-052 | CODE | Allowlisted bounded log tails with explicit missing/permission errors | Author `toolkit-diagnostics-logs.test.js` |
| OSJT-053 | TEST | Log-evidence regression | `cd OSJS && node --test tests/toolkit-diagnostics-logs.test.js` |
| OSJT-054 | CODE | Problems/Ports/Logs UI, filters, Copy Report; historical logs are not current failures | Author `toolkit-diagnostics-presentation.test.js` |
| OSJT-055 | TEST | Diagnostics-presentation regression | `cd OSJS && node --test tests/toolkit-diagnostics-presentation.test.js` |
| OSJT-056 | VERIFY | Synthetic actionable diagnostics with unavailable host evidence labeled honestly | Packet-defined UI/runtime evidence and cleanup |
| OSJT-057 | CODE | Point desktop shortcut to unified Toolkit; retain legacy packages | Author `toolkit-launcher.test.js` |
| OSJT-058 | TEST | Launcher regression | `cd OSJS && node --test tests/toolkit-launcher.test.js` |
| OSJT-059 | VERIFY | Final rendered unified Toolkit acceptance; legacy apps remain available | Packet-defined UI/runtime evidence; no legacy deletion |

## CODE checkpoint checklist

- Confirm every changed path is named by the ACTIVE task; tests do not count against the three
  production-file limit.
- Inspect the focused diff and ensure the next task's exact named test exists and has nonzero cases.
- Run formatting plus that exact named test as preliminary evidence. Do not substitute a broad suite
  for the named gate.
- Record source SHA, changed paths, commands, exits, and remaining caveats in the archive record.
- If the next named test cannot be authored without another production path or architecture decision,
  stop and repair/split the task before editing outside scope.

## TEST checkpoint checklist

- Confirm clean checkout, sole ACTIVE task, source-checkpoint ancestry, activation changed-path
  allowlist, and current remote-main equality before the product command.
- Capture raw stdout/stderr and exit for each preflight command, the one product command, and all
  post-checks. A named test that does not run is not PASS.
- No source edits, report files, commits, pushes, cleanup, retries, or successor activation during
  CWAL unless the exact packet explicitly authorizes them.

## VERIFY checkpoint checklist

- The packet must identify the disposable target, synthetic fixture, starting revision/state,
  permitted services/actions, evidence capture, expected observations, rollback/cleanup, and final
  state check.
- Missing target credentials/session, ambiguous service ownership, unavailable cleanup, or any risk
  to production/operator data is BLOCKED.
