package mssqlApplication

import (
	"context"
	"database/sql"
	"fmt"
	"socialsecurity/internal/types"
	"time"
)

func (r *ApplicationRepository) AddApplication(ctx context.Context, application *types.Application) error {
	// Define the SQL query
	query := `
        INSERT INTO applications (
            id, 
            user_id, 
            benefit_id, 
            service_id, 
            application_date, 
            status, 
            approval_date, 
            rejection_reason, 
            description
        ) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)
    `

	// Set default values for optional fields
	if application.Date == "" {
		application.Date = time.Now().Format("2006-01-02") // Default to current date
	}
	if application.Status == "" {
		application.Status = types.StatusPending // Default status
	}

	// Execute the query
	_, err := r.db.ExecContext(ctx, query,
		application.ID,
		application.UserID,
		application.BenefitID,
		application.ServiceID,
		application.Date,
		application.Status,
		sql.NullString{String: application.ApprovalDate, Valid: application.ApprovalDate != ""},       // Handle nullable ApprovalDate
		sql.NullString{String: application.RejectionReason, Valid: application.RejectionReason != ""}, // Handle nullable RejectionReason
		application.Description,
	)
	if err != nil {
		return fmt.Errorf("failed to insert application: %w", err)
	}

	return nil
}
