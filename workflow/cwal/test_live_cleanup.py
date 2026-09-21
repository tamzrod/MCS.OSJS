import importlib.util
from pathlib import Path
import unittest
from unittest.mock import patch

spec = importlib.util.spec_from_file_location('live_runner', Path(__file__).with_name('umig-em-003-v-live.py'))
runner = importlib.util.module_from_spec(spec)
spec.loader.exec_module(runner)


class CleanupTests(unittest.TestCase):
    def test_snapshot_failure_still_attempts_cleanup(self):
        with patch.object(runner, 'snap', side_effect=RuntimeError('container stopped')), patch.object(runner, 'safe_down') as cleanup, patch.object(runner, 'stamp'):
            status, reason = runner.finalize(True, 'test-container', 'PASS', 'checks passed')
        cleanup.assert_called_once_with(True)
        self.assertEqual(status, 'INCOMPLETE')
        self.assertIn('container stopped', reason)

    def test_cleanup_failure_preserves_product_failure(self):
        with patch.object(runner, 'safe_down', side_effect=RuntimeError('daemon unavailable')), patch.object(runner, 'stamp'):
            status, reason = runner.finalize(True, None, 'FAIL', 'product contradiction')
        self.assertEqual(status, 'FAIL')
        self.assertIn('product contradiction', reason)
        self.assertIn('daemon unavailable', reason)

    def test_inspection_failure_never_calls_down(self):
        def docker(*args, **kwargs):
            if args[1] == 'ls':
                return runner.PROJECT + '_verify-data'
            raise runner.Gate('inspection denied')
        with patch.object(runner, 'd', side_effect=docker), patch.object(runner, 'dc') as compose:
            with self.assertRaises(runner.Gate):
                runner.safe_down(True)
        compose.assert_not_called()

    def test_inventory_failure_never_calls_down(self):
        with patch.object(runner, 'd', side_effect=runner.Gate('daemon unavailable')), patch.object(runner, 'dc') as compose:
            with self.assertRaises(runner.Gate):
                runner.safe_down(True)
        compose.assert_not_called()

    def test_empty_or_foreign_label_never_calls_down(self):
        for label in ['', '<no value>', 'another-project']:
            with self.subTest(label=label):
                with patch.object(runner, 'd', side_effect=[runner.PROJECT + '_verify-data', label]), patch.object(runner, 'dc') as compose:
                    with self.assertRaises(runner.Gate):
                        runner.safe_down(True)
                compose.assert_not_called()

    def test_absent_networks_are_distinguished_from_inspection_errors(self):
        def docker(*args, **kwargs):
            if args[1] == 'ls':
                return runner.PROJECT + '_verify-data' if args[0] == 'volume' else ''
            return runner.PROJECT
        with patch.object(runner, 'd', side_effect=docker), patch.object(runner, 'dc', return_value='') as compose, patch.object(runner, 'stamp'):
            runner.safe_down(True)
        compose.assert_any_call('down', '--remove-orphans', timeout=120)


if __name__ == '__main__':
    unittest.main()
