(function(){
  const form = document.getElementById('register-form');
  const msg = document.getElementById('message');
  const btn = form.querySelector('button[type=submit]');

  form.addEventListener('submit', async (e)=>{
    e.preventDefault();
    msg.textContent = '';
    window.ui && window.ui.setLoading(btn, true, 'Creating account...');

    const payload = {
      first_name: document.getElementById('first_name').value.trim(),
      last_name: document.getElementById('last_name').value.trim(),
      email: document.getElementById('email').value.trim(),
      password: document.getElementById('password').value,
      role: document.getElementById('role').value || 'student'
    };

    try{
      const res = await window.api.register(payload);
      window.ui && window.ui.showToast(res.message || 'Registered', 'success');
      setTimeout(()=> window.location.href = 'index.html', 700);
    }catch(err){
      console.error(err);
      const m = (err && err.message) || (err && err.body && err.body.error) || 'Registration failed';
      window.ui && window.ui.showToast(m, 'error');
      window.ui && window.ui.setLoading(btn, false);
    }
  });
})();
