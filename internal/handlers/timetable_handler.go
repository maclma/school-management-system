package handlers

import (
	"school-management-system/internal/models"
	"school-management-system/internal/service"
	"school-management-system/pkg/errors"
	"school-management-system/pkg/response"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TimeTableHandler struct {
	service service.TimeTableService
}

func NewTimeTableHandler(svc service.TimeTableService) *TimeTableHandler {
	return &TimeTableHandler{service: svc}
}

func (h *TimeTableHandler) Create(c *gin.Context) {
	var req struct {
		CourseID  uint   `json:"course_id" binding:"required"`
		TeacherID uint   `json:"teacher_id" binding:"required"`
		DayOfWeek string `json:"day_of_week" binding:"required"`
		StartTime string `json:"start_time" binding:"required"`
		EndTime   string `json:"end_time" binding:"required"`
		Classroom string `json:"classroom"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	timetable := &models.TimeTable{
		CourseID:  req.CourseID,
		TeacherID: req.TeacherID,
		DayOfWeek: req.DayOfWeek,
		StartTime: req.StartTime,
		EndTime:   req.EndTime,
		Classroom: req.Classroom,
		IsActive:  true,
	}

	if err := h.service.Create(timetable); err != nil {
		response.Error(c, errors.InternalError("Failed to create timetable"))
		return
	}
	response.Created(c, "Timetable created", timetable)
}

func (h *TimeTableHandler) GetByCourseID(c *gin.Context) {
	courseID, _ := strconv.ParseUint(c.Param("course_id"), 10, 32)
	timetables, err := h.service.GetByCourseID(uint(courseID))
	if err != nil {
		response.Error(c, errors.InternalError("Failed to fetch timetable"))
		return
	}
	response.Success(c, "Timetable fetched", timetables)
}

func (h *TimeTableHandler) GetByTeacherID(c *gin.Context) {
	teacherID, _ := strconv.ParseUint(c.Param("teacher_id"), 10, 32)
	timetables, err := h.service.GetByTeacherID(uint(teacherID))
	if err != nil {
		response.Error(c, errors.InternalError("Failed to fetch timetable"))
		return
	}
	response.Success(c, "Timetable fetched", timetables)
}

func (h *TimeTableHandler) GetByDay(c *gin.Context) {
	day := c.Param("day")
	timetables, err := h.service.GetByDayOfWeek(day)
	if err != nil {
		response.Error(c, errors.InternalError("Failed to fetch timetable"))
		return
	}
	response.Success(c, "Timetable fetched", timetables)
}

func (h *TimeTableHandler) GetAll(c *gin.Context) {
	timetables, err := h.service.GetAll()
	if err != nil {
		response.Error(c, errors.InternalError("Failed to fetch timetable"))
		return
	}
	response.Success(c, "Timetable fetched", timetables)
}

func (h *TimeTableHandler) Update(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	var req struct {
		DayOfWeek string `json:"day_of_week"`
		StartTime string `json:"start_time"`
		EndTime   string `json:"end_time"`
		Classroom string `json:"classroom"`
		IsActive  bool   `json:"is_active"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	timetable, _ := h.service.GetByID(uint(id))
	if req.DayOfWeek != "" {
		timetable.DayOfWeek = req.DayOfWeek
	}
	if req.StartTime != "" {
		timetable.StartTime = req.StartTime
	}
	if req.EndTime != "" {
		timetable.EndTime = req.EndTime
	}
	if req.Classroom != "" {
		timetable.Classroom = req.Classroom
	}
	timetable.IsActive = req.IsActive

	if err := h.service.Update(timetable); err != nil {
		response.Error(c, errors.InternalError("Failed to update timetable"))
		return
	}
	response.Success(c, "Timetable updated", timetable)
}

func (h *TimeTableHandler) Delete(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	if err := h.service.Delete(uint(id)); err != nil {
		response.Error(c, errors.InternalError("Failed to delete timetable"))
		return
	}
	response.NoContent(c)
}
