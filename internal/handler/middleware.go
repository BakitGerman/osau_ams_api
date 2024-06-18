package handler

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "user_id"
	roleCtx             = "user_role"
	groupCtx            = "group_id"
	teacherCtx          = "teacher_id"
)

func (h *Handler) parseAuthHeader(c *gin.Context) (string, error) {
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		return "", errors.New("empty auth header")
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		return "", errors.New("invalid auth header")
	}

	if len(headerParts[1]) == 0 {
		return "", errors.New("token is empty")
	}

	return h.TokenManager.Parse(headerParts[1])
}

func (h *Handler) userIdentity(c *gin.Context) {
	id, err := h.parseAuthHeader(c)
	if err != nil {
		respondWithError(h.logger, c, http.StatusForbidden, err.Error())
		return
	}
	userID, err := uuid.Parse(id)
	if err != nil {
		respondWithError(h.logger, c, http.StatusForbidden, err.Error())
		return
	}
	user, err := h.services.UserService.GetByID(c.Request.Context(), userID)
	if err != nil {
		respondWithError(h.logger, c, http.StatusForbidden, err.Error())
		return
	}

	c.Set(userCtx, userID)
	c.Set(roleCtx, user.User.Role)
	if user.UserSub.GroupID != nil {
		c.Set(groupCtx, *user.UserSub.GroupID)
	}
	if user.User.TeacherID != nil {
		c.Set(teacherCtx, *user.User.TeacherID)
	}
}

func RoleMiddleware(role string) gin.HandlerFunc {
	return func(c *gin.Context) {
		userRole, exists := c.Get(roleCtx)
		if !exists || userRole != role {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"message": "Forbidden"})
			return
		}
		c.Next()
	}
}
