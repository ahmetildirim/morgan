package post

import "context"

type Service struct {
	repo repository
}

func NewService(repo repository) *Service {
	return &Service{repo: repo}
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
