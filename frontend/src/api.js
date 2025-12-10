// small API wrapper used by React app
const base = '/api'

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
  recordGrade: (payload) => request('/grades', { method:'POST', body: JSON.stringify(payload) }),
  getAttendanceByStudent: (studentId) => request(`/attendance/by-student/${studentId}`),
  recordAttendance: (payload) => request('/attendance', { method:'POST', body: JSON.stringify(payload) }),
  getAdminStats: () => request('/admin/dashboard'),
  getAdminUsers: () => request('/admin/users'),
  createUserAdmin: (payload) => request('/admin/users', { method:'POST', body: JSON.stringify(payload) }),
  deleteUser: (userId) => request(`/admin/users/${userId}`, { method:'DELETE' }),
}
