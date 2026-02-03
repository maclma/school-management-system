# Comprehensive Test Suite for School Management System - Fixed Endpoints

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
        Write-Host "  FAIL: $Name - $($_.Exception.Message.Split("`n")[0])" -ForegroundColor Red
        $failedTests++
    }
}

# ============================================
Write-Host "`n[CORE FUNCTIONALITY]" -ForegroundColor Yellow
Test-Endpoint "System Health Check" "GET" "/health"
Test-Endpoint "Admin Dashboard" "GET" "/admin/health"

# ============================================
Write-Host "`n[USER MANAGEMENT]" -ForegroundColor Yellow
Test-Endpoint "Get All Users" "GET" "/users"
Test-Endpoint "Get Admin View - All Users" "GET" "/admin/users"
Test-Endpoint "Get Profile" "GET" "/profile"

# ============================================
Write-Host "`n[COURSES & ENROLLMENT]" -ForegroundColor Yellow
Test-Endpoint "Get All Courses" "GET" "/courses"
Test-Endpoint "Get All Students" "GET" "/students"
Test-Endpoint "Get All Enrollments (Admin)" "GET" "/admin/enrollments"

# ============================================
Write-Host "`n[GRADES & PERFORMANCE]" -ForegroundColor Yellow
Test-Endpoint "Search Grades by Range" "GET" "/search/grades"

# ============================================
Write-Host "`n[ANNOUNCEMENTS]" -ForegroundColor Yellow
Test-Endpoint "Get All Announcements" "GET" "/announcements"
Test-Endpoint "Get Active Announcements" "GET" "/announcements/active"
Test-Endpoint "Search Announcements" "GET" "/search/announcements"

# ============================================
Write-Host "`n[NOTIFICATIONS]" -ForegroundColor Yellow
Test-Endpoint "Get My Notifications" "GET" "/notifications"
Test-Endpoint "Get Unread Notifications" "GET" "/notifications/unread"

# ============================================
Write-Host "`n[MESSAGING]" -ForegroundColor Yellow
Test-Endpoint "Get Messages Inbox" "GET" "/messages/inbox"
Test-Endpoint "Get Unread Messages Count" "GET" "/messages/unread"

# ============================================
Write-Host "`n[PAYMENTS & FINANCE]" -ForegroundColor Yellow
Test-Endpoint "Get All Payments" "GET" "/payments"
Test-Endpoint "Search Payments" "GET" "/search/payments"
Test-Endpoint "Search Overdue Payments" "GET" "/search/overdue-payments"

# ============================================
Write-Host "`n[TIMETABLE & SCHEDULE]" -ForegroundColor Yellow
Test-Endpoint "Get All Timetables" "GET" "/timetable"

# ============================================
Write-Host "`n[TEACHER MANAGEMENT]" -ForegroundColor Yellow
Test-Endpoint "Get All Teachers (Admin)" "GET" "/admin/teachers"

# ============================================
Write-Host "`n[EXPORT & REPORTING]" -ForegroundColor Yellow
Test-Endpoint "Export Payments CSV" "GET" "/export/payments"
Test-Endpoint "Export Grades CSV" "GET" "/export/grades"
Test-Endpoint "Export Attendance CSV" "GET" "/export/attendance"
Test-Endpoint "Export Enrollments CSV" "GET" "/export/enrollments"

# ============================================
Write-Host "`n[SYSTEM SETTINGS]" -ForegroundColor Yellow
Test-Endpoint "Get System Settings" "GET" "/admin/settings"

# ============================================
Write-Host "`n========================================" -ForegroundColor Cyan
Write-Host "TEST RESULTS SUMMARY" -ForegroundColor Cyan
Write-Host "========================================" -ForegroundColor Cyan
Write-Host "Total Tests Run: $totalTests" -ForegroundColor White
Write-Host "Passed: $passedTests" -ForegroundColor Green
Write-Host "Failed: $failedTests" -ForegroundColor Red

if ($totalTests -gt 0) {
    $percentage = [math]::Round(($passedTests / $totalTests) * 100, 2)
    Write-Host "Success Rate: $percentage%" -ForegroundColor Cyan
}

Write-Host "`n========================================" -ForegroundColor Cyan

if ($failedTests -eq 0) {
    Write-Host "ALL TESTS PASSED! System is fully functional." -ForegroundColor Green
} else {
    Write-Host "$failedTests test(s) failed" -ForegroundColor Yellow
}
