$ErrorActionPreference = 'Stop'
$ts = Get-Date -Format yyyyMMddHHmmss
$email = "testuser_$ts@example.com"
$password = "Pass123!"

$reg = @{ email = $email; password = $password; first_name = 'Test'; last_name = 'User'; role = 'student' } | ConvertTo-Json
Write-Host "Registering $email..."
$regResp = Invoke-RestMethod -Uri 'http://localhost:8080/api/auth/register' -Method Post -ContentType 'application/json' -Body $reg
Write-Host "REGISTER_RESPONSE:"; $regResp | ConvertTo-Json

$login = @{ email = $email; password = $password } | ConvertTo-Json
$resp = Invoke-RestMethod -Uri 'http://localhost:8080/api/auth/login' -Method Post -ContentType 'application/json' -Body $login
Write-Host "LOGIN_SUCCESS:"; $resp | ConvertTo-Json
