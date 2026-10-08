# NPE-06H — Restore MMA2 Validation CLI Compatibility

Status: CODE COMPLETE
Stage: HOTFIX / CODE
Owner: ChatGPT operator-assisted hotfix
Previous: NPE-06W Windows acceptance failure
Next: NPE-06W rerun

## Defect

Electron Save & Apply invokes the bundled MMA2 binary with:

```
mma2 --validate-stdin
```

The NPE-05A donor refresh replaced the embedded `MMA2/cmd/mma2/main.go` with the upstream donor version, which no longer contained the MCS-local generic validation-only CLI mode.

As a result, MMA2 interpreted `--validate-stdin` as a config filename and returned:

```
config path must end in .yaml or .yml
```

## Fix

Restore the generic `--validate-stdin` mode in embedded MMA2:
- read at most 4 MiB from stdin;
- validate through `pkg/configvalidate.YAML`;
- exit without starting listeners or mutating runtime state;
- preserve normal `mma2 <config.yaml>` behavior.

This restores the pre-refresh Electron/MMA2 contract without adding persistence semantics to Electron.

## Delivered

Product commit:
`8c1a805a868dfaec41c5d2d11ab0bc7ec4a53ba0`

Changed:
- `MMA2/cmd/mma2/main.go`

## Next

Rebuild the Windows Electron package so the bundled `mma2.exe` includes this commit, then rerun NPE-06W from Save & Apply.
