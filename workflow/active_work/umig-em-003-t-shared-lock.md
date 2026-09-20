# UMIG-EM-003-T — TEST: Shared writer lock (one OpenCode JR trial)

Status: QUEUED — NOT executable until the coding agent has completed/archived UMIG-EM-003 and activated this TEST with a new exact handoff packet.
Stage / owner: TEST / independent OpenCode V2 local JR on a separate disposable Git worktree, an expressly human-approved ONE-TASK experimental exception to the usual OpenHands runner. OpenHands remains the fallback; no automatic role replacement.
Previous: UMIG-EM-003.
Next: UMIG-EM-003-V (PLANNED; requires separate human promotion and separate safe runtime target).

## Primary outcome
Independently build/test the committed shared Linux lock implementation and Simulator/Replicator writer integrations using ONLY Go unit/fixture/concurrency/race tests on the operator's Legion JR worktree. No Docker, host Modbus/network endpoints, installed config, process restarts or real operator data.

## Acceptance
1. Preflight one clean isolated worktree at the coding-agent-pinned source SHA, Go >=1.25, exact lock tests and unchanged report authority. Local test cache/temp dirs may be written only after explicit command approval; no sudo, auto-approve, interactive daemon, source changes or host services. If environment/checkout is ambiguous BLOCKED before any test.
2. Run only exact commands from the activated `handoff.md` packet: `mma2composer`, `simulator`, `replicator` unit/fixture suites plus named focused concurrency/race tests. Verify lock timeout, process crash release, no nesting hang, foreign ownership and no false success after failed restart; report raw command stdout/stderr, exit codes and whether any dependencies were unavailable. Do not improvise downloads or broaden scope.
3. Report PASS/FAIL/BLOCKED truthfully in the sole authorized handoff report section; no source fix, task advancement, ICC or production access. If any gate fails stop without converting static source inspection into a test PASS.

Exact current packet, fixed preflight, command sequence, expected outputs, evidence and report-write/transport authority are REQUIRED at activation; this queued file is not an executable instruction. Existing `~/opencode-smoke` experimental agent currently says not to invoke Operation CWAL; user must deliberately update its local system text for this single test without weakening shell approval/edit denial. OpenCode approval is not a sandbox; review every shell command individually. Sizing 0/1/0/1/1=3. Never infer live VERIFY/production from Go TEST.
