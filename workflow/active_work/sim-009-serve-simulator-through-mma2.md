# SIM-009 — Serve Simulator Values Through MMA2

Status: SUPERSEDED — 2026-09-08 by SIM-010 through SIM-017.

## Supersession reason

This task was based on an incorrect architecture assumption that the Simulator owns and activates a managed MMA2 child process. Human clarification establishes the authoritative boundary:

- MMA2 is an independent appliance component and auto-starts on system boot.
- There is no `simbridge` service and no Simulator configuration API.
- The Simulator must not START, STOP, SPAWN, KILL, REPLACE, or otherwise own MMA2.
- The only MMA2 lifecycle/control action the Simulator may request is RESTART, and only after a valid shared MMA2 configuration change is committed.
- Simulator values continue to enter MMA2 through Raw Ingest.

The remaining intended outcomes of SIM-009 are decomposed and replaced by the sequential active tasks SIM-010 through SIM-017. Do not execute this superseded task.