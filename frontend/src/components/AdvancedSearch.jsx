import { useState } from 'react'
import api from '../../api'
import './AdvancedSearch.css'

export default function AdvancedSearch() {
  const [searchType, setSearchType] = useState('announcements')
  const [filters, setFilters] = useState({})
  const [results, setResults] = useState([])
  const [loading, setLoading] = useState(false)
  const [error, setError] = useState(null)
  const [page, setPage] = useState(1)

  const handleFilterChange = (e) => {
    const { name, value } = e.target
    setFilters(prev => ({ ...prev, [name]: value }))
  }

  const handleSearch = async (e) => {
    e.preventDefault()
    setLoading(true)
    setError(null)

    try {
      let data
      switch (searchType) {
        case 'announcements':
          data = await api.searchAnnouncements(
            filters.query || '',
            filters.audience || '',
            filters.priority || '',
            page
          )
          break
        case 'payments':
          data = await api.searchPayments(
            filters.studentId || 0,
            filters.status || '',
            page
          )
          break
        case 'students':
          data = await api.searchStudents(
            filters.query || '',
            page
          )
          break
        case 'grades':
          data = await api.searchGradesByRange(
            filters.courseId || 0,
            filters.minScore || 0,
            filters.maxScore || 100,
            page
          )
          break
        case 'overdue':
          data = await api.searchOverduePayments()
          break
        default:
          throw new Error('Invalid search type')
      }
      setResults(data.data || [])
    } catch (err) {
      setError(err.message || 'Search failed')
    } finally {
      setLoading(false)
    }
  }

  return (
    <div className="advanced-search-container">
      <h2>Advanced Search</h2>
      
      <form onSubmit={handleSearch} className="search-form">
        <div className="search-type-selector">
          <label>Search Type:</label>
          <select value={searchType} onChange={(e) => setSearchType(e.target.value)}>
            <option value="announcements">Announcements</option>
            <option value="payments">Payments</option>
            <option value="students">Students</option>
            <option value="grades">Grades</option>
            <option value="overdue">Overdue Payments</option>
          </select>
        </div>

        {/* Dynamic Filters based on search type */}
        <div className="filters">
          {(searchType === 'announcements' || searchType === 'students') && (
            <input
              type="text"
              name="query"
              placeholder="Search query..."
              value={filters.query || ''}
              onChange={handleFilterChange}
            />
          )}

          {searchType === 'announcements' && (
            <>
              <select name="audience" value={filters.audience || ''} onChange={handleFilterChange}>
                <option value="">All Audiences</option>
                <option value="students">Students</option>
                <option value="teachers">Teachers</option>
                <option value="all">All</option>
              </select>
              <select name="priority" value={filters.priority || ''} onChange={handleFilterChange}>
                <option value="">All Priorities</option>
                <option value="high">High</option>
                <option value="medium">Medium</option>
                <option value="low">Low</option>
              </select>
            </>
          )}

          {searchType === 'payments' && (
            <>
              <input
                type="number"
                name="studentId"
                placeholder="Student ID (optional)"
                value={filters.studentId || ''}
                onChange={handleFilterChange}
              />
              <select name="status" value={filters.status || ''} onChange={handleFilterChange}>
                <option value="">All Statuses</option>
                <option value="pending">Pending</option>
                <option value="paid">Paid</option>
                <option value="overdue">Overdue</option>
                <option value="cancelled">Cancelled</option>
              </select>
            </>
          )}

          {searchType === 'grades' && (
            <>
              <input
                type="number"
                name="courseId"
                placeholder="Course ID"
                required
                value={filters.courseId || ''}
                onChange={handleFilterChange}
              />
              <input
                type="number"
                name="minScore"
                placeholder="Min Score"
                value={filters.minScore || '0'}
                onChange={handleFilterChange}
              />
              <input
                type="number"
                name="maxScore"
                placeholder="Max Score"
                value={filters.maxScore || '100'}
                onChange={handleFilterChange}
              />
            </>
          )}
        </div>

        <button type="submit" disabled={loading}>
          {loading ? 'Searching...' : 'Search'}
        </button>
      </form>

      {error && <div className="error-message">{error}</div>}

      {results.length > 0 && (
        <div className="results">
          <h3>Results ({results.length})</h3>
          <div className="results-list">
            {searchType === 'announcements' && results.map(item => (
              <div key={item.id} className="result-card">
                <h4>{item.title}</h4>
                <p>{item.content}</p>
                <span className="badge">{item.priority}</span>
              </div>
            ))}

            {searchType === 'payments' && results.map(item => (
              <div key={item.id} className="result-card">
                <h4>Payment #{item.id}</h4>
                <p>Amount: ${item.amount.toFixed(2)}</p>
                <p>Status: <span className={`status-${item.status}`}>{item.status}</span></p>
              </div>
            ))}

            {searchType === 'students' && results.map(item => (
              <div key={item.id} className="result-card">
                <h4>{item.user?.first_name} {item.user?.last_name}</h4>
                <p>ID: {item.student_id}</p>
                <p>Email: {item.user?.email}</p>
              </div>
            ))}

            {searchType === 'grades' && results.map(item => (
              <div key={item.id} className="result-card">
                <h4>Grade #{item.id}</h4>
                <p>Score: {item.score.toFixed(2)}</p>
                <p>Grade: <strong>{item.grade}</strong></p>
              </div>
            ))}
          </div>

          <div className="pagination">
            <button 
              onClick={() => setPage(Math.max(1, page - 1))}
              disabled={page === 1}
            >
              Previous
            </button>
            <span>Page {page}</span>
            <button onClick={() => setPage(page + 1)}>
              Next
            </button>
          </div>
        </div>
      )}

      {!loading && results.length === 0 && !error && (
        <div className="no-results">No results found</div>
      )}
    </div>
  )
}
