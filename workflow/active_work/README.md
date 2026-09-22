# Active Work — OTR

Goal: deliver the familiar Electron Toolkit user workflows in the unified OS.js Toolkit without
duplicating working implementation or importing the Electron runtime.

Root `handoff.md` and `OTR_PROMOTION_QUEUE.md` are the only routing authorities.

## OpenCode loop

```text
ACTIVE task → execute exact packet → record COMPLETE / FAIL / BLOCKED
→ update dependencies and queue → run next eligible task
→ repeat until no eligible task remains → exhaustion summary
```

- COMPLETE activates a dependency-satisfied successor immediately.
- FAIL/BLOCKED never disappears: preserve its evidence and mark dependent work SKIPPED-BLOCKED.
- Continue with independent eligible work instead of stopping the whole invocation.
- QUEUED is not runnable until its dependencies are COMPLETE.
- SUPERSEDED / NOT EXECUTABLE is never runnable and cannot create children.
- Commit/push is optional unless explicitly requested.
- No product task may edit ICC; BLACK SHEEP WALL owns ICC.

Each task remains atomic: finish its evidence and state before moving to another. Never fabricate a
PASS, broaden file scope, retry to force success, or perform destructive/live actions without exact
authority.
