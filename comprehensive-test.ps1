# Comprehensive Test Suite for School Management System

Write-Host "╔════════════════════════════════════════════════════════════════════════╗" -ForegroundColor Cyan
Write-Host "║   SCHOOL MANAGEMENT SYSTEM - COMPREHENSIVE FUNCTIONALITY TEST          ║" -ForegroundColor Cyan
Write-Host "║   Date: $(Get-Date -Format 'yyyy-MM-dd HH:mm:ss')                                     ║" -ForegroundColor Cyan
Write-Host "╚════════════════════════════════════════════════════════════════════════╝" -ForegroundColor Cyan

# Admin token
$token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImFkbWluQHNjaG9vbC5jb20iLCJleHAiOjE3NzAxMzczMzgsImlhdCI6MTc3MDA1MDkzOCwicm9sZSI6ImFkbWluIiwidXNlcl9pZCI6MX0.QV4q5ahtUAJdMbEWOj3hYcmF7yx-vfjc_pxpXUFB2Wo"
$headers = @{"Authorization" = "Bearer $token"; "Content-Type" = "application/json"}
$baseUrl = "http://localhost:8080/api"

# Test counters
$totalTests = 0
$passedTests = 0
$failedTests = 0

function Test-Endpoint {
    param(
        [string]$Name,
        [string]$Method,
        [string]$Endpoint,
        [object]$Body = $null,
        [int]$ExpectedStatus = 200
    )
    
    $totalTests++
    $url = "$baseUrl$Endpoint"
    
    try {
        $params = @{
            Uri = $url
            Method = $Method
            Headers = $headers
            UseBasicParsing = $true
            ErrorAction = 'Stop'
        }
        
        if ($Body) {
            $params['Body'] = $Body | ConvertTo-Json -Depth 10
        }
        
        $response = Invoke-WebRequest @params
        
        if ($response.StatusCode -eq $ExpectedStatus) {
            Write-Host "  ✅ $Name" -ForegroundColor Green
            $passedTests++
            return $true
        } else {
            Write-Host "  ❌ $Name (Expected $ExpectedStatus, got $($response.StatusCode))" -ForegroundColor Red
            $failedTests++
            return $false
        }
    }
    catch {
        $statusCode = $_.Exception.Response.StatusCode.Value__
        if ($statusCode -eq $ExpectedStatus) {
            Write-Host "  ✅ $Name (Status: $statusCode)" -ForegroundColor Green
            $passedTests++
            return $true
        } else {
            Write-Host "  ❌ $Name (Status: $statusCode)" -ForegroundColor Red
            $failedTests++
            return $false
        }
    }
}

# ============================================================================
Write-Host "`n📋 AUTHENTICATION & AUTHORIZATION TESTS" -ForegroundColor Yellow
Write-Host "─────────────────────────────────────────────────────────────────────────" -ForegroundColor Gray

Test-Endpoint -Name "Health Check" -Method "GET" -Endpoint "/health" -ExpectedStatus 200
Test-Endpoint -Name "Get Current Profile" -Method "GET" -Endpoint "/profile" -ExpectedStatus 200

# ============================================================================
Write-Host "`n👥 USER MANAGEMENT TESTS" -ForegroundColor Yellow
Write-Host "─────────────────────────────────────────────────────────────────────────" -ForegroundColor Gray

Test-Endpoint -Name "List All Users" -Method "GET" -Endpoint "/users" -ExpectedStatus 200
Test-Endpoint -Name "Get User by ID (ID: 1)" -Method "GET" -Endpoint "/users/1" -ExpectedStatus 200
Test-Endpoint -Name "Update User Profile" -Method "PUT" -Endpoint "/profile" -Body @{first_name="Mark"; last_name="Admin"} -ExpectedStatus 200
Test-Endpoint -Name "Get Super Admins List" -Method "GET" -Endpoint "/admin/super-admins" -ExpectedStatus 200

# ============================================================================
Write-Host "`n📚 COURSES & ENROLLMENT TESTS" -ForegroundColor Yellow
Write-Host "─────────────────────────────────────────────────────────────────────────" -ForegroundColor Gray

Test-Endpoint -Name "List All Courses" -Method "GET" -Endpoint "/courses" -ExpectedStatus 200
Test-Endpoint -Name "List All Students" -Method "GET" -Endpoint "/students" -ExpectedStatus 200
Test-Endpoint -Name "Admin Dashboard Stats" -Method "GET" -Endpoint "/admin/dashboard" -ExpectedStatus 200
Test-Endpoint -Name "Get All Enrollments" -Method "GET" -Endpoint "/admin/enrollments" -ExpectedStatus 200

# ============================================================================
Write-Host "`n📊 GRADES & TRANSCRIPTS TESTS" -ForegroundColor Yellow
Write-Host "─────────────────────────────────────────────────────────────────────────" -ForegroundColor Gray

Test-Endpoint -Name "Get Student Grades (Student 1)" -Method "GET" -Endpoint "/grades/by-student/1" -ExpectedStatus 200
Test-Endpoint -Name "Get Course Grades (Course 1)" -Method "GET" -Endpoint "/grades/by-course/1" -ExpectedStatus 200
Test-Endpoint -Name "Get Student Transcript (Student 1)" -Method "GET" -Endpoint "/transcripts/student/1" -ExpectedStatus 200
Test-Endpoint -Name "Get Grade Transcript List" -Method "GET" -Endpoint "/transcripts/latest/1" -ExpectedStatus 200

# ============================================================================
Write-Host "`n📍 ATTENDANCE TESTS" -ForegroundColor Yellow
Write-Host "─────────────────────────────────────────────────────────────────────────" -ForegroundColor Gray

Test-Endpoint -Name "Get Student Attendance (Student 1)" -Method "GET" -Endpoint "/attendance/by-student/1" -ExpectedStatus 200
Test-Endpoint -Name "Get Course Attendance (Course 1)" -Method "GET" -Endpoint "/attendance/by-course/1" -ExpectedStatus 200
Test-Endpoint -Name "Get Attendance Stats" -Method "GET" -Endpoint "/attendance/stats/1/1" -ExpectedStatus 200

# ============================================================================
Write-Host "`n📝 ASSIGNMENTS TESTS" -ForegroundColor Yellow
Write-Host "─────────────────────────────────────────────────────────────────────────" -ForegroundColor Gray

Test-Endpoint -Name "Get Assignments by Course (Course 1)" -Method "GET" -Endpoint "/assignments/course/1" -ExpectedStatus 200

# ============================================================================
Write-Host "`n📢 ANNOUNCEMENTS TESTS" -ForegroundColor Yellow
Write-Host "─────────────────────────────────────────────────────────────────────────" -ForegroundColor Gray

Test-Endpoint -Name "List All Announcements" -Method "GET" -Endpoint "/announcements" -ExpectedStatus 200
Test-Endpoint -Name "Get Active Announcements" -Method "GET" -Endpoint "/announcements/active" -ExpectedStatus 200
Test-Endpoint -Name "Create Announcement" -Method "POST" -Endpoint "/announcements" -Body @{title="Test"; content="Test announcement"; audience="all"; priority="normal"} -ExpectedStatus 201

# ============================================================================
Write-Host "`n🔔 NOTIFICATIONS TESTS" -ForegroundColor Yellow
Write-Host "─────────────────────────────────────────────────────────────────────────" -ForegroundColor Gray

Test-Endpoint -Name "Get My Notifications" -Method "GET" -Endpoint "/notifications" -ExpectedStatus 200
Test-Endpoint -Name "Get Unread Notifications" -Method "GET" -Endpoint "/notifications/unread" -ExpectedStatus 200

# ============================================================================
Write-Host "`n💬 MESSAGES TESTS" -ForegroundColor Yellow
Write-Host "─────────────────────────────────────────────────────────────────────────" -ForegroundColor Gray

Test-Endpoint -Name "Get Message Inbox" -Method "GET" -Endpoint "/messages/inbox" -ExpectedStatus 200
Test-Endpoint -Name "Get Unread Messages Count" -Method "GET" -Endpoint "/messages/unread" -ExpectedStatus 200

# ============================================================================
Write-Host "`n💰 PAYMENTS TESTS" -ForegroundColor Yellow
Write-Host "─────────────────────────────────────────────────────────────────────────" -ForegroundColor Gray

Test-Endpoint -Name "Get All Payments" -Method "GET" -Endpoint "/payments" -ExpectedStatus 200

# ============================================================================
Write-Host "`n📅 TIMETABLE/SCHEDULE TESTS" -ForegroundColor Yellow
Write-Host "─────────────────────────────────────────────────────────────────────────" -ForegroundColor Gray

Test-Endpoint -Name "Get All Timetables" -Method "GET" -Endpoint "/timetable" -ExpectedStatus 200
Test-Endpoint -Name "Get Timetable by Course (Course 1)" -Method "GET" -Endpoint "/timetable/course/1" -ExpectedStatus 200

# ============================================================================
Write-Host "`n👨‍🏫 TEACHER MANAGEMENT TESTS" -ForegroundColor Yellow
Write-Host "─────────────────────────────────────────────────────────────────────────" -ForegroundColor Gray

Test-Endpoint -Name "Get All Teachers" -Method "GET" -Endpoint "/admin/teachers" -ExpectedStatus 200

# ============================================================================
Write-Host "`n📊 ADVANCED SEARCH TESTS" -ForegroundColor Yellow
Write-Host "─────────────────────────────────────────────────────────────────────────" -ForegroundColor Gray

Test-Endpoint -Name "Search Announcements" -Method "GET" -Endpoint "/search/announcements?q=test" -ExpectedStatus 200
Test-Endpoint -Name "Search Students" -Method "GET" -Endpoint "/search/students?q=john" -ExpectedStatus 200

# ============================================================================
Write-Host "`n💾 EXPORT TESTS" -ForegroundColor Yellow
Write-Host "─────────────────────────────────────────────────────────────────────────" -ForegroundColor Gray

Test-Endpoint -Name "Export Payments CSV" -Method "GET" -Endpoint "/export/payments" -ExpectedStatus 200
Test-Endpoint -Name "Export Grades CSV" -Method "GET" -Endpoint "/export/grades" -ExpectedStatus 200
Test-Endpoint -Name "Export Attendance CSV" -Method "GET" -Endpoint "/export/attendance" -ExpectedStatus 200
Test-Endpoint -Name "Export Enrollments" -Method "GET" -Endpoint "/export/enrollments" -ExpectedStatus 200

# ============================================================================
Write-Host "`n" -ForegroundColor Cyan
Write-Host "╔════════════════════════════════════════════════════════════════════════╗" -ForegroundColor Cyan
Write-Host "║                         TEST RESULTS SUMMARY                          ║" -ForegroundColor Cyan
Write-Host "╠════════════════════════════════════════════════════════════════════════╣" -ForegroundColor Cyan

Write-Host "║ Total Tests:     $($totalTests.ToString().PadRight(50)) ║" -ForegroundColor Cyan
Write-Host "║ Passed:          $($passedTests.ToString().PadRight(50)) ║" -ForegroundColor Green
Write-Host "║ Failed:          $($failedTests.ToString().PadRight(50)) ║" -ForegroundColor Red

$passPercentage = if ($totalTests -gt 0) { [math]::Round(($passedTests / $totalTests) * 100, 2) } else { 0 }
Write-Host "║ Success Rate:    $($passPercentage.ToString().PadRight(50))% ║" -ForegroundColor Cyan

Write-Host "╚════════════════════════════════════════════════════════════════════════╝" -ForegroundColor Cyan

if ($failedTests -eq 0) {
    Write-Host "`n🎉 ALL TESTS PASSED! System is fully functional." -ForegroundColor Green
} else {
    Write-Host "`n⚠️  $failedTests test(s) failed. Please review the results above." -ForegroundColor Yellow
}
