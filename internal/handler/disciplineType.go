package handler

import (
	"net/http"
	"strconv"

	"github.com/BeRebornBng/OsauAmsApi/domain"
	"github.com/gin-gonic/gin"
)

type CreateDisciplineTypeRequest struct {
	DisciplineTypeName string `json:"discipline_type_name" validate:"required,customfieldrusnumregex"`
}

type PutDisciplineTypeRequest struct {
	DisciplineTypeID   int64  `json:"discipline_type_id" validate:"required,numeric"`
	DisciplineTypeName string `json:"discipline_type_name" validate:"required,customfieldrusnumregex"`
}

type PatchDisciplineTypeRequest struct {
	DisciplineTypeID   int64  `json:"discipline_type_id" validate:"required,numeric"`
	DisciplineTypeName string `json:"discipline_type_name" validate:"omitempty,customfieldrusnumregex"`
}

// CreateDisciplineType godoc
// @Summary Create a discipline type
// @Description Create a new discipline type
// @Tags DisciplineTypes
// @Accept json
// @Produce json
// @Param discipline_type body CreateDisciplineTypeRequest true "Discipline type info"
// @Success 201 {object} SuccessResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /discipline_types [post]
func (h *Handler) CreateDisciplineType(c *gin.Context) {
	var req CreateDisciplineTypeRequest
	if err := c.BindJSON(&req); err != nil {
		respondWithError(h.logger, c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.validate.Struct(req); err != nil {
		respondWithError(h.logger, c, http.StatusBadRequest, err.Error())
		return
	}

	disciplineType := domain.DisciplineType{
		DisciplineTypeName: req.DisciplineTypeName,
	}

	err := h.services.DisciplineTypeService.Create(c.Request.Context(), disciplineType)
	if err != nil {
		respondWithError(h.logger, c, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithSuccess(c, http.StatusCreated, "Discipline type created successfully")
}

// PutDisciplineType godoc
// @Summary Update a discipline type
// @Description Update an existing discipline type
// @Tags DisciplineTypes
// @Accept json
// @Produce json
// @Param discipline_type body PutDisciplineTypeRequest true "Discipline type info"
// @Success 200 {object} SuccessResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /discipline_types [put]
func (h *Handler) PutDisciplineType(c *gin.Context) {
	var req PutDisciplineTypeRequest
	if err := c.BindJSON(&req); err != nil {
		respondWithError(h.logger, c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.validate.Struct(req); err != nil {
		respondWithError(h.logger, c, http.StatusBadRequest, err.Error())
		return
	}

	disciplineType := domain.DisciplineType{
		DisciplineTypeID:   req.DisciplineTypeID,
		DisciplineTypeName: req.DisciplineTypeName,
	}

	err := h.services.DisciplineTypeService.Put(c.Request.Context(), disciplineType)
	if err != nil {
		respondWithError(h.logger, c, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithSuccess(c, http.StatusOK, "Discipline type updated successfully")
}

// PatchDisciplineType godoc
// @Summary Partially update a discipline type
// @Description Partially update an existing discipline type
// @Tags DisciplineTypes
// @Accept json
// @Produce json
// @Param discipline_type body PatchDisciplineTypeRequest true "Discipline type info"
// @Success 200 {object} SuccessResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /discipline_types [patch]
func (h *Handler) PatchDisciplineType(c *gin.Context) {
	var req PatchDisciplineTypeRequest
	if err := c.BindJSON(&req); err != nil {
		respondWithError(h.logger, c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.validate.Struct(req); err != nil {
		respondWithError(h.logger, c, http.StatusBadRequest, err.Error())
		return
	}

	disciplineType := domain.DisciplineType{
		DisciplineTypeID:   req.DisciplineTypeID,
		DisciplineTypeName: req.DisciplineTypeName,
	}

	err := h.services.DisciplineTypeService.Patch(c.Request.Context(), disciplineType)
	if err != nil {
		respondWithError(h.logger, c, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithSuccess(c, http.StatusOK, "Discipline type patched successfully")
}

// DeleteDisciplineType godoc
// @Summary Delete a discipline type
// @Description Delete an existing discipline type
// @Tags DisciplineTypes
// @Accept json
// @Produce json
// @Param id path int64 true "Discipline type ID"
// @Success 200 {object} SuccessResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /discipline_types/{id} [delete]
func (h *Handler) DeleteDisciplineType(c *gin.Context) {
	disciplineTypeID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		respondWithError(h.logger, c, http.StatusBadRequest, "Invalid discipline type ID")
		return
	}

	err = h.services.DisciplineTypeService.Delete(c.Request.Context(), disciplineTypeID)
	if err != nil {
		respondWithError(h.logger, c, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithSuccess(c, http.StatusOK, "Discipline type deleted successfully")
}

// GetDisciplineTypeByID godoc
// @Summary Get a discipline type by ID
// @Description Get a discipline type by its ID
// @Tags DisciplineTypes
// @Accept json
// @Produce json
// @Param id path int64 true "Discipline type ID"
// @Success 200 {object} domain.DisciplineType
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /discipline_types/{id} [get]
func (h *Handler) GetDisciplineTypeByID(c *gin.Context) {
	disciplineTypeID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		respondWithError(h.logger, c, http.StatusBadRequest, "Invalid discipline type ID")
		return
	}

	disciplineType, err := h.services.DisciplineTypeService.GetByID(c.Request.Context(), disciplineTypeID)
	if err != nil {
		respondWithError(h.logger, c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, disciplineType)
}

// GetAllDisciplineTypes godoc
// @Summary Get all discipline types
// @Description Get a list of all discipline types
// @Tags DisciplineTypes
// @Accept json
// @Produce json
// @Success 200 {array} domain.DisciplineType
// @Failure 500 {object} ErrorResponse
// @Router /discipline_types [get]
func (h *Handler) GetAllDisciplineTypes(c *gin.Context) {
	disciplineTypes, err := h.services.DisciplineTypeService.GetAll(c.Request.Context())
	if err != nil {
		respondWithError(h.logger, c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, disciplineTypes)
}
