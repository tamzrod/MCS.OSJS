# Experimental Editor

A small deterministic text editor for MCS.OSJS authoring safety and diagnosis.

It is not an AI repair tool. It performs bounded line edits, writes to a temporary file, verifies exact bytes, then atomically replaces the target. It reports SHA-256 values for the supplied input, original file, and final file so the agent can distinguish malformed authored input from write-path mutation.

## Operational status

Use this as the preferred write path for substantive source, configuration, and document edits, especially multiline replacements or generated blocks.

Do not route ordinary reads, git operations, directory creation, deletion, or trivial exact edits through it unnecessarily.

A `SAFEEDIT_OK` result proves that the editor wrote exactly the bytes it received. It does **not** prove that the input text was authored correctly.

## Commands

```bash
python3 tools/experimental-editor/editor.py view FILE --start 10 --end 20
python3 tools/experimental-editor/editor.py hash FILE
python3 tools/experimental-editor/editor.py replace-lines FILE 10 15 --input /tmp/new.txt
python3 tools/experimental-editor/editor.py insert-lines FILE 15 --input /tmp/new.txt
python3 tools/experimental-editor/editor.py delete-lines FILE 10 15
```

When `--input` is omitted, `replace-lines` and `insert-lines` read replacement text from standard input.

Line numbers are 1-based. `insert-lines FILE 0` inserts before the first line.

## Required usage pattern

For a substantive edit:

```text
READ authoritative target
→ prepare smallest coherent replacement input
→ inspect the actual prepared input
→ apply with experimental editor
→ require SAFEEDIT_OK
→ inspect resulting target
→ run repository-native formatter/compiler/test gate
```

If the prepared input is already malformed, do not apply it and do not normalize it broadly. Follow the `AGENTS.md` authoring-failure rules.

If patterned corruption appears after one exact correction, do not switch to heredocs, `str_replace`, generator scripts, ad-hoc Python rewrites, or another editor as a workaround chain. Enter Fresh Authoring Recovery instead.

## Diagnostic interpretation

A successful modification prints:

```text
SAFEEDIT_OK path=...
input_sha256=... input_bytes=...
before_sha256=... before_bytes=...
after_sha256=... after_bytes=...
```

The editor independently verifies the temporary file and final file byte-for-byte against its received input/result before reporting success.

Interpret results narrowly:

- malformed text already present in `--input` means the problem is upstream of this editor for that attempt;
- `SAFEEDIT_OK` with correct input means this editor preserved those bytes for that attempt;
- an editor byte-verification failure is evidence of mutation inside or below this write path;
- a clean transport/editor probe does not prove the authoring layer is healthy.

## Control fixture

The immutable control fixture is:

```text
tools/experimental-editor/fixtures/canonical.txt
```

Its expected SHA-256 is stored in:

```text
tools/experimental-editor/fixtures/canonical.sha256
```

The fixture is for diagnostics only. Normal implementation work must use the actual authoritative repository source and task requirements, not this fixture as content.

## Safety boundaries

- Make bounded edits only.
- Prefer the smallest coherent transformation.
- Never auto-normalize malformed source.
- Never use global punctuation cleanup or Unicode stripping as corruption recovery.
- Do not claim a lower-level cause unless a reproducible probe isolates that layer.
- Stop after repeated patterned malformed authoring according to `AGENTS.md`.
