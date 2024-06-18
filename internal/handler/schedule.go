package handler

import (
	"net/http"
	"strconv"
	"time"

	"github.com/BeRebornBng/OsauAmsApi/domain"
	"github.com/gin-gonic/gin"
)

// Error messages constants
const (
	ErrInvalidRequestBody = "Invalid request body"
	ErrInvalidScheduleID  = "Invalid schedule ID"
	ErrScheduleNotFound   = "Schedule not found"
	ErrSchedulesNotFound  = "Schedules not found"
	ErrInvalidGroupID     = "Invalid group ID"
	ErrInvalidWeekType    = "Invalid week type"
	ErrInvalidDayOfWeek   = "Invalid day of the week"
	ErrInvalid
)

// CreateScheduleRequest represents the request body for creating a schedule
type CreateScheduleRequest struct {
	GroupID          string `json:"group_id" validate:"required,customgroupidregex"`
	DisciplineID     int64  `json:"discipline_id" validate:"required,numeric"`
	TeacherID        int64  `json:"teacher_id" validate:"required,numeric"`
	DisciplineTypeID int64  `json:"discipline_type_id" validate:"required,numeric"`
	ClassroomID      int64  `json:"classroom_id" validate:"required,numeric"`
	Semester         int    `json:"semester" validate:"required,numeric"`
	BeginStudies     string `json:"begin_studies" validate:"required,datetime=2006-01-02"`
	WeekType         string `json:"week_type" validate:"required,oneof=Верхняя Нижняя"`
	DayOfWeek        string `json:"day_of_week" validate:"required,oneof=Понедельник Вторник Среда Четверг Пятница Суббота Воскресенье"`
	StartTime        string `json:"start_time" validate:"required,time"`
	IsActual         *bool  `json:"is_actual" validate:"required,boolean"`
}

// PutScheduleRequest represents the request body for updating a schedule
type PutScheduleRequest struct {
	ScheduleID       int64  `json:"schedule_id" validate:"required,numeric"`
	GroupID          string `json:"group_id" validate:"required,customgroupidregex"`
	DisciplineID     int64  `json:"discipline_id" validate:"required,numeric"`
	TeacherID        int64  `json:"teacher_id" validate:"required,numeric"`
	DisciplineTypeID int64  `json:"discipline_type_id" validate:"required,numeric"`
	ClassroomID      int64  `json:"classroom_id" validate:"required,numeric"`
	Semester         int    `json:"semester" validate:"required,numeric"`
	BeginStudies     string `json:"begin_studies" validate:"required,datetime=2006-01-02"`
	WeekType         string `json:"week_type" validate:"required,oneof=Верхняя Нижняя"`
	DayOfWeek        string `json:"day_of_week" validate:"required,oneof=Понедельник Вторник Среда Четверг Пятница Суббота Воскресенье"`
	StartTime        string `json:"start_time" validate:"required,time"`
	IsActual         *bool  `json:"is_actual" validate:"required,boolean"`
}

// PatchScheduleRequest represents the request body for partially updating a schedule
type PatchScheduleRequest struct {
	ScheduleID       int64  `json:"schedule_id" validate:"required,numeric"`
	GroupID          string `json:"group_id" validate:"omitempty,customgroupidregex"`
	DisciplineID     int64  `json:"discipline_id" validate:"omitempty,numeric"`
	TeacherID        int64  `json:"teacher_id" validate:"omitempty,numeric"`
	DisciplineTypeID int64  `json:"discipline_type_id" validate:"omitempty,numeric"`
	ClassroomID      int64  `json:"classroom_id" validate:"omitempty,numeric"`
	Semester         int    `json:"semester" validate:"omitempty,numeric"`
	BeginStudies     string `json:"begin_studies" validate:"omitempty,datetime=2006-01-02"`
	WeekType         string `json:"week_type" validate:"omitempty,oneof=Верхняя Нижняя"`
	DayOfWeek        string `json:"day_of_week" validate:"omitempty,oneof=Понедельник Вторник Среда Четверг Пятница Суббота Воскресенье"`
	StartTime        string `json:"start_time" validate:"omitempty,time"`
	IsActual         *bool  `json:"is_actual" validate:"omitempty,boolean"`
}

// GetActualSchedulesByGroupAndWeekType represents the request body
type GetASByGWRequest struct {
	GroupID  string `json:"group_id" validate:"required,customgroupidregex"`
	WeekType string `json:"week_type" validate:"required,oneof=Верхняя Нижняя"`
}

// GetActualSchedulesByGroupAndWeekType represents the request body
type GetASByGWDRequest struct {
	GroupID   string `json:"group_id" validate:"required,customgroupidregex"`
	WeekType  string `json:"week_type" validate:"required,oneof=Верхняя Нижняя"`
	DayOfWeek string `json:"day_of_week" validate:"omitempty,oneof=Понедельник Вторник Среда Четверг Пятница Суббота Воскресенье"`
}

// GetActualSchedulesByGroupAndWeekType represents the request body
type GetASByTWRequest struct {
	TeacherID int64  `json:"teacher_id" validate:"required,numeric"`
	WeekType  string `json:"week_type" validate:"required,oneof=Верхняя Нижняя"`
}

// GetActualSchedulesByGroupAndWeekType represents the request body
type GetASByTWDRequest struct {
	TeacherID int64  `json:"teacher_id" validate:"required,numeric"`
	WeekType  string `json:"week_type" validate:"required,oneof=Верхняя Нижняя"`
	DayOfWeek string `json:"day_of_week" validate:"omitempty,oneof=Понедельник Вторник Среда Четверг Пятница Суббота Воскресенье"`
}

// CreateSchedule godoc
// @Summary Create a schedule
// @Description Create a new schedule
// @Tags Schedules
// @Accept json
// @Produce json
// @Param schedule body CreateScheduleRequest true "Schedule info"
// @Success 201 {object} SuccessResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /schedules [post]
func (h *Handler) CreateSchedule(c *gin.Context) {
	var req CreateScheduleRequest
	if err := c.BindJSON(&req); err != nil {
		respondWithError(h.logger, c, http.StatusBadRequest, ErrInvalidRequestBody)
		return
	}

	if err := h.validate.Struct(req); err != nil {
		respondWithError(h.logger, c, http.StatusBadRequest, ErrValidationFailed)
		return
	}

	date, err := time.Parse("2006-01-02", req.BeginStudies)
	if err != nil {
		respondWithError(h.logger, c, http.StatusBadRequest, ErrInvalidRequestBody)
		return
	}

	startTime, err := time.Parse("15:04", req.StartTime)
	if err != nil {
		respondWithError(h.logger, c, http.StatusBadRequest, ErrInvalidRequestBody)
		return
	}

	schedule := domain.Schedule{
		GroupID:          req.GroupID,
		DisciplineID:     req.DisciplineID,
		TeacherID:        req.TeacherID,
		DisciplineTypeID: req.DisciplineTypeID,
		ClassroomID:      req.ClassroomID,
		Semester:         req.Semester,
		BeginStudies:     date,
		WeekType:         req.WeekType,
		DayOfWeek:        req.DayOfWeek,
		StartTime:        startTime,
		IsActual:         req.IsActual,
	}

	err = h.services.ScheduleService.Create(c.Request.Context(), schedule)
	if err != nil {
		respondWithError(h.logger, c, http.StatusInternalServerError, ErrInternalServerError)
		return
	}

	c.JSON(http.StatusCreated, SuccessResponse{Message: "Schedule created successfully"})
}

// PutSchedule godoc
// @Summary Update a schedule
// @Description Update an existing schedule
// @Tags Schedules
// @Accept json
// @Produce json
// @Param schedule body PutScheduleRequest true "Schedule info"
// @Success 200 {object} SuccessResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /schedules [put]
func (h *Handler) PutSchedule(c *gin.Context) {
	var req PutScheduleRequest
	if err := c.BindJSON(&req); err != nil {
		respondWithError(h.logger, c, http.StatusBadRequest, ErrInvalidRequestBody)
		return
	}

	if err := h.validate.Struct(req); err != nil {
		respondWithError(h.logger, c, http.StatusBadRequest, ErrValidationFailed)
		return
	}

	date, err := time.Parse("2006-01-02", req.BeginStudies)
	if err != nil {
		respondWithError(h.logger, c, http.StatusBadRequest, ErrInvalidRequestBody)
		return
	}

	startTime, err := time.Parse("15:04", req.StartTime)
	if err != nil {
		respondWithError(h.logger, c, http.StatusBadRequest, ErrInvalidRequestBody)
		return
	}

	schedule := domain.Schedule{
		ScheduleID:       req.ScheduleID,
		GroupID:          req.GroupID,
		DisciplineID:     req.DisciplineID,
		TeacherID:        req.TeacherID,
		DisciplineTypeID: req.DisciplineTypeID,
		ClassroomID:      req.ClassroomID,
		Semester:         req.Semester,
		BeginStudies:     date,
		WeekType:         req.WeekType,
		DayOfWeek:        req.DayOfWeek,
		StartTime:        startTime,
		IsActual:         req.IsActual,
	}

	err = h.services.ScheduleService.Put(c.Request.Context(), schedule)
	if err != nil {
		respondWithError(h.logger, c, http.StatusInternalServerError, ErrInternalServerError)
		return
	}

	c.JSON(http.StatusOK, SuccessResponse{Message: "Schedule updated successfully"})
}

// PatchSchedule godoc
// @Summary Partially update a schedule
// @Description Partially update an existing schedule
// @Tags Schedules
// @Accept json
// @Produce json
// @Param schedule body PatchScheduleRequest true "Schedule info"
// @Success 200 {object} SuccessResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /schedules [patch]
func (h *Handler) PatchSchedule(c *gin.Context) {
	var req PatchScheduleRequest
	if err := c.BindJSON(&req); err != nil {
		respondWithError(h.logger, c, http.StatusBadRequest, ErrInvalidRequestBody)
		return
	}

	if err := h.validate.Struct(req); err != nil {
		respondWithError(h.logger, c, http.StatusBadRequest, ErrValidationFailed)
		return
	}

	date, err := time.Parse("2006-01-02", req.BeginStudies)
	if err != nil {
		respondWithError(h.logger, c, http.StatusBadRequest, ErrInvalidRequestBody)
		return
	}

	startTime, err := time.Parse("15:04", req.StartTime)
	if err != nil {
		respondWithError(h.logger, c, http.StatusBadRequest, ErrInvalidRequestBody)
		return
	}

	schedule := domain.Schedule{
		ScheduleID:       req.ScheduleID,
		GroupID:          req.GroupID,
		DisciplineID:     req.DisciplineID,
		TeacherID:        req.TeacherID,
		DisciplineTypeID: req.DisciplineTypeID,
		ClassroomID:      req.ClassroomID,
		Semester:         req.Semester,
		BeginStudies:     date,
		WeekType:         req.WeekType,
		DayOfWeek:        req.DayOfWeek,
		StartTime:        startTime,
		IsActual:         req.IsActual,
	}

	err = h.services.ScheduleService.Patch(c.Request.Context(), schedule)
	if err != nil {
		respondWithError(h.logger, c, http.StatusInternalServerError, ErrInternalServerError)
		return
	}

	c.JSON(http.StatusOK, SuccessResponse{Message: "Schedule patched successfully"})
}

// DeleteSchedule godoc
// @Summary Delete a schedule
// @Description Delete an existing schedule
// @Tags Schedules
// @Accept json
// @Produce json
// @Param id path int64 true "Schedule ID"
// @Success 200 {object} SuccessResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /schedules/{id} [delete]
func (h *Handler) DeleteSchedule(c *gin.Context) {
	scheduleID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		respondWithError(h.logger, c, http.StatusBadRequest, ErrInvalidScheduleID)
		return
	}

	err = h.services.ScheduleService.Delete(c.Request.Context(), scheduleID)
	if err != nil {
		respondWithError(h.logger, c, http.StatusInternalServerError, ErrInternalServerError)
		return
	}

	c.JSON(http.StatusOK, SuccessResponse{Message: "Schedule deleted successfully"})
}

// GetScheduleByID godoc
// @Summary Get a schedule by ID
// @Description Get a schedule by its ID
// @Tags Schedules
// @Accept json
// @Produce json
// @Param id path int64 true "Schedule ID"
// @Success 200 {object} domain.Schedule
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /schedules/{id} [get]
func (h *Handler) GetScheduleByID(c *gin.Context) {
	scheduleID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		respondWithError(h.logger, c, http.StatusBadRequest, ErrInvalidScheduleID)
		return
	}

	schedule, err := h.services.ScheduleService.GetByID(c.Request.Context(), scheduleID)
	if err != nil {
		respondWithError(h.logger, c, http.StatusInternalServerError, ErrScheduleNotFound)
		return
	}

	c.JSON(http.StatusOK, schedule)
}

// GetAllSchedules godoc
// @Summary Get all schedules
// @Description Get a list of all schedules
// @Tags Schedules
// @Accept json
// @Produce json
// @Success 200 {array} domain.Schedule
// @Failure 500 {object} ErrorResponse
// @Router /schedules [get]
func (h *Handler) GetAllSchedules(c *gin.Context) {
	schedules, err := h.services.ScheduleService.GetAll(c.Request.Context())
	if err != nil {
		respondWithError(h.logger, c, http.StatusInternalServerError, ErrSchedulesNotFound)
		return
	}

	c.JSON(http.StatusOK, schedules)
}

// GetSchedulesByGroupID godoc
// @Summary Get schedules by group ID
// @Description Get schedules by group ID
// @Tags Schedules
// @Accept json
// @Produce json
// @Param group_id path string true "Group ID"
// @Success 200 {array} domain.Schedule
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /schedules/group/{group_id} [get]
func (h *Handler) GetSchedulesByGroupID(c *gin.Context) {
	groupID := c.Param("group_id")

	schedules, err := h.services.ScheduleService.GetByGroupID(c.Request.Context(), groupID)
	if err != nil {
		respondWithError(h.logger, c, http.StatusInternalServerError, ErrSchedulesNotFound)
		return
	}

	c.JSON(http.StatusOK, schedules)
}

// GetSchedulesByTeacherID godoc
// @Summary Get schedules by teacher ID
// @Description Get schedules by teacher ID
// @Tags Schedules
// @Accept json
// @Produce json
// @Param teacher_id path int64 true "Teacher ID"
// @Success 200 {array} domain.Schedule
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /schedules/teacher/{teacher_id} [get]
func (h *Handler) GetSchedulesByTeacherID(c *gin.Context) {
	teacherID, err := strconv.ParseInt(c.Param("teacher_id"), 10, 64)
	if err != nil {
		respondWithError(h.logger, c, http.StatusBadRequest, ErrInvalidTeacherID)
		return
	}

	schedules, err := h.services.ScheduleService.GetByTeacherID(c.Request.Context(), teacherID)
	if err != nil {
		respondWithError(h.logger, c, http.StatusInternalServerError, ErrSchedulesNotFound)
		return
	}

	c.JSON(http.StatusOK, schedules)
}

// GetSchedulesByGroupAndWeekType godoc
// @Summary Get schedules by group ID and week type
// @Description Get schedules by group ID and week type
// @Tags Schedules
// @Accept json
// @Produce json
// @Param group_id path string true "Group ID"
// @Param week_type path string true "Week Type"
// @Success 200 {array} domain.Schedule
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /schedules/group/{group_id}/week/{week_type} [get]
func (h *Handler) GetSchedulesByGroupAndWeekType(c *gin.Context) {
	groupID := c.Param("group_id")
	weekType := c.Param("week_type")

	schedules, err := h.services.ScheduleService.GetByGroupAndWeekType(c.Request.Context(), groupID, weekType)
	if err != nil {
		respondWithError(h.logger, c, http.StatusInternalServerError, ErrSchedulesNotFound)
		return
	}

	c.JSON(http.StatusOK, schedules)
}

// GetSchedulesByTeacherAndWeekType godoc
// @Summary Get schedules by teacher ID and week type
// @Description Get schedules by teacher ID and week type
// @Tags Schedules
// @Accept json
// @Produce json
// @Param teacher_id path int64 true "Teacher ID"
// @Param week_type path string true "Week Type"
// @Success 200 {array} domain.Schedule
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /schedules/teacher/{teacher_id}/week/{week_type} [get]
func (h *Handler) GetSchedulesByTeacherAndWeekType(c *gin.Context) {
	teacherID, err := strconv.ParseInt(c.Param("teacher_id"), 10, 64)
	if err != nil {
		respondWithError(h.logger, c, http.StatusBadRequest, ErrInvalidTeacherID)
		return
	}
	weekType := c.Param("week_type")

	schedules, err := h.services.ScheduleService.GetByTeacherAndWeekType(c.Request.Context(), teacherID, weekType)
	if err != nil {
		respondWithError(h.logger, c, http.StatusInternalServerError, ErrSchedulesNotFound)
		return
	}

	c.JSON(http.StatusOK, schedules)
}

// GetSchedulesByGroupWeekTypeAndDay godoc
// @Summary Get schedules by group ID, week type, and day of the week
// @Description Get schedules by group ID, week type, and day of the week
// @Tags Schedules
// @Accept json
// @Produce json
// @Param group_id path string true "Group ID"
// @Param week_type path string true "Week Type"
// @Param day_of_week path string true "Day of the Week"
// @Success 200 {array} domain.Schedule
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /schedules/group/{group_id}/week/{week_type}/day/{day_of_week} [get]
func (h *Handler) GetSchedulesByGroupWeekTypeAndDay(c *gin.Context) {
	groupID := c.Param("group_id")
	weekType := c.Param("week_type")
	dayOfWeek := c.Param("day_of_week")

	schedules, err := h.services.ScheduleService.GetByGroupWeekTypeAndDay(c.Request.Context(), groupID, weekType, dayOfWeek)
	if err != nil {
		respondWithError(h.logger, c, http.StatusInternalServerError, ErrSchedulesNotFound)
		return
	}

	c.JSON(http.StatusOK, schedules)
}

// GetSchedulesByTeacherWeekTypeAndDay godoc
// @Summary Get schedules by teacher ID, week type, and day of the week
// @Description Get schedules by teacher ID, week type, and day of the week
// @Tags Schedules
// @Accept json
// @Produce json
// @Param teacher_id path int64 true "Teacher ID"
// @Param week_type path string true "Week Type"
// @Param day_of_week path string true "Day of the Week"
// @Success 200 {array} domain.Schedule
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /schedules/teacher/{teacher_id}/week/{week_type}/day/{day_of_week} [get]
func (h *Handler) GetSchedulesByTeacherWeekTypeAndDay(c *gin.Context) {
	teacherID, err := strconv.ParseInt(c.Param("teacher_id"), 10, 64)
	if err != nil {
		respondWithError(h.logger, c, http.StatusBadRequest, ErrInvalidTeacherID)
		return
	}
	weekType := c.Param("week_type")
	dayOfWeek := c.Param("day_of_week")

	schedules, err := h.services.ScheduleService.GetByTeacherWeekTypeAndDay(c.Request.Context(), teacherID, weekType, dayOfWeek)
	if err != nil {
		respondWithError(h.logger, c, http.StatusInternalServerError, ErrSchedulesNotFound)
		return
	}

	c.JSON(http.StatusOK, schedules)
}

// GetActualSchedulesByGroupID godoc
// @Summary Get actual schedules by group ID
// @Description Get actual schedules by group ID
// @Tags Schedules
// @Accept json
// @Produce json
// @Param group_id path string true "Group ID"
// @Success 200 {array} domain.Schedule
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /schedules/group/{group_id}/actual [get]
func (h *Handler) GetActualSchedulesByGroupID(c *gin.Context) {
	groupID := c.Param("group_id")

	schedules, err := h.services.ScheduleService.GetActualByGroupID(c.Request.Context(), groupID)
	if err != nil {
		respondWithError(h.logger, c, http.StatusInternalServerError, ErrSchedulesNotFound)
		return
	}

	c.JSON(http.StatusOK, schedules)
}

// GetActualSchedulesByTeacherID godoc
// @Summary Get actual schedules by teacher ID
// @Description Get actual schedules by teacher ID
// @Tags Schedules
// @Accept json
// @Produce json
// @Param teacher_id path int64 true "Teacher ID"
// @Success 200 {array} domain.Schedule
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /schedules/teacher/{teacher_id}/actual [get]
func (h *Handler) GetActualSchedulesByTeacherID(c *gin.Context) {
	teacherID, err := strconv.ParseInt(c.Param("teacher_id"), 10, 64)
	if err != nil {
		respondWithError(h.logger, c, http.StatusBadRequest, ErrInvalidTeacherID)
		return
	}

	schedules, err := h.services.ScheduleService.GetActualByTeacherID(c.Request.Context(), teacherID)
	if err != nil {
		respondWithError(h.logger, c, http.StatusInternalServerError, ErrSchedulesNotFound)
		return
	}

	c.JSON(http.StatusOK, schedules)
}

// GetSchedulesByGroupAndWeekType godoc
// @Summary Get schedules by group ID and week type
// @Description Get schedules by group ID and week type
// @Tags Schedules
// @Accept json
// @Produce json
// @Param group_id path string true "Group ID"
// @Param week_type path string true "Week Type"
// @Success 200 {array} domain.Schedule
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /schedules/group/{group_id}/week/{week_type} [get]
func (h *Handler) GetActualSchedulesByGroupAndWeekType(c *gin.Context) {
	data, ok := c.Get(groupCtx)
	if !ok {
		respondWithError(h.logger, c, http.StatusUnauthorized, "Group ID not found in context")
		return
	}

	groupID, ok := data.(string)
	if !ok {
		respondWithError(h.logger, c, http.StatusInternalServerError, "Failed to convert group ID")
		return
	}

	req := GetASByGWRequest{GroupID: groupID, WeekType: c.Param("week")}
	if err := h.validate.Struct(req); err != nil {
		respondWithError(h.logger, c, http.StatusBadRequest, ErrValidationFailed)
		return
	}

	schedules, err := h.services.ScheduleService.GetActualByGroupAndWeekType(c.Request.Context(), req.GroupID, req.WeekType)
	if err != nil {
		respondWithError(h.logger, c, http.StatusInternalServerError, ErrSchedulesNotFound)
		return
	}

	c.JSON(http.StatusOK, schedules)
}

// GetSchedulesByGroupAndWeekType godoc
// @Summary Get schedules by group ID and week type
// @Description Get schedules by group ID and week type
// @Tags Schedules
// @Accept json
// @Produce json
// @Param group_id path string true "Group ID"
// @Param week_type path string true "Week Type"
// @Success 200 {array} domain.Schedule
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /schedules/group/{group_id}/week/{week_type} [get]
func (h *Handler) GetActualSchedulesByGroupWeekTypeAndDay(c *gin.Context) {

	data, ok := c.Get(groupCtx)
	if !ok {
		respondWithError(h.logger, c, http.StatusUnauthorized, "Group ID not found in context")
		return
	}

	groupID, ok := data.(string)
	if !ok {
		respondWithError(h.logger, c, http.StatusInternalServerError, "Failed to convert group ID")
		return
	}

	req := GetASByGWDRequest{GroupID: groupID, WeekType: c.Param("week"), DayOfWeek: c.Param("day")}
	if err := h.validate.Struct(req); err != nil {
		respondWithError(h.logger, c, http.StatusBadRequest, ErrValidationFailed)
		return
	}

	schedules, err := h.services.ScheduleService.GetActualByGroupWeekTypeAndDay(c.Request.Context(), req.GroupID, req.WeekType, req.DayOfWeek)
	if err != nil {
		respondWithError(h.logger, c, http.StatusInternalServerError, ErrSchedulesNotFound)
		return
	}

	c.JSON(http.StatusOK, schedules)
}

// GetActualSchedulesByTeacherID godoc
// @Summary Get actual schedules by teacher ID
// @Description Get actual schedules by teacher ID
// @Tags Schedules
// @Accept json
// @Produce json
// @Param teacher_id path int64 true "Teacher ID"
// @Success 200 {array} domain.Schedule
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /schedules/teacher/{teacher_id}/actual [get]
func (h *Handler) GetActualSchedulesByTeacherIDWeekType(c *gin.Context) {

	data, ok := c.Get(teacherCtx)
	if !ok {
		respondWithError(h.logger, c, http.StatusUnauthorized, "Teacher ID not found in context")
		return
	}

	teacherID, ok := data.(int64)
	if !ok {
		respondWithError(h.logger, c, http.StatusInternalServerError, "Failed to convert teacher ID")
		return
	}

	req := GetASByTWRequest{TeacherID: teacherID, WeekType: c.Param("week")}
	if err := h.validate.Struct(req); err != nil {
		respondWithError(h.logger, c, http.StatusBadRequest, ErrValidationFailed)
		return
	}

	schedules, err := h.services.ScheduleService.GetActualByTeacherAndWeekType(c.Request.Context(), teacherID, req.WeekType)
	if err != nil {
		respondWithError(h.logger, c, http.StatusInternalServerError, ErrSchedulesNotFound)
		return
	}

	c.JSON(http.StatusOK, schedules)
}

// GetActualSchedulesByTeacherID godoc
// @Summary Get actual schedules by teacher ID
// @Description Get actual schedules by teacher ID
// @Tags Schedules
// @Accept json
// @Produce json
// @Param teacher_id path int64 true "Teacher ID"
// @Success 200 {array} domain.Schedule
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /schedules/teacher/{teacher_id}/actual [get]
func (h *Handler) GetActualSchedulesByTeacherIDWeekTypeAndDay(c *gin.Context) {

	data, ok := c.Get(teacherCtx)
	if !ok {
		respondWithError(h.logger, c, http.StatusUnauthorized, "Teacher ID not found in context")
		return
	}

	teacherID, ok := data.(int64)
	if !ok {
		respondWithError(h.logger, c, http.StatusInternalServerError, "Failed to convert teacher ID")
		return
	}

	req := GetASByTWDRequest{TeacherID: teacherID, WeekType: c.Param("week"), DayOfWeek: c.Param("day")}
	if err := h.validate.Struct(req); err != nil {
		respondWithError(h.logger, c, http.StatusBadRequest, ErrValidationFailed)
		return
	}

	schedules, err := h.services.ScheduleService.GetActualByTeacherWeekTypeAndDay(c.Request.Context(), teacherID, req.WeekType, req.DayOfWeek)
	if err != nil {
		respondWithError(h.logger, c, http.StatusInternalServerError, ErrSchedulesNotFound)
		return
	}

	c.JSON(http.StatusOK, schedules)
}
