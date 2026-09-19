# UMIG-EM-007 — CODE: Memory per-device Advanced editor
Status: PLANNED; human promotion required. Stage/owner CODE / ChatGPT. Previous: UMIG-EM-006-V. Next: UMIG-EM-007-T.

Outcome: Toolkit-owned Device Definition/Advanced folder strip with RBE Rules, State Sealing and Access Policy editing of canonical Go device fields. Reuse Electron behavior as model, not Electron IPC/component imports.

Acceptance: (1) field mapping, ordered policy/FC/CIDR, sealing disabled/exception and RBE IDs 1..255 with globally unique duplicate allocation, allocated range validation; (2) separate unsaved device draft, Discard and rejected apply preserve old canonical snapshot/foreign data; (3) only existing Go Simulator `load/apply/status` via contract, no auto apply, no RBE network activation. Keep shared MMA modal OUT of scope. Source diff/readback only. Size 2/0/2/1/1=6: split RBE/sealing/policy editor into separately promoted CODE microtasks if more than 3 implementation outcomes before promotion.
