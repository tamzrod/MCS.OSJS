# OTR-010 — Safe configuration save adapter
Status: PLANNING / UNDER REVIEW. Stage: CODE. Owner: OpenCode. Previous: OTR-009. Next: OTR-011.

Outcome: Implement one verified OS.js/Linux save contract for one runtime, using only synthetic/disposable fixture targets. Pin actual backend interface, permission boundary, revision/conflict semantics and ≤3 production files at promotion. No production configuration write is authorized by this planning task.

Non-scope: restart, blanket filesystem writes, new transport, unrelated YAML schema changes. Acceptance: (1) Validated draft persists and reloads accurately, preserving supported unknown fields. (2) Invalid/stale/conflicting draft cannot overwrite existing fixture. (3) Failed write preserves original fixture and surfaces error; atomic behavior only claimed if demonstrated. Evidence: diff, disposable paths, exact tests/exits and post-check; independent JR test separately assigned. Dependencies OTR-008/009 and verified save contract. Size 2/1/1/1/1=6: split into bounded backend-write and UI-wiring tasks before promotion, each with own tests. STOP.
