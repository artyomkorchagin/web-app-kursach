package application

import (
	"context"
	"socialsecurity/internal/types"
	"time"

	"github.com/google/uuid"
)

func (s *Service) CreateApplication(ctx context.Context, req types.CreateApplicationRequest) (types.Application, error) {
	application := types.Application{
		ID:        uuid.New().String(),
		UserID:    req.UserID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := s.repo.AddApplication(ctx, &application); err != nil {
		return types.Application{}, err
	}

	return application, nil
}
