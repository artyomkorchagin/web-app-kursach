package mssqlApplication

import (
	"context"
	"socialsecurity/internal/types"
)

func (r *ApplicationRepository) ListUsersApplications(ctx context.Context, user_id string) ([]*types.Application, error) {
	return nil, nil
}
