import { useState } from 'react'
import api from '../../api'
import './RubricsManager.css'

export default function RubricsManager() {
  const [activeTab, setActiveTab] = useState('create')
  const [assignmentId, setAssignmentId] = useState(1)
  const [rubricId, setRubricId] = useState(1)
  const [submissionId, setSubmissionId] = useState(1)
  
  const [rubrics, setRubrics] = useState([])
  const [selectedRubric, setSelectedRubric] = useState(null)
  const [submissionScore, setSubmissionScore] = useState(null)
  
  const [loading, setLoading] = useState(false)
  const [error, setError] = useState(null)
  const [success, setSuccess] = useState(false)

  const [formData, setFormData] = useState({
    name: '',
    assignment_id: '',
    criteria: [{ name: '', max_points: 0, description: '' }],
    is_active: true
  })

  const [scoreData, setScoreData] = useState({
    submission_id: '',
    rubric_id: '',
    criterion_scores: [],
    feedback: '',
    graded_by: ''
  })

  const handleFormChange = (e) => {
    const { name, value, checked, type } = e.target
    setFormData(prev => ({
      ...prev,
      [name]: type === 'checkbox' ? checked : value
    }))
  }

  const handleCriterionChange = (index, field, value) => {
    const newCriteria = [...formData.criteria]
    newCriteria[index][field] = field === 'max_points' ? parseFloat(value) : value
    setFormData(prev => ({ ...prev, criteria: newCriteria }))
  }

  const addCriterion = () => {
    setFormData(prev => ({
      ...prev,
      criteria: [...prev.criteria, { name: '', max_points: 0, description: '' }]
    }))
  }

  const removeCriterion = (index) => {
    setFormData(prev => ({
      ...prev,
      criteria: prev.criteria.filter((_, i) => i !== index)
    }))
  }

  const handleCreateRubric = async (e) => {
    e.preventDefault()
    setLoading(true)
    setError(null)
    setSuccess(false)

    try {
      if (!formData.name || !formData.assignment_id || formData.criteria.length === 0) {
        throw new Error('Name, Assignment ID, and at least one criterion are required')
      }

      await api.createRubric({
        name: formData.name,
        assignment_id: parseInt(formData.assignment_id),
        criteria: formData.criteria,
        is_active: formData.is_active
      })

      setSuccess(true)
      setFormData({ 
        name: '', 
        assignment_id: '', 
        criteria: [{ name: '', max_points: 0, description: '' }],
        is_active: true 
      })
    } catch (err) {
      setError(err.message)
    } finally {
      setLoading(false)
    }
  }

  const loadRubricsByAssignment = async () => {
    setLoading(true)
    setError(null)
    try {
      const data = await api.getRubricsByAssignment(assignmentId)
      setRubrics(data.data?.rubrics || [])
    } catch (err) {
      setError(err.message)
    } finally {
      setLoading(false)
    }
  }

  const loadRubric = async () => {
    setLoading(true)
    setError(null)
    try {
      const data = await api.getRubric(rubricId)
      setSelectedRubric(data.data)
    } catch (err) {
      setError(err.message)
    } finally {
      setLoading(false)
    }
  }

  const handleDeleteRubric = async (id) => {
    if (window.confirm('Are you sure you want to delete this rubric?')) {
      setLoading(true)
      try {
        await api.deleteRubric(id)
        setSuccess(true)
        setSelectedRubric(null)
        loadRubricsByAssignment()
      } catch (err) {
        setError(err.message)
      } finally {
        setLoading(false)
      }
    }
  }

  const handleScoreSubmission = async (e) => {
    e.preventDefault()
    setLoading(true)
    setError(null)
    setSuccess(false)

    try {
      if (!scoreData.submission_id || !scoreData.rubric_id || scoreData.criterion_scores.length === 0) {
        throw new Error('Submission ID, Rubric ID, and criterion scores are required')
      }

      await api.scoreSubmission({
        submission_id: parseInt(scoreData.submission_id),
        rubric_id: parseInt(scoreData.rubric_id),
        criterion_scores: scoreData.criterion_scores.map(cs => ({
          criterion_name: cs.criterion_name,
          points_earned: parseFloat(cs.points_earned)
        })),
        feedback: scoreData.feedback,
        graded_by: parseInt(scoreData.graded_by)
      })

      setSuccess(true)
      setScoreData({
        submission_id: '',
        rubric_id: '',
        criterion_scores: [],
        feedback: '',
        graded_by: ''
      })
    } catch (err) {
      setError(err.message)
    } finally {
      setLoading(false)
    }
  }

  const loadSubmissionScore = async () => {
    setLoading(true)
    setError(null)
    try {
      const data = await api.getSubmissionScore(submissionId)
      setSubmissionScore(data.data)
    } catch (err) {
      setError(err.message)
    } finally {
      setLoading(false)
    }
  }

  return (
    <div className="rubrics-manager-container">
      <h2>Rubrics Management</h2>

      <div className="tabs">
        <button 
          className={`tab ${activeTab === 'create' ? 'active' : ''}`}
          onClick={() => setActiveTab('create')}
        >
          Create Rubric
        </button>
        <button 
          className={`tab ${activeTab === 'list' ? 'active' : ''}`}
          onClick={() => setActiveTab('list')}
        >
          View Rubrics
        </button>
        <button 
          className={`tab ${activeTab === 'score' ? 'active' : ''}`}
          onClick={() => setActiveTab('score')}
        >
          Score Submission
        </button>
        <button 
          className={`tab ${activeTab === 'view-score' ? 'active' : ''}`}
          onClick={() => setActiveTab('view-score')}
        >
          View Score
        </button>
      </div>

      {error && <div className="error-message">{error}</div>}
      {success && <div className="success-message">✓ Operation completed successfully!</div>}

      {/* Create Rubric Tab */}
      {activeTab === 'create' && (
        <div className="tab-content">
          <h3>Create New Rubric</h3>
          <form onSubmit={handleCreateRubric} className="rubric-form">
            <div className="form-group">
              <label>Rubric Name *</label>
              <input
                type="text"
                name="name"
                required
                value={formData.name}
                onChange={handleFormChange}
                placeholder="e.g., Essay Grading Rubric"
              />
            </div>

            <div className="form-group">
              <label>Assignment ID *</label>
              <input
                type="number"
                name="assignment_id"
                required
                value={formData.assignment_id}
                onChange={handleFormChange}
                placeholder="Assignment ID"
              />
            </div>

            <div className="criteria-section">
              <h4>Grading Criteria</h4>
              {formData.criteria.map((criterion, index) => (
                <div key={index} className="criterion-form">
                  <input
                    type="text"
                    placeholder="Criterion name (e.g., Organization)"
                    value={criterion.name}
                    onChange={(e) => handleCriterionChange(index, 'name', e.target.value)}
                  />
                  <input
                    type="number"
                    placeholder="Max points"
                    step="0.5"
                    value={criterion.max_points}
                    onChange={(e) => handleCriterionChange(index, 'max_points', e.target.value)}
                  />
                  <textarea
                    placeholder="Description (optional)"
                    value={criterion.description}
                    onChange={(e) => handleCriterionChange(index, 'description', e.target.value)}
                  />
                  {formData.criteria.length > 1 && (
                    <button
                      type="button"
                      className="remove-btn"
                      onClick={() => removeCriterion(index)}
                    >
                      Remove
                    </button>
                  )}
                </div>
              ))}
              <button type="button" className="add-btn" onClick={addCriterion}>
                + Add Criterion
              </button>
            </div>

            <div className="form-group">
              <label>
                <input
                  type="checkbox"
                  name="is_active"
                  checked={formData.is_active}
                  onChange={handleFormChange}
                />
                Active
              </label>
            </div>

            <button type="submit" disabled={loading}>
              {loading ? 'Creating...' : 'Create Rubric'}
            </button>
          </form>
        </div>
      )}

      {/* List Rubrics Tab */}
      {activeTab === 'list' && (
        <div className="tab-content">
          <div className="input-group">
            <input
              type="number"
              placeholder="Assignment ID"
              value={assignmentId}
              onChange={(e) => setAssignmentId(e.target.value)}
            />
            <button onClick={loadRubricsByAssignment} disabled={loading}>
              {loading ? 'Loading...' : 'Load Rubrics'}
            </button>
          </div>

          {rubrics.length > 0 && (
            <div className="rubrics-list">
              <h3>Rubrics for Assignment {assignmentId}</h3>
              {rubrics.map((rubric) => (
                <div key={rubric.id} className="rubric-card">
                  <h4>{rubric.name}</h4>
                  <p className={`status ${rubric.is_active ? 'active' : 'inactive'}`}>
                    {rubric.is_active ? 'Active' : 'Inactive'}
                  </p>
                  <div className="criteria">
                    <h5>Criteria:</h5>
                    <ul>
                      {rubric.criteria?.map((c, idx) => (
                        <li key={idx}>
                          <strong>{c.name}</strong> - {c.max_points} points
                          {c.description && <p>{c.description}</p>}
                        </li>
                      ))}
                    </ul>
                  </div>
                  <div className="actions">
                    <button 
                      onClick={() => { setRubricId(rubric.id); loadRubric(); }}
                      className="view-btn"
                    >
                      View Details
                    </button>
                    <button 
                      onClick={() => handleDeleteRubric(rubric.id)}
                      className="delete-btn"
                    >
                      Delete
                    </button>
                  </div>
                </div>
              ))}
            </div>
          )}

          {selectedRubric && (
            <div className="rubric-details">
              <h3>Rubric Details - {selectedRubric.name}</h3>
              <p>ID: {selectedRubric.id}</p>
              <p>Assignment ID: {selectedRubric.assignment_id}</p>
              <div className="criteria">
                <h4>Criteria:</h4>
                <ul>
                  {selectedRubric.criteria?.map((c, idx) => (
                    <li key={idx}>
                      <strong>{c.name}</strong> - {c.max_points} points
                      {c.description && <p className="desc">{c.description}</p>}
                    </li>
                  ))}
                </ul>
              </div>
            </div>
          )}
        </div>
      )}

      {/* Score Submission Tab */}
      {activeTab === 'score' && (
        <div className="tab-content">
          <h3>Score Submission Using Rubric</h3>
          <form onSubmit={handleScoreSubmission} className="score-form">
            <div className="form-group">
              <label>Submission ID *</label>
              <input
                type="number"
                value={scoreData.submission_id}
                onChange={(e) => setScoreData(prev => ({ ...prev, submission_id: e.target.value }))}
                placeholder="Submission ID"
                required
              />
            </div>

            <div className="form-group">
              <label>Rubric ID *</label>
              <input
                type="number"
                value={scoreData.rubric_id}
                onChange={(e) => setScoreData(prev => ({ ...prev, rubric_id: e.target.value }))}
                placeholder="Rubric ID"
                required
              />
            </div>

            <div className="form-group">
              <label>Criterion Scores *</label>
              <textarea
                placeholder='JSON format: [{"criterion_name": "Organization", "points_earned": 23}]'
                value={JSON.stringify(scoreData.criterion_scores)}
                onChange={(e) => {
                  try {
                    setScoreData(prev => ({ ...prev, criterion_scores: JSON.parse(e.target.value) }))
                  } catch (err) {
                    // Invalid JSON, ignore
                  }
                }}
                rows="4"
              />
            </div>

            <div className="form-group">
              <label>Feedback</label>
              <textarea
                value={scoreData.feedback}
                onChange={(e) => setScoreData(prev => ({ ...prev, feedback: e.target.value }))}
                placeholder="Grading feedback"
                rows="4"
              />
            </div>

            <div className="form-group">
              <label>Graded By (User ID)</label>
              <input
                type="number"
                value={scoreData.graded_by}
                onChange={(e) => setScoreData(prev => ({ ...prev, graded_by: e.target.value }))}
                placeholder="Teacher/Admin ID"
              />
            </div>

            <button type="submit" disabled={loading}>
              {loading ? 'Scoring...' : 'Score Submission'}
            </button>
          </form>

          <div className="help-text">
            <h4>Format Example:</h4>
            <pre>{`[
  {"criterion_name": "Organization", "points_earned": 23},
  {"criterion_name": "Content", "points_earned": 48},
  {"criterion_name": "Grammar", "points_earned": 24}
]`}</pre>
          </div>
        </div>
      )}

      {/* View Score Tab */}
      {activeTab === 'view-score' && (
        <div className="tab-content">
          <div className="input-group">
            <input
              type="number"
              placeholder="Submission ID"
              value={submissionId}
              onChange={(e) => setSubmissionId(e.target.value)}
            />
            <button onClick={loadSubmissionScore} disabled={loading}>
              {loading ? 'Loading...' : 'Get Score'}
            </button>
          </div>

          {submissionScore && (
            <div className="score-details">
              <h3>Submission Score Details</h3>
              
              <div className="score-summary">
                <div className="score-stat">
                  <span className="label">Total Score</span>
                  <span className="value">{submissionScore.total_points_earned} / {submissionScore.total_points_available}</span>
                </div>
                <div className="score-stat">
                  <span className="label">Percentage</span>
                  <span className="value">{submissionScore.percentage?.toFixed(1)}%</span>
                </div>
                <div className="score-stat">
                  <span className="label">Letter Grade</span>
                  <span className="value grade">{submissionScore.letter_grade}</span>
                </div>
              </div>

              {submissionScore.criterion_scores?.length > 0 && (
                <div className="criterion-breakdown">
                  <h4>Criterion Breakdown</h4>
                  <table>
                    <thead>
                      <tr>
                        <th>Criterion</th>
                        <th>Points Earned</th>
                        <th>Max Points</th>
                        <th>%</th>
                      </tr>
                    </thead>
                    <tbody>
                      {submissionScore.criterion_scores.map((cs, idx) => (
                        <tr key={idx}>
                          <td>{cs.criterion_name}</td>
                          <td>{cs.points_earned}</td>
                          <td>{cs.max_points}</td>
                          <td>{((cs.points_earned / cs.max_points) * 100).toFixed(1)}%</td>
                        </tr>
                      ))}
                    </tbody>
                  </table>
                </div>
              )}

              {submissionScore.feedback && (
                <div className="feedback">
                  <h4>Feedback</h4>
                  <p>{submissionScore.feedback}</p>
                </div>
              )}

              <p className="graded-at">
                Graded: {new Date(submissionScore.graded_at).toLocaleString()}
              </p>
            </div>
          )}
        </div>
      )}
    </div>
  )
}
