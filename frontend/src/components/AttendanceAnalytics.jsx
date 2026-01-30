import { useState, useEffect } from 'react'
import api from '../../api'
import './AttendanceAnalytics.css'

export default function AttendanceAnalytics() {
  const [activeTab, setActiveTab] = useState('stats')
  const [courseId, setCourseId] = useState(1)
  const [studentId, setStudentId] = useState(1)
  const [threshold, setThreshold] = useState(75)
  
  const [stats, setStats] = useState(null)
  const [percentage, setPercentage] = useState(null)
  const [lowAttendanceStudents, setLowAttendanceStudents] = useState([])
  const [report, setReport] = useState(null)
  
  const [loading, setLoading] = useState(false)
  const [error, setError] = useState(null)

  const loadStats = async () => {
    setLoading(true)
    setError(null)
    try {
      const data = await api.getAttendanceStatsByCourse(courseId)
      setStats(data.data)
    } catch (err) {
      setError(err.message)
    } finally {
      setLoading(false)
    }
  }

  const loadPercentage = async () => {
    setLoading(true)
    setError(null)
    try {
      const data = await api.getStudentAttendancePercentage(studentId, courseId)
      setPercentage(data.data)
    } catch (err) {
      setError(err.message)
    } finally {
      setLoading(false)
    }
  }

  const loadLowAttendance = async () => {
    setLoading(true)
    setError(null)
    try {
      const data = await api.getStudentsWithLowAttendance(threshold, courseId)
      setLowAttendanceStudents(data.data?.students || [])
    } catch (err) {
      setError(err.message)
    } finally {
      setLoading(false)
    }
  }

  const loadReport = async () => {
    setLoading(true)
    setError(null)
    try {
      const data = await api.getAttendanceReport(courseId)
      setReport(data.data)
    } catch (err) {
      setError(err.message)
    } finally {
      setLoading(false)
    }
  }

  useEffect(() => {
    if (activeTab === 'stats') loadStats()
    else if (activeTab === 'percentage') loadPercentage()
    else if (activeTab === 'low-attendance') loadLowAttendance()
    else if (activeTab === 'report') loadReport()
  }, [activeTab, courseId])

  const getAttendanceColor = (percentage) => {
    if (percentage >= 90) return '#4CAF50'
    if (percentage >= 75) return '#FFC107'
    return '#F44336'
  }

  const AttendanceBar = ({ percentage }) => (
    <div className="attendance-bar-container">
      <div 
        className="attendance-bar" 
        style={{ 
          width: `${percentage}%`,
          backgroundColor: getAttendanceColor(percentage)
        }}
      />
      <span className="percentage">{percentage.toFixed(1)}%</span>
    </div>
  )

  return (
    <div className="attendance-analytics-container">
      <h2>Attendance Analytics</h2>

      <div className="tabs">
        <button 
          className={`tab ${activeTab === 'stats' ? 'active' : ''}`}
          onClick={() => setActiveTab('stats')}
        >
          Course Stats
        </button>
        <button 
          className={`tab ${activeTab === 'percentage' ? 'active' : ''}`}
          onClick={() => setActiveTab('percentage')}
        >
          Student %
        </button>
        <button 
          className={`tab ${activeTab === 'low-attendance' ? 'active' : ''}`}
          onClick={() => setActiveTab('low-attendance')}
        >
          Low Attendance
        </button>
        <button 
          className={`tab ${activeTab === 'report' ? 'active' : ''}`}
          onClick={() => setActiveTab('report')}
        >
          Full Report
        </button>
      </div>

      {error && <div className="error-message">{error}</div>}

      {/* Course Stats Tab */}
      {activeTab === 'stats' && (
        <div className="tab-content">
          <div className="input-group">
            <label>Course ID:</label>
            <input
              type="number"
              value={courseId}
              onChange={(e) => setCourseId(e.target.value)}
            />
            <button onClick={loadStats} disabled={loading}>
              {loading ? 'Loading...' : 'Load Stats'}
            </button>
          </div>

          {stats && (
            <div className="stats-grid">
              <div className="stat-card">
                <h3>Total Sessions</h3>
                <p className="stat-value">{stats.total_sessions}</p>
              </div>
              <div className="stat-card">
                <h3>Total Students</h3>
                <p className="stat-value">{stats.total_students}</p>
              </div>
              <div className="stat-card">
                <h3>Average Attendance</h3>
                <p className="stat-value">{stats.average_attendance?.toFixed(1)}%</p>
                <AttendanceBar percentage={stats.average_attendance || 0} />
              </div>
            </div>
          )}
        </div>
      )}

      {/* Student Attendance % Tab */}
      {activeTab === 'percentage' && (
        <div className="tab-content">
          <div className="input-group">
            <input
              type="number"
              placeholder="Student ID"
              value={studentId}
              onChange={(e) => setStudentId(e.target.value)}
            />
            <input
              type="number"
              placeholder="Course ID"
              value={courseId}
              onChange={(e) => setCourseId(e.target.value)}
            />
            <button onClick={loadPercentage} disabled={loading}>
              {loading ? 'Loading...' : 'Get Percentage'}
            </button>
          </div>

          {percentage && (
            <div className="percentage-card">
              <h3>Student Attendance</h3>
              <div className="percentage-stat">
                <p>Student ID: {percentage.student_id}</p>
                <p>Course ID: {percentage.course_id}</p>
              </div>
              <AttendanceBar percentage={percentage.attendance_percentage || 0} />
              <div className="details">
                <p>Classes Attended: {percentage.classes_attended}</p>
                <p>Total Classes: {percentage.total_classes}</p>
              </div>
            </div>
          )}
        </div>
      )}

      {/* Low Attendance Tab */}
      {activeTab === 'low-attendance' && (
        <div className="tab-content">
          <div className="input-group">
            <input
              type="number"
              placeholder="Threshold (%)"
              value={threshold}
              onChange={(e) => setThreshold(e.target.value)}
            />
            <input
              type="number"
              placeholder="Course ID (optional)"
              value={courseId}
              onChange={(e) => setCourseId(e.target.value)}
            />
            <button onClick={loadLowAttendance} disabled={loading}>
              {loading ? 'Loading...' : 'Find Low Attendance'}
            </button>
          </div>

          {lowAttendanceStudents.length > 0 && (
            <div className="low-attendance-list">
              <h3>Students Below {threshold}% Attendance</h3>
              <table>
                <thead>
                  <tr>
                    <th>Student ID</th>
                    <th>Name</th>
                    <th>Email</th>
                    <th>Attendance</th>
                  </tr>
                </thead>
                <tbody>
                  {lowAttendanceStudents.map((student, idx) => (
                    <tr key={idx}>
                      <td>{student.student_id}</td>
                      <td>{student.name}</td>
                      <td>{student.email}</td>
                      <td>
                        <AttendanceBar percentage={student.attendance_percentage} />
                      </td>
                    </tr>
                  ))}
                </tbody>
              </table>
            </div>
          )}
        </div>
      )}

      {/* Full Report Tab */}
      {activeTab === 'report' && (
        <div className="tab-content">
          <div className="input-group">
            <input
              type="number"
              placeholder="Course ID"
              value={courseId}
              onChange={(e) => setCourseId(e.target.value)}
            />
            <button onClick={loadReport} disabled={loading}>
              {loading ? 'Loading...' : 'Generate Report'}
            </button>
          </div>

          {report && (
            <div className="full-report">
              <h3>Attendance Report - {report.course_name}</h3>
              
              <div className="report-section">
                <h4>Overall Statistics</h4>
                <div className="stats-grid">
                  <div className="stat-card">
                    <p>Total Sessions: {report.stats?.total_sessions}</p>
                  </div>
                  <div className="stat-card">
                    <p>Total Students: {report.stats?.total_students}</p>
                  </div>
                  <div className="stat-card">
                    <p>Average Attendance: {report.stats?.average_attendance?.toFixed(1)}%</p>
                  </div>
                </div>
              </div>

              {report.students_below_80pct?.length > 0 && (
                <div className="report-section">
                  <h4>Students Below 80% Attendance</h4>
                  <ul>
                    {report.students_below_80pct.map((student, idx) => (
                      <li key={idx}>
                        {student.name} ({student.attendance_percentage?.toFixed(1)}%)
                      </li>
                    ))}
                  </ul>
                </div>
              )}

              <p className="report-generated">
                Report Generated: {new Date(report.report_generated_at).toLocaleString()}
              </p>
            </div>
          )}
        </div>
      )}
    </div>
  )
}
