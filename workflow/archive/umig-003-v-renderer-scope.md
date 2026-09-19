# UMIG-003-V — VERIFY: Fixture Tabs and Shell Isolation

Status: COMPLETE — independent rendered VERIFY PASS reviewed by ChatGPT, 2026-09-19.
Stage / owner: VERIFY / OpenHands (JR via `operation cwal.md`); evidence review and closure / ChatGPT
Previous: UMIG-003-T (archived BUILD/STATIC PASS)
Next: UMIG-004 (promoted ACTIVE in the workflow-advancement commit)

## Evidence and review

OpenHands/JR executed the CURRENT UMIG-003-V packet in a disposable checkout at `1d4cd826f32f6fb9f0972d8dd4278dc6dd0567ea` and committed the full JR PASS report in the immutable `handoff.md` version at `6f82196700b1c312652fdd7584a508e5452c4822`. ChatGPT read the complete report and compared `1d4cd82..6f82196`: exactly `handoff.md` was modified, and the commit patch changed only the authorized JR TEST REPORT — UMIG-003-V section; no product, workflow or ICC edits by JR. This review accepts JR's direct evidence, not a second test run.

Required results reported: Node 16/npm 8, clean disposable checkout and both ancestor checks; build:local-packages, package:discover and build exit 0, isolated serve and `/healthz` HTTP 200. In real Chromium 152 (1280x800) with fresh data/session, zero existing windows and exactly one actual Start-menu click produced exactly one `MCSModbusToolkitWindow` with intact chrome/taskbar. Real Memory → Replicator → Diagnostics → Memory clicks yielded exactly one computed-visible active panel and correct aria-selected states. Memory showed `Fixture Sim-PLC-1` and FC1–FC4; Replicator showed `Fixture Rep-PLC-1`, destination/pull block FC3, unknown ownership and four neutral-gray COMMS LEDs; Diagnostics showed UNKNOWN Simulator status, UNAVAILABLE native runtime/paths and fixture log. Header UNKNOWN, readonly fields and disabled unsafe actions persisted. Browser observations and computed styles found no shell CSS leakage, no required-content clipping, working Start menu, taskbar minimize/restore, advancing clock, and both legacy apps still listed. Zero post-launch/backend requests and no page errors. Four existing icon/sound 404s were documented as non-gating. Test-only server/data cleaned, port 18209 freed, tracked checkout clean. Screenshots and structured logs were reported outside checkout under `/tmp/jr-shots/` and `/tmp/jr-verify-results.json`; they are not stored in this repository.

## Scope boundary

PASS proves fixture-only rendered tabs and OS.js shell isolation at tested HEAD; it does NOT prove live Simulator/Replicator/MMA2 connectivity, configuration writes, dormant contract/model tests, Docker deployment, Windows/Electron behavior, or final donor pixel parity. UMIG-004 is next CODE stage with separate TEST/VERIFY gates. Preserve operator deployment and data, legacy applications, and all later cutover approval boundaries.

## Sizing
Surface 0, environment 1, behavior 0, verification 1, recovery 0 = 2.
