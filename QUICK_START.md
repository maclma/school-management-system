# School Management System - Quick Access Guide

## 🚀 Start Here

### 1. Ensure Both Servers Are Running

**Backend Server** (must run first)
```bash
cd c:\Users\dell\school-management-system
go run cmd/server/main.go
# Expected output: Server listening on port 8080
```

**Frontend Server**
```bash
cd c:\Users\dell\school-management-system\frontend
node simple-server.js
# Expected output: Frontend server running at http://localhost:3001
```

### 2. Open Browser
```
URL: http://localhost:3001
```

### 3. Login with Test Account
```
Email: admin@school.com
Password: admin
```

---

## 📑 All Available Pages

### Main Entry Points
| Page | URL | Purpose |
|------|-----|---------|
| Login | `/login.html` | User authentication |
| Register | `/register.html` | New user signup |
| Features | `/features.html` | System overview |
| Home | `/home.html` | Main dashboard |

### Feature Pages
| Page | URL | Role | Feature |
|------|-----|------|---------|
| Grades | `/grades.html` | All | Grade management |
| Attendance | `/attendance.html` | All | Attendance tracking |
| Assignments | `/assignments.html` | All | Assignment system |
| Announcements | `/announcements.html` | All | Announcements |
| Notifications | `/notifications.html` | All | Notification center |
| Messages | `/messages.html` | All | Direct messaging |
| Payments | `/payments.html` | All | Payment tracking |
| Transcripts | `/transcripts.html` | All | Academic records |
| Timetables | `/timetables.html` | All | Class schedule |
| Dashboard | `/dashboard.html` | All | Courses |
| Admin Panel | `/admin-dashboard.html` | Admin | System admin |

---

## 🎯 Feature Quick Links

### Click These Links After Login

```
Home Dashboard
└── Grades
    └── View: http://localhost:3001/grades.html
    └── View Your Grades & GPA
    
└── Attendance  
    └── View: http://localhost:3001/attendance.html
    └── Check Attendance Record
    
└── Assignments
    └── View: http://localhost:3001/assignments.html
    └── Submit & Track Assignments
    
└── Announcements
    └── View: http://localhost:3001/announcements.html
    └── Read School Updates
    
└── Notifications
    └── View: http://localhost:3001/notifications.html
    └── Check All Alerts
    
└── Messages
    └── View: http://localhost:3001/messages.html
    └── Send/Receive Messages
    
└── Payments
    └── View: http://localhost:3001/payments.html
    └── Manage Fees & Payments
    
└── Transcripts
    └── View: http://localhost:3001/transcripts.html
    └── Download Academic Records
    
└── Schedule
    └── View: http://localhost:3001/timetables.html
    └── View Class Timetable
    
└── Admin Panel (if admin)
    └── View: http://localhost:3001/admin-dashboard.html
    └── System Management
```

---

## 🔑 Test Credentials

### Admin Account
```
Email: admin@school.com
Password: admin
Role: admin
Access: All features including admin panel
```

### Create Your Own Account
1. Go to `/register.html`
2. Fill in details
3. Select role (Student/Teacher)
4. Click Register
5. Login with new credentials

---

## ✅ Feature Testing Checklist

### Login & Authentication
- [ ] Login with admin account
- [ ] Check auto-redirect when logged in
- [ ] Logout and verify redirect to login
- [ ] Register new account
- [ ] Login with new account

### Academic Features
- [ ] View grades page
- [ ] View attendance page
- [ ] View assignments page
- [ ] View transcripts page
- [ ] Check GPA calculation

### Communication Features
- [ ] View notifications
- [ ] View announcements
- [ ] Check messages
- [ ] (Optional) Send a message

### Administrative Features
- [ ] View admin dashboard (if admin)
- [ ] Check system statistics
- [ ] View payment page
- [ ] View timetables page

### General
- [ ] Click all navigation links
- [ ] Test logout from all pages
- [ ] Check responsive design (resize browser)
- [ ] Verify all buttons work

---

## 🐛 Troubleshooting

### Problem: Can't connect to localhost:3001

**Solution:**
```bash
# Check if Node.js server is running
Get-Process node

# If not running, start it:
cd c:\Users\dell\school-management-system\frontend
node simple-server.js
```

### Problem: Login fails

**Solution:**
```bash
# Check if backend is running
Get-Process go

# Verify admin account exists
# Use credentials: admin@school.com / admin

# Check browser console (F12) for errors
```

### Problem: Blank page after login

**Solution:**
1. Clear browser cache: Ctrl + Shift + Delete
2. Clear localStorage: 
   ```javascript
   // In browser console (F12)
   localStorage.clear()
   location.reload()
   ```
3. Try login again

### Problem: API calls failing

**Solution:**
```bash
# Verify backend is running on port 8080
netstat -ano | findstr :8080

# Restart backend:
go run cmd/server/main.go

# Check browser Network tab (F12) for API calls
```

---

## 📊 System Status Check

### Check Backend
```bash
curl http://localhost:8080/health
# Should return: {"status":"ok"}
```

### Check Frontend
```
Open: http://localhost:3001
# Should load login page
```

### Check Database
```bash
# Login to PostgreSQL (if you have psql installed)
psql -U postgres

# List databases
\l

# Connect to school_db
\c school_db

# List tables
\dt
```

---

## 🔧 Common Commands

### Backend
```bash
# Start server
go run cmd/server/main.go

# Build binary
go build -o sms-server cmd/server/main.go

# Run tests
go test ./...
```

### Frontend
```bash
# Start server
node simple-server.js

# View logs
# Check browser console (F12)

# Test API
# Use test-integration.html page
```

### Database
```bash
# Reset admin password (if needed)
go run reset_admin.go

# Backup database
pg_dump school_db > backup.sql

# Restore database
psql school_db < backup.sql
```

---

## 📱 Responsive Testing

### Test on Different Sizes
1. **Desktop** (1920x1080)
   - All features visible
   - Navigation horizontal

2. **Tablet** (768x1024)
   - Responsive grid
   - Touch-friendly buttons

3. **Mobile** (375x667)
   - Stack vertically
   - Touch-optimized

### Browser DevTools
```
Press: F12 → Device Toggle → Select Device
```

---

## 🎯 Feature Matrix

| Feature | Student | Teacher | Admin |
|---------|---------|---------|-------|
| View Grades | ✅ | ✅ | ✅ |
| Record Grades | ❌ | ✅ | ✅ |
| Attendance | ✅ | ✅ | ✅ |
| Assignments | ✅ | ✅ | ✅ |
| Announcements | ✅ | ✅ | ✅ |
| Messages | ✅ | ✅ | ✅ |
| Notifications | ✅ | ✅ | ✅ |
| Payments | ✅ | ✅ | ✅ |
| Transcripts | ✅ | ✅ | ✅ |
| Schedule | ✅ | ✅ | ✅ |
| Admin Panel | ❌ | ❌ | ✅ |

---

## 📞 Quick Help

**Need to reset a password?**
```bash
go run reset_admin.go
# Creates admin@school.com with password: admin
```

**Want to see API responses?**
1. Open browser DevTools (F12)
2. Go to Network tab
3. Perform an action
4. Click on API request
5. View Response tab

**Want to check stored data?**
```javascript
// In browser console (F12)
console.log(localStorage.getItem('sms_user'))
console.log(localStorage.getItem('sms_token'))
```

**Want to test API directly?**
```bash
# In PowerShell
$token = "your-token-here"
$headers = @{ Authorization = "Bearer $token" }
Invoke-WebRequest http://localhost:8080/api/profile -Headers $headers
```

---

## 🎓 Learning Path

### For First-Time Users
1. Start at `/features.html` to understand what's available
2. Login to access `/home.html`
3. Explore each feature page from the dashboard
4. Click the "View" buttons to access features

### For Administrators
1. Login as admin
2. Go to `/admin-dashboard.html`
3. Manage enrollments
4. View system statistics
5. Manage announcements

### For Teachers
1. Login with teacher account
2. Go to `/grades.html` to record grades
3. Go to `/attendance.html` to mark attendance
4. Go to `/assignments.html` to create assignments
5. Post announcements from `/announcements.html`

### For Students
1. Login with student account
2. View your dashboard
3. Check `/grades.html` for your grades
4. Check `/attendance.html` for attendance
5. Submit assignments from `/assignments.html`
6. View transcripts from `/transcripts.html`

---

## 📚 Documentation Files

- **PROJECT_COMPLETION_SUMMARY.md** - Overall project status
- **FEATURES_COMPLETE.md** - Detailed feature documentation
- **FRONTEND_FILES_COMPLETE.md** - Frontend pages list
- **DEPLOYMENT_COMPLETE.md** - Deployment guide
- **README.md** - General information

---

## ✨ Pro Tips

1. **Bookmark important pages**
   - Save http://localhost:3001/home.html
   - Save http://localhost:3001/admin-dashboard.html

2. **Use browser dev tools**
   - F12 → Console to see errors
   - F12 → Network to see API calls
   - F12 → Application to check localStorage

3. **Create multiple accounts**
   - Test as different roles
   - Verify access control works

4. **Keep servers running**
   - Use separate terminal windows
   - Minimize to taskbar
   - Don't close windows

5. **Export important data**
   - Download transcripts
   - Save announcements
   - Keep payment records

---

## 🎉 You're Ready!

Everything is set up and ready to go. Start both servers, open the browser, and begin exploring the School Management System!

**Happy Learning! 🚀**
