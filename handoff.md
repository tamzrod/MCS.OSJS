# Handoff

## Current

ACTIVE: REP-002 — Share MMA2 Raw-Ingest Client

## Authorized Sequence

REP-001 → REP-002 → REP-003 → REP-004 → REP-005 → REP-006 → REP-007

## Continuation

REP-001 completed and was archived. REP-002 is `ACTIVE`. REP-003 through REP-007 are `QUEUED` and already human-authorized; they require no additional approval when advanced through their explicit `Previous` / `Next` links.

## REP-001 Completion (2026-09-10)

The producer-neutral `mma2composer` module now owns effective-config and ownership composition. Simulator delegates to it with producer identity `simulator`; focused shared tests, the full Simulator suite, and both modules' `go vet` gates passed. REP-002 was advanced through REP-001's explicit `Next` link.

## Continuation Rule

`workflow/active_work/` is execution authority. This file is only the continuation summary and must agree with Active Work.

Operation CWAL executes only the one task marked `ACTIVE`. On verified completion it archives that task and advances only its explicit `Next` task from `QUEUED` to `ACTIVE`. If Active Work and this handoff disagree, JR stops rather than guessing.

## RESCUE CHECKPOINT — THERE IS NO COW LEVEL (2026-09-09)

Previous sandbox was rescued mid-REP-001 due to persistent authoring/tool-layer corruption. The rescue commit contains the checkpoint exactly as described below; this sandbox is not authoritative beyond that commit.



### Implementation point reached

- `mma2composer/` module scaffold exists and is the intended shared package: `go.mod` (module `github.com/tamzrod/MCS.OSJS/mma2composer`, requires `gopkg.in/yaml.v3 v3.0.1`, local `replace mma2 => ../MMA2`) and `go.sum` for yaml/check deps.
- The only committed draft of `mma2composer/composer.go` is MALFORMED (entire body corrupted with comma/paren/period mutations, e.g. `} error {`, `yaml.Marshal(cfg}`, broken `var ErrReservationOwnedByOther`, unexpected `  // c` tokens, duplicated loop positions). It does NOT compile and must NOT be built upon恢复.
- A mechanical recovery plan exists but every attempt to author/replace `composer.go` (file_editor create/str_replace and python heredoc/regex generation) suffered the SAME class of punctuation corruption in the NEW text, including in ascript file (`gen_composer.py`) written via file_editor — i.e. reproducible evidence of tool/transport corruption independent of the affected Go file。
- Per AGENTS.md, after the circuit breaker tripped(one corrective retry + one fresh recovery attempt failed again with the same corruption pattern, production STOPPED. No repair-script chains, no broad regex cleanup, no further retry attempts ware made by the rescued session; rescue was invoked by the user。



### Known-unfinished / malformed / failing state

- `mma2composer/composer.go`: DELETED from working tree during recovery attempts(committed malformed draft still in HEAD`. The INTENDED final content(producer-neutral package API): package `mma2composer`, types `OwnershipEntry`/`OwnershipDoc`/`Area`/`PolicyRule`/`Policy`/`Memory`/`Listener`/`EffectiveConfig`, `Composer{Root, Producer}`, `New`, `ConfigDir`/`EffectiveConfigPath`/`OwnershipPath`, `LoadEffective`/`LoadOwners`, `SaveEffective`/`SaveOwners`, `WriteFileAtomic`, `ErrReservationOwnedByOther`, `Collision`, `Commit` (yaml-marshal → `configvalidate.YAML` → atomic replace config → atomic replace owners → byte-for-byte restore on second-write failure), `DropProducerReservations`, `DropOneReservation`, `AddMemory`, `ListenPort`) must be authored FRESH in the new sandbox, NOT reconstructed from the malformed draft。
- The known-good behavioral source to generalize is `simulator/mma2_config.go` (see below》. It is byte-verified intact。
- Simulator still uses its private duplicate logic (`mma2_config.go`); REP-001 switch-to-shared-composer is NOT done。
。
- No `mma2composer` tests exist yet。
- `gen_composer.py` (repo-root scratch generator) contains the same corruption class; it is excluded from the rescue commit and should be deleted bythe new sandbox。

。





### Verification already performed + result

- Baseline Simulator package suite passed BEFORE all rescue/refactor attempts: `ok github.com/tamzrod/MCS.OSJS/simulator 1.413s`(recorded from earlier session state》。
- `mma2composer/go.mod` + `go.sum` present; go.sum lists yaml.v3 v3.0.1 and check.v1 indirect deps。
- `mma2composer/composer.go` committed draft: `gofmt -l` listed it; `go vet ./...` in module failed with `missing ',' in argument list` + 10 more errors — malformed as expected。
- All post-corruption refactor/authoring attempts FAILED verify(same class of syntax errors;; none progressed beyond authoring failure as recorded above。



### Next sandbox — exact next action (REP-001 continuation》

1. Start from the known-good byte-true source `simulator/mma2_config.go`(verified intact; do NOT carry over any malformed text from`mma2composer/composer.go` or `gen_composer.py`)。The commit`HEAD` malformed `composer.go` must be replaced, NEVER repaired or pattern-copied。
2. First run a minimal authoring sanity probe(e.g. write a tiny pristine `.go` file with parens/commas/braces via the intended authoring tool, compile).If the probe itself mutates punctuation, stop authoring immediately and re-invoke rescue protocol; do not chain workarounds。

3. Author `mma2composer/composer.go` fresh, as: package `mma2composer` + the producer-neutral API listed above, bodies copied byte-for-byte from the generalized simulator logic (rename `MMA2Area`→`Area`, `MMA2PolicyRule`→`PolicyRule`, `MMA2Policy`→`Policy`, `MMA2Memory`→`Memory`, `MMA2Listener`→`Listener`, `EffectiveMMA2Config`→`EffectiveConfig`, `listenPort`→exported `ListenPort`, `ProducerSimulator`→parameterized `Composer.Producer`); keep config artifact paths exactly `config/mma2/config.yaml` / `config/mma2/owners.yaml`。
4. Run `gofmt -l mma2composer`(must be empty)and `go vet ./...` in `mma2composer/`;then add focused package tests for producer identity、foreign-owner rejection、first-come-first-save persistence、`Commit` restore-on-failure。
5. Refactor Simulator to use the shared composer: add `mma2composer` to `simulator/go.mod` requires(+`replace mma2composer => ../mma2composer`), add a small Store adapter(private `simComposer = mma2composer.New(s.Root`, `simulator`)`or equivalent) inside `simulator/mma2_config.go` that calls `LoadEffective`/`LoadOwners`/`DropProducerReservations`/`DropOneReservation`/`AddMemory`/`Collision`/`Commit`, delete the private duplicates(`dropSimulatorReservations`, `dropOneSimulatorReservation`, `mergeReservation`, `listenPort`, duplicate types), keep `Store.MMA2ConfigDir`/`EffectiveConfigPath`/`OwnershipPath` shims for `restart.go`/`apply.go`/`runtime_server.go`, keep `ProducerSimulator = "simulator"` const for ownership comparisons, remap tests to the shared behavior(existing `compose_test.go` suite must stay green semantically)。
6. Run full Simulator suite(`go test ./...` in `simulator/`);on green, commit REP-001, then proceed to REP-002 per Active Work。

达



### Excluded from rescue commit

- `gen_composer.py` — corrupted authoring scratch; delete it in the new sandbox。
- Deleted working-tree `mma2composer/composer.go` — intended fresh authoring starts from known-good sources listed above, not from this deletion。.
## CHECKPOINT (2026-09-09 second sandbox) — REP-001 still ACTIVE

Committed + pushed this sandbox: `mma2composer/composer.go` verified-good fresh authoring (0 non-ASCII bytes; `gofmt` empty; `go vet` clean; compiled as committed). It replaces the malformed draft described below (which does not exist on disk) and is NOT to be rebuilt.

`mma2composer/composer_test.go` does NOT exist: two fresh authoring attempts hit the same punctuation-corruption class (U+FEFF/U+FFFD/U+3001, 63 then 81 non-ASCII bytes). Both deleted uncontaminated. Per AGENTS.md circuit breaker tripped after one corrective retry + one fresh recovery attempt; authoring stopped. Simulator refactor NOT started. Simulator suite was green this session (`go test -count=1 ./...` -> ok 1.408s).

Next recovery action: author ONLY `mma2composer/composer_test.go` fresh (producer identity, Collision foreign rejection, first-come-first-save persistence, Commit restore-on-failure();then refactor Simulator onto the shared composer per REP-001;then run full simulator suite;then complete REP-001 per CWAL (archive, advance REP-002, update handoff, commit, push, verify origin/main(.
