# Brainstorm: Replicator UI Layout

Status: brainstorm material only. Non-authoritative. No implementation authorized until human promotion.

## Current Layout Direction

```text
+---------------------------------------------------------------+
| Modbus Replicator                                             |
+----------------------+----------------------------------------+
| Devices              | Device Definition                      |
|                      |                                        |
| [ Search devices ]   | Name          [ Pump-PLC-1        ]    |
| [Add] [Duplicate]    | Enabled       [x]                      |
| [Delete]             |                                        |
|                      | Source                                 |
| Pump-PLC-1           | Endpoint      [192.168.1.20:502   ]    |
| Meter-01             | Unit ID       [1]                      |
| Boiler-PLC           | FC            [3]                      |
|                      | Start         [0]                      |
|                      | Count         [16]                     |
|                      | Scan Rate     [1000 ms]                |
|                      |                                        |
|                      | Destination                            |
|                      | Port          [Auto: 5021]             |
|                      | Unit ID       [Auto: 2]                |
|                      | Owner         Replicator               |
|                      | Status        AVAILABLE                |
|                      |                                        |
|                      | [ Save & Apply ]   [ Discard ]         |
+----------------------+----------------------------------------+
| Replicator: RUNNING   Source: OK   Last Poll: 0.4s ago        |
+---------------------------------------------------------------+
```

## Device Model

Left side remains the device list. Selecting a device opens its definition on the right.

First UI version keeps the source definition aligned with the current simple Replicator backend:

- Name
- Enabled
- Endpoint
- Source Unit ID
- FC
- Start
- Count
- Scan Rate

Scan Rate is device-wide for the first version. Do not introduce per-block scan rates until multiple read ranges are intentionally designed.

## Destination Behavior

Destination allocation should be automatic by default.

- Port: auto-fill the next available destination port/reservation unless the user explicitly specifies one.
- Unit ID: auto-fill the next available Unit ID for the selected port unless the user explicitly specifies one.
- Manual override remains available.
- User should not normally need to manage destination allocation.

Suggested simple presentation:

```text
Destination
Port       [ Auto: 5021 ]
Unit ID    [ Auto: 2    ]
Owner      Replicator
Status     AVAILABLE
```

If the manually selected `(port, unit_id)` is already owned by another producer:

```text
Port       [5020]
Unit ID    [1]
Owner      Simulator
Status     IN USE
```

Save & Apply must reject the collision rather than overwrite the existing reservation.

## Ownership Rule

Shared MMA2 destination ownership follows first-come, first-served reservation semantics.

- Ownership key is `(port, unit_id)`.
- A free reservation may be claimed by Replicator.
- A reservation already owned by Replicator may be updated by Replicator.
- A reservation owned by Simulator, Memory Appliance, or another producer must not be overwritten by Replicator.
- Replicator may delete/rebuild only Replicator-owned memory/reservations.
- Replicator must never delete foreign-owned memory/reservations.
- The inverse applies to other producers: Simulator must not delete Replicator-owned reservations.
- Collision must be reported truthfully to the UI.

This keeps shared MMA2 configuration cooperative: first owner keeps the reservation until that owner releases it.

## Status Area

Keep operational status lightweight at the bottom of the window. Initial useful fields:

```text
Replicator: RUNNING   Source: OK   Last Poll: 0.4s ago
```

Possible error state:

```text
Replicator: RUNNING   Source: ERROR   Last Error: connection timeout
```

No history database, advanced metrics, transforms, or unrelated runtime features are implied by this brainstorm.
