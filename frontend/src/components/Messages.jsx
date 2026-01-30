import React, { useState, useEffect } from 'react'
import api from '../api'

export default function Messages() {
  const [messages, setMessages] = useState([])
  const [unreadCount, setUnreadCount] = useState(0)
  const [loading, setLoading] = useState(false)
  const [showCompose, setShowCompose] = useState(false)
  const [selectedConversation, setSelectedConversation] = useState(null)
  const [formData, setFormData] = useState({
    recipient_id: '',
    subject: '',
    content: ''
  })

  useEffect(() => {
    loadMessages()
  }, [])

  const loadMessages = async () => {
    setLoading(true)
    try {
      const res = await api.getInbox()
      setMessages(res.data || [])
      
      const unread = await api.countUnreadMessages()
      setUnreadCount(unread.data?.unread_count || 0)
    } catch (err) {
      console.error('Failed to load messages:', err)
    }
    setLoading(false)
  }

  const handleSendMessage = async (e) => {
    e.preventDefault()
    if (!formData.recipient_id || !formData.subject || !formData.content) {
      alert('All fields are required')
      return
    }

    try {
      await api.sendMessage({
        recipient_id: parseInt(formData.recipient_id),
        subject: formData.subject,
        content: formData.content
      })
      setFormData({ recipient_id: '', subject: '', content: '' })
      setShowCompose(false)
      loadMessages()
      alert('Message sent successfully')
    } catch (err) {
      alert('Failed to send message: ' + err.message)
    }
  }

  const handleMarkAsRead = async (id) => {
    try {
      await api.markMessageAsRead(id)
      loadMessages()
    } catch (err) {
      console.error('Failed to mark message as read:', err)
    }
  }

  return (
    <div className="messages-container">
      <div className="messages-header">
        <h2>Messages</h2>
        {unreadCount > 0 && (
          <span className="unread-badge">{unreadCount}</span>
        )}
      </div>

      <button 
        className="btn btn-primary"
        onClick={() => setShowCompose(!showCompose)}
      >
        {showCompose ? 'Cancel' : 'New Message'}
      </button>

      {showCompose && (
        <form onSubmit={handleSendMessage} className="message-form">
          <div className="form-group">
            <label>Recipient ID *</label>
            <input
              type="number"
              value={formData.recipient_id}
              onChange={(e) => setFormData({...formData, recipient_id: e.target.value})}
              placeholder="User ID"
              required
            />
          </div>

          <div className="form-group">
            <label>Subject *</label>
            <input
              type="text"
              value={formData.subject}
              onChange={(e) => setFormData({...formData, subject: e.target.value})}
              placeholder="Message subject"
              required
            />
          </div>

          <div className="form-group">
            <label>Message *</label>
            <textarea
              value={formData.content}
              onChange={(e) => setFormData({...formData, content: e.target.value})}
              placeholder="Your message"
              rows="5"
              required
            ></textarea>
          </div>

          <button type="submit" className="btn btn-success">Send Message</button>
        </form>
      )}

      {loading ? (
        <p>Loading messages...</p>
      ) : messages.length === 0 ? (
        <p className="no-data">No messages</p>
      ) : (
        <div className="messages-list">
          {messages.map((msg) => (
            <div 
              key={msg.id} 
              className={`message-item ${msg.is_read ? 'read' : 'unread'}`}
              onClick={() => setSelectedConversation(msg.id)}
            >
              <div className="message-header">
                <h4>{msg.subject}</h4>
                {!msg.is_read && (
                  <button
                    className="btn-sm btn-mark-read"
                    onClick={(e) => {
                      e.stopPropagation()
                      handleMarkAsRead(msg.id)
                    }}
                  >
                    Mark as Read
                  </button>
                )}
              </div>
              <p className="message-preview">{msg.content.substring(0, 100)}...</p>
              <small className="message-date">
                {new Date(msg.created_at).toLocaleDateString()}
              </small>
            </div>
          ))}
        </div>
      )}
    </div>
  )
}
