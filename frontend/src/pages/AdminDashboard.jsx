import React, { useEffect, useState, useRef } from 'react'
import api from '../api'
import { showToast, setLoading } from '../ui'

export default function AdminDashboard() {
  const [stats, setStats] = useState(null)
  const [users, setUsers] = useState([])
  const [tab, setTab] = useState('stats')
  const [showModal, setShowModal] = useState(false)
  const [search, setSearch] = useState('')

  const emailRef = useRef(); const fnRef = useRef(); const lnRef = useRef()
  const passwordRef = useRef(); const roleRef = useRef(); const btnRef = useRef()

  useEffect(() => { loadStats(); loadUsers() }, [])

  async function loadStats() { try { const s = await api.getAdminStats(); setStats(s.data || s) } catch (e) { console.error('Failed to load admin stats:', e) } }

  async function loadUsers() { try { const u = await api.getAdminUsers(); setUsers(u.data || u || []) } catch (e) { console.error('Failed to load users:', e) } }

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

  const filteredUsers = users.filter(u => u.email.toLowerCase().includes(search.toLowerCase()) || u.first_name.toLowerCase().includes(search.toLowerCase()) || u.last_name.toLowerCase().includes(search.toLowerCase()))

  return (
    <main className="container">
      <div className="card">
        <h1>Admin Dashboard</h1>
        <p className="muted-small">System overview and user management</p>

        <div className="tabs">
          <button className={`tab ${tab === 'stats' ? 'active' : ''}`} onClick={() => setTab('stats')}>Stats</button>
          <button className={`tab ${tab === 'users' ? 'active' : ''}`} onClick={() => setTab('users')}>Users</button>
          <button className={`tab ${tab === 'enrollments' ? 'active' : ''}`} onClick={() => window.navigate('/admin/enrollments')}>Enrollments</button>
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

