#!/usr/bin/env python3
"""Check ICC link/tree integrity without source-code discovery or ICC writes.

Usage: python3 scripts/check_icc_map.py [repository-root]
"""

from __future__ import annotations

import re
import sys
from pathlib import Path

ROW = re.compile(r"^\|\s*\[([^]]+)\]\((context/[^)]+\.md)\)\s*\|\s*([^|]+)\|\s*([^|]+)\|", re.M)
LINK = re.compile(r"\[[^]]+\]\(([^)]+)\)")
PARENT = re.compile(r"^\s*(?:Parent(?:\s*/\s*Zoom Out)?|Zoom Out):\s*(.*)$", re.I | re.M)


def name_of(raw: str) -> str:
    raw = raw.strip().strip('`')
    link = re.search(r"\[([^]]+)\]\([^)]+\)", raw)
    if link:
        return link.group(1).strip()
    if raw.lower().startswith(('none', '—', '-')):
        return 'none'
    return re.split(r"[,;\s]", raw, maxsplit=1)[0].strip('`')


def check(root: Path) -> list[str]:
    root = root.resolve()
    icc = root / 'ICC'
    index = icc / 'INDEX.md'
    if not index.is_file():
        return ['missing ICC/INDEX.md']
    index_text = index.read_text(encoding='utf-8')
    rows = ROW.findall(index_text)
    if not rows:
        return ['no parseable semantic registry rows in ICC/INDEX.md']
    errors: list[str] = []
    nodes: dict[str, tuple[Path, str, list[str]]] = {}
    for name, relative, parent, children in rows:
        name = name.strip()
        path = icc / relative
        kids = [] if children.strip() in ('—', '-', 'none') else [x.strip() for x in children.split(',')]
        if name in nodes:
            errors.append(f'duplicate registry node: {name}')
        nodes[name] = (path, parent.strip(), kids)
        if not path.is_file():
            errors.append(f'missing node file: ICC/{relative}')
    for name, (path, parent, kids) in nodes.items():
        if parent != 'none' and parent not in nodes:
            errors.append(f'{name}: unknown parent {parent}')
        for child in kids:
            if child not in nodes:
                errors.append(f'{name}: unknown child {child}')
            elif nodes[child][1] != name:
                errors.append(f'{name}: child {child} declares parent {nodes[child][1]} in index')
        if parent in nodes and name not in nodes[parent][2]:
            errors.append(f'{name}: parent {parent} does not list this child')
        if not path.is_file():
            continue
        text = path.read_text(encoding='utf-8')
        declared = PARENT.search(text)
        if not declared:
            errors.append(f'{name}: no Parent/Zoom Out declaration')
        elif name_of(declared.group(1)) != parent:
            errors.append(f'{name}: file parent {name_of(declared.group(1))} != index parent {parent}')
        # Only inspect explicit Markdown targets in ICC, not repository source dependencies.
        for target in LINK.findall(text):
            target = target.partition('#')[0]
            if not target or '://' in target or target.startswith('mailto:'):
                continue
            resolved = (path.parent / target).resolve()
            if not resolved.is_relative_to(icc):
                continue  # external repository document, not an ICC node
            if not resolved.is_file():
                errors.append(f'{name}: broken ICC link {target}')
    for target in LINK.findall(index_text):
        target = target.partition('#')[0]
        if not target or '://' in target:
            continue
        resolved = (icc / target).resolve()
        if resolved.is_relative_to(icc) and not resolved.is_file():
            errors.append(f'index: broken ICC link {target}')
    return sorted(set(errors))


def main() -> int:
    root = Path(sys.argv[1]) if len(sys.argv) > 1 else Path(__file__).resolve().parent.parent
    problems = check(root)
    for problem in problems:
        print('ICC MAP ERROR:', problem)
    if problems:
        print(f'ICC map check: FAIL ({len(problems)} distinct issue(s)); no files changed')
        return 1
    print('ICC map check: PASS (registry, parent/child and Markdown links); no files changed')
    return 0


if __name__ == '__main__':
    raise SystemExit(main())
