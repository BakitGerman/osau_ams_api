package service

import (
	"context"
	"errors"
	"time"

	"github.com/BeRebornBng/OsauAmsApi/domain"
	"github.com/BeRebornBng/OsauAmsApi/internal/repository"
	"github.com/BeRebornBng/OsauAmsApi/pkg/auth"
	"github.com/BeRebornBng/OsauAmsApi/pkg/myhash"
	"github.com/google/uuid"
	"github.com/jackc/pgx"
)

type UserService struct {
	TokenManager   auth.TokenManager
	Hasher         myhash.PasswordHasher
	UserRepo       repository.IUser
	StudentRepo    repository.IStudent
	AccessTokenTTL time.Duration
}

func NewUserService(
	TokenManager auth.TokenManager,
	Hasher myhash.PasswordHasher,
	UserRepo repository.IUser,
	AccessTokenTTL time.Duration,
) *UserService {
	return &UserService{
		TokenManager:   TokenManager,
		Hasher:         Hasher,
		UserRepo:       UserRepo,
		AccessTokenTTL: AccessTokenTTL,
	}
}

func (s *UserService) SignIn(ctx context.Context, username, password string) (Tokens, error) {
	user, err := s.UserRepo.GetByName(ctx, username)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return Tokens{}, ErrUserNamePassNotExists
		}
		return Tokens{}, err
	}

	if !(s.Hasher.ComparePassword(user.User.Password, password)) {
		return Tokens{}, ErrUserNamePassNotExists
	}

	accessToken, err := s.TokenManager.NewJWT(user.User.UserID.String(), user.User.Role, s.AccessTokenTTL)
	if err != nil {
		return Tokens{}, err
	}

	return Tokens{AccessToken: accessToken}, nil
}

func (s *UserService) Create(ctx context.Context, user domain.User) error {
	_, err := s.UserRepo.GetByName(ctx, user.Username)
	if err != nil {
		if err.Error() != pgx.ErrNoRows.Error() {
			return err
		}
	} else {
		return ErrUserNameExists
	}
	err = s.userWithRoleIDExists(ctx, user.StudentID, user.TeacherID, user.HeadmanID)
	if err != nil {
		return err
	}

	hashpassword, err := s.Hasher.HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashpassword

	return s.UserRepo.Create(ctx, user)
}

func (s *UserService) Put(ctx context.Context, user domain.User) error {

	hashpassword, err := s.Hasher.HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashpassword
	return s.UserRepo.Put(ctx, user)
}

func (s *UserService) Patch(ctx context.Context, user domain.User) error {

	updates := make(map[string]interface{})
	if user.Username != "" {
		updates["username"] = user.Username
	}
	if user.Role != "" {
		updates["user_role"] = user.Role
	}
	if user.HeadmanID != nil {
		updates["headman_id"] = user.HeadmanID
	}
	if user.StudentID != nil {
		updates["student_id"] = user.StudentID
	}
	if user.TeacherID != nil {
		updates["teacher_id"] = user.TeacherID
	}
	if user.Password != "" {
		hashpassword, err := s.Hasher.HashPassword(user.Password)
		if err != nil {
			return err
		}
		updates["password"] = hashpassword
	}
	if len(updates) == 0 {
		return ErrNoUpdates
	}
	return s.UserRepo.Patch(ctx, user.UserID, updates)
}

func (s *UserService) Delete(ctx context.Context, userID uuid.UUID) error {
	return s.UserRepo.Delete(ctx, userID)
}

func (s *UserService) GetByID(ctx context.Context, userID uuid.UUID) (domain.UserInfo, error) {
	return s.UserRepo.GetByID(ctx, userID)
}

func (s *UserService) GetByName(ctx context.Context, username string) (domain.UserInfo, error) {
	return s.UserRepo.GetByName(ctx, username)
}

func (s *UserService) GetByStudentID(ctx context.Context, studentID int64) (domain.UserInfo, error) {
	return s.UserRepo.GetByStudentID(ctx, studentID)
}

func (s *UserService) GetByHeadmanID(ctx context.Context, teacherID int64) (domain.UserInfo, error) {
	return s.UserRepo.GetByHeadmanID(ctx, teacherID)
}

func (s *UserService) GetByTeacherID(ctx context.Context, teacherID int64) (domain.UserInfo, error) {
	return s.UserRepo.GetByTeacherID(ctx, teacherID)
}

func (s *UserService) GetAllByRole(ctx context.Context, role string) ([]domain.UserInfo, error) {
	return s.UserRepo.GetAllByRole(ctx, role)
}

func (s *UserService) GetAll(ctx context.Context) ([]domain.UserInfo, error) {
	return s.UserRepo.GetAll(ctx)
}

func (s *UserService) userWithRoleIDExists(ctx context.Context, studentID *int64, teacherID *int64, headmanID *int64) error {

	var err error
	if studentID != nil {
		_, err = s.UserRepo.GetByStudentID(ctx, *studentID)
		if err != nil {
			if err.Error() != pgx.ErrNoRows.Error() {
				return err
			}
		} else {
			return ErrStudentIDExists
		}
	} else if teacherID != nil {
		_, err = s.UserRepo.GetByTeacherID(ctx, *teacherID)
		if err != nil {
			if err.Error() != pgx.ErrNoRows.Error() {
				return err
			}
		} else {
			return ErrTeacherIDExists
		}
	} else if headmanID != nil {
		_, err = s.UserRepo.GetByHeadmanID(ctx, *headmanID)
		if err != nil {
			if err.Error() != pgx.ErrNoRows.Error() {
				return err
			}
		} else {
			return ErrHeadmanIDExists
		}
	}

	return nil
}
