package service

import (
	"context"

	"github.com/BeRebornBng/OsauAmsApi/domain"
	"github.com/BeRebornBng/OsauAmsApi/internal/repository"
)

type TeacherService struct {
	TeacherRepo repository.ITeacher
}

func NewTeacherService(teacherRepo repository.ITeacher) *TeacherService {
	return &TeacherService{TeacherRepo: teacherRepo}
}

func (s *TeacherService) Create(ctx context.Context, teacher domain.Teacher) error {
	return s.TeacherRepo.Create(ctx, teacher)
}

func (s *TeacherService) Put(ctx context.Context, teacher domain.Teacher) error {
	return s.TeacherRepo.Put(ctx, teacher)
}

func (s *TeacherService) Patch(ctx context.Context, teacher domain.Teacher) error {
	updates := make(map[string]interface{})
	if teacher.DepartamentID != 0 {
		updates["departament_id"] = teacher.DepartamentID
	}
	if teacher.LastName != "" {
		updates["last_name"] = teacher.LastName
	}
	if teacher.FirstName != "" {
		updates["first_name"] = teacher.FirstName
	}
	if teacher.MiddleName != "" {
		updates["middle_name"] = teacher.MiddleName
	}
	if teacher.TeacherEmail != "" {
		updates["teacher_email"] = teacher.TeacherEmail
	}
	if len(updates) == 0 {
		return ErrNoUpdates
	}
	return s.TeacherRepo.Patch(ctx, teacher.TeacherID, updates)
}

func (s *TeacherService) Delete(ctx context.Context, teacherID int64) error {
	return s.TeacherRepo.Delete(ctx, teacherID)
}

func (s *TeacherService) GetByID(ctx context.Context, teacherID int64) (domain.TeacherInfo, error) {
	return s.TeacherRepo.GetByID(ctx, teacherID)
}

func (s *TeacherService) GetByEmail(ctx context.Context, teacherEmail string) (domain.TeacherInfo, error) {
	return s.TeacherRepo.GetByEmail(ctx, teacherEmail)
}

func (s *TeacherService) GetAll(ctx context.Context) ([]domain.TeacherInfo, error) {
	return s.TeacherRepo.GetAll(ctx)
}

func (s *TeacherService) GetAllByDepartamentID(ctx context.Context, departamentID int64) ([]domain.TeacherInfo, error) {
	return s.TeacherRepo.GetAllByDepartamentID(ctx, departamentID)
}
