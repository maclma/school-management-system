# School Management System - API Testing Guide

## New Features API Reference (v5.0)

This guide covers testing for all 11 newly implemented features.

---

## 1. System Settings API

**Base URL**: `http://localhost:8080/api/admin/settings`

### Get All Settings
```bash
curl -X GET http://localhost:8080/api/admin/settings \
  -H "Authorization: Bearer YOUR_TOKEN"
```

**Response** (200 OK):
```json
{
  "success": true,
  "message": "System settings retrieved",
  "data": [
    {
      "id": 1,
      "key": "app.name",
      "value": "School Management System",
      "description": "Application name",
      "created_at": "2026-01-29T16:23:58Z"
    }
  ]
}
```

### Get Setting by Key
```bash
curl -X GET http://localhost:8080/api/admin/settings/app.name \
  -H "Authorization: Bearer YOUR_TOKEN"
```

### Create Setting
```bash
curl -X POST http://localhost:8080/api/admin/settings \
  -H "Authorization: Bearer YOUR_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "key": "maintenance.mode",
    "value": "false",
    "description": "Enable/disable maintenance mode"
  }'
```

### Update Setting
```bash
curl -X PUT http://localhost:8080/api/admin/settings/1 \
  -H "Authorization: Bearer YOUR_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "value": "true"
  }'
```

### Delete Setting
```bash
curl -X DELETE http://localhost:8080/api/admin/settings/1 \
  -H "Authorization: Bearer YOUR_TOKEN"
```

---

## 2. Announcements API

**Base URL**: `http://localhost:8080/api/announcements`

### Get All Announcements (Paginated)
```bash
curl -X GET "http://localhost:8080/api/announcements?page=1" \
  -H "Authorization: Bearer YOUR_TOKEN"
```

### Get Active Announcements
```bash
curl -X GET "http://localhost:8080/api/announcements/active?page=1" \
  -H "Authorization: Bearer YOUR_TOKEN"
```

**Response** (200 OK):
```json
{
  "success": true,
  "message": "Active announcements fetched",
  "data": [
    {
      "id": 1,
      "title": "Welcome to New Semester",
      "content": "We are excited to announce the start of the new academic year...",
      "created_by": 1,
      "audience": "all",
      "priority": "high",
      "is_active": true,
      "expires_at": 1735689600,
      "created_at": "2026-01-29T16:23:58Z"
    }
  ],
  "pagination": {
    "page": 1,
    "limit": 20,
    "total": 5
  }
}
```

### Create Announcement
```bash
curl -X POST http://localhost:8080/api/announcements \
  -H "Authorization: Bearer YOUR_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "title": "Campus Maintenance",
    "content": "The campus will be closed for maintenance on Saturday.",
    "audience": "all",
    "priority": "medium",
    "expires_at": 1738368000
  }'
```

### Update Announcement
```bash
curl -X PUT http://localhost:8080/api/announcements/1 \
  -H "Authorization: Bearer YOUR_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "title": "Updated Announcement",
    "is_active": false
  }'
```

### Delete Announcement
```bash
curl -X DELETE http://localhost:8080/api/announcements/1 \
  -H "Authorization: Bearer YOUR_TOKEN"
```

---

## 3. Messages API

**Base URL**: `http://localhost:8080/api/messages`

### Send Message
```bash
curl -X POST http://localhost:8080/api/messages \
  -H "Authorization: Bearer YOUR_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "recipient_id": 3,
    "subject": "Project Feedback",
    "content": "Great work on your project submission!"
  }'
```

**Response** (201 Created):
```json
{
  "success": true,
  "message": "Message created",
  "data": {
    "id": 6,
    "sender_id": 2,
    "recipient_id": 3,
    "subject": "Project Feedback",
    "content": "Great work on your project submission!",
    "is_read": false,
    "created_at": "2026-01-29T16:30:00Z"
  }
}
```

### Get Inbox
```bash
curl -X GET http://localhost:8080/api/messages/inbox \
  -H "Authorization: Bearer YOUR_TOKEN"
```

### Get Conversation with User
```bash
curl -X GET http://localhost:8080/api/messages/conversation/3 \
  -H "Authorization: Bearer YOUR_TOKEN"
```

### Count Unread Messages
```bash
curl -X GET http://localhost:8080/api/messages/unread \
  -H "Authorization: Bearer YOUR_TOKEN"
```

**Response**:
```json
{
  "success": true,
  "message": "Unread message count retrieved",
  "data": {
    "unread_count": 3
  }
}
```

### Mark Message as Read
```bash
curl -X PUT http://localhost:8080/api/messages/5/read \
  -H "Authorization: Bearer YOUR_TOKEN"
```

---

## 4. Notifications API

**Base URL**: `http://localhost:8080/api/notifications`

### Get My Notifications
```bash
curl -X GET http://localhost:8080/api/notifications \
  -H "Authorization: Bearer YOUR_TOKEN"
```

### Get Unread Notifications
```bash
curl -X GET http://localhost:8080/api/notifications/unread \
  -H "Authorization: Bearer YOUR_TOKEN"
```

**Response**:
```json
{
  "success": true,
  "message": "Unread notifications retrieved",
  "data": [
    {
      "id": 1,
      "user_id": 2,
      "title": "Grade Posted",
      "message": "Your grade for Advanced Calculus has been posted",
      "type": "in-app",
      "is_read": false,
      "created_at": "2026-01-29T15:30:00Z"
    }
  ]
}
```

### Create Notification
```bash
curl -X POST http://localhost:8080/api/notifications \
  -H "Authorization: Bearer YOUR_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "title": "Assignment Submitted",
    "message": "Your assignment has been submitted successfully",
    "type": "in-app"
  }'
```

### Mark Notification as Read
```bash
curl -X PUT http://localhost:8080/api/notifications/1/read \
  -H "Authorization: Bearer YOUR_TOKEN"
```

### Mark All Notifications as Read
```bash
curl -X PUT http://localhost:8080/api/notifications/mark-all-read \
  -H "Authorization: Bearer YOUR_TOKEN"
```

### Delete Notification
```bash
curl -X DELETE http://localhost:8080/api/notifications/1 \
  -H "Authorization: Bearer YOUR_TOKEN"
```

---

## 5. Payments API

**Base URL**: `http://localhost:8080/api/payments`

### Create Payment
```bash
curl -X POST http://localhost:8080/api/payments \
  -H "Authorization: Bearer YOUR_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "student_id": 2,
    "amount": 2500.00,
    "status": "pending",
    "description": "Spring 2026 Tuition",
    "due_date": 1738368000
  }'
```

### Get Student Payments
```bash
curl -X GET http://localhost:8080/api/payments/student/2 \
  -H "Authorization: Bearer YOUR_TOKEN"
```

**Response**:
```json
{
  "success": true,
  "message": "Student payments retrieved",
  "data": [
    {
      "id": 1,
      "student_id": 2,
      "amount": 2500.00,
      "status": "paid",
      "description": "Spring 2026 Tuition",
      "due_date": 1735689600,
      "created_at": "2026-01-19T16:23:58Z"
    },
    {
      "id": 2,
      "student_id": 2,
      "amount": 150.00,
      "status": "paid",
      "description": "Laboratory Fee",
      "due_date": 1735689600,
      "created_at": "2026-01-21T16:23:58Z"
    }
  ]
}
```

### Get All Payments
```bash
curl -X GET http://localhost:8080/api/payments \
  -H "Authorization: Bearer YOUR_TOKEN"
```

### Update Payment
```bash
curl -X PUT http://localhost:8080/api/payments/1 \
  -H "Authorization: Bearer YOUR_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "status": "paid"
  }'
```

### Get Student Balance
```bash
curl -X GET http://localhost:8080/api/payments/balance/2 \
  -H "Authorization: Bearer YOUR_TOKEN"
```

**Response**:
```json
{
  "success": true,
  "message": "Student balance retrieved",
  "data": {
    "student_id": 2,
    "total_due": 0.00,
    "total_paid": 2650.00,
    "balance": 0.00
  }
}
```

---

## 6. TimeTable (Class Schedule) API

**Base URL**: `http://localhost:8080/api/timetable`

### Get All TimeTable Entries
```bash
curl -X GET http://localhost:8080/api/timetable \
  -H "Authorization: Bearer YOUR_TOKEN"
```

### Get TimeTable by Course
```bash
curl -X GET http://localhost:8080/api/timetable/course/1 \
  -H "Authorization: Bearer YOUR_TOKEN"
```

**Response**:
```json
{
  "success": true,
  "message": "Timetable entries retrieved",
  "data": [
    {
      "id": 1,
      "course_id": 1,
      "teacher_id": 2,
      "day_of_week": "Monday",
      "start_time": "09:00",
      "end_time": "10:30",
      "classroom": "101",
      "location": "Building A",
      "created_at": "2026-01-29T16:23:58Z"
    }
  ]
}
```

### Get TimeTable by Teacher
```bash
curl -X GET http://localhost:8080/api/timetable/teacher/2 \
  -H "Authorization: Bearer YOUR_TOKEN"
```

### Get TimeTable by Day
```bash
curl -X GET http://localhost:8080/api/timetable/day/Monday \
  -H "Authorization: Bearer YOUR_TOKEN"
```

### Create TimeTable Entry
```bash
curl -X POST http://localhost:8080/api/timetable \
  -H "Authorization: Bearer YOUR_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "course_id": 1,
    "teacher_id": 2,
    "day_of_week": "Tuesday",
    "start_time": "10:00",
    "end_time": "11:30",
    "classroom": "102",
    "location": "Building A"
  }'
```

### Update TimeTable Entry
```bash
curl -X PUT http://localhost:8080/api/timetable/1 \
  -H "Authorization: Bearer YOUR_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "classroom": "103"
  }'
```

### Delete TimeTable Entry
```bash
curl -X DELETE http://localhost:8080/api/timetable/1 \
  -H "Authorization: Bearer YOUR_TOKEN"
```

---

## 7. Grade Transcript (Academic Records) API

**Base URL**: `http://localhost:8080/api/transcripts`

### Get Student Transcripts
```bash
curl -X GET http://localhost:8080/api/transcripts/student/2 \
  -H "Authorization: Bearer YOUR_TOKEN"
```

**Response**:
```json
{
  "success": true,
  "message": "Student transcripts retrieved",
  "data": [
    {
      "id": 1,
      "student_id": 2,
      "academic_year": "2024-2025",
      "term": "Fall",
      "gpa": 3.75,
      "total_credits": 30,
      "completed_credits": 30,
      "created_at": "2026-01-29T16:23:58Z"
    }
  ]
}
```

### Get Latest Transcript
```bash
curl -X GET http://localhost:8080/api/transcripts/latest/2 \
  -H "Authorization: Bearer YOUR_TOKEN"
```

### Get GPA
```bash
curl -X GET http://localhost:8080/api/transcripts/gpa/2 \
  -H "Authorization: Bearer YOUR_TOKEN"
```

**Response**:
```json
{
  "success": true,
  "message": "GPA retrieved",
  "data": {
    "student_id": 2,
    "current_gpa": 3.80,
    "cumulative_gpa": 3.77
  }
}
```

---

## 8. Backup API

**Base URL**: `http://localhost:8080/api/admin/backups`

### Get All Backups
```bash
curl -X GET http://localhost:8080/api/admin/backups \
  -H "Authorization: Bearer YOUR_TOKEN"
```

**Response**:
```json
{
  "success": true,
  "message": "Backups fetched",
  "data": [
    {
      "id": 1,
      "backup_name": "backup_2026_01_29_automatic",
      "status": "completed",
      "size_bytes": 5242880,
      "file_path": "/backups/2026/01/29/backup_automatic.sql",
      "backed_up_at": 1737986638,
      "created_at": "2026-01-29T14:23:58Z"
    }
  ]
}
```

### Get Latest Backup
```bash
curl -X GET http://localhost:8080/api/admin/backups/latest \
  -H "Authorization: Bearer YOUR_TOKEN"
```

### Get Backup by ID
```bash
curl -X GET http://localhost:8080/api/admin/backups/1 \
  -H "Authorization: Bearer YOUR_TOKEN"
```

### Delete Backup
```bash
curl -X DELETE http://localhost:8080/api/admin/backups/1 \
  -H "Authorization: Bearer YOUR_TOKEN"
```

---

## 9. Import Batch API

**Base URL**: `http://localhost:8080/api/admin/imports`

### Get All Import Batches
```bash
curl -X GET http://localhost:8080/api/admin/imports \
  -H "Authorization: Bearer YOUR_TOKEN"
```

**Response**:
```json
{
  "success": true,
  "message": "Batches fetched",
  "data": [
    {
      "id": 1,
      "entity_type": "students",
      "status": "completed",
      "total_records": 50,
      "successful_records": 50,
      "failed_records": 0,
      "error_log": null,
      "created_at": "2026-01-22T16:23:58Z"
    }
  ]
}
```

### Get Batch by ID
```bash
curl -X GET http://localhost:8080/api/admin/imports/1 \
  -H "Authorization: Bearer YOUR_TOKEN"
```

### Get Batches by Status
```bash
curl -X GET http://localhost:8080/api/admin/imports/status/completed \
  -H "Authorization: Bearer YOUR_TOKEN"
```

### Delete Batch
```bash
curl -X DELETE http://localhost:8080/api/admin/imports/1 \
  -H "Authorization: Bearer YOUR_TOKEN"
```

---

## Testing Checklist

### System Settings
- [ ] Retrieve all settings
- [ ] Get setting by key
- [ ] Create new setting
- [ ] Update existing setting
- [ ] Delete setting

### Announcements
- [ ] Get all announcements (pagination)
- [ ] Get active announcements only
- [ ] Create announcement with different audiences (all/students/teachers)
- [ ] Update announcement
- [ ] Delete announcement

### Messages
- [ ] Send message between users
- [ ] Get inbox
- [ ] Get conversation with specific user
- [ ] Count unread messages
- [ ] Mark message as read

### Notifications
- [ ] Get my notifications
- [ ] Get unread notifications
- [ ] Create notification
- [ ] Mark single notification as read
- [ ] Mark all notifications as read
- [ ] Delete notification

### Payments
- [ ] Create payment
- [ ] Get student payments
- [ ] Get all payments (admin)
- [ ] Update payment status
- [ ] Calculate student balance

### TimeTable
- [ ] Get all schedule entries
- [ ] Filter by course
- [ ] Filter by teacher
- [ ] Filter by day of week
- [ ] Create schedule entry
- [ ] Update schedule entry
- [ ] Delete schedule entry

### Grade Transcript
- [ ] Get student transcripts
- [ ] Get latest transcript
- [ ] Get current GPA
- [ ] Get cumulative GPA

### Backup
- [ ] List all backups
- [ ] Get latest backup
- [ ] Get backup by ID
- [ ] Verify backup file size

### Import Batch
- [ ] List all import batches
- [ ] Get batch by ID
- [ ] Filter by status
- [ ] Check error logs for failed batches

---

## Error Handling

All endpoints follow consistent error response format:

```json
{
  "success": false,
  "message": "Error description",
  "error": "error_code"
}
```

**Common HTTP Status Codes**:
- `200` - OK
- `201` - Created
- `400` - Bad Request
- `401` - Unauthorized
- `403` - Forbidden
- `404` - Not Found
- `500` - Internal Server Error

---

## Authentication

All endpoints require a valid JWT token in the Authorization header:

```
Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...
```

Obtain a token by logging in:

```bash
curl -X POST http://localhost:8080/api/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "admin@school.com",
    "password": "admin123"
  }'
```

---

## Notes

- All timestamps are Unix epoch (seconds since Jan 1, 1970)
- Pagination defaults: page=1, limit=20
- All new features require admin or appropriate user role authorization
- Transaction IDs and user IDs must exist in the system
