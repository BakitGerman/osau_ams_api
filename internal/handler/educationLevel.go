package handler

import (
	"net/http"
	"strconv"

	"github.com/BeRebornBng/OsauAmsApi/domain"
	"github.com/gin-gonic/gin"
)

type CreateEducationLevelRequest struct {
	EducationLevelName string `json:"education_level_name" validate:"required,customfieldrusnumregex"`
}

type PutEducationLevelRequest struct {
	EducationLevelID   int64  `json:"education_level_id" validate:"required,numeric"`
	EducationLevelName string `json:"education_level_name" validate:"required,customfieldrusnumregex"`
}

type PatchEducationLevelRequest struct {
	EducationLevelID   int64  `json:"education_level_id" validate:"required,numeric"`
	EducationLevelName string `json:"education_level_name" validate:"omitempty,customfieldrusnumregex"`
}

// CreateEducationLevel godoc
// @Summary Create an education level
// @Description Create a new education level
// @Tags EducationLevels
// @Accept json
// @Produce json
// @Param education_level body CreateEducationLevelRequest true "Education level info"
// @Success 201 {object} SuccessResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /education_levels [post]
func (h *Handler) CreateEducationLevel(c *gin.Context) {
	var req CreateEducationLevelRequest
	if err := c.BindJSON(&req); err != nil {
		respondWithError(h.logger, c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.validate.Struct(req); err != nil {
		respondWithError(h.logger, c, http.StatusBadRequest, err.Error())
		return
	}

	educationLevel := domain.EducationLevel{
		EducationLevelName: req.EducationLevelName,
	}

	err := h.services.EducationLevelService.Create(c.Request.Context(), educationLevel)
	if err != nil {
		respondWithError(h.logger, c, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithSuccess(c, http.StatusCreated, "Education level created successfully")
}

// PutEducationLevel godoc
// @Summary Update an education level
// @Description Update an existing education level
// @Tags EducationLevels
// @Accept json
// @Produce json
// @Param education_level body PutEducationLevelRequest true "Education level info"
// @Success 200 {object} SuccessResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /education_levels [put]
func (h *Handler) PutEducationLevel(c *gin.Context) {
	var req PutEducationLevelRequest
	if err := c.BindJSON(&req); err != nil {
		respondWithError(h.logger, c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.validate.Struct(req); err != nil {
		respondWithError(h.logger, c, http.StatusBadRequest, err.Error())
		return
	}

	educationLevel := domain.EducationLevel{
		EducationLevelID:   req.EducationLevelID,
		EducationLevelName: req.EducationLevelName,
	}

	err := h.services.EducationLevelService.Put(c.Request.Context(), educationLevel)
	if err != nil {
		respondWithError(h.logger, c, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithSuccess(c, http.StatusOK, "Education level updated successfully")
}

// PatchEducationLevel godoc
// @Summary Partially update an education level
// @Description Partially update an existing education level
// @Tags EducationLevels
// @Accept json
// @Produce json
// @Param education_level body PatchEducationLevelRequest true "Education level info"
// @Success 200 {object} SuccessResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /education_levels [patch]
func (h *Handler) PatchEducationLevel(c *gin.Context) {
	var req PatchEducationLevelRequest
	if err := c.BindJSON(&req); err != nil {
		respondWithError(h.logger, c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.validate.Struct(req); err != nil {
		respondWithError(h.logger, c, http.StatusBadRequest, err.Error())
		return
	}

	educationLevel := domain.EducationLevel{
		EducationLevelID:   req.EducationLevelID,
		EducationLevelName: req.EducationLevelName,
	}

	err := h.services.EducationLevelService.Patch(c.Request.Context(), educationLevel)
	if err != nil {
		respondWithError(h.logger, c, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithSuccess(c, http.StatusOK, "Education level patched successfully")
}

// DeleteEducationLevel godoc
// @Summary Delete an education level
// @Description Delete an existing education level
// @Tags EducationLevels
// @Accept json
// @Produce json
// @Param id path int64 true "Education level ID"
// @Success 200 {object} SuccessResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /education_levels/{id} [delete]
func (h *Handler) DeleteEducationLevel(c *gin.Context) {
	educationLevelID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		respondWithError(h.logger, c, http.StatusBadRequest, "Invalid education level ID")
		return
	}

	err = h.services.EducationLevelService.Delete(c.Request.Context(), educationLevelID)
	if err != nil {
		respondWithError(h.logger, c, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithSuccess(c, http.StatusOK, "Education level deleted successfully")
}

// GetEducationLevelByID godoc
// @Summary Get an education level by ID
// @Description Get an education level by its ID
// @Tags EducationLevels
// @Accept json
// @Produce json
// @Param id path int64 true "Education level ID"
// @Success 200 {object} domain.EducationLevel
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /education_levels/{id} [get]
func (h *Handler) GetEducationLevelByID(c *gin.Context) {
	educationLevelID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		respondWithError(h.logger, c, http.StatusBadRequest, "Invalid education level ID")
		return
	}

	educationLevel, err := h.services.EducationLevelService.GetByID(c.Request.Context(), educationLevelID)
	if err != nil {
		respondWithError(h.logger, c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, educationLevel)
}

// GetAllEducationLevels godoc
// @Summary Get all education levels
// @Description Get a list of all education levels
// @Tags EducationLevels
// @Accept json
// @Produce json
// @Success 200 {array} domain.EducationLevel
// @Failure 500 {object} ErrorResponse
// @Router /education_levels [get]
func (h *Handler) GetAllEducationLevels(c *gin.Context) {
	educationLevels, err := h.services.EducationLevelService.GetAll(c.Request.Context())
	if err != nil {
		respondWithError(h.logger, c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, educationLevels)
}
