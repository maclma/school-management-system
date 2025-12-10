// Small UI helpers: toast notifications and simple loading state
(function(){
  function showToast(message, type='info', timeout=3000){
    let container = document.getElementById('sms-toast-container');
    if (!container){
      container = document.createElement('div');
      container.id = 'sms-toast-container';
      container.style.position = 'fixed';
      container.style.right = '16px';
      container.style.top = '16px';
      container.style.zIndex = 9999;
      document.body.appendChild(container);
    }

    const el = document.createElement('div');
    el.className = 'sms-toast ' + type;
    el.textContent = message;
    container.appendChild(el);
    setTimeout(()=>{
      el.style.opacity = 0;
      setTimeout(()=> el.remove(), 300);
    }, timeout);
  }

  function setLoading(btn, state=true, text){
    if (!btn) return;
    if (state){
      btn.dataset.orig = btn.textContent;
      btn.disabled = true;
      btn.textContent = text || 'Please wait...';
    } else {
      btn.disabled = false;
      btn.textContent = btn.dataset.orig || text || 'Submit';
    }
  }

  window.ui = { showToast, setLoading };
})();
