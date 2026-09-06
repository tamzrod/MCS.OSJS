# OS.js Base Desktop Shell (Implemented

Baseline commit: 500376cfb5c222298aadfcf035aad0af0a635773
Working tree: clean
Source dependencies: OSJS/README.md, OSJS/package.json, OSJS/Dockerfile, OSJS/webpack.config.js, OSJS/scripts/build-local-packages.js, OSJS/src/server/config.js, OSJS/src/server/index.js, OSJS/src/server/providers/health.js, OSJS/src/server/providers/classic-icons.js, OSJS/src/client/config.js, OSJS/src/client/index.ejs, OSJS/src/packages/NamelessClassicIcons/metadata.json, OSJS/src/packages/NamelessWorkstationTheme/metadata.json
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
