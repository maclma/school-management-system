// Simple API helper for the demo frontend
// Attaches `window.api` with convenience methods
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
    login: (email, password) => request('/api/auth/login', { method: 'POST', body: JSON.stringify({ email, password }) }),
    register: (payload) => request('/api/auth/register', { method: 'POST', body: JSON.stringify(payload) }),
    getProfile: () => request('/api/profile', { method: 'GET' }),
    getCourses: (page=1, limit=50) => request(`/api/courses?page=${page}&limit=${limit}`),
    createEnrollment: (studentId, courseId) => request('/api/enrollments', { method: 'POST', body: JSON.stringify({ student_id: studentId, course_id: courseId }) }),
    getEnrollmentsByStudent: (studentId) => request(`/api/enrollments/by-student/${studentId}`),
    getGradesByStudent: (studentId) => request(`/api/grades/by-student/${studentId}`),
    getAttendanceByStudent: (studentId) => request(`/api/attendance/by-student/${studentId}`),
  };
})();
