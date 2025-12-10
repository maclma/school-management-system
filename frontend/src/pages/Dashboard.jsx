import React, { useEffect, useState } from 'react'
import api from '../api'
import { showToast } from '../ui'

export default function Dashboard(){
  const [user, setUser] = useState(null)
  const [courses, setCourses] = useState([])
  const [enrollments, setEnrollments] = useState([])
  const [grades, setGrades] = useState([])
  const [attendance, setAttendance] = useState([])

  useEffect(()=>{ load() }, [])

  async function load(){
    try{ const p = await api.getProfile(); setUser(p.data||p) }catch(e){ showToast('Auth required','error'); return }
    const uid = (user && (user.ID||user.id)) || (JSON.parse(localStorage.getItem('sms_user')||'{}').id)
    try{ const c = await api.getCourses(); setCourses(c.data||c||[]) }catch(e){ console.error(e) }
    try{ const e = await api.getEnrollmentsByStudent(uid); setEnrollments(e.data||e||[]) }catch(e){ console.error(e) }
    try{ const g = await api.getGradesByStudent(uid); setGrades(g.data||g||[]) }catch(e){ console.error(e) }
    try{ const a = await api.getAttendanceByStudent(uid); setAttendance(a.data||a||[]) }catch(e){ console.error(e) }
  }

  async function enroll(courseId){
    const uid = (user && (user.ID||user.id)) || (JSON.parse(localStorage.getItem('sms_user')||'{}').id)
    try{ await api.enroll(uid, courseId); showToast('Enrolled','success'); load() }catch(e){ showToast(e.message||'Failed to enroll','error') }
  }

  return (
    <main className="container">
      <div className="card"><pre>{JSON.stringify(user||{},null,2)}</pre></div>
      <div className="grid">
        <div className="panel">
          <h2>Courses</h2>
          {courses.length===0? <p>No courses</p> : courses.map(c=> (
            <div key={c.id} className="course-item">
              <strong>{c.title}</strong>
              <div className="muted-small">{c.department}</div>
              <p>{c.description}</p>
              <div className="flex"><button onClick={()=>enroll(c.id)}>Enroll</button><a className="small muted-small" href={`/course.html?id=${c.id}`}>Details</a></div>
            </div>
          ))}

          <h3 style={{marginTop:12}}>Your Enrollments</h3>
          {enrollments.length===0? <p>No enrollments</p> : enrollments.map(e=> (<div key={e.id} className="row"><div className="col">Course: <strong>{e.course_id}</strong><div className="muted-small">{e.status}</div></div></div>))}
        </div>

        <aside className="panel">
          <h2>Grades</h2>
          {grades.length===0? <p>No grades</p> : (
            <div>
              <div className="summary"><div className="tile"><div className="small">Entries</div><strong>{grades.length}</strong></div></div>
              <div style={{marginTop:8}}>{grades.slice(0,8).map(g=> <div key={g.id} className="row"><div className="col"><strong>{g.grade}</strong><div className="muted-small">Course: {g.course_id}</div></div><div className="small">{g.score}/{g.max_score}</div></div>)}</div>
            </div>
          )}

          <h2 style={{marginTop:12}}>Attendance</h2>
          {attendance.length===0? <p>No attendance</p> : (
            <div>
              <div className="summary"><div className="tile"><div className="small">Total</div><strong>{attendance.length}</strong></div></div>
              <div style={{marginTop:8}}>{/* simple list */} {attendance.slice(0,8).map(a=> <div key={a.id} className="row"><div className="col"><div className="muted-small">{a.date} - {a.status}</div></div></div>)}</div>
            </div>
          )}
        </aside>
      </div>
    </main>
  )
}
