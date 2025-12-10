import React, { useState } from 'react'
import Login from './pages/Login'
import Register from './pages/Register'
import Dashboard from './pages/Dashboard'
import Course from './pages/Course'
import Profile from './pages/Profile'
import TeacherPanel from './pages/TeacherPanel'
import AdminDashboard from './pages/AdminDashboard'
import EnrollmentApproval from './pages/EnrollmentApproval'
import Header from './components/Header'

// Minimal router based on path
function App(){
  const [path, setPath] = useState(window.location.pathname)
  window.navigate = (to)=>{ window.history.pushState({},'',to); setPath(to); }
  window.onpopstate = ()=> setPath(window.location.pathname)

  const token = localStorage.getItem('sms_token')
  const role = localStorage.getItem('sms_role') || 'guest'

  return (
    <>
      <Header />
      {path.startsWith('/register') || path === '/register.html' ? <Register/> :
       path.startsWith('/dashboard') || path === '/dashboard.html' ? <Dashboard/> :
       path.startsWith('/course') ? <Course /> :
       path.startsWith('/profile') ? <Profile /> :
       path.startsWith('/teacher') && role === 'teacher' ? <TeacherPanel /> :
       path.startsWith('/admin/enrollments') && role === 'admin' ? <EnrollmentApproval /> :
       path.startsWith('/admin') && role === 'admin' ? <AdminDashboard /> :
       // default to login
       <Login />}
    </>
  )
}

export default App
