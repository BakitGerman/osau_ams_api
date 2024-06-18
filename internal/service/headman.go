package service

import (
	"context"

	"github.com/BeRebornBng/OsauAmsApi/domain"
	"github.com/BeRebornBng/OsauAmsApi/internal/repository"
)

type HeadmanService struct {
	HeadmanRepo repository.IHeadman
}

func NewHeadmanService(headmanRepo repository.IHeadman) *HeadmanService {
	return &HeadmanService{HeadmanRepo: headmanRepo}
}

func (s *HeadmanService) Create(ctx context.Context, headman domain.Headman) error {
	return s.HeadmanRepo.Create(ctx, headman)
}

func (s *HeadmanService) Put(ctx context.Context, headman domain.Headman) error {
	return s.HeadmanRepo.Put(ctx, headman)
}

func (s *HeadmanService) Patch(ctx context.Context, headman domain.Headman) error {
	updates := make(map[string]interface{})
	if headman.StudentID != 0 {
		updates["student_id"] = headman.StudentID
	}
	if headman.GroupID != "" {
		updates["group_id"] = headman.GroupID
	}
	if len(updates) == 0 {
		return ErrNoUpdates
	}
	return s.HeadmanRepo.Patch(ctx, headman.HeadmanID, updates)
}

func (s *HeadmanService) Delete(ctx context.Context, headmanID int64) error {
	return s.HeadmanRepo.Delete(ctx, headmanID)
}

func (s *HeadmanService) GetByID(ctx context.Context, headmanID int64) (domain.HeadmanInfo, error) {
	return s.HeadmanRepo.GetByID(ctx, headmanID)
}

func (s *HeadmanService) GetByStudentID(ctx context.Context, studentID int64) (domain.HeadmanInfo, error) {
	return s.HeadmanRepo.GetByStudentID(ctx, studentID)
}

func (s *HeadmanService) GetAll(ctx context.Context) ([]domain.HeadmanInfo, error) {
	return s.HeadmanRepo.GetAll(ctx)
}
