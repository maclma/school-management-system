(function(){
  const params = new URLSearchParams(window.location.search);
  const id = params.get('id');
  const card = document.getElementById('course-card');
  const enrollBtn = document.getElementById('enroll-btn');

  if (!id) {
    card.innerHTML = '<p>Missing course id</p>';
  } else {
    async function load(){
      try{
        const res = await window.api.getCourses(1,100);
        const courses = res.data || res;
        const c = (courses || []).find(x => String(x.id) === String(id));
        if (!c) { card.innerHTML = '<p>Course not found</p>'; return; }
        card.innerHTML = `<h2>${c.title}</h2><div class="muted-small">${c.department||''}</div><p>${c.description||''}</p>`;
      }catch(err){
        card.innerHTML = '<p>Failed to load course</p>';
        console.error(err);
      }
    }

    load();
  }

  enrollBtn.addEventListener('click', async ()=>{
    const userRaw = localStorage.getItem('sms_user');
    let uid = null;
    try{ uid = JSON.parse(userRaw).id }catch(e){}
    if (!uid) return window.ui && window.ui.showToast('You must be logged in to enroll', 'error');
    window.ui && window.ui.setLoading(enrollBtn, true, 'Enrolling...');
    try{
      await window.api.createEnrollment(uid, id);
      window.ui && window.ui.showToast('Enrolled', 'success');
      window.ui && window.ui.setLoading(enrollBtn, false);
    }catch(err){
      console.error(err);
      window.ui && window.ui.showToast(err.message || 'Failed to enroll', 'error');
      window.ui && window.ui.setLoading(enrollBtn, false);
    }
  });
})();
