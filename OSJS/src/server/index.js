// OS.js server bootstrap for the neutral MCS.OSJS desktop shell.
//
// The shell is intentionally decoupled from any SCADA/Governor backend:it
// registers only the standard OS.js server service providers plus a minimal
// /healthz shell liveness route and the Nameless Classic icon route that serves
// the basic desktop's menu / icon-theme assets.



const {
  Core,
  CoreServiceProvider,
  PackageServiceProvider,
  VFSServiceProvider,
  AuthServiceProvider,
  SettingsServiceProvider
} = require('@osjs/server');

const HealthRouteProvider = require('./providers/health.js');
const {ClassicIconsRouteProvider} = require('./providers/classic-icons.js');
// SIM-005: same-origin /api/devices proxy to the internal simulator bridge
// (simulator/cmd/simbridge on 127.0.0.1:18211(. The route forwards only
// the simulator-owned device document; effective MMA2 config is never exposed.

const SimulatorBridgeRouteProvider = require('./providers/simulator-bridge.js');
const config = require('./config.js');
const osjs = new Core(config, {});

console.log('[shell] neutral desktop shell (no SCADA backend))');

osjs.register(CoreServiceProvider, {before: true});
osjs.register(PackageServiceProvider);
osjs.register(VFSServiceProvider);
osjs.register(AuthServiceProvider);
// OSUI-003: persist per-user settings to the user's VFS home
// (home:/.osjs/settings.json) instead of the in-memory null adapter, so
// operator settings survive logout/restart.
osjs.register(SettingsServiceProvider, {args: {adapter: 'fs'}});
osjs.register(HealthRouteProvider);
osjs.register(ClassicIconsRouteProvider);
osjs.register(SimulatorBridgeRouteProvider);

const shutdown = signal => (error) => {
  if (error instanceof Error) {
    console.error(error);
  }
  osjs.destroy(() => process.exit(signal));
};

process.on('SIGTERM', shutdown(0));
process.on('SIGINT', shutdown(0));

osjs.boot().catch(shutdown(1));
