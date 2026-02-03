// Complete API helper for School Management System
// Attaches `window.api` with all convenience methods for frontend features
(function(){
  const base = 'http://localhost:8080';

  async function request(path, opts = {}){
    const headers = opts.headers || {};
    headers['Content-Type'] = headers['Content-Type'] || 'application/json';

    const token = localStorage.getItem('sms_token');
    if (token) headers['Authorization'] = 'Bearer ' + token;

    const res = await fetch(base + path, Object.assign({}, opts, { headers }));
    let body = null;
    try { body = await res.json(); } catch(e) { body = null; }
    if (!res.ok) {
      const err = (body && (body.error || body.message)) || res.statusText || 'Request failed';
      const e = new Error(err);
      e.status = res.status;
      e.body = body;
      throw e;
    }
    return body;
  }

  window.api = {
    // ========== AUTHENTICATION ==========
    login: (email, password) => request('/api/auth/login', { method: 'POST', body: JSON.stringify({ email, password }) }),
    register: (payload) => request('/api/auth/register', { method: 'POST', body: JSON.stringify(payload) }),
    
    // ========== PROFILE & USER ==========
    getProfile: () => request('/api/profile', { method: 'GET' }),
    updateProfile: (data) => request('/api/profile', { method: 'PUT', body: JSON.stringify(data) }),
    getUser: (id) => request(`/api/users/${id}`, { method: 'GET' }),
    updateUser: (id, data) => request(`/api/users/${id}`, { method: 'PUT', body: JSON.stringify(data) }),
    
    // ========== COURSES & ENROLLMENT ==========
    getCourses: (page=1, limit=50) => request(`/api/courses?page=${page}&limit=${limit}`),
    getCourse: (id) => request(`/api/courses/${id}`, { method: 'GET' }),
    createCourse: (data) => request('/api/courses', { method: 'POST', body: JSON.stringify(data) }),
    updateCourse: (id, data) => request(`/api/courses/${id}`, { method: 'PUT', body: JSON.stringify(data) }),
    deleteCourse: (id) => request(`/api/courses/${id}`, { method: 'DELETE' }),
    
    createEnrollment: (studentId, courseId) => request('/api/enrollments', { method: 'POST', body: JSON.stringify({ student_id: studentId, course_id: courseId }) }),
    getEnrollmentsByStudent: (studentId) => request(`/api/enrollments/by-student/${studentId}`),
    getEnrollmentsByCourse: (courseId) => request(`/api/enrollments/by-course/${courseId}`),
    updateEnrollmentStatus: (id, status) => request(`/api/enrollments/${id}`, { method: 'PUT', body: JSON.stringify({ status }) }),
    deleteEnrollment: (id) => request(`/api/enrollments/${id}`, { method: 'DELETE' }),
    
    // ========== GRADES ==========
    getMyGrades: () => request('/api/student/grades', { method: 'GET' }),
    getGradesByStudent: (studentId) => request(`/api/grades/by-student/${studentId}`),
    getGradesByStudentCourse: (studentId, courseId) => request(`/api/grades/by-student/${studentId}/course/${courseId}`),
    recordGrade: (data) => request('/api/teacher/grades', { method: 'POST', body: JSON.stringify(data) }),
    updateGrade: (id, data) => request(`/api/grades/${id}`, { method: 'PUT', body: JSON.stringify(data) }),
    deleteGrade: (id) => request(`/api/grades/${id}`, { method: 'DELETE' }),
    getGradeAverages: (studentId) => request(`/api/grades/averages/${studentId}`, { method: 'GET' }),
    
    // ========== ATTENDANCE ==========
    getMyAttendance: () => request('/api/student/attendance', { method: 'GET' }),
    getAttendanceByStudent: (studentId) => request(`/api/attendance/by-student/${studentId}`),
    getAttendanceByStudentCourse: (studentId, courseId) => request(`/api/attendance/by-student/${studentId}/course/${courseId}`),
    markAttendance: (data) => request('/api/teacher/attendance', { method: 'POST', body: JSON.stringify(data) }),
    updateAttendance: (id, data) => request(`/api/attendance/${id}`, { method: 'PUT', body: JSON.stringify(data) }),
    deleteAttendance: (id) => request(`/api/attendance/${id}`, { method: 'DELETE' }),
    getAttendanceStats: (studentId) => request(`/api/attendance/stats/${studentId}`, { method: 'GET' }),
    
    // ========== ASSIGNMENTS ==========
    getMyAssignments: () => request('/api/student/assignments', { method: 'GET' }),
    getAssignmentsByCourse: (courseId) => request(`/api/assignments/by-course/${courseId}`),
    getAssignmentsByStudent: (studentId) => request(`/api/assignments/by-student/${studentId}`),
    createAssignment: (data) => request('/api/assignments', { method: 'POST', body: JSON.stringify(data) }),
    updateAssignment: (id, data) => request(`/api/assignments/${id}`, { method: 'PUT', body: JSON.stringify(data) }),
    deleteAssignment: (id) => request(`/api/assignments/${id}`, { method: 'DELETE' }),
    submitAssignment: (assignmentId, data) => request(`/api/assignments/${assignmentId}/submit`, { method: 'POST', body: JSON.stringify(data) }),
    gradeAssignment: (assignmentId, grade) => request(`/api/assignments/${assignmentId}/grade`, { method: 'POST', body: JSON.stringify({ grade }) }),
    
    // ========== ANNOUNCEMENTS ==========
    getAnnouncements: () => request('/api/announcements', { method: 'GET' }),
    getAnnouncement: (id) => request(`/api/announcements/${id}`, { method: 'GET' }),
    createAnnouncement: (data) => request('/api/announcements', { method: 'POST', body: JSON.stringify(data) }),
    updateAnnouncement: (id, data) => request(`/api/announcements/${id}`, { method: 'PUT', body: JSON.stringify(data) }),
    deleteAnnouncement: (id) => request(`/api/announcements/${id}`, { method: 'DELETE' }),
    
    // ========== NOTIFICATIONS ==========
    getNotifications: () => request('/api/notifications', { method: 'GET' }),
    markNotificationAsRead: (id) => request(`/api/notifications/${id}/read`, { method: 'POST' }),
    markAllNotificationsAsRead: () => request('/api/notifications/mark-all-read', { method: 'POST' }),
    deleteNotification: (id) => request(`/api/notifications/${id}`, { method: 'DELETE' }),
    clearAllNotifications: () => request('/api/notifications/clear-all', { method: 'POST' }),
    
    // ========== MESSAGES & CONVERSATIONS ==========
    getConversations: () => request('/api/messages/conversations', { method: 'GET' }),
    getConversation: (id) => request(`/api/messages/conversations/${id}`, { method: 'GET' }),
    getMessages: (conversationId) => request(`/api/messages/conversation/${conversationId}`, { method: 'GET' }),
    sendMessage: (data) => request('/api/messages', { method: 'POST', body: JSON.stringify(data) }),
    startConversation: (recipientId) => request('/api/messages/conversations', { method: 'POST', body: JSON.stringify({ recipient_id: recipientId }) }),
    
    // ========== PAYMENTS ==========
    getPayments: () => request('/api/payments', { method: 'GET' }),
    getPaymentHistory: () => request('/api/payments/history', { method: 'GET' }),
    makePayment: (data) => request('/api/payments', { method: 'POST', body: JSON.stringify(data) }),
    getPaymentStatus: (id) => request(`/api/payments/${id}`, { method: 'GET' }),
    
    // ========== TIMETABLES ==========
    getTimetables: () => request('/api/timetable', { method: 'GET' }),
    getTimetablesByCourse: (courseId) => request(`/api/timetable/course/${courseId}`),
    getTimetablesByTeacher: (teacherId) => request(`/api/timetable/teacher/${teacherId}`),
    getTimetablesByDay: (day) => request(`/api/timetable/day/${day}`),
    createTimetable: (data) => request('/api/timetable', { method: 'POST', body: JSON.stringify(data) }),
    updateTimetable: (id, data) => request(`/api/timetable/${id}`, { method: 'PUT', body: JSON.stringify(data) }),
    deleteTimetable: (id) => request(`/api/timetable/${id}`, { method: 'DELETE' }),
    
    // ========== TRANSCRIPTS ==========
    getTranscripts: () => request('/api/transcripts/student/:student_id', { method: 'GET' }),
    getTranscriptBySemester: (studentId, semester) => request(`/api/transcripts/student/${studentId}?semester=${semester}`),
    getTranscriptGPA: (studentId) => request(`/api/transcripts/gpa/${studentId}`),
    downloadTranscript: (studentId) => request(`/api/export/transcript/${studentId}`, { method: 'GET' }),
    
    // ========== ADMIN FEATURES ==========
    adminGetUsers: (page=1, limit=50) => request(`/api/admin/users?page=${page}&limit=${limit}`),
    adminCreateUser: (data) => request('/api/admin/users', { method: 'POST', body: JSON.stringify(data) }),
    adminDeleteUser: (id) => request(`/api/admin/users/${id}`, { method: 'DELETE' }),
    adminGetDashboard: () => request('/api/admin/dashboard', { method: 'GET' }),
    adminGetEnrollments: (status) => request(`/api/admin/enrollments${status ? '?status=' + status : ''}`, { method: 'GET' }),
    adminApproveEnrollment: (id) => request(`/api/admin/enrollments/${id}/approve`, { method: 'POST' }),
    adminRejectEnrollment: (id) => request(`/api/admin/enrollments/${id}/reject`, { method: 'POST' }),
    adminHealth: () => request('/api/admin/health', { method: 'GET' }),

    // ========== CONVENIENCE ALIASES FOR ADMIN DASHBOARD ==========
    getAllUsers: (page=1, limit=50) => request(`/api/admin/users?page=${page}&limit=${limit}`),
    createUser: (data) => request('/api/admin/users', { method: 'POST', body: JSON.stringify(data) }),
    deleteUser: (id) => request(`/api/admin/users/${id}`, { method: 'DELETE' }),
    getAllEnrollments: (status) => request(`/api/admin/enrollments${status ? '?status=' + status : ''}`, { method: 'GET' }),
    approveEnrollment: (id) => request(`/api/admin/enrollments/${id}/approve`, { method: 'POST' }),
    rejectEnrollment: (id) => request(`/api/admin/enrollments/${id}/reject`, { method: 'POST' }),
    getAdminHealth: () => request('/api/admin/health', { method: 'GET' }),
  };
})();
