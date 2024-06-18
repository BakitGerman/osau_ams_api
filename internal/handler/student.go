package handler

import (
	"net/http"
	"strconv"

	"github.com/BeRebornBng/OsauAmsApi/domain"
	"github.com/gin-gonic/gin"
)

type CreateStudentRequest struct {
	LastName   string `json:"last_name" validate:"required,customfieldrusnumregex"`
	FirstName  string `json:"first_name" validate:"required,customfieldrusnumregex"`
	MiddleName string `json:"middle_name" validate:"required,customfieldrusnumregex"`
	GroupID    string `json:"group_id" validate:"required,customgroupidregex"`
}

type PutStudentRequest struct {
	StudentID  int64  `json:"student_id" validate:"required,numeric"`
	LastName   string `json:"last_name" validate:"required,customfieldrusnumregex"`
	FirstName  string `json:"first_name" validate:"required,customfieldrusnumregex"`
	MiddleName string `json:"middle_name" validate:"required,customfieldrusnumregex"`
	GroupID    string `json:"group_id" validate:"required,customgroupidregex"`
}

type PatchStudentRequest struct {
	StudentID  int64  `json:"student_id" validate:"required,numeric"`
	LastName   string `json:"last_name" validate:"omitempty,customfieldrusnumregex"`
	FirstName  string `json:"first_name" validate:"omitempty,customfieldrusnumregex"`
	MiddleName string `json:"middle_name" validate:"omitempty,customfieldrusnumregex"`
	GroupID    string `json:"group_id" validate:"omitempty,customgroupidregex"`
}

// CreateStudent godoc
// @Summary Create a student
// @Description Create a new student
// @Tags Students
// @Accept json
// @Produce json
// @Param student body CreateStudentRequest true "Student info"
// @Success 201 {object} SuccessResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /students [post]
func (h *Handler) CreateStudent(c *gin.Context) {
	var req CreateStudentRequest
	if err := c.BindJSON(&req); err != nil {
		respondWithError(h.logger, c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.validate.Struct(req); err != nil {
		respondWithError(h.logger, c, http.StatusBadRequest, err.Error())
		return
	}

	student := domain.Student{
		LastName:   req.LastName,
		FirstName:  req.FirstName,
		MiddleName: req.MiddleName,
		GroupID:    req.GroupID,
	}

	err := h.services.StudentService.Create(c.Request.Context(), student)
	if err != nil {
		respondWithError(h.logger, c, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithSuccess(c, http.StatusCreated, "Student created successfully")
}

// PutStudent godoc
// @Summary Update a student
// @Description Update an existing student
// @Tags Students
// @Accept json
// @Produce json
// @Param student body PutStudentRequest true "Student info"
// @Success 200 {object} SuccessResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /students [put]
func (h *Handler) PutStudent(c *gin.Context) {
	var req PutStudentRequest
	if err := c.BindJSON(&req); err != nil {
		respondWithError(h.logger, c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.validate.Struct(req); err != nil {
		respondWithError(h.logger, c, http.StatusBadRequest, err.Error())
		return
	}

	student := domain.Student{
		StudentID:  req.StudentID,
		LastName:   req.LastName,
		FirstName:  req.FirstName,
		MiddleName: req.MiddleName,
		GroupID:    req.GroupID,
	}

	err := h.services.StudentService.Put(c.Request.Context(), student)
	if err != nil {
		respondWithError(h.logger, c, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithSuccess(c, http.StatusOK, "Student updated successfully")
}

// PatchStudent godoc
// @Summary Partially update a student
// @Description Partially update an existing student
// @Tags Students
// @Accept json
// @Produce json
// @Param student body PatchStudentRequest true "Student info"
// @Success 200 {object} SuccessResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /students [patch]
func (h *Handler) PatchStudent(c *gin.Context) {
	var req PatchStudentRequest
	if err := c.BindJSON(&req); err != nil {
		respondWithError(h.logger, c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.validate.Struct(req); err != nil {
		respondWithError(h.logger, c, http.StatusBadRequest, err.Error())
		return
	}

	student := domain.Student{
		StudentID:  req.StudentID,
		LastName:   req.LastName,
		FirstName:  req.FirstName,
		MiddleName: req.MiddleName,
		GroupID:    req.GroupID,
	}

	err := h.services.StudentService.Patch(c.Request.Context(), student)
	if err != nil {
		respondWithError(h.logger, c, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithSuccess(c, http.StatusOK, "Student patched successfully")
}

// DeleteStudent godoc
// @Summary Delete a student
// @Description Delete an existing student
// @Tags Students
// @Accept json
// @Produce json
// @Param id path int64 true "Student ID"
// @Success 200 {object} SuccessResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /students/{id} [delete]
func (h *Handler) DeleteStudent(c *gin.Context) {
	studentID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		respondWithError(h.logger, c, http.StatusBadRequest, "Invalid student ID")
		return
	}

	err = h.services.StudentService.Delete(c.Request.Context(), studentID)
	if err != nil {
		respondWithError(h.logger, c, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithSuccess(c, http.StatusOK, "Student deleted successfully")
}

// GetStudentByID godoc
// @Summary Get a student by ID
// @Description Get a student by its ID
// @Tags Students
// @Accept json
// @Produce json
// @Param id path int64 true "Student ID"
// @Success 200 {object} domain.Student
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /students/{id} [get]
func (h *Handler) GetStudentByID(c *gin.Context) {
	studentID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		respondWithError(h.logger, c, http.StatusBadRequest, "Invalid student ID")
		return
	}

	student, err := h.services.StudentService.GetByID(c.Request.Context(), studentID)
	if err != nil {
		respondWithError(h.logger, c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, student)
}

// GetAllStudents godoc
// @Summary Get all students
// @Description Get a list of all students
// @Tags Students
// @Accept json
// @Produce json
// @Success 200 {array} domain.Student
// @Failure 500 {object} ErrorResponse
// @Router /students [get]
func (h *Handler) GetAllStudents(c *gin.Context) {
	students, err := h.services.StudentService.GetAll(c.Request.Context())
	if err != nil {
		respondWithError(h.logger, c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, students)
}

// GetStudentsByGroupID godoc
// @Summary Get students by group ID
// @Description Get a list of students by group ID
// @Tags Students
// @Accept json
// @Produce json
// @Param group_id path string true "Group ID"
// @Success 200 {array} domain.Student
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /students/group/{group_id} [get]
func (h *Handler) GetStudentsByGroupID(c *gin.Context) {
	groupID := c.Param("group_id")

	students, err := h.services.StudentService.GetAllByGroupID(c.Request.Context(), groupID)
	if err != nil {
		respondWithError(h.logger, c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, students)
}
