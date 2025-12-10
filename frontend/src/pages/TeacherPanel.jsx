import React, { useEffect, useState, useRef } from 'react'
import api from '../api'
import { showToast, setLoading } from '../ui'

export default function TeacherPanel() {
  const [courses, setCourses] = useState([])
  const [selectedCourse, setSelectedCourse] = useState(null)
  const [enrollments, setEnrollments] = useState([])
  const [tab, setTab] = useState('grades')
  const [showModal, setShowModal] = useState(false)
  const [modalType, setModalType] = useState('grade')
  const [selectedStudent, setSelectedStudent] = useState(null)
  const [search, setSearch] = useState('')

  const gradeRef = useRef(); const dateRef = useRef(); const statusRef = useRef()
  const courseNameRef = useRef(); const courseCodeRef = useRef(); const deptRef = useRef()
  const creditsRef = useRef(); const capacityRef = useRef(); const btnRef = useRef()

  useEffect(() => { loadCourses() }, [])
  useEffect(() => { if (selectedCourse) loadEnrollments() }, [selectedCourse])

  async function loadCourses() {
    try { const c = await api.getCourses(); setCourses(c.data || c || []); if ((c.data || c || []).length > 0) setSelectedCourse((c.data || c)[0].id) } catch (e) { showToast('Failed to load courses', 'error') }
  }

  async function loadEnrollments() {
    try { const e = await api.getCourseEnrollments(selectedCourse); setEnrollments(e.data || e || []) } catch (e) { console.error('Failed to load enrollments:', e) }
  }

  async function submitGrade(e) {
    e.preventDefault(); setLoading(btnRef, true, 'Recording...')
    try {
      const payload = { student_id: selectedStudent.student_id || selectedStudent.id, course_id: selectedCourse, grade_value: parseFloat(gradeRef.current.value), date_recorded: dateRef.current.value }
      await api.recordGrade(payload); showToast('Grade recorded', 'success'); setShowModal(false); await loadEnrollments()
    } catch (e) { showToast(e.message || 'Failed to record grade', 'error') } finally { setLoading(btnRef, false) }
  }

  async function submitAttendance(e) {
    e.preventDefault(); setLoading(btnRef, true, 'Recording...')
    try {
      const payload = { student_id: selectedStudent.student_id || selectedStudent.id, course_id: selectedCourse, attendance_date: dateRef.current.value, status: statusRef.current.value }
      await api.recordAttendance(payload); showToast('Attendance recorded', 'success'); setShowModal(false); await loadEnrollments()
    } catch (e) { showToast(e.message || 'Failed to record attendance', 'error') } finally { setLoading(btnRef, false) }
  }

  async function submitCreateCourse(e) {
    e.preventDefault(); setLoading(btnRef, true, 'Creating...')
    try {
      const payload = { name: courseNameRef.current.value, code: courseCodeRef.current.value, department: deptRef.current.value, credits: parseInt(creditsRef.current.value), max_capacity: parseInt(capacityRef.current.value) }
      await api.createCourse(payload); showToast('Course created', 'success'); setShowModal(false); await loadCourses()
    } catch (e) { showToast(e.message || 'Failed to create course', 'error') } finally { setLoading(btnRef, false) }
  }

  function openModal(student, type) { setSelectedStudent(student); setModalType(type); setShowModal(true); if (dateRef.current) dateRef.current.value = new Date().toISOString().split('T')[0] }

  const filteredEnrollments = enrollments.filter(e => e.student_id.toString().includes(search) || e.id.toString().includes(search))
  const course = courses.find(c => c.id === selectedCourse)

  return (
    <main className="container">
      <div className="card">
        <h1>Teacher Panel</h1>
        <p className="muted-small">Manage grades, attendance, and courses</p>

        {tab !== 'courses' && (<div style={{ marginBottom: '20px' }}><label>Select Course</label><select value={selectedCourse || ''} onChange={(e) => setSelectedCourse(Number(e.target.value))}><option value="">— Choose a course —</option>{courses.map(c => (<option key={c.id} value={c.id}>{c.name || c.title} ({c.code || ''})</option>))}</select></div>)}

        {selectedCourse && tab !== 'courses' || courses.length > 0 ? (
          <>
            <div className="tabs">
              <button className={`tab ${tab === 'grades' ? 'active' : ''}`} onClick={() => setTab('grades')}>Grades</button>
              <button className={`tab ${tab === 'attendance' ? 'active' : ''}`} onClick={() => setTab('attendance')}>Attendance</button>
              <button className={`tab ${tab === 'courses' ? 'active' : ''}`} onClick={() => setTab('courses')}>My Courses</button>
            </div>

            {tab !== 'courses' ? (
              <div className="panel">
                <h2>{course?.name || course?.title}</h2>
                <div style={{ marginBottom: '12px' }}><input type="text" placeholder="Search by student ID..." value={search} onChange={(e) => setSearch(e.target.value)} style={{ maxWidth: '300px' }} /></div>
                {filteredEnrollments.length === 0 ? <p className="muted-small">{search ? 'No matching students' : 'No enrolled students'}</p> : (
                  <div className="table-container">
                    <table className="simple-table">
                      <thead><tr><th>Student ID</th><th>Status</th><th>Action</th></tr></thead>
                      <tbody>{filteredEnrollments.map(e => (<tr key={e.id}><td>{e.student_id}</td><td>{e.status || 'active'}</td><td><button className="small" onClick={() => openModal(e, tab === 'grades' ? 'grade' : 'attendance')}>{tab === 'grades' ? 'Grade' : 'Attend'}</button></td></tr>))}</tbody>
                    </table>
                  </div>
                )}
              </div>
            ) : (
              <div className="panel">
                <h2>My Courses</h2>
                <button className="secondary" onClick={() => { setModalType('createCourse'); setShowModal(true); }} style={{ marginBottom: '12px' }}>+ New Course</button>
                {courses.length === 0 ? <p className="muted-small">No courses created yet</p> : (
                  <div className="grid" style={{ gridTemplateColumns: 'repeat(auto-fill, minmax(250px, 1fr))' }}>
                    {courses.map(c => (<div key={c.id} className="card" style={{ marginBottom: '12px' }}><h3 style={{ marginTop: 0 }}>{c.name || c.title}</h3><div className="muted-small">{c.code}</div><div className="muted-small">{c.department}</div><div style={{ marginTop: '8px' }}><span className="badge">{c.credits} credits</span><span className="badge" style={{ marginLeft: '4px' }}>Cap: {c.max_capacity}</span></div></div>))}
                  </div>
                )}
              </div>
            )}
          </>
        ) : null}
      </div>

      {showModal && (
        <div className="modal-overlay" onClick={() => setShowModal(false)}>
          <div className="modal-content" onClick={(e) => e.stopPropagation()}>
            {modalType === 'grade' ? (
              <><h2>Record Grade</h2><form onSubmit={submitGrade}><label>Student ID: {selectedStudent?.student_id || selectedStudent?.id}</label><label>Grade (0-100)</label><input ref={gradeRef} type="number" min="0" max="100" step="0.5" required /><label>Date</label><input ref={dateRef} type="date" required /><div style={{ display: 'flex', gap: '8px', marginTop: '16px' }}><button ref={btnRef} type="submit">Save</button><button type="button" className="secondary" onClick={() => setShowModal(false)}>Cancel</button></div></form></>
            ) : modalType === 'attendance' ? (
              <><h2>Record Attendance</h2><form onSubmit={submitAttendance}><label>Student ID: {selectedStudent?.student_id || selectedStudent?.id}</label><label>Status</label><select ref={statusRef} required><option value="">— Select —</option><option value="present">Present</option><option value="absent">Absent</option><option value="late">Late</option></select><label>Date</label><input ref={dateRef} type="date" required /><div style={{ display: 'flex', gap: '8px', marginTop: '16px' }}><button ref={btnRef} type="submit">Save</button><button type="button" className="secondary" onClick={() => setShowModal(false)}>Cancel</button></div></form></>
            ) : (
              <><h2>Create New Course</h2><form onSubmit={submitCreateCourse}><label>Course Name</label><input ref={courseNameRef} required /><label>Course Code</label><input ref={courseCodeRef} required /><label>Department</label><input ref={deptRef} required /><label>Credits</label><input ref={creditsRef} type="number" min="1" max="10" required /><label>Max Capacity</label><input ref={capacityRef} type="number" min="1" required /><div style={{ display: 'flex', gap: '8px', marginTop: '16px' }}><button ref={btnRef} type="submit">Create</button><button type="button" className="secondary" onClick={() => setShowModal(false)}>Cancel</button></div></form></>
            )}
          </div>
        </div>
      )}
    </main>
  )
}
