# Detailed Step-by-Step Execution Guide

This guide walks you through running the School Management System API from scratch.

## Part 1: Initial Setup

### Step 1: Navigate to Project Directory
```powershell
cd C:\Users\dell\school-management-system
```

### Step 2: Verify Go is Installed
```powershell
go version
```
Expected output: `go version go1.16` or higher

### Step 3: Install Dependencies
```powershell
go mod download
go mod tidy
```

This downloads all required Go packages (Gin, GORM, JWT, SQLite driver, etc.)

---

## Part 2: Running the Server

### Option A: Quick Start (Recommended for First Time)

#### Step 1: Start the Server with Default Settings
```powershell
$env:DB_DRIVER='sqlite'
$env:DB_PATH='school.db'
$env:JWT_SECRET='changeme'
go run ./cmd/server
```

**What happens:**
- Creates/uses `school.db` SQLite database in project root
- Runs auto-migrations (creates database schema)
- Seeds an admin user (email: `admin@school.com`, password: `admin123`)
- Server starts on `http://localhost:8080`

**Expected output:**
```
FATA[2025-12-08 11:00:00] Database connected successfully
FATA[2025-12-08 11:00:00] Running database migrations...
FATA[2025-12-08 11:00:00] Database migrations completed
FATA[2025-12-08 11:00:00] Admin user created successfully
FATA[2025-12-08 11:00:00] Server starting on port 8080
```

#### Step 2: Leave Server Running
Keep this terminal open while you test the API.

---

### Option B: Using the Run Script

If port 8080 is already in use, use the included script:

#### Step 1: Open a New PowerShell Terminal
```powershell
cd C:\Users\dell\school-management-system
```

#### Step 2: Run the Script
```powershell
.\scripts\run_local.ps1
```

This automatically:
- Sets SQLite environment variables
- Checks if port 8080 is available
- Starts the server
- Displays server startup messages

---

## Part 3: Testing the API

### Open a Second Terminal for API Testing

Open a **new PowerShell terminal** (keep the server running in the first terminal).

```powershell
cd C:\Users\dell\school-management-system
```

---

## Testing Workflow

### Test 1: Register a New User

```powershell


$response = Invoke-RestMethod -Method Post 
  -Uri "http://localhost:8080/api/auth/register" `
  -ContentType 'application/json' `
  -Body $body

---
$response | ConvertTo-Json
```

**Expected output:**
```json
{
  "message": "User registered successfully",
  "user": {
    "id": 2,
    "email": "testuser_12345@example.com",
    "role": "student"
  }
}
```

**If you see errors:**
- Port not available: Change to port 8081 in server (`$env:SERVER_PORT='8081'`)
- Email already exists: Use a different email (the script generates unique ones)

---

### Test 2: Login with Registered User

```powershell
$loginBody = @{
    email = $email
    password = "password123"
} | ConvertTo-Json

$loginResponse = Invoke-RestMethod -Method Post `
  -Uri "http://localhost:8080/api/auth/login" `
  -ContentType 'application/json' `
  -Body $loginBody

$loginResponse | ConvertTo-Json
```

**Expected output:**
```json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1$email = "testuser_$(Get-Random)@example.com"
$body = @{
    first_name = "Test"
    last_name = "User"
    email = $email
    password = "password123"
    role = "student"
} | ConvertTo-Jsonc2VyX2lkIjoyLCJlbWFpbCI6InRlc3R1c2VyXzEyMzQ1QGV4YW1wbGUuY29tIiwicm9sZSI6InN0dWRlbnQiLCJleHAiOjE3NjUyNzc2MzUsImlhdCI6MTc2NTI0NzYzNX0.kWO8...",
  "user": {
    "id": 2,
    "email": "testuser_12345@example.com",
    "first_name": "Test",
    "last_name": "User",
    "role": "student"
  }
}
```

**Save the token for next steps:**
```powershell
$token = $loginResponse.token
Write-Host "Token saved: $token"
```

---

### Test 3: Get User Profile (Protected Route)

```powershell
$headers = @{
    Authorization = "Bearer $token"
}

$profile = Invoke-RestMethod -Method Get `
  -Uri "http://localhost:8080/api/profile" `
  -Headers $headers

$profile | ConvertTo-Json
```

**Expected output:**
```json
{
  "id": 2,
  "first_name": "Test",
  "last_name": "User",
  "email": "testuser_12345@example.com",
  "role": "student"
}
```

**If you get a 401 Unauthorized:**
- Token is invalid or expired
- Authorization header format is wrong: must be `"Bearer <token>"`
- Re-login to get a fresh token

---

### Test 4: Update Profile

```powershell
$updateBody = @{
    first_name = "UpdatedName"
    phone = "+1234567890"
} | ConvertTo-Json

$updateResponse = Invoke-RestMethod -Method Put `
  -Uri "http://localhost:8080/api/profile" `
  -ContentType 'application/json' `
  -Headers $headers `
  -Body $updateBody

$updateResponse | ConvertTo-Json
```

**Expected output:**
```json
{
  "message": "Profile updated successfully"
}
```

**Verify update:**
```powershell
# Get profile again to confirm changes
$profile = Invoke-RestMethod -Method Get `
  -Uri "http://localhost:8080/api/profile" `
  -Headers $headers

$profile | ConvertTo-Json
```

---

## Part 4: Running Automated Tests

### Run Integration Tests

```powershell
go test ./tests -v
```

**Expected output:**
```
=== RUN   TestRegister
--- PASS: TestRegister (0.13s)
=== RUN   TestLogin
--- PASS: TestLogin (0.30s)
=== RUN   TestGetProfile
--- PASS: TestGetProfile (0.31s)
=== RUN   TestUpdateProfile
--- PASS: TestUpdateProfile (0.35s)
=== RUN   TestUnauthorizedAccess
--- PASS: TestUnauthorizedAccess (0.00s)
PASS
ok      school-management-system/tests  6.518s
```

---

## Part 5: Using the Demo Script

Run a fully automated demo that registers, logs in, and tests all endpoints:

```powershell
.\scripts\demo.ps1
```

**What happens:**
1. Starts server on port 8081
2. Waits 3 seconds for server to initialize
3. Registers a new user
4. Logs in and extracts token
5. Gets profile
6. Updates profile
7. Stops server automatically

**Output includes:**
```
School Management System - Demo Script
======================================

Starting server on port 8081...
Server PID: 12345
Waiting for server to start...

=== 1. REGISTER USER ===
Email: demo_5432_1765194843@example.com
Response:
{
  "message": "User registered successfully",
  ...
}

=== 2. LOGIN ===
Response:
{
  "token": "eyJhbGc...",
  ...
}

=== 3. GET PROFILE ===
...

=== 4. UPDATE PROFILE ===
...

=== Demo Complete ===
```

---

## Part 6: Using Postman

### Step 1: Import Collection
1. Open Postman
2. Click **Import**
3. Select file: `api/postman_collection.json`

### Step 2: Import Environment
1. Click **Environments** (left sidebar)
2. Click **Import**
3. Select file: `api/postman_environment.json`

### Step 3: Select Environment
1. Top-right dropdown
2. Select `School Management - Local`

### Step 4: Run Requests
1. Click **Register** request
2. Click **Send**
3. Click **Login** request
4. Postman auto-extracts token into the `jwt` environment variable
5. Click **Get Profile** - uses the saved token automatically

---

## Troubleshooting

### Problem: "Port 8080 already in use"

**Solution 1: Use different port**
```powershell
$env:SERVER_PORT='8081'
go run ./cmd/server
# Then test on http://localhost:8081
```

**Solution 2: Kill process using port 8080**
```powershell
# Find the process
Get-NetTCPConnection -LocalPort 8080 | Select-Object -ExpandProperty OwningProcess | ForEach-Object { Get-Process -Id $_ }

# Kill it (replace 12345 with actual PID)
Stop-Process -Id 12345 -Force
```

---

### Problem: "Invalid credentials" on login

**Possible causes:**
1. Wrong email address
2. Wrong password
3. User doesn't exist

**Solution:**
```powershell
# Register user again with correct details
$email = "newuser_$(Get-Random)@example.com"
# Then login with that email and password
```

---

### Problem: "401 Unauthorized" on protected routes

**Causes:**
1. Missing Authorization header
2. Wrong token format (must be "Bearer <token>")
3. Token expired (default: 24 hours)

**Solution:**
```powershell
# Re-login to get fresh token
$loginBody = @{
    email = $email
    password = "password123"
} | ConvertTo-Json

$loginResponse = Invoke-RestMethod -Method Post `
  -Uri "http://localhost:8080/api/auth/login" `
  -ContentType 'application/json' `
  -Body $loginBody

$token = $loginResponse.token
```

---

### Problem: Database issues

**Reset database:**
```powershell
# Stop the server (Ctrl+C in server terminal)

# Remove old database
Remove-Item -Path school.db -Force -ErrorAction SilentlyContinue

# Restart server
$env:DB_DRIVER='sqlite'
$env:DB_PATH='school.db'
$env:JWT_SECRET='changeme'
go run ./cmd/server
```

---

## Quick Reference Commands

```powershell
# Start server
$env:DB_DRIVER='sqlite'; $env:DB_PATH='school.db'; $env:JWT_SECRET='changeme'; go run ./cmd/server

# Run tests
go test ./tests -v

# Run demo
.\scripts\demo.ps1

# Check if port is available
Get-NetTCPConnection -LocalPort 8080 -ErrorAction SilentlyContinue

# View logs
Get-Content -Path logs/app.log -Tail 50

# Build executable
go build -o school-management ./cmd/server
```

---

## Next Steps After Verification

Once the above tests pass:

1. **Course Management**: Implement endpoint to create/list/update courses
2. **Student Enrollment**: Build enrollment workflows
3. **Grade Management**: Add grade tracking and reporting
4. **Attendance Tracking**: Implement attendance logging
5. **Admin Dashboard**: Create admin-only endpoints
6. **File Uploads**: Add profile image upload support
7. **Notifications**: Implement email/SMS notifications

---

## Environment Variables Reference

Create a `.env` file in project root:

```
DB_DRIVER=sqlite
DB_PATH=school.db
SERVER_PORT=8080
JWT_SECRET=your-secure-secret
JWT_EXPIRY=24

# For Postgres (if switching from SQLite):
# DB_DRIVER=postgres
# DB_HOST=localhost
# DB_PORT=5432
# DB_USER=postgres
# DB_PASSWORD=yourpassword
# DB_NAME=school_db
```

Load automatically:
```powershell
$env | Where-Object { $_.Name -match '^(DB_|SERVER_|JWT_)' } | Format-Table Name, Value
```

---

## Support

For issues or questions:
1. Check `API_QUICKSTART.md` for endpoint documentation
2. Review test file: `tests/integration_test.go` for examples
3. Check server logs in the running terminal
4. Review code in `internal/handlers/` for implementation details
