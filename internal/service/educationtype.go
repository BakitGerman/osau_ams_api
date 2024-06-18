package service

import (
	"context"

	"github.com/BeRebornBng/OsauAmsApi/domain"
	"github.com/BeRebornBng/OsauAmsApi/internal/repository"
)

type EducationTypeService struct {
	EducationTypeRepo repository.IEducationType
}

func NewEducationTypeService(educationTypeRepo repository.IEducationType) *EducationTypeService {
	return &EducationTypeService{EducationTypeRepo: educationTypeRepo}
}

func (s *EducationTypeService) Create(ctx context.Context, educationType domain.EducationType) error {
	return s.EducationTypeRepo.Create(ctx, educationType)
}

func (s *EducationTypeService) Put(ctx context.Context, educationType domain.EducationType) error {
	return s.EducationTypeRepo.Put(ctx, educationType)
}

func (s *EducationTypeService) Patch(ctx context.Context, educationType domain.EducationType) error {
	updates := make(map[string]interface{})
	if educationType.EducationTypeName != "" {
		updates["education_type_name"] = educationType.EducationTypeName
	}
	if len(updates) == 0 {
		return ErrNoUpdates
	}
	return s.EducationTypeRepo.Patch(ctx, educationType.EducationTypeID, updates)
}

func (s *EducationTypeService) Delete(ctx context.Context, educationTypeID int64) error {
	return s.EducationTypeRepo.Delete(ctx, educationTypeID)
}

func (s *EducationTypeService) GetByID(ctx context.Context, educationTypeID int64) (domain.EducationType, error) {
	return s.EducationTypeRepo.GetByID(ctx, educationTypeID)
}

func (s *EducationTypeService) GetByName(ctx context.Context, educationTypeName string) (domain.EducationType, error) {
	return s.EducationTypeRepo.GetByName(ctx, educationTypeName)
}

func (s *EducationTypeService) GetAll(ctx context.Context) ([]domain.EducationType, error) {
	return s.EducationTypeRepo.GetAll(ctx)
}
