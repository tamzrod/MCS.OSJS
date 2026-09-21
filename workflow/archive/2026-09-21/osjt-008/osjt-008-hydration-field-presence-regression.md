# OSJT-008 — Hydration field presence regression

Status: COMPLETE
Stage: TEST
Owner: independent JR under OPERATION CWAL
Previous: OSJT-007
Next: OSJT-009

## Independent verdict

PASS at tested activation HEAD `644a50cc21e984b562e6add1d6bf9c600cbe1274`.

## Evidence

- Target: `/home/sysadmin/apps/MCS.OSJS-jr`, branch `temp-main`
- Source checkpoint `7e3e4fd` is an ancestor of the tested HEAD.
- Local `origin/main` and read-only remote `refs/heads/main` both equaled the tested HEAD.
- Initial and post-test `git status --short` were empty.
- Go environment: `go version go1.26.0 linux/amd64`.
- Exact test:

  ```text
  go test -count=1 -timeout=90s -v -run '^TestAdvancedProjectionPresence' .
  === RUN   TestAdvancedProjectionPresence
  --- PASS: TestAdvancedProjectionPresence (0.00s)
  PASS
  ok github.com/tamzrod/MCS.OSJS/simulator 0.006s
  ```

  Exit code: 0.

- Post-check HEAD remained `644a50cc21e984b562e6add1d6bf9c600cbe1274`;
  `git status --short` remained empty; `git diff --check` exited 0.
- No source edits, report file, commit, push, or unexpected effects occurred during CWAL.

The coding agent reviewed the raw command evidence and accepted the scoped PASS.
