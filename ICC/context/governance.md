# Governance — authority and directive routing

Parent: [L0-project](L0-project.md)
Zoom Out: [L0-project](L0-project.md)
Zoom In: none
Source dependencies: `AGENTS.md`, `BLACK_SHEEP_WALL.md`, `operation cwal.md`, `the gathering.md`, `there is no cow level.md`, `workflow/IDENTITY_MAP.md`, `workflow/active_work/README.md`.

## Stable authority boundaries

- `AGENTS.md` routes the StarCraft directives and agent identity. Read the invoked directive's actual file; a cached ICC summary is not authority.
- `BLACK_SHEEP_WALL.md` alone defines ICC maintenance and alone owns ICC writes. **Bootstrap occurs only if no valid baseline exists or the map is genuinely unusable. After bootstrap, direct invocation means delta-driven incremental maintenance, never an automatic full repository audit.** This corrects the superseded rule formerly copied into this context.
- `operation cwal.md` and the active task packet define bounded execution/test permissions. ICC cannot promote tasks, authorize code edits, invent evidence or advance workflow.
- `handoff.md` and canonical workflow files supply continuation state; Planning and Brainstorm do not authorize implementation. A summarized task status belongs in `active-work`, not governance.
- Git/source is final authority for implementation facts; ICC is a cache. A conflict means verify and refresh the smallest affected semantic node, not blindly trust either cached prose or a different branch.

## Maintenance and navigation

Consult the authoritative directive for exact sequencing, delegation and failure rules. Only navigate to a sibling if the authorized question explicitly requires it. Never reproduce full directives, per-task checklists or historical incident timelines in this node. Source baseline of the historical governance node was `ee19b8a`; this concise boundary summary was checked against the `AGENTS.md` and `BLACK_SHEEP_WALL.md` retrieved from remote `main` snapshot `45cd3d2d81831306c6943f844e49b7deccbfc5ad`. Other dependencies have not been re-audited here; compare their delta before treating those details as current. Local overlay state remains unknown.
