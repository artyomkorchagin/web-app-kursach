package user

import (
	"context"
	"socialsecurity/internal/types"
)

type Reader interface {
	GetUser(ctx context.Context, username, password string) (*types.User, error)
}

type Writer interface {
	AddUser(ctx context.Context, a *types.User) error
}

type ReadWriter interface {
	Reader
	Writer
}
