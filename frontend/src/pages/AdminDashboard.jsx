import React, { useEffect, useState, useRef } from 'react'
import api from '../api'
import { showToast, setLoading } from '../ui'

export default function AdminDashboard() {
  const [stats, setStats] = useState(null)
  const [users, setUsers] = useState([])
  const [courses, setCourses] = useState([])
  const [enrollments, setEnrollments] = useState([])
  const [tab, setTab] = useState('stats')
  const [showModal, setShowModal] = useState(false)
  const [search, setSearch] = useState('')
  const [loading, setLoadingState] = useState(false)

  const emailRef = useRef(); const fnRef = useRef(); const lnRef = useRef()
  const passwordRef = useRef(); const roleRef = useRef(); const btnRef = useRef()
  const courseNameRef = useRef(); const courseDeptRef = useRef(); const courseDescRef = useRef(); const courseCodeRef = useRef()

  useEffect(() => { loadStats(); loadUsers(); loadCourses(); loadEnrollments() }, [])

  async function loadStats() { try { const s = await api.getAdminStats(); setStats(s.data || s) } catch (e) { console.error('Failed to load admin stats:', e) } }

  async function loadUsers() { try { const u = await api.getAdminUsers(); setUsers(u.data || u || []) } catch (e) { console.error('Failed to load users:', e) } }

  async function loadCourses() { try { const c = await api.getCourses(); setCourses(Array.isArray(c) ? c : c.data || []) } catch (e) { console.error('Failed to load courses:', e) } }

  async function loadEnrollments() { try { const e = await api.getAdminEnrollments(); setEnrollments(Array.isArray(e) ? e : e.data || []) } catch (e) { console.error('Failed to load enrollments:', e) } }

  async function createUser(e) {
    e.preventDefault(); setLoading(btnRef, true, 'Creating...')
    try {
      const payload = { email: emailRef.current.value, first_name: fnRef.current.value, last_name: lnRef.current.value, password: passwordRef.current.value, role: roleRef.current.value }
      await api.createUserAdmin(payload); showToast('User created', 'success'); setShowModal(false); emailRef.current.value = ''; fnRef.current.value = ''; lnRef.current.value = ''; passwordRef.current.value = ''; await loadUsers()
    } catch (e) { showToast(e.message || 'Failed to create user', 'error') } finally { setLoading(btnRef, false) }
  }

  async function deleteUser(userId) {
    if (!confirm('Are you sure? This cannot be undone.')) return
    try { await api.deleteUser(userId); showToast('User deleted', 'success'); await loadUsers() } catch (e) { showToast(e.message || 'Failed to delete user', 'error') }
  }

  async function createCourse(e) {
    e.preventDefault(); setLoading(btnRef, true, 'Creating...')
    try {
      const payload = { name: courseNameRef.current.value, course_code: courseCodeRef.current.value, department: courseDeptRef.current.value, description: courseDescRef.current.value }
      await api.createCourse(payload); showToast('Course created', 'success'); setShowModal(false); courseNameRef.current.value = ''; courseDeptRef.current.value = ''; courseDescRef.current.value = ''; courseCodeRef.current.value = ''; await loadCourses()
    } catch (e) { showToast(e.message || 'Failed to create course', 'error') } finally { setLoading(btnRef, false) }
  }

  async function deleteCourse(courseId) {
    if (!confirm('Are you sure? This will remove the course.')) return
    try { await api.deleteCourse(courseId); showToast('Course deleted', 'success'); await loadCourses() } catch (e) { showToast(e.message || 'Failed to delete course', 'error') }
  }

  async function approveEnrollment(enrollmentId) {
    try { await api.approveEnrollment(enrollmentId); showToast('Enrollment approved', 'success'); await loadEnrollments() } catch (e) { showToast(e.message || 'Failed to approve', 'error') }
  }

  async function rejectEnrollment(enrollmentId) {
    try { await api.rejectEnrollment(enrollmentId); showToast('Enrollment rejected', 'success'); await loadEnrollments() } catch (e) { showToast(e.message || 'Failed to reject', 'error') }
  }

  const filteredUsers = users.filter(u => u.email.toLowerCase().includes(search.toLowerCase()) || u.first_name.toLowerCase().includes(search.toLowerCase()) || u.last_name.toLowerCase().includes(search.toLowerCase()))

  return (
    <main className="container">
      <div className="card">
        <h1>Admin Dashboard</h1>
        <p className="muted-small">System overview and user management</p>

        <div className="tabs">
          <button className={`tab ${tab === 'stats' ? 'active' : ''}`} onClick={() => setTab('stats')}>Stats</button>
          <button className={`tab ${tab === 'users' ? 'active' : ''}`} onClick={() => setTab('users')}>Users</button>
          <button className={`tab ${tab === 'courses' ? 'active' : ''}`} onClick={() => setTab('courses')}>Courses</button>
          <button className={`tab ${tab === 'enrollments' ? 'active' : ''}`} onClick={() => setTab('enrollments')}>Enrollments</button>
          <button className={`tab ${tab === 'create' ? 'active' : ''}`} onClick={() => setTab('create')}>Create User</button>
        </div>

        {tab === 'stats' && stats && (
          <div className="grid">
            <div className="stat-tile"><div className="stat-label">Total Users</div><div className="stat-value">{stats.total_users || 0}</div></div>
            <div className="stat-tile"><div className="stat-label">Total Courses</div><div className="stat-value">{stats.total_courses || 0}</div></div>
            <div className="stat-tile"><div className="stat-label">Total Enrollments</div><div className="stat-value">{stats.total_enrollments || 0}</div></div>
            <div className="stat-tile"><div className="stat-label">Active Students</div><div className="stat-value">{stats.active_students || 0}</div></div>
          </div>
        )}

        {tab === 'users' && (
          <div className="panel">
            <h2>Users</h2>
            <div style={{ marginBottom: '12px' }}><input type="text" placeholder="Search by email or name..." value={search} onChange={(e) => setSearch(e.target.value)} style={{ maxWidth: '400px' }} /></div>
            {filteredUsers.length === 0 ? <p className="muted-small">{search ? 'No matching users' : 'No users'}</p> : (
              <div className="table-container">
                <table className="simple-table">
                  <thead><tr><th>ID</th><th>Email</th><th>Name</th><th>Role</th><th>Status</th><th>Action</th></tr></thead>
                  <tbody>{filteredUsers.map(u => (<tr key={u.id}><td>{u.id}</td><td>{u.email}</td><td>{u.first_name} {u.last_name}</td><td><span className="badge">{u.role}</span></td><td><select value={u.is_active ? 'active' : 'inactive'} onChange={async (e) => { try { await api.updateUserStatus(u.id, e.target.value === 'active'); showToast('Status updated', 'success'); await loadUsers() } catch (e) { showToast('Failed to update', 'error') } }}><option value="active">Active</option><option value="inactive">Inactive</option></select></td><td><button className="small danger" onClick={() => deleteUser(u.id)}>Delete</button></td></tr>))}</tbody>
                </table>
              </div>
            )}
          </div>
        )}

        {tab === 'courses' && (
          <div className="panel">
            <h2>Courses</h2>
            <button className="secondary" onClick={() => { setShowModal(true); }} style={{ marginBottom: '12px' }}>+ Create Course</button>
            {courses.length === 0 ? <p className="muted-small">No courses</p> : (
              <div className="table-container">
                <table className="simple-table">
                  <thead><tr><th>ID</th><th>Name</th><th>Code</th><th>Department</th><th>Description</th><th>Action</th></tr></thead>
                  <tbody>{courses.map(c => (<tr key={c.id}><td>{c.id}</td><td>{c.name}</td><td>{c.course_code || 'N/A'}</td><td>{c.department}</td><td>{c.description}</td><td><button className="small danger" onClick={() => deleteCourse(c.id)}>Delete</button></td></tr>))}</tbody>
                </table>
              </div>
            )}
            {showModal && (
              <div className="modal-overlay" onClick={() => setShowModal(false)}>
                <div className="modal-content" onClick={(e) => e.stopPropagation()}>
                  <h2>Create New Course</h2>
                  <form onSubmit={createCourse}>
                    <label>Course Title</label><input ref={courseNameRef} required />
                    <label>Course Code</label><input ref={courseCodeRef} />
                    <label>Department</label><input ref={courseDeptRef} required />
                    <label>Description</label><textarea ref={courseDescRef} style={{ minHeight: '100px' }}></textarea>
                    <button ref={btnRef} type="submit" style={{ marginTop: '12px' }}>Create Course</button>
                    <button type="button" onClick={() => setShowModal(false)} style={{ marginTop: '12px', marginLeft: '8px' }}>Cancel</button>
                  </form>
                </div>
              </div>
            )}
          </div>
        )}

        {tab === 'enrollments' && (
          <div className="panel">
            <h2>Enrollment Approvals</h2>
            {enrollments.length === 0 ? <p className="muted-small">No enrollments</p> : (
              <div className="table-container">
                <table className="simple-table">
                  <thead><tr><th>ID</th><th>Student ID</th><th>Course ID</th><th>Status</th><th>Actions</th></tr></thead>
                  <tbody>{enrollments.map(e => (
                    <tr key={e.id}>
                      <td>{e.id}</td>
                      <td>{e.student_id}</td>
                      <td>{e.course_id}</td>
                      <td><span className={`badge ${e.status === 'pending' ? 'warning' : e.status === 'approved' ? 'success' : 'danger'}`}>{e.status}</span></td>
                      <td>
                        {e.status === 'pending' && (
                          <>
                            <button className="small success" onClick={() => approveEnrollment(e.id)}>Approve</button>
                            <button className="small danger" onClick={() => rejectEnrollment(e.id)} style={{ marginLeft: '4px' }}>Reject</button>
                          </>
                        )}
                      </td>
                    </tr>
                  ))}</tbody>
                </table>
              </div>
            )}
          </div>
        )}

        {tab === 'create' && (
          <div className="panel">
            <h2>Create New User</h2>
            <form onSubmit={createUser}>
              <label>Email</label><input ref={emailRef} type="email" required />
              <label>First Name</label><input ref={fnRef} required />
              <label>Last Name</label><input ref={lnRef} required />
              <label>Password</label><input ref={passwordRef} type="password" required />
              <label>Role</label><select ref={roleRef} required><option value="student">Student</option><option value="teacher">Teacher</option><option value="admin">Admin</option></select>
              <button ref={btnRef} type="submit" style={{ marginTop: '12px' }}>Create User</button>
            </form>
          </div>
        )}
      </div>
    </main>
  )
}

