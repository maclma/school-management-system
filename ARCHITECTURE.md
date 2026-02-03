# Frontend-Backend Architecture

## System Architecture Diagram

```
┌─────────────────────────────────────────────────────────────────┐
│                         USER BROWSER                             │
│  (Chrome, Firefox, Safari, Edge)                                │
└──────────────────────┬──────────────────────────────────────────┘
                       │
                       │ HTTP/HTTPS
                       │ Port 3000
                       │
       ┌───────────────┴──────────────────┐
       │                                  │
┌──────▼──────┐                    ┌──────▼──────┐
│  index.html │  Static Files      │  login.html │
│ dashboard   │  (HTML/CSS/JS)     │  profile    │
│ course.html │                    │  register   │
└──────┬──────┘                    └──────┬──────┘
       │                                  │
       └──────────────┬───────────────────┘
                      │
        ┌─────────────▼─────────────┐
        │   Frontend Application     │
        │   (React + Vite)          │
        │                           │
        │  ┌──────────────────────┐ │
        │  │  api.js              │ │
        │  │  (API Client)        │ │
        │  │                      │ │
        │  │ - Stores JWT token   │ │
        │  │ - Makes HTTP calls   │ │
        │  │ - Handles responses  │ │
        │  └──────┬───────────────┘ │
        └─────────┼──────────────────┘
                  │
                  │ REST API Calls
                  │ Port 8080
                  │ Content-Type: application/json
                  │
        ┌─────────▼──────────────────┐
        │   Backend Application       │
        │   (Go + Gin Framework)      │
        │                            │
        │   ┌────────────────────┐  │
        │   │  Router/Routes     │  │
        │   │  /api/auth/*       │  │
        │   │  /api/courses/*    │  │
        │   │  /api/students/*   │  │
        │   │  /api/grades/*     │  │
        │   │  ... 25+ routes    │  │
        │   └────────────┬───────┘  │
        │                │          │
        │   ┌────────────▼───────┐  │
        │   │  Middleware Layer  │  │
        │   │  - Auth (JWT)      │  │
        │   │  - CORS            │  │
        │   │  - Rate Limiting   │  │
        │   │  - Validation      │  │
        │   └────────────┬───────┘  │
        │                │          │
        │   ┌────────────▼─────┐   │
        │   │ Handler Layer    │   │
        │   │ (25+ Handlers)   │   │
        │   │                  │   │
        │   │ - AuthHandler    │   │
        │   │ - UserHandler    │   │
        │   │ - CourseHandler  │   │
        │   │ - GradeHandler   │   │
        │   │ - etc.           │   │
        │   └────────────┬─────┘   │
        │                │         │
        │   ┌────────────▼────┐    │
        │   │ Service Layer   │    │
        │   │ Business Logic  │    │
        │   └────────────┬────┘    │
        │                │         │
        │   ┌────────────▼────┐    │
        │   │Repository Layer │    │
        │   │Data Access      │    │
        │   └────────────┬────┘    │
        └────────────────┼─────────┘
                         │
                         │ SQL Queries
                         │ Connection Pooling
                         │
        ┌────────────────▼────────────────┐
        │   PostgreSQL Database            │
        │   (or configured in .env)        │
        │                                  │
        │  Tables:                         │
        │  - users (auth data)             │
        │  - students (enrollment)         │
        │  - courses (course data)         │
        │  - enrollments (registrations)   │
        │  - grades (student grades)       │
        │  - attendance (tracking)         │
        │  - assignments (coursework)      │
        │  - notifications (system)        │
        │  ... (20+ tables total)          │
        └─────────────────────────────────┘
```

## Request Flow

### 1. Unauthenticated Request (Login)

```
User Browser
    │
    │ 1. User enters credentials
    │ 2. Clicks "Login"
    │
    ▼
login.html (JavaScript)
    │
    │ 3. Calls: window.api.login(email, password)
    │
    ▼
api.js (API Client)
    │
    │ 4. POST /api/auth/login
    │    Content-Type: application/json
    │    { "email": "admin@school.com", "password": "admin" }
    │
    ▼
Backend Server (Port 8080)
    │
    │ 5. Router matches POST /api/auth/login
    │
    ▼
CORS Middleware
    │ 6. Checks: Is origin allowed?
    │    ✓ http://localhost:3000 is allowed
    │
    ▼
Auth Middleware
    │ 7. This is public endpoint, no token needed
    │
    ▼
AuthHandler.Login()
    │
    │ 8. Validates email/password
    │
    ▼
AuthService.Login()
    │
    │ 9. Generates JWT token
    │
    ▼
Database Query
    │
    │ 10. SELECT * FROM users WHERE email = ?
    │
    ▼
Response to Client
    │
    │ 11. 200 OK
    │     {
    │       "data": {
    │         "token": "eyJhbGc...",
    │         "user": { "id": 1, "email": "admin@school.com", ... }
    │       }
    │     }
    │
    ▼
api.js Receives Response
    │
    │ 12. Stores token in localStorage.sms_token
    │
    ▼
Dashboard
    │
    │ 13. Redirect to dashboard.html
    │
    ▼
User Logged In ✅
```

### 2. Authenticated Request (Get Courses)

```
User Browser
    │
    │ 1. Clicks "View Courses"
    │
    ▼
dashboard.html
    │
    │ 2. Calls: window.api.getCourses(1, 50)
    │
    ▼
api.js (API Client)
    │
    │ 3. GET /api/courses?page=1&limit=50
    │
    │ 4. Headers:
    │    Content-Type: application/json
    │    Authorization: Bearer eyJhbGc...  ◄─── Token from localStorage
    │
    ▼
Backend Server (Port 8080)
    │
    │ 5. Router matches GET /api/courses
    │
    ▼
CORS Middleware
    │ 6. ✓ Origin allowed
    │
    ▼
Auth Middleware
    │
    │ 7. Extract token from Authorization header
    │    Authorization: Bearer eyJhbGc...
    │
    │ 8. Validate JWT signature
    │    ✓ Token valid, expires: 2026-01-31
    │
    │ 9. Extract user info from token
    │    user_id: 1, role: "admin"
    │
    │ 10. Attach to context (c.Set("user_id", 1))
    │
    ▼
CourseHandler.GetAllCourses()
    │
    │ 11. Extract pagination params (page=1, limit=50)
    │
    ▼
CourseService.GetAllCourses()
    │
    │ 12. Calculate offset (page-1)*limit
    │
    ▼
Repository.GetAllCourses()
    │
    │ 13. Database Query:
    │     SELECT * FROM courses
    │     LIMIT 50 OFFSET 0
    │     ORDER BY created_at DESC
    │
    ▼
Database (PostgreSQL)
    │
    │ 14. Execute query
    │     Returns: [
    │       { id: 1, title: "Math 101", description: "...", ... },
    │       { id: 2, title: "Physics 101", description: "...", ... },
    │       ...
    │     ]
    │
    ▼
Response to Client
    │
    │ 15. 200 OK
    │     {
    │       "data": [
    │         { "id": 1, "title": "Math 101", ... },
    │         { "id": 2, "title": "Physics 101", ... }
    │       ]
    │     }
    │
    ▼
api.js Receives Response
    │
    │ 16. Parse JSON response
    │
    ▼
dashboard.js
    │
    │ 17. Render courses in HTML
    │     <div class="course-item">...</div>
    │
    ▼
User Sees Courses ✅
```

## Error Handling Flow

```
Invalid Token / Unauthorized

    User Browser
        │
        │ Makes request with expired token
        │
        ▼
    Backend Auth Middleware
        │
        │ Token validation fails
        │
        ▼
    Response:
        401 Unauthorized
        {
          "error": "Invalid or expired token"
        }
    │
    ▼
    api.js Error Handler
        │
        │ Catches error
        │ Status code: 401
        │
        ▼
    Throw Error to Frontend
        │
        │ catch (err) {
        │   showToast(err.message, 'error')
        │   // Optionally redirect to login
        │ }
        │
        ▼
    User sees error message
    and can login again ✅
```

## Component Communication

```
┌─────────────────────────────────────────┐
│           Frontend (Browser)             │
│                                         │
│  ┌──────────────────────────────────┐  │
│  │  login.html / dashboard.html     │  │
│  │  (User Interface)                │  │
│  │                                  │  │
│  │  DOM Events:                     │  │
│  │  - button.onclick                │  │
│  │  - form.onsubmit                 │  │
│  └────────────┬─────────────────────┘  │
│               │                        │
│               │ JavaScript             │
│               │ Function Calls         │
│               │                        │
│  ┌────────────▼─────────────────────┐  │
│  │  api.js                          │  │
│  │  (API Client Layer)              │  │
│  │                                  │  │
│  │  window.api.login()              │  │
│  │  window.api.getCourses()         │  │
│  │  window.api.getProfile()         │  │
│  │  etc.                            │  │
│  │                                  │  │
│  │  Features:                       │  │
│  │  - Token management              │  │
│  │  - HTTP requests                 │  │
│  │  - Error handling                │  │
│  │  - Response parsing              │  │
│  └────────────┬─────────────────────┘  │
└───────────────┼───────────────────────┘
                │
                │ HTTP Requests
                │ JSON Data
                │ JWT Token
                │
┌───────────────▼───────────────────────┐
│        Backend (Go Server)             │
│        (http://localhost:8080)         │
│                                       │
│  ┌──────────────────────────────────┐ │
│  │  Route Matching                  │ │
│  │  /api/auth/login                 │ │
│  │  /api/courses                    │ │
│  │  /api/profile                    │ │
│  └────────────┬─────────────────────┘ │
│               │                       │
│  ┌────────────▼─────────────────────┐ │
│  │  Middleware Chain                │ │
│  │  1. CORS Check                   │ │
│  │  2. Auth Validation              │ │
│  │  3. Request Validation           │ │
│  │  4. Rate Limiting                │ │
│  └────────────┬─────────────────────┘ │
│               │                       │
│  ┌────────────▼─────────────────────┐ │
│  │  Handlers (25+)                  │ │
│  │  - AuthHandler                   │ │
│  │  - UserHandler                   │ │
│  │  - CourseHandler                 │ │
│  │  - etc.                          │ │
│  │                                  │ │
│  │  Responsibility:                 │ │
│  │  - Parse request                 │ │
│  │  - Validate input                │ │
│  │  - Call service layer            │ │
│  │  - Return response               │ │
│  └────────────┬─────────────────────┘ │
│               │                       │
│  ┌────────────▼─────────────────────┐ │
│  │  Service Layer                   │ │
│  │  (Business Logic)                │ │
│  │                                  │ │
│  │  Responsibility:                 │ │
│  │  - Implement business rules      │ │
│  │  - Orchestrate operations        │ │
│  │  - Handle complex logic          │ │
│  └────────────┬─────────────────────┘ │
│               │                       │
│  ┌────────────▼─────────────────────┐ │
│  │  Repository Layer                │ │
│  │  (Data Access)                   │ │
│  │                                  │ │
│  │  Responsibility:                 │ │
│  │  - Database queries              │ │
│  │  - CRUD operations               │ │
│  │  - Query optimization            │ │
│  └────────────┬─────────────────────┘ │
└────────────────┼──────────────────────┘
                 │
                 │ SQL Queries
                 │ Connection Pool
                 │
           ┌─────▼─────┐
           │ PostgreSQL │
           │ Database   │
           └────────────┘
```

## Technology Stack

### Frontend
```
│
├─ Framework: React 18.2
├─ Build Tool: Vite
├─ HTTP Client: Fetch API (native)
├─ Storage: localStorage
├─ Language: JavaScript
└─ Styling: CSS3
```

### Backend
```
│
├─ Language: Go 1.21+
├─ Web Framework: Gin-gonic
├─ Database: PostgreSQL (configurable)
├─ Authentication: JWT (JSON Web Tokens)
├─ Validation: Custom middleware
├─ Logging: Logrus
└─ ORM: GORM
```

### Infrastructure
```
│
├─ Frontend Server: Port 3000 (HTTP)
├─ Backend Server: Port 8080 (HTTP)
├─ Database: Default (configured in .env)
├─ Protocol: HTTP/REST
└─ Authentication: JWT Bearer Token
```

---

**Architecture Version**: 1.0
**Last Updated**: January 30, 2026
**Status**: ✅ Complete and Integrated
