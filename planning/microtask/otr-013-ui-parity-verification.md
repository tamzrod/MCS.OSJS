# OTR-013 — Independent UI parity verification
Status: PLANNING / UNDER REVIEW. Stage: VERIFY. Owner: OpenHands / independent JR. Previous: OTR-012 (and completed UI port source checkpoints). Next: OTR-014.

## Primary outcome
Produce an evidence-backed side-by-side parity report for ONE selected Electron Toolkit screen against its OS.js replica.

## Scope
Human review/promotion must select one screen and pin Electron and OS.js source revisions, safe disposable launch targets, exact repeatable launch/capture/navigation actions, expected visual and interaction states, screenshots, evidence destination and report-only permissions. Compare the reference's actual layout, labels, tabs, dialogs, navigation, draft/discard and error states applicable to that screen.

## Non-scope
No product fixes, guessed visual expectations, backend mutation, live operator data, deployment or blanket verification of uninspected screens. Additional screens get separate VERIFY microtasks.

## Acceptance
1. Matching reference and replica captures document the selected screen at equivalent viewport and state.
2. Each inspected navigation/interaction state has observed PASS/FAIL/BLOCKED evidence, with concrete differences and no inferred PASS.
3. Report names source revisions, environment, exact actions, screenshots and unverified states; JR does not repair or advance workflow.

## Evidence / handoff
Before promotion, replace the screen placeholder with exact actions and expected observations and a complete CWAL JR packet. Return original observations/screenshots and verdict; STOP. Do not execute this planning template as a test packet.

## Dependencies and size
OTR-001..007 and applicable CODE/TEST checkpoints. Size implementation 0, environment 1, behavior 1, verification 2, decision 1 = 5; split one screen per execution, and separate materially different scenarios.
