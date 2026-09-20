# JR TEST TASK — OSJT-010 Hydration projection regression

## Status and authority

- Sole ACTIVE task: `OSJT-010` in `workflow/active_work/osjt-010-hydration-projection-regression.md`.
- Stage: TEST. Owner: independent JR under `OPERATION CWAL`.
- Goal: independently verify matching-memory advanced projection, explicit-value preservation,
  unknown-extension preservation, and malformed recognized-value rejection.
- Source checkpoint: `c289f2f3872b792c45813ac7a515f8fe90e77671`.
- Test-activation revision: current HEAD containing only the OSJT-009 archive, OSJT-010 activation,
  this packet, bounded ICC active-work refresh, and the human-requested active-work queue cleanup
  after the source checkpoint.
- Chat-only report. No source edits, fixes, task advancement, ICC writes, report-file writes,
  commits, or pushes. STOP after returning the verdict and evidence.

## Exact target and safety boundary

- Checkout: `/home/sysadmin/apps/MCS.OSJS-jr`; branch: `temp-main`.
- Local Ubuntu only; temporary Go test data and existing Go build cache are allowed.
- No Docker, services, browser, devices, production/operator data, dependency install/upgrade,
  sudo, destructive action, or external runtime target.
- Clean initial worktree required. Preserve every stash without applying, dropping, or rewriting it.
- Read-only `git ls-remote origin refs/heads/main` is authorized.
- Read `workflow/active_work/OSJT_QUEUE.md` only as a routing/checklist aid; this packet and the
  ACTIVE OSJT-010 task remain execution authority.

## Preflight

Run in order, capturing complete stdout/stderr and exit status separately:

```sh
pwd
git branch --show-current
git rev-parse HEAD
git status --short
git merge-base --is-ancestor c289f2f3872b792c45813ac7a515f8fe90e77671 HEAD
git diff --name-only c289f2f3872b792c45813ac7a515f8fe90e77671..HEAD
git rev-parse origin/main
git ls-remote origin refs/heads/main
go version
```

Expected: exact checkout and branch; empty status; source checkpoint is an ancestor; activation diff
contains only `handoff.md`, `ICC/INDEX.md`, `ICC/context/active-work.md`,
`ICC/context/simulator-device-config.md`, the removed/archived OSJT-009 paths, and the OSJT-010
active task, plus `workflow/active_work/README.md`, `workflow/active_work/OSJT_QUEUE.md`, and removal
of the stale duplicate active-work copies of archived OSJT-002 through OSJT-004; local tracking and
remote main both equal HEAD; Go is available. Any mismatch, unexpected path, unavailable remote
query, or dirty checkout is BLOCKED.

## Exact product test

From `/home/sysadmin/apps/MCS.OSJS-jr/simulator`, run exactly once:

```sh
go test -count=1 -timeout=90s -v -run '^TestAdvancedProjectionMerge' .
```

Expected: exit 0, named test visibly runs, all three focused subtests PASS, and package reports PASS.
Zero matching tests is not PASS. Nonzero or contradictory output is FAIL; do not rerun or substitute.

## Required post-check

Return to repository root and run even after product FAIL:

```sh
git rev-parse HEAD
git status --short
git diff --check
```

Expected: unchanged HEAD, empty status, and exit 0 from `git diff --check`. Report unexpected changes
without cleanup.

## Verdict and report

- PASS only when all preflight, product-test, and post-check expectations are directly confirmed.
- FAIL on contradictory product evidence; BLOCKED on unavailable/unsafe/stale required evidence.
- Return tested HEAD, OS/Go environment, exact commands with raw outputs and exits, post-state,
  verdict, caveats, and unexpected effects in chat only.
