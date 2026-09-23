# ICC format v1 — map, knowledge, validity

This specifies the **format**, not repository facts or workflow authorization. Only BLACK SHEEP WALL writes ICC. Git source and an authorized task take precedence over ICC. Mode selection and the short return contract are in `BLACK_SHEEP_WALL.md`.

## Three logical layers, one small on-disk model

1. **Map:** `ICC/manifest.json` is canonical machine-readable routing: `schema_version`, `root`, `snapshot`, `nodes[]`. Each node has a unique `id`, Markdown `file`, `parent` (string or null), `children[]`, `connectors[]` (`target`, `relation`), and repository-relative `sources[]`.
2. **Knowledge:** `ICC/context/<id>.md` contains boundaries, compact established facts/contracts, unresolved questions and source/evidence links. Its legacy parent/zoom/source headers remain readable; the manifest is authoritative for structural metadata once reconciled. Do not duplicate source, task histories or entire parents.
3. **Validity:** Each node's `state` has `baseline` (full 40-hex SHA or null), `overlay` (`unknown`, `clean`, `audited`), `validity` (`unverified`, `partial`, `stale`, `current`), and `fingerprints` only for `audited` overlays. Fingerprints map repository-relative paths to deterministic full hashes. This is per-node; the index-level HEAD certifies nothing globally.

`ICC/INDEX.md` is the concise human routing mirror of the manifest, **not an independent status source or task log**. The checker enforces registry-table parity. The current index snapshot describes provenance, not node validity.

### Example metadata (illustrative, not an actual project node)

```json
{
  "id": "sample-component",
  "file": "context/sample-component.md",
  "parent": "L0-project",
  "children": [],
  "connectors": [{"target": "replicator", "relation": "contract"}],
  "sources": ["sample/component.go"],
  "state": {"baseline": null, "overlay": "unknown", "validity": "unverified"}
}
```

Connector relations: `contract`, `dependency`, `data-flow`, `ownership`. Connectors are not parentage or cross-boundary permission. Source paths are repository-relative; `directory/**` means all descendants; prefer precise actual files. Never guess a nonexistent path.

## Safe one-time metadata migration

All 18 existing node identities, Markdown knowledge files and index routing rows were migrated to the v1 manifest **without repository rediscovery**. Historical baseline notes are provenance, NOT current certification; GitHub remote review could not determine the local overlay. Thus **every migrated node begins with** `baseline: null`, `overlay: unknown`, `validity: unverified`. Do not audit every node simply to clear that label or assume a remote branch has a clean local working tree.

**Crucial dependency-coverage rule:** Migrated manifest `sources[]` are candidate routing seeds, **not guaranteed complete source coverage**. A node can reference additional dependencies in its existing Markdown header or semantic boundary. On the *first request for that node only*, compare its candidate source list with its existing Markdown source-dependency declarations, resolve missing/renamed source paths in the smallest relevant scope, and update that node's manifest `sources[]` to cover all **material** dependencies. Do not mark it `current` until source coverage, node facts, actual checkout HEAD, and audited overlay are verified together. If coverage cannot be established, retain `partial`/`unverified`; never allow a subset of source paths to produce a false REUSE. This is bounded reconciliation, not permission to scout the whole repository.

After successful verification, record a full baseline SHA and deterministic fingerprints (or a verified clean overlay). Verify only a requested node or node intersecting a genuine source delta. Unrelated nodes stay unchanged. Missing required territory is revealed narrowly without restarting bootstrap.

## Coherent write/finalization protocol

For an affected node: (1) observe branch, HEAD and overlay; (2) establish its bounded actual source coverage; (3) stage only necessary knowledge/map changes; (4) re-read written nodes and check claims against inspected source; (5) update manifest validity only **after** verification; (6) mirror changed topology in `INDEX.md`; (7) run the read-only checker; (8) publish one coherent Git commit **only if authorized**. Interruption/failure leaves impacted state stale or unverified; never advance baseline or claim completion. The checker validates syntax and relationships, not semantic truth.

A fact change may affect knowledge without topology; a connector change may affect the map without unrelated source reads; a validity change must not rewrite knowledge. No trigger, scheduler, auto-commit or service operation is implemented here.

## Checker and deterministic tests

`python3 scripts/check_icc_map.py` checks schema, IDs/paths, reciprocal tree, cycles, connectors, source-path syntax, plausible state claims, index parity and links. It does **not** establish that declared sources are complete, facts match source, HEAD equals the local checkout, or hashes were really audited. Those are Black Sheep Wall responsibilities.

`python3 scripts/test_icc_process.py` tests unchanged REUSE, one changed-file UPDATE, missing bounded-node REVEAL, stale-node revalidation, unknown-overlay safety and broken schema on isolated fixtures. They do not prove the full repository ICC or live application.

Optional efficiency counters in the short response: nodes read/written, source paths inspected and reason for expansion. No append-only activity log or mandatory metrics work on cache hits.
