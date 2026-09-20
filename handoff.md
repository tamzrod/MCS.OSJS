# JR TEST TASK — OSJT-008 Hydration field presence regression

## Status and authority

- Sole ACTIVE task: `OSJT-008` in `workflow/active_work/osjt-008-hydration-field-presence-regression.md`.
- Stage: TEST. Owner: independent JR under `OPERATION CWAL`.
- Goal: independently verify the repaired advanced-projection presence fixture.
- Source checkpoint: `7e3e4fd` (`test: repair OSJT-008 advanced projection fixture`).
- Test-activation revision: the current checkout HEAD containing only the OSJT-007 archive,
  OSJT-008 activation, this packet, and the required bounded ICC active-work refresh after source
  checkpoint `7e3e4fd`.
- No source edits, fixes, task advancement, ICC changes, report-file writes, commits, or pushes.
  Report in chat only and STOP.

## Exact target and safety boundary

- Checkout/worktree: `/home/sysadmin/apps/MCS.OSJS-jr` on branch `temp-main`.
- Local Ubuntu environment only; temporary Go test data and the existing Go build cache are allowed.
- No Docker, services, browser, devices, production/operator data, dependency installation or upgrade,
  sudo, destructive commands, or external runtime target.
- A clean initial worktree is required. Preserve all stashes; do not apply, drop, or rewrite them.
- Read-only remote freshness is authorized with `git ls-remote origin refs/heads/main`.

## Preflight and activation gate

Run each command in order and capture its complete stdout/stderr and exit code:

```sh
pwd
git branch --show-current
git rev-parse HEAD
git status --short
git merge-base --is-ancestor 7e3e4fd HEAD
git diff --name-only 7e3e4fd..HEAD
git rev-parse origin/main
git ls-remote origin refs/heads/main
go version
```

Expected:

- `pwd` is exactly `/home/sysadmin/apps/MCS.OSJS-jr` and branch is `temp-main`.
- `git status --short` is empty.
- The source checkpoint is an ancestor of HEAD.
- `git diff --name-only 7e3e4fd..HEAD` contains only `handoff.md`, `ICC/INDEX.md`,
  `ICC/context/active-work.md`, the archived and removed active OSJT-007 paths, and
  `workflow/active_work/osjt-008-hydration-field-presence-regression.md`.
- `origin/main` and the remote main SHA returned by `git ls-remote` both equal HEAD. Any mismatch,
  unavailable remote query, unexpected path, or dirty checkout is BLOCKED; do not run the product test.
- Go is available. Do not install or download anything if it is unavailable.

## Exact product test

From `/home/sysadmin/apps/MCS.OSJS-jr/simulator`, run exactly once:

```sh
go test -count=1 -timeout=90s -v -run '^TestAdvancedProjectionPresence' .
```

Expected: exit 0, the named test visibly runs, and it reports PASS. Zero matching tests is not PASS.
Any nonzero exit or contradictory output is FAIL. Do not rerun or substitute another test.

## Required post-check

Return to the repository root and run even if the product test fails:

```sh
git rev-parse HEAD
git status --short
git diff --check
```

Expected: HEAD is unchanged from preflight, status remains empty, and `git diff --check` exits 0.
Unexpected changes are reported without cleanup.

## Verdict and report

- PASS only when every preflight item, the exact named test, and every post-check meet expectations.
- FAIL when the product test contradicts its expectation; preserve that result even if a later check fails.
- BLOCKED when freshness, environment, safety, or required evidence is unavailable.
- Return the tested HEAD, OS/Go environment, every exact command with raw output and exit status,
  post-state, verdict, caveats, and unexpected effects in chat only.
- Do not write a report file, commit, push, archive OSJT-008, or activate OSJT-009.
