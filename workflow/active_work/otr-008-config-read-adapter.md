# OTR-008 — Verified configuration read adapter
Status: PLANNING / UNDER REVIEW. Stage: CODE. Owner: OpenCode. Previous: OTR-007. Next: OTR-009.

Outcome: Implement one source-verified Linux/OS.js read-only configuration adapter for one runtime, selected from OTR-003 mapping. Pin exact contract, schema, fixture and ≤3 production files at promotion; repeat as separate microtasks for other runtimes. No guessed endpoint.

Non-scope: writes, restart, live service actions, broad API refactor. Acceptance: (1) Correctly loads representative synthetic YAML/config preserving fields. (2) Missing/malformed/unavailable source gives explicit errors, not fabricated defaults. (3) Read operation has no write side effects. Evidence: source diff, fixture, targeted commands/exits; independent JR test separately authorized. Dependencies: OTR-003 verified contract and corresponding UI. Size 1/1/1/1/1=5; one tightly coupled read contract only; split per runtime. STOP.
