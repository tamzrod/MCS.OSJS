# ICC format v1 — map, knowledge, validity

This file specifies the **format**, not repository facts or workflow authorization. Only BLACK SHEEP WALL writes ICC. Git source and an authorized task take precedence over ICC. The mode-selection and return contract are in `BLACK_SHEEP_WALL.md`.

## Three logical layers, one small on-disk model

1. **Map:** `ICC/manifest.json` is the canonical machine-readable registry: `schema_version`, `root`, `snapshot`, and `nodes[]`. Each node has a unique `id`, Markdown `file`, `parent` (string or null), `children[]`, `connectors[]` (`target`, `relation`), and repository-relative `sources[]`.
2. **Knowledge:** `ICC/context/<id>.md` contains only the boundary, compact established facts/contracts, unresolved questions and links to actual source/evidence. Its legacy parent/zoom header remains readable, but the manifest is authoritative for structural metadata. Do not duplicate Git source, task timelines or the entire parent in a child.
3. **Validity:** Each manifest node has `state` with `baseline` (full 40-hex commit SHA or null), `overlay` (`unknown`, `clean`, or `audited`), `validity` (`unverified`, `partial`, `stale`, or `current`), and `fingerprints` only for an `audited` overlay. Fingerprints are repository-relative path to deterministic full SHA hash. This is *per-node*; no global HEAD claims all nodes current.

`ICC/INDEX.md` is a compact, human-readable routing mirror of `manifest.json`, **not an independent status source or a task log**. Keep its registry table synchronized when changing map topology. The structural checker compares it against the manifest. Its prose may summarize provenance but does not certify validity.

### Example metadata (illustrative; not an actual project node)

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

Valid connector relations: `contract`, `dependency`, `data-flow`, `ownership`. A connector is not parentage or permission to traverse another branch. All source paths are relative to the repository root. `directory/**` means all descendants; more precise file lists are preferred. Do not record an unverified guessed path as fact.

## Migration rule

All existing ICC node identities, Markdown knowledge files and index rows were migrated into the v1 manifest **without re-inventorying the repository**. This migration does not prove historic baselines or local worktree state. Therefore migrated nodes initially have `baseline: null`, `overlay: unknown`, `validity: unverified`. Prior Markdown baseline notes are historic provenance, **not an implicit current-state claim**. The first operation needing a node performs only bounded verification of that node's declared dependencies, checks the actual local branch/HEAD/overlay, and then records a verified full SHA and fingerprints (or clean overlay). Do not audit all nodes just to clear `unverified`. An unknown local overlay must never be rewritten to `clean` based on a GitHub-only snapshot.

When an existing node lacks sufficient knowledge for a question, verify only its relevant source scope; when a required semantic node is absent, reveal only bounded new territory. Unrelated nodes retain state unchanged.

## Write/finalization protocol

For any affected node: (1) establish branch/HEAD and changed dependency paths, (2) stage only the smallest knowledge or map changes, (3) re-read each written node and verify its claims against inspected sources, (4) update its manifest validity **only after verification**, (5) mirror topology changes in `INDEX.md`, (6) run the read-only checker, and (7) publish the coherent batch in one Git commit if explicitly authorized. If interrupted or a check fails, leave impacted state `stale`/`unverified`, do not advance baseline and report the failure. The tool must never silently change a current node to claim success.

A changed source fact can update the knowledge layer without changing map topology; a new verified connector can update the map without rereading unrelated source; changing only validity must not rewrite knowledge. No scheduled trigger or auto-commit is implemented by this format.

## Checker and tests

`python3 scripts/check_icc_map.py` verifies JSON schema, unique IDs/files, reciprocal parent/children, reachability/cycles, connectors, source-path syntax, plausible state claims, index parity and links. It **does not verify that node facts match product source, that the remote branch equals the local HEAD, or that any hash was audited**. That semantic proof belongs to a bounded BLACK SHEEP WALL operation.

`python3 scripts/test_icc_process.py` runs deterministic offline scenarios for REUSE, UPDATE on a changed file, REVEAL of a bounded missing node, stale-node verification, unknown-overlay safety, and malformed metadata. Tests use isolated temporary fixtures; no product or operator data is touched.

Optional efficiency counters belong in an operation's short response: nodes read/written, source paths inspected, and reason for any expansion. Do not maintain an append-only log or require metrics collection for a cache hit.
