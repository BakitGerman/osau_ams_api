package service

import (
	"context"

	"github.com/BeRebornBng/OsauAmsApi/domain"
	"github.com/BeRebornBng/OsauAmsApi/internal/repository"
)

type StudentService struct {
	StudentRepo repository.IStudent
}

func NewStudentService(studentRepo repository.IStudent) *StudentService {
	return &StudentService{StudentRepo: studentRepo}
}

func (s *StudentService) Create(ctx context.Context, student domain.Student) error {
	return s.StudentRepo.Create(ctx, student)
}

func (s *StudentService) Put(ctx context.Context, student domain.Student) error {
	return s.StudentRepo.Put(ctx, student)
}

func (s *StudentService) Patch(ctx context.Context, student domain.Student) error {
	updates := make(map[string]interface{})
	if student.LastName != "" {
		updates["last_name"] = student.LastName
	}
	if student.FirstName != "" {
		updates["first_name"] = student.FirstName
	}
	if student.MiddleName != "" {
		updates["middle_name"] = student.MiddleName
	}
	if student.GroupID != "" {
		updates["group_id"] = student.GroupID
	}
	if len(updates) == 0 {
		return ErrNoUpdates
	}
	return s.StudentRepo.Patch(ctx, student.StudentID, updates)
}

func (s *StudentService) Delete(ctx context.Context, studentID int64) error {
	return s.StudentRepo.Delete(ctx, studentID)
}

func (s *StudentService) GetByID(ctx context.Context, studentID int64) (domain.Student, error) {
	return s.StudentRepo.GetByID(ctx, studentID)
}

func (s *StudentService) GetByName(ctx context.Context, lastName, firstName, middleName string) (domain.Student, error) {
	return s.StudentRepo.GetByName(ctx, lastName, firstName, middleName)
}

func (s *StudentService) GetAll(ctx context.Context) ([]domain.Student, error) {
	return s.StudentRepo.GetAll(ctx)
}

func (s *StudentService) GetAllByGroupID(ctx context.Context, groupID string) ([]domain.Student, error) {
	return s.StudentRepo.GetAllByGroupID(ctx, groupID)
}
