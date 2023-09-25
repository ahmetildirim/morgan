package auth

import "context"

type Service struct {
	userService userService
	secretKey   string
}

func NewService(userService userService, secretKey string) *Service {
	return &Service{
		userService: userService,
		secretKey:   secretKey,
	}
}

func (s Service) Login(ctx context.Context, params *LoginServiceParams) (string, error) {
	user, err := s.userService.Authenticate(ctx, params.Email, params.Password)
	if err != nil {
		return "", err
	}

	token, err := NewAccessToken(user, s.secretKey)
	if err != nil {
		return "", err
	}

	return token, nil
}
