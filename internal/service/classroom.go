package service

import (
	"context"

	"github.com/BeRebornBng/OsauAmsApi/domain"
	"github.com/BeRebornBng/OsauAmsApi/internal/repository"
)

type ClassroomService struct {
	ClassroomRepo repository.IClassroom
}

func NewClassroomService(classroomRepo repository.IClassroom) *ClassroomService {
	return &ClassroomService{ClassroomRepo: classroomRepo}
}

func (s *ClassroomService) Create(ctx context.Context, classroom domain.Classroom) error {
	return s.ClassroomRepo.Create(ctx, classroom)
}

func (s *ClassroomService) Put(ctx context.Context, classroom domain.Classroom) error {
	return s.ClassroomRepo.Put(ctx, classroom)
}

func (s *ClassroomService) Patch(ctx context.Context, classroom domain.Classroom) error {
	updates := make(map[string]interface{})
	if classroom.ClassroomName != "" {
		updates["classroom_name"] = classroom.ClassroomName
	}
	if len(updates) == 0 {
		return ErrNoUpdates
	}
	return s.ClassroomRepo.Patch(ctx, classroom.ClassroomID, updates)
}

func (s *ClassroomService) Delete(ctx context.Context, classroomID int64) error {
	return s.ClassroomRepo.Delete(ctx, classroomID)
}

func (s *ClassroomService) GetByID(ctx context.Context, classroomID int64) (domain.Classroom, error) {
	return s.ClassroomRepo.GetByID(ctx, classroomID)
}

func (s *ClassroomService) GetAll(ctx context.Context) ([]domain.Classroom, error) {
	return s.ClassroomRepo.GetAll(ctx)
}
