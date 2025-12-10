$ErrorActionPreference = 'Stop'

$loginBody = @{ 
    email = "john.doe@example.com"; 
    password = "SecurePass123!" 
} | ConvertTo-Json

try {
    $resp = Invoke-RestMethod -Uri 'http://localhost:8080/api/auth/login' -Method Post -ContentType 'application/json' -Body $loginBody
    Write-Host "LOGIN_SUCCESS:"
    $resp | ConvertTo-Json
} catch {
    Write-Host "LOGIN_FAILED, attempting registration..."
    $reg = @{ 
        email = "john.doe@example.com"; 
        password = "SecurePass123!"; 
        first_name = "John"; 
        last_name = "Doe"; 
        role = "student" 
    } | ConvertTo-Json

    $registerResp = Invoke-RestMethod -Uri 'http://localhost:8080/api/auth/register' -Method Post -ContentType 'application/json' -Body $reg
    Write-Host "REGISTER_RESPONSE:"
    $registerResp | ConvertTo-Json

    Start-Sleep -Seconds 1

    $resp = Invoke-RestMethod -Uri 'http://localhost:8080/api/auth/login' -Method Post -ContentType 'application/json' -Body $loginBody
    Write-Host "LOGIN_AFTER_REGISTER_SUCCESS:"
    $resp | ConvertTo-Json
}
