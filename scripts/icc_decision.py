"""Pure Black Sheep Wall routing decision: no Git I/O, source reads or ICC writes.

Caller supplies observed committed + working-tree changed paths; baseline/overlay
comparison is per node. This is not an automation trigger or a replacement for
semantic analysis of changed source.
"""
from dataclasses import dataclass
from fnmatch import fnmatchcase

@dataclass(frozen=True)
class Decision:
    mode: str
    reason: str
    affected_paths: tuple[str, ...] = ()

def intersects(source: str, changed: str) -> bool:
    """Compare repository-relative paths without opening source files."""
    if source.endswith('/**'):
        return changed.startswith(source[:-3] + '/')
    return fnmatchcase(changed, source)

def decide(node: dict | None, changed_paths: list[str], *,
           baseline_compared: bool, overlay_compared: bool,
           required: bool = True, bounded_source: str | None = None) -> Decision:
    """Choose exactly one route; do not perform the operation."""
    if not required:
        return Decision('REUSE', 'Outside requested boundary; do not inspect')
    if node is None:
        if bounded_source:
            return Decision('REVEAL', 'Required territory unmapped; scoped discovery only', (bounded_source,))
        return Decision('BLOCKED', 'Unmapped territory has no verified bounded source')
    sources = node.get('sources', [])
    affected = tuple(sorted({p for p in changed_paths if any(intersects(s, p) for s in sources)}))
    validity = node.get('state', {}).get('validity')
    if validity in ('stale', 'partial', 'unverified') and sources:
        return Decision('UPDATE', 'Scoped verification of existing non-current node', affected)
    if not baseline_compared or not overlay_compared:
        return Decision('BLOCKED', 'Baseline or local overlay has not been compared for current node')
    if affected:
        return Decision('UPDATE', 'Relevant committed or working-tree source delta', affected)
    if validity == 'current':
        return Decision('REUSE', 'Verified node; no intersecting source delta')
    return Decision('BLOCKED', 'Node is not proven current and has no bounded source dependencies')
