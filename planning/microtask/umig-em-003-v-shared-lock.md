# UMIG-EM-003-V — VERIFY: Cross-process writer serialization
Status: PLANNED. Stage/owner VERIFY / OpenHands JR. Previous: UMIG-EM-003-T. Next: UMIG-EM-004.

Outcome: in separately approved NEW disposable Docker project, run exactly bounded competing synthetic Simulator/Replicator applies and verify no lost update, preserved foreign owner and at most one serialized config/restart chain per apply. Packet must specify unique token, local daemon provenance, exact harmless requests, maximum attempts/timing, hashes before/after and label-safe `down --remove-orphans` without `-v`. Never use old volumes or customer ports, RBE output or production. Contradiction FAIL/STOP; unavailable environment BLOCKED. Return raw evidence, no fixes. Size 0/2/1/1/1=5; split preflight if unbounded.
