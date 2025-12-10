# School Management System — Familiarization Guide

## Overview

This is a **Go + React web application** for managing a school with students, courses, enrollments, grades, and attendance tracking. It has role-based access control (Admin, Teacher, Student).

---

## Quick Start

**Backend is currently running on `http://localhost:8080`**

### Option 1: Test API with PowerShell (immediate exploration)
### Option 2: Run Frontend React App (visual exploration)

---

## Architecture

### Backend Stack
- **Language:** Go
- **Framework:** Gin (HTTP router)
- **Database:** SQLite (default, configured in `internal/config/config.go`)
- **ORM:** GORM
- **Auth:** JWT tokens
- **Access Control:** Middleware-based role checking (Admin, Teacher, Student)

### Frontend Stack
- **Framework:** React 18
- **Build Tool:** Vite
- **Development:** Dev server with API proxy to backend (port 3000 by default)
- **Production:** Built SPA served by backend (from `frontend/dist`)

### Architecture Pattern
```
Handler (HTTP) → Service (Business Logic) → Repository (Database Access) → Database
                                  ↓
                            Middleware (Auth, Roles)
```

---

## Core Models

### User
- Fields: `id`, `email`, `password` (hashed), `first_name`, `last_name`, `role` (admin/teacher/student), `status` (active/inactive)
- Default admin: email `admin@example.com`, password `admin123`

### Student
- Fields: `id`, `user_id`, `grade_level`, `enrollment_status`=- Links to User; used for student-specific data

### Teacher
- Fields: `id`, `user_id`, `department`, `specialization`
- Links to User; teacher-specific profile

### Course
- Fields: `id`, `name`, `code`, `department`, `credits`, `max_capacity`, `teacher_id`
- Represents a class/course; taught by a teacher

### Enrollment
- Fields: `id`, `student_id`, `course_id`, `enrollment_date`, `status` (active/completed/dropped)
- Links students to courses

### Grade
- Fields: `id`, `student_id`, `course_id`, `grade_value` (0–100), `date_recorded`
- Records a student's grade in a course

### Attendance
- Fields: `id`, `student_id`, `course_id`, `attendance_date`, `status` (present/absent/late)
- Tracks daily attendance

---

## Endpoints Overview

### Authentication (No Auth Required)
- `POST /api/auth/login` — Login with email/password → returns JWT token
- `POST /api/auth/register` — Register a new user

### User Management (Auth Required)
- `GET /api/users` — List all users
- `GET /api/users/:id` — Get user by ID
- `PUT /api/users/:id` — Update user (Admin)
- `GET /api/profile` — Get current user's profile
- `PUT /api/profile` — Update current user's profile

### Courses (Auth Required)
- `POST /api/courses` — Create course (Admin/Teacher)
- `GET /api/courses` — List all courses
- `GET /api/courses/:id` — Get course details
- `PUT /api/courses/:id` — Update course
- `DELETE /api/courses/:id` — Delete course (Admin)

### Students (Auth Required)
- `POST /api/students` — Create student (Admin)
- `GET /api/students` — List all students
- `GET /api/students/:id` — Get student details
- `PUT /api/students/:id` — Update student

### Enrollments (Auth Required)
- `POST /api/enrollments` — Enroll student in course
- `GET /api/enrollments/:id` — Get enrollment details
- `GET /api/enrollments/by-student/:studentId` — Get student's enrollments
- `GET /api/enrollments/by-course/:courseId` — Get course enrollments
- `PUT /api/enrollments/:id/status` — Update enrollment status (Active/Completed/Dropped)

### Grades (Auth Required)
- `POST /api/grades` — Record a grade (Teacher)
- `GET /api/grades/:id` — Get grade
- `GET /api/grades/by-student/:studentId` — Get student's grades
- `GET /api/grades/by-course/:courseId` — Get course grades
- `GET /api/grades/average/:studentId` — Get student's average grade
- `PUT /api/grades/:id` — Update grade

### Attendance (Auth Required)
- `POST /api/attendance` — Record attendance (Teacher)
- `GET /api/attendance/:id` — Get attendance record
- `GET /api/attendance/by-student/:studentId` — Get student's attendance
- `GET /api/attendance/by-course/:courseId` — Get course attendance
- `GET /api/attendance/student-course/:studentId/:courseId` — Get specific student-course attendance
- `GET /api/attendance/stats/:studentId/:courseId` — Get attendance stats

### Admin Only
- `GET /api/admin/dashboard` — Dashboard stats (total users, courses, enrollments, etc.)
- `GET /api/admin/health` — System health check
- `POST /api/admin/users` — Create user directly (bypasses normal registration)

### Teacher Routes (Shortcuts)
- `POST /api/teacher/grades` — Same as grade creation
- `POST /api/teacher/attendance` — Same as attendance creation

### Student Routes (Shortcuts)
- `GET /api/student/enrollments` — Get current student's enrollments
- `GET /api/student/grades` — Get current student's grades
- `GET /api/student/attendance` — Get current student's attendance

---

## Test Workflows

### Workflow 1: Register → Login → Update Profile

**1. Register a new user**
```powershell
$body = @{
  email = "john.doe@example.com"
  password = "SecurePass123!"
  first_name = "John"
  last_name = "Doe"
  role = "student"
} | ConvertTo-Json

Invoke-RestMethod -Uri http://localhost:8080/api/auth/register `
  -Method Post -ContentType application/json -Body $body
```

**Response:** `{ "id": 1, "email": "john.doe@example.com", "role": "student", ... }`

**2. Login**
```powershell
john.doe@example.com"
    password = "SecurePass123!"
} | ConvertTo-Json

$response = Invoke-RestMethod -Uri http://localhost:8080/api/auth/lo$body = @{
    email = "jogin `
  -Method Post -ContentType application/json -Body $body

$token = $response.token
Write-Host "Token: $token"
```

**3. Update your profile**
```powershell
$headers = @{
    Authorization = "Bearer $token"
    "Content-Type" = "application/json"
}

$body = @{
    first_name = "John Updated"
    last_name = "Doe Updated"
} | ConvertTo-Json

Invoke-RestMethod -Uri http://localhost:8080/api/profile `
  -Method Put -Headers $headers -Body $body
```

---

### Workflow 2: Create a Course (Admin)

**Login as admin first:**
```powershell
$body = @{
    email = "admin@example.com"
    password = "admin123"
} | ConvertTo-Json

$response = Invoke-RestMethod -Uri http://localhost:8080/api/auth/login `
  -Method Post -ContentType application/json -Body $body

$adminToken = $response.token
```

**Create a course:**
```powershell
$headers = @{
    Authorization = "Bearer $adminToken"
    "Content-Type" = "application/json"
}

$body = @{
    name = "Mathematics 101"
    code = "MATH101"
    department = "Science"
    credits = 3
    max_capacity = 40
    teacher_id = 1
} | ConvertTo-Json

Invoke-RestMethod -Uri http://localhost:8080/api/courses `
  -Method Post -Headers $headers -Body $body
```

---

### Workflow 3: Enroll & Grade a Student

**1. Enroll student in course:**
```powershell
$body = @{
    student_id = 1
    course_id = 1
} | ConvertTo-Json

Invoke-RestMethod -Uri http://localhost:8080/api/enrollments `
  -Method Post -Headers $headers -Body $body
```

**2. Record a grade (as teacher):**
```powershell
$body = @{
    student_id = 1
    course_id = 1
    grade_value = 92
    date_recorded = (Get-Date -Format "2006-01-02")
} | ConvertTo-Json

Invoke-RestMethod -Uri http://localhost:8080/api/grades `
  -Method Post -Headers $headers -Body $body
```

**3. Get student's grades:**
```powershell
Invoke-RestMethod -Uri http://localhost:8080/api/grades/by-student/1 `
  -Method Get -Headers $headers | ConvertTo-Json
```

---

### Workflow 4: Track Attendance

**Record attendance (teacher):**
```powershell
$body = @{
    student_id = 1
    course_id = 1
    attendance_date = (Get-Date -Format "2006-01-02")
    status = "present"
} | ConvertTo-Json

Invoke-RestMethod -Uri http://localhost:8080/api/attendance `
  -Method Post -Headers $headers -Body $body
```

**Get attendance stats:**
```powershell
Invoke-RestMethod -Uri http://localhost:8080/api/attendance/stats/1/1 `
  -Method Get -Headers $headers | ConvertTo-Json
```

---

### Workflow 5: Admin Dashboard

**Get system overview:**
```powershell
Invoke-RestMethod -Uri http://localhost:8080/api/admin/dashboard `
  -Method Get -Headers @{ Authorization = "Bearer $adminToken" } | ConvertTo-Json
```

**Response includes:**
- Total users
- Total active courses
- Total enrollments
- System health status

---

## Running the React Frontend

If you want to test the UI (not just API):

### Development Mode
```powershell
# In a new terminal, from repo root:
cd frontend
npm install
npm run dev
```
Then open `http://localhost:3000` in your browser.

**Features in React app:**
- Login / Register pages
- Profile page (view & edit user info)
- Dashboard (view enrolled courses, grades summary, attendance)
- Course details page (enroll, view status)

### Production Build
```powershell
cd frontend
npm install
npm run build
# Backend will serve the built app from http://localhost:8080
```

---

## Project Structure

```
school-management-system/
├── cmd/server/
│   └── main.go                 # App entry, router setup, middleware
├── internal/
│   ├── config/
│   │   └── config.go           # Database & app config
│   ├── models/                 # All data structures (User, Course, Student, etc.)
│   ├── repository/             # Database access layer
│   ├── service/                # Business logic layer
│   ├── handlers/               # HTTP handlers (API endpoints)
│   └── middleware/             # Auth & role-checking middleware
├── pkg/
│   ├── database/               # Database initialization
│   ├── logger/                 # Logging utilities
│   └── utils/                  # Password hashing, validation
├── frontend/                   # React app (src/, package.json, vite.config.js)
├── migrations/                 # Database schema migrations
├── tests/                      # Integration tests
└── go.mod / go.sum            # Go dependencies
```

---

## Role-Based Features

### Admin
- Create/manage all users
- Create/update courses
- View dashboard stats
- Manage system health

### Teacher
- Create courses
- Record grades and attendance
- View enrollments in their courses
- Update student grades/attendance

### Student
- View enrolled courses
- See own grades
- Check attendance record
- Update own profile

---

## Common Issues & Solutions

| Issue | Solution |
|-------|----------|
| `401 Unauthorized` on protected endpoints | Include `Authorization: Bearer <token>` header |
| Port 8080 already in use | Kill the process: `netstat -ano \| findstr :8080` → `taskkill /PID <PID>` |
| `404` on React pages | Ensure backend is running; Vite proxy needs `/api` calls to work |
| CORS errors in frontend dev | Vite proxy should handle `/api` (see `frontend/vite.config.js`); if not, backend CORS is enabled for `localhost:3000` |

---

## Next Steps

1. **Explore the API** — Copy test workflows above into PowerShell and run them
2. **Check the frontend** — Install React dependencies and run `npm run dev`
3. **Review source code** — Look at `internal/handlers/`, `internal/service/`, and `frontend/src/` for implementation details
4. **Customize** — Modify models, add endpoints, update UI to match your requirements

---

## Files to Review First

1. `cmd/server/main.go` — See how routes & middleware are set up
2. `internal/models/` — Understand data structures
3. `internal/handlers/auth_handler.go` — See authentication flow
4. `internal/service/auth_service.go` — Business logic
5. `frontend/src/App.jsx` — React app structure
6. `frontend/src/api.js` — How frontend calls the API

