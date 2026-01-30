import React, { useState, useEffect } from 'react'
import api from '../api'

export default function Notifications() {
  const [notifications, setNotifications] = useState([])
  const [unreadCount, setUnreadCount] = useState(0)
  const [loading, setLoading] = useState(false)

  useEffect(() => {
    loadNotifications()
  }, [])

  const loadNotifications = async () => {
    setLoading(true)
    try {
      const res = await api.getMyNotifications()
      setNotifications(res.data || [])
      
      const unread = await api.getUnreadNotifications()
      setUnreadCount((unread.data || []).length)
    } catch (err) {
      console.error('Failed to load notifications:', err)
    }
    setLoading(false)
  }

  const handleMarkAsRead = async (id) => {
    try {
      await api.markNotificationAsRead(id)
      loadNotifications()
    } catch (err) {
      console.error('Failed to mark notification as read:', err)
    }
  }

  const handleMarkAllAsRead = async () => {
    try {
      await api.markAllNotificationsAsRead()
      loadNotifications()
    } catch (err) {
      console.error('Failed to mark all as read:', err)
    }
  }

  const handleDelete = async (id) => {
    if (confirm('Delete this notification?')) {
      try {
        await api.deleteNotification(id)
        loadNotifications()
      } catch (err) {
        console.error('Failed to delete notification:', err)
      }
    }
  }

  return (
    <div className="notifications-container">
      <div className="notifications-header">
        <h2>Notifications</h2>
        {unreadCount > 0 && (
          <span className="unread-badge">{unreadCount} new</span>
        )}
      </div>

      {unreadCount > 0 && (
        <button 
          className="btn btn-sm btn-primary"
          onClick={handleMarkAllAsRead}
        >
          Mark All as Read
        </button>
      )}

      {loading ? (
        <p>Loading notifications...</p>
      ) : notifications.length === 0 ? (
        <p className="no-data">No notifications yet</p>
      ) : (
        <div className="notifications-list">
          {notifications.map((notif) => (
            <div 
              key={notif.id} 
              className={`notification-item ${notif.is_read ? 'read' : 'unread'}`}
            >
              <div className="notification-content">
                <h4>{notif.title}</h4>
                <p>{notif.message}</p>
                <small className="notification-time">
                  {new Date(notif.created_at).toLocaleDateString()}
                </small>
              </div>
              <div className="notification-actions">
                {!notif.is_read && (
                  <button
                    className="btn-icon"
                    onClick={() => handleMarkAsRead(notif.id)}
                    title="Mark as read"
                  >
                    ✓
                  </button>
                )}
                <button
                  className="btn-icon btn-danger"
                  onClick={() => handleDelete(notif.id)}
                  title="Delete"
                >
                  ✕
                </button>
              </div>
            </div>
          ))}
        </div>
      )}
    </div>
  )
}
