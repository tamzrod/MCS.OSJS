# UMIG-EM-002-R — TEST: current Replicator Go advanced-settings regression

Status: COMPLETE / PASS — coding-agent evidence review 2026-09-20.
Stage / owner: TEST / independent OpenHands JR; adjudication / ChatGPT.
Previous: UMIG-EM-002-B (archived PASS).
Next: UMIG-EM-002-V (human-authorized queued successor, now promoted separately).

## Immutable evidence and disposition

The original once-only Go test ran against the product baseline `c98eacef8a6af8b0786a09766a4a034e25e82ce6` at JR checkout `62d05fe94568fbf9903848d08b2233861d0f4f2e`. JR's report commit `9f9746e0a2a3d99bc09dc27259440b316f297653` initially contained required PASS lines and summaries but lacked the required entire stdout/stderr. The coding agent did NOT prematurely close the evidence gate; a formal Operation CWAL accounting packet was committed at `d5b68c6f5d38dcabe43dcde49281c8060f18a43d`.

JR's second, evidence-only CWAL report at `0290e6b3bb29552c04fedbc6651d99b399a570d0` appended the retained original 128-line stdout/stderr transcript from `/tmp/cwal_test.log` with its provenance and hash, and expressly did not rerun the product test. GitHub comparison `d5b68c6..0290e6b` confirms ONLY `handoff.md` changed: 139 additions, 0 deletions. Direct readback of that commit confirms `TEST_EXIT=0`, `real 0m30.620s`, both Go packages `ok`, the three named advanced-settings tests PASS, no FAIL or race report, plus the separately reported original pre/post clean tracked tree. The earlier packet recorded Go 1.25.0 and permitted sandbox-local setup. The coding agent accepts the independent TEST PASS on this bounded evidence, not on a fresh self-run.

## Scope limit

PASS covers the current Replicator Go unit/regression suite and evidence transport only. It does NOT prove an actual new Linux disposable stack, rendered OS.js Toolkit, runtime sockets, live MMA2 ownership, raw ingest, actual COMMS, global service health, RBE listener, deployment or visual parity. `UMIG-EM-002-V` is the separately authorized next VERIFY gate. No production resource was touched or product source edited by this closure. ICC remains stale and may be updated only through BLACK SHEEP WALL.
