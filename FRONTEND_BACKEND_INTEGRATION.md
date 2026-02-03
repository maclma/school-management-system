# Frontend-Backend Integration Guide

## Current Status ✅

### Backend (Go/Gin)
- **Status**: Running on port 8080
- **Location**: `cmd/server/main.go`
- **Database**: Connected and migrated
- **Admin User**: Auto-created (admin/admin by default)
- **CORS**: Configured for localhost:3000, localhost:5173, localhost:8080

### Frontend (React/Vite)
- **Status**: Ready on port 3000
- **Location**: `frontend/` directory
- **Framework**: React 18.2 with Vite
- **API Integration**: Already configured via `frontend/api.js`

## Starting the Services

### Terminal 1: Backend Server
```bash
cd c:\Users\dell\school-management-system
go run ./cmd/server/main.go
```

### Terminal 2: Frontend Development Server (Vite)
```bash
cd c:\Users\dell\school-management-system\frontend
npm install  # Only first time
npm run dev
```

## API Configuration

The frontend API client is already configured in [frontend/api.js](frontend/api.js):

```javascript
const base = 'http://localhost:8080';
```

### Available API Methods

#### Authentication
- `api.login(email, password)` - POST /api/auth/login
- `api.register(payload)` - POST /api/auth/register

#### Profile & Users
- `api.getProfile()` - GET /api/profile
- `api.updateProfile(data)` - PUT /api/profile

#### Courses
- `api.getCourses(page, limit)` - GET /api/courses
- `api.getCourse(id)` - GET /api/courses/:id

#### Enrollments
- `api.createEnrollment(studentId, courseId)` - POST /api/enrollments
- `api.getEnrollmentsByStudent(studentId)` - GET /api/enrollments/by-student/:studentId

#### Grades
- `api.getGradesByStudent(studentId)` - GET /api/grades/by-student/:studentId

#### Attendance
- `api.getAttendanceByStudent(studentId)` - GET /api/attendance/by-student/:studentId

## Testing the Integration

### 1. Login Test
**URL**: http://localhost:3000/login.html
**Default Admin Credentials**:
- Email: admin@school.com
- Password: admin

### 2. Dashboard Test
**URL**: http://localhost:3000/dashboard.html
- View profile information
- List available courses
- Enroll in courses
- View enrollments

### 3. Course Details Test
**URL**: http://localhost:3000/course.html?id=1
- View course information
- Enroll in course
- View related content

## Token Management

Tokens are automatically stored in `localStorage` under the key `sms_token` and attached to all authenticated requests via the Authorization header.

## CORS Setup

The backend is configured to accept requests from:
- http://localhost:3000
- http://localhost:5173 (Vite dev server)
- http://127.0.0.1:3000
- http://127.0.0.1:5173
- http://localhost:8080
- http://127.0.0.1:8080

## Common Issues & Solutions

### Issue: CORS Error
**Solution**: Ensure backend CORS middleware is active and frontend is accessing from an allowed origin.

### Issue: Token Not Persisting
**Solution**: Check browser's localStorage is enabled and not in private mode.

### Issue: API Endpoint Not Found
**Solution**: Verify backend routes are registered in `cmd/server/main.go`

### Issue: 401 Unauthorized
**Solution**: Token may have expired - user should login again.

## Next Steps

1. **Start the backend** on Terminal 1
2. **Start the frontend dev server** on Terminal 2
3. **Open** http://localhost:3000 in browser
4. **Login** with admin credentials
5. **Test** the dashboard and course features

## Admin Reset

If you need to reset admin credentials:
```bash
cd c:\Users\dell\school-management-system
go run reset_admin.go
```

## Database Inspection

To check users in the database:
```sql
SELECT id, email, name, role FROM users LIMIT 10;
```

## Frontend File Structure

```
frontend/
├── index.html           # Main entry point
├── login.html          # Login page
├── register.html       # Registration page
├── dashboard.html      # User dashboard
├── course.html         # Course details
├── profile.html        # User profile
├── api.js              # API client (sets base URL)
├── server.go           # Simple HTTP server
├── styles.css          # Global styles
├── ui.js               # UI utilities
├── package.json        # NPM dependencies
├── vite.config.js      # Vite configuration
└── src/
    ├── main.jsx       # React entry point
    ├── App.jsx        # Main app component
    └── components/    # React components
```

## Backend API Routes

All API routes are prefixed with `/api/` and are documented in `cmd/server/main.go`.

Key route groups:
- `/api/health` - Health check (public)
- `/api/auth/` - Authentication (public)
- `/api/courses` - Course management
- `/api/students` - Student management
- `/api/enrollments` - Enrollment management
- `/api/grades` - Grade management
- `/api/attendance` - Attendance tracking
- `/api/admin/` - Admin operations

## Performance Notes

- Token validation happens on every protected request
- Database indexes are automatically created on startup
- Request IDs are tracked for logging
- Rate limiting is applied to auth endpoints

---

**Last Updated**: 2026-01-30
**Integration Status**: ✅ Complete and Tested
