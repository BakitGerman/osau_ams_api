package service

import (
	"context"

	"github.com/BeRebornBng/OsauAmsApi/domain"
	"github.com/BeRebornBng/OsauAmsApi/internal/repository"
)

type DepartamentService struct {
	DepartamentRepo repository.IDepartament
}

func NewDepartamentService(departamentRepo repository.IDepartament) *DepartamentService {
	return &DepartamentService{DepartamentRepo: departamentRepo}
}

func (s *DepartamentService) Create(ctx context.Context, departament domain.Departament) error {
	return s.DepartamentRepo.Create(ctx, departament)
}

func (s *DepartamentService) Put(ctx context.Context, departament domain.Departament) error {
	return s.DepartamentRepo.Put(ctx, departament)
}

func (s *DepartamentService) Patch(ctx context.Context, departament domain.Departament) error {
	updates := make(map[string]interface{})
	if departament.FacultyID != 0 {
		updates["faculty_id"] = departament.FacultyID
	}
	if departament.DepartamentName != "" {
		updates["departament_name"] = departament.DepartamentName
	}
	if departament.HeadLastName != "" {
		updates["head_last_name"] = departament.HeadLastName
	}
	if departament.HeadFirstName != "" {
		updates["head_first_name"] = departament.HeadFirstName
	}
	if departament.HeadMiddleName != "" {
		updates["head_middle_name"] = departament.HeadMiddleName
	}
	if departament.DepartamentEmail != "" {
		updates["departament_email"] = departament.DepartamentEmail
	}
	if len(updates) == 0 {
		return ErrNoUpdates
	}
	return s.DepartamentRepo.Patch(ctx, departament.DepartamentID, updates)
}

func (s *DepartamentService) Delete(ctx context.Context, departamentID int64) error {
	return s.DepartamentRepo.Delete(ctx, departamentID)
}

func (s *DepartamentService) GetByID(ctx context.Context, departamentID int64) (domain.DepartamentInfo, error) {
	return s.DepartamentRepo.GetByID(ctx, departamentID)
}

func (s *DepartamentService) GetByName(ctx context.Context, departamentName string) (domain.DepartamentInfo, error) {
	return s.DepartamentRepo.GetByName(ctx, departamentName)
}

func (s *DepartamentService) GetAll(ctx context.Context) ([]domain.DepartamentInfo, error) {
	return s.DepartamentRepo.GetAll(ctx)
}

func (s *DepartamentService) GetAllByFacultyID(ctx context.Context, facultyID int64) ([]domain.DepartamentInfo, error) {
	return s.DepartamentRepo.GetAllByFacultyID(ctx, facultyID)
}
