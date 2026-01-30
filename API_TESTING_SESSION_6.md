# API Testing Guide - Session 6 Advanced Features

This guide covers all 26 new API endpoints added in Session 6 for advanced customization features.

**Base URL:** `http://localhost:8080/api`

---

## Table of Contents

1. [Advanced Search Endpoints](#advanced-search-endpoints)
2. [CSV Export Endpoints](#csv-export-endpoints)
3. [Attendance Automation Endpoints](#attendance-automation-endpoints)
4. [Grade Auto-Calculation Endpoints](#grade-auto-calculation-endpoints)
5. [Rubrics Endpoints](#rubrics-endpoints)

---

## Advanced Search Endpoints

### 1. Search Announcements
**Endpoint:** `GET /api/search/announcements`

**Description:** Search announcements with filters for title/content, audience, and priority.

**Query Parameters:**
- `query` (optional): Search term for title or content
- `audience` (optional): Filter by audience (e.g., "students", "teachers", "all")
- `priority` (optional): Filter by priority (e.g., "high", "medium", "low")
- `page` (optional): Page number (default: 1)
- `limit` (optional): Results per page (default: 10)

**Example Request:**
```bash
curl -X GET "http://localhost:8080/api/search/announcements?query=exam&audience=students&priority=high&page=1&limit=10" \
  -H "Authorization: Bearer YOUR_TOKEN"
```

**Expected Response:**
```json
{
  "success": true,
  "data": {
    "announcements": [
      {
        "id": 1,
        "title": "Final Exam Schedule",
        "content": "Exam dates have been announced...",
        "audience": "students",
        "priority": "high",
        "is_active": true,
        "created_at": 1706486400
      }
    ],
    "total": 1,
    "page": 1,
    "limit": 10
  }
}
```

---

### 2. Search Payments
**Endpoint:** `GET /api/search/payments`

**Description:** Search payments by student ID and status.

**Query Parameters:**
- `student_id` (optional): Filter by student
- `status` (optional): Filter by status (pending, paid, overdue, cancelled)
- `page` (optional): Page number (default: 1)
- `limit` (optional): Results per page (default: 10)

**Example Request:**
```bash
curl -X GET "http://localhost:8080/api/search/payments?student_id=5&status=pending&page=1&limit=10" \
  -H "Authorization: Bearer YOUR_TOKEN"
```

**Expected Response:**
```json
{
  "success": true,
  "data": {
    "payments": [
      {
        "id": 1,
        "student_id": 5,
        "amount": 5000.00,
        "status": "pending",
        "description": "Tuition Fee",
        "due_date": 1706572800,
        "created_at": 1706486400
      }
    ],
    "total": 1,
    "page": 1,
    "limit": 10
  }
}
```

---

### 3. Search Students
**Endpoint:** `GET /api/search/students`

**Description:** Search students by name, email, or student ID.

**Query Parameters:**
- `query` (required): Search term (name, email, or student_id)
- `page` (optional): Page number (default: 1)
- `limit` (optional): Results per page (default: 10)

**Example Request:**
```bash
curl -X GET "http://localhost:8080/api/search/students?query=John&page=1&limit=10" \
  -H "Authorization: Bearer YOUR_TOKEN"
```

**Expected Response:**
```json
{
  "success": true,
  "data": {
    "students": [
      {
        "id": 1,
        "user_id": 2,
        "student_id": "STU001",
        "grade_level": "10",
        "enrollment_date": "2024-01-15T00:00:00Z",
        "user": {
          "id": 2,
          "first_name": "John",
          "last_name": "Doe",
          "email": "john@school.com"
        }
      }
    ],
    "total": 1,
    "page": 1,
    "limit": 10
  }
}
```

---

### 4. Search Grades by Range
**Endpoint:** `GET /api/search/grades`

**Description:** Search grades within a score range.

**Query Parameters:**
- `course_id` (required): Course ID
- `min_score` (optional): Minimum score (default: 0)
- `max_score` (optional): Maximum score (default: 100)
- `page` (optional): Page number (default: 1)
- `limit` (optional): Results per page (default: 10)

**Example Request:**
```bash
curl -X GET "http://localhost:8080/api/search/grades?course_id=3&min_score=80&max_score=100&page=1&limit=10" \
  -H "Authorization: Bearer YOUR_TOKEN"
```

**Expected Response:**
```json
{
  "success": true,
  "data": {
    "grades": [
      {
        "id": 5,
        "student_id": 1,
        "course_id": 3,
        "score": 92.5,
        "grade": "A",
        "graded_at": "2024-01-20T10:30:00Z"
      }
    ],
    "total": 1,
    "page": 1,
    "limit": 10
  }
}
```

---

### 5. Search Overdue Payments
**Endpoint:** `GET /api/search/overdue-payments`

**Description:** Find all overdue payments.

**Query Parameters:** None

**Example Request:**
```bash
curl -X GET "http://localhost:8080/api/search/overdue-payments" \
  -H "Authorization: Bearer YOUR_TOKEN"
```

**Expected Response:**
```json
{
  "success": true,
  "data": {
    "payments": [
      {
        "id": 2,
        "student_id": 4,
        "amount": 3000.00,
        "status": "pending",
        "description": "Lab Fee",
        "due_date": 1706486400,
        "created_at": 1706400000
      }
    ],
    "total": 1
  }
}
```

---

## CSV Export Endpoints

### 6. Export Payments CSV
**Endpoint:** `GET /api/export/payments`

**Description:** Export payments as CSV file.

**Query Parameters:**
- `student_id` (optional): Filter by student
- `status` (optional): Filter by status

**Example Request:**
```bash
curl -X GET "http://localhost:8080/api/export/payments?status=paid" \
  -H "Authorization: Bearer YOUR_TOKEN" \
  -o payments.csv
```

**Response:** CSV file with columns: ID, Student ID, Amount, Status, Description, Due Date, Created

---

### 7. Export Grades CSV
**Endpoint:** `GET /api/export/grades`

**Description:** Export grades as CSV file.

**Query Parameters:**
- `course_id` (required): Course ID

**Example Request:**
```bash
curl -X GET "http://localhost:8080/api/export/grades?course_id=3" \
  -H "Authorization: Bearer YOUR_TOKEN" \
  -o grades.csv
```

**Response:** CSV file with columns: ID, Student ID, Course ID, Score, Grade, Graded At

---

### 8. Export Attendance CSV
**Endpoint:** `GET /api/export/attendance`

**Description:** Export attendance records as CSV file.

**Query Parameters:**
- `course_id` (required): Course ID

**Example Request:**
```bash
curl -X GET "http://localhost:8080/api/export/attendance?course_id=3" \
  -H "Authorization: Bearer YOUR_TOKEN" \
  -o attendance.csv
```

**Response:** CSV file with columns: ID, Student ID, Course ID, Date, Present, Remarks

---

### 9. Export Student Transcript CSV
**Endpoint:** `GET /api/export/transcript/:student_id`

**Description:** Export student academic transcript as CSV file.

**Path Parameters:**
- `student_id` (required): Student ID

**Example Request:**
```bash
curl -X GET "http://localhost:8080/api/export/transcript/1" \
  -H "Authorization: Bearer YOUR_TOKEN" \
  -o transcript.csv
```

**Response:** CSV file with columns: Semester, Year, GPA, Total Credits, Earned Credits, Generated At

---

### 10. Export Enrollments CSV
**Endpoint:** `GET /api/export/enrollments`

**Description:** Export course enrollments as CSV file.

**Query Parameters:**
- `course_id` (required): Course ID

**Example Request:**
```bash
curl -X GET "http://localhost:8080/api/export/enrollments?course_id=3" \
  -H "Authorization: Bearer YOUR_TOKEN" \
  -o enrollments.csv
```

**Response:** CSV file with columns: ID, Student ID, Course ID, Status, Enrolled At

---

## Attendance Automation Endpoints

### 11. Get Attendance Stats by Course
**Endpoint:** `GET /api/attendance/stats/course/:course_id`

**Description:** Get comprehensive attendance statistics for a course.

**Path Parameters:**
- `course_id` (required): Course ID

**Example Request:**
```bash
curl -X GET "http://localhost:8080/api/attendance/stats/course/3" \
  -H "Authorization: Bearer YOUR_TOKEN"
```

**Expected Response:**
```json
{
  "success": true,
  "data": {
    "total_sessions": 20,
    "total_students": 30,
    "average_attendance": 85.5,
    "last_updated_at": "2024-01-20T15:30:00Z"
  }
}
```

---

### 12. Get Student Attendance Percentage
**Endpoint:** `GET /api/attendance/percentage/:student_id/:course_id`

**Description:** Calculate attendance percentage for a student in a course.

**Path Parameters:**
- `student_id` (required): Student ID
- `course_id` (required): Course ID

**Example Request:**
```bash
curl -X GET "http://localhost:8080/api/attendance/percentage/1/3" \
  -H "Authorization: Bearer YOUR_TOKEN"
```

**Expected Response:**
```json
{
  "success": true,
  "data": {
    "student_id": 1,
    "course_id": 3,
    "attendance_percentage": 92.5,
    "classes_attended": 37,
    "total_classes": 40
  }
}
```

---

### 13. Check Low Attendance
**Endpoint:** `POST /api/attendance/check-low`

**Description:** Check if student attendance is below threshold and trigger alerts.

**Request Body:**
```json
{
  "student_id": 1,
  "course_id": 3,
  "threshold": 75
}
```

**Example Request:**
```bash
curl -X POST "http://localhost:8080/api/attendance/check-low" \
  -H "Authorization: Bearer YOUR_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "student_id": 1,
    "course_id": 3,
    "threshold": 75
  }'
```

**Expected Response:**
```json
{
  "success": true,
  "data": {
    "below_threshold": false,
    "current_attendance": 92.5,
    "threshold": 75,
    "alert_sent": false
  }
}
```

---

### 14. Get Students with Low Attendance
**Endpoint:** `GET /api/attendance/low/:threshold`

**Description:** Get all students with attendance below threshold.

**Path Parameters:**
- `threshold` (required): Attendance threshold (0-100)

**Query Parameters:**
- `course_id` (optional): Filter by course

**Example Request:**
```bash
curl -X GET "http://localhost:8080/api/attendance/low/75?course_id=3" \
  -H "Authorization: Bearer YOUR_TOKEN"
```

**Expected Response:**
```json
{
  "success": true,
  "data": {
    "threshold": 75,
    "course_id": 3,
    "students": [
      {
        "student_id": 2,
        "name": "Jane Smith",
        "email": "jane@school.com",
        "attendance_percentage": 60.0
      }
    ],
    "total": 1
  }
}
```

---

### 15. Get Attendance Report
**Endpoint:** `GET /api/attendance/report/:course_id`

**Description:** Generate comprehensive attendance report for a course.

**Path Parameters:**
- `course_id` (required): Course ID

**Example Request:**
```bash
curl -X GET "http://localhost:8080/api/attendance/report/3" \
  -H "Authorization: Bearer YOUR_TOKEN"
```

**Expected Response:**
```json
{
  "success": true,
  "data": {
    "course_id": 3,
    "course_name": "Mathematics",
    "stats": {
      "total_sessions": 20,
      "total_students": 30,
      "average_attendance": 85.5,
      "last_updated_at": "2024-01-20T15:30:00Z"
    },
    "students_below_80pct": [
      {
        "student_id": 2,
        "name": "Jane Smith",
        "email": "jane@school.com",
        "attendance_percentage": 70.0
      }
    ],
    "report_generated_at": "2024-01-20T16:00:00Z"
  }
}
```

---

## Grade Auto-Calculation Endpoints

### 16. Record Grade with Auto-Calculation
**Endpoint:** `POST /api/grades/auto`

**Description:** Record a grade and automatically calculate letter grade, update transcript.

**Request Body:**
```json
{
  "student_id": 1,
  "course_id": 3,
  "score": 92.5,
  "max_score": 100,
  "remarks": "Excellent work",
  "graded_by": 10
}
```

**Example Request:**
```bash
curl -X POST "http://localhost:8080/api/grades/auto" \
  -H "Authorization: Bearer YOUR_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "student_id": 1,
    "course_id": 3,
    "score": 92.5,
    "max_score": 100,
    "remarks": "Excellent work",
    "graded_by": 10
  }'
```

**Expected Response:**
```json
{
  "success": true,
  "data": {
    "id": 8,
    "student_id": 1,
    "course_id": 3,
    "score": 92.5,
    "grade": "A",
    "graded_at": "2024-01-20T16:15:00Z",
    "message": "Grade recorded and auto-calculated. Letter grade: A. Email notification sent."
  }
}
```

---

### 17. Get Course Average Grade
**Endpoint:** `GET /api/grades/course-average/:course_id`

**Description:** Calculate average grade for a course.

**Path Parameters:**
- `course_id` (required): Course ID

**Example Request:**
```bash
curl -X GET "http://localhost:8080/api/grades/course-average/3" \
  -H "Authorization: Bearer YOUR_TOKEN"
```

**Expected Response:**
```json
{
  "success": true,
  "data": {
    "course_id": 3,
    "course_name": "Mathematics",
    "average_score": 82.4,
    "total_grades": 28,
    "highest_score": 98.0,
    "lowest_score": 62.5
  }
}
```

---

### 18. Get Grade Distribution
**Endpoint:** `GET /api/grades/distribution/:course_id`

**Description:** Get grade distribution (A, B, C, D, F counts) for a course.

**Path Parameters:**
- `course_id` (required): Course ID

**Example Request:**
```bash
curl -X GET "http://localhost:8080/api/grades/distribution/3" \
  -H "Authorization: Bearer YOUR_TOKEN"
```

**Expected Response:**
```json
{
  "success": true,
  "data": {
    "course_id": 3,
    "course_name": "Mathematics",
    "distribution": {
      "A": 8,
      "B": 12,
      "C": 5,
      "D": 2,
      "F": 1
    },
    "total_students": 28
  }
}
```

---

### 19. Get Student Grade Statistics
**Endpoint:** `GET /api/grades/student-stats/:student_id`

**Description:** Get comprehensive grade statistics for a student.

**Path Parameters:**
- `student_id` (required): Student ID

**Example Request:**
```bash
curl -X GET "http://localhost:8080/api/grades/student-stats/1" \
  -H "Authorization: Bearer YOUR_TOKEN"
```

**Expected Response:**
```json
{
  "success": true,
  "data": {
    "student_id": 1,
    "grade_count": 6,
    "average": "88.50",
    "a_count": 3,
    "b_count": 2,
    "c_count": 1,
    "d_count": 0,
    "f_count": 0,
    "highest_grade": 95.0,
    "lowest_grade": 80.5
  }
}
```

---

## Rubrics Endpoints

### 20. Create Rubric
**Endpoint:** `POST /api/rubrics`

**Description:** Create a new grading rubric for an assignment.

**Request Body:**
```json
{
  "assignment_id": 5,
  "name": "Essay Grading Rubric",
  "criteria": [
    {
      "name": "Organization",
      "max_points": 25,
      "description": "Logical flow and structure"
    },
    {
      "name": "Content",
      "max_points": 50,
      "description": "Quality and accuracy of information"
    },
    {
      "name": "Grammar",
      "max_points": 25,
      "description": "Spelling and grammar correctness"
    }
  ],
  "is_active": true
}
```

**Example Request:**
```bash
curl -X POST "http://localhost:8080/api/rubrics" \
  -H "Authorization: Bearer YOUR_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "assignment_id": 5,
    "name": "Essay Grading Rubric",
    "criteria": [
      {"name": "Organization", "max_points": 25, "description": "Logical flow"},
      {"name": "Content", "max_points": 50, "description": "Quality of info"},
      {"name": "Grammar", "max_points": 25, "description": "Spelling and grammar"}
    ],
    "is_active": true
  }'
```

**Expected Response:**
```json
{
  "success": true,
  "data": {
    "id": 1,
    "assignment_id": 5,
    "name": "Essay Grading Rubric",
    "criteria": [...],
    "is_active": true,
    "created_at": "2024-01-20T17:00:00Z"
  }
}
```

---

### 21. Get Rubric by ID
**Endpoint:** `GET /api/rubrics/:id`

**Description:** Get a specific rubric.

**Path Parameters:**
- `id` (required): Rubric ID

**Example Request:**
```bash
curl -X GET "http://localhost:8080/api/rubrics/1" \
  -H "Authorization: Bearer YOUR_TOKEN"
```

**Expected Response:**
```json
{
  "success": true,
  "data": {
    "id": 1,
    "assignment_id": 5,
    "name": "Essay Grading Rubric",
    "criteria": [
      {
        "name": "Organization",
        "max_points": 25,
        "description": "Logical flow and structure"
      },
      {
        "name": "Content",
        "max_points": 50,
        "description": "Quality and accuracy of information"
      },
      {
        "name": "Grammar",
        "max_points": 25,
        "description": "Spelling and grammar correctness"
      }
    ],
    "is_active": true,
    "created_at": "2024-01-20T17:00:00Z"
  }
}
```

---

### 22. Get Rubrics by Assignment
**Endpoint:** `GET /api/rubrics/assignment/:assignment_id`

**Description:** Get all rubrics for an assignment.

**Path Parameters:**
- `assignment_id` (required): Assignment ID

**Example Request:**
```bash
curl -X GET "http://localhost:8080/api/rubrics/assignment/5" \
  -H "Authorization: Bearer YOUR_TOKEN"
```

**Expected Response:**
```json
{
  "success": true,
  "data": {
    "assignment_id": 5,
    "rubrics": [
      {
        "id": 1,
        "name": "Essay Grading Rubric",
        "criteria": [...],
        "is_active": true
      }
    ],
    "total": 1
  }
}
```

---

### 23. Update Rubric
**Endpoint:** `PUT /api/rubrics/:id`

**Description:** Update a rubric.

**Path Parameters:**
- `id` (required): Rubric ID

**Request Body:**
```json
{
  "name": "Essay Grading Rubric (Updated)",
  "criteria": [...],
  "is_active": true
}
```

**Example Request:**
```bash
curl -X PUT "http://localhost:8080/api/rubrics/1" \
  -H "Authorization: Bearer YOUR_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Essay Grading Rubric (Updated)",
    "is_active": true
  }'
```

**Expected Response:**
```json
{
  "success": true,
  "data": {
    "id": 1,
    "assignment_id": 5,
    "name": "Essay Grading Rubric (Updated)",
    "updated_at": "2024-01-20T17:30:00Z"
  }
}
```

---

### 24. Delete Rubric
**Endpoint:** `DELETE /api/rubrics/:id`

**Description:** Delete a rubric.

**Path Parameters:**
- `id` (required): Rubric ID

**Example Request:**
```bash
curl -X DELETE "http://localhost:8080/api/rubrics/1" \
  -H "Authorization: Bearer YOUR_TOKEN"
```

**Expected Response:**
```json
{
  "success": true,
  "message": "Rubric deleted successfully"
}
```

---

### 25. Score Submission with Rubric
**Endpoint:** `POST /api/rubrics/score`

**Description:** Grade a submission using a rubric.

**Request Body:**
```json
{
  "submission_id": 10,
  "rubric_id": 1,
  "criterion_scores": [
    {
      "criterion_name": "Organization",
      "points_earned": 23
    },
    {
      "criterion_name": "Content",
      "points_earned": 48
    },
    {
      "criterion_name": "Grammar",
      "points_earned": 24
    }
  ],
  "feedback": "Excellent essay with minor grammar issues.",
  "graded_by": 10
}
```

**Example Request:**
```bash
curl -X POST "http://localhost:8080/api/rubrics/score" \
  -H "Authorization: Bearer YOUR_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "submission_id": 10,
    "rubric_id": 1,
    "criterion_scores": [
      {"criterion_name": "Organization", "points_earned": 23},
      {"criterion_name": "Content", "points_earned": 48},
      {"criterion_name": "Grammar", "points_earned": 24}
    ],
    "feedback": "Excellent essay with minor grammar issues.",
    "graded_by": 10
  }'
```

**Expected Response:**
```json
{
  "success": true,
  "data": {
    "id": 1,
    "submission_id": 10,
    "rubric_id": 1,
    "total_points_earned": 95,
    "total_points_available": 100,
    "percentage": 95.0,
    "letter_grade": "A",
    "feedback": "Excellent essay with minor grammar issues.",
    "graded_at": "2024-01-20T17:45:00Z"
  }
}
```

---

### 26. Get Submission Score
**Endpoint:** `GET /api/rubrics/score/:submission_id`

**Description:** Get rubric score for a submission.

**Path Parameters:**
- `submission_id` (required): Submission ID

**Example Request:**
```bash
curl -X GET "http://localhost:8080/api/rubrics/score/10" \
  -H "Authorization: Bearer YOUR_TOKEN"
```

**Expected Response:**
```json
{
  "success": true,
  "data": {
    "id": 1,
    "submission_id": 10,
    "rubric_id": 1,
    "total_points_earned": 95,
    "total_points_available": 100,
    "percentage": 95.0,
    "letter_grade": "A",
    "feedback": "Excellent essay with minor grammar issues.",
    "criterion_scores": [
      {
        "criterion_name": "Organization",
        "points_earned": 23,
        "max_points": 25
      },
      {
        "criterion_name": "Content",
        "points_earned": 48,
        "max_points": 50
      },
      {
        "criterion_name": "Grammar",
        "points_earned": 24,
        "max_points": 25
      }
    ],
    "graded_at": "2024-01-20T17:45:00Z"
  }
}
```

---

## Testing Checklist

Use this checklist to verify all endpoints:

### Search Endpoints (5)
- [ ] Search Announcements with filters
- [ ] Search Payments by student/status
- [ ] Search Students by name/email
- [ ] Search Grades by score range
- [ ] Search Overdue Payments

### Export Endpoints (5)
- [ ] Export Payments CSV
- [ ] Export Grades CSV
- [ ] Export Attendance CSV
- [ ] Export Student Transcript CSV
- [ ] Export Enrollments CSV

### Attendance Automation (5)
- [ ] Get Course Attendance Stats
- [ ] Get Student Attendance Percentage
- [ ] Check Low Attendance (POST)
- [ ] Get Students with Low Attendance
- [ ] Get Attendance Report

### Grade Auto-Calculation (4)
- [ ] Record Grade with Auto-Calculation
- [ ] Get Course Average Grade
- [ ] Get Grade Distribution
- [ ] Get Student Grade Statistics

### Rubrics (7)
- [ ] Create Rubric
- [ ] Get Rubric by ID
- [ ] Get Rubrics by Assignment
- [ ] Update Rubric
- [ ] Delete Rubric
- [ ] Score Submission with Rubric
- [ ] Get Submission Score

---

## Authentication

All endpoints require a valid JWT token in the `Authorization` header:
```
Authorization: Bearer <YOUR_JWT_TOKEN>
```

To get a token:
```bash
curl -X POST "http://localhost:8080/api/auth/login" \
  -H "Content-Type: application/json" \
  -d '{
    "email": "admin@school.com",
    "password": "admin123"
  }'
```

---

## Error Handling

All endpoints return consistent error responses:

```json
{
  "success": false,
  "error": "Error message describing what went wrong",
  "code": "ERROR_CODE"
}
```

Common HTTP Status Codes:
- `200` - OK (successful GET)
- `201` - Created (successful POST)
- `400` - Bad Request (invalid parameters)
- `401` - Unauthorized (missing/invalid token)
- `403` - Forbidden (insufficient permissions)
- `404` - Not Found (resource doesn't exist)
- `500` - Server Error

---

## Performance Notes

- **Search endpoints:** Support pagination (page/limit) for large result sets
- **Export endpoints:** Stream CSV data directly to client for memory efficiency
- **Attendance/Grade endpoints:** Cache statistics for better performance
- **Rubrics:** Support flexible JSON criteria storage

---

## Next Steps

1. Start the server: `go run ./cmd/server`
2. Test endpoints with provided curl examples
3. Monitor logs for any errors
4. Integrate with frontend using API methods
5. Run `go test ./...` to verify all tests pass

