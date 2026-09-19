# Handoff

## Current direction and source checkpoint — 2026-09-19

Human directive: follow the CURRENT Electron MCS Toolkit functional MODEL, not the old frozen donor as a complete feature target. Latest source-reviewed main `449cfda29d23406295cd0ccbc35fa8f5d6293681`: Memory layout/RBE at `ff6846f`, Go Replicator COMMS at `1ca1740`, and Windows-only settings-owner/installer changes at `449cfda`. Re-pin on later relevant commits. Old `1c971b9` donor is historical provenance. Preserve UMIG-006-V immutable JR evidence `f6b7549` and reviewer archive `ec390c0`; those passes do not cover subsequently changed Go/MMA2/Replicator, features, production or parity.

## Sole ACTIVE DESIGN — UMIG-EM-001 / NO JR TEST TASK

Human approved proceeding with the contract stage. Promoted `workflow/active_work/umig-em-001-electron-model-contract.md`; detailed versioned compatibility matrix, proposed Go ownership/authentication/revision/recovery and truthful COMMS semantics are in `docs/TOOLKIT_ELECTRON_MODEL_CONTRACT.md`. Architecture/security contract remains PROPOSED pending explicit human decision; planning successor tasks are NOT promoted or executable. UMIG-007 visual parity QUEUED/DEFERRED; UMIG-007A QUEUED. Do NOT issue `OPERATION CWAL` or use an earlier JR packet while DESIGN is ACTIVE. A later design review/approval is not a test PASS.

## Guardrails and next

ChatGPT owns design/CODE, task-specific JR packet authorship, evidence adjudication and advancement; OpenHands/JR TEST/VERIFY ONLY. Only BLACK SHEEP WALL edits ICC. Do not edit `operation cwal.md` to repair a task packet. No product source/Go/Compose changes or Docker runs authorized by this design promotion; no production `osjs-data`, customer devices, Windows installations, legacy removal, RBE listener activation or retained-volume reuse/removal. Proposed future sequence: upgraded Go baseline TEST/VERIFY -> shared Go management CODE/T/V -> per-device Advanced CODE/T/V -> shared modal CODE/T/V -> COMMS CODE/T/V -> current-model visual parity -> separate deployment. Prior test evidence never silently migrates to new SHA.

Next action: read and approve/amend architecture decisions in `docs/TOOLKIT_ELECTRON_MODEL_CONTRACT.md`; once all planned successor microtasks are linked and a human explicitly promotes the ordered stages, ChatGPT archives DESIGN and issues only the next exact JR packet. Recommendation: maintain the stop gate until security, multi-writer serialization and post-restart-failure behavior are approved.
