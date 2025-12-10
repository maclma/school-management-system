#!/usr/bin/env bash
set -euo pipefail
cd "$(dirname "$0")/.."

echo "Installing developer tools..."

echo "Installing golangci-lint..."
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

echo "Installing air (optional live reload)..."
go install github.com/cosmtrek/air@latest

echo "Done. Ensure your PATH includes: "; go env GOPATH
