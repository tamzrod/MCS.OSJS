# REP-003 — Replicator Device Configuration

Status: COMPLETED
Previous: REP-002
Next: REP-004

## Primary Outcome

Add a minimal persisted Replicator configuration model describing one external Modbus source and one MMA2 destination reservation, using the same storage conventions and validation style already proven by the Simulator where applicable.

## Scope

- Add Replicator-owned persisted configuration under `$OSJS_DATA_DIR/config/replicator/`.
- Define source host, port, unit ID, function/area, start address, count, and polling interval.
- Define destination MMA2 listener port, unit ID, area/start/count.
- Validate obvious invalid/zero/out-of-range values before persistence.
- Keep the first configuration shape limited to one simple 1:1 register range.

## Completion Evidence

JR verification PASS on commit `c11ae364652e1f2b2e65b9d74397545050b25edb`:

- clean working tree before test;
- `gofmt -l .` produced no output;
- `go test -count=1 ./...` passed for `github.com/tamzrod/MCS.OSJS/replicator`;
- `go vet ./...` exited 0 with no diagnostics;
- Go toolchain was sandbox/user-local only and did not modify project code.

REP-003 is complete and REP-004 is the authorized next task.
