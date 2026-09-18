# OS.js Base Desktop Shell (Implemented

Baseline commit: ee19b8a
Working tree: clean
Source dependencies: OSJS/README.md, OSJS/package.json, OSJS/Dockerfile, OSJS/webpack.config.js, OSJS/scripts/build-local-packages.js, OSJS/src/server/config.js, OSJS/src/server/index.js, OSJS/src/server/providers/health.js, OSJS/src/server/providers/classic-icons.js, OSJS/src/client/config.js, OSJS/src/client/index.js, OSJS/src/client/index.ejs, OSJS/src/client/providers/nameless-app-shortcuts.js, OSJS/src/packages/MCSModbusToolkit/, OSJS/src/packages/NamelessClassicIcons/metadata.json, OSJS/src/packages/NamelessWorkstationTheme/metadata.json
Parent: L0-project
Zoom In:(none; leaf node(
Zoom Out: L0-project

## What Exists

A runnable basic OS.js desktop shell for MCS.OSJS, carved from the Nameless SCADA `desktop/osjs-prototype/` neutral shell (donor commit 13e66df0f4f2137e436b9dbf0f48cc2f0e85f67c(. No SCADA applications, no SCADA backend coupling; only the basic desktop environment (panel/taskbar, start menu, desktop icons, session snapshot, auto-start launcher( + Nameless Workstation theme + Nameless Classic icon sheet.

 Provenance/license recorded in THIRD_PARTY_NOTICES.md per harvest gate; complete license texts still to be shipped alongside redistribution (see donor-licensing context(..

## Layout

- `OSJS/src/client/` — OS.js client bootstrap + config + scss + html template (index.ejs( plus custom providers/panels (autostart, desktop-icons, icon-fallback, session, session-snapshot, taskbar, menu-panel-item, start-menu-band/dnd/layout(.
- `OSJS/src/server/` — OS.js server bootstrap + config + `/healthz` liveness route + classic-icon route serving local icon theme dist.
- `OSJS/src/packages/NamelessClassicIcons` — original Windows-2000-inspired greyscale SVG icon theme (icons/*.svg, main.scss, metadata.json, webpack.config.js(.
- `OSJS/src/packages/NamelessWorkstationTheme` — classic workstation theme (index.scss, src/theme.js, metadata.json, webpack.config.js(.
- `OSJS/scripts/build-local-packages.js` — deterministic local package build harness.
- Build/run scaffolding:  webpack.config.js, .babelrc, package.json, Dockerfile, .gitignore.

## Key Configuration Facts (server

- `OSJS/src/server/config.js`:  root=OSJS/, port = Number(PORT( || 18209 (OSJS-003 deterministic unoccupied management/UI port;public=dist/; VFS root = {OSJS_DATA_DIR||cwd}/vfs; session store connect-loki db behaveside vfs under same data dir; OSJS_DATA_DIR relocates persistence to stable data dir (container mounts persistent volume at /data(deploy/docker-compose.yml(.
- `OSJS/src/server/index.js`:  registers Core, Package, VFS, Auth, Settings (adapter fs — per-user settings persisted to VFS home home:/.osjs/settings.json(, HealthRouteProvider, ClassicIconsRouteProvider; boot/shutdown via SIGTERM/SIGINT; neutral shell (no SCADA backend(.
- `OSJS/src/server/providers/health.js`:  GET /healthz -> {status:ok,shell:neutral}; never echoes credentials; skipped gracefully if osjs/express unavailable.

## Key Configuration Facts (client

- `OSJS/src/client/config.js`:  auto-login demo/demo; locale en_EN (deterministic regardless of host browser locale(; settings adapter server (per-user VFS home(; desktop theme NamelessWorkstationTheme, icons NamelessClassicIcons; background override via style color + theme root background-image none (fixes stock wallpaper/purple defect(; bottom panel items menu (icon URL /icons/NamelessClassicIcons/icons/start-here.svg(, windows, tray, clock..
- Menu icon rendered verbatim as <img src>, so must be a URL not a bare icon name (fs.icon resolver not applied thus(.

## Build / Run

Local (Node 10-16; OS.js v3 builds with webpack 4/MD4; Node 17+ set NODE_OPTIONS=--openssl-legacy-provider(:
```bash
cd OSJS
npm install
npm run build:local-packages
npm run package:discover
NODE_OPTIONS=--openssl-legacy-provider npm run build
npm run serve
```
Docker (Node 16 bullseye image; two-stage layer caching, local packages built before discovery so dist/ exists; EXPOSE 18209; CMD npm run serve(:
```bash
sudo docker build -t mcs-osjs-shell .
sudo docker run --rm -p 18209:18209 mcs-osjs-shell
```
Deployment:  deploy/docker-compose.yml (osjs-shell image, ports OSJS_PORT:-18209:18209, PORT=18209, OSJS_DATA_DIR=/data, osjs-data named volume, /healthz healthcheck on 127.0.0.1:18209(. Note:  standalone shell compose currently uses Docker `ports:` publishing — final appliance must switch to host networking per network-exposure directive (this publishing exists for the standalone shell deployment only(..

## Added since the harvested shell (delta to `ee19b8a`)

- Pinned desktop shortcut provider `OSJS/src/client/providers/nameless-app-shortcuts.js`, registered
  in `OSJS/src/client/index.js`. It injects one explicit desktop launcher, `ModbusReplicator`, into
  `.osjs-desktop-iconview__wrapper` on `osjs/core:started` and re-ensures it via a `MutationObserver`
  on `document.body`. Launch uses `core.make('osjs/packages').launch(name)`. The stock Start/Application
  menu remains package-driven; this provider only adds desktop launchers.
- First independent Toolkit package `OSJS/src/packages/MCSModbusToolkit/` (icon.svg, index.js,
  index.scss, metadata.json, webpack.config.js). Its `metadata.json` declares type `application`,
  name `MCSModbusToolkit`, files `main.js`/`main.css`, and title "MCS Modbus Toolkit". `index.js`
  registers one `MCSModbusToolkitWindow` (960x640, centered) whose body is a placeholder that
  explicitly reads `NOT CONNECTED — PLACEHOLDER ONLY`; it imports no Electron source and calls no
  backend. Styling is confined to `.mcs-toolkit-placeholder`.
- `OSJS/src/packages/MCSModbusToolkit/webpack.config.js` follows the local-package convention
  (entry `index.js`, output `dist/main.js`, externals `{osjs: 'OSjs'}`, MiniCssExtract to
  `main.css`). The build/discovery commands the ACTIVE UMIG-002-T packet names are
  `npm run build:local-packages` and `npm run package:discover`, both present in
  `OSJS/package.json` (`build:manifest` and `package:discover` both map to
  `osjs-cli package:discover`).
- Each package carries its own `webpack.config.js`. `OSJS/webpack.config.js` was not modified by
  this delta; only `OSJS/src/client/index.js`, the new shortcut provider, and the new Toolkit
  package are new in the OSJS tree.
- The previously harvested `ModbusSimulator` and `ModbusReplicator` packages were added before this
  baseline (at `500376cf`..`ee19b8a`) and are not part of the `ee19b8a..HEAD` delta. Note for later
  Toolkit work: both still relay over a `$OSJS_DATA_DIR/run/*.sock` Unix socket, which no longer
  matches the committed named-pipe Go runtimes; see `simulator-memory-none` and `replicator`.
