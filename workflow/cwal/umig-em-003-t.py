#!/usr/bin/env python3
"""UMIG-EM-003-T only. Execute after human review/approval in detached JR worktree.

The script makes no product edits; its sole permitted repository modification is
replacing the final JR TEST REPORT section of handoff.md and committing/pushing
that single file. Refuses stale GitHub main, dirty checkout or unsafe preflight.
"""
import datetime
import os
from pathlib import Path
import re
import shutil
import subprocess
import sys

SOURCE = "949d7a8eb328dfa66f8359ffe81cfe359b8db2b7"
REPORT_HEADER = "## JR TEST REPORT — UMIG-EM-003-T"
ROOT = Path.home() / "apps" / "MCS.OSJS-jr"
EXPECTED_DIFF = {
    "handoff.md",
    "operation cwal.md",
    "workflow/active_work/umig-em-003-shared-lock.md",
    "workflow/active_work/umig-em-003-t-shared-lock.md",
    "workflow/archive/umig-em-003-shared-lock.md",
    "workflow/cwal/umig-em-003-t.py",
}
SUITES = (
    ("mma2composer", (
        "TestWriterLockTimeoutAndRelease", "TestWriterLockCrashRelease",
        "TestWriterLockRejectsAmbiguousPath", "TestWriterLockConcurrentForeignOwnershipConflict",
        "TestProducerIdentityAndForeignCollision",
    )),
    ("simulator", (
        "TestSimulatorComposeRejectsBusySharedWriterLock",
        "TestComposeRejectsForeignCollisionUnchanged",
    )),
    ("replicator", (
        "TestReplicatorComposeRejectsBusySharedWriterLock",
        "TestManagerCommittedUnacknowledgedRestartFailsClosed",
        "TestRuntimeManagerApplyLifecycleAndStatus",
        "TestRunOnceRejectsForeignOwnedDestination",
    )),
)


def run(command, cwd=ROOT, timeout=30):
    """Never use a shell: argv is always an explicitly enumerated command."""
    try:
        result = subprocess.run(command, cwd=cwd, text=True, stdout=subprocess.PIPE,
                                stderr=subprocess.STDOUT, timeout=timeout, check=False)
        return result.returncode, result.stdout
    except (OSError, subprocess.TimeoutExpired) as exc:
        return 125, f"COMMAND UNAVAILABLE OR PYTHON TIMEOUT: {exc}\n"


def record(log, label, command, cwd=ROOT, timeout=30):
    code, output = run(command, cwd, timeout)
    log.append((label, command, code, output))
    print(f"[{label}] exit={code}", flush=True)
    if code != 0:
        print(output[-2200:], flush=True)
    return code, output


def fail(reason):
    print(f"BLOCKED — {reason}. No tests, repository edit, commit or push beyond actions already shown.", flush=True)
    return 2


def main():
    if Path.cwd().resolve() != ROOT.resolve():
        return fail("must start in the exact separate JR worktree")
    if os.environ.get("MCS_RUN_E2E") == "1":
        return fail("MCS_RUN_E2E is enabled")
    log = []
    checks = (
        ("pwd", ["pwd", "-P"]),
        ("status", ["git", "status", "--porcelain", "--untracked-files=all"]),
        ("HEAD", ["git", "rev-parse", "HEAD"]),
        ("origin/main", ["git", "rev-parse", "origin/main"]),
        ("worktrees", ["git", "worktree", "list", "--porcelain"]),
        ("remote", ["git", "ls-remote", "--exit-code", "origin", "refs/heads/main"]),
        ("source ancestor", ["git", "merge-base", "--is-ancestor", SOURCE, "HEAD"]),
        ("source diff", ["git", "diff", "--name-only", SOURCE, "HEAD"]),
        ("system", ["uname", "-s"]),
        ("Go", ["go", "version"]),
        ("disk", ["df", "-Pk", ".", str(Path.home()), "/tmp"]),
    )
    values = {}
    for label, command in checks:
        code, output = record(log, label, command)
        if code != 0:
            return fail(f"preflight {label} exit {code}")
        values[label] = output.strip()
        if label == "status" and output.strip():
            return fail("initial worktree is dirty")
    head = values["HEAD"]
    remote_line = values["remote"].splitlines()
    if len(remote_line) != 1 or remote_line[0].split() != [head, "refs/heads/main"]:
        return fail("live GitHub main differs from local HEAD")
    if values["origin/main"] != head:
        return fail("origin/main differs from GitHub main or HEAD")
    if not re.search(rf"(?m)^worktree {re.escape(str(ROOT))}$\nHEAD {head}$\ndetached$", values["worktrees"]):
        return fail("JR worktree is not the expected detached HEAD")
    if set(values["source diff"].splitlines()) != EXPECTED_DIFF:
        return fail("source-to-activation changes differ from the six authorized workflow paths")
    if values["system"] != "Linux":
        return fail("non-Linux platform")
    version = re.search(r"\bgo(\d+)\.(\d+)(?:\.|\b)", values["Go"])
    if not version or (int(version.group(1)), int(version.group(2))) < (1, 25):
        return fail("installed Go is not >=1.25")
    for directory in (ROOT, Path.home(), Path("/tmp")):
        if shutil.disk_usage(directory).free < 3 * 1024**3:
            return fail(f"under 3 GiB available on {directory}")
    active = []
    for path in (ROOT / "workflow" / "active_work").glob("*.md"):
        if re.search(r"(?m)^Status:\s*ACTIVE\b", path.read_text()):
            active.append(path.name)
    if active != ["umig-em-003-t-shared-lock.md"]:
        return fail(f"sole ACTIVE gate failed: {active}")
    if not (ROOT / "workflow/archive/umig-em-003-shared-lock.md").is_file():
        return fail("CODE predecessor is not archived")
    if REPORT_HEADER not in (ROOT / "handoff.md").read_text():
        return fail("report anchor missing")
    print("PREFLIGHT PASS: latest live GitHub main, pinned source, one ACTIVE, clean and safe environment", flush=True)

    verdict = "PASS"
    reason = "All exact offline Go race suites, required named tests and post-check passed."
    ran = []
    for module, required in SUITES:
        command = ["timeout", "360s", "env", "GOTOOLCHAIN=local", "GOPROXY=off",
                   "GOSUMDB=off", "GOFLAGS=-mod=readonly", "go", "test", "-race",
                   "-count=1", "-timeout=300s", "-v", "./..."]
        code, output = record(log, "TEST " + module, command, ROOT / module, timeout=370)
        ran.append(module)
        if code:
            missing = any(item in output.lower() for item in (
                "module lookup disabled", "missing go.sum entry", "c compiler not found",
                "gcc: executable file not found", "toolchain not available"))
            verdict = "BLOCKED" if missing or code == 125 else "FAIL"
            reason = f"{module} exact suite returned {code}; stopped without retry."
            break
        if not re.search(r"(?m)^ok\s+", output) or re.search(r"(?m)^(?:FAIL|--- FAIL:|WARNING: DATA RACE|panic:)", output):
            verdict, reason = "FAIL", f"{module} lacks clean package OK or has failure/race evidence."
            break
        missing = [name for name in required if not re.search(r"(?m)^\s*--- PASS: " + re.escape(name) + r"(?:\s|$)", output)]
        if missing:
            verdict, reason = "BLOCKED", f"{module} required named-test evidence missing: {missing}."
            break
    for label, command in (("post status", ["git", "status", "--porcelain", "--untracked-files=all"]),
                           ("post HEAD", ["git", "rev-parse", "HEAD"])):
        code, output = record(log, label, command)
        if code or (label == "post status" and output.strip()) or (label == "post HEAD" and output.strip() != head):
            return fail(f"post-check {label} invalid; verdict so far {verdict}; no report push")
    code, current = record(log, "pre-report remote", ["git", "ls-remote", "--exit-code", "origin", "refs/heads/main"])
    if code or current.split() != [head, "refs/heads/main"]:
        return fail("remote main advanced during tests; no report push")

    report = [REPORT_HEADER, "", f"Verdict: **{verdict}**", f"Reason: {reason}",
              f"UTC: {datetime.datetime.now(datetime.timezone.utc).isoformat()}",
              f"Worktree: `{ROOT}`; source checkpoint: `{SOURCE}`; test HEAD / original remote: `{head}`.",
              "Runner: `python3 workflow/cwal/umig-em-003-t.py` (one approved invocation).",
              "Safety: offline Go modules, MCS_RUN_E2E != 1, no Docker, sudo or operator endpoints.",
              f"Test modules executed in order: {', '.join(ran)}; unrun: {', '.join(n for n, _ in SUITES if n not in ran) or 'none'}.",
              "Preflight and post-check stdout/stderr plus exit codes (exact commands are recorded below):", ""]
    for label, command, code, output in log:
        report += [f"### {label}: exit {code}", "", "Command: `" + " ".join(command) + "`", "", "```text", output.rstrip("\n") or "(no output)", "```", ""]
    report += ["Scope: Go TEST only; no live VERIFY, Electron or production claim.",
               "JR STOP: no source edits, task advancement or further tests authorized.", ""]
    handoff = ROOT / "handoff.md"
    original = handoff.read_text()
    marker = "\n" + REPORT_HEADER
    if original.count(marker) != 1:
        return fail("handoff report anchor ambiguous; no edit")
    prefix, _ = original.split(marker, 1)
    # Script-mediated handoff-only report write is expressly authorized by this
    # exact packet and the human approving this script; edit-tool access stays denied.
    handoff.write_text(prefix + "\n" + "\n".join(report))
    code, changed = record(log, "report changed paths", ["git", "diff", "--name-only"])
    if code or changed.splitlines() != ["handoff.md"]:
        return fail("report modified unexpected paths; do not commit")
    code, output = record(log, "report whitespace", ["git", "diff", "--check"])
    if code:
        return fail("report diff check failed; do not commit")
    code, output = record(log, "stage report", ["git", "add", "--", "handoff.md"])
    if code:
        return fail("cannot stage report")
    code, staged = record(log, "staged paths", ["git", "diff", "--cached", "--name-only"])
    if code or staged.splitlines() != ["handoff.md"]:
        return fail("unexpected staged path; no commit")
    code, output = record(log, "commit report", ["git", "commit", "-m", "JR: report UMIG-EM-003-T independent OpenCode test", "--only", "handoff.md"], timeout=60)
    if code:
        return fail("report commit failed; no push")
    code, commit = record(log, "report commit SHA", ["git", "rev-parse", "HEAD"])
    if code:
        return fail("cannot identify report commit; no push")
    code, paths = record(log, "report commit scope", ["git", "diff-tree", "--no-commit-id", "--name-only", "-r", "HEAD"])
    if code or paths.splitlines() != ["handoff.md"]:
        return fail("commit changed anything besides handoff; no push")
    code, current = record(log, "pre-push remote", ["git", "ls-remote", "--exit-code", "origin", "refs/heads/main"])
    if code or current.split() != [head, "refs/heads/main"]:
        return fail("remote moved; report committed locally but not pushed")
    code, output = record(log, "push report ONLY", ["git", "push", "origin", "HEAD:refs/heads/main"], timeout=120)
    if code:
        return fail("non-force report push failed; no retry")
    code, remote = record(log, "confirm pushed report", ["git", "ls-remote", "--exit-code", "origin", "refs/heads/main"])
    if code or remote.split() != [commit.strip(), "refs/heads/main"]:
        return fail("push verification failed; report status uncertain")
    code, dirty = record(log, "final status", ["git", "status", "--porcelain", "--untracked-files=all"])
    if code or dirty.strip():
        return fail("report pushed but worktree became dirty; report to coding agent")
    print(f"CWAL REPORT PUSHED: verdict={verdict}, commit={commit.strip()}; JR STOP", flush=True)
    return 0 if verdict == "PASS" else 1


if __name__ == "__main__":
    sys.exit(main())
