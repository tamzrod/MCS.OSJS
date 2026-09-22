# OTR-002A — OS.js Toolkit UI baseline

Status: ACTIVE
Stage: DISCOVERY
Owner: OpenCode
Previous: OTR-001C
Next: OTR-002B
Size: I0 E1 B0 V1 D0 = 2

## One outcome
Write a source-linked inventory of the EXISTING OS.js MCS Modbus Toolkit window, tabs, and UI files. Do not change product code.

## Allowed files to write
- `workflow/active_work/evidence/otr-002a-report.md` (create)
- `handoff.md` (only the Current task section, after the report exists)
- `workflow/active_work/OTR_PROMOTION_QUEUE.md` (only the 002A/002B lines)
- this file (Status line only: ACTIVE → COMPLETE after acceptance passes)
- `workflow/active_work/otr-002b-osjs-backend-baseline.md` (Status line only: QUEUED → ACTIVE)

## Forbidden
Any other path. No `OSJS/` edits. No `electron/` edits. No archive moves. No ICC. No services. No build. No merge, force-push, PR, or push to `main`.

## Exact files to read (only these)
1. `OSJS/src/packages/MCSModbusToolkit/metadata.json`
2. `OSJS/src/packages/MCSModbusToolkit/package.json`
3. `OSJS/src/packages/MCSModbusToolkit/index.js`
4. `OSJS/src/packages/MCSModbusToolkit/toolkit-renderer.js`
5. `OSJS/src/packages/MCSModbusToolkit/index.scss`
6. `OSJS/src/packages/MCSModbusToolkit/renderer.css`
7. `OSJS/src/packages/MCSModbusToolkit/webpack.config.js`

If a listed file is missing, write BLOCKED evidence, publish it through the BLOCKED transport below,
then STOP. Do not search other trees for a replacement path.

## Numbered steps
1. Run `git branch --show-current`. If not `opencode`, STOP.
2. Run `git rev-parse HEAD`, `git status --short`, `git rev-parse origin/opencode`, and
   `git ls-remote origin refs/heads/opencode`. Status must be empty and all three SHAs must match;
   otherwise write BLOCKED evidence only and do not advance.
3. Read the seven files above.
4. In the report file write these headings only:
   - HEAD
   - git status
   - Window id and title (from index.js / metadata.json)
   - Tab names and `data-tab` values (from toolkit-renderer.js)
   - Panel element ids
   - Which modules index.js mounts into which roots
   - CSS files used by the window
   - What is fixture/read-only vs live (quote the source comment or string)
   - Unknowns for OTR-002B (backend only; do not inspect backend in this task)
   - Verdict: SOURCE INVENTORY COMPLETE or BLOCKED
5. Use path and line anchors. Do not invent directories.
6. Do not claim visual PASS. Do not claim Electron parity.
7. If and only if verdict is SOURCE INVENTORY COMPLETE, set handoff Current task to OTR-002B, mark
   002A COMPLETE and 002B ACTIVE in the queue, set this Status to COMPLETE, and set the OTR-002B
   packet Status to ACTIVE. Do not execute OTR-002B.
8. Run `git diff --check`; verify `git diff --name-only` contains only the five allowed paths.
9. Stage only the five allowed paths and run `git diff --cached --name-only` plus
   `git diff --cached --check`. Commit with exactly `git commit -m "OTR-002A: record OS.js UI baseline"`.
10. Re-run `git ls-remote origin refs/heads/opencode`; it must still equal the pre-task HEAD. Then run
    `git push origin HEAD:opencode` and re-run the remote query; it must equal the new commit SHA.
11. Report the tested HEAD, report commit, pushed remote SHA, and changed paths. STOP.

If evidence is BLOCKED, create/update only the report, commit it with
`git commit -m "OTR-002A: record blocked UI baseline"`, push and verify it the same way, leave task
and queue statuses unchanged, then STOP. A failed/unverified push is INCOMPLETE.

## Acceptance (max 3)
1. Window launch path and window id are named with file anchors.
2. The three tabs and panel ids are named with file anchors.
3. Unknowns for backend are listed without inspecting backend files.
