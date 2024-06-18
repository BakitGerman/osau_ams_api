package handler

import (
	"net/http"
	"strconv"

	"github.com/BeRebornBng/OsauAmsApi/domain"
	"github.com/gin-gonic/gin"
)

type CreateProfileRequest struct {
	ProfileName     string `json:"profile_name" validate:"required,customfieldrusnumregex"`
	SpecialtyCode   string `json:"specialty_code" validate:"required,customspecialtycoderegex"`
	EducationTypeID int64  `json:"education_type_id" validate:"required,numeric"`
}

type PutProfileRequest struct {
	ProfileID       int64  `json:"profile_id" validate:"required,numeric"`
	ProfileName     string `json:"profile_name" validate:"required,customfieldrusnumregex"`
	SpecialtyCode   string `json:"specialty_code" validate:"required,customspecialtycoderegex"`
	EducationTypeID int64  `json:"education_type_id" validate:"required,numeric"`
}

type PatchProfileRequest struct {
	ProfileID       int64  `json:"profile_id" validate:"required,numeric"`
	ProfileName     string `json:"profile_name" validate:"omitempty,customfieldrusnumregex"`
	SpecialtyCode   string `json:"specialty_code" validate:"omitempty,customspecialtycoderegex"`
	EducationTypeID int64  `json:"education_type_id" validate:"omitempty,numeric"`
}

// CreateProfile godoc
// @Summary Create a profile
// @Description Create a new profile
// @Tags Profiles
// @Accept json
// @Produce json
// @Param profile body CreateProfileRequest true "Profile info"
// @Success 201 {object} SuccessResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /profiles [post]
func (h *Handler) CreateProfile(c *gin.Context) {
	var req CreateProfileRequest
	if err := c.BindJSON(&req); err != nil {
		respondWithError(h.logger, c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.validate.Struct(req); err != nil {
		respondWithError(h.logger, c, http.StatusBadRequest, err.Error())
		return
	}

	profile := domain.Profile{
		ProfileName:     req.ProfileName,
		SpecialtyCode:   req.SpecialtyCode,
		EducationTypeID: req.EducationTypeID,
	}

	err := h.services.ProfileService.Create(c.Request.Context(), profile)
	if err != nil {
		respondWithError(h.logger, c, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithSuccess(c, http.StatusCreated, "Profile created successfully")
}

// PutProfile godoc
// @Summary Update a profile
// @Description Update an existing profile
// @Tags Profiles
// @Accept json
// @Produce json
// @Param profile body PutProfileRequest true "Profile info"
// @Success 200 {object} SuccessResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /profiles [put]
func (h *Handler) PutProfile(c *gin.Context) {
	var req PutProfileRequest
	if err := c.BindJSON(&req); err != nil {
		respondWithError(h.logger, c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.validate.Struct(req); err != nil {
		respondWithError(h.logger, c, http.StatusBadRequest, err.Error())
		return
	}

	profile := domain.Profile{
		ProfileID:       req.ProfileID,
		ProfileName:     req.ProfileName,
		SpecialtyCode:   req.SpecialtyCode,
		EducationTypeID: req.EducationTypeID,
	}

	err := h.services.ProfileService.Put(c.Request.Context(), profile)
	if err != nil {
		respondWithError(h.logger, c, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithSuccess(c, http.StatusOK, "Profile updated successfully")
}

// PatchProfile godoc
// @Summary Partially update a profile
// @Description Partially update an existing profile
// @Tags Profiles
// @Accept json
// @Produce json
// @Param profile body PatchProfileRequest true "Profile info"
// @Success 200 {object} SuccessResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /profiles [patch]
func (h *Handler) PatchProfile(c *gin.Context) {
	var req PatchProfileRequest
	if err := c.BindJSON(&req); err != nil {
		respondWithError(h.logger, c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.validate.Struct(req); err != nil {
		respondWithError(h.logger, c, http.StatusBadRequest, err.Error())
		return
	}

	profile := domain.Profile{
		ProfileID:       req.ProfileID,
		ProfileName:     req.ProfileName,
		SpecialtyCode:   req.SpecialtyCode,
		EducationTypeID: req.EducationTypeID,
	}

	err := h.services.ProfileService.Patch(c.Request.Context(), profile)
	if err != nil {
		respondWithError(h.logger, c, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithSuccess(c, http.StatusOK, "Profile patched successfully")
}

// DeleteProfile godoc
// @Summary Delete a profile
// @Description Delete an existing profile
// @Tags Profiles
// @Accept json
// @Produce json
// @Param id path int64 true "Profile ID"
// @Success 200 {object} SuccessResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /profiles/{id} [delete]
func (h *Handler) DeleteProfile(c *gin.Context) {
	profileID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		respondWithError(h.logger, c, http.StatusBadRequest, "Invalid profile ID")
		return
	}

	err = h.services.ProfileService.Delete(c.Request.Context(), profileID)
	if err != nil {
		respondWithError(h.logger, c, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithSuccess(c, http.StatusOK, "Profile deleted successfully")
}

// GetProfileByID godoc
// @Summary Get a profile by ID
// @Description Get a profile by its ID
// @Tags Profiles
// @Accept json
// @Produce json
// @Param id path int64 true "Profile ID"
// @Success 200 {object} domain.ProfileInfo
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /profiles/{id} [get]
func (h *Handler) GetProfileByID(c *gin.Context) {
	profileID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		respondWithError(h.logger, c, http.StatusBadRequest, "Invalid profile ID")
		return
	}

	profile, err := h.services.ProfileService.GetByID(c.Request.Context(), profileID)
	if err != nil {
		respondWithError(h.logger, c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, profile)
}

// GetProfileByName godoc
// @Summary Get a profile by name
// @Description Get a profile by its name
// @Tags Profiles
// @Accept json
// @Produce json
// @Param name path string true "Profile Name"
// @Success 200 {object} domain.ProfileInfo
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /profiles/name/{name} [get]
func (h *Handler) GetProfileByName(c *gin.Context) {
	profileName := c.Param("name")

	profile, err := h.services.ProfileService.GetByName(c.Request.Context(), profileName)
	if err != nil {
		respondWithError(h.logger, c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, profile)
}

// GetAllProfiles godoc
// @Summary Get all profiles
// @Description Get a list of all profiles
// @Tags Profiles
// @Accept json
// @Produce json
// @Success 200 {array} domain.ProfileInfo
// @Failure 500 {object} ErrorResponse
// @Router /profiles [get]
func (h *Handler) GetAllProfiles(c *gin.Context) {
	profiles, err := h.services.ProfileService.GetAll(c.Request.Context())
	if err != nil {
		respondWithError(h.logger, c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, profiles)
}

// GetProfilesBySpecialtyCode godoc
// @Summary Get profiles by specialty code
// @Description Get a list of profiles by specialty code
// @Tags Profiles
// @Accept json
// @Produce json
// @Param specialty_code path string true "Specialty Code"
// @Success 200 {array} domain.ProfileInfo
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /profiles/specialty/{specialty_code} [get]
func (h *Handler) GetProfilesBySpecialtyCode(c *gin.Context) {
	specialtyCode := c.Param("specialty_code")

	profiles, err := h.services.ProfileService.GetAllBySpecialtyCode(c.Request.Context(), specialtyCode)
	if err != nil {
		respondWithError(h.logger, c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, profiles)
}

// GetProfilesByEducationTypeID godoc
// @Summary Get profiles by education type ID
// @Description Get a list of profiles by education type ID
// @Tags Profiles
// @Accept json
// @Produce json
// @Param education_type_id path int64 true "Education Type ID"
// @Success 200 {array} domain.ProfileInfo
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /profiles/education_type/{education_type_id} [get]
func (h *Handler) GetProfilesByEducationTypeID(c *gin.Context) {
	educationTypeID, err := strconv.ParseInt(c.Param("education_type_id"), 10, 64)
	if err != nil {
		respondWithError(h.logger, c, http.StatusBadRequest, "Invalid education type ID")
		return
	}

	profiles, err := h.services.ProfileService.GetByEducationTypeID(c.Request.Context(), educationTypeID)
	if err != nil {
		respondWithError(h.logger, c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, profiles)
}
