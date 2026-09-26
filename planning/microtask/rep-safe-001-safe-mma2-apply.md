# REP-SAFE-001 — Safe MMA2 Apply

Status: PLANNING / NOT ACTIVE
Purpose: prevent a Replicator Save & Apply from taking down working MMA2 memories when a requested destination port is already in use.

## Safety invariant

Replicator must not mutate the live MMA2 configuration until the proposed destination passes pre-flight validation. If activation later fails for another reason, the previous known-good config must be recoverable.

Do not redesign MMA2. Do not implement one MMA instance per memory. Do not refactor unrelated code.

## REP-SAFE-001A — Trace current Save & Apply path

**Type:** DISCOVERY / READ ONLY

**Goal:** identify the real current path before implementation.

Trace:
1. UI Save & Apply handler.
2. request validation.
3. destination port and Unit ID handling.
4. ownership/port checking.
5. MMA2 config read.
6. first MMA2 config mutation/write.
7. MMA2 restart/activation.
8. result/error returned to UI.

Record exact file + function for each step. Identify any existing ownership/port checker that can be reused. Identify the smallest insertion point immediately before the first config mutation.

**Forbidden:** code changes, config changes, service restart.

**Evidence:** concise call-path map with file/function names and recommended insertion point.

**Stop:** report findings and STOP.

---

## REP-SAFE-001B — Pre-flight before config mutation

**Prerequisite:** REP-SAFE-001A evidence.

**Goal:** reject an unsafe destination before live MMA2 config changes.

Immediately before the first config mutation:
1. Resolve requested destination port and Unit ID.
2. Check whether the destination port is already bound/in use.
3. Distinguish an expected/compatible MMA2 ownership case from a conflicting owner using existing ownership logic where possible.
4. Reject an incompatible destination/Unit-ID conflict.

Required ordering:

```
validate request
-> pre-flight destination
-> only if safe: mutate/write MMA2 config
```

On pre-flight failure:
- return clear APPLY_FAILED/conflict information;
- do not modify MMA2 config;
- do not restart or stop MMA2;
- do not disturb existing memories.

**Do not implement rollback in this task.**

**Test evidence:** intentionally occupy the requested port and prove:
- Save & Apply fails;
- MMA2 config is byte-for-byte unchanged;
- MMA2 was not restarted;
- existing MMA2 remains running.

**Stop:** return changed files + evidence and STOP.

---

## REP-SAFE-001C — Last-known-good rollback

**Prerequisite:** REP-SAFE-001B PASS.

**Goal:** recover from activation failures that pre-flight cannot predict.

Before candidate config mutation, preserve the current known-good config. Then:

```
preserve old config
-> write candidate safely
-> activate candidate
   -> success: candidate remains current
   -> failure:
        restore old config
        reactivate old config
        return APPLY_FAILED
```

Rollback success must never turn the failed Replicator apply into success.

**Evidence required:**
- successful apply;
- forced activation failure;
- previous config restored;
- MMA2 successfully running the restored config.

**Stop:** return changed files + evidence and STOP.

---

## REP-SAFE-001D — Precise existing-UI error states

**Prerequisite:** REP-SAFE-001C PASS.

**Goal:** report the safety outcome without redesigning the UI.

Use the existing Save & Apply status/error area to distinguish:

1. **PRE-FLIGHT CONFLICT** — destination rejected; MMA2 config was not changed.
2. **ACTIVATION FAILED / RECOVERED** — candidate failed; previous config restored and reactivated.
3. **ACTIVATION FAILED / RECOVERY FAILED** — candidate failed and automatic recovery failed; operator intervention required.

No unrelated UI work.

**Evidence:** demonstrate each reachable state.

**Stop:** return changed files + evidence and STOP.

---

## REP-SAFE-001E — Regression verification

**Prerequisite:** REP-SAFE-001D PASS.
**Type:** TEST ONLY

Start with at least one known-working MMA2 memory.

### Case 1: occupied destination port

Attempt Replicator Save & Apply to an intentionally conflicting port.

Expected:
- apply fails;
- live MMA2 config unchanged;
- no MMA2 restart for the rejection;
- MMA2 remains running;
- previous memory remains operational.

### Case 2: activation failure after successful pre-flight

Force a safe test activation failure after pre-flight.

Expected:
- candidate activation fails;
- previous config restored;
- previous config reactivates;
- previous memory remains/becomes operational;
- result remains APPLY_FAILED and reports recovery.

**Forbidden:** production config damage, unrelated code changes, architecture changes.

**Evidence:** exact test actions/commands and observed results.

**Stop:** report verdict and STOP.

## Entire-series out of scope

- one MMA2 instance per memory;
- MMA2 process-management redesign;
- Replicator architecture redesign;
- simulator changes;
- unrelated cleanup/refactoring;
- unrelated configuration behavior.

If any child appears to require an out-of-scope change, STOP and report BLOCKED rather than expanding scope.
