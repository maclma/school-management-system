# Demo script for School Management System API
# Starts the server and runs a sample workflow

Write-Host "School Management System - Demo Script" -ForegroundColor Green
Write-Host "======================================" -ForegroundColor Green
Write-Host ""

# Start server in background
Write-Host "Starting server on port 8081..." -ForegroundColor Cyan
$serverEnv = @{
    DB_DRIVER   = 'sqlite'
    DB_PATH     = 'school_demo.db'
    JWT_SECRET  = 'changeme'
    SERVER_PORT = '8081'
}
$serverEnv.GetEnumerator() | ForEach-Object { Set-Item "env:$($_.Key)" $_.Value }

# Start server process
$server = Start-Process -FilePath "go" -ArgumentList "run", "./cmd/server" -NoNewWindow -PassThru
Write-Host "Server PID: $($server.Id)" -ForegroundColor Yellow
Write-Host "Waiting for server to start..." -ForegroundColor Cyan
Start-Sleep -Seconds 3

$baseUrl = "http://localhost:8081"

try {
    # Generate unique email
    $timestamp = [DateTimeOffset]::UtcNow.ToUnixTimeMilliseconds()
    $randomId = Get-Random -Minimum 1000 -Maximum 9999
    $email = "demo_${randomId}_${timestamp}@example.com"
    
    Write-Host ""
    Write-Host "=== 1. REGISTER USER ===" -ForegroundColor Green
    Write-Host "Email: $email" -ForegroundColor Yellow
    
    $registerPayload = @{
        first_name = "Demo"
        last_name  = "User"
        email      = $email
        password   = "password123"
        role       = "student"
    } | ConvertTo-Json
    
    $registerResponse = Invoke-RestMethod -Method Post `
        -Uri "$baseUrl/api/auth/register" `
        -ContentType 'application/json' `
        -Body $registerPayload
    
    Write-Host "Response:" -ForegroundColor Cyan
    $registerResponse | ConvertTo-Json -Depth 4 | Write-Host
    
    Write-Host ""
    Write-Host "=== 2. LOGIN ===" -ForegroundColor Green
    
    $loginPayload = @{
        email    = $email
        password = "password123"
    } | ConvertTo-Json
    
    $loginResponse = Invoke-RestMethod -Method Post `
        -Uri "$baseUrl/api/auth/login" `
        -ContentType 'application/json' `
        -Body $loginPayload
    
    Write-Host "Response:" -ForegroundColor Cyan
    $loginResponse | ConvertTo-Json -Depth 4 | Write-Host
    
    $token = $loginResponse.token
    
    Write-Host ""
    Write-Host "=== 3. GET PROFILE ===" -ForegroundColor Green
    Write-Host "Token: $($token.Substring(0, 20))..." -ForegroundColor Yellow
    
    $profileResponse = Invoke-RestMethod -Method Get `
        -Uri "$baseUrl/api/profile" `
        -Headers @{ Authorization = "Bearer $token" }
    
    Write-Host "Response:" -ForegroundColor Cyan
    $profileResponse | ConvertTo-Json -Depth 4 | Write-Host
    
    Write-Host ""
    Write-Host "=== 4. UPDATE PROFILE ===" -ForegroundColor Green
    
    $updatePayload = @{
        first_name = "Updated"
        last_name  = "Demo User"
        phone      = "+1234567890"
    } | ConvertTo-Json
    
    $updateResponse = Invoke-RestMethod -Method Put `
        -Uri "$baseUrl/api/profile" `
        -ContentType 'application/json' `
        -Headers @{ Authorization = "Bearer $token" } `
        -Body $updatePayload
    
    Write-Host "Response:" -ForegroundColor Cyan
    $updateResponse | ConvertTo-Json -Depth 4 | Write-Host
    
    Write-Host ""
    Write-Host "=== Demo Complete ===" -ForegroundColor Green
    
} catch {
    Write-Host "Error: $_" -ForegroundColor Red
} finally {
    Write-Host ""
    Write-Host "Stopping server..." -ForegroundColor Cyan
    Stop-Process -Id $server.Id -Force -ErrorAction SilentlyContinue
    Write-Host "Done!" -ForegroundColor Green
}
