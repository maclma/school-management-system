Simple static frontend to test login/profile endpoints.

Files:
- `index.html` — Login page (POSTs to `http://localhost:8080/api/auth/login`).
- `login.js` — Handles the form submit and stores `sms_token` in localStorage.
- `profile.html` — Fetches `GET /api/profile` with `Authorization: Bearer <token>` and shows user info.
- `styles.css` — Small CSS to make it look pleasant.
 - `api.js` — Small wrapper for API calls (handles token header and error parsing)
 - `register.html`, `register.js` — Simple registration flow
 - `dashboard.html`, `dashboard.js` — Basic dashboard to view courses & enroll

How to run locally (PowerShell):

```powershell
# From repo root
# 1) Ensure backend is running on http://localhost:8080
# 2) Serve the frontend directory (needs a simple static server to avoid CORS issues)
# If you have Python installed:
python -m http.server 3000 --directory frontend
# Then open http://localhost:3000 in your browser.

# Alternative (if you have Node installed):
# npm i -g serve
serve -s frontend -l 3000
```

CORS note:
- If the backend blocks cross-origin requests, either run the frontend from the same origin (same host:port) or enable CORS in the backend Gin server (middleware). For quick testing, run the frontend server and call the API; if you see CORS errors in the browser console, enable a permissive CORS middleware while testing.

Security note:
- This is a demo template. Do not use localStorage for highly sensitive tokens in production without proper protections.

Quick usage:

1. Start backend: `go run ./cmd/server`
2. Start frontend: `go run ./frontend/server.go` or `python -m http.server 3000 --directory frontend`
3. Open `http://localhost:3000`

Pages:
- `/index.html` — login
- `/register.html` — create account
- `/dashboard.html` — after login: dashboard showing courses and enrollments
 - `/course.html?id=<id>` — course details and enroll button

UI notes:
- `ui.js` provides `ui.showToast(message, type)` and `ui.setLoading(button, state, text)` for simple notifications and button loading states.
- `api.js` is a small wrapper that attaches `window.api` helper methods for common API calls.
 - Dashboard now includes grades and attendance summaries.
 - Styles are responsive and use simple tiles for key metrics.
