package application

import (
	"context"
	"socialsecurity/internal/types"
)

type Reader interface {
	// FindApplicationByID(ctx context.Context, id string) (*types.Application, error)
	ListUsersApplications(ctx context.Context, user_id string) ([]*types.Application, error)
}

type Writer interface {
	AddApplication(ctx context.Context, a *types.Application) error
	EditApplicaton(ctx context.Context, a *types.Application) error
}

type ReadWriter interface {
	Reader
	Writer
}
