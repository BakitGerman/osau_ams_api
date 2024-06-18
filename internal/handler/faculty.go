package handler

import (
	"net/http"
	"strconv"

	"github.com/BeRebornBng/OsauAmsApi/domain"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type CreateFacultyRequest struct {
	FacultyName    string `json:"faculty_name" validate:"required,min=2,max=100,customfieldrusregex"`
	UniversityID   int64  `json:"university_id" validate:"required,numeric,min=1"`
	HeadLastName   string `json:"head_last_name" validate:"required,min=2,max=60,customfieldrusregex"`
	HeadFirstName  string `json:"head_first_name" validate:"required,min=2,max=60,customfieldrusregex"`
	HeadMiddleName string `json:"head_middle_name" validate:"required,min=2,max=60,customfieldrusregex"`
	FacultyEmail   string `json:"faculty_email" validate:"required,min=8,max=60,email"`
}

type PutFacultyRequest struct {
	FacultyID      int64  `json:"faculty_id" validate:"required,numeric,min=1"`
	FacultyName    string `json:"faculty_name" validate:"required,min=2,max=60,customfieldrusregex"`
	UniversityID   int64  `json:"university_id" validate:"required,numeric,min=1"`
	HeadLastName   string `json:"head_last_name" validate:"required,min=2,max=60,customfieldrusregex"`
	HeadFirstName  string `json:"head_first_name" validate:"required,min=2,max=60,customfieldrusregex"`
	HeadMiddleName string `json:"head_middle_name" validate:"required,min=2,max=60,customfieldrusregex"`
	FacultyEmail   string `json:"faculty_email" validate:"required,min=8,max=60,email"`
}

type PatchFacultyRequest struct {
	FacultyID      int64  `json:"faculty_id" validate:"required,numeric,min=1"`
	FacultyName    string `json:"faculty_name" validate:"required,min=2,max=60,customfieldrusregex"`
	UniversityID   int64  `json:"university_id" validate:"omitempty,numeric,min=1"`
	HeadLastName   string `json:"head_last_name" validate:"omitempty,min=2,max=60,customfieldrusregex"`
	HeadFirstName  string `json:"head_first_name" validate:"omitempty,min=2,max=60,customfieldrusregex"`
	HeadMiddleName string `json:"head_middle_name" validate:"omitempty,min=2,max=60,customfieldrusregex"`
	FacultyEmail   string `json:"faculty_email" validate:"omitempty,min=8,max=60,email"`
}

// CreateFaculty godoc
// @Security ApiKeyAuth
// @Summary Create a faculty
// @Description Create a new faculty
// @Tags Faculties
// @Accept json
// @Produce json
// @Param faculty body CreateFacultyRequest true "Faculty info"
// @Success 201 {object} SuccessResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /faculties [post]
func (h *Handler) CreateFaculty(c *gin.Context) {
	var req CreateFacultyRequest
	if err := c.BindJSON(&req); err != nil {
		respondWithError(h.logger, c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.validate.Struct(req); err != nil {
		errs := translateValidationErrors(err.(validator.ValidationErrors), h.translator)
		respondWithError(h.logger, c, http.StatusBadRequest, errs[0])
		return
	}

	faculty := domain.Faculty{
		FacultyName:    req.FacultyName,
		UniversityID:   req.UniversityID,
		HeadLastName:   req.HeadLastName,
		HeadFirstName:  req.HeadFirstName,
		HeadMiddleName: req.HeadMiddleName,
		FacultyEmail:   req.FacultyEmail,
	}

	err := h.services.FacultyService.Create(c.Request.Context(), faculty)
	if err != nil {
		respondWithError(h.logger, c, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithSuccess(c, http.StatusCreated, "Faculty created successfully")
}

// PutFaculty godoc
// @Security ApiKeyAuth
// @Summary Update a faculty
// @Description Update an existing faculty
// @Tags Faculties
// @Accept json
// @Produce json
// @Param faculty body PutFacultyRequest true "Faculty info"
// @Success 200 {object} SuccessResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /faculties [put]
func (h *Handler) PutFaculty(c *gin.Context) {
	var req PutFacultyRequest
	if err := c.BindJSON(&req); err != nil {
		respondWithError(h.logger, c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.validate.Struct(req); err != nil {
		errs := translateValidationErrors(err.(validator.ValidationErrors), h.translator)
		respondWithError(h.logger, c, http.StatusBadRequest, errs[0])
		return
	}

	faculty := domain.Faculty{
		FacultyID:      req.FacultyID,
		FacultyName:    req.FacultyName,
		UniversityID:   req.UniversityID,
		HeadLastName:   req.HeadLastName,
		HeadFirstName:  req.HeadFirstName,
		HeadMiddleName: req.HeadMiddleName,
		FacultyEmail:   req.FacultyEmail,
	}

	err := h.services.FacultyService.Put(c.Request.Context(), faculty)
	if err != nil {
		respondWithError(h.logger, c, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithSuccess(c, http.StatusOK, "Faculty updated successfully")
}

// PatchFaculty godoc
// @Security ApiKeyAuth
// @Summary Partially update a faculty
// @Description Partially update an existing faculty
// @Tags Faculties
// @Accept json
// @Produce json
// @Param faculty body PatchFacultyRequest true "Faculty info"
// @Success 200 {object} SuccessResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /faculties [patch]
func (h *Handler) PatchFaculty(c *gin.Context) {
	var req PatchFacultyRequest
	if err := c.BindJSON(&req); err != nil {
		respondWithError(h.logger, c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.validate.Struct(req); err != nil {
		errs := translateValidationErrors(err.(validator.ValidationErrors), h.translator)
		respondWithError(h.logger, c, http.StatusBadRequest, errs[0])
		return
	}

	faculty := domain.Faculty{
		FacultyID:      req.FacultyID,
		FacultyName:    req.FacultyName,
		UniversityID:   req.UniversityID,
		HeadLastName:   req.HeadLastName,
		HeadFirstName:  req.HeadFirstName,
		HeadMiddleName: req.HeadMiddleName,
		FacultyEmail:   req.FacultyEmail,
	}

	err := h.services.FacultyService.Patch(c.Request.Context(), faculty)
	if err != nil {
		respondWithError(h.logger, c, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithSuccess(c, http.StatusOK, "Faculty patched successfully")
}

// DeleteFaculty godoc
// @Security ApiKeyAuth
// @Summary Delete a faculty
// @Description Delete an existing faculty
// @Tags Faculties
// @Accept json
// @Produce json
// @Param id path int64 true "Faculty ID"
// @Success 200 {object} SuccessResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /faculties/{id} [delete]
func (h *Handler) DeleteFaculty(c *gin.Context) {
	facultyID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		respondWithError(h.logger, c, http.StatusBadRequest, "Invalid faculty ID")
		return
	}

	err = h.services.FacultyService.Delete(c.Request.Context(), facultyID)
	if err != nil {
		respondWithError(h.logger, c, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithSuccess(c, http.StatusOK, "Faculty deleted successfully")
}

// GetFacultyByID godoc
// @Security ApiKeyAuth
// @Summary Get a faculty by ID
// @Description Get a faculty by its ID
// @Tags Faculties
// @Accept json
// @Produce json
// @Param id path int64 true "Faculty ID"
// @Success 200 {object} domain.Faculty
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /faculties/id/{id} [get]
func (h *Handler) GetFacultyByID(c *gin.Context) {
	facultyID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		respondWithError(h.logger, c, http.StatusBadRequest, "Invalid faculty ID")
		return
	}

	faculty, err := h.services.FacultyService.GetByID(c.Request.Context(), facultyID)
	if err != nil {
		respondWithError(h.logger, c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, faculty)
}

// GetFacultyByName godoc
// @Security ApiKeyAuth
// @Summary Get a faculty by ID
// @Description Get a faculty by its ID
// @Tags Faculties
// @Accept json
// @Produce json
// @Param id path int64 true "Faculty ID"
// @Success 200 {object} domain.Faculty
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /faculties/name/{name} [get]
func (h *Handler) GetFacultyByName(c *gin.Context) {
	facultyName := c.Param("name")
	// if err != nil {
	// 	respondWithError(h.logger, c, http.StatusBadRequest, "Invalid faculty ID")
	// 	return
	// }

	faculty, err := h.services.FacultyService.GetByName(c.Request.Context(), facultyName)
	if err != nil {
		respondWithError(h.logger, c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, faculty)
}

// GetAllFaculties godoc
// @Security ApiKeyAuth
// @Summary Get all faculties
// @Description Get a list of all faculties
// @Tags Faculties
// @Accept json
// @Produce json
// @Success 200 {array} domain.Faculty
// @Failure 500 {object} ErrorResponse
// @Router /faculties [get]
func (h *Handler) GetAllFaculties(c *gin.Context) {
	faculties, err := h.services.FacultyService.GetAll(c.Request.Context())
	if err != nil {
		respondWithError(h.logger, c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, faculties)
}

// GetFacultiesByUniversityID godoc
// @Security ApiKeyAuth
// @Summary Get faculties by university ID
// @Description Get a list of faculties by university ID
// @Tags Faculties
// @Accept json
// @Produce json
// @Param university_id path int64 true "University ID"
// @Success 200 {array} domain.Faculty
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /faculties/university/id/{university_id} [get]
func (h *Handler) GetFacultiesByUniversityID(c *gin.Context) {
	universityID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		respondWithError(h.logger, c, http.StatusBadRequest, "Invalid university ID")
		return
	}

	faculties, err := h.services.FacultyService.GetAllByUniversityID(c.Request.Context(), universityID)
	if err != nil {
		respondWithError(h.logger, c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, faculties)
}
