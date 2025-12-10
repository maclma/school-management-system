(function(){
  const userEl = document.getElementById('user');
  const coursesEl = document.getElementById('courses-list');
  const enrollmentsEl = document.getElementById('enrollments-list');
  const logoutBtn = document.getElementById('logout');

  let currentUser = null;

  async function load(){
    try{
      const profile = await window.api.getProfile();
      currentUser = profile.data || profile;
      userEl.innerHTML = '<pre>'+JSON.stringify(currentUser, null, 2)+'</pre>';
    }catch(err){
      userEl.innerHTML = '<p>Failed to load profile. Please <a href="index.html">login</a>.</p>';
      console.error(err);
      return;
    }

    try{
      const coursesRes = await window.api.getCourses(1, 50);
      const courses = coursesRes.data || coursesRes;
      if (!courses || courses.length === 0) coursesEl.innerHTML = '<p>No courses found.</p>';
      else {
        coursesEl.innerHTML = '';
        courses.forEach(c => {
          const div = document.createElement('div');
          div.className = 'course-item';
          div.innerHTML = `<strong>${c.title}</strong> <div class="muted">${c.department || ''}</div> <p>${c.description || ''}</p>`;
          const btnWrap = document.createElement('div');
          btnWrap.className = 'flex';
          const btn = document.createElement('button');
          btn.textContent = 'Enroll';
          btn.addEventListener('click', ()=> enroll(c.id, btn));
          const details = document.createElement('a');
          details.href = `course.html?id=${c.id}`;
          details.textContent = 'Details';
          details.className = 'small muted-small';
          btnWrap.appendChild(btn);
          btnWrap.appendChild(details);
          div.appendChild(btnWrap);
          coursesEl.appendChild(div);
        });
      }
    }catch(err){
      coursesEl.innerHTML = '<p>Failed to fetch courses.</p>';
      console.error(err);
    }

    // load enrollments
    try{
      const enr = await window.api.getEnrollmentsByStudent(currentUser.ID || currentUser.id || currentUser.user_id || currentUser.student_id);
      const data = enr.data || enr;
      if (!data || data.length === 0) enrollmentsEl.innerHTML = '<p>No enrollments yet.</p>';
      else {
        enrollmentsEl.innerHTML = '';
        data.forEach(e => {
          const d = document.createElement('div');
          d.className = 'row';
          d.innerHTML = `<div class="col">Course: <strong>${e.course_id}</strong><div class="muted-small">Status: ${e.status}</div></div>`;
          enrollmentsEl.appendChild(d);
        });
      }
    }catch(err){
      enrollmentsEl.innerHTML = '<p>Failed to load enrollments.</p>';
      console.error(err);
    }

      // load grades summary
      try{
        const gradesRes = await window.api.getGradesByStudent(currentUser.ID || currentUser.id || currentUser.user_id || currentUser.student_id);
        const grades = gradesRes.data || gradesRes || [];
        const gradesListEl = document.getElementById('grades-list');
        if (!grades || grades.length === 0) {
          gradesListEl.innerHTML = '<p>No grades recorded.</p>';
        } else {
          // compute average score if scores available
          let total = 0; let count = 0;
          grades.forEach(g=>{ if (typeof g.score === 'number') { total += g.score; count++ } });
          const avg = count ? (total / count).toFixed(2) : 'N/A';
          let html = `<div class="summary"><div class="tile"><div class="small">Avg Score</div><strong>${avg}</strong></div><div class="tile"><div class="small">Entries</div><strong>${grades.length}</strong></div></div>`;
          html += '<div style="margin-top:.75rem">';
          grades.slice(0,6).forEach(g=>{
            html += `<div class="row"><div class="col"><strong>${g.grade || g.Grade || 'Grade'}</strong><div class="muted-small">Course: ${g.course_id || g.courseId || '-'}</div></div><div class="small">${g.score!=null?g.score:''}/${g.max_score||g.maxScore||''}</div></div>`;
          });
          html += '</div>';
          gradesListEl.innerHTML = html;
        }
      }catch(err){
        const gradesListEl = document.getElementById('grades-list');
        gradesListEl.innerHTML = '<p>Failed to load grades.</p>';
        console.error(err);
      }

      // load attendance summary
      try{
        const attRes = await window.api.getAttendanceByStudent(currentUser.ID || currentUser.id || currentUser.user_id || currentUser.student_id);
        const attend = attRes.data || attRes || [];
        const attEl = document.getElementById('attendance-list');
        if (!attend || attend.length === 0) {
          attEl.innerHTML = '<p>No attendance records.</p>';
        } else {
          const counts = { present:0, absent:0, late:0, excused:0 };
          attend.forEach(a=>{ const s=(a.status||'').toLowerCase(); if (counts[s]!==undefined) counts[s]++; });
          const total = attend.length;
          const presentPct = total? Math.round((counts.present/total)*100):0;
          let html = `<div class="summary"><div class="tile"><div class="small">Present</div><strong>${counts.present}</strong></div><div class="tile"><div class="small">Absent</div><strong>${counts.absent}</strong></div><div class="tile"><div class="small">Late</div><strong>${counts.late}</strong></div></div>`;
          html += `<div style="margin-top:.6rem"><div class="small">Presence %</div><strong>${presentPct}%</strong></div>`;
          attEl.innerHTML = html;
        }
      }catch(err){
        const attEl = document.getElementById('attendance-list');
        attEl.innerHTML = '<p>Failed to load attendance.</p>';
        console.error(err);
      }
  }

  async function enroll(courseId, btn){
    if (!currentUser) return alert('Not loaded');
    const sid = currentUser.ID || currentUser.id || currentUser.user_id || currentUser.student_id;
    btn.disabled = true;
    btn.textContent = 'Enrolling...';
    try{
      await window.api.createEnrollment(sid, courseId);
      btn.textContent = 'Enrolled';
      load();
    }catch(err){
      btn.disabled = false;
      btn.textContent = 'Enroll';
      alert(err.message || 'Failed to enroll');
    }
  }

  logoutBtn.addEventListener('click', ()=>{
    localStorage.removeItem('sms_token');
    localStorage.removeItem('sms_user');
    window.location.href = 'index.html';
  });

  load();
})();
