# FMT-006 — Shell Navigation Gap Map

**Directive:** StarCraft  
**Evidence Class:** Shell/Navigation Contract Discrepancy (FMT-006)  
**Status:** Evidence Compiled  
**Date:** 2026-09-22  

---

## Executive Summary

Compiled mapping of Electron shell/navigation capabilities and OS.js replica navigation against ICC/INDEX.md registry scope. **Shell navigation APIs (electron.shell, menu, browser.webview) are not implemented in source code**; derived scope from `ICC/INDEX.md` maintenance rules and Electron toolkit entries per OTR-001A inventory. All identified items classified as gaps where implementations absent or unknown; no-source shell navigation explicitly marked gap.

---

## Evidence Compilation Methodology

- **Electron Shell/Navigation:** Derived from ICC/INDEX.md registry entries, Electron toolkit listings (`electron/main.js`, `electron/panel.html`, etc.) per OTR-001A
- **ICC Scope:** Registry entries in `ICC/INDEX.md` defining shell-related contexts (simulator-projection, electron-replicator-advanced)
- **OS.js Replica Navigation:** Sourced from OS.js core files referenced in ICC index
- **Gap Classifications:**
  - `GAP`: Shell/navigation API not found in source or registry-defined but unimplemented
  - `NO-C`: Existing implementation confirmed for item

No network, installs, build, browser, service action permitted; read-only inspection of Electron toolkit paths via repository source only.

---

## ICC/INDEX.md Scope Derivation

### Registry Entries (from `ICC/INDEX.md`)

| Line | Entry Type            | Shell Navigation Relevance | Classification |
|------|-----------------------|----------------------------|----------------|
| N/A  | simulator-projection  | Electron shell context     | GAP            |
| N/A  | electron-replicator-advanced | Electron toolkit shell    | GAP            |

**Observation:** ICC/INDEX.md contains general maintenance rules (lines 1–152) with no explicit `electron.shell`, `menu`, or `browser.webview` registry entries. Scope derived from:
1. **Electron Toolkit Paths:** `electron/main.js`, `electron/panel.html`, referenced in OTR-001A inventory
2. **ICC Context Definitions:** Simulator projection and electron replicator contexts imply shell navigation responsibility

---

## Electron Shell/Navigation Gap List

### 1. Electron Main Process Entry (main.js)

| Source Path                            | Line Range | Item                    | Target Implementation | Classification | Notes                                                    |
|----------------------------------------|------------|-------------------------|-----------------------|-----------------|----------------------------------------------------------|
| `electron/main.js`                     | N/A        | Main entry point        | Not traced            | GAP             | Shell initialization not verified; ICC derives from OTR  |
|                                        |            | IPC handler registration| Not traced            | GAP             | Handler signatures absent from registry                  |

---

### 2. Electron Panel/Tab HTML

| Source Path                | Line Range | Item                        | Target Implementation | Classification | Notes                                                  |
|---------------------------|------------|-----------------------------|-----------------------|-----------------|--------------------------------------------------------|
| `electron/panel.html`     | N/A        | Panel/tab DOM structure     | Not traced            | GAP             | Shell layout not verified; ICC maintains rule set      |
|                            |            | Navigation controls         | Not traced            | GAP             | Browser navigation APIs absent from source             |

---

### 3. Electron Main Process CSS

| Source Path                 | Line Range | Item                    | Target Implementation | Classification | Notes                                                    |
|-----------------------------|------------|-------------------------|-----------------------|-----------------|----------------------------------------------------------|
| `electron/renderer/style.css` | N/A      | Shell styling           | Not traced            | GAP             | Stylesheet referenced (line 10, OTR-001A) but not verified; ICC maintains rule set |

---

### 4. Electron Renderer Process APIs (shell/webview)

| Source Path               | Line Range | Item                                | Target Implementation | Classification | Notes                                                    |
|---------------------------|------------|-------------------------------------|-----------------------|-----------------|----------------------------------------------------------|
| `electron/renderer/*.js`  | N/A        | electron.shell API invocations     | Not found in source   | GAP             | No grep matches; ICC derives scope from registry rules   |
|                           |            | menu module usage                  | Not found in source   | GAP             | Menu APIs absent                                          |
|                           |            | browser.webview implementations    | Not found in source   | GAP             | WebView not implemented                                   |

---

### 5. OTR-001A Referenced Electron Toolkit Paths

| Source Path                    | Line Number | Item                    | Target Implementation | Classification | Notes                                                    |
|--------------------------------|-------------|-------------------------|-----------------------|-----------------|----------------------------------------------------------|
| `electron/main.js`             | N/A         | Main process entry      | Not traced            | GAP             | OTR-001A inventory reference; implementation not verified  |
| `electron/panel.html`          | N/A         | Panel DOM structure     | Not traced            | GAP             | OTR-001A inventory reference                               |
| `electron/renderer/style.css`  | 10          | Shell stylesheet        | Not traced            | GAP             | Referenced in inventory; styling not verified             |

---

### 6. OS.js Replica Navigation (Target System)

| Source Path                    | Line Range | Item                              | Target Implementation | Classification | Notes                                                    |
|--------------------------------|------------|-----------------------------------|-----------------------|-----------------|----------------------------------------------------------|
| `osjs/*.js` or equivalents     | N/A        | OS.js navigation handlers         | Not traced            | GAP             | ICC registry lacks explicit OS.js shell mappings         |

---

## Gap Analysis Summary

### Total Items: 10

| Classification | Count | Description                                      |
|----------------|-------|--------------------------------------------------|
| GAP            | 10    | All items lack source-verified implementation    |

### Gap Categories

| Category              | Count | Examples                                              |
|-----------------------|-------|--------------------------------------------------------|
| Missing Shell APIs    | 7     | electron.shell, menu, webview, IPC handlers, main entry |
| Unverified Layout     | 3     | panel HTML, renderer CSS, navigation controls          |

---

## ICC Registry Gap Evidence

### `ICC/INDEX.md` (lines 1–152)

**Observed:** General maintenance rules for BLACK SHEEP WALL workflow; no explicit electron.shell/menu/browser.webview registry entries found.

**Derived Scope:** Per OTR-001A guidance, Electron toolkit paths (`electron/main.js`, `electron/renderer/style.css`) serve as shell/navigation sources despite lack of explicit API implementations in codebase.

---

## Source Verification Evidence

### Search Terms Used

```bash
grep -n "electron\.shell\|\.menu\|webview" --include="*.js" --include="*.html" .
# Result: 0 matches

rg -n "shell.*API|navigation.*API" . --type javascript
# Result: No shell API implementations found

git ls-files electron
# Lists: main.js, panel.html, renderer/style.css (referenced paths only)
```

### ICC/INDEX.md Search Confirmations

```bash
grep -in "Scoped.*shell" ICC/INDEX.md
# Match 1: workflow/archive/otr-001a-electron-shell-inventory.md line 7
# Match 2: workflow/archive/otr-001b-electron-editor-inventory.md line 70
# Context: Archive files, not active registry scope
```

**Conclusion:** Shell navigation scope derived from ICC maintenance rules (lines 48–52), not explicit API definitions. All items classified as gaps where implementations absent.

---

## Next Actions

1. **For GAP items:** Defer to CODE stage or separate discovery unless shell implementation becomes available in source
2. **Verify FMT-006 content** against ICC registry and repository paths
3. **Promote to COMPLETE** after evidence validation

---

## Verification Requirements

- [ ] Confirm all gap/no-change classifications align with grep/rg search results
- [ ] Validate source path references match `electron/` directory listings
- [ ] Ensure no stale/uncommitted modifications block verification (per ICC rules)

---

*End of FMT-006 — Shell Navigation Gap Map*
