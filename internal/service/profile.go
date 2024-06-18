package service

import (
	"context"

	"github.com/BeRebornBng/OsauAmsApi/domain"
	"github.com/BeRebornBng/OsauAmsApi/internal/repository"
)

type ProfileService struct {
	ProfileRepo repository.IProfile
}

func NewProfileService(profileRepo repository.IProfile) *ProfileService {
	return &ProfileService{ProfileRepo: profileRepo}
}

func (s *ProfileService) Create(ctx context.Context, profile domain.Profile) error {
	return s.ProfileRepo.Create(ctx, profile)
}

func (s *ProfileService) Put(ctx context.Context, profile domain.Profile) error {
	return s.ProfileRepo.Put(ctx, profile)
}

func (s *ProfileService) Patch(ctx context.Context, profile domain.Profile) error {
	updates := make(map[string]interface{})
	if profile.SpecialtyCode != "" {
		updates["specialty_code"] = profile.SpecialtyCode
	}
	if profile.EducationTypeID != 0 {
		updates["education_type_id"] = profile.EducationTypeID
	}
	if profile.ProfileName != "" {
		updates["profile_name"] = profile.ProfileName
	}
	if len(updates) == 0 {
		return ErrNoUpdates
	}
	return s.ProfileRepo.Patch(ctx, profile.ProfileID, updates)
}

func (s *ProfileService) Delete(ctx context.Context, profileID int64) error {
	return s.ProfileRepo.Delete(ctx, profileID)
}

func (s *ProfileService) GetByID(ctx context.Context, profileID int64) (domain.ProfileInfo, error) {
	return s.ProfileRepo.GetByID(ctx, profileID)
}

func (s *ProfileService) GetByName(ctx context.Context, profileName string) (domain.ProfileInfo, error) {
	return s.ProfileRepo.GetByName(ctx, profileName)
}

func (s *ProfileService) GetAll(ctx context.Context) ([]domain.ProfileInfo, error) {
	return s.ProfileRepo.GetAll(ctx)
}

func (s *ProfileService) GetAllBySpecialtyCode(ctx context.Context, specialtyCode string) ([]domain.ProfileInfo, error) {
	return s.ProfileRepo.GetAllBySpecialtyCode(ctx, specialtyCode)
}

func (s *ProfileService) GetByEducationTypeID(ctx context.Context, educationTypeID int64) ([]domain.ProfileInfo, error) {
	return s.ProfileRepo.GetByEducationTypeID(ctx, educationTypeID)
}
