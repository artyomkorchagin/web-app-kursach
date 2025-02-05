package mssqlApplication

import (
	"context"
	"socialsecurity/internal/types"
)

func (r *ApplicationRepository) GetApplicationsByUserID(ctx context.Context, userID string) ([]*types.Application, error) {
	rows, err := r.db.QueryContext(ctx, `
   	SELECT id, user_id, description, status, type, created_at, updated_at
   	FROM applications
   	WHERE user_id = ?
   `, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var applications []*types.Application
	for rows.Next() {
		var application types.Application
		err := rows.Scan(&application.ID, &application.UserID, &application.Description, &application.Status, &application.Type, &application.CreatedAt, &application.UpdatedAt)
		if err != nil {
			return nil, err
		}
		applications = append(applications, &application)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return applications, nil
}
