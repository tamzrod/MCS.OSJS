// SIM-005 — same-origin /api/devices proxy to the internal simulator bridge.
//
// The OS.js server (public management/UI port) exposes /api/devices so the
// browser client can GET/PUT the simulator-owned device document without directly
// reaching the loopback-only Go bridge (simulator/cmd/simbridge, bound to
// 127.0.0.1:18211). This route never touches effective MMA2 runtime
// configuration: it forwards only the simulator-owned document. If the bridge
// is not running, the route answers 503 so the UI can surface a bridge-unavailable
// state rather than silently failing.

'use strict';

const http = require('http');

const DEFAULT_BRIDGE = '127.0.0.1:18211';

class SimulatorBridgeRouteProvider {
  constructor(core, options = {}) {
    this.core = core;
    this.options = options;
  }

  provides() {
    return ['simulator/bridge'];
  }

  init() {
    const express = this.core.make('osjs/express');
    if (!express || typeof express.route !== 'function') {
      console.warn('SimulatorBridgeRouteProvider: osjs/express not available; route skipped');
      return Promise.resolve(true);
    }
    const target = this.options.bridgeAddress || process.env.SIMBRIDGE_ADDR || DEFAULT_BRIDGE;
    express.route('get', '/api/devices', (req, res) => this.forward(req, res, target, 'GET'));
    express.route('put', '/api/devices', (req, res) => this.forward(req, res, target, 'PUT'));
    return Promise.resolve(true);
  }

  forward(req, res, target, method) {
    const outbound = http.request({hostname: '127.0.0.1', port: Number(target.split(':').pop()) || 18211, path:'/api/devices', method, headers: {'content-type': 'application/json'}}, (upstream) => {
      const chunks = [];
      upstream.on('data', (c) => chunks.push(c));
      upstream.on('end', () => {
        const body = Buffer.concat(chunks);
        res.status(upstream.statusCode || 502).set('content-type', upstream.headers['content-type'] || 'application/json').send(body);
      });
      upstream.on('error', (err) => {
        console.error('SimulatorBridgeRouteProvider: upstream error:', err.message);
        res.status(502).json({error: 'simulator bridge unavailable'});
      });
    });
    outbound.on('error', (err) => {
      console.error('SimulatorBridgeRouteProvider: bridge unreachable:', err.message);
      res.status(503).json({error: 'simulator bridge not running'});
    });
    if (method === 'PUT' && req.body && typeof req.body === 'object') {
      outbound.write(JSON.stringify(req.body);
    }
    outbound.end();
  }

  start() { return Promise.resolve(true); }
  destroy() {}
}

module.exports = SimulatorBridgeRouteProvider;