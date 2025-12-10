import React, { useEffect, useState, useRef } from 'react'
import api from '../api'
import { showToast, setLoading } from '../ui'

export default function Profile() {
  const [user, setUser] = useState(null)
  const [isEditing, setIsEditing] = useState(false)
  const fnRef = useRef()
  const lnRef = useRef()
  const phoneRef = useRef()
  const addressRef = useRef()
  const btnRef = useRef()

  useEffect(() => {
    loadProfile()
  }, [])

  async function loadProfile() {
    try {
      const p = await api.getProfile()
      const userData = p.data || p
      setUser(userData)
      if (fnRef.current) fnRef.current.value = userData.first_name || ''
      if (lnRef.current) lnRef.current.value = userData.last_name || ''
      if (phoneRef.current) phoneRef.current.value = userData.phone || ''
      if (addressRef.current) addressRef.current.value = userData.address || ''
    } catch (e) {
      showToast('Failed to load profile', 'error')
      setTimeout(() => window.navigate('/'), 1000)
    }
  }

  async function updateProfile(e) {
    e.preventDefault()
    setLoading(btnRef, true, 'Saving...')
    try {
      const payload = {
        first_name: fnRef.current.value,
        last_name: lnRef.current.value,
        phone: phoneRef.current.value,
        address: addressRef.current.value,
      }
      await api.updateProfile(payload)
      showToast('Profile updated', 'success')
      setIsEditing(false)
      await loadProfile()
    } catch (e) {
      showToast(e.message || 'Failed to update profile', 'error')
    } finally {
      setLoading(btnRef, false)
    }
  }

  if (!user) return <main className="container"><div className="card"><p>Loading...</p></div></main>

  return (
    <main className="container">
      <div className="card">
        <h1>My Profile</h1>
        {!isEditing ? (
          <>
            <div className="info-grid">
              <div className="info-item">
                <div className="label">Name</div>
                <div className="value">{user.first_name} {user.last_name}</div>
              </div>
              <div className="info-item">
                <div className="label">Email</div>
                <div className="value">{user.email}</div>
              </div>
              <div className="info-item">
                <div className="label">Role</div>
                <div className="value badge">{user.role}</div>
              </div>
              <div className="info-item">
                <div className="label">Phone</div>
                <div className="value">{user.phone || '—'}</div>
              </div>
              <div className="info-item">
                <div className="label">Address</div>
                <div className="value">{user.address || '—'}</div>
              </div>
              <div className="info-item">
                <div className="label">Status</div>
                <div className="value">{user.is_active ? '✓ Active' : '⊘ Inactive'}</div>
              </div>
            </div>
            <button className="secondary" onClick={() => setIsEditing(true)}>Edit Profile</button>
          </>
        ) : (
          <form onSubmit={updateProfile}>
            <label>First Name</label>
            <input ref={fnRef} required />
            <label>Last Name</label>
            <input ref={lnRef} required />
            <label>Phone</label>
            <input ref={phoneRef} type="tel" />
            <label>Address</label>
            <textarea ref={addressRef} rows="3"></textarea>
            <div style={{ display: 'flex', gap: '8px', marginTop: '12px' }}>
              <button ref={btnRef} type="submit">Save Changes</button>
              <button type="button" className="secondary" onClick={() => setIsEditing(false)}>Cancel</button>
            </div>
          </form>
        )}
      </div>
    </main>
  )
}
