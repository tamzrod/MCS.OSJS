# Electron Replicator LEDs

## Boundary

Owns the Electron communication-status contract and Go runtime evidence feeding it; does not own OS.js migration or workflow execution state.
Parent / Zoom Out: [L0-project](L0-project.md). No children.
Connector: [replicator](replicator.md) for runtime lifecycle context with its own older baseline; follow only for explicitly required contract questions.

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
Audited historical overlay: CycleComms observations, observed source/cycle helpers, manager JSON aggregation, renderer legacy-backend message and loopback tests. This node's routing change does not certify latest checkout source or local overlay.

## Established Facts

Previously the renderer expected comms and per-block layer observations that the Go manager did not emit. The managed poll cycle now records TCP connection/request activity, validated Modbus responses or exceptions, and acknowledged MMA2 Raw Ingest writes. Network reachability is based on TCP connection, not ICMP; an unsuccessful connection leaves network reachability unknown.
Each cycle replaces observations so skipped downstream writes cannot reuse prior green evidence. Aggregate state prioritizes ERROR, WARNING, UNKNOWN, then OK. Renderer rejects disabled, unavailable, wrong-device and stopped snapshots, and explains missing telemetry on older backends. Backend rebuild/restart is needed for the new schema; isolated UI review does not demonstrate live LEDs.

## Historical Verification

go test -count=1 -timeout=90s ./... in replicator passed; node --test electron/test/*.test.js passed 32 tests. This historical evidence is not installed-service or human visual acceptance, nor a fresh test in this branch.
