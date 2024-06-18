package handler

import (
	"net/http"
	"strconv"

	"github.com/BeRebornBng/OsauAmsApi/domain"
	"github.com/gin-gonic/gin"
)

type CreateDepartamentRequest struct {
	FacultyID        int64  `json:"faculty_id" validate:"required,numeric"`
	DepartamentName  string `json:"departament_name" validate:"required,customfieldrusnumregex"`
	HeadLastName     string `json:"head_last_name" validate:"required,customfieldrusnumregex"`
	HeadFirstName    string `json:"head_first_name" validate:"required,customfieldrusnumregex"`
	HeadMiddleName   string `json:"head_middle_name" validate:"required,customfieldrusnumregex"`
	DepartamentEmail string `json:"departament_email" validate:"required,email"`
}

type PutDepartamentRequest struct {
	DepartamentID    int64  `json:"departament_id" validate:"required,numeric"`
	FacultyID        int64  `json:"faculty_id" validate:"required,numeric"`
	DepartamentName  string `json:"departament_name" validate:"required,customfieldrusnumregex"`
	HeadLastName     string `json:"head_last_name" validate:"required,customfieldrusnumregex"`
	HeadFirstName    string `json:"head_first_name" validate:"required,customfieldrusnumregex"`
	HeadMiddleName   string `json:"head_middle_name" validate:"required,customfieldrusnumregex"`
	DepartamentEmail string `json:"departament_email" validate:"required,email"`
}

type PatchDepartamentRequest struct {
	DepartamentID    int64  `json:"departament_id" validate:"required,numeric"`
	FacultyID        int64  `json:"faculty_id" validate:"omitempty,numeric"`
	DepartamentName  string `json:"departament_name" validate:"omitempty,customfieldrusnumregex"`
	HeadLastName     string `json:"head_last_name" validate:"omitempty,customfieldrusnumregex"`
	HeadFirstName    string `json:"head_first_name" validate:"omitempty,customfieldrusnumregex"`
	HeadMiddleName   string `json:"head_middle_name" validate:"omitempty,customfieldrusnumregex"`
	DepartamentEmail string `json:"departament_email" validate:"omitempty,email"`
}

// CreateDepartament godoc
// @Summary Create a departament
// @Description Create a new departament
// @Tags Departaments
// @Accept json
// @Produce json
// @Param departament body CreateDepartamentRequest true "Departament info"
// @Success 201 {object} SuccessResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /departaments [post]
func (h *Handler) CreateDepartament(c *gin.Context) {
	var req CreateDepartamentRequest
	if err := c.BindJSON(&req); err != nil {
		respondWithError(h.logger, c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.validate.Struct(req); err != nil {
		respondWithError(h.logger, c, http.StatusBadRequest, err.Error())
		return
	}

	departament := domain.Departament{
		FacultyID:        req.FacultyID,
		DepartamentName:  req.DepartamentName,
		HeadLastName:     req.HeadLastName,
		HeadFirstName:    req.HeadFirstName,
		HeadMiddleName:   req.HeadMiddleName,
		DepartamentEmail: req.DepartamentEmail,
	}

	err := h.services.DepartamentService.Create(c.Request.Context(), departament)
	if err != nil {
		respondWithError(h.logger, c, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithSuccess(c, http.StatusCreated, "Departament created successfully")
}

// PutDepartament godoc
// @Summary Update a departament
// @Description Update an existing departament
// @Tags Departaments
// @Accept json
// @Produce json
// @Param departament body PutDepartamentRequest true "Departament info"
// @Success 200 {object} SuccessResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /departaments [put]
func (h *Handler) PutDepartament(c *gin.Context) {
	var req PutDepartamentRequest
	if err := c.BindJSON(&req); err != nil {
		respondWithError(h.logger, c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.validate.Struct(req); err != nil {
		respondWithError(h.logger, c, http.StatusBadRequest, err.Error())
		return
	}

	departament := domain.Departament{
		DepartamentID:    req.DepartamentID,
		FacultyID:        req.FacultyID,
		DepartamentName:  req.DepartamentName,
		HeadLastName:     req.HeadLastName,
		HeadFirstName:    req.HeadFirstName,
		HeadMiddleName:   req.HeadMiddleName,
		DepartamentEmail: req.DepartamentEmail,
	}

	err := h.services.DepartamentService.Put(c.Request.Context(), departament)
	if err != nil {
		respondWithError(h.logger, c, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithSuccess(c, http.StatusOK, "Departament updated successfully")
}

// PatchDepartament godoc
// @Summary Partially update a departament
// @Description Partially update an existing departament
// @Tags Departaments
// @Accept json
// @Produce json
// @Param departament body PatchDepartamentRequest true "Departament info"
// @Success 200 {object} SuccessResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /departaments [patch]
func (h *Handler) PatchDepartament(c *gin.Context) {
	var req PatchDepartamentRequest
	if err := c.BindJSON(&req); err != nil {
		respondWithError(h.logger, c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.validate.Struct(req); err != nil {
		respondWithError(h.logger, c, http.StatusBadRequest, err.Error())
		return
	}

	departament := domain.Departament{
		DepartamentID:    req.DepartamentID,
		FacultyID:        req.FacultyID,
		DepartamentName:  req.DepartamentName,
		HeadLastName:     req.HeadLastName,
		HeadFirstName:    req.HeadFirstName,
		HeadMiddleName:   req.HeadMiddleName,
		DepartamentEmail: req.DepartamentEmail,
	}

	err := h.services.DepartamentService.Patch(c.Request.Context(), departament)
	if err != nil {
		respondWithError(h.logger, c, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithSuccess(c, http.StatusOK, "Departament patched successfully")
}

// DeleteDepartament godoc
// @Summary Delete a departament
// @Description Delete an existing departament
// @Tags Departaments
// @Accept json
// @Produce json
// @Param id path int64 true "Departament ID"
// @Success 200 {object} SuccessResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /departaments/{id} [delete]
func (h *Handler) DeleteDepartament(c *gin.Context) {
	departamentID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		respondWithError(h.logger, c, http.StatusBadRequest, "Invalid departament ID")
		return
	}

	err = h.services.DepartamentService.Delete(c.Request.Context(), departamentID)
	if err != nil {
		respondWithError(h.logger, c, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithSuccess(c, http.StatusOK, "Departament deleted successfully")
}

// GetDepartamentByID godoc
// @Summary Get a departament by ID
// @Description Get a departament by its ID
// @Tags Departaments
// @Accept json
// @Produce json
// @Param id path int64 true "Departament ID"
// @Success 200 {object} domain.DepartamentInfo
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /departaments/{id} [get]
func (h *Handler) GetDepartamentByID(c *gin.Context) {
	departamentID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		respondWithError(h.logger, c, http.StatusBadRequest, "Invalid departament ID")
		return
	}

	departament, err := h.services.DepartamentService.GetByID(c.Request.Context(), departamentID)
	if err != nil {
		respondWithError(h.logger, c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, departament)
}

// GetDepartamentByName godoc
// @Summary Get a departament by name
// @Description Get a departament by its name
// @Tags Departaments
// @Accept json
// @Produce json
// @Param name path string true "Departament Name"
// @Success 200 {object} domain.DepartamentInfo
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /departaments/name/{name} [get]
func (h *Handler) GetDepartamentByName(c *gin.Context) {
	departamentName := c.Param("name")

	departament, err := h.services.DepartamentService.GetByName(c.Request.Context(), departamentName)
	if err != nil {
		respondWithError(h.logger, c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, departament)
}

// GetAllDepartaments godoc
// @Summary Get all departaments
// @Description Get a list of all departaments
// @Tags Departaments
// @Accept json
// @Produce json
// @Success 200 {array} domain.DepartamentInfo
// @Failure 500 {object} ErrorResponse
// @Router /departaments [get]
func (h *Handler) GetAllDepartaments(c *gin.Context) {
	departaments, err := h.services.DepartamentService.GetAll(c.Request.Context())
	if err != nil {
		respondWithError(h.logger, c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, departaments)
}

// GetDepartamentsByFacultyID godoc
// @Summary Get departaments by faculty ID
// @Description Get a list of departaments by faculty ID
// @Tags Departaments
// @Accept json
// @Produce json
// @Param faculty_id path int64 true "Faculty ID"
// @Success 200 {array} domain.DepartamentInfo
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /departaments/faculty/{faculty_id} [get]
func (h *Handler) GetDepartamentsByFacultyID(c *gin.Context) {
	facultyID, err := strconv.ParseInt(c.Param("faculty_id"), 10, 64)
	if err != nil {
		respondWithError(h.logger, c, http.StatusBadRequest, "Invalid faculty ID")
		return
	}

	departaments, err := h.services.DepartamentService.GetAllByFacultyID(c.Request.Context(), facultyID)
	if err != nil {
		respondWithError(h.logger, c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, departaments)
}
