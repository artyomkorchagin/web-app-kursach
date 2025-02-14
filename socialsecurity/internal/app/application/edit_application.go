package application

import (
	"context"

	"socialsecurity/internal/types"
)

func (s *Service) EditApplication(ctx context.Context, req types.EditApplicatonRequest, a *types.Application) error {
	a.Description = req.Description
	if err := s.repo.EditApplication(ctx, a); err != nil {
		return err
	}
	return nil
}
