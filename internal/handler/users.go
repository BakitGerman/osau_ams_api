package handler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/BeRebornBng/OsauAmsApi/domain"
	"github.com/BeRebornBng/OsauAmsApi/internal/service"
	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

const (
	ErrDuplicateUserName     = "the user with this username already exists"
	ErrDuplicateHeadmanID    = "the user with this headman ID already exists"
	ErrDuplicateStudentID    = "the user with this student ID already exists"
	ErrDuplicateTeacherID    = "the user with this teacher ID already exists"
	ErrDuplicateKeyViolation = "Duplicate key violation"
	ErrValidationFailed      = "Validation failed"
	ErrUserNotFound          = "user not found"
	ErrUsersNotFound         = "users not found"
	ErrInvalidUserID         = "Invalid user ID"
	ErrInvalidUsername       = "Invalid username"
	ErrInvalidStudentID      = "Invalid student ID"
	ErrInvalidHeadmanID      = "Invalid headman ID"
	ErrInvalidTeacherID      = "Invalid teacher ID"
	ErrInvalidUserRole       = "Invalid user role"
	ErrUserNamePassNotExists = "username or password does not exist"
	ErrInternalServerError   = "Internal server error"
	ConstraintUserName       = "U_users_username"
	ConstraintHeadmanID      = "U_users_headman_id"
	ConstraintStudentID      = "U_users_student_id"
	ConstraintTeacherID      = "U_users_teacher_id"
)

func parsePostgresError(err error) (string, bool) {
	if pgErr, ok := err.(*pgconn.PgError); ok {
		switch pgErr.Code {
		case "23505":
			switch pgErr.ConstraintName {
			case ConstraintUserName:
				return ErrDuplicateUserName, true
			case ConstraintHeadmanID:
				return ErrDuplicateHeadmanID, true
			case ConstraintStudentID:
				return ErrDuplicateStudentID, true
			case ConstraintTeacherID:
				return ErrDuplicateTeacherID, true
			default:
				return ErrDuplicateKeyViolation, true
			}
		default:
			return pgErr.Message, true
		}
	}
	return "", false
}

// SignInUserRequest represents the request body for signing in a user
type SignInUserRequest struct {
	Username string `json:"username" validate:"required,min=8,max=40,alphanum"`
	Password string `json:"password" validate:"required,min=8,max=40,custompasswordregex"`
}

// CreateUserRequest represents the request body for creating a user
type CreateUserRequest struct {
	Username  string `json:"username" validate:"required,min=8,max=40,alphanum"`
	Password  string `json:"password" validate:"required,min=8,max=40,custompasswordregex"`
	Role      string `json:"role" validate:"required,oneof=Студент Староста Админ Преподаватель,roledependentfields"`
	HeadmanID *int64 `json:"headman_id" validate:"required_if=Role Староста,omitempty,min=1,number"`
	StudentID *int64 `json:"student_id" validate:"required_if=Role Студент,omitempty,min=1,number"`
	TeacherID *int64 `json:"teacher_id" validate:"required_if=Role Преподаватель,omitempty,min=1,number"`
}

// PutUserRequest represents the request body for updating a user
type PutUserRequest struct {
	UserID    uuid.UUID `json:"user_id" validate:"required,uuid4"`
	Username  string    `json:"username" validate:"required,min=8,max=40,alphanum"`
	Password  string    `json:"password" validate:"required,min=8,max=40,custompasswordregex"`
	Role      string    `json:"role" validate:"required,oneof=Студент Староста Админ Преподаватель,roledependentfields"`
	HeadmanID *int64    `json:"headman_id" validate:"required_if=Role Староста,omitempty,min=1,number"`
	StudentID *int64    `json:"student_id" validate:"required_if=Role Студент,omitempty,min=1,number"`
	TeacherID *int64    `json:"teacher_id" validate:"required_if=Role Преподаватель,omitempty,min=1,number"`
}

// PatchUserRequest represents the request body for partially updating a user
type PatchUserRequest struct {
	UserID    uuid.UUID `json:"user_id" validate:"required,uuid4"`
	Username  string    `json:"username" validate:"omitempty,min=8,max=40,alphanum"`
	Password  string    `json:"password" validate:"omitempty,min=8,max=40,custompasswordregex"`
	Role      string    `json:"role" validate:"omitempty,oneof=Студент Староста Админ Преподаватель,roledependentfields"`
	HeadmanID *int64    `json:"headman_id" validate:"required_if=Role Староста,omitempty,min=1,number"`
	StudentID *int64    `json:"student_id" validate:"required_if=Role Студент,omitempty,min=1,number"`
	TeacherID *int64    `json:"teacher_id" validate:"required_if=Role Преподаватель,omitempty,min=1,number"`
}

type GetUserByNameRequest struct {
	Username string `json:"username" validate:"required,min=8,max=40,alphanum"`
}

type GetUserByRoleRequest struct {
	Role string `json:"role" validate:"omitempty,oneof=Студент Староста Админ Преподаватель,roledependentfields"`
}

func translateValidationErrors(validationErrors validator.ValidationErrors, trans ut.Translator) []string {
	var errs []string
	for _, fe := range validationErrors {
		errs = append(errs, fe.Translate(trans))
	}
	return errs
}

// SignInUser godoc
// @Summary Sign in a user
// @Description Sign in a user
// @Tags Users
// @Accept json
// @Produce json
// @Param user body SignInUserRequest true "User sign-in info"
// @Success 200 {object} service.Tokens
// @Failure 400 {object} ErrorResponse
// @Failure 401 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /auth/signin [post]
func (h *Handler) SignInUser(c *gin.Context) {
	userReq := SignInUserRequest{}
	if err := c.BindJSON(&userReq); err != nil {
		respondWithError(h.logger, c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.validate.Struct(userReq); err != nil {
		errs := translateValidationErrors(err.(validator.ValidationErrors), h.translator)
		respondWithError(h.logger, c, http.StatusBadRequest, errs[0])
		return
	}

	tokens, err := h.services.UserService.SignIn(c.Request.Context(), userReq.Username, userReq.Password)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			respondWithError(h.logger, c, http.StatusUnauthorized, err.Error())
			return
		}

		if errors.Is(err, service.ErrUserNamePassNotExists) {
			respondWithError(h.logger, c, http.StatusUnauthorized, err.Error())
			return
		} else {
			respondWithError(h.logger, c, http.StatusInternalServerError, err.Error())
			return
		}
	}

	c.JSON(http.StatusOK, tokens)
}

// CreateUser godoc
// @Security ApiKeyAuth
// @Summary Create a user
// @Description Create a new user
// @Tags Users
// @Accept json
// @Produce json
// @Param user body CreateUserRequest true "User info"
// @Success 201 {object} SuccessResponse
// @Failure 400 {object} ErrorResponse
// @Failure 409 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /admin/user [post]
func (h *Handler) CreateUser(c *gin.Context) {
	userReq := CreateUserRequest{}
	if err := c.BindJSON(&userReq); err != nil {
		respondWithError(h.logger, c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.validate.Struct(userReq); err != nil {
		respondWithError(h.logger, c, http.StatusBadRequest, err.Error())
		return
	}

	user := domain.User{
		Username:  userReq.Username,
		Password:  userReq.Password,
		Role:      userReq.Role,
		HeadmanID: userReq.HeadmanID,
		StudentID: userReq.StudentID,
		TeacherID: userReq.TeacherID,
	}

	if err := h.services.UserService.Create(c.Request.Context(), user); err != nil {
		if errors.Is(err, service.ErrUserNameExists) {
			respondWithError(h.logger, c, http.StatusConflict, ErrDuplicateUserName)
			return
		}
		if errors.Is(err, service.ErrHeadmanIDExists) {
			respondWithError(h.logger, c, http.StatusConflict, ErrDuplicateHeadmanID)
			return
		}
		if errors.Is(err, service.ErrTeacherIDExists) {
			respondWithError(h.logger, c, http.StatusConflict, ErrDuplicateTeacherID)
			return
		}
		if errors.Is(err, service.ErrStudentIDExists) {
			respondWithError(h.logger, c, http.StatusConflict, ErrDuplicateStudentID)
			return
		}
		respondWithError(h.logger, c, http.StatusInternalServerError, err.Error())
		return
	}

	c.Status(http.StatusCreated)
}

// PutUser godoc
// @Security ApiKeyAuth
// @Summary Update a user
// @Description Update an existing user
// @Tags Users
// @Accept json
// @Produce json
// @Param user body PutUserRequest true "User info"
// @Success 200 {object} SuccessResponse
// @Failure 400 {object} ErrorResponse
// @Failure 409 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /admin/user [put]
func (h *Handler) PutUser(c *gin.Context) {
	var userReq PutUserRequest
	if err := c.BindJSON(&userReq); err != nil {
		respondWithError(h.logger, c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.validate.Struct(userReq); err != nil {
		respondWithError(h.logger, c, http.StatusBadRequest, err.Error())
		return
	}

	user := domain.User{
		UserID:    userReq.UserID,
		Username:  userReq.Username,
		Password:  userReq.Password,
		Role:      userReq.Role,
		HeadmanID: userReq.HeadmanID,
		StudentID: userReq.StudentID,
		TeacherID: userReq.TeacherID,
	}

	if err := h.services.UserService.Put(c.Request.Context(), user); err != nil {
		if pgErrMsg, ok := parsePostgresError(err); ok {
			respondWithError(h.logger, c, http.StatusConflict, pgErrMsg)
			return
		}
		respondWithError(h.logger, c, http.StatusInternalServerError, err.Error())
		return
	}

	c.Status(http.StatusOK)
}

// PatchUser godoc
// @Security ApiKeyAuth
// @Summary Partially update a user
// @Description Partially update an existing user
// @Tags Users
// @Accept json
// @Produce json
// @Param user body PatchUserRequest true "User info"
// @Success 200 {object} SuccessResponse
// @Failure 400 {object} ErrorResponse
// @Failure 409 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /admin/user [patch]
func (h *Handler) PatchUser(c *gin.Context) {
	var userReq PatchUserRequest
	if err := c.BindJSON(&userReq); err != nil {
		respondWithError(h.logger, c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.validate.Struct(userReq); err != nil {
		respondWithError(h.logger, c, http.StatusBadRequest, err.Error())
		return
	}

	user := domain.User{
		UserID:    userReq.UserID,
		Username:  userReq.Username,
		Password:  userReq.Password,
		Role:      userReq.Role,
		HeadmanID: userReq.HeadmanID,
		StudentID: userReq.StudentID,
		TeacherID: userReq.TeacherID,
	}

	if err := h.services.UserService.Patch(c.Request.Context(), user); err != nil {
		if pgErrMsg, ok := parsePostgresError(err); ok {
			respondWithError(h.logger, c, http.StatusConflict, pgErrMsg)
			return
		}
		respondWithError(h.logger, c, http.StatusInternalServerError, err.Error())
		return
	}

	c.Status(http.StatusOK)
}

// DeleteUser godoc
// @Security ApiKeyAuth
// @Summary Delete a user
// @Description Delete an existing user
// @Tags Users
// @Accept json
// @Produce json
// @Param id path uuid.UUID true "User ID"
// @Success 200 {object} SuccessResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /admin/user/{id} [delete]
func (h *Handler) DeleteUser(c *gin.Context) {
	userID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		respondWithError(h.logger, c, http.StatusBadRequest, ErrInvalidUserID)
		return
	}

	if err := h.services.UserService.Delete(c.Request.Context(), userID); err != nil {
		respondWithError(h.logger, c, http.StatusInternalServerError, err.Error())
		return
	}

	c.Status(http.StatusOK)
}

// GetUserByID godoc
// @Security ApiKeyAuth
// @Summary Get a user by ID
// @Description Get a user by its ID
// @Tags Users
// @Accept json
// @Produce json
// @Param id path uuid.UUID true "User ID"
// @Success 200 {object} domain.UserInfo
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /admin/user/{id} [get]
func (h *Handler) GetUserByID(c *gin.Context) {
	userID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		respondWithError(h.logger, c, http.StatusBadRequest, ErrInvalidUserID)
		return
	}

	user, err := h.services.UserService.GetByID(c.Request.Context(), userID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			respondWithError(h.logger, c, http.StatusNotFound, ErrUserNotFound)
			return
		}
		respondWithError(h.logger, c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, user)
}

// GetUserByName godoc
// @Security ApiKeyAuth
// @Summary Get a user by name
// @Description Get a user by its name
// @Tags Users
// @Accept json
// @Produce json
// @Param name path string true "User name"
// @Success 200 {object} domain.UserInfo
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /admin/user/name/{name} [get]
func (h *Handler) GetUserByName(c *gin.Context) {
	userReq := GetUserByNameRequest{Username: c.Param("name")}
	if err := h.validate.Struct(userReq); err != nil {
		respondWithError(h.logger, c, http.StatusBadRequest, ErrInvalidUsername)
		return
	}

	user, err := h.services.UserService.GetByName(c.Request.Context(), userReq.Username)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			respondWithError(h.logger, c, http.StatusNotFound, ErrUserNotFound)
			return
		}
		respondWithError(h.logger, c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, user)
}

// GetUserByStudentID godoc
// @Security ApiKeyAuth
// @Summary Get a user by student ID
// @Description Get a user by student ID
// @Tags Users
// @Accept json
// @Produce json
// @Param student_id path int64 true "Student ID"
// @Success 200 {object} domain.UserInfo
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /admin/user/student/{student_id} [get]
func (h *Handler) GetUserByStudentID(c *gin.Context) {
	studentID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		respondWithError(h.logger, c, http.StatusBadRequest, ErrInvalidStudentID)
		return
	}

	user, err := h.services.UserService.GetByStudentID(c.Request.Context(), studentID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			respondWithError(h.logger, c, http.StatusNotFound, ErrUserNotFound)
			return
		}
		respondWithError(h.logger, c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, user)
}

// GetUserByHeadmanID godoc
// @Security ApiKeyAuth
// @Summary Get a user by headman ID
// @Description Get a user by headman ID
// @Tags Users
// @Accept json
// @Produce json
// @Param headman_id path int64 true "Headman ID"
// @Success 200 {object} domain.UserInfo
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /admin/user/headman/{headman_id} [get]
func (h *Handler) GetUserByHeadmanID(c *gin.Context) {
	headmanID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		respondWithError(h.logger, c, http.StatusBadRequest, ErrInvalidHeadmanID)
		return
	}

	user, err := h.services.UserService.GetByHeadmanID(c.Request.Context(), headmanID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			respondWithError(h.logger, c, http.StatusNotFound, ErrUserNotFound)
			return
		}
		respondWithError(h.logger, c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, user)
}

// GetUserByTeacherID godoc
// @Security ApiKeyAuth
// @Summary Get a user by teacher ID
// @Description Get a user by teacher ID
// @Tags Users
// @Accept json
// @Produce json
// @Param teacher_id path int64 true "Teacher ID"
// @Success 200 {object} domain.UserInfo
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /admin/user/teacher/{teacher_id} [get]
func (h *Handler) GetUserByTeacherID(c *gin.Context) {
	teacherID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		respondWithError(h.logger, c, http.StatusBadRequest, ErrInvalidTeacherID)
		return
	}

	user, err := h.services.UserService.GetByTeacherID(c.Request.Context(), teacherID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			respondWithError(h.logger, c, http.StatusNotFound, ErrUserNotFound)
			return
		}
		respondWithError(h.logger, c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, user)
}

// GetAllUsersByRole godoc
// @Security ApiKeyAuth
// @Summary Get all users by role
// @Description Get a list of all users by role
// @Tags Users
// @Accept json
// @Produce json
// @Param role path string true "User role"
// @Success 200 {array} domain.UserInfo
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /admin/user/role/{role} [get]
func (h *Handler) GetAllUsersByRole(c *gin.Context) {
	userReq := GetUserByRoleRequest{Role: c.Param("role")}
	if err := h.validate.Struct(userReq); err != nil {
		respondWithError(h.logger, c, http.StatusBadRequest, ErrInvalidUserRole)
		return
	}

	users, err := h.services.UserService.GetAllByRole(c.Request.Context(), userReq.Role)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			respondWithError(h.logger, c, http.StatusNotFound, ErrUsersNotFound)
			return
		}
		respondWithError(h.logger, c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, users)
}

// GetAll godoc
// @Security ApiKeyAuth
// @Summary Get all users
// @Description Get a list of all users
// @Tags Users
// @Produce json
// @Success 200 {array} domain.UserInfo
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /admin/user [get]
func (h *Handler) GetAll(c *gin.Context) {
	users, err := h.services.UserService.GetAll(c.Request.Context())
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			respondWithError(h.logger, c, http.StatusNotFound, ErrUsersNotFound)
			return
		}
		respondWithError(h.logger, c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, users)
}
