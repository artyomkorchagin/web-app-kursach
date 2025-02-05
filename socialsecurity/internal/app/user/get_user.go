package user

import (
	"context"
	"socialsecurity/internal/types"
)

func (s *Service) GetUser(ctx context.Context, username, password string) (*types.User, error) {

	user, err := s.repo.GetUser(ctx, username, password)
	if err != nil {
		return nil, err
	}

	return user, nil
}
