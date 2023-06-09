package useCase

import (
	"Ass4ADVP/services/contact/internal/domain"
	"context"
	"time"
)

type groupUsecase struct {
	groupRepo      domain.GroupRepository
	contextTimeout time.Duration
}

// Create implements domain.GroupUseCase
func (uc *groupUsecase) Create(group *domain.Group) error {
	ctx, cancel := context.WithTimeout(context.Background(), uc.contextTimeout)
	defer cancel()

	return uc.groupRepo.Update(group, ctx)
}

// GetByID implements domain.GroupUseCase
func (uc *groupUsecase) GetByID(id int64) (*domain.Group, error) {
	ctx, cancel := context.WithTimeout(context.Background(), uc.contextTimeout)
	defer cancel()

	return uc.groupRepo.GetByID(id, ctx)
}

// Update implements domain.GroupUseCase
func (uc *groupUsecase) Update(group *domain.Group) error {
	ctx, cancel := context.WithTimeout(context.Background(), uc.contextTimeout)
	defer cancel()

	return uc.groupRepo.Update(group, ctx)
}

func NewGroupUsecase(c domain.GroupRepository, timeout time.Duration) domain.GroupUseCase {
	return &groupUsecase{
		groupRepo:      c,
		contextTimeout: timeout,
	}
}
