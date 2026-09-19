# Electron Replicator Advanced Settings

## Boundary

Owns Replicator destination advanced settings and the shared Electron IP/CIDR editor.
Parent / Zoom Out: INDEX.md. No children.
Connector: electron-memory-layout.md for the existing shared MMA Settings popup.

## Source Dependencies

- electron/renderer/app.js
- electron/renderer/memory-advanced.js
- electron/renderer/style.css
- electron/main.js
- electron/replicator-runtime.js
- electron/test/memory-advanced.test.js
- electron/test/replicator-advanced.test.js
- electron/test/replicator-runtime.test.js
- electron/REPLICATOR_ADVANCED.md
- replicator/advanced_settings.go
- replicator/advanced_settings_test.go
- replicator/document.go
- replicator/document_store.go
- replicator/compose_document.go
- replicator/manager.go
- replicator/runtime_api.go

## Baseline / Overlay

Source baseline: 8b5541f.
Scoped overlay: Replicator destination editor and persistence, runtime capability guard, reusable preset dropdown and comma parsing.
Unrelated workflow records remain untouched.

## Facts

Replicator has Device Definition / Advanced Settings folder tabs. The latter reuses RBE Rules, State Sealing and Access Policy.
Advanced fields apply to destination MMA2 memory and persist in mma2_advanced. Ranges derive from Pull Blocks, not editable duplicate area fields.
Composition inherits existing policy, RBE and sealing when omitted, respects explicit null removal, and validates the complete candidate before writing it.
Cloned runtime documents copy nested decoded advanced values. The disk migration loader preserves the new field.
Advanced saves require the runtime load capability mma2_advanced, preventing silent loss with older backends.
RBE ID allocation and duplication consider loaded devices from both editors.
The IP/CIDR dropdown uses an independent native preset select, not a filtered datalist. Input accepts comma-separated trimmed values, ignoring empty comma segments.
Shared global RBE output remains in Memory / Advanced Settings / MMA Settings.

## Verification

node --test electron/test/*.test.js passed 38 tests.
go test -race -count=1 -timeout=90s ./... in replicator passed.
These tests do not establish human visual acceptance or deployment to installed services.
