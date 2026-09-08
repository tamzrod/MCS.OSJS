# SIM-003 — Implement Per-FC Random Runtime Scheduler

Status: DONE — completed+verified 2026-09-07 (CWAL).

Source intent: `planning/Brainstorm/osjs-modbus-simulator.md`.

## Completion evidence (recorded by CWAL 2026-09-07)

- Scheduler implemented in `simulator/scheduler.go` (committed `1662227`(: per-FC independent schedule map,`math/rand` value generation, per-FC `last`/`next` timing state,`Stop()` amid `run()` loop with `changed` wake channel,and `Timing()` snapshot surface for SIM-007 runtime status. `UpdateTiming(RandomRuntimeParams)` recomputes only the affected FC's next deadline from now,preserves existing last-update state,and never touches MMA2 (no restart,no config writes(.
- Consumed domain: only the simulator-owned `random_runtime` interval domain(FC1IntervalMS..FC4IntervalMS),loaded from `DeviceDefinition.RandomRuntime`;FC area count is read once from the device's MMA2 parameters merely to size generated batches. FC1/FC2 generate boolean coils values; FC3/FC4 generate uint16 register values within 16-bit range.
(FC1=coils, FC2=discrete inputs, FC3=holding registers;, FC4=input registers,(.
- Type/range contract matched repository MMA2 mnemonic mapping from SIM-002A (`memoryFromMMA2Params()`: FC1→coils, FC2→discrete_inputs, FC3→holding_registers, FC4→input_registers;unconfigured count 0 FC areas get count 0 schedule entries and never fire(.
- Verification (`cd simulator; export GOCACHE/GOPATH under /tmp; PATH=/tmp/go-1.22.2/bin`;`go test -count=1 -v`(: 
  - `TestSchedulerRunsAllFCsConcurrentlyAtOwnCadence` PASS (0.12s(: device(5ms,10ms,0, 0) with FC1/FC2 counts configured;observed FC1 fired 8+ times and FC2 4+ times in ~105ms with cadence ratio ~2x (5ms vs 10ms(; FC3/FC4 (count 0; never fired( — proving concurrent per-FC cadence and correct value-type paths(.
  - `TestSchedulerIntervalChangeUpdatesScheduleWithoutRestart` PASS (0.03s(: waits all four FCs to fire,then `UpdateTiming` only FC3 9ms→2ms;proved FC3 next deadline recomputed to after the old one while FC1/FC2/FC4 schedules retained (no MMA2 restart;`Stop()` clean(.
  - `go vet ./...` clean. Full package: `go test -count=1 ./...` → `ok github.com/tamzrod/MCS.OSJS/simulator 0.162s`.
- Ownership/lifecycle non-impacts: scheduler touches no MMA2 files(config/owners/memory/process(;no ownership boundary weakened.no raw-ingest transport,no OS.js UI,no ramps/sine/scripts/manual values (per Non-Scope..
- This closure commit: `handoff.md` + this task file as SIM-003 DONE`. SIM-004 promoted CURRENT per handoff dependency order.

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