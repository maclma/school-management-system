# Quick Feature Status Check
# Testing which endpoints actually work

$token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImFkbWluQHNjaG9vbC5jb20iLCJleHAiOjE3NzAxMzczMzgsImlhdCI6MTc3MDA1MDkzOCwicm9sZSI6ImFkbWluIiwidXNlcl9pZCI6MX0.QV4q5ahtUAJdMbEWOj3hYcmF7yx-vfjc_pxpXUFB2Wo"
$headers = @{"Authorization" = "Bearer $token"; "Content-Type" = "application/json"}

$endpoints = @(
    @{name = "Users"; url = "http://localhost:8080/api/users"; method = "GET"},
    @{name = "Courses"; url = "http://localhost:8080/api/courses"; method = "GET"},
    @{name = "Enrollments"; url = "http://localhost:8080/api/admin/enrollments"; method = "GET"},
    @{name = "Timetables"; url = "http://localhost:8080/api/timetable"; method = "GET"},
    @{name = "Grades (by student)"; url = "http://localhost:8080/api/grades/by-student/1"; method = "GET"},
    @{name = "Attendance (by student)"; url = "http://localhost:8080/api/attendance/by-student/1"; method = "GET"},
    @{name = "Announcements"; url = "http://localhost:8080/api/announcements"; method = "GET"},
    @{name = "Notifications"; url = "http://localhost:8080/api/notifications"; method = "GET"},
    @{name = "Messages"; url = "http://localhost:8080/api/messages/inbox"; method = "GET"},
    @{name = "Assignments"; url = "http://localhost:8080/api/assignments/course/1"; method = "GET"},
    @{name = "Payments"; url = "http://localhost:8080/api/payments"; method = "GET"},
    @{name = "Transcripts"; url = "http://localhost:8080/api/transcripts/gpa/1"; method = "GET"},
    @{name = "Teachers"; url = "http://localhost:8080/api/admin/teachers"; method = "GET"},
    @{name = "Students"; url = "http://localhost:8080/api/students"; method = "GET"}
)

Write-Host "FEATURE STATUS REPORT`n" -ForegroundColor Cyan
Write-Host "=" * 50

foreach ($endpoint in $endpoints) {
    try {
        $result = Invoke-WebRequest -Uri $endpoint.url -Method $endpoint.method -Headers $headers -UseBasicParsing -ErrorAction Stop
        $status = if ($result.StatusCode -eq 200) { "✅ WORKING" } else { "⚠️  CODE " + $result.StatusCode }
        Write-Host "$($endpoint.name.PadRight(25)) : $status" -ForegroundColor Green
    } catch {
        $code = $_.Exception.Response.StatusCode.Value__
        $status = "❌ ERROR $code"
        Write-Host "$($endpoint.name.PadRight(25)) : $status" -ForegroundColor Red
    }
}

Write-Host "`n" + "=" * 50
Write-Host "Legend: ✅ = Working | ⚠️  = Different code | ❌ = Error" -ForegroundColor Yellow
