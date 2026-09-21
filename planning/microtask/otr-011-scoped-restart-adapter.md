# OTR-011 — Scoped runtime restart adapter
Status: PLANNING / UNDER REVIEW. Stage: CODE. Owner: OpenCode. Previous: OTR-010. Next: OTR-012.

Outcome: Connect one runtime's existing, verified Linux restart mechanism to the corresponding Toolkit action. Identify actual service ownership, permitted target and command/API from source first; if unavailable, report BLOCKED and create discovery task rather than inventing one. ≤3 production files plus tests per promoted slice; one runtime only.

Non-scope: blanket restart, production service action during coding/tests, config writes, unsupported service control. Acceptance: (1) Only selected runtime targeted in disposable/mock environment. (2) Request acknowledgement and failure shown accurately. (3) No restart occurs on failed validation/save or mere draft edit. Evidence: verified endpoint citation, diff, mock/disposable test exits and side-effect post-check; independent JR separately tasked. Dependencies OTR-003/010. Size 2/2/1/1/1=7: split mechanism discovery from implementation and per-runtime adapters before promotion. STOP.
