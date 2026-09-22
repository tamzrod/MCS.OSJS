# Operation CWAL

## Mandatory entry and authority
Read AGENTS.md and root handoff.md. The handoff names the only executable task and its queue. If Current task is NONE, report NO ACTIVE TASK and STOP. A retired queue, historical packet, archive, or evidence report never grants execution authority. Missing packet, inactive status, contradictory routing or stale checkout means BLOCKED / STOP.

Exactly ONE explicitly ACTIVE microtask per invocation. Execute only its exact authorized actions, capture genuine evidence, record COMPLETE, FAIL or BLOCKED as permitted, perform packet-required delivery checks, and STOP. Do not activate or execute the successor, search for independent work, loop until exhaustion, or reinterpret a packet STOP. Fresh task planning and activation require separate human authorization.

## Roles and boundaries
OpenCode CODE/DISCOVERY/DESIGN: follow workflow/adapters/opencode.md and the one current packet. Independent JR TEST/VERIFY: follow the JR lifecycle below and its exact test packet. Calling Operation CWAL does not assign JR identity. No role may edit ICC except through BLACK SHEEP WALL authority. Never infer missing endpoints, tests, permissions or PASS. No destructive cleanup, production access, unrelated edits, merge, force-push, or main push without explicit authority.

## JR lifecycle
JR requires a complete current test packet naming goal, target, safe environment, pinned source and activation revision, exact commands and expected results, raw evidence, post-check, verdict rules, report path/schema, report-only commit, non-force push/ref and remote delivery verification. Thin, chat-only, stale, mismatched or incomplete packets are BLOCKED / STOP.

Before testing verify exact worktree, branch, clean state, packet ACTIVE status, source ancestry and current remote SHA using git ls-remote. Never repair stale state by autonomously fetching, switching, resetting, cleaning, restoring, rebasing or merging. Run authorized product tests once; stop on the first failed gate. Capture exact commands, exit codes and full stdout/stderr or packet-approved durable transcript. Perform safe required post-check even after FAIL. Never rerun tests to force PASS.

JR cannot edit product source, configuration, manifests, tasks, handoff, ICC or directives. Its only write exception is the exact packet-named report. Verify changed and staged paths equal that report alone. Commit the report, verify remote has not advanced, push non-force to the specified branch and independently verify remote SHA. If delivery fails, report INCOMPLETE with the underlying test result preserved. An OpenCode self-test is not independent JR PASS.

PASS requires all packet checks directly observed. Contradicted expectation means FAIL; unavailable authority, unsafe environment, stale source or missing mandatory evidence means BLOCKED. Missing report, post-check or required delivery means INCOMPLETE. Report evidence and STOP after any outcome. No automatic successor promotion.
