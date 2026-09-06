# MCS.OSJS — Third-Party Notices

MCS.OSJS is distributed under the Apache License, Version 2.0, except for components, source files, assets, libraries, and other materials that retain their own licenses as identified here or in their accompanying license files.

This document records third-party provenance and redistribution obligations. Inclusion in MCS.OSJS does not relicense third-party material under Apache-2.0.

## OS.js

MCS.OSJS may incorporate or derive portions of its desktop/application environment from the OS.js project.

OS.js material must retain the copyright notices, license terms, and disclaimer supplied by the applicable upstream OS.js repository or package. OS.js code identified as BSD-2-Clause remains under BSD-2-Clause when incorporated or modified in this repository.

Do not replace upstream OS.js license headers or notices with MCS.OSJS Apache-2.0 headers.

Before importing an OS.js package, application, theme, icon set, font, sound, asset, or other separately distributed component, verify its exact upstream license. The license of the OS.js framework must not be assumed to cover separately licensed dependencies or assets.

## Project-owned donor repositories

MCS.OSJS is expected to reuse or derive material from repositories controlled by the MCS.OSJS project owner, including:

- `tamzrod/modbus`
- `tamzrod/mma2`
- `tamzrod/modbus-replicator`
- `tamzrod/replicator-stack`
- `tamzrod/namelessscada`

Known licensing at project initialization:

- `tamzrod/modbus` — MIT License. Imported MIT-licensed material must retain its original copyright and MIT license notice.
- `tamzrod/mma2` — Apache License, Version 2.0.
- Other project-owned donor material — provenance and applicable licensing must be recorded when material is actually harvested. Ownership of a donor repository does not authorize an agent to silently assign or change a license.

## External repositories and dependencies

Every externally sourced component must be reviewed before incorporation. Record, at minimum:

1. upstream project or repository;
2. exact component or material imported;
3. version, tag, commit, or other stable source reference when practical;
4. copyright holder or upstream attribution when supplied;
5. applicable license;
6. required notices or redistribution conditions.

Unknown licensing blocks transplantation into MCS.OSJS until resolved. Investigation and evaluation are allowed; copying unknown-license material into the distributable application is not.

## Agent rule

Donor status is not license authority.

No automated agent may infer that public availability, repository ownership, dependency use, or technical compatibility implies license compatibility. Preserve upstream notices and stop at unresolved provenance or licensing boundaries.

## Maintenance

Update this document whenever a new third-party component is incorporated or an existing component's provenance, version, or licensing changes. Where a license requires distribution of its full text, preserve that text in an appropriate license/notice file alongside the distribution.

## Nameless SCADA — OS.js desktop shell slice

**Source:** `tamzrod/namelessscada` → `desktop/osjs-prototype/` (project-owned donor)

**Exact material harvested** into `OSJS/` at commit `13e66df0f4f2137e436b9dbf0f48cc2f0e85f67c` (main:

- OS.js client bootstrap + config + scss + html templateand custom panel/session/autostart/taskbar/desktop-icon providers.

- OS.js server bootstrap + config + a minimal `/healthz` route and a classic-icon route serving the local icon theme dist.


- Local packages: `NamelessClassicIcons` (original Windows-2000-inspired greyscale SVG icon set;and `NamelessWorkstationTheme` (classic workstation theme;, plus `scripts/build-local-packages.js` (deterministic local package build harness..
- Build/run scaffolding: `webpack.config.js`, `.babelrc`, `package.json`, `Dockerfile`, `.gitignore`.

**License:** the donor shell prototype package declaration `"license": "BSD-2-Clause"` applies to the harvested slice (donor repo ships no LICENSE file; upstream OS.js framework packages retain their own BSD-2-Clause licenses,. BSD-2-Clause requires retaining the copyright notice, license conditions,and disclaimer; this notice records the obligation..

**Retained notices:** no per-file copyright/license headers existin the donor slice; the package.json license declaration was retained in the harvested copy. A copy of the BSD-2-Clause text should be shipped alongside any redistribution of OSJS/ (see "License texts" note below..

**Excluded from harvest:** all donor application packages (AutoStart,, Dnp3Editor,, InfluxExplorer,, Ingestor,, ModbusEditor,, TagManager,, TaskbarSettings,, TestApp,,, sample/test apps, SCADA/Governor/Ingestor/Modbus/DNP3/Tag-Manager server+client integration providers,bridge/realtime/data providers, donor domain docs,tests, wallpapers,and other assets without build/run referencesin the kept shell.



## License texts

License texts required for redistribution of the harvested OS.js slice:the BSD-2-Clause text (declared by the donor `desktop/osjs-prototype/package.json`)and the OS.js framework's BSD-2-Clause texts supplied with the corresponding npm packages.. These texts should be retained in a dedicated license/notice area alongside the distribution as the application becomes distributable per the licensing policy above..
