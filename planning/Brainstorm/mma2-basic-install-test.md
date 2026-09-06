# Brainstorm: MMA2 Basic Install and Test

Status: brainstorm material only. This does not authorize implementation.

## Intent

Bring MMA2 into MCS.OSJS as the next standalone component and prove that the basic MMA2 runtime works before any Replicator, Orchestrator, or OS.js integration is attempted.

## Candidate Outcome

A local MCS.OSJS checkout contains a runnable MMA2 component sourced from `tamzrod/mma2`, and MMA2 passes one simple end-to-end Modbus memory test.

## Basic Scope

- add MMA2 under a clear repository component directory;
- preserve MMA2 as its own component boundary;
- build MMA2 successfully;
- start MMA2 with one minimal configuration;
- expose one Modbus TCP listener for the test;
- configure one Unit ID with a small memory range;
- perform one write and one read against the same address;
- confirm the read returns the value that was written;
- stop and restart MMA2, then confirm the configured listener starts cleanly again.

## Basic Verification Shape

The first useful proof should be deliberately small:

```text
BUILD MMA2
→ START MMA2
→ VERIFY LISTENER
→ WRITE ONE REGISTER
→ READ SAME REGISTER
→ VALUE MATCHES
→ RESTART MMA2
→ LISTENER RETURNS
```

This is enough to establish that the imported MMA2 component can build, start, expose Modbus TCP, and perform a basic deterministic memory operation.

## Non-Scope

- no Replicator integration;
- no Orchestrator integration;
- no OS.js configuration UI;
- no multi-port or multi-Unit-ID stress testing;
- no persistence design changes;
- no container consolidation work;
- no protocol expansion beyond the basic Modbus TCP test.

## Donor Facts Relevant to This Brainstorm

The MMA2 donor is a Go project with `go.mod`, `cmd/mma2/main.go`, a Dockerfile, tests, and an existing MMA2 binary in the repository. Its architecture defines configuration as startup-loaded and immutable, with configured ports, Unit IDs, and memory sizes feeding the core memory and transport/runtime layers.

## Open Questions for Later Planning

- Which exact MCS.OSJS directory should own the imported MMA2 component?
- Should the first execution use `go run`, a built binary, or the donor Dockerfile?
- Which minimal configuration file should be used for the smoke test?
- Which test client/tool should perform the single Modbus write/read?

These questions should be resolved only when converting this brainstorm into bounded microtask(s).
