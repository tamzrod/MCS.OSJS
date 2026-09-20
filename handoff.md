# Handoff

## Current authority — 2026-09-20

Human-approved sequence: finish `UMIG-EM-003-V` independent live VERIFY using the EXISTING disposable OpenHands environment, then prepare the separate OpenCode coding trial for `UMIG-EM-004` after accepted VERIFY and human promotion. `UMIG-EM-003-R-T` Go race TEST is archived PASS; original evidence in `workflow/archive/umig-em-003-r-t-passed-report.md` (report commit `04a31475c6427b185c09767703564894bd2791ed`). Go TEST is not live VERIFY.

SOLE ACTIVE: `workflow/active_work/umig-em-003-v-shared-lock.md` — VERIFY, preparation authorized; live concurrency execution not yet authorized. `UMIG-EM-004` remains PLANNED. OpenHands is the independent JR for this task; JR never changes product, workflow or ICC. Only BLACK SHEEP WALL updates ICC.

## CURRENT OPENHANDS ACTION — ONE-PASS SANDBOX PREPARATION, NOT LIVE VERIFY

Proceed in the EXISTING disposable OpenHands sandbox (reuse the clean `/tmp/em003v-discovery-clone` checkout if still present; a new clone/workspace is NOT required). Execute the entire bounded preparation packet `workflow/cwal/umig-em-003-v-sandbox-preflight.md`, including its narrowly authorized one-time clean Git fast-forward if needed and, ONLY when the sandbox's local Docker daemon is missing, one bounded `sudo -n dockerd` startup as sandbox tooling preparation. Do not repeat already evidenced discovery merely to create a new sandbox. Fresh HEAD is checked against live remote at execution time. No operator environment or existing Docker project may be targeted. If startup fails, report the exact failure once and STOP; no retries/alternative environment investigation. No product source changes, no synthetic applies, no Compose up/build/down, no live test or report push under this preparation packet. Return the outputs and command exit codes in chat; STOP.

The existing verify-only Compose template is `deploy/verify/compose.yaml`, NOT `deploy/docker-compose.yml`. After the preparation report, the coding agent will pin ONE exact bounded concurrency-execution packet with request bodies, timing, isolation, hashes, cleanup and report delivery; this requires a separate explicit execution gate. Do not invoke Operation CWAL to guess the absent live packet and do not mark VERIFY PASS from preparation. `UMIG-EM-004` coding trial remains on hold until actual live VERIFY evidence is reviewed.
