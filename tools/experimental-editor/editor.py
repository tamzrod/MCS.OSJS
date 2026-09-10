#!/usr/bin/env python3
"""Experimental deterministic text editor for MCS.OSJS.

This tool is intentionally small. It does not repair text, run regex cleanup,
or infer intent. It applies bounded line edits to a temporary file, verifies the
exact bytes written, then atomically replaces the target.
"""

from __future__ import annotations

import argparse
import hashlib
import os
from pathlib import Path
import sys
import tempfile
from typing import Iterable


def sha256(data: bytes) -> str:
    return hashlib.sha256(data).hexdigest()


def read_bytes(path: Path) -> bytes:
    return path.read_bytes()


def read_input(path: str | None) -> bytes:
    if path:
        return Path(path).read_bytes()
    return sys.stdin.buffer.read()


def split_lines(data: bytes) -> list[bytes]:
    return data.splitlines(keepends=True)


def normalize_insert(data: bytes) -> list[bytes]:
    if not data:
        return []
    lines = split_lines(data)
    if lines and not lines[-1].endswith((b"\n", b"\r")):
        lines[-1] += b"\n"
    return lines


def validate_range(start: int, end: int, line_count: int, allow_empty: bool = False) -> None:
    if start < 1:
        raise ValueError("start line must be >= 1")
    if end < start and not allow_empty:
        raise ValueError("end line must be >= start line")
    if end > line_count:
        raise ValueError(f"end line {end} exceeds file line count {line_count}")


def atomic_write_verified(path: Path, output: bytes) -> None:
    parent = path.parent
    parent.mkdir(parents=True, exist_ok=True)

    original_mode = None
    if path.exists():
        original_mode = path.stat().st_mode & 0o7777

    fd, tmp_name = tempfile.mkstemp(prefix=f".{path.name}.safeedit.", dir=parent)
    tmp = Path(tmp_name)
    try:
        with os.fdopen(fd, "wb") as handle:
            handle.write(output)
            handle.flush()
            os.fsync(handle.fileno())

        actual = tmp.read_bytes()
        if actual != output:
            raise RuntimeError(
                "SAFEEDIT_REJECTED: temporary-file byte verification failed "
                f"expected={sha256(output)} actual={sha256(actual)}"
            )

        if original_mode is not None:
            os.chmod(tmp, original_mode)

        os.replace(tmp, path)

        final = path.read_bytes()
        if final != output:
            raise RuntimeError(
                "SAFEEDIT_REJECTED: final-file byte verification failed "
                f"expected={sha256(output)} actual={sha256(final)}"
            )
    finally:
        if tmp.exists():
            tmp.unlink()


def print_write_report(path: Path, before: bytes, after: bytes, input_bytes: bytes) -> None:
    print(f"SAFEEDIT_OK path={path}")
    print(f"input_sha256={sha256(input_bytes)} input_bytes={len(input_bytes)}")
    print(f"before_sha256={sha256(before)} before_bytes={len(before)}")
    print(f"after_sha256={sha256(after)} after_bytes={len(after)}")


def cmd_view(args: argparse.Namespace) -> int:
    path = Path(args.file)
    lines = split_lines(read_bytes(path))
    start = args.start or 1
    end = args.end or len(lines)
    validate_range(start, end, len(lines))
    for number in range(start, end + 1):
        raw = lines[number - 1]
        text = raw.decode("utf-8", errors="backslashreplace").rstrip("\r\n")
        print(f"{number:6d}  {text}")
    return 0


def cmd_hash(args: argparse.Namespace) -> int:
    path = Path(args.file)
    data = read_bytes(path)
    print(f"sha256={sha256(data)} bytes={len(data)} path={path}")
    return 0


def cmd_replace_lines(args: argparse.Namespace) -> int:
    path = Path(args.file)
    before = read_bytes(path)
    lines = split_lines(before)
    validate_range(args.start, args.end, len(lines))
    incoming = read_input(args.input)
    replacement = normalize_insert(incoming)
    output = b"".join(lines[: args.start - 1] + replacement + lines[args.end :])
    atomic_write_verified(path, output)
    print_write_report(path, before, output, incoming)
    return 0


def cmd_insert_lines(args: argparse.Namespace) -> int:
    path = Path(args.file)
    before = read_bytes(path)
    lines = split_lines(before)
    if args.after < 0 or args.after > len(lines):
        raise ValueError(f"after line must be between 0 and {len(lines)}")
    incoming = read_input(args.input)
    insertion = normalize_insert(incoming)
    output = b"".join(lines[: args.after] + insertion + lines[args.after :])
    atomic_write_verified(path, output)
    print_write_report(path, before, output, incoming)
    return 0


def cmd_delete_lines(args: argparse.Namespace) -> int:
    path = Path(args.file)
    before = read_bytes(path)
    lines = split_lines(before)
    validate_range(args.start, args.end, len(lines))
    output = b"".join(lines[: args.start - 1] + lines[args.end :])
    atomic_write_verified(path, output)
    print_write_report(path, before, output, b"")
    return 0


def parser() -> argparse.ArgumentParser:
    p = argparse.ArgumentParser(prog="experimental-editor")
    sub = p.add_subparsers(dest="command", required=True)

    view = sub.add_parser("view", help="view a numbered line range")
    view.add_argument("file")
    view.add_argument("--start", type=int)
    view.add_argument("--end", type=int)
    view.set_defaults(func=cmd_view)

    hash_cmd = sub.add_parser("hash", help="print exact file SHA-256 and size")
    hash_cmd.add_argument("file")
    hash_cmd.set_defaults(func=cmd_hash)

    replace = sub.add_parser("replace-lines", help="replace an inclusive line range")
    replace.add_argument("file")
    replace.add_argument("start", type=int)
    replace.add_argument("end", type=int)
    replace.add_argument("--input", help="replacement text file; stdin when omitted")
    replace.set_defaults(func=cmd_replace_lines)

    insert = sub.add_parser("insert-lines", help="insert text after a line; 0 means before line 1")
    insert.add_argument("file")
    insert.add_argument("after", type=int)
    insert.add_argument("--input", help="text file to insert; stdin when omitted")
    insert.set_defaults(func=cmd_insert_lines)

    delete = sub.add_parser("delete-lines", help="delete an inclusive line range")
    delete.add_argument("file")
    delete.add_argument("start", type=int)
    delete.add_argument("end", type=int)
    delete.set_defaults(func=cmd_delete_lines)

    return p


def main() -> int:
    try:
        args = parser().parse_args()
        return args.func(args)
    except (OSError, ValueError, RuntimeError) as exc:
        print(f"SAFEEDIT_ERROR: {exc}", file=sys.stderr)
        return 2


if __name__ == "__main__":
    raise SystemExit(main())
