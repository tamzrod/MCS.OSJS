# Handoff — OSJT-010 hydration projection regression

Status: BLOCKED — packet reconciliation required before JR execution.

## Authoritative task selection

- Queue map: `workflow/active_work/OSJT_QUEUE.md` identifies OSJT-010 as ACTIVE TEST.
- Task: `workflow/active_work/osjt-010-hydration-projection-regression.md` identifies OSJT-010 as ACTIVE TEST, owned by OpenHands / independent JR.
- OSJT-011 is QUEUED successor CODE work. Do not activate or implement it before OSJT-010's independent evidence is reviewed and the transition is separately authorized.
- Prior handoff text contradicted itself by naming OSJT-011 active while retaining OSJT-010's test packet and directing ICC edits/task deletion. Those directions are withdrawn; no task deletion, ICC rewrite, source repair or invented PASS is authorized.

## Known source and test intent (not an executable packet)

- Previous recorded CODE source checkpoint: `c289f2f3872b792c45813ac7a515f8fe90e77671`.
- Target: focused `TestAdvancedProjectionMerge` regression in `simulator`.
- Historical test command: `go test -count=1 -timeout=90s -v -run '^TestAdvancedProjectionMerge' .`.
- Report mode: chat only; no JR repository writes, commits or pushes.
- Previous packet named local checkout `/home/sysadmin/apps/MCS.OSJS-jr`, branch `temp-main`, and required remote main, local tracking ref and HEAD to agree. The remote merge changes the baseline; those values are not presumed current or valid.

## Required authorized reconciliation

The coding/workflow owner must inspect actual independent OSJT-010 test evidence, current branch/worktree, source ancestry and remote HEAD. If already independently PASS, record the real evidence and follow the existing authorized completion/transition gate. Otherwise issue ONE complete current source-pinned OSJT-010 JR packet with exact checkout, branch, source and activation revisions, changed-path allowlist, freshness method, ordered preflight, exact test, expected result, safe post-check, raw evidence, verdict, and report permissions. Do not infer values from this note or run the historical command as an executable packet.

OpenCode may autonomously code and self-test only when an ACTIVE CODE task explicitly assigns it ownership and includes the fields required by `workflow/adapters/opencode.md`. Current OSJT-010 is TEST, not CODE. A failed or blocked gate does not create or authorize a repair task automatically.

STOP: return the missing evidence or packet requirements; preserve task files and ICC unchanged.
