(function(){
  const form = document.getElementById('login-form');
  const msg = document.getElementById('message');
  const btn = document.getElementById('login-btn');

  form.addEventListener('submit', async (e) => {
    e.preventDefault();
    msg.textContent = '';
    window.ui && window.ui.setLoading(btn, true, 'Signing in...');

    const email = document.getElementById('email').value.trim();
    const password = document.getElementById('password').value;

    if (!email || !password) {
      window.ui && window.ui.showToast('Please enter email and password', 'error');
      window.ui && window.ui.setLoading(btn, false);
      return;
    }

    try {
      const body = await window.api.login(email, password).catch(async (e) => { throw e; });

      // extract token
      const token = (body.data && body.data.token) || body.token || '';
      if (!token) {
        window.ui && window.ui.showToast('Login succeeded but token missing', 'error');
        window.ui && window.ui.setLoading(btn, false);
        return;
      }

      localStorage.setItem('sms_token', token);
      if (body.data && body.data.user) localStorage.setItem('sms_user', JSON.stringify(body.data.user));

      window.ui && window.ui.showToast('Logged in', 'success');
      setTimeout(()=> window.location.href = 'dashboard.html', 500);
    } catch (err) {
      console.error(err);
      const msgText = (err && err.message) || 'Network or server error';
      window.ui && window.ui.showToast(msgText, 'error');
      window.ui && window.ui.setLoading(btn, false);
    }
  });
})();
