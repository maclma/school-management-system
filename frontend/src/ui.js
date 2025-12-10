// Simple UI helpers for React components
export function showToast(message, type='info', timeout=3000){
  const containerId = 'sms-toast-container'
  let container = document.getElementById(containerId)
  if (!container){
    container = document.createElement('div')
    container.id = containerId
    container.style.position='fixed'
    container.style.right='16px'
    container.style.top='16px'
    container.style.zIndex=9999
    document.body.appendChild(container)
  }
  const el = document.createElement('div')
  el.className = 'sms-toast ' + type
  el.textContent = message
  container.appendChild(el)
  setTimeout(()=>{ el.style.opacity='0'; setTimeout(()=>el.remove(),300) }, timeout)
}

export function setLoading(buttonRef, state=true, text='Please wait...'){
  if (!buttonRef || !buttonRef.current) return
  const btn = buttonRef.current
  if (state){ btn.dataset.orig = btn.textContent; btn.disabled = true; btn.textContent = text }
  else { btn.disabled = false; btn.textContent = btn.dataset.orig || 'Submit' }
}
