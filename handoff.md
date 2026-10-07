# Handoff — MCS.OSJS

## Autonomous routing authority
Human has approved the PERSIST-001..022 persistence roadmap and explicitly promoted **OpenHands to JR DEV** for OPERATION CWAL CODE/DISCOVERY work.

OpenHands modes remain separate:
- **JR DEV** — assigned CODE/DISCOVERY implementation and bounded source transport.
- **Independent JR** — separately assigned TEST/VERIFY only.

## Completed predecessor
**PERSIST-001 — Persistence Configuration Schema**
CODE COMPLETE at `faa33929429a0382a64b78bc773f0074e11cb6b1`.

## Current task — ACTIVE / TRANSPORT RECOVERY
**PERSIST-002 — State Sealing Prerequisite Validation**
Mode / owner: **CODE / OpenHands JR DEV**

Local implementation commit reported:
`62234991a19e323aef9bcecb2161e486ccba474f`

Its first non-force push correctly failed because remote main had advanced with workflow-only commits. The local commit is preserved. This invocation is authorized only to integrate that already-created PERSIST-002 commit onto current remote main and push it non-force.

### Exact recovery authority

1. Start from the existing OpenHands workspace where local `main` contains commit `62234991a19e323aef9bcecb2161e486ccba474f`.
2. Require clean working tree before integration.
3. Run one `git fetch origin main`.
4. Verify fetched `origin/main` is a descendant of `0256c1ff4e94f522279404516081f1dafc162025`.
5. Verify commits on remote main after `0256c1f`, if any, are workflow-only and do not modify these three product/test paths:
   - `mma2composer/memory_validation.go`
   - `mma2composer/persistence_sealing_validation_test.go`
   - `simulator/persistence_config_schema_test.go`
   If any of those three paths changed upstream, BLOCKED / STOP.
6. Rebase the single local PERSIST-002 commit onto fetched `origin/main`:
   `git rebase origin/main`

### Exact conflict rule

A conflict is permitted **only** in:
`workflow/active_work/persist-002-state-sealing-prerequisite-validation.md`

If that one file conflicts:
- resolve it by keeping the **remote/origin-main version** of the task packet;
- do not carry the local task-packet version into the rebased source commit;
- stage that resolution and continue the rebase once.

Any conflict in any other file => abort the rebase and report BLOCKED / STOP. Do not invent a merge.

Expected rebased source commit should therefore change only:
- `mma2composer/memory_validation.go`
- `mma2composer/persistence_sealing_validation_test.go`
- `simulator/persistence_config_schema_test.go`

7. After successful rebase, verify the rebased commit diff contains only those three paths.
8. Push `main` to `origin/main` **non-force**.
9. Verify `git ls-remote origin refs/heads/main` equals the new local HEAD.
10. Report the new pushed SHA and STOP.

### Forbidden in this invocation
- no force or force-with-lease;
- no merge commit;
- no second fetch/rebase attempt;
- no source reimplementation;
- no additional tests/retests;
- no PERSIST-003;
- no ICC edits;
- no unrelated cleanup or workflow advancement.

If the exact local commit is missing, the tree is dirty, remote touched any of the three source/test paths, rebase conflicts outside the one authorized task packet, or push again fails: report the precise blocker and STOP.

## Successor routing
PERSIST-003 remains QUEUED. It may be activated only after GitHub main contains the rebased PERSIST-002 source commit and the workflow owner confirms repository completion.

Only BLACK SHEEP WALL edits ICC.


## Operator loop

For OpenHands, the intended human workflow is:

```text
git pull
OPERATION CWAL
```

JR DEV has standing authority for routine in-scope Git synchronization, commit, non-force push, evidence recording, current-task closure, and selection of the next already-promoted eligible packet for the next invocation.

One invocation still executes exactly one product task. Independent TEST/VERIFY remains a separate mode and cannot be self-certified by JR DEV.
