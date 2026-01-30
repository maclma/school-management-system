package service

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"time"

	"school-management-system/internal/models"

	"gorm.io/gorm"
)

// ExportService handles CSV and report generation
type ExportService struct {
	db *gorm.DB
}

// NewExportService creates a new export service
func NewExportService(db *gorm.DB) *ExportService {
	return &ExportService{db: db}
}

// ExportPaymentsCSV generates CSV of payments
func (es *ExportService) ExportPaymentsCSV(studentID uint, status string) ([]byte, error) {
	db := es.db
	var payments []models.Payment

	q := db.Where("1=1")
	if studentID > 0 {
		q = q.Where("student_id = ?", studentID)
	}
	if status != "" {
		q = q.Where("status = ?", status)
	}

	if err := q.Find(&payments).Error; err != nil {
		return nil, err
	}

	var b bytes.Buffer
	w := csv.NewWriter(&b)

	// Write header
	w.Write([]string{"ID", "Student ID", "Amount", "Status", "Description", "Due Date", "Created"})

	// Write data
	for _, p := range payments {
		w.Write([]string{
			fmt.Sprintf("%d", p.ID),
			fmt.Sprintf("%d", p.StudentID),
			fmt.Sprintf("%.2f", p.Amount),
			p.Status,
			p.Description,
			fmt.Sprintf("%d", p.DueDate),
			fmt.Sprintf("%d", p.CreatedAt),
		})
	}

	w.Flush()
	if err := w.Error(); err != nil {
		return nil, err
	}

	return b.Bytes(), nil
}

// ExportGradesCSV generates CSV of grades
func (es *ExportService) ExportGradesCSV(courseID uint) ([]byte, error) {
	db := es.db
	var grades []models.Grade

	if err := db.Where("course_id = ?", courseID).Find(&grades).Error; err != nil {
		return nil, err
	}

	var b bytes.Buffer
	w := csv.NewWriter(&b)

	// Write header
	w.Write([]string{"ID", "Student ID", "Course ID", "Score", "Grade", "Graded At"})

	// Write data
	for _, g := range grades {
		w.Write([]string{
			fmt.Sprintf("%d", g.ID),
			fmt.Sprintf("%d", g.StudentID),
			fmt.Sprintf("%d", g.CourseID),
			fmt.Sprintf("%.2f", g.Score),
			g.Grade,
			g.GradedAt.Format(time.RFC3339),
		})
	}

	w.Flush()
	if err := w.Error(); err != nil {
		return nil, err
	}

	return b.Bytes(), nil
}

// ExportAttendanceCSV generates CSV of attendance records
func (es *ExportService) ExportAttendanceCSV(courseID uint) ([]byte, error) {
	db := es.db
	var attendance []models.Attendance

	if err := db.Where("course_id = ?", courseID).Find(&attendance).Error; err != nil {
		return nil, err
	}

	var b bytes.Buffer
	w := csv.NewWriter(&b)

	// Write header
	w.Write([]string{"ID", "Student ID", "Course ID", "Date", "Present", "Remarks"})

	// Write data
	for _, a := range attendance {
		present := "No"
		if a.Status == "present" {
			present = "Yes"
		}
		w.Write([]string{
			fmt.Sprintf("%d", a.ID),
			fmt.Sprintf("%d", a.StudentID),
			fmt.Sprintf("%d", a.CourseID),
			a.Date.Format(time.RFC3339),
			present,
			a.Remarks,
		})
	}

	w.Flush()
	if err := w.Error(); err != nil {
		return nil, err
	}

	return b.Bytes(), nil
}

// ExportStudentTranscriptCSV generates student academic transcript
func (es *ExportService) ExportStudentTranscriptCSV(studentID uint) ([]byte, error) {
	db := es.db
	var transcripts []models.GradeTranscript

	if err := db.Where("student_id = ?", studentID).Order("academic_year DESC, term DESC").Find(&transcripts).Error; err != nil {
		return nil, err
	}

	var b bytes.Buffer
	w := csv.NewWriter(&b)

	// Write header
	w.Write([]string{"Semester", "Year", "GPA", "Total Credits", "Earned Credits", "Generated At"})

	// Write data
	for _, t := range transcripts {
		w.Write([]string{
			t.TranscriptSemester,
			fmt.Sprintf("%d", t.Year),
			fmt.Sprintf("%.2f", t.GPA),
			fmt.Sprintf("%.2f", t.TotalCredits),
			fmt.Sprintf("%.2f", t.EarnedCredits),
			fmt.Sprintf("%d", t.GeneratedAt),
		})
	}

	w.Flush()
	if err := w.Error(); err != nil {
		return nil, err
	}

	return b.Bytes(), nil
}

// ExportEnrollmentsCSV generates CSV of course enrollments
func (es *ExportService) ExportEnrollmentsCSV(courseID uint) ([]byte, error) {
	db := es.db
	var enrollments []models.Enrollment

	if err := db.Where("course_id = ?", courseID).Find(&enrollments).Error; err != nil {
		return nil, err
	}

	var b bytes.Buffer
	w := csv.NewWriter(&b)

	// Write header
	w.Write([]string{"ID", "Student ID", "Course ID", "Status", "Enrolled At"})

	// Write data
	for _, e := range enrollments {
		w.Write([]string{
			fmt.Sprintf("%d", e.ID),
			fmt.Sprintf("%d", e.StudentID),
			fmt.Sprintf("%d", e.CourseID),
			e.Status,
			e.EnrolledAt.Format(time.RFC3339),
		})
	}

	w.Flush()
	if err := w.Error(); err != nil {
		return nil, err
	}

	return b.Bytes(), nil
}
