# School Management System - Quick Test Guide

## Setup Quick Commands

### Start Backend
```powershell
cd c:\Users\dell\school-management-system
go build -o server.exe ./cmd/server/main.go
.\server.exe
```

### Start Frontend (in another terminal)
```powershell
cd c:\Users\dell\school-management-system\frontend
npm run dev
```

Both servers should be running:
- Backend: http://localhost:8080
- Frontend: http://localhost:3000

---

## Test Workflow

### 1. Test Admin Enrollment Approval Feature

#### Create Test Data via PowerShell
```powershell
# 1. Register an Admin user
$adminPayload = @{
    email = "admin@test.com"
    password = "admin123"
    first_name = "Admin"
    last_name = "User"
    role = "admin"
} | ConvertTo-Json

$adminRes = Invoke-WebRequest -Uri "http://localhost:8080/api/auth/register" `
  -Method POST -Body $adminPayload -ContentType "application/json"
$adminToken = ($adminRes.Content | ConvertFrom-Json).token

# 2. Register a Teacher user
$teacherPayload = @{
    email = "teacher@test.com"
    password = "teacher123"
    first_name = "Teacher"
    last_name = "User"
    role = "teacher"
} | ConvertTo-Json

$teacherRes = Invoke-WebRequest -Uri "http://localhost:8080/api/auth/register" `
  -Method POST -Body $teacherPayload -ContentType "application/json"
$teacherToken = ($teacherRes.Content | ConvertFrom-Json).token

# 3. Register a Student user
$studentPayload = @{
    email = "student@test.com"
    password = "student123"
    first_name = "Student"
    last_name = "User"
    role = "student"
} | ConvertTo-Json

$studentRes = Invoke-WebRequest -Uri "http://localhost:8080/api/auth/register" `
  -Method POST -Body $studentPayload -ContentType "application/json"
$studentToken = ($studentRes.Content | ConvertFrom-Json).token
```

#### Create a Course (as Teacher)
```powershell
$coursePayload = @{
    name = "Mathematics 101"
    code = "MATH101"
    department = "Science"
    credits = 3
    capacity = 30
    teacher_id = 2  # Teacher ID from registration
} | ConvertTo-Json

$courseRes = Invoke-WebRequest -Uri "http://localhost:8080/api/courses" `
  -Method POST -Body $coursePayload -ContentType "application/json" `
  -Headers @{"Authorization" = "Bearer $teacherToken"}

$courseId = ($courseRes.Content | ConvertFrom-Json).id
```

#### Enroll Student in Course
```powershell
$enrollPayload = @{
    student_id = 3
    course_id = $courseId
    status = "pending"
} | ConvertTo-Json

$enrollRes = Invoke-WebRequest -Uri "http://localhost:8080/api/enrollments" `
  -Method POST -Body $enrollPayload -ContentType "application/json" `
  -Headers @{"Authorization" = "Bearer $studentToken"}

$enrollmentId = ($enrollRes.Content | ConvertFrom-Json).enrollment.id
```

#### Admin Approves Enrollment
```powershell
# List all enrollments
$allRes = Invoke-WebRequest -Uri "http://localhost:8080/api/admin/enrollments" `
  -Method GET -Headers @{"Authorization" = "Bearer $adminToken"}

$allRes.Content | ConvertFrom-Json

# Approve enrollment
$approveRes = Invoke-WebRequest -Uri "http://localhost:8080/api/admin/enrollments/$enrollmentId/approve" `
  -Method POST -ContentType "application/json" `
  -Headers @{"Authorization" = "Bearer $adminToken"}

$approveRes.Content
```

---

### 2. Test UI via Browser

#### Login Flow
1. Go to http://localhost:3000
2. Register as admin user with email/password
3. See "Admin" button in header
4. Click "Admin" button

#### Admin Dashboard
1. View "Stats" tab with system metrics
2. View "Users" tab and search by name/email
3. Search bar shows users matching query
4. Create new user in "Create User" tab

#### Enrollment Approval (NEW)
1. Click "Enrollments" button from Admin Dashboard
2. See stat counters for Pending/Approved/Rejected
3. Use filter buttons to show different statuses
4. Use search bar to find enrollments by student/course
5. Click "Approve" or "Reject" for pending enrollments
6. Verify status changes in real-time

#### Teacher Panel
1. Login as teacher user (role = "teacher")
2. Click "Teach" button in header
3. **Grades Tab**:
   - Select course from dropdown
   - See enrolled students
   - Search students by ID
   - Click "Record Grade" to open modal
   - Submit grade form
4. **Attendance Tab**:
   - Same interface
   - Record attendance status (Present/Absent/Late)
5. **My Courses Tab**:
   - View created courses as cards
   - Click "+ New Course" to create course
   - Fill form and submit

---

## Expected Results

✅ Backend Routes Registered (visible in logs)
✅ Enrollment approval endpoints working
✅ Admin can view all enrollments
✅ Admin can approve/reject enrollments
✅ Admin Dashboard has enrollment tab
✅ EnrollmentApproval page filters and searches correctly
✅ Toast notifications show on actions
✅ Tables and badges display with correct styling

---

## Troubleshooting

### Frontend not reflecting changes?
- Check browser console (F12) for errors
- Verify API calls in Network tab
- Check backend logs in `server_out.log` and `server_err.log`

### Backend routes not showing?
- Rebuild: `go build -o server.exe ./cmd/server/main.go`
- Check `server_out.log` for route registration messages
- Ensure process is running: `Get-Process server` in PowerShell

### Database issues?
- Delete `school.db` to reset database
- Restart backend to recreate schema
- Check logs for migration errors

---

## Key Features Demonstrated

### Session 3 Additions
1. ✅ Enrollment Approval Workflow for Admins
   - List all enrollments with status filter
   - Search by student/course/email
   - Approve/Reject buttons with confirmation
   - Real-time stat counters

2. ✅ Enhanced Admin Dashboard
   - Search/filter users by name or email
   - Link to enrollment approval page
   - Color-coded stat tiles

3. ✅ New Backend Routes
   - Admin-only enrollment management endpoints
   - Proper role-based access control

4. ✅ UI Improvements
   - Better button styling (success/danger colors)
   - Responsive table layouts
   - Modal forms for actions
   - Search/filter on multiple pages
