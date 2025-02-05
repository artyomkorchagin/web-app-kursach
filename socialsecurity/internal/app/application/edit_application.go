package application

import (
	"context"
	"time"

	"socialsecurity/internal/types"
)

func (s *Service) EditApplicaton(ctx context.Context, req types.EditApplicatonRequest, a *types.Application) error {
	a.Type = req.Type
	a.Description = req.Description
	a.UpdatedAt = time.Now()
	if err := s.repo.EditApplicaton(ctx, a); err != nil {
		return err
	}
	return nil
}
