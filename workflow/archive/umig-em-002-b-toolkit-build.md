# UMIG-EM-002-B — TEST: upgraded Toolkit Node/build regression

Status: COMPLETE / PASS, adjudicated 2026-09-20 from independent JR evidence committed at `6612784f458339521e86ec822c4516c9af15df2b` (see that commit's `handoff.md` for the full immutable packet and report).
Stage / owner: TEST / OpenHands JR; adjudication / ChatGPT.
Previous: UMIG-EM-002-T (archived COMPLETE / Go UNIT PASS).
Next: UMIG-EM-002-R (ACTIVE independent regression of newer Replicator Go changes).

## Evidence and bounded finding

The 2026-09-19 JR report records clean guarded source sync to `5c5894d715e1acd4ed0e35ceda5b29bce2d1ea5b`, compatible Node 16, all twelve specified Toolkit Node tests, three local build/discovery stages and six artifact checks at exit 0. Final tracked status was clean. Accept PASS for the OS.js/Toolkit source at checkpoint `8b5541fcefeba8df82bea675b15bd0d03c7c6c20`, not as a live UI or backend pass.

The comparison `5c5894d..c98eace` shows no OSJS source changes but DOES show changes to `replicator/advanced_settings.go`, `replicator/advanced_settings_test.go`, `replicator/compose_document.go`, `replicator/document.go`, `replicator/document_store.go`, `replicator/manager.go`, and `replicator/runtime_api.go`. The prior Go test predates these changes and does not establish their regression result. An independent focused current-Go TEST is therefore next, before disposable live verification or further migration CODE. No rerun is claimed here.

Non-scope: no rendered GUI, live service, RBE/access-event network exposure, Windows installer, final visual parity, or production certification was established. The old current packet/report is retained in Git history at the evidence commit rather than duplicated into the next `handoff.md` packet.
