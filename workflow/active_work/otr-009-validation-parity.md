# OTR-009 — Configuration validation parity
Status: PLANNING / UNDER REVIEW. Stage: CODE. Owner: OpenCode. Previous: OTR-008. Next: OTR-010.

Outcome: Port or adapt the Electron reference's verified validation behavior for one configuration type using a synthetic fixture. Exact schema, source/target paths and ≤3 production files pinned from OTR-003 before promotion; separate task per runtime/config type.

Non-scope: saving, restart, schema invention, unknown-field deletion, production data. Acceptance: (1) Valid fixture accepted and invalid fixture rejected with reference-equivalent errors. (2) Unknown fields preserved on round trip where reference/schema permits. (3) Validation failure leaves persisted configuration unchanged. Evidence: diff, fixture cases, commands/exits, source revision; independent JR test separate. Dependencies OTR-003/008. Size 1/1/1/1/1=5 tightly coupled one validator; split distinct serializers or schemas. STOP.
