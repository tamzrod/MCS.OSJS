# NPE-01 — Native Persistence Contract Realignment

Status: PROMOTED / CURRENT
Stage: CODE
Owner: Codex JR DEV
Previous: NP-01 historical implementation
Next: NPE-02

## Purpose
Realign the MCS-side persistence schema/model with the CURRENT native MMA2 persistence contract before any UI work continues.

## Authoritative contract
Persistence belongs to each MMA2 memory identity.

- `persistence.enabled: true` enables native MMA2 persistence.
- `directory` is OPTIONAL. Omitted means MMA2 uses its native default beside the loaded YAML.
- omitted `ranges` means ALL allocated areas of that memory.
- optional custom ranges must be contained within their allocated area.
- persistence does NOT require State Sealing.
- persistence does NOT require RBE or RBE TCP.
- persistence does NOT require Raw Ingest restore orchestration.
- MMA2 owns disk snapshot, restore, flush, backup and recovery.

## Scope
Update only the shared MCS configuration/model/validation layer needed by Electron to represent the current MMA2 contract correctly.

Correct the stale NP-01 assumptions, especially:
- remove "enabled requires nonempty directory";
- remove persistence-specific State Sealing prerequisite;
- remove persistence-generated/system RBE assumptions from the native model;
- preserve native per-memory enabled/directory/ranges round-trip;
- preserve unrelated memory/RBE/State Sealing configuration.

Do NOT implement UI in this packet.
Do NOT implement persistence runtime outside MMA2.

## Acceptance
Focused schema/model tests prove:
1. enabled + omitted directory is valid;
2. omitted ranges means no duplicate range projection is created by MCS;
3. explicit valid custom ranges round-trip;
4. invalid out-of-area/overlapping ranges reject;
5. persistence and State Sealing validate independently;
6. persistence and RBE validate independently.

Record exact changed paths, commands, outputs and delivered SHA. STOP after delivery.
