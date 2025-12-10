import React, { useRef } from 'react'
import api from '../api'
import { showToast, setLoading } from '../ui'

export default function Register(){
  const fnRef = useRef();
  const lnRef = useRef();
  const emailRef = useRef();
  const passRef = useRef();
  const roleRef = useRef();
  const btnRef = useRef();

  async function submit(e){
    e && e.preventDefault()
    setLoading(btnRef, true, 'Creating...')
    try{
      const payload = { first_name: fnRef.current.value, last_name: lnRef.current.value, email: emailRef.current.value, password: passRef.current.value, role: roleRef.current.value }
      const res = await api.register(payload)
      showToast(res.message || 'Registered', 'success')
      setTimeout(()=> window.navigate('/'), 600)
    }catch(err){
      console.error(err)
      showToast(err.message || 'Registration failed', 'error')
      setLoading(btnRef, false)
    }
  }

  return (
    <main className="container">
      <div className="card">
        <h1>Register</h1>
        <form onSubmit={submit}>
          <label>First name</label>
          <input ref={fnRef} required />
          <label>Last name</label>
          <input ref={lnRef} required />
          <label>Email</label>
          <input ref={emailRef} type="email" required />
          <label>Password</label>
          <input ref={passRef} type="password" required />
          <label>Role</label>
          <select ref={roleRef} defaultValue="student"><option value="student">Student</option><option value="teacher">Teacher</option></select>
          <div style={{display:'flex',gap:'8px',marginTop:'8px'}}>
            <button ref={btnRef} type="submit">Create account</button>
            <button type="button" className="secondary" onClick={()=>window.navigate('/')}>Back</button>
          </div>
        </form>
      </div>
    </main>
  )
}
