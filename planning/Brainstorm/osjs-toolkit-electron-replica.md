# Brainstorm — OS.js Toolkit as Electron Toolkit Replica

Status: BRAINSTORM / REVIEW ONLY. No microtask promoted; no active execution authorized by this document.
Date: 2026-09-21

## Master goal

Replicate the existing Electron MCS ModbusToolkit in OS.js: same visual UI, navigation, interaction, configuration concepts, and user workflows. Keep the Electron and OS.js applications separately deployed and separately integrated with their respective backends. The Electron implementation is the reference, not merely inspiration. Port or adapt proven code where appropriate; do not redesign an equivalent UI from scratch.

## Intended user journey

Open Toolkit -> select Simulator, Replicator, or MMA memory management -> load actual configuration -> edit in familiar Electron-equivalent UI -> validate complete proposed YAML/configuration without mutating running state -> save safely -> request restart only for the affected runtime/service where supported and authorized -> observe actual applied configuration and service status. Preserve draft/discard, unknown fields, errors, and read-only boundaries. Exact restart capabilities and endpoint contracts must be established from source, not assumed.

## Non-negotiable boundaries

- Same UI and concept, different backend. OS.js remains a distinct deployment; no shared runtime dependency on Electron or Windows endpoints.
- Preserve OS.js shell/start menu/taskbar and the unified Toolkit entry; do not remove legacy packages until a separately reviewed task authorizes it.
- Keep existing configuration schema and validation semantics compatible where evidence confirms compatibility; never silently rewrite or discard unknown YAML fields.
- No speculative APIs, new transport architecture, blanket restart permissions, production/operator-data writes, or service actions during inventory and planning.
- Work in small reviewable microtasks. `planning/microtask/` is under review; moving an approved task to `workflow/active_work/` is the authorization gate. One assigned task per agent run; implement, self-test, report, STOP. Independent JR evidence is separate.
- Only BLACK SHEEP WALL may update ICC. This brainstorm does not update ICC, handoff, active_work, or product code.

## Proposed sequential microtask plan — candidates, not yet approved

01. Reference inventory (inspection only): identify Electron Toolkit entry point, actual rendered screens, navigation, dialogs, CSS/assets, configuration editor modules, validation/save/restart/status flows, and exact source paths. Record screenshots or UI references if available. Deliver an evidence-linked feature/screen inventory, not a guessed architecture.
02. OS.js baseline inventory (inspection only): identify existing unified Toolkit entry, UI files, existing config/backend capabilities, deployment boundaries and test harness. Record reusable pieces and mismatches against Electron.
03. Parity map and port boundary (design only): for each Electron UI component/workflow, identify copy/adapt/reimplement decisions, OS.js target file, Electron-specific dependency, required Linux backend operation and verified evidence. Flag unknown contracts; no invented endpoint.
04. UI shell replica (CODE): port the Electron Toolkit's actual outer layout, navigation, visual assets and styles into the unified OS.js Toolkit. Keep backend calls stubbed/read-only as explicitly defined in the task; test rendering and navigation against the reference.
05. Simulator editor UI replica (CODE): port its existing forms, tabs, draft/discard interactions and visual states. No new persistence or runtime effects yet; targeted UI tests.
06. Replicator editor UI replica (CODE): same approach using Electron's actual editor and shared reusable UI primitives; targeted UI tests.
07. MMA memory UI replica (CODE): port the existing memory management views and editor behaviors; targeted UI tests.
08. Config read adapter (CODE): map verified Linux/OS.js backend to the existing YAML/configuration read contract, with read-only tests and explicit error states.
09. Edit/validation parity (CODE): reuse verified validation and serialization logic; test preservation of unknown fields, invalid inputs, draft/discard and no mutation on validation failure.
10. Safe save adapter (CODE): implement the verified OS.js/Linux write path with authorization, revision/conflict protection and atomicity where supported; test with synthetic/disposable fixtures only.
11. Restart adapter (CODE): wire only supported, explicitly scoped MMA/Simulator/Replicator restart operations; no global or speculative service control. Verify acknowledgement and failure handling with disposable targets.
12. Applied-state/status parity (CODE): render independently observed service/configuration state; never infer green or success from a successful save/restart request alone.
13. UI parity review (independent VERIFY): compare Electron and OS.js side by side for each inventoried screen, navigation and draft/error state; record differences and create bounded follow-up microtasks only where evidence warrants.
14. End-to-end acceptance (independent VERIFY): on a safe synthetic environment, load -> edit -> validate -> save -> restart -> confirm applied state, plus failure/conflict/discard scenarios; no production data.

## Review gates before decomposition

Inventory must establish actual Electron paths and behaviors; do not treat candidate steps 04-14 as source-verified scope. Split each candidate further to fit the microtask rule (at most three production files per task, plus tests). Name exact paths, test commands, agent owner, evidence, allowed side effects and STOP in each reviewed microtask. Do not activate the entire sequence in advance. No work should depend on a fabricated PASS from previous OSJT tasks.

## Open questions to resolve from source

- Which Electron UI modules/assets can be copied directly without importing Electron/Windows APIs?
- Which configuration files and schemas are actually shared, and which require translation?
- What exact Linux/OS.js backend endpoints and restart mechanisms exist today for each runtime?
- How is applied-state confirmation obtained independently of a save/restart acknowledgement?
- Which screen-level parity details are intentional platform differences rather than regressions?

## Definition of done

A user can open the unified OS.js Toolkit and see the same layout, screens, labels, tabs, interactions and configuration concepts as the reference Electron Toolkit. On supported Linux backends, equivalent load/edit/validate/save/restart/verify workflows behave consistently with genuine status/error reporting. Both applications build and deploy independently; parity is demonstrated by screen comparisons, targeted tests and safe end-to-end evidence.


## Candidate enhancement — Docker Network CIDR Detection for Access Policy UI

Status: BRAINSTORM ONLY. UI/configuration convenience; not authorized implementation.

### Objective

Improve the Access Policy `Source IP / CIDR` selector so the MCS Modbus Toolkit can discover Docker networks available on the host and offer their IPAM subnets as selectable source CIDRs.

This capability belongs entirely to the Toolkit/application side. MMA2.0 must remain Docker-unaware and continue receiving, storing, and matching ordinary IP/CIDR values only.

### Desired selector behavior

Keep the normal choices available:

- All IPv4
- All IPv6
- Localhost IPv4
- Localhost IPv6
- Custom

When Docker discovery succeeds, insert detected network/subnet suggestions such as:

- `Docker: bridge (172.17.0.0/16)`
- `Docker: mcs_default (172.30.10.0/24)`
- `Docker: monitoring_default (172.31.0.0/24)`

The displayed network name is UI metadata only. Selecting a Docker suggestion must insert/write only its resolved CIDR, for example `172.30.10.0/24`.

MMA configuration remains ordinary CIDR data:

```yaml
source_ip:
  - 172.30.10.0/24
```

Never persist Docker network names such as `mcs_default` as access-policy source values.

### Discovery requirements

- Query the local Docker Engine from the Toolkit/application side using the existing host/system-operation architecture where possible.
- Enumerate all usable Docker bridge networks.
- Obtain each network name and each IPAM subnet/CIDR.
- Expose multiple IPAM subnets from the same network independently.
- Support both IPv4 and IPv6 CIDRs as reported by Docker.
- Do not assume Docker subnets use any particular private range such as `172.x.x.x`.
- Keep Docker discovery isolated from MMA policy/configuration semantics.

### Refresh behavior

Use the simplest refresh mechanism consistent with the existing UI architecture. At minimum, rediscover Docker networks when the Access Policy editor/device configuration is opened or explicitly refreshed. Do not add continuous polling unless the application already has a suitable refresh mechanism.

### Failure behavior

Docker discovery is optional enrichment. If Docker is absent, stopped, inaccessible, permission-denied, or network inspection fails:

- Access Policy editing must continue normally.
- All normal non-Docker source choices must remain usable.
- Docker availability must never become a Toolkit requirement.
- Discovery failure must not corrupt or block policy editing.

### Existing configuration behavior

Existing CIDR entries remain ordinary CIDRs regardless of whether they originally corresponded to a Docker network. If `172.30.10.0/24` is stored and that Docker network later disappears, preserve the value unchanged and allow it to be displayed/edited as a custom CIDR.

### Security and architecture boundary

```text
Docker Engine
    ↓
MCS Toolkit detects networks
    ↓
UI presents friendly Docker network choices
    ↓
Explicit user selection resolves to CIDR
    ↓
MMA configuration stores CIDR
    ↓
MMA performs normal source-IP/CIDR matching
```

MMA2.0 remains completely unaware of Docker network names, container IDs, Compose projects, Docker sockets, and Docker APIs.

Docker networks are suggestions only. Discovery must never automatically trust, add, or modify policy sources.

### Explicit non-goals

- Do not modify MMA2.0 networking behavior.
- Do not modify MMA2.0 access-control matching.
- Do not add Docker dependencies to MMA2.0.
- Do not change MMA configuration format.
- Do not store Docker network names in MMA YAML.
- Do not automatically trust all detected Docker networks.
- Do not automatically add Docker CIDRs to existing policies.
- Do not change an existing policy without explicit user selection.
- Do not redesign the Access Policy editor.
- Do not refactor unrelated code.

Existing Add source / Remove source behavior must remain intact.

### Source inspection required before any implementation microtask

Before promoting this idea to executable work:

1. Inspect the existing Access Policy UI implementation.
2. Locate where All IPv4 / All IPv6 / localhost / Custom choices are generated.
3. Identify the existing application-side mechanism for host/system operations.
4. Reuse that architecture instead of introducing a parallel Docker subsystem.
5. Establish the smallest clean integration point that keeps discovery separate from policy/configuration logic.

### Candidate validation checklist

1. Open Access Policy.
2. Existing All IPv4 / All IPv6 / Custom behavior still works.
3. Detected Docker networks appear in the selector.
4. Network name and CIDR are visible to the user.
5. Selecting a Docker network inserts its CIDR.
6. Saved MMA configuration contains only the CIDR.
7. Reloading preserves the CIDR.
8. Removing the Docker network does not corrupt an existing policy.
9. Docker daemon unavailable does not break the editor.
10. Multiple Docker networks and multiple subnets are handled correctly.
11. IPv4 and IPv6 CIDRs are preserved without incorrect transformation.
12. No MMA2.0 source code or access-control semantics are changed.

Implementation discipline: make the smallest clean change necessary after source inspection. This brainstorm does not authorize code changes or microtask promotion.
