# School Management System - Complete API Reference

## 🔌 Frontend-Accessible API Endpoints

All endpoints are available through the `api.js` client and are fully integrated into the frontend pages.

---

## 🔐 Authentication Endpoints

### POST /api/login
```javascript
window.api.login(email, password)
```
**Used In**: login.html
**Response**: JWT token + user data
**Stores**: Token in localStorage.sms_token

### POST /api/register
```javascript
window.api.register({
  first_name: "John",
  last_name: "Doe",
  email: "john@example.com",
  password: "password",
  role: "student"
})
```
**Used In**: register.html
**Response**: New user created

---

## 👤 Profile Endpoints

### GET /api/profile
```javascript
window.api.getProfile()
```
**Used In**: home.html, all pages
**Response**: Current user profile
**Requires**: Authentication token

---

## 📚 Course Endpoints

### GET /api/courses
```javascript
window.api.getCourses(page, limit)
```
**Used In**: dashboard.html, home.html
**Response**: List of courses
**Pagination**: Supported

### POST /api/enrollments
```javascript
window.api.createEnrollment(studentId, courseId)
```
**Used In**: dashboard.html
**Response**: Enrollment created

---

## 📊 Grades Endpoints

### GET /api/grades/me
```javascript
window.api.getMyGrades()
```
**Used In**: grades.html, transcripts.html, home.html
**Response**: Student's grades

### POST /api/grades
```javascript
window.api.recordGrade({
  student_id: 1,
  course_id: 2,
  grade: 85
})
```
**Used In**: grades.html
**Response**: Grade recorded
**Requires**: Teacher/Admin role

### GET /api/grades/student/:id
```javascript
window.api.getGradesByStudent(studentId)
```
**Used In**: Admin dashboard
**Response**: Student's grade history

---

## 📍 Attendance Endpoints

### GET /api/attendance/me
```javascript
window.api.getMyAttendance()
```
**Used In**: attendance.html, transcripts.html
**Response**: Student's attendance records

### POST /api/attendance
```javascript
window.api.markAttendance({
  student_id: 1,
  course_id: 2,
  status: "present"
})
```
**Used In**: attendance.html
**Response**: Attendance marked
**Status Options**: present, absent, late, excused

### GET /api/attendance/student/:id
```javascript
window.api.getAttendanceByStudent(studentId)
```
**Used In**: Admin dashboard
**Response**: Student's attendance history

---

## 📝 Assignment Endpoints

### GET /api/assignments/me
```javascript
window.api.getMyAssignments()
```
**Used In**: assignments.html, home.html
**Response**: Student's assignments

### POST /api/assignments
```javascript
window.api.createAssignment({
  title: "Project 1",
  description: "Build a calculator",
  course_id: 1,
  due_date: "2024-12-31"
})
```
**Used In**: assignments.html
**Response**: Assignment created
**Requires**: Teacher/Admin role

### POST /api/assignments/:id/submit
```javascript
window.api.submitAssignment({
  assignment_id: 1,
  submission: "Project code or URL"
})
```
**Used In**: assignments.html
**Response**: Submission recorded

### POST /api/assignments/:id/grade
```javascript
window.api.gradeAssignment(assignmentId, grade)
```
**Used In**: assignments.html
**Response**: Grade recorded

---

## 📢 Announcement Endpoints

### GET /api/announcements
```javascript
window.api.getAnnouncements()
```
**Used In**: announcements.html, home.html
**Response**: List of announcements

### POST /api/announcements
```javascript
window.api.createAnnouncement({
  title: "Holiday Notice",
  content: "School closed on...",
  priority: "high",
  target_audience: "all"
})
```
**Used In**: announcements.html
**Response**: Announcement created
**Requires**: Teacher/Admin role
**Priority**: low, medium, high
**Target**: all, students, teachers, staff, admin

### DELETE /api/announcements/:id
```javascript
window.api.deleteAnnouncement(id)
```
**Used In**: announcements.html
**Response**: Announcement deleted

---

## 🔔 Notification Endpoints

### GET /api/notifications
```javascript
window.api.getNotifications()
```
**Used In**: notifications.html, home.html
**Response**: List of notifications

### POST /api/notifications/:id/read
```javascript
window.api.markNotificationAsRead(notificationId)
```
**Used In**: notifications.html
**Response**: Marked as read

### DELETE /api/notifications/:id
```javascript
window.api.deleteNotification(notificationId)
```
**Used In**: notifications.html
**Response**: Notification deleted

---

## 💬 Message Endpoints

### GET /api/conversations
```javascript
window.api.getConversations()
```
**Used In**: messages.html
**Response**: List of user conversations

### POST /api/conversations
```javascript
window.api.createConversation({
  recipient_email: "user@example.com"
})
```
**Used In**: messages.html
**Response**: Conversation created

### GET /api/messages/:conversationId
```javascript
window.api.getMessages(conversationId)
```
**Used In**: messages.html
**Response**: Messages in conversation

### POST /api/messages
```javascript
window.api.sendMessage({
  conversation_id: 1,
  content: "Hello!"
})
```
**Used In**: messages.html
**Response**: Message sent

---

## 💳 Payment Endpoints

### GET /api/payments
```javascript
window.api.getPayments()
```
**Used In**: payments.html
**Response**: Payment history and balance

### POST /api/payments
```javascript
window.api.makePayment({
  amount: 100.00,
  method: "credit_card"
})
```
**Used In**: payments.html
**Response**: Payment submitted
**Methods**: credit_card, bank_transfer, check, cash

---

## 📜 Transcript Endpoints

### GET /api/transcripts/me
```javascript
window.api.getTranscript()
```
**Used In**: transcripts.html
**Response**: Complete academic transcript

---

## 📅 Timetable Endpoints

### GET /api/timetables
```javascript
window.api.getTimetable()
```
**Used In**: timetables.html
**Response**: Class schedule

### POST /api/timetables
```javascript
window.api.createTimetable({
  course_id: 1,
  day: "monday",
  start_time: "09:00",
  end_time: "10:30",
  room: "101"
})
```
**Used In**: timetables.html
**Response**: Schedule created
**Requires**: Teacher/Admin role

---

## 👥 User/Enrollment Endpoints

### GET /api/students
```javascript
window.api.getAllStudents()
```
**Used In**: admin-dashboard.html
**Response**: List of students
**Requires**: Admin role

### GET /api/enrollments
```javascript
window.api.getAllEnrollments()
```
**Used In**: admin-dashboard.html
**Response**: All enrollments

### GET /api/enrollments/:studentId
```javascript
window.api.getEnrollmentsByStudent(studentId)
```
**Used In**: dashboard.html
**Response**: Student's enrollments

### POST /api/enrollments/:id/approve
```javascript
window.api.approveEnrollment(enrollmentId)
```
**Used In**: admin-dashboard.html
**Response**: Enrollment approved
**Requires**: Admin role

### POST /api/enrollments/:id/reject
```javascript
window.api.rejectEnrollment(enrollmentId)
```
**Used In**: admin-dashboard.html
**Response**: Enrollment rejected
**Requires**: Admin role

---

## ⚙️ Admin Endpoints

### GET /api/admin/health
```javascript
window.api.getAdminHealth()
```
**Used In**: admin-dashboard.html
**Response**: System statistics
```json
{
  "total_users": 100,
  "total_students": 80,
  "total_teachers": 20,
  "total_courses": 15,
  "pending_enrollments": 5
}
```

### GET /api/admin/users
```javascript
window.api.getAllUsers()
```
**Used In**: admin-dashboard.html
**Response**: List of all users
**Requires**: Admin role

---

## 🔗 Complete API Integration Map

### Dashboard (home.html)
- ✅ getProfile() → Display user info
- ✅ getCourses() → Display course count
- ✅ getMyAssignments() → Count pending
- ✅ getMyGrades() → Calculate GPA
- ✅ getNotifications() → Count unread

### Grades (grades.html)
- ✅ getMyGrades() → Display grades
- ✅ recordGrade() → Teacher/Admin add grades

### Attendance (attendance.html)
- ✅ getMyAttendance() → Display attendance
- ✅ markAttendance() → Teacher mark attendance

### Assignments (assignments.html)
- ✅ getMyAssignments() → Display assignments
- ✅ createAssignment() → Teacher create
- ✅ submitAssignment() → Student submit

### Announcements (announcements.html)
- ✅ getAnnouncements() → Display all
- ✅ createAnnouncement() → Post new
- ✅ deleteAnnouncement() → Remove

### Notifications (notifications.html)
- ✅ getNotifications() → Display all
- ✅ markNotificationAsRead() → Mark read
- ✅ deleteNotification() → Remove

### Messages (messages.html)
- ✅ getConversations() → Display chats
- ✅ createConversation() → Start new
- ✅ getMessages() → Load history
- ✅ sendMessage() → Send message

### Payments (payments.html)
- ✅ getPayments() → Display history
- ✅ makePayment() → Submit payment

### Transcripts (transcripts.html)
- ✅ getProfile() → Student info
- ✅ getMyGrades() → Grade history

### Timetables (timetables.html)
- ✅ getTimetable() → Display schedule
- ✅ createTimetable() → Add class
- ✅ getCourses() → Load courses

### Admin Dashboard (admin-dashboard.html)
- ✅ getAdminHealth() → Display stats
- ✅ getAllEnrollments() → Enrollment list
- ✅ approveEnrollment() → Approve
- ✅ rejectEnrollment() → Reject
- ✅ getAllUsers() → User list

---

## 📊 API Statistics

| Category | Count |
|----------|-------|
| Total Endpoints | 25+ |
| Authentication | 2 |
| Profile | 1 |
| Courses | 2 |
| Grades | 3 |
| Attendance | 3 |
| Assignments | 4 |
| Announcements | 3 |
| Notifications | 3 |
| Messages | 4 |
| Payments | 2 |
| Transcripts | 1 |
| Timetables | 2 |
| Users/Enrollments | 5 |
| Admin | 2 |

---

## 🔐 Authentication

All API endpoints (except login/register) require:
- **Header**: `Authorization: Bearer [token]`
- **Token Storage**: `localStorage.sms_token`
- **Token Expiration**: Configurable
- **Auto-Refresh**: Via api.js client

---

## 🎯 Frontend Page to API Mapping

```
login.html              → POST /api/login
register.html           → POST /api/register
home.html              → GET /api/profile, courses, assignments, grades, notifications
dashboard.html         → GET /api/courses, enrollments; POST /api/enrollments
grades.html            → GET /api/grades/me; POST /api/grades
attendance.html        → GET /api/attendance/me; POST /api/attendance
assignments.html       → GET /api/assignments/me; POST /api/assignments, /submit
announcements.html     → GET /api/announcements; POST /api/announcements; DELETE
notifications.html     → GET /api/notifications; POST /read; DELETE
messages.html          → GET /api/conversations, /messages; POST /api/conversations, /messages
payments.html          → GET /api/payments; POST /api/payments
transcripts.html       → GET /api/transcripts/me, /api/grades/me
timetables.html        → GET /api/timetables; POST /api/timetables; GET /api/courses
admin-dashboard.html   → GET /api/admin/health, /enrollments, /users; POST /approve, /reject
```

---

## ✅ All Features Connected

Every frontend page is fully connected to the backend API with:
- ✅ Automatic token handling
- ✅ Error management
- ✅ Loading states
- ✅ Success/failure feedback
- ✅ Data persistence
- ✅ Real-time updates (where applicable)

---

**API Version**: 1.0.0  
**Status**: ✅ Fully Operational  
**Authentication**: JWT Bearer Tokens  
**Base URL**: http://localhost:8080  
**Frontend URL**: http://localhost:3001  
