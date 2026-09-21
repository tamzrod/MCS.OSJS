# OTR-014 — End-to-end acceptance on disposable configuration
Status: PLANNING / UNDER REVIEW. Stage: VERIFY. Owner: OpenHands / independent JR. Previous: OTR-013. Next: none (follow-ups only through new reviewed microtasks).

## Primary outcome
Independently verify ONE selected runtime's OS.js Toolkit workflow on a safe synthetic/disposable target: load → edit → validate → save → restart (only if supported and authorized) → confirm applied configuration and status.

## Scope
Review/promotion must select Simulator, Replicator or MMA, name the exact synthetic YAML fixture, disposable runtime, supported restart method, source revision, preflight, exact ordered actions, expected results, read-only post-check and evidence/report permissions. Restart is not assumed available for every runtime. Create separate microtasks for conflict, invalid input, discard and other independently verifiable failure scenarios.

## Non-scope
No production/operator configuration, global restart, live network devices, dependency changes, product fixes, unapproved service actions or cross-runtime acceptance claims.

## Acceptance
1. Recorded before/after fixture and UI evidence establishes the approved load/edit/validate/save path without silent loss of unknown fields.
2. Where supported, scoped restart is explicitly authorized and observed; unsupported restart is recorded as a contract limitation, never simulated as success.
3. Independent readback/status proves actual applied state or returns FAIL/BLOCKED; acknowledgement alone cannot establish PASS.

## Evidence / handoff
Prior to promotion author an exact source-pinned JR CWAL packet with commands/actions, safety constraints, expected output, raw transcript/screenshots, post-check and report transport. JR executes once, reports verdict and STOPs; no automatic successor or product edit.

## Dependencies and size
OTR-001..013 and independent stage evidence for applicable features. Size implementation 0, environment 2, behavior 2, verification 2, decision 2 = 8: MUST split by runtime and by success/failure workflow before promotion. This file is a review template, not an executable combined task.
