package handler

import (
	"net/http"
	"strconv"

	"github.com/BeRebornBng/OsauAmsApi/domain"
	"github.com/gin-gonic/gin"
)

// CreateHeadmanRequest represents the request body for creating a headman
type CreateHeadmanRequest struct {
	StudentID int64  `json:"student_id" validate:"required,numeric"`
	GroupID   string `json:"group_id" validate:"required,customgroupidregex"`
}

// PutHeadmanRequest represents the request body for updating a headman
type PutHeadmanRequest struct {
	HeadmanID int64  `json:"headman_id" validate:"required,numeric"`
	StudentID int64  `json:"student_id" validate:"required,numeric"`
	GroupID   string `json:"group_id" validate:"required,customgroupidregex"`
}

// PatchHeadmanRequest represents the request body for partially updating a headman
type PatchHeadmanRequest struct {
	HeadmanID int64  `json:"headman_id" validate:"required,numeric"`
	StudentID int64  `json:"student_id" validate:"omitempty,numeric"`
	GroupID   string `json:"group_id" validate:"omitempty,customgroupidregex"`
}

// CreateHeadman godoc
// @Summary Create a headman
// @Description Create a new headman
// @Tags Headmen
// @Accept json
// @Produce json
// @Param headman body CreateHeadmanRequest true "Headman info"
// @Success 201 {object} SuccessResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /headmen [post]
func (h *Handler) CreateHeadman(c *gin.Context) {
	var req CreateHeadmanRequest
	if err := c.BindJSON(&req); err != nil {
		respondWithError(h.logger, c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.validate.Struct(req); err != nil {
		respondWithError(h.logger, c, http.StatusBadRequest, err.Error())
		return
	}

	headman := domain.Headman{
		StudentID: req.StudentID,
		GroupID:   req.GroupID,
	}

	err := h.services.HeadmanService.Create(c.Request.Context(), headman)
	if err != nil {
		respondWithError(h.logger, c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, SuccessResponse{Message: "Headman created successfully"})
}

// PutHeadman godoc
// @Summary Update a headman
// @Description Update an existing headman
// @Tags Headmen
// @Accept json
// @Produce json
// @Param headman body PutHeadmanRequest true "Headman info"
// @Success 200 {object} SuccessResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /headmen [put]
func (h *Handler) PutHeadman(c *gin.Context) {
	var req PutHeadmanRequest
	if err := c.BindJSON(&req); err != nil {
		respondWithError(h.logger, c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.validate.Struct(req); err != nil {
		respondWithError(h.logger, c, http.StatusBadRequest, err.Error())
		return
	}

	headman := domain.Headman{
		HeadmanID: req.HeadmanID,
		StudentID: req.StudentID,
		GroupID:   req.GroupID,
	}

	err := h.services.HeadmanService.Put(c.Request.Context(), headman)
	if err != nil {
		respondWithError(h.logger, c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, SuccessResponse{Message: "Headman updated successfully"})
}

// PatchHeadman godoc
// @Summary Partially update a headman
// @Description Partially update an existing headman
// @Tags Headmen
// @Accept json
// @Produce json
// @Param headman body PatchHeadmanRequest true "Headman info"
// @Success 200 {object} SuccessResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /headmen [patch]
func (h *Handler) PatchHeadman(c *gin.Context) {
	var req PatchHeadmanRequest
	if err := c.BindJSON(&req); err != nil {
		respondWithError(h.logger, c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.validate.Struct(req); err != nil {
		respondWithError(h.logger, c, http.StatusBadRequest, err.Error())
		return
	}

	headman := domain.Headman{
		HeadmanID: req.HeadmanID,
		StudentID: req.StudentID,
		GroupID:   req.GroupID,
	}

	err := h.services.HeadmanService.Patch(c.Request.Context(), headman)
	if err != nil {
		respondWithError(h.logger, c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, SuccessResponse{Message: "Headman patched successfully"})
}

// DeleteHeadman godoc
// @Summary Delete a headman
// @Description Delete an existing headman
// @Tags Headmen
// @Accept json
// @Produce json
// @Param id path int64 true "Headman ID"
// @Success 200 {object} SuccessResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /headmen/{id} [delete]
func (h *Handler) DeleteHeadman(c *gin.Context) {
	headmanID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		respondWithError(h.logger, c, http.StatusBadRequest, "Invalid headman ID")
		return
	}

	err = h.services.HeadmanService.Delete(c.Request.Context(), headmanID)
	if err != nil {
		respondWithError(h.logger, c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, SuccessResponse{Message: "Headman deleted successfully"})
}

// GetHeadmanByID godoc
// @Summary Get a headman by ID
// @Description Get a headman by its ID
// @Tags Headmen
// @Accept json
// @Produce json
// @Param id path int64 true "Headman ID"
// @Success 200 {object} domain.HeadmanInfo
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /headmen/{id} [get]
func (h *Handler) GetHeadmanByID(c *gin.Context) {
	headmanID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		respondWithError(h.logger, c, http.StatusBadRequest, "Invalid headman ID")
		return
	}

	headman, err := h.services.HeadmanService.GetByID(c.Request.Context(), headmanID)
	if err != nil {
		respondWithError(h.logger, c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, headman)
}

// GetHeadmanByID godoc
// @Summary Get a headman by ID
// @Description Get a headman by its ID
// @Tags Headmen
// @Accept json
// @Produce json
// @Param id path int64 true "Headman ID"
// @Success 200 {object} domain.HeadmanInfo
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /headmen/{id} [get]
func (h *Handler) GetHeadmanByStudentID(c *gin.Context) {
	headmanID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		respondWithError(h.logger, c, http.StatusBadRequest, "Invalid headman ID")
		return
	}

	headman, err := h.services.HeadmanService.GetByID(c.Request.Context(), headmanID)
	if err != nil {
		respondWithError(h.logger, c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, headman)
}

// GetAllHeadmen godoc
// @Summary Get all headmen
// @Description Get a list of all headmen
// @Tags Headmen
// @Accept json
// @Produce json
// @Success 200 {array} domain.HeadmanInfo
// @Failure 500 {object} ErrorResponse
// @Router /headmen [get]
func (h *Handler) GetAllHeadmen(c *gin.Context) {
	headmen, err := h.services.HeadmanService.GetAll(c.Request.Context())
	if err != nil {
		respondWithError(h.logger, c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, headmen)
}
