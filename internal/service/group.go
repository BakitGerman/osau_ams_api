package service

import (
	"context"

	"github.com/BeRebornBng/OsauAmsApi/domain"
	"github.com/BeRebornBng/OsauAmsApi/internal/repository"
)

type GroupService struct {
	GroupRepo repository.IGroup
}

func NewGroupService(groupRepo repository.IGroup) *GroupService {
	return &GroupService{GroupRepo: groupRepo}
}

func (s *GroupService) Create(ctx context.Context, group domain.Group) error {
	return s.GroupRepo.Create(ctx, group)
}

func (s *GroupService) Put(ctx context.Context, group domain.Group) error {
	return s.GroupRepo.Put(ctx, group)
}

func (s *GroupService) Patch(ctx context.Context, group domain.Group) error {
	updates := make(map[string]interface{})
	if group.ProfileID != 0 {
		updates["profile_id"] = group.ProfileID
	}
	if len(updates) == 0 {
		return ErrNoUpdates
	}
	return s.GroupRepo.Patch(ctx, group.GroupID, updates)
}

func (s *GroupService) Delete(ctx context.Context, groupID string) error {
	return s.GroupRepo.Delete(ctx, groupID)
}

func (s *GroupService) GetByID(ctx context.Context, groupID string) (domain.GroupInfo, error) {
	return s.GroupRepo.GetByID(ctx, groupID)
}

func (s *GroupService) GetByName(ctx context.Context, profileName string) (domain.GroupInfo, error) {
	return s.GroupRepo.GetByName(ctx, profileName)
}

func (s *GroupService) GetAll(ctx context.Context) ([]domain.GroupInfo, error) {
	return s.GroupRepo.GetAll(ctx)
}

func (s *GroupService) GetAllByProfileID(ctx context.Context, profileID int64) ([]domain.GroupInfo, error) {
	return s.GroupRepo.GetAllByProfileID(ctx, profileID)
}
