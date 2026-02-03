# Test admin API endpoints

$baseUrl = "http://localhost:8080/api"

# 1. Login
Write-Host "1. Testing LOGIN endpoint..." -ForegroundColor Cyan
$loginResp = Invoke-WebRequest -Uri "$baseUrl/auth/login" -Method POST -Headers @{"Content-Type"="application/json"} -Body '{"email":"admin@school.com","password":"admin123"}' -UseBasicParsing -ErrorAction SilentlyContinue
$loginJson = $loginResp.Content | ConvertFrom-Json
$token = $loginJson.token
Write-Host "✓ Login successful. Token obtained." -ForegroundColor Green
Write-Host ""

# 2. Test Admin Dashboard Stats
Write-Host "2. Testing ADMIN DASHBOARD STATS endpoint..." -ForegroundColor Cyan
try {
    $statsResp = Invoke-WebRequest -Uri "$baseUrl/admin/dashboard" -Method GET -Headers @{"Authorization"="Bearer $token"; "Content-Type"="application/json"} -UseBasicParsing
    $statsJson = $statsResp.Content | ConvertFrom-Json
    Write-Host "✓ Stats endpoint works!" -ForegroundColor Green
    Write-Host "Response: $($statsJson | ConvertTo-Json)"
} catch {
    Write-Host "✗ Stats endpoint failed: $($_.Exception.Message)" -ForegroundColor Red
}
Write-Host ""

# 3. Test Get All Users
Write-Host "3. Testing GET ALL USERS endpoint..." -ForegroundColor Cyan
try {
    $usersResp = Invoke-WebRequest -Uri "$baseUrl/admin/users" -Method GET -Headers @{"Authorization"="Bearer $token"; "Content-Type"="application/json"} -UseBasicParsing
    $usersJson = $usersResp.Content | ConvertFrom-Json
    Write-Host "✓ Get Users endpoint works!" -ForegroundColor Green
    Write-Host "Found $($usersJson.data.count) users" -ForegroundColor Green
} catch {
    Write-Host "✗ Get Users endpoint failed: $($_.Exception.Message)" -ForegroundColor Red
}
Write-Host ""

# 4. Test Get All Courses
Write-Host "4. Testing GET ALL COURSES endpoint..." -ForegroundColor Cyan
try {
    $coursesResp = Invoke-WebRequest -Uri "$baseUrl/courses" -Method GET -Headers @{"Authorization"="Bearer $token"; "Content-Type"="application/json"} -UseBasicParsing
    $coursesJson = $coursesResp.Content | ConvertFrom-Json
    Write-Host "✓ Get Courses endpoint works!" -ForegroundColor Green
    if ($coursesJson -is [array]) {
        Write-Host "Found $($coursesJson.Count) courses" -ForegroundColor Green
    } else {
        Write-Host "Response: $($coursesJson | ConvertTo-Json)" -ForegroundColor Green
    }
} catch {
    Write-Host "✗ Get Courses endpoint failed: $($_.Exception.Message)" -ForegroundColor Red
}
Write-Host ""

# 5. Test Get All Enrollments
Write-Host "5. Testing GET ALL ENROLLMENTS endpoint..." -ForegroundColor Cyan
try {
    $enrollResp = Invoke-WebRequest -Uri "$baseUrl/admin/enrollments" -Method GET -Headers @{"Authorization"="Bearer $token"; "Content-Type"="application/json"} -UseBasicParsing
    $enrollJson = $enrollResp.Content | ConvertFrom-Json
    Write-Host "✓ Get Enrollments endpoint works!" -ForegroundColor Green
    Write-Host "Response: $($enrollJson | ConvertTo-Json)" -ForegroundColor Green
} catch {
    Write-Host "✗ Get Enrollments endpoint failed: $($_.Exception.Message)" -ForegroundColor Red
}
Write-Host ""

# 6. Test Create Course
Write-Host "6. Testing CREATE COURSE endpoint..." -ForegroundColor Cyan
try {
    $coursePayload = @{
        title = "Test Course $(Get-Date -Format 'HHmmss')"
        code = "TEST001"
        department = "Computer Science"
        description = "Test course for admin panel"
    } | ConvertTo-Json
    
    $createCourseResp = Invoke-WebRequest -Uri "$baseUrl/courses" -Method POST -Headers @{"Authorization"="Bearer $token"; "Content-Type"="application/json"} -Body $coursePayload -UseBasicParsing
    $courseJson = $createCourseResp.Content | ConvertFrom-Json
    Write-Host "✓ Create Course endpoint works!" -ForegroundColor Green
    Write-Host "Created course ID: $($courseJson.course.id)" -ForegroundColor Green
} catch {
    Write-Host "✗ Create Course endpoint failed: $($_.Exception.Message)" -ForegroundColor Red
}
Write-Host ""

Write-Host "=" -ForegroundColor Yellow | % { $_ * 50 }
Write-Host "Admin API Testing Complete!" -ForegroundColor Cyan
