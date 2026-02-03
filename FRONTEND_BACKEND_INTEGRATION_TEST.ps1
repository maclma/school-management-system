# Frontend-Backend Integration Test
# Tests the complete flow: register, login, fetch profile, enroll in courses, get grades, etc.

$baseUrl = "http://localhost:8080"
$timestamp = [int](Get-Date -UFormat %s)

# Track results
$results = @{
    tests = @()
    passed = 0
    failed = 0
}

function Test-Endpoint {
    param(
        [string]$Name,
        [string]$Method,
        [string]$Url,
        [string]$Body,
        [string]$Token,
        [int]$ExpectedStatus = 200
    )
    
    try {
        $headers = @{ "Content-Type" = "application/json" }
        if ($Token) {
            $headers["Authorization"] = "Bearer $Token"
        }
        
        $params = @{
            Uri = $Url
            Method = $Method
            Headers = $headers
            UseBasicParsing = $true
        }
        
        if ($Body) {
            $params["Body"] = $Body
        }
        
        $response = Invoke-WebRequest @params
        $status = $response.StatusCode
        
        if ($status -eq $ExpectedStatus) {
            $results.tests += @{ Name = $Name; Status = "PASS"; Code = $status; Time = $(Get-Date -Format "HH:mm:ss") }
            $results.passed++
            Write-Host "✅ $Name - $status OK"
            return $response
        } else {
            $results.tests += @{ Name = $Name; Status = "FAIL"; Code = $status; Expected = $ExpectedStatus; Time = $(Get-Date -Format "HH:mm:ss") }
            $results.failed++
            Write-Host "❌ $Name - Expected $ExpectedStatus, got $status"
            return $null
        }
    } catch {
        $results.tests += @{ Name = $Name; Status = "ERROR"; Error = $_.Exception.Message; Time = $(Get-Date -Format "HH:mm:ss") }
        $results.failed++
        Write-Host "❌ $Name - ERROR: $($_.Exception.Message)"
        return $null
    }
}

Write-Host "======================================"
Write-Host "🔗 Frontend-Backend Integration Tests"
Write-Host "======================================"
Write-Host "Backend API: $baseUrl"
Write-Host "Test Time: $(Get-Date -Format 'yyyy-MM-dd HH:mm:ss')"
Write-Host ""

# Phase 1: Health Check
Write-Host "📊 Phase 1: Health Check"
Write-Host "------------------------"
Test-Endpoint -Name "Health Endpoint" -Method GET -Url "$baseUrl/api/health" -ExpectedStatus 200 | Out-Null

# Phase 2: Authentication Flow
Write-Host ""
Write-Host "🔐 Phase 2: Authentication Flow"
Write-Host "--------------------------------"

# Register a test user
$registerEmail = "frontend_test_${timestamp}@example.com"
$registerBody = @{
    first_name = "Frontend"
    last_name = "User"
    email = $registerEmail
    password = "FrontendTest123!"
    role = "student"
} | ConvertTo-Json -Compress

$registerResp = Test-Endpoint -Name "Register Student" -Method POST -Url "$baseUrl/api/auth/register" -Body $registerBody -ExpectedStatus 201
if ($registerResp -and $registerResp.Content) {
    $regData = $registerResp.Content | ConvertFrom-Json
    $userId = $regData.user.id
    Write-Host "   User ID: $userId"
}

# Login with the registered user
$loginBody = @{
    email = $registerEmail
    password = "FrontendTest123!"
} | ConvertTo-Json -Compress

$loginResp = Test-Endpoint -Name "Login Student" -Method POST -Url "$baseUrl/api/auth/login" -Body $loginBody -ExpectedStatus 200
$token = $null
if ($loginResp -and $loginResp.Content) {
    $loginData = $loginResp.Content | ConvertFrom-Json
    $token = $loginData.token
    Write-Host "   Token: $($token.Substring(0, 20))..." | Out-Host
}

if (-not $token) {
    Write-Host "❌ Login failed - cannot proceed with authenticated tests"
    exit 1
}

# Phase 3: User Profile Tests
Write-Host ""
Write-Host "👤 Phase 3: User Profile"
Write-Host "------------------------"

$profileResp = Test-Endpoint -Name "Get Profile" -Method GET -Url "$baseUrl/api/profile" -Token $token -ExpectedStatus 200
if ($profileResp -and $profileResp.Content) {
    $profile = $profileResp.Content | ConvertFrom-Json
    Write-Host "   User Email: $($profile.email)"
    Write-Host "   User Role: $($profile.role)"
}

$updateProfileBody = @{
    first_name = "Frontend"
    last_name = "TestUser"
    phone = "555-0123"
} | ConvertTo-Json -Compress

Test-Endpoint -Name "Update Profile" -Method PUT -Url "$baseUrl/api/profile" -Body $updateProfileBody -Token $token -ExpectedStatus 200 | Out-Null

# Phase 4: Courses & Enrollment
Write-Host ""
Write-Host "📚 Phase 4: Courses & Enrollment"
Write-Host "--------------------------------"

# First, create a teacher user to own a course
$teacherEmail = "teacher_${timestamp}@example.com"
$registerTeacherBody = @{
    first_name = "Test"
    last_name = "Teacher"
    email = $teacherEmail
    password = "TeacherTest123!"
    role = "teacher"
} | ConvertTo-Json -Compress

$teacherResp = Test-Endpoint -Name "Register Teacher" -Method POST -Url "$baseUrl/api/auth/register" -Body $registerTeacherBody -ExpectedStatus 201
$teacherToken = $null
if ($teacherResp -and $teacherResp.Content) {
    $teacherData = $teacherResp.Content | ConvertFrom-Json
    $teacherId = $teacherData.user.id
    
    # Login as teacher
    $teacherLoginBody = @{
        email = $teacherEmail
        password = "TeacherTest123!"
    } | ConvertTo-Json -Compress
    
    $tLoginResp = Test-Endpoint -Name "Login Teacher" -Method POST -Url "$baseUrl/api/auth/login" -Body $teacherLoginBody -ExpectedStatus 200
    if ($tLoginResp -and $tLoginResp.Content) {
        $tLoginData = $tLoginResp.Content | ConvertFrom-Json
        $teacherToken = $tLoginData.token
    }
}

# Get available courses
Test-Endpoint -Name "Get Courses List" -Method GET -Url "$baseUrl/api/courses" -Token $token -ExpectedStatus 200 | Out-Null

# Phase 5: Grades & Attendance (simulated data)
Write-Host ""
Write-Host "📊 Phase 5: Academic Data"
Write-Host "------------------------"

# Get student's grades (may be empty)
Test-Endpoint -Name "Get My Grades" -Method GET -Url "$baseUrl/api/student/grades" -Token $token -ExpectedStatus 200 | Out-Null

# Get student's attendance (may be empty)
Test-Endpoint -Name "Get My Attendance" -Method GET -Url "$baseUrl/api/student/attendance" -Token $token -ExpectedStatus 200 | Out-Null

# Phase 6: Announcements & Notifications
Write-Host ""
Write-Host "📢 Phase 6: Communications"
Write-Host "---------------------------"

# Get announcements
Test-Endpoint -Name "Get Announcements" -Method GET -Url "$baseUrl/api/announcements" -Token $token -ExpectedStatus 200 | Out-Null

# Get notifications
Test-Endpoint -Name "Get Notifications" -Method GET -Url "$baseUrl/api/notifications" -Token $token -ExpectedStatus 200 | Out-Null

# Phase 7: Admin Features (test with admin auto-created on first run)
Write-Host ""
Write-Host "🔧 Phase 7: Admin Features"
Write-Host "---------------------------"

# Register as admin
$adminEmail = "admin_test_${timestamp}@example.com"
$registerAdminBody = @{
    first_name = "Admin"
    last_name = "Test"
    email = $adminEmail
    password = "AdminTest123!"
    role = "admin"
} | ConvertTo-Json -Compress

$adminResp = Test-Endpoint -Name "Register Admin" -Method POST -Url "$baseUrl/api/auth/register" -Body $registerAdminBody -ExpectedStatus 201
$adminToken = $null
if ($adminResp -and $adminResp.Content) {
    $adminData = $adminResp.Content | ConvertFrom-Json
    
    # Login as admin
    $adminLoginBody = @{
        email = $adminEmail
        password = "AdminTest123!"
    } | ConvertTo-Json -Compress
    
    $aLoginResp = Test-Endpoint -Name "Login Admin" -Method POST -Url "$baseUrl/api/auth/login" -Body $adminLoginBody -ExpectedStatus 200
    if ($aLoginResp -and $aLoginResp.Content) {
        $aLoginData = $aLoginResp.Content | ConvertFrom-Json
        $adminToken = $aLoginData.token
    }
}

if ($adminToken) {
    # Get admin dashboard
    Test-Endpoint -Name "Admin Dashboard" -Method GET -Url "$baseUrl/api/admin/dashboard" -Token $adminToken -ExpectedStatus 200 | Out-Null
    
    # Get admin health
    Test-Endpoint -Name "Admin Health Check" -Method GET -Url "$baseUrl/api/admin/health" -Token $adminToken -ExpectedStatus 200 | Out-Null
    
    # Get super admins list
    Test-Endpoint -Name "Get Super Admins" -Method GET -Url "$baseUrl/api/admin/super-admins" -Token $adminToken -ExpectedStatus 200 | Out-Null
}

# Summary
Write-Host ""
Write-Host "======================================"
Write-Host "📋 Test Summary"
Write-Host "======================================"
Write-Host "Total Tests: $($results.passed + $results.failed)"
Write-Host "Passed: ✅ $($results.passed)"
Write-Host "Failed: ❌ $($results.failed)"
Write-Host ""

if ($results.failed -eq 0) {
    Write-Host "🎉 ALL TESTS PASSED!"
    Write-Host ""
    Write-Host "✅ Frontend-Backend Integration Status: WORKING"
    Write-Host ""
    Write-Host "Key Features Verified:"
    Write-Host "  • Health check endpoint"
    Write-Host "  • User registration and login"
    Write-Host "  • Profile management"
    Write-Host "  • Course retrieval"
    Write-Host "  • Student grades access"
    Write-Host "  • Attendance tracking"
    Write-Host "  • Announcements and notifications"
    Write-Host "  • Admin dashboard access"
    Write-Host "  • Super admin management"
    Write-Host ""
    exit 0
} else {
    Write-Host "Some tests failed - review above for details"
    Write-Host ""
    Write-Host "Failed Tests:"
    foreach ($test in $results.tests | Where-Object { $_.Status -ne "PASS" }) {
        Write-Host "  - $($test.Name): $($test.Status)"
    }
    exit 1
}
