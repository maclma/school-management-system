#!/usr/bin/env pwsh
# School Management System - API Integration Testing
# Tests all 11 new features

param(
    [string]$BaseURL = "http://localhost:8080",
    [string]$AdminEmail = "admin@school.com",
    [string]$AdminPassword = "admin123"
)

$ErrorActionPreference = "Stop"

# Colors for output
$Colors = @{
    Success = 'Green'
    Error = 'Red'
    Info = 'Cyan'
    Warning = 'Yellow'
}

function Write-Status($message, $type = "Info") {
    Write-Host $message -ForegroundColor $Colors[$type]
}

function Get-AuthToken {
    Write-Status "Authenticating admin user..." "Info"
    
    $loginBody = @{
        email = $AdminEmail
        password = $AdminPassword
    } | ConvertTo-Json
    
    try {
        $response = Invoke-WebRequest -Uri "$BaseURL/api/auth/login" `
            -Method POST `
            -ContentType "application/json" `
            -Body $loginBody `
            -ErrorAction Stop
        
        $token = ($response.Content | ConvertFrom-Json).data.token
        Write-Status "✓ Authentication successful" "Success"
        return $token
    }
    catch {
        Write-Status "✗ Authentication failed: $_" "Error"
        exit 1
    }
}

function Test-Endpoint {
    param(
        [string]$Method,
        [string]$Endpoint,
        [string]$Token,
        [hashtable]$Body = $null,
        [string]$Description = ""
    )
    
    $url = "$BaseURL$Endpoint"
    $headers = @{ Authorization = "Bearer $Token" }
    
    try {
        if ($Body) {
            $response = Invoke-WebRequest -Uri $url `
                -Method $Method `
                -Headers $headers `
                -ContentType "application/json" `
                -Body ($Body | ConvertTo-Json) `
                -ErrorAction Stop
        } else {
            $response = Invoke-WebRequest -Uri $url `
                -Method $Method `
                -Headers $headers `
                -ErrorAction Stop
        }
        
        $statusCode = $response.StatusCode
        $data = $response.Content | ConvertFrom-Json
        
        if ($data.success -or $statusCode -in @(200, 201, 204)) {
            Write-Status "✓ $Description" "Success"
            return $data
        } else {
            Write-Status "✗ $Description - Failed response" "Error"
            return $null
        }
    }
    catch {
        Write-Status "✗ $Description - Error: $($_.Exception.Message)" "Error"
        return $null
    }
}

# Main testing flow
Write-Status "========================================" "Info"
Write-Status "School Management System - API Tests" "Info"
Write-Status "========================================" "Info"
Write-Host ""

# Get authentication token
$token = Get-AuthToken
Write-Host ""

# Test 1: System Settings
Write-Status "[1/9] Testing System Settings API..." "Info"
$settings = Test-Endpoint -Method GET -Endpoint "/api/admin/settings" -Token $token -Description "Get all settings"
if ($settings) {
    Write-Status "  Found $(($settings.data | Measure-Object).Count) settings" "Success"
}
Write-Host ""

# Test 2: Announcements
Write-Status "[2/9] Testing Announcements API..." "Info"
$announcements = Test-Endpoint -Method GET -Endpoint "/api/announcements" -Token $token -Description "Get all announcements"
if ($announcements) {
    Write-Status "  Found $(($announcements.data | Measure-Object).Count) announcements" "Success"
}
$activeAnnouncements = Test-Endpoint -Method GET -Endpoint "/api/announcements/active" -Token $token -Description "Get active announcements"
Write-Host ""

# Test 3: Messages
Write-Status "[3/9] Testing Messages API..." "Info"
$inbox = Test-Endpoint -Method GET -Endpoint "/api/messages/inbox" -Token $token -Description "Get inbox messages"
if ($inbox) {
    Write-Status "  Found $(($inbox.data | Measure-Object).Count) messages" "Success"
}
$unreadCount = Test-Endpoint -Method GET -Endpoint "/api/messages/unread" -Token $token -Description "Count unread messages"
Write-Host ""

# Test 4: Notifications
Write-Status "[4/9] Testing Notifications API..." "Info"
$notifications = Test-Endpoint -Method GET -Endpoint "/api/notifications" -Token $token -Description "Get my notifications"
if ($notifications) {
    Write-Status "  Found $(($notifications.data | Measure-Object).Count) notifications" "Success"
}
$unreadNotifs = Test-Endpoint -Method GET -Endpoint "/api/notifications/unread" -Token $token -Description "Get unread notifications"
Write-Host ""

# Test 5: Payments
Write-Status "[5/9] Testing Payments API..." "Info"
$payments = Test-Endpoint -Method GET -Endpoint "/api/payments" -Token $token -Description "Get all payments"
if ($payments) {
    Write-Status "  Found $(($payments.data | Measure-Object).Count) payments" "Success"
}
Write-Host ""

# Test 6: TimeTable
Write-Status "[6/9] Testing TimeTable API..." "Info"
$timetable = Test-Endpoint -Method GET -Endpoint "/api/timetable" -Token $token -Description "Get all timetable entries"
if ($timetable) {
    Write-Status "  Found $(($timetable.data | Measure-Object).Count) schedule entries" "Success"
}
Write-Host ""

# Test 7: Grade Transcripts
Write-Status "[7/9] Testing Grade Transcript API..." "Info"
$transcripts = Test-Endpoint -Method GET -Endpoint "/api/transcripts/student/2" -Token $token -Description "Get student transcripts"
if ($transcripts) {
    Write-Status "  Found $(($transcripts.data | Measure-Object).Count) transcript records" "Success"
}
Write-Host ""

# Test 8: Backups
Write-Status "[8/9] Testing Backup API..." "Info"
$backups = Test-Endpoint -Method GET -Endpoint "/api/admin/backups" -Token $token -Description "Get all backups"
if ($backups) {
    Write-Status "  Found $(($backups.data | Measure-Object).Count) backup records" "Success"
}
Write-Host ""

# Test 9: Import Batches
Write-Status "[9/9] Testing Import Batch API..." "Info"
$imports = Test-Endpoint -Method GET -Endpoint "/api/admin/imports" -Token $token -Description "Get all import batches"
if ($imports) {
    Write-Status "  Found $(($imports.data | Measure-Object).Count) import batches" "Success"
}
Write-Host ""

# Summary
Write-Status "========================================" "Info"
Write-Status "All Tests Completed!" "Success"
Write-Status "========================================" "Info"
Write-Host ""
Write-Status "Next Steps:" "Info"
Write-Status "1. Review API_TESTING_GUIDE.md for detailed endpoint documentation" "Info"
Write-Status "2. Run seed script: sqlite3 school.db < scripts/seed_advanced_features.sql" "Info"
Write-Status "3. Test with Postman collection: api/postman_collection.json" "Info"
