# OTR-012 — Applied-state and status parity
Status: PLANNING / UNDER REVIEW. Stage: CODE. Owner: ChatGPT coding agent or explicitly assigned OpenCode. Previous: OTR-011. Next: OTR-013.

## Primary outcome
Display independently observed applied configuration and service status in the OS.js Toolkit using Electron-equivalent UI states.

## Scope
Using the verified status contracts from OTR-003 and outcomes of OTR-008 through OTR-011, adapt the existing UI/status logic for ONE selected runtime (Simulator, Replicator or MMA) per promoted execution packet. Pin exact <=3 production paths and targeted test files before promotion. Subsequent runtimes require separate bounded tasks, not silent scope expansion.

## Non-scope
No speculative endpoint, new transport, service control, production data, Electron edits, deployment or claim of success based only on save/restart acknowledgement.

## Acceptance
1. Status is derived from the verified applied-state/runtime observation, not from a successful request alone.
2. Loading, running, stopped, disconnected, stale and error states are shown only where supported by the reference and backend; unknown remains explicit.
3. Targeted synthetic tests demonstrate successful observation, failed observation and stale/unknown status without a false green state.

## Evidence / handoff
Return pinned source revision, changed-path allowlist, targeted test commands and raw outcomes; self-tests are not independent JR PASS. Separate TEST and rendered VERIFY packets must be decomposed before promotion. STOP after this task.

## Dependencies and size
OTR-003, OTR-008..011; size implementation 1, environment 1, behavior 1, verification 1, decision 1 = 5. Split per runtime and stage during review; no blanket three-runtime implementation authorized.
