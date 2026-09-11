param(
  [switch]$NoPush
)

$ErrorActionPreference = 'Stop'

$ElectronDir = $PSScriptRoot
$RepoRoot = Split-Path -Parent $ElectronDir
$BinDir = Join-Path $ElectronDir 'bin'
$DistDir = Join-Path $ElectronDir 'dist'

function Require-Command {
  param([string]$Name)

  if (-not (Get-Command $Name -ErrorAction SilentlyContinue)) {
    throw "Required command not found: $Name"
  }
}

function Step {
  param([string]$Text)
  Write-Host "`n==> $Text" -ForegroundColor Cyan
}

Require-Command git
Require-Command go
Require-Command npm

Step 'Checking repository'
Set-Location $RepoRoot
$branch = (git branch --show-current).Trim()
if ($branch -ne 'main') {
  throw "Expected branch 'main', current branch is '$branch'."
}

git pull --ff-only origin main

Step 'Preparing Electron binary directory'
New-Item -ItemType Directory -Force -Path $BinDir | Out-Null

$nssm = Join-Path $BinDir 'nssm.exe'
if (-not (Test-Path $nssm)) {
  throw "Missing $nssm. Copy the Windows NSSM executable there before building."
}

Step 'Building MMA2 for Windows'
go build -o (Join-Path $BinDir 'mma2.exe') .\MMA2\cmd\mma2

Step 'Building Simulator runtime for Windows'
go build -o (Join-Path $BinDir 'modbus-simulator-runtime.exe') .\simulator\cmd\modbus-simulator-runtime

Step 'Building Replicator runtime for Windows'
go build -o (Join-Path $BinDir 'modbus-replicator-runtime.exe') .\replicator\cmd\modbus-replicator-runtime

Step 'Installing Electron dependencies'
Set-Location $ElectronDir
npm install

Step 'Building Windows installer'
npm run dist:win

$installer = Get-ChildItem -Path $DistDir -Filter 'MCS-Modbus-Toolkit-*-Setup.exe' -File |
  Sort-Object LastWriteTime -Descending |
  Select-Object -First 1

if (-not $installer) {
  throw "Build completed but no installer was found in $DistDir"
}

Write-Host "`nInstaller: $($installer.FullName)" -ForegroundColor Green

if (-not $NoPush) {
  Step 'Pushing committed source changes to origin/main'
  Set-Location $RepoRoot

  $dirty = git status --porcelain
  if ($dirty) {
    Write-Host 'Working tree has uncommitted changes. They will NOT be committed by this script.' -ForegroundColor Yellow
    git status --short
  }

  git push origin main
}

Write-Host "`nDONE" -ForegroundColor Green
