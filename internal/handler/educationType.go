package handler

import (
	"net/http"
	"strconv"

	"github.com/BeRebornBng/OsauAmsApi/domain"
	"github.com/gin-gonic/gin"
)

type CreateEducationTypeRequest struct {
	EducationTypeName string `json:"education_type_name" validate:"required,customfieldrusnumregex"`
}

type PutEducationTypeRequest struct {
	EducationTypeID   int64  `json:"education_type_id" validate:"required,numeric"`
	EducationTypeName string `json:"education_type_name" validate:"required,customfieldrusnumregex"`
}

type PatchEducationTypeRequest struct {
	EducationTypeID   int64  `json:"education_type_id" validate:"required,numeric"`
	EducationTypeName string `json:"education_type_name" validate:"omitempty,customfieldrusnumregex"`
}

// CreateEducationType godoc
// @Summary Create an education type
// @Description Create a new education type
// @Tags EducationTypes
// @Accept json
// @Produce json
// @Param education_type body CreateEducationTypeRequest true "Education type info"
// @Success 201 {object} SuccessResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /education_types [post]
func (h *Handler) CreateEducationType(c *gin.Context) {
	var req CreateEducationTypeRequest
	if err := c.BindJSON(&req); err != nil {
		respondWithError(h.logger, c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.validate.Struct(req); err != nil {
		respondWithError(h.logger, c, http.StatusBadRequest, err.Error())
		return
	}

	educationType := domain.EducationType{
		EducationTypeName: req.EducationTypeName,
	}

	err := h.services.EducationTypeService.Create(c.Request.Context(), educationType)
	if err != nil {
		respondWithError(h.logger, c, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithSuccess(c, http.StatusCreated, "Education type created successfully")
}

// PutEducationType godoc
// @Summary Update an education type
// @Description Update an existing education type
// @Tags EducationTypes
// @Accept json
// @Produce json
// @Param education_type body PutEducationTypeRequest true "Education type info"
// @Success 200 {object} SuccessResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /education_types [put]
func (h *Handler) PutEducationType(c *gin.Context) {
	var req PutEducationTypeRequest
	if err := c.BindJSON(&req); err != nil {
		respondWithError(h.logger, c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.validate.Struct(req); err != nil {
		respondWithError(h.logger, c, http.StatusBadRequest, err.Error())
		return
	}

	educationType := domain.EducationType{
		EducationTypeID:   req.EducationTypeID,
		EducationTypeName: req.EducationTypeName,
	}

	err := h.services.EducationTypeService.Put(c.Request.Context(), educationType)
	if err != nil {
		respondWithError(h.logger, c, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithSuccess(c, http.StatusOK, "Education type updated successfully")
}

// PatchEducationType godoc
// @Summary Partially update an education type
// @Description Partially update an existing education type
// @Tags EducationTypes
// @Accept json
// @Produce json
// @Param education_type body PatchEducationTypeRequest true "Education type info"
// @Success 200 {object} SuccessResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /education_types [patch]
func (h *Handler) PatchEducationType(c *gin.Context) {
	var req PatchEducationTypeRequest
	if err := c.BindJSON(&req); err != nil {
		respondWithError(h.logger, c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.validate.Struct(req); err != nil {
		respondWithError(h.logger, c, http.StatusBadRequest, err.Error())
		return
	}

	educationType := domain.EducationType{
		EducationTypeID:   req.EducationTypeID,
		EducationTypeName: req.EducationTypeName,
	}

	err := h.services.EducationTypeService.Patch(c.Request.Context(), educationType)
	if err != nil {
		respondWithError(h.logger, c, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithSuccess(c, http.StatusOK, "Education type patched successfully")
}

// DeleteEducationType godoc
// @Summary Delete an education type
// @Description Delete an existing education type
// @Tags EducationTypes
// @Accept json
// @Produce json
// @Param id path int64 true "Education type ID"
// @Success 200 {object} SuccessResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /education_types/{id} [delete]
func (h *Handler) DeleteEducationType(c *gin.Context) {
	educationTypeID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		respondWithError(h.logger, c, http.StatusBadRequest, "Invalid education type ID")
		return
	}

	err = h.services.EducationTypeService.Delete(c.Request.Context(), educationTypeID)
	if err != nil {
		respondWithError(h.logger, c, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithSuccess(c, http.StatusOK, "Education type deleted successfully")
}

// GetEducationTypeByID godoc
// @Summary Get an education type by ID
// @Description Get an education type by its ID
// @Tags EducationTypes
// @Accept json
// @Produce json
// @Param id path int64 true "Education type ID"
// @Success 200 {object} domain.EducationType
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /education_types/{id} [get]
func (h *Handler) GetEducationTypeByID(c *gin.Context) {
	educationTypeID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		respondWithError(h.logger, c, http.StatusBadRequest, "Invalid education type ID")
		return
	}

	educationType, err := h.services.EducationTypeService.GetByID(c.Request.Context(), educationTypeID)
	if err != nil {
		respondWithError(h.logger, c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, educationType)
}

// GetEducationTypeByName godoc
// @Summary Get an education type by name
// @Description Get an education type by its name
// @Tags EducationTypes
// @Accept json
// @Produce json
// @Param name path string true "Education type Name"
// @Success 200 {object} domain.EducationType
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /education_types/name/{name} [get]
func (h *Handler) GetEducationTypeByName(c *gin.Context) {
	educationTypeName := c.Param("name")

	educationType, err := h.services.EducationTypeService.GetByName(c.Request.Context(), educationTypeName)
	if err != nil {
		respondWithError(h.logger, c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, educationType)
}

// GetAllEducationTypes godoc
// @Summary Get all education types
// @Description Get a list of all education types
// @Tags EducationTypes
// @Accept json
// @Produce json
// @Success 200 {array} domain.EducationType
// @Failure 500 {object} ErrorResponse
// @Router /education_types [get]
func (h *Handler) GetAllEducationTypes(c *gin.Context) {
	educationTypes, err := h.services.EducationTypeService.GetAll(c.Request.Context())
	if err != nil {
		respondWithError(h.logger, c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, educationTypes)
}
