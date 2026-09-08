# MMA2 Basic Install and Test — Archived Work

Status: COMPLETED — archived 2026-09-08 after completion evidence was reconfirmed.

## Execution Order

1. `MMA2-001` — Import MMA2 and prove a clean build. **DONE** — completed+verified per evidence below.
2. `MMA2-002` — Run MMA2 and prove one basic Modbus memory path. **DONE** — completed+verified per evidence below.

JR must complete and verify `MMA2-001` before starting `MMA2-002`.

---

## MMA2-001 — Import MMA2 and prove a clean build

**Completion evidence (recorded by CWAL 2026-09-06(:**
- Imported `tamzrod/mma2` @ `12311c1d06510840b42723a83438574e6b3ed06f` into repository-root `MMA2/`; delivered as an independent component boundary (`cmd/`+`internal/`, `go.mod/go.sum/Dockerfile/.gitignore`, `README.md`, `docs/`, `LICENSE`; prebuilt donor binary + donor `test/` e2e scaffolding excluded).
- Clean build from `MMA2/`:  `go build -o /tmp/mma2-build-check ./cmd/mma2` — exit 0, binary produced, Go 1.25.0 toolchain.
- No-donor-checkout dependency proven:  build succeeds with `/tmp/mma2-probe` donor checkout renamed away (exit 0).
- Provenance/license recorded in `MMA2/README.md` and `THIRD_PARTY_NOTICES.md` per harvest gate; Apache-2.0 text retained in `MMA2/LICENSE`.


### Primary outcome

A repository-root `MMA2/` component sourced from `tamzrod/mma2` exists as an independent component and builds cleanly from repository-owned source without requiring the donor checkout at build or runtime.

### Scope

- Create repository-root `MMA2/` as the MMA2 component boundary.
- Import only donor material required to build and run MMA2 from `tamzrod/mma2`.
- Preserve MMA2 as an independently bounded component.
- Make only the minimum repository-local adjustments required for the imported component to build.
- Perform a clean build from `MMA2/`.

### Non-scope

- No MMA2 runtime smoke test.
- No Modbus client write/read test.
- No listener verification.
- No restart verification.
- No Replicator, Orchestrator, or OS.js integration.
- No unrelated donor refactoring.

### Acceptance criteria

1. `MMA2/` contains the donor-derived source and build/runtime material required for MMA2 as its own component.
2. A clean build succeeds from the imported repository source.
3. The build does not depend on the separate donor repository checkout.

### Verification method

```text
VERIFY MMA2/ COMPONENT BOUNDARY
→ CLEAN BUILD FROM MMA2/
→ VERIFY BUILD SUCCESS
→ VERIFY NO DONOR CHECKOUT DEPENDENCY
```

Required evidence:

- imported source location;
- exact clean-build command;
- successful build result;
- evidence that the build uses repository-root `MMA2/`, not the donor checkout.

### Dependencies

- MMA2 donor repository `tamzrod/mma2` as implementation source.

### Sizing assessment

Size: **3 / 10**.

One bounded implementation surface and one build-verification workflow.

---

## MMA2-002 — Run MMA2 and prove one basic Modbus memory path

**Completion evidence (recorded by CWAL 2026-09-06(:**
- Minimal repo-owned config: `MMA2/testdata/smoke-test.yaml` (listener `127.0.0.1:5020`, unit 1, holding registers 0..10; explicit allow-all policy — MMA2 authority defaults-to-deny with no rules per `authority.go`, excluding the "No Policy = Allow" doc claim(.)
- Start command: `/tmp/mma2 /workspace/project/MCS.OSJS/MMA2/testdata/smoke-test.yaml` (binary built via `go build -o /tmp/mma2 ./cmd/mma2`(. Verify: log `ingress smoke-test listening on 127.0.0.1:5020` + socket connect BOUND-OK.
- Real Modbus TCP client (raw standard Modbus TCP PDU — FC6 write single register(( wrote value 1234 to holding addr 0, unit 1; response `0001 0000 0006 01 06 0000 04d2` (write-echo, no exception(.
- Read back via FC3: response `0002 0000 0005 01 03 02 04d2` — value 0x04D2 = 1234 — exact match.
- Second independent real client, pymodbus 3.15.0 (ModbusTcpClient(: wrote 1234, read back 1234, explicit equality `1234 == 1234 True`.
- Stop/restart with same config:  both log + socket confirm listener rebinds `127.0.0.1:5020`; post-restart FC3 read path responds (addr 0 =value 0 — fresh in-memory zeroed store(; no startup/config/listener errors in restart log.
- Docker: n/a — direct binary run, no Docker used;trivially no `ports:` mapping introduced.


### Primary outcome

The repository-owned MMA2 build starts from one minimal configuration and passes one end-to-end Modbus TCP write/read/restart smoke test.

### Scope

- Add one minimal repository-owned MMA2 configuration.
- Configure one Modbus TCP ingress listener, one Unit ID, and a small memory range sufficient for the test.
- Start the repository-owned MMA2 build with that configuration.
- Verify the configured TCP listener is bound.
- Write one known value to one configured holding-register address using a real Modbus TCP client.
- Read the same address back and verify an exact value match.
- Stop and restart MMA2 with the same configuration.
- Verify the listener returns and the Modbus read path responds after restart.
- If Docker participates in runtime verification, use `network_mode: "host"` and do not introduce Docker `ports:` mappings for MMA2 Modbus TCP.

### Non-scope

- No donor import work except a defect discovered in the already-completed `MMA2-001` result.
- No Replicator integration.
- No Orchestrator integration.
- No OS.js MMA2 configuration UI.
- No multi-port or multi-Unit-ID stress testing.
- No persistence redesign.
- No protocol expansion beyond the basic Modbus TCP smoke test.

### Acceptance criteria

1. MMA2 starts from a minimal repository-owned configuration and the configured Modbus TCP listener is confirmed bound.
2. A real Modbus TCP client writes a known holding-register value and reads the same address back with an exact value match.
3. After stop/restart with the same configuration, the listener binds again and the Modbus read path responds without startup, configuration, or listener errors.
4. If Docker is used, MMA2 uses host networking and no Docker bridge/NAT port publishing is introduced.

### Verification method

```text
START MMA2 WITH TEST CONFIG
→ VERIFY LISTENER
→ MODBUS WRITE ONE REGISTER
→ MODBUS READ SAME REGISTER
→ VERIFY EXACT VALUE MATCH
→ STOP MMA2
→ RESTART MMA2
→ VERIFY LISTENER RETURNS
→ VERIFY MODBUS READ PATH RESPONDS
```

Required evidence:

- exact start command and configuration path;
- listener inspection showing the configured address/port is bound by MMA2;
- exact Modbus client/tool and target Unit ID/address/value;
- returned read value and explicit equality check;
- stop/restart evidence;
- post-restart listener and Modbus response evidence;
- if Docker is used, host-network evidence and absence of MMA2 `ports:` publishing.

File-copy or build evidence alone is not completion evidence for this task.

### Dependencies

- `MMA2-001` completed and verified.
- Existing MCS.OSJS network directive requiring Docker `network_mode: "host"` when containerized.

### Sizing assessment

Size: **3 / 10**.

One runtime behavior surface and one continuous end-to-end verification workflow. Toolchain/import uncertainty belongs to `MMA2-001`, not this task.
