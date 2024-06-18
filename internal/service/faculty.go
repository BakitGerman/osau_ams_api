package service

import (
	"context"

	"github.com/BeRebornBng/OsauAmsApi/domain"
	"github.com/BeRebornBng/OsauAmsApi/internal/repository"
)

type FacultyService struct {
	FacultyRepo repository.IFaculty
}

func NewFacultyService(facultyRepo repository.IFaculty) *FacultyService {
	return &FacultyService{FacultyRepo: facultyRepo}
}

func (s *FacultyService) Create(ctx context.Context, faculty domain.Faculty) error {
	return s.FacultyRepo.Create(ctx, faculty)
}

func (s *FacultyService) Put(ctx context.Context, faculty domain.Faculty) error {
	return s.FacultyRepo.Put(ctx, faculty)
}

func (s *FacultyService) Patch(ctx context.Context, faculty domain.Faculty) error {
	updates := make(map[string]interface{})
	if faculty.UniversityID != 0 {
		updates["university_id"] = faculty.UniversityID
	}
	if faculty.FacultyName != "" {
		updates["faculty_name"] = faculty.FacultyName
	}
	if faculty.HeadLastName != "" {
		updates["head_last_name"] = faculty.HeadLastName
	}
	if faculty.HeadFirstName != "" {
		updates["head_first_name"] = faculty.HeadFirstName
	}
	if faculty.HeadMiddleName != "" {
		updates["head_middle_name"] = faculty.HeadMiddleName
	}
	if faculty.FacultyEmail != "" {
		updates["faculty_email"] = faculty.FacultyEmail
	}
	if len(updates) == 0 {
		return ErrNoUpdates
	}
	return s.FacultyRepo.Patch(ctx, faculty.FacultyID, updates)
}

func (s *FacultyService) Delete(ctx context.Context, facultyID int64) error {
	return s.FacultyRepo.Delete(ctx, facultyID)
}

func (s *FacultyService) GetByID(ctx context.Context, facultyID int64) (domain.FacultyInfo, error) {
	return s.FacultyRepo.GetByID(ctx, facultyID)
}

func (s *FacultyService) GetByName(ctx context.Context, facultyName string) (domain.FacultyInfo, error) {
	return s.FacultyRepo.GetByName(ctx, facultyName)
}

func (s *FacultyService) GetAll(ctx context.Context) ([]domain.FacultyInfo, error) {
	return s.FacultyRepo.GetAll(ctx)
}

func (s *FacultyService) GetAllByUniversityID(ctx context.Context, universityID int64) ([]domain.FacultyInfo, error) {
	return s.FacultyRepo.GetAllByUniversityID(ctx, universityID)
}
