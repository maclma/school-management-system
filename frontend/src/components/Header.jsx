import React from 'react'

export default function Header() {
  const user = JSON.parse(localStorage.getItem('sms_user') || '{}')
  const role = localStorage.getItem('sms_role') || 'guest'

  function logout() {
    localStorage.removeItem('sms_token')
    localStorage.removeItem('sms_user')
    localStorage.removeItem('sms_role')
    window.navigate('/')
  }

  return (
    <header className="header">
      <div className="header-container">
        <div className="logo">
          <h1>📚 School Management</h1>
        </div>
        <nav className="nav">
          {user.email && (
            <>
              <button className="nav-btn" onClick={() => window.navigate('/dashboard')}>Dashboard</button>
              {role === 'teacher' && <button className="nav-btn" onClick={() => window.navigate('/teacher')}>Teach</button>}
              {role === 'admin' && <button className="nav-btn" onClick={() => window.navigate('/admin')}>Admin</button>}
              <button className="nav-btn" onClick={() => window.navigate('/profile')}>Profile</button>
              <button className="nav-btn logout" onClick={logout}>Logout ({user.first_name || 'User'})</button>
            </>
          )}
          {!user.email && (
            <>
              <button className="nav-btn" onClick={() => window.navigate('/')}>Login</button>
              <button className="nav-btn" onClick={() => window.navigate('/register')}>Register</button>
            </>
          )}
        </nav>
      </div>
    </header>
  )
}
