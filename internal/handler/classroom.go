package handler

import (
	"net/http"
	"strconv"

	"github.com/BeRebornBng/OsauAmsApi/domain"
	"github.com/gin-gonic/gin"
)

type CreateClassroomRequest struct {
	ClassroomName string `json:"classroom_name" validate:"required,customfieldrusnumregex"`
}

type PutClassroomRequest struct {
	ClassroomID   int64  `json:"classroom_id" validate:"required,numeric"`
	ClassroomName string `json:"classroom_name" validate:"required,customfieldrusnumregex"`
}

type PatchClassroomRequest struct {
	ClassroomID   int64  `json:"classroom_id" validate:"required,numeric"`
	ClassroomName string `json:"classroom_name" validate:"omitempty,customfieldrusnumregex"`
}

// CreateClassroom godoc
// @Summary Create a classroom
// @Description Create a new classroom
// @Tags Classrooms
// @Accept json
// @Produce json
// @Param classroom body CreateClassroomRequest true "Classroom info"
// @Success 201 {object} SuccessResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /classrooms [post]
func (h *Handler) CreateClassroom(c *gin.Context) {
	var req CreateClassroomRequest
	if err := c.BindJSON(&req); err != nil {
		respondWithError(h.logger, c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.validate.Struct(req); err != nil {
		respondWithError(h.logger, c, http.StatusBadRequest, err.Error())
		return
	}

	classroom := domain.Classroom{
		ClassroomName: req.ClassroomName,
	}

	err := h.services.ClassroomService.Create(c.Request.Context(), classroom)
	if err != nil {
		respondWithError(h.logger, c, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithSuccess(c, http.StatusCreated, "Classroom created successfully")
}

// PutClassroom godoc
// @Summary Update a classroom
// @Description Update an existing classroom
// @Tags Classrooms
// @Accept json
// @Produce json
// @Param classroom body PutClassroomRequest true "Classroom info"
// @Success 200 {object} SuccessResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /classrooms [put]
func (h *Handler) PutClassroom(c *gin.Context) {
	var req PutClassroomRequest
	if err := c.BindJSON(&req); err != nil {
		respondWithError(h.logger, c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.validate.Struct(req); err != nil {
		respondWithError(h.logger, c, http.StatusBadRequest, err.Error())
		return
	}

	classroom := domain.Classroom{
		ClassroomID:   req.ClassroomID,
		ClassroomName: req.ClassroomName,
	}

	err := h.services.ClassroomService.Put(c.Request.Context(), classroom)
	if err != nil {
		respondWithError(h.logger, c, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithSuccess(c, http.StatusOK, "Classroom updated successfully")
}

// PatchClassroom godoc
// @Summary Partially update a classroom
// @Description Partially update an existing classroom
// @Tags Classrooms
// @Accept json
// @Produce json
// @Param classroom body PatchClassroomRequest true "Classroom info"
// @Success 200 {object} SuccessResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /classrooms [patch]
func (h *Handler) PatchClassroom(c *gin.Context) {
	var req PatchClassroomRequest
	if err := c.BindJSON(&req); err != nil {
		respondWithError(h.logger, c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.validate.Struct(req); err != nil {
		respondWithError(h.logger, c, http.StatusBadRequest, err.Error())
		return
	}

	classroom := domain.Classroom{
		ClassroomID:   req.ClassroomID,
		ClassroomName: req.ClassroomName,
	}

	err := h.services.ClassroomService.Patch(c.Request.Context(), classroom)
	if err != nil {
		respondWithError(h.logger, c, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithSuccess(c, http.StatusOK, "Classroom patched successfully")
}

// DeleteClassroom godoc
// @Summary Delete a classroom
// @Description Delete an existing classroom
// @Tags Classrooms
// @Accept json
// @Produce json
// @Param id path int64 true "Classroom ID"
// @Success 200 {object} SuccessResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /classrooms/{id} [delete]
func (h *Handler) DeleteClassroom(c *gin.Context) {
	classroomID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		respondWithError(h.logger, c, http.StatusBadRequest, "Invalid classroom ID")
		return
	}

	err = h.services.ClassroomService.Delete(c.Request.Context(), classroomID)
	if err != nil {
		respondWithError(h.logger, c, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithSuccess(c, http.StatusOK, "Classroom deleted successfully")
}

// GetClassroomByID godoc
// @Summary Get a classroom by ID
// @Description Get a classroom by its ID
// @Tags Classrooms
// @Accept json
// @Produce json
// @Param id path int64 true "Classroom ID"
// @Success 200 {object} domain.Classroom
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /classrooms/{id} [get]
func (h *Handler) GetClassroomByID(c *gin.Context) {
	classroomID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		respondWithError(h.logger, c, http.StatusBadRequest, "Invalid classroom ID")
		return
	}

	classroom, err := h.services.ClassroomService.GetByID(c.Request.Context(), classroomID)
	if err != nil {
		respondWithError(h.logger, c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, classroom)
}

// GetAllClassrooms godoc
// @Summary Get all classrooms
// @Description Get a list of all classrooms
// @Tags Classrooms
// @Accept json
// @Produce json
// @Success 200 {array} domain.Classroom
// @Failure 500 {object} ErrorResponse
// @Router /classrooms [get]
func (h *Handler) GetAllClassrooms(c *gin.Context) {
	classrooms, err := h.services.ClassroomService.GetAll(c.Request.Context())
	if err != nil {
		respondWithError(h.logger, c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, classrooms)
}
