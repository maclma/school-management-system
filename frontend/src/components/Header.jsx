import React, { useState } from 'react'

export default function Header() {
  const user = JSON.parse(localStorage.getItem('sms_user') || '{}')
  const role = localStorage.getItem('sms_role') || 'guest'
  const [isMenuOpen, setIsMenuOpen] = useState(false)

  function logout() {
    localStorage.removeItem('sms_token')
    localStorage.removeItem('sms_user')
    localStorage.removeItem('sms_role')
    window.navigate('/')
  }

  return (
    <header className="header" style={{visibility: 'visible'}}>
      <div className="header-container">
        <div className="logo">
          <h1>📚 School Management</h1>
        </div>
        <button className="hamburger" onClick={() => setIsMenuOpen(!isMenuOpen)}>
          ☰
        </button>
        <nav className={`nav ${isMenuOpen ? 'open' : ''}`}>
          {user.email && (
            <>
              <button className="nav-btn" onClick={() => { window.navigate('/dashboard'); setIsMenuOpen(false) }}>Dashboard</button>
              {(role === 'student' || role === 'teacher') && <button className="nav-btn" onClick={() => { window.navigate('/assignments'); setIsMenuOpen(false) }}>Assignments</button>}
              {role === 'teacher' && <button className="nav-btn" onClick={() => { window.navigate('/teacher'); setIsMenuOpen(false) }}>Teach</button>}
              {role === 'admin' && <button className="nav-btn" onClick={() => { window.navigate('/admin'); setIsMenuOpen(false) }}>Admin</button>}
              <button className="nav-btn" onClick={() => { window.navigate('/profile'); setIsMenuOpen(false) }}>Profile</button>
              <button className="nav-btn logout" onClick={() => { logout(); setIsMenuOpen(false) }}>Logout ({user.first_name || 'User'})</button>
            </>
          )}
          {!user.email && (
            <>
              <button className="nav-btn" onClick={() => { window.navigate('/'); setIsMenuOpen(false) }}>Login</button>
              <button className="nav-btn" onClick={() => { window.navigate('/register'); setIsMenuOpen(false) }}>Register</button>
            </>
          )}
        </nav>
      </div>
    </header>
  )
}
