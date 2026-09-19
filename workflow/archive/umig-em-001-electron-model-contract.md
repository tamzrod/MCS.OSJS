# UMIG-EM-001 — DESIGN: Current Electron Model Compatibility Contract

Status: COMPLETE — human-approved design on 2026-09-19; archived by ChatGPT. This is DESIGN acceptance, NOT a product test, Go runtime acceptance or production release.
Stage/owner: DESIGN / ChatGPT. Previous: UMIG-006-V archived PASS for the earlier SHA. Next: UMIG-EM-002-T (first separately promoted TEST).

Outcome: Source-level Electron-to-OS.js matrix, Go management/ownership/transaction/failure design, COMMS truth requirements and future microtask sequence documented in `docs/TOOLKIT_ELECTRON_MODEL_CONTRACT.md`. Approved baseline: repository `449cfda29d23406295cd0ccbc35fa8f5d6293681`; Electron layout `ff6846f`, Go COMMS `1ca1740`, Windows-only settings owner at `449cfda`. Human explicitly approved the four architecture/security decisions in conversation on 2026-09-19. Revised sequence splits Go baseline from Toolkit Node/build verification to satisfy one-verification-workflow-per-task.

Approved decisions: Go Simulator Unix v1 shared manager; authenticated load plus separately checked manage-MMA capability and fail-closed apply; shared lock around all producers with revision/CAS/validation and truthful RECOVERY_REQUIRED after uncertain restart; RBE/access-event listeners off until separately authorized. Root network-output editing remains disabled; no production or customer data authorization.

Evidence: source readback of contract, Go Store.Load vs Electron hydration, new Go per-block COMMS shape, Go composer validation, Toolkit v1 allowlist and updated Windows-only installer. No product changes, builds, JR tests, Docker or independent live verification performed as part of this design. Future tasks in Planning are not automatically promoted.
