# 🧪 LIVE TESTING SESSION - January 31, 2026

## 🚀 Testing Guide

Follow these steps to test each feature:

---

## ✅ Step 1: Login Test

**Page**: http://localhost:3000/system-test.html

1. See the login section at the top
2. Email field shows: `admin@school.com`
3. Password field shows: `admin123`
4. Click **"Test Login"** button
5. Check output box for:
   - ✅ "Login successful!"
   - Token displayed (first 20 chars + ...)
   - User object shown

**Expected**: Green success message with token

---

## ✅ Step 2: Backend Health Test

**After login**, click **"Test Backend"** button in Backend & API card

1. Should show: ✅ Backend is healthy
2. Response shows service status and timestamp
3. Confirms database and API are working

**Expected**: Green checkmark with health status

---

## ✅ Step 3: User Profile Test

Click **"Test Users"** button in User Management card

1. Should retrieve your profile
2. Shows: Name, Email, Role
3. Confirms authentication is working

**Expected**: Your user profile data

---

## ✅ Step 4: Courses & Enrollment Test

Click **"Test Courses"** button in Courses & Enrollment card

1. Should list all available courses
2. Shows course title, code, department
3. Displays number of courses found

**Expected**: Course list with details

---

## ✅ Step 5: Timetables Test ⭐ (KEY FEATURE)

Click **"Test Timetables"** button in Timetables & Schedule card

**This tests**:
- GET /api/timetable (all timetables)
- GET /api/timetable/course/:id (filter by course)
- GET /api/timetable/teacher/:id (filter by teacher)
- GET /api/timetable/day/:day (filter by day)

1. Should show timetable count
2. Sample timetable entries displayed
3. Contains: course, day, time, room, teacher

**Expected**: Timetable data with schedules

---

## ✅ Step 6: Grades Test

Click **"Test Grades"** button in Grades & Academic card

1. Shows number of grades found
2. Sample grade entries
3. Contains: course, score, grade

**Expected**: Grade records for student

---

## ✅ Step 7: Attendance Test

Click **"Test Attendance"** button in Attendance card

1. Shows attendance records
2. Contains: date, course, status (present/absent/late)
3. Displays record count

**Expected**: Attendance data

---

## ✅ Step 8: Assignments Test

Click **"Test Assignments"** button in Assignments card

1. Shows assignments list
2. Contains: title, course, due date
3. Submission status

**Expected**: Assignment records

---

## ✅ Step 9: Payments Test

Click **"Test Payments"** button in Payments card

1. Shows payment records
2. Contains: amount, date, status
3. Payment history

**Expected**: Payment data

---

## ✅ Step 10: Communications Test

Click **"Test Communications"** button in Communications card

1. Shows announcements
2. Shows notifications
3. Both should list records

**Expected**: Announcements and notifications

---

## 🎯 Manual Testing - Timetables Feature

After API testing, manually test the timetable page:

1. Go to: http://localhost:3000/timetables.html
2. Login if needed (admin@school.com / admin123)
3. You should see:
   - Weekly timetable grid (Monday-Saturday)
   - Time slots (8:00 AM - 6:00 PM)
   - Color-coded class blocks with:
     - Course name
     - Room number
     - Teacher name
4. Try filters:
   - Select a course from dropdown
   - Click filter button
   - Schedule should update
5. Try adding a class:
   - Select course
   - Choose day
   - Set start/end time
   - Enter room number
   - Click "Add to Schedule"
   - Verify it appears

**Expected**: Full timetable interface working

---

## 🎯 Manual Testing - Course Enrollment

Test the enrollment feature:

1. Go to: http://localhost:3000/dashboard.html
2. Login if needed
3. You should see:
   - User info card
   - "Courses" section with list
   - "Your Enrollments" section
4. For each course:
   - Click "Enroll" button
   - Check message says "Enrolled"
   - Course appears in "Your Enrollments"
5. Click "Details" on a course:
   - Goes to course.html?id=X
   - Shows full course info
   - Can also enroll from here

**Expected**: Enrollment works end-to-end

---

## 📊 What to Look For

### Green ✅ Indicators
- API returns data successfully
- No errors in response
- Data displays properly
- Features are accessible
- Operations complete

### Red ❌ Warnings
- Error messages in output
- Empty responses
- 404 or 500 errors
- Failed operations
- Missing data

---

## 🔍 Checking API Responses

In each test output box, look for:

```
✅ [Action completed]
JSON data showing returned values
```

If you see:
```
❌ Error: [message]
```

Then check:
1. Are both servers running?
2. Is token valid?
3. Are there console errors (F12)?
4. Check network tab for API response

---

## 📱 Browser Tools (F12)

Press F12 to open Developer Tools:

1. **Console Tab**: Check for JavaScript errors
2. **Network Tab**: See API calls and responses
3. **Storage Tab**: Check localStorage for token
4. **Application Tab**: View saved data

---

## ✅ Success Checklist

After testing, confirm:

- [ ] Login successful with token
- [ ] Backend health check passes
- [ ] Can see user profile
- [ ] Can list courses
- [ ] Can see timetables
- [ ] Can view grades
- [ ] Can see attendance
- [ ] Can view assignments
- [ ] Can see payments
- [ ] Can see communications
- [ ] Timetable page loads
- [ ] Can filter timetables
- [ ] Can add to schedule
- [ ] Can view courses on dashboard
- [ ] Can enroll in courses
- [ ] Enrollments appear in "My Enrollments"
- [ ] All data persists after refresh

---

## 🆘 Troubleshooting

**If backend not responding**:
```powershell
# Check if running
netstat -ano | findstr :8080

# Restart
cd c:\Users\dell\school-management-system
go run cmd/server/main.go
```

**If frontend not loading**:
```powershell
# Check if running
netstat -ano | findstr :3000

# Restart
cd c:\Users\dell\school-management-system\frontend
npm run dev
```

**If token issues**:
```javascript
// In browser console
localStorage.clear()
// Then refresh page and login again
```

---

## 📞 Need Help?

1. Check SYSTEM_TESTING_GUIDE.md for details
2. Review TIMETABLE_COURSE_QUICK_GUIDE.md for features
3. Check browser console (F12) for errors
4. Verify servers are running
5. Try clearing localStorage and re-login

---

## 🎉 Ready?

1. Open http://localhost:3000/system-test.html
2. Follow the steps above
3. Click each test button
4. Navigate to feature pages
5. Test manual operations
6. Report results!

---

**Session**: January 31, 2026
**Status**: Ready for Testing
**All Systems**: GO ✅
