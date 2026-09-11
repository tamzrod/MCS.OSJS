param(
  [switch]$NoPush
)

$ErrorActionPreference = 'Stop'

$ElectronDir = $PSScriptRoot
$RepoRoot = Split-Path -Parent $ElectronDir
$BinDir = Join-Path $ElectronDir 'bin'
$DistDir = Join-Path $ElectronDir 'dist'

$NssmVersion = '2.24-101-g897c7ad'
$NssmUrl = "https://nssm.cc/ci/nssm-$NssmVersion.zip"
$NssmSha256 = '99F5045FFFBFFB745D67FE3A065A953C4A3D9C253B868892D9B685B0EE7D07B8'

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

function Assert-LastExitCode {
  param([string]$CommandName)

  if ($LASTEXITCODE -ne 0) {
    throw "$CommandName failed with exit code $LASTEXITCODE"
  }
}

function Build-GoModule {
  param(
    [string]$Label,
    [string]$ModuleDir,
    [string]$Package,
    [string]$Output
  )

  Step "Building $Label for Windows"
  Push-Location $ModuleDir
  try {
    go build -o $Output $Package
    Assert-LastExitCode "$Label go build"
  }
  finally {
    Pop-Location
  }
}

function Ensure-Nssm {
  param([string]$Destination)

  if (Test-Path $Destination) {
    Write-Host "NSSM already present: $Destination" -ForegroundColor DarkGray
    return
  }

  Step "Downloading NSSM $NssmVersion"

  $tempRoot = Join-Path ([System.IO.Path]::GetTempPath()) ("mcs-nssm-" + [guid]::NewGuid().ToString('N'))
  $zipPath = Join-Path $tempRoot 'nssm.zip'
  $extractDir = Join-Path $tempRoot 'extract'

  New-Item -ItemType Directory -Force -Path $tempRoot | Out-Null

  try {
    Invoke-WebRequest -Uri $NssmUrl -OutFile $zipPath -UseBasicParsing

    $actualHash = (Get-FileHash -Path $zipPath -Algorithm SHA256).Hash.ToUpperInvariant()
    if ($actualHash -ne $NssmSha256) {
      throw "NSSM checksum mismatch. Expected $NssmSha256, got $actualHash"
    }

    Expand-Archive -Path $zipPath -DestinationPath $extractDir -Force

    $archDir = if ([Environment]::Is64BitOperatingSystem) { 'win64' } else { 'win32' }
    $source = Join-Path $extractDir "nssm-$NssmVersion\$archDir\nssm.exe"

    if (-not (Test-Path $source)) {
      throw "Downloaded NSSM archive did not contain expected file: $source"
    }

    Copy-Item -Path $source -Destination $Destination -Force
    Write-Host "NSSM ready: $Destination" -ForegroundColor Green
  }
  finally {
    if (Test-Path $tempRoot) {
      Remove-Item -Recurse -Force $tempRoot
    }
  }
}

Require-Command git
Require-Command go
Require-Command npm

Step 'Checking repository'
Set-Location $RepoRoot
$branch = (git branch --show-current).Trim()
Assert-LastExitCode 'git branch'
if ($branch -ne 'main') {
  throw "Expected branch 'main', current branch is '$branch'."
}

git pull --ff-only origin main
Assert-LastExitCode 'git pull'

Step 'Preparing Electron binary directory'
New-Item -ItemType Directory -Force -Path $BinDir | Out-Null

$nssm = Join-Path $BinDir 'nssm.exe'
Ensure-Nssm -Destination $nssm

Build-GoModule -Label 'MMA2' `
  -ModuleDir (Join-Path $RepoRoot 'MMA2') `
  -Package '.\cmd\mma2' `
  -Output (Join-Path $BinDir 'mma2.exe')

Build-GoModule -Label 'MMA2 supervisor' `
  -ModuleDir (Join-Path $RepoRoot 'MMA2') `
  -Package '.\cmd\mma2-supervisor' `
  -Output (Join-Path $BinDir 'mma2-supervisor.exe')

Build-GoModule -Label 'Simulator runtime' `
  -ModuleDir (Join-Path $RepoRoot 'simulator') `
  -Package '.\cmd\modbus-simulator-runtime' `
  -Output (Join-Path $BinDir 'modbus-simulator-runtime.exe')

Build-GoModule -Label 'Replicator runtime' `
  -ModuleDir (Join-Path $RepoRoot 'replicator') `
  -Package '.\cmd\modbus-replicator-runtime' `
  -Output (Join-Path $BinDir 'modbus-replicator-runtime.exe')

Step 'Installing Electron dependencies'
Set-Location $ElectronDir
npm install
Assert-LastExitCode 'npm install'

Step 'Building Windows installer'
npm run dist:win
Assert-LastExitCode 'npm run dist:win'

$installer = Get-ChildItem -Path $DistDir -Filter 'MCS-Modbus-Toolkit-*-Setup.exe' -File |
  Sort-Object LastWriteTime -Descending |
  Select-Object -First 1

if (-not $installer) {
  throw "electron-builder reported success but no installer was found in $DistDir"
}

Write-Host "`nInstaller: $($installer.FullName)" -ForegroundColor Green

if (-not $NoPush) {
  Step 'Pushing committed source changes to origin/main'
  Set-Location $RepoRoot

  $dirty = git status --porcelain
  Assert-LastExitCode 'git status'
  if ($dirty) {
    Write-Host 'Working tree has uncommitted changes. They will NOT be committed by this script.' -ForegroundColor Yellow
    git status --short
    Assert-LastExitCode 'git status --short'
  }

  git push origin main
  Assert-LastExitCode 'git push'
}

Write-Host "`nDONE" -ForegroundColor Green
