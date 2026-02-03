# Test Login
$loginData = @{
    email = "admin@school.com"
    password = "admin123"
} | ConvertTo-Json

$response = Invoke-RestMethod -Uri "http://localhost:8080/api/auth/login" -Method POST -Body $loginData -ContentType "application/json"
Write-Host "Login Response:"
$response | ConvertTo-Json

# Save token for testing
$token = $response.token
Write-Host "`nToken: $token"

# Test dashboard endpoint
Write-Host "`nTesting dashboard endpoint..."
$headers = @{
    "Authorization" = "Bearer $token"
    "Content-Type" = "application/json"
}

try {
    $dashResponse = Invoke-RestMethod -Uri "http://localhost:8080/api/admin/dashboard" -Method GET -Headers $headers
    Write-Host "Dashboard Response:"
    $dashResponse | ConvertTo-Json -Depth 3
} catch {
    Write-Host "Error: $($_.Exception.Message)"
}

# Test get users
Write-Host "`nTesting get users..."
try {
    $usersResponse = Invoke-RestMethod -Uri "http://localhost:8080/api/admin/users?page=1&limit=10" -Method GET -Headers $headers
    Write-Host "Users Response:"
    $usersResponse | ConvertTo-Json -Depth 3
} catch {
    Write-Host "Error: $($_.Exception.Message)"
}
