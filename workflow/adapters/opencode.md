# OpenCode adapter — Ubuntu / Qwen

This adapter changes execution mechanics, NOT task authority. Read `workflow/IDENTITY_MAP.md`, `AGENTS.md`, `operation cwal.md`, `handoff.md` and `workflow/active_work/OTR_PROMOTION_QUEUE.md` as needed. OPERATION CWAL is an instruction to route and execute, not a request to ask the user what to do. Human promotion of the OTR queue has already occurred. `ACTIVE`/`QUEUED` are scheduling states, not additional approval gates. If handoff identifies an assigned valid task, execute it immediately without asking for confirmation. If the handoff is stale or its task is genuinely complete, on a NEW invocation use the queue's deterministic first-eligible successor rule only when actual predecessor evidence and a complete packet are available. Never infer completion from numbering or a chat claim alone. Exactly one task per invocation; STOP after its report. A later invocation may route the next task without another human selection. No self-triggering or background work is implied.

## Autonomous CODE / DISCOVERY lifecycle

1. Resolve ONE eligible execution packet with explicit ID, owner, objective, scope, acceptance, evidence, permissions and STOP. DISCOVERY is read-only unless explicitly stated. Do not execute copied parent templates OTR-004..014 or treat an obsolete planning header as a second approval requirement; missing material execution details remain a real blocker. If no valid packet exists, report its exact missing fields once and STOP, not 'what would you like me to do?'.
2. Read only named files and relevant predecessor evidence. Do not repeat completed investigations or explore unrelated source.
3. Check `git status --short` and HEAD as packet requires. Preserve pre-existing changes; overlapping edits stop work. Do not reset, restore, clean, stash, switch branches or reconcile unrelated Git state without authorization.
4. For CODE only, edit the exact task allowlist, inspect diff and perform bounded in-scope fixes. No installs, services, sudo, production data, network operations or unrelated refactors without explicit packet authority.
5. Run only packet-authorized targeted tests; capture commands, exit codes and actual output. CODE self-tests are never independent JR PASS. DISCOVERY does not run builds/tests unless its packet says so.
6. Report task ID, observed HEAD, scope, paths, evidence, actual outcomes, blockers and successor eligibility. Write task state/handoff or commit/push ONLY if packet explicitly authorizes those exact paths and transport; otherwise report in chat and STOP. Do not claim persistence that did not occur.
7. STOP after exactly one task. A subsequent autonomous invocation can select the next eligible packet from verified evidence; this STOP does not require the human to reauthorize the approved queue. FAIL/BLOCKED never auto-advances. Do not edit ICC; BLACK SHEEP WALL alone owns ICC.

## Independent JR

CODE and independent TEST/VERIFY remain separate. A JR packet requires exact commands/actions, source checkpoint, safety and report permissions. OpenCode must not independently certify its own work. No fabricated test evidence or automatic live-device action. The Ubuntu operator workstation is not a disposable sandbox.
