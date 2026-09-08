# Simulator Runtime Integration

Status: accepted architecture for SIM-019 through SIM-022.

## Decision

The Modbus Simulator uses three local layers:

1. The OS.js application runs in the operator's browser.
2. An authenticated OS.js server provider relays Simulator messages over the existing OS.js session WebSocket. It does not register a Simulator HTTP route.
3. A separately supervised Go `modbus-simulator-runtime` process owns `ApplyRouter`, schedulers, and Raw Ingest clients. The provider reaches it only through a Unix-domain socket at `$OSJS_DATA_DIR/run/modbus-simulator.sock`.

The runtime and OS.js shell share `$OSJS_DATA_DIR`. The runtime's `Store` at `$OSJS_DATA_DIR/config/simulator/devices.yaml` is the sole canonical Simulator document. Per-user OS.js settings are not a second configuration store.

MMA2 remains independently supervised and auto-started. The Simulator runtime never starts, stops, spawns, kills, or replaces MMA2. After a committed structural change it may write the existing restart request and wait for MMA2 readiness. Scheduled values enter MMA2 only through Raw Ingest.

## Why this boundary

- A browser cannot directly retain Go schedulers, access the appliance filesystem, or inspect MMA2 readiness.
- The long-lived runtime process preserves schedules when the application window closes.
- A Unix socket provides a host-local process boundary without a TCP listener.
- Reusing the authenticated OS.js session WebSocket avoids restoring `/api/devices`, `/api/devices/status`, a standalone `simbridge`, or any Simulator configuration/control HTTP API.
- One canonical Go store keeps UI reload, boot restore, apply rollback, and scheduler state aligned.

Rejected alternatives:

- Browser-only OS.js settings: can persist form data but cannot execute `ApplyRouter`, retain schedules, or prove Raw Ingest success.
- A standalone Simulator HTTP service: recreates the removed bridge/config API.
- Starting a Go command per click: loses scheduler ownership when the command exits.
- Polling shared JSON/settings files as a command mailbox: creates ambiguous ownership, weak request ordering, and poor failure semantics.
- Running MMA2 as a child of Simulator: violates independent MMA2 lifecycle ownership.

## Process and startup model

`modbus-simulator-runtime` is supervised independently with restart-on-failure behavior and shares the appliance data volume. On startup it:

1. creates the socket directory with appliance-local permissions;
2. removes only its own stale socket after proving no live listener owns it;
3. constructs `NewRuntimeApplyRouter(Store)`;
4. restores the canonical document, waits for independently started MMA2, and arms schedules only when ready;
5. accepts local RPC messages.

The OS.js provider connects lazily and reconnects after runtime restart. A disconnected runtime yields an explicit `RUNTIME_UNAVAILABLE` response; the UI never converts that state to STOPPED or RUNNING. Closing the browser window does not stop the runtime or its schedules.

## Message contract

Every message is a JSON object with a protocol version and request identifier. The OS.js provider copies the authenticated session identity into server-side audit context; it does not trust identity fields supplied by the browser.

Request envelope:

```json
{"version":1,"request_id":"uuid","operation":"load|apply|status","payload":{}}
```

Response envelope:

```json
{"version":1,"request_id":"uuid","ok":true,"result":{},"error":null}
```

Errors use stable codes plus human-readable detail: `INVALID_REQUEST`, `VALIDATION_FAILED`, `RESERVATION_CONFLICT`, `MMA2_RESTART_FAILED`, `MMA2_NOT_READY`, `RAW_INGEST_FAILED`, `RUNTIME_UNAVAILABLE`, and `INTERNAL`.

Operations:

- `load`: returns the canonical `Document` and the last apply result. It does not mutate runtime state.
- `apply`: accepts one complete `Document`, calls `ApplyRouter.Apply`, and returns its classification (`mma2-structural`, `random-runtime`, or `no-change`), completion time, and resulting canonical document. Only a successful response advances the UI's Discard snapshot.
- `status`: accepts a device name and returns `DeviceRuntimeStatus` plus the last apply result. It is read-only.

Requests and responses are length-prefixed on the Unix socket, have bounded payload size and deadlines, and are processed with mutation serialization. Status reads may run concurrently against snapshot-safe runtime methods. Duplicate `request_id` values return the cached completed response for a short bounded interval so a reconnect cannot apply the same mutation twice.

## Browser behavior

The package uses OS.js's authenticated socket connection, not `fetch()` and not a new HTTP endpoint. The provider permits only the three versioned operations and validates message size/schema before forwarding them to the Unix socket.

On window open, the client loads the canonical document. Save & Apply remains pending until an `apply` response arrives. Status refreshes once per second while the window exists, stops when destroyed, and marks stale/unavailable data visibly after a missed deadline.

The UI may say RUNNING or Raw Ingest OK only from runtime evidence. An enabled checkbox, locally edited form, saved settings value, open Unix socket, or reachable MMA2 port alone is insufficient.

## Security and exposure

- No Simulator TCP port and no Simulator HTTP route exist.
- The Unix socket is inside the shared appliance data volume and is inaccessible to the browser/network.
- The OS.js provider accepts messages only from authenticated OS.js sessions and enforces operation allowlists, payload bounds, request deadlines, and error redaction.
- The provider does not expose arbitrary filesystem, process, MMA2 lifecycle, or socket operations.

## Implementation boundaries

- SIM-019 implements the supervised Go runtime, Unix-socket protocol, OS.js relay provider, and runtime-level verification.
- SIM-020 moves document load/apply from OS.js settings to the approved message contract.
- SIM-021 renders and refreshes truthful runtime/apply status.
- SIM-022 verifies the whole visible workflow with real MMA2, Raw Ingest, and Modbus reads.
