package service

import (
	"fmt"
	"time"

	"school-management-system/internal/models"
	"school-management-system/pkg/database"
)

// GradeAutoCalculationService handles automatic grade calculations
type GradeAutoCalculationService struct {
	gradeTranscriptService GradeTranscriptService
	emailService           *EmailService
}

// NewGradeAutoCalculationService creates a new service
func NewGradeAutoCalculationService(
	gradeTranscriptService GradeTranscriptService,
	emailService *EmailService,
) *GradeAutoCalculationService {
	return &GradeAutoCalculationService{
		gradeTranscriptService: gradeTranscriptService,
		emailService:           emailService,
	}
}

// CalculateLetterGrade converts numeric score to letter grade
func (gacs *GradeAutoCalculationService) CalculateLetterGrade(score float64) string {
	if score >= 90 {
		return "A"
	} else if score >= 80 {
		return "B"
	} else if score >= 70 {
		return "C"
	} else if score >= 60 {
		return "D"
	}
	return "F"
}

// CalculateGradePoints converts letter grade to grade points
func (gacs *GradeAutoCalculationService) CalculateGradePoints(letterGrade string) float64 {
	gradePoints := map[string]float64{
		"A": 4.0,
		"B": 3.0,
		"C": 2.0,
		"D": 1.0,
		"F": 0.0,
	}
	return gradePoints[letterGrade]
}

// RecordGradeAndAutoCalculate records grade and auto-calculates letter grade
func (gacs *GradeAutoCalculationService) RecordGradeAndAutoCalculate(grade *models.Grade) error {
	db := database.DB

	// Auto-calculate letter grade
	grade.Grade = gacs.CalculateLetterGrade(grade.Score)

	// Save grade
	if err := db.Create(grade).Error; err != nil {
		return err
	}

	// Trigger automatic actions
	go func() {
		// Get student info for email
		var student models.Student
		if err := db.Preload("User").First(&student, grade.StudentID).Error; err == nil {
			// Get course info
			var course models.Course
			if err := db.First(&course, grade.CourseID).Error; err == nil {
				// Send grade notification email
				gacs.emailService.SendGradeNotification(
					student.User.Email,
					student.User.FirstName+" "+student.User.LastName,
					course.Name,
					fmt.Sprintf("%s (%.1f%%)", grade.Grade, grade.Score),
				)
			}
		}

		// Update transcript
		gacs.updateStudentTranscript(grade.StudentID)
	}()

	return nil
}

// updateStudentTranscript updates student's grade transcript
func (gacs *GradeAutoCalculationService) updateStudentTranscript(studentID uint) error {
	db := database.DB

	// Get all grades for student in current term
	var grades []models.Grade
	if err := db.Where("student_id = ?", studentID).Find(&grades).Error; err != nil {
		return err
	}

	if len(grades) == 0 {
		return nil
	}

	// Calculate GPA
	var totalGradePoints float64
	var totalCredits float64

	for _, grade := range grades {
		// Get course for credits
		var course models.Course
		if err := db.First(&course, grade.CourseID).Error; err == nil {
			gradePoints := gacs.CalculateGradePoints(grade.Grade)
			totalGradePoints += gradePoints * float64(course.CreditHours)
			totalCredits += float64(course.CreditHours)
		}
	}

	var gpa float64
	if totalCredits > 0 {
		gpa = totalGradePoints / totalCredits
	}

	// Update or create transcript
	var transcript models.GradeTranscript
	// Determine current semester and year
	now := time.Now()
	year := now.Year()
	semester := "Fall"
	if now.Month() <= 6 {
		semester = "Spring"
	}

	if err := db.Where("student_id = ? AND transcript_semester = ? AND year = ?",
		studentID, semester, year).First(&transcript).Error; err != nil {
		// Create new transcript
		transcript = models.GradeTranscript{
			StudentID:          studentID,
			TranscriptSemester: semester,
			Year:               year,
			GPA:                gpa,
			TotalCredits:       totalCredits,
			EarnedCredits:      totalCredits,
			GradePointsSum:     totalGradePoints,
			GeneratedAt:        time.Now().Unix(),
		}
		db.Create(&transcript)
	} else {
		// Update existing transcript
		transcript.GPA = gpa
		transcript.TotalCredits = totalCredits
		transcript.EarnedCredits = totalCredits
		transcript.GradePointsSum = totalGradePoints
		transcript.GeneratedAt = time.Now().Unix()
		db.Save(&transcript)
	}

	return nil
}

// CalculateCourseAverage calculates average grade for a course
func (gacs *GradeAutoCalculationService) CalculateCourseAverage(courseID uint) (float64, error) {
	db := database.DB

	var avg float64
	if err := db.Model(&models.Grade{}).
		Where("course_id = ?", courseID).
		Select("AVG(score)").
		Row().
		Scan(&avg); err != nil {
		return 0, err
	}

	return avg, nil
}

// CalculateClassGradeDistribution calculates grade distribution for a class
func (gacs *GradeAutoCalculationService) CalculateClassGradeDistribution(courseID uint) (map[string]int, error) {
	db := database.DB

	var grades []models.Grade
	if err := db.Where("course_id = ?", courseID).Find(&grades).Error; err != nil {
		return nil, err
	}

	distribution := map[string]int{
		"A": 0,
		"B": 0,
		"C": 0,
		"D": 0,
		"F": 0,
	}

	for _, grade := range grades {
		distribution[grade.Grade]++
	}

	return distribution, nil
}

// GetStudentGradeStats returns comprehensive grade statistics for a student
func (gacs *GradeAutoCalculationService) GetStudentGradeStats(studentID uint) (map[string]interface{}, error) {
	db := database.DB

	var grades []models.Grade
	if err := db.Where("student_id = ?", studentID).Find(&grades).Error; err != nil {
		return nil, err
	}

	if len(grades) == 0 {
		return map[string]interface{}{
			"student_id":  studentID,
			"grade_count": 0,
			"average":     0,
		}, nil
	}

	var totalScore float64
	var aCount, bCount, cCount, dCount, fCount int

	for _, grade := range grades {
		totalScore += grade.Score
		switch grade.Grade {
		case "A":
			aCount++
		case "B":
			bCount++
		case "C":
			cCount++
		case "D":
			dCount++
		case "F":
			fCount++
		}
	}

	avg := totalScore / float64(len(grades))

	return map[string]interface{}{
		"student_id":    studentID,
		"grade_count":   len(grades),
		"average":       fmt.Sprintf("%.2f", avg),
		"a_count":       aCount,
		"b_count":       bCount,
		"c_count":       cCount,
		"d_count":       dCount,
		"f_count":       fCount,
		"highest_grade": getMaxGrade(grades),
		"lowest_grade":  getMinGrade(grades),
	}, nil
}

// Helper functions
func getMaxGrade(grades []models.Grade) float64 {
	if len(grades) == 0 {
		return 0
	}
	max := grades[0].Score
	for _, g := range grades {
		if g.Score > max {
			max = g.Score
		}
	}
	return max
}

func getMinGrade(grades []models.Grade) float64 {
	if len(grades) == 0 {
		return 0
	}
	min := grades[0].Score
	for _, g := range grades {
		if g.Score < min {
			min = g.Score
		}
	}
	return min
}
