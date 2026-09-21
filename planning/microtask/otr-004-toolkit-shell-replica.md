# OTR-004 — Toolkit shell visual replica
Status: PLANNING / UNDER REVIEW. Stage: CODE. Owner: OpenCode. Previous: OTR-003. Next: OTR-005.

Outcome: Port one bounded slice of Electron Toolkit's outer layout/navigation into unified OS.js Toolkit. Scope: use OTR-003 verified source-to-target map; at most three production files plus tests, exact files to be pinned on promotion. Preserve OS.js desktop, start menu, taskbar and separate deployment. No backend behavior in this task.

Non-scope: editor internals, YAML, Electron runtime imports, service actions, legacy removal. Acceptance: (1) Shell structure matches mapped Electron reference for assigned slice. (2) Existing OS.js launch/navigation remains intact. (3) Targeted rendering/navigation tests demonstrate behavior or report BLOCKED without claiming PASS.

Evidence: source revision, changed-path diff, test commands/exits, visual comparison if actually observed; independent JR TEST/VERIFY must be separately tasked and cannot be inferred from self-tests. Dependencies: OTR-003; if shell requires >3 files split this planning task before promotion. Size 2/1/1/1/1 = 6 → MUST SPLIT at review into smaller CODE slices, each with its own task and test packet. STOP.
