# Experimental Editor

A small deterministic text editor for diagnosing MCS.OSJS authoring corruption.

It is intentionally not an AI repair tool. It performs bounded line edits, writes to a temporary file, verifies exact bytes, then atomically replaces the target. It reports SHA-256 values for the supplied input, original file, and final file so we can distinguish upstream malformed authoring from write-path mutation.

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

## Diagnostic interpretation

A successful modification prints:

```text
SAFEEDIT_OK path=...
input_sha256=... input_bytes=...
before_sha256=... before_bytes=...
after_sha256=... after_bytes=...
```

The editor independently verifies the temporary file and final file byte-for-byte against its received input/result before reporting success.

If suspicious text such as doubled commas, joined words, or wrong delimiters is already present in the input file supplied to the editor, the problem occurred upstream of this editor. If the editor ever reports a byte-verification failure, that is evidence of mutation inside or below the write path.

## Scope

Experimental only. Do not make this the default repository editor yet, and do not use it to auto-normalize or broadly repair malformed source. The purpose is to create a controlled alternate authoring path and collect evidence.
