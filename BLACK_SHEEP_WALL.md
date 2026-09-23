# BLACK SHEEP WALL

## Purpose and authority

BLACK SHEEP WALL alone maintains `ICC/` as a compact, verified semantic map of repository truth. **Zoom is navigation; delta is maintenance.** Reveal an unknown area once; thereafter scout only newly required territory or source changes. Git is final source authority; ICC is not execution authority. BLACK SHEEP WALL may not select/promote work, grant permissions, change workflow state, implement product code, widen an invocation, or silently trigger another directive. After a delegated refresh, return control to the caller; direct invocation stops after maintenance.

## Semantic model

`ICC/INDEX.md` routes to a semantic tree. Each node has a semantic boundary, parent/Zoom Out, direct children/Zoom In, established facts, unresolved questions, source dependencies and audited baseline/overlay state. Parents summarize children without duplicating their detail. Semantic hierarchy follows ownership, not necessarily filesystem directories. Cross-tree connectors describe verified contracts, dependencies, runtime/data flows or ownership crossings; they are secondary navigation hints, **never automatic imports or authority to enter another branch**.

### Navigation: read-only by default

- **Zoom In:** read only a required child when the current node lacks detail.
- **Zoom Out:** read a parent to understand ownership or boundaries, never above the invoking operation's navigation ceiling.
- **Pan:** enter a sibling only when the question/task explicitly requires that concern; possible relevance is insufficient.
- **Trace:** follow a verified explicit connector only when the question requires that dependency/contract; do not recursively load connected nodes.

```text
QUESTION / TASK -> READ ICC/INDEX.md -> ROUTE TO RELEVANT NODE -> ENOUGH?
  YES -> STOP GATHERING
  NO  -> ZOOM / PAN / TRACE ONLY FOR MISSING DETAIL
  SOURCE -> open only if ICC lacks the required fact, conflicts with source,
            source verification is required, or bounded discovery is authorized
```

Stop as soon as the smallest sufficient semantic view is loaded. Do not gather unrelated context for completeness or future possibilities.

## Bootstrap: only for a missing or genuinely unusable map

If no valid ICC baseline exists or the current semantic map is genuinely unusable, inspect only as broadly as required to construct a navigable semantic model. Build compact nodes, established connectors and index routing; record verified baseline/overlay state, re-read written nodes, verify and STOP. Once bootstrap succeeds, **normal full repository rebuilding is disabled**. An individual missing node does not by itself justify re-bootstrap.

## Incremental maintenance: normal mode

```text
READ CURRENT CHECKOUT BRANCH + HEAD AND THE SELECTED NODE'S AUDITED STATE
-> DETERMINE COMMITTED DIFF + WORKING-TREE DELTA RELATIVE TO AUDITED OVERLAY
-> RESOLVE ONLY RELEVANT SOURCE DEPENDENCIES / CHANGED PATHS
-> NO INTERSECTION? REUSE CURRENT NODE; NO SOURCE INSPECTION OR ICC REWRITE
-> INTERSECTION? INSPECT ONLY CHANGED / NEW / DELETED / RENAMED DEPENDENCIES
-> PATCH THE DEEPEST AFFECTED NODE
-> PROPAGATE ONLY IF PARENT SEMANTIC TRUTH CHANGED
-> UPDATE INDEX ONLY IF ROUTING, STRUCTURE, CONNECTOR OR INDEX STATE CHANGED
-> RE-READ AND VERIFY AFFECTED WRITES AGAINST INSPECTED SOURCE
-> ONLY THEN ADVANCE AFFECTED BASELINE / OVERLAY
-> VERIFY AFFECTED CONTEXT ONLY -> STOP
```

A direct invocation after bootstrap means this delta-driven maintenance, **not a full audit**. An unchanged HEAD still requires checking working-tree changes against the *audited* overlay, not assuming a clean tree. If both deltas are empty, do not reopen source, regenerate context or rewrite the index. A different branch/HEAD is not automatically fresh or stale: compare its relevant dependencies with each node's recorded baseline. Do not copy another branch's status or clean-tree claim into this checkout. When a connector can see remote commits but not a local working tree, record overlay as **unknown/unverified**; never invent `clean` or an overlay hash.

### Bounded missing-territory exception

When the authorized question requires a semantic component never mapped in ICC, its absence can occur **without any Git delta**. This does not invalidate the rest of the map. Identify the smallest required source path/component from the request and existing parent/connector hints; discover only that territory; create one compact node (and only required children/connectors); register its verified parent/source dependencies in the index; re-read and verify the new node. Do **not** explore every unmapped folder, repeat bootstrap, or authorize additional task work. If the required territory cannot be bounded from evidence, report precisely what boundary is missing instead of guessing.

### Dependency mapping precision

Declare the narrowest known material source files or directory patterns for each node. A broad pattern such as `simulator/*` is permitted only if the entire scope really affects that node; otherwise enumerate the relevant exact files. Resolve a pattern against tracked paths to decide whether the delta intersects; **do not inspect unchanged matching files** merely because the pattern is broad. On a rename/deletion, update affected references and source dependencies only after verifying the move. Missing/ambiguous paths are unresolved facts, not invented source. Inspect a descendant or connector only when a changed fact demonstrably affects it.

## Delegated branch refresh

A caller such as CWAL supplies the selected semantic branch as both a navigation and refresh ceiling. Read its ICC context; check whether the committed/overlay delta intersects its declared dependencies; if not, return that context unchanged. If yes, inspect only changed dependencies, patch the deepest affected node and propagate only inside the selected branch as required. Ignore unrelated stale siblings. A caller's ceiling cannot be raised by Zoom Out, Pan, Trace or maintenance convenience. Return to the invoking directive without executing, completing, promoting or reinterpreting the caller's task.

## ICC state and index stability

ICC represents **committed truth at its audited baseline commit + an audited uncommitted working-tree overlay**. Each node records sufficient deterministic source state to detect whether its relevant dependencies changed; per-path content hashes or equivalent fingerprints are preferred for overlays. An index-level baseline does not certify every node at that HEAD: mixed-baseline nodes must be validated individually. Record full SHA when practical, actual branch/snapshot provenance and any unknown local overlay explicitly. Never advance a baseline simply because an ICC file was committed.

`ICC/INDEX.md` is a stable routing/state map, not a task log or a checklist. Update it only for created/deleted/moved/renamed nodes, parent/child routing, verified connectors, or actual index-level state metadata changes. A changed fact within an existing node should not rewrite unrelated routing or index prose. Node target size is a concise semantic boundary: split independently useful concerns into children; line count is a guardrail rather than the only split trigger. Do not duplicate large source bodies, chat history or historical task narratives in ICC.

## Authority, write integrity and failure

If ICC conflicts with authoritative source, mark only the affected node stale, verify the smallest required source scope, and refresh only actual semantic impact; never silently settle ambiguity. Other operations may read ICC or request a bounded Black Sheep Wall refresh but must **never write `ICC/` themselves**.

All writes inherit the `AGENTS.md` Execution Guardrail:

```text
PATCH AFFECTED NODE -> RE-READ ACTUAL WRITTEN CONTENT
-> VERIFY REQUIRED FACTS AGAINST INSPECTED SOURCE DELTA
-> VERIFY ROUTING AND SOURCE DEPENDENCIES IF CHANGED
-> ONLY THEN MARK CURRENT / ADVANCE BASELINE OR OVERLAY
```

If a write is malformed, incomplete or cannot be verified, **do not advance baseline, mark current or propagate uncertain results**. Apply the `AGENTS.md` authoring-failure limit and STOP/report when reached. Never call a generic syntax check equivalent to an authoritative product/build gate.

## Verification and STOP

Bootstrap verifies the semantic hierarchy, routing, source mappings, connectors and baseline. Incremental maintenance verifies only that changed paths map to the right deepest node, affected facts match source, parent propagation is warranted, index changes are structural/state-only, unaffected branches were untouched, and each written node was re-read before state advancement. A structure-only checker may validate references but must not perform discovery or write ICC.

**Invariants:** one ICC writer; Git source authority; tree before connectors; navigation is read-only when context current; after bootstrap maintain only deltas or narrowly required unmapped territory; unchanged files and unaffected nodes cost no source reread; deepest node first; impact-only propagation; stable index; caller ceiling; minimum sufficient context; recorded branch/baseline/overlay truth; verified re-read before advancing state; STOP/RETURN at the invoking boundary.
