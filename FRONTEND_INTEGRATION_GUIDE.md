# Frontend Integration Guide - Session 6 Advanced Features

Complete guide for integrating 26 new API endpoints into the React frontend.

---

## Overview

**6 Advanced Feature Areas:**
1. **Advanced Search** - Multi-criteria search across data
2. **CSV Export Reports** - Download data as CSV files
3. **Attendance Automation** - Attendance tracking and analytics
4. **Grade Auto-Calculation** - Automatic grading and statistics
5. **Assignment Rubrics** - Criterion-based grading system
6. **Dashboard Analytics** - Visual analytics and insights

**Total New Endpoints:** 26  
**New API Methods:** 26 (all in `frontend/src/api.js`)  
**New Components Needed:** 6-8 React components

---

## Project Structure

```
frontend/
├── src/
│   ├── api.js                          # ✅ Updated with 26 new methods
│   ├── App.jsx                         # Update: Add new routes
│   ├── components/
│   │   ├── (existing components)
│   │   ├── AdvancedSearch.jsx          # New: Search interface
│   │   ├── ReportsExport.jsx           # New: Export CSV reports
│   │   ├── AttendanceAnalytics.jsx     # New: Attendance stats
│   │   ├── GradeAnalytics.jsx          # New: Grade stats & distribution
│   │   ├── RubricsManager.jsx          # New: Rubric CRUD & scoring
│   │   └── AdvancedFeatures.jsx        # New: Main dashboard for all features
│   └── pages/
│       └── (existing pages)
```

---

## Installation & Setup

### 1. Ensure Backend Server is Running

```bash
cd school-management-system
go run ./cmd/server
# Server should start on http://localhost:8080
```

### 2. Install Frontend Dependencies

```bash
cd frontend
npm install
```

### 3. Start Development Server

```bash
npm run dev
# Frontend will run on http://localhost:5173
```

### 4. Build for Production

```bash
npm run build
npm run preview
```

---

## API Methods Reference

All 26 new methods are available in `api.js`:

### Search Methods (5)
```javascript
api.searchAnnouncements(query, audience, priority, page, limit)
api.searchPayments(studentId, status, page, limit)
api.searchStudents(query, page, limit)
api.searchGradesByRange(courseId, minScore, maxScore, page, limit)
api.searchOverduePayments()
```

### Export Methods (5)
```javascript
api.exportPaymentsCSV(studentId, status)
api.exportGradesCSV(courseId)
api.exportAttendanceCSV(courseId)
api.exportStudentTranscriptCSV(studentId)
api.exportEnrollmentsCSV(courseId)
```

### Attendance Methods (5)
```javascript
api.getAttendanceStatsByCourse(courseId)
api.getStudentAttendancePercentage(studentId, courseId)
api.checkLowAttendance({ student_id, course_id, threshold })
api.getStudentsWithLowAttendance(threshold, courseId)
api.getAttendanceReport(courseId)
```

### Grade Methods (4)
```javascript
api.recordGradeWithAutoCalc({ student_id, course_id, score, max_score, graded_by })
api.getCourseAverageGrade(courseId)
api.getGradeDistribution(courseId)
api.getStudentGradeStats(studentId)
```

### Rubric Methods (7)
```javascript
api.createRubric({ assignment_id, name, criteria, is_active })
api.getRubric(id)
api.getRubricsByAssignment(assignmentId)
api.updateRubric(id, { name, criteria, is_active })
api.deleteRubric(id)
api.scoreSubmission({ submission_id, rubric_id, criterion_scores, feedback, graded_by })
api.getSubmissionScore(submissionId)
```

---

## Component Examples

### 1. Advanced Search Component

**Purpose:** Multi-criteria search interface for finding announcements, payments, students, grades

**Features:**
- Search by type (announcements, payments, students, grades)
- Filters (audience, priority, status, score range)
- Pagination
- Results display in table/card format

**Usage:**
```jsx
import AdvancedSearch from './components/AdvancedSearch'

export default function App() {
  return (
    <AdvancedSearch />
  )
}
```

**Key Methods Used:**
- `api.searchAnnouncements()`
- `api.searchPayments()`
- `api.searchStudents()`
- `api.searchGradesByRange()`
- `api.searchOverduePayments()`

---

### 2. Reports & Export Component

**Purpose:** Generate and download CSV reports

**Features:**
- Select report type (payments, grades, attendance, transcripts, enrollments)
- Apply filters (student, course, status)
- Download CSV file
- Preview data before export

**Usage:**
```jsx
import ReportsExport from './components/ReportsExport'

export default function App() {
  return (
    <ReportsExport />
  )
}
```

**Key Methods Used:**
- `api.exportPaymentsCSV()`
- `api.exportGradesCSV()`
- `api.exportAttendanceCSV()`
- `api.exportStudentTranscriptCSV()`
- `api.exportEnrollmentsCSV()`

---

### 3. Attendance Analytics Component

**Purpose:** Track and visualize attendance data

**Features:**
- Course attendance statistics
- Student attendance percentage
- Low attendance alerts
- Bulk attendance report
- Student list by attendance threshold

**Usage:**
```jsx
import AttendanceAnalytics from './components/AttendanceAnalytics'

export default function App() {
  return (
    <AttendanceAnalytics />
  )
}
```

**Key Methods Used:**
- `api.getAttendanceStatsByCourse()`
- `api.getStudentAttendancePercentage()`
- `api.checkLowAttendance()`
- `api.getStudentsWithLowAttendance()`
- `api.getAttendanceReport()`

---

### 4. Grade Analytics Component

**Purpose:** View grade statistics and distributions

**Features:**
- Student grade statistics (count, average, A/B/C/D/F distribution)
- Course average grades
- Grade distribution chart (pie/bar chart)
- Auto-grade recording with letter calculation

**Usage:**
```jsx
import GradeAnalytics from './components/GradeAnalytics'

export default function App() {
  return (
    <GradeAnalytics />
  )
}
```

**Key Methods Used:**
- `api.recordGradeWithAutoCalc()`
- `api.getCourseAverageGrade()`
- `api.getGradeDistribution()`
- `api.getStudentGradeStats()`

---

### 5. Rubrics Manager Component

**Purpose:** Create and manage grading rubrics

**Features:**
- Create rubric with criteria
- View rubric details
- Edit/delete rubrics
- Score submissions using rubric
- View submission scores

**Usage:**
```jsx
import RubricsManager from './components/RubricsManager'

export default function App() {
  return (
    <RubricsManager />
  )
}
```

**Key Methods Used:**
- `api.createRubric()`
- `api.getRubric()`
- `api.getRubricsByAssignment()`
- `api.updateRubric()`
- `api.deleteRubric()`
- `api.scoreSubmission()`
- `api.getSubmissionScore()`

---

### 6. Advanced Features Dashboard

**Purpose:** Main hub for all advanced features

**Features:**
- Navigation to all advanced features
- Quick stats/summaries
- Recent activity
- Feature overview cards

**Usage:**
```jsx
import AdvancedFeatures from './components/AdvancedFeatures'

export default function App() {
  return (
    <AdvancedFeatures />
  )
}
```

---

## Step-by-Step Integration

### Step 1: Create Components Directory Structure

```bash
mkdir -p frontend/src/components/AdvancedFeatures
```

### Step 2: Create Individual Components

Create files:
- `frontend/src/components/AdvancedFeatures/AdvancedSearch.jsx`
- `frontend/src/components/AdvancedFeatures/ReportsExport.jsx`
- `frontend/src/components/AdvancedFeatures/AttendanceAnalytics.jsx`
- `frontend/src/components/AdvancedFeatures/GradeAnalytics.jsx`
- `frontend/src/components/AdvancedFeatures/RubricsManager.jsx`
- `frontend/src/components/AdvancedFeatures/index.jsx`

### Step 3: Update App.jsx Routes

Add new routes to main navigation:

```jsx
import { BrowserRouter, Routes, Route, Navigate } from 'react-router-dom'
import AdvancedFeatures from './components/AdvancedFeatures'

export default function App() {
  return (
    <BrowserRouter>
      <Routes>
        {/* Existing routes */}
        
        {/* New Advanced Features Routes */}
        <Route path="/advanced/*" element={<AdvancedFeatures />} />
      </Routes>
    </BrowserRouter>
  )
}
```

### Step 4: Update Navigation Menu

Add link to advanced features in header/navigation:

```jsx
<nav>
  {/* Existing nav items */}
  <Link to="/advanced">Advanced Features</Link>
</nav>
```

### Step 5: Test Each Component

Use test data or existing database records:

```javascript
// Test search
const results = await api.searchStudents('John')

// Test export
const csvData = await api.exportGradesCSV(1)

// Test attendance
const stats = await api.getAttendanceStatsByCourse(1)
```

---

## Common Integration Patterns

### Pattern 1: Search with Filters

```jsx
const [searchType, setSearchType] = useState('announcements')
const [filters, setFilters] = useState({})
const [results, setResults] = useState([])
const [loading, setLoading] = useState(false)

const handleSearch = async () => {
  setLoading(true)
  try {
    let data
    if (searchType === 'announcements') {
      data = await api.searchAnnouncements(
        filters.query, 
        filters.audience, 
        filters.priority
      )
    } else if (searchType === 'students') {
      data = await api.searchStudents(filters.query)
    }
    // ... more search types
    setResults(data.data)
  } catch (error) {
    console.error('Search failed:', error)
  } finally {
    setLoading(false)
  }
}
```

### Pattern 2: Export CSV

```jsx
const handleExport = async (reportType, filters) => {
  try {
    let response
    if (reportType === 'payments') {
      response = await api.exportPaymentsCSV(filters.studentId, filters.status)
    } else if (reportType === 'grades') {
      response = await api.exportGradesCSV(filters.courseId)
    }
    
    // Create download link
    const blob = new Blob([response], { type: 'text/csv' })
    const url = window.URL.createObjectURL(blob)
    const a = document.createElement('a')
    a.href = url
    a.download = `report-${Date.now()}.csv`
    a.click()
  } catch (error) {
    console.error('Export failed:', error)
  }
}
```

### Pattern 3: Load and Display Analytics

```jsx
const [courseId, setCourseId] = useState(1)
const [stats, setStats] = useState(null)
const [loading, setLoading] = useState(false)

useEffect(() => {
  const loadStats = async () => {
    setLoading(true)
    try {
      const data = await api.getAttendanceStatsByCourse(courseId)
      setStats(data.data)
    } catch (error) {
      console.error('Failed to load stats:', error)
    } finally {
      setLoading(false)
    }
  }
  
  loadStats()
}, [courseId])
```

### Pattern 4: Form Submission with Validation

```jsx
const [formData, setFormData] = useState({
  submission_id: '',
  rubric_id: '',
  criterion_scores: [],
  feedback: '',
  graded_by: ''
})

const handleSubmit = async (e) => {
  e.preventDefault()
  
  // Validate
  if (!formData.submission_id || !formData.rubric_id) {
    alert('Please fill all required fields')
    return
  }
  
  try {
    const result = await api.scoreSubmission(formData)
    alert('Score recorded successfully!')
    // Reset form or redirect
  } catch (error) {
    alert('Error: ' + error.message)
  }
}
```

---

## Error Handling Best Practices

```jsx
const handleError = (error) => {
  if (error.status === 401) {
    // Redirect to login
    window.location.href = '/login'
  } else if (error.status === 403) {
    alert('You do not have permission to perform this action')
  } else if (error.status === 404) {
    alert('Resource not found')
  } else {
    alert('Error: ' + (error.message || 'Unknown error'))
  }
}

try {
  const data = await api.searchStudents(query)
} catch (error) {
  handleError(error)
}
```

---

## Loading States & UI Feedback

```jsx
const [loading, setLoading] = useState(false)
const [error, setError] = useState(null)
const [success, setSuccess] = useState(false)

// Show loading spinner
{loading && <LoadingSpinner />}

// Show error message
{error && <ErrorAlert message={error} />}

// Show success message
{success && <SuccessAlert message="Operation completed!" />}
```

---

## Performance Optimization

### 1. Pagination
```jsx
const [page, setPage] = useState(1)
const [limit, setLimit] = useState(10)

const results = await api.searchStudents(query, page, limit)
```

### 2. Debounce Search Input
```jsx
import { useMemo } from 'react'

const debouncedSearch = useMemo(() => {
  return debounce((query) => {
    api.searchStudents(query)
  }, 500)
}, [])
```

### 3. Caching
```jsx
const [cache, setCache] = useState({})

const getCachedData = async (key, fn) => {
  if (cache[key]) return cache[key]
  
  const data = await fn()
  setCache(prev => ({ ...prev, [key]: data }))
  return data
}
```

---

## Testing Checklist

- [ ] Search functionality works with all filter types
- [ ] Export downloads CSV files correctly
- [ ] Attendance stats display correctly
- [ ] Grade calculations are accurate
- [ ] Rubric creation and scoring works
- [ ] Error messages display appropriately
- [ ] Pagination works for search results
- [ ] Form validation prevents invalid submissions
- [ ] Loading states show/hide correctly
- [ ] Responsive design works on mobile
- [ ] Authentication required for all endpoints
- [ ] No console errors or warnings

---

## Deployment Notes

### Environment Variables

Create `.env` in frontend directory:
```
VITE_API_BASE=http://localhost:8080/api  # Development
VITE_API_BASE=https://api.school.com/api  # Production
```

### Build Optimization
```bash
npm run build
# Analyze bundle size
npm install -D rollup-plugin-visualizer
```

### Production Server
```bash
# Use backend server to serve frontend
# Backend serves static files from frontend/dist/
go run ./cmd/server
```

---

## Troubleshooting

### CORS Issues
- Ensure backend CORS middleware allows frontend origin
- Check backend `.env` or config

### API Methods Not Available
- Verify `api.js` includes all 26 new methods
- Check import statement: `import api from './api'`

### Authentication Errors
- Ensure JWT token is stored in localStorage
- Check token expiration
- Re-login if token invalid

### Data Not Loading
- Check browser console for errors
- Verify backend server is running
- Check network tab in DevTools
- Verify correct API endpoint paths

---

## Next Steps

1. ✅ API methods added to `api.js` (Done)
2. ⭕ Create React components (Next)
3. ⭕ Update App.jsx routes
4. ⭕ Test all components
5. ⭕ Deploy to production

---

## Support & Documentation

- **Backend API Docs:** [API_TESTING_SESSION_6.md](API_TESTING_SESSION_6.md)
- **Component Props:** See individual component files
- **API Reference:** Check `frontend/src/api.js`
- **Example Usage:** See pattern examples above

