# Donor Licensing and Provenance

Baseline commit: 500376cfb5c222298aadfcf035aad0af0a635773
Working tree: clean
Source dependencies: docs/LICENSING.md, THIRD_PARTY_NOTICES.md, LICENSE
Parent: L0-project
Zoom In:(none; leaf node(
Zoom Out: L0-project

## Project License

MCS.OSJS original work: Apache-2.0 (authoritative text in repository-root LICENSE). Does not erase/relicense third-party material under another compatible license.

## Donor Lineages

Donor 1 = Modbus Replicator stack:  `tamzrod/modbus` (MIT, preserve notice(, `tamzrod/mma2` (Apache-2.0, compatible, preserve required notices(, `tamzrod/modbus-replicator` (no root license; unresolved until provenance/authority established(, `tamzrod/replicator-stack` (no root license; audit material individually before transplant(.
Donor 2 = Nameless SCADA:  primary donor for OS.js shell/generic desktop infrastructure; whole repo NOT auto-transplantable under Apache-2.0; distinguish integration code from upstream OS.js, third-party packages, themes, icons, fonts, sounds, assets. Known:  `tamzrod/namelessscada` → `desktop/osjs-prototype/` declared BSD-2-Clause (project-owned donor(.

## Harvest Gate (before copying in

IDENTIFY SOURCE -> IDENTIFY EXACT MATERIAL -> IDENTIFY COPYRIGHT/PROVENANCE -> IDENTIFY LICENSE -> IDENTIFY REDISTRIBUTION OBLIGATIONS -> CHECK COMPATIBILITY -> RECORD NOTICE -> HARVEST. Unknown license:  STOP TRANSPLANTATION -> INVESTIGATE -> RESOLVE -> THEN HARVEST. Unknown licensing blocks copying into distributable project, but does not block reading/auditing/designing interfaces around/independent reimplementation.

 Agent constraints:  never infer license from visibility/dependency relationships; never remove required notices; never replace third-party headers with Apache-2.0; never assign license to donor material without authority; never import unresolved-provenance material. When uncertain:  stop at licensing boundary, report what must be resolved..

## Executed Harvest:  OS.js desktop shell slice (recorded in THIRD_PARTY_NOTICES.md

Harvested `tamzrod/namelessscada` → `desktop/osjs-prototype/` at commit `13e66df0f4f2137e436b9dbf0f48cc2f0e85f67c` into `OSJS/`:

- client bootstrap+config+scss+html template, custom panel/session/autostart/taskbar/desktop-icon providers;
- server bootstrap+config, minimal /healthz route, classic-icon route serving local icon dist;
- local packages NamelessClassicIcons (original Windows-2000-inspired greyscale SVG set( + NamelessWorkstationTheme (classic workstation theme(, + scripts/build-local-packages.js determininistic local package build harness;
- build/run scaffolding (webpack.config.js, .babelrc, package.json, Dockerfile, .gitignore(.

License:  donor shell package declaration `BSD-2-Clause` applies to harvested slice (donor repo ships no LICENSE file; upstream OS.js framework packages retain their own BSD-2-Clause licenses(. BSD-2-Clause requires retaining copyright notice, license conditions, disclaimer; recorded obligation. Retained notices:  no per-file copyright/license headers exist in donor slice; package.json license declaration retained; BSD-2-Clause text should ship alongside any redistribution of OSJS/. Excluded:  all donor application packages, sample/test apps, SCADA/Governor/Ingestor/Modbus/DNP3/Tag-Manager integration providers, bridge/realtime/data providers, donor domain docs, tests, wallpapers, and assets without build/run references in the kept shell..

## License Texts

License texts required for redistribution of harvested OS.js slice:  BSD-2-Clause text (declared by donor desktop/osjs-prototype/package.json( and OS.js framework's BSD-2-Clause texts supplied with corresponding npm packages. Retain in dedicated license/notice area alongside distribution as application becomes distributable..

## Third-Party Notices File

THIRD_PARTY_NOTICES.md records provenance/redistribution obligations; inclusion under Apache-2.0 does not relicense third-party material. Update whenever new component incorporated or provenance/version/licensing changes; preserve full license texts where required alongside distribution..
