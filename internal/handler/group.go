package handler

import (
	"net/http"
	"strconv"

	"github.com/BeRebornBng/OsauAmsApi/domain"
	"github.com/gin-gonic/gin"
)

type CreateGroupRequest struct {
	GroupID   string `json:"group_id" validate:"required,customgroupidregex"`
	ProfileID int64  `json:"profile_id" validate:"required,numeric"`
}

type PutGroupRequest struct {
	GroupID   string `json:"group_id" validate:"required,customgroupidregex"`
	ProfileID int64  `json:"profile_id" validate:"required,numeric"`
}

type PatchGroupRequest struct {
	GroupID   string `json:"group_id" validate:"required,customgroupidregex"`
	ProfileID int64  `json:"profile_id" validate:"omitempty,numeric"`
}

// CreateGroup godoc
// @Summary Create a group
// @Description Create a new group
// @Tags Groups
// @Accept json
// @Produce json
// @Param group body CreateGroupRequest true "Group info"
// @Success 201 {object} SuccessResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /groups [post]
func (h *Handler) CreateGroup(c *gin.Context) {
	var req CreateGroupRequest
	if err := c.BindJSON(&req); err != nil {
		respondWithError(h.logger, c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.validate.Struct(req); err != nil {
		respondWithError(h.logger, c, http.StatusBadRequest, err.Error())
		return
	}

	group := domain.Group{
		GroupID:   req.GroupID,
		ProfileID: req.ProfileID,
	}

	err := h.services.GroupService.Create(c.Request.Context(), group)
	if err != nil {
		respondWithError(h.logger, c, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithSuccess(c, http.StatusCreated, "Group created successfully")
}

// PutGroup godoc
// @Summary Update a group
// @Description Update an existing group
// @Tags Groups
// @Accept json
// @Produce json
// @Param group body PutGroupRequest true "Group info"
// @Success 200 {object} SuccessResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /groups [put]
func (h *Handler) PutGroup(c *gin.Context) {
	var req PutGroupRequest
	if err := c.BindJSON(&req); err != nil {
		respondWithError(h.logger, c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.validate.Struct(req); err != nil {
		respondWithError(h.logger, c, http.StatusBadRequest, err.Error())
		return
	}

	group := domain.Group{
		GroupID:   req.GroupID,
		ProfileID: req.ProfileID,
	}

	err := h.services.GroupService.Put(c.Request.Context(), group)
	if err != nil {
		respondWithError(h.logger, c, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithSuccess(c, http.StatusOK, "Group updated successfully")
}

// PatchGroup godoc
// @Summary Partially update a group
// @Description Partially update an existing group
// @Tags Groups
// @Accept json
// @Produce json
// @Param group body PatchGroupRequest true "Group info"
// @Success 200 {object} SuccessResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /groups [patch]
func (h *Handler) PatchGroup(c *gin.Context) {
	var req PatchGroupRequest
	if err := c.BindJSON(&req); err != nil {
		respondWithError(h.logger, c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.validate.Struct(req); err != nil {
		respondWithError(h.logger, c, http.StatusBadRequest, err.Error())
		return
	}

	group := domain.Group{
		GroupID:   req.GroupID,
		ProfileID: req.ProfileID,
	}

	err := h.services.GroupService.Patch(c.Request.Context(), group)
	if err != nil {
		respondWithError(h.logger, c, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithSuccess(c, http.StatusOK, "Group patched successfully")
}

// DeleteGroup godoc
// @Summary Delete a group
// @Description Delete an existing group
// @Tags Groups
// @Accept json
// @Produce json
// @Param id path string true "Group ID"
// @Success 200 {object} SuccessResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /groups/{id} [delete]
func (h *Handler) DeleteGroup(c *gin.Context) {
	groupID := c.Param("id")

	err := h.services.GroupService.Delete(c.Request.Context(), groupID)
	if err != nil {
		respondWithError(h.logger, c, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithSuccess(c, http.StatusOK, "Group deleted successfully")
}

// GetGroupByID godoc
// @Summary Get a group by ID
// @Description Get a group by its ID
// @Tags Groups
// @Accept json
// @Produce json
// @Param id path string true "Group ID"
// @Success 200 {object} domain.GroupInfo
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /groups/{id} [get]
func (h *Handler) GetGroupByID(c *gin.Context) {
	groupID := c.Param("id")

	group, err := h.services.GroupService.GetByID(c.Request.Context(), groupID)
	if err != nil {
		respondWithError(h.logger, c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, group)
}

// GetGroupByName godoc
// @Summary Get a group by name
// @Description Get a group by its name
// @Tags Groups
// @Accept json
// @Produce json
// @Param name path string true "Group Name"
// @Success 200 {object} domain.GroupInfo
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /groups/name/{name} [get]
func (h *Handler) GetGroupByName(c *gin.Context) {
	groupName := c.Param("name")

	group, err := h.services.GroupService.GetByName(c.Request.Context(), groupName)
	if err != nil {
		respondWithError(h.logger, c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, group)
}

// GetAllGroups godoc
// @Summary Get all groups
// @Description Get a list of all groups
// @Tags Groups
// @Accept json
// @Produce json
// @Success 200 {array} domain.GroupInfo
// @Failure 500 {object} ErrorResponse
// @Router /groups [get]
func (h *Handler) GetAllGroups(c *gin.Context) {
	groups, err := h.services.GroupService.GetAll(c.Request.Context())
	if err != nil {
		respondWithError(h.logger, c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, groups)
}

// GetGroupsByProfileID godoc
// @Summary Get groups by profile ID
// @Description Get a list of groups by profile ID
// @Tags Groups
// @Accept json
// @Produce json
// @Param profile_id path int64 true "Profile ID"
// @Success 200 {array} domain.GroupInfo
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /groups/profile/{profile_id} [get]
func (h *Handler) GetGroupsByProfileID(c *gin.Context) {
	profileID, err := strconv.ParseInt(c.Param("profile_id"), 10, 64)
	if err != nil {
		respondWithError(h.logger, c, http.StatusBadRequest, "Invalid profile ID")
		return
	}

	groups, err := h.services.GroupService.GetAllByProfileID(c.Request.Context(), profileID)
	if err != nil {
		respondWithError(h.logger, c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, groups)
}
