#!/usr/bin/env python3
"""Pinned UMIG-EM-003-R-T only: one approved offline Replicator test/report run.

Runs only in the clean detached Legion JR worktree. External access is limited
 to read-only Git main freshness and the final handoff-only non-force push.
"""
import datetime
import os
from pathlib import Path
import re
import shutil
import subprocess
import sys

ROOT = Path.home() / "apps" / "MCS.OSJS-jr"
SOURCE = "538324a472eb15ff8ef97ee66f826fa02ba9462f"
REPORT = "## JR TEST REPORT — UMIG-EM-003-R-T"
ALLOWED = {
    "handoff.md",
    "workflow/active_work/umig-em-003-r-recovery-regression.md",
    "workflow/archive/umig-em-003-r-recovery-regression.md",
    "workflow/active_work/umig-em-003-r-t-recovery-test.md",
    "workflow/cwal/umig-em-003-r-t.py",
}
REQUIRED = (
    "TestApplyPreCommitFailureRestoresPreviousPollers",
    "TestApplyPostCommitRestartFailureStopsPreviousPollers",
    "TestManagerCommittedUnacknowledgedRestartFailsClosed",
    "TestRuntimeManagerApplyLifecycleAndStatus",
)
TEST = ["timeout", "360s", "env", "GOTOOLCHAIN=local", "GOPROXY=off",
        "GOSUMDB=off", "GOFLAGS=-mod=readonly", "go", "test", "-race",
        "-count=1", "-timeout=300s", "-v", "./..."]


def command(argv, log, name, cwd=ROOT, limit=30):
    try:
        p = subprocess.run(argv, cwd=cwd, text=True, stdout=subprocess.PIPE,
                           stderr=subprocess.STDOUT, timeout=limit, check=False)
        code, output = p.returncode, p.stdout
    except (OSError, subprocess.TimeoutExpired) as exc:
        code, output = 125, f"Command unavailable or Python timeout: {exc}\n"
    log.append((name, argv, code, output))
    print(f"[{name}] exit={code}", flush=True)
    if code: print(output[-2500:], flush=True)
    return code, output.strip()


def blocked(reason):
    print(f"CWAL BLOCKED/INCOMPLETE: {reason}; no test retry or unauthorized cleanup. JR STOP", flush=True)
    return 2


def main():
    if Path.cwd().resolve() != ROOT.resolve():
        return blocked("not the exact detached JR worktree")
    if os.environ.get("MCS_RUN_E2E") == "1":
        return blocked("MCS_RUN_E2E enabled")
    log, values = [], {}
    preflight = (
        ("pwd", ["pwd", "-P"]),
        ("initial status", ["git", "status", "--porcelain", "--untracked-files=all"]),
        ("HEAD", ["git", "rev-parse", "HEAD"]),
        ("tracking", ["git", "rev-parse", "origin/main"]),
        ("worktrees", ["git", "worktree", "list", "--porcelain"]),
        ("live remote", ["git", "ls-remote", "--exit-code", "origin", "refs/heads/main"]),
        ("source ancestor", ["git", "merge-base", "--is-ancestor", SOURCE, "HEAD"]),
        ("source diff", ["git", "diff", "--name-only", SOURCE, "HEAD"]),
        ("platform", ["uname", "-s"]),
        ("Go version", ["go", "version"]),
        ("disk", ["df", "-Pk", ".", str(Path.home()), "/tmp"]),
    )
    for name, argv in preflight:
        code, output = command(argv, log, name)
        if code: return blocked(f"preflight {name} exit {code}")
        values[name] = output
        if name == "initial status" and output: return blocked("dirty initial repository")
    head = values["HEAD"]
    if not re.fullmatch(r"[a-f0-9]{40}", head): return blocked("invalid HEAD")
    if values["tracking"] != head or values["live remote"].split() != [head, "refs/heads/main"]:
        return blocked("HEAD, origin/main and current GitHub main do not match")
    if not re.search(rf"(?m)^worktree {re.escape(str(ROOT))}$\nHEAD {head}$\ndetached$", values["worktrees"]):
        return blocked("JR worktree is not detached at pinned HEAD")
    if set(values["source diff"].splitlines()) != ALLOWED:
        return blocked("source-to-activation files differ from authorized workflow paths")
    if values["platform"] != "Linux": return blocked("Linux required")
    v = re.search(r"\bgo(\d+)\.(\d+)(?:\.|\b)", values["Go version"])
    if not v or (int(v.group(1)), int(v.group(2))) < (1, 25):
        return blocked("Go >=1.25 required")
    for path in (ROOT, Path.home(), Path("/tmp")):
        if shutil.disk_usage(path).free < 3 * 1024**3:
            return blocked(f"less than 3 GiB free on {path}")
    active = [p.name for p in (ROOT / "workflow/active_work").glob("*.md")
              if re.search(r"(?m)^Status:\s*ACTIVE\b", p.read_text())]
    if active != ["umig-em-003-r-t-recovery-test.md"]:
        return blocked(f"active task mismatch: {active}")
    if not (ROOT / "workflow/archive/umig-em-003-r-recovery-regression.md").is_file():
        return blocked("CODE predecessor not archived")
    original = (ROOT / "handoff.md").read_text()
    marker = "\n" + REPORT
    if original.count(marker) != 1:
        return blocked("report anchor absent or ambiguous")
    print("PREFLIGHT PASS — live GitHub main checked, clean pinned source and sole ACTIVE verified", flush=True)

    code, test_output = command(TEST, log, "TEST replicator", ROOT / "replicator", limit=370)
    reason = "Replicator race suite and all mandatory named tests passed."
    verdict = "PASS"
    if code:
        unavailable = any(s in test_output.lower() for s in (
            "module lookup disabled", "missing go.sum entry", "c compiler not found",
            "gcc: executable file not found", "toolchain not available"))
        verdict = "BLOCKED" if code == 125 or unavailable else "FAIL"
        reason = f"Replicator suite exit {code}; original output preserved below."
    elif (not re.search(r"(?m)^ok\s+github\.com/tamzrod/MCS\.OSJS/replicator\s", test_output)
          or re.search(r"(?m)^(?:FAIL|--- FAIL:|WARNING: DATA RACE|panic:)", test_output)):
        verdict, reason = "FAIL", "Replicator suite lacks main package OK or contains failure evidence."
    else:
        missing = [name for name in REQUIRED if not re.search(
            r"(?m)^\s*--- PASS: " + re.escape(name) + r"(?:\s|$)", test_output)]
        if missing: verdict, reason = "BLOCKED", f"Missing actual named-test PASS evidence: {missing}."

    for name, argv in (("post status", ["git", "status", "--porcelain", "--untracked-files=all"]),
                       ("post HEAD", ["git", "rev-parse", "HEAD"]),
                       ("pre-report remote", ["git", "ls-remote", "--exit-code", "origin", "refs/heads/main"])):
        check, output = command(argv, log, name)
        if check or (name == "post status" and output) or (name == "post HEAD" and output != head) or (name == "pre-report remote" and output.split() != [head, "refs/heads/main"]):
            return blocked(f"{name} unsafe after product verdict={verdict} reason={reason}; report transport stopped")

    report = [REPORT, "", f"Verdict: **{verdict}**", f"Reason: {reason}",
              f"UTC: {datetime.datetime.now(datetime.timezone.utc).isoformat()}",
              f"Source checkpoint: `{SOURCE}`; tested activation HEAD/live main: `{head}`; worktree: `{ROOT}`.",
              "Runner: `python3 workflow/cwal/umig-em-003-r-t.py` (one human-approved invocation).",
              "Test scope: Replicator offline Go race suite ONLY; earlier Composer/Simulator results are historical, not rerun.",
              "E2E disabled; test-owned ephemeral loopback fixtures only; no Docker, service or operator device action.",
              "Complete command transcripts (merged original stdout/stderr) and exit codes:", ""]
    for name, argv, result, output in log:
        report.extend([f"### {name} — exit {result}", "", "Command: `" + " ".join(argv) + "`",
                       "", "```text", output.rstrip("\n") or "(no output)", "```", ""])
    report.extend(["Scope: Go TEST only; does not establish OS.js/Electron live VERIFY or production readiness.",
                   "JR STOP: no code or task changes and no additional tests authorized.", ""])
    handoff = ROOT / "handoff.md"
    if handoff.read_text() != original:
        return blocked("handoff changed since preflight; no report write")
    prefix, _ = original.split(marker, 1)
    handoff.write_text(prefix + "\n" + "\n".join(report))
    for name, argv, expect in (
        ("report changed files", ["git", "diff", "--name-only"], ["handoff.md"]),
        ("report whitespace", ["git", "diff", "--check"], []),
        ("stage report", ["git", "add", "--", "handoff.md"], []),
        ("staged paths", ["git", "diff", "--cached", "--name-only"], ["handoff.md"]),
    ):
        check, output = command(argv, log, name)
        if check or (name in ("report changed files", "staged paths") and output.splitlines() != expect):
            return blocked(f"report transport {name} failed; report file may be modified locally; do not clean or retry")
    check, _ = command(["git", "commit", "-m", "JR: report UMIG-EM-003-R-T OpenCode recovery tests", "--only", "handoff.md"], log, "commit", limit=60)
    if check: return blocked("report commit failed; do not push or retry")
    check, commit = command(["git", "rev-parse", "HEAD"], log, "report HEAD")
    if check: return blocked("report commit SHA unavailable")
    check, changed = command(["git", "diff-tree", "--no-commit-id", "--name-only", "-r", "HEAD"], log, "commit scope")
    if check or changed.splitlines() != ["handoff.md"]: return blocked("commit modified non-report file; no push")
    check, remote = command(["git", "ls-remote", "--exit-code", "origin", "refs/heads/main"], log, "pre-push remote")
    if check or remote.split() != [head, "refs/heads/main"]:
        return blocked(f"remote advanced; local report commit={commit}, not pushed")
    check, _ = command(["git", "push", "origin", "HEAD:refs/heads/main"], log, "report-only push", limit=120)
    if check: return blocked(f"non-force push failed; local report commit={commit}; no retry")
    check, remote = command(["git", "ls-remote", "--exit-code", "origin", "refs/heads/main"], log, "confirm remote")
    if check or remote.split() != [commit, "refs/heads/main"]:
        return blocked(f"remote confirmation failed; report commit={commit}; delivery uncertain")
    check, dirty = command(["git", "status", "--porcelain", "--untracked-files=all"], log, "final status")
    if check or dirty: return blocked(f"report pushed ({commit}) but worktree not clean; inspect without cleanup")
    print(f"CWAL REPORT PUSHED: verdict={verdict}, commit={commit}; JR STOP", flush=True)
    return 0 if verdict == "PASS" else 1


if __name__ == "__main__":
    sys.exit(main())
