# MMA2 Basic Install and Test — Microtask

Status: planning material only. Human promotion is required before execution.

## MMA2-001 — Install MMA2 and prove one basic Modbus memory path

### Primary outcome

A repository-root `MMA2/` component sourced from `tamzrod/mma2` builds and runs successfully, and one minimal Modbus TCP write/read smoke test proves the imported MMA2 memory path works end to end.

### Scope

- Create repository-root `MMA2/` as the MMA2 component boundary.
- Import only the donor material required to build and run MMA2 from `tamzrod/mma2`.
- Preserve MMA2 as an independently bounded component; do not merge its implementation into OS.js or another component.
- Add one minimal MMA2 configuration for the smoke test.
- Configure one Modbus TCP ingress listener, one Unit ID, and a small memory range sufficient for the test.
- Build MMA2 from the imported source.
- Start MMA2 with the minimal configuration.
- Verify the configured TCP listener is actually bound.
- Write one known value to one configured holding-register address through Modbus TCP.
- Read the same address back and verify the returned value exactly matches the written value.
- Stop MMA2, start it again with the same configuration, and verify the listener returns cleanly.
- If Docker is used for any runtime verification, use `network_mode: "host"`; do not introduce Docker `ports:` mappings for MMA2 Modbus TCP.

### Non-scope

- No Replicator integration.
- No Orchestrator integration.
- No OS.js MMA2 configuration UI.
- No multi-port testing.
- No multi-Unit-ID stress testing.
- No persistence redesign.
- No container consolidation work beyond respecting the existing host-network requirement if Docker is used for verification.
- No protocol expansion beyond the basic Modbus TCP smoke test.
- No unrelated refactoring of donor MMA2 behavior.

### Acceptance criteria

1. `MMA2/` contains the donor-derived source and build/runtime material required for MMA2 to compile and start as its own component.
2. A clean MMA2 build succeeds from the imported source without relying on the donor repository checkout at runtime.
3. MMA2 starts from a minimal repository-owned configuration and the configured Modbus TCP listener is confirmed bound.
4. A real Modbus TCP client writes a known register value and reads the same address back with an exact value match.
5. After MMA2 is stopped and restarted with the same configuration, the listener binds again and a read of the configured memory path succeeds without startup, configuration, or listener errors.
6. No Docker bridge/NAT port publishing is introduced for MMA2; if Docker participates in the verification path, host networking is used.

### Verification method

Perform one end-to-end verification workflow and record evidence for each stage:

```text
CLEAN BUILD
→ START MMA2 WITH TEST CONFIG
→ VERIFY LISTENER
→ MODBUS WRITE ONE REGISTER
→ MODBUS READ SAME REGISTER
→ VERIFY EXACT VALUE MATCH
→ STOP MMA2
→ RESTART MMA2
→ VERIFY LISTENER RETURNS
→ VERIFY MODBUS READ PATH STILL RESPONDS
```

Required evidence:

- exact build command and successful result;
- exact MMA2 start command and configuration path;
- listener inspection showing the configured address/port is bound by MMA2;
- exact Modbus client command/tool used for the write and its target Unit ID/address/value;
- exact Modbus client command/tool used for the read and the returned value;
- explicit comparison showing written value equals read value;
- stop/restart evidence;
- post-restart listener and Modbus response evidence;
- if Docker is used, runtime/config evidence showing host networking and absence of MMA2 `ports:` publishing.

File-copy evidence alone is not completion evidence.

### Dependencies

- Existing MCS.OSJS network directive requiring Docker `network_mode: "host"` for the appliance when containerized.
- MMA2 donor repository `tamzrod/mma2` as the implementation source.
- Human promotion to Active Work.

### Sizing assessment

Size: **4 / 10**.

This is acceptable as one JR task because import, build, launch, and the single Modbus write/read/restart check all serve one primary outcome and one continuous verification workflow. Split only if execution exposes an independent blocker that can be verified separately.
