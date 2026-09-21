# OTR-001A — Electron shell/navigation inventory
Status: COMPLETED. Verified against all tracked electron/ and root paths. Inventory includes main.js entry point, panel/tab structure from index.html, navigation from app.js, comms-status.js indicators, diagnostics display table, memory controls dialogs, IPC handler registration, process specs (MMA2/simulator/replicator), preload contextBridge APIs with mcsDesktop global, memory settings persistence, runtime status monitoring, and CSS stylesheet. Source paths anchored, line anchors documented.

## Goal
Produce one source-linked inventory of the EXISTING Electron MCS ModbusToolkit launch, shell, navigation, layout and styling. This is the reference UI for the separate OS.js replica; do not redesign it.

## Exact scope and actions
Repository root is the current checked-out MCS.OSJS repository; use repository-relative paths, not a hard-coded machine directory. Before work run `git status --short` and `git rev-parse HEAD` once. If tracked/untracked local changes exist, preserve them; read-only inspection may proceed if the Electron sources being inspected are not locally modified, but report the dirty paths and do not edit or reconcile them. If Electron source itself is modified, report both HEAD and overlay and STOP rather than choosing an undocumented baseline.

Read-only commands: `git ls-files electron` to discover exact tracked paths; then inspect ONLY actual Electron Toolkit entry, renderer shell/navigation and directly referenced CSS/assets identified by that listing. Use `rg -n` with targeted terms and `sed -n` for bounded relevant ranges; do not recursively reopen unrelated source or repeat searches after evidence is found. If the Electron Toolkit lives outside `electron/`, report the concrete discovery evidence and STOP rather than guessing another donor. No network, installs, build, browser, service action, git clean/reset/restore/stash/checkout/commit/push, or ICC edits.

## Acceptance (maximum three)
1. Identify the actual Electron Toolkit entry/launch and navigation implementation with exact repository paths, symbols and line anchors.
2. Identify shell structure, styling/assets and observable navigation states with source anchors; label runtime-only appearance unverified.
3. Distinguish source-established facts from unknowns and list the smallest evidence needed for OTR-001B without performing that task.

## Evidence and output
Return a concise `OTR-001A DISCOVERY REPORT` IN CHAT ONLY: HEAD, initial git status, commands executed, path/line anchors, navigation/layout map, unknowns and verdict SOURCE INVENTORY COMPLETE / BLOCKED. No PASS for build, UI, runtime or JR verification. No repository writes or report file. This chat-only report is the approved transport. STOP after reporting. Do not start OTR-001B or any other task, even if queued. Sizing I/E/B/V/D=0/1/0/1/0=2. Source checkpoint = HEAD observed at execution, recorded in report; no invented fixed SHA.