import React, { useState } from 'react'
import { ErrorBoundary } from './components/ErrorBoundary'
import Login from './pages/Login'
import Register from './pages/Register'
import Dashboard from './pages/Dashboard'
import Course from './pages/Course'
import Profile from './pages/Profile'
import TeacherPanel from './pages/TeacherPanel'
import AdminDashboard from './pages/AdminDashboard'
import EnrollmentApproval from './pages/EnrollmentApproval'
import TeacherAssignments from './pages/TeacherAssignments'
import AssignmentSubmissions from './pages/AssignmentSubmissions'
import StudentAssignments from './pages/StudentAssignments'
import Header from './components/Header'

// Minimal router based on path
function AppContent(){
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
       path.startsWith('/teacher/assignments') && role === 'teacher' ? <TeacherAssignments /> :
       path.match(/^\/teacher\/assignments\/\d+\/submissions/) && role === 'teacher' ? <AssignmentSubmissions /> :
       path.startsWith('/teacher') && role === 'teacher' ? <TeacherPanel /> :
       path.startsWith('/admin/enrollments') && role === 'admin' ? <EnrollmentApproval /> :
       path.startsWith('/admin') && role === 'admin' ? <AdminDashboard /> :
       path.startsWith('/assignments') && (role === 'student' || role === 'teacher') ? <StudentAssignments /> :
       // default to login
       <Login />}
    </>
  )
}

function App() {
  return (
    <ErrorBoundary>
      <AppContent />
    </ErrorBoundary>
  )
}

export default App
