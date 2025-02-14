package mssqlApplication

import (
	"context"
	"database/sql"
	"fmt"
	"socialsecurity/internal/types"
)

func (r *ApplicationRepository) GetApplicationsByUserID(ctx context.Context, userID string) ([]*types.Application, error) {
	// Define the SQL query
	query := `
        SELECT 
            id, 
            user_id, 
            benefit_id, 
            service_id, 
            application_date, 
            status, 
            approval_date, 
            rejection_reason, 
            description
        FROM applications
        WHERE user_id = ?
    `

	// Execute the query
	rows, err := r.db.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}
	defer rows.Close()

	// Create a slice to store the applications
	var applications []*types.Application

	// Iterate over the rows and map them to Application structs
	for rows.Next() {
		var application types.Application
		var approvalDate sql.NullString
		var rejectionReason sql.NullString

		// Scan the row into the struct fields
		err := rows.Scan(
			&application.ID,
			&application.UserID,
			&application.BenefitID,
			&application.ServiceID,
			&application.Date,
			&application.Status,
			&approvalDate,
			&rejectionReason,
			&application.Description,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}

		// Handle nullable fields
		if approvalDate.Valid {
			application.ApprovalDate = approvalDate.String
		}
		if rejectionReason.Valid {
			application.RejectionReason = rejectionReason.String
		}

		// Append the application to the slice
		applications = append(applications, &application)
	}

	// Check for errors during iteration
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error during row iteration: %w", err)
	}

	// If no applications are found, return an empty slice
	if len(applications) == 0 {
		return []*types.Application{}, nil
	}

	return applications, nil
}
