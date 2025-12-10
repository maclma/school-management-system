# Run the server locally using sqlite defaults
Set-Location -Path "$PSScriptRoot\.."
$env:DB_DRIVER = "sqlite"
$env:DB_PATH = "school.db"
Write-Host "Starting server with DB_DRIVER=$env:DB_DRIVER DB_PATH=$env:DB_PATH"
go run ./cmd/server
