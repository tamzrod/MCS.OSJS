> SUPERSEDED by human backlog reset. Historical only; no executable authority or new PASS claim.

# UMIG-EM-009 — CODE: Measured Replicator COMMS indicators
Status: PLANNED; human promotion required. Stage/owner CODE / ChatGPT. Previous: UMIG-EM-008-V. Next: UMIG-EM-009-T.

Outcome: map current Go per-block observed Network/TCP/Modbus/MMA2 evidence into OS.js LEDs without inventing probes or inferring each layer from `source_status`.

Acceptance: (1) selected canonical device/name, block identity/index/count, enabled/running and recent `last_poll/observed_at` required for non-UNKNOWN; stale/older/missing/future/cross-device responses immediately UNKNOWN; (2) real ERROR/WARNING/OK per layer with ERROR > WARNING > UNKNOWN > OK aggregation and details; network success means TCP connect, NOT ping; (3) new activity pulses only on fresh acknowledged evidence and no repaint after teardown, respect reduced motion. No new Go poll, IPC, service health or production. Source diff/readback only. Size 1/0/2/1/1=5; split visual strip from model if needed.
