# Comprehensive Test Suite for School Management System

Write-Host "========================================" -ForegroundColor Cyan
Write-Host "COMPREHENSIVE FUNCTIONALITY TESTS" -ForegroundColor Cyan
Write-Host "========================================" -ForegroundColor Cyan

# Admin token
$token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImFkbWluQHNjaG9vbC5jb20iLCJleHAiOjE3NzAxMzczMzgsImlhdCI6MTc3MDA1MDkzOCwicm9sZSI6ImFkbWluIiwidXNlcl9pZCI6MX0.QV4q5ahtUAJdMbEWOj3hYcmF7yx-vfjc_pxpXUFB2Wo"
$headers = @{
    "Authorization" = "Bearer $token"
    "Content-Type" = "application/json"
}
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
            Write-Host "  PASS: $Name" -ForegroundColor Green
            $passedTests++
        } else {
            Write-Host "  FAIL: $Name (Expected $ExpectedStatus, got $($response.StatusCode))" -ForegroundColor Red
            $failedTests++
        }
    }
    catch {
        Write-Host "  FAIL: $Name - $($_.Exception.Message)" -ForegroundColor Red
        $failedTests++
    }
}

# ============================================
Write-Host "`n[AUTHENTICATION]" -ForegroundColor Yellow
Test-Endpoint "Health Check" "GET" "/health"

# ============================================
Write-Host "`n[USER MANAGEMENT]" -ForegroundColor Yellow
Test-Endpoint "Get All Users" "GET" "/users"
Test-Endpoint "Get Admin Users" "GET" "/admin/users"
Test-Endpoint "Admin Health" "GET" "/admin/health"

# ============================================
Write-Host "`n[COURSES & ENROLLMENT]" -ForegroundColor Yellow
Test-Endpoint "Get All Courses" "GET" "/courses"
Test-Endpoint "Get All Students" "GET" "/students"

# ============================================
Write-Host "`n[GRADES]" -ForegroundColor Yellow
Test-Endpoint "Get Grades (all)" "GET" "/grades"

# ============================================
Write-Host "`n[ATTENDANCE]" -ForegroundColor Yellow
Test-Endpoint "Get Attendance (all)" "GET" "/attendance"

# ============================================
Write-Host "`n[ANNOUNCEMENTS]" -ForegroundColor Yellow
Test-Endpoint "Get All Announcements" "GET" "/announcements"
Test-Endpoint "Get Active Announcements" "GET" "/announcements/active"

# ============================================
Write-Host "`n[NOTIFICATIONS]" -ForegroundColor Yellow
Test-Endpoint "Get My Notifications" "GET" "/notifications"
Test-Endpoint "Get Unread Notifications" "GET" "/notifications/unread"

# ============================================
Write-Host "`n[MESSAGES]" -ForegroundColor Yellow
Test-Endpoint "Get Messages Inbox" "GET" "/messages/inbox"
Test-Endpoint "Get Unread Messages Count" "GET" "/messages/unread"

# ============================================
Write-Host "`n[PAYMENTS]" -ForegroundColor Yellow
Test-Endpoint "Get All Payments" "GET" "/payments"

# ============================================
Write-Host "`n[TRANSCRIPTS]" -ForegroundColor Yellow
Test-Endpoint "Get Transcript GPA (Student 1)" "GET" "/transcripts/gpa/1"

# ============================================
Write-Host "`n[TIMETABLE]" -ForegroundColor Yellow
Test-Endpoint "Get All Timetables" "GET" "/timetable"

# ============================================
Write-Host "`n[ASSIGNMENTS]" -ForegroundColor Yellow
Test-Endpoint "Get All Assignments" "GET" "/assignments"

# ============================================
Write-Host "`n[SEARCH]" -ForegroundColor Yellow
Test-Endpoint "Search Announcements" "GET" "/search/announcements"
Test-Endpoint "Search Payments" "GET" "/search/payments"
Test-Endpoint "Search Students" "GET" "/search/students"

# ============================================
Write-Host "`n[EXPORTS]" -ForegroundColor Yellow
Test-Endpoint "Export Payments CSV" "GET" "/export/payments"
Test-Endpoint "Export Grades CSV" "GET" "/export/grades"
Test-Endpoint "Export Attendance CSV" "GET" "/export/attendance"

# ============================================
Write-Host "`n========================================" -ForegroundColor Cyan
Write-Host "TEST RESULTS SUMMARY" -ForegroundColor Cyan
Write-Host "========================================" -ForegroundColor Cyan
Write-Host "Total Tests: $totalTests" -ForegroundColor White
Write-Host "Passed: $passedTests" -ForegroundColor Green
Write-Host "Failed: $failedTests" -ForegroundColor Red

if ($failedTests -eq 0) {
    Write-Host "`nALL TESTS PASSED!" -ForegroundColor Green
} else {
    Write-Host "`n$failedTests test(s) failed" -ForegroundColor Yellow
}
