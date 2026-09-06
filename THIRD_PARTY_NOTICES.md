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
