# OSJT-001 — Inventory OpenCode's existing Ubuntu checkout

Status: ACTIVE
Stage: PREP
Owner: OpenCode for read-only inventory; coding agent reviews and advances
Previous: NONE
Next: OSJT-002

## Fixed target and permissions

The human confirmed OpenCode is already running on Ubuntu. Use ONLY its current MCS.OSJS repository checkout. Resolve the checkout root with git rev-parse; do not choose a different machine, create a worktree/clone, use WSL, or request an OpenHands sandbox. This is a read-only inventory of that Ubuntu checkout, not approval for live Docker or operator-data tests.

No sudo, installs, downloads, git fetch/pull/push, service operations, Docker startup, runtime requests or source edits. Existing dirty files must remain untouched. Do not run the cancelled EM-003 packet. This PREP task is executable directly by OpenCode; it does not require a separate JR TEST TASK packet.

## Exact actions

Run this Bash block once from the existing repository checkout. Every probe prints its exit code. A missing tool is an inventory result; continue the remaining probes. If resolving the repository root fails, stop and report that error.

```bash
repo_root=$(git rev-parse --show-toplevel) || exit 1
cd "$repo_root" || exit 1
probe() {
  printf '\nCOMMAND:'
  printf ' %q' "$@"
  printf '\n'
  "$@"
  result=$?
  printf 'EXIT=%s\n' "$result"
  return 0
}
probe pwd -P
probe uname -s
probe cat /etc/os-release
probe git rev-parse HEAD
probe git status --porcelain --untracked-files=all
probe node --version
probe npm --version
probe env GOTOOLCHAIN=local go version
probe test -f OSJS/package.json
probe test -f OSJS/node_modules/.bin/webpack
probe test -f OSJS/node_modules/.bin/osjs-cli
probe test -d OSJS/node_modules/jsdom
probe test -f OSJS/tests/toolkit-memory-relay.test.js
probe test -f OSJS/tests/toolkit-replicator-relay.test.js
probe test -f deploy/verify/compose.yaml
probe git status --porcelain --untracked-files=all
```

Read OSJS/package.json engines and simulator/go.mod plus replicator/go.mod go versions to report required versus observed versions. These reads do not authorize installing or changing them. The current repository OS.js engine is >=10 <17; its Docker build uses Node 16. Go modules currently require 1.25.0. Report mismatches; do not silently change Node or enable automatic Go toolchain downloads.

## Acceptance — exactly three outcomes

1. Report the actual Ubuntu release, absolute checkout path and full source SHA; or report precisely which identity probe failed.
2. Report observed Node/npm/Go versions and every dependency-file probe as PRESENT or MISSING, retaining command exit codes. Missing prerequisites must be explicit, not fabricated and not installed.
3. Return the complete probe output and compare initial/final git status. State that no tests, builds, Docker actions or product changes were performed. List prerequisites needed for OSJT-002/003 separately from live-test setup, which is NOT authorized here.

## Completion and evidence

INVENTORY COMPLETE means every permitted probe was attempted, available identity evidence and missing prerequisites are honestly recorded, and no unapproved mutation occurred. It does NOT mean the application or toolchain passed. Missing Node/dependencies does not make this inventory incomplete; it makes later execution NOT READY. Unavailable checkout, non-Ubuntu target, truncated required evidence or unapproved mutation means BLOCKED/INCOMPLETE.

OpenCode returns the report in chat and stops. It need not edit any file or ask for file-edit permission. The coding agent stores the reviewed report in docs/osjs-toolkit/test-target.md, decides whether OSJT-002 prerequisites are satisfied, and alone updates handoff/task state. If later prerequisites are missing, record the exact gap and do not dispatch the next test until separately resolved. No automatic environment preparation or live-run authorization is implied.

## Size

One outcome: factual read-only Ubuntu inventory. Five dimensions (implementation/environment/behavior/verification/decision): 0/0/0/1/0 = 1. No product files are changed.
