# PERSIST-019 — Atomic Unseal / Commit Step

Status: CODE COMPLETE — 2026-10-08 (OpenHands JR DEV); awaiting independent TEST/VERIFY (PERSIST-022)
Stage: CODE
Owner: OpenHands JR DEV
Previous: PERSIST-018
Next: PERSIST-020

## Primary outcome
Make the final successful restore action an explicit write of the existing State Sealing flag to 1.

## Scope
Implement or verify only the primary outcome above against the current authoritative MCS.OSJS source and existing MMA2 contracts. Preserve the Port → Unit ID → Memory authority model and existing State Sealing/RBE semantics unless this packet explicitly says otherwise.

## Non-scope
Do not add runtime reseal or alternate commit flags.

## Acceptance
1. Unseal occurs only after restore verification success.
2. The flag location comes only from current State Sealing configuration.
3. No earlier restore step can expose memory to Modbus.

## Evidence / handoff
Coding owner records the exact source diff, changed paths, targeted self-checks and resulting source checkpoint. Do not claim independent TEST/VERIFY PASS. STOP after this task.

## Dependencies
Requires genuine completion evidence for PERSIST-018. Being present in `workflow/active_work/` authorizes this written scope but does not bypass predecessor or current-task selection gates.

## Sizing
1/0/2/1/2=6.

## Coding evidence (PERSIST-019, OpenHands JR DEV)

- Source checkpoint base: `261faa89a36e00a6166a06086acd1bce5f7785c1` (`main`, clean, = `origin/main` = `git ls-remote origin refs/heads/main` at run).
- Changed paths (product):
  - M `mma2composer/persistence_restore.go` (blob `7302908c117091ef3755304e9101f12c9527583e`) — after `VerifyPersistenceRestore` success, `RestorePersistencePlan` performs the explicit final commit: it writes the authoritative State Sealing flag to unsealed (1) through the existing Raw Ingest v1 contract as a single-coil write (`RawIngestCoils`, address = `plan.SealingFlag.Address`, count 1, payload `0x01`) — the final action after every area restore. The flag location comes solely from the State Sealing configuration (`PersistenceSealingFlag`); no runtime reseal or alternate commit flag is added. `PersistenceRestoreResult` gains `Committed`, true only when the unseal write was acknowledged with `0x00`; a non-zero unseal response leaves the restore completed but not committed (still sealed). New helper `encodePersistenceSealingFlag`.
  - M `mma2composer/persistence_restore_test.go` (blob `fb4f9e77b28095e08abaf2f55ca3b5c652d12117`) — PERSIST-019 self-check.
- Targeted self-check: `cd mma2composer && go test -mod=readonly -run 'TestPersistenceRestore|TestPersistenceSealing|TestForcePersistenceSealingFlag|TestVerifyPersistenceRestore|TestEncodePersistenceSealingFlag' -v .` → 23 tests PASS, exit 0.
- Additional bounded real-socket check (temporary, removed before commit): over a real Raw Ingest v1 server the final wire packet was the single-coil unseal (value 1) at the configured address, and a server rejecting the unseal left the restore completed but not committed; passed with `-race`, exit 0.
- Bounded regression (`-mod=readonly`): full `mma2composer`, `simulator`, `replicator`, `MMA2`, `mma2raw` suites → all `ok`, exit 0. `gofmt -l` clean for changed files; `go vet` exit 0.
- Acceptance mapping: (1) unseal occurs only after restore verification success — the unseal block runs only after `VerifyPersistenceRestore` returns true, and a failed area returns before it; (2) the flag location comes only from current State Sealing configuration — the write uses `plan.SealingFlag.Address` and no other source; (3) no earlier restore step can expose memory to Modbus — every area write is performed while the memory is still sealed (the sealing bit is forced to 0 during area restore per PERSIST-017), and unseal is the last action.
- Design boundary: this task performs the single explicit commit/unseal write only; restore-failure behavior is PERSIST-020. No runtime reseal or alternate commit flag is introduced.
- This is a JR DEV self-check only. It is **not** an independent TEST/VERIFY PASS; independent verification is deferred to PERSIST-022.
- Delivered source commit: `e22459ecdd10433ef21f63df25e3a22a43fb3662` on GitHub main.

## CWAL
PERSIST-018 is delivered on GitHub main at `3eb118f507688d483c5d4985f74fe6a98b531bc3`. PERSIST-019 CODE is now complete and no longer ACTIVE; PERSIST-020 is the next eligible CODE packet. This is JR DEV evidence only, not an independent TEST/VERIFY PASS.
