"""Offline ICC process and schema regressions. No product source or repository writes."""
import json
import sys
import tempfile
import unittest
from pathlib import Path

sys.path.insert(0, str(Path(__file__).resolve().parent))
from check_icc_map import check
from icc_decision import decide

class DecisionTests(unittest.TestCase):
    def setUp(self):
        self.node = {'sources': ['simulator/apply.go'], 'state': {'validity': 'current'}}

    def test_unchanged_reuses(self):
        result = decide(self.node, [], baseline_compared=True, overlay_compared=True)
        self.assertEqual(result.mode, 'REUSE')
        self.assertEqual(result.affected_paths, ())

    def test_one_changed_file_updates_exact_path(self):
        result = decide(self.node, ['electron/main.js', 'simulator/apply.go'],
                        baseline_compared=True, overlay_compared=True)
        self.assertEqual((result.mode, result.affected_paths), ('UPDATE', ('simulator/apply.go',)))

    def test_required_unmapped_reveals_only_bounded_source(self):
        result = decide(None, [], baseline_compared=False, overlay_compared=False,
                        bounded_source='simulator/new_component.go')
        self.assertEqual((result.mode, result.affected_paths),
                         ('REVEAL', ('simulator/new_component.go',)))

    def test_stale_without_delta_needs_scoped_verification(self):
        node = {**self.node, 'state': {'validity': 'stale'}}
        self.assertEqual(decide(node, [], baseline_compared=True,
                                overlay_compared=True).mode, 'UPDATE')

    def test_migrated_unknown_state_triggers_scoped_update(self):
        node = {**self.node, 'state': {'validity': 'unverified'}}
        self.assertEqual(decide(node, [], baseline_compared=False,
                                overlay_compared=False).mode, 'UPDATE')

    def test_unknown_overlay_blocks_not_false_reuse(self):
        self.assertEqual(decide(self.node, [], baseline_compared=True,
                                overlay_compared=False).mode, 'BLOCKED')

    def test_unbounded_missing_territory_blocks(self):
        self.assertEqual(decide(None, [], baseline_compared=True,
                                overlay_compared=True).mode, 'BLOCKED')

class SchemaTests(unittest.TestCase):
    def setUp(self):
        self.tmp = tempfile.TemporaryDirectory()
        self.addCleanup(self.tmp.cleanup)
        self.root = Path(self.tmp.name)
        icc = self.root / 'ICC'
        (icc / 'context').mkdir(parents=True)
        self.model = {'schema_version': 1, 'root': 'project',
                      'snapshot': {'reviewed_branch': 'main',
                                   'reviewed_head': 'a' * 40, 'local_overlay': 'unknown'},
                      'nodes': [
                          {'id': 'project', 'file': 'context/project.md', 'parent': None,
                           'children': ['simulator'], 'connectors': [], 'sources': ['README.md'],
                           'state': {'baseline': None, 'overlay': 'unknown', 'validity': 'unverified'}},
                          {'id': 'simulator', 'file': 'context/simulator.md', 'parent': 'project',
                           'children': [], 'connectors': [], 'sources': ['simulator/apply.go'],
                           'state': {'baseline': None, 'overlay': 'unknown', 'validity': 'unverified'}}]}
        (icc / 'context/project.md').write_text('Parent: none\nZoom In: simulator\n')
        (icc / 'context/simulator.md').write_text('Parent: project\nZoom Out: project\n')
        self.write()

    def write(self):
        icc = self.root / 'ICC'
        (icc / 'manifest.json').write_text(json.dumps(self.model))
        rows = ['| Node | Parent | Direct children | Purpose |', '| --- | --- | --- | --- |']
        for n in self.model['nodes']:
            rows.append(f"| [{n['id']}]({n['file']}) | {n['parent'] or 'none'} | {', '.join(n['children']) or '—'} | test |")
        (icc / 'INDEX.md').write_text('\n'.join(rows) + '\n')

    def test_valid(self):
        self.assertEqual(check(self.root), [])

    def test_wrong_parent_is_rejected(self):
        self.model['nodes'][1]['parent'] = None
        self.write()
        self.assertTrue(any('parent' in e for e in check(self.root)))

    def test_false_current_is_rejected(self):
        self.model['nodes'][1]['state']['validity'] = 'current'
        self.write()
        self.assertTrue(any('cannot claim current' in e for e in check(self.root)))

    def test_missing_node_is_rejected(self):
        (self.root / 'ICC/context/simulator.md').unlink()
        self.assertTrue(any('missing context' in e for e in check(self.root)))

if __name__ == '__main__':
    unittest.main(verbosity=2)
