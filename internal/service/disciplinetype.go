package service

import (
	"context"

	"github.com/BeRebornBng/OsauAmsApi/domain"
	"github.com/BeRebornBng/OsauAmsApi/internal/repository"
)

type DisciplineTypeService struct {
	DisciplineTypeRepo repository.IDisciplineType
}

func NewDisciplineTypeService(disciplineTypeRepo repository.IDisciplineType) *DisciplineTypeService {
	return &DisciplineTypeService{DisciplineTypeRepo: disciplineTypeRepo}
}

func (s *DisciplineTypeService) Create(ctx context.Context, disciplineType domain.DisciplineType) error {
	return s.DisciplineTypeRepo.Create(ctx, disciplineType)
}

func (s *DisciplineTypeService) Put(ctx context.Context, disciplineType domain.DisciplineType) error {
	return s.DisciplineTypeRepo.Put(ctx, disciplineType)
}

func (s *DisciplineTypeService) Patch(ctx context.Context, disciplineType domain.DisciplineType) error {
	updates := make(map[string]interface{})
	if disciplineType.DisciplineTypeName != "" {
		updates["discipline_type_name"] = disciplineType.DisciplineTypeName
	}
	if len(updates) == 0 {
		return ErrNoUpdates
	}
	return s.DisciplineTypeRepo.Patch(ctx, disciplineType.DisciplineTypeID, updates)
}

func (s *DisciplineTypeService) Delete(ctx context.Context, disciplineTypeID int64) error {
	return s.DisciplineTypeRepo.Delete(ctx, disciplineTypeID)
}

func (s *DisciplineTypeService) GetByID(ctx context.Context, disciplineTypeID int64) (domain.DisciplineType, error) {
	return s.DisciplineTypeRepo.GetByID(ctx, disciplineTypeID)
}

func (s *DisciplineTypeService) GetAll(ctx context.Context) ([]domain.DisciplineType, error) {
	return s.DisciplineTypeRepo.GetAll(ctx)
}
