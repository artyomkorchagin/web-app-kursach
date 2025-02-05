package mssqlUser

import (
	"context"
	"database/sql"

	"socialsecurity/internal/types"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) GetUser(ctx context.Context, username, password string) (*types.User, error) {
	return nil, nil
}

func (r *UserRepository) AddUser(ctx context.Context, user *types.User) error {
	query := `
		INSERT INTO Users (name, email, password)
		VALUES (?, ?, ?)`

	_, err := r.db.ExecContext(ctx, query,
		user.Name,
		user.Email,
		user.Password,
	)
	return err
}
