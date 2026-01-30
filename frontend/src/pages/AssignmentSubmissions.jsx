import { useState, useEffect } from 'react'
import api from '../api'
import ui from '../ui'

export default function AssignmentSubmissions() {
  const path = window.location.pathname
  const assignmentId = path.match(/^\/teacher\/assignments\/(\d+)\/submissions/)?.[1]
  const [assignment, setAssignment] = useState(null)
  const [submissions, setSubmissions] = useState([])
  const [loading, setLoading] = useState(true)
  const [gradingSubmission, setGradingSubmission] = useState(null)

  useEffect(() => {
    loadData()
  }, [assignmentId])

  const loadData = async () => {
    try {
      setLoading(true)
      const [assignmentRes, submissionsRes] = await Promise.all([
        api.getAssignment(assignmentId),
        api.getSubmissionsByAssignment(assignmentId)
      ])
      setAssignment(assignmentRes)
      setSubmissions(submissionsRes.submissions || [])
    } catch (error) {
      ui.toast.error('Failed to load assignment data')
    } finally {
      setLoading(false)
    }
  }

  const handleGradeSubmission = async (submissionId, score, feedback) => {
    try {
      await api.gradeSubmission(submissionId, { score: parseFloat(score), feedback })
      ui.toast.success('Submission graded successfully')
      setGradingSubmission(null)
      loadData()
    } catch (error) {
      ui.toast.error('Failed to grade submission')
    }
  }

  const getStatusBadge = (status) => {
    const statusClasses = {
      pending: 'badge warning',
      submitted: 'badge info',
      graded: 'badge success'
    }
    return statusClasses[status] || 'badge'
  }

  if (loading) {
    return <div className="loading">Loading submissions...</div>
  }

  if (!assignment) {
    return <div className="error">Assignment not found</div>
  }

  return (
    <div className="submissions-page">
      <div className="page-header">
        <div className="header-info">
          <h1>{assignment.title}</h1>
          <p className="course-info">{assignment.course.name} ({assignment.course.course_code})</p>
          <p className="assignment-details">
            Due: {new Date(assignment.due_date).toLocaleDateString()} •
            Max Score: {assignment.max_score}
          </p>
        </div>
        <button
          className="button secondary"
          onClick={() => window.navigate('/teacher/assignments')}
        >
          ← Back to Assignments
        </button>
      </div>

      <div className="submissions-summary">
        <div className="summary-item">
          <span className="label">Total Submissions:</span>
          <span className="value">{submissions.length}</span>
        </div>
        <div className="summary-item">
          <span className="label">Graded:</span>
          <span className="value">{submissions.filter(s => s.status === 'graded').length}</span>
        </div>
        <div className="summary-item">
          <span className="label">Pending:</span>
          <span className="value">{submissions.filter(s => s.status === 'pending').length}</span>
        </div>
      </div>

      {submissions.length === 0 ? (
        <div className="empty-state">
          <p>No submissions yet for this assignment.</p>
        </div>
      ) : (
        <div className="submissions-table">
          <table>
            <thead>
              <tr>
                <th>Student</th>
                <th>Submitted</th>
                <th>Status</th>
                <th>Score</th>
                <th>Actions</th>
              </tr>
            </thead>
            <tbody>
              {submissions.map(submission => (
                <tr key={submission.id}>
                  <td>
                    <div className="student-info">
                      <div className="student-name">
                        {submission.student.first_name} {submission.student.last_name}
                      </div>
                      <div className="student-email">{submission.student.email}</div>
                    </div>
                  </td>
                  <td>
                    {submission.submitted_at ?
                      new Date(submission.submitted_at).toLocaleString() :
                      'Not submitted'
                    }
                  </td>
                  <td>
                    <span className={getStatusBadge(submission.status)}>
                      {submission.status}
                    </span>
                  </td>
                  <td>
                    {submission.score !== null ? `${submission.score}/${assignment.max_score}` : '-'}
                  </td>
                  <td>
                    {submission.status === 'submitted' && (
                      <button
                        className="button primary small"
                        onClick={() => setGradingSubmission(submission)}
                      >
                        Grade
                      </button>
                    )}
                    {submission.status === 'graded' && (
                      <button
                        className="button secondary small"
                        onClick={() => setGradingSubmission(submission)}
                      >
                        Update Grade
                      </button>
                    )}
                  </td>
                </tr>
              ))}
            </tbody>
          </table>
        </div>
      )}

      {gradingSubmission && (
        <GradeSubmissionModal
          submission={gradingSubmission}
          assignment={assignment}
          onSubmit={handleGradeSubmission}
          onClose={() => setGradingSubmission(null)}
        />
      )}
    </div>
  )
}

function GradeSubmissionModal({ submission, assignment, onSubmit, onClose }) {
  const [score, setScore] = useState(submission.score || '')
  const [feedback, setFeedback] = useState(submission.feedback || '')

  const handleSubmit = (e) => {
    e.preventDefault()
    if (score === '' || score < 0 || score > assignment.max_score) {
      ui.toast.error(`Score must be between 0 and ${assignment.max_score}`)
      return
    }

    onSubmit(submission.id, score, feedback)
  }

  return (
    <div className="modal-overlay" onClick={onClose}>
      <div className="modal-content" onClick={e => e.stopPropagation()}>
        <div className="modal-header">
          <h2>Grade Submission</h2>
          <button className="close-button" onClick={onClose}>×</button>
        </div>

        <div className="submission-details">
          <h3>{assignment.title}</h3>
          <p><strong>Student:</strong> {submission.student.first_name} {submission.student.last_name}</p>
          <p><strong>Submitted:</strong> {new Date(submission.submitted_at).toLocaleString()}</p>
          {submission.file_url && (
            <p><strong>File:</strong> <a href={submission.file_url} target="_blank" rel="noopener noreferrer">View Submission</a></p>
          )}
        </div>

        <form onSubmit={handleSubmit} className="grading-form">
          <div className="form-group">
            <label>Score (0 - {assignment.max_score}) *</label>
            <input
              type="number"
              value={score}
              onChange={(e) => setScore(e.target.value)}
              min="0"
              max={assignment.max_score}
              step="0.1"
              required
            />
          </div>

          <div className="form-group">
            <label>Feedback</label>
            <textarea
              value={feedback}
              onChange={(e) => setFeedback(e.target.value)}
              placeholder="Provide feedback for the student..."
              rows="4"
            />
          </div>

          <div className="form-actions">
            <button type="button" className="button secondary" onClick={onClose}>
              Cancel
            </button>
            <button type="submit" className="button primary">
              {submission.status === 'graded' ? 'Update Grade' : 'Submit Grade'}
            </button>
          </div>
        </form>
      </div>
    </div>
  )
}