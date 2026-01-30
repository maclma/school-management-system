/**
 * Toast Notification Component
 * Displays notifications with auto-dismiss
 */

import React, { useEffect, useState } from 'react';
import { Toast } from '../utils/toast';

const toastStyles = {
  success: {
    bg: '#10b981',
    icon: '✓',
    borderColor: '#059669',
  },
  error: {
    bg: '#ef4444',
    icon: '✕',
    borderColor: '#dc2626',
  },
  warning: {
    bg: '#f59e0b',
    icon: '⚠',
    borderColor: '#d97706',
  },
  info: {
    bg: '#3b82f6',
    icon: 'ℹ',
    borderColor: '#2563eb',
  },
  loading: {
    bg: '#6366f1',
    icon: '⟳',
    borderColor: '#4f46e5',
  },
};

const ToastItem = ({ notification, onDismiss }) => {
  const style = toastStyles[notification.type] || toastStyles.info;

  return (
    <div
      style={{
        display: 'flex',
        alignItems: 'center',
        backgroundColor: style.bg,
        color: 'white',
        padding: '12px 16px',
        borderRadius: '4px',
        marginBottom: '8px',
        boxShadow: '0 4px 6px rgba(0, 0, 0, 0.1)',
        animation: 'slideIn 0.3s ease-in-out',
        borderLeft: `4px solid ${style.borderColor}`,
      }}
    >
      <span
        style={{
          fontSize: '18px',
          marginRight: '12px',
          display: 'flex',
          alignItems: 'center',
        }}
      >
        {notification.type === 'loading' ? (
          <span style={{ animation: 'spin 1s linear infinite' }}>
            {style.icon}
          </span>
        ) : (
          style.icon
        )}
      </span>

      <span style={{ flex: 1, fontSize: '14px' }}>
        {notification.message}
      </span>

      {notification.duration > 0 && (
        <button
          onClick={() => onDismiss(notification.id)}
          style={{
            background: 'none',
            border: 'none',
            color: 'white',
            cursor: 'pointer',
            fontSize: '18px',
            padding: '0 0 0 12px',
          }}
        >
          ✕
        </button>
      )}
    </div>
  );
};

export function ToastContainer() {
  const [notifications, setNotifications] = useState([]);

  useEffect(() => {
    const unsubscribe = Toast.subscribe(setNotifications);
    return unsubscribe;
  }, []);

  const handleDismiss = (id) => {
    Toast.getInstance().remove(id);
  };

  return (
    <div
      style={{
        position: 'fixed',
        top: '20px',
        right: '20px',
        zIndex: 9999,
        maxWidth: '400px',
        maxHeight: '90vh',
        overflowY: 'auto',
      }}
    >
      <style>{`
        @keyframes slideIn {
          from {
            transform: translateX(400px);
            opacity: 0;
          }
          to {
            transform: translateX(0);
            opacity: 1;
          }
        }

        @keyframes spin {
          from {
            transform: rotate(0deg);
          }
          to {
            transform: rotate(360deg);
          }
        }
      `}</style>

      {notifications.map(notification => (
        <ToastItem
          key={notification.id}
          notification={notification}
          onDismiss={handleDismiss}
        />
      ))}
    </div>
  );
}

export default ToastContainer;
