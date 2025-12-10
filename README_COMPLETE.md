# School Management System - Complete Documentation

A comprehensive, full-stack school management system with role-based access control, built with Go (backend) and React (frontend).

---

## 🎯 Project Overview

This system enables institutions to manage:
- **User Management** - Create and manage users with roles (Admin, Teacher, Student)
- **Course Management** - Teachers create courses; admins manage catalog
- **Enrollment Management** - Students enroll; admins approve enrollments
- **Grade Tracking** - Teachers record grades; students view grades
- **Attendance Tracking** - Teachers record attendance; students view records
- **Admin Dashboard** - System-wide statistics and user management
- **Teacher Panel** - Grade and attendance recording, course creation
- **Student Dashboard** - View enrollments, grades, attendance

---

## 🏗️ Architecture

### Backend Architecture (Go + Gin)
```
Request → CORS Middleware → Auth Middleware → Role Middleware → Handler
         ↓
    Handler (validates input)
         ↓
    Service (business logic, authorization)
         ↓
    Repository (data access)
         ↓
    Database (SQLite + GORM)
```

### Frontend Architecture (React + Vite)
```
App.jsx (Router)
    ↓
Header.jsx (Navigation + Logout)
    ↓
[Login | Register | Dashboard | Profile | TeacherPanel | AdminDashboard | EnrollmentApproval]
    ↓
api.js (Centralized API calls)
    ↓
Backend (HTTP)
```

---

## 🚀 Quick Start

### Prerequisites
- Go 1.19+
- Node.js 16+
- npm/yarn

### Setup Steps

#### 1. Clone and Navigate
```bash
cd c:\Users\dell\school-management-system
```

#### 2. Start Backend
```powershell
# Terminal 1
go build -o server.exe ./cmd/server/main.go
.\server.exe
# Server runs on http://localhost:8080
```

#### 3. Start Frontend
```powershell
# Terminal 2
cd frontend
npm install  # First time only
npm run dev
# Frontend runs on http://localhost:3000
```

#### 4. Access Application
Open http://localhost:3000 in browser

---

## 📋 Database Schema

### Core Tables
- **users** - System users with role and status
- **students** - Student profiles (extends users)
- **teachers** - Teacher profiles (extends users)
- **courses** - Academic courses with capacity
- **enrollments** - Student → Course relationships (with approval status)
- **grades** - Student grades per course
- **attendance** - Attendance records per student per course

### Relationships
```
User (1) ──→ (Many) Enrollment
         ──→ (Many) Grade
         ──→ (Many) Attendance

Course (1) ──→ (Many) Enrollment
          ──→ (Many) Grade
          ──→ (Many) Attendance

Student (1) ──→ (Many) Enrollment
Enrollment (Many) ──→ (1) Course
```

---

## 🔐 Authentication & Authorization

### JWT Authentication
- Login endpoint returns JWT token
- Token stored in `localStorage` as `sms_token`
- Auto-injected in request headers: `Authorization: Bearer {token}`
- Token validated in all protected routes

### Role-Based Access Control
- **Admin**: System management, user creation/deletion, enrollment approval
- **Teacher**: Course creation, grade/attendance recording
- **Student**: View own enrollments, grades, attendance, enroll in courses

### Routes by Role
```
Public:
  POST /api/auth/login
  POST /api/auth/register

Admin Only:
  GET    /api/admin/users
  POST   /api/admin/users
  DELETE /api/admin/users/:id
  GET    /api/admin/dashboard
  GET    /api/admin/enrollments          [NEW]
  POST   /api/admin/enrollments/:id/approve [NEW]
  POST   /api/admin/enrollments/:id/reject  [NEW]

Teacher Only:
  POST   /api/teacher/grades
  PUT    /api/teacher/grades/:id
  POST   /api/teacher/attendance
  PUT    /api/teacher/attendance/:id

Student Only:
  GET    /api/student/enrollments
  GET    /api/student/grades
  GET    /api/student/attendance

All Authenticated:
  [All CRUD operations on public entities]
```

---

## 📊 API Endpoints Reference

### Authentication
| Method | Endpoint | Body | Response |
|--------|----------|------|----------|
| POST | `/api/auth/login` | `{email, password}` | `{token, user}` |
| POST | `/api/auth/register` | `{email, password, first_name, last_name, role}` | `{token, user}` |

### Profiles
| Method | Endpoint | Body | Response |
|--------|----------|------|----------|
| GET | `/api/profile` | - | `{user}` |
| PUT | `/api/profile` | `{first_name, last_name, phone, address}` | `{success}` |

### Courses
| Method | Endpoint | Body | Response |
|--------|----------|------|----------|
| POST | `/api/courses` | `{name, code, department, credits, capacity, teacher_id}` | `{course}` |
| GET | `/api/courses` | - | `{data: [courses], total}` |
| GET | `/api/courses/:id` | - | `{course}` |
| PUT | `/api/courses/:id` | `{name, code, department, credits, capacity}` | `{success}` |
| DELETE | `/api/courses/:id` | - | `{success}` |

### Enrollments
| Method | Endpoint | Body | Response |
|--------|----------|------|----------|
| POST | `/api/enrollments` | `{student_id, course_id, status?}` | `{enrollment}` |
| GET | `/api/enrollments/:id` | - | `{enrollment}` |
| GET | `/api/enrollments/by-student/:studentId` | - | `{data: [enrollments], total}` |
| GET | `/api/enrollments/by-course/:courseId` | - | `{data: [enrollments], total}` |
| PUT | `/api/enrollments/:id/status` | `{status}` | `{success}` |
| DELETE | `/api/enrollments/:id` | - | `{success}` |

### Grades
| Method | Endpoint | Body | Response |
|--------|----------|------|----------|
| POST | `/api/grades` | `{student_id, course_id, score, recorded_date}` | `{grade}` |
| GET | `/api/grades/:id` | - | `{grade}` |
| GET | `/api/grades/by-student/:studentId` | - | `{data: [grades], total}` |
| GET | `/api/grades/by-course/:courseId` | - | `{data: [grades], total}` |
| GET | `/api/grades/average/:studentId` | - | `{average}` |
| PUT | `/api/grades/:id` | `{score, recorded_date}` | `{success}` |
| DELETE | `/api/grades/:id` | - | `{success}` |

### Attendance
| Method | Endpoint | Body | Response |
|--------|----------|------|----------|
| POST | `/api/attendance` | `{student_id, course_id, status, recorded_date}` | `{attendance}` |
| GET | `/api/attendance/:id` | - | `{attendance}` |
| GET | `/api/attendance/by-student/:studentId` | - | `{data: [records], total}` |
| GET | `/api/attendance/by-course/:courseId` | - | `{data: [records], total}` |
| GET | `/api/attendance/student-course/:studentId/:courseId` | - | `{data: [records], total}` |
| GET | `/api/attendance/stats/:studentId/:courseId` | - | `{present, absent, late, total}` |
| PUT | `/api/attendance/:id` | `{status, recorded_date}` | `{success}` |
| DELETE | `/api/attendance/:id` | - | `{success}` |

### Admin Operations
| Method | Endpoint | Body | Response |
|--------|----------|------|----------|
| GET | `/api/admin/users` | - | `{data: [users], total}` |
| POST | `/api/admin/users` | `{email, password, first_name, last_name, role}` | `{user}` |
| DELETE | `/api/admin/users/:id` | - | `{success}` |
| GET | `/api/admin/dashboard` | - | `{total_users, total_courses, total_enrollments, active_students}` |
| **GET** | **/api/admin/enrollments** | - | `{data: [enrollments], total}` |
| **POST** | **/api/admin/enrollments/:id/approve** | - | `{success}` |
| **POST** | **/api/admin/enrollments/:id/reject** | - | `{success}` |

---

## 🎨 Frontend Pages

### 1. Login (`/`)
- Email and password fields
- Form validation
- Redirects to dashboard on success
- Toast error notifications

### 2. Register (`/register`)
- First/Last name, email, password fields
- Role selection dropdown (Student/Teacher/Admin)
- Form validation
- Success redirects to login

### 3. Dashboard (`/dashboard`)
- Shows student's enrolled courses (cards)
- Course details: name, code, department, credits
- "View" button for course details
- "Enroll" button with course selector modal
- Student-specific view

### 4. Course Details (`/course/:id`)
- Full course information
- Enrollment status
- Availability info
- Enroll button

### 5. Profile (`/profile`)
- **View Mode**: Display all profile fields
  - Name, email, phone, address, role, status
  - "Edit Profile" button
- **Edit Mode**: Form with all fields editable
  - Save/Cancel buttons
  - Real-time updates via API

### 6. Teacher Panel (`/teacher`)
**Grades Tab**:
- Course selector dropdown
- Student enrollments table for course
- Search/filter by student ID
- "Record Grade" button → modal form
- Grade: student ID, score (0-100), date

**Attendance Tab**:
- Same course selector and student table
- "Record Attendance" button → modal form
- Attendance: status (Present/Absent/Late), date

**My Courses Tab**:
- List of teacher's courses as cards
- Course info: name, code, department, credits, capacity
- "+ New Course" button → modal form
- Course creation fields

### 7. Admin Dashboard (`/admin`)
**Stats Tab**:
- Four stat tiles with system metrics
- Total Users, Courses, Enrollments, Active Students
- Color-coded tiles with large values

**Users Tab**:
- Table of all system users
- Columns: ID, Email, Name, Role, Status, Delete action
- Search/filter by email or name
- Delete button for user removal

**Create User Tab**:
- Form to add new users
- Fields: email, first name, last name, password, role
- Form submission

**Enrollments Button**:
- Link to enrollment approval page

### 8. Enrollment Approval (`/admin/enrollments`) **NEW**
**Stat Summary**:
- Pending count (orange)
- Approved count (green)
- Rejected count (red)
- Real-time updates

**Filter Buttons**:
- Pending, Approved, Rejected, All
- Shows count in button label

**Search**:
- By student name, email, or course name
- Real-time filtering

**Enrollment Table**:
- ID, Student Name, Email, Course, Status (color badge), Date, Actions
- For pending: Approve (green) and Reject (red) buttons
- For approved/rejected: Status indicator (✓ or ✗)

---

## 🎨 UI Components & Styling

### Design System
- **Colors**: Blue (#2563eb), Red (#dc2626), Green (#059669), Orange (#f39c12)
- **Typography**: Inter font family
- **Spacing**: 0.6rem base unit, consistent grid
- **Shadows**: Subtle drop shadows on cards
- **Borders**: 1px solid #e6e9ee

### Component Library
- **Header**: Fixed navigation with role buttons
- **Cards**: Containers with shadow and rounded corners
- **Forms**: Vertical layouts with validation
- **Tables**: Striped rows, hover effects, responsive
- **Tabs**: Active state with bottom border
- **Modals**: Centered overlay with close-on-outside
- **Badges**: Inline status with color coding
- **Buttons**: Multiple variants (primary, secondary, success, danger)
- **Toast**: Auto-dismiss notifications (success, error, info)

### Responsive Design
- Mobile: Single column, stacked layout
- Tablet: Two columns, adjusted spacing
- Desktop: Multi-column grids, full width tables

---

## 📁 Project Structure

```
school-management-system/
├── cmd/
│   └── server/
│       └── main.go                    # Entry point, routing, server config
├── internal/
│   ├── config/
│   │   └── config.go                  # Configuration from environment
│   ├── handlers/
│   │   ├── auth_handler.go            # Login/register
│   │   ├── user_handler.go            # User CRUD
│   │   ├── course_handler.go          # Course CRUD
│   │   ├── student_handler.go         # Student CRUD
│   │   ├── enrollment_handler.go      # Enrollment + Approval [ENHANCED]
│   │   ├── grade_handler.go           # Grade recording
│   │   ├── attendance_handler.go      # Attendance recording
│   │   ├── teacher_handler.go         # Teacher-specific
│   │   └── admin_handler.go           # Admin operations
│   ├── middleware/
│   │   └── auth.go                    # JWT validation, role checking
│   ├── models/
│   │   ├── user.go
│   │   ├── student.go
│   │   ├── teacher.go
│   │   ├── course.go
│   │   ├── enrollment.go
│   │   ├── grade.go
│   │   └── attendance.go
│   ├── repository/
│   │   ├── user_repository.go
│   │   ├── course_repository.go
│   │   ├── student_repository.go
│   │   ├── enrollment_repository.go
│   │   ├── grade_repository.go
│   │   └── attendance_repository.go
│   ├── service/
│   │   ├── auth_service.go
│   │   ├── user_service.go
│   │   ├── course_service.go
│   │   ├── student_service.go
│   │   ├── enrollment_service.go
│   │   ├── grade_service.go
│   │   └── attendance_service.go
│   └── config/
│       └── config.go
├── pkg/
│   ├── database/
│   │   └── database.go                # GORM initialization
│   ├── logger/
│   │   └── logger.go                  # Logrus setup
│   └── utils/
│       ├── password.go                # Hashing
│       └── validator.go               # Input validation
├── migrations/
│   └── 001_init_schema.sql           # Initial schema
├── scripts/
│   ├── setup_db.sql                   # DB initialization
│   └── seed_db.sql                    # Test data
├── frontend/
│   ├── src/
│   │   ├── pages/
│   │   │   ├── Login.jsx
│   │   │   ├── Register.jsx
│   │   │   ├── Dashboard.jsx
│   │   │   ├── Course.jsx
│   │   │   ├── Profile.jsx
│   │   │   ├── TeacherPanel.jsx
│   │   │   ├── AdminDashboard.jsx
│   │   │   └── EnrollmentApproval.jsx [NEW]
│   │   ├── components/
│   │   │   └── Header.jsx
│   │   ├── api.js                     # HTTP wrapper (19 methods)
│   │   ├── ui.js                      # Toast & helpers
│   │   ├── App.jsx                    # Router
│   │   ├── styles.css                 # Global CSS
│   │   └── index.html
│   ├── package.json
│   └── vite.config.js
├── go.mod
├── makefile
├── README.md
├── FEATURE_SUMMARY.md                 # Feature documentation
└── QUICK_TEST_GUIDE.md               # Testing guide
```

---

## 🔄 Common Workflows

### Teacher Recording Grades
1. Login as teacher
2. Click "Teach" in header → TeacherPanel
3. Select "Grades" tab
4. Select course from dropdown
5. Click "Record Grade" for a student
6. Enter score and date
7. Submit form → saves to database

### Admin Approving Enrollments
1. Login as admin
2. Click "Admin" → Admin Dashboard
3. Click "Enrollments" button
4. Filter by "Pending" status
5. Search for student/course if needed
6. Click "Approve" or "Reject"
7. Confirm action → status updates

### Student Enrolling in Course
1. Login as student
2. Dashboard shows enrolled courses
3. Click "Enroll in Course"
4. Select course from modal dropdown
5. Submit → enrollment created (pending admin approval)
6. Check admin approval status in enrollments page

### Admin Creating User
1. Login as admin
2. Click "Admin" → Admin Dashboard
3. Click "Create User" tab
4. Fill form: email, name, password, role
5. Select role: Student/Teacher/Admin
6. Submit → user created and can login

---

## 🧪 Testing

### Manual Testing via PowerShell
See `QUICK_TEST_GUIDE.md` for detailed commands to:
- Register users with different roles
- Create courses
- Enroll students
- Record grades/attendance
- Approve enrollments

### Test Accounts
Create these via registration:
- Admin: email="admin@test.com", role="admin"
- Teacher: email="teacher@test.com", role="teacher"
- Student: email="student@test.com", role="student"
- Password: "test123" (or any password)

### Browser DevTools Testing
1. Open http://localhost:3000
2. Press F12 for DevTools
3. Check "Network" tab for API requests
4. Check "Console" for errors
5. Check "Application" → localStorage for token

---

## ⚙️ Configuration

### Environment Variables (`.env` file)
```
GIN_MODE=debug                 # or release
SERVER_PORT=8080
JWT_SECRET=your-secret-key
JWT_EXPIRY=72h
DATABASE_PATH=./school.db
LOG_LEVEL=info
```

### Build & Run
```bash
# Build
go build -o server.exe ./cmd/server/main.go

# Run
./server.exe

# Or use make
make build
make run
```

---

## 🐛 Troubleshooting

### Port Already in Use
```powershell
# Find and kill process on port 8080
Get-NetTCPConnection -LocalPort 8080 | Select-Object OwningProcess
Stop-Process -Id <PID> -Force
```

### Frontend Not Updating
1. Clear browser cache (Ctrl+Shift+Delete)
2. Check Vite dev server is running: `npm run dev`
3. Check console for errors (F12)
4. Rebuild frontend: `npm run build`

### Backend Not Responding
1. Check logs: `Get-Content server_out.log -Tail 50`
2. Verify port: `netstat -an | findstr 8080`
3. Rebuild: `go build -o server.exe ./cmd/server/main.go`
4. Restart: Stop process and run binary again

### Database Issues
1. Delete `school.db` file
2. Restart backend (will recreate schema)
3. Check `server_err.log` for migration errors
4. Verify database path in config

---

## 📈 Performance Considerations

- **Pagination**: Courses and enrollments support pagination (page, limit params)
- **Indexing**: Database has indexes on frequently queried fields
- **Caching**: Could be added for courses and user profiles
- **Rate Limiting**: Could be implemented for auth endpoints
- **Database**: SQLite suitable for development; consider PostgreSQL for production

---

## 🔒 Security

### Implemented
- JWT authentication for all protected routes
- Password hashing (bcrypt)
- Role-based access control (RBAC) with middleware
- CORS configuration for cross-origin requests
- Input validation on all endpoints
- Error messages don't leak sensitive info

### Recommendations
- Use HTTPS in production
- Implement rate limiting on auth endpoints
- Add CSRF tokens for state-changing operations
- Implement email verification for registration
- Add password strength requirements
- Implement account lockout after failed login attempts
- Regular security audits and dependency updates

---

## 📚 Additional Resources

- [Gin Framework Documentation](https://github.com/gin-gonic/gin)
- [GORM Documentation](https://gorm.io/)
- [React Documentation](https://react.dev/)
- [Vite Documentation](https://vitejs.dev/)
- [JWT Introduction](https://jwt.io/introduction)

---

## 📝 License

This project is provided as-is for educational purposes.

---

## 🎓 Learning Outcomes

This project demonstrates:
- Full-stack web application architecture
- RESTful API design and implementation
- Database design and ORM usage
- Authentication and authorization patterns
- Frontend-backend integration
- React component patterns and hooks
- Go backend patterns and middleware
- Role-based access control
- Error handling and validation
- Responsive UI/UX design

---

**Last Updated**: December 2024
**Version**: 1.0.0
**Status**: Feature Complete with Session 3 Enhancements
