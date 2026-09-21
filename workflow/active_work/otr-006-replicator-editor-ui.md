# OTR-006 — Replicator editor UI replica
Status: PLANNING / UNDER REVIEW. Stage: CODE. Owner: OpenCode. Previous: OTR-005. Next: OTR-007.

Outcome: Port source-verified Electron Replicator editor UI into OS.js in one bounded component slice per execution. OTR-003 supplies exact Electron/OS.js paths; max three production files plus tests. Preserve forms, folder/tab structure, validation-display states and draft/discard without backend writes.

Non-scope: new tab architecture by guesswork, config writes, restart, Electron dependencies. Acceptance: reference-equivalent assigned view; no persistence on draft/discard; targeted UI evidence with remaining parity gaps. Evidence: diff, revision, exact tests/exits, observed screenshots if possible. Dependencies OTR-003/005 shared primitives. Size 2/1/1/1/1=6: split by independent editor view/tab before promotion; create separate independent TEST/VERIFY packets. STOP.
