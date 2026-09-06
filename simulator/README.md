# Simulator device configuration (SIM-001)

Simulator-owned device definitions live here as a component boundary separate
from MMA2 effective runtime configuration and from the OS.js packaged source
tree.

Persistent files are written under the verified host-mounted data root:

```text
$OSJS_DATA_DIR/config/simulator/devices.yaml
```

`OSJS_DATA_DIR` is the established appliance mount (`deploy/docker-compose.yml`
volume `osjs-data` → `/data`, `OSJS/src/server/config.js`). This component
refuses to invent a host path when the variable is unset.
