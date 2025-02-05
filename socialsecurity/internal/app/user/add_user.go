package user

import (
	"context"
	"socialsecurity/internal/types"
	"time"

	"github.com/google/uuid"
)

func (s *Service) AddUser(ctx context.Context, req types.CreateUserRequest) (types.User, error) {
	hashpass, err := HashPassword(req.Password)
	if err != nil {
		return types.User{}, err
	}
	user := types.User{
		ID:        uuid.New().String(),
		Name:      req.Name,
		Email:     req.Email,
		Password:  hashpass,
		CreatedAt: time.Now().String(),
	}

	if err := s.repo.AddUser(ctx, &user); err != nil {
		return types.User{}, err
	}

	return user, nil
}
