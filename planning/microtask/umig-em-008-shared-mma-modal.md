# UMIG-EM-008 — CODE: Shared MMA Settings modal and draft
Status: PLANNED; human promotion required. Stage/owner CODE / ChatGPT. Previous: UMIG-EM-007-V. Next: UMIG-EM-008-T.

Outcome: in Memory Advanced, implement separate folder-styled MMA Settings dialog using authenticated EM-006 contract; global root settings are not per-device. Keep unsafe RBE TCP/access-events listeners unavailable until separate safety authorization.

Acceptance: (1) modal blocks background editing, keyboard/Close retain unsaved shared draft and no apply; (2) shared Save invokes exactly one revision-checked `mma-apply`, independent shared Discard, truthful conflict/failure and no auto reload/overwrite; (3) disabled unsafe output controls and no Windows IPC, filesystem or service calls. No backend changes or new RBE exposure. Source diff/readback, not test PASS. Size 2/0/1/1/1=5; separate UX vs transport if implementation grows.
