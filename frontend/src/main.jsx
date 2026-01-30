import React from 'react'
import { createRoot } from 'react-dom/client'
import App from './App'
import './styles.css'

const rootElement = document.getElementById('root')
if (!rootElement) {
  console.error('Root element not found!')
  document.body.innerHTML = '<h1>Error: Root element not found</h1>'
} else {
  createRoot(rootElement).render(
    <React.StrictMode>
      <App />
    </React.StrictMode>
  )
}
