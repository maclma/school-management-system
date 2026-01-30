import React, { useEffect, useState } from 'react'
import api from '../api'
import { showToast } from '../ui'

export default function Dashboard(){
  const [user, setUser] = useState(null)
  const [courses, setCourses] = useState([])
  const [enrollments, setEnrollments] = useState([])
  const [grades, setGrades] = useState([])
  const [attendance, setAttendance] = useState([])
  const [search, setSearch] = useState('')

  useEffect(()=>{ load() }, [])

  async function load(){
    try{ const p = await api.getProfile(); setUser(p.data||p) }catch(e){ showToast('Auth required','error'); return }
    const uid = (user && (user.ID||user.id)) || (JSON.parse(localStorage.getItem('sms_user')||'{}').id)
    try{ const c = await api.getCourses(); setCourses(c.data||c||[]) }catch(e){ console.error(e) }
    try{ const e = await api.getStudentEnrollments(); setEnrollments(e.data||e||[]) }catch(e){ console.error(e) }
    try{ const g = await api.getStudentGrades(); setGrades(g.data||g||[]) }catch(e){ console.error(e) }
    try{ const a = await api.getStudentAttendance(); setAttendance(a.data||a||[]) }catch(e){ console.error(e) }
  }

  async function load(){
    try{ const p = await api.getProfile(); setUser(p.data||p) }catch(e){ showToast('Auth required','error'); return }
    const uid = (user && (user.ID||user.id)) || (JSON.parse(localStorage.getItem('sms_user')||'{}').id)
    try{ const c = await api.getCourses(); setCourses(c.data||c||[]) }catch(e){ console.error(e) }
    try{ const e = await api.getStudentEnrollments(); setEnrollments(e.data||e||[]) }catch(e){ console.error(e) }
    try{ const g = await api.getStudentGrades(); setGrades(g.data||g||[]) }catch(e){ console.error(e) }
    try{ const a = await api.getStudentAttendance(); setAttendance(a.data||a||[]) }catch(e){ console.error(e) }
  }

  async function enroll(courseId){
    const uid = (user && (user.ID||user.id)) || (JSON.parse(localStorage.getItem('sms_user')||'{}').id)
    try{ await api.enroll(uid, courseId); showToast('Enrolled','success'); load() }catch(e){ showToast(e.message||'Failed to enroll','error') }
  }

  const filteredCourses = courses.filter(c => 
    c.title.toLowerCase().includes(search.toLowerCase()) || 
    c.department.toLowerCase().includes(search.toLowerCase()) ||
    c.description.toLowerCase().includes(search.toLowerCase())
  )

  return (
    <main className="container">
      <div className="card"><pre>{JSON.stringify(user||{},null,2)}</pre></div>
      <div className="grid">
        <div className="panel">
          <h2>Courses</h2>
          <input type="text" placeholder="Search courses..." value={search} onChange={e => setSearch(e.target.value)} style={{marginBottom:8, width:'100%'}} />
          {filteredCourses.length===0? <p>No courses</p> : filteredCourses.map(c=> (
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
            <div className="table-container">
              <table className="simple-table">
                <thead>
                  <tr>
                    <th>Course</th>
                    <th>Grade</th>
                    <th>Score</th>
                    <th>Date</th>
                  </tr>
                </thead>
                <tbody>
                  {grades.map(g => (
                    <tr key={g.id}>
                      <td>{g.course_id}</td>
                      <td>{g.grade_value || g.grade}</td>
                      <td>{g.score || 'N/A'}</td>
                      <td>{g.date_recorded || g.created_at}</td>
                    </tr>
                  ))}
                </tbody>
              </table>
            </div>
          )}

          <h2 style={{marginTop:12}}>Attendance</h2>
          {attendance.length===0? <p>No attendance</p> : (
            <div className="table-container">
              <table className="simple-table">
                <thead>
                  <tr>
                    <th>Course</th>
                    <th>Date</th>
                    <th>Status</th>
                  </tr>
                </thead>
                <tbody>
                  {attendance.map(a => (
                    <tr key={a.id}>
                      <td>{a.course_id}</td>
                      <td>{a.attendance_date || a.date}</td>
                      <td>{a.status}</td>
                    </tr>
                  ))}
                </tbody>
              </table>
            </div>
          )}
        </aside>
      </div>
    </main>
  )
}
