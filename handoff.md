# Handoff — MCS.OSJS

## Autonomous routing authority
Human has approved the PERSIST-001..022 persistence roadmap and explicitly promoted **OpenHands to JR DEV** for OPERATION CWAL CODE/DISCOVERY work.

OpenHands modes remain separate:
- **JR DEV** — assigned CODE/DISCOVERY implementation plus bounded workflow continuation.
- **Independent JR** — separately assigned TEST/VERIFY only.

The normal operator loop is:

```text
git pull
OPERATION CWAL
```

One invocation executes exactly one product task, may deliver/record it, may prepare the next already-promoted eligible task, then STOPS.

## Completed predecessors

**PERSIST-001 — Persistence Configuration Schema**
CODE COMPLETE at `faa33929429a0382a64b78bc773f0074e11cb6b1`.

**PERSIST-002 — State Sealing Prerequisite Validation**
CODE COMPLETE and delivered on GitHub main at
`fe3bda9e5c863c959500514e8c6083d2514526d1`.

**PERSIST-003 — Derived Persistence RBE Generation**
CODE COMPLETE at `4e8f774eb946cc9a946938e4f3a28eaa13de5c9d`, delivered on main.

**PERSIST-004 — Locked System RBE Behavior**
CODE COMPLETE at `936ab24560d57cbf8aacc8180bc00a2527dcfad5`, delivered on main.

**PERSIST-005 — Memory Range Synchronization**
CODE COMPLETE at `c743174204455afa83bc92b859fd47582a7afdbc`, delivered on main.

**PERSIST-006 — Memory Area Removal Synchronization**
CODE COMPLETE at `9ac48cb23e369fb0384b3d8f00811435d5bf7b69`, delivered on main.

**PERSIST-007 — Persistence Disable Cleanup**
CODE COMPLETE at `feb1e622d38cf8107a1df9ee182e5df17f95790f`, delivered on main.

**PERSIST-008 — User RBE Compatibility**
CODE COMPLETE at `4783e0d6e1d3dbd55b29b8647a115905eb5bc285`, delivered on main.

**PERSIST-009 — RBE ID Collision Handling**
CODE COMPLETE at `19057f02d94ff404b74a9f01348fc8c31235e374`, delivered on main.

**PERSIST-010 — Persistence Configuration UI**
CODE COMPLETE at `190e464f6106c3e640d21b1f9352328447464bef`, delivered on main by this invocation.

PERSIST-001..PERSIST-010 CODE tasks are complete. Their self-check/regression evidence remains JR DEV evidence only; independent TEST/VERIFY is now the PERSIST-011 gate below.

## Current task — ACTIVE (independent JR TEST/VERIFY)

**PERSIST-011 — Persistence Configuration Tests**
Mode / owner: **independent JR (OpenHands JR TEST runner)** — NOT JR DEV
Packet: `workflow/active_work/persist-011-persistence-configuration-tests.md`
Queue: `workflow/active_work/PERSISTENCE_PROMOTION_QUEUE.md`

Goal: independently verify the persistence configuration behavior delivered by PERSIST-001..010 on the pinned committed source checkpoint. No product source edits, no retests to force PASS.

Pinned product checkpoint: `190e464f6106c3e640d21b1f9352328447464bef` (PERSIST-010). Predecessor source commits: `faa3392` (001), `fe3bda9` (002), `4e8f774` (003), `936ab24` (004), `c743174` (005), `9ac48cb` (006), `feb1e62` (007), `4783e0d` (008), `19057f0` (009).

Safe environment: disposable OpenHands sandbox. No production/customer data, no live services, no destructive cleanup. Sandbox-local toolchain preparation is permitted only as declared below.

Freshness gate (before any test): verify `git rev-parse HEAD` is `190e464...` or a descendant that changes only workflow files, and `git status --short` is clean. If product/test paths under test changed after `190e464`, report BLOCKED.

Changed-path allowlist under test:
- `mma2composer/composer.go`, `mma2composer/memory_validation.go`, `mma2composer/persistence_rbe.go`
- `mma2composer/persistence_sealing_validation_test.go`, `mma2composer/persistence_rbe_*_test.go`
- `simulator/device.go`, `simulator/mma2_config.go`, `simulator/persistence_config_schema_test.go`
- `OSJS/src/packages/MCSModbusToolkit/memory-advanced.js`, `OSJS/tests/toolkit-persistence-ui.test.js`

Declared preparation (safe, no manifest change): Go 1.25.x toolchain; in `OSJS`, `npm install --no-save jsdom` (jsdom is already a declared devDependency; do not modify `package.json`/lockfiles).

Exact commands (run once, in order; stop at first failed product gate):
1. `cd mma2composer && go test -mod=readonly ./...`
2. `cd simulator && go test -mod=readonly ./...`
3. `cd replicator && go test -mod=readonly ./...`
4. `cd OSJS && node tests/toolkit-persistence-ui.test.js`
5. `cd OSJS && node tests/toolkit-ui-parity.test.js`
6. `cd OSJS && node tests/toolkit-fc43.test.js`

Expected: each `go test` prints `ok` and exits 0; each `node` test prints its PASS line and exits 0.

Raw evidence: capture each exact command, exit code and full stdout/stderr.

Post-check (read-only, even after FAIL): `git status --short` shows no product/test file modified by JR; report `git rev-parse HEAD`.

Verdict: PASS only if every command exits 0 and the post-check is clean; FAIL if any executed test fails; BLOCKED if the checkpoint is not current or a required toolchain is unavailable within the declared safe scope; INCOMPLETE if the report/post-check cannot be delivered.

Report/transport: return the full report in chat (verdict, tested HEAD, environment, command transcript with exits, post-state, caveats). No product edits. No commit/push unless separately authorized; if authorized, update only a `## JR TEST REPORT — PERSIST-011` section of this handoff.

## Successor routing

PERSIST-012..022 are already human-promoted and remain QUEUED/dependency-gated. PERSIST-011 is the ACTIVE independent JR TEST/VERIFY gate above; it must not be self-certified by JR DEV.

OpenHands JR DEV may autonomously close and deliver CODE/DISCOVERY packets and select the next already-promoted eligible CODE/DISCOVERY packet for the next invocation.

TEST/VERIFY packets switch to independent JR mode and require their exact current packet. JR DEV must not self-certify those gates.

Only BLACK SHEEP WALL edits ICC. No force push, destructive history rewrite, production/operator-data mutation, or scope expansion.
