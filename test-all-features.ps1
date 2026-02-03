# Test all features
$token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImFkbWluQHNjaG9vbC5jb20iLCJleHAiOjE3NzAxMzczMzgsImlhdCI6MTc3MDA1MDkzOCwicm9sZSI6ImFkbWluIiwidXNlcl9pZCI6MX0.QV4q5ahtUAJdMbEWOj3hYcmF7yx-vfjc_pxpXUFB2Wo"

$headers = @{
    "Authorization" = "Bearer $token"
    "Content-Type" = "application/json"
}

# 1. Test Courses
Write-Host "=== TESTING COURSES ===" 
try {
    $response = Invoke-RestMethod -Uri "http://localhost:8080/api/courses" -Method GET -Headers $headers
    Write-Host "✅ Courses: $($response.data.Count) courses found"
} catch {
    Write-Host "❌ Courses Error: $($_.Exception.Message)"
}

# 2. Test Grades
Write-Host "`n=== TESTING GRADES ===" 
try {
    $response = Invoke-RestMethod -Uri "http://localhost:8080/api/grades" -Method GET -Headers $headers
    Write-Host "✅ Grades: Working"
} catch {
    Write-Host "❌ Grades Error: $($_.Exception.Message)"
}

# 3. Test Attendance
Write-Host "`n=== TESTING ATTENDANCE ===" 
try {
    $response = Invoke-RestMethod -Uri "http://localhost:8080/api/attendance" -Method GET -Headers $headers
    Write-Host "✅ Attendance: Working"
} catch {
    Write-Host "❌ Attendance Error: $($_.Exception.Message)"
}

# 4. Test Timetables
Write-Host "`n=== TESTING TIMETABLES ===" 
try {
    $response = Invoke-RestMethod -Uri "http://localhost:8080/api/timetable" -Method GET -Headers $headers
    Write-Host "✅ Timetables: $($response.data.Count) timetables found"
} catch {
    Write-Host "❌ Timetables Error: $($_.Exception.Message)"
}

# 5. Test Announcements
Write-Host "`n=== TESTING ANNOUNCEMENTS ===" 
try {
    $response = Invoke-RestMethod -Uri "http://localhost:8080/api/announcements" -Method GET -Headers $headers
    Write-Host "✅ Announcements: $($response.data.Count) announcements found"
} catch {
    Write-Host "❌ Announcements Error: $($_.Exception.Message)"
}

# 6. Test Notifications
Write-Host "`n=== TESTING NOTIFICATIONS ===" 
try {
    $response = Invoke-RestMethod -Uri "http://localhost:8080/api/notifications" -Method GET -Headers $headers
    Write-Host "✅ Notifications: $($response.data.Count) notifications found"
} catch {
    Write-Host "❌ Notifications Error: $($_.Exception.Message)"
}

# 7. Test Messages
Write-Host "`n=== TESTING MESSAGES ===" 
try {
    $response = Invoke-RestMethod -Uri "http://localhost:8080/api/messages/conversations" -Method GET -Headers $headers
    Write-Host "✅ Messages: Working"
} catch {
    Write-Host "❌ Messages Error: $($_.Exception.Message)"
}

# 8. Test Enrollments
Write-Host "`n=== TESTING ENROLLMENTS ===" 
try {
    $response = Invoke-RestMethod -Uri "http://localhost:8080/api/admin/enrollments" -Method GET -Headers $headers
    Write-Host "✅ Enrollments: $($response.data.Count) enrollments found"
} catch {
    Write-Host "❌ Enrollments Error: $($_.Exception.Message)"
}

# 9. Test Assignments
Write-Host "`n=== TESTING ASSIGNMENTS ===" 
try {
    $response = Invoke-RestMethod -Uri "http://localhost:8080/api/assignments" -Method GET -Headers $headers
    Write-Host "✅ Assignments: Working"
} catch {
    Write-Host "❌ Assignments Error: $($_.Exception.Message)"
}

# 10. Test Transcripts
Write-Host "`n=== TESTING TRANSCRIPTS ===" 
try {
    $response = Invoke-RestMethod -Uri "http://localhost:8080/api/transcripts/gpa/1" -Method GET -Headers $headers
    Write-Host "✅ Transcripts: Working"
} catch {
    Write-Host "❌ Transcripts Error: $($_.Exception.Message)"
}

# 11. Test Payments
Write-Host "`n=== TESTING PAYMENTS ===" 
try {
    $response = Invoke-RestMethod -Uri "http://localhost:8080/api/payments" -Method GET -Headers $headers
    Write-Host "✅ Payments: Working"
} catch {
    Write-Host "❌ Payments Error: $($_.Exception.Message)"
}
