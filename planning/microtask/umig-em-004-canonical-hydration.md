# UMIG-EM-004 — CODE: Canonical advanced-field hydration
Status: PLANNED. Stage/owner CODE / ChatGPT. Previous: UMIG-EM-003-V. Next: UMIG-EM-004-T.

Outcome: Go Simulator `load` exposes matching effective `policy`, `state_sealing`, `rbe` when absent in raw devices YAML, WITHOUT silently rewriting it. `Store.Load` currently returns only raw document; mirror Electron effective projection in Go with `(port,unit_id)` identity and explicit omitted vs false/empty handling.

Acceptance: (1) legacy device loads existing effective advanced fields, preserving extensions; (2) explicit disable/empty field wins over inherited value; (3) corrupt/unmatched effective settings produce truthful bounded fallback/error, never fixture/default overwrite. Limit to Simulator read projection and Go tests; do not modify runtime apply/Windows/producer config or enable output. Source checkpoint/readback only. Size 1/0/1/1/1=4: keep error behavior bounded.
