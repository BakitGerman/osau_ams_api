package handler

import (
	"net/http"
	"strconv"

	"github.com/BeRebornBng/OsauAmsApi/domain"
	"github.com/gin-gonic/gin"
)

type CreateSpecialtyRequest struct {
	SpecialtyCode    string `json:"specialty_code" validate:"required,customspecialtycoderegex"`
	SpecialtyName    string `json:"specialty_name" validate:"required,customfieldrusnumregex"`
	DepartamentID    int64  `json:"departament_id" validate:"required,numeric"`
	EducationLevelID int64  `json:"education_level_id" validate:"required,numeric"`
}

type PutSpecialtyRequest struct {
	SpecialtyCode    string `json:"specialty_code" validate:"required,customspecialtycoderegex"`
	SpecialtyName    string `json:"specialty_name" validate:"required,customfieldrusnumregex"`
	DepartamentID    int64  `json:"departament_id" validate:"required,numeric"`
	EducationLevelID int64  `json:"education_level_id" validate:"required,numeric"`
}

type PatchSpecialtyRequest struct {
	SpecialtyCode    string `json:"specialty_code" validate:"required,customspecialtycoderegex"`
	SpecialtyName    string `json:"specialty_name" validate:"omitempty,customfieldrusnumregex"`
	DepartamentID    int64  `json:"departament_id" validate:"omitempty,numeric"`
	EducationLevelID int64  `json:"education_level_id" validate:"required,numeric"`
}

// CreateSpecialty godoc
// @Summary Create a specialty
// @Description Create a new specialty
// @Tags Specialties
// @Accept json
// @Produce json
// @Param specialty body CreateSpecialtyRequest true "Specialty info"
// @Success 201 {object} SuccessResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /specialties [post]
func (h *Handler) CreateSpecialty(c *gin.Context) {
	var req CreateSpecialtyRequest
	if err := c.BindJSON(&req); err != nil {
		respondWithError(h.logger, c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.validate.Struct(req); err != nil {
		respondWithError(h.logger, c, http.StatusBadRequest, err.Error())
		return
	}

	specialty := domain.Specialty{
		SpecialtyCode:    req.SpecialtyCode,
		SpecialtyName:    req.SpecialtyName,
		DepartamentID:    req.DepartamentID,
		EducationLevelID: req.EducationLevelID,
	}

	err := h.services.SpecialtyService.Create(c.Request.Context(), specialty)
	if err != nil {
		respondWithError(h.logger, c, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithSuccess(c, http.StatusCreated, "Specialty created successfully")
}

// PutSpecialty godoc
// @Summary Update a specialty
// @Description Update an existing specialty
// @Tags Specialties
// @Accept json
// @Produce json
// @Param specialty body PutSpecialtyRequest true "Specialty info"
// @Success 200 {object} SuccessResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /specialties [put]
func (h *Handler) PutSpecialty(c *gin.Context) {
	var req PutSpecialtyRequest
	if err := c.BindJSON(&req); err != nil {
		respondWithError(h.logger, c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.validate.Struct(req); err != nil {
		respondWithError(h.logger, c, http.StatusBadRequest, err.Error())
		return
	}

	specialty := domain.Specialty{
		SpecialtyCode:    req.SpecialtyCode,
		SpecialtyName:    req.SpecialtyName,
		DepartamentID:    req.DepartamentID,
		EducationLevelID: req.EducationLevelID,
	}

	err := h.services.SpecialtyService.Put(c.Request.Context(), specialty)
	if err != nil {
		respondWithError(h.logger, c, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithSuccess(c, http.StatusOK, "Specialty updated successfully")
}

// PatchSpecialty godoc
// @Summary Partially update a specialty
// @Description Partially update an existing specialty
// @Tags Specialties
// @Accept json
// @Produce json
// @Param specialty body PatchSpecialtyRequest true "Specialty info"
// @Success 200 {object} SuccessResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /specialties [patch]
func (h *Handler) PatchSpecialty(c *gin.Context) {
	var req PatchSpecialtyRequest
	if err := c.BindJSON(&req); err != nil {
		respondWithError(h.logger, c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.validate.Struct(req); err != nil {
		respondWithError(h.logger, c, http.StatusBadRequest, err.Error())
		return
	}

	specialty := domain.Specialty{
		SpecialtyCode:    req.SpecialtyCode,
		SpecialtyName:    req.SpecialtyName,
		DepartamentID:    req.DepartamentID,
		EducationLevelID: req.EducationLevelID,
	}

	err := h.services.SpecialtyService.Patch(c.Request.Context(), specialty)
	if err != nil {
		respondWithError(h.logger, c, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithSuccess(c, http.StatusOK, "Specialty patched successfully")
}

// DeleteSpecialty godoc
// @Summary Delete a specialty
// @Description Delete an existing specialty
// @Tags Specialties
// @Accept json
// @Produce json
// @Param code path string true "Specialty Code"
// @Success 200 {object} SuccessResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /specialties/{code} [delete]
func (h *Handler) DeleteSpecialty(c *gin.Context) {
	specialtyCode := c.Param("code")

	err := h.services.SpecialtyService.Delete(c.Request.Context(), specialtyCode)
	if err != nil {
		respondWithError(h.logger, c, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithSuccess(c, http.StatusOK, "Specialty deleted successfully")
}

// GetSpecialtyByCode godoc
// @Summary Get a specialty by code
// @Description Get a specialty by its code
// @Tags Specialties
// @Accept json
// @Produce json
// @Param code path string true "Specialty Code"
// @Success 200 {object} domain.SpecialtyInfo
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /specialties/{code} [get]
func (h *Handler) GetSpecialtyByCode(c *gin.Context) {
	specialtyCode := c.Param("code")

	specialty, err := h.services.SpecialtyService.GetByCode(c.Request.Context(), specialtyCode)
	if err != nil {
		respondWithError(h.logger, c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, specialty)
}

// GetSpecialtyByName godoc
// @Summary Get a specialty by name
// @Description Get a specialty by its name
// @Tags Specialties
// @Accept json
// @Produce json
// @Param name path string true "Specialty Name"
// @Success 200 {object} domain.SpecialtyInfo
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /specialties/name/{name} [get]
func (h *Handler) GetSpecialtyByName(c *gin.Context) {
	specialtyName := c.Param("name")

	specialty, err := h.services.SpecialtyService.GetByName(c.Request.Context(), specialtyName)
	if err != nil {
		respondWithError(h.logger, c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, specialty)
}

// GetAllSpecialties godoc
// @Summary Get all specialties
// @Description Get a list of all specialties
// @Tags Specialties
// @Accept json
// @Produce json
// @Success 200 {array} domain.SpecialtyInfo
// @Failure 500 {object} ErrorResponse
// @Router /specialties [get]
func (h *Handler) GetAllSpecialties(c *gin.Context) {
	specialties, err := h.services.SpecialtyService.GetAll(c.Request.Context())
	if err != nil {
		respondWithError(h.logger, c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, specialties)
}

// GetSpecialtiesByDepartamentID godoc
// @Summary Get specialties by departament ID
// @Description Get a list of specialties by departament ID
// @Tags Specialties
// @Accept json
// @Produce json
// @Param departament_id path int64 true "Departament ID"
// @Success 200 {array} domain.SpecialtyInfo
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /specialties/departament/{departament_id} [get]
func (h *Handler) GetSpecialtiesByDepartamentID(c *gin.Context) {
	departamentID, err := strconv.ParseInt(c.Param("departament_id"), 10, 64)
	if err != nil {
		respondWithError(h.logger, c, http.StatusBadRequest, "Invalid departament ID")
		return
	}

	specialties, err := h.services.SpecialtyService.GetAllByDepartamentID(c.Request.Context(), departamentID)
	if err != nil {
		respondWithError(h.logger, c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, specialties)
}
