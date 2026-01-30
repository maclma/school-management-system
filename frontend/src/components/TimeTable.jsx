import React, { useState, useEffect } from 'react'
import api from '../api'

export default function TimeTable() {
  const [timetable, setTimetable] = useState([])
  const [loading, setLoading] = useState(false)
  const [filterType, setFilterType] = useState('all')
  const [filterId, setFilterId] = useState('')
  const [selectedDay, setSelectedDay] = useState('Monday')

  const days = ['Monday', 'Tuesday', 'Wednesday', 'Thursday', 'Friday', 'Saturday', 'Sunday']

  useEffect(() => {
    loadTimetable()
  }, [])

  const loadTimetable = async () => {
    setLoading(true)
    try {
      const res = await api.getTimetable()
      setTimetable(res.data || [])
    } catch (err) {
      console.error('Failed to load timetable:', err)
    }
    setLoading(false)
  }

  const handleFilter = async () => {
    setLoading(true)
    try {
      let res
      if (filterType === 'course' && filterId) {
        res = await api.getTimetableByCourse(parseInt(filterId))
      } else if (filterType === 'teacher' && filterId) {
        res = await api.getTimetableByTeacher(parseInt(filterId))
      } else if (filterType === 'day') {
        res = await api.getTimetableByDay(selectedDay)
      } else {
        res = await api.getTimetable()
      }
      setTimetable(res.data || [])
    } catch (err) {
      console.error('Failed to filter timetable:', err)
      alert('Failed to filter: ' + err.message)
    }
    setLoading(false)
  }

  const getFilteredSchedule = () => {
    if (filterType === 'day') {
      return timetable.filter(t => t.day_of_week === selectedDay)
    }
    return timetable
  }

  const groupByDay = (schedule) => {
    const grouped = {}
    schedule.forEach(item => {
      if (!grouped[item.day_of_week]) {
        grouped[item.day_of_week] = []
      }
      grouped[item.day_of_week].push(item)
    })
    return grouped
  }

  const filteredSchedule = getFilteredSchedule()
  const groupedByDay = groupByDay(filteredSchedule)

  return (
    <div className="timetable-container">
      <h2>Class Schedule / TimeTable</h2>

      <div className="timetable-filters">
        <div className="filter-group">
          <label>Filter By:</label>
          <select value={filterType} onChange={(e) => setFilterType(e.target.value)}>
            <option value="all">All Classes</option>
            <option value="course">Course</option>
            <option value="teacher">Teacher</option>
            <option value="day">Day</option>
          </select>
        </div>

        {filterType === 'course' && (
          <div className="filter-group">
            <input
              type="number"
              placeholder="Course ID"
              value={filterId}
              onChange={(e) => setFilterId(e.target.value)}
            />
          </div>
        )}

        {filterType === 'teacher' && (
          <div className="filter-group">
            <input
              type="number"
              placeholder="Teacher ID"
              value={filterId}
              onChange={(e) => setFilterId(e.target.value)}
            />
          </div>
        )}

        {filterType === 'day' && (
          <div className="filter-group">
            <select value={selectedDay} onChange={(e) => setSelectedDay(e.target.value)}>
              {days.map(day => (
                <option key={day} value={day}>{day}</option>
              ))}
            </select>
          </div>
        )}

        <button className="btn btn-primary" onClick={handleFilter}>
          Apply Filter
        </button>
      </div>

      {loading ? (
        <p>Loading schedule...</p>
      ) : filteredSchedule.length === 0 ? (
        <p className="no-data">No classes found</p>
      ) : filterType === 'day' ? (
        <div className="timetable-view">
          <h3>{selectedDay} Classes</h3>
          <div className="day-schedule">
            {filteredSchedule.map((entry) => (
              <div key={entry.id} className="class-card">
                <div className="class-time">
                  <strong>{entry.start_time} - {entry.end_time}</strong>
                </div>
                <div className="class-details">
                  <p><strong>Course ID:</strong> {entry.course_id}</p>
                  <p><strong>Teacher ID:</strong> {entry.teacher_id}</p>
                  <p><strong>Classroom:</strong> {entry.classroom}</p>
                  <p><strong>Location:</strong> {entry.location}</p>
                </div>
              </div>
            ))}
          </div>
        </div>
      ) : (
        <div className="timetable-weekly">
          {days.map(day => (
            groupedByDay[day] && groupedByDay[day].length > 0 && (
              <div key={day} className="day-section">
                <h3>{day}</h3>
                {groupedByDay[day].map((entry) => (
                  <div key={entry.id} className="class-entry">
                    <span className="time-slot">{entry.start_time} - {entry.end_time}</span>
                    <span className="course">Course {entry.course_id}</span>
                    <span className="teacher">Teacher {entry.teacher_id}</span>
                    <span className="room">{entry.classroom} ({entry.location})</span>
                  </div>
                ))}
              </div>
            )
          ))}
        </div>
      )}
    </div>
  )
}
