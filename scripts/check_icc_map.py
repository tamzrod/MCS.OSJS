#!/usr/bin/env python3
"""Validate ICC schema, tree, state claims and links, without source reads or writes.

Usage: python3 scripts/check_icc_map.py [repository-root]
"""
from __future__ import annotations
import json
import re
import sys
from pathlib import Path, PurePosixPath

ROW = re.compile(r'^\|\s*\[([^]]+)\]\((context/[^)]+\.md)\)\s*\|\s*([^|]+)\|\s*([^|]+)\|', re.M)
LINK = re.compile(r'\[[^]]+\]\(([^)]+)\)')
PARENT = re.compile(r'^\s*(?:Parent(?:\s*/\s*Zoom Out)?|Zoom Out):\s*(.*)$', re.I | re.M)
SHA = re.compile(r'^[0-9a-f]{40}$')
ID = re.compile(r'^[a-zA-Z0-9][a-zA-Z0-9-]*$')
VALIDITY = {'current', 'unverified', 'stale', 'partial'}
RELATIONS = {'contract', 'dependency', 'data-flow', 'ownership'}

def safe_path(raw: str) -> bool:
    return (isinstance(raw, str) and bool(raw) and not raw.startswith('/')
            and '\\' not in raw and not any(p in ('', '.', '..') for p in raw.split('/'))
            and not PurePosixPath(raw).is_absolute())

def name_of(raw: str) -> str:
    raw = raw.strip().strip('`')
    match = re.search(r'\[([^]]+)\]\([^)]+\)', raw)
    if match:
        return match.group(1).strip()
    token = re.split(r'[,;\s]', raw, maxsplit=1)[0].strip('`')
    return 'none' if token.lower() in ('none', '—', '-') else token

def check(root: Path) -> list[str]:
    icc = root.resolve() / 'ICC'
    index, manifest = icc / 'INDEX.md', icc / 'manifest.json'
    errors: list[str] = []
    if not index.is_file() or not manifest.is_file():
        return [f'missing {p}' for p in ('ICC/INDEX.md', 'ICC/manifest.json')
                if not (root.resolve() / p).is_file()]
    index_text = index.read_text(encoding='utf-8')
    try:
        model = json.loads(manifest.read_text(encoding='utf-8'))
    except (ValueError, UnicodeError) as exc:
        return [f'invalid ICC/manifest.json: {exc}']
    if not isinstance(model, dict) or model.get('schema_version') != 1 or not isinstance(model.get('nodes'), list):
        return ['manifest: expected schema_version=1 and nodes array']
    root_id = model.get('root')
    snapshot = model.get('snapshot')
    if not isinstance(snapshot, dict) or snapshot.get('local_overlay') not in ('unknown', 'clean', 'audited'):
        errors.append('manifest: snapshot needs local_overlay unknown, clean or audited')
    if isinstance(snapshot, dict) and not SHA.fullmatch(str(snapshot.get('reviewed_head', ''))):
        errors.append('manifest: reviewed_head must be full SHA')
    nodes: dict[str, dict] = {}
    files: set[str] = set()
    for position, item in enumerate(model['nodes']):
        if not isinstance(item, dict):
            errors.append(f'node {position}: expected object')
            continue
        name = item.get('id')
        if not isinstance(name, str) or not ID.fullmatch(name):
            errors.append(f'node {position}: invalid id')
            continue
        if name in nodes:
            errors.append(f'duplicate registry id: {name}')
        nodes[name] = item
        file = item.get('file')
        if not safe_path(file) or not file.startswith('context/') or not file.endswith('.md'):
            errors.append(f'{name}: invalid context path')
            continue
        if file in files:
            errors.append(f'{name}: duplicate context path {file}')
        files.add(file)
        if not (icc / file).is_file():
            errors.append(f'{name}: missing context {file}')
        for field in ('children', 'connectors', 'sources'):
            if not isinstance(item.get(field), list):
                errors.append(f'{name}: {field} must be array')
        sources = item.get('sources', [])
        if isinstance(sources, list):
            if not sources:
                errors.append(f'{name}: missing bounded source dependencies')
            for source in sources:
                if not safe_path(source) or source.startswith('ICC/'):
                    errors.append(f'{name}: invalid/self-referential source {source!r}')
        state = item.get('state')
        if not isinstance(state, dict) or state.get('validity') not in VALIDITY:
            errors.append(f'{name}: state needs valid validity')
        else:
            baseline, overlay = state.get('baseline'), state.get('overlay')
            if baseline is not None and not SHA.fullmatch(str(baseline)):
                errors.append(f'{name}: baseline must be full SHA or null')
            if overlay not in ('unknown', 'clean', 'audited'):
                errors.append(f'{name}: invalid overlay')
            if state['validity'] == 'current' and (baseline is None or overlay == 'unknown'
                                                  or not SHA.fullmatch(str(baseline))):
                errors.append(f'{name}: cannot claim current without full baseline and audited overlay')
            if overlay == 'audited' and not isinstance(state.get('fingerprints'), dict):
                errors.append(f'{name}: audited overlay requires path:fingerprint mapping')
            if overlay == 'audited' and isinstance(state.get('fingerprints'), dict):
                if not state['fingerprints'] or any(not safe_path(k) or not SHA.fullmatch(str(v))
                                                   for k, v in state['fingerprints'].items()):
                    errors.append(f'{name}: invalid audited fingerprints')
    if root_id not in nodes or nodes[root_id].get('parent') is not None:
        errors.append('manifest: root must exist with parent null')
    for name, item in nodes.items():
        parent, children = item.get('parent'), item.get('children', [])
        if parent is not None and parent not in nodes:
            errors.append(f'{name}: unknown parent {parent}')
        if not isinstance(children, list):
            continue
        if len(children) != len(set(map(str, children))):
            errors.append(f'{name}: duplicate child')
        for child in children:
            if child not in nodes:
                errors.append(f'{name}: unknown child {child}')
            elif nodes[child].get('parent') != name:
                errors.append(f'{name}: child {child} has wrong parent')
        if parent in nodes and name not in nodes[parent].get('children', []):
            errors.append(f'{name}: parent {parent} does not list child')
        for connector in item.get('connectors', []):
            if (not isinstance(connector, dict) or connector.get('target') not in nodes
                    or connector.get('relation') not in RELATIONS):
                errors.append(f'{name}: invalid connector {connector!r}')
        file = item.get('file')
        if not safe_path(file) or not (icc / file).is_file():
            continue
        text = (icc / file).read_text(encoding='utf-8')
        match = PARENT.search(text)
        if not match:
            errors.append(f'{name}: no parent/Zoom Out metadata in knowledge file')
        elif name_of(match.group(1)) != (parent or 'none'):
            errors.append(f'{name}: legacy parent differs from manifest')
        for target in LINK.findall(text):
            target = target.partition('#')[0]
            if not target or '://' in target or target.startswith('mailto:'):
                continue
            resolved = ((icc / file).parent / target).resolve()
            if resolved.is_relative_to(icc) and not resolved.is_file():
                errors.append(f'{name}: broken ICC link {target}')
    if root_id in nodes:
        seen: set[str] = set()
        stack = [root_id]
        while stack:
            candidate = stack.pop()
            if candidate in seen:
                errors.append(f'tree cycle or multiple traversal: {candidate}')
                continue
            seen.add(candidate)
            stack.extend(nodes[candidate].get('children', []))
        if set(nodes) != seen:
            errors.append(f'unreachable nodes: {sorted(set(nodes) - seen)}')
    rows = ROW.findall(index_text)
    table: dict[str, tuple[str, str, list[str]]] = {}
    for name, file, parent, children in rows:
        name = name.strip()
        if name in table:
            errors.append(f'index duplicate node: {name}')
        table[name] = (file, parent.strip(), [] if children.strip() in ('—', 'none', '-')
                       else [x.strip() for x in children.split(',')])
    if set(table) != set(nodes):
        errors.append(f'index/manifest registry difference: {sorted(set(table) ^ set(nodes))}')
    for name in set(table) & set(nodes):
        item = nodes[name]
        if table[name] != (item.get('file'), item.get('parent') or 'none', item.get('children')):
            errors.append(f'{name}: index routing differs from manifest')
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
        print(f'ICC schema check: FAIL ({len(problems)}); no files changed')
        return 1
    print('ICC schema check: PASS (schema, tree, routing, state claims, links); no files changed')
    return 0

if __name__ == '__main__':
    raise SystemExit(main())
