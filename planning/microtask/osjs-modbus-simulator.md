# OS.js Modbus Simulator — Microtask Plan

Status: planning material only. Human promotion is required before execution.

Source intent: `planning/Brainstorm/osjs-modbus-simulator.md`.

The simulator is intentionally decomposed along independently verifiable execution boundaries. MMA2 structural configuration and simulator random-runtime configuration are separate parameter domains even when presented together in one UI.

## SIM-001 — Establish simulator device configuration model

### Primary outcome

A simulator-owned persistent device definition can represent the approved UI parameters while preserving the boundary between MMA2 structural parameters and random-runtime parameters.

### Scope

- Define the simulator device configuration model for Name and Enabled state.
- Define the MMA2 parameter section: Listen Port, Unit ID, and FC1-FC4 Start/Count.
- Define the random-runtime parameter section: FC1-FC4 `Randomize Every (ms)`.
- Store simulator-owned persistent configuration under the designated host-mounted configuration root/convention established by the repository/runtime.
- Add deterministic validation for the simulator-owned fields.
- Keep the effective MMA2 runtime configuration separate from the simulator-owned definition.

### Non-scope

- No OS.js simulator window yet.
- No MMA2 restart/reload integration.
- No random value generation.
- No raw-ingest writes.
- No Replicator implementation.
- No direct editing of effective MMA2 configuration.

### Acceptance criteria

1. One simulator device definition can be created, persisted, loaded, and validated with separate `mma2` and `random_runtime` parameter domains.
2. Persistent simulator configuration is stored under the established host-mounted configuration location rather than inside packaged application/MMA2 source directories.
3. Invalid Port, Unit ID, FC Start/Count, or configured randomization interval is rejected without replacing the last valid persisted definition.

### Verification method

Create one valid definition, persist/reload it, compare the loaded values exactly, then submit representative invalid fields and prove the valid persisted definition remains intact.

### Dependencies

- Approved simulator brainstorm.
- Existing appliance host-mount convention must be located from repository/runtime truth; do not invent a host path.

### Sizing assessment

Size: **3 / 10** — one bounded configuration-model outcome with one persistence/validation verification path.

---

## SIM-002 — Implement MMA2 configuration-intent validation and activation path

### Primary outcome

Simulator MMA2 parameters can be submitted as simulator-owned intent, checked against shared MMA2 ownership, and activated without allowing the simulator UI/runtime to overwrite another producer's MMA2 resources.

### Scope

- Accept the simulator MMA2 parameter domain from a valid simulator definition.
- Validate requested listener/Unit ID/FC ranges against the effective/shared MMA2 namespace using the actual MMA2 addressing model established in repository truth.
- Reject collisions before active MMA2 configuration is changed.
- Compose/update only the simulator-owned contribution to effective MMA2 configuration.
- Restart/reload MMA2 when an accepted structural MMA2 change requires it.
- Preserve the previously active configuration when validation or activation fails.

### Non-scope

- No random scheduler.
- No random value generation/raw ingest.
- No OS.js window.
- No Replicator implementation beyond respecting its future/shared ownership boundary.

### Acceptance criteria

1. A non-conflicting simulator MMA2 definition can be activated and MMA2 exposes the requested listener/Unit/ranges.
2. A conflicting request is rejected before active configuration is replaced.
3. Failed validation/activation leaves the previously working MMA2 configuration operational.

### Verification method

Activate one valid simulator MMA2 structure and verify it through MMA2 runtime behavior; then submit a deliberate collision/invalid structural request and prove the previous active structure remains available.

### Dependencies

- SIM-001.
- Existing MMA2 runtime/config behavior in the repository.
- Shared MMA2 configuration-authority boundary from the approved brainstorm.

### Sizing assessment

Size: **4 / 10** — medium because it crosses configuration composition and MMA2 lifecycle behavior, but they form one tightly coupled activation transaction. Split during promotion if repository inspection shows config composition and lifecycle control are independently unresolved.

---

## SIM-003 — Implement per-FC random simulator runtime

### Primary outcome

The simulator independently schedules FC1-FC4 random value generation using each FC's configured millisecond interval.

### Scope

- Load the simulator-owned `random_runtime` parameters.
- Maintain independent FC1-FC4 schedules.
- Generate boolean values for FC1/FC2 and uint16 values for FC3/FC4.
- Apply timing-only configuration changes to the scheduler without restarting MMA2.
- Expose per-FC runtime timing state sufficient for later UI status: last update and next update.

### Non-scope

- No MMA2 structural config changes.
- No OS.js UI.
- No direct MMA2 memory writes in this task.
- No waveform/ramp/sine/script behavior.

### Acceptance criteria

1. FC1-FC4 can run at different configured millisecond intervals.
2. Each FC generates values of the correct basic type/range on its own schedule.
3. Changing only a randomization interval updates that FC schedule without restarting MMA2 and exposes updated last/next timing state.

### Verification method

Configure visibly different short intervals for the four FCs, observe multiple cycles, verify independent cadence/type behavior, then change one interval and prove only scheduler behavior changes.

### Dependencies

- SIM-001.

### Sizing assessment

Size: **3 / 10** — one bounded simulator-runtime outcome with deterministic scheduler verification.

---

## SIM-004 — Connect random runtime to MMA2 raw ingest

### Primary outcome

Every simulator-generated FC value reaches the configured MMA2 memory exclusively through MMA2 raw ingest and is externally readable through Modbus TCP.

### Scope

- Convert generated FC1-FC4 random values into the actual MMA2 raw-ingest format required by repository/runtime truth.
- Send simulator-generated values only through MMA2 raw ingest.
- Target only simulator-owned MMA2 ranges.
- Verify values arriving through raw ingest are observable through the corresponding Modbus TCP read path.

### Non-scope

- No direct-memory bypass.
- No Modbus-write-based population of simulator values.
- No OS.js UI.
- No structural MMA2 config redesign.

### Acceptance criteria

1. Simulator FC1-FC4 generated values are accepted through MMA2 raw ingest for configured simulator-owned ranges.
2. At least one generated value from each enabled FC area can be observed through the corresponding external Modbus TCP read path.
3. The implementation contains no simulator direct-memory or Modbus-write population bypass.

### Verification method

Run one configured simulator device, capture/identify generated values entering raw ingest, read each enabled FC area through a real Modbus TCP client, and verify the observed values correspond to the simulator update path.

### Dependencies

- SIM-002.
- SIM-003.
- Actual MMA2 raw-ingest contract in repository truth.

### Sizing assessment

Size: **3 / 10** — one integration boundary and one end-to-end data-path verification workflow.

---

## SIM-005 — Build approved OS.js Modbus Simulator window

### Primary outcome

An OS.js Modbus Simulator application presents the approved device-list/editor UI and can create/edit simulator-owned definitions without exposing raw MMA2 configuration.

### Scope

- Add/open the Modbus Simulator as an OS.js application using established repository OS.js application conventions.
- Implement the device list/search surface.
- Implement Add, Duplicate, Delete, Save & Apply, and Discard interactions.
- Implement Device Settings: Name, Listen Port, Unit ID, Enabled.
- Implement FC1-FC4 rows with Start, Count, and `Randomize Every (ms)`.
- Show calculated address ranges.
- Bind the editor to simulator-owned definitions.

### Non-scope

- No raw YAML editor.
- No individual register/coil editor.
- No live memory table.
- No charts/waveforms/scripts.
- No Replicator configuration.
- No direct effective-MMA2-config editing.

### Acceptance criteria

1. The OS.js simulator window opens and matches the approved functional layout: device list plus selected-device editor.
2. All approved device and FC fields can be edited with the two parameter domains preserved behind the UI.
3. Add/Duplicate/Delete/Discard operate only on simulator-owned definitions and do not directly modify effective MMA2 configuration.

### Verification method

Open the application, exercise the device-list/editor operations against simulator-owned definitions, and verify field values round-trip through the configuration model without touching effective MMA2 config directly.

### Dependencies

- SIM-001.
- Existing OS.js application/runtime conventions in MCS.OSJS.

### Sizing assessment

Size: **4 / 10** — medium UI surface but one coherent window/editor outcome. Keep separate from backend activation/runtime behavior to avoid crossing UI and protocol verification workflows.

---

## SIM-006 — Wire Save & Apply to both parameter domains

### Primary outcome

`Save & Apply` validates one edited simulator device, classifies changed parameters, and routes MMA2 changes and random-runtime changes to their correct consumers.

### Scope

- Validate the complete edited simulator definition.
- Detect/classify MMA2 structural changes separately from random-runtime timing changes.
- Route MMA2 structural changes through the SIM-002 activation path.
- Route timing changes through the SIM-003 scheduler update path.
- Ensure timing-only changes do not restart MMA2.
- Surface validation/activation errors while preserving the last active valid state.

### Non-scope

- No new MMA2 config semantics.
- No new random generation behavior.
- No additional simulator UI features.

### Acceptance criteria

1. A structural change follows the MMA2 activation path and restarts/reloads MMA2 only when required.
2. A timing-only `Randomize Every (ms)` change updates the random scheduler without restarting MMA2.
3. A rejected change is surfaced to the user and leaves the previously active valid runtime state intact.

### Verification method

Perform three UI saves: one valid structural change, one timing-only change, and one invalid/conflicting change; verify each follows the expected path and runtime state remains deterministic.

### Dependencies

- SIM-002.
- SIM-003.
- SIM-005.

### Sizing assessment

Size: **3 / 10** — one routing/transaction outcome with three related acceptance cases.

---

## SIM-007 — Add simulator runtime status to the OS.js window

### Primary outcome

The simulator window displays enough runtime feedback to prove MMA2, raw ingest, and each FC scheduler are operating.

### Scope

- Show device runtime status (`RUNNING`, `STOPPED`, or `ERROR`).
- Show MMA2 status and raw-ingest status.
- Show per-FC last random update and next update timing.
- Show total configured points for the selected device.
- Refresh status without turning the UI into a live register viewer.

### Non-scope

- No live register table.
- No charts.
- No historical trending.
- No waveform controls.

### Acceptance criteria

1. The selected device shows current MMA2/raw-ingest/device state.
2. FC1-FC4 show last/next update timing consistent with their independent schedules.
3. Total points reflects the selected device's configured FC counts.

### Verification method

Run a simulator device with different FC intervals, observe status across multiple update cycles, and compare displayed state with runtime evidence from the simulator/MMA2 paths.

### Dependencies

- SIM-004.
- SIM-005.
- SIM-006.

### Sizing assessment

Size: **2 / 10** — one bounded observability/UI outcome using already-established runtime state.

## Planned Execution Order

```text
SIM-001  Configuration model
   ↓
SIM-002  MMA2 structural activation
   ↓
SIM-003  Random runtime scheduler
   ↓
SIM-004  Raw-ingest integration
   ↓
SIM-005  OS.js simulator window
   ↓
SIM-006  Save & Apply routing
   ↓
SIM-007  Runtime status
```

The ordering intentionally proves backend boundaries before wiring the final UI control path. Each task remains planning material until human promotion.
