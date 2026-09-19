# Electron Replicator LEDs

## Boundary

Owns the Electron communication-status contract and Go runtime evidence feeding it.
Does not own OS.js migration or workflow execution state.
Parent / Zoom Out: INDEX.md. No children.
Connector: replicator.md for runtime lifecycle context, which has its own older baseline.

## Source Dependencies

- electron/renderer/comms-status.js
- electron/test/comms-status.test.js
- electron/COMMS_STATUS.md
- replicator/comms.go
- replicator/comms_test.go
- replicator/reader.go
- replicator/config_cycle.go
- replicator/manager.go
- replicator/runtime.go

## Baseline / Overlay

Source baseline: ff6846f.
Audited overlay: new CycleComms observations, observed source/cycle helpers, manager JSON aggregation, renderer legacy-backend message, and loopback tests.
This bounded refresh does not advance unrelated ICC baselines or active-work records.

## Established Facts

Previously the renderer expected comms and per-block layer observations that the Go manager did not emit.
The managed poll cycle now records TCP connection/request activity, validated Modbus responses or exceptions, and acknowledged MMA2 Raw Ingest writes.
Network reachability is established by successful TCP connection, not an ICMP probe; unsuccessful connection leaves network reachability unknown.
Each cycle replaces its observations, so skipped downstream writes cannot reuse prior green evidence.
Aggregate state prioritizes ERROR, WARNING, UNKNOWN, then OK.
The renderer rejects disabled, unavailable, wrong-device, and stopped snapshots; it explains telemetry absence on older backends.
The backend must be rebuilt and restarted to deliver the new schema. Isolated UI review disables backend connections and cannot demonstrate live LEDs.

## Verification

go test -count=1 -timeout=90s ./... in replicator passed.
node --test electron/test/*.test.js passed 32 tests.
These establish runtime loopback evidence and renderer behavior, not installed-service or human visual acceptance.
