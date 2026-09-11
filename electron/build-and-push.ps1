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
if ($branch -ne 'main') {
  throw "Expected branch 'main', current branch is '$branch'."
}

git pull --ff-only origin main

Step 'Preparing Electron binary directory'
New-Item -ItemType Directory -Force -Path $BinDir | Out-Null

$nssm = Join-Path $BinDir 'nssm.exe'
Ensure-Nssm -Destination $nssm

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
