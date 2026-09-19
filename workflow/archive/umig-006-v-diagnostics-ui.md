# UMIG-006-V — VERIFY: Rendered Diagnostics Safety and Toolkit Reopen

Status: COMPLETE / PASS — reviewed 2026-09-19; independent live report `f6b7549b2c31e4fcf22427d8dee00b181f4b7186`.
Stage / owner: VERIFY / OpenHands JR; evidence review and closure / ChatGPT
Previous: UMIG-006-T (archived COMPLETE/PASS)
Next: UMIG-007 (QUEUED; approved donor screenshot baseline not yet available)

## Immutable evidence and disposition

JR report: `handoff.md` at `f6b7549b2c31e4fcf22427d8dee00b181f4b7186`. GitHub compare `f4994d435e1b3f02d05f2355abcd2e5c09b14901..f6b7549b2c31e4fcf22427d8dee00b181f4b7186` confirms ONE changed path, `handoff.md`, containing the authorized report. Reviewer read the exact packet and report; accepts JR's direct browser/backend evidence without claiming an independent rerun.

Sandbox-local `mcsverify-diag-006-20260919` was new, label-scoped and separate from prior retained volumes. Five-service test Compose startup, `/healthz`, two Unix sockets, initial empty canonical files and no fixture leakage were reported PASS. Real Chromium showed Diagnostics only read the first labelled canonical Memory/Replicator devices, Memory RUNNING/IDLE and Replicator RUNNING/OK with real FC3 polling, while all three global service indicators stayed UNKNOWN, native paths UNAVAILABLE, Start/Stop disabled, and text explicitly labelled observations rather than service logs. Diagnostics refresh preserved four canonical config/owner hashes and MMA2 restart count.

Previously outstanding UMIG-005-V lifecycle proof is now CLOSED: JR directly closed the actual Toolkit window via OS.js chrome, observed zero Toolkit windows, relaunched once from Start, observed fresh canonical Replicator device and RUNNING/OK, unchanged four hashes/owners and unchanged MMA2 restart count, no auto-apply. Real unbound synthetic source 127.0.0.1:15999 produced connection-refused ERROR in Replicator and Diagnostics without masking independent Memory IDLE or inventing global service health; restoring 127.0.0.1:15020 recovered OK after MMA2 reload. Two Memory plus four Replicator explicit Save & Apply clicks, test-only deletions restored all three seeded config hashes and released owners. Ownership-verified `down --remove-orphans` (without -v) removed five test containers and two networks, retaining only `mcsverify-diag-006-20260919_verify-data` for this project; previous `mcsverify-rep-005-20260919_verify-data` left intact. JR reported clean tracked tree and no product/Compose/workflow/ICC edits.

Non-blocking environment notes: `ss` absent (`/proc/net/tcp` alternative authorized); no `--no-color` flag used; browser field re-render required fresh element state; an 08:22:31 pre-reload snapshot was superseded by directly observed successful recovery after 08:22:35 without repeating apply.

## Scope boundary and successor gate

PASS is for isolated Linux real-browser Diagnostics, synthetic Go polling/errors, read-only safety, lifecycle and scoped cleanup. It does NOT prove actual Windows controls, global Docker/MMA2 supervisor health, real four-layer COMMS probes, genuine service logs, production deployment, final visual parity, or launcher cutover. The frozen donor renderer source SHA is `1c971b9a6e00bafadf329df8821421a40cfc079c` (`workflow/archive/umig-001-freeze-electron-ui-donor.md`); however approved comparable donor screenshots for all three tabs are not located in repo, and UMIG-003-V reports its screenshots were saved only outside the repo under `/tmp/jr-shots/`. UMIG-007 remains QUEUED pending an approved retrievable donor image baseline and author-written exact JR packet; no autonomous JR invocation authorized. No new source, production, general `operation cwal.md` or ICC edits.
