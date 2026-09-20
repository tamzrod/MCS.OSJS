# Handoff

## Current authority — 2026-09-20

Human-approved sequence: finish `UMIG-EM-003-V` independent live VERIFY with OpenHands, then prepare `UMIG-EM-004` OpenCode coding trial after accepted VERIFY and separate human promotion. Previous Go race TEST `UMIG-EM-003-R-T` is archived PASS; that is not live VERIFY. SOLE ACTIVE: `workflow/active_work/umig-em-003-v-shared-lock.md`. `UMIG-EM-004` is PLANNED; no coding authority. Only BLACK SHEEP WALL updates ICC.

**SOLE OPENHANDS ENTRY POINT: `OPERATION CWAL`.** OpenHands must not be instructed to execute a preparation document or any other workflow directly. Operation CWAL reads this CURRENT JR packet and follows the exact subordinate commands named below. No second directive, new sandbox or new clone is required. This packet authorizes preparation and report delivery only, NOT the live concurrency test.

## JR TEST TASK — CURRENT: UMIG-EM-003-V — CWAL PREPARATION GATE ONLY

GOAL: in the existing disposable OpenHands environment, unblock the missing local Docker daemon once, confirm the verify-only topology and collect the remaining evidence needed for the subsequent exact live VERIFY. This is a PREPARATION stage of the sole ACTIVE VERIFY task, not a product test or a VERIFY PASS.

TARGET / SOURCE: reuse the existing clean `/tmp/em003v-discovery-clone` if present, otherwise the already existing OpenHands checkout. Product source checkpoint `538324a472eb15ff8ef97ee66f826fa02ba9462f` must be an ancestor of the final preparation HEAD. Require a clean checkout and `HEAD == origin/main == git ls-remote origin refs/heads/main` after any strictly authorized fast-forward. The current activation revision is the live GitHub main revision proven by this equality at run time; record its full SHA. No operator/Legion Docker, remote daemon, other deployments or old test volumes.

EXACT COMMANDS / ACTIONS, IN ORDER: Under `OPERATION CWAL`, read `workflow/cwal/umig-em-003-v-sandbox-preflight.md` and execute its numbered steps **1 through 6 exactly once**. That file is incorporated into this packet as the exact command list, NOT a separately invocable directive. It explicitly authorizes ONE clean local-main fast-forward if stale and ONE bounded sandbox-local `sudo -n dockerd` startup if the daemon is absent; it does not permit other Git repairs or privileged operations. No application container start/build, synthetic applies, product tests, production Compose, installation, source edits or live concurrency. If an unsafe prerequisite fails, stop actions and preserve actual BLOCKED evidence; perform only safe post-checks from step 6 where possible.

EXPECTED: source clean/current; sandbox-local daemon accessible with identity; verify-only rendered Compose has no host-network or operator mounts and no public Modbus bind; inventory shows no collisions; local runtime apply/restart interfaces and unknowns reported; final checkout unchanged by preparation (except permitted clean fast-forward). Missing environment or ambiguous provenance = BLOCKED, not product FAIL. Full completion of preparation = `PREPARATION COMPLETE`, **never** `UMIG-EM-003-V PASS`. Interrupted or undelivered mandatory report = `INCOMPLETE` with observed underlying result preserved.

EVIDENCE: command strings, original stdout/stderr or appropriately secret-redacted durable transcript, exit codes, initial/final source SHA/status and live remote, daemon startup result and provenance, rendered verify topology, port/resource inventory, interface file paths and schemas, any side effects. Do not print token-bearing remote URLs. No need to rerun earlier completed product tests or create another sandbox.

REPORT / DELIVERY EXCEPTION: If the checkout is fresh, clean and source-verified after safe post-check, replace ONLY the section `## JR PREPARATION REPORT — UMIG-EM-003-V` below with actual evidence and verdict (including a reportable BLOCKED after a successful freshness check). Preserve all other handoff bytes. Check `git diff --check` and require `git diff --name-only` to equal ONLY `handoff.md`; `git add -- handoff.md`; require `git diff --cached --name-only` to equal ONLY `handoff.md`; `git commit -m 'JR: EM-003-V sandbox preparation report' -- handoff.md`; verify the new commit changes only `handoff.md`. Require live `git ls-remote --exit-code origin refs/heads/main` still equals the tested preparation HEAD immediately before ONE non-force `git push origin HEAD:refs/heads/main`. Re-query live remote and require it equals the new report commit SHA; check clean worktree. No force, no retry, no unrelated commits. If stale/dirty prevents safe report, return evidence IN CHAT and stop; if transport fails, preserve verdict and state INCOMPLETE. Return report commit SHA and STOP. JR cannot archive/promote tasks or authorize live execution.

NEXT (not authorized now): coding agent reviews the pushed preparation evidence and authors a distinct complete live VERIFY packet with exact synthetic requests, bounds, hashes, acknowledgement and label-scoped cleanup. No automatic advance to EM-004.

## JR PREPARATION REPORT — UMIG-EM-003-V

PENDING — OpenHands Operation CWAL has not executed this current packet.
