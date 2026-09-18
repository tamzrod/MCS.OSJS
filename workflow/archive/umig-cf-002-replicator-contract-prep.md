# UMIG-CF-002 — CODE: Prepare Dormant Replicator Runtime Contract

Status: COMPLETE — source-only CODE gate, 2026-09-19. No unit/build/GUI/live test has run or passed.
Stage / owner: CODE / ChatGPT
Previous: none — independently human-approved prep, not an advancement from UMIG-003-T/V or UMIG-005.
Next: none — STOP; explicit next task selection is required.

## Approval and source checkpoint
Human replied `recommendations approved continue` to the proposed isolated Replicator preparation while OpenHands is occupied. Task selection and scope were committed at `2ebe63c24539c3d029f3dfe8acd62758c1e64836`, from prior main `2f124dee63bce62f95aa2c631bb5536e8177fa67`. Source-only code checkpoint: `b00bb29dfae4597016820483d00209383a91cf8b`. Read back both committed files and compared `2ebe63c..b00bb29`; the diff contains exactly these two added files:
- `OSJS/src/packages/MCSModbusToolkit/replicator-contract.js`: unimported, transport-injected version-1 Replicator contract for load/apply/status/suggest, unique request IDs, envelope/result checks, explicit runtime error propagation, invalid input rejection and apply document snapshot.
- `OSJS/tests/toolkit-replicator-contract.test.js`: authored future JR-only test cases; NOT RUN. Proposed command: `cd OSJS && node tests/toolkit-replicator-contract.test.js`.

## Wire contract and integration boundary
Derived from current `replicator/runtime_api.go` and OS.js `ModbusReplicator/server.js`: `version=1`, `request_id`, `operation`, `payload`, response `ok/result/error`. `load` returns `{document,suggestion}`; `apply` returns `{document,structural,message,completed_at}`; `status` returns the direct device status object with `name`, `enabled`, `running`, `source_status`, `blocks` (NOT a Simulator `{status}` wrapper); `suggest` returns a destination ownership object or performs `{inspect:true,port,unit_id}` lookup. Errors retain actual code/message, missing/mismatched responses reject. The transport is injected and is not wired to OS.js, network or any backend. No test output, live response, or working UI behavior is inferred from source review.

No Toolkit entry point or renderer import; no UI enablement, runtime server provider, Electron/legacy UI import, persistent write, Go, Docker, MMA2, existing deployment or ICC edit. UMIG-CF-001 Memory contract remains dormant. UMIG-003-T and UMIG-003-V are QUEUED/NOT RUN. UMIG-005 remains blocked on its original predecessor chain. Source-only preparation is archived; independent TEST/VERIFY and full integration are still pending.
