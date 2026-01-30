import React from 'react'

export default function TestApp() {
  return (
    <div style={{ padding: '20px', fontFamily: 'sans-serif' }}>
      <h1>Test Page - Server is Running!</h1>
      <p style={{ fontSize: '16px' }}>
        If you can see this, the frontend is working correctly.
      </p>
      <div style={{ marginTop: '20px', padding: '10px', background: '#f0f0f0', borderRadius: '5px' }}>
        <p><strong>Backend:</strong> http://localhost:8080</p>
        <p><strong>Frontend:</strong> http://localhost:3000</p>
      </div>
      <button 
        onClick={() => {
          try {
            console.log('Testing API call...')
            fetch('/api/health')
              .then(r => r.json())
              .then(data => {
                console.log('API Response:', data)
                alert('API Connected!\n\n' + JSON.stringify(data, null, 2))
              })
              .catch(e => {
                console.error('API Error:', e)
                alert('API Error: ' + e.message)
              })
          } catch (e) {
            console.error('Test error:', e)
            alert('Error: ' + e.message)
          }
        }}
        style={{
          padding: '10px 20px',
          fontSize: '16px',
          background: '#2563eb',
          color: 'white',
          border: 'none',
          borderRadius: '5px',
          cursor: 'pointer'
        }}
      >
        Test API Connection
      </button>
    </div>
  )
}
