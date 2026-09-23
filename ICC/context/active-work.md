# Active Work — workflow routing context

Parent: L0-project
Zoom Out: L0-project
Zoom In: none
Connector: osjs-shell (Toolkit implementation context; follow only if required by the task)
Source dependencies: `handoff.md`, `workflow/active_work/OTR_PROMOTION_QUEUE.md`, `workflow/active_work/otr-*.md`, `workflow/active_work/README.md`.

## Verified remote snapshot

Source branch: `main`, HEAD `45cd3d2d81831306c6943f844e49b7deccbfc5ad` (2026-09-23 review). This is a remote committed-source snapshot; the local working tree and uncommitted overlay are **not inspected** and must not be described as clean.

At this snapshot `handoff.md` names **OTR-001A** as the current read-only discovery task, with `workflow/active_work/otr-001a-electron-shell-inventory.md` as its packet. It identifies OTR-001B, OTR-001C, OTR-002A, OTR-002B, OTR-003A, and OTR-003B as candidates in sequence, subject to real predecessor evidence and packet authorization; OTR-004 through OTR-014 are non-executable parent topics. Handoff prohibits product edits, builds, services, ICC changes, commits, and pushes during OTR-001A. It does not establish that a task was actually completed or that current local state matches the remote snapshot.

The current handoff supersedes the old OSJT-010 assertion previously cached in this file. To determine the actionable task in any checkout, inspect the checkout's actual handoff, task packet, queue and evidence under their normal authority rules. Do **not** promote work or infer PASS from this context. If the task packet and handoff disagree, report the conflict rather than inventing a resolution.

## Maintenance

This node is a semantic summary, not an authoritative queue. If a changed `handoff.md` or task/queue dependency intersects this node, refresh it from **only those changed paths**. Do not refresh Simulator or Replicator product contexts merely because workflow state changed. The node is current only for the cited committed snapshot, not for another branch or a dirty local overlay.
