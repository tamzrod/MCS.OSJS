// Minimal /healthz liveness route for the neutral desktop shell.
//
// Replaces the former SCADA bridge's /healthz so deployment healthchecks keep
// working without coupling the shell to any platform backend. Never echoes
// credentials.

class HealthRouteProvider {
  constructor(core, options = {}) {
    this.core = core;
    this.options = options;
  }

  provides() {
    return ['osjs/health'];
  }

  init() {
    const express = this.core.make('osjs/express');
    if (!express || typeof express.route !== 'function') {
      console.warn('HealthRouteProvider: osjs/express not available; route skipped');
      return Promise.resolve(true);
    }
    express.route('get', '/healthz', (req, res) => {
      const body = {status: 'ok', shell: 'neutral'};
      if (this.options.realtimeHealth && this.options.realtimeHealth.realtime) {
        const relay = this.options.realtimeHealth.realtime;
        body.realtime = typeof relay === 'function' ? relay() : relay;
      }
      res.status(200).json(body);
    });
    return Promise.resolve(true);
  }

  start() {
    return Promise.resolve(true);
  }

  destroy() {}
}

module.exports = HealthRouteProvider;
