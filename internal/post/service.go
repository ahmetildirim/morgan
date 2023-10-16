package post

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"morgan.io/internal/post/comment"
	"morgan.io/internal/post/like"
)

var (
	ErrPostNotFound      = errors.New("post not found")
	ErrLikeAlreadyExists = errors.New("like already exists")
)

type CreatePostServiceParams struct {
	OwnerID uuid.UUID
	Content string
}

type CreateCommentServiceParams struct {
	PostID  uuid.UUID
	OwnerID uuid.UUID
	Content string
}

type repository interface {
	CreatePost(ctx context.Context, post *Post) error
	GetPostsByUserIDs(ctx context.Context, userIDs []uuid.UUID) ([]*Post, error)
	AddLike(ctx context.Context, postID uuid.UUID) error
	Exists(ctx context.Context, postID uuid.UUID) (bool, error)
}

type commentRepository interface {
	Create(ctx context.Context, comment *comment.Comment) error
}

type likeRepository interface {
	Create(ctx context.Context, like *like.Like) error
	Exists(ctx context.Context, postID, ownerID uuid.UUID) (bool, error)
}

type Service struct {
	repo        repository
	commentRepo commentRepository
	likeRepo    likeRepository
}

func NewService(repo repository, commentRepo commentRepository, likeRepo likeRepository) *Service {
	return &Service{
		repo:        repo,
		commentRepo: commentRepo,
		likeRepo:    likeRepo,
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

func (s *Service) CreateComment(ctx context.Context, params *CreateCommentServiceParams) (*comment.Comment, error) {
	exists, err := s.repo.Exists(ctx, params.PostID)
	if err != nil {
		return nil, err
	}

	if !exists {
		return nil, ErrPostNotFound
	}

	comment, err := comment.NewComment(params.PostID, params.OwnerID, params.Content)
	if err != nil {
		return nil, err
	}

	err = s.commentRepo.Create(ctx, comment)
	if err != nil {
		return nil, err
	}

	return comment, nil
}

func (s *Service) AddLike(ctx context.Context, postID, userID uuid.UUID) error {
	exists, err := s.repo.Exists(ctx, postID)
	if err != nil {
		return err
	}

	if !exists {
		return ErrPostNotFound
	}

	exists, err = s.likeRepo.Exists(ctx, postID, userID)
	if err != nil {
		return err
	}

	if exists {
		return ErrLikeAlreadyExists
	}

	like, err := like.NewLike(postID, userID)
	if err != nil {
		return err
	}

	// create a like record
	err = s.likeRepo.Create(ctx, like)
	if err != nil {
		return err
	}

	// increment the post's like count
	err = s.repo.AddLike(ctx, postID)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) GetPostsByUserIDs(ctx context.Context, userIDs []uuid.UUID) ([]*Post, error) {
	posts, err := s.repo.GetPostsByUserIDs(ctx, userIDs)
	if err != nil {
		return nil, err
	}

	return posts, nil
}
