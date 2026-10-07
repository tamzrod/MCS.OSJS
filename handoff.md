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

PERSIST-001..PERSIST-010 CODE tasks are complete. Their self-check/regression evidence remains JR DEV evidence only.

**PERSIST-012 — RBE-Triggered Snapshot Writer**
CODE COMPLETE at `2f18f8196f5497c5435b697b8b1570f59cab2642`, delivered on main.

**PERSIST-013 — Snapshot File Format**
CODE COMPLETE and delivered on GitHub main at `ccccb891b737995f6a24a58e844e5bd1d794b609` via GitHub connector recovery after the OpenHands push credential expired. OpenHands' local-only commits `9343d9e` / `ab27b27` are superseded by this remote recovery and must not be treated as additional product work.

## Completed independent gate

**PERSIST-011 — Persistence Configuration Tests**
Independent JR TEST/VERIFY **PASS** on tested HEAD `313e3be60a29fff6c967d78c0cf677b8b888ed56`, with pinned product checkpoint `190e464f6106c3e640d21b1f9352328447464bef`.

Verified once, in order:
- `cd mma2composer && go test -mod=readonly ./...` — exit 0
- `cd simulator && go test -mod=readonly ./...` — exit 0
- `cd replicator && go test -mod=readonly ./...` — exit 0
- `cd OSJS && node tests/toolkit-persistence-ui.test.js` — exit 0
- `cd OSJS && node tests/toolkit-ui-parity.test.js` — exit 0
- `cd OSJS && node tests/toolkit-fc43.test.js` — exit 0

Freshness/post-check passed: commits after `190e464` were workflow-only, working tree clean, no product/test modification by independent JR. This PASS covers PERSIST-001..010 configuration behavior only; runtime persistence remains for PERSIST-012+.

## Current task — ACTIVE

**PERSIST-014 — Snapshot Manifest / Compatibility Metadata**
Mode / owner: **CODE / OpenHands JR DEV**
Packet: `workflow/active_work/persist-014-snapshot-manifest-compatibility-metadata.md`
Queue: `workflow/active_work/PERSISTENCE_PROMOTION_QUEUE.md`

Execute PERSIST-014 exactly as written:
- bind each snapshot to Port, Unit ID, area, start, count and format version;
- add integrity metadata sufficient to detect corrupted/incomplete payloads;
- reject incompatible changed memory layouts explicitly rather than truncate/remap;
- do not restore or unseal memory;
- run only task-bounded self-checks and bounded in-scope corrective retests;
- commit and non-force push under standing JR DEV authority;
- record changed paths/checks/source checkpoint;
- prepare PERSIST-015 for the next invocation only after genuine PERSIST-014 completion;
- STOP after PERSIST-014.

Recovery note for a reused OpenHands workspace: if local-only commits `9343d9e` / `ab27b27` are still present, they represent the transport-blocked PERSIST-013 attempt already superseded by remote checkpoint `ccccb891`. Do not replay them as new work; reconcile only through bounded safe Git mechanics consistent with standing JR DEV authority.

## Successor routing

PERSIST-015..022 are already human-promoted and remain QUEUED/dependency-gated. PERSIST-011 passed independently; PERSIST-012 and PERSIST-013 are CODE complete; PERSIST-014 is now the sole current ACTIVE JR DEV CODE task.

OpenHands JR DEV may autonomously close and deliver CODE/DISCOVERY packets and select the next already-promoted eligible CODE/DISCOVERY packet for the next invocation.

TEST/VERIFY packets switch to independent JR mode and require their exact current packet. JR DEV must not self-certify those gates.

Only BLACK SHEEP WALL edits ICC. No force push, destructive history rewrite, production/operator-data mutation, or scope expansion.
