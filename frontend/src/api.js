// small API wrapper used by React app
// Use Vite env var `VITE_API_BASE` when available (dev/prod), fallback to `/api` for same-origin
let base = (typeof import.meta !== 'undefined' && import.meta.env && import.meta.env.VITE_API_BASE) || '/api'

// If running in local dev and no explicit VITE_API_BASE is set, point API to backend server
if (base === '/api' && typeof window !== 'undefined') {
  const host = window.location.hostname
  const port = window.location.port
  if (host === 'localhost' || host === '127.0.0.1') {
    // If the frontend dev server port is different from backend (8080), target backend directly
    if (port && port !== '8080') {
      base = 'http://localhost:8080/api'
    }
  }
}

async function request(path, opts={}){
  const headers = opts.headers || {}
  if (!headers['Content-Type'] && !(opts.body instanceof FormData)) headers['Content-Type']='application/json'

  const token = localStorage.getItem('sms_token')
  if (token) headers['Authorization'] = 'Bearer ' + token

  const res = await fetch(base + path, Object.assign({}, opts, { headers }))
  let body = null
  try{ body = await res.json() }catch(e){}
  if (!res.ok) {
    const err = (body && (body.error || body.message)) || res.statusText
    const e = new Error(err)
    e.status = res.status
    e.body = body
    throw e
  }
  return body
}

export default {
  login: (email, password) => request('/auth/login',{ method:'POST', body: JSON.stringify({email,password}) }),
  register: (payload) => request('/auth/register',{ method:'POST', body: JSON.stringify(payload) }),
  getProfile: () => request('/profile'),
  updateProfile: (payload) => request('/profile', { method:'PUT', body: JSON.stringify(payload) }),
  getCourses: (page=1, limit=50) => request(`/courses?page=${page}&limit=${limit}`),
  getCourse: (id) => request(`/courses/${id}`),
  createCourse: (payload) => request('/courses', { method:'POST', body: JSON.stringify(payload) }),
  enroll: (studentId, courseId) => request('/enrollments', { method:'POST', body: JSON.stringify({ student_id: studentId, course_id: courseId }) }),
  getEnrollmentsByStudent: (studentId) => request(`/enrollments/by-student/${studentId}`),
  getCourseEnrollments: (courseId) => request(`/enrollments/by-course/${courseId}`),
  getGradesByStudent: (studentId) => request(`/grades/by-student/${studentId}`),
  recordGrade: (payload) => request('/teacher/grades', { method:'POST', body: JSON.stringify(payload) }),
  getAttendanceByStudent: (studentId) => request(`/attendance/by-student/${studentId}`),
  recordAttendance: (payload) => request('/teacher/attendance', { method:'POST', body: JSON.stringify(payload) }),
  getAdminStats: () => request('/admin/dashboard'),
  getAdminUsers: () => request('/admin/users'),
  createUserAdmin: (payload) => request('/admin/users', { method:'POST', body: JSON.stringify(payload) }),
  updateUserAdmin: (userId, payload) => request(`/admin/users/${userId}`, { method:'PUT', body: JSON.stringify(payload) }),
  deleteUser: (userId) => request(`/admin/users/${userId}`, { method:'DELETE' }),
  updateUserStatus: (userId, status) => request(`/users/${userId}/status`, { method:'PATCH', body: JSON.stringify({ status }) }),
  getAdminEnrollments: () => request('/admin/enrollments'),
  approveEnrollment: (enrollmentId) => request(`/admin/enrollments/${enrollmentId}/approve`, { method:'POST' }),
  rejectEnrollment: (enrollmentId) => request(`/admin/enrollments/${enrollmentId}/reject`, { method:'POST' }),
  getStudentEnrollments: () => request('/student/enrollments'),
  getStudentGrades: () => request('/student/grades'),
  getStudentAttendance: () => request('/student/attendance'),
  // Assignments
  createAssignment: (payload) => request('/assignments', { method:'POST', body: JSON.stringify(payload) }),
  getAssignment: (id) => request(`/assignments/${id}`),
  updateAssignment: (id, payload) => request(`/assignments/${id}`, { method:'PUT', body: JSON.stringify(payload) }),
  deleteAssignment: (id) => request(`/assignments/${id}`, { method:'DELETE' }),
  getAssignmentsByCourse: (courseId) => request(`/assignments/course/${courseId}`),
  submitAssignment: (payload) => request('/assignments/submit', { method:'POST', body: JSON.stringify(payload) }),
  getSubmissionsByAssignment: (assignmentId) => request(`/submissions/assignment/${assignmentId}`),
  gradeSubmission: (submissionId, payload) => request(`/submissions/${submissionId}/grade`, { method:'PUT', body: JSON.stringify(payload) }),
  getAssignmentsByTeacher: () => request('/teacher/assignments'),
  getSubmissionsByStudent: () => request('/student/assignments/submissions'),
  
  // New Feature APIs - System Settings
  getSystemSettings: () => request('/admin/settings'),
  getSystemSetting: (key) => request(`/admin/settings/${key}`),
  createSystemSetting: (payload) => request('/admin/settings', { method:'POST', body: JSON.stringify(payload) }),
  updateSystemSetting: (id, payload) => request(`/admin/settings/${id}`, { method:'PUT', body: JSON.stringify(payload) }),
  deleteSystemSetting: (id) => request(`/admin/settings/${id}`, { method:'DELETE' }),
  
  // New Feature APIs - Announcements
  getAnnouncements: (page=1) => request(`/announcements?page=${page}`),
  getActiveAnnouncements: (page=1) => request(`/announcements/active?page=${page}`),
  createAnnouncement: (payload) => request('/announcements', { method:'POST', body: JSON.stringify(payload) }),
  updateAnnouncement: (id, payload) => request(`/announcements/${id}`, { method:'PUT', body: JSON.stringify(payload) }),
  deleteAnnouncement: (id) => request(`/announcements/${id}`, { method:'DELETE' }),
  
  // New Feature APIs - Messages
  sendMessage: (payload) => request('/messages', { method:'POST', body: JSON.stringify(payload) }),
  getInbox: () => request('/messages/inbox'),
  getConversation: (userId) => request(`/messages/conversation/${userId}`),
  countUnreadMessages: () => request('/messages/unread'),
  markMessageAsRead: (id) => request(`/messages/${id}/read`, { method:'PUT' }),
  
  // New Feature APIs - Notifications
  getMyNotifications: (page=1) => request(`/notifications?page=${page}`),
  getUnreadNotifications: () => request('/notifications/unread'),
  createNotification: (payload) => request('/notifications', { method:'POST', body: JSON.stringify(payload) }),
  markNotificationAsRead: (id) => request(`/notifications/${id}/read`, { method:'PUT' }),
  markAllNotificationsAsRead: () => request('/notifications/mark-all-read', { method:'PUT' }),
  deleteNotification: (id) => request(`/notifications/${id}`, { method:'DELETE' }),
  
  // New Feature APIs - Payments
  createPayment: (payload) => request('/payments', { method:'POST', body: JSON.stringify(payload) }),
  getStudentPayments: (studentId) => request(`/payments/student/${studentId}`),
  getAllPayments: (page=1) => request(`/payments?page=${page}`),
  updatePayment: (id, payload) => request(`/payments/${id}`, { method:'PUT', body: JSON.stringify(payload) }),
  getStudentBalance: (studentId) => request(`/payments/balance/${studentId}`),
  
  // New Feature APIs - TimeTable
  getTimetable: (page=1) => request(`/timetable?page=${page}`),
  getTimetableByCourse: (courseId) => request(`/timetable/course/${courseId}`),
  getTimetableByTeacher: (teacherId) => request(`/timetable/teacher/${teacherId}`),
  getTimetableByDay: (day) => request(`/timetable/day/${day}`),
  createTimetableEntry: (payload) => request('/timetable', { method:'POST', body: JSON.stringify(payload) }),
  updateTimetableEntry: (id, payload) => request(`/timetable/${id}`, { method:'PUT', body: JSON.stringify(payload) }),
  deleteTimetableEntry: (id) => request(`/timetable/${id}`, { method:'DELETE' }),
  
  // New Feature APIs - Grade Transcripts
  getStudentTranscripts: (studentId) => request(`/transcripts/student/${studentId}`),
  getLatestTranscript: (studentId) => request(`/transcripts/latest/${studentId}`),
  getStudentGPA: (studentId) => request(`/transcripts/gpa/${studentId}`),
  
  // New Feature APIs - Backups
  getAllBackups: (page=1) => request(`/admin/backups?page=${page}`),
  getLatestBackup: () => request('/admin/backups/latest'),
  getBackupById: (id) => request(`/admin/backups/${id}`),
  deleteBackup: (id) => request(`/admin/backups/${id}`, { method:'DELETE' }),
  
  // New Feature APIs - Import Batches
  getAllImportBatches: (page=1) => request(`/admin/imports?page=${page}`),
  getImportBatchById: (id) => request(`/admin/imports/${id}`),
  getImportBatchesByStatus: (status, page=1) => request(`/admin/imports/status/${status}?page=${page}`),
  deleteImportBatch: (id) => request(`/admin/imports/${id}`, { method:'DELETE' }),
  
  // ===== SESSION 6: ADVANCED FEATURES =====
  
  // Advanced Search APIs (5 endpoints)
  searchAnnouncements: (query='', audience='', priority='', page=1, limit=10) => 
    request(`/search/announcements?query=${query}&audience=${audience}&priority=${priority}&page=${page}&limit=${limit}`),
  searchPayments: (studentId=0, status='', page=1, limit=10) => 
    request(`/search/payments?student_id=${studentId}&status=${status}&page=${page}&limit=${limit}`),
  searchStudents: (query, page=1, limit=10) => 
    request(`/search/students?query=${encodeURIComponent(query)}&page=${page}&limit=${limit}`),
  searchGradesByRange: (courseId, minScore=0, maxScore=100, page=1, limit=10) => 
    request(`/search/grades?course_id=${courseId}&min_score=${minScore}&max_score=${maxScore}&page=${page}&limit=${limit}`),
  searchOverduePayments: () => 
    request('/search/overdue-payments'),
  
  // CSV Export APIs (5 endpoints)
  exportPaymentsCSV: (studentId=0, status='') => 
    request(`/export/payments?student_id=${studentId}&status=${status}`),
  exportGradesCSV: (courseId) => 
    request(`/export/grades?course_id=${courseId}`),
  exportAttendanceCSV: (courseId) => 
    request(`/export/attendance?course_id=${courseId}`),
  exportStudentTranscriptCSV: (studentId) => 
    request(`/export/transcript/${studentId}`),
  exportEnrollmentsCSV: (courseId) => 
    request(`/export/enrollments?course_id=${courseId}`),
  
  // Attendance Automation APIs (5 endpoints)
  getAttendanceStatsByCourse: (courseId) => 
    request(`/attendance/stats/course/${courseId}`),
  getStudentAttendancePercentage: (studentId, courseId) => 
    request(`/attendance/percentage/${studentId}/${courseId}`),
  checkLowAttendance: (payload) => 
    request('/attendance/check-low', { method:'POST', body: JSON.stringify(payload) }),
  getStudentsWithLowAttendance: (threshold, courseId=0) => {
    const url = courseId ? `/attendance/low/${threshold}?course_id=${courseId}` : `/attendance/low/${threshold}`;
    return request(url);
  },
  getAttendanceReport: (courseId) => 
    request(`/attendance/report/${courseId}`),
  
  // Grade Auto-Calculation APIs (4 endpoints)
  recordGradeWithAutoCalc: (payload) => 
    request('/grades/auto', { method:'POST', body: JSON.stringify(payload) }),
  getCourseAverageGrade: (courseId) => 
    request(`/grades/course-average/${courseId}`),
  getGradeDistribution: (courseId) => 
    request(`/grades/distribution/${courseId}`),
  getStudentGradeStats: (studentId) => 
    request(`/grades/student-stats/${studentId}`),
  
  // Rubrics APIs (7 endpoints)
  createRubric: (payload) => 
    request('/rubrics', { method:'POST', body: JSON.stringify(payload) }),
  getRubric: (id) => 
    request(`/rubrics/${id}`),
  getRubricsByAssignment: (assignmentId) => 
    request(`/rubrics/assignment/${assignmentId}`),
  updateRubric: (id, payload) => 
    request(`/rubrics/${id}`, { method:'PUT', body: JSON.stringify(payload) }),
  deleteRubric: (id) => 
    request(`/rubrics/${id}`, { method:'DELETE' }),
  scoreSubmission: (payload) => 
    request('/rubrics/score', { method:'POST', body: JSON.stringify(payload) }),
  getSubmissionScore: (submissionId) => 
    request(`/rubrics/score/${submissionId}`),
}
