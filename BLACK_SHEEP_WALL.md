# BLACK SHEEP WALL — process v1

## Purpose, ownership and stop boundary

BLACK SHEEP WALL alone maintains `ICC/` as a compact, verified semantic map of Git repository truth. **Zoom navigates; delta maintains; reveal only required unmapped territory.** It provides knowledge, not task selection, execution permission, implementation, test acceptance or workflow advancement. A delegated call returns to its caller; direct invocation ends after maintenance. No scheduler, Git hook, background watcher or automatic task execution is installed by this directive.

The machine-readable map, Markdown knowledge and per-node validity schema are specified in `ICC/FORMAT.md` and `ICC/manifest.json`. `ICC/INDEX.md` is the concise routing mirror. Git is the final source authority; a current task/handoff is the execution authority. Other directives may read ICC but may never write it.

## Deterministic entry and decision sequence

1. Resolve the **actual** checkout branch, HEAD and local working-tree overlay if accessible. Read `ICC/INDEX.md` and the relevant manifest node metadata; never assume a GitHub remote view establishes local cleanliness. A different branch does not inherit validity by name. Record an unavailable overlay as `unknown`, never `clean`.
2. Resolve the calling operation's semantic boundary. On a direct invocation after bootstrap, identify only nodes intersecting the committed and working-tree deltas; do not visit every `unverified` node for completeness. A delegated invocation stays inside the caller's navigation and refresh ceiling, including when following connectors.
3. Classify only the required node/territory using the routes below. The pure helper `scripts/icc_decision.py` implements the four-way routing decision given **observed** paths and per-node comparison facts; it performs no Git inspection, source reads or writes. Never invent changed paths, a baseline comparison or a clean overlay to satisfy its inputs.

```text
NO USABLE ICC MAP?              -> BOOTSTRAP (one-time, bounded to usable semantic model)
REQUIRED NODE DOES NOT EXIST?   -> REVEAL if source territory can be bounded; else BLOCKED
EXISTING NODE UNVERIFIED/STALE/PARTIAL? -> UPDATE: verify only its declared dependencies
EXISTING NODE CURRENT BUT BASELINE/OVERLAY NOT COMPARABLE? -> BLOCKED: report missing comparison
RELEVANT COMMITTED OR WORKTREE DELTA? -> UPDATE only intersecting dependencies
CURRENT + NO RELEVANT DELTA?    -> REUSE: no source reads and no ICC writes
```

When nothing is required outside the selected boundary, do not inspect it. If direct invocation finds no relevant source delta and no selected stale node or new territory, stop even if unrelated nodes are marked unverified. An ICC-only maintenance commit is not a product-source delta: exclude `ICC/**` and the checker from source dependency matching unless the task is governance about those exact sources. Checking Git paths/status is not a semantic repository inventory.

### REUSE — cheapest successful outcome

Return the smallest relevant current semantic node. Start with `INDEX.md`, zoom in only when its facts are insufficient, zoom out only within the caller ceiling; PAN and TRACE require an explicitly needed concern or verified connector. Never recursively load connected nodes. **Stop gathering as soon as sufficient context exists.** REUSE reads no source and rewrites no ICC; it is a complete success, not a skipped audit.

### UPDATE — changed or non-current territory

Compare each affected node's recorded full baseline with current HEAD and compare working-tree changes to its **audited** overlay. Compute changed/new/deleted/renamed repository paths first; intersect with declared source dependencies. Narrow path patterns to exact material files where possible. Open only changed dependencies, except that a legacy `unverified` node with no verifiable baseline may require a **one-time bounded** read of its own declared dependencies to establish a baseline. Never turn this into a repository-wide audit.

Patch the deepest affected knowledge node first. Propagate to a parent only if its summary truth changes; inspect descendants/connectors only when demonstrated impact requires them. If only validity changes, do not rewrite the Markdown summary; if only topology changes, update map metadata and the index mirror, not unrelated knowledge. Record missing/ambiguous paths as unresolved rather than inventing them.

### REVEAL — required territory missing, even without Git delta

Locate the smallest previously unmapped semantic boundary from the actual question, parent and connector hints. Identify verified source paths, create only the required compact node and established connectors, register its parent and source ownership, and verify it. An absent node never triggers another full bootstrap. If the source boundary cannot be established, BLOCKED with the precise missing fact. Do not scout every unexplored directory.

### BOOTSTRAP — exceptional

Only when there is no valid semantic map, or the map is genuinely unusable, inspect broadly enough to build one navigable hierarchy. Record verified facts and source dependencies, write compact nodes, establish routing and audited state, re-read and validate before declaring success; stop. Once established, broad rebuild is disabled as normal maintenance.

## ICC structure, state and navigation invariants

The semantic tree is primary; connectors represent verified contract/dependency/data-flow/ownership crossings and never grant authority. Each node must declare its semantic boundary, parent, direct children, source dependencies, facts and unresolved questions. Parents summarize children instead of repeating detail. Historical task narratives and verbatim source do not belong in nodes. Split independently navigable concerns only when the current use case requires it; line count is a guardrail, not a target to fill.

A node is `current` only after its relevant committed sources are verified against its **full commit SHA baseline**, and relevant uncommitted paths match a recorded clean overlay or deterministic per-path audited fingerprints. The index-level reviewed HEAD is provenance, never blanket validity. A remote-only snapshot cannot certify a local overlay. During v1 migration existing Markdown remains intact and node validity stays `unverified` until its own bounded source verification is performed. Do not promote all nodes merely because metadata is syntactically valid.

## Verified update and interruption protocol

```text
OBSERVE BRANCH/HEAD/OVERLAY -> SELECT BOUNDARY + MODE
-> INSPECT SMALLEST REQUIRED SOURCE SCOPE (UPDATE/REVEAL ONLY)
-> STAGE ONLY AFFECTED KNOWLEDGE/MAP CHANGES
-> RE-READ WRITTEN NODES AND VERIFY CLAIMS AGAINST INSPECTED SOURCE
-> UPDATE AFFECTED MANIFEST STATE ONLY AFTER FACT VERIFICATION
-> MIRROR ACTUAL TOPOLOGY CHANGES IN INDEX (OTHERWISE LEAVE INDEX ALONE)
-> RUN python3 scripts/check_icc_map.py
-> PUBLISH AS ONE COHERENT GIT COMMIT ONLY IF SEPARATELY AUTHORIZED
-> RETURN / STOP
```

The process should not publish partially synchronized file-by-file changes as a claimed completed refresh. If a write, verification or checker fails, leave the affected node stale/unverified, **do not advance its baseline/overlay**, apply `AGENTS.md` authoring-failure guardrail and stop with an exact failure. Git remains authority if source and ICC disagree. Do not perform build/runtime tests as a substitute for verifying ICC semantics, and do not present structural checker PASS as product acceptance.

## Exact short return contract

```text
BLACK SHEEP WALL RESULT
MODE: REUSE | UPDATE | REVEAL | BLOCKED | BOOTSTRAP
BOUNDARY: selected node or bounded source territory
CHECKOUT: observed branch + HEAD; overlay known/unknown
SOURCE DELTA: relevant paths only; "none" if proven
NODES READ: exact affected nodes
NODES WRITTEN: exact affected nodes or none
VALIDITY: current | partial | stale | unverified
CONTEXT ROUTE: INDEX -> selected node [-> required child/connector]
UNRESOLVED: exact missing source/comparison or none
RESULT: sufficient context | blocked (never task completion)
STOP
```

Keep normal reports short: optional counters are node reads/writes and source paths inspected. These counters belong in the current response, not a permanent task log. Never invent evidence or repeat reconnaissance just to fill report fields. If no local worktree is available, say so and return `unverified` when relevant; a remote review alone is not a full local maintenance completion.

## Verification coverage

`python3 scripts/check_icc_map.py` checks metadata structure, routing consistency, cycles, links, and implausible current-state claims without reading product source or writing ICC. `python3 scripts/test_icc_process.py` tests unchanged-node REUSE, relevant-file UPDATE, bounded missing-node REVEAL, stale and unknown-overlay handling, and metadata rejection. These checks establish the **process and format only**; source semantic truth and live runtime remain separate checks. The navigation ceiling, one ICC writer, Git source authority, minimum sufficient context and verified-state-before-advance are mandatory in every mode.
