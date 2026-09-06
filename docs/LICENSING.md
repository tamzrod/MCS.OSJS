# MCS.OSJS Licensing and Provenance Policy

## Project license

MCS.OSJS original work is intended to be distributed under the Apache License, Version 2.0. The authoritative license text is the repository-root `LICENSE` file.

The Apache-2.0 project license does not erase, replace, or relicense third-party material that is incorporated under another compatible license.

## Donor model

MCS.OSJS is assembled from two primary donor lineages plus external dependencies.

### Donor 1 — Modbus Replicator stack

The donor lineage currently includes:

- `tamzrod/modbus`
- `tamzrod/mma2`
- `tamzrod/modbus-replicator`
- `tamzrod/replicator-stack`

Known repository-level licensing at initialization:

| Repository | Observed license | Treatment |
| --- | --- | --- |
| `tamzrod/modbus` | MIT | Preserve MIT copyright/license notice for imported material. |
| `tamzrod/mma2` | Apache-2.0 | Compatible with the MCS.OSJS Apache-2.0 baseline; preserve required notices. |
| `tamzrod/modbus-replicator` | No root license established during initial audit | Treat licensing as unresolved until provenance/authority is explicitly established. |
| `tamzrod/replicator-stack` | No root license established during initial audit | Audit material individually before transplantation. |

### Donor 2 — Nameless SCADA

`tamzrod/namelessscada` is a project-owned donor, principally for reusable OS.js shell/application infrastructure and other explicitly selected generic components.

Do not treat the entire Nameless SCADA repository as automatically transplantable under the MCS.OSJS license. Distinguish project-owned integration code from upstream OS.js code, third-party packages, themes, icons, fonts, sounds, and other assets.

### External sources

Other repositories and packages may be used when deliberately selected. They are not donors merely because they are technically useful or publicly accessible.

## OS.js boundary

OS.js framework material used under BSD-2-Clause may coexist with MCS.OSJS Apache-2.0 code. The applicable OS.js copyright notice, BSD license conditions, and disclaimer must be retained.

Do not assume every item distributed in an OS.js environment uses the framework's license. Applications, packages, themes, icon sets, fonts, sounds, and transitive dependencies require their own provenance check.

Where OS.js-derived source is modified, retain its applicable upstream license/notice and clearly maintain provenance rather than silently converting the source to Apache-2.0.

## Harvest gate

Before source code, assets, binaries, packages, or substantial derived material are copied into MCS.OSJS:

```text
IDENTIFY SOURCE
→ IDENTIFY EXACT MATERIAL
→ IDENTIFY COPYRIGHT / PROVENANCE
→ IDENTIFY LICENSE
→ IDENTIFY REDISTRIBUTION OBLIGATIONS
→ CHECK COMPATIBILITY
→ RECORD NOTICE
→ HARVEST
```

If provenance or licensing is unknown:

```text
UNKNOWN
→ STOP TRANSPLANTATION
→ INVESTIGATE
→ RESOLVE
→ THEN HARVEST
```

Unknown licensing does not prevent reading, auditing, designing interfaces around, or independently reimplementing behavior where legally appropriate. It blocks copying the unresolved material into the distributable project.

## Agent constraints

JR and other automated agents must not:

- infer a license from repository visibility;
- infer permission from technical dependency relationships;
- remove required copyright, attribution, license, or disclaimer text;
- replace third-party license headers with the MCS.OSJS license;
- assign a new license to a donor repository or donor material without explicit authority;
- import material whose provenance/license is unresolved.

When uncertain, the correct execution result is to stop at the licensing boundary and report what must be resolved.

## Distribution records

`THIRD_PARTY_NOTICES.md` is the project-level registry for incorporated third-party material and donor-license obligations.

As the application becomes distributable, full upstream license texts required for redistribution should be retained in a dedicated license/notice area or otherwise shipped in the form required by the applicable license.

## Scope

This document is an engineering provenance and compliance policy for MCS.OSJS. It is not a substitute for legal advice. Questions involving ambiguous ownership, incompatible licenses, trademarks, patents, or commercial redistribution conditions should be resolved before release.
