# Active Work

Baseline commit: 71d729c34a9a1dfa8cf62f5ff367a059c7e9c8f1
Working tree: clean at audit start.
Source dependencies: handoff.md, operation cwal.md, workflow/adapters/opencode.md,
workflow/active_work/README.md, workflow/active_work/OTR_PROMOTION_QUEUE.md,
workflow/active_work/otr-*.md, workflow/archive/otr-001a-electron-shell-inventory.md,
workflow/archive/otr-001b-electron-editor-inventory.md,
workflow/archive/otr-001c-electron-backend-boundary-inventory.md
Parent: L0-project
Zoom In: replicator, simulator-device-config
Zoom Out: L0-project

## Current OTR execution state

OTR-001A and OTR-001B are COMPLETE. OTR-001C is COMPLETE-UNTRUSTED because its retained
artifact is an IPC-note summary rather than independently usable discovery evidence. OTR-002A is
the sole ACTIVE task and `handoff.md` names its exact packet. OTR-002B, OTR-003A and OTR-003B are
QUEUED. OTR-004 through OTR-014 are NOT EXECUTABLE parent placeholders and require evidence-based
child packets before any product work.

The current CWAL route permits only the handoff-named ACTIVE packet. A successful packet must
write its named evidence, make only its allowlisted status/handoff transition, commit, push
non-force to `origin/opencode`, verify the remote commit, and stop. QUEUED tasks, parents, planning
files and archives are not fallback assignments.

## Evidence integrity

`workflow/archive/otr-002b-baseline-inventory.md` and
`workflow/archive/otr-003a-ui-parity-map-inventory.md` are explicitly invalid and cannot satisfy
gates. They refer to invented paths and were produced by the failed `db48c41` run. The recovery
commits restored executable discovery/design packets but deliberately did not revert all product
or directive damage from that run.

The OTR implementation parents were drafted after the existing UMIG implementation was already
present in this branch. Current Toolkit source already mounts live Memory, Replicator and
Diagnostics editors through `OSJS/src/packages/MCSModbusToolkit/index.js`, with contracts,
transports, relay allowlists, validation, load/apply/status behavior and focused tests. Therefore
OTR-004 through OTR-012 cannot be promoted from their current descriptions without first
reconciling them against implemented source and archived UMIG evidence; several stated outcomes
would duplicate existing work.

## Known branch contamination

Commit `db48c41` added `simulator/devices.yaml` and the typo directive file `operation c wal.md`,
changed handoff/directive content, deleted the original OTR queue, and created invalid completion
artifacts. Recovery commits `eb7b6acc`, `a80ba339` and `71d729c` restored and hardened workflow
routing, but the new YAML and typo directive file remain tracked. Their retention is not an
acceptance decision and must be resolved explicitly before implementation uses this branch as a
trusted product baseline.

## Review conclusion

OTR-002A and OTR-002B remain useful as current-state inventories. OTR-003A and OTR-003B are useful
only if reframed as reconciliation maps between Electron intent, current OS.js implementation and
verified gaps. The numbered implementation parents are roadmap topics, not an executable design.
After discovery, planning should be regrouped around verified deltas and vertical capabilities,
with exact source/test paths and separate TEST/VERIFY gates. Do not discard the repaired routing,
but do not continue the implementation sequence as presently described.

## Scope note

This refresh records current workflow truth only. It does not execute OTR-002A, validate product
behavior, accept or revert `db48c41`, create child packets, promote a successor, or establish any
build/runtime/UI PASS.
