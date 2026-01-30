import React, { useRef } from 'react'
import api from '../api'
import { showToast, setLoading } from '../ui'

export default function Login(){
  const emailRef = useRef();
  const passRef = useRef();
  const btnRef = useRef();

  async function submit(e){
    e && e.preventDefault()
    setLoading(btnRef, true, 'Signing in...')
    try{
      const body = await api.login(emailRef.current.value, passRef.current.value)
      const token = (body.data && body.data.token) || body.token
      const user = (body.data && body.data.user) || body.user
      if (!token) throw new Error('Token missing')
      localStorage.setItem('sms_token', token)
      if (user) {
        localStorage.setItem('sms_user', JSON.stringify(user))
        localStorage.setItem('sms_role', user.role)
      }
      showToast('Logged in', 'success')
      const targetPath = user.role === 'admin' ? '/admin' : user.role === 'teacher' ? '/teacher' : '/dashboard'
      setTimeout(()=> window.navigate(targetPath), 400)
    }catch(err){
      console.error(err)
      showToast(err.message || 'Login failed', 'error')
      setLoading(btnRef, false)
    }
  }

  return (
    <main className="container">
      <div className="card">
        <h1>Login</h1>
        <form onSubmit={submit}>
          <label>Email</label>
          <input ref={emailRef} type="email" required />
          <label>Password</label>
          <input ref={passRef} type="password" required />
          <div style={{display:'flex',gap:'8px',marginTop:'8px'}}>
            <button ref={btnRef} type="submit">Login</button>
            <button type="button" className="secondary" onClick={()=>window.navigate('/register')}>Register</button>
          </div>
        </form>
      </div>
    </main>
  )
}
