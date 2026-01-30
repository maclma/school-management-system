import { useState, useEffect } from 'react'
import api from '../api'
import ui from '../ui'

export default function TeacherAssignments() {
  const [assignments, setAssignments] = useState([])
  const [courses, setCourses] = useState([])
  const [showCreateForm, setShowCreateForm] = useState(false)
  const [selectedCourse, setSelectedCourse] = useState('')
  const [loading, setLoading] = useState(true)

  useEffect(() => {
    loadData()
  }, [])

  const loadData = async () => {
    try {
      setLoading(true)
      const [assignmentsRes, coursesRes] = await Promise.all([
        api.getAssignmentsByTeacher(),
        api.getCourses()
      ])
      setAssignments(assignmentsRes.assignments || [])
      setCourses(coursesRes || [])
    } catch (error) {
      ui.toast.error('Failed to load assignments')
    } finally {
      setLoading(false)
    }
  }

  const handleCreateAssignment = async (formData) => {
    try {
      await api.createAssignment(formData)
      ui.toast.success('Assignment created successfully')
      setShowCreateForm(false)
      loadData()
    } catch (error) {
      ui.toast.error(error.message || 'Failed to create assignment')
    }
  }

  const handleDeleteAssignment = async (assignmentId) => {
    if (!confirm('Are you sure you want to delete this assignment?')) return

    try {
      await api.deleteAssignment(assignmentId)
      ui.toast.success('Assignment deleted successfully')
      loadData()
    } catch (error) {
      ui.toast.error('Failed to delete assignment')
    }
  }

  const viewSubmissions = (assignmentId) => {
    window.navigate(`/teacher/assignments/${assignmentId}/submissions`)
  }

  if (loading) {
    return <div className="loading">Loading assignments...</div>
  }

  return (
    <div className="assignments-page">
      <div className="page-header">
        <h1>My Assignments</h1>
        <button
          className="button primary"
          onClick={() => setShowCreateForm(true)}
        >
          + Create Assignment
        </button>
      </div>

      {assignments.length === 0 ? (
        <div className="empty-state">
          <p>You haven't created any assignments yet.</p>
          <button
            className="button primary"
            onClick={() => setShowCreateForm(true)}
          >
            Create Your First Assignment
          </button>
        </div>
      ) : (
        <div className="assignments-grid">
          {assignments.map(assignment => (
            <div key={assignment.id} className="assignment-card">
              <div className="assignment-header">
                <h3>{assignment.title}</h3>
                <span className="course-name">{assignment.course.name}</span>
              </div>
              <div className="assignment-details">
                <p className="description">{assignment.description}</p>
                <div className="assignment-meta">
                  <span>Due: {new Date(assignment.due_date).toLocaleDateString()}</span>
                  <span>Max Score: {assignment.max_score}</span>
                </div>
              </div>
              <div className="assignment-actions">
                <button
                  className="button secondary small"
                  onClick={() => viewSubmissions(assignment.id)}
                >
                  View Submissions ({assignment.submissions?.length || 0})
                </button>
                <button
                  className="button danger small"
                  onClick={() => handleDeleteAssignment(assignment.id)}
                >
                  Delete
                </button>
              </div>
            </div>
          ))}
        </div>
      )}

      {showCreateForm && (
        <CreateAssignmentModal
          courses={courses}
          onSubmit={handleCreateAssignment}
          onClose={() => setShowCreateForm(false)}
        />
      )}
    </div>
  )
}

function CreateAssignmentModal({ courses, onSubmit, onClose }) {
  const [formData, setFormData] = useState({
    course_id: '',
    title: '',
    description: '',
    due_date: '',
    max_score: 100
  })

  const handleSubmit = (e) => {
    e.preventDefault()
    if (!formData.course_id || !formData.title || !formData.due_date) {
      ui.toast.error('Please fill in all required fields')
      return
    }

    onSubmit({
      ...formData,
      course_id: parseInt(formData.course_id),
      max_score: parseFloat(formData.max_score),
      due_date: new Date(formData.due_date).toISOString()
    })
  }

  const handleChange = (e) => {
    setFormData({
      ...formData,
      [e.target.name]: e.target.value
    })
  }

  return (
    <div className="modal-overlay" onClick={onClose}>
      <div className="modal-content" onClick={e => e.stopPropagation()}>
        <div className="modal-header">
          <h2>Create New Assignment</h2>
          <button className="close-button" onClick={onClose}>×</button>
        </div>

        <form onSubmit={handleSubmit} className="assignment-form">
          <div className="form-group">
            <label>Course *</label>
            <select
              name="course_id"
              value={formData.course_id}
              onChange={handleChange}
              required
            >
              <option value="">Select a course</option>
              {courses.map(course => (
                <option key={course.id} value={course.id}>
                  {course.name} ({course.course_code})
                </option>
              ))}
            </select>
          </div>

          <div className="form-group">
            <label>Title *</label>
            <input
              type="text"
              name="title"
              value={formData.title}
              onChange={handleChange}
              placeholder="Assignment title"
              required
            />
          </div>

          <div className="form-group">
            <label>Description</label>
            <textarea
              name="description"
              value={formData.description}
              onChange={handleChange}
              placeholder="Assignment description"
              rows="4"
            />
          </div>

          <div className="form-row">
            <div className="form-group">
              <label>Due Date *</label>
              <input
                type="datetime-local"
                name="due_date"
                value={formData.due_date}
                onChange={handleChange}
                required
              />
            </div>

            <div className="form-group">
              <label>Max Score</label>
              <input
                type="number"
                name="max_score"
                value={formData.max_score}
                onChange={handleChange}
                min="0"
                step="0.1"
              />
            </div>
          </div>

          <div className="form-actions">
            <button type="button" className="button secondary" onClick={onClose}>
              Cancel
            </button>
            <button type="submit" className="button primary">
              Create Assignment
            </button>
          </div>
        </form>
      </div>
    </div>
  )
}