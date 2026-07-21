# Cross-build sshctl for Linux / Windows / macOS
$ErrorActionPreference = "Stop"
$Version = if ($env:VERSION) { $env:VERSION } else { "0.1.0" }
$ld = "-s -w -X github.com/Fracizz/sshctl/cmd.Version=$Version"
New-Item -ItemType Directory -Force -Path dist | Out-Null

$targets = @(
  @{ GOOS="linux";   GOARCH="amd64"; Out="dist/sshctl-linux-amd64" },
  @{ GOOS="linux";   GOARCH="arm64"; Out="dist/sshctl-linux-arm64" },
  @{ GOOS="windows"; GOARCH="amd64"; Out="dist/sshctl-windows-amd64.exe" },
  @{ GOOS="windows"; GOARCH="arm64"; Out="dist/sshctl-windows-arm64.exe" },
  @{ GOOS="darwin";  GOARCH="amd64"; Out="dist/sshctl-darwin-amd64" },
  @{ GOOS="darwin";  GOARCH="arm64"; Out="dist/sshctl-darwin-arm64" }
)

foreach ($t in $targets) {
  $env:GOOS = $t.GOOS
  $env:GOARCH = $t.GOARCH
  $env:CGO_ENABLED = "0"
  Write-Host "building $($t.Out)"
  go build -ldflags $ld -o $t.Out .
}
Remove-Item Env:GOOS, Env:GOARCH -ErrorAction SilentlyContinue
Write-Host "done"
