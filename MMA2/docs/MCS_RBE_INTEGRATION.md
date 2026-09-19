# MCS RBE integration

Pinned source: tamzrod/mma2 at 84934d540b7201f96f1ab411177bc62ec6d2165c.
License: Apache-2.0, unchanged from the original harvest; MMA2/LICENSE is retained.

## Imported scope

The changed runtime slice comprises cmd/mma2/main.go; internal/authority
(authority.go, decision.go, removal of sealing.go); internal/config
(config.go, build_rbe_rules.go, reject_influx.go and RBE/legacy-notify tests);
internal/ingress/listener.go and tests; internal/memorycore/observed_write.go
and tests; internal/notify/engine_test.go and removal of influx_adapter.go;
internal/rbe/*; internal/transport/modbus dispatch/connection changes and
RBE/notify tests; internal/transport/rawingest connection changes and RBE tests.
Supporting material: test/test_notify.yaml, examples/rbe-tcp.yaml,
docs/RBE_V1_DRAFT.md. Other upstream tools, deployment and large test fixtures
were not imported. This is not a whole-folder replacement.

## Local compatibility adaptations

- Preserve cmd/mma2-supervisor, internal/restartwatch, Dockerfile.supervised
  and pkg/configvalidate. Preserve module/dependency files and existing data paths.
- pkg/configvalidate.YAML performs structural validation followed by BuildRBERules,
  rejecting invalid RBE settings before composer writes/restart requests.
- Load and Validate explicitly reject removed notify.influx and rbe.influx settings.
  They are not silently discarded or migrated; external subscribers own Influx delivery.
- Retain fatal ingress failure behavior so a bind failure cannot leave a seemingly
  running appliance with a missing Modbus listener.
- The serving goroutine and shutdown closure share the same Listener pointer,
  avoiding a copy of its mutex and ensuring Close targets the serving listener.

## Configuration contract

RBE rules belong to listeners[].memory[].rbe.<area>. Memory identity is port + Unit ID.
Each rule has a globally unique ID in 1..255, name, start and count within its allocation.
The shared output is rbe.tcp.listen, required when RBE is configured; there is no implicit
default port. It must not reuse a Modbus/Raw Ingest listener port. Legacy notify and RBE
cannot coexist in one configuration. TCP events contain one rule-ID byte, not values.

State sealing remains per memory. Raw Ingest writes beginning sealed do not emit RBE,
including the unsealing write. Modbus access is rejected while sealed. Upstream's draft
documents subscriber reconciliation and lack of authentication/replay. Restrict exposure
to trusted peers; this integration does not enable an RBE listener in installed data.

## Boundaries

This step integrates the backend only. Electron controls, producer persistence, packaging,
installed-service verification, human UI acceptance and production latency qualification
are separate tasks. Passing unit/race tests does not claim those later gates have passed.
