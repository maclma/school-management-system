import React, { useState, useEffect } from 'react'
import api from '../api'

export default function Announcements() {
  const [announcements, setAnnouncements] = useState([])
  const [loading, setLoading] = useState(false)
  const [page, setPage] = useState(1)
  const [showForm, setShowForm] = useState(false)
  const [formData, setFormData] = useState({
    title: '',
    content: '',
    audience: 'all',
    priority: 'medium'
  })

  useEffect(() => {
    loadAnnouncements()
  }, [page])

  const loadAnnouncements = async () => {
    setLoading(true)
    try {
      const res = await api.getActiveAnnouncements(page)
      setAnnouncements(res.data || [])
    } catch (err) {
      console.error('Failed to load announcements:', err)
    }
    setLoading(false)
  }

  const handleSubmit = async (e) => {
    e.preventDefault()
    if (!formData.title || !formData.content) {
      alert('Title and content are required')
      return
    }

    try {
      await api.createAnnouncement(formData)
      setFormData({ title: '', content: '', audience: 'all', priority: 'medium' })
      setShowForm(false)
      loadAnnouncements()
      alert('Announcement created successfully')
    } catch (err) {
      alert('Failed to create announcement: ' + err.message)
    }
  }

  return (
    <div className="announcements-container">
      <h2>Announcements</h2>
      
      <button 
        className="btn btn-primary"
        onClick={() => setShowForm(!showForm)}
      >
        {showForm ? 'Cancel' : 'New Announcement'}
      </button>

      {showForm && (
        <form onSubmit={handleSubmit} className="announcement-form">
          <div className="form-group">
            <label>Title *</label>
            <input
              type="text"
              value={formData.title}
              onChange={(e) => setFormData({...formData, title: e.target.value})}
              placeholder="Announcement title"
              required
            />
          </div>

          <div className="form-group">
            <label>Content *</label>
            <textarea
              value={formData.content}
              onChange={(e) => setFormData({...formData, content: e.target.value})}
              placeholder="Announcement content"
              rows="5"
              required
            ></textarea>
          </div>

          <div className="form-row">
            <div className="form-group">
              <label>Audience</label>
              <select
                value={formData.audience}
                onChange={(e) => setFormData({...formData, audience: e.target.value})}
              >
                <option value="all">Everyone</option>
                <option value="students">Students Only</option>
                <option value="teachers">Teachers Only</option>
              </select>
            </div>

            <div className="form-group">
              <label>Priority</label>
              <select
                value={formData.priority}
                onChange={(e) => setFormData({...formData, priority: e.target.value})}
              >
                <option value="low">Low</option>
                <option value="medium">Medium</option>
                <option value="high">High</option>
              </select>
            </div>
          </div>

          <button type="submit" className="btn btn-success">Create Announcement</button>
        </form>
      )}

      {loading ? (
        <p>Loading announcements...</p>
      ) : announcements.length === 0 ? (
        <p className="no-data">No active announcements</p>
      ) : (
        <div className="announcements-list">
          {announcements.map((ann) => (
            <div key={ann.id} className="announcement-card">
              <div className="announcement-header">
                <h3>{ann.title}</h3>
                <span className={`priority-badge priority-${ann.priority}`}>
                  {ann.priority.toUpperCase()}
                </span>
              </div>
              <p className="announcement-content">{ann.content}</p>
              <div className="announcement-meta">
                <small>Audience: {ann.audience}</small>
                <small>Posted: {new Date(ann.created_at).toLocaleDateString()}</small>
              </div>
            </div>
          ))}
        </div>
      )}
    </div>
  )
}
