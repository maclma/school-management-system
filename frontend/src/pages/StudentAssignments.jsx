import { useState, useEffect } from 'react'
import api from '../api'
import ui from '../ui'

export default function StudentAssignments() {
  const [assignments, setAssignments] = useState([])
  const [submissions, setSubmissions] = useState([])
  const [courses, setCourses] = useState([])
  const [loading, setLoading] = useState(true)
  const [selectedCourse, setSelectedCourse] = useState('')
  const [submittingAssignment, setSubmittingAssignment] = useState(null)

  useEffect(() => {
    loadData()
  }, [])

  const loadData = async () => {
    try {
      setLoading(true)
      const [coursesRes, submissionsRes] = await Promise.all([
        api.getCourses(),
        api.getSubmissionsByStudent()
      ])

      setCourses(coursesRes || [])
      setSubmissions(submissionsRes.submissions || [])

      // Get assignments for all enrolled courses
      const enrolledCourseIds = [...new Set(submissionsRes.submissions?.map(s => s.assignment.course_id) || [])]
      if (enrolledCourseIds.length > 0) {
        const assignmentPromises = enrolledCourseIds.map(courseId => api.getAssignmentsByCourse(courseId))
        const assignmentResults = await Promise.all(assignmentPromises)
        const allAssignments = assignmentResults.flatMap(result => result.assignments || [])
        setAssignments(allAssignments)
      }
    } catch (error) {
      ui.toast.error('Failed to load assignments')
    } finally {
      setLoading(false)
    }
  }

  const handleSubmitAssignment = async (assignmentId, fileUrl) => {
    try {
      await api.submitAssignment({
        assignment_id: assignmentId,
        file_url: fileUrl
      })
      ui.toast.success('Assignment submitted successfully')
      setSubmittingAssignment(null)
      loadData()
    } catch (error) {
      ui.toast.error(error.message || 'Failed to submit assignment')
    }
  }

  const getSubmissionForAssignment = (assignmentId) => {
    return submissions.find(s => s.assignment_id === assignmentId)
  }

  const getStatusBadge = (assignment, submission) => {
    if (!submission) return <span className="badge danger">Not Submitted</span>
    if (submission.status === 'graded') return <span className="badge success">Graded</span>
    if (submission.status === 'submitted') return <span className="badge info">Submitted</span>
    return <span className="badge warning">Pending</span>
  }

  const isOverdue = (dueDate) => {
    return new Date(dueDate) < new Date()
  }

  const filteredAssignments = selectedCourse
    ? assignments.filter(a => a.course_id === parseInt(selectedCourse))
    : assignments

  if (loading) {
    return <div className="loading">Loading assignments...</div>
  }

  return (
    <div className="student-assignments-page">
      <div className="page-header">
        <h1>My Assignments</h1>
        <div className="filters">
          <select
            value={selectedCourse}
            onChange={(e) => setSelectedCourse(e.target.value)}
          >
            <option value="">All Courses</option>
            {courses.map(course => (
              <option key={course.id} value={course.id}>
                {course.name} ({course.course_code})
              </option>
            ))}
          </select>
        </div>
      </div>

      {filteredAssignments.length === 0 ? (
        <div className="empty-state">
          <p>No assignments found.</p>
          {selectedCourse && (
            <p>Try selecting a different course or check back later for new assignments.</p>
          )}
        </div>
      ) : (
        <div className="assignments-list">
          {filteredAssignments.map(assignment => {
            const submission = getSubmissionForAssignment(assignment.id)
            const overdue = isOverdue(assignment.due_date)

            return (
              <div key={assignment.id} className={`assignment-item ${overdue && !submission ? 'overdue' : ''}`}>
                <div className="assignment-header">
                  <h3>{assignment.title}</h3>
                  <div className="assignment-meta">
                    <span className="course-name">{assignment.course.name}</span>
                    {getStatusBadge(assignment, submission)}
                  </div>
                </div>

                <div className="assignment-details">
                  <p className="description">{assignment.description}</p>
                  <div className="assignment-info">
                    <span>Due: {new Date(assignment.due_date).toLocaleDateString()}</span>
                    <span>Max Score: {assignment.max_score}</span>
                  </div>
                </div>

                <div className="assignment-status">
                  {submission ? (
                    <div className="submission-info">
                      <p>
                        <strong>Submitted:</strong> {new Date(submission.submitted_at).toLocaleString()}
                      </p>
                      {submission.status === 'graded' && (
                        <p>
                          <strong>Score:</strong> {submission.score}/{assignment.max_score}
                        </p>
                      )}
                      {submission.feedback && (
                        <div className="feedback">
                          <strong>Feedback:</strong>
                          <p>{submission.feedback}</p>
                        </div>
                      )}
                    </div>
                  ) : overdue ? (
                    <div className="overdue-notice">
                      <span className="badge danger">Overdue</span>
                      <p>This assignment is past its due date.</p>
                    </div>
                  ) : (
                    <button
                      className="button primary"
                      onClick={() => setSubmittingAssignment(assignment)}
                    >
                      Submit Assignment
                    </button>
                  )}
                </div>
              </div>
            )
          })}
        </div>
      )}

      {submittingAssignment && (
        <SubmitAssignmentModal
          assignment={submittingAssignment}
          onSubmit={handleSubmitAssignment}
          onClose={() => setSubmittingAssignment(null)}
        />
      )}
    </div>
  )
}

function SubmitAssignmentModal({ assignment, onSubmit, onClose }) {
  const [fileUrl, setFileUrl] = useState('')

  const handleSubmit = (e) => {
    e.preventDefault()
    if (!fileUrl.trim()) {
      ui.toast.error('Please provide a file URL or link to your submission')
      return
    }

    onSubmit(assignment.id, fileUrl.trim())
  }

  return (
    <div className="modal-overlay" onClick={onClose}>
      <div className="modal-content" onClick={e => e.stopPropagation()}>
        <div className="modal-header">
          <h2>Submit Assignment</h2>
          <button className="close-button" onClick={onClose}>×</button>
        </div>

        <div className="assignment-info">
          <h3>{assignment.title}</h3>
          <p><strong>Course:</strong> {assignment.course.name}</p>
          <p><strong>Due:</strong> {new Date(assignment.due_date).toLocaleDateString()}</p>
          <p><strong>Max Score:</strong> {assignment.max_score}</p>
        </div>

        <form onSubmit={handleSubmit} className="submit-form">
          <div className="form-group">
            <label>Submission URL *</label>
            <input
              type="url"
              value={fileUrl}
              onChange={(e) => setFileUrl(e.target.value)}
              placeholder="https://drive.google.com/... or https://github.com/..."
              required
            />
            <small className="help-text">
              Provide a link to your submitted work (Google Drive, GitHub, etc.)
            </small>
          </div>

          <div className="form-actions">
            <button type="button" className="button secondary" onClick={onClose}>
              Cancel
            </button>
            <button type="submit" className="button primary">
              Submit Assignment
            </button>
          </div>
        </form>
      </div>
    </div>
  )
}