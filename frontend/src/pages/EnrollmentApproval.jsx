import React, { useEffect, useState, useRef } from 'react'
import api from '../api'
import { showToast, setLoading } from '../ui'

export default function EnrollmentApproval() {
  const [enrollments, setEnrollments] = useState([])
  const [filter, setFilter] = useState('pending')
  const [search, setSearch] = useState('')

  useEffect(() => { loadEnrollments() }, [])

  async function loadEnrollments() {
    try {
      const res = await api.request('/admin/enrollments', { method: 'GET' })
      setEnrollments(res.data || res || [])
    } catch (e) {
      showToast('Failed to load enrollments', 'error')
    }
  }

  async function approveEnrollment(enrollmentId) {
    try {
      await api.request(`/admin/enrollments/${enrollmentId}/approve`, { method: 'POST' })
      showToast('Enrollment approved', 'success')
      await loadEnrollments()
    } catch (e) {
      showToast(e.message || 'Failed to approve enrollment', 'error')
    }
  }

  async function rejectEnrollment(enrollmentId) {
    if (!confirm('Reject this enrollment? This cannot be undone.')) return
    try {
      await api.request(`/admin/enrollments/${enrollmentId}/reject`, { method: 'POST' })
      showToast('Enrollment rejected', 'success')
      await loadEnrollments()
    } catch (e) {
      showToast(e.message || 'Failed to reject enrollment', 'error')
    }
  }

  const filtered = enrollments.filter(e => {
    const matchStatus = filter === 'all' || e.status === filter
    const matchSearch = search === '' || 
      e.student_name.toLowerCase().includes(search.toLowerCase()) ||
      e.course_name.toLowerCase().includes(search.toLowerCase()) ||
      e.student_email.toLowerCase().includes(search.toLowerCase())
    return matchStatus && matchSearch
  })

  const stats = {
    pending: enrollments.filter(e => e.status === 'pending').length,
    approved: enrollments.filter(e => e.status === 'approved').length,
    rejected: enrollments.filter(e => e.status === 'rejected').length,
  }

  return (
    <main className="container">
      <div className="card">
        <h1>Enrollment Approvals</h1>
        <p className="muted-small">Review and manage student enrollment requests</p>

        <div style={{ display: 'grid', gridTemplateColumns: 'repeat(auto-fit, minmax(120px, 1fr))', gap: '12px', marginBottom: '20px' }}>
          <div style={{ padding: '12px', border: '1px solid #ddd', borderRadius: '4px', textAlign: 'center' }}>
            <div style={{ fontSize: '24px', fontWeight: 'bold', color: '#f39c12' }}>{stats.pending}</div>
            <div style={{ fontSize: '12px', color: '#666' }}>Pending</div>
          </div>
          <div style={{ padding: '12px', border: '1px solid #ddd', borderRadius: '4px', textAlign: 'center' }}>
            <div style={{ fontSize: '24px', fontWeight: 'bold', color: '#27ae60' }}>{stats.approved}</div>
            <div style={{ fontSize: '12px', color: '#666' }}>Approved</div>
          </div>
          <div style={{ padding: '12px', border: '1px solid #ddd', borderRadius: '4px', textAlign: 'center' }}>
            <div style={{ fontSize: '24px', fontWeight: 'bold', color: '#e74c3c' }}>{stats.rejected}</div>
            <div style={{ fontSize: '12px', color: '#666' }}>Rejected</div>
          </div>
        </div>

        <div style={{ display: 'flex', gap: '8px', marginBottom: '16px', flexWrap: 'wrap' }}>
          <button
            className={`tab ${filter === 'pending' ? 'active' : ''}`}
            onClick={() => setFilter('pending')}
          >
            Pending ({stats.pending})
          </button>
          <button
            className={`tab ${filter === 'approved' ? 'active' : ''}`}
            onClick={() => setFilter('approved')}
          >
            Approved ({stats.approved})
          </button>
          <button
            className={`tab ${filter === 'rejected' ? 'active' : ''}`}
            onClick={() => setFilter('rejected')}
          >
            Rejected ({stats.rejected})
          </button>
          <button
            className={`tab ${filter === 'all' ? 'active' : ''}`}
            onClick={() => setFilter('all')}
          >
            All
          </button>
        </div>

        <div style={{ marginBottom: '12px' }}>
          <input
            type="text"
            placeholder="Search by student name, email, or course..."
            value={search}
            onChange={(e) => setSearch(e.target.value)}
            style={{ maxWidth: '500px' }}
          />
        </div>

        {filtered.length === 0 ? (
          <p className="muted-small">{search ? 'No matching enrollments' : 'No enrollments'}</p>
        ) : (
          <div className="table-container">
            <table className="simple-table">
              <thead>
                <tr>
                  <th>ID</th>
                  <th>Student</th>
                  <th>Email</th>
                  <th>Course</th>
                  <th>Status</th>
                  <th>Enrolled Date</th>
                  <th>Actions</th>
                </tr>
              </thead>
              <tbody>
                {filtered.map(e => (
                  <tr key={e.id}>
                    <td>{e.id}</td>
                    <td>{e.student_name}</td>
                    <td>{e.student_email}</td>
                    <td>{e.course_name}</td>
                    <td>
                      <span
                        className="badge"
                        style={{
                          backgroundColor:
                            e.status === 'pending'
                              ? '#f39c12'
                              : e.status === 'approved'
                              ? '#27ae60'
                              : '#e74c3c',
                          color: 'white',
                        }}
                      >
                        {e.status.charAt(0).toUpperCase() + e.status.slice(1)}
                      </span>
                    </td>
                    <td>{new Date(e.created_at).toLocaleDateString()}</td>
                    <td style={{ display: 'flex', gap: '4px' }}>
                      {e.status === 'pending' && (
                        <>
                          <button
                            className="small success"
                            onClick={() => approveEnrollment(e.id)}
                          >
                            Approve
                          </button>
                          <button
                            className="small danger"
                            onClick={() => rejectEnrollment(e.id)}
                          >
                            Reject
                          </button>
                        </>
                      )}
                      {e.status !== 'pending' && (
                        <span style={{ fontSize: '12px', color: '#666' }}>
                          {e.status === 'approved' ? '✓' : '✗'} {e.status}
                        </span>
                      )}
                    </td>
                  </tr>
                ))}
              </tbody>
            </table>
          </div>
        )}
      </div>
    </main>
  )
}
