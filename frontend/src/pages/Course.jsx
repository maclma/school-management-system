import React, { useEffect, useState } from 'react'
import api from '../api'
import { showToast, setLoading } from '../ui'

export default function Course(){
  const [course, setCourse] = useState(null)
  const [loading, setLoadingState] = useState(false)
  useEffect(()=>{ const p = new URLSearchParams(window.location.search); const id = p.get('id'); if (id) load(id) }, [])

  async function load(id){
    try{ const res = await api.getCourse(id); setCourse(res) }catch(e){ console.error(e); showToast('Failed to load', 'error') }
  }

  async function enroll(){
    const user = JSON.parse(localStorage.getItem('sms_user')||'{}')
    const id = (course && (course.id || course.ID))
    if (!user.id) return showToast('Login required','error')
    setLoadingState(true)
    try{ await api.enroll(user.id, id); showToast('Enrolled','success') }catch(e){ showToast(e.message||'Failed','error') }
    setLoadingState(false)
  }

  if (!course) return (<main className="container"><div className="card">Loading course...</div></main>)
  return (
    <main className="container">
      <div className="card">
        <h1>{course.title}</h1>
        <div className="muted-small">{course.department}</div>
        <p>{course.description}</p>
        <div className="controls"><button onClick={enroll} disabled={loading}>{loading? 'Enrolling...':'Enroll'}</button></div>
      </div>
    </main>
  )
}
