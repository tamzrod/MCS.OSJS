# OTR-005 — Simulator editor UI replica
Status: PLANNING / UNDER REVIEW. Stage: CODE. Owner: OpenCode. Previous: OTR-004. Next: OTR-006.

Outcome: Port one source-verified Simulator editor UI slice from Electron into OS.js, preserving layout, labels, tabs, input and draft/discard behavior. Source/target paths must be supplied by OTR-003; maximum three production files plus tests. No persistence or runtime effects.

Non-scope: backend, YAML writes, restart, new UI design. Acceptance: (1) Assigned view renders same controls and states as reference. (2) Draft/discard never writes configuration. (3) Targeted UI tests/evidence and explicit unverified differences recorded. Evidence: exact diff, source revision, test commands/exits and observed screenshots if available. Dependencies: OTR-003/004. Size 2/1/1/1/1=6: split by independent view/tab before promotion; each child gets own file and CODE/TEST/VERIFY stages as applicable. STOP.
