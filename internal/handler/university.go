package handler

import (
	"net/http"
	"strconv"

	"github.com/BeRebornBng/OsauAmsApi/domain"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type CreateUniversityRequest struct {
	UniversityName  string `json:"university_name" validate:"required,min=2,max=100,customfieldrusregex"`
	HeadLastName    string `json:"head_last_name" validate:"required,min=2,max=60,customfieldrusregex"`
	HeadFirstName   string `json:"head_first_name" validate:"required,min=2,max=60,customfieldrusregex"`
	HeadMiddleName  string `json:"head_middle_name" validate:"required,min=2,max=60,customfieldrusregex"`
	UniversityEmail string `json:"university_email" validate:"required,min=8,max=60,email"`
}

type PutUniversityRequest struct {
	UniversityID    int64  `json:"university_id" validate:"required,numeric"`
	UniversityName  string `json:"university_name" validate:"required,min=2,max=100,customfieldrusregex"`
	HeadLastName    string `json:"head_last_name" validate:"required,min=2,max=60,customfieldrusregex"`
	HeadFirstName   string `json:"head_first_name" validate:"required,min=2,max=60,customfieldrusregex"`
	HeadMiddleName  string `json:"head_middle_name" validate:"required,min=2,max=60,customfieldrusregex"`
	UniversityEmail string `json:"university_email" validate:"required,min=8,max=60,email"`
}

type PatchUniversityRequest struct {
	UniversityID    int64  `json:"university_id" validate:"required,numeric"`
	UniversityName  string `json:"university_name" validate:"omitempty,min=2,max=100,customfieldrusregex"`
	HeadLastName    string `json:"head_last_name" validate:"omitempty,min=2,max=60,customfieldrusregex"`
	HeadFirstName   string `json:"head_first_name" validate:"omitempty,min=2,max=60,customfieldrusregex"`
	HeadMiddleName  string `json:"head_middle_name" validate:"omitempty,min=2,max=60,customfieldrusregex"`
	UniversityEmail string `json:"university_email" validate:"omitempty,min=8,max=60,email"`
}

// CreateUniversity godoc
// @Security ApiKeyAuth
// @Summary Create a university
// @Description Create a new university
// @Tags Universities
// @Accept json
// @Produce json
// @Param university body CreateUniversityRequest true "University info"
// @Success 201 {object} SuccessResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /admins/universities [post]
func (h *Handler) CreateUniversity(c *gin.Context) {
	var req CreateUniversityRequest
	if err := c.BindJSON(&req); err != nil {
		respondWithError(h.logger, c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.validate.Struct(req); err != nil {
		errs := translateValidationErrors(err.(validator.ValidationErrors), h.translator)
		respondWithError(h.logger, c, http.StatusBadRequest, errs[0])
		return
	}

	university := domain.University{
		UniversityName:  req.UniversityName,
		HeadLastName:    req.HeadLastName,
		HeadFirstName:   req.HeadFirstName,
		HeadMiddleName:  req.HeadMiddleName,
		UniversityEmail: req.UniversityEmail,
	}

	err := h.services.UniversityService.Create(c.Request.Context(), university)
	if err != nil {
		respondWithError(h.logger, c, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithSuccess(c, http.StatusCreated, "University created successfully")
}

// PutUniversity godoc
// @Security ApiKeyAuth
// @Summary Update a university
// @Description Update an existing university
// @Tags Universities
// @Accept json
// @Produce json
// @Param university body PutUniversityRequest true "University info"
// @Success 200 {object} SuccessResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /admins/universities [put]
func (h *Handler) PutUniversity(c *gin.Context) {
	var req PutUniversityRequest
	if err := c.BindJSON(&req); err != nil {
		respondWithError(h.logger, c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.validate.Struct(req); err != nil {
		errs := translateValidationErrors(err.(validator.ValidationErrors), h.translator)
		respondWithError(h.logger, c, http.StatusBadRequest, errs[0])
		return
	}

	university := domain.University{
		UniversityID:    req.UniversityID,
		UniversityName:  req.UniversityName,
		HeadLastName:    req.HeadLastName,
		HeadFirstName:   req.HeadFirstName,
		HeadMiddleName:  req.HeadMiddleName,
		UniversityEmail: req.UniversityEmail,
	}

	err := h.services.UniversityService.Put(c.Request.Context(), university)
	if err != nil {
		respondWithError(h.logger, c, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithSuccess(c, http.StatusOK, "University updated successfully")
}

// PatchUniversity godoc
// @Security ApiKeyAuth
// @Summary Partially update a university
// @Description Partially update an existing university
// @Tags Universities
// @Accept json
// @Produce json
// @Param university body PatchUniversityRequest true "University info"
// @Success 200 {object} SuccessResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /admins/universities [patch]
func (h *Handler) PatchUniversity(c *gin.Context) {
	var req PatchUniversityRequest
	if err := c.BindJSON(&req); err != nil {
		respondWithError(h.logger, c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.validate.Struct(req); err != nil {
		errs := translateValidationErrors(err.(validator.ValidationErrors), h.translator)
		respondWithError(h.logger, c, http.StatusBadRequest, errs[0])
		return
	}

	university := domain.University{
		UniversityID:    req.UniversityID,
		UniversityName:  req.UniversityName,
		HeadLastName:    req.HeadLastName,
		HeadFirstName:   req.HeadFirstName,
		HeadMiddleName:  req.HeadMiddleName,
		UniversityEmail: req.UniversityEmail,
	}

	err := h.services.UniversityService.Patch(c.Request.Context(), university)
	if err != nil {
		respondWithError(h.logger, c, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithSuccess(c, http.StatusOK, "University patched successfully")
}

// DeleteUniversity godoc
// @Security ApiKeyAuth
// @Summary Delete a university
// @Description Delete an existing university
// @Tags Universities
// @Accept json
// @Produce json
// @Param id path int64 true "University ID"
// @Success 200 {object} SuccessResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /admins/universities/{id} [delete]
func (h *Handler) DeleteUniversity(c *gin.Context) {
	universityID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		respondWithError(h.logger, c, http.StatusBadRequest, "Invalid university ID")
		return
	}

	err = h.services.UniversityService.Delete(c.Request.Context(), universityID)
	if err != nil {
		respondWithError(h.logger, c, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithSuccess(c, http.StatusOK, "University deleted successfully")
}

// GetUniversityByID godoc
// @Security ApiKeyAuth
// @Summary Get a university by ID
// @Description Get a university by its ID
// @Tags Universities
// @Accept json
// @Produce json
// @Param id path int64 true "University ID"
// @Success 200 {object} domain.University
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /admins/universities/{id} [get]
func (h *Handler) GetUniversityByID(c *gin.Context) {
	universityID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		respondWithError(h.logger, c, http.StatusBadRequest, "Invalid university ID")
		return
	}

	university, err := h.services.UniversityService.GetByID(c.Request.Context(), universityID)
	if err != nil {
		respondWithError(h.logger, c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, university)
}

// GetUniversityByName godoc
// @Security ApiKeyAuth
// @Summary Get a university by name
// @Description Get a university by its name
// @Tags Universities
// @Accept json
// @Produce json
// @Param name path string true "University Name"
// @Success 200 {object} domain.University
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /admins/universities/name/{name} [get]
func (h *Handler) GetUniversityByName(c *gin.Context) {
	universityName := c.Param("name")

	university, err := h.services.UniversityService.GetByName(c.Request.Context(), universityName)
	if err != nil {
		respondWithError(h.logger, c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, university)
}

// GetAllUniversities godoc
// @Security ApiKeyAuth
// @Summary Get all universities
// @Description Get a list of all universities
// @Tags Universities
// @Accept json
// @Produce json
// @Success 200 {array} domain.University
// @Failure 500 {object} ErrorResponse
// @Router /admins/universities [get]
func (h *Handler) GetAllUniversities(c *gin.Context) {
	universities, err := h.services.UniversityService.GetAll(c.Request.Context())
	if err != nil {
		respondWithError(h.logger, c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, universities)
}
