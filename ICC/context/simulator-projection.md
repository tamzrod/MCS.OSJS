# Simulator Projection Settings

## Boundary

Owns MMA2 listener matching and device-specific settings hydration in the Simulator's advanced-settings context. Handles state_sealing, RBE Rules, and extension value precedence when omitted from a Simulator device definition. No children.

Parent / Zoom Out: `INDEX.md`.
Connector: `simulator-device-config.md` for device metadata used as lookup keys in listener matching.

## Source Dependencies

- simulator/advanced_projection.go
- simulator/advanced_projection_test.go
- simulator/advanced_simulator_test.go
- simulator/listener.go
- simulatormain.js

## Baseline / Overlay

Source baseline: `c289f2f`.

The `projectAdvancedSettings` function hydrates advanced settings omitted by a Simulator device from the matching effective MMA2 memory. Device values always take precedence. Listener matching uses port and unit ID as keys. Extra fields are populated with device-specific values unless explicitly set, avoiding clobbering. State sealing, RBE Rules, and extensions inherit from memory only when not already explicit in the device configuration.

## Facts

- Listener matching: scans all listeners for a match where `listener.Listen.Port == params.Port` AND `memory.UnitID == params.UnitID`.
- Precedence: explicit device values override hydrated values; if device has no value, the hydration logic uses memory defaults unless already set elsewhere in Extra.
- State sealing (`state_sealing`) and RBE rules are populated only when the device lacks an existing value (not already explicit).
- Extensions populate `params.Extra` with missing keys from memory, skipping keys already present in the device's Extra map.
- The hydration applies to Simulator devices; it does not affect Electron Replicator settings ownership which uses a different domain and access policy system.

## Verification

`go test -race -count=1 -timeout=90s ./...` in simulator passes projection tests covering listener matching precedence and hydration of missing fields without affecting explicit device values.
