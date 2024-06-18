package service

import (
	"context"

	"github.com/BeRebornBng/OsauAmsApi/domain"
	"github.com/BeRebornBng/OsauAmsApi/internal/repository"
)

type UniversityService struct {
	UniversityRepo repository.IUniversity
}

func NewUniversityService(universityRepo repository.IUniversity) *UniversityService {
	return &UniversityService{UniversityRepo: universityRepo}
}

func (s *UniversityService) Create(ctx context.Context, university domain.University) error {
	return s.UniversityRepo.Create(ctx, university)
}

func (s *UniversityService) Put(ctx context.Context, university domain.University) error {
	return s.UniversityRepo.Put(ctx, university)
}

func (s *UniversityService) Patch(ctx context.Context, university domain.University) error {
	updates := make(map[string]interface{})
	if university.UniversityName != "" {
		updates["university_name"] = university.UniversityName
	}
	if university.HeadLastName != "" {
		updates["head_last_name"] = university.HeadLastName
	}
	if university.HeadFirstName != "" {
		updates["head_first_name"] = university.HeadFirstName
	}
	if university.HeadMiddleName != "" {
		updates["head_middle_name"] = university.HeadMiddleName
	}
	if university.UniversityEmail != "" {
		updates["university_email"] = university.UniversityEmail
	}
	if len(updates) == 0 {
		return ErrNoUpdates
	}
	return s.UniversityRepo.Patch(ctx, university.UniversityID, updates)
}

func (s *UniversityService) Delete(ctx context.Context, universityID int64) error {
	return s.UniversityRepo.Delete(ctx, universityID)
}

func (s *UniversityService) GetByID(ctx context.Context, universityID int64) (domain.University, error) {
	return s.UniversityRepo.GetByID(ctx, universityID)
}

func (s *UniversityService) GetByName(ctx context.Context, universityName string) (domain.University, error) {
	return s.UniversityRepo.GetByName(ctx, universityName)
}

func (s *UniversityService) GetAll(ctx context.Context) ([]domain.University, error) {
	return s.UniversityRepo.GetAll(ctx)
}
