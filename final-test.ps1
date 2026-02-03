# Comprehensive Test Suite - Final Version
Write-Host "`n================== SCHOOL MANAGEMENT SYSTEM - TEST REPORT ==================`n" -ForegroundColor Cyan

$token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImFkbWluQHNjaG9vbC5jb20iLCJleHAiOjE3NzAxMzczMzgsImlhdCI6MTc3MDA1MDkzOCwicm9sZSI6ImFkbWluIiwidXNlcl9pZCI6MX0.QV4q5ahtUAJdMbEWOj3hYcmF7yx-vfjc_pxpXUFB2Wo"
$headers = @{"Authorization" = "Bearer $token"; "Content-Type" = "application/json"}
$baseUrl = "http://localhost:8080/api"

$passed = 0
$failed = 0

function Test-API {
    param([string]$Name, [string]$Method, [string]$Endpoint)
    
    $url = "$baseUrl$Endpoint"
    try {
        $resp = Invoke-WebRequest -Uri $url -Method $Method -Headers $headers -UseBasicParsing -ErrorAction Stop
        if ($resp.StatusCode -eq 200) {
            Write-Host "  [OK]     $Name" -ForegroundColor Green
            return 1
        }
    }
    catch {
        Write-Host "  [FAIL]   $Name" -ForegroundColor Red
        return 0
    }
    return 0
}

Write-Host "[1] CORE & AUTH" -ForegroundColor Yellow
$passed += Test-API "Health Check" "GET" "/health"
$passed += Test-API "Admin Dashboard" "GET" "/admin/health"

Write-Host "`n[2] USER & PROFILE MANAGEMENT" -ForegroundColor Yellow
$passed += Test-API "List All Users" "GET" "/users"
$passed += Test-API "Admin Users View" "GET" "/admin/users"
$passed += Test-API "My Profile" "GET" "/profile"

Write-Host "`n[3] COURSES & STUDENTS" -ForegroundColor Yellow
$passed += Test-API "All Courses" "GET" "/courses"
$passed += Test-API "All Students" "GET" "/students"

Write-Host "`n[4] ENROLLMENT SYSTEM" -ForegroundColor Yellow
$passed += Test-API "All Enrollments" "GET" "/admin/enrollments"

Write-Host "`n[5] ANNOUNCEMENTS" -ForegroundColor Yellow
$passed += Test-API "All Announcements" "GET" "/announcements"
$passed += Test-API "Active Announcements" "GET" "/announcements/active"
$passed += Test-API "Search Announcements" "GET" "/search/announcements"

Write-Host "`n[6] NOTIFICATIONS" -ForegroundColor Yellow
$passed += Test-API "My Notifications" "GET" "/notifications"
$passed += Test-API "Unread Notifications" "GET" "/notifications/unread"

Write-Host "`n[7] MESSAGING SYSTEM" -ForegroundColor Yellow
$passed += Test-API "Message Inbox" "GET" "/messages/inbox"
$passed += Test-API "Unread Count" "GET" "/messages/unread"

Write-Host "`n[8] PAYMENTS & FINANCE" -ForegroundColor Yellow
$passed += Test-API "All Payments" "GET" "/payments"
$passed += Test-API "Search Payments" "GET" "/search/payments"
$passed += Test-API "Overdue Payments" "GET" "/search/overdue-payments"

Write-Host "`n[9] SCHEDULE & TIMETABLE" -ForegroundColor Yellow
$passed += Test-API "All Timetables" "GET" "/timetable"

Write-Host "`n[10] TEACHER MANAGEMENT" -ForegroundColor Yellow
$passed += Test-API "List Teachers" "GET" "/admin/teachers"

Write-Host "`n[11] EXPORTS & REPORTING" -ForegroundColor Yellow
$passed += Test-API "Export Payments" "GET" "/export/payments"
$passed += Test-API "Export Grades" "GET" "/export/grades"
$passed += Test-API "Export Attendance" "GET" "/export/attendance"
$passed += Test-API "Export Enrollments" "GET" "/export/enrollments"

Write-Host "`n[12] SYSTEM SETTINGS" -ForegroundColor Yellow
$passed += Test-API "Get Settings" "GET" "/admin/settings"

Write-Host "`n=================================================================================" -ForegroundColor Cyan
Write-Host "TOTAL TESTS PASSED: $passed out of 28" -ForegroundColor Green
Write-Host "SUCCESS RATE: $(([math]::Round(($passed/28)*100,1)))%" -ForegroundColor Cyan
Write-Host "=================================================================================" -ForegroundColor Cyan
