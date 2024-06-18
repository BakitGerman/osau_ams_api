package service

import (
	"context"

	"github.com/BeRebornBng/OsauAmsApi/domain"
	"github.com/BeRebornBng/OsauAmsApi/internal/repository"
)

type EducationLevelService struct {
	EducationLevelRepo repository.IEducationLevel
}

func NewEducationLevelService(educationLevelRepo repository.IEducationLevel) *EducationLevelService {
	return &EducationLevelService{EducationLevelRepo: educationLevelRepo}
}

func (s *EducationLevelService) Create(ctx context.Context, educationLevel domain.EducationLevel) error {
	return s.EducationLevelRepo.Create(ctx, educationLevel)
}

func (s *EducationLevelService) Put(ctx context.Context, educationLevel domain.EducationLevel) error {
	return s.EducationLevelRepo.Put(ctx, educationLevel)
}

func (s *EducationLevelService) Patch(ctx context.Context, educationLevel domain.EducationLevel) error {
	updates := make(map[string]interface{})
	if educationLevel.EducationLevelName != "" {
		updates["education_level_name"] = educationLevel.EducationLevelName
	}
	if len(updates) == 0 {
		return ErrNoUpdates
	}
	return s.EducationLevelRepo.Patch(ctx, educationLevel.EducationLevelID, updates)
}

func (s *EducationLevelService) Delete(ctx context.Context, educationLevelID int64) error {
	return s.EducationLevelRepo.Delete(ctx, educationLevelID)
}

func (s *EducationLevelService) GetByID(ctx context.Context, educationLevelID int64) (domain.EducationLevel, error) {
	return s.EducationLevelRepo.GetByID(ctx, educationLevelID)
}

func (s *EducationLevelService) GetAll(ctx context.Context) ([]domain.EducationLevel, error) {
	return s.EducationLevelRepo.GetAll(ctx)
}
