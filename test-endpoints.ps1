# Test specific endpoints to see what's wrong
$token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImFkbWluQHNjaG9vbC5jb20iLCJleHAiOjE3NzAxMzczMzgsImlhdCI6MTc3MDA1MDkzOCwicm9sZSI6ImFkbWluIiwidXNlcl9pZCI6MX0.QV4q5ahtUAJdMbEWOj3hYcmF7yx-vfjc_pxpXUFB2Wo"

$headers = @{
    "Authorization" = "Bearer $token"
    "Content-Type" = "application/json"
}

# 1. Test Grades - List all grades for student
Write-Host "=== TESTING: GET /api/grades/by-student/1 ===" 
try {
    $response = Invoke-RestMethod -Uri "http://localhost:8080/api/grades/by-student/1" -Method GET -Headers $headers
    Write-Host "✅ Success: $(($response.data | Measure-Object).Count) grades"
} catch {
    Write-Host "❌ Error: $($_.Exception.Response.StatusCode) - $($_.Exception.Message)"
}

# 2. Test Attendance - List all attendance for student
Write-Host "`n=== TESTING: GET /api/attendance/by-student/1 ===" 
try {
    $response = Invoke-RestMethod -Uri "http://localhost:8080/api/attendance/by-student/1" -Method GET -Headers $headers
    Write-Host "✅ Success: $(($response.data | Measure-Object).Count) attendance records"
} catch {
    Write-Host "❌ Error: $($_.Exception.Response.StatusCode) - $($_.Exception.Message)"
}

# 3. Test Announcements - List all
Write-Host "`n=== TESTING: GET /api/announcements ===" 
try {
    $response = Invoke-RestMethod -Uri "http://localhost:8080/api/announcements" -Method GET -Headers $headers -ErrorAction Stop
    Write-Host "✅ Success: $(($response.data | Measure-Object).Count) announcements"
} catch {
    Write-Host "❌ Error: $($_.Exception.Response.StatusCode)"
    Write-Host "Full error: $($_.Exception | ConvertTo-Json)"
}

# 4. Test Messages - Conversations
Write-Host "`n=== TESTING: GET /api/messages/inbox ===" 
try {
    $response = Invoke-RestMethod -Uri "http://localhost:8080/api/messages/inbox" -Method GET -Headers $headers
    Write-Host "✅ Success: $(($response.data | Measure-Object).Count) messages"
} catch {
    Write-Host "❌ Error: $($_.Exception.Response.StatusCode) - $($_.Exception.Message)"
}

# 5. Test Assignments - List all
Write-Host "`n=== TESTING: GET /api/assignments/course/1 ===" 
try {
    $response = Invoke-RestMethod -Uri "http://localhost:8080/api/assignments/course/1" -Method GET -Headers $headers
    Write-Host "✅ Success: $(($response.data | Measure-Object).Count) assignments"
} catch {
    Write-Host "❌ Error: $($_.Exception.Response.StatusCode) - $($_.Exception.Message)"
}
