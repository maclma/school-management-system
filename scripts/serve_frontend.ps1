# Run this from the repository root in PowerShell
# Starts the small Go static server serving the ./frontend directory on port 3000

Write-Host "Starting frontend server (http://localhost:3000)"
pushd (Split-Path -Path $PSScriptRoot -Parent) | Out-Null
# If you run this script directly, repo root may already be current dir
# Try to run go if available
if (Get-Command go -ErrorAction SilentlyContinue) {
    go run .\frontend\server.go
} else {
    Write-Host "Go is not installed or not on PATH. Install Go or run another static server (python/http-server)."
}
popd | Out-Null
