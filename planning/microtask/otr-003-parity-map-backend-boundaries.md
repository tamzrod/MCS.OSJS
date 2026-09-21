# OTR-003 — Evidence-based UI parity and backend boundary map
Status: PLANNING / UNDER REVIEW. Stage: DESIGN. Owner: ChatGPT. Previous: OTR-002. Next: OTR-004.

Outcome: Produce a component-by-component Electron→OS.js port map using OTR-001/002 evidence. Scope: documentation only: identify actual source/target paths, reuse/copy/adapt decisions, Electron-only dependencies, verified Linux backend contracts, and explicit unknowns. Partition the remaining candidate implementation work into at-most-three-production-file slices; do not pretend unverified paths or endpoints exist.

Non-scope: product code, invented transport, runtime, configuration or service changes; no activation. Acceptance: (1) Each inventoried screen mapped to target and port decision. (2) Each workflow operation mapped to verified backend contract or marked BLOCKED/unknown. (3) Source-backed bounded implementation slices and safe test targets proposed without expanding scope.

Evidence: path/line citations, parity matrix, unresolved questions, source revision. Dependencies: OTR-001 and OTR-002. Size: 0/1/1/1/2 = 5, one design decision artifact; split any independent unresolved architecture decision into a new review task. STOP, no successor execution.
