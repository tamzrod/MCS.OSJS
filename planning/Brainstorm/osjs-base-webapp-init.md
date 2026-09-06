# Brainstorm: OS.js Base Web App Initialization

Status: brainstorm material. No architectural decision is made here. This topic frames intent, donor sourcing, port assignment, and open questions for later planning.

## Intent Framing

Initialize the OS.js base web app as the appliance presentation shell, sourcing it from the donor, copying the color schemes and sample app from the donor source, assigning a random port for the OS.js management/UI endpoint, and documenting that port into the documentation.

aligned with the project's staged construction, this topic touches donor inventory construction stages 2-3,and UI planning stage 5, and is preparation for later microtasks formulation.

It does not authorize any copying or code writing.



## Donor Sourcing (Nameless SCADA,

Per project identity: Nameless SCADA is the primary expected donor for the OS.js shell and generic desktop/application infrastructure.

The exact source path, branch, commit, and package set must be identified during the donor inventory stage,: not inferred here.



The OS.js licensing harvest gate applies before any material is copied:

- OS.js framework material is BSD-2-Clause and may coexist with Apache-2.0 code, retaining its copyright/license noticeand
- Do not assume the sample app, color schemes, themes, icons, fonts, sounds, or other assets share the OS.js framework license,: each requires its own provenance/license audit:
 IDENTIFY SOURCE ->
 IDENTIFY EXACT MATERIAL ->
 IDENTIFY COPYRIGHT / PROVENANCE ->
 IDENTIFY LICENSE ->
 IDENTIFY REDISTRIBUTION OBLIGATIONS ->
 CHECK COMPATIBILITY ->
 RECORD NOTICE (:docs/LICENSING.md.
- Record donor additions in THIRD_PARTY_NOTICES.md before importing.



## Scope Candidates

Brainstorm only. No selection is implied by listing:

- Initialize a base OS.js web app shell from the donor,(minimal application loading the shell..
- Copy the sample app from the donor source,: shipped as-is as reference, or as a structural/pattern template for our application? This distinction matters for licensing and for how much of it is retained.


- Copy the color schemes and theme configuration from the donor sample:(application-level theming, OS.js desktop theme packages, or bundled assets? Affects what is copied and what licenses apply.

- Assign a random port for the OS.js management/UI endpoint.

- Document the chosen port into the documentation,: candidate home: docs/NETWORK_EXPOSURE.md and a configuration/README reference.
- Record donor provenance and license notices for every copied asset (THIRD_PARTY_NOTICES.md.



## Port Assignment (Random,

Repository authority fixes the externally exposed surface: exactly one OS.js management/UI port exposed, while Orchestrator and Modbus Replicator remain internal-only(:docs/NETWORK_EXPOSURE.md. The directive explicitly leaves exact TCP port numbers to the architecture tasks, so this brainstorm frames the assignment dynamics, it does not settle them.

"Assign a random port" raises select-dynamics questions::

- Random once at scaffold/init time,: then recorded as a fixed documented default for the deployment?
- Random at each boot and surfaced dynamically,: and in that case, what exactly is documented?
- Who selects,: and how is availability validated against other appliance ports,: host ports,: and the ephemeral range?
- How does a random assignment interact with the open port-configuration-authority question in port-config-and-socket-deployment.md?
- "Document the port into the documentation" implies the chosen value becomes a recorded artifact,: which leans toward a one-time-select-then-document model,: but that stays an open question until planned.



## Open Questions

- From which exact donor path/commit do we copy the OS.js base,: sample app,: and color schemes?
- Are the sample app and color schemes shipped as-is or adapted,: and does each adapted asset get its own provenance/license notice?
- Which OS.js version and packages does the donor use,: and are transitive dependencies audited for license/compatibility?
- Does "base web app" mean purely the frontend OS.js shell for now,: or does it already include backend stem/(e.g., Orchestrator wiring): Keep separation of responsibilities explicitly planned.:
- Random port: selected once and documented,: or dynamic per boot,(and if dynamic,: what exactly gets documented?
- Where does the documented port live: NETWORK_EXPOSURE.md,: a config file,: a README,: or a combination?
- How is the random port validated at assignment time to avoid collisions:(with ephemeral ports,: host adapter ports,: and the configuration-dependent MMA2 Modbus TCP ports).



## Constraints from Repository Authority

- External exposure directive,: the appliance exposes one OS.js management endpoint and one or more MMA2 Modbus TCP endpoints; Orchestrator and Modbus Replicator remain internal-only,and donor code must not be re-exposed merely because it listens on TCP (:docs/NETWORK_EXPOSURE.md.
- Licensing harvest gate,: stop at the licensing boundary until provenance is established(:docs/LICENSING.md.When uncertain, the correct execution result is to report what must be resolved,: not to assume.



## Candidate Directions (not decisions)

- Vendor the donor-derived OS.js base into a controlled project location with a recorded provenance manifest entry,: rather than ad-hoc copies.
- Choose the random port once at scaffold time,: validate it against the known port landscape,: and record it as the documented default alongside the directive's one-OS.js-port rule.
- Keep the base OS.js app minimal,:with the donor sample app and color schemes captured as reference/:theme material,: not necessarily shipped as the product UI.



## Out of Scope Here

- Choosing the design(:left for the architecture planning stage:
- Microtask decomposition(:comes after planning assumes shape:
- Writing or running candidate code now(:brainstorm grants no authority:
- Resolving licensing questions by assumption,:stop at the licensing boundary until provenance/authority is established(:docs/LICENSING.md:
