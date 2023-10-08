package follow

import (
	"context"
	"errors"

	"morgan.io/internal/user"

	"github.com/google/uuid"
)

var (
	ErrFolloweeNotFound = errors.New("followee not found")
	ErrAlreadyFollowing = errors.New("already following")
)

type CreateFollowServiceParams struct {
	FollowerID uuid.UUID
	FolloweeID uuid.UUID
}

type repository interface {
	Create(ctx context.Context, follow *Follow) error
	FindByFollowerAndFollowee(ctx context.Context, followerID uuid.UUID, followeeID uuid.UUID) (*Follow, error)
	FindByFollower(ctx context.Context, followerID uuid.UUID) ([]*Follow, error)
}

type userService interface {
	GetUser(ctx context.Context, id uuid.UUID) (*user.User, error)
}

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

func (s *Service) GetFollowees(ctx context.Context, userID uuid.UUID) ([]uuid.UUID, error) {
	follows, err := s.repo.FindByFollower(ctx, userID)
	if err != nil {
		return nil, err
	}

	var followees []uuid.UUID
	for _, follow := range follows {
		followees = append(followees, follow.FolloweeID)
	}

	return followees, nil
}
