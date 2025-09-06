$ErrorActionPreference = "Stop"
$APP = "ctfmini"
$VERSION = $env:VERSION
if (-not $VERSION) { $VERSION = "0.1.0" }
$OUT = "dist"
mkdir $OUT -Force | Out-Null

$targets = @(
  @{ GOOS="windows"; GOARCH="amd64"; OUTFILE="$APP-windows-amd64.exe" },
  @{ GOOS="linux";   GOARCH="amd64"; OUTFILE="$APP-linux-amd64" },
  @{ GOOS="darwin";  GOARCH="arm64"; OUTFILE="$APP-macos-arm64" },
  @{ GOOS="darwin";  GOARCH="amd64"; OUTFILE="$APP-macos-amd64" }
)

foreach ($t in $targets) {
  $env:GOOS = $t.GOOS
  $env:GOARCH = $t.GOARCH
  go build -ldflags="-s -w" -o "$OUT/$($t.OUTFILE)"
  Compress-Archive -Path "$OUT/$($t.OUTFILE)" -DestinationPath "$OUT/$APP-$VERSION-$($t.GOOS)-$($t.GOARCH).zip" -Force
}

Write-Host "Built zips in $OUT/"
