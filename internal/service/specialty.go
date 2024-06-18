package service

import (
	"context"

	"github.com/BeRebornBng/OsauAmsApi/domain"
	"github.com/BeRebornBng/OsauAmsApi/internal/repository"
)

type SpecialtyService struct {
	SpecialtyRepo repository.ISpecialty
}

func NewSpecialtyService(specialtyRepo repository.ISpecialty) *SpecialtyService {
	return &SpecialtyService{SpecialtyRepo: specialtyRepo}
}

func (s *SpecialtyService) Create(ctx context.Context, specialty domain.Specialty) error {
	return s.SpecialtyRepo.Create(ctx, specialty)
}

func (s *SpecialtyService) Put(ctx context.Context, specialty domain.Specialty) error {
	return s.SpecialtyRepo.Put(ctx, specialty)
}

func (s *SpecialtyService) Patch(ctx context.Context, specialty domain.Specialty) error {
	updates := make(map[string]interface{})
	if specialty.SpecialtyName != "" {
		updates["specialty_name"] = specialty.SpecialtyName
	}
	if specialty.DepartamentID != 0 {
		updates["departament_id"] = specialty.DepartamentID
	}
	if specialty.EducationLevelID != 0 {
		updates["education_level_id"] = specialty.EducationLevelID
	}
	if len(updates) == 0 {
		return ErrNoUpdates
	}
	return s.SpecialtyRepo.Patch(ctx, specialty.SpecialtyCode, updates)
}

func (s *SpecialtyService) Delete(ctx context.Context, specialtyCode string) error {
	return s.SpecialtyRepo.Delete(ctx, specialtyCode)
}

func (s *SpecialtyService) GetByCode(ctx context.Context, specialtyCode string) (domain.SpecialtyInfo, error) {
	return s.SpecialtyRepo.GetByCode(ctx, specialtyCode)
}

func (s *SpecialtyService) GetByName(ctx context.Context, specialtyName string) (domain.SpecialtyInfo, error) {
	return s.SpecialtyRepo.GetByName(ctx, specialtyName)
}

func (s *SpecialtyService) GetAll(ctx context.Context) ([]domain.SpecialtyInfo, error) {
	return s.SpecialtyRepo.GetAll(ctx)
}

func (s *SpecialtyService) GetAllByDepartamentID(ctx context.Context, departamentID int64) ([]domain.SpecialtyInfo, error) {
	return s.SpecialtyRepo.GetAllByDepartamentID(ctx, departamentID)
}
