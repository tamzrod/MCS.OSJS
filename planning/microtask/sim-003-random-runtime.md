# SIM-003 — Implement Per-FC Random Runtime Scheduler

Status: planning material only. Human promotion is required before execution.

Source intent: `planning/Brainstorm/osjs-modbus-simulator.md`.

## Primary Outcome

The simulator runs independent FC1-FC4 randomization schedules using each FC's configured millisecond interval.

## Scope

- Load only the simulator-owned `random_runtime` parameter domain.
- Maintain an independent timer/schedule for each configured FC.
- Generate boolean values for FC1 and FC2.
- Generate uint16 values for FC3 and FC4.
- Track per-FC last-update and next-update timing state.
- Apply a timing-only parameter change without restarting MMA2.

## Non-Scope

- No MMA2 structural configuration changes.
- No shared MMA2 collision validation.
- No raw-ingest transport in this task.
- No OS.js UI.
- No ramps, sine waves, scripts, or manual values.

## Acceptance Criteria

1. FC1-FC4 can run concurrently at four different configured millisecond intervals.
2. Each FC produces values of the correct basic type/range when its own timer fires.
3. Changing one FC interval updates that FC schedule without restarting MMA2 and updates its last/next timing state.

## Verification

Use visibly different short intervals for all four FCs, observe multiple cycles, and record per-FC firing cadence and generated value types. Change one interval and prove only scheduler timing changes.

## Dependencies

- SIM-001.

## Sizing

**3 / 10 — Small.** One scheduler outcome with one deterministic runtime verification path.