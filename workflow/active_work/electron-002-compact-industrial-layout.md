# ELECTRON-002 — Compact Industrial Desktop Layout

Status: ACTIVE

## Primary Outcome

Make the MCS desktop UI use the information-dense mentality of classic Modbus engineering tools: compact controls, minimal decorative spacing, table-first editing, small status indicators, and a substantially smaller useful default window without changing Simulator or Replicator behavior.

## Scope

- Reduce the Electron desktop default/minimum window footprint so the toolkit feels like a small utility rather than a dashboard.
- Compress the Electron shell title/status strip, tabs, and content padding.
- Compress the canonical Simulator editor layout: narrower device list, smaller controls, tighter identity fields, dense FC rows, compact status line, and left-aligned actions.
- Compress the canonical Replicator editor using the same industrial-tool density, including its Device/Pull Blocks surfaces.
- Keep status/activity indicators updating in place; this task must not introduce editor reconstruction for background status changes.
- Keep existing fields, actions, backend calls, persistence behavior, and runtime semantics unchanged.

## Non-Scope

- No backend/runtime protocol changes.
- No installer/service changes.
- No new Simulator or Replicator fields.
- No new navigation model.
- No removal of runtime status indicators.
- No completion claim for Windows packaging; rendered Windows acceptance is still required.

## Acceptance Criteria

1. Electron defaults to a compact utility-sized window around 900 x 560 and remains resizable.
2. Header/runtime strip and top tabs consume materially less vertical space than the current build.
3. Simulator device list is narrower and its controls/rows use compact desktop-tool spacing.
4. Simulator Name/Enabled/Port/Unit plus all four FC rows fit comfortably in the normal window without large blank regions.
5. Replicator Device/Pull Blocks UI uses the same compact control density and preserves the spreadsheet-first Pull Blocks design.
6. Save/Discard and list/block actions remain immediately visible and functionally unchanged.
7. Runtime/status activity continues to patch only status elements and must not steal focus from inputs/selects.

## Verification

- Re-read the modified Electron and canonical Simulator/Replicator layout sources to confirm only sizing/styling/window defaults changed.
- Windows rendered test: build/run the Electron package and compare against the current screenshot. Confirm useful operation at the new default size, no clipped required controls, FC/select interaction remains stable while status updates run, and resizing larger still works.

## Dependencies

None. This task is a human-prioritized UI pass. ELECTRON-001 and REP-BLOCK-002/003 remain queued for their pending Windows/JR verification and are not completed by this task.

## Sizing

Implementation surface 2; environment uncertainty 1; behavioral surface 0; verification surface 1; decision/recovery surface 0. Total 4 — tightly coupled visual-density changes with one rendered acceptance workflow.
