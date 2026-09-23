# Simulator — device configuration and runtime boundary

Parent: [L0-project](L0-project.md)
Zoom Out: [L0-project](L0-project.md)
Zoom In: [simulator-memory-none](simulator-memory-none.md), [simulator-projection](simulator-projection.md)
Connectors: [osjs-shell](osjs-shell.md) for UI/relay integration; [replicator](replicator.md) for shared MMA2 reservations. Follow connectors only when authorized work crosses these contracts.

## Semantic boundary

The repository-root `simulator/` owns Simulator device definitions and the runtime integration with MMA2. It does not own the OS.js desktop, the Windows Electron installer or the Replicator's polling implementation. The primary source map is `simulator/device.go`, `simulator/store.go`, `simulator/apply.go`, `simulator/runtime_server.go`, `simulator/scheduler.go`, `simulator/raw_ingest.go`, `simulator/advanced_projection.go`, `docs/SIMULATOR_RUNTIME_INTEGRATION.md` and `deploy/docker-compose.yml`. The historical SIM-001–SIM-024 evidence belongs in `workflow/archive/sim-*.md`, not duplicated as twenty-four implementation narratives in this node.

## Established contracts from earlier audited context

- Simulator definitions are separate from effective MMA2 configuration. The host-mounted data root is `OSJS_DATA_DIR`; device definitions were recorded at `$OSJS_DATA_DIR/config/simulator/devices.yaml` and the shared MMA2 effective config at `$OSJS_DATA_DIR/config/mma2/config.yaml`. Verify paths against current runtime source before authorizing a write; do not invent alternate host paths.
- MMA2 reservation identity is `(port, unit_id)`, shared with Replicator. Composition must not overwrite a different owner's reservation. Recognized FC1–FC4 areas map to coils, discrete inputs, holding registers and input registers. Invalid saves must preserve prior persisted data.
- Simulator value generation uses per-FC schedules; `RawIngestClient` writes generated data via MMA2 raw ingest rather than pretending Modbus read-only areas are writable using standard Modbus writes.
- Simulator does **not** own the MMA2 process lifecycle under the superseding design. MMA2 auto-starts independently; Simulator may request a restart only after a valid shared-config change. Earlier SIM-008/009 child-process and HTTP bridge descriptions were superseded by later work and are not current implementation instructions.
- `simulator/advanced_projection.go` contains a projection helper for omitted advanced properties; the narrow inheritance semantics and provenance live in [simulator-projection](simulator-projection.md).
- None-mode-specific runtime and persistence behavior belongs in [simulator-memory-none](simulator-memory-none.md); do not rely on superseded pre-None status or Unix-socket text from historical SIM-018/019.

## Evidence boundary and freshness

These are concise, historically audited contracts, **not a fresh full-source or Linux-runtime audit**. Historical source checkpoint: `c289f2f3872b792c45813ac7a515f8fe90e77671` for the projection-related material. On the reviewed remote `main` snapshot `45cd3d2d81831306c6943f844e49b7deccbfc5ad`, this entire Simulator dependency set was not re-audited; before relying on a detail in a new checkout, compare the affected source dependencies against their audited baseline and refresh only intersecting paths. Local working-tree state and overlay are unknown to this remote-only review.

## Navigation

For a Simulator task, first decide whether it concerns device persistence/runtime, None mode or advanced settings projection. Load only the corresponding node. Read specific authoritative source when this summary lacks the precise function/contract or a relevant delta makes it stale. The older long-form history remains recoverable from Git history and archived SIM task records; no historical task evidence has been deleted.
