package handler

import (
	"net/http"
	"strconv"
	"time"

	"github.com/BeRebornBng/OsauAmsApi/domain"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// CreateAttendanceRequest represents the request body for creating an attendance
type CreateAttendanceRequest struct {
	StudentID      int64   `json:"student_id" validate:"required,numeric"`
	ScheduleID     int64   `json:"schedule_id" validate:"required,numeric"`
	Presence       *bool   `json:"presence"`
	LateArrival    *bool   `json:"late_arrival"`
	Respectfulness *bool   `json:"respectfulness"`
	Reason         *string `json:"reason" validate:"omitempty,customfieldrusregex"`
	Created        string  `json:"created" validate:"required,datetime=2006-01-02"`
}

type CreateAttendancesRequest struct {
	Attendances []CreateAttendanceRequest `json:"attendances" binding:"required,dive,required"`
}

// PutAttendanceRequest represents the request body for updating an attendance
type PutAttendanceRequest struct {
	AttendanceID   int64   `json:"attendance_id" validate:"required,numeric"`
	StudentID      int64   `json:"student_id" validate:"required,numeric"`
	ScheduleID     int64   `json:"schedule_id" validate:"required,numeric"`
	Presence       *bool   `json:"presence"`
	LateArrival    *bool   `json:"late_arrival"`
	Respectfulness *bool   `json:"respectfulness"`
	Reason         *string `json:"reason" validate:"omitempty,customfieldrusregex"`
}

// PatchAttendanceRequest represents the request body for partially updating an attendance
type PatchAttendanceRequest struct {
	AttendanceID   int64   `json:"attendance_id" validate:"required,numeric"`
	StudentID      int64   `json:"student_id" validate:"omitempty,numeric"`
	ScheduleID     int64   `json:"schedule_id" validate:"omitempty,numeric"`
	Presence       *bool   `json:"presence"`
	LateArrival    *bool   `json:"late_arrival"`
	Respectfulness *bool   `json:"respectfulness"`
	Reason         *string `json:"reason"`
	Created        string  `json:"created" validate:"omitempty,datetime=2006-01-02"`
}

type GetHeadmanAllAttendancesRequest struct {
	GroupID    string `json:"group_id" validate:"required,customgroupidregex"`
	ScheduleID int64  `json:"schedule_id" validate:"omitempty,numeric"`
	Created    string `json:"created" validate:"required,datetime=2006-01-02"`
}

// CreateAttendance godoc
// @Summary Create an attendance
// @Description Create a new attendance
// @Tags Attendance
// @Accept json
// @Produce json
// @Param attendance body CreateAttendanceRequest true "Attendance info"
// @Success 201 {object} SuccessResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /attendances [post]
func (h *Handler) CreateAttendance(c *gin.Context) {

	var req CreateAttendanceRequest
	if err := c.BindJSON(&req); err != nil {
		respondWithError(h.logger, c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.validate.Struct(req); err != nil {
		errs := translateValidationErrors(err.(validator.ValidationErrors), h.translator)
		respondWithError(h.logger, c, http.StatusBadRequest, errs[0])
		return
	}

	date, err := time.Parse("2006-01-02", req.Created)
	if err != nil {
		respondWithError(h.logger, c, http.StatusBadRequest, err.Error())
		return
	}

	attendance := domain.Attendance{
		StudentID:      req.StudentID,
		ScheduleID:     req.ScheduleID,
		Presence:       req.Presence,
		LateArrival:    req.LateArrival,
		Respectfulness: req.Respectfulness,
		Reason:         req.Reason,
		Created:        date,
	}

	err = h.services.AttendanceService.Create(c.Request.Context(), attendance)
	if err != nil {
		respondWithError(h.logger, c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, SuccessResponse{Message: "Attendance created successfully"})
}

func (h *Handler) CreateAttendances(c *gin.Context) {

	var req CreateAttendancesRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		respondWithError(h.logger, c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.validate.Struct(req); err != nil {
		respondWithError(h.logger, c, http.StatusBadRequest, err.Error())
		return
	}

	for _, attendance := range req.Attendances {
		if err := h.validate.Struct(attendance); err != nil {
			errs := translateValidationErrors(err.(validator.ValidationErrors), h.translator)
			respondWithError(h.logger, c, http.StatusBadRequest, errs[0])
			return
		}
	}

	for _, attendance := range req.Attendances {
		date, err := time.Parse("2006-01-02", attendance.Created)
		if err != nil {
			respondWithError(h.logger, c, http.StatusBadRequest, err.Error())
			return
		}
		attendance := domain.Attendance{
			StudentID:      attendance.StudentID,
			ScheduleID:     attendance.ScheduleID,
			Presence:       attendance.Presence,
			LateArrival:    attendance.LateArrival,
			Respectfulness: attendance.Respectfulness,
			Reason:         attendance.Reason,
			Created:        date,
		}
		if err := h.services.AttendanceService.Create(c.Request.Context(), attendance); err != nil {
			respondWithError(h.logger, c, http.StatusBadRequest, err.Error())
			return
		}
	}

	c.JSON(http.StatusCreated, SuccessResponse{Message: "Attendance created successfully"})
}

// PutAttendance godoc
// @Summary Update an attendance
// @Description Update an existing attendance
// @Tags Attendance
// @Accept json
// @Produce json
// @Param attendance body PutAttendanceRequest true "Attendance info"
// @Success 200 {object} SuccessResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /attendances [put]
func (h *Handler) PutAttendance(c *gin.Context) {
	var req PutAttendanceRequest
	if err := c.BindJSON(&req); err != nil {
		respondWithError(h.logger, c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.validate.Struct(req); err != nil {
		errs := translateValidationErrors(err.(validator.ValidationErrors), h.translator)
		respondWithError(h.logger, c, http.StatusBadRequest, errs[0])
		return
	}

	attendance := domain.Attendance{
		StudentID:      req.StudentID,
		ScheduleID:     req.ScheduleID,
		Presence:       req.Presence,
		LateArrival:    req.LateArrival,
		Respectfulness: req.Respectfulness,
		Reason:         req.Reason,
	}

	err := h.services.AttendanceService.Put(c.Request.Context(), attendance)
	if err != nil {
		respondWithError(h.logger, c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, SuccessResponse{Message: "Attendance updated successfully"})
}

type PutAttendancesRequest struct {
	Attendances []PutAttendanceRequest `json:"attendances" binding:"required,dive,required"`
}

// PutAttendances godoc
// @Summary Update multiple attendances
// @Description Update multiple existing attendances
// @Tags Attendance
// @Accept json
// @Produce json
// @Param attendance body PutAttendancesRequest true "Attendances info"
// @Success 200 {object} SuccessResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /attendances [put]
func (h *Handler) PutAttendances(c *gin.Context) {
	var req PutAttendancesRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		respondWithError(h.logger, c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.validate.Struct(req); err != nil {
		errs := translateValidationErrors(err.(validator.ValidationErrors), h.translator)
		respondWithError(h.logger, c, http.StatusBadRequest, errs[0])
		return
	}

	for _, attendance := range req.Attendances {
		if err := h.validate.Struct(attendance); err != nil {
			errs := translateValidationErrors(err.(validator.ValidationErrors), h.translator)
			respondWithError(h.logger, c, http.StatusBadRequest, errs[0])
			return
		}
	}

	for _, attendance := range req.Attendances {

		attendance := domain.Attendance{
			AttendanceID:   attendance.AttendanceID,
			StudentID:      attendance.StudentID,
			ScheduleID:     attendance.ScheduleID,
			Presence:       attendance.Presence,
			LateArrival:    attendance.LateArrival,
			Respectfulness: attendance.Respectfulness,
			Reason:         attendance.Reason,
		}
		if err := h.services.AttendanceService.Put(c.Request.Context(), attendance); err != nil {
			respondWithError(h.logger, c, http.StatusInternalServerError, err.Error())
			return
		}
	}

	c.JSON(http.StatusOK, SuccessResponse{Message: "Attendances updated successfully"})
}

// PatchAttendance godoc
// @Summary Partially update an attendance
// @Description Partially update an existing attendance
// @Tags Attendance
// @Accept json
// @Produce json
// @Param attendance body PatchAttendanceRequest true "Attendance info"
// @Success 200 {object} SuccessResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /attendances [patch]
func (h *Handler) PatchAttendance(c *gin.Context) {
	var req PatchAttendanceRequest
	if err := c.BindJSON(&req); err != nil {
		respondWithError(h.logger, c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.validate.Struct(req); err != nil {
		respondWithError(h.logger, c, http.StatusBadRequest, err.Error())
		return
	}

	date, err := time.Parse("2006-01-02", req.Created)
	if err != nil {
		respondWithError(h.logger, c, http.StatusBadRequest, err.Error())
		return
	}

	attendance := domain.Attendance{
		StudentID:      req.StudentID,
		ScheduleID:     req.ScheduleID,
		Presence:       req.Presence,
		LateArrival:    req.LateArrival,
		Respectfulness: req.Respectfulness,
		Reason:         req.Reason,
		Created:        date,
	}

	err = h.services.AttendanceService.Patch(c.Request.Context(), attendance)
	if err != nil {
		respondWithError(h.logger, c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, SuccessResponse{Message: "Attendance patched successfully"})
}

// DeleteAttendance godoc
// @Summary Delete an attendance
// @Description Delete an existing attendance
// @Tags Attendance
// @Accept json
// @Produce json
// @Param id path int64 true "Attendance ID"
// @Success 200 {object} SuccessResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /attendances/{id} [delete]
func (h *Handler) DeleteAttendance(c *gin.Context) {
	attendanceID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		respondWithError(h.logger, c, http.StatusBadRequest, "Invalid attendance ID")
		return
	}

	err = h.services.AttendanceService.Delete(c.Request.Context(), attendanceID)
	if err != nil {
		respondWithError(h.logger, c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, SuccessResponse{Message: "Attendance deleted successfully"})
}

// GetAttendanceByID godoc
// @Summary Get an attendance by ID
// @Description Get an attendance by its ID
// @Tags Attendance
// @Accept json
// @Produce json
// @Param id path int64 true "Attendance ID"
// @Success 200 {object} domain.Attendance
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /attendances/{id} [get]
func (h *Handler) GetAttendanceByID(c *gin.Context) {
	attendanceID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		respondWithError(h.logger, c, http.StatusBadRequest, "Invalid attendance ID")
		return
	}

	attendance, err := h.services.AttendanceService.GetByID(c.Request.Context(), attendanceID)
	if err != nil {
		respondWithError(h.logger, c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, attendance)
}

// GetAttendancesByStudentID godoc
// @Summary Get attendances by student ID
// @Description Get attendances by student ID
// @Tags Attendance
// @Accept json
// @Produce json
// @Param student_id path int64 true "Student ID"
// @Success 200 {array} domain.Attendance
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /attendances/student/{student_id} [get]
func (h *Handler) GetAttendancesByStudentID(c *gin.Context) {
	studentID, err := strconv.ParseInt(c.Param("student_id"), 10, 64)
	if err != nil {
		respondWithError(h.logger, c, http.StatusBadRequest, "Invalid student ID")
		return
	}

	attendances, err := h.services.AttendanceService.GetByStudentID(c.Request.Context(), studentID)
	if err != nil {
		respondWithError(h.logger, c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, attendances)
}

/*
// GetAllAttendances godoc
// @Summary Get all attendances
// @Description Get a list of all attendances
// @Tags Attendance
// @Accept json
// @Produce json
// @Success 200 {array} domain.Attendance
// @Failure 500 {object} ErrorResponse
// @Router /attendances [get]
func (h *Handler) GetAllAttendances(c *gin.Context) {
	attendances, err := h.services.AttendanceService.GetAll(c.Request.Context())
	if err != nil {
		respondWithError(h.logger, c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, attendances)
}
*/

// GetAllAttendances godoc
// @Summary Get all attendances
// @Description Get a list of all attendances
// @Tags Attendance
// @Accept json
// @Produce json
// @Success 200 {array} domain.Attendance
// @Failure 500 {object} ErrorResponse
// @Router /attendances [get]
func (h *Handler) GetHeadmanAllAttendances(c *gin.Context) {

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

	scheduleID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		respondWithError(h.logger, c, http.StatusInternalServerError, "Failed to convert schedule ID")
		return
	}

	req := GetHeadmanAllAttendancesRequest{GroupID: groupID, ScheduleID: scheduleID, Created: c.Param("date")}

	if err := h.validate.Struct(req); err != nil {
		respondWithError(h.logger, c, http.StatusBadRequest, err.Error())
		return
	}

	date, err := time.Parse("2006-01-02", req.Created)
	if err != nil {
		respondWithError(h.logger, c, http.StatusBadRequest, err.Error())
		return
	}

	attendances, err := h.services.AttendanceService.GetAllByGroupIDAndCreated(c.Request.Context(), groupID, scheduleID, date)
	if err != nil {
		respondWithError(h.logger, c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, attendances)
}

// GetAllAttendances godoc
// @Summary Get all attendances
// @Description Get a list of all attendances
// @Tags Attendance
// @Accept json
// @Produce json
// @Success 200 {array} domain.Attendance
// @Failure 500 {object} ErrorResponse
// @Router /attendances [get]
func (h *Handler) GetTeacherAllAttendances(c *gin.Context) {

	scheduleID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		respondWithError(h.logger, c, http.StatusInternalServerError, "Failed to convert schedule ID")
		return
	}

	req := GetHeadmanAllAttendancesRequest{GroupID: c.Param("group_id"), ScheduleID: scheduleID, Created: c.Param("date")}

	if err := h.validate.Struct(req); err != nil {
		respondWithError(h.logger, c, http.StatusBadRequest, err.Error())
		return
	}

	date, err := time.Parse("2006-01-02", req.Created)
	if err != nil {
		respondWithError(h.logger, c, http.StatusBadRequest, err.Error())
		return
	}

	attendances, err := h.services.AttendanceService.GetAllByGroupIDAndCreated(c.Request.Context(), req.GroupID, scheduleID, date)
	if err != nil {
		respondWithError(h.logger, c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, attendances)
}
