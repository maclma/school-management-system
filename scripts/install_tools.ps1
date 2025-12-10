#!/usr/bin/env pwsh
Set-Location -Path "$PSScriptRoot\.."

Write-Host "Installing developer tools..."

Write-Host "Installing golangci-lint..."
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

Write-Host "Installing air (optional live reload)..."
go install github.com/cosmtrek/air@latest

Write-Host "Done. Ensure your PATH includes $(go env GOPATH)\bin"
