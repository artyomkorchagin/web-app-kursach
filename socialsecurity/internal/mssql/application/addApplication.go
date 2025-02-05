package mssqlApplication

import (
	"context"
	"socialsecurity/internal/types"
)

func (r *ApplicationRepository) AddApplication(ctx context.Context, application *types.Application) error {
	_, err := r.db.ExecContext(ctx, `
		INSERT INTO applications (id, user_id, description, status, type, created_at, updated_at)
		VALUES (?, ?, ?, ?, ?, ?, ?)
	`, application.ID, application.UserID, application.Description, types.StatusPending, application.Type, application.CreatedAt, application.UpdatedAt)
	if err != nil {
		return err
	}

	return nil
}
