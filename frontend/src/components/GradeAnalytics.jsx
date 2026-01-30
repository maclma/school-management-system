import { useState, useEffect } from 'react'
import api from '../../api'
import './GradeAnalytics.css'

export default function GradeAnalytics() {
  const [activeTab, setActiveTab] = useState('record')
  const [courseId, setCourseId] = useState(1)
  const [studentId, setStudentId] = useState(1)
  
  const [courseAverage, setCourseAverage] = useState(null)
  const [gradeDistribution, setGradeDistribution] = useState(null)
  const [studentStats, setStudentStats] = useState(null)
  
  const [loading, setLoading] = useState(false)
  const [error, setError] = useState(null)
  const [success, setSuccess] = useState(false)

  const [formData, setFormData] = useState({
    student_id: '',
    course_id: '',
    score: '',
    max_score: 100,
    remarks: '',
    graded_by: ''
  })

  const loadCourseAverage = async () => {
    setLoading(true)
    setError(null)
    try {
      const data = await api.getCourseAverageGrade(courseId)
      setCourseAverage(data.data)
    } catch (err) {
      setError(err.message)
    } finally {
      setLoading(false)
    }
  }

  const loadGradeDistribution = async () => {
    setLoading(true)
    setError(null)
    try {
      const data = await api.getGradeDistribution(courseId)
      setGradeDistribution(data.data)
    } catch (err) {
      setError(err.message)
    } finally {
      setLoading(false)
    }
  }

  const loadStudentStats = async () => {
    setLoading(true)
    setError(null)
    try {
      const data = await api.getStudentGradeStats(studentId)
      setStudentStats(data.data)
    } catch (err) {
      setError(err.message)
    } finally {
      setLoading(false)
    }
  }

  const handleRecordGrade = async (e) => {
    e.preventDefault()
    setLoading(true)
    setError(null)
    setSuccess(false)

    try {
      if (!formData.student_id || !formData.course_id || !formData.score) {
        throw new Error('Student ID, Course ID, and Score are required')
      }

      await api.recordGradeWithAutoCalc({
        student_id: parseInt(formData.student_id),
        course_id: parseInt(formData.course_id),
        score: parseFloat(formData.score),
        max_score: parseFloat(formData.max_score),
        remarks: formData.remarks,
        graded_by: parseInt(formData.graded_by)
      })

      setSuccess(true)
      setFormData({ student_id: '', course_id: '', score: '', max_score: 100, remarks: '', graded_by: '' })
      
      // Reload stats
      setTimeout(() => {
        loadStudentStats()
        loadGradeDistribution()
      }, 1000)
    } catch (err) {
      setError(err.message)
    } finally {
      setLoading(false)
    }
  }

  const handleInputChange = (e) => {
    const { name, value } = e.target
    setFormData(prev => ({ ...prev, [name]: value }))
  }

  const GradeDistributionChart = ({ distribution, total }) => {
    if (!distribution) return null

    const getGradeColor = (grade) => {
      const colors = { A: '#4CAF50', B: '#8BC34A', C: '#FFC107', D: '#FF9800', F: '#F44336' }
      return colors[grade] || '#999'
    }

    const maxCount = Math.max(...Object.values(distribution))

    return (
      <div className="distribution-chart">
        <h4>Grade Distribution</h4>
        {Object.entries(distribution).map(([grade, count]) => (
          <div key={grade} className="distribution-bar">
            <span className="grade-label">{grade}</span>
            <div className="bar-container">
              <div
                className="bar"
                style={{
                  width: `${(count / maxCount) * 100}%`,
                  backgroundColor: getGradeColor(grade)
                }}
              />
            </div>
            <span className="count">{count}</span>
          </div>
        ))}
        <p className="total-students">Total: {total} students</p>
      </div>
    )
  }

  return (
    <div className="grade-analytics-container">
      <h2>Grade Analytics</h2>

      <div className="tabs">
        <button 
          className={`tab ${activeTab === 'record' ? 'active' : ''}`}
          onClick={() => setActiveTab('record')}
        >
          Record Grade
        </button>
        <button 
          className={`tab ${activeTab === 'average' ? 'active' : ''}`}
          onClick={() => setActiveTab('average')}
        >
          Course Average
        </button>
        <button 
          className={`tab ${activeTab === 'distribution' ? 'active' : ''}`}
          onClick={() => setActiveTab('distribution')}
        >
          Grade Distribution
        </button>
        <button 
          className={`tab ${activeTab === 'student-stats' ? 'active' : ''}`}
          onClick={() => setActiveTab('student-stats')}
        >
          Student Stats
        </button>
      </div>

      {error && <div className="error-message">{error}</div>}
      {success && <div className="success-message">✓ Grade recorded successfully!</div>}

      {/* Record Grade Tab */}
      {activeTab === 'record' && (
        <div className="tab-content">
          <h3>Record New Grade (Auto-Calculated)</h3>
          <form onSubmit={handleRecordGrade} className="grade-form">
            <div className="form-group">
              <label>Student ID *</label>
              <input
                type="number"
                name="student_id"
                required
                value={formData.student_id}
                onChange={handleInputChange}
                placeholder="Student ID"
              />
            </div>

            <div className="form-group">
              <label>Course ID *</label>
              <input
                type="number"
                name="course_id"
                required
                value={formData.course_id}
                onChange={handleInputChange}
                placeholder="Course ID"
              />
            </div>

            <div className="form-group">
              <label>Score *</label>
              <input
                type="number"
                name="score"
                required
                step="0.1"
                value={formData.score}
                onChange={handleInputChange}
                placeholder="Score"
              />
            </div>

            <div className="form-group">
              <label>Max Score</label>
              <input
                type="number"
                name="max_score"
                step="0.1"
                value={formData.max_score}
                onChange={handleInputChange}
              />
            </div>

            <div className="form-group">
              <label>Remarks</label>
              <textarea
                name="remarks"
                value={formData.remarks}
                onChange={handleInputChange}
                placeholder="Optional comments"
                rows="3"
              />
            </div>

            <div className="form-group">
              <label>Graded By (User ID)</label>
              <input
                type="number"
                name="graded_by"
                value={formData.graded_by}
                onChange={handleInputChange}
                placeholder="Teacher/Admin ID"
              />
            </div>

            <button type="submit" disabled={loading}>
              {loading ? 'Recording...' : 'Record Grade'}
            </button>
          </form>

          <div className="auto-calc-info">
            <h4>Auto-Calculation Information</h4>
            <ul>
              <li>✓ Letter grade automatically calculated from score</li>
              <li>✓ A: 90-100, B: 80-89, C: 70-79, D: 60-69, F: 0-59</li>
              <li>✓ Student transcript automatically updated</li>
              <li>✓ Email notification sent to student</li>
            </ul>
          </div>
        </div>
      )}

      {/* Course Average Tab */}
      {activeTab === 'average' && (
        <div className="tab-content">
          <div className="input-group">
            <input
              type="number"
              placeholder="Course ID"
              value={courseId}
              onChange={(e) => setCourseId(e.target.value)}
            />
            <button onClick={loadCourseAverage} disabled={loading}>
              {loading ? 'Loading...' : 'Get Average'}
            </button>
          </div>

          {courseAverage && (
            <div className="average-card">
              <h3>{courseAverage.course_name}</h3>
              <div className="stats-grid">
                <div className="stat">
                  <span className="label">Average Score</span>
                  <span className="value">{courseAverage.average_score?.toFixed(2)}</span>
                </div>
                <div className="stat">
                  <span className="label">Total Grades</span>
                  <span className="value">{courseAverage.total_grades}</span>
                </div>
                <div className="stat">
                  <span className="label">Highest</span>
                  <span className="value">{courseAverage.highest_score?.toFixed(2)}</span>
                </div>
                <div className="stat">
                  <span className="label">Lowest</span>
                  <span className="value">{courseAverage.lowest_score?.toFixed(2)}</span>
                </div>
              </div>
            </div>
          )}
        </div>
      )}

      {/* Grade Distribution Tab */}
      {activeTab === 'distribution' && (
        <div className="tab-content">
          <div className="input-group">
            <input
              type="number"
              placeholder="Course ID"
              value={courseId}
              onChange={(e) => setCourseId(e.target.value)}
            />
            <button onClick={loadGradeDistribution} disabled={loading}>
              {loading ? 'Loading...' : 'Get Distribution'}
            </button>
          </div>

          {gradeDistribution && (
            <GradeDistributionChart
              distribution={gradeDistribution.distribution}
              total={gradeDistribution.total_students}
            />
          )}
        </div>
      )}

      {/* Student Stats Tab */}
      {activeTab === 'student-stats' && (
        <div className="tab-content">
          <div className="input-group">
            <input
              type="number"
              placeholder="Student ID"
              value={studentId}
              onChange={(e) => setStudentId(e.target.value)}
            />
            <button onClick={loadStudentStats} disabled={loading}>
              {loading ? 'Loading...' : 'Get Stats'}
            </button>
          </div>

          {studentStats && (
            <div className="student-stats-card">
              <h3>Student Grade Statistics</h3>
              <div className="stats-grid">
                <div className="stat">
                  <span className="label">Total Grades</span>
                  <span className="value">{studentStats.grade_count}</span>
                </div>
                <div className="stat">
                  <span className="label">Average</span>
                  <span className="value">{studentStats.average}</span>
                </div>
                <div className="stat">
                  <span className="label">Highest</span>
                  <span className="value">{studentStats.highest_grade?.toFixed(2)}</span>
                </div>
                <div className="stat">
                  <span className="label">Lowest</span>
                  <span className="value">{studentStats.lowest_grade?.toFixed(2)}</span>
                </div>
              </div>

              <div className="grade-distribution">
                <h4>Grade Breakdown</h4>
                <div className="grades">
                  <div className="grade-item">
                    <span>A:</span> {studentStats.a_count}
                  </div>
                  <div className="grade-item">
                    <span>B:</span> {studentStats.b_count}
                  </div>
                  <div className="grade-item">
                    <span>C:</span> {studentStats.c_count}
                  </div>
                  <div className="grade-item">
                    <span>D:</span> {studentStats.d_count}
                  </div>
                  <div className="grade-item">
                    <span>F:</span> {studentStats.f_count}
                  </div>
                </div>
              </div>
            </div>
          )}
        </div>
      )}
    </div>
  )
}
