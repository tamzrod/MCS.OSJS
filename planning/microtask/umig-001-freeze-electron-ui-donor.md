# UMIG-001 — Freeze the Electron UI Donor Snapshot

Status: PLANNED / BLOCKED — early OS.js scaffold is authorized, but no Electron donor snapshot has been selected or approved for copying. Not ACTIVE.
Previous: UMIG-002 (independent OS.js placeholder scaffold)
Next: UMIG-003

## Primary outcome
Record one human-approved, immutable Electron renderer commit as the OS.js UI copy reference.

## Scope
After UMIG-002 is verified, obtain human approval for the renderer snapshot to copy; pin the exact SHA and list Memory, Replicator, Diagnostics, HTML/CSS/JS and UI assets. Identify Electron-only host calls requiring an OS.js adapter. The reference is provenance for one-time copying, not a shared build or automatic sync. The human's direction to begin OS.js migration does not by itself certify Windows installer or COMMS LED runtime acceptance; any later Windows fixes are tracked independently and require explicit donor reapproval only if the copied UI changes.

## Non-scope
No Windows LED repairs, Electron changes, OS.js renderer copy, promotion, runtime changes or automatic synchronization.

## Acceptance
1. Human approval of the chosen UI snapshot and exact commit SHA is recorded; Windows runtime verification is not misrepresented as complete.
2. A donor note identifies UI files/assets, native host boundary and one-time independent-copy provenance.

## Verification
Inspect the approved SHA and listed donor paths; verify the donor note states copy-only transfer and documents any known backend/UI status mismatch instead of claiming success.

## Dependencies
UMIG-002 verified, plus explicit human selection for promotion and an approved visual donor commit. This gate is for UI source identity, not blanket Windows LED/installer acceptance. Refresh stale affected ICC context through BLACK SHEEP WALL before promotion.

## Sizing
Surface 1, environment 0, behavior 0, verification 1, recovery 0 = 2.
