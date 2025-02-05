package user

import (
	"context"
	"socialsecurity/internal/types"
)

func (s *Service) AddUser(ctx context.Context, req types.CreateUserRequest) error {
	hashpass, err := HashPassword(req.Password)
	if err != nil {
		return err
	}
	user := types.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: hashpass,
	}

	if err := s.repo.AddUser(ctx, &user); err != nil {
		return err
	}

	return nil
}
