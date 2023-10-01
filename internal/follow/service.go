package follow

import (
	"context"
	"errors"
)

var (
	ErrFolloweeNotFound = errors.New("followee not found")
	ErrAlreadyFollowing = errors.New("already following")
)

type Service struct {
	repo    repository
	userSvc userService
}

func NewService(repo repository, userSvc userService) *Service {
	return &Service{
		repo:    repo,
		userSvc: userSvc,
	}
}

func (s *Service) Follow(ctx context.Context, params *CreateFollowServiceParams) error {
	_, err := s.userSvc.GetUser(ctx, params.FolloweeID)
	if err != nil {
		return ErrFolloweeNotFound
	}

	follow, err := s.repo.FindByFollowerAndFollowee(ctx, params.FollowerID, params.FolloweeID)
	if err != nil && err != ErrNotFound {
		return err
	}
	if follow != nil {
		return ErrAlreadyFollowing
	}

	follow, err = NewFollow(params.FollowerID, params.FolloweeID)
	if err != nil {
		return err
	}

	err = s.repo.Create(ctx, follow)
	if err != nil {
		return err
	}

	return nil
}
