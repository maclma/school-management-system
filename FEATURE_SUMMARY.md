# School Management System - Feature Summary

## Overview
A full-stack school management system built with Go (backend) and React (frontend), featuring role-based access control (Admin, Teacher, Student) with comprehensive course, enrollment, grade, and attendance management.

---

## Backend Features (Go + Gin)

### Authentication & Authorization
- **JWT-based authentication** with token validation
- **Role-based middleware** (Admin, Teacher, Student)
- **Secure password hashing** and validation
- **User profile management** with updateable fields

### Core Entities
1. **Users** - System users with roles and status
2. **Students** - Student profiles linked to users
3. **Teachers** - Teacher profiles linked to users
4. **Courses** - Academic courses with department, credits, capacity
5. **Enrollments** - Student enrollment in courses with approval workflow
6. **Grades** - Student performance records with scores and dates
7. **Attendance** - Attendance tracking per course with status (present/absent/late)

### API Endpoints

#### Public Routes
- `POST /api/auth/login` - User authentication
- `POST /api/auth/register` - User registration

#### Protected Routes (All authenticated users)
- **Users**: GET all, GET by ID, UPDATE, PATCH status
- **Profiles**: GET own profile, PUT update profile
- **Courses**: POST create, GET all, GET by ID, PUT update, DELETE, filter by department
- **Students**: Full CRUD operations
- **Enrollments**: Enroll, GET by student/course, UPDATE status, DELETE
- **Grades**: Record grade, GET by student/course, GET averages, UPDATE, DELETE
- **Attendance**: Record attendance, GET by student/course, GET stats, UPDATE, DELETE

#### Role-Restricted Routes

**Admin Only** (`/api/admin/`)
- `GET /admin/users` - List all users with pagination
- `POST /admin/users` - Create new user with role assignment
- `DELETE /admin/users/:id` - Delete user
- `GET /admin/dashboard` - System stats (total users, courses, enrollments, active students)
- `GET /admin/health` - System health check
- **NEW**: `GET /admin/enrollments` - List all enrollments with optional status filter
- **NEW**: `POST /admin/enrollments/:id/approve` - Approve pending enrollment
- **NEW**: `POST /admin/enrollments/:id/reject` - Reject pending enrollment

**Teacher Only** (`/api/teacher/`)
- `POST /teacher/grades` - Record grades
- `PUT /teacher/grades/:id` - Update grades
- `POST /teacher/attendance` - Record attendance
- `PUT /teacher/attendance/:id` - Update attendance

**Student Only** (`/api/student/`)
- `GET /student/enrollments` - View own enrollments
- `GET /student/grades` - View own grades
- `GET /student/attendance` - View own attendance

### Database
- **SQLite** with GORM ORM
- Auto-migration on startup
- Relational schema with foreign keys
- Admin user auto-seeding

---

## Frontend Features (React + Vite)

### Pages & Components

#### 1. **Header Component** (`Header.jsx`)
- Navigation bar with logo and role-based menu
- "Teach" button for teachers → redirects to `/teacher`
- "Admin" button for admins → redirects to `/admin`
- Profile button → redirects to `/profile`
- Logout button → clears auth tokens and localStorage

#### 2. **Login Page** (`Login.jsx`)
- Email and password input fields
- Stores JWT token and user role in localStorage
- Role-based redirection (teacher → dashboard, admin → admin panel)
- Error handling with toast notifications

#### 3. **Register Page** (`Register.jsx`)
- New user registration form
- Fields: first name, last name, email, password, role selection
- Role options: Student, Teacher, Admin
- Form validation and error feedback

#### 4. **Dashboard Page** (`Dashboard.jsx`)
- Student-focused view showing enrolled courses
- Displays course list with names, codes, departments, credits
- "View" button to see course details
- "Enroll in Course" button for new enrollments with course selection modal

#### 5. **Course Details Page** (`Course.jsx`)
- Course information display (name, code, department, credits, capacity)
- Current enrollment count and availability status
- Student enrollment status indication

#### 6. **Profile Page** (`Profile.jsx`)
- **View Mode**: Display user information in info grid
  - Name, email, phone, address, role, account status
- **Edit Mode**: Editable form with all profile fields
  - "Edit Profile" button to toggle edit mode
  - Save and Cancel buttons
  - API integration for profile updates

#### 7. **Teacher Panel** (`TeacherPanel.jsx`)
- **Three-Tab Interface**:
  1. **Grades Tab**:
     - Course selector dropdown
     - Student enrollments table for selected course
     - Search/filter by student ID
     - "Record Grade" button → modal form with student ID, score (0-100), date
     - Submission saves grade to backend
  
  2. **Attendance Tab**:
     - Same course selector and student table
     - Search/filter functionality
     - "Record Attendance" button → modal form with status (Present/Absent/Late), date
  
  3. **My Courses Tab**:
     - List of courses created by the teacher
     - Course cards displaying: name, code, department, credits, capacity
     - "+ New Course" button → modal form for course creation
     - Course creation fields: name, code, department, credits, capacity

#### 8. **Admin Dashboard** (`AdminDashboard.jsx`)
- **Four-Tab Interface**:
  1. **Stats Tab**:
     - Four stat tiles displaying:
       - Total Users
       - Total Courses
       - Total Enrollments
       - Active Students
     - Color-coded stat tiles with values
  
  2. **Users Tab**:
     - Full user list in table format
     - Columns: ID, Email, Name, Role, Status, Actions
     - **NEW**: Search/filter by email or name
     - Delete button for user removal
     - Responsive table with hover effects
  
  3. **Create User Tab**:
     - Form to create new users
     - Fields: email, first name, last name, password, role
     - Role dropdown (Student, Teacher, Admin)
     - Form submission with validation
  
  4. **Enrollments Button**:
     - Link to new Enrollment Approval page

#### 9. **Enrollment Approval Page** (`EnrollmentApproval.jsx`) **NEW**
- **Stat Summary Section**:
  - Pending enrollment count (orange badge)
  - Approved enrollment count (green badge)
  - Rejected enrollment count (red badge)
  - Real-time stat updates

- **Filter Buttons**:
  - Filter by status: Pending, Approved, Rejected, All
  - Shows count for each status in button label

- **Search Functionality**:
  - Search by student name, email, or course name
  - Real-time filtering as user types

- **Enrollment Table**:
  - Columns: ID, Student Name, Email, Course Name, Status (color-coded badge), Enrolled Date, Actions
  - Status badges with color coding:
    - Pending: Orange
    - Approved: Green
    - Rejected: Red
  
- **Action Buttons**:
  - For pending enrollments: "Approve" (green) and "Reject" (red) buttons
  - For approved/rejected: Displays status indicator (✓ or ✗)

- **API Integration**:
  - `GET /admin/enrollments` - Load all enrollments
  - `POST /admin/enrollments/:id/approve` - Approve enrollment
  - `POST /admin/enrollments/:id/reject` - Reject enrollment

### Global API Wrapper (`api.js`)
Centralized request utility with:
- **19 API methods** including:
  - Authentication: `login`, `register`
  - Profile: `getProfile`, `updateProfile`
  - Courses: `getCourses`, `getCourse`, `createCourse`
  - Enrollments: `enroll`, `getEnrollmentsByStudent`, `getCourseEnrollments`
  - Grades: `getGradesByStudent`, `recordGrade`
  - Attendance: `getAttendanceByStudent`, `recordAttendance`
  - Admin: `getAdminStats`, `getAdminUsers`, `createUserAdmin`, `deleteUser`
- JWT token auto-injection from localStorage
- Standard error handling and JSON parsing
- Raw `request()` method for custom API calls

### UI/UX Components & Styling

#### Style System
- **Color scheme**: Blue accent (#2563eb), gray muted (#6b7280)
- **Responsive grid**: Adapts from 1 column on mobile to multi-column on desktop
- **Typography**: Inter font stack with system fallbacks
- **Spacing**: Consistent 0.6rem base unit

#### UI Patterns
1. **Header**: Fixed top navigation with role-based buttons
2. **Cards**: Consistent card containers with shadow and rounded corners
3. **Forms**: Vertical layout with labels, input validation, buttons
4. **Tables**: Hover effects, alternating row backgrounds, sortable columns
5. **Tabs**: Active/inactive states with bottom border indicator
6. **Modals**: Overlay with centered content, max-width constraints, scrolling support
7. **Badges**: Inline status indicators with color coding
8. **Stat Tiles**: Large, centered numbers with labels (gradients and borders)
9. **Info Grids**: Multi-column layout for profile/detail information

#### Interactive Features
- **Toast notifications**: Success/error/info messages (auto-dismiss)
- **Modal forms**: For complex data entry (grades, attendance, course creation)
- **Search/filter**: Real-time input filtering on tables and lists
- **Tab switching**: Conditional rendering of tab content
- **Dropdowns**: Course/role selection in forms
- **Button states**: Hover effects, disabled states, loading states

### State Management
- **React Hooks**: `useState`, `useRef`, `useEffect`
- **localStorage**: Persistent storage for JWT token and user role
- **Local component state**: Tab selection, form data, modal visibility

### Navigation
- **Client-side routing**: Simple path-based router in `App.jsx`
- **Protected routes**: Role-based access checks
- **Navigation function**: `window.navigate()` for programmatic routing
- **Browser back button**: Supported via `onpopstate` listener

---

## Technology Stack

### Backend
- **Language**: Go 1.19+
- **Framework**: Gin Web Framework
- **ORM**: GORM v2
- **Database**: SQLite
- **Authentication**: JWT (github.com/dgrijalvo/jwt-go)
- **Logging**: Logrus
- **CORS**: gin-contrib/cors
- **Environment**: .env support via godotenv

### Frontend
- **Framework**: React 18
- **Build Tool**: Vite
- **Language**: JavaScript (ES6+)
- **Styling**: Plain CSS with responsive design
- **Storage**: localStorage for auth persistence
- **HTTP**: Fetch API for backend communication

### Infrastructure
- **Development**:
  - Backend: `go run cmd/server/main.go` (port 8080)
  - Frontend: `npm run dev` (port 3000 with Vite proxy)
- **Production**:
  - Built frontend in `frontend/dist`
  - Backend serves SPA fallback to index.html via NoRoute handler

---

## How to Use

### Setup
1. **Backend**: `go build -o server.exe ./cmd/server/main.go` then `./server.exe`
2. **Frontend**: `npm install` in `frontend/` then `npm run dev`
3. **Database**: Auto-created on first run (SQLite)

### Testing Accounts
- **Admin**: Create via registration with "Admin" role, or seed admin user
- **Teacher**: Create via registration with "Teacher" role
- **Student**: Create via registration with "Student" role

### Workflow Examples

**Teacher Recording Grades**:
1. Login as teacher
2. Click "Teach" in header
3. Go to "Grades" tab
4. Select course from dropdown
5. Click "Record Grade" for a student
6. Enter score (0-100) and date
7. Submit

**Admin Approving Enrollments**:
1. Login as admin
2. Click "Admin" in header
3. Click "Enrollments" button
4. Review pending enrollments
5. Click "Approve" or "Reject" for each
6. View approved/rejected enrollments in tabs

**Student Viewing Profile**:
1. Login as student
2. Click "Profile" in header
3. View current profile information
4. Click "Edit Profile" to modify fields
5. Click "Save Changes" to update

---

## File Structure
```
frontend/
  src/
    pages/
      Login.jsx              # Login page
      Register.jsx           # Registration page
      Dashboard.jsx          # Student dashboard
      Course.jsx             # Course details
      Profile.jsx            # User profile view/edit
      TeacherPanel.jsx       # Teacher grade/attendance/course management
      AdminDashboard.jsx     # Admin stats and user management
      EnrollmentApproval.jsx # NEW: Admin enrollment approval workflow
    components/
      Header.jsx             # Navigation header
    api.js                   # Centralized API wrapper
    App.jsx                  # Router and layout
    styles.css               # Global CSS styling
    ui.js                    # Toast and loading utilities
    index.html               # HTML entry point
  package.json
  vite.config.js

internal/
  config/                    # Configuration loading
  handlers/                  # HTTP handlers for all entities
    enrollment_handler.go    # NEW: Enrollment approval endpoints
  service/                   # Business logic layer
  repository/                # Data access layer
  models/                    # Data models
  middleware/                # Auth and role middleware

cmd/
  server/
    main.go                  # Entry point with routing

pkg/
  database/                  # Database connection
  logger/                    # Logging setup
  utils/                     # Utilities (password, validation)
```

---

## Recent Enhancements

### Session 3 Additions
1. ✅ **Enrollment Approval UI**
   - New `EnrollmentApproval.jsx` page for admin enrollment review
   - Status filtering (Pending/Approved/Rejected/All)
   - Search by student/course/email
   - Approve/Reject action buttons with confirmation
   - Real-time stat counters

2. ✅ **Admin Dashboard Improvements**
   - Added search/filter to Users tab (by email or name)
   - Added "Enrollments" tab button linking to approval page
   - Improved user list with color-coded badges

3. ✅ **Backend Routes**
   - `GET /api/admin/enrollments` - List all enrollments with optional filter
   - `POST /api/admin/enrollments/:id/approve` - Approve enrollment
   - `POST /api/admin/enrollments/:id/reject` - Reject enrollment

4. ✅ **CSS Enhancements**
   - Added `.button.success` styling for approve buttons
   - Stat counter styling for enrollment approval page
   - Improved table and badge layouts

---

## Future Enhancement Opportunities
- [ ] Icons via react-icons or feather-icons library
- [ ] Advanced animations and transitions
- [ ] Unit tests for React components
- [ ] E2E testing with Cypress
- [ ] Error boundary components
- [ ] Form validation enhancements (email uniqueness, password strength)
- [ ] Email notifications for enrollments
- [ ] Bulk user import (CSV)
- [ ] Grade distribution analytics
- [ ] Attendance reports and visualizations
- [ ] Export functionality (CSV, PDF)
- [ ] Multi-language support (i18n)
- [ ] Dark mode theme
- [ ] Mobile app (React Native)
