# ICC — semantic routing index

**Start here.** `manifest.json` is the canonical machine-readable map and per-node validity state; this table is its compact human-readable routing mirror. [FORMAT.md](FORMAT.md) defines the schema and three logical layers. [BLACK SHEEP WALL](../BLACK_SHEEP_WALL.md) alone maintains ICC. Git source is authoritative; ICC never grants task execution permission.

## Snapshot and migration safety

The metadata was migrated from existing ICC Markdown **without repository rediscovery**. Historical baseline notes in node Markdown are provenance only: the manifest marks migrated nodes `unverified` with `baseline: null` and `overlay: unknown` until each required node receives its own bounded verification. Do not blanket-refresh these nodes, invent local cleanliness, or treat the snapshot HEAD as proof of any node's validity. The `snapshot` block records the checkout observed for the current maintenance pass (`5a89194e4acc710032c00a26bb5d2a1a095eacd1`, overlay clean at observation); only `governance`, `planning-workflow` and `brainstorm-topics` have been verified against it and set `current`, while every other node remains `unverified`. That pass left only ICC maintenance uncommitted, which is excluded from source dependency matching, so no declared source path is modified.

## Registry — mirror of manifest.json

| Node | Parent | Direct children | Purpose |
| --- | --- | --- | --- |
| [L0-project](context/L0-project.md) | none | governance, donor-licensing, network-exposure, planning-workflow, osjs-shell, simulator-device-config, replicator, electron-memory-layout, electron-replicator-advanced, electron-replicator-leds, electron-settings-owner, active-work | Identity and system boundaries |
| [governance](context/governance.md) | L0-project | — | Directive ownership |
| [donor-licensing](context/donor-licensing.md) | L0-project | — | Provenance and licensing |
| [network-exposure](context/network-exposure.md) | L0-project | — | Network boundaries |
| [planning-workflow](context/planning-workflow.md) | L0-project | brainstorm-topics | Planning contracts |
| [brainstorm-topics](context/brainstorm-topics.md) | planning-workflow | — | Brainstorm index |
| [osjs-shell](context/osjs-shell.md) | L0-project | osjs-toolkit-advanced-status, osjs-shared-mma-settings | OS.js desktop and Toolkit wiring |
| [osjs-toolkit-advanced-status](context/osjs-toolkit-advanced-status.md) | osjs-shell | — | Toolkit editor and status |
| [osjs-shared-mma-settings](context/osjs-shared-mma-settings.md) | osjs-shell | — | Shared MMA settings transaction |
| [simulator-device-config](context/simulator-device-config.md) | L0-project | simulator-memory-none, simulator-projection | Simulator boundaries |
| [simulator-memory-none](context/simulator-memory-none.md) | simulator-device-config | — | None mode and transport |
| [simulator-projection](context/simulator-projection.md) | simulator-device-config | — | Settings inheritance |
| [replicator](context/replicator.md) | L0-project | — | Replicator runtime boundaries |
| [electron-memory-layout](context/electron-memory-layout.md) | L0-project | — | Windows Memory presentation |
| [electron-replicator-advanced](context/electron-replicator-advanced.md) | L0-project | — | Windows Replicator advanced settings |
| [electron-replicator-leds](context/electron-replicator-leds.md) | L0-project | — | Windows status display |
| [electron-settings-owner](context/electron-settings-owner.md) | L0-project | — | Windows installer settings |
| [active-work](context/active-work.md) | L0-project | — | Workflow context only |

Explicit connector relationships reside in the manifest. They are verified navigation hints, not parentage, automatic imports or permission to leave the caller's semantic ceiling.

## Read route

```text
TASK -> INDEX + relevant manifest metadata -> smallest useful node
  CURRENT + NO RELEVANT DELTA -> REUSE -> STOP
  STALE/UNVERIFIED -> bounded UPDATE verification -> return
  REQUIRED UNMAPPED NODE -> bounded REVEAL -> return
  NOT ENOUGH -> ZOOM / PAN / TRACE only within authorized boundary
  NO SAFE SOURCE BOUNDARY OR NO VALIDITY COMPARISON -> BLOCKED
```

Do not browse every node because it is listed. Use only the context necessary to answer the question. Parent summaries do not replace precise child facts. A connector never authorizes an unrelated task. Current workflow authority lives in the actual `handoff.md` and task packet, never in cached ICC prose.

## Maintenance route

For existing verified nodes, compare their **own** source baseline with actual HEAD and working-tree changes with audited overlays; intersect changed paths with that node's material source dependencies. No relevant delta means no source inspection or ICC rewrite. Change only the deepest impacted knowledge node and propagate only genuine semantic impact. Missing required territory can be revealed without Git delta, but never rebuild the entire map. `unverified` migration state is resolved only when a caller needs that specific node; it does not trigger global verification. An ICC-only commit is not a product-source delta.

Update `manifest.json` for map or per-node validity changes. Change this index only for routing/topology changes or genuinely changed index-level provenance. Re-read and verify changed nodes before recording `current`; keep unknown overlays unknown. Run `python3 scripts/check_icc_map.py` for read-only structural/state-claim validation, not semantic source certification. Publishing changes, if authorized, should be one coherent Git commit. No automation is installed.
