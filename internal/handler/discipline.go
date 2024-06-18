package handler

import (
	"net/http"
	"strconv"

	"github.com/BeRebornBng/OsauAmsApi/domain"
	"github.com/gin-gonic/gin"
)

type CreateDisciplineRequest struct {
	DepartamentID  int64  `json:"departament_id" validate:"required,numeric"`
	DisciplineName string `json:"discipline_name" validate:"required,customfieldrusnumregex"`
}

type PutDisciplineRequest struct {
	DisciplineID   int64  `json:"discipline_id" validate:"required,numeric"`
	DepartamentID  int64  `json:"departament_id" validate:"required,numeric"`
	DisciplineName string `json:"discipline_name" validate:"required,customfieldrusnumregex"`
}

type PatchDisciplineRequest struct {
	DisciplineID   int64  `json:"discipline_id" validate:"required,numeric"`
	DepartamentID  int64  `json:"departament_id" validate:"omitempty,numeric"`
	DisciplineName string `json:"discipline_name" validate:"omitempty,customfieldrusnumregex"`
}

// CreateDiscipline godoc
// @Summary Create a discipline
// @Description Create a new discipline
// @Tags Disciplines
// @Accept json
// @Produce json
// @Param discipline body CreateDisciplineRequest true "Discipline info"
// @Success 201 {object} SuccessResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /disciplines [post]
func (h *Handler) CreateDiscipline(c *gin.Context) {
	var req CreateDisciplineRequest
	if err := c.BindJSON(&req); err != nil {
		respondWithError(h.logger, c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.validate.Struct(req); err != nil {
		respondWithError(h.logger, c, http.StatusBadRequest, err.Error())
		return
	}

	discipline := domain.Discipline{
		DepartamentID:  req.DepartamentID,
		DisciplineName: req.DisciplineName,
	}

	err := h.services.DisciplineService.Create(c.Request.Context(), discipline)
	if err != nil {
		respondWithError(h.logger, c, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithSuccess(c, http.StatusCreated, "Discipline created successfully")
}

// PutDiscipline godoc
// @Summary Update a discipline
// @Description Update an existing discipline
// @Tags Disciplines
// @Accept json
// @Produce json
// @Param discipline body PutDisciplineRequest true "Discipline info"
// @Success 200 {object} SuccessResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /disciplines [put]
func (h *Handler) PutDiscipline(c *gin.Context) {
	var req PutDisciplineRequest
	if err := c.BindJSON(&req); err != nil {
		respondWithError(h.logger, c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.validate.Struct(req); err != nil {
		respondWithError(h.logger, c, http.StatusBadRequest, err.Error())
		return
	}

	discipline := domain.Discipline{
		DisciplineID:   req.DisciplineID,
		DepartamentID:  req.DepartamentID,
		DisciplineName: req.DisciplineName,
	}

	err := h.services.DisciplineService.Put(c.Request.Context(), discipline)
	if err != nil {
		respondWithError(h.logger, c, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithSuccess(c, http.StatusOK, "Discipline updated successfully")
}

// PatchDiscipline godoc
// @Summary Partially update a discipline
// @Description Partially update an existing discipline
// @Tags Disciplines
// @Accept json
// @Produce json
// @Param discipline body PatchDisciplineRequest true "Discipline info"
// @Success 200 {object} SuccessResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /disciplines [patch]
func (h *Handler) PatchDiscipline(c *gin.Context) {
	var req PatchDisciplineRequest
	if err := c.BindJSON(&req); err != nil {
		respondWithError(h.logger, c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.validate.Struct(req); err != nil {
		respondWithError(h.logger, c, http.StatusBadRequest, err.Error())
		return
	}

	discipline := domain.Discipline{
		DisciplineID:   req.DisciplineID,
		DepartamentID:  req.DepartamentID,
		DisciplineName: req.DisciplineName,
	}

	err := h.services.DisciplineService.Patch(c.Request.Context(), discipline)
	if err != nil {
		respondWithError(h.logger, c, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithSuccess(c, http.StatusOK, "Discipline patched successfully")
}

// DeleteDiscipline godoc
// @Summary Delete a discipline
// @Description Delete an existing discipline
// @Tags Disciplines
// @Accept json
// @Produce json
// @Param id path int64 true "Discipline ID"
// @Success 200 {object} SuccessResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /disciplines/{id} [delete]
func (h *Handler) DeleteDiscipline(c *gin.Context) {
	disciplineID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		respondWithError(h.logger, c, http.StatusBadRequest, "Invalid discipline ID")
		return
	}

	err = h.services.DisciplineService.Delete(c.Request.Context(), disciplineID)
	if err != nil {
		respondWithError(h.logger, c, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithSuccess(c, http.StatusOK, "Discipline deleted successfully")
}

// GetDisciplineByID godoc
// @Summary Get a discipline by ID
// @Description Get a discipline by its ID
// @Tags Disciplines
// @Accept json
// @Produce json
// @Param id path int64 true "Discipline ID"
// @Success 200 {object} domain.DisciplineInfo
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /disciplines/{id} [get]
func (h *Handler) GetDisciplineByID(c *gin.Context) {
	disciplineID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		respondWithError(h.logger, c, http.StatusBadRequest, "Invalid discipline ID")
		return
	}

	discipline, err := h.services.DisciplineService.GetByID(c.Request.Context(), disciplineID)
	if err != nil {
		respondWithError(h.logger, c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, discipline)
}

// GetDisciplineByName godoc
// @Summary Get a discipline by name
// @Description Get a discipline by its name
// @Tags Disciplines
// @Accept json
// @Produce json
// @Param name path string true "Discipline Name"
// @Success 200 {object} domain.DisciplineInfo
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /disciplines/name/{name} [get]
func (h *Handler) GetDisciplineByName(c *gin.Context) {
	disciplineName := c.Param("name")

	discipline, err := h.services.DisciplineService.GetByName(c.Request.Context(), disciplineName)
	if err != nil {
		respondWithError(h.logger, c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, discipline)
}

// GetAllDisciplines godoc
// @Summary Get all disciplines
// @Description Get a list of all disciplines
// @Tags Disciplines
// @Accept json
// @Produce json
// @Success 200 {array} domain.DisciplineInfo
// @Failure 500 {object} ErrorResponse
// @Router /disciplines [get]
func (h *Handler) GetAllDisciplines(c *gin.Context) {
	disciplines, err := h.services.DisciplineService.GetAll(c.Request.Context())
	if err != nil {
		respondWithError(h.logger, c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, disciplines)
}

// GetDisciplinesByDepartamentID godoc
// @Summary Get disciplines by departament ID
// @Description Get a list of disciplines by departament ID
// @Tags Disciplines
// @Accept json
// @Produce json
// @Param departament_id path int64 true "Departament ID"
// @Success 200 {array} domain.DisciplineInfo
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /disciplines/departament/{departament_id} [get]
func (h *Handler) GetDisciplinesByDepartamentID(c *gin.Context) {
	departamentID, err := strconv.ParseInt(c.Param("departament_id"), 10, 64)
	if err != nil {
		respondWithError(h.logger, c, http.StatusBadRequest, "Invalid departament ID")
		return
	}

	disciplines, err := h.services.DisciplineService.GetAllByDepartamentID(c.Request.Context(), departamentID)
	if err != nil {
		respondWithError(h.logger, c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, disciplines)
}
