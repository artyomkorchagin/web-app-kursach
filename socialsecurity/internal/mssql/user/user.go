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
		INSERT INTO users (id, name, email, created_at)
		VALUES (@p1, @p2, @p3, @p4)`

	_, err := r.db.ExecContext(ctx, query,
		user.ID,
		user.Name,
		user.Email,
		user.CreatedAt,
	)
	return err
}
