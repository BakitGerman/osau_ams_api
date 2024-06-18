package handler

import (
	"net/http"
	"strconv"

	"github.com/BeRebornBng/OsauAmsApi/domain"
	"github.com/gin-gonic/gin"
)

type CreateTeacherRequest struct {
	DepartamentID int64  `json:"departament_id" validate:"required,numeric"`
	LastName      string `json:"last_name" validate:"required,customfieldrusnumregex"`
	FirstName     string `json:"first_name" validate:"required,customfieldrusnumregex"`
	MiddleName    string `json:"middle_name" validate:"required,customfieldrusnumregex"`
	TeacherEmail  string `json:"teacher_email" validate:"required,email"`
}

type PutTeacherRequest struct {
	TeacherID     int64  `json:"teacher_id" validate:"required,numeric"`
	DepartamentID int64  `json:"departament_id" validate:"required,numeric"`
	LastName      string `json:"last_name" validate:"required,customfieldrusnumregex"`
	FirstName     string `json:"first_name" validate:"required,customfieldrusnumregex"`
	MiddleName    string `json:"middle_name" validate:"required,customfieldrusnumregex"`
	TeacherEmail  string `json:"teacher_email" validate:"required,email"`
}

type PatchTeacherRequest struct {
	TeacherID     int64  `json:"teacher_id" validate:"required,numeric"`
	DepartamentID int64  `json:"departament_id" validate:"omitempty,numeric"`
	LastName      string `json:"last_name" validate:"omitempty,customfieldrusnumregex"`
	FirstName     string `json:"first_name" validate:"omitempty,customfieldrusnumregex"`
	MiddleName    string `json:"middle_name" validate:"omitempty,customfieldrusnumregex"`
	TeacherEmail  string `json:"teacher_email" validate:"omitempty,email"`
}

// CreateTeacher godoc
// @Summary Create a teacher
// @Description Create a new teacher
// @Tags Teachers
// @Accept json
// @Produce json
// @Param teacher body CreateTeacherRequest true "Teacher info"
// @Success 201 {object} SuccessResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /teachers [post]
func (h *Handler) CreateTeacher(c *gin.Context) {
	var req CreateTeacherRequest
	if err := c.BindJSON(&req); err != nil {
		respondWithError(h.logger, c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.validate.Struct(req); err != nil {
		respondWithError(h.logger, c, http.StatusBadRequest, err.Error())
		return
	}

	teacher := domain.Teacher{
		DepartamentID: req.DepartamentID,
		LastName:      req.LastName,
		FirstName:     req.FirstName,
		MiddleName:    req.MiddleName,
		TeacherEmail:  req.TeacherEmail,
	}

	err := h.services.TeacherService.Create(c.Request.Context(), teacher)
	if err != nil {
		respondWithError(h.logger, c, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithSuccess(c, http.StatusCreated, "Teacher created successfully")
}

// PutTeacher godoc
// @Summary Update a teacher
// @Description Update an existing teacher
// @Tags Teachers
// @Accept json
// @Produce json
// @Param teacher body PutTeacherRequest true "Teacher info"
// @Success 200 {object} SuccessResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /teachers [put]
func (h *Handler) PutTeacher(c *gin.Context) {
	var req PutTeacherRequest
	if err := c.BindJSON(&req); err != nil {
		respondWithError(h.logger, c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.validate.Struct(req); err != nil {
		respondWithError(h.logger, c, http.StatusBadRequest, err.Error())
		return
	}

	teacher := domain.Teacher{
		TeacherID:     req.TeacherID,
		DepartamentID: req.DepartamentID,
		LastName:      req.LastName,
		FirstName:     req.FirstName,
		MiddleName:    req.MiddleName,
		TeacherEmail:  req.TeacherEmail,
	}

	err := h.services.TeacherService.Put(c.Request.Context(), teacher)
	if err != nil {
		respondWithError(h.logger, c, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithSuccess(c, http.StatusOK, "Teacher updated successfully")
}

// PatchTeacher godoc
// @Summary Partially update a teacher
// @Description Partially update an existing teacher
// @Tags Teachers
// @Accept json
// @Produce json
// @Param teacher body PatchTeacherRequest true "Teacher info"
// @Success 200 {object} SuccessResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /teachers [patch]
func (h *Handler) PatchTeacher(c *gin.Context) {
	var req PatchTeacherRequest
	if err := c.BindJSON(&req); err != nil {
		respondWithError(h.logger, c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.validate.Struct(req); err != nil {
		respondWithError(h.logger, c, http.StatusBadRequest, err.Error())
		return
	}

	teacher := domain.Teacher{
		TeacherID:     req.TeacherID,
		DepartamentID: req.DepartamentID,
		LastName:      req.LastName,
		FirstName:     req.FirstName,
		MiddleName:    req.MiddleName,
		TeacherEmail:  req.TeacherEmail,
	}

	err := h.services.TeacherService.Patch(c.Request.Context(), teacher)
	if err != nil {
		respondWithError(h.logger, c, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithSuccess(c, http.StatusOK, "Teacher patched successfully")
}

// DeleteTeacher godoc
// @Summary Delete a teacher
// @Description Delete an existing teacher
// @Tags Teachers
// @Accept json
// @Produce json
// @Param id path int64 true "Teacher ID"
// @Success 200 {object} SuccessResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /teachers/{id} [delete]
func (h *Handler) DeleteTeacher(c *gin.Context) {
	teacherID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		respondWithError(h.logger, c, http.StatusBadRequest, "Invalid teacher ID")
		return
	}

	err = h.services.TeacherService.Delete(c.Request.Context(), teacherID)
	if err != nil {
		respondWithError(h.logger, c, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithSuccess(c, http.StatusOK, "Teacher deleted successfully")
}

// GetTeacherByID godoc
// @Summary Get a teacher by ID
// @Description Get a teacher by its ID
// @Tags Teachers
// @Accept json
// @Produce json
// @Param id path int64 true "Teacher ID"
// @Success 200 {object} domain.TeacherInfo
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /teachers/{id} [get]
func (h *Handler) GetTeacherByID(c *gin.Context) {
	teacherID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		respondWithError(h.logger, c, http.StatusBadRequest, "Invalid teacher ID")
		return
	}

	teacher, err := h.services.TeacherService.GetByID(c.Request.Context(), teacherID)
	if err != nil {
		respondWithError(h.logger, c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, teacher)
}

// GetTeacherByEmail godoc
// @Summary Get a teacher by email
// @Description Get a teacher by its email
// @Tags Teachers
// @Accept json
// @Produce json
// @Param email path string true "Teacher Email"
// @Success 200 {object} domain.TeacherInfo
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /teachers/email/{email} [get]
func (h *Handler) GetTeacherByEmail(c *gin.Context) {
	teacherEmail := c.Param("email")

	teacher, err := h.services.TeacherService.GetByEmail(c.Request.Context(), teacherEmail)
	if err != nil {
		respondWithError(h.logger, c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, teacher)
}

// GetAllTeachers godoc
// @Summary Get all teachers
// @Description Get a list of all teachers
// @Tags Teachers
// @Accept json
// @Produce json
// @Success 200 {array} domain.TeacherInfo
// @Failure 500 {object} ErrorResponse
// @Router /teachers [get]
func (h *Handler) GetAllTeachers(c *gin.Context) {
	teachers, err := h.services.TeacherService.GetAll(c.Request.Context())
	if err != nil {
		respondWithError(h.logger, c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, teachers)
}

// GetTeachersByDepartamentID godoc
// @Summary Get teachers by departament ID
// @Description Get a list of teachers by departament ID
// @Tags Teachers
// @Accept json
// @Produce json
// @Param departament_id path int64 true "Departament ID"
// @Success 200 {array} domain.TeacherInfo
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /teachers/departament/{departament_id} [get]
func (h *Handler) GetTeachersByDepartamentID(c *gin.Context) {
	departamentID, err := strconv.ParseInt(c.Param("departament_id"), 10, 64)
	if err != nil {
		respondWithError(h.logger, c, http.StatusBadRequest, "Invalid departament ID")
		return
	}

	teachers, err := h.services.TeacherService.GetAllByDepartamentID(c.Request.Context(), departamentID)
	if err != nil {
		respondWithError(h.logger, c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, teachers)
}
