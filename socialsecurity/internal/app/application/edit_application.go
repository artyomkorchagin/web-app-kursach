package application

import (
	"context"
	"time"

	"socialsecurity/internal/types"
)

func (s *Service) EditApplication(ctx context.Context, req types.EditApplicatonRequest, a *types.Application) error {
	a.Type = req.Type
	a.Description = req.Description
	a.UpdatedAt = time.Now()
	if err := s.repo.EditApplication(ctx, a); err != nil {
		return err
	}
	return nil
}
