# API Quick Start Guide

This document shows how to interact with the School Management System API locally.

## Prerequisites

- Go 1.16+ installed
- Port 8080 (or 8081) available

## Running the Server

### Option 1: Using the local run script (Windows)
```powershell
.\scripts\run_local.ps1
```

### Option 2: Using environment variables (Windows PowerShell)
```powershell
$env:DB_DRIVER='sqlite'
$env:DB_PATH='school.db'
$env:JWT_SECRET='changeme'
go run ./cmd/server
```

### Option 3: Custom port
```powershell
$env:SERVER_PORT='8081'
go run ./cmd/server
```

## API Endpoints

All endpoints use JSON for request/response. Base URL: `http://localhost:8080`

### Authentication

#### Register User
```bash
curl -X POST http://localhost:8080/api/auth/register \
  -H "Content-Type: application/json" \
  -d '{
    "first_name": "John",
    "last_name": "Doe",
    "email": "john@example.com",
    "password": "password123",
    "role": "student"
  }'
```

**Roles**: `admin`, `teacher`, `student`, `parent`

**Response (201)**:
```json
{
  "message": "User registered successfully",
  "user": {
    "id": 1,
    "email": "john@example.com",
    "role": "student"
  }
}
```

#### Login
```bash
curl -X POST http://localhost:8080/api/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "john@example.com",
    "password": "password123"
  }'
```

**Response (200)**:
```json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
  "user": {
    "id": 1,
    "email": "john@example.com",
    "first_name": "John",
    "last_name": "Doe",
    "role": "student"
  }
}
```

### User Profile (Protected Routes)

All protected routes require the `Authorization: Bearer <token>` header.

#### Get Profile
```bash
curl -X GET http://localhost:8080/api/profile \
  -H "Authorization: Bearer <your-token-here>"
```

**Response (200)**:
```json
{
  "id": 1,
  "first_name": "John",
  "last_name": "Doe",
  "email": "john@example.com",
  "role": "student"
}
```

#### Update Profile
```bash
curl -X PUT http://localhost:8080/api/profile \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer <your-token-here>" \
  -d '{
    "first_name": "Jane",
    "last_name": "Smith",
    "phone": "+1234567890"
  }'
```

**Response (200)**:
```json
{
  "message": "Profile updated successfully"
}
```

## Testing with Postman

1. Import `api/postman_collection.json` into Postman
2. Import `api/postman_environment.json` as the environment
3. Update the `jwt` variable after login with your token
4. Run requests using the available endpoints

## Testing with PowerShell

Run the included demo script:
```powershell
.\scripts\demo.ps1
```

This script automatically:
1. Starts the server on port 8081
2. Registers a new user
3. Logs in
4. Retrieves the user profile
5. Updates the profile
6. Stops the server

## Running Tests

Run all integration tests:
```bash
go test ./tests -v
```

Tests use an in-memory SQLite database and are fully isolated. Each test clears the database before running.

## Troubleshooting

### Port already in use
```powershell
# List processes using port 8080
Get-NetTCPConnection -LocalPort 8080 | Select-Object -ExpandProperty OwningProcess | ForEach-Object { Get-Process -Id $_ }

# Kill the process
Stop-Process -Id <PID> -Force

# Or use a different port
$env:SERVER_PORT='8081'
go run ./cmd/server
```

### Database issues
```powershell
# Use SQLite (default, requires no setup)
$env:DB_DRIVER='sqlite'
$env:DB_PATH='school.db'

# Remove old database file
Remove-Item -Path school.db -Force -ErrorAction SilentlyContinue
go run ./cmd/server
```

### JWT Secret
For local development, the default secret is `changeme`. For production, set:
```powershell
$env:JWT_SECRET='your-secure-secret-key-here'
```

## Environment Variables

| Variable | Default | Purpose |
|----------|---------|---------|
| `DB_DRIVER` | `sqlite` | Database driver (`sqlite` or `postgres`) |
| `DB_PATH` | `school.db` | SQLite file path (ignored for postgres) |
| `DB_HOST` | `localhost` | Postgres host |
| `DB_PORT` | `5432` | Postgres port |
| `DB_USER` | `postgres` | Postgres user |
| `DB_PASSWORD` | empty | Postgres password |
| `DB_NAME` | `school_db` | Postgres database name |
| `SERVER_PORT` | `8080` | HTTP server port |
| `JWT_SECRET` | `changeme` | JWT signing secret |
| `JWT_EXPIRY` | `24` | Token expiry in hours |

## Next Steps

- Implement course management endpoints
- Add student/teacher enrollment flows
- Build admin dashboard routes
- Add file upload for profile images
