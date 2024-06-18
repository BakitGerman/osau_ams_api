package service

import (
	"context"

	"github.com/BeRebornBng/OsauAmsApi/domain"
	"github.com/BeRebornBng/OsauAmsApi/internal/repository"
)

type DisciplineService struct {
	DisciplineRepo repository.IDiscipline
}

func NewDisciplineService(disciplineRepo repository.IDiscipline) *DisciplineService {
	return &DisciplineService{DisciplineRepo: disciplineRepo}
}

func (s *DisciplineService) Create(ctx context.Context, discipline domain.Discipline) error {
	return s.DisciplineRepo.Create(ctx, discipline)
}

func (s *DisciplineService) Put(ctx context.Context, discipline domain.Discipline) error {
	return s.DisciplineRepo.Put(ctx, discipline)
}

func (s *DisciplineService) Patch(ctx context.Context, discipline domain.Discipline) error {
	updates := make(map[string]interface{})
	if discipline.DepartamentID != 0 {
		updates["departament_id"] = discipline.DepartamentID
	}
	if discipline.DisciplineName != "" {
		updates["discipline_name"] = discipline.DisciplineName
	}
	if len(updates) == 0 {
		return ErrNoUpdates
	}
	return s.DisciplineRepo.Patch(ctx, discipline.DisciplineID, updates)
}

func (s *DisciplineService) Delete(ctx context.Context, disciplineID int64) error {
	return s.DisciplineRepo.Delete(ctx, disciplineID)
}

func (s *DisciplineService) GetByID(ctx context.Context, disciplineID int64) (domain.DisciplineInfo, error) {
	return s.DisciplineRepo.GetByID(ctx, disciplineID)
}

func (s *DisciplineService) GetByName(ctx context.Context, disciplineName string) (domain.DisciplineInfo, error) {
	return s.DisciplineRepo.GetByName(ctx, disciplineName)
}

func (s *DisciplineService) GetAll(ctx context.Context) ([]domain.DisciplineInfo, error) {
	return s.DisciplineRepo.GetAll(ctx)
}

func (s *DisciplineService) GetAllByDepartamentID(ctx context.Context, departamentID int64) ([]domain.DisciplineInfo, error) {
	return s.DisciplineRepo.GetAllByDepartamentID(ctx, departamentID)
}
