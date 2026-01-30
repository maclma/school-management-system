import { useState } from 'react'
import api from '../../api'
import './ReportsExport.css'

export default function ReportsExport() {
  const [reportType, setReportType] = useState('payments')
  const [filters, setFilters] = useState({})
  const [loading, setLoading] = useState(false)
  const [error, setError] = useState(null)
  const [success, setSuccess] = useState(false)

  const handleFilterChange = (e) => {
    const { name, value } = e.target
    setFilters(prev => ({ ...prev, [name]: value }))
  }

  const downloadCSV = (data, filename) => {
    const blob = new Blob([data], { type: 'text/csv' })
    const url = window.URL.createObjectURL(blob)
    const link = document.createElement('a')
    link.href = url
    link.download = filename
    document.body.appendChild(link)
    link.click()
    document.body.removeChild(link)
    window.URL.revokeObjectURL(url)
  }

  const handleExport = async (e) => {
    e.preventDefault()
    setLoading(true)
    setError(null)
    setSuccess(false)

    try {
      let response
      let filename
      
      switch (reportType) {
        case 'payments':
          response = await api.exportPaymentsCSV(
            filters.studentId || 0,
            filters.status || ''
          )
          filename = `payments_${new Date().toISOString().split('T')[0]}.csv`
          break
        case 'grades':
          response = await api.exportGradesCSV(filters.courseId)
          filename = `grades_${new Date().toISOString().split('T')[0]}.csv`
          break
        case 'attendance':
          response = await api.exportAttendanceCSV(filters.courseId)
          filename = `attendance_${new Date().toISOString().split('T')[0]}.csv`
          break
        case 'transcript':
          response = await api.exportStudentTranscriptCSV(filters.studentId)
          filename = `transcript_student_${filters.studentId}.csv`
          break
        case 'enrollments':
          response = await api.exportEnrollmentsCSV(filters.courseId)
          filename = `enrollments_${new Date().toISOString().split('T')[0]}.csv`
          break
        default:
          throw new Error('Invalid report type')
      }

      if (response && response.data) {
        downloadCSV(response.data, filename)
        setSuccess(true)
      }
    } catch (err) {
      setError(err.message || 'Export failed')
    } finally {
      setLoading(false)
    }
  }

  return (
    <div className="reports-export-container">
      <h2>Export Reports</h2>

      <form onSubmit={handleExport} className="export-form">
        <div className="report-type-selector">
          <label>Report Type:</label>
          <select value={reportType} onChange={(e) => setReportType(e.target.value)}>
            <option value="payments">Payments Report</option>
            <option value="grades">Grades Report</option>
            <option value="attendance">Attendance Report</option>
            <option value="transcript">Student Transcript</option>
            <option value="enrollments">Enrollments Report</option>
          </select>
        </div>

        {/* Dynamic Filters based on report type */}
        <div className="filters">
          {reportType === 'payments' && (
            <>
              <input
                type="number"
                name="studentId"
                placeholder="Student ID (optional)"
                value={filters.studentId || ''}
                onChange={handleFilterChange}
              />
              <select name="status" value={filters.status || ''} onChange={handleFilterChange}>
                <option value="">All Statuses</option>
                <option value="pending">Pending</option>
                <option value="paid">Paid</option>
                <option value="overdue">Overdue</option>
                <option value="cancelled">Cancelled</option>
              </select>
            </>
          )}

          {(reportType === 'grades' || reportType === 'attendance' || reportType === 'enrollments') && (
            <input
              type="number"
              name="courseId"
              placeholder="Course ID (required)"
              required
              value={filters.courseId || ''}
              onChange={handleFilterChange}
            />
          )}

          {reportType === 'transcript' && (
            <input
              type="number"
              name="studentId"
              placeholder="Student ID (required)"
              required
              value={filters.studentId || ''}
              onChange={handleFilterChange}
            />
          )}
        </div>

        <button type="submit" disabled={loading}>
          {loading ? 'Exporting...' : 'Export as CSV'}
        </button>
      </form>

      {error && <div className="error-message">{error}</div>}
      {success && <div className="success-message">✓ Report exported successfully!</div>}

      <div className="report-info">
        <h3>Report Descriptions</h3>
        <ul>
          <li><strong>Payments Report:</strong> All payment records with amount, status, due date</li>
          <li><strong>Grades Report:</strong> Student grades for a course with scores and letter grades</li>
          <li><strong>Attendance Report:</strong> Attendance records with present/absent status and remarks</li>
          <li><strong>Student Transcript:</strong> Academic transcript with GPA and credit information</li>
          <li><strong>Enrollments Report:</strong> Course enrollment records with student and status</li>
        </ul>
      </div>

      <div className="report-features">
        <h3>Features</h3>
        <ul>
          <li>✓ Filter by student, course, or status</li>
          <li>✓ Download CSV files for spreadsheet software</li>
          <li>✓ Generate reports for compliance and analysis</li>
          <li>✓ Include all relevant data fields</li>
        </ul>
      </div>
    </div>
  )
}
