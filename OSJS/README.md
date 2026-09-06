# MCS.OSJS — Basic OS.js Desktop Shell

A runnable basic [OS.js](https://www.osjs.org/) desktop shell for MCS.OSJS,
carved from the *Nameless SCADA* `desktop/osjs-prototype/` neutral shell.
It contains **no SCADA applications and no SCADA backend coupling**:only the
basic desktop environment (panel/taskbar, start menu, desktop icons, session
snapshot, auto-start launcher, the Nameless Workstation theme and Nameless Classic
icon sheet).

## Provenance (donor harvest

- Source: `tamzrod/namelessscada` → `desktop/osjs-prototype/`
- Donor commit: `13e66df0f4f2137e436b9dbf0f48cc2f0e85f67c` (main)
- Donor licensing: base shell material declared `BSD-2-Clause` (see package.json;
  OS.js framework packages retain their own BSD-2-Clause licenses).
- Excluded: all donor application packages, sample apps, SCADA/Governor/
  Ingestor/Modbus/DNP3/Tag Manager integration, donor domain docs,and tests.
 Recorded here in THIRD_PARTY_NOTICES.md per the harvest gate.



## Layout

OSJS/
  src/client/        OS.js client bootstrap + config + scss + html template
  src/server/        OS.js server bootstrap + config + /healthz + classic icon route
  src/packages/NamelessClassicIcons    Original era-style SVG icon theme
  src/packages/NamelessWorkstationTheme  Classic workstation theme
  webpack.config.js  Client bundler config
  package.json       @osjs/* v3 dependencies
  Dockerfile         Node 16 build/run image

## Run

Local, Node 10-16 (OS.js v3 builds with webpack 4 / MD4; on Node 17+
set `NODE_OPTIONS=--openssl-legacy-provider`):



```bash
cd OSJS
npm install
npm run build:local-packages
npm run package:discover
NODE_OPTIONS=--openssl-legacy-provider npm run build
npm run serve
```

The OS.js management/UI endpoint listens on port **18209** by default
(`src/server/config.js`; override with `PORT` env var)。 It is recorded in
`docs/NETWORK_EXPOSURE.md`。

## Docker



```bash
sudo docker build -t mcs-osjs-shell .
sudo docker run --rm -p 18209:18209 mcs-osjs-shell
```
