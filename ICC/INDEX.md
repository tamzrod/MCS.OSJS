# ICC — semantic routing map

ICC is a compact cache of repository truth, not execution authority. Git and the authorized task packet win on conflicts. BLACK SHEEP WALL alone writes ICC. Start here, load the smallest relevant node, stop once sufficient. Zoom is read-only navigation; delta drives maintenance. Never read the registry as a checklist.

## Snapshot and validity

- Last broadly recorded historical baseline: `c289f2f` (from the previous index). **This does not certify the repository or every node at current HEAD.**
- Remote branch reviewed for this repair: `main` at `45cd3d2d81831306c6943f844e49b7deccbfc5ad` (2026-09-23). Branch-local ICC repairs are being prepared separately; this is not an assertion about another checkout's HEAD.
- Remote-only access cannot inspect local uncommitted changes; working-tree status and overlay are **unknown**, not `clean` or `none`. Do not reuse a remote-only snapshot to certify a local overlay.
- [L0-project](context/L0-project.md), [osjs-shell](context/osjs-shell.md), and [active-work](context/active-work.md) contain narrowly inspected facts for the cited `main` snapshot. Other nodes retain their own recorded baselines; **check changed source dependencies before treating them as current**. Some are historical and explicitly describe superseded implementations.
- A context is current for a checkout only if its relevant committed dependencies have no unincorporated diff from that node's audited baseline and its relevant working-tree paths match the audited fingerprints. A changed unrelated branch or unrelated file does not invalidate it. If that proof is unavailable, state `unverified/stale` and inspect the smallest relevant source scope; do not assume either validity or global invalidity.

## Registry: primary semantic tree

| Node | Parent | Direct children | Purpose |
| --- | --- | --- | --- |
| [L0-project](context/L0-project.md) | none | governance, donor-licensing, network-exposure, planning-workflow, osjs-shell, simulator-device-config, replicator, electron-memory-layout, electron-replicator-advanced, electron-replicator-leds, electron-settings-owner, active-work | System identity and boundaries |
| [governance](context/governance.md) | L0-project | — | Directive ownership, not a copy of runtime instructions |
| [donor-licensing](context/donor-licensing.md) | L0-project | — | Provenance and licensing |
| [network-exposure](context/network-exposure.md) | L0-project | — | Network boundaries |
| [planning-workflow](context/planning-workflow.md) | L0-project | brainstorm-topics | Planning contracts |
| [brainstorm-topics](context/brainstorm-topics.md) | planning-workflow | — | Brainstorm index |
| [osjs-shell](context/osjs-shell.md) | L0-project | osjs-toolkit-advanced-status, osjs-shared-mma-settings | OS.js desktop and Toolkit wiring |
| [osjs-toolkit-advanced-status](context/osjs-toolkit-advanced-status.md) | osjs-shell | — | Toolkit editor and status behavior |
| [osjs-shared-mma-settings](context/osjs-shared-mma-settings.md) | osjs-shell | — | Shared MMA settings transactions |
| [simulator-device-config](context/simulator-device-config.md) | L0-project | simulator-memory-none, simulator-projection | Simulator architecture and configuration |
| [simulator-memory-none](context/simulator-memory-none.md) | simulator-device-config | — | None-mode runtime specifics |
| [simulator-projection](context/simulator-projection.md) | simulator-device-config | — | Advanced-settings inheritance |
| [replicator](context/replicator.md) | L0-project | — | Replicator runtime and transport facts |
| [electron-memory-layout](context/electron-memory-layout.md) | L0-project | — | Windows Memory presentation |
| [electron-replicator-advanced](context/electron-replicator-advanced.md) | L0-project | — | Windows Replicator advanced settings |
| [electron-replicator-leds](context/electron-replicator-leds.md) | L0-project | — | Windows telemetry display |
| [electron-settings-owner](context/electron-settings-owner.md) | L0-project | — | Windows installer settings ownership |
| [active-work](context/active-work.md) | L0-project | — | Workflow routing summary only |

Cross-tree connectors are deliberate, not parentage or permission to load a second branch: `osjs-shell` -> `simulator-device-config` (Memory/Simulator contract), `osjs-shell` -> `replicator` (Replicator contract), `osjs-shared-mma-settings` -> `osjs-toolkit-advanced-status` (shared editor wiring). No connector is an automatic import.

## Read route

```text
TASK / QUESTION -> INDEX -> authorized semantic boundary -> smallest useful node
  enough? STOP
  more? ZOOM IN / ZOOM OUT within ceiling; PAN / TRACE only when explicitly required
  stale? BLACK SHEEP WALL refreshes affected dependencies only -> read refreshed node
```

The task defines the navigation ceiling. Do not browse unrelated siblings or expand permissions through a connector. Only consult authoritative repository source where context is absent, insufficient, conflicted, or source verification is required.

## Maintenance route and state integrity

```text
NO VALID MAP -> one scoped bootstrap
VALID MAP -> compare node baseline to checkout HEAD + audited working-tree overlay
NO RELEVANT DELTA -> no source inspection, no regenerated ICC
RELEVANT DELTA -> inspect changed/new/deleted/renamed dependencies only
  -> patch deepest affected node -> propagate only actual semantic changes
MISSING BUT REQUIRED NODE -> narrowly discover only its required source territory;
  register it; do NOT bootstrap or tour the whole repository
```

For each written node, re-read and verify against inspected source before advancing its baseline or overlay state. Record full commit SHA and deterministic fingerprints for audited uncommitted paths. Do not mark an entire index or branch current on the strength of a different node's audit. Update this index only for routing/structure/connector changes or genuine index-level state changes. When source path patterns are ambiguous, resolve the smallest matching tracked paths first; record a missing path as unresolved rather than fabricating a dependency.

## Scope limits

This index does not promote tasks, assign agents, approve merges, establish product tests, or certify live Ubuntu acceptance. See `BLACK_SHEEP_WALL.md` for authoritative maintenance procedure and `handoff.md`/workflow packets for execution. The OTR and OSJT prose in the old index was historical cache, not authoritative current state; task details belong in `active-work` and the actual handoff.
