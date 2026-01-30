import React, { useState, useEffect } from 'react'
import api from '../api'

export default function Payments() {
  const [payments, setPayments] = useState([])
  const [studentBalance, setStudentBalance] = useState(null)
  const [loading, setLoading] = useState(false)
  const [selectedStudent, setSelectedStudent] = useState(null)

  useEffect(() => {
    loadPayments()
  }, [])

  const loadPayments = async () => {
    setLoading(true)
    try {
      const res = await api.getAllPayments()
      setPayments(res.data || [])
    } catch (err) {
      console.error('Failed to load payments:', err)
    }
    setLoading(false)
  }

  const loadStudentBalance = async (studentId) => {
    try {
      const res = await api.getStudentBalance(studentId)
      setStudentBalance(res.data)
      setSelectedStudent(studentId)
    } catch (err) {
      console.error('Failed to load student balance:', err)
    }
  }

  const getStatusClass = (status) => {
    return `status-${status.toLowerCase()}`
  }

  const getStatusBadge = (status) => {
    const colors = {
      'paid': 'success',
      'pending': 'warning',
      'overdue': 'danger',
      'cancelled': 'secondary'
    }
    return colors[status.toLowerCase()] || 'info'
  }

  return (
    <div className="payments-container">
      <h2>Payment Management</h2>

      {studentBalance && (
        <div className="student-balance card">
          <h3>Student Balance</h3>
          <div className="balance-details">
            <div className="balance-item">
              <label>Total Due:</label>
              <span className="amount">${studentBalance.total_due?.toFixed(2)}</span>
            </div>
            <div className="balance-item">
              <label>Total Paid:</label>
              <span className="amount success">${studentBalance.total_paid?.toFixed(2)}</span>
            </div>
            <div className="balance-item">
              <label>Balance:</label>
              <span className="amount">${studentBalance.balance?.toFixed(2)}</span>
            </div>
          </div>
        </div>
      )}

      <div className="search-section">
        <label>Find Student Balance:</label>
        <input
          type="number"
          placeholder="Enter Student ID"
          onBlur={(e) => {
            if (e.target.value) {
              loadStudentBalance(parseInt(e.target.value))
            }
          }}
        />
      </div>

      {loading ? (
        <p>Loading payments...</p>
      ) : payments.length === 0 ? (
        <p className="no-data">No payments found</p>
      ) : (
        <div className="payments-table">
          <table>
            <thead>
              <tr>
                <th>Student ID</th>
                <th>Amount</th>
                <th>Status</th>
                <th>Description</th>
                <th>Due Date</th>
                <th>Created</th>
              </tr>
            </thead>
            <tbody>
              {payments.map((payment) => (
                <tr key={payment.id}>
                  <td>{payment.student_id}</td>
                  <td className="amount">${payment.amount?.toFixed(2)}</td>
                  <td>
                    <span className={`badge badge-${getStatusBadge(payment.status)}`}>
                      {payment.status}
                    </span>
                  </td>
                  <td>{payment.description}</td>
                  <td>{new Date(payment.due_date * 1000).toLocaleDateString()}</td>
                  <td>{new Date(payment.created_at).toLocaleDateString()}</td>
                </tr>
              ))}
            </tbody>
          </table>
        </div>
      )}
    </div>
  )
}
