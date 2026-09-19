# UMIG-006-V — VERIFY: Rendered Diagnostics Safety

Status: QUEUED — wait for UMIG-006-T PASS and a separately authored exact JR packet.
Stage / owner: VERIFY / OpenHands (JR)
Previous: UMIG-006-T
Next: UMIG-007

## Primary outcome
Observe a truthful, safe Diagnostics tab in an actual isolated OS.js window, and close the one outstanding Toolkit window-reopen evidence gap from UMIG-005-V during the same rendered session when a safe canonical Replicator test device is available.

## Instruction / expected / evidence
In isolated OS.js, open Diagnostics with a reachable test runtime and one intentionally unavailable diagnostic source; inspect actual status/errors and Windows-only controls. Expected: accessible real observations, clearly UNKNOWN/UNAVAILABLE missing data, no usable unsupported native service control. Record exact actions, screenshots, logs, HEAD and target. GUI/runtime unavailable = BLOCKED; wrong product state = FAIL. The current file is QUEUED, NOT an executable JR packet.

Mandatory carried-forward check from the reviewed UMIG-005-V report `1a04664e5fdc21e3bd323a97a4f12f3f9d39abdd`: the reported 'close/reopen' used tab switching/re-selection; actual Toolkit WINDOW destruction and Start-menu relaunch were not demonstrated. In this already planned browser verification, use ONLY a separately authorized disposable canonical Replicator device/runtime and exact future JR packet to capture: persisted device before close; close Toolkit window via OS.js chrome; reopen it through Start menu; verify a new Toolkit window reloads that same canonical Replicator definition from the Go runtime with correct destination, no automatic apply and no fixture fallback. Verify unchanged persisted config hash/owners across close/reopen. A tab switch, reload inferred from YAML, or mock-only fixture is NOT this check. If the target cannot safely provide it, mark this subcheck NOT VERIFIED/BLOCKED and do not certify overall migration/cutover; no separate VM prerequisite, no reuse/deletion of retained volumes. ChatGPT must author the safe commands/limits in the specific JR packet, NOT modify general `operation cwal.md`.

## Non-scope
No service starts/stops or product debugging by JR, no source fixes, full visual parity or cutover; no additional unscripted exploratory tests.

## Dependencies
UMIG-006-T PASS and current JR packet. A safe disposable real runtime is required for the reopen regression; if not available, explicit evidence gap persists.

## Sizing
Surface 0, environment 1, behavior 0, verification 1, recovery 0 = 2.
