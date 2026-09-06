# Donor Licensing and Provenance

Baseline commit: cb869da3caeed040478e67ecb0a01c5f93b3a66b
Working tree: clean
Source dependencies: docs/LICENSING.md, THIRD_PARTY_NOTICES.md, LICENSE

## Project License

MCS.OSJS original work: Apache-2.0 (authoritative text in repository-root LICENSE). Does not erase/relicense third-party material under another compatible license.

## Donor Lineages

Donor 1 = Modbus Replicator stack:`tamzrod/modbus` (MIT, preserve notice),`tamzrod/mma2` (Apache-2.0, compatible, preserve required notices),`tamzrod/modbus-replicator` (no root license; unresolved until provenance/authority established),`tamzrod/replicator-stack`(no root license; audit material individually before transplant). Donor 2 = Nameless SCADA: primary donor for OS.js shell/generic desktop infrastructure; whole repo NOT auto-transplantable under Apache-2.0; distinguish integration code from upstream OS.js, third-party packages, themes, icons, fonts, sounds, assets.

## OS.js Boundary

Framework material BSD-2-Clause;may coexist with Apache-2.0 code while retaining copyright/license/disclaimer. Do not assume every item in OS.js env is BSD-2-Clause; apps/packages/themes/icons/fonts/sounds/deps need own provenance check. Modified OS.js source keeps upstream license/notice; provenance maintained;

## Harvest Gate (before copying in

IDENTIFY SOURCE -> IDENTIFY EXACT MATERIAL -> IDENTIFY COPYRIGHT/PROVENANCE -> IDENTIFY LICENSE -> IDENTIFY REDISTRIBUTION OBLIGATIONS -> CHECK COMPATIBILITY -> RECORD NOTICE -> HARVEST. Unknown license: STOP TRANSPLANTATION -> INVESTIGATE -> RESOLVE -> THEN HARVEST. Unknown licensing blocks copying into distributable project, but does not block reading/auditing/designing interfaces around/independent reimplementation.

 Agent constraints: never infer license from visibility/dependency relationships; never remove required notices; never replace third-party headers with Apache-2.0; never assign license to donor material without authority; never import unresolved-provenance material. When uncertain: stop at licensing boundary, report what must be resolved.



## Third-Party Notices File

THIRD_PARTY_NOTICES.md records provenance/redistribution obligations; inclusion under Apache-2.0 does not relicense third-party material. Update whenever new component incorporated or provenance/version/licensing changes; preserve full license texts where required alongside distribution.