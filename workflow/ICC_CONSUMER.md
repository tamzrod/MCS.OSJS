# ICC consumer contract — read path for other operations

This is a shared read contract, not a new directive or an execution permission. `BLACK_SHEEP_WALL.md` and `ICC/FORMAT.md` own maintenance/schema rules; `ICC/manifest.json` owns structural metadata and per-node validity; `ICC/INDEX.md` is its human-readable routing mirror. Git source and the assigned task/handoff remain authoritative. Only BLACK SHEEP WALL writes `ICC/`.

## Repository-dependent CODE, DISCOVERY, Brainstorm, Planning and Microtask work

1. Resolve the assigned task/question and its semantic ceiling FIRST. Read `ICC/INDEX.md` and only the relevant entry in `ICC/manifest.json`; do not tour the registry or select new tasks from it.
2. Check the selected node's **own** state (`baseline`, `overlay`, `validity`) against the actual checkout HEAD and relevant working-tree changes, when access and task permissions allow. A historic index HEAD, remote snapshot or another branch does not certify a local checkout. `unverified`, `partial`, `stale`, or incomparable state is not `current`.
3. If verified `current` with no relevant source delta, REUSE its smallest sufficient Markdown context and STOP gathering. If a required node is absent, or a selected node needs verification/refresh, request a bounded BLACK SHEEP WALL REVEAL/UPDATE **only where the caller's directive and task authorize that delegation**. The call receives the same semantic ceiling and returns control; it does not approve implementation or switch tasks.
4. If maintenance is not authorized or cannot be done safely, inspect only the authoritative files already permitted by the task when this is sufficient, explicitly flag the ICC as unverified, or report the exact blocker. Do not treat an ICC failure alone as a reason to expand source access, run tests, alter task state, or write ICC yourself.
5. Source verification required by the authorized task still applies: ICC is a navigation cache, never proof of build, test, live runtime, current workflow authorization or actual source behavior.

## Exact-packet TEST / VERIFY exception

Independent JR follows `operation cwal.md` and its current packet first. ICC is optional read-only context only when the packet requires it; a migrated `unverified` node is NOT an extra test gate. JR must not invoke BLACK SHEEP WALL, run extra Git probes, rewrite ICC, or delay an otherwise complete test packet merely to certify cache validity. If the packet explicitly requires ICC truth that cannot be established within its allowlist, report BLOCKED with the exact missing evidence for the coding/workflow owner to resolve separately. No product test substitutions or reruns.

## THE GATHERING exception

THE GATHERING observes **committed HEAD only**, and never inspects or verifies the uncommitted overlay. Read `ICC/INDEX.md` / relevant committed ICC status only as navigation or a summary when the precise status is demonstrably supported at the observed HEAD. A manifest node with null baseline, `unverified`, `partial`, `stale`, or an unproved committed dependency comparison is NOT certified current. For each of its four permitted views, read the corresponding authoritative committed handoff/active-work/planning/brainstorm file instead. Do not invoke BLACK SHEEP WALL or inspect product code, archives, unrelated paths or working-tree state. Report contradictions rather than resolve them by guessing.

## THERE IS NO COW LEVEL exception

Preserve/checkpoint the actual state according to its rescue directive. ICC can guide a bounded read, but must not postpone emergency preservation to certify the cache. Only the rescue directive's already-authorized narrow delegation may call BLACK SHEEP WALL, and only when essential. Rescue never writes ICC directly.

## Stop and reporting

Return to the original operation after any delegated refresh. Report `ICC: reused | refreshed | unverified/fallback | not needed` only when materially useful; never claim a Black Sheep Wall call occurred unless it did. No implied automation, repo-wide rescan, task promotion, commit, push or branch switch.