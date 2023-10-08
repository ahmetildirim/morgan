package post

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"morgan.io/internal/like"
)

var (
	ErrPostNotFound          = errors.New("post not found")
	ErrPostLikeAlreadyExists = like.ErrLikeAlreadyExists
)

type CreatePostServiceParams struct {
	OwnerID uuid.UUID
	Content string
}

type repository interface {
	CreatePost(ctx context.Context, post *Post) error
	GetPostsByUserIDs(ctx context.Context, userIDs []uuid.UUID) ([]*Post, error)
	AddLike(ctx context.Context, postID uuid.UUID) error
	Exists(ctx context.Context, postID uuid.UUID) (bool, error)
}

type likeService interface {
	Create(ctx context.Context, postID, ownerID uuid.UUID) error
}

type Service struct {
	repo        repository
	likeService likeService
}

func NewService(repo repository, likeService likeService) *Service {
	return &Service{
		repo:        repo,
		likeService: likeService,
	}
}

func (s *Service) CreatePost(ctx context.Context, params *CreatePostServiceParams) (*Post, error) {
	post, err := NewPost(params.OwnerID, params.Content)
	if err != nil {
		return nil, err
	}

	err = s.repo.CreatePost(ctx, post)
	if err != nil {
		return nil, err
	}

	return post, nil
}

func (s *Service) GetPostsByUserIDs(ctx context.Context, userIDs []uuid.UUID) ([]*Post, error) {
	posts, err := s.repo.GetPostsByUserIDs(ctx, userIDs)
	if err != nil {
		return nil, err
	}

	return posts, nil
}

func (s *Service) AddLike(ctx context.Context, postID, userID uuid.UUID) error {
	err := s.likeService.Create(ctx, postID, userID)
	if err != nil {
		return err
	}

	exists, err := s.repo.Exists(ctx, postID)
	if err != nil {
		return err
	}

	if !exists {
		return ErrPostNotFound
	}

	err = s.repo.AddLike(ctx, postID)
	if err != nil {
		return err
	}

	return nil
}
