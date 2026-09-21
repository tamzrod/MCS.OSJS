# OS.js Test Target Metadata

## Repository Baseline

- **Branch:** `temp-main` (aligned to origin/main)
- **Commit SHA:** `b546a8f` — "populate micro task"
- **Working Tree:** clean
- **Audit Source:** Read-only inspection

---

## Approved Disposable Target

### Identity: OS.js Base Desktop Shell (Test Only)

**Purpose:** Standalone disposable Compose environment for safe Node/Go package testing. Never reuse without verified clean-vol check; do not deploy with operator Compose.

### Toolchain & Source Manifests

#### Node Environment (OSJS)

- **Package Manager:** `npm` (via `@osjs/cli`)
- **Node Version Boundary:** `>=10.0.0 <17` (per package.json engines)
- **OS.js CLI:** `^3.1.3`
- **Core Modules Sampled:**
  - `@osjs/client`: `^3.7.1`
  - `@osjs/server`: `^3.2.3`
  - `@osjs/panels`: `^3.0.31`
  - `@osjs/gui`: `^4.0.36`
- **Webpack Build:** v4 (`webpack-cli: ^3.3.12`)

#### Go Environment (Simulator & Replicator)

- **Build Mode:** Docker build from context in parent directory
- **Runtime Binding:** Socket under `$OSJS_DATA_DIR/run/*.sock` (Docker-managed, not host)
- **Network Isolation:** Internal-only networks (`verify-runtime`, `verify-ui`)
- **Data Volume:** Dedicated `verify-data` for scoped volumes only

---

## Permission & Network Boundaries

- **Port Exposure:** `127.0.0.1:${MCS_VERIFY_OSJS_PORT:-18219}:18209` (UI-only)
- **Host Access:** None (no host network mode, no published Modbus ports)
- **Data Boundary:** Write confined to `verify-data` volume; test refuse reuse of any populated volume

---

## Exact Commands for Later Stages

### Initialize Target Environment

```bash
# Clean start: compose up with empty verify-data only
docker-compose -f deploy/verify/compose.yaml up --build --remove-orphans --force-recreate
# First run seeds /data config; subsequent runs fail if verify-data not empty (by design)
```

### Read-Only Inspect (Current Stage Only)

```bash
# Verify volume state
docker inspect compose_mcs-osjs-jr_verify-data 2>&1 | grep -i '"MountSource"'

# Confirm service health readiness (no live verification required for PREP gate pass)
docker exec compose_modbus-simulator-runtime cat /data/config/simulator/devices.yaml 2>/dev/null || echo "No devices configured"
```

---

## Target Metadata Status: **APPROVED**

- **Recording Authority:** Human decision to approve new disposable target creation per OSJT-001 block relief.
- **Reuse Policy:** Allowed only after confirming clean `verify-data` volume state and matching baseline commit evidence; never share with operator Compose stacks.
- **Lifecycle Gate:** PREP stage completed under read-only inspection; TEST/VERIFY activation requires separate authorization.

