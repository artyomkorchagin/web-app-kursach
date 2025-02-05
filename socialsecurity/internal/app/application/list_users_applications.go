package application

import (
	"context"
	"socialsecurity/internal/types"
)

func (s Service) ListUsersApplications(ctx context.Context, user_id string) ([]*types.Application, error) {

	apps, err := s.repo.ListUsersApplications(ctx, user_id)
	if err != nil {
		return []*types.Application{}, err
	}
	return apps, nil
}
